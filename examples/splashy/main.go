//
// Based on the example: https://github.com/stdiopt/gowasm-experiments/tree/master/splashy
//
package main

import (
	"fmt"
	"github.com/nuberu/webgl"
	"github.com/nuberu/webgl/types"
	"log"
	"math/rand"
	"strconv"

	"syscall/js"
	// this box2d throws some unexpected panics
	"github.com/ByteArena/box2d"

	"github.com/lucasb-eyer/go-colorful"
)

var (
	width      int
	height     int
	ctx        js.Value
	simSpeed   = 1.0
	worldScale = 0.0125
	resDiv     = 8
	maxBodies  = 120
)

func main() {
	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("createElement", "canvas")
	doc.Get("body").Call("appendChild", canvasEl)
	width := 800
	height := 600
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)

	gl, err := webgl.FromCanvas(canvasEl)

	done := make(chan struct{}, 0)

	thing := Thing{}
	mouseDown := false

	mouseDownEvt := js.NewCallback(func(args []js.Value) {
		mouseDown = true
		evt := args[0]
		if evt.Get("target") != canvasEl {
			return
		}
		mx := evt.Get("clientX").Float() * worldScale
		my := evt.Get("clientY").Float() * worldScale
		thing.AddCircle(mx, my)
	})
	defer mouseDownEvt.Release() // go1.11Beta1 is Close() latest is Release()

	mouseUpEvt := js.NewCallback(func(args []js.Value) {
		mouseDown = false
	})
	defer mouseUpEvt.Release()

	mouseMoveEvt := js.NewCallback(func(args []js.Value) {
		if !mouseDown {
			return
		}
		evt := args[0]
		if evt.Get("target") != canvasEl {
			return
		}
		mx := evt.Get("clientX").Float() * worldScale
		my := evt.Get("clientY").Float() * worldScale
		thing.AddCircle(mx, my)
	})
	defer mouseMoveEvt.Release()

	speedInputEvt := js.NewCallback(func(args []js.Value) {
		evt := args[0]
		fval, err := strconv.ParseFloat(evt.Get("target").Get("value").String(), 64)
		if err != nil {
			log.Println("Invalid value", err)
			return
		}
		simSpeed = fval
	})
	defer speedInputEvt.Release()
	// Events
	doc.Call("addEventListener", "mousedown", mouseDownEvt)
	doc.Call("addEventListener", "mouseup", mouseUpEvt)
	doc.Call("addEventListener", "mousemove", mouseMoveEvt)
	speedInput := doc.Call("createElement", "input")
	doc.Get("body").Call("appendChild", speedInput)
	speedInput.Call("addEventListener", "input", speedInputEvt)

	err = thing.Init(gl)
	if err != nil {
		log.Println("Err Initializing thing:", err)
		return
	}

	// Draw things
	var renderFrame js.Callback
	var tmark float64
	var markCount = 0
	var tdiffSum float64

	renderFrame = js.NewCallback(func(args []js.Value) {
		// Update the DOM less frequently TODO: func on this
		now := args[0].Float()
		tdiff := now - tmark
		tdiffSum += tdiff
		markCount++
		if markCount > 10 {
			doc.Call("getElementById", "fps").Set("innerHTML", fmt.Sprintf("FPS: %.01f", 1000/(tdiffSum/float64(markCount))))
			tdiffSum, markCount = 0, 0
		}
		tmark = now
		// --
		thing.Render(gl, tdiff/1000)

		js.Global().Call("requestAnimationFrame", renderFrame)
	})
	defer renderFrame.Release()

	// Start running
	js.Global().Call("requestAnimationFrame", renderFrame)

	<-done

}

type Thing struct {
	// dot shaders
	prog        *types.Program
	aPosition   int
	uFragColor  *types.UniformLocation
	uResolution *types.UniformLocation
	dotBuf      *types.Buffer
	qBlur       *QuadFX
	qThreshold  *QuadFX
	rtTex       [2]*types.Texture     // render target Texture
	rt          [2]*types.FrameBuffer // framebuffer(render target)
	world       box2d.B2World
}

func (t *Thing) Init(gl *webgl.RenderingContext) error {
	// Drawing program
	var err error
	t.prog, err = programFromSrc(gl, dotVertShader, dotFragShader)
	if err != nil {
		return err
	}
	t.aPosition = gl.GetAttribLocation(t.prog, "a_position")
	t.uFragColor = gl.GetUniformLocation(t.prog, "uFragColor")
	t.uResolution = gl.GetUniformLocation(t.prog, "uResolution")

	t.dotBuf = gl.CreateBuffer()
	//renderer targets
	for i := 0; i < 2; i++ {
		t.rtTex[i] = createTexture(gl, width/resDiv, height/resDiv)
		t.rt[i] = createFB(gl, t.rtTex[i])
	}

	t.qBlur = &QuadFX{}
	err = t.qBlur.Init(gl, blurShader)
	if err != nil {
		log.Fatal("Error:", err)
	}
	t.qThreshold = &QuadFX{}
	err = t.qThreshold.Init(gl, thresholdShader)
	if err != nil {
		log.Fatal("Error:", err)
	}

	//////////////////////////
	// Physics
	// ///////////
	t.world = box2d.MakeB2World(box2d.B2Vec2{X: 0, Y: 9.8})
	floor := t.world.CreateBody(&box2d.B2BodyDef{
		Type:     box2d.B2BodyType.B2_kinematicBody,
		Position: box2d.B2Vec2{X: 0, Y: float64(height+10) * worldScale},
		Active:   true,
	})
	floorShape := &box2d.B2PolygonShape{}
	floorShape.SetAsBox(float64(width)*worldScale, 20*worldScale)
	ft := floor.CreateFixture(floorShape, 1)
	ft.M_friction = 0.3

	// Walls
	wallShape := &box2d.B2PolygonShape{}
	wallShape.SetAsBox(20*worldScale, float64(height)*worldScale)

	wallL := t.world.CreateBody(&box2d.B2BodyDef{
		Type:     box2d.B2BodyType.B2_kinematicBody,
		Position: box2d.B2Vec2{X: 0, Y: 0},
		Active:   true,
	})
	wlf := wallL.CreateFixture(wallShape, 1)
	wlf.M_friction = 0.3

	wallR := t.world.CreateBody(&box2d.B2BodyDef{
		Type:     box2d.B2BodyType.B2_kinematicBody,
		Position: box2d.B2Vec2{X: float64(width) * worldScale, Y: 0},
		Active:   true,
	})
	wrt := wallR.CreateFixture(wallShape, 1)
	wrt.M_friction = 0.3

	for i := 0; i < 10; i++ {
		t.AddCircle(rand.Float64()*float64(width)*worldScale, rand.Float64()*float64(height)*worldScale)
	}

	return nil
}

func (t *Thing) Render(gl *webgl.RenderingContext, dtTime float64) {
	texWidth := width / resDiv
	texHeight := height / resDiv
	t.world.Step(dtTime*simSpeed, 3, 3)

	gl.BindFrameBuffer(webgl.FRAMEBUFFER, t.rt[0])
	gl.Viewport(0, 0, texWidth, texHeight)
	gl.ClearColor(0, 0, 0, 0)
	gl.Clear(uint32(webgl.COLOR_BUFFER_BIT))

	// DotRenderer
	gl.UseProgram(t.prog)

	count := 0
	for curBody := t.world.GetBodyList(); curBody != nil; curBody = curBody.M_next {
		ft := curBody.M_fixtureList
		if _, ok := ft.M_shape.(*box2d.B2CircleShape); !ok {
			continue
		}
		x := float32(curBody.M_xf.P.X / (float64(width) * worldScale))  /* 0-1 */
		y := float32(curBody.M_xf.P.Y / (float64(height) * worldScale)) /*0-1*/

		col := colorful.Hsv(float64(360*count/maxBodies), 1, 1)
		gl.VertexAttrib2f(t.aPosition, x, y)
		gl.Uniform4f(t.uFragColor, float32(col.R), float32(col.G), float32(col.B), 1.0)
		gl.DrawArrays(webgl.POINTS, 0, 1)

		count++
		// Stop processing
		if count > maxBodies {
			break
		}
	}

	/// FX Blurx4 TODO: better blur
	for i := 0; i < 4; i++ {
		gl.BindFrameBuffer(webgl.FRAMEBUFFER, t.rt[1])
		gl.Viewport(0, 0, width, height)
		gl.BindTexture(webgl.TEXTURE_2D, t.rtTex[0])
		t.qBlur.Render(gl)

		gl.BindFrameBuffer(webgl.FRAMEBUFFER, t.rt[0])
		gl.Viewport(0, 0, width, height)
		gl.BindTexture(webgl.TEXTURE_2D, t.rtTex[1])
		t.qBlur.Render(gl)
	}

	/// FX Threshold to Screen
	gl.BindFrameBuffer(webgl.FRAMEBUFFER, nil)
	gl.Viewport(0, 0, width, height)
	gl.BindTexture(webgl.TEXTURE_2D, t.rtTex[0])
	t.qThreshold.Render(gl)
}

func (t *Thing) AddCircle(mx, my float64) {
	if t.world.GetBodyCount() > maxBodies {
		// Check for the last on list and delete backwards:o
		var b *box2d.B2Body
		// theres is no M_last but we could cache it somewhere
		for b = t.world.GetBodyList(); b.M_next != nil; b = b.M_next {
		}
		// Search backwards for a circle (ignoring the walls/floors)
		for ; b != nil; b = b.M_prev {
			if _, ok := b.M_fixtureList.M_shape.(*box2d.B2CircleShape); ok {
				t.world.DestroyBody(b) // Destroy first found body
				break
			}
		}
	}
	obj1 := t.world.CreateBody(&box2d.B2BodyDef{
		Type:         box2d.B2BodyType.B2_dynamicBody,
		Position:     box2d.B2Vec2{X: mx, Y: my},
		Awake:        true,
		Active:       true,
		GravityScale: 1.0,
	})
	shape := box2d.NewB2CircleShape()
	shape.M_radius = 10 * worldScale
	ft := obj1.CreateFixture(shape, 1)
	ft.M_friction = 0.2
	ft.M_restitution = 0.6
}

//// SHADERS & Utils
const dotVertShader = `
attribute vec4 a_position;
void main () {
	vec4 lpos= vec4(a_position.xy*2.0-1.0, 0, 1);
	lpos.y = -lpos.y;
	gl_Position = lpos;
	gl_PointSize = 22.0/4.0;
}
`
const dotFragShader = `
precision mediump float;
uniform vec4 uFragColor;
void main () {
	vec2 pt = gl_PointCoord - vec2(0.5);
	if(pt.x*pt.x+pt.y*pt.y > 0.25)
	  discard;
	gl_FragColor = uFragColor;
}
`

const blurShader = `
precision mediump float;
uniform sampler2D u_image;
uniform vec2 u_textureSize;
varying vec2 v_texCoord;
void main() {
	vec2 onePixel = vec2(1,1) / u_textureSize;
	vec4 colorSum =
     texture2D(u_image, v_texCoord + onePixel * vec2(-1, -1)) + 
     texture2D(u_image, v_texCoord + onePixel * vec2( 0, -1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 1, -1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2(-1,  0)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 0,  0)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 1,  0)) +
     texture2D(u_image, v_texCoord + onePixel * vec2(-1,  1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 0,  1)) +
     texture2D(u_image, v_texCoord + onePixel * vec2( 1,  1));
  gl_FragColor = colorSum / 9.0;
}
`

const thresholdShader = `
precision mediump float;
uniform sampler2D u_image;
uniform vec2 u_textureSize;
varying vec2 v_texCoord;
void main() {
	float a;
	vec2 onePixel = vec2(1,1) / u_textureSize;
	vec4 col = texture2D(u_image,v_texCoord);
	if (col.a < 0.4) discard;
	if (col.a < 0.8 && col.a > 0.72) {
		a = texture2D(u_image, v_texCoord + onePixel * vec2(-1, 1)).a;
		if (a < col.a ) {
			col += 0.4;
		}
	} 
	gl_FragColor = vec4(col.rgb,1.0);
}
`

const vertQuad = `
attribute vec2 a_position;
attribute vec2 a_texCoord;
varying vec2 v_texCoord;
void main() {
   gl_Position = vec4((a_position * 2.0 - 1.0), 0, 1);
   v_texCoord = a_texCoord;
 }
`

type QuadFX struct {
	prog         *types.Program
	aPosition    int
	aTexCoord    int
	uTextureSize *types.UniformLocation
	quadBuf      *types.Buffer
	vertexData   []float32
}

func (q *QuadFX) Init(gl *webgl.RenderingContext, frag string) error {
	var err error
	q.prog, err = programFromSrc(gl, vertQuad, frag)
	if err != nil {
		return err
	}
	q.vertexData = []float32{
		0.0, 0.0, 1.0, 0.0, 0.0, 1.0,
		0.0, 1.0, 1.0, 0.0, 1.0, 1.0,
	}

	q.aPosition = gl.GetAttribLocation(q.prog, "a_position")
	q.aTexCoord = gl.GetAttribLocation(q.prog, "a_texCoord")
	q.uTextureSize = gl.GetUniformLocation(q.prog, "u_textureSize")

	q.quadBuf = gl.CreateBuffer()
	gl.BindBuffer(webgl.ARRAY_BUFFER, q.quadBuf)
	gl.BufferData(webgl.ARRAY_BUFFER, q.vertexData, webgl.STATIC_DRAW)
	return nil

}
func (q *QuadFX) Render(gl *webgl.RenderingContext) {
	gl.UseProgram(q.prog)
	gl.BindBuffer(webgl.ARRAY_BUFFER, q.quadBuf)

	gl.EnableVertexAttribArray(q.aPosition)
	gl.VertexAttribPointer(q.aPosition, 2, webgl.FLOAT, false, 0, 0)
	gl.EnableVertexAttribArray(q.aTexCoord)
	gl.VertexAttribPointer(q.aTexCoord, 2, webgl.FLOAT, false, 0, 0)

	gl.Uniform2f(q.uTextureSize, float32(width)/float32(resDiv), float32(height)/float32(resDiv))

	gl.DrawArrays(webgl.TRIANGLES, 0, 6)
	gl.DisableVertexAttribArray(q.aPosition)
	gl.DisableVertexAttribArray(q.aTexCoord)

}

// Helper funcs

// Render to framebuffer first, then framebuffer to screen
func compileShader(gl *webgl.RenderingContext, shaderType types.GLEnum, shaderSrc string) (*types.Shader, error) {
	shader := gl.CreateShader(shaderType)
	gl.ShaderSource(shader, shaderSrc)
	gl.CompileShader(shader)

	if !gl.GetShaderParameterCompileStatus(shader) {
		return nil, fmt.Errorf("could not compile shader: %v", gl.GetShaderInfoLog(shader))
	} else {
		return shader, nil
	}
}

func linkProgram(gl *webgl.RenderingContext, vertexShader, fragmentShader *types.Shader) (*types.Program, error) {
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	if !gl.GetProgramParameterLinkStatus(program) {
		return nil, fmt.Errorf("could not link program: %v", gl.GetProgramInfoLog(program))
	} else {
		return program, nil
	}
}

func programFromSrc(gl *webgl.RenderingContext, vertShaderSrc, fragShaderSrc string) (*types.Program, error) {
	vertexShader, err := compileShader(gl, webgl.VERTEX_SHADER, vertShaderSrc)
	if err != nil {
		return nil, err
	}
	fragShader, err := compileShader(gl, webgl.FRAGMENT_SHADER, fragShaderSrc)
	if err != nil {
		return nil, err
	}
	prog, err := linkProgram(gl, vertexShader, fragShader)
	if err != nil {
		return nil, err
	}
	return prog, nil
}

func createTexture(gl *webgl.RenderingContext, width, height int) *types.Texture {
	tex := gl.CreateTexture()
	gl.BindTexture(webgl.TEXTURE_2D, tex)
	gl.TexImage2Dui8(webgl.TEXTURE_2D, 0, webgl.RGBA, width, height, 0, webgl.RGBA, nil)

	// set the filtering so we don't need mips
	gl.TexParameterMinFilter(webgl.TEXTURE_2D, webgl.LINEAR)
	gl.TexParameterMagFilter(webgl.TEXTURE_2D, webgl.LINEAR)
	gl.TexParameterWrapS(webgl.TEXTURE_2D, webgl.CLAMP_TO_EDGE)
	gl.TexParameterWrapT(webgl.TEXTURE_2D, webgl.CLAMP_TO_EDGE)

	return tex
}

// Create framebuffer binded to texture
func createFB(gl *webgl.RenderingContext, tex *types.Texture) *types.FrameBuffer {
	frameBuffer := gl.CreateFrameBuffer()
	gl.BindFrameBuffer(webgl.FRAMEBUFFER, frameBuffer)
	gl.FramebufferTexture2D(webgl.FRAMEBUFFER, webgl.COLOR_ATTACHMENT0, webgl.TEXTURE_2D, tex, 0)
	return frameBuffer
}

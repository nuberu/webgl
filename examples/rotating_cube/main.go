//
// Based on the example: https://github.com/bobcob7/wasm-rotating-cube
// Buffers and shaders, shamelessly copied from: https://www.tutorialspoint.com/webgl/webgl_cube_rotation.html
//
package main

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/nuberu/webgl"
	"syscall/js"
	"unsafe"
)

var vertices = []float32{
	-1, -1, -1, 1, -1, -1, 1, 1, -1, -1, 1, -1,
	-1, -1, 1, 1, -1, 1, 1, 1, 1, -1, 1, 1,
	-1, -1, -1, -1, 1, -1, -1, 1, 1, -1, -1, 1,
	1, -1, -1, 1, 1, -1, 1, 1, 1, 1, -1, 1,
	-1, -1, -1, -1, -1, 1, 1, -1, 1, 1, -1, -1,
	-1, 1, -1, -1, 1, 1, 1, 1, 1, 1, 1, -1,
}
var colors = []float32{
	5, 3, 7, 5, 3, 7, 5, 3, 7, 5, 3, 7,
	1, 1, 3, 1, 1, 3, 1, 1, 3, 1, 1, 3,
	0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1,
	1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0,
	1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0,
	0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0,
}
var indices = []uint16{
	0, 1, 2, 0, 2, 3, 4, 5, 6, 4, 6, 7,
	8, 9, 10, 8, 10, 11, 12, 13, 14, 12, 14, 15,
	16, 17, 18, 16, 18, 19, 20, 21, 22, 20, 22, 23,
}

const vertShaderCode = `
attribute vec3 position;
uniform mat4 Pmatrix;
uniform mat4 Vmatrix;
uniform mat4 Mmatrix;
attribute vec3 color;
varying vec3 vColor;
void main(void) {
	gl_Position = Pmatrix*Vmatrix*Mmatrix*vec4(position, 1.);
	vColor = color;
}
`
const fragShaderCode = `
precision mediump float;
varying vec3 vColor;
void main(void) {
	gl_FragColor = vec4(vColor, 1.);
}
`

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

	if err == nil {
		// Create vertex buffer
		vertexBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
		gl.BufferData(webgl.ARRAY_BUFFER, vertices, webgl.STATIC_DRAW)

		// Create color buffer
		colorBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ARRAY_BUFFER, colorBuffer)
		gl.BufferData(webgl.ARRAY_BUFFER, colors, webgl.STATIC_DRAW)

		// Create index buffer
		indexBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)
		gl.BufferDataUI16(webgl.ELEMENT_ARRAY_BUFFER, indices, webgl.STATIC_DRAW)

		// Create a vertex shader object
		vertShader := gl.CreateVertexShader()
		gl.ShaderSource(vertShader, vertShaderCode)
		gl.CompileShader(vertShader)

		// Create fragment shader object
		fragShader := gl.CreateFragmentShader()
		gl.ShaderSource(fragShader, fragShaderCode)
		gl.CompileShader(fragShader)

		// Create a shader program object to store the combined shader program
		shaderProgram := gl.CreateProgram()
		gl.AttachShader(shaderProgram, vertShader)
		gl.AttachShader(shaderProgram, fragShader)
		gl.LinkProgram(shaderProgram)

		// Associate attributes to vertex shader
		PositionMatrix := gl.GetUniformLocation(shaderProgram, "Pmatrix")
		ViewMatrix := gl.GetUniformLocation(shaderProgram, "Vmatrix")
		ModelMatrix := gl.GetUniformLocation(shaderProgram, "Mmatrix")

		gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
		position := gl.GetAttribLocation(shaderProgram, "position")
		gl.VertexAttribPointer(position, 3, webgl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(position)

		gl.BindBuffer(webgl.ARRAY_BUFFER, colorBuffer)
		color := gl.GetAttribLocation(shaderProgram, "color")
		gl.VertexAttribPointer(color, 3, webgl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(color)

		gl.UseProgram(shaderProgram)

		// Set WebGL properties
		gl.ClearColor(0.5, 0.5, 0.5, 0.9)
		gl.ClearDepth(1.0)
		gl.Viewport(0, 0, width, height)
		gl.DepthFunc(webgl.LEQUAL)

		// Create Matrices
		ratio := float32(width / height)

		// Generate and apply projection matrix
		projMatrix := mgl32.Perspective(mgl32.DegToRad(45.0), ratio, 1, 100.0)
		var projMatrixBuffer *[16]float32
		projMatrixBuffer = (*[16]float32)(unsafe.Pointer(&projMatrix))
		gl.UniformMatrix4fv(PositionMatrix, false, []float32((*projMatrixBuffer)[:]))

		// Generate and apply view matrix
		viewMatrix := mgl32.LookAtV(mgl32.Vec3{3.0, 3.0, 3.0}, mgl32.Vec3{0.0, 0.0, 0.0}, mgl32.Vec3{0.0, 1.0, 0.0})
		var viewMatrixBuffer *[16]float32
		viewMatrixBuffer = (*[16]float32)(unsafe.Pointer(&viewMatrix))
		gl.UniformMatrix4fv(ViewMatrix, false, []float32((*viewMatrixBuffer)[:]))

		// Drawing the Cube
		movMatrix := mgl32.Ident4()
		var renderFrame js.Callback
		var tmark float32
		var rotation = float32(0)

		// Bind to element array for draw function
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)

		renderFrame = js.NewCallback(func(args []js.Value) {
			// Calculate rotation rate
			now := float32(args[0].Float())
			tdiff := now - tmark
			tmark = now
			rotation = rotation + float32(tdiff)/500

			// Do new model matrix calculations
			movMatrix = mgl32.HomogRotate3DX(0.5 * rotation)
			movMatrix = movMatrix.Mul4(mgl32.HomogRotate3DY(0.3 * rotation))
			movMatrix = movMatrix.Mul4(mgl32.HomogRotate3DZ(0.2 * rotation))

			// Convert model matrix to a JS TypedArray
			var modelMatrixBuffer *[16]float32
			modelMatrixBuffer = (*[16]float32)(unsafe.Pointer(&movMatrix))

			// Apply the model matrix
			gl.UniformMatrix4fv(ModelMatrix, false, []float32((*modelMatrixBuffer)[:]))

			// Clear the screen
			gl.Enable(webgl.DEPTH_TEST)
			gl.Clear(uint32(webgl.COLOR_BUFFER_BIT) & uint32(webgl.DEPTH_BUFFER_BIT))

			// Draw the cube
			gl.DrawElements(webgl.TRIANGLES, len(indices), webgl.UNSIGNED_SHORT, 0)

			// Call next frame
			js.Global().Call("requestAnimationFrame", renderFrame)
		})
		defer renderFrame.Release()

		js.Global().Call("requestAnimationFrame", renderFrame)

		done := make(chan struct{}, 0)
		<-done
	}
}

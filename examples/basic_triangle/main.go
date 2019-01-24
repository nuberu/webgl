//
// Based on the example: https://github.com/bobcob7/wasm-basic-triangle
//
package main

import (
	"github.com/nuberu/webgl"
	"syscall/js"
)

const vertexShaderCode = `
attribute vec3 coordinates;
		
void main(void) {
	gl_Position = vec4(coordinates, 1.0);
}
`

const fragmentShaderCode = `
void main(void) {
	gl_FragColor = vec4(0.0, 0.0, 1.0, 1.0);
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
		vertices := []float32{
			-0.5, 0.5, 0,
			-0.5, -0.5, 0,
			0.5, -0.5, 0,
		}

		vertexBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
		gl.BufferData(webgl.ARRAY_BUFFER, vertices, webgl.STATIC_DRAW)
		gl.BindBuffer(webgl.ARRAY_BUFFER, nil)

		indices := []uint32{
			2, 1, 0,
		}

		indexBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)
		gl.BufferDataUI(webgl.ELEMENT_ARRAY_BUFFER, indices, webgl.STATIC_DRAW)
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, nil)

		vertShader := gl.CreateVertexShader()
		gl.ShaderSource(vertShader, vertexShaderCode)
		gl.CompileShader(vertShader)

		fragShader := gl.CreateFragmentShader()
		gl.ShaderSource(fragShader, fragmentShaderCode)
		gl.CompileShader(fragShader)

		shaderProgram := gl.CreateProgram()
		gl.AttachShader(shaderProgram, vertShader)
		gl.AttachShader(shaderProgram, fragShader)
		gl.LinkProgram(shaderProgram)
		gl.UseProgram(shaderProgram)

		gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)

		coordinates := gl.GetAttribLocation(shaderProgram, "coordinates")
		gl.VertexAttribPointer(coordinates, 3, webgl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(coordinates)

		gl.ClearColor(0.5, 0.5, 0.5, 0.9)
		gl.Clear(uint32(webgl.COLOR_BUFFER_BIT))
		gl.Enable(webgl.DEPTH_TEST)
		gl.Viewport(0, 0, width, height)

		gl.DrawElements(webgl.TRIANGLES, len(indices), webgl.UNSIGNED_SHORT, 0)
	}
}

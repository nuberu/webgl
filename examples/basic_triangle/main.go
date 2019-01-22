//
// Based on the example: https://github.com/bobcob7/wasm-basic-triangle
//
package main

import (
	"fmt"
	"github.com/nuberu/webgl"
	"syscall/js"
)

func main() {
	fmt.Println("Iniciando...")

	// Init Canvas stuff
	doc := js.Global().Get("document")
	canvasEl := doc.Call("getElementById", "gocanvas")
	width := doc.Get("body").Get("clientWidth").Int()
	height := doc.Get("body").Get("clientHeight").Int()
	canvasEl.Set("width", width)
	canvasEl.Set("height", height)

	gl, err := webgl.FromCanvas(canvasEl)

	if err == nil {
		verticesNative := []float32{
			-0.5, 0.5, 0,
			-0.5, -0.5, 0,
			0.5, -0.5, 0,
		}
		vertices := js.TypedArrayOf(verticesNative)

		vertexBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
		gl.BufferData(webgl.ARRAY_BUFFER, vertices, webgl.STATIC_DRAW)
		gl.BindBuffer(webgl.ARRAY_BUFFER, nil)

		indicesNative := []uint32{
			2, 1, 0,
		}
		indices := js.TypedArrayOf(indicesNative)

		indexBuffer := gl.CreateBuffer()
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, indexBuffer)
		gl.BufferData(webgl.ELEMENT_ARRAY_BUFFER, indices, webgl.STATIC_DRAW)
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, nil)

		vertCode := `
	attribute vec3 coordinates;
		
	void main(void) {
		gl_Position = vec4(coordinates, 1.0);
	}`
		vertShader := gl.CreateVertexShader()
		gl.ShaderSource(vertShader, vertCode)
		gl.CompileShader(vertShader)

		fragCode := `
	void main(void) {
		gl_FragColor = vec4(0.0, 0.0, 1.0, 1.0);
	}`
		fragShader := gl.CreateFragmentShader()
		gl.ShaderSource(fragShader, fragCode)
		gl.CompileShader(fragShader)

		shaderProgram := gl.CreateProgram()
		gl.AttachShader(shaderProgram, vertShader)
		gl.AttachShader(shaderProgram, fragShader)
		gl.LinkProgram(shaderProgram)
		gl.UseProgram(shaderProgram)

		gl.BindBuffer(webgl.ARRAY_BUFFER, vertexBuffer)
		gl.BindBuffer(webgl.ELEMENT_ARRAY_BUFFER, vertexBuffer)

		coordinates := gl.GetAttribLocation(shaderProgram, "coordinates")
		gl.VertexAttribPointer(coordinates, 3, webgl.FLOAT, false, 0, 0)
		gl.EnableVertexAttribArray(coordinates)

		gl.ClearColor(0.5, 0.5, 0.5, 0.9)
		gl.Clear(uint32(webgl.COLOR_BUFFER_BIT))
		gl.Enable(webgl.DEPTH_TEST)
		gl.Viewport(0, 0, width, height)

		gl.DrawElements(webgl.TRIANGLES, len(indicesNative), webgl.UNSIGNED_SHORT, 0)
	} else {
		fmt.Println(err.Error())
	}
}

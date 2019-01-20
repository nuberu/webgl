package webgl

import (
	"errors"
	"syscall/js"
)

// WebGL context wrapper
type RenderingContext struct {
	loaded  bool
	js      js.Value
	version uint

	// Constant values
}

func WrapContext(jsContext js.Value) *RenderingContext {
	context := &RenderingContext{
		loaded: true,
		js:     jsContext,
	}

	return context
}

func FromCanvas(canvasEl js.Value) (*RenderingContext, error) {
	jsContext := canvasEl.Call("getContext", "webgl")
	if jsContext == js.Undefined() {
		jsContext = canvasEl.Call("getContext", "experimental-webgl")
	}
	if jsContext == js.Undefined() {
		return nil, errors.New("browser might not support webgl")
	}
	return WrapContext(jsContext), nil
}

// Specifies which texture unit to make active
func (c *RenderingContext) ActiveTexture(textureUnit uint32) {
	c.js.Call("activeTexture", textureUnit)
}

func (c *RenderingContext) AttachShader(program *Program, shader *Shader) {
	c.js.Call("attachShader", program.js, shader.js)
}

func (c *RenderingContext) BindAttribLocation(program *Program, index uint32, name string) {
	c.js.Call("bindAttribLocation", program.js, index, name)
}

func (c *RenderingContext) BindBuffer(target GLEnum, buffer *Buffer) {
	c.js.Call("bindBuffer", target, buffer.js)
}

func (c *RenderingContext) BindFramebuffer(target GLEnum, buffer *FrameBuffer) {
	c.js.Call("bindFramebuffer", target, buffer.js)
}

func (c *RenderingContext) BindRenderbuffer(target GLEnum, buffer *RenderBuffer) {
	c.js.Call("bindRenderbuffer", target, buffer.js)
}

func (c *RenderingContext) BindTexture(target GLEnum, texture *Texture) {
	c.js.Call("bindTexture", target, texture.js)
}

func (c *RenderingContext) BlendColor(r, g, b, a float32) {
	c.js.Call("blendColor", r, g, b, a)
}

func (c *RenderingContext) BlendEquation(mode GLEnum) {
	c.js.Call("blendEquation", mode)
}

func (c *RenderingContext) BlendEquationSeparate(modeRGB GLEnum, modeAlpha GLEnum) {
	c.js.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

func (c *RenderingContext) BlendFunc(sFactor GLEnum, dFactor GLEnum) {
	c.js.Call("blendFunc", sFactor, dFactor)
}

func (c *RenderingContext) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLEnum) {
	c.js.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (c *RenderingContext) BufferDataBySize(target GLEnum, size int, usage GLEnum) {
	c.js.Call("bufferData", target, size, usage)
}

func (c *RenderingContext) BufferData(target GLEnum, srcData js.Value, usage GLEnum) {
	c.js.Call("bufferData", target, srcData, usage)
}

// WebGL 2.0
func (c *RenderingContext) BufferDataWithOffset(target GLEnum, srcData js.Value, usage GLEnum, srcOffset, length uint32) {
	c.js.Call("bufferData", target, srcData, usage, srcOffset, length)
}

func (c *RenderingContext) BufferSubData(target GLEnum, offset int64, srcData js.Value) {
	c.js.Call("bufferSubData", target, offset, srcData)
}

// WebGL 2.0
func (c *RenderingContext) BufferSubDataWithOffset(target GLEnum, dstByteOffset int64, srcData js.Value, srcOffset, length uint32) {
	c.js.Call("bufferSubData", target, dstByteOffset, srcData, srcOffset, length)
}

func (c *RenderingContext) CheckFramebufferStatus(target GLEnum) {
	c.js.Call("checkFramebufferStatus", target)
}

func (c *RenderingContext) Clear(mask uint32) {
	c.js.Call("clear", mask)
}

func (c *RenderingContext) ClearColor(r, g, b, a float32) {
	c.js.Call("clearColor", r, g, b, a)
}

func (c *RenderingContext) ClearDepth(depth float32) {
	c.js.Call("clearDepth", depth)
}

func (c *RenderingContext) ClearStencil(s int) {
	c.js.Call("clearStencil", s)
}

func (c *RenderingContext) ColorMask(r, g, b, a float32) {
	c.js.Call("colorMask", r, g, b, a)
}

func (c *RenderingContext) Commit() {
	c.js.Call("commit")
}

func (c *RenderingContext) CompileShader(shader *Shader) {
	c.js.Call("compileShader", shader.js)
}

func (c *RenderingContext) CompressedTexImage2D(target GLEnum, level int, internalFormat GLEnum, width int, height int, border int) {
	c.js.Call("compressedTexImage2D", target, level, internalFormat, width, height, border)
}

func (c *RenderingContext) CompressedTexImage2DIn(target GLEnum, level int, internalFormat GLEnum, width int, height int, border int, pixels js.Value) {
	c.js.Call("compressedTexImage2D", target, level, internalFormat, width, height, border, pixels)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage2DOffset(target GLEnum, level int, internalFormat GLEnum, width int, height int, border int, imageSize int, offset int64) {
	c.js.Call("compressedTexImage2D", target, level, internalFormat, width, height, border, imageSize, offset)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage2DFromOffset(target GLEnum, level int, internalFormat GLEnum, width int, height int, border int, srcData js.Value, srcOffset int64, srcLengthOverride int) {
	c.js.Call("compressedTexImage2D", target, level, internalFormat, width, height, border, srcData, srcOffset, srcLengthOverride)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage3DOffset(target GLEnum, level int, internalFormat GLEnum, width int, height int, depth int, border int, imageSize int, offset int64) {
	c.js.Call("compressedTexImage3D", target, level, internalFormat, width, height, depth, border, imageSize, offset)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage3DFromOffset(target GLEnum, level int, internalFormat GLEnum, width int, height int, depth int, border int, srcData js.Value, srcOffset int64, srcLengthOverride int) {
	c.js.Call("compressedTexImage3D", target, level, internalFormat, width, height, depth, border, srcData, srcOffset, srcLengthOverride)
}

func (c *RenderingContext) CompressedTexSubImage2D(target GLEnum, level int, xOffset, yOffset int, width, height int, format GLEnum) {
	c.js.Call("compressedTexSubImage2D", target, level, xOffset, yOffset, width, height, format)
}

func (c *RenderingContext) CompressedTexSubImage2DIn(target GLEnum, level int, xOffset, yOffset int, width, height int, format GLEnum, pixels js.Value) {
	c.js.Call("compressedTexSubImage2D", target, level, xOffset, yOffset, width, height, format, pixels)
}

func (c *RenderingContext) CompressedTexSubImage2DFrom(target GLEnum, level int, xOffset, yOffset int, width, height int, format GLEnum, imageSize int, offset int64) {
	c.js.Call("compressedTexSubImage2D", target, level, xOffset, yOffset, width, height, format, imageSize, offset)
}

func (c *RenderingContext) CompressedTexSubImage2DFromOffset(target GLEnum, level int, xOffset, yOffset int, width, height int, format GLEnum, srcData js.Value, srcOffset int64, srcLengthOverride int) {
	c.js.Call("compressedTexSubImage2D", target, level, xOffset, yOffset, width, height, format, srcData, srcOffset, srcLengthOverride)
}

func (c *RenderingContext) CopyTexImage2D(target GLEnum, level int, internalFormat GLEnum, x, y int, width, height int, border int) {
	c.js.Call("copyTexImage2D", target, level, internalFormat, x, y, width, height, border)
}

func (c *RenderingContext) CopyTexSubImage2D(target GLEnum, level int, xOffset, yOffset int, x, y int, width, height int) {
	c.js.Call("copyTexSubImage2D", target, level, xOffset, yOffset, x, y, width, height)
}

func (c *RenderingContext) CreateBuffer() *Buffer {
	return &Buffer{
		js: c.js.Call("createBuffer"),
	}
}

func (c *RenderingContext) CreateFrameBuffer() *FrameBuffer {
	return &FrameBuffer{
		js: c.js.Call("createFramebuffer"),
	}
}

func (c *RenderingContext) CreateProgram() *Program {
	return &Program{
		js: c.js.Call("createProgram"),
	}
}

func (c *RenderingContext) CreateRenderBuffer() *RenderBuffer {
	return &RenderBuffer{
		js: c.js.Call("createRenderbuffer"),
	}
}

func (c *RenderingContext) CreateShader(shaderType ShaderType) *Shader {
	return &Shader{
		js:         c.js.Call("createShader", uint32(shaderType)),
		shaderType: shaderType,
	}
}

func (c *RenderingContext) CreateVertexShader() *Shader {
	return &Shader{
		js: c.js.Call("createShader", VERTEX_SHADER),
	}
}

func (c *RenderingContext) CreateTexture() *Texture {
	return &Texture{
		js: c.js.Call("createTexture"),
	}
}

func (c *RenderingContext) CullFace(mode GLEnum) {
	c.js.Call("cullFace", mode)
}

func (c *RenderingContext) GetDrawingBufferWidth() int {
	return c.js.Get("drawingBufferWidth").Int()
}

func (c *RenderingContext) GetDrawingBufferHeight() int {
	return c.js.Get("drawingBufferHeight").Int()
}

func (c *RenderingContext) GetCanvas() js.Value {
	return c.js.Get("canvas")
}

/*func (c *RenderingContext) GetContextAttributes() *Attributes {
	jsAttributes := c.js.Call("getContextAttributes")

	if jsAttributes != js.Null() {
		return &Attributes{
			Alpha:                        jsAttributes.Get("alpha").Bool(),
			Antialias:                    jsAttributes.Get("antialias").Bool(),
			Depth:                        jsAttributes.Get("depth").Bool(),
			PremultipliedAlpha:           jsAttributes.Get("premultipliedAlpha").Bool(),
			PreserveDrawingBuffer:        jsAttributes.Get("preserveDrawingBuffer").Bool(),
			Stencil:                      jsAttributes.Get("stencil").Bool(),
			PowerPreference:              webgl.CastPowerPreference(jsAttributes.Get("powerPreference").String()),
			FailIfMajorPerformanceCaveat: jsAttributes.Get("failIfMajorPerformanceCaveat").Bool(),
		}
	} else {
		return nil
	}
}*/

func (c *RenderingContext) IsContextLost() bool {
	return c.js.Call("isContextLost").Bool()
}

func (c *RenderingContext) DepthFunc(depth GLEnum) {
	c.js.Call("depthFunc", depth)
}

func (c *RenderingContext) DepthMask(flag bool) {
	c.js.Call("depthMask", flag)
}

func (c *RenderingContext) Scissor(x int, y int, width int, height int) {
	c.js.Call("scissor", x, y, width, height)
}

func (c *RenderingContext) StencilMask(mask uint) {
	c.js.Call("stencilMask", mask)
}

func (c *RenderingContext) StencilFunc(function GLEnum, ref int, mask uint) {
	c.js.Call("stencilFunc", function, ref, mask)
}

func (c *RenderingContext) StencilOp(fail GLEnum, zfail GLEnum, zpass GLEnum) {
	c.js.Call("stencilOp", c.js.Get(string(fail)), c.js.Get(string(zfail)), c.js.Get(string(zpass)))
}

func (c *RenderingContext) Viewport(x int, y int, width int, height int) {
	c.js.Call("viewport", x, y, width, height) // TODO: Add error handler
}

func (c *RenderingContext) DeleteTexture(texture *Texture) {
	c.js.Call("deleteTexture", texture.js)
}

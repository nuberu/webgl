package webgl

import (
	"errors"
	"github.com/nuberu/webgl/extensions"
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

func (c *RenderingContext) GetDrawingBufferWidth() int {
	return c.js.Get("drawingBufferWidth").Int()
}

func (c *RenderingContext) GetDrawingBufferHeight() int {
	return c.js.Get("drawingBufferHeight").Int()
}

func (c *RenderingContext) GetCanvas() js.Value {
	return c.js.Get("canvas")
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

func (c *RenderingContext) DeleteBuffer(buffer *Buffer) {
	c.js.Call("deleteBuffer", buffer.js)
	buffer.js = js.Undefined()
}

func (c *RenderingContext) DeleteFrameBuffer(framebuffer *FrameBuffer) {
	c.js.Call("deleteFramebuffer", framebuffer.js)
	framebuffer.js = js.Undefined()
}

func (c *RenderingContext) DeleteProgram(program *Program) {
	c.js.Call("deleteProgram", program.js)
	program.js = js.Undefined()
}

func (c *RenderingContext) DeleteRenderBuffer(renderbuffer *RenderBuffer) {
	c.js.Call("deleteFramebuffer", renderbuffer.js)
	renderbuffer.js = js.Undefined()
}

func (c *RenderingContext) DeleteShader(shader *Shader) {
	c.js.Call("deleteShader", shader.js)
	shader.js = js.Undefined()
}

func (c *RenderingContext) DeleteTexture(texture *Shader) {
	c.js.Call("deleteTexture", texture.js)
	texture.js = js.Undefined()
}

func (c *RenderingContext) DepthFunc(depth GLEnum) {
	c.js.Call("depthFunc", depth)
}

func (c *RenderingContext) DepthMask(flag bool) {
	c.js.Call("depthMask", flag)
}

func (c *RenderingContext) DepthRange(zNear, zFar float32) {
	c.js.Call("depthRange", zNear, zFar)
}

func (c *RenderingContext) DetachShader(program *Program, shader *Shader) {
	c.js.Call("detachShader", program.js, shader.js)
}

func (c *RenderingContext) Disable(cap GLEnum) {
	c.js.Call("disable", cap)
}

func (c *RenderingContext) DisableVertexAttribArray(index uint) {
	c.js.Call("disableVertexAttribArray", index)
}

func (c *RenderingContext) DrawArrays(mode GLEnum, first int, count int) {
	c.js.Call("drawArrays", mode, first, count)
}

func (c *RenderingContext) DrawElements(mode GLEnum, count int, valueType GLEnum, offset int64) {
	c.js.Call("drawElements", mode, count, valueType, offset)
}

func (c *RenderingContext) Enable(cap GLEnum) {
	c.js.Call("enable", cap)
}

func (c *RenderingContext) EnableVertexAttribArray(index uint32) {
	c.js.Call("enableVertexAttribArray", index)
}

func (c *RenderingContext) Finnish() {
	c.js.Call("finnish")
}

func (c *RenderingContext) Flush() {
	c.js.Call("flush")
}

func (c *RenderingContext) FramebufferRenderbuffer(target GLEnum, attachment GLEnum, renderBufferTarget GLEnum, renderBuffer *RenderBuffer) {
	c.js.Call("framebufferRenderbuffer", target, attachment, renderBufferTarget, renderBuffer.js)
}

func (c *RenderingContext) FramebufferTexture2D(target GLEnum, attachment GLEnum, texTarget GLEnum, texture *Texture, level int) {
	c.js.Call("framebufferTexture2D", target, attachment, texTarget, texture.js, level)
}

func (c *RenderingContext) FrontFace(mode GLEnum) {
	c.js.Call("frontFace", mode)
}

func (c *RenderingContext) GenerateMipmap(target GLEnum) {
	c.js.Call("generateMipmap", target)
}

func (c *RenderingContext) GetActiveAttrib(program *Program, index uint32) *ActiveInfo {
	info := c.js.Call("getActiveAttrib", program.js, index)
	return &ActiveInfo{
		name:     info.Get("name").String(),
		size:     info.Get("size").Int(),
		infoType: GLEnum(info.Get("type").Int()),
	}
}

func (c *RenderingContext) GetActiveUniform(program *Program, index uint32) *ActiveInfo {
	info := c.js.Call("getActiveUniform", program.js, index)
	return &ActiveInfo{
		name:     info.Get("name").String(),
		size:     info.Get("size").Int(),
		infoType: GLEnum(info.Get("type").Int()),
	}
}

func (c *RenderingContext) GetAttachedShaders(program *Program) []*Shader {
	shadersJs := c.js.Call("getAttachedShaders", program.js)
	shaders := make([]*Shader, 0, shadersJs.Length())
	for i := 0; i < shadersJs.Length(); i++ {
		shaders[i] = &Shader{js: shadersJs.Index(i)}
	}
	return shaders
}

func (c *RenderingContext) GetAttribLocation(program *Program, name string) int {
	return c.js.Call("getAttribLocation", program.js, name).Int()
}

func (c *RenderingContext) GetBufferParameter(target GLEnum, pName GLEnum) int {
	return c.js.Call("getBufferParameter", target, pName).Int()
}

func (c *RenderingContext) GetContextAttributes() *Attributes {
	attrJs := c.js.Call("getContextAttributes")
	if attrJs == js.Undefined() {
		return nil
	} else {
		return &Attributes{
			Alpha:                        attrJs.Get("alpha").Bool(),
			Antialias:                    attrJs.Get("alpha").Bool(),
			Depth:                        attrJs.Get("antialias").Bool(),
			FailIfMajorPerformanceCaveat: attrJs.Get("failIfMajorPerformanceCaveat").Bool(),
			PowerPreference:              PowerPreference(attrJs.Get("powerPreference").String()),
			PremultipliedAlpha:           attrJs.Get("premultipliedAlpha").Bool(),
			PreserveDrawingBuffer:        attrJs.Get("preserveDrawingBuffer").Bool(),
			Stencil:                      attrJs.Get("stencil").Bool(),
			Storage:                      attrJs.Get("storage").String(),
			WillReadFrequently:           attrJs.Get("willReadFrequently").Bool(),
		}
	}
}

func (c *RenderingContext) GetError() error {
	errorJs := c.js.Call("getError")

	switch GLEnum(errorJs.Int()) {
	case NO_ERROR:
		return nil
	case INVALID_ENUM:
		return errors.New("invalid enum")
	case INVALID_VALUE:
		return errors.New("invalid value")
	case INVALID_OPERATION:
		return errors.New("invalid operation")
	case INVALID_FRAMEBUFFER_OPERATION:
		return errors.New("invalid framebuffer operation")
	case OUT_OF_MEMORY:
		return errors.New("out of memory")
	case CONTEXT_LOST_WEBGL:
		return errors.New("context lost webgl")
	}

	return errors.New("unknown error")
}

func (c *RenderingContext) GetExtension(name string) *extensions.Extension {
	return extensions.LoadGenericExtension(c.js, name)
}

func (c *RenderingContext) GetExtensionLoseContext() *extensions.LoseContext {
	return extensions.LoadLoseContextExtension(c.js)
}

// TODO: Add other extensions

func (c *RenderingContext) GetFramebufferAttachmentParameterInt(target GLEnum, attachment GLEnum, pName GLEnum) int {
	return c.js.Call("getFramebufferAttachmentParameter", target, attachment, pName).Int()
}

func (c *RenderingContext) GetFramebufferAttachmentParameterEnum(target GLEnum, attachment GLEnum, pName GLEnum) GLEnum {
	return GLEnum(c.js.Call("getFramebufferAttachmentParameter", target, attachment, pName).Int())
}

func (c *RenderingContext) GetFramebufferAttachmentParameterRenderBuffer(target GLEnum, attachment GLEnum, pName GLEnum) *RenderBuffer {
	bufferJs := c.js.Call("getFramebufferAttachmentParameter", target, attachment, pName)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return &RenderBuffer{
			js: bufferJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetFramebufferAttachmentParameterTexture(target GLEnum, attachment GLEnum, pName GLEnum) *Texture {
	textureJs := c.js.Call("getFramebufferAttachmentParameter", target, attachment, pName)
	if textureJs != js.Undefined() && textureJs != js.Null() {
		return &Texture{
			js: textureJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameter(pName GLEnum) js.Value {
	return c.js.Call("getParameter", pName)
}

func (c *RenderingContext) GetParameterActiveTexture() GLEnum {
	return GLEnum(c.js.Call("getParameter", ACTIVE_TEXTURE).Int())
}

func (c *RenderingContext) GetParameterAliasedLineWidthRange() [2]float32 {
	arrJs := c.js.Call("getParameter", ALIASED_LINE_WIDTH_RANGE)
	var arr [2]float32
	arr[0] = float32(arrJs.Index(0).Float())
	arr[1] = float32(arrJs.Index(1).Float())
	return arr
}

func (c *RenderingContext) GetParameterAliasedPointSizeRange() [2]float32 {
	arrJs := c.js.Call("getParameter", ALIASED_POINT_SIZE_RANGE)
	var arr [2]float32
	arr[0] = float32(arrJs.Index(0).Float())
	arr[1] = float32(arrJs.Index(1).Float())
	return arr
}

func (c *RenderingContext) GetParameterAlphaBits() int {
	return c.js.Call("getParameter", ALPHA_BITS).Int()
}

func (c *RenderingContext) GetParameterArrayBufferBinding() *Buffer {
	bufferJs := c.js.Call("getParameter", ARRAY_BUFFER_BINDING)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return &Buffer{
			js: bufferJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterBlend() bool {
	return c.js.Call("getParameter", BLEND).Bool()
}

func (c *RenderingContext) GetParameterBlendColor() [4]float32 {
	arrJs := c.js.Call("getParameter", BLEND_COLOR)
	var arr [4]float32
	arr[0] = float32(arrJs.Index(0).Float())
	arr[1] = float32(arrJs.Index(1).Float())
	arr[2] = float32(arrJs.Index(2).Float())
	arr[3] = float32(arrJs.Index(3).Float())
	return arr
}

func (c *RenderingContext) GetParameterBlendDstAlpha() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_DST_ALPHA).Int())
}

func (c *RenderingContext) GetParameterBlendDstRgb() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_DST_RGB).Int())
}

func (c *RenderingContext) GetParameterBlendEquation() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_EQUATION).Int())
}

func (c *RenderingContext) GetParameterBlendEquationAlpha() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_EQUATION_ALPHA).Int())
}

func (c *RenderingContext) GetParameterBlendEquationRgb() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_EQUATION_RGB).Int())
}

func (c *RenderingContext) GetParameterBlendSrcAlpha() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_SRC_ALPHA).Int())
}

func (c *RenderingContext) GetParameterBlendSrcRgb() GLEnum {
	return GLEnum(c.js.Call("getParameter", BLEND_SRC_RGB).Int())
}

func (c *RenderingContext) GetParameterBlueBits() int {
	return c.js.Call("getParameter", BLUE_BITS).Int()
}

func (c *RenderingContext) GetParameterColorClearValue() [4]float32 {
	arrJs := c.js.Call("getParameter", COLOR_CLEAR_VALUE)
	var arr [4]float32
	arr[0] = float32(arrJs.Index(0).Float())
	arr[1] = float32(arrJs.Index(1).Float())
	arr[2] = float32(arrJs.Index(2).Float())
	arr[3] = float32(arrJs.Index(3).Float())
	return arr
}

func (c *RenderingContext) GetParameterColorWritemask() [4]bool {
	arrJs := c.js.Call("getParameter", COLOR_WRITEMASK)
	var arr [4]bool
	arr[0] = arrJs.Index(0).Bool()
	arr[1] = arrJs.Index(1).Bool()
	arr[2] = arrJs.Index(2).Bool()
	arr[3] = arrJs.Index(3).Bool()
	return arr
}

func (c *RenderingContext) GetParameterCompressedTextureFormats(pName string) []GLEnum {
	arrJs := c.js.Call("getParameter", COMPRESSED_TEXTURE_FORMATS)
	arr := make([]GLEnum, arrJs.Length())
	for i := 0; i < arrJs.Length(); i++ {
		arr[i] = GLEnum(arrJs.Index(i).Int())
	}
	return arr
}

func (c *RenderingContext) GetParameterCullFace() bool {
	return c.js.Call("getParameter", CULL_FACE).Bool()
}

func (c *RenderingContext) GetParameterCullFaceMode() GLEnum {
	return GLEnum(c.js.Call("getParameter", CULL_FACE_MODE).Int())
}

func (c *RenderingContext) GetParameterCurrentProgram() *Program {
	programJs := c.js.Call("getParameter", CURRENT_PROGRAM)
	if programJs != js.Undefined() && programJs != js.Null() {
		return &Program{
			js: programJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterDepthBits() float32 {
	return float32(c.js.Call("getParameter", DEPTH_BITS).Float())
}

func (c *RenderingContext) GetParameterDepthFunc() GLEnum {
	return GLEnum(c.js.Call("getParameter", DEPTH_FUNC).Int())
}

func (c *RenderingContext) GetParameterElementArrayBufferBinding() *Buffer {
	bufferJs := c.js.Call("getParameter", ELEMENT_ARRAY_BUFFER_BINDING)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return &Buffer{
			js: bufferJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterFrameBufferBinding() *FrameBuffer {
	frameBufferJs := c.js.Call("getParameter", FRAMEBUFFER_BINDING)
	if frameBufferJs != js.Undefined() && frameBufferJs != js.Null() {
		return &FrameBuffer{
			js: frameBufferJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterFrontFace() GLEnum {
	return GLEnum(c.js.Call("getParameter", FRONT_FACE).Int())
}

func (c *RenderingContext) GetParameterGenerateMipmapHint() GLEnum {
	return GLEnum(c.js.Call("getParameter", GENERATE_MIPMAP_HINT).Int())
}

func (c *RenderingContext) GetParameterGreenBits() int {
	return c.js.Call("getParameter", GREEN_BITS).Int()
}

func (c *RenderingContext) GetParameterImplementationColorReadFormat() GLEnum {
	return GLEnum(c.js.Call("getParameter", IMPLEMENTATION_COLOR_READ_FORMAT).Int())
}

func (c *RenderingContext) GetParameterImplementationColorReadType() GLEnum {
	return GLEnum(c.js.Call("getParameter", IMPLEMENTATION_COLOR_READ_TYPE).Int())
}

func (c *RenderingContext) GetParameterLineWidth() float32 {
	return float32(c.js.Call("getParameter", LINE_WIDTH).Float())
}

func (c *RenderingContext) GetParameterCombinedTextureImageUnits() int {
	return c.js.Call("getParameter", MAX_COMBINED_TEXTURE_IMAGE_UNITS).Int()
}

func (c *RenderingContext) GetParameterMaxCubeMapTextureSize() int {
	return c.js.Call("getParameter", MAX_CUBE_MAP_TEXTURE_SIZE).Int()
}

func (c *RenderingContext) GetParameterMaxFragmentUniformVectors() int {
	return c.js.Call("getParameter", MAX_FRAGMENT_UNIFORM_VECTORS).Int()
}

func (c *RenderingContext) GetParameterMaxRenderBufferSize() int {
	return c.js.Call("getParameter", MAX_RENDERBUFFER_SIZE).Int()
}

func (c *RenderingContext) GetParameterMaxTextureImageUnits() int {
	return c.js.Call("getParameter", MAX_TEXTURE_IMAGE_UNITS).Int()
}

func (c *RenderingContext) GetParameterMaxTextureSize() int {
	return c.js.Call("getParameter", MAX_TEXTURE_SIZE).Int()
}

func (c *RenderingContext) GetParameterMaxVaryingVectors() int {
	return c.js.Call("getParameter", MAX_VARYING_VECTORS).Int()
}

func (c *RenderingContext) GetParameterMaxVertexAttribs() int {
	return c.js.Call("getParameter", MAX_VERTEX_ATTRIBS).Int()
}

func (c *RenderingContext) GetParameterMaxVertexTextureImageUnits() int {
	return c.js.Call("getParameter", MAX_VERTEX_TEXTURE_IMAGE_UNITS).Int()
}

func (c *RenderingContext) GetParameterMaxVertexUniformVectors() int {
	return c.js.Call("getParameter", MAX_VERTEX_UNIFORM_VECTORS).Int()
}

func (c *RenderingContext) GetParameterMaxViewportDims() [2]float32 {
	arrJs := c.js.Call("getParameter", MAX_VIEWPORT_DIMS)
	var arr [2]float32
	arr[0] = float32(arrJs.Index(0).Float())
	arr[1] = float32(arrJs.Index(1).Float())
	return arr
}

func (c *RenderingContext) GetParameterPackAlignment() int {
	return c.js.Call("getParameter", PACK_ALIGNMENT).Int()
}

func (c *RenderingContext) GetParameterPolygonOffsetFactor() float32 {
	return float32(c.js.Call("getParameter", POLYGON_OFFSET_FACTOR).Float())
}

func (c *RenderingContext) GetParameterPolygonOffsetFill() bool {
	return c.js.Call("getParameter", POLYGON_OFFSET_FILL).Bool()
}

func (c *RenderingContext) GetParameterPolygonOffsetUnits() float32 {
	return float32(c.js.Call("getParameter", POLYGON_OFFSET_UNITS).Float())
}

func (c *RenderingContext) GetParameterRedBits() int {
	return c.js.Call("getParameter", RED_BITS).Int()
}

func (c *RenderingContext) GetParameterRenderBufferBinding() *RenderBuffer {
	bufferJs := c.js.Call("getParameter", RENDERBUFFER_BINDING)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return &RenderBuffer{
			js: bufferJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterRenderer() string {
	return c.js.Call("getParameter", RENDERER).String()
}

func (c *RenderingContext) GetParameterSampleBuffers() int {
	return c.js.Call("getParameter", SAMPLE_BUFFERS).Int()
}

func (c *RenderingContext) GetParameterSampleCoverageInvert() bool {
	return c.js.Call("getParameter", SAMPLE_COVERAGE_INVERT).Bool()
}

func (c *RenderingContext) GetParameterSampleCoverageValue() float32 {
	return float32(c.js.Call("getParameter", SAMPLE_COVERAGE_VALUE).Float())
}

func (c *RenderingContext) GetParameterSamples() int {
	return c.js.Call("getParameter", SAMPLES).Int()
}

func (c *RenderingContext) GetParameterScissorBox() [4]bool {
	arrJs := c.js.Call("getParameter", SCISSOR_BOX)
	var arr [4]bool
	arr[0] = arrJs.Index(0).Bool()
	arr[1] = arrJs.Index(1).Bool()
	arr[2] = arrJs.Index(2).Bool()
	arr[3] = arrJs.Index(3).Bool()
	return arr
}

func (c *RenderingContext) GetParameterScissorTest() bool {
	return c.js.Call("getParameter", SCISSOR_TEST).Bool()
}

func (c *RenderingContext) GetParameterShadingLanguageVersion() string {
	return c.js.Call("getParameter", SHADING_LANGUAGE_VERSION).String()
}

func (c *RenderingContext) GetParameterStencilBackFail() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_BACK_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilBackFunc() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_BACK_FUNC).Int())
}

func (c *RenderingContext) GetParameterStencilBackPassDepthFail() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_BACK_PASS_DEPTH_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilBackPassDepthPass() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_BACK_PASS_DEPTH_PASS).Int())
}

func (c *RenderingContext) GetParameterStencilBackRef() int {
	return c.js.Call("getParameter", STENCIL_BACK_REF).Int()
}

func (c *RenderingContext) GetParameterStencilBackValueMask() uint32 {
	return uint32(c.js.Call("getParameter", STENCIL_BACK_VALUE_MASK).Int())
}

func (c *RenderingContext) GetParameterStencilBackWritemask() uint32 {
	return uint32(c.js.Call("getParameter", STENCIL_BACK_WRITEMASK).Int())
}

func (c *RenderingContext) GetParameterStencilBits() int {
	return c.js.Call("getParameter", STENCIL_BITS).Int()
}

func (c *RenderingContext) GetParameterStencilClearValue() int {
	return c.js.Call("getParameter", STENCIL_CLEAR_VALUE).Int()
}

func (c *RenderingContext) GetParameterStencilFail() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilFunc() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_FUNC).Int())
}

func (c *RenderingContext) GetParameterStencilPassDepthFail() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_PASS_DEPTH_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilPassDepthPass() GLEnum {
	return GLEnum(c.js.Call("getParameter", STENCIL_PASS_DEPTH_PASS).Int())
}

func (c *RenderingContext) GetParameterStencilRef() int {
	return c.js.Call("getParameter", STENCIL_REF).Int()
}

func (c *RenderingContext) GetParameterStencilTest() bool {
	return c.js.Call("getParameter", STENCIL_TEST).Bool()
}

func (c *RenderingContext) GetParameterStencilValueMask() uint32 {
	return uint32(c.js.Call("getParameter", STENCIL_VALUE_MASK).Int())
}

func (c *RenderingContext) GetParameterStencilWritemask() uint32 {
	return uint32(c.js.Call("getParameter", STENCIL_WRITEMASK).Int())
}

func (c *RenderingContext) GetParameterSubpixelBits() int {
	return c.js.Call("getParameter", SUBPIXEL_BITS).Int()
}

func (c *RenderingContext) GetParameterTextureBinding2D() *Texture {
	textureJs := c.js.Call("getParameter", TEXTURE_BINDING_2D)
	if textureJs != js.Undefined() && textureJs != js.Null() {
		return &Texture{
			js: textureJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterTextureBindingCubeMap() *Texture {
	textureJs := c.js.Call("getParameter", TEXTURE_BINDING_CUBE_MAP)
	if textureJs != js.Undefined() && textureJs != js.Null() {
		return &Texture{
			js: textureJs,
		}
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterUnpackAlignment() int {
	return c.js.Call("getParameter", UNPACK_ALIGNMENT).Int()
}

func (c *RenderingContext) GetParameterUnpackColorspaceConversionWebGL() GLEnum {
	return GLEnum(c.js.Call("getParameter", UNPACK_COLORSPACE_CONVERSION_WEBGL).Int())
}

func (c *RenderingContext) GetParameterUnpackFlipYWebGL() bool {
	return c.js.Call("getParameter", UNPACK_FLIP_Y_WEBGL).Bool()
}

func (c *RenderingContext) GetParameterUnpackPremultiplyAlphaWebGL() bool {
	return c.js.Call("getParameter", UNPACK_PREMULTIPLY_ALPHA_WEBGL).Bool()
}

func (c *RenderingContext) GetParameterVendor() string {
	return c.js.Call("getParameter", VENDOR).String()
}

func (c *RenderingContext) GetParameterVersion() string {
	return c.js.Call("getParameter", VERSION).String()
}

func (c *RenderingContext) GetParameterViewport() [4]bool {
	arrJs := c.js.Call("getParameter", VIEWPORT)
	var arr [4]bool
	arr[0] = arrJs.Index(0).Bool()
	arr[1] = arrJs.Index(1).Bool()
	arr[2] = arrJs.Index(2).Bool()
	arr[3] = arrJs.Index(3).Bool()
	return arr
}

// TODO: Add WebGL 2.0 parameters

// TODO: Continue porting

func (c *RenderingContext) IsContextLost() bool {
	return c.js.Call("isContextLost").Bool()
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

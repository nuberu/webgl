package webgl

import (
	"errors"
	"github.com/nuberu/webgl/extensions"
	"github.com/nuberu/webgl/types"
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

func (c *RenderingContext) AttachShader(program *types.Program, shader *types.Shader) {
	c.js.Call("attachShader", program.GetJs(), shader.GetJs())
}

func (c *RenderingContext) BindAttribLocation(program *types.Program, index int, name string) {
	if program == nil {
		c.js.Call("bindAttribLocation", js.Null(), index, name)
	} else {
		c.js.Call("bindAttribLocation", program.GetJs(), index, name)
	}
}

func (c *RenderingContext) BindBuffer(target types.GLEnum, buffer *types.Buffer) {
	if buffer == nil {
		c.js.Call("bindBuffer", uint32(target), js.Null())
	} else {
		c.js.Call("bindBuffer", uint32(target), buffer.GetJs())
	}
}

func (c *RenderingContext) BindFramebuffer(target types.GLEnum, buffer *types.FrameBuffer) {
	if buffer == nil {
		c.js.Call("bindFramebuffer", uint32(target), js.Null())
	} else {
		c.js.Call("bindFramebuffer", uint32(target), buffer.GetJs())
	}
}

func (c *RenderingContext) BindRenderbuffer(target types.GLEnum, buffer *types.RenderBuffer) {
	if buffer == nil {
		c.js.Call("bindFramebuffer", uint32(target), js.Null())
	} else {
		c.js.Call("bindRenderbuffer", uint32(target), buffer.GetJs())
	}
}

func (c *RenderingContext) BindTexture(target types.GLEnum, texture *types.Texture) {
	if texture == nil {
		c.js.Call("bindTexture", uint32(target), js.Null())
	} else {
		c.js.Call("bindTexture", uint32(target), texture.GetJs())
	}
}

func (c *RenderingContext) BlendColor(r, g, b, a float32) {
	c.js.Call("blendColor", r, g, b, a)
}

func (c *RenderingContext) BlendEquation(mode types.GLEnum) {
	c.js.Call("blendEquation", uint32(mode))
}

func (c *RenderingContext) BlendEquationSeparate(modeRGB types.GLEnum, modeAlpha types.GLEnum) {
	c.js.Call("blendEquationSeparate", uint32(modeRGB), uint32(modeAlpha))
}

func (c *RenderingContext) BlendFunc(sFactor types.GLEnum, dFactor types.GLEnum) {
	c.js.Call("blendFunc", uint32(sFactor), uint32(dFactor))
}

func (c *RenderingContext) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha types.GLEnum) {
	c.js.Call("blendFuncSeparate", uint32(srcRGB), uint32(dstRGB), uint32(srcAlpha), uint32(dstAlpha))
}

func (c *RenderingContext) BufferDataBySize(target types.GLEnum, size int, usage types.GLEnum) {
	c.js.Call("bufferData", uint32(target), size, uint32(usage))
}

func (c *RenderingContext) BufferData(target types.GLEnum, srcData []float32, usage types.GLEnum) {
	c.js.Call("bufferData", uint32(target), js.TypedArrayOf(srcData), uint32(usage))
}

func (c *RenderingContext) BufferDataI(target types.GLEnum, srcData []int, usage types.GLEnum) {
	c.js.Call("bufferData",  uint32(target), js.TypedArrayOf(srcData), uint32(usage))
}

func (c *RenderingContext) BufferDataUI(target types.GLEnum, srcData []uint32, usage types.GLEnum) {
	c.js.Call("bufferData",  uint32(target), js.TypedArrayOf(srcData), uint32(usage))
}

func (c *RenderingContext) BufferDataUI16(target types.GLEnum, srcData []uint16, usage types.GLEnum) {
	c.js.Call("bufferData",  uint32(target), js.TypedArrayOf(srcData), uint32(usage))
}

// WebGL 2.0
func (c *RenderingContext) BufferDataWithOffset(target types.GLEnum, srcData []float32, usage types.GLEnum, srcOffset, length uint) {
	c.js.Call("bufferData",  uint32(target), js.TypedArrayOf(srcData), uint32(usage), srcOffset, length)
}

// WebGL 2.0
func (c *RenderingContext) BufferDataIWithOffset(target types.GLEnum, srcData []int, usage types.GLEnum, srcOffset, length uint) {
	c.js.Call("bufferData",  uint32(target), js.TypedArrayOf(srcData), uint32(usage), srcOffset, length)
}

// WebGL 2.0
func (c *RenderingContext) BufferDataUIWithOffset(target types.GLEnum, srcData []uint, usage types.GLEnum, srcOffset, length uint) {
	c.js.Call("bufferData",  uint32(target), js.TypedArrayOf(srcData), uint32(usage), srcOffset, length)
}

func (c *RenderingContext) BufferSubData(target types.GLEnum, offset int, srcData []float32) {
	c.js.Call("bufferSubData",  uint32(target), offset, js.TypedArrayOf(srcData))
}

func (c *RenderingContext) BufferSubDataI(target types.GLEnum, offset int, srcData []int) {
	c.js.Call("bufferSubData",  uint32(target), offset, js.TypedArrayOf(srcData))
}

func (c *RenderingContext) BufferSubDataUI(target types.GLEnum, offset int, srcData []uint) {
	c.js.Call("bufferSubData",  uint32(target), offset, js.TypedArrayOf(srcData))
}

// WebGL 2.0
func (c *RenderingContext) BufferSubDataWithOffset(target types.GLEnum, dstByteOffset int, srcData []float32, srcOffset, length uint) {
	c.js.Call("bufferSubData",  uint32(target), dstByteOffset, js.TypedArrayOf(srcData), srcOffset, length)
}

// WebGL 2.0
func (c *RenderingContext) BufferSubDataIWithOffset(target types.GLEnum, dstByteOffset int, srcData []int, srcOffset, length uint) {
	c.js.Call("bufferSubData",  uint32(target), dstByteOffset, js.TypedArrayOf(srcData), srcOffset, length)
}

// WebGL 2.0
func (c *RenderingContext) BufferSubDataUIWithOffset(target types.GLEnum, dstByteOffset int, srcData []uint, srcOffset, length uint) {
	c.js.Call("bufferSubData",  uint32(target), dstByteOffset, js.TypedArrayOf(srcData), srcOffset, length)
}

func (c *RenderingContext) CheckFramebufferStatus(target types.GLEnum) {
	c.js.Call("checkFramebufferStatus",  uint32(target))
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

func (c *RenderingContext) CompileShader(shader *types.Shader) {
	c.js.Call("compileShader", shader.GetJs())
}

func (c *RenderingContext) CompressedTexImage2D(target types.GLEnum, level int, internalFormat types.GLEnum, width int, height int, border int) {
	c.js.Call("compressedTexImage2D", uint32(target), level, uint32(internalFormat), width, height, border)
}

func (c *RenderingContext) CompressedTexImage2DIn(target types.GLEnum, level int, internalFormat types.GLEnum, width int, height int, border int, pixels []float32) {
	c.js.Call("compressedTexImage2D", uint32(target), level, uint32(internalFormat), width, height, border, js.TypedArrayOf(pixels))
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage2DOffset(target types.GLEnum, level int, internalFormat types.GLEnum, width int, height int, border int, imageSize int, offset int) {
	c.js.Call("compressedTexImage2D", uint32(target), level, uint32(internalFormat), width, height, border, imageSize, offset)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage2DFromOffset(target types.GLEnum, level int, internalFormat types.GLEnum, width int, height int, border int, srcData []float32, srcOffset int, srcLengthOverride int) {
	c.js.Call("compressedTexImage2D", uint32(target), level, uint32(internalFormat), width, height, border, js.TypedArrayOf(srcData), srcOffset, srcLengthOverride)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage3DOffset(target types.GLEnum, level int, internalFormat types.GLEnum, width int, height int, depth int, border int, imageSize int, offset int) {
	c.js.Call("compressedTexImage3D", uint32(target), level, uint32(internalFormat), width, height, depth, border, imageSize, offset)
}

// WebGL 2.0
func (c *RenderingContext) CompressedTexImage3DFromOffset(target types.GLEnum, level int, internalFormat types.GLEnum, width int, height int, depth int, border int, srcData []float32, srcOffset int, srcLengthOverride int) {
	c.js.Call("compressedTexImage3D", uint32(target), level, uint32(internalFormat), width, height, depth, border, js.TypedArrayOf(srcData), srcOffset, srcLengthOverride)
}

func (c *RenderingContext) CompressedTexSubImage2D(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum) {
	c.js.Call("compressedTexSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format))
}

func (c *RenderingContext) CompressedTexSubImage2DIn(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum, pixels []float32) {
	c.js.Call("compressedTexSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format), js.TypedArrayOf(pixels))
}

func (c *RenderingContext) CompressedTexSubImage2DFrom(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum, imageSize int, offset int) {
	c.js.Call("compressedTexSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format), imageSize, offset)
}

func (c *RenderingContext) CompressedTexSubImage2DFromOffset(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum, srcData []float32, srcOffset int, srcLengthOverride int) {
	c.js.Call("compressedTexSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format), js.TypedArrayOf(srcData), srcOffset, srcLengthOverride)
}

func (c *RenderingContext) CopyTexImage2D(target types.GLEnum, level int, internalFormat types.GLEnum, x, y int, width, height int, border int) {
	c.js.Call("copyTexImage2D", uint32(target), level, uint32(internalFormat), x, y, width, height, border)
}

func (c *RenderingContext) CopyTexSubImage2D(target types.GLEnum, level int, xOffset, yOffset int, x, y int, width, height int) {
	c.js.Call("copyTexSubImage2D", uint32(target), level, xOffset, yOffset, x, y, width, height)
}

func (c *RenderingContext) CreateBuffer() *types.Buffer {
	return types.NewBuffer(c.js.Call("createBuffer"))
}

func (c *RenderingContext) CreateFrameBuffer() *types.FrameBuffer {
	return types.NewFrameBuffer(c.js.Call("createFramebuffer"))
}

func (c *RenderingContext) CreateProgram() *types.Program {
	return types.NewProgram(c.js.Call("createProgram"))
}

func (c *RenderingContext) CreateRenderBuffer() *types.RenderBuffer {
	return types.NewRenderBuffer(c.js.Call("createRenderbuffer"))
}

func (c *RenderingContext) CreateShader(shaderType types.GLEnum) *types.Shader {
	return types.NewShader(c.js.Call("createShader", uint32(shaderType)))
}

func (c *RenderingContext) CreateFragmentShader() *types.Shader {
	return c.CreateShader(FRAGMENT_SHADER)
}

func (c *RenderingContext) CreateVertexShader() *types.Shader {
	return c.CreateShader(VERTEX_SHADER)
}

func (c *RenderingContext) CreateTexture() *types.Texture {
	return types.NewTexture(c.js.Call("createTexture"))
}

func (c *RenderingContext) CullFace(mode types.GLEnum) {
	c.js.Call("cullFace", uint32(mode))
}

func (c *RenderingContext) DeleteBuffer(buffer *types.Buffer) {
	c.js.Call("deleteBuffer", buffer.GetJs())
}

func (c *RenderingContext) DeleteFrameBuffer(framebuffer *types.FrameBuffer) {
	c.js.Call("deleteFramebuffer", framebuffer.GetJs())
}

func (c *RenderingContext) DeleteProgram(program *types.Program) {
	c.js.Call("deleteProgram", program.GetJs())
}

func (c *RenderingContext) DeleteRenderBuffer(renderbuffer *types.RenderBuffer) {
	c.js.Call("deleteFramebuffer", renderbuffer.GetJs())
}

func (c *RenderingContext) DeleteShader(shader *types.Shader) {
	c.js.Call("deleteShader", shader.GetJs())
}

func (c *RenderingContext) DeleteTexture(texture *types.Shader) {
	c.js.Call("deleteTexture", texture.GetJs())
}

func (c *RenderingContext) DepthFunc(depth types.GLEnum) {
	c.js.Call("depthFunc", uint32(depth))
}

func (c *RenderingContext) DepthMask(flag bool) {
	c.js.Call("depthMask", flag)
}

func (c *RenderingContext) DepthRange(zNear, zFar float32) {
	c.js.Call("depthRange", zNear, zFar)
}

func (c *RenderingContext) DetachShader(program *types.Program, shader *types.Shader) {
	c.js.Call("detachShader", program.GetJs(), shader.GetJs())
}

func (c *RenderingContext) Disable(cap types.GLEnum) {
	c.js.Call("disable", uint32(cap))
}

func (c *RenderingContext) DisableVertexAttribArray(index uint) {
	c.js.Call("disableVertexAttribArray", index)
}

func (c *RenderingContext) DrawArrays(mode types.GLEnum, first int, count int) {
	c.js.Call("drawArrays", uint32(mode), first, count)
}

func (c *RenderingContext) DrawElements(mode types.GLEnum, count int, valueType types.GLEnum, offset int64) {
	c.js.Call("drawElements", uint32(mode), count, uint32(valueType), offset)
}

func (c *RenderingContext) Enable(cap types.GLEnum) {
	c.js.Call("enable", uint32(cap))
}

func (c *RenderingContext) EnableVertexAttribArray(index int) {
	c.js.Call("enableVertexAttribArray", index)
}

func (c *RenderingContext) Finnish() {
	c.js.Call("finnish")
}

func (c *RenderingContext) Flush() {
	c.js.Call("flush")
}

func (c *RenderingContext) FramebufferRenderbuffer(target types.GLEnum, attachment types.GLEnum, renderBufferTarget types.GLEnum, renderBuffer *types.RenderBuffer) {
	c.js.Call("framebufferRenderbuffer", uint32(target), uint32(attachment), uint32(renderBufferTarget), renderBuffer.GetJs())
}

func (c *RenderingContext) FramebufferTexture2D(target types.GLEnum, attachment types.GLEnum, texTarget types.GLEnum, texture *types.Texture, level int) {
	c.js.Call("framebufferTexture2D", uint32(target), uint32(attachment), uint32(texTarget), texture.GetJs(), level)
}

func (c *RenderingContext) FrontFace(mode types.GLEnum) {
	c.js.Call("frontFace", uint32(mode))
}

func (c *RenderingContext) GenerateMipmap(target types.GLEnum) {
	c.js.Call("generateMipmap", uint32(target))
}

func (c *RenderingContext) GetActiveAttrib(program *types.Program, index uint) *types.ActiveInfo {
	info := c.js.Call("getActiveAttrib", program.GetJs(), index)
	return types.NewActiveInfo(
		info.Get("name").String(),
		info.Get("size").Int(),
		types.GLEnum(info.Get("type").Int()),
	)
}

func (c *RenderingContext) GetActiveUniform(program *types.Program, index uint) *types.ActiveInfo {
	info := c.js.Call("getActiveUniform", program.GetJs(), index)
	return types.NewActiveInfo(
		info.Get("name").String(),
		info.Get("size").Int(),
		types.GLEnum(info.Get("type").Int()),
	)
}

func (c *RenderingContext) GetAttachedShaders(program *types.Program) []*types.Shader {
	shadersJs := c.js.Call("getAttachedShaders", program.GetJs())
	shaders := make([]*types.Shader, 0, shadersJs.Length())
	for i := 0; i < shadersJs.Length(); i++ {
		shaders[i] = types.NewShader(shadersJs.Index(i))
	}
	return shaders
}

func (c *RenderingContext) GetAttribLocation(program *types.Program, name string) int {
	return c.js.Call("getAttribLocation", program.GetJs(), name).Int()
}

func (c *RenderingContext) GetBufferParameter(target types.GLEnum, pName types.GLEnum) int {
	return c.js.Call("getBufferParameter", uint32(target), uint32(pName)).Int()
}

func (c *RenderingContext) GetContextAttributes() *types.Attributes {
	attrJs := c.js.Call("getContextAttributes")
	if attrJs == js.Undefined() {
		return nil
	} else {
		return &types.Attributes{
			Alpha:                        attrJs.Get("alpha").Bool(),
			Antialias:                    attrJs.Get("alpha").Bool(),
			Depth:                        attrJs.Get("antialias").Bool(),
			FailIfMajorPerformanceCaveat: attrJs.Get("failIfMajorPerformanceCaveat").Bool(),
			PowerPreference:              types.PowerPreference(attrJs.Get("powerPreference").String()),
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

	switch types.GLEnum(errorJs.Int()) {
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

func (c *RenderingContext) GetFramebufferAttachmentParameterInt(target types.GLEnum, attachment types.GLEnum, pName types.GLEnum) int {
	return c.js.Call("getFramebufferAttachmentParameter", uint32(target), uint32(attachment), uint32(pName)).Int()
}

func (c *RenderingContext) GetFramebufferAttachmentParameterEnum(target types.GLEnum, attachment types.GLEnum, pName types.GLEnum) types.GLEnum {
	return types.GLEnum(c.js.Call("getFramebufferAttachmentParameter", uint32(target), uint32(attachment), uint32(pName)).Int())
}

func (c *RenderingContext) GetFramebufferAttachmentParameterRenderBuffer(target types.GLEnum, attachment types.GLEnum, pName types.GLEnum) *types.RenderBuffer {
	bufferJs := c.js.Call("getFramebufferAttachmentParameter", uint32(target), uint32(attachment), uint32(pName))
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return types.NewRenderBuffer(bufferJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetFramebufferAttachmentParameterTexture(target types.GLEnum, attachment types.GLEnum, pName types.GLEnum) *types.Texture {
	textureJs := c.js.Call("getFramebufferAttachmentParameter", uint32(target), uint32(attachment), uint32(pName))
	if textureJs != js.Undefined() && textureJs != js.Null() {
		return types.NewTexture(textureJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameter(pName types.GLEnum) js.Value {
	return c.js.Call("getParameter", uint32(pName))
}

func (c *RenderingContext) GetParameterActiveTexture() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", ACTIVE_TEXTURE).Int())
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

func (c *RenderingContext) GetParameterArrayBufferBinding() *types.Buffer {
	bufferJs := c.js.Call("getParameter", ARRAY_BUFFER_BINDING)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return types.NewBuffer(bufferJs)
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

func (c *RenderingContext) GetParameterBlendDstAlpha() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_DST_ALPHA).Int())
}

func (c *RenderingContext) GetParameterBlendDstRgb() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_DST_RGB).Int())
}

func (c *RenderingContext) GetParameterBlendEquation() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_EQUATION).Int())
}

func (c *RenderingContext) GetParameterBlendEquationAlpha() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_EQUATION_ALPHA).Int())
}

func (c *RenderingContext) GetParameterBlendEquationRgb() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_EQUATION_RGB).Int())
}

func (c *RenderingContext) GetParameterBlendSrcAlpha() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_SRC_ALPHA).Int())
}

func (c *RenderingContext) GetParameterBlendSrcRgb() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", BLEND_SRC_RGB).Int())
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

func (c *RenderingContext) GetParameterCompressedTextureFormats() []types.GLEnum {
	arrJs := c.js.Call("getParameter", COMPRESSED_TEXTURE_FORMATS)
	arr := make([]types.GLEnum, arrJs.Length())
	for i := 0; i < arrJs.Length(); i++ {
		arr[i] = types.GLEnum(arrJs.Index(i).Int())
	}
	return arr
}

func (c *RenderingContext) GetParameterCullFace() bool {
	return c.js.Call("getParameter", CULL_FACE).Bool()
}

func (c *RenderingContext) GetParameterCullFaceMode() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", CULL_FACE_MODE).Int())
}

func (c *RenderingContext) GetParameterCurrentProgram() *types.Program {
	programJs := c.js.Call("getParameter", CURRENT_PROGRAM)
	if programJs != js.Undefined() && programJs != js.Null() {
		return types.NewProgram(programJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterDepthBits() float32 {
	return float32(c.js.Call("getParameter", DEPTH_BITS).Float())
}

func (c *RenderingContext) GetParameterDepthFunc() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", DEPTH_FUNC).Int())
}

func (c *RenderingContext) GetParameterElementArrayBufferBinding() *types.Buffer {
	bufferJs := c.js.Call("getParameter", ELEMENT_ARRAY_BUFFER_BINDING)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return types.NewBuffer(bufferJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterFrameBufferBinding() *types.FrameBuffer {
	frameBufferJs := c.js.Call("getParameter", FRAMEBUFFER_BINDING)
	if frameBufferJs != js.Undefined() && frameBufferJs != js.Null() {
		return types.NewFrameBuffer(frameBufferJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterFrontFace() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", FRONT_FACE).Int())
}

func (c *RenderingContext) GetParameterGenerateMipmapHint() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", GENERATE_MIPMAP_HINT).Int())
}

func (c *RenderingContext) GetParameterGreenBits() int {
	return c.js.Call("getParameter", GREEN_BITS).Int()
}

func (c *RenderingContext) GetParameterImplementationColorReadFormat() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", IMPLEMENTATION_COLOR_READ_FORMAT).Int())
}

func (c *RenderingContext) GetParameterImplementationColorReadType() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", IMPLEMENTATION_COLOR_READ_TYPE).Int())
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

func (c *RenderingContext) GetParameterRenderBufferBinding() *types.RenderBuffer {
	bufferJs := c.js.Call("getParameter", RENDERBUFFER_BINDING)
	if bufferJs != js.Undefined() && bufferJs != js.Null() {
		return types.NewRenderBuffer(bufferJs)
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

func (c *RenderingContext) GetParameterStencilBackFail() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_BACK_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilBackFunc() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_BACK_FUNC).Int())
}

func (c *RenderingContext) GetParameterStencilBackPassDepthFail() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_BACK_PASS_DEPTH_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilBackPassDepthPass() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_BACK_PASS_DEPTH_PASS).Int())
}

func (c *RenderingContext) GetParameterStencilBackRef() int {
	return c.js.Call("getParameter", STENCIL_BACK_REF).Int()
}

func (c *RenderingContext) GetParameterStencilBackValueMask() uint {
	return uint(c.js.Call("getParameter", STENCIL_BACK_VALUE_MASK).Int())
}

func (c *RenderingContext) GetParameterStencilBackWritemask() uint {
	return uint(c.js.Call("getParameter", STENCIL_BACK_WRITEMASK).Int())
}

func (c *RenderingContext) GetParameterStencilBits() int {
	return c.js.Call("getParameter", STENCIL_BITS).Int()
}

func (c *RenderingContext) GetParameterStencilClearValue() int {
	return c.js.Call("getParameter", STENCIL_CLEAR_VALUE).Int()
}

func (c *RenderingContext) GetParameterStencilFail() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilFunc() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_FUNC).Int())
}

func (c *RenderingContext) GetParameterStencilPassDepthFail() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_PASS_DEPTH_FAIL).Int())
}

func (c *RenderingContext) GetParameterStencilPassDepthPass() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", STENCIL_PASS_DEPTH_PASS).Int())
}

func (c *RenderingContext) GetParameterStencilRef() int {
	return c.js.Call("getParameter", STENCIL_REF).Int()
}

func (c *RenderingContext) GetParameterStencilTest() bool {
	return c.js.Call("getParameter", STENCIL_TEST).Bool()
}

func (c *RenderingContext) GetParameterStencilValueMask() uint {
	return uint(c.js.Call("getParameter", STENCIL_VALUE_MASK).Int())
}

func (c *RenderingContext) GetParameterStencilWritemask() uint {
	return uint(c.js.Call("getParameter", STENCIL_WRITEMASK).Int())
}

func (c *RenderingContext) GetParameterSubpixelBits() int {
	return c.js.Call("getParameter", SUBPIXEL_BITS).Int()
}

func (c *RenderingContext) GetParameterTextureBinding2D() *types.Texture {
	textureJs := c.js.Call("getParameter", TEXTURE_BINDING_2D)
	if textureJs != js.Undefined() && textureJs != js.Null() {
		return types.NewTexture(textureJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterTextureBindingCubeMap() *types.Texture {
	textureJs := c.js.Call("getParameter", TEXTURE_BINDING_CUBE_MAP)
	if textureJs != js.Undefined() && textureJs != js.Null() {
		return types.NewTexture(textureJs)
	} else {
		return nil
	}
}

func (c *RenderingContext) GetParameterUnpackAlignment() int {
	return c.js.Call("getParameter", UNPACK_ALIGNMENT).Int()
}

func (c *RenderingContext) GetParameterUnpackColorspaceConversionWebGL() types.GLEnum {
	return types.GLEnum(c.js.Call("getParameter", UNPACK_COLORSPACE_CONVERSION_WEBGL).Int())
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

func (c *RenderingContext) GetProgramInfoLog(program *types.Program) string {
	return c.js.Call("getProgramInfoLog", program.GetJs()).String()
}

func (c *RenderingContext) GetProgramParameter(program *types.Program, pName types.GLEnum) js.Value {
	return c.js.Call("getProgramParameter", program.GetJs(), uint32(pName))
}

func (c *RenderingContext) GetProgramParameterDeleteStatus(program *types.Program) bool {
	return c.GetProgramParameter(program, DELETE_STATUS).Bool()
}

func (c *RenderingContext) GetProgramParameterLinkStatus(program *types.Program) bool {
	return c.GetProgramParameter(program, LINK_STATUS).Bool()
}

func (c *RenderingContext) GetProgramParameterValidateStatus(program *types.Program) bool {
	return c.GetProgramParameter(program, VALIDATE_STATUS).Bool()
}

func (c *RenderingContext) GetProgramParameterAttachedShaders(program *types.Program) int {
	return c.GetProgramParameter(program, ATTACHED_SHADERS).Int()
}

func (c *RenderingContext) GetProgramParameterActiveAttributes(program *types.Program) int {
	return c.GetProgramParameter(program, ACTIVE_ATTRIBUTES).Int()
}

func (c *RenderingContext) GetProgramParameterActiveUniforms(program *types.Program) int {
	return c.GetProgramParameter(program, ACTIVE_UNIFORMS).Int()
}

// WebGL 2.0
func (c *RenderingContext) GetProgramParameterTransformFeedbackBufferMode(program *types.Program) types.GLEnum {
	return types.GLEnum(c.GetProgramParameter(program, TRANSFORM_FEEDBACK_BUFFER_MODE).Int())
}

// WebGL 2.0
func (c *RenderingContext) GetProgramParameterTransformFeedbackVaryings(program *types.Program) int {
	return c.GetProgramParameter(program, TRANSFORM_FEEDBACK_VARYINGS).Int()
}

// WebGL 2.0
func (c *RenderingContext) GetProgramParameterActiveUniformBlocks(program *types.Program) int {
	return c.GetProgramParameter(program, ACTIVE_UNIFORM_BLOCKS).Int()
}

func (c *RenderingContext) GetRenderbufferParameter(target types.GLEnum, pName types.GLEnum) js.Value {
	return c.js.Call("getRenderbufferParameter", uint32(target), uint32(pName))
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferWidth(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_WIDTH).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferHeight(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_HEIGHT).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferInternalFormat(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetRenderbufferParameter(target, RENDERBUFFER_INTERNAL_FORMAT).Int())
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferGreenSize(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_GREEN_SIZE).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferBlueSize(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_BLUE_SIZE).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferRedSize(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_RED_SIZE).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferAlphaSize(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_ALPHA_SIZE).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferDepthSize(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_DEPTH_SIZE).Int()
}

func (c *RenderingContext) GetRenderbufferParameterRenderBufferStencilSize(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_STENCIL_SIZE).Int()
}

// WebGL 2.0
func (c *RenderingContext) GetRenderbufferParameterRenderBufferSamples(target types.GLEnum) int {
	return c.GetRenderbufferParameter(target, RENDERBUFFER_SAMPLES).Int()
}

func (c *RenderingContext) GetShaderInfoLog(shader *types.Shader) string {
	return c.js.Call("getShaderInfoLog", shader.GetJs()).String()
}

func (c *RenderingContext) GetShaderParameter(shader *types.Shader, pName types.GLEnum) js.Value {
	return c.js.Call("getShaderParameter", shader.GetJs(), uint32(pName))
}

func (c *RenderingContext) GetShaderParameterDeleteStatus(shader *types.Shader) bool {
	return c.GetShaderParameter(shader, DELETE_STATUS).Bool()
}

func (c *RenderingContext) GetShaderParameterCompileStatus(shader *types.Shader) bool {
	return c.GetShaderParameter(shader, COMPILE_STATUS).Bool()
}

func (c *RenderingContext) GetShaderParameterShaderType(shader *types.Shader) types.GLEnum {
	return types.GLEnum(c.GetShaderParameter(shader, SHADER_TYPE).Int())
}

func (c *RenderingContext) GetShaderPrecisionFormat(shaderType types.GLEnum, precisionType types.GLEnum) *types.ShaderPrecisionFormat {
	pFormatJs := c.js.Call("getShaderPrecisionFormat", uint32(shaderType), uint32(precisionType))
	return types.NewShaderPrecisionFormat(
		pFormatJs.Get("rangeMin").Int(),
		pFormatJs.Get("rangeMax").Int(),
		pFormatJs.Get("precision").Int(),
	)
}

func (c *RenderingContext) GetShaderSource(shader *types.Shader) string {
	return c.js.Call("getShaderSource", shader.GetJs()).String()
}

func (c *RenderingContext) GetSupportedExtensions() []extensions.Name {
	arrJs := c.js.Call("getSupportedExtensions")
	arr := make([]extensions.Name, arrJs.Length())
	for i := 0; i < arrJs.Length(); i++ {
		arr[i] = extensions.Name(arrJs.Index(i).String())
	}
	return arr
}

func (c *RenderingContext) GetTexParameter(target types.GLEnum, pName types.GLEnum) js.Value {
	return c.js.Call("getTexParameter", uint32(target), uint32(pName))
}

func (c *RenderingContext) GetTexParameterMagFilter(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_MAG_FILTER).Int())
}

func (c *RenderingContext) GetTexParameterMinFilter(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_MIN_FILTER).Int())
}

func (c *RenderingContext) GetTexParameterWrapS(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_WRAP_S).Int())
}

func (c *RenderingContext) GetTexParameterWrapT(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_WRAP_T).Int())
}

func (c *RenderingContext) GetTexParameterMaxAnisotropyExt(target types.GLEnum) float32 {
	return float32(c.GetTexParameter(target, extensions.TEXTURE_MAX_ANISOTROPY_EXT).Float())
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterBaseLevel(target types.GLEnum) int {
	return c.GetTexParameter(target, TEXTURE_BASE_LEVEL).Int()
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterCompareFunc(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_COMPARE_FUNC).Int())
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterCompareMode(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_COMPARE_MODE).Int())
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterImmutableFormat(target types.GLEnum) bool {
	return c.GetTexParameter(target, TEXTURE_IMMUTABLE_FORMAT).Bool()
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterImmutableLevels(target types.GLEnum) uint {
	return uint(c.GetTexParameter(target, TEXTURE_IMMUTABLE_LEVELS).Int())
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterMaxLever(target types.GLEnum) int {
	return c.GetTexParameter(target, TEXTURE_MAX_LEVEL).Int()
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterMaxLOD(target types.GLEnum) float32 {
	return float32(c.GetTexParameter(target, TEXTURE_MAX_LOD).Int())
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterMinLOD(target types.GLEnum) float32 {
	return float32(c.GetTexParameter(target, TEXTURE_MIN_LOD).Int())
}

// WebGL 2.0
func (c *RenderingContext) GetTexParameterWrapR(target types.GLEnum) types.GLEnum {
	return types.GLEnum(c.GetTexParameter(target, TEXTURE_WRAP_R).Int())
}

func (c *RenderingContext) GetUniform(program *types.Program, location *types.UniformLocation) js.Value {
	return c.js.Call("getUniform", program.GetJs(), location.GetJs())
}

func (c *RenderingContext) GetUniformLocation(program *types.Program, name string) *types.UniformLocation {
	return types.NewUniformLocation(c.js.Call("getUniformLocation", program.GetJs(), name))
}

func (c *RenderingContext) GetVertexAttrib(index int, pName types.GLEnum) js.Value {
	return c.js.Call("getVertexAttrib", index, uint32(pName))
}

func (c *RenderingContext) GetVertexAttribArrayBufferBinding(index int) *types.Buffer {
	return types.NewBuffer(c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_BUFFER_BINDING))
}

func (c *RenderingContext) GetVertexAttribArrayBufferEnabled(index int) bool {
	return c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_ENABLED).Bool()
}

func (c *RenderingContext) GetVertexAttribArraySize(index int) int {
	return c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_SIZE).Int()
}

func (c *RenderingContext) GetVertexAttribArrayStride(index int) int {
	return c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_STRIDE).Int()
}

func (c *RenderingContext) GetVertexAttribArrayType(index int) types.GLEnum {
	return types.GLEnum(c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_TYPE).Int())
}

func (c *RenderingContext) GetVertexAttribArrayNormalized(index int) bool {
	return c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_NORMALIZED).Bool()
}

func (c *RenderingContext) GetVertexAttribCurrentVertexAttrib(index int) [4]float32 {
	arrJs := c.js.Call("getVertexAttrib", index, CURRENT_VERTEX_ATTRIB)
	var arr [4]float32
	arr[0] = float32(arrJs.Index(0).Float())
	arr[1] = float32(arrJs.Index(1).Float())
	arr[2] = float32(arrJs.Index(2).Float())
	arr[3] = float32(arrJs.Index(3).Float())
	return arr
}

// WebGL 2.0
func (c *RenderingContext) GetVertexAttribArrayInteger(index int) bool {
	return c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_INTEGER).Bool()
}

func (c *RenderingContext) GetVertexAttribArrayDivisor(index int) int {
	return c.js.Call("getVertexAttrib", index, VERTEX_ATTRIB_ARRAY_DIVISOR).Int()
}

func (c *RenderingContext) GetVertexAttribArrayDivisorAngle(index int) int {
	return c.js.Call("getVertexAttrib", index, extensions.VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE).Int()
}

func (c *RenderingContext) GetVertexAttribOffset(index int, pName types.GLEnum) int {
	return c.js.Call("getVertexAttribOffset", index, uint32(pName)).Int()
}

func (c *RenderingContext) Hint(target types.GLEnum, mode types.GLEnum) {
	c.js.Call("hint", uint32(target), uint32(mode))
}

func (c *RenderingContext) IsBuffer(buffer js.Value) bool {
	return c.js.Call("isBuffer", buffer).Bool()
}

func (c *RenderingContext) IsContextLost() bool {
	return c.js.Call("isContextLost").Bool()
}

func (c *RenderingContext) IsEnabled(cap types.GLEnum) bool {
	return c.js.Call("isEnabled", uint32(cap)).Bool()
}

func (c *RenderingContext) IsFramebuffer(framebuffer js.Value) bool {
	return c.js.Call("isFramebuffer", framebuffer).Bool()
}

func (c *RenderingContext) IsProgram(program js.Value) bool {
	return c.js.Call("isProgram", program).Bool()
}

func (c *RenderingContext) IsRenderbuffer(renderbuffer js.Value) bool {
	return c.js.Call("isRenderbuffer", renderbuffer).Bool()
}

func (c *RenderingContext) IsShader(shader js.Value) bool {
	return c.js.Call("isShader", shader).Bool()
}

func (c *RenderingContext) IsTexture(texture js.Value) bool {
	return c.js.Call("isTexture", texture).Bool()
}

// Deprecated: Most browsers only support 1.0 value
func (c *RenderingContext) LineWidth(width float32) {
	c.js.Call("lineWidth", width)
}

func (c *RenderingContext) LinkProgram(program *types.Program) {
	c.js.Call("linkProgram", program.GetJs())
}

func (c *RenderingContext) PixelStorei(pName types.GLEnum, param int) {
	c.js.Call("pixelStorei", uint32(pName), param)
}

func (c *RenderingContext) PolygonOffset(factor float32, units float32) {
	c.js.Call("polygonOffset", factor, units)
}

func (c *RenderingContext) ReadPixels(x, y int, width, height int, format types.GLEnum, pixels js.TypedArray) {
	c.js.Call("readPixels", x, y, width, height, uint32(format), pixels)
}

// WebGL 2.0
func (c *RenderingContext) ReadPixelsOffset(x, y int, width, height int, format types.GLEnum, pixels js.TypedArray, dstOffset uint) {
	c.js.Call("readPixels", x, y, width, height, uint32(format), pixels, dstOffset)
}

// WebGL 2.0
func (c *RenderingContext) ReadPixelsOffsetPointer(x, y int, width, height int, format types.GLEnum, offset int) {
	c.js.Call("readPixels", x, y, width, height, uint32(format), offset)
}

func (c *RenderingContext) RenderbufferStorage(target types.GLEnum, internalFormat types.GLEnum, width, height int) {
	c.js.Call("renderbufferStorage", uint32(target), uint32(internalFormat), width, height)
}

func (c *RenderingContext) SampleCoverage(value float32, invert bool) {
	c.js.Call("sampleCoverage", value, invert)
}

func (c *RenderingContext) Scissor(x int, y int, width int, height int) {
	c.js.Call("scissor", x, y, width, height)
}

func (c *RenderingContext) ShaderSource(shader *types.Shader, source string) {
	c.js.Call("shaderSource", shader.GetJs(), source)
}

func (c *RenderingContext) StencilFunc(function types.GLEnum, ref int, mask uint32) {
	c.js.Call("stencilFunc", uint32(function), ref, mask)
}

func (c *RenderingContext) StencilFuncSeparate(face types.GLEnum, function types.GLEnum, ref int, mask uint32) {
	c.js.Call("stencilFuncSeparate", uint32(face), function, ref, mask)
}

func (c *RenderingContext) StencilMask(mask uint32) {
	c.js.Call("stencilMask", mask)
}

func (c *RenderingContext) StencilMaskSeparate(face types.GLEnum, mask uint32) {
	c.js.Call("stencilMaskSeparate", uint32(face), mask)
}

func (c *RenderingContext) StencilOp(fail types.GLEnum, zFail types.GLEnum, zPass types.GLEnum) {
	c.js.Call("stencilOp", uint32(fail), uint32(zFail), zPass)
}

func (c *RenderingContext) StencilOpSeparate(face types.GLEnum, fail types.GLEnum, zFail types.GLEnum, zPass types.GLEnum) {
	c.js.Call("stencilOpSeparate", face, uint32(fail), uint32(zFail), zPass)
}

func (c *RenderingContext) TexImage2D(target types.GLEnum, level int, internalFormat types.GLEnum, width, height int, border int, format types.GLEnum, dataType types.GLEnum, pixels js.TypedArray) {
	c.js.Call("texImage2D", uint32(target), level, uint32(internalFormat), width, height, border, uint32(format), uint32(dataType), pixels)
}

func (c *RenderingContext) TexImage2DHtmlElement(target types.GLEnum, level int, internalFormat types.GLEnum, format types.GLEnum, dataType types.GLEnum, pixels js.Value) {
	c.js.Call("texImage2D", uint32(target), level, uint32(internalFormat), uint32(format), uint32(dataType), pixels)
}

// WebGL 2.0
func (c *RenderingContext) TexImage2DOffset(target types.GLEnum, level int, internalFormat types.GLEnum, width, height int, border int, format types.GLEnum, dataType types.GLEnum, offset int) {
	c.js.Call("texImage2D", uint32(target), level, uint32(internalFormat), width, height, border, uint32(format), uint32(dataType), offset)
}

// WebGL 2.0
func (c *RenderingContext) TexImage2DHtmlElement2(target types.GLEnum, level int, internalFormat types.GLEnum, width, height int, border int, format types.GLEnum, dataType types.GLEnum, source js.Value) {
	c.js.Call("texImage2D", uint32(target), level, uint32(internalFormat), width, height, border, uint32(format), uint32(dataType), source)
}

// WebGL 2.0
func (c *RenderingContext) TexImage2D2(target types.GLEnum, level int, internalFormat types.GLEnum, width, height int, border int, format types.GLEnum, dataType types.GLEnum, srcData []float32, srcOffset int) {
	c.js.Call("texImage2D", uint32(target), level, uint32(internalFormat), width, height, border, uint32(format), uint32(dataType), js.TypedArrayOf(srcData), srcOffset)
}

func (c *RenderingContext) TexParameterf(target types.GLEnum, pName types.GLEnum, param float32) {
	c.js.Call("texParameterf", uint32(target), uint32(pName), param)
}

func (c *RenderingContext) TexParameteri(target types.GLEnum, pName types.GLEnum, param int) {
	c.js.Call("texParameteri", uint32(target), uint32(pName), param)
}

func (c *RenderingContext) texParameterEnum(target types.GLEnum, pName types.GLEnum, param types.GLEnum) {
	c.js.Call("texParameteri", uint32(target), uint32(pName), uint32(param))
}

func (c *RenderingContext) TexParameterMagFilter(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_MAG_FILTER, param)
}

func (c *RenderingContext) TexParameterMinFilter(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_MIN_FILTER, param)
}

func (c *RenderingContext) TexParameterWrapS(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_WRAP_S, param)
}

func (c *RenderingContext) TexParameterWrapT(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_WRAP_T, param)
}

func (c *RenderingContext) TexParameterMaxAnisotropyExt(target types.GLEnum, param float32) {
	c.TexParameterf(target, extensions.TEXTURE_MAX_ANISOTROPY_EXT, param)
}

func (c *RenderingContext) TexParameterBaseLevel(target types.GLEnum, param int) {
	c.TexParameteri(target, TEXTURE_BASE_LEVEL, param)
}

func (c *RenderingContext) TexParameterCompareFunc(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_COMPARE_FUNC, param)
}

func (c *RenderingContext) TexParameterCompareMode(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_COMPARE_MODE, param)
}

func (c *RenderingContext) TexParameterMaxLevel(target types.GLEnum, param int) {
	c.TexParameteri(target, TEXTURE_MAX_LEVEL, param)
}

func (c *RenderingContext) TexParameterMaxLOD(target types.GLEnum, param int) {
	c.TexParameteri(target, TEXTURE_MAX_LOD, param)
}

func (c *RenderingContext) TexParameterMinLOD(target types.GLEnum, param int) {
	c.TexParameteri(target, TEXTURE_MIN_LOD, param)
}

func (c *RenderingContext) TexParameterWrapR(target types.GLEnum, param types.GLEnum) {
	c.texParameterEnum(target, TEXTURE_WRAP_R, param)
}

func (c *RenderingContext) TexSubImage2D(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum, dataType types.GLEnum, pixels js.TypedArray) {
	c.js.Call("texSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format), uint32(dataType), pixels)
}

func (c *RenderingContext) TexSubImage2DHtmlElement(target types.GLEnum, level int, xOffset, yOffset int, format types.GLEnum, dataType types.GLEnum, pixels js.Value) {
	c.js.Call("texSubImage2D", uint32(target), level, xOffset, yOffset, uint32(format), uint32(dataType), pixels)
}

// WebGL 2.0
func (c *RenderingContext) TexSubImage2DOffset(target types.GLEnum, level int, xOffset, yOffset int, format types.GLEnum, dataType types.GLEnum, offset int) {
	c.js.Call("texSubImage2D", uint32(target), level, xOffset, yOffset, uint32(format), uint32(dataType), offset)
}

// WebGL 2.0
func (c *RenderingContext) TexSubImage2DOffset2(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum, dataType types.GLEnum, offset int) {
	c.js.Call("texSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format), uint32(dataType), offset)
}

// WebGL 2.0
func (c *RenderingContext) TexSubImage2DHtmlElement2(target types.GLEnum, level int, xOffset, yOffset int, width, height int, format types.GLEnum, dataType types.GLEnum, source js.Value) {
	c.js.Call("texSubImage2D", uint32(target), level, xOffset, yOffset, width, height, uint32(format), uint32(dataType), source)
}

func (c *RenderingContext) Uniform1f(location *types.UniformLocation, v0 float32) {
	c.js.Call("uniform1f", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform1fv(location *types.UniformLocation, value []float32) {
	c.js.Call("uniform1fv", location.GetJs(), js.TypedArrayOf(value))
}

func (c *RenderingContext) Uniform1i(location *types.UniformLocation, v0 int) {
	c.js.Call("uniform1i", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform1iv(location *types.UniformLocation, value []int) {
	c.js.Call("uniform1iv", location.GetJs(), value)
}

func (c *RenderingContext) Uniform2f(location *types.UniformLocation, v0 float32) {
	c.js.Call("uniform2f", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform2fv(location *types.UniformLocation, value []float32) {
	c.js.Call("uniform2fv", location.GetJs(), js.TypedArrayOf(value))
}

func (c *RenderingContext) Uniform2i(location *types.UniformLocation, v0 int) {
	c.js.Call("uniform2i", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform2iv(location *types.UniformLocation, value []int) {
	c.js.Call("uniform2iv", location.GetJs(), value)
}

func (c *RenderingContext) Uniform3f(location *types.UniformLocation, v0 float32) {
	c.js.Call("uniform3f", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform3fv(location *types.UniformLocation, value []float32) {
	c.js.Call("uniform3fv", location.GetJs(), js.TypedArrayOf(value))
}

func (c *RenderingContext) Uniform3i(location *types.UniformLocation, v0 int) {
	c.js.Call("uniform3i", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform3iv(location *types.UniformLocation, value []int) {
	c.js.Call("uniform3iv", location.GetJs(), value)
}

func (c *RenderingContext) Uniform4f(location *types.UniformLocation, v0 float32) {
	c.js.Call("uniform4f", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform4fv(location *types.UniformLocation, value []float32) {
	c.js.Call("uniform4fv", location.GetJs(), js.TypedArrayOf(value))
}

func (c *RenderingContext) Uniform4i(location *types.UniformLocation, v0 int) {
	c.js.Call("uniform4i", location.GetJs(), v0)
}

func (c *RenderingContext) Uniform4iv(location *types.UniformLocation, value []int) {
	c.js.Call("uniform4iv", location.GetJs(), value)
}

func (c *RenderingContext) UniformMatrix2fv(location *types.UniformLocation, transpose bool, value []float32) {
	c.js.Call("uniformMatrix2fv", location.GetJs(), transpose, js.TypedArrayOf(value))
}

func (c *RenderingContext) UniformMatrix3fv(location *types.UniformLocation, transpose bool, value []float32) {
	c.js.Call("uniformMatrix3fv", location.GetJs(), transpose, js.TypedArrayOf(value))
}

func (c *RenderingContext) UniformMatrix4fv(location *types.UniformLocation, transpose bool, value []float32) {
	c.js.Call("uniformMatrix4fv", location.GetJs(), transpose, js.TypedArrayOf(value))
}

func (c *RenderingContext) UseProgram(program *types.Program) {
	c.js.Call("useProgram", program.GetJs())
}

func (c *RenderingContext) ValidateProgram(program *types.Program) {
	c.js.Call("validateProgram", program.GetJs())
}

func (c *RenderingContext) VertexAttrib1f(index int, v0 float32) {
	c.js.Call("vertexAttrib1f", index, v0)
}

func (c *RenderingContext) VertexAttrib2f(index int, v0, v1 float32) {
	c.js.Call("vertexAttrib2f", index, v0, v1)
}

func (c *RenderingContext) VertexAttrib3f(index int, v0, v1, v2 float32) {
	c.js.Call("vertexAttrib3f", index, v0, v1, v2)
}

func (c *RenderingContext) VertexAttrib4f(index int, v0, v1, v2, v3 float32) {
	c.js.Call("vertexAttrib4f", index, v0, v1, v2, v3)
}

func (c *RenderingContext) VertexAttrib1fv(index int, value []float32) {
	c.js.Call("vertexAttrib1fv", index, js.TypedArrayOf(value))
}

func (c *RenderingContext) VertexAttrib2fv(index int, value []float32) {
	c.js.Call("vertexAttrib2fv", index, js.TypedArrayOf(value))
}

func (c *RenderingContext) VertexAttrib3fv(index int, value []float32) {
	c.js.Call("vertexAttrib3fv", index, js.TypedArrayOf(value))
}

func (c *RenderingContext) VertexAttrib4fv(index int, value []float32) {
	c.js.Call("vertexAttrib4fv", index, js.TypedArrayOf(value))
}

func (c *RenderingContext) VertexAttribPointer(index int, size int, aType types.GLEnum, normalized bool, stride int, offset int) {
	c.js.Call("vertexAttribPointer", index, size, uint32(aType), normalized, stride, offset)
}

func (c *RenderingContext) Viewport(x int, y int, width int, height int) {
	c.js.Call("viewport", x, y, width, height)
}

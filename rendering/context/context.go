package context

import (
	"github.com/nuberu/webgl"
	"syscall/js"
)

// WebGL context wrapper
type Context struct {
	loaded  bool
	js      js.Value
	version uint
}

func (c *Context) GetDrawingBufferWidth() int {
	return c.js.Get("drawingBufferWidth").Int()
}

func (c *Context) GetDrawingBufferHeight() int {
	return c.js.Get("drawingBufferHeight").Int()
}

func (c *Context) GetCanvas() js.Value {
	return c.js.Get("canvas")
}

func (c *Context) GetContextAttributes() *Attributes {
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
}

func (c *Context) Commit() {
	c.js.Call("commit")
}

func (c *Context) IsContextLost() bool {
	return c.js.Call("isContextLost").Bool()
}

func (c *Context) ClearColor(r float32, g float32, b float32, a float32) {
	c.js.Call("clearColor", r, g, b, a)
}

func (c *Context) ClearDepth(depth float32) {
	c.js.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(index int) {
	c.js.Call("clearStencil", index)
}

func (c *Context) ColorMask(r float32, g float32, b float32, a float32) {
	c.js.Call("colorMask", r, g, b, a)
}

func (c *Context) DepthFunc(depth webgl.Condition) {
	c.js.Call("depthFunc", c.js.Get(string(depth)))
}

func (c *Context) DepthMask(flag bool) {
	c.js.Call("depthMask", flag)
}

func (c *Context) Scissor(x int, y int, width int, height int) {
	c.js.Call("scissor", x, y, width, height)
}

func (c *Context) StencilMask(mask uint) {
	c.js.Call("stencilMask", mask)
}

func (c *Context) StencilFunc(function webgl.Condition, ref int, mask uint) {
	c.js.Call("stencilFunc", c.js.Get(string(function)), ref, mask)
}

func (c *Context) StencilOp(fail webgl.StencilFunc, zfail webgl.StencilFunc, zpass webgl.StencilFunc) {
	c.js.Call("stencilOp", c.js.Get(string(fail)), c.js.Get(string(zfail)), c.js.Get(string(zpass)))
}

func (c *Context) Viewport(x int, y int, width int, height int) {
	c.js.Call("viewport", x, y, width, height) // TODO: Add error handler
}

func (c *Context) NewBuffer() *webgl.Buffer {
	return webgl.NewBuffer(c.js)
}

func (c *Context) NewFramebuffer() *webgl.Framebuffer {
	return webgl.NewFramebuffer(c.js)
}

func (c *Context) NewTexture(textureType webgl.TextureType) *webgl.Texture {
	return webgl.NewTexture(c.js, textureType)
}

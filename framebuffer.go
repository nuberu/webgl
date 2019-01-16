package webgl

import (
	"syscall/js"
)

type Framebuffer struct {
	glc js.Value
	ptr js.Value
}

func NewFramebuffer(glContext js.Value) *Framebuffer {
	return &Framebuffer{
		glc: glContext,
		ptr: glContext.Call("createFramebuffer"),
	}
}

func (buf *Framebuffer) Delete() {
	buf.glc.Call("deleteFramebuffer", buf.ptr)
	buf.ptr = js.Null()
}

func (buf *Framebuffer) Bind(target FramebufferMode) {
	/*if c.version <= 1 && target != FramebufferDefaultMode {
		// TODO: WebGL version warning
	}*/
	buf.glc.Call("bindFramebuffer", buf.glc.Get(string(target)), buf.ptr)
}

func (buf *Framebuffer) CheckStatus() FramebufferStatus {
	status := buf.glc.Call("checkFramebufferStatus", buf.ptr)
	return CastFramebufferStatus(status.String())
}

func (buf *Framebuffer) ReadPixels() []byte {
	return nil // FIXME
}

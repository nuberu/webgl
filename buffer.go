package webgl

import (
	"syscall/js"
)

type Buffer struct {
	glc js.Value
	ptr js.Value
}

func NewBuffer(glContext js.Value) *Buffer {
	return &Buffer{
		glc: glContext,
		ptr: glContext.Call("createBuffer"),
	}
}

func (buf *Buffer) Delete() {
	buf.glc.Call("deleteBuffer", buf.ptr)
	buf.ptr = js.Null()
}

func (buf *Buffer) Bind(target BufferType) {
	/*if c.version <= 1 && target != FramebufferDefaultMode && target != ElementArrayBuffer {
		// TODO: WebGL version warning
	}*/
	buf.glc.Call("bindBuffer", buf.glc.Get(string(target)), buf.ptr)
}

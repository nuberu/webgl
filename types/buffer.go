package types

import "syscall/js"

type Buffer struct {
	js js.Value
}

func NewBuffer(pointer js.Value) *Buffer {
	return &Buffer{
		js: pointer,
	}
}

func (buffer *Buffer) GetJs() js.Value {
	return buffer.js
}

type FrameBuffer struct {
	js js.Value
}

func NewFrameBuffer(pointer js.Value) *FrameBuffer {
	return &FrameBuffer{
		js: pointer,
	}
}

func (buffer *FrameBuffer) GetJs() js.Value {
	return buffer.js
}

type RenderBuffer struct {
	js js.Value
}

func NewRenderBuffer(pointer js.Value) *RenderBuffer {
	return &RenderBuffer{
		js: pointer,
	}
}

func (buffer *RenderBuffer) GetJs() js.Value {
	return buffer.js
}

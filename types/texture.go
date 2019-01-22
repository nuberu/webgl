package types

import "syscall/js"

type Texture struct {
	js js.Value
}

func NewTexture(pointer js.Value) *Texture {
	return &Texture{
		js: pointer,
	}
}

func (tex *Texture) GetJs() js.Value {
	return tex.js
}

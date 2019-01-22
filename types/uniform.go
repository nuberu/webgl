package types

import "syscall/js"

type UniformLocation struct {
	js js.Value
}

func NewUniformLocation(pointer js.Value) *UniformLocation {
	return &UniformLocation{
		js: pointer,
	}
}

func (ul *UniformLocation) GetJs() js.Value {
	return ul.js
}

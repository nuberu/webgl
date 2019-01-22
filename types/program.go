package types

import "syscall/js"

type Program struct {
	js js.Value
}

func NewProgram(pointer js.Value) *Program {
	return &Program{
		js: pointer,
	}
}

func (program *Program) GetJs() js.Value {
	return program.js
}

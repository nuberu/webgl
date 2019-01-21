package extensions

import "syscall/js"

type ExtensionName string

type IExtension interface {
	GetJs() js.Value
}

type Extension struct {
	js js.Value
}

func LoadGenericExtension(glContext js.Value, name string) *Extension {
	return &Extension{
		js: glContext.Call("getExtension", name),
	}
}

func (ext *Extension) GetJs() js.Value {
	return ext.js
}

package webgl

import (
	"syscall/js"
)

type Buffer struct {
	js js.Value
}

type FrameBuffer struct {
	js js.Value
}

type RenderBuffer struct {
	js js.Value
}

type Program struct {
	js js.Value
}

type Shader struct {
	shaderType ShaderType
	js         js.Value
}

func (shader *Shader) GetType() ShaderType {
	return shader.shaderType
}

type Texture struct {
	js js.Value
}

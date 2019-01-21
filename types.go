package webgl

import (
	"syscall/js"
)

type ActiveInfo struct {
	name     string
	size     int
	infoType GLEnum
}

func (ai *ActiveInfo) GetName() string {
	return ai.name
}

func (ai *ActiveInfo) GetSize() int {
	return ai.size
}

func (ai *ActiveInfo) GetType() GLEnum {
	return ai.infoType
}

type Attributes struct {
	Alpha                        bool
	Antialias                    bool
	Depth                        bool
	FailIfMajorPerformanceCaveat bool
	PowerPreference              PowerPreference
	PremultipliedAlpha           bool
	PreserveDrawingBuffer        bool
	Stencil                      bool
	Storage                      string // Only Chromium
	WillReadFrequently           bool   // Only Firefox
}

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

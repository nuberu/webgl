package types

import "syscall/js"

type Shader struct {
	js js.Value
}

func NewShader(pointer js.Value) *Shader {
	return &Shader{
		js: pointer,
	}
}

func (shader *Shader) GetJs() js.Value {
	return shader.js
}

type ShaderPrecisionFormat struct {
	rangeMin  int
	rangeMax  int
	precision int
}

func NewShaderPrecisionFormat(rangeMin int, rangeMax int, precision int) *ShaderPrecisionFormat {
	return &ShaderPrecisionFormat{
		rangeMin:  rangeMin,
		rangeMax:  rangeMax,
		precision: precision,
	}
}

func (spf *ShaderPrecisionFormat) GetRangeMin() int {
	return spf.rangeMin
}

func (spf *ShaderPrecisionFormat) GetRangeMax() int {
	return spf.rangeMax
}

func (spf *ShaderPrecisionFormat) GetPrecision() int {
	return spf.precision
}

package webgl

import (
	"syscall/js"
)

type Texture struct {
	glContext js.Value
	pointer   js.Value
	ttype     js.Value
}

func NewTexture(glContext js.Value, textureType TextureType) *Texture {
	ttype := glContext.Get(string(textureType))
	return &Texture{
		glContext: glContext,
		pointer:   glContext.Call("createTexture"),
		ttype:     ttype,
	}
}

func (tex *Texture) Delete() {
	tex.glContext.Call("deleteTexture", tex.pointer)
	tex.pointer = js.Null()
}

func (tex *Texture) Bind() {
	/*if c.version <= 1 && target != FramebufferDefaultMode && target != ElementArrayBuffer {
		// TODO: WebGL version warning
	}*/
	tex.glContext.Call("bindTexture", tex.ttype, tex.pointer)
}

func (tex *Texture) SetMagFilter(filter TextureFilter) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureMagFilter)), tex.glContext.Get(string(filter)))
}

func (tex *Texture) GetMagFilter() TextureFilter {
	retValue := tex.glContext.Call("getTexParameter", tex.ttype, tex.glContext.Get(string(TextureMagFilter)))
	return CastTextureFilter(retValue.String())
}

func (tex *Texture) SetMinFilter(filter TextureFilter) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureMinFilter)), tex.glContext.Get(string(filter)))
}

func (tex *Texture) SetWrapS(filter TextureWrapMode) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureWrapS)), tex.glContext.Get(string(filter)))
}

func (tex *Texture) SetWrapT(filter TextureWrapMode) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureWrapT)), tex.glContext.Get(string(filter)))
}

func (tex *Texture) SetMaxAnisotropy(anisotropy float32) {
	tex.glContext.Call("texParameterf", tex.pointer, tex.glContext.Get(string(TextureMaxAnisotropyExt)), anisotropy)
}

func (tex *Texture) SetBaseLevel(level int) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureBaseLevel)), level)
}

func (tex *Texture) SetCompareFunc(comp Condition) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureCompareFunc)), tex.glContext.Get(string(comp)))
}

func (tex *Texture) SetCompareMode(mode TextureComparisonMode) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureCompareMode)), tex.glContext.Get(string(mode)))
}

func (tex *Texture) SetMaxLevel(level int) {
	tex.glContext.Call("texParameteri", tex.pointer, tex.glContext.Get(string(TextureMaxLevel)), level)
}

func (tex *Texture) SetMaxLevelOfDetail(levelOfDetail float32) {
	tex.glContext.Call("texParameterf", tex.pointer, tex.glContext.Get(string(TextureMaxLod)), levelOfDetail)
}

func (tex *Texture) SetMinLevelOfDetail(levelOfDetail float32) {
	tex.glContext.Call("texParameterf", tex.pointer, tex.glContext.Get(string(TextureMinLod)), levelOfDetail)
}

func (tex *Texture) SetWrapR(wrapMode TextureWrapMode) {
	tex.glContext.Call("texParameterf", tex.pointer, tex.glContext.Get(string(TextureWrapR)), tex.glContext.Get(string(wrapMode)))
}

func (tex *Texture) CompressedTexImage2D(level int, format Format, width int, height int, border int, pixels js.Value) {
	tex.glContext.Call("compressedTexSubImage2D", tex.pointer, level, tex.glContext.Get(string(format)), width, height, border, pixels)
}

func (tex *Texture) CompressedTexSubImage2D(level int, xOffset int, yOffset int, width int, height int, format Format, pixels js.Value) {
	tex.glContext.Call("compressedTexSubImage2D", tex.pointer, level, xOffset, yOffset, width, height, tex.glContext.Get(string(format)), pixels)
}

func (tex *Texture) CopyImage2D(level int, format Format, x int, y int, width int, height int, border int) {
	tex.glContext.Call("copyTexImage2D", tex.ttype, level, tex.glContext.Get(string(format)), x, y, width, height, border)
}

func (tex *Texture) CopySubImage2D(level int, xOffset int, yOffset int, x int, y int, width int, height int) {
	tex.glContext.Call("copyTexSubImage2D", tex.ttype, level, xOffset, yOffset, x, y, width, height)
}

func (tex *Texture) GenerateMipmap() {
	tex.glContext.Call("generateMipmap", tex.ttype)
}

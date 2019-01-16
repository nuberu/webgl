package webgl

type Condition string

const (
	Never        Condition = "NEVER"
	Always       Condition = "ALWAYS"
	Less         Condition = "LESS"
	LessEqual    Condition = "LESSEQ"
	Equal        Condition = "EQUAL"
	GreaterEqual Condition = "GREATEREQ"
	Greater      Condition = "GREATER"
	NotEqual     Condition = "NOTEQUAL"
)

type Precision string

const (
	HighPrecision   = "highp"
	MediumPrecision = "mediump"
	LowPrecision    = "lowp"
)

type StencilFunc string

const (
	Keep         StencilFunc = "KEEP"
	Replace      StencilFunc = "REPLACE"
	Increment    StencilFunc = "INCR"
	Decrement    StencilFunc = "DECR"
	Inverr       StencilFunc = "INVERT"
	IncreaseWrap StencilFunc = "INCR_WRAP"
	DecreaseWrap StencilFunc = "DECR_WRAP"
)

type BufferType string

const (
	ArrayBuffer        BufferType = "ARRAY_BUFFER"
	ElementArrayBuffer BufferType = "ELEMENT_ARRAY_BUFFER"
	// WebGL 2.0
	CopyReadBuffer          BufferType = "COPY_READ_BUFFER"
	CopyWriteBuffer         BufferType = "COPY_WRITE_BUFFER"
	TransformFeedbackBuffer BufferType = "TRANSFORM_FEEDBACK_BUFFER"
	UniformBuffer           BufferType = "UNIFORM_BUFFER"
	PixelPackBuffer         BufferType = "PIXEL_PACK_BUFFER"
	PixelUnpackBuffer       BufferType = "PIXEL_UNPACK_BUFFER"
)

type FramebufferMode string

const (
	FramebufferDefaultMode FramebufferMode = "FRAMEBUFFER"
	// WebGL 2.0
	FramebufferDrawMode FramebufferMode = "DRAW_FRAMEBUFFER"
	FramebufferReadMode FramebufferMode = "READ_FRAMEBUFFER"
)

type FramebufferStatus string

const (
	FramebufferStatusComplete                    FramebufferStatus = "FRAMEBUFFER_COMPLETE"
	FramebufferStatusIncompleteAttachment        FramebufferStatus = "FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	FramebufferStatusIncompleteMissingAttachment FramebufferStatus = "FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	FramebufferStatusIncompleteDimensions        FramebufferStatus = "FRAMEBUFFER_INCOMPLETE_DIMENSIONS"
	FramebufferStatusUnsupported                 FramebufferStatus = "FRAMEBUFFER_UNSUPPORTED"
	// WebGL 2.0
	FramebufferStatusIncompleteMultisample FramebufferStatus = "FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"
	RenderBufferSamples                    FramebufferStatus = "RENDERBUFFER_SAMPLES"
)

func CastFramebufferStatus(status string) FramebufferStatus {
	switch status {
	case string(FramebufferStatusComplete):
		return FramebufferStatusComplete
	case string(FramebufferStatusIncompleteAttachment):
		return FramebufferStatusIncompleteAttachment
	case string(FramebufferStatusIncompleteMissingAttachment):
		return FramebufferStatusIncompleteMissingAttachment
	case string(FramebufferStatusIncompleteDimensions):
		return FramebufferStatusIncompleteDimensions
	case string(FramebufferStatusUnsupported):
		return FramebufferStatusUnsupported
	case string(FramebufferStatusIncompleteMultisample):
		return FramebufferStatusIncompleteMultisample
	case string(RenderBufferSamples):
		return RenderBufferSamples
	}
	return FramebufferStatusIncompleteMissingAttachment
}

type PowerPreference string

const (
	DefaultPower PowerPreference = "default"
	// WebGL 2.0
	LowPower             PowerPreference = "low-power"
	HighPerformancePower PowerPreference = "high-performance"
)

func CastPowerPreference(name string) PowerPreference {
	var powerPreference PowerPreference

	switch name {
	case string(LowPower):
		powerPreference = LowPower
		break
	case string(HighPerformancePower):
		powerPreference = HighPerformancePower
		break
	case string(DefaultPower):
	default:
		powerPreference = DefaultPower
		break
	}

	return powerPreference
}

type TextureType string

const (
	Texture2d               TextureType = "TEXTURE_2D"
	TextureCubeMap          TextureType = "TEXTURE_CUBE_MAP"
	TextureCubeMapPositiveX TextureType = "TEXTURE_CUBE_MAP_POSITIVE_X"
	TextureCubeMapNegativeX TextureType = "TEXTURE_CUBE_MAP_NEGATIVE_X"
	TextureCubeMapPositiveY TextureType = "TEXTURE_CUBE_MAP_POSITIVE_Y"
	TextureCubeMapNegativeY TextureType = "TEXTURE_CUBE_MAP_NEGATIVE_Y"
	TextureCubeMapPositiveZ TextureType = "TEXTURE_CUBE_MAP_POSITIVE_Z"
	TextureCubeMapNegativeZ TextureType = "TEXTURE_CUBE_MAP_NEGATIVE_Z"
	// WebGL 2.0
	Texture3d      TextureType = "TEXTURE_3D"
	Texture2dArray TextureType = "TEXTURE_2D_ARRAY"
)

type Format string

const (
	CompressedRgbS3tcDxt1Ext  Format = "COMPRESSED_RGB_S3TC_DXT1_EXT"
	CompressedRgbaS3tcDxt1Ext Format = "COMPRESSED_RGBA_S3TC_DXT1_EXT"
	CompressedRgbaS3tcDxt3Ext Format = "COMPRESSED_RGBA_S3TC_DXT3_EXT"
	CompressedRgbaS3tcDxt5Ext Format = "COMPRESSED_RGBA_S3TC_DXT5_EXT"
)

type TextureFilter string

const (
	Linear               TextureFilter = "LINEAR"
	Nearest              TextureFilter = "NEAREST"
	NearestMipmapNearest TextureFilter = "NEAREST_MIPMAP_NEAREST"
	LinearMipmapNearest  TextureFilter = "LINEAR_MIPMAP_NEAREST"
	NearestMipmapLinear  TextureFilter = "NEAREST_MIPMAP_LINEAR "
	LinearMipmapLinear   TextureFilter = "LINEAR_MIPMAP_LINEAR "
)

func CastTextureFilter(name string) TextureFilter {
	switch name {
	case string(Nearest):
		return Nearest
	case string(Linear):
		return Linear
	case string(NearestMipmapNearest):
		return NearestMipmapNearest
	case string(LinearMipmapNearest):
		return LinearMipmapNearest
	case string(NearestMipmapLinear):
		return NearestMipmapLinear
	case string(LinearMipmapLinear):
		return LinearMipmapLinear
	}
	return Linear
}

type TextureWrapMode string

const (
	Repeat         TextureWrapMode = "REPEAT"
	ClampToEdge    TextureWrapMode = "CLAMP_TO_EDGE"
	MirroredRepeat TextureWrapMode = "MIRRORED_REPEAT"
)

type TextureComparisonMode string

const (
	None                TextureComparisonMode = "NONE"
	CompareRefToTexture TextureComparisonMode = "COMPARE_REF_TO_TEXTURE"
)

type TextureParameter string

const (
	TextureMagFilter        TextureParameter = "TEXTURE_MAG_FILTER"
	TextureMinFilter        TextureParameter = "TEXTURE_MIN_FILTER"
	TextureWrapS            TextureParameter = "TEXTURE_WRAP_S"
	TextureWrapT            TextureParameter = "TEXTURE_WRAP_T"
	TextureMaxAnisotropyExt TextureParameter = "TEXTURE_MAX_ANISOTROPY_EXT"
	TextureBaseLevel        TextureParameter = "TEXTURE_BASE_LEVEL"
	TextureCompareFunc      TextureParameter = "TEXTURE_COMPARE_FUNC"
	TextureCompareMode      TextureParameter = "TEXTURE_COMPARE_MODE"
	TextureMaxLevel         TextureParameter = "TEXTURE_MAX_LEVEL"
	TextureMaxLod           TextureParameter = "TEXTURE_MAX_LOD"
	TextureMinLod           TextureParameter = "TEXTURE_MIN_LOD"
	TextureWrapR            TextureParameter = "TEXTURE_WRAP_R"
)

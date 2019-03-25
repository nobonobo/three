package three

type AnimationActionLoopStyles js.Value
type Blending js.Value
type BlendingDstFactor js.Value
type BlendingEquation js.Value
type BlendingSrcFactor js.Value
type Colors js.Value
type Combine js.Value
type CompressedPixelFormat js.Value
type CullFace js.Value
type DepthModes js.Value
type DepthPackingStrategies js.Value
type FrontFaceDirection js.Value
type InterpolationEndingModes js.Value
type InterpolationModes js.Value
type MOUSE js.Value

var (
	LEFT   MOUSE = get("LEFT")
	MIDDLE MOUSE = get("MIDDLE")
	RIGHT  MOUSE = get("RIGHT")
)

type Mapping js.Value
type PixelFormat js.Value
type PixelType js.Value
type Shading js.Value
type ShadowMapType js.Value
type Side js.Value
type TextureDataType js.Value
type TextureEncoding js.Value
type TextureFilter js.Value
type ToneMapping js.Value
type TrianglesDrawModes js.Value
type Wrapping js.Value

var (
	AddEquation                             = &*BlendingEquation{Value: get("AddEquation")}
	AddOperation                            = &*Combine{Value: get("AddOperation")}
	AdditiveBlending                        = &*Blending{Value: get("AdditiveBlending")}
	AlphaFormat                             = &*PixelFormat{Value: get("AlphaFormat")}
	AlwaysDepth                             = &*DepthModes{Value: get("AlwaysDepth")}
	BackSide                                = &*Side{Value: get("BackSide")}
	BasicDepthPacking                       = &*DepthPackingStrategies{Value: get("BasicDepthPacking")}
	BasicShadowMap                          = &*ShadowMapType{Value: get("BasicShadowMap")}
	ByteType                                = &*TextureDataType{Value: get("ByteType")}
	CineonToneMapping                       = &*ToneMapping{Value: get("CineonToneMapping")}
	ClampToEdgeWrapping                     = &*Wrapping{Value: get("ClampToEdgeWrapping")}
	CubeReflectionMapping                   = &*Mapping{Value: get("CubeReflectionMapping")}
	CubeRefractionMapping                   = &*Mapping{Value: get("CubeRefractionMapping")}
	CubeUVReflectionMapping                 = &*Mapping{Value: get("CubeUVReflectionMapping")}
	CubeUVRefractionMapping                 = &*Mapping{Value: get("CubeUVRefractionMapping")}
	CullFaceBack                            = &*CullFace{Value: get("CullFaceBack")}
	CullFaceFront                           = &*CullFace{Value: get("CullFaceFront")}
	CullFaceFrontBack                       = &*CullFace{Value: get("CullFaceFrontBack")}
	CullFaceNone                            = &*CullFace{Value: get("CullFaceNone")}
	CustomBlending                          = &*Blending{Value: get("CustomBlending")}
	DepthFormat                             = &*PixelFormat{Value: get("DepthFormat")}
	DepthStencilFormat                      = &*PixelFormat{Value: get("DepthStencilFormat")}
	DoubleSide                              = &*Side{Value: get("DoubleSide")}
	DstAlphaFactor                          = &*BlendingDstFactor{Value: get("DstAlphaFactor")}
	DstColorFactor                          = &*BlendingDstFactor{Value: get("DstColorFactor")}
	EqualDepth                              = &*DepthModes{Value: get("EqualDepth")}
	EquirectangularReflectionMapping        = &*Mapping{Value: get("EquirectangularReflectionMapping")}
	EquirectangularRefractionMapping        = &*Mapping{Value: get("EquirectangularRefractionMapping")}
	FaceColors                              = &*Colors{Value: get("FaceColors")}
	FlatShading                             = &*Shading{Value: get("FlatShading")}
	FloatType                               = &*TextureDataType{Value: get("FloatType")}
	FrontFaceDirectionCCW                   = &*FrontFaceDirection{Value: get("FrontFaceDirectionCCW")}
	FrontFaceDirectionCW                    = &*FrontFaceDirection{Value: get("FrontFaceDirectionCW")}
	FrontSide                               = &*Side{Value: get("FrontSide")}
	GammaEncoding                           = &*TextureEncoding{Value: get("GammaEncoding")}
	GreaterDepth                            = &*DepthModes{Value: get("GreaterDepth")}
	GreaterEqualDepth                       = &*DepthModes{Value: get("GreaterEqualDepth")}
	HalfFloatType                           = &*TextureDataType{Value: get("HalfFloatType")}
	IntType                                 = &*TextureDataType{Value: get("IntType")}
	InterpolateDiscrete                     = &*InterpolationModes{Value: get("InterpolateDiscrete")}
	InterpolateLinear                       = &*InterpolationModes{Value: get("InterpolateLinear")}
	InterpolateSmooth                       = &*InterpolationModes{Value: get("InterpolateSmooth")}
	LessDepth                               = &*DepthModes{Value: get("LessDepth")}
	LessEqualDepth                          = &*DepthModes{Value: get("LessEqualDepth")}
	LinearEncoding                          = &*TextureEncoding{Value: get("LinearEncoding")}
	LinearFilter                            = &*TextureFilter{Value: get("LinearFilter")}
	LinearMipMapLinearFilter                = &*TextureFilter{Value: get("LinearMipMapLinearFilter")}
	LinearMipMapNearestFilter               = &*TextureFilter{Value: get("LinearMipMapNearestFilter")}
	LinearToneMapping                       = &*ToneMapping{Value: get("LinearToneMapping")}
	LogLuvEncoding                          = &*TextureEncoding{Value: get("LogLuvEncoding")}
	LoopOnce                                = &*AnimationActionLoopStyles{Value: get("LoopOnce")}
	LoopPingPong                            = &*AnimationActionLoopStyles{Value: get("LoopPingPong")}
	LoopRepeat                              = &*AnimationActionLoopStyles{Value: get("LoopRepeat")}
	LuminanceAlphaFormat                    = &*PixelFormat{Value: get("LuminanceAlphaFormat")}
	LuminanceFormat                         = &*PixelFormat{Value: get("LuminanceFormat")}
	MaxEquation                             = &*BlendingEquation{Value: get("MaxEquation")}
	MinEquation                             = &*BlendingEquation{Value: get("MinEquation")}
	MirroredRepeatWrapping                  = &*Wrapping{Value: get("MirroredRepeatWrapping")}
	MixOperation                            = &*Combine{Value: get("MixOperation")}
	MultiplyBlending                        = &*Blending{Value: get("MultiplyBlending")}
	MultiplyOperation                       = &*Combine{Value: get("MultiplyOperation")}
	NearestFilter                           = &*TextureFilter{Value: get("NearestFilter")}
	NearestMipMapLinearFilter               = &*TextureFilter{Value: get("NearestMipMapLinearFilter")}
	NearestMipMapNearestFilter              = &*TextureFilter{Value: get("NearestMipMapNearestFilter")}
	NeverDepth                              = &*DepthModes{Value: get("NeverDepth")}
	NoBlending                              = &*Blending{Value: get("NoBlending")}
	NoColors                                = &*Colors{Value: get("NoColors")}
	NoToneMapping                           = &*ToneMapping{Value: get("NoToneMapping")}
	NormalBlending                          = &*Blending{Value: get("NormalBlending")}
	NotEqualDepth                           = &*DepthModes{Value: get("NotEqualDepth")}
	OneFactor                               = &*BlendingDstFactor{Value: get("OneFactor")}
	OneMinusDstAlphaFactor                  = &*BlendingDstFactor{Value: get("OneMinusDstAlphaFactor")}
	OneMinusDstColorFactor                  = &*BlendingDstFactor{Value: get("OneMinusDstColorFactor")}
	OneMinusSrcAlphaFactor                  = &*BlendingDstFactor{Value: get("OneMinusSrcAlphaFactor")}
	OneMinusSrcColorFactor                  = &*BlendingDstFactor{Value: get("OneMinusSrcColorFactor")}
	PCFShadowMap                            = &*ShadowMapType{Value: get("PCFShadowMap")}
	PCFSoftShadowMap                        = &*ShadowMapType{Value: get("PCFSoftShadowMap")}
	REVISION                         string = get("REVISION").String()
	RGBADepthPacking                        = &*DepthPackingStrategies{Value: get("RGBADepthPacking")}
	RGBAFormat                              = &*PixelFormat{Value: get("RGBAFormat")}
	RGBA_ASTC_10x10_Format                  = &*CompressedPixelFormat{Value: get("RGBA_ASTC_10x10_Format")}
	RGBA_ASTC_10x5_Format                   = &*CompressedPixelFormat{Value: get("RGBA_ASTC_10x5_Format")}
	RGBA_ASTC_10x6_Format                   = &*CompressedPixelFormat{Value: get("RGBA_ASTC_10x6_Format")}
	RGBA_ASTC_10x8_Format                   = &*CompressedPixelFormat{Value: get("RGBA_ASTC_10x8_Format")}
	RGBA_ASTC_12x10_Format                  = &*CompressedPixelFormat{Value: get("RGBA_ASTC_12x10_Format")}
	RGBA_ASTC_12x12_Format                  = &*CompressedPixelFormat{Value: get("RGBA_ASTC_12x12_Format")}
	RGBA_ASTC_4x4_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_4x4_Format")}
	RGBA_ASTC_5x4_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_5x4_Format")}
	RGBA_ASTC_5x5_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_5x5_Format")}
	RGBA_ASTC_6x5_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_6x5_Format")}
	RGBA_ASTC_6x6_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_6x6_Format")}
	RGBA_ASTC_8x5_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_8x5_Format")}
	RGBA_ASTC_8x6_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_8x6_Format")}
	RGBA_ASTC_8x8_Format                    = &*CompressedPixelFormat{Value: get("RGBA_ASTC_8x8_Format")}
	RGBA_PVRTC_2BPPV1_Format                = &*CompressedPixelFormat{Value: get("RGBA_PVRTC_2BPPV1_Format")}
	RGBA_PVRTC_4BPPV1_Format                = &*CompressedPixelFormat{Value: get("RGBA_PVRTC_4BPPV1_Format")}
	RGBA_S3TC_DXT1_Format                   = &*CompressedPixelFormat{Value: get("RGBA_S3TC_DXT1_Format")}
	RGBA_S3TC_DXT3_Format                   = &*CompressedPixelFormat{Value: get("RGBA_S3TC_DXT3_Format")}
	RGBA_S3TC_DXT5_Format                   = &*CompressedPixelFormat{Value: get("RGBA_S3TC_DXT5_Format")}
	RGBDEncoding                            = &*TextureEncoding{Value: get("RGBDEncoding")}
	RGBEEncoding                            = &*TextureEncoding{Value: get("RGBEEncoding")}
	RGBEFormat                              = &*PixelFormat{Value: get("RGBEFormat")}
	RGBFormat                               = &*PixelFormat{Value: get("RGBFormat")}
	RGBM16Encoding                          = &*TextureEncoding{Value: get("RGBM16Encoding")}
	RGBM7Encoding                           = &*TextureEncoding{Value: get("RGBM7Encoding")}
	RGB_ETC1_Format                         = &*CompressedPixelFormat{Value: get("RGB_ETC1_Format")}
	RGB_PVRTC_2BPPV1_Format                 = &*CompressedPixelFormat{Value: get("RGB_PVRTC_2BPPV1_Format")}
	RGB_PVRTC_4BPPV1_Format                 = &*CompressedPixelFormat{Value: get("RGB_PVRTC_4BPPV1_Format")}
	RGB_S3TC_DXT1_Format                    = &*CompressedPixelFormat{Value: get("RGB_S3TC_DXT1_Format")}
	RedFormat                               = &*PixelFormat{Value: get("RedFormat")}
	ReinhardToneMapping                     = &*ToneMapping{Value: get("ReinhardToneMapping")}
	RepeatWrapping                          = &*Wrapping{Value: get("RepeatWrapping")}
	ReverseSubtractEquation                 = &*BlendingEquation{Value: get("ReverseSubtractEquation")}
	ShortType                               = &*TextureDataType{Value: get("ShortType")}
	SmoothShading                           = &*Shading{Value: get("SmoothShading")}
	SphericalReflectionMapping              = &*Mapping{Value: get("SphericalReflectionMapping")}
	SrcAlphaFactor                          = &*BlendingDstFactor{Value: get("SrcAlphaFactor")}
	SrcAlphaSaturateFactor                  = &*BlendingSrcFactor{Value: get("SrcAlphaSaturateFactor")}
	SrcColorFactor                          = &*BlendingDstFactor{Value: get("SrcColorFactor")}
	SubtractEquation                        = &*BlendingEquation{Value: get("SubtractEquation")}
	SubtractiveBlending                     = &*Blending{Value: get("SubtractiveBlending")}
	TriangleFanDrawMode                     = &*TrianglesDrawModes{Value: get("TriangleFanDrawMode")}
	TriangleStripDrawMode                   = &*TrianglesDrawModes{Value: get("TriangleStripDrawMode")}
	TrianglesDrawMode                       = &*TrianglesDrawModes{Value: get("TrianglesDrawMode")}
	UVMapping                               = &*Mapping{Value: get("UVMapping")}
	Uncharted2ToneMapping                   = &*ToneMapping{Value: get("Uncharted2ToneMapping")}
	UnsignedByteType                        = &*TextureDataType{Value: get("UnsignedByteType")}
	UnsignedInt248Type                      = &*PixelType{Value: get("UnsignedInt248Type")}
	UnsignedIntType                         = &*TextureDataType{Value: get("UnsignedIntType")}
	UnsignedShort4444Type                   = &*PixelType{Value: get("UnsignedShort4444Type")}
	UnsignedShort5551Type                   = &*PixelType{Value: get("UnsignedShort5551Type")}
	UnsignedShort565Type                    = &*PixelType{Value: get("UnsignedShort565Type")}
	UnsignedShortType                       = &*TextureDataType{Value: get("UnsignedShortType")}
	VertexColors                            = &*Colors{Value: get("VertexColors")}
	WrapAroundEnding                        = &*InterpolationEndingModes{Value: get("WrapAroundEnding")}
	ZeroCurvatureEnding                     = &*InterpolationEndingModes{Value: get("ZeroCurvatureEnding")}
	ZeroFactor                              = &*BlendingDstFactor{Value: get("ZeroFactor")}
	ZeroSlopeEnding                         = &*InterpolationEndingModes{Value: get("ZeroSlopeEnding")}
	SRGBEncoding                            = &*TextureEncoding{Value: get("sRGBEncoding")}
)

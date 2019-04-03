// Code generated from "constants.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

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
	LEFT   MOUSE = MOUSE(get("LEFT"))
	MIDDLE MOUSE = MOUSE(get("MIDDLE"))
	RIGHT  MOUSE = MOUSE(get("RIGHT"))
)

type Mapping js.Value
type NormalMapTypes js.Value
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
	AddEquation                      BlendingEquation          = BlendingEquation(get("AddEquation"))
	AddOperation                     Combine                   = Combine(get("AddOperation"))
	AdditiveBlending                 Blending                  = Blending(get("AdditiveBlending"))
	AlphaFormat                      PixelFormat               = PixelFormat(get("AlphaFormat"))
	AlwaysDepth                      DepthModes                = DepthModes(get("AlwaysDepth"))
	BackSide                         Side                      = Side(get("BackSide"))
	BasicDepthPacking                DepthPackingStrategies    = DepthPackingStrategies(get("BasicDepthPacking"))
	BasicShadowMap                   ShadowMapType             = ShadowMapType(get("BasicShadowMap"))
	ByteType                         TextureDataType           = TextureDataType(get("ByteType"))
	CineonToneMapping                ToneMapping               = ToneMapping(get("CineonToneMapping"))
	ClampToEdgeWrapping              Wrapping                  = Wrapping(get("ClampToEdgeWrapping"))
	CubeReflectionMapping            Mapping                   = Mapping(get("CubeReflectionMapping"))
	CubeRefractionMapping            Mapping                   = Mapping(get("CubeRefractionMapping"))
	CubeUVReflectionMapping          Mapping                   = Mapping(get("CubeUVReflectionMapping"))
	CubeUVRefractionMapping          Mapping                   = Mapping(get("CubeUVRefractionMapping"))
	CullFaceBack                     CullFace                  = CullFace(get("CullFaceBack"))
	CullFaceFront                    CullFace                  = CullFace(get("CullFaceFront"))
	CullFaceFrontBack                CullFace                  = CullFace(get("CullFaceFrontBack"))
	CullFaceNone                     CullFace                  = CullFace(get("CullFaceNone"))
	CustomBlending                   Blending                  = Blending(get("CustomBlending"))
	DepthFormat                      PixelFormat               = PixelFormat(get("DepthFormat"))
	DepthStencilFormat               PixelFormat               = PixelFormat(get("DepthStencilFormat"))
	DoubleSide                       Side                      = Side(get("DoubleSide"))
	DstAlphaFactor                   BlendingDstFactor         = BlendingDstFactor(get("DstAlphaFactor"))
	DstColorFactor                   BlendingDstFactor         = BlendingDstFactor(get("DstColorFactor"))
	EqualDepth                       DepthModes                = DepthModes(get("EqualDepth"))
	EquirectangularReflectionMapping Mapping                   = Mapping(get("EquirectangularReflectionMapping"))
	EquirectangularRefractionMapping Mapping                   = Mapping(get("EquirectangularRefractionMapping"))
	FaceColors                       Colors                    = Colors(get("FaceColors"))
	FlatShading                      Shading                   = Shading(get("FlatShading"))
	FloatType                        TextureDataType           = TextureDataType(get("FloatType"))
	FrontFaceDirectionCCW            FrontFaceDirection        = FrontFaceDirection(get("FrontFaceDirectionCCW"))
	FrontFaceDirectionCW             FrontFaceDirection        = FrontFaceDirection(get("FrontFaceDirectionCW"))
	FrontSide                        Side                      = Side(get("FrontSide"))
	GammaEncoding                    TextureEncoding           = TextureEncoding(get("GammaEncoding"))
	GreaterDepth                     DepthModes                = DepthModes(get("GreaterDepth"))
	GreaterEqualDepth                DepthModes                = DepthModes(get("GreaterEqualDepth"))
	HalfFloatType                    TextureDataType           = TextureDataType(get("HalfFloatType"))
	IntType                          TextureDataType           = TextureDataType(get("IntType"))
	InterpolateDiscrete              InterpolationModes        = InterpolationModes(get("InterpolateDiscrete"))
	InterpolateLinear                InterpolationModes        = InterpolationModes(get("InterpolateLinear"))
	InterpolateSmooth                InterpolationModes        = InterpolationModes(get("InterpolateSmooth"))
	LessDepth                        DepthModes                = DepthModes(get("LessDepth"))
	LessEqualDepth                   DepthModes                = DepthModes(get("LessEqualDepth"))
	LinearEncoding                   TextureEncoding           = TextureEncoding(get("LinearEncoding"))
	LinearFilter                     TextureFilter             = TextureFilter(get("LinearFilter"))
	LinearMipMapLinearFilter         TextureFilter             = TextureFilter(get("LinearMipMapLinearFilter"))
	LinearMipMapNearestFilter        TextureFilter             = TextureFilter(get("LinearMipMapNearestFilter"))
	LinearToneMapping                ToneMapping               = ToneMapping(get("LinearToneMapping"))
	LogLuvEncoding                   TextureEncoding           = TextureEncoding(get("LogLuvEncoding"))
	LoopOnce                         AnimationActionLoopStyles = AnimationActionLoopStyles(get("LoopOnce"))
	LoopPingPong                     AnimationActionLoopStyles = AnimationActionLoopStyles(get("LoopPingPong"))
	LoopRepeat                       AnimationActionLoopStyles = AnimationActionLoopStyles(get("LoopRepeat"))
	LuminanceAlphaFormat             PixelFormat               = PixelFormat(get("LuminanceAlphaFormat"))
	LuminanceFormat                  PixelFormat               = PixelFormat(get("LuminanceFormat"))
	MaxEquation                      BlendingEquation          = BlendingEquation(get("MaxEquation"))
	MinEquation                      BlendingEquation          = BlendingEquation(get("MinEquation"))
	MirroredRepeatWrapping           Wrapping                  = Wrapping(get("MirroredRepeatWrapping"))
	MixOperation                     Combine                   = Combine(get("MixOperation"))
	MultiplyBlending                 Blending                  = Blending(get("MultiplyBlending"))
	MultiplyOperation                Combine                   = Combine(get("MultiplyOperation"))
	NearestFilter                    TextureFilter             = TextureFilter(get("NearestFilter"))
	NearestMipMapLinearFilter        TextureFilter             = TextureFilter(get("NearestMipMapLinearFilter"))
	NearestMipMapNearestFilter       TextureFilter             = TextureFilter(get("NearestMipMapNearestFilter"))
	NeverDepth                       DepthModes                = DepthModes(get("NeverDepth"))
	NoBlending                       Blending                  = Blending(get("NoBlending"))
	NoColors                         Colors                    = Colors(get("NoColors"))
	NoToneMapping                    ToneMapping               = ToneMapping(get("NoToneMapping"))
	NormalBlending                   Blending                  = Blending(get("NormalBlending"))
	NotEqualDepth                    DepthModes                = DepthModes(get("NotEqualDepth"))
	ObjectSpaceNormalMap             NormalMapTypes            = NormalMapTypes(get("ObjectSpaceNormalMap"))
	OneFactor                        BlendingDstFactor         = BlendingDstFactor(get("OneFactor"))
	OneMinusDstAlphaFactor           BlendingDstFactor         = BlendingDstFactor(get("OneMinusDstAlphaFactor"))
	OneMinusDstColorFactor           BlendingDstFactor         = BlendingDstFactor(get("OneMinusDstColorFactor"))
	OneMinusSrcAlphaFactor           BlendingDstFactor         = BlendingDstFactor(get("OneMinusSrcAlphaFactor"))
	OneMinusSrcColorFactor           BlendingDstFactor         = BlendingDstFactor(get("OneMinusSrcColorFactor"))
	PCFShadowMap                     ShadowMapType             = ShadowMapType(get("PCFShadowMap"))
	PCFSoftShadowMap                 ShadowMapType             = ShadowMapType(get("PCFSoftShadowMap"))
	REVISION                         string                    = get("REVISION").String()
	RGBADepthPacking                 DepthPackingStrategies    = DepthPackingStrategies(get("RGBADepthPacking"))
	RGBAFormat                       PixelFormat               = PixelFormat(get("RGBAFormat"))
	RGBA_ASTC_10x10_Format           CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_10x10_Format"))
	RGBA_ASTC_10x5_Format            CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_10x5_Format"))
	RGBA_ASTC_10x6_Format            CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_10x6_Format"))
	RGBA_ASTC_10x8_Format            CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_10x8_Format"))
	RGBA_ASTC_12x10_Format           CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_12x10_Format"))
	RGBA_ASTC_12x12_Format           CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_12x12_Format"))
	RGBA_ASTC_4x4_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_4x4_Format"))
	RGBA_ASTC_5x4_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_5x4_Format"))
	RGBA_ASTC_5x5_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_5x5_Format"))
	RGBA_ASTC_6x5_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_6x5_Format"))
	RGBA_ASTC_6x6_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_6x6_Format"))
	RGBA_ASTC_8x5_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_8x5_Format"))
	RGBA_ASTC_8x6_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_8x6_Format"))
	RGBA_ASTC_8x8_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_ASTC_8x8_Format"))
	RGBA_PVRTC_2BPPV1_Format         CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_PVRTC_2BPPV1_Format"))
	RGBA_PVRTC_4BPPV1_Format         CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_PVRTC_4BPPV1_Format"))
	RGBA_S3TC_DXT1_Format            CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_S3TC_DXT1_Format"))
	RGBA_S3TC_DXT3_Format            CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_S3TC_DXT3_Format"))
	RGBA_S3TC_DXT5_Format            CompressedPixelFormat     = CompressedPixelFormat(get("RGBA_S3TC_DXT5_Format"))
	RGBDEncoding                     TextureEncoding           = TextureEncoding(get("RGBDEncoding"))
	RGBEEncoding                     TextureEncoding           = TextureEncoding(get("RGBEEncoding"))
	RGBEFormat                       PixelFormat               = PixelFormat(get("RGBEFormat"))
	RGBFormat                        PixelFormat               = PixelFormat(get("RGBFormat"))
	RGBM16Encoding                   TextureEncoding           = TextureEncoding(get("RGBM16Encoding"))
	RGBM7Encoding                    TextureEncoding           = TextureEncoding(get("RGBM7Encoding"))
	RGB_ETC1_Format                  CompressedPixelFormat     = CompressedPixelFormat(get("RGB_ETC1_Format"))
	RGB_PVRTC_2BPPV1_Format          CompressedPixelFormat     = CompressedPixelFormat(get("RGB_PVRTC_2BPPV1_Format"))
	RGB_PVRTC_4BPPV1_Format          CompressedPixelFormat     = CompressedPixelFormat(get("RGB_PVRTC_4BPPV1_Format"))
	RGB_S3TC_DXT1_Format             CompressedPixelFormat     = CompressedPixelFormat(get("RGB_S3TC_DXT1_Format"))
	RedFormat                        PixelFormat               = PixelFormat(get("RedFormat"))
	ReinhardToneMapping              ToneMapping               = ToneMapping(get("ReinhardToneMapping"))
	RepeatWrapping                   Wrapping                  = Wrapping(get("RepeatWrapping"))
	ReverseSubtractEquation          BlendingEquation          = BlendingEquation(get("ReverseSubtractEquation"))
	ShortType                        TextureDataType           = TextureDataType(get("ShortType"))
	SmoothShading                    Shading                   = Shading(get("SmoothShading"))
	SphericalReflectionMapping       Mapping                   = Mapping(get("SphericalReflectionMapping"))
	SrcAlphaFactor                   BlendingDstFactor         = BlendingDstFactor(get("SrcAlphaFactor"))
	SrcAlphaSaturateFactor           js.Value                  = get("SrcAlphaSaturateFactor")
	SrcColorFactor                   BlendingDstFactor         = BlendingDstFactor(get("SrcColorFactor"))
	SubtractEquation                 BlendingEquation          = BlendingEquation(get("SubtractEquation"))
	SubtractiveBlending              Blending                  = Blending(get("SubtractiveBlending"))
	TangentSpaceNormalMap            NormalMapTypes            = NormalMapTypes(get("TangentSpaceNormalMap"))
	TriangleFanDrawMode              TrianglesDrawModes        = TrianglesDrawModes(get("TriangleFanDrawMode"))
	TriangleStripDrawMode            TrianglesDrawModes        = TrianglesDrawModes(get("TriangleStripDrawMode"))
	TrianglesDrawMode                TrianglesDrawModes        = TrianglesDrawModes(get("TrianglesDrawMode"))
	UVMapping                        Mapping                   = Mapping(get("UVMapping"))
	Uncharted2ToneMapping            ToneMapping               = ToneMapping(get("Uncharted2ToneMapping"))
	UnsignedByteType                 TextureDataType           = TextureDataType(get("UnsignedByteType"))
	UnsignedInt248Type               PixelType                 = PixelType(get("UnsignedInt248Type"))
	UnsignedIntType                  TextureDataType           = TextureDataType(get("UnsignedIntType"))
	UnsignedShort4444Type            PixelType                 = PixelType(get("UnsignedShort4444Type"))
	UnsignedShort5551Type            PixelType                 = PixelType(get("UnsignedShort5551Type"))
	UnsignedShort565Type             PixelType                 = PixelType(get("UnsignedShort565Type"))
	UnsignedShortType                TextureDataType           = TextureDataType(get("UnsignedShortType"))
	VertexColors                     Colors                    = Colors(get("VertexColors"))
	WrapAroundEnding                 InterpolationEndingModes  = InterpolationEndingModes(get("WrapAroundEnding"))
	ZeroCurvatureEnding              InterpolationEndingModes  = InterpolationEndingModes(get("ZeroCurvatureEnding"))
	ZeroFactor                       BlendingDstFactor         = BlendingDstFactor(get("ZeroFactor"))
	ZeroSlopeEnding                  InterpolationEndingModes  = InterpolationEndingModes(get("ZeroSlopeEnding"))
	SRGBEncoding                     TextureEncoding           = TextureEncoding(get("sRGBEncoding"))
)

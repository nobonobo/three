package main

import (
	"os"
	"path/filepath"
)

var (
	prefix       = os.Getenv("PWD") + "/three.js/src/"
	outputPrefix = "github.com/nobonobo/three"
	outputDir    = filepath.Clean(filepath.Join(os.Getenv("PWD"), "..", ".."))

	// BlackList ...
	BlackList = map[string]struct{}{
		"Array":                       struct{}{},
		"ArrayLike":                   struct{}{},
		"Iterable":                    struct{}{},
		"ArrayBuffer":                 struct{}{},
		"MimeType":                    struct{}{},
		"Uint8Array":                  struct{}{},
		"Float32Array":                struct{}{},
		"RegExp":                      struct{}{},
		"Object":                      struct{}{},
		"TypedArray":                  struct{}{},
		"MediaElementAudioSourceNode": struct{}{},
		"AudioContext":                struct{}{},
		"GainNode":                    struct{}{},
		"AudioBufferSourceNode":       struct{}{},
		"ArrayBufferView":             struct{}{},
		"ImageData":                   struct{}{},
		"PannerNode":                  struct{}{},
		"Curve":                       struct{}{},
		"HTMLImageElement":            struct{}{},
		"HTMLVideoElement":            struct{}{},
		"HTMLCanvasElement":           struct{}{},
		"Shader":                      struct{}{},
		"WebGLRenderingContext":       struct{}{},
		"VRDisplay":                   struct{}{},
		"UVGenerator":                 struct{}{},
		"BlendingSrcFactor":           struct{}{},
		"LoaderHandler":               struct{}{},
		"WebVRManager":                struct{}{},
		"IFog":                        struct{}{},
		"RenderTarget":                struct{}{},
	}
	// BlackListArrayType ...
	BlackListArrayType = map[string]struct{}{
		"null":                struct{}{},
		"[]Object3D":          struct{}{},
		"[]js.Value":          struct{}{},
		"[]KeyframeTrack":     struct{}{},
		"[]AnimationClip":     struct{}{},
		"[]int":               struct{}{},
		"[]Vector2":           struct{}{},
		"[][]Vector2":         struct{}{},
		"[][][]Vector2":       struct{}{},
		"[]PerspectiveCamera": struct{}{},
		"[]Bone":              struct{}{},
		"[]Color":             struct{}{},
		"[]MorphNormals":      struct{}{},
		"[]MorphTarget":       struct{}{},
		"[]Vector4":           struct{}{},
		"[]Vector3":           struct{}{},
		"[]Face3":             struct{}{},
		"[]string":            struct{}{},
		"[]Plane":             struct{}{},
		"[]Material":          struct{}{},
		"[]Intersection":      struct{}{},
		"[]Path":              struct{}{},
		"[]Matrix4":           struct{}{},
		"[]WebGLProgram":      struct{}{},
		"[]Shape":             struct{}{},
		"[][]int":             struct{}{},
		"[]float64":           struct{}{},
		"[][]float64":         struct{}{},
	}
	// ConstantFiles definition files
	ConstantFiles = map[string]struct{}{
		"constants.d.ts":                        struct{}{},
		"math/ColorKeywords/colorkeywords.d.ts": struct{}{},
		"audio/AudioContext.d.ts":               struct{}{},
	}
	// ExcludeFiles exclude files
	ExcludeFiles = map[string]struct{}{
		"extras/core/Curve.d.ts":             struct{}{}, // generics included
		"extras/core/CurvePath.d.ts":         struct{}{}, // generics included
		"loaders/ObjectLoader.d.ts":          struct{}{}, // generics included
		"renderers/shaders/ShaderChunk.d.ts": struct{}{},
		"renderers/shaders/ShaderLib.d.ts":   struct{}{},
	}
)

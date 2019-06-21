// Code generated from "renderers/webgl/WebGLUniforms.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLUniforms struct {
	js.Value
}

func NewWebGLUniforms(gl js.Value, program *WebGLProgram) *WebGLUniforms {
	return &WebGLUniforms{Value: get("WebGLUniforms").New(gl, program)}
}
func (wglu *WebGLUniforms) SetOptional(gl js.Value, object js.Value, name string) {
	wglu.Call("setOptional", gl, object, name)
}
func (wglu *WebGLUniforms) SetValue2(gl js.Value, name string, value js.Value, textures *WebGLTextures) {
	wglu.Call("setValue", gl, name, value, textures)
}
func (wglu *WebGLUniforms) EvalDynamic(seq js.Value, values js.Value, object js.Value, camera js.Value) js.Value {
	return wglu.Call("evalDynamic", seq, values, object, camera)
}
func (wglu *WebGLUniforms) SeqWithValue(seq js.Value, values js.Value) js.Value {
	return wglu.Call("seqWithValue", seq, values)
}
func (wglu *WebGLUniforms) SplitDynamic(seq js.Value, values js.Value) js.Value {
	return wglu.Call("splitDynamic", seq, values)
}
func (wglu *WebGLUniforms) Upload(gl js.Value, seq js.Value, values js.Value, textures *WebGLTextures) {
	wglu.Call("upload", gl, seq, values, textures)
}
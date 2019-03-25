package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/renderers"
)

type WebGLUniforms struct {
	js.Value
}

func NewWebGLUniforms(gl js.Value, program *WebGLProgram, renderer *renderers.WebGLRenderer) *WebGLUniforms {
	return &WebGLUniforms{Value: get("WebGLUniforms").New(gl, program, renderer)}
}
func (wglu *WebGLUniforms) Renderer() *renderers.WebGLRenderer {
	return &renderers.WebGLRenderer{Value: wglu.Get("renderer")}
}
func (wglu *WebGLUniforms) SetRenderer(v *renderers.WebGLRenderer) {
	wglu.Set("renderer", v)
}
func (wglu *WebGLUniforms) Set(gl js.Value, object js.Value, name string) {
	wglu.Call("set", gl, object, name)
}
func (wglu *WebGLUniforms) SetOptional(gl js.Value, object js.Value, name string) {
	wglu.Call("setOptional", gl, object, name)
}
func (wglu *WebGLUniforms) SetValue(gl js.Value, value js.Value, renderer js.Value) {
	wglu.Call("setValue", gl, value, renderer)
}
func (wglu *WebGLUniforms) EvalDynamic(seq js.Value, values []js.Value, object js.Value, camera js.Value) []js.Value {
	return []js.Value(wglu.Call("evalDynamic", seq, values, object, camera))
}
func (wglu *WebGLUniforms) SeqWithValue(seq js.Value, values []js.Value) []js.Value {
	return []js.Value(wglu.Call("seqWithValue", seq, values))
}
func (wglu *WebGLUniforms) SplitDynamic(seq js.Value, values []js.Value) []js.Value {
	return []js.Value(wglu.Call("splitDynamic", seq, values))
}
func (wglu *WebGLUniforms) Upload(gl js.Value, seq js.Value, values []js.Value, renderer js.Value) {
	wglu.Call("upload", gl, seq, values, renderer)
}

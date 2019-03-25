package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLProgram struct {
	js.Value
}

func NewWebGLProgram(renderer *renderers.WebGLRenderer, code string, material *materials.ShaderMaterial, parameters renderers.WebGLRendererParameters) *WebGLProgram {
	return &WebGLProgram{Value: get("WebGLProgram").New(renderer, code, material, parameters)}
}
func (wglp *WebGLProgram) Attributes() js.Value {
	return wglp.Get("attributes")
}
func (wglp *WebGLProgram) SetAttributes(v js.Value) {
	wglp.Set("attributes", v)
}
func (wglp *WebGLProgram) Code() string {
	return wglp.Get("code").String()
}
func (wglp *WebGLProgram) SetCode(v string) {
	wglp.Set("code", v)
}
func (wglp *WebGLProgram) FragmentShader() *WebGLShader {
	return &WebGLShader{Value: wglp.Get("fragmentShader")}
}
func (wglp *WebGLProgram) SetFragmentShader(v *WebGLShader) {
	wglp.Set("fragmentShader", v)
}
func (wglp *WebGLProgram) Id() int {
	return wglp.Get("id").Int()
}
func (wglp *WebGLProgram) SetId(v int) {
	wglp.Set("id", v)
}
func (wglp *WebGLProgram) Program() js.Value {
	return wglp.Get("program")
}
func (wglp *WebGLProgram) SetProgram(v js.Value) {
	wglp.Set("program", v)
}
func (wglp *WebGLProgram) Uniforms() js.Value {
	return wglp.Get("uniforms")
}
func (wglp *WebGLProgram) SetUniforms(v js.Value) {
	wglp.Set("uniforms", v)
}
func (wglp *WebGLProgram) UsedTimes() int {
	return wglp.Get("usedTimes").Int()
}
func (wglp *WebGLProgram) SetUsedTimes(v int) {
	wglp.Set("usedTimes", v)
}
func (wglp *WebGLProgram) VertexShader() *WebGLShader {
	return &WebGLShader{Value: wglp.Get("vertexShader")}
}
func (wglp *WebGLProgram) SetVertexShader(v *WebGLShader) {
	wglp.Set("vertexShader", v)
}
func (wglp *WebGLProgram) Destroy() {
	wglp.Call("destroy")
}
func (wglp *WebGLProgram) GetAttributes() js.Value {
	return wglp.Call("getAttributes")
}
func (wglp *WebGLProgram) GetUniforms() *WebGLUniforms {
	return &WebGLUniforms{Value: wglp.Call("getUniforms")}
}

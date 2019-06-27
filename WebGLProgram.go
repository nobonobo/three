// Code generated from "renderers/webgl/WebGLProgram.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLProgram extend: []
type WebGLProgram struct {
	js.Value
}

func NewWebGLProgram(renderer *WebGLRenderer, extensions *WebGLExtensions, code string, material *ShaderMaterial, shader *WebGLShader, parameters WebGLRendererParameters, capabilities *WebGLCapabilities, textures *WebGLTextures) *WebGLProgram {
	return &WebGLProgram{Value: get("WebGLProgram").New(renderer, extensions, code, material, shader, parameters, capabilities, textures)}
}
func (wglp *WebGLProgram) JSValue() js.Value {
	return wglp.Value
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
	wglp.Set("fragmentShader", v.Value)
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
func (wglp *WebGLProgram) UsedTimes() float64 {
	return wglp.Get("usedTimes").Float()
}
func (wglp *WebGLProgram) SetUsedTimes(v float64) {
	wglp.Set("usedTimes", v)
}
func (wglp *WebGLProgram) VertexShader() *WebGLShader {
	return &WebGLShader{Value: wglp.Get("vertexShader")}
}
func (wglp *WebGLProgram) SetVertexShader(v *WebGLShader) {
	wglp.Set("vertexShader", v.Value)
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

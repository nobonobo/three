// Code generated from "renderers/webgl/WebGLPrograms.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLPrograms extend: []
type WebGLPrograms struct {
	js.Value
}

func NewWebGLPrograms(renderer *WebGLRenderer, extensions *WebGLExtensions, capabilities *WebGLCapabilities, textures *WebGLTextures) *WebGLPrograms {
	return &WebGLPrograms{Value: get("WebGLPrograms").New(renderer.JSValue(), extensions.JSValue(), capabilities.JSValue(), textures.JSValue())}
}
func (wglp *WebGLPrograms) JSValue() js.Value {
	return wglp.Value
}
func (wglp *WebGLPrograms) Programs() js.Value {
	return wglp.Get("programs")
}
func (wglp *WebGLPrograms) SetPrograms(v js.Value) {
	wglp.Set("programs", v)
}
func (wglp *WebGLPrograms) AcquireProgram(material *ShaderMaterial, parameters js.Value, code string) *WebGLProgram {
	return &WebGLProgram{Value: wglp.Call("acquireProgram", material, parameters, code)}
}
func (wglp *WebGLPrograms) GetParameters(material *ShaderMaterial, lights js.Value, fog js.Value, nClipPlanes float64, object js.Value) js.Value {
	return wglp.Call("getParameters", material, lights, fog, nClipPlanes, object)
}
func (wglp *WebGLPrograms) GetProgramCode(material *ShaderMaterial, parameters js.Value) string {
	return wglp.Call("getProgramCode", material, parameters).String()
}
func (wglp *WebGLPrograms) ReleaseProgram(program *WebGLProgram) {
	wglp.Call("releaseProgram", program.JSValue())
}

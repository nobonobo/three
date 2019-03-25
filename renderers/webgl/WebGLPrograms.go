package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/materials"
)

type WebGLPrograms struct {
	js.Value
}

func NewWebGLPrograms(renderer *renderers.WebGLRenderer, capabilities js.Value) *WebGLPrograms {
	return &WebGLPrograms{Value: get("WebGLPrograms").New(renderer, capabilities)}
}
func (wglp *WebGLPrograms) Programs() []WebGLProgram {
	return []WebGLProgram(wglp.Get("programs"))
}
func (wglp *WebGLPrograms) SetPrograms(v []WebGLProgram) {
	wglp.Set("programs", v)
}
func (wglp *WebGLPrograms) AcquireProgram(material *materials.ShaderMaterial, parameters js.Value, code string) *WebGLProgram {
	return &WebGLProgram{Value: wglp.Call("acquireProgram", material, parameters, code)}
}
func (wglp *WebGLPrograms) GetParameters(material *materials.ShaderMaterial, lights js.Value, fog js.Value, nClipPlanes int, object js.Value) js.Value {
	return wglp.Call("getParameters", material, lights, fog, nClipPlanes, object)
}
func (wglp *WebGLPrograms) GetProgramCode(material *materials.ShaderMaterial, parameters js.Value) string {
	return wglp.Call("getProgramCode", material, parameters).String()
}
func (wglp *WebGLPrograms) ReleaseProgram(program *WebGLProgram) {
	wglp.Call("releaseProgram", program)
}

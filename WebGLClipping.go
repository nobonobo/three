// Code generated from "renderers/webgl/WebGLClipping.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLClipping extend: []
type WebGLClipping struct {
	js.Value
}

func (wglc *WebGLClipping) JSValue() js.Value {
	return wglc.Value
}
func (wglc *WebGLClipping) NumPlanes() float64 {
	return wglc.Get("numPlanes").Float()
}
func (wglc *WebGLClipping) SetNumPlanes(v float64) {
	wglc.Set("numPlanes", v)
}
func (wglc *WebGLClipping) Uniform() js.Value {
	return wglc.Get("uniform")
}
func (wglc *WebGLClipping) SetUniform(v js.Value) {
	wglc.Set("uniform", v)
}
func (wglc *WebGLClipping) BeginShadows() {
	wglc.Call("beginShadows")
}
func (wglc *WebGLClipping) EndShadows() {
	wglc.Call("endShadows")
}
func (wglc *WebGLClipping) Init(planes js.Value, enableLocalClipping bool, camera Camera) bool {
	return wglc.Call("init", planes, enableLocalClipping, camera.JSValue()).Bool()
}
func (wglc *WebGLClipping) SetState(planes js.Value, clipShadows bool, camera Camera, cache bool, fromCache bool) {
	wglc.Call("setState", planes, clipShadows, camera.JSValue(), cache, fromCache)
}

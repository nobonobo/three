// Code generated from "renderers/webgl/WebGLInfo.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type WebGLInfo struct {
	js.Value
}

func (wgli *WebGLInfo) AutoReset() bool {
	return wgli.Get("autoReset").Bool()
}
func (wgli *WebGLInfo) SetAutoReset(v bool) {
	wgli.Set("autoReset", v)
}
func (wgli *WebGLInfo) Memory() js.Value {
	return wgli.Get("memory")
}
func (wgli *WebGLInfo) SetMemory(v js.Value) {
	wgli.Set("memory", v)
}
func (wgli *WebGLInfo) Programs() js.Value {
	return wgli.Get("programs")
}
func (wgli *WebGLInfo) SetPrograms(v js.Value) {
	wgli.Set("programs", v)
}
func (wgli *WebGLInfo) Render() js.Value {
	return wgli.Get("render")
}
func (wgli *WebGLInfo) SetRender(v js.Value) {
	wgli.Set("render", v)
}
func (wgli *WebGLInfo) Reset() {
	wgli.Call("reset")
}

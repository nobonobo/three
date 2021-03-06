// Code generated from "renderers/webgl/WebGLProperties.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLProperties extend: []
type WebGLProperties struct {
	js.Value
}

func NewWebGLProperties() *WebGLProperties {
	return &WebGLProperties{Value: get("WebGLProperties").New()}
}
func (wglp *WebGLProperties) JSValue() js.Value {
	return wglp.Value
}
func (wglp *WebGLProperties) Clear() {
	wglp.Call("clear")
}
func (wglp *WebGLProperties) Delete(object js.Value) {
	wglp.Call("delete", object)
}
func (wglp *WebGLProperties) Get2(object js.Value) js.Value {
	return wglp.Call("get", object)
}

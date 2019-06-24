// Code generated from "renderers/webgl/WebGLExtensions.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type WebGLExtensions struct {
	js.Value
}

func NewWebGLExtensions(gl js.Value) *WebGLExtensions {
	return &WebGLExtensions{Value: get("WebGLExtensions").New(gl)}
}
func (wgle *WebGLExtensions) Get2(name string) js.Value {
	return wgle.Call("get", name)
}

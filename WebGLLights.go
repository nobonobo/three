// Code generated from "renderers/webgl/WebGLLights.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type WebGLLights struct {
	js.Value
}

func NewWebGLLights(gl js.Value, properties js.Value, info js.Value) *WebGLLights {
	return &WebGLLights{Value: get("WebGLLights").New(gl, properties, info)}
}
func (wgll *WebGLLights) Get2(light js.Value) js.Value {
	return wgll.Call("get", light)
}

// Code generated from "renderers/webgl/WebGLShader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLShader extend: []
type WebGLShader struct {
	js.Value
}

func NewWebGLShader(gl js.Value, typ string, string string, debug bool) *WebGLShader {
	return &WebGLShader{Value: get("WebGLShader").New(gl, typ, string, debug)}
}
func (wgls *WebGLShader) JSValue() js.Value {
	return wgls.Value
}

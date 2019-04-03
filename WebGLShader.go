// Code generated from "renderers/webgl/WebGLShader.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLShader struct {
	js.Value
}

func NewWebGLShader(gl js.Value, typ string, string string, debug bool) *WebGLShader {
	return &WebGLShader{Value: get("WebGLShader").New(gl, typ, string, debug)}
}

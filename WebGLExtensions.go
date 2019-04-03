// Code generated from "renderers/webgl/WebGLExtensions.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
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

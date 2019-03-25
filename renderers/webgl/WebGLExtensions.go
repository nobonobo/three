package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLExtensions struct {
	js.Value
}

func NewWebGLExtensions(gl *WebGLRenderingContext) *WebGLExtensions {
	return &WebGLExtensions{Value: get("WebGLExtensions").New(gl)}
}
func (wgle *WebGLExtensions) Get(name string) js.Value {
	return wgle.Call("get", name)
}

package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLLights struct {
	js.Value
}

func NewWebGLLights(gl *WebGLRenderingContext, properties js.Value, info js.Value) *WebGLLights {
	return &WebGLLights{Value: get("WebGLLights").New(gl, properties, info)}
}
func (wgll *WebGLLights) Get(light js.Value) js.Value {
	return wgll.Call("get", light)
}

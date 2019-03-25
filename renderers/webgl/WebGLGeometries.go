package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLGeometries struct {
	js.Value
}

func NewWebGLGeometries(gl *WebGLRenderingContext, extensions js.Value, _infoRender js.Value) *WebGLGeometries {
	return &WebGLGeometries{Value: get("WebGLGeometries").New(gl, extensions, _infoRender)}
}
func (wglg *WebGLGeometries) Get(object js.Value) js.Value {
	return wglg.Call("get", object)
}

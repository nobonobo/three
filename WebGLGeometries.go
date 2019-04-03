// Code generated from "renderers/webgl/WebGLGeometries.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLGeometries struct {
	js.Value
}

func NewWebGLGeometries(gl js.Value, extensions js.Value, _infoRender js.Value) *WebGLGeometries {
	return &WebGLGeometries{Value: get("WebGLGeometries").New(gl, extensions, _infoRender)}
}
func (wglg *WebGLGeometries) Get2(object js.Value) js.Value {
	return wglg.Call("get", object)
}

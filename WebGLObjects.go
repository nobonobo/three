// Code generated from "renderers/webgl/WebGLObjects.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLObjects struct {
	js.Value
}

func NewWebGLObjects(gl js.Value, properties js.Value, info js.Value) *WebGLObjects {
	return &WebGLObjects{Value: get("WebGLObjects").New(gl, properties, info)}
}
func (wglo *WebGLObjects) GetAttributeBuffer(attribute js.Value) js.Value {
	return wglo.Call("getAttributeBuffer", attribute)
}
func (wglo *WebGLObjects) GetWireframeAttribute(geometry js.Value) js.Value {
	return wglo.Call("getWireframeAttribute", geometry)
}
func (wglo *WebGLObjects) Update(object js.Value) {
	wglo.Call("update", object)
}
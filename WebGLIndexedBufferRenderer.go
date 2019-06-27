// Code generated from "renderers/webgl/WebGLIndexedBufferRenderer.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLIndexedBufferRenderer extend: []
type WebGLIndexedBufferRenderer struct {
	js.Value
}

func NewWebGLIndexedBufferRenderer(gl js.Value, properties js.Value, info js.Value) *WebGLIndexedBufferRenderer {
	return &WebGLIndexedBufferRenderer{Value: get("WebGLIndexedBufferRenderer").New(gl, properties, info)}
}
func (wglibr *WebGLIndexedBufferRenderer) JSValue() js.Value {
	return wglibr.Value
}
func (wglibr *WebGLIndexedBufferRenderer) Render(start js.Value, count int) {
	wglibr.Call("render", start, count)
}
func (wglibr *WebGLIndexedBufferRenderer) RenderInstances(geometry js.Value, start js.Value, count int) {
	wglibr.Call("renderInstances", geometry, start, count)
}
func (wglibr *WebGLIndexedBufferRenderer) SetIndex(index js.Value) {
	wglibr.Call("setIndex", index)
}
func (wglibr *WebGLIndexedBufferRenderer) SetMode(value js.Value) {
	wglibr.Call("setMode", value)
}

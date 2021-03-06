// Code generated from "renderers/webgl/WebGLBufferRenderer.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLBufferRenderer extend: []
type WebGLBufferRenderer struct {
	js.Value
}

func NewWebGLBufferRenderer(_gl js.Value, extensions js.Value, _infoRender js.Value) *WebGLBufferRenderer {
	return &WebGLBufferRenderer{Value: get("WebGLBufferRenderer").New(_gl, extensions, _infoRender)}
}
func (wglbr *WebGLBufferRenderer) JSValue() js.Value {
	return wglbr.Value
}
func (wglbr *WebGLBufferRenderer) Render(start js.Value, count int) {
	wglbr.Call("render", start, count)
}
func (wglbr *WebGLBufferRenderer) RenderInstances(geometry js.Value) {
	wglbr.Call("renderInstances", geometry)
}
func (wglbr *WebGLBufferRenderer) SetMode(value js.Value) {
	wglbr.Call("setMode", value)
}

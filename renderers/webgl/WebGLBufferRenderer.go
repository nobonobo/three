package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLBufferRenderer struct {
	js.Value
}

func NewWebGLBufferRenderer(_gl *WebGLRenderingContext, extensions js.Value, _infoRender js.Value) *WebGLBufferRenderer {
	return &WebGLBufferRenderer{Value: get("WebGLBufferRenderer").New(_gl, extensions, _infoRender)}
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

package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLIndexedBufferRenderer struct {
	js.Value
}

func NewWebGLIndexedBufferRenderer(gl *WebGLRenderingContext, properties js.Value, info js.Value) *WebGLIndexedBufferRenderer {
	return &WebGLIndexedBufferRenderer{Value: get("WebGLIndexedBufferRenderer").New(gl, properties, info)}
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

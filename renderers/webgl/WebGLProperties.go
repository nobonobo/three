package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLProperties struct {
	js.Value
}

func NewWebGLProperties() *WebGLProperties {
	return &WebGLProperties{Value: get("WebGLProperties").New()}
}
func (wglp *WebGLProperties) Clear() {
	wglp.Call("clear")
}
func (wglp *WebGLProperties) Delete(object js.Value) {
	wglp.Call("delete", object)
}
func (wglp *WebGLProperties) Get(object js.Value) js.Value {
	return wglp.Call("get", object)
}

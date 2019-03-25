package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
)

type WebGLClipping struct {
	js.Value
}

func (wglc *WebGLClipping) NumPlanes() int {
	return wglc.Get("numPlanes").Int()
}
func (wglc *WebGLClipping) SetNumPlanes(v int) {
	wglc.Set("numPlanes", v)
}
func (wglc *WebGLClipping) Uniform() js.Value {
	return wglc.Get("uniform")
}
func (wglc *WebGLClipping) SetUniform(v js.Value) {
	wglc.Set("uniform", v)
}
func (wglc *WebGLClipping) BeginShadows() {
	wglc.Call("beginShadows")
}
func (wglc *WebGLClipping) EndShadows() {
	wglc.Call("endShadows")
}
func (wglc *WebGLClipping) Init(planes []js.Value, enableLocalClipping bool, camera *cameras.Camera) bool {
	return wglc.Call("init", planes, enableLocalClipping, camera).Bool()
}
func (wglc *WebGLClipping) SetState(planes []js.Value, clipShadows bool, camera *cameras.Camera, cache bool, fromCache bool) {
	wglc.Call("setState", planes, clipShadows, camera, cache, fromCache)
}

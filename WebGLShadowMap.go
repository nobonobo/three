// Code generated from "renderers/webgl/WebGLShadowMap.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// WebGLShadowMap extend: []
type WebGLShadowMap struct {
	js.Value
}

func NewWebGLShadowMap(_renderer *WebGLRenderer, _lights js.Value, _objects js.Value, capabilities js.Value) *WebGLShadowMap {
	return &WebGLShadowMap{Value: get("WebGLShadowMap").New(_renderer.JSValue(), _lights, _objects, capabilities)}
}
func (wglsm *WebGLShadowMap) JSValue() js.Value {
	return wglsm.Value
}
func (wglsm *WebGLShadowMap) AutoUpdate() bool {
	return wglsm.Get("autoUpdate").Bool()
}
func (wglsm *WebGLShadowMap) SetAutoUpdate(v bool) {
	wglsm.Set("autoUpdate", v)
}
func (wglsm *WebGLShadowMap) CullFace() js.Value {
	return wglsm.Get("cullFace")
}
func (wglsm *WebGLShadowMap) SetCullFace(v js.Value) {
	wglsm.Set("cullFace", v)
}
func (wglsm *WebGLShadowMap) Enabled() bool {
	return wglsm.Get("enabled").Bool()
}
func (wglsm *WebGLShadowMap) SetEnabled(v bool) {
	wglsm.Set("enabled", v)
}
func (wglsm *WebGLShadowMap) NeedsUpdate() bool {
	return wglsm.Get("needsUpdate").Bool()
}
func (wglsm *WebGLShadowMap) SetNeedsUpdate(v bool) {
	wglsm.Set("needsUpdate", v)
}
func (wglsm *WebGLShadowMap) Type() ShadowMapType {
	return ShadowMapType(wglsm.Get("type"))
}
func (wglsm *WebGLShadowMap) SetType(v ShadowMapType) {
	wglsm.Set("type", v)
}
func (wglsm *WebGLShadowMap) Render(scene Scene, camera Camera) {
	wglsm.Call("render", scene.JSValue(), camera.JSValue())
}

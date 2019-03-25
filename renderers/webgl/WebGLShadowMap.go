package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/scenes"
)

type WebGLShadowMap struct {
	js.Value
}

func NewWebGLShadowMap(_renderer *renderers.WebGLRenderer, _lights []js.Value, _objects []js.Value, capabilities js.Value) *WebGLShadowMap {
	return &WebGLShadowMap{Value: get("WebGLShadowMap").New(_renderer, _lights, _objects, capabilities)}
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
func (wglsm *WebGLShadowMap) Type() *ShadowMapType {
	return &ShadowMapType{Value: wglsm.Get("type")}
}
func (wglsm *WebGLShadowMap) SetType(v *ShadowMapType) {
	wglsm.Set("type", v)
}
func (wglsm *WebGLShadowMap) Render(scene *scenes.Scene, camera *cameras.Camera) {
	wglsm.Call("render", scene, camera)
}

// Code generated from "renderers/webgl/WebGLRenderLists.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type RenderItem interface {
}
type RenderTarget interface {
}
type WebGLRenderList struct {
	js.Value
}

func (wglrl *WebGLRenderList) Opaque() js.Value {
	return wglrl.Get("opaque")
}
func (wglrl *WebGLRenderList) SetOpaque(v js.Value) {
	wglrl.Set("opaque", v)
}
func (wglrl *WebGLRenderList) Transparent() js.Value {
	return wglrl.Get("transparent")
}
func (wglrl *WebGLRenderList) SetTransparent(v js.Value) {
	wglrl.Set("transparent", v)
}
func (wglrl *WebGLRenderList) Init() {
	wglrl.Call("init")
}
func (wglrl *WebGLRenderList) Push(object *Object3D, geometry *BufferGeometry, material *Material, groupOrder int, z float64, group *Group) {
	wglrl.Call("push", object, geometry, material, groupOrder, z, group)
}
func (wglrl *WebGLRenderList) Sort() {
	wglrl.Call("sort")
}
func (wglrl *WebGLRenderList) Unshift(object *Object3D, geometry *BufferGeometry, material *Material, groupOrder int, z int, group *Group) {
	wglrl.Call("unshift", object, geometry, material, groupOrder, z, group)
}

type WebGLRenderLists struct {
	js.Value
}

func (wglrl *WebGLRenderLists) Dispose() {
	wglrl.Call("dispose")
}
func (wglrl *WebGLRenderLists) Get2(scene *Scene, camera *Camera) *WebGLRenderList {
	return &WebGLRenderList{Value: wglrl.Call("get", scene, camera)}
}
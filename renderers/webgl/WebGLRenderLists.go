package webgl

import (
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/scenes"
)

type RenderItem interface {
}
type RenderTarget interface {
}
type WebGLRenderList struct {
	js.Value
}

func (wglrl *WebGLRenderList) Opaque() *Array {
	return &Array{Value: wglrl.Get("opaque")}
}
func (wglrl *WebGLRenderList) SetOpaque(v *Array) {
	wglrl.Set("opaque", v)
}
func (wglrl *WebGLRenderList) Transparent() *Array {
	return &Array{Value: wglrl.Get("transparent")}
}
func (wglrl *WebGLRenderList) SetTransparent(v *Array) {
	wglrl.Set("transparent", v)
}
func (wglrl *WebGLRenderList) Init() {
	wglrl.Call("init")
}
func (wglrl *WebGLRenderList) Push(object *core.Object3D, geometry core.Geometry, material *materials.Material, z int, group *objects.Group) {
	wglrl.Call("push", object, geometry, material, z, group)
}
func (wglrl *WebGLRenderList) Sort() {
	wglrl.Call("sort")
}

type WebGLRenderLists struct {
	js.Value
}

func (wglrl *WebGLRenderLists) Dispose() {
	wglrl.Call("dispose")
}
func (wglrl *WebGLRenderLists) Get(scene *scenes.Scene, camera *cameras.Camera) *WebGLRenderList {
	return &WebGLRenderList{Value: wglrl.Call("get", scene, camera)}
}

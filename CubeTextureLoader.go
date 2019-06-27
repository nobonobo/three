// Code generated from "loaders/CubeTextureLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CubeTextureLoader extend: []
type CubeTextureLoader struct {
	js.Value
}

func NewCubeTextureLoader(manager *LoadingManager) *CubeTextureLoader {
	return &CubeTextureLoader{Value: get("CubeTextureLoader").New(manager)}
}
func (ctl *CubeTextureLoader) JSValue() js.Value {
	return ctl.Value
}
func (ctl *CubeTextureLoader) CrossOrigin() string {
	return ctl.Get("crossOrigin").String()
}
func (ctl *CubeTextureLoader) SetCrossOrigin(v string) {
	ctl.Set("crossOrigin", v)
}
func (ctl *CubeTextureLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: ctl.Get("manager")}
}
func (ctl *CubeTextureLoader) SetManager(v *LoadingManager) {
	ctl.Set("manager", v.Value)
}
func (ctl *CubeTextureLoader) Path() string {
	return ctl.Get("path").String()
}
func (ctl *CubeTextureLoader) SetPath(v string) {
	ctl.Set("path", v)
}
func (ctl *CubeTextureLoader) Load(urls js.Value, onLoad js.Value, onProgress js.Value, onError js.Value) *CubeTexture {
	return &CubeTexture{Value: ctl.Call("load", urls, onLoad, onProgress, onError)}
}
func (ctl *CubeTextureLoader) SetCrossOrigin2(crossOrigin string) *CubeTextureLoader {
	return &CubeTextureLoader{Value: ctl.Call("setCrossOrigin", crossOrigin)}
}
func (ctl *CubeTextureLoader) SetPath2(path string) *CubeTextureLoader {
	return &CubeTextureLoader{Value: ctl.Call("setPath", path)}
}

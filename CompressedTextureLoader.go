// Code generated from "loaders/CompressedTextureLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CompressedTextureLoader extend: []
type CompressedTextureLoader struct {
	js.Value
}

func NewCompressedTextureLoader(manager *LoadingManager) *CompressedTextureLoader {
	return &CompressedTextureLoader{Value: get("CompressedTextureLoader").New(manager.JSValue())}
}
func (ctl *CompressedTextureLoader) JSValue() js.Value {
	return ctl.Value
}
func (ctl *CompressedTextureLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: ctl.Get("manager")}
}
func (ctl *CompressedTextureLoader) SetManager(v *LoadingManager) {
	ctl.Set("manager", v.JSValue())
}
func (ctl *CompressedTextureLoader) Path() string {
	return ctl.Get("path").String()
}
func (ctl *CompressedTextureLoader) SetPath(v string) {
	ctl.Set("path", v)
}
func (ctl *CompressedTextureLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	ctl.Call("load", url, onLoad, onProgress, onError)
}
func (ctl *CompressedTextureLoader) SetPath2(path string) *CompressedTextureLoader {
	return &CompressedTextureLoader{Value: ctl.Call("setPath", path)}
}

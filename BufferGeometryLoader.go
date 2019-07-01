// Code generated from "loaders/BufferGeometryLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// BufferGeometryLoader extend: []
type BufferGeometryLoader struct {
	js.Value
}

func NewBufferGeometryLoader(manager *LoadingManager) *BufferGeometryLoader {
	return &BufferGeometryLoader{Value: get("BufferGeometryLoader").New(manager.JSValue())}
}
func (bgl *BufferGeometryLoader) JSValue() js.Value {
	return bgl.Value
}
func (bgl *BufferGeometryLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: bgl.Get("manager")}
}
func (bgl *BufferGeometryLoader) SetManager(v *LoadingManager) {
	bgl.Set("manager", v.JSValue())
}
func (bgl *BufferGeometryLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	bgl.Call("load", url, onLoad, onProgress, onError)
}
func (bgl *BufferGeometryLoader) Parse(json js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: bgl.Call("parse", json)}
}

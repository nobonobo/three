// Code generated from "loaders/BufferGeometryLoader.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type BufferGeometryLoader struct {
	js.Value
}

func NewBufferGeometryLoader(manager *LoadingManager) *BufferGeometryLoader {
	return &BufferGeometryLoader{Value: get("BufferGeometryLoader").New(manager)}
}
func (bgl *BufferGeometryLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: bgl.Get("manager")}
}
func (bgl *BufferGeometryLoader) SetManager(v *LoadingManager) {
	bgl.Set("manager", v)
}
func (bgl *BufferGeometryLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	bgl.Call("load", url, onLoad, onProgress, onError)
}
func (bgl *BufferGeometryLoader) Parse(json js.Value) *BufferGeometry {
	return &BufferGeometry{Value: bgl.Call("parse", json)}
}
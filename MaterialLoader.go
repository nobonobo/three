// Code generated from "loaders/MaterialLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// MaterialLoader extend: []
type MaterialLoader struct {
	js.Value
}

func NewMaterialLoader(manager *LoadingManager) *MaterialLoader {
	return &MaterialLoader{Value: get("MaterialLoader").New(manager)}
}
func (ml *MaterialLoader) JSValue() js.Value {
	return ml.Value
}
func (ml *MaterialLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: ml.Get("manager")}
}
func (ml *MaterialLoader) SetManager(v *LoadingManager) {
	ml.Set("manager", v.Value)
}
func (ml *MaterialLoader) Textures() js.Value {
	return ml.Get("textures")
}
func (ml *MaterialLoader) SetTextures(v js.Value) {
	ml.Set("textures", v)
}
func (ml *MaterialLoader) GetTexture(name string) *Texture {
	return &Texture{Value: ml.Call("getTexture", name)}
}
func (ml *MaterialLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	ml.Call("load", url, onLoad, onProgress, onError)
}
func (ml *MaterialLoader) Parse(json js.Value) Material {
	return &MaterialImpl{Value: ml.Call("parse", json)}
}
func (ml *MaterialLoader) SetTextures2(textures js.Value) {
	ml.Call("setTextures", textures)
}

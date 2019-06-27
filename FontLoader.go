// Code generated from "loaders/FontLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// FontLoader extend: []
type FontLoader struct {
	js.Value
}

func NewFontLoader(manager *LoadingManager) *FontLoader {
	return &FontLoader{Value: get("FontLoader").New(manager)}
}
func (fl *FontLoader) JSValue() js.Value {
	return fl.Value
}
func (fl *FontLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: fl.Get("manager")}
}
func (fl *FontLoader) SetManager(v *LoadingManager) {
	fl.Set("manager", v.Value)
}
func (fl *FontLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	fl.Call("load", url, onLoad, onProgress, onError)
}
func (fl *FontLoader) Parse(json js.Value) *Font {
	return &Font{Value: fl.Call("parse", json)}
}

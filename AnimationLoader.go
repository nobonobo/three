// Code generated from "loaders/AnimationLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// AnimationLoader extend: []
type AnimationLoader struct {
	js.Value
}

func NewAnimationLoader(manager *LoadingManager) *AnimationLoader {
	return &AnimationLoader{Value: get("AnimationLoader").New(manager.JSValue())}
}
func (al *AnimationLoader) JSValue() js.Value {
	return al.Value
}
func (al *AnimationLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: al.Get("manager")}
}
func (al *AnimationLoader) SetManager(v *LoadingManager) {
	al.Set("manager", v.JSValue())
}
func (al *AnimationLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value {
	return al.Call("load", url, onLoad, onProgress, onError)
}
func (al *AnimationLoader) Parse(json js.Value) js.Value {
	return al.Call("parse", json)
}
func (al *AnimationLoader) SetPath(path string) *AnimationLoader {
	return &AnimationLoader{Value: al.Call("setPath", path)}
}

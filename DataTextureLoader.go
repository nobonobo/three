// Code generated from "loaders/DataTextureLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// DataTextureLoader extend: []
type DataTextureLoader struct {
	js.Value
}

func NewDataTextureLoader(manager *LoadingManager) *DataTextureLoader {
	return &DataTextureLoader{Value: get("DataTextureLoader").New(manager.JSValue())}
}
func (dtl *DataTextureLoader) JSValue() js.Value {
	return dtl.Value
}
func (dtl *DataTextureLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: dtl.Get("manager")}
}
func (dtl *DataTextureLoader) SetManager(v *LoadingManager) {
	dtl.Set("manager", v.JSValue())
}
func (dtl *DataTextureLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	dtl.Call("load", url, onLoad, onProgress, onError)
}

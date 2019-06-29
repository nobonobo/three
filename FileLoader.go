// Code generated from "loaders/FileLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type AnyLoader interface {
	Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value
}
type LoaderHandler interface {
	Add(regex js.Value, loader AnyLoader)
	Get(file string) *AnyLoader
}

// FileLoader extend: []
type FileLoader struct {
	js.Value
}

func NewFileLoader(manager *LoadingManager) *FileLoader {
	return &FileLoader{Value: get("FileLoader").New(manager.JSValue())}
}
func (fl *FileLoader) JSValue() js.Value {
	return fl.Value
}
func (fl *FileLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: fl.Get("manager")}
}
func (fl *FileLoader) SetManager(v *LoadingManager) {
	fl.Set("manager", v.JSValue())
}
func (fl *FileLoader) MimeType() js.Value {
	return fl.Get("mimeType")
}
func (fl *FileLoader) SetMimeType(v js.Value) {
	fl.Set("mimeType", v)
}
func (fl *FileLoader) Path() string {
	return fl.Get("path").String()
}
func (fl *FileLoader) SetPath(v string) {
	fl.Set("path", v)
}
func (fl *FileLoader) ResponseType() string {
	return fl.Get("responseType").String()
}
func (fl *FileLoader) SetResponseType(v string) {
	fl.Set("responseType", v)
}
func (fl *FileLoader) WithCredentials() string {
	return fl.Get("withCredentials").String()
}
func (fl *FileLoader) SetWithCredentials(v string) {
	fl.Set("withCredentials", v)
}
func (fl *FileLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value {
	return fl.Call("load", url, onLoad, onProgress, onError)
}
func (fl *FileLoader) SetMimeType2(mimeType js.Value) *FileLoader {
	return &FileLoader{Value: fl.Call("setMimeType", mimeType)}
}
func (fl *FileLoader) SetPath2(path string) *FileLoader {
	return &FileLoader{Value: fl.Call("setPath", path)}
}
func (fl *FileLoader) SetRequestHeader(value js.Value) *FileLoader {
	return &FileLoader{Value: fl.Call("setRequestHeader", value)}
}
func (fl *FileLoader) SetResponseType2(responseType string) *FileLoader {
	return &FileLoader{Value: fl.Call("setResponseType", responseType)}
}
func (fl *FileLoader) SetWithCredentials2(value string) *FileLoader {
	return &FileLoader{Value: fl.Call("setWithCredentials", value)}
}

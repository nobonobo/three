// Code generated from "loaders/LoadingManager.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func DefaultLoadingManager() *LoadingManager {
	return &LoadingManager{Value: get("DefaultLoadingManager")}
}
func SetDefaultLoadingManager(v *LoadingManager) {
	set("DefaultLoadingManager", v)
}

// JSONLoader extend: [Loader]
type JSONLoader struct {
	js.Value
}

func NewJSONLoader(manager *LoadingManager) *JSONLoader {
	return &JSONLoader{Value: get("JSONLoader").New(manager)}
}
func (jsonl *JSONLoader) JSValue() js.Value {
	return jsonl.Value
}
func (jsonl *JSONLoader) CrossOrigin() string {
	return jsonl.Get("crossOrigin").String()
}
func (jsonl *JSONLoader) SetCrossOrigin(v string) {
	jsonl.Set("crossOrigin", v)
}
func (jsonl *JSONLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: jsonl.Get("manager")}
}
func (jsonl *JSONLoader) SetManager(v *LoadingManager) {
	jsonl.Set("manager", v.Value)
}
func (jsonl *JSONLoader) OnLoadComplete() js.Value {
	return jsonl.Get("onLoadComplete")
}
func (jsonl *JSONLoader) SetOnLoadComplete(v js.Value) {
	jsonl.Set("onLoadComplete", v)
}
func (jsonl *JSONLoader) OnLoadProgress() js.Value {
	return jsonl.Get("onLoadProgress")
}
func (jsonl *JSONLoader) SetOnLoadProgress(v js.Value) {
	jsonl.Set("onLoadProgress", v)
}
func (jsonl *JSONLoader) OnLoadStart() js.Value {
	return jsonl.Get("onLoadStart")
}
func (jsonl *JSONLoader) SetOnLoadStart(v js.Value) {
	jsonl.Set("onLoadStart", v)
}
func (jsonl *JSONLoader) WithCredentials() bool {
	return jsonl.Get("withCredentials").Bool()
}
func (jsonl *JSONLoader) SetWithCredentials(v bool) {
	jsonl.Set("withCredentials", v)
}
func (jsonl *JSONLoader) Handlers() js.Value {
	return jsonl.Get("Handlers")
}
func (jsonl *JSONLoader) SetHandlers(v js.Value) {
	jsonl.Set("Handlers", v)
}
func (jsonl *JSONLoader) CreateMaterial(m Material, texturePath string, crossOrigin string) bool {
	return jsonl.Call("createMaterial", m.JSValue(), texturePath, crossOrigin).Bool()
}
func (jsonl *JSONLoader) ExtractUrlBase(url string) string {
	return jsonl.Call("extractUrlBase", url).String()
}
func (jsonl *JSONLoader) InitMaterials(materials js.Value, texturePath string) js.Value {
	return jsonl.Call("initMaterials", materials, texturePath)
}
func (jsonl *JSONLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	jsonl.Call("load", url, onLoad, onProgress, onError)
}
func (jsonl *JSONLoader) Parse(json js.Value, texturePath string) js.Value {
	return jsonl.Call("parse", json, texturePath)
}
func (jsonl *JSONLoader) SetTexturePath(value string) {
	jsonl.Call("setTexturePath", value)
}

// LoadingManager extend: []
type LoadingManager struct {
	js.Value
}

func NewLoadingManager(onLoad js.Value, onProgress js.Value, onError js.Value) *LoadingManager {
	return &LoadingManager{Value: get("LoadingManager").New(onLoad, onProgress, onError)}
}
func (lm *LoadingManager) JSValue() js.Value {
	return lm.Value
}
func (lm *LoadingManager) OnError() js.Value {
	return lm.Get("onError")
}
func (lm *LoadingManager) SetOnError(v js.Value) {
	lm.Set("onError", v)
}
func (lm *LoadingManager) OnLoad() js.Value {
	return lm.Get("onLoad")
}
func (lm *LoadingManager) SetOnLoad(v js.Value) {
	lm.Set("onLoad", v)
}
func (lm *LoadingManager) OnProgress() js.Value {
	return lm.Get("onProgress")
}
func (lm *LoadingManager) SetOnProgress(v js.Value) {
	lm.Set("onProgress", v)
}
func (lm *LoadingManager) OnStart() js.Value {
	return lm.Get("onStart")
}
func (lm *LoadingManager) SetOnStart(v js.Value) {
	lm.Set("onStart", v)
}
func (lm *LoadingManager) ItemEnd(url string) {
	lm.Call("itemEnd", url)
}
func (lm *LoadingManager) ItemError(url string) {
	lm.Call("itemError", url)
}
func (lm *LoadingManager) ItemStart(url string) {
	lm.Call("itemStart", url)
}
func (lm *LoadingManager) ResolveURL(url string) string {
	return lm.Call("resolveURL", url).String()
}
func (lm *LoadingManager) SetURLModifier(callback js.Value) {
	lm.Call("setURLModifier", callback)
}

// Code generated from "loaders/Loader.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Loader struct {
	js.Value
}

func NewLoader() *Loader {
	return &Loader{Value: get("Loader").New()}
}
func (ll *Loader) CrossOrigin() string {
	return ll.Get("crossOrigin").String()
}
func (ll *Loader) SetCrossOrigin(v string) {
	ll.Set("crossOrigin", v)
}
func (ll *Loader) OnLoadComplete() js.Value {
	return ll.Get("onLoadComplete")
}
func (ll *Loader) SetOnLoadComplete(v js.Value) {
	ll.Set("onLoadComplete", v)
}
func (ll *Loader) OnLoadProgress() js.Value {
	return ll.Get("onLoadProgress")
}
func (ll *Loader) SetOnLoadProgress(v js.Value) {
	ll.Set("onLoadProgress", v)
}
func (ll *Loader) OnLoadStart() js.Value {
	return ll.Get("onLoadStart")
}
func (ll *Loader) SetOnLoadStart(v js.Value) {
	ll.Set("onLoadStart", v)
}
func (ll *Loader) Handlers() js.Value {
	return ll.Get("Handlers")
}
func (ll *Loader) SetHandlers(v js.Value) {
	ll.Set("Handlers", v)
}
func (ll *Loader) CreateMaterial(m *Material, texturePath string, crossOrigin string) bool {
	return ll.Call("createMaterial", m, texturePath, crossOrigin).Bool()
}
func (ll *Loader) ExtractUrlBase(url string) string {
	return ll.Call("extractUrlBase", url).String()
}
func (ll *Loader) InitMaterials(materials js.Value, texturePath string) js.Value {
	return ll.Call("initMaterials", materials, texturePath)
}

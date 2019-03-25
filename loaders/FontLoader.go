package loaders

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/extras/core"
)

type FontLoader struct {
	js.Value
}

func NewFontLoader(manager *LoadingManager) *FontLoader {
	return &FontLoader{Value: get("FontLoader").New(manager)}
}
func (fl *FontLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: fl.Get("manager")}
}
func (fl *FontLoader) SetManager(v *LoadingManager) {
	fl.Set("manager", v)
}
func (fl *FontLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	fl.Call("load", url, onLoad, onProgress, onError)
}
func (fl *FontLoader) Parse(json js.Value) *core.Font {
	return &core.Font{Value: fl.Call("parse", json)}
}

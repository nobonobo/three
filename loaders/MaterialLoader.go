package loaders

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/textures"
)

type MaterialLoader struct {
	js.Value
}

func NewMaterialLoader(manager *LoadingManager) *MaterialLoader {
	return &MaterialLoader{Value: get("MaterialLoader").New(manager)}
}
func (ml *MaterialLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: ml.Get("manager")}
}
func (ml *MaterialLoader) SetManager(v *LoadingManager) {
	ml.Set("manager", v)
}
func (ml *MaterialLoader) Textures() js.Value {
	return ml.Get("textures")
}
func (ml *MaterialLoader) SetTextures(v js.Value) {
	ml.Set("textures", v)
}
func (ml *MaterialLoader) GetTexture(name string) *textures.Texture {
	return &textures.Texture{Value: ml.Call("getTexture", name)}
}
func (ml *MaterialLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	ml.Call("load", url, onLoad, onProgress, onError)
}
func (ml *MaterialLoader) Parse(json js.Value) *materials.Material {
	return &materials.Material{Value: ml.Call("parse", json)}
}
func (ml *MaterialLoader) SetTextures(textures js.Value) {
	ml.Call("setTextures", textures)
}

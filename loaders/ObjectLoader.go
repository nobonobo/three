package loaders

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/loaders"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/textures"
)

type ObjectLoader struct {
	js.Value
}

func NewObjectLoader(manager *loaders.LoadingManager) *loaders.ObjectLoader {
	return &loaders.ObjectLoader{Value: get("ObjectLoader").New(manager)}
}
func (ol *ObjectLoader) CrossOrigin() string {
	return ol.Get("crossOrigin").String()
}
func (ol *ObjectLoader) SetCrossOrigin(v string) {
	ol.Set("crossOrigin", v)
}
func (ol *ObjectLoader) Manager() *loaders.LoadingManager {
	return &loaders.LoadingManager{Value: ol.Get("manager")}
}
func (ol *ObjectLoader) SetManager(v *loaders.LoadingManager) {
	ol.Set("manager", v)
}
func (ol *ObjectLoader) TexturePass() string {
	return ol.Get("texturePass").String()
}
func (ol *ObjectLoader) SetTexturePass(v string) {
	ol.Set("texturePass", v)
}
func (ol *ObjectLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	ol.Call("load", url, onLoad, onProgress, onError)
}
func (ol *ObjectLoader) Parse(json js.Value, onLoad js.Value) {
	ol.Call("parse", json, onLoad)
}
func (ol *ObjectLoader) ParseAnimations(json js.Value) []animation.AnimationClip {
	return &[]animation.AnimationClip{Value: ol.Call("parseAnimations", json)}
}
func (ol *ObjectLoader) ParseGeometries(json js.Value) []js.Value {
	return &[]js.Value{Value: ol.Call("parseGeometries", json)}
}
func (ol *ObjectLoader) ParseImages(json js.Value, onLoad js.Value) js.Value {
	return ol.Call("parseImages", json, onLoad)
}
func (ol *ObjectLoader) ParseMaterials(json js.Value, textures []textures.Texture) []materials.Material {
	return &[]materials.Material{Value: ol.Call("parseMaterials", json, textures)}
}
func (ol *ObjectLoader) ParseObject(data js.Value, geometries []js.Value, materials []materials.Material) {
	ol.Call("parseObject", data, geometries, materials)
}
func (ol *ObjectLoader) ParseTextures(json js.Value, images js.Value) []textures.Texture {
	return &[]textures.Texture{Value: ol.Call("parseTextures", json, images)}
}
func (ol *ObjectLoader) SetCrossOrigin(crossOrigin string) {
	ol.Call("setCrossOrigin", crossOrigin)
}
func (ol *ObjectLoader) SetTexturePath(value string) {
	ol.Call("setTexturePath", value)
}

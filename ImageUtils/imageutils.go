// Code generated from "ImageUtils/imageutils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package ImageUtils

import (
	"github.com/nobonobo/three"
	"syscall/js"
)

func CrossOrigin() string {
	return get("crossOrigin").String()
}
func SetCrossOrigin(v string) {
	set("crossOrigin", v)
}
func LoadTexture(url string, mapping three.Mapping, onLoad js.Value, onError js.Value) *three.Texture {
	return &three.Texture{Value: global().Call("loadTexture", url, mapping, onLoad, onError)}
}
func LoadTextureCube(array js.Value, mapping three.Mapping, onLoad js.Value, onError js.Value) *three.Texture {
	return &three.Texture{Value: global().Call("loadTextureCube", array, mapping, onLoad, onError)}
}

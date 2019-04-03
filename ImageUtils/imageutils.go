// Code generated from "ImageUtils/imageutils.d.ts"; DO NOT EDIT.

package ImageUtils

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three"
)

func CrossOrigin() string {
	return get("crossOrigin").String()
}
func SetCrossOrigin(v string) {
	set("crossOrigin", v)
}
func LoadTexture(url string, mapping three.Mapping, onLoad js.Value, onError js.Value) *three.Texture {
	return &three.Texture{Value: _Global.Call("loadTexture", url, mapping, onLoad, onError)}
}
func LoadTextureCube(array js.Value, mapping three.Mapping, onLoad js.Value, onError js.Value) *three.Texture {
	return &three.Texture{Value: _Global.Call("loadTextureCube", array, mapping, onLoad, onError)}
}

package imageutils

import (
	"github.com/nobonobo/three/textures"
)

func CrossOrigin() string {
	return get("crossOrigin").String()
}
func SetCrossOrigin(v string) {
	set("crossOrigin", v)
}
func LoadTexture(url string, mapping *Mapping, onLoad js.Value, onError js.Value) *textures.Texture {
	return &textures.Texture{Value: _Global.Call("loadTexture", url, mapping, onLoad, onError)}
}
func LoadTextureCube(array []string, mapping *Mapping, onLoad js.Value, onError js.Value) *textures.Texture {
	return &textures.Texture{Value: _Global.Call("loadTextureCube", array, mapping, onLoad, onError)}
}

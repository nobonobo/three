package loaders

import (
	"github.com/nobonobo/three/textures"
)

type TextureLoader struct {
	js.Value
}

func NewTextureLoader(manager *LoadingManager) *TextureLoader {
	return &TextureLoader{Value: get("TextureLoader").New(manager)}
}
func (tl *TextureLoader) CrossOrigin() string {
	return tl.Get("crossOrigin").String()
}
func (tl *TextureLoader) SetCrossOrigin(v string) {
	tl.Set("crossOrigin", v)
}
func (tl *TextureLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: tl.Get("manager")}
}
func (tl *TextureLoader) SetManager(v *LoadingManager) {
	tl.Set("manager", v)
}
func (tl *TextureLoader) Path() string {
	return tl.Get("path").String()
}
func (tl *TextureLoader) SetPath(v string) {
	tl.Set("path", v)
}
func (tl *TextureLoader) WithCredentials() string {
	return tl.Get("withCredentials").String()
}
func (tl *TextureLoader) SetWithCredentials(v string) {
	tl.Set("withCredentials", v)
}
func (tl *TextureLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) *textures.Texture {
	return &textures.Texture{Value: tl.Call("load", url, onLoad, onProgress, onError)}
}
func (tl *TextureLoader) SetCrossOrigin(crossOrigin string) *TextureLoader {
	return &TextureLoader{Value: tl.Call("setCrossOrigin", crossOrigin)}
}
func (tl *TextureLoader) SetPath(path string) *TextureLoader {
	return &TextureLoader{Value: tl.Call("setPath", path)}
}
func (tl *TextureLoader) SetWithCredentials(value string) *TextureLoader {
	return &TextureLoader{Value: tl.Call("setWithCredentials", value)}
}

package loaders

import (
	"github.com/nobonobo/three/textures"
)

type CubeTextureLoader struct {
	js.Value
}

func NewCubeTextureLoader(manager *LoadingManager) *CubeTextureLoader {
	return &CubeTextureLoader{Value: get("CubeTextureLoader").New(manager)}
}
func (ctl *CubeTextureLoader) CrossOrigin() string {
	return ctl.Get("crossOrigin").String()
}
func (ctl *CubeTextureLoader) SetCrossOrigin(v string) {
	ctl.Set("crossOrigin", v)
}
func (ctl *CubeTextureLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: ctl.Get("manager")}
}
func (ctl *CubeTextureLoader) SetManager(v *LoadingManager) {
	ctl.Set("manager", v)
}
func (ctl *CubeTextureLoader) Path() string {
	return ctl.Get("path").String()
}
func (ctl *CubeTextureLoader) SetPath(v string) {
	ctl.Set("path", v)
}
func (ctl *CubeTextureLoader) Load(urls *Array, onLoad js.Value, onProgress js.Value, onError js.Value) *textures.CubeTexture {
	return &textures.CubeTexture{Value: ctl.Call("load", urls, onLoad, onProgress, onError)}
}
func (ctl *CubeTextureLoader) SetCrossOrigin(crossOrigin string) this {
	return this(ctl.Call("setCrossOrigin", crossOrigin))
}
func (ctl *CubeTextureLoader) SetPath(path string) this {
	return this(ctl.Call("setPath", path))
}

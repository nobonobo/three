package loaders

import (
	"github.com/nobonobo/three/textures"
)

type DataTextureLoader struct {
	js.Value
}

func NewDataTextureLoader(manager *LoadingManager) *DataTextureLoader {
	return &DataTextureLoader{Value: get("DataTextureLoader").New(manager)}
}
func (dtl *DataTextureLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: dtl.Get("manager")}
}
func (dtl *DataTextureLoader) SetManager(v *LoadingManager) {
	dtl.Set("manager", v)
}
func (dtl *DataTextureLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) {
	dtl.Call("load", url, onLoad, onProgress, onError)
}

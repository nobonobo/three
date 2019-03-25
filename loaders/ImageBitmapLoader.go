package loaders

import (
	"github.com/gopherjs/gopherwasm/js"
)

type ImageBitmapLoader struct {
	js.Value
}

func NewImageBitmapLoader(manager *LoadingManager) *ImageBitmapLoader {
	return &ImageBitmapLoader{Value: get("ImageBitmapLoader").New(manager)}
}
func (ibl *ImageBitmapLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: ibl.Get("manager")}
}
func (ibl *ImageBitmapLoader) SetManager(v *LoadingManager) {
	ibl.Set("manager", v)
}
func (ibl *ImageBitmapLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value {
	return ibl.Call("load", url, onLoad, onProgress, onError)
}
func (ibl *ImageBitmapLoader) SetCrossOrigin() *ImageBitmapLoader {
	return &ImageBitmapLoader{Value: ibl.Call("setCrossOrigin")}
}
func (ibl *ImageBitmapLoader) SetOptions(options js.Value) *ImageBitmapLoader {
	return &ImageBitmapLoader{Value: ibl.Call("setOptions", options)}
}
func (ibl *ImageBitmapLoader) SetPath(path string) *ImageBitmapLoader {
	return &ImageBitmapLoader{Value: ibl.Call("setPath", path)}
}

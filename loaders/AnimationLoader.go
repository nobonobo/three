package loaders

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
)

type AnimationLoader struct {
	js.Value
}

func NewAnimationLoader(manager *LoadingManager) *AnimationLoader {
	return &AnimationLoader{Value: get("AnimationLoader").New(manager)}
}
func (al *AnimationLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: al.Get("manager")}
}
func (al *AnimationLoader) SetManager(v *LoadingManager) {
	al.Set("manager", v)
}
func (al *AnimationLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value {
	return al.Call("load", url, onLoad, onProgress, onError)
}
func (al *AnimationLoader) Parse(json js.Value) []animation.AnimationClip {
	return []animation.AnimationClip(al.Call("parse", json))
}
func (al *AnimationLoader) SetPath(path string) *AnimationLoader {
	return &AnimationLoader{Value: al.Call("setPath", path)}
}

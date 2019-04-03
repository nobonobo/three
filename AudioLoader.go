// Code generated from "loaders/AudioLoader.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type AudioLoader struct {
	js.Value
}

func NewAudioLoader(manager *LoadingManager) *AudioLoader {
	return &AudioLoader{Value: get("AudioLoader").New(manager)}
}
func (al *AudioLoader) Load(url string, onLoad js.Value, onPrgress js.Value, onError js.Value) {
	al.Call("load", url, onLoad, onPrgress, onError)
}

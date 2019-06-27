// Code generated from "loaders/AudioLoader.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// AudioLoader extend: []
type AudioLoader struct {
	js.Value
}

func NewAudioLoader(manager *LoadingManager) *AudioLoader {
	return &AudioLoader{Value: get("AudioLoader").New(manager)}
}
func (al *AudioLoader) JSValue() js.Value {
	return al.Value
}
func (al *AudioLoader) Load(url string, onLoad js.Value, onPrgress js.Value, onError js.Value) {
	al.Call("load", url, onLoad, onPrgress, onError)
}

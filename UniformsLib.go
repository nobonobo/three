// Code generated from "renderers/shaders/UniformsLib.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func UniformsLib() js.Value {
	return get("UniformsLib")
}
func SetUniformsLib(v js.Value) {
	set("UniformsLib", v)
}

type IUniform interface {
}

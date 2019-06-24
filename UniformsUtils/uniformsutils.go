// Code generated from "UniformsUtils/uniformsutils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package UniformsUtils

import (
	"syscall/js"
)

func Clone(uniforms_src js.Value) js.Value {
	return global().Call("clone", uniforms_src)
}
func Merge(uniforms js.Value) js.Value {
	return global().Call("merge", uniforms)
}

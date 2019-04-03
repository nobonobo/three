// Code generated from "UniformsUtils/uniformsutils.d.ts"; DO NOT EDIT.

package UniformsUtils

import (
	"github.com/gopherjs/gopherwasm/js"
)

func Clone(uniforms_src js.Value) js.Value {
	return _Global.Call("clone", uniforms_src)
}
func Merge(uniforms js.Value) js.Value {
	return _Global.Call("merge", uniforms)
}

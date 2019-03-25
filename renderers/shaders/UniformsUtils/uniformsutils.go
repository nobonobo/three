package uniformsutils

import (
	"github.com/gopherjs/gopherwasm/js"
)

func Clone(uniforms_src js.Value) js.Value {
	return _Global.Call("clone", uniforms_src)
}
func Merge(uniforms []js.Value) js.Value {
	return _Global.Call("merge", uniforms)
}

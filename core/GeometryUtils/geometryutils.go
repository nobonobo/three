package geometryutils

import (
	"github.com/gopherjs/gopherwasm/js"
)

func Center(geometry js.Value) js.Value {
	return _Global.Call("center", geometry)
}
func Merge(geometry1 js.Value, geometry2 js.Value, materialIndexOffset js.Value) js.Value {
	return _Global.Call("merge", geometry1, geometry2, materialIndexOffset)
}

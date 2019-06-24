// Code generated from "GeometryUtils/geometryutils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package GeometryUtils

import (
	"syscall/js"
)

func Center(geometry js.Value) js.Value {
	return global().Call("center", geometry)
}
func Merge(geometry1 js.Value, geometry2 js.Value, materialIndexOffset js.Value) js.Value {
	return global().Call("merge", geometry1, geometry2, materialIndexOffset)
}

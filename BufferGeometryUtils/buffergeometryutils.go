// Code generated from "BufferGeometryUtils/buffergeometryutils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package BufferGeometryUtils

import (
	"github.com/nobonobo/three"
	"syscall/js"
)

func ComputeTangents(geometry three.BufferGeometry) js.Value {
	return global().Call("computeTangents", geometry)
}
func MergeBufferAttributes(attributes []three.BufferAttribute) *three.BufferAttribute {
	return &three.BufferAttribute{Value: global().Call("mergeBufferAttributes", attributes)}
}
func MergeBufferGeometries(geometries []three.BufferGeometry) three.BufferGeometry {
	return &three.BufferGeometryImpl{Value: global().Call("mergeBufferGeometries", geometries)}
}

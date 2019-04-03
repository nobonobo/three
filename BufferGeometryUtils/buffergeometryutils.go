// Code generated from "BufferGeometryUtils/buffergeometryutils.d.ts"; DO NOT EDIT.

package BufferGeometryUtils

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three"
)

func ComputeTangents(geometry *three.BufferGeometry) js.Value {
	return _Global.Call("computeTangents", geometry)
}
func MergeBufferAttributes(attributes []three.BufferAttribute) *three.BufferAttribute {
	return &three.BufferAttribute{Value: _Global.Call("mergeBufferAttributes", attributes)}
}
func MergeBufferGeometries(geometries []three.BufferGeometry) *three.BufferGeometry {
	return &three.BufferGeometry{Value: _Global.Call("mergeBufferGeometries", geometries)}
}

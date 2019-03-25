package buffergeometryutils

import (
	"github.com/nobonobo/three/core"
)

func ComputeTangents(geometry *core.BufferGeometry) null {
	return null(_Global.Call("computeTangents", geometry))
}
func MergeBufferAttributes(attributes []core.BufferAttribute) *core.BufferAttribute {
	return &core.BufferAttribute{Value: _Global.Call("mergeBufferAttributes", attributes)}
}
func MergeBufferGeometries(geometries []core.BufferGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: _Global.Call("mergeBufferGeometries", geometries)}
}

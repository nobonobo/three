// Code generated from "lights/SpotLightShadow.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// SpotLightShadow extend: [LightShadow]
type SpotLightShadow struct {
	js.Value
}

func NewSpotLightShadow(camera Camera) *SpotLightShadow {
	return &SpotLightShadow{Value: get("SpotLightShadow").New(camera.JSValue())}
}
func (sls *SpotLightShadow) JSValue() js.Value {
	return sls.Value
}
func (sls *SpotLightShadow) Bias() float64 {
	return sls.Get("bias").Float()
}
func (sls *SpotLightShadow) SetBias(v float64) {
	sls.Set("bias", v)
}
func (sls *SpotLightShadow) Camera() *PerspectiveCamera {
	return &PerspectiveCamera{Value: sls.Get("camera")}
}
func (sls *SpotLightShadow) SetCamera(v *PerspectiveCamera) {
	sls.Set("camera", v.Value)
}
func (sls *SpotLightShadow) Map() js.Value {
	return sls.Get("map")
}
func (sls *SpotLightShadow) SetMap(v js.Value) {
	sls.Set("map", v)
}
func (sls *SpotLightShadow) MapSize() *Vector2 {
	return &Vector2{Value: sls.Get("mapSize")}
}
func (sls *SpotLightShadow) SetMapSize(v *Vector2) {
	sls.Set("mapSize", v.Value)
}
func (sls *SpotLightShadow) Matrix() *Matrix4 {
	return &Matrix4{Value: sls.Get("matrix")}
}
func (sls *SpotLightShadow) SetMatrix(v *Matrix4) {
	sls.Set("matrix", v.Value)
}
func (sls *SpotLightShadow) Radius() float64 {
	return sls.Get("radius").Float()
}
func (sls *SpotLightShadow) SetRadius(v float64) {
	sls.Set("radius", v)
}
func (sls *SpotLightShadow) Clone(recursive bool) *SpotLightShadow {
	return &SpotLightShadow{Value: sls.Call("clone", recursive)}
}
func (sls *SpotLightShadow) Copy(source *LightShadow) *SpotLightShadow {
	return &SpotLightShadow{Value: sls.Call("copy", source)}
}
func (sls *SpotLightShadow) ToJSON() js.Value {
	return sls.Call("toJSON")
}
func (sls *SpotLightShadow) Update(light *Light) {
	sls.Call("update", light)
}

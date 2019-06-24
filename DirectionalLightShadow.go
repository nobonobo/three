// Code generated from "lights/DirectionalLightShadow.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type DirectionalLightShadow struct {
	js.Value
}

func NewDirectionalLightShadow(camera *Camera) *DirectionalLightShadow {
	return &DirectionalLightShadow{Value: get("DirectionalLightShadow").New(camera)}
}
func (dls *DirectionalLightShadow) Bias() float64 {
	return dls.Get("bias").Float()
}
func (dls *DirectionalLightShadow) SetBias(v float64) {
	dls.Set("bias", v)
}
func (dls *DirectionalLightShadow) Camera() *OrthographicCamera {
	return &OrthographicCamera{Value: dls.Get("camera")}
}
func (dls *DirectionalLightShadow) SetCamera(v *OrthographicCamera) {
	dls.Set("camera", v)
}
func (dls *DirectionalLightShadow) Map() js.Value {
	return dls.Get("map")
}
func (dls *DirectionalLightShadow) SetMap(v js.Value) {
	dls.Set("map", v)
}
func (dls *DirectionalLightShadow) MapSize() *Vector2 {
	return &Vector2{Value: dls.Get("mapSize")}
}
func (dls *DirectionalLightShadow) SetMapSize(v *Vector2) {
	dls.Set("mapSize", v)
}
func (dls *DirectionalLightShadow) Matrix() *Matrix4 {
	return &Matrix4{Value: dls.Get("matrix")}
}
func (dls *DirectionalLightShadow) SetMatrix(v *Matrix4) {
	dls.Set("matrix", v)
}
func (dls *DirectionalLightShadow) Radius() float64 {
	return dls.Get("radius").Float()
}
func (dls *DirectionalLightShadow) SetRadius(v float64) {
	dls.Set("radius", v)
}
func (dls *DirectionalLightShadow) Clone(recursive bool) *DirectionalLightShadow {
	return &DirectionalLightShadow{Value: dls.Call("clone", recursive)}
}
func (dls *DirectionalLightShadow) Copy(source *LightShadow) *DirectionalLightShadow {
	return &DirectionalLightShadow{Value: dls.Call("copy", source)}
}
func (dls *DirectionalLightShadow) ToJSON() js.Value {
	return dls.Call("toJSON")
}

// Code generated from "lights/LightShadow.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type LightShadow struct {
	js.Value
}

func NewLightShadow(camera *Camera) *LightShadow {
	return &LightShadow{Value: get("LightShadow").New(camera)}
}
func (ls *LightShadow) Bias() float64 {
	return ls.Get("bias").Float()
}
func (ls *LightShadow) SetBias(v float64) {
	ls.Set("bias", v)
}
func (ls *LightShadow) Camera() *Camera {
	return &Camera{Value: ls.Get("camera")}
}
func (ls *LightShadow) SetCamera(v *Camera) {
	ls.Set("camera", v)
}
func (ls *LightShadow) Map() js.Value {
	return ls.Get("map")
}
func (ls *LightShadow) SetMap(v js.Value) {
	ls.Set("map", v)
}
func (ls *LightShadow) MapSize() *Vector2 {
	return &Vector2{Value: ls.Get("mapSize")}
}
func (ls *LightShadow) SetMapSize(v *Vector2) {
	ls.Set("mapSize", v)
}
func (ls *LightShadow) Matrix() *Matrix4 {
	return &Matrix4{Value: ls.Get("matrix")}
}
func (ls *LightShadow) SetMatrix(v *Matrix4) {
	ls.Set("matrix", v)
}
func (ls *LightShadow) Radius() float64 {
	return ls.Get("radius").Float()
}
func (ls *LightShadow) SetRadius(v float64) {
	ls.Set("radius", v)
}
func (ls *LightShadow) Clone(recursive bool) *LightShadow {
	return &LightShadow{Value: ls.Call("clone", recursive)}
}
func (ls *LightShadow) Copy(source *LightShadow) *LightShadow {
	return &LightShadow{Value: ls.Call("copy", source)}
}
func (ls *LightShadow) ToJSON() js.Value {
	return ls.Call("toJSON")
}

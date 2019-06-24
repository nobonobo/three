// Code generated from "scenes/FogExp2.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type FogExp2 struct {
	js.Value
}

func NewFogExp2(hex int, density float64) *FogExp2 {
	return &FogExp2{Value: get("FogExp2").New(hex, density)}
}
func (fe *FogExp2) Color() *Color {
	return &Color{Value: fe.Get("color")}
}
func (fe *FogExp2) SetColor(v *Color) {
	fe.Set("color", v)
}
func (fe *FogExp2) Density() float64 {
	return fe.Get("density").Float()
}
func (fe *FogExp2) SetDensity(v float64) {
	fe.Set("density", v)
}
func (fe *FogExp2) Name() string {
	return fe.Get("name").String()
}
func (fe *FogExp2) SetName(v string) {
	fe.Set("name", v)
}
func (fe *FogExp2) Clone() *FogExp2 {
	return &FogExp2{Value: fe.Call("clone")}
}
func (fe *FogExp2) ToJSON() js.Value {
	return fe.Call("toJSON")
}

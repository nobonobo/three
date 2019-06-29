// Code generated from "scenes/Fog.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type IFog interface {
	Clone() *IFog
	ToJSON() js.Value
}

// Fog extend: []
type Fog struct {
	js.Value
}

func NewFog(hex int, near float64, far float64) *Fog {
	return &Fog{Value: get("Fog").New(hex, near, far)}
}
func (ff *Fog) JSValue() js.Value {
	return ff.Value
}
func (ff *Fog) Color() *Color {
	return &Color{Value: ff.Get("color")}
}
func (ff *Fog) SetColor(v *Color) {
	ff.Set("color", v.JSValue())
}
func (ff *Fog) Far() float64 {
	return ff.Get("far").Float()
}
func (ff *Fog) SetFar(v float64) {
	ff.Set("far", v)
}
func (ff *Fog) Name() string {
	return ff.Get("name").String()
}
func (ff *Fog) SetName(v string) {
	ff.Set("name", v)
}
func (ff *Fog) Near() float64 {
	return ff.Get("near").Float()
}
func (ff *Fog) SetNear(v float64) {
	ff.Set("near", v)
}
func (ff *Fog) Clone() *Fog {
	return &Fog{Value: ff.Call("clone")}
}
func (ff *Fog) ToJSON() js.Value {
	return ff.Call("toJSON")
}

// Code generated from "helpers/RectAreaLightHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// RectAreaLightHelper extend: []
type RectAreaLightHelper struct {
	js.Value
}

func NewRectAreaLightHelper(light *RectAreaLight, color *Color) *RectAreaLightHelper {
	return &RectAreaLightHelper{Value: get("RectAreaLightHelper").New(light.JSValue(), color)}
}
func (ralh *RectAreaLightHelper) JSValue() js.Value {
	return ralh.Value
}
func (ralh *RectAreaLightHelper) Color() *Color {
	return &Color{Value: ralh.Get("color")}
}
func (ralh *RectAreaLightHelper) SetColor(v *Color) {
	ralh.Set("color", v.JSValue())
}
func (ralh *RectAreaLightHelper) Light() *RectAreaLight {
	return &RectAreaLight{Value: ralh.Get("light")}
}
func (ralh *RectAreaLightHelper) SetLight(v *RectAreaLight) {
	ralh.Set("light", v.JSValue())
}

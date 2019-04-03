// Code generated from "helpers/RectAreaLightHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type RectAreaLightHelper struct {
	js.Value
}

func NewRectAreaLightHelper(light *RectAreaLight, color *Color) *RectAreaLightHelper {
	return &RectAreaLightHelper{Value: get("RectAreaLightHelper").New(light, color)}
}
func (ralh *RectAreaLightHelper) Color() *Color {
	return &Color{Value: ralh.Get("color")}
}
func (ralh *RectAreaLightHelper) SetColor(v *Color) {
	ralh.Set("color", v)
}
func (ralh *RectAreaLightHelper) Light() *RectAreaLight {
	return &RectAreaLight{Value: ralh.Get("light")}
}
func (ralh *RectAreaLightHelper) SetLight(v *RectAreaLight) {
	ralh.Set("light", v)
}

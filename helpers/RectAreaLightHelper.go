package helpers

import (
	"github.com/nobonobo/three/lights"
	"github.com/nobonobo/three/math"
)

type RectAreaLightHelper struct {
	js.Value
}

func NewRectAreaLightHelper(light *lights.RectAreaLight, color math.Color) *RectAreaLightHelper {
	return &RectAreaLightHelper{Value: get("RectAreaLightHelper").New(light, color)}
}
func (ralh *RectAreaLightHelper) Color() math.Color {
	return math.Color(ralh.Get("color"))
}
func (ralh *RectAreaLightHelper) SetColor(v math.Color) {
	ralh.Set("color", v)
}
func (ralh *RectAreaLightHelper) Light() *lights.RectAreaLight {
	return &lights.RectAreaLight{Value: ralh.Get("light")}
}
func (ralh *RectAreaLightHelper) SetLight(v *lights.RectAreaLight) {
	ralh.Set("light", v)
}

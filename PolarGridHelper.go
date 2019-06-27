// Code generated from "helpers/PolarGridHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PolarGridHelper extend: []
type PolarGridHelper struct {
	js.Value
}

func NewPolarGridHelper(radius float64, radials float64, circles float64, divisions float64, color1 *Color, color2 *Color) *PolarGridHelper {
	return &PolarGridHelper{Value: get("PolarGridHelper").New(radius, radials, circles, divisions, color1, color2)}
}
func (pgh *PolarGridHelper) JSValue() js.Value {
	return pgh.Value
}

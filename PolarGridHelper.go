// Code generated from "helpers/PolarGridHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type PolarGridHelper struct {
	js.Value
}

func NewPolarGridHelper(radius float64, radials float64, circles float64, divisions float64, color1 *Color, color2 *Color) *PolarGridHelper {
	return &PolarGridHelper{Value: get("PolarGridHelper").New(radius, radials, circles, divisions, color1, color2)}
}

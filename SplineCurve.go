// Code generated from "extras/curves/SplineCurve.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type SplineCurve struct {
	js.Value
}

func NewSplineCurve(points js.Value) *SplineCurve {
	return &SplineCurve{Value: get("SplineCurve").New(points)}
}
func (sc *SplineCurve) ArcLengthDivisions() float64 {
	return sc.Get("arcLengthDivisions").Float()
}
func (sc *SplineCurve) SetArcLengthDivisions(v float64) {
	sc.Set("arcLengthDivisions", v)
}
func (sc *SplineCurve) Points() js.Value {
	return sc.Get("points")
}
func (sc *SplineCurve) SetPoints(v js.Value) {
	sc.Set("points", v)
}
func (sc *SplineCurve) GetLength() float64 {
	return sc.Call("getLength").Float()
}
func (sc *SplineCurve) GetLengths(divisions float64) js.Value {
	return sc.Call("getLengths", divisions)
}
func (sc *SplineCurve) GetPoint(t float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: sc.Call("getPoint", t, optionalTarget)}
}
func (sc *SplineCurve) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: sc.Call("getPointAt", u, optionalTarget)}
}
func (sc *SplineCurve) GetPoints(divisions float64) js.Value {
	return sc.Call("getPoints", divisions)
}
func (sc *SplineCurve) GetSpacedPoints(divisions float64) js.Value {
	return sc.Call("getSpacedPoints", divisions)
}
func (sc *SplineCurve) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: sc.Call("getTangent", t)}
}
func (sc *SplineCurve) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: sc.Call("getTangentAt", u)}
}
func (sc *SplineCurve) GetUtoTmapping(u float64, distance float64) float64 {
	return sc.Call("getUtoTmapping", u, distance).Float()
}
func (sc *SplineCurve) UpdateArcLengths() {
	sc.Call("updateArcLengths")
}
func (sc *SplineCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return sc.Call("create", constructorFunc, getPointFunc)
}

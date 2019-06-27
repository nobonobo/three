// Code generated from "extras/curves/CubicBezierCurve.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CubicBezierCurve extend: [Curve]
type CubicBezierCurve struct {
	js.Value
}

func NewCubicBezierCurve(v0 *Vector2, v1 *Vector2, v2 *Vector2, v3 *Vector2) *CubicBezierCurve {
	return &CubicBezierCurve{Value: get("CubicBezierCurve").New(v0, v1, v2, v3)}
}
func (cbc *CubicBezierCurve) JSValue() js.Value {
	return cbc.Value
}
func (cbc *CubicBezierCurve) ArcLengthDivisions() float64 {
	return cbc.Get("arcLengthDivisions").Float()
}
func (cbc *CubicBezierCurve) SetArcLengthDivisions(v float64) {
	cbc.Set("arcLengthDivisions", v)
}
func (cbc *CubicBezierCurve) V0() *Vector2 {
	return &Vector2{Value: cbc.Get("v0")}
}
func (cbc *CubicBezierCurve) SetV0(v *Vector2) {
	cbc.Set("v0", v.Value)
}
func (cbc *CubicBezierCurve) V1() *Vector2 {
	return &Vector2{Value: cbc.Get("v1")}
}
func (cbc *CubicBezierCurve) SetV1(v *Vector2) {
	cbc.Set("v1", v.Value)
}
func (cbc *CubicBezierCurve) V2() *Vector2 {
	return &Vector2{Value: cbc.Get("v2")}
}
func (cbc *CubicBezierCurve) SetV2(v *Vector2) {
	cbc.Set("v2", v.Value)
}
func (cbc *CubicBezierCurve) V3() *Vector2 {
	return &Vector2{Value: cbc.Get("v3")}
}
func (cbc *CubicBezierCurve) SetV3(v *Vector2) {
	cbc.Set("v3", v.Value)
}
func (cbc *CubicBezierCurve) GetLength() float64 {
	return cbc.Call("getLength").Float()
}
func (cbc *CubicBezierCurve) GetLengths(divisions float64) js.Value {
	return cbc.Call("getLengths", divisions)
}
func (cbc *CubicBezierCurve) GetPoint(t float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: cbc.Call("getPoint", t, optionalTarget)}
}
func (cbc *CubicBezierCurve) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: cbc.Call("getPointAt", u, optionalTarget)}
}
func (cbc *CubicBezierCurve) GetPoints(divisions float64) js.Value {
	return cbc.Call("getPoints", divisions)
}
func (cbc *CubicBezierCurve) GetSpacedPoints(divisions float64) js.Value {
	return cbc.Call("getSpacedPoints", divisions)
}
func (cbc *CubicBezierCurve) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: cbc.Call("getTangent", t)}
}
func (cbc *CubicBezierCurve) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: cbc.Call("getTangentAt", u)}
}
func (cbc *CubicBezierCurve) GetUtoTmapping(u float64, distance float64) float64 {
	return cbc.Call("getUtoTmapping", u, distance).Float()
}
func (cbc *CubicBezierCurve) UpdateArcLengths() {
	cbc.Call("updateArcLengths")
}
func (cbc *CubicBezierCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return cbc.Call("create", constructorFunc, getPointFunc)
}

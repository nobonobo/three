// Code generated from "extras/curves/CubicBezierCurve3.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CubicBezierCurve3 extend: [Curve]
type CubicBezierCurve3 struct {
	js.Value
}

func NewCubicBezierCurve3(v0 *Vector3, v1 *Vector3, v2 *Vector3, v3 *Vector3) *CubicBezierCurve3 {
	return &CubicBezierCurve3{Value: get("CubicBezierCurve3").New(v0.JSValue(), v1.JSValue(), v2.JSValue(), v3.JSValue())}
}
func (cbc *CubicBezierCurve3) JSValue() js.Value {
	return cbc.Value
}
func (cbc *CubicBezierCurve3) ArcLengthDivisions() float64 {
	return cbc.Get("arcLengthDivisions").Float()
}
func (cbc *CubicBezierCurve3) SetArcLengthDivisions(v float64) {
	cbc.Set("arcLengthDivisions", v)
}
func (cbc *CubicBezierCurve3) V0() *Vector3 {
	return &Vector3{Value: cbc.Get("v0")}
}
func (cbc *CubicBezierCurve3) SetV0(v *Vector3) {
	cbc.Set("v0", v.JSValue())
}
func (cbc *CubicBezierCurve3) V1() *Vector3 {
	return &Vector3{Value: cbc.Get("v1")}
}
func (cbc *CubicBezierCurve3) SetV1(v *Vector3) {
	cbc.Set("v1", v.JSValue())
}
func (cbc *CubicBezierCurve3) V2() *Vector3 {
	return &Vector3{Value: cbc.Get("v2")}
}
func (cbc *CubicBezierCurve3) SetV2(v *Vector3) {
	cbc.Set("v2", v.JSValue())
}
func (cbc *CubicBezierCurve3) V3() *Vector3 {
	return &Vector3{Value: cbc.Get("v3")}
}
func (cbc *CubicBezierCurve3) SetV3(v *Vector3) {
	cbc.Set("v3", v.JSValue())
}
func (cbc *CubicBezierCurve3) GetLength() float64 {
	return cbc.Call("getLength").Float()
}
func (cbc *CubicBezierCurve3) GetLengths(divisions float64) js.Value {
	return cbc.Call("getLengths", divisions)
}
func (cbc *CubicBezierCurve3) GetPoint(t float64) *Vector3 {
	return &Vector3{Value: cbc.Call("getPoint", t)}
}
func (cbc *CubicBezierCurve3) GetPointAt(u float64, optionalTarget *Vector3) *Vector3 {
	return &Vector3{Value: cbc.Call("getPointAt", u, optionalTarget)}
}
func (cbc *CubicBezierCurve3) GetPoints(divisions float64) js.Value {
	return cbc.Call("getPoints", divisions)
}
func (cbc *CubicBezierCurve3) GetSpacedPoints(divisions float64) js.Value {
	return cbc.Call("getSpacedPoints", divisions)
}
func (cbc *CubicBezierCurve3) GetTangent(t float64) *Vector3 {
	return &Vector3{Value: cbc.Call("getTangent", t)}
}
func (cbc *CubicBezierCurve3) GetTangentAt(u float64) *Vector3 {
	return &Vector3{Value: cbc.Call("getTangentAt", u)}
}
func (cbc *CubicBezierCurve3) GetUtoTmapping(u float64, distance float64) float64 {
	return cbc.Call("getUtoTmapping", u, distance).Float()
}
func (cbc *CubicBezierCurve3) UpdateArcLengths() {
	cbc.Call("updateArcLengths")
}
func (cbc *CubicBezierCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return cbc.Call("create", constructorFunc, getPointFunc)
}

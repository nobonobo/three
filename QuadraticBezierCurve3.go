// Code generated from "extras/curves/QuadraticBezierCurve3.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// QuadraticBezierCurve3 extend: [Curve]
type QuadraticBezierCurve3 struct {
	js.Value
}

func NewQuadraticBezierCurve3(v0 *Vector3, v1 *Vector3, v2 *Vector3) *QuadraticBezierCurve3 {
	return &QuadraticBezierCurve3{Value: get("QuadraticBezierCurve3").New(v0.JSValue(), v1.JSValue(), v2.JSValue())}
}
func (qbc *QuadraticBezierCurve3) JSValue() js.Value {
	return qbc.Value
}
func (qbc *QuadraticBezierCurve3) ArcLengthDivisions() float64 {
	return qbc.Get("arcLengthDivisions").Float()
}
func (qbc *QuadraticBezierCurve3) SetArcLengthDivisions(v float64) {
	qbc.Set("arcLengthDivisions", v)
}
func (qbc *QuadraticBezierCurve3) V0() *Vector3 {
	return &Vector3{Value: qbc.Get("v0")}
}
func (qbc *QuadraticBezierCurve3) SetV0(v *Vector3) {
	qbc.Set("v0", v.JSValue())
}
func (qbc *QuadraticBezierCurve3) V1() *Vector3 {
	return &Vector3{Value: qbc.Get("v1")}
}
func (qbc *QuadraticBezierCurve3) SetV1(v *Vector3) {
	qbc.Set("v1", v.JSValue())
}
func (qbc *QuadraticBezierCurve3) V2() *Vector3 {
	return &Vector3{Value: qbc.Get("v2")}
}
func (qbc *QuadraticBezierCurve3) SetV2(v *Vector3) {
	qbc.Set("v2", v.JSValue())
}
func (qbc *QuadraticBezierCurve3) GetLength() float64 {
	return qbc.Call("getLength").Float()
}
func (qbc *QuadraticBezierCurve3) GetLengths(divisions float64) js.Value {
	return qbc.Call("getLengths", divisions)
}
func (qbc *QuadraticBezierCurve3) GetPoint(t float64) *Vector3 {
	return &Vector3{Value: qbc.Call("getPoint", t)}
}
func (qbc *QuadraticBezierCurve3) GetPointAt(u float64, optionalTarget *Vector3) *Vector3 {
	return &Vector3{Value: qbc.Call("getPointAt", u, optionalTarget)}
}
func (qbc *QuadraticBezierCurve3) GetPoints(divisions float64) js.Value {
	return qbc.Call("getPoints", divisions)
}
func (qbc *QuadraticBezierCurve3) GetSpacedPoints(divisions float64) js.Value {
	return qbc.Call("getSpacedPoints", divisions)
}
func (qbc *QuadraticBezierCurve3) GetTangent(t float64) *Vector3 {
	return &Vector3{Value: qbc.Call("getTangent", t)}
}
func (qbc *QuadraticBezierCurve3) GetTangentAt(u float64) *Vector3 {
	return &Vector3{Value: qbc.Call("getTangentAt", u)}
}
func (qbc *QuadraticBezierCurve3) GetUtoTmapping(u float64, distance float64) float64 {
	return qbc.Call("getUtoTmapping", u, distance).Float()
}
func (qbc *QuadraticBezierCurve3) UpdateArcLengths() {
	qbc.Call("updateArcLengths")
}
func (qbc *QuadraticBezierCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return qbc.Call("create", constructorFunc, getPointFunc)
}

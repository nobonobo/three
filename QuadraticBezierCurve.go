// Code generated from "extras/curves/QuadraticBezierCurve.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type QuadraticBezierCurve struct {
	js.Value
}

func NewQuadraticBezierCurve(v0 *Vector2, v1 *Vector2, v2 *Vector2) *QuadraticBezierCurve {
	return &QuadraticBezierCurve{Value: get("QuadraticBezierCurve").New(v0, v1, v2)}
}
func (qbc *QuadraticBezierCurve) ArcLengthDivisions() float64 {
	return qbc.Get("arcLengthDivisions").Float()
}
func (qbc *QuadraticBezierCurve) SetArcLengthDivisions(v float64) {
	qbc.Set("arcLengthDivisions", v)
}
func (qbc *QuadraticBezierCurve) V0() *Vector2 {
	return &Vector2{Value: qbc.Get("v0")}
}
func (qbc *QuadraticBezierCurve) SetV0(v *Vector2) {
	qbc.Set("v0", v)
}
func (qbc *QuadraticBezierCurve) V1() *Vector2 {
	return &Vector2{Value: qbc.Get("v1")}
}
func (qbc *QuadraticBezierCurve) SetV1(v *Vector2) {
	qbc.Set("v1", v)
}
func (qbc *QuadraticBezierCurve) V2() *Vector2 {
	return &Vector2{Value: qbc.Get("v2")}
}
func (qbc *QuadraticBezierCurve) SetV2(v *Vector2) {
	qbc.Set("v2", v)
}
func (qbc *QuadraticBezierCurve) GetLength() float64 {
	return qbc.Call("getLength").Float()
}
func (qbc *QuadraticBezierCurve) GetLengths(divisions float64) js.Value {
	return qbc.Call("getLengths", divisions)
}
func (qbc *QuadraticBezierCurve) GetPoint(t float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: qbc.Call("getPoint", t, optionalTarget)}
}
func (qbc *QuadraticBezierCurve) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: qbc.Call("getPointAt", u, optionalTarget)}
}
func (qbc *QuadraticBezierCurve) GetPoints(divisions float64) js.Value {
	return qbc.Call("getPoints", divisions)
}
func (qbc *QuadraticBezierCurve) GetSpacedPoints(divisions float64) js.Value {
	return qbc.Call("getSpacedPoints", divisions)
}
func (qbc *QuadraticBezierCurve) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: qbc.Call("getTangent", t)}
}
func (qbc *QuadraticBezierCurve) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: qbc.Call("getTangentAt", u)}
}
func (qbc *QuadraticBezierCurve) GetUtoTmapping(u float64, distance float64) float64 {
	return qbc.Call("getUtoTmapping", u, distance).Float()
}
func (qbc *QuadraticBezierCurve) UpdateArcLengths() {
	qbc.Call("updateArcLengths")
}
func (qbc *QuadraticBezierCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return qbc.Call("create", constructorFunc, getPointFunc)
}

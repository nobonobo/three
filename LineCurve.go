// Code generated from "extras/curves/LineCurve.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type LineCurve struct {
	js.Value
}

func NewLineCurve(v1 *Vector2, v2 *Vector2) *LineCurve {
	return &LineCurve{Value: get("LineCurve").New(v1, v2)}
}
func (lc *LineCurve) ArcLengthDivisions() float64 {
	return lc.Get("arcLengthDivisions").Float()
}
func (lc *LineCurve) SetArcLengthDivisions(v float64) {
	lc.Set("arcLengthDivisions", v)
}
func (lc *LineCurve) V1() *Vector2 {
	return &Vector2{Value: lc.Get("v1")}
}
func (lc *LineCurve) SetV1(v *Vector2) {
	lc.Set("v1", v)
}
func (lc *LineCurve) V2() *Vector2 {
	return &Vector2{Value: lc.Get("v2")}
}
func (lc *LineCurve) SetV2(v *Vector2) {
	lc.Set("v2", v)
}
func (lc *LineCurve) GetLength() float64 {
	return lc.Call("getLength").Float()
}
func (lc *LineCurve) GetLengths(divisions float64) js.Value {
	return lc.Call("getLengths", divisions)
}
func (lc *LineCurve) GetPoint(t float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: lc.Call("getPoint", t, optionalTarget)}
}
func (lc *LineCurve) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: lc.Call("getPointAt", u, optionalTarget)}
}
func (lc *LineCurve) GetPoints(divisions float64) js.Value {
	return lc.Call("getPoints", divisions)
}
func (lc *LineCurve) GetSpacedPoints(divisions float64) js.Value {
	return lc.Call("getSpacedPoints", divisions)
}
func (lc *LineCurve) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: lc.Call("getTangent", t)}
}
func (lc *LineCurve) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: lc.Call("getTangentAt", u)}
}
func (lc *LineCurve) GetUtoTmapping(u float64, distance float64) float64 {
	return lc.Call("getUtoTmapping", u, distance).Float()
}
func (lc *LineCurve) UpdateArcLengths() {
	lc.Call("updateArcLengths")
}
func (lc *LineCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return lc.Call("create", constructorFunc, getPointFunc)
}

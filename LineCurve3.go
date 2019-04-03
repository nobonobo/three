// Code generated from "extras/curves/LineCurve3.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type LineCurve3 struct {
	js.Value
}

func NewLineCurve3(v1 *Vector3, v2 *Vector3) *LineCurve3 {
	return &LineCurve3{Value: get("LineCurve3").New(v1, v2)}
}
func (lc *LineCurve3) ArcLengthDivisions() float64 {
	return lc.Get("arcLengthDivisions").Float()
}
func (lc *LineCurve3) SetArcLengthDivisions(v float64) {
	lc.Set("arcLengthDivisions", v)
}
func (lc *LineCurve3) V1() *Vector3 {
	return &Vector3{Value: lc.Get("v1")}
}
func (lc *LineCurve3) SetV1(v *Vector3) {
	lc.Set("v1", v)
}
func (lc *LineCurve3) V2() *Vector3 {
	return &Vector3{Value: lc.Get("v2")}
}
func (lc *LineCurve3) SetV2(v *Vector3) {
	lc.Set("v2", v)
}
func (lc *LineCurve3) GetLength() float64 {
	return lc.Call("getLength").Float()
}
func (lc *LineCurve3) GetLengths(divisions float64) js.Value {
	return lc.Call("getLengths", divisions)
}
func (lc *LineCurve3) GetPoint(t float64) *Vector3 {
	return &Vector3{Value: lc.Call("getPoint", t)}
}
func (lc *LineCurve3) GetPointAt(u float64, optionalTarget *Vector3) *Vector3 {
	return &Vector3{Value: lc.Call("getPointAt", u, optionalTarget)}
}
func (lc *LineCurve3) GetPoints(divisions float64) js.Value {
	return lc.Call("getPoints", divisions)
}
func (lc *LineCurve3) GetSpacedPoints(divisions float64) js.Value {
	return lc.Call("getSpacedPoints", divisions)
}
func (lc *LineCurve3) GetTangent(t float64) *Vector3 {
	return &Vector3{Value: lc.Call("getTangent", t)}
}
func (lc *LineCurve3) GetTangentAt(u float64) *Vector3 {
	return &Vector3{Value: lc.Call("getTangentAt", u)}
}
func (lc *LineCurve3) GetUtoTmapping(u float64, distance float64) float64 {
	return lc.Call("getUtoTmapping", u, distance).Float()
}
func (lc *LineCurve3) UpdateArcLengths() {
	lc.Call("updateArcLengths")
}
func (lc *LineCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return lc.Call("create", constructorFunc, getPointFunc)
}

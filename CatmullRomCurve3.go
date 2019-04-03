// Code generated from "extras/curves/CatmullRomCurve3.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type CatmullRomCurve3 struct {
	js.Value
}

func NewCatmullRomCurve3(points js.Value, closed bool, curveType string, tension float64) *CatmullRomCurve3 {
	return &CatmullRomCurve3{Value: get("CatmullRomCurve3").New(points, closed, curveType, tension)}
}
func (crc *CatmullRomCurve3) ArcLengthDivisions() float64 {
	return crc.Get("arcLengthDivisions").Float()
}
func (crc *CatmullRomCurve3) SetArcLengthDivisions(v float64) {
	crc.Set("arcLengthDivisions", v)
}
func (crc *CatmullRomCurve3) Points() js.Value {
	return crc.Get("points")
}
func (crc *CatmullRomCurve3) SetPoints(v js.Value) {
	crc.Set("points", v)
}
func (crc *CatmullRomCurve3) GetLength() float64 {
	return crc.Call("getLength").Float()
}
func (crc *CatmullRomCurve3) GetLengths(divisions float64) js.Value {
	return crc.Call("getLengths", divisions)
}
func (crc *CatmullRomCurve3) GetPoint(t float64) *Vector3 {
	return &Vector3{Value: crc.Call("getPoint", t)}
}
func (crc *CatmullRomCurve3) GetPointAt(u float64, optionalTarget *Vector3) *Vector3 {
	return &Vector3{Value: crc.Call("getPointAt", u, optionalTarget)}
}
func (crc *CatmullRomCurve3) GetPoints(divisions float64) js.Value {
	return crc.Call("getPoints", divisions)
}
func (crc *CatmullRomCurve3) GetSpacedPoints(divisions float64) js.Value {
	return crc.Call("getSpacedPoints", divisions)
}
func (crc *CatmullRomCurve3) GetTangent(t float64) *Vector3 {
	return &Vector3{Value: crc.Call("getTangent", t)}
}
func (crc *CatmullRomCurve3) GetTangentAt(u float64) *Vector3 {
	return &Vector3{Value: crc.Call("getTangentAt", u)}
}
func (crc *CatmullRomCurve3) GetUtoTmapping(u float64, distance float64) float64 {
	return crc.Call("getUtoTmapping", u, distance).Float()
}
func (crc *CatmullRomCurve3) UpdateArcLengths() {
	crc.Call("updateArcLengths")
}
func (crc *CatmullRomCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return crc.Call("create", constructorFunc, getPointFunc)
}

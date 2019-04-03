// Code generated from "extras/curves/ArcCurve.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type ArcCurve struct {
	js.Value
}

func NewArcCurve(aX float64, aY float64, aRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool) *ArcCurve {
	return &ArcCurve{Value: get("ArcCurve").New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)}
}
func (ac *ArcCurve) AClockwise() bool {
	return ac.Get("aClockwise").Bool()
}
func (ac *ArcCurve) SetAClockwise(v bool) {
	ac.Set("aClockwise", v)
}
func (ac *ArcCurve) AEndAngle() float64 {
	return ac.Get("aEndAngle").Float()
}
func (ac *ArcCurve) SetAEndAngle(v float64) {
	ac.Set("aEndAngle", v)
}
func (ac *ArcCurve) ARotation() float64 {
	return ac.Get("aRotation").Float()
}
func (ac *ArcCurve) SetARotation(v float64) {
	ac.Set("aRotation", v)
}
func (ac *ArcCurve) AStartAngle() float64 {
	return ac.Get("aStartAngle").Float()
}
func (ac *ArcCurve) SetAStartAngle(v float64) {
	ac.Set("aStartAngle", v)
}
func (ac *ArcCurve) AX() float64 {
	return ac.Get("aX").Float()
}
func (ac *ArcCurve) SetAX(v float64) {
	ac.Set("aX", v)
}
func (ac *ArcCurve) AY() float64 {
	return ac.Get("aY").Float()
}
func (ac *ArcCurve) SetAY(v float64) {
	ac.Set("aY", v)
}
func (ac *ArcCurve) ArcLengthDivisions() float64 {
	return ac.Get("arcLengthDivisions").Float()
}
func (ac *ArcCurve) SetArcLengthDivisions(v float64) {
	ac.Set("arcLengthDivisions", v)
}
func (ac *ArcCurve) XRadius() float64 {
	return ac.Get("xRadius").Float()
}
func (ac *ArcCurve) SetXRadius(v float64) {
	ac.Set("xRadius", v)
}
func (ac *ArcCurve) YRadius() float64 {
	return ac.Get("yRadius").Float()
}
func (ac *ArcCurve) SetYRadius(v float64) {
	ac.Set("yRadius", v)
}
func (ac *ArcCurve) GetLength() float64 {
	return ac.Call("getLength").Float()
}
func (ac *ArcCurve) GetLengths(divisions float64) js.Value {
	return ac.Call("getLengths", divisions)
}
func (ac *ArcCurve) GetPoint(t float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: ac.Call("getPoint", t, optionalTarget)}
}
func (ac *ArcCurve) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: ac.Call("getPointAt", u, optionalTarget)}
}
func (ac *ArcCurve) GetPoints(divisions float64) js.Value {
	return ac.Call("getPoints", divisions)
}
func (ac *ArcCurve) GetSpacedPoints(divisions float64) js.Value {
	return ac.Call("getSpacedPoints", divisions)
}
func (ac *ArcCurve) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: ac.Call("getTangent", t)}
}
func (ac *ArcCurve) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: ac.Call("getTangentAt", u)}
}
func (ac *ArcCurve) GetUtoTmapping(u float64, distance float64) float64 {
	return ac.Call("getUtoTmapping", u, distance).Float()
}
func (ac *ArcCurve) UpdateArcLengths() {
	ac.Call("updateArcLengths")
}
func (ac *ArcCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return ac.Call("create", constructorFunc, getPointFunc)
}

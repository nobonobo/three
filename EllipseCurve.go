// Code generated from "extras/curves/EllipseCurve.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type EllipseCurve struct {
	js.Value
}

func NewEllipseCurve(aX float64, aY float64, xRadius float64, yRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool, aRotation float64) *EllipseCurve {
	return &EllipseCurve{Value: get("EllipseCurve").New(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)}
}
func (ec *EllipseCurve) AClockwise() bool {
	return ec.Get("aClockwise").Bool()
}
func (ec *EllipseCurve) SetAClockwise(v bool) {
	ec.Set("aClockwise", v)
}
func (ec *EllipseCurve) AEndAngle() float64 {
	return ec.Get("aEndAngle").Float()
}
func (ec *EllipseCurve) SetAEndAngle(v float64) {
	ec.Set("aEndAngle", v)
}
func (ec *EllipseCurve) ARotation() float64 {
	return ec.Get("aRotation").Float()
}
func (ec *EllipseCurve) SetARotation(v float64) {
	ec.Set("aRotation", v)
}
func (ec *EllipseCurve) AStartAngle() float64 {
	return ec.Get("aStartAngle").Float()
}
func (ec *EllipseCurve) SetAStartAngle(v float64) {
	ec.Set("aStartAngle", v)
}
func (ec *EllipseCurve) AX() float64 {
	return ec.Get("aX").Float()
}
func (ec *EllipseCurve) SetAX(v float64) {
	ec.Set("aX", v)
}
func (ec *EllipseCurve) AY() float64 {
	return ec.Get("aY").Float()
}
func (ec *EllipseCurve) SetAY(v float64) {
	ec.Set("aY", v)
}
func (ec *EllipseCurve) ArcLengthDivisions() float64 {
	return ec.Get("arcLengthDivisions").Float()
}
func (ec *EllipseCurve) SetArcLengthDivisions(v float64) {
	ec.Set("arcLengthDivisions", v)
}
func (ec *EllipseCurve) XRadius() float64 {
	return ec.Get("xRadius").Float()
}
func (ec *EllipseCurve) SetXRadius(v float64) {
	ec.Set("xRadius", v)
}
func (ec *EllipseCurve) YRadius() float64 {
	return ec.Get("yRadius").Float()
}
func (ec *EllipseCurve) SetYRadius(v float64) {
	ec.Set("yRadius", v)
}
func (ec *EllipseCurve) GetLength() float64 {
	return ec.Call("getLength").Float()
}
func (ec *EllipseCurve) GetLengths(divisions float64) js.Value {
	return ec.Call("getLengths", divisions)
}
func (ec *EllipseCurve) GetPoint(t float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: ec.Call("getPoint", t, optionalTarget)}
}
func (ec *EllipseCurve) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: ec.Call("getPointAt", u, optionalTarget)}
}
func (ec *EllipseCurve) GetPoints(divisions float64) js.Value {
	return ec.Call("getPoints", divisions)
}
func (ec *EllipseCurve) GetSpacedPoints(divisions float64) js.Value {
	return ec.Call("getSpacedPoints", divisions)
}
func (ec *EllipseCurve) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: ec.Call("getTangent", t)}
}
func (ec *EllipseCurve) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: ec.Call("getTangentAt", u)}
}
func (ec *EllipseCurve) GetUtoTmapping(u float64, distance float64) float64 {
	return ec.Call("getUtoTmapping", u, distance).Float()
}
func (ec *EllipseCurve) UpdateArcLengths() {
	ec.Call("updateArcLengths")
}
func (ec *EllipseCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return ec.Call("create", constructorFunc, getPointFunc)
}

// Code generated from "extras/core/Shape.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Shape struct {
	js.Value
}

func NewShape(points js.Value) *Shape {
	return &Shape{Value: get("Shape").New(points)}
}
func (ss *Shape) ArcLengthDivisions() float64 {
	return ss.Get("arcLengthDivisions").Float()
}
func (ss *Shape) SetArcLengthDivisions(v float64) {
	ss.Set("arcLengthDivisions", v)
}
func (ss *Shape) AutoClose() bool {
	return ss.Get("autoClose").Bool()
}
func (ss *Shape) SetAutoClose(v bool) {
	ss.Set("autoClose", v)
}
func (ss *Shape) CurrentPoint() *Vector2 {
	return &Vector2{Value: ss.Get("currentPoint")}
}
func (ss *Shape) SetCurrentPoint(v *Vector2) {
	ss.Set("currentPoint", v)
}
func (ss *Shape) Curves() js.Value {
	return ss.Get("curves")
}
func (ss *Shape) SetCurves(v js.Value) {
	ss.Set("curves", v)
}
func (ss *Shape) Holes() js.Value {
	return ss.Get("holes")
}
func (ss *Shape) SetHoles(v js.Value) {
	ss.Set("holes", v)
}
func (ss *Shape) Absarc(aX float64, aY float64, aRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool) {
	ss.Call("absarc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (ss *Shape) Absellipse(aX float64, aY float64, xRadius float64, yRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool, aRotation float64) {
	ss.Call("absellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (ss *Shape) Add(curve js.Value) {
	ss.Call("add", curve)
}
func (ss *Shape) Arc(aX float64, aY float64, aRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool) {
	ss.Call("arc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (ss *Shape) BezierCurveTo(aCP1x float64, aCP1y float64, aCP2x float64, aCP2y float64, aX float64, aY float64) {
	ss.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
}
func (ss *Shape) CheckConnection() bool {
	return ss.Call("checkConnection").Bool()
}
func (ss *Shape) ClosePath() {
	ss.Call("closePath")
}
func (ss *Shape) CreateGeometry(points js.Value) *Geometry {
	return &Geometry{Value: ss.Call("createGeometry", points)}
}
func (ss *Shape) CreatePointsGeometry(divisions float64) *Geometry {
	return &Geometry{Value: ss.Call("createPointsGeometry", divisions)}
}
func (ss *Shape) CreateSpacedPointsGeometry(divisions float64) *Geometry {
	return &Geometry{Value: ss.Call("createSpacedPointsGeometry", divisions)}
}
func (ss *Shape) Ellipse(aX float64, aY float64, xRadius float64, yRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool, aRotation float64) {
	ss.Call("ellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (ss *Shape) ExtractAllPoints(divisions float64) js.Value {
	return ss.Call("extractAllPoints", divisions)
}
func (ss *Shape) ExtractPoints(divisions float64) js.Value {
	return ss.Call("extractPoints", divisions)
}
func (ss *Shape) Extrude(options js.Value) *ExtrudeGeometry {
	return &ExtrudeGeometry{Value: ss.Call("extrude", options)}
}
func (ss *Shape) FromPoints(vectors js.Value) {
	ss.Call("fromPoints", vectors)
}
func (ss *Shape) GetCurveLengths() js.Value {
	return ss.Call("getCurveLengths")
}
func (ss *Shape) GetLength() float64 {
	return ss.Call("getLength").Float()
}
func (ss *Shape) GetLengths(divisions float64) js.Value {
	return ss.Call("getLengths", divisions)
}
func (ss *Shape) GetPoint(t float64) *Vector2 {
	return &Vector2{Value: ss.Call("getPoint", t)}
}
func (ss *Shape) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: ss.Call("getPointAt", u, optionalTarget)}
}
func (ss *Shape) GetPoints(divisions float64) js.Value {
	return ss.Call("getPoints", divisions)
}
func (ss *Shape) GetPointsHoles(divisions float64) js.Value {
	return ss.Call("getPointsHoles", divisions)
}
func (ss *Shape) GetSpacedPoints(divisions float64) js.Value {
	return ss.Call("getSpacedPoints", divisions)
}
func (ss *Shape) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: ss.Call("getTangent", t)}
}
func (ss *Shape) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: ss.Call("getTangentAt", u)}
}
func (ss *Shape) GetUtoTmapping(u float64, distance float64) float64 {
	return ss.Call("getUtoTmapping", u, distance).Float()
}
func (ss *Shape) LineTo(x float64, y float64) {
	ss.Call("lineTo", x, y)
}
func (ss *Shape) MakeGeometry(options js.Value) *ShapeGeometry {
	return &ShapeGeometry{Value: ss.Call("makeGeometry", options)}
}
func (ss *Shape) MoveTo(x float64, y float64) {
	ss.Call("moveTo", x, y)
}
func (ss *Shape) QuadraticCurveTo(aCPx float64, aCPy float64, aX float64, aY float64) {
	ss.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
}
func (ss *Shape) SetFromPoints(vectors js.Value) {
	ss.Call("setFromPoints", vectors)
}
func (ss *Shape) SplineThru(pts js.Value) {
	ss.Call("splineThru", pts)
}
func (ss *Shape) UpdateArcLengths() {
	ss.Call("updateArcLengths")
}
func (ss *Shape) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return ss.Call("create", constructorFunc, getPointFunc)
}

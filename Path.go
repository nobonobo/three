// Code generated from "extras/core/Path.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type PathActions js.Value

var (
	ARC                PathActions = PathActions(get("ARC"))
	BEZIER_CURVE_TO    PathActions = PathActions(get("BEZIER_CURVE_TO"))
	CSPLINE_THRU       PathActions = PathActions(get("CSPLINE_THRU"))
	ELLIPSE            PathActions = PathActions(get("ELLIPSE"))
	LINE_TO            PathActions = PathActions(get("LINE_TO"))
	MOVE_TO            PathActions = PathActions(get("MOVE_TO"))
	QUADRATIC_CURVE_TO PathActions = PathActions(get("QUADRATIC_CURVE_TO"))
)

type PathAction interface {
}
type Path struct {
	js.Value
}

func NewPath(points js.Value) *Path {
	return &Path{Value: get("Path").New(points)}
}
func (pp *Path) ArcLengthDivisions() float64 {
	return pp.Get("arcLengthDivisions").Float()
}
func (pp *Path) SetArcLengthDivisions(v float64) {
	pp.Set("arcLengthDivisions", v)
}
func (pp *Path) AutoClose() bool {
	return pp.Get("autoClose").Bool()
}
func (pp *Path) SetAutoClose(v bool) {
	pp.Set("autoClose", v)
}
func (pp *Path) CurrentPoint() *Vector2 {
	return &Vector2{Value: pp.Get("currentPoint")}
}
func (pp *Path) SetCurrentPoint(v *Vector2) {
	pp.Set("currentPoint", v)
}
func (pp *Path) Curves() js.Value {
	return pp.Get("curves")
}
func (pp *Path) SetCurves(v js.Value) {
	pp.Set("curves", v)
}
func (pp *Path) Absarc(aX float64, aY float64, aRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool) {
	pp.Call("absarc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (pp *Path) Absellipse(aX float64, aY float64, xRadius float64, yRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool, aRotation float64) {
	pp.Call("absellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (pp *Path) Add(curve js.Value) {
	pp.Call("add", curve)
}
func (pp *Path) Arc(aX float64, aY float64, aRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool) {
	pp.Call("arc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (pp *Path) BezierCurveTo(aCP1x float64, aCP1y float64, aCP2x float64, aCP2y float64, aX float64, aY float64) {
	pp.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
}
func (pp *Path) CheckConnection() bool {
	return pp.Call("checkConnection").Bool()
}
func (pp *Path) ClosePath() {
	pp.Call("closePath")
}
func (pp *Path) CreateGeometry(points js.Value) *Geometry {
	return &Geometry{Value: pp.Call("createGeometry", points)}
}
func (pp *Path) CreatePointsGeometry(divisions float64) *Geometry {
	return &Geometry{Value: pp.Call("createPointsGeometry", divisions)}
}
func (pp *Path) CreateSpacedPointsGeometry(divisions float64) *Geometry {
	return &Geometry{Value: pp.Call("createSpacedPointsGeometry", divisions)}
}
func (pp *Path) Ellipse(aX float64, aY float64, xRadius float64, yRadius float64, aStartAngle float64, aEndAngle float64, aClockwise bool, aRotation float64) {
	pp.Call("ellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (pp *Path) FromPoints(vectors js.Value) {
	pp.Call("fromPoints", vectors)
}
func (pp *Path) GetCurveLengths() js.Value {
	return pp.Call("getCurveLengths")
}
func (pp *Path) GetLength() float64 {
	return pp.Call("getLength").Float()
}
func (pp *Path) GetLengths(divisions float64) js.Value {
	return pp.Call("getLengths", divisions)
}
func (pp *Path) GetPoint(t float64) *Vector2 {
	return &Vector2{Value: pp.Call("getPoint", t)}
}
func (pp *Path) GetPointAt(u float64, optionalTarget *Vector2) *Vector2 {
	return &Vector2{Value: pp.Call("getPointAt", u, optionalTarget)}
}
func (pp *Path) GetPoints(divisions float64) js.Value {
	return pp.Call("getPoints", divisions)
}
func (pp *Path) GetSpacedPoints(divisions float64) js.Value {
	return pp.Call("getSpacedPoints", divisions)
}
func (pp *Path) GetTangent(t float64) *Vector2 {
	return &Vector2{Value: pp.Call("getTangent", t)}
}
func (pp *Path) GetTangentAt(u float64) *Vector2 {
	return &Vector2{Value: pp.Call("getTangentAt", u)}
}
func (pp *Path) GetUtoTmapping(u float64, distance float64) float64 {
	return pp.Call("getUtoTmapping", u, distance).Float()
}
func (pp *Path) LineTo(x float64, y float64) {
	pp.Call("lineTo", x, y)
}
func (pp *Path) MoveTo(x float64, y float64) {
	pp.Call("moveTo", x, y)
}
func (pp *Path) QuadraticCurveTo(aCPx float64, aCPy float64, aX float64, aY float64) {
	pp.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
}
func (pp *Path) SetFromPoints(vectors js.Value) {
	pp.Call("setFromPoints", vectors)
}
func (pp *Path) SplineThru(pts js.Value) {
	pp.Call("splineThru", pts)
}
func (pp *Path) UpdateArcLengths() {
	pp.Call("updateArcLengths")
}
func (pp *Path) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return pp.Call("create", constructorFunc, getPointFunc)
}

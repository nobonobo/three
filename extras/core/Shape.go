package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/geometries"
	"github.com/nobonobo/three/math"
)

type Shape struct {
	js.Value
}

func NewShape(points []math.Vector2) *Shape {
	return &Shape{Value: get("Shape").New(points)}
}
func (s *Shape) ArcLengthDivisions() int {
	return s.Get("arcLengthDivisions").Int()
}
func (s *Shape) SetArcLengthDivisions(v int) {
	s.Set("arcLengthDivisions", v)
}
func (s *Shape) AutoClose() bool {
	return s.Get("autoClose").Bool()
}
func (s *Shape) SetAutoClose(v bool) {
	s.Set("autoClose", v)
}
func (s *Shape) CurrentPoint() *math.Vector2 {
	return &math.Vector2{Value: s.Get("currentPoint")}
}
func (s *Shape) SetCurrentPoint(v *math.Vector2) {
	s.Set("currentPoint", v)
}
func (s *Shape) Curves() []Curve {
	return []Curve(s.Get("curves"))
}
func (s *Shape) SetCurves(v []Curve) {
	s.Set("curves", v)
}
func (s *Shape) Holes() []Path {
	return []Path(s.Get("holes"))
}
func (s *Shape) SetHoles(v []Path) {
	s.Set("holes", v)
}
func (s *Shape) Absarc(aX int, aY int, aRadius int, aStartAngle int, aEndAngle int, aClockwise bool) {
	s.Call("absarc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (s *Shape) Absellipse(aX int, aY int, xRadius int, yRadius int, aStartAngle int, aEndAngle int, aClockwise bool, aRotation int) {
	s.Call("absellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (s *Shape) Add(curve *Curve) {
	s.Call("add", curve)
}
func (s *Shape) Arc(aX int, aY int, aRadius int, aStartAngle int, aEndAngle int, aClockwise bool) {
	s.Call("arc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (s *Shape) BezierCurveTo(aCP1x int, aCP1y int, aCP2x int, aCP2y int, aX int, aY int) {
	s.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
}
func (s *Shape) CheckConnection() bool {
	return s.Call("checkConnection").Bool()
}
func (s *Shape) ClosePath() {
	s.Call("closePath")
}
func (s *Shape) CreateGeometry(points []math.Vector2) *core.Geometry {
	return &core.Geometry{Value: s.Call("createGeometry", points)}
}
func (s *Shape) CreatePointsGeometry(divisions int) *core.Geometry {
	return &core.Geometry{Value: s.Call("createPointsGeometry", divisions)}
}
func (s *Shape) CreateSpacedPointsGeometry(divisions int) *core.Geometry {
	return &core.Geometry{Value: s.Call("createSpacedPointsGeometry", divisions)}
}
func (s *Shape) Ellipse(aX int, aY int, xRadius int, yRadius int, aStartAngle int, aEndAngle int, aClockwise bool, aRotation int) {
	s.Call("ellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (s *Shape) ExtractAllPoints(divisions int) js.Value {
	return s.Call("extractAllPoints", divisions)
}
func (s *Shape) ExtractPoints(divisions int) []math.Vector2 {
	return []math.Vector2(s.Call("extractPoints", divisions))
}
func (s *Shape) Extrude(options js.Value) *geometries.ExtrudeGeometry {
	return &geometries.ExtrudeGeometry{Value: s.Call("extrude", options)}
}
func (s *Shape) FromPoints(vectors []math.Vector2) {
	s.Call("fromPoints", vectors)
}
func (s *Shape) GetCurveLengths() []int {
	return []int(s.Call("getCurveLengths"))
}
func (s *Shape) GetLength() int {
	return s.Call("getLength").Int()
}
func (s *Shape) GetLengths(divisions int) []int {
	return []int(s.Call("getLengths", divisions))
}
func (s *Shape) GetPoint(t int) *math.Vector2 {
	return &math.Vector2{Value: s.Call("getPoint", t)}
}
func (s *Shape) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: s.Call("getPointAt", u, optionalTarget)}
}
func (s *Shape) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(s.Call("getPoints", divisions))
}
func (s *Shape) GetPointsHoles(divisions int) [][]math.Vector2 {
	return [][]math.Vector2(s.Call("getPointsHoles", divisions))
}
func (s *Shape) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(s.Call("getSpacedPoints", divisions))
}
func (s *Shape) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: s.Call("getTangent", t)}
}
func (s *Shape) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: s.Call("getTangentAt", u)}
}
func (s *Shape) GetUtoTmapping(u int, distance int) int {
	return s.Call("getUtoTmapping", u, distance).Int()
}
func (s *Shape) LineTo(x int, y int) {
	s.Call("lineTo", x, y)
}
func (s *Shape) MakeGeometry(options js.Value) *geometries.ShapeGeometry {
	return &geometries.ShapeGeometry{Value: s.Call("makeGeometry", options)}
}
func (s *Shape) MoveTo(x int, y int) {
	s.Call("moveTo", x, y)
}
func (s *Shape) QuadraticCurveTo(aCPx int, aCPy int, aX int, aY int) {
	s.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
}
func (s *Shape) SetFromPoints(vectors []math.Vector2) {
	s.Call("setFromPoints", vectors)
}
func (s *Shape) SplineThru(pts []math.Vector2) {
	s.Call("splineThru", pts)
}
func (s *Shape) UpdateArcLengths() {
	s.Call("updateArcLengths")
}
func (s *Shape) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return s.Call("create", constructorFunc, getPointFunc)
}

package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type PathActions js.Value

var (
	ARC                PathActions = get("ARC")
	BEZIER_CURVE_TO    PathActions = get("BEZIER_CURVE_TO")
	CSPLINE_THRU       PathActions = get("CSPLINE_THRU")
	ELLIPSE            PathActions = get("ELLIPSE")
	LINE_TO            PathActions = get("LINE_TO")
	MOVE_TO            PathActions = get("MOVE_TO")
	QUADRATIC_CURVE_TO PathActions = get("QUADRATIC_CURVE_TO")
)

type PathAction interface {
}
type Path struct {
	js.Value
}

func NewPath(points []math.Vector2) *Path {
	return &Path{Value: get("Path").New(points)}
}
func (p *Path) ArcLengthDivisions() int {
	return p.Get("arcLengthDivisions").Int()
}
func (p *Path) SetArcLengthDivisions(v int) {
	p.Set("arcLengthDivisions", v)
}
func (p *Path) AutoClose() bool {
	return p.Get("autoClose").Bool()
}
func (p *Path) SetAutoClose(v bool) {
	p.Set("autoClose", v)
}
func (p *Path) CurrentPoint() *math.Vector2 {
	return &math.Vector2{Value: p.Get("currentPoint")}
}
func (p *Path) SetCurrentPoint(v *math.Vector2) {
	p.Set("currentPoint", v)
}
func (p *Path) Curves() []Curve {
	return []Curve(p.Get("curves"))
}
func (p *Path) SetCurves(v []Curve) {
	p.Set("curves", v)
}
func (p *Path) Absarc(aX int, aY int, aRadius int, aStartAngle int, aEndAngle int, aClockwise bool) {
	p.Call("absarc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (p *Path) Absellipse(aX int, aY int, xRadius int, yRadius int, aStartAngle int, aEndAngle int, aClockwise bool, aRotation int) {
	p.Call("absellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (p *Path) Add(curve *Curve) {
	p.Call("add", curve)
}
func (p *Path) Arc(aX int, aY int, aRadius int, aStartAngle int, aEndAngle int, aClockwise bool) {
	p.Call("arc", aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)
}
func (p *Path) BezierCurveTo(aCP1x int, aCP1y int, aCP2x int, aCP2y int, aX int, aY int) {
	p.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
}
func (p *Path) CheckConnection() bool {
	return p.Call("checkConnection").Bool()
}
func (p *Path) ClosePath() {
	p.Call("closePath")
}
func (p *Path) CreateGeometry(points []math.Vector2) *core.Geometry {
	return &core.Geometry{Value: p.Call("createGeometry", points)}
}
func (p *Path) CreatePointsGeometry(divisions int) *core.Geometry {
	return &core.Geometry{Value: p.Call("createPointsGeometry", divisions)}
}
func (p *Path) CreateSpacedPointsGeometry(divisions int) *core.Geometry {
	return &core.Geometry{Value: p.Call("createSpacedPointsGeometry", divisions)}
}
func (p *Path) Ellipse(aX int, aY int, xRadius int, yRadius int, aStartAngle int, aEndAngle int, aClockwise bool, aRotation int) {
	p.Call("ellipse", aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)
}
func (p *Path) FromPoints(vectors []math.Vector2) {
	p.Call("fromPoints", vectors)
}
func (p *Path) GetCurveLengths() []int {
	return []int(p.Call("getCurveLengths"))
}
func (p *Path) GetLength() int {
	return p.Call("getLength").Int()
}
func (p *Path) GetLengths(divisions int) []int {
	return []int(p.Call("getLengths", divisions))
}
func (p *Path) GetPoint(t int) *math.Vector2 {
	return &math.Vector2{Value: p.Call("getPoint", t)}
}
func (p *Path) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: p.Call("getPointAt", u, optionalTarget)}
}
func (p *Path) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(p.Call("getPoints", divisions))
}
func (p *Path) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(p.Call("getSpacedPoints", divisions))
}
func (p *Path) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: p.Call("getTangent", t)}
}
func (p *Path) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: p.Call("getTangentAt", u)}
}
func (p *Path) GetUtoTmapping(u int, distance int) int {
	return p.Call("getUtoTmapping", u, distance).Int()
}
func (p *Path) LineTo(x int, y int) {
	p.Call("lineTo", x, y)
}
func (p *Path) MoveTo(x int, y int) {
	p.Call("moveTo", x, y)
}
func (p *Path) QuadraticCurveTo(aCPx int, aCPy int, aX int, aY int) {
	p.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
}
func (p *Path) SetFromPoints(vectors []math.Vector2) {
	p.Call("setFromPoints", vectors)
}
func (p *Path) SplineThru(pts []math.Vector2) {
	p.Call("splineThru", pts)
}
func (p *Path) UpdateArcLengths() {
	p.Call("updateArcLengths")
}
func (p *Path) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return p.Call("create", constructorFunc, getPointFunc)
}

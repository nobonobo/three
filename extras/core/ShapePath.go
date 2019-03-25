package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type ShapePath struct {
	js.Value
}

func NewShapePath() *ShapePath {
	return &ShapePath{Value: get("ShapePath").New()}
}
func (sp *ShapePath) CurrentPath() js.Value {
	return sp.Get("currentPath")
}
func (sp *ShapePath) SetCurrentPath(v js.Value) {
	sp.Set("currentPath", v)
}
func (sp *ShapePath) SubPaths() []js.Value {
	return []js.Value(sp.Get("subPaths"))
}
func (sp *ShapePath) SetSubPaths(v []js.Value) {
	sp.Set("subPaths", v)
}
func (sp *ShapePath) BezierCurveTo(aCP1x int, aCP1y int, aCP2x int, aCP2y int, aX int, aY int) {
	sp.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
}
func (sp *ShapePath) LineTo(x int, y int) {
	sp.Call("lineTo", x, y)
}
func (sp *ShapePath) MoveTo(x int, y int) {
	sp.Call("moveTo", x, y)
}
func (sp *ShapePath) QuadraticCurveTo(aCPx int, aCPy int, aX int, aY int) {
	sp.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
}
func (sp *ShapePath) SplineThru(pts []math.Vector2) {
	sp.Call("splineThru", pts)
}
func (sp *ShapePath) ToShapes(isCCW bool, noHoles js.Value) []Shape {
	return []Shape(sp.Call("toShapes", isCCW, noHoles))
}

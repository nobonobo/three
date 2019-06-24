// Code generated from "extras/core/ShapePath.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
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
func (sp *ShapePath) SubPaths() js.Value {
	return sp.Get("subPaths")
}
func (sp *ShapePath) SetSubPaths(v js.Value) {
	sp.Set("subPaths", v)
}
func (sp *ShapePath) BezierCurveTo(aCP1x float64, aCP1y float64, aCP2x float64, aCP2y float64, aX float64, aY float64) {
	sp.Call("bezierCurveTo", aCP1x, aCP1y, aCP2x, aCP2y, aX, aY)
}
func (sp *ShapePath) LineTo(x float64, y float64) {
	sp.Call("lineTo", x, y)
}
func (sp *ShapePath) MoveTo(x float64, y float64) {
	sp.Call("moveTo", x, y)
}
func (sp *ShapePath) QuadraticCurveTo(aCPx float64, aCPy float64, aX float64, aY float64) {
	sp.Call("quadraticCurveTo", aCPx, aCPy, aX, aY)
}
func (sp *ShapePath) SplineThru(pts js.Value) {
	sp.Call("splineThru", pts)
}
func (sp *ShapePath) ToShapes(isCCW bool, noHoles js.Value) js.Value {
	return sp.Call("toShapes", isCCW, noHoles)
}

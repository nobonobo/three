// Code generated from "ShapeUtils/shapeutils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package ShapeUtils

import (
	"github.com/nobonobo/three"
	"syscall/js"
)

func Area(contour []three.Vec2) float64 {
	return global().Call("area", contour).Float()
}
func IsClockWise(pts []three.Vec2) bool {
	return global().Call("isClockWise", pts).Bool()
}
func Triangulate(contour []three.Vec2, indices bool) js.Value {
	return global().Call("triangulate", contour, indices)
}
func TriangulateShape(contour []three.Vec2, holes []three.Vec2) js.Value {
	return global().Call("triangulateShape", contour, holes)
}

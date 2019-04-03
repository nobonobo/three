// Code generated from "ShapeUtils/shapeutils.d.ts"; DO NOT EDIT.

package ShapeUtils

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three"
)

func Area(contour []three.Vec2) float64 {
	return _Global.Call("area", contour).Float()
}
func IsClockWise(pts []three.Vec2) bool {
	return _Global.Call("isClockWise", pts).Bool()
}
func Triangulate(contour []three.Vec2, indices bool) js.Value {
	return _Global.Call("triangulate", contour, indices)
}
func TriangulateShape(contour []three.Vec2, holes []three.Vec2) js.Value {
	return _Global.Call("triangulateShape", contour, holes)
}

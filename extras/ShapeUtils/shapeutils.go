package shapeutils

import (
	"github.com/nobonobo/three/extras"
)

func Area(contour []extras.Vec2) int {
	return _Global.Call("area", contour).Int()
}
func IsClockWise(pts []extras.Vec2) bool {
	return _Global.Call("isClockWise", pts).Bool()
}
func Triangulate(contour []extras.Vec2, indices bool) []int {
	return []int(_Global.Call("triangulate", contour, indices))
}
func TriangulateShape(contour []extras.Vec2, holes []extras.Vec2) [][]int {
	return [][]int(_Global.Call("triangulateShape", contour, holes))
}

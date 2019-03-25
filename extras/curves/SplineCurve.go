package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type SplineCurve struct {
	js.Value
}

func NewSplineCurve(points []math.Vector2) *SplineCurve {
	return &SplineCurve{Value: get("SplineCurve").New(points)}
}
func (sc *SplineCurve) ArcLengthDivisions() int {
	return sc.Get("arcLengthDivisions").Int()
}
func (sc *SplineCurve) SetArcLengthDivisions(v int) {
	sc.Set("arcLengthDivisions", v)
}
func (sc *SplineCurve) Points() []math.Vector2 {
	return []math.Vector2(sc.Get("points"))
}
func (sc *SplineCurve) SetPoints(v []math.Vector2) {
	sc.Set("points", v)
}
func (sc *SplineCurve) GetLength() int {
	return sc.Call("getLength").Int()
}
func (sc *SplineCurve) GetLengths(divisions int) []int {
	return []int(sc.Call("getLengths", divisions))
}
func (sc *SplineCurve) GetPoint(t int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: sc.Call("getPoint", t, optionalTarget)}
}
func (sc *SplineCurve) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: sc.Call("getPointAt", u, optionalTarget)}
}
func (sc *SplineCurve) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(sc.Call("getPoints", divisions))
}
func (sc *SplineCurve) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(sc.Call("getSpacedPoints", divisions))
}
func (sc *SplineCurve) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: sc.Call("getTangent", t)}
}
func (sc *SplineCurve) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: sc.Call("getTangentAt", u)}
}
func (sc *SplineCurve) GetUtoTmapping(u int, distance int) int {
	return sc.Call("getUtoTmapping", u, distance).Int()
}
func (sc *SplineCurve) UpdateArcLengths() {
	sc.Call("updateArcLengths")
}
func (sc *SplineCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return sc.Call("create", constructorFunc, getPointFunc)
}

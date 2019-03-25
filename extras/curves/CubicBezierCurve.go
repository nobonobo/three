package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type CubicBezierCurve struct {
	js.Value
}

func NewCubicBezierCurve(v0 *math.Vector2, v1 *math.Vector2, v2 *math.Vector2, v3 *math.Vector2) *CubicBezierCurve {
	return &CubicBezierCurve{Value: get("CubicBezierCurve").New(v0, v1, v2, v3)}
}
func (cbc *CubicBezierCurve) ArcLengthDivisions() int {
	return cbc.Get("arcLengthDivisions").Int()
}
func (cbc *CubicBezierCurve) SetArcLengthDivisions(v int) {
	cbc.Set("arcLengthDivisions", v)
}
func (cbc *CubicBezierCurve) V0() *math.Vector2 {
	return &math.Vector2{Value: cbc.Get("v0")}
}
func (cbc *CubicBezierCurve) SetV0(v *math.Vector2) {
	cbc.Set("v0", v)
}
func (cbc *CubicBezierCurve) V1() *math.Vector2 {
	return &math.Vector2{Value: cbc.Get("v1")}
}
func (cbc *CubicBezierCurve) SetV1(v *math.Vector2) {
	cbc.Set("v1", v)
}
func (cbc *CubicBezierCurve) V2() *math.Vector2 {
	return &math.Vector2{Value: cbc.Get("v2")}
}
func (cbc *CubicBezierCurve) SetV2(v *math.Vector2) {
	cbc.Set("v2", v)
}
func (cbc *CubicBezierCurve) V3() *math.Vector2 {
	return &math.Vector2{Value: cbc.Get("v3")}
}
func (cbc *CubicBezierCurve) SetV3(v *math.Vector2) {
	cbc.Set("v3", v)
}
func (cbc *CubicBezierCurve) GetLength() int {
	return cbc.Call("getLength").Int()
}
func (cbc *CubicBezierCurve) GetLengths(divisions int) []int {
	return []int(cbc.Call("getLengths", divisions))
}
func (cbc *CubicBezierCurve) GetPoint(t int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: cbc.Call("getPoint", t, optionalTarget)}
}
func (cbc *CubicBezierCurve) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: cbc.Call("getPointAt", u, optionalTarget)}
}
func (cbc *CubicBezierCurve) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(cbc.Call("getPoints", divisions))
}
func (cbc *CubicBezierCurve) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(cbc.Call("getSpacedPoints", divisions))
}
func (cbc *CubicBezierCurve) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: cbc.Call("getTangent", t)}
}
func (cbc *CubicBezierCurve) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: cbc.Call("getTangentAt", u)}
}
func (cbc *CubicBezierCurve) GetUtoTmapping(u int, distance int) int {
	return cbc.Call("getUtoTmapping", u, distance).Int()
}
func (cbc *CubicBezierCurve) UpdateArcLengths() {
	cbc.Call("updateArcLengths")
}
func (cbc *CubicBezierCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return cbc.Call("create", constructorFunc, getPointFunc)
}

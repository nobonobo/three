package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type CubicBezierCurve3 struct {
	js.Value
}

func NewCubicBezierCurve3(v0 *math.Vector3, v1 *math.Vector3, v2 *math.Vector3, v3 *math.Vector3) *CubicBezierCurve3 {
	return &CubicBezierCurve3{Value: get("CubicBezierCurve3").New(v0, v1, v2, v3)}
}
func (cbc *CubicBezierCurve3) ArcLengthDivisions() int {
	return cbc.Get("arcLengthDivisions").Int()
}
func (cbc *CubicBezierCurve3) SetArcLengthDivisions(v int) {
	cbc.Set("arcLengthDivisions", v)
}
func (cbc *CubicBezierCurve3) V0() *math.Vector3 {
	return &math.Vector3{Value: cbc.Get("v0")}
}
func (cbc *CubicBezierCurve3) SetV0(v *math.Vector3) {
	cbc.Set("v0", v)
}
func (cbc *CubicBezierCurve3) V1() *math.Vector3 {
	return &math.Vector3{Value: cbc.Get("v1")}
}
func (cbc *CubicBezierCurve3) SetV1(v *math.Vector3) {
	cbc.Set("v1", v)
}
func (cbc *CubicBezierCurve3) V2() *math.Vector3 {
	return &math.Vector3{Value: cbc.Get("v2")}
}
func (cbc *CubicBezierCurve3) SetV2(v *math.Vector3) {
	cbc.Set("v2", v)
}
func (cbc *CubicBezierCurve3) V3() *math.Vector3 {
	return &math.Vector3{Value: cbc.Get("v3")}
}
func (cbc *CubicBezierCurve3) SetV3(v *math.Vector3) {
	cbc.Set("v3", v)
}
func (cbc *CubicBezierCurve3) GetLength() int {
	return cbc.Call("getLength").Int()
}
func (cbc *CubicBezierCurve3) GetLengths(divisions int) []int {
	return []int(cbc.Call("getLengths", divisions))
}
func (cbc *CubicBezierCurve3) GetPoint(t int) *math.Vector3 {
	return &math.Vector3{Value: cbc.Call("getPoint", t)}
}
func (cbc *CubicBezierCurve3) GetPointAt(u int, optionalTarget *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: cbc.Call("getPointAt", u, optionalTarget)}
}
func (cbc *CubicBezierCurve3) GetPoints(divisions int) []math.Vector3 {
	return []math.Vector3(cbc.Call("getPoints", divisions))
}
func (cbc *CubicBezierCurve3) GetSpacedPoints(divisions int) []math.Vector3 {
	return []math.Vector3(cbc.Call("getSpacedPoints", divisions))
}
func (cbc *CubicBezierCurve3) GetTangent(t int) *math.Vector3 {
	return &math.Vector3{Value: cbc.Call("getTangent", t)}
}
func (cbc *CubicBezierCurve3) GetTangentAt(u int) *math.Vector3 {
	return &math.Vector3{Value: cbc.Call("getTangentAt", u)}
}
func (cbc *CubicBezierCurve3) GetUtoTmapping(u int, distance int) int {
	return cbc.Call("getUtoTmapping", u, distance).Int()
}
func (cbc *CubicBezierCurve3) UpdateArcLengths() {
	cbc.Call("updateArcLengths")
}
func (cbc *CubicBezierCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return cbc.Call("create", constructorFunc, getPointFunc)
}

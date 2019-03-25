package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type LineCurve3 struct {
	js.Value
}

func NewLineCurve3(v1 *math.Vector3, v2 *math.Vector3) *LineCurve3 {
	return &LineCurve3{Value: get("LineCurve3").New(v1, v2)}
}
func (lc *LineCurve3) ArcLengthDivisions() int {
	return lc.Get("arcLengthDivisions").Int()
}
func (lc *LineCurve3) SetArcLengthDivisions(v int) {
	lc.Set("arcLengthDivisions", v)
}
func (lc *LineCurve3) V1() *math.Vector3 {
	return &math.Vector3{Value: lc.Get("v1")}
}
func (lc *LineCurve3) SetV1(v *math.Vector3) {
	lc.Set("v1", v)
}
func (lc *LineCurve3) V2() *math.Vector3 {
	return &math.Vector3{Value: lc.Get("v2")}
}
func (lc *LineCurve3) SetV2(v *math.Vector3) {
	lc.Set("v2", v)
}
func (lc *LineCurve3) GetLength() int {
	return lc.Call("getLength").Int()
}
func (lc *LineCurve3) GetLengths(divisions int) []int {
	return []int(lc.Call("getLengths", divisions))
}
func (lc *LineCurve3) GetPoint(t int) *math.Vector3 {
	return &math.Vector3{Value: lc.Call("getPoint", t)}
}
func (lc *LineCurve3) GetPointAt(u int, optionalTarget *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: lc.Call("getPointAt", u, optionalTarget)}
}
func (lc *LineCurve3) GetPoints(divisions int) []math.Vector3 {
	return []math.Vector3(lc.Call("getPoints", divisions))
}
func (lc *LineCurve3) GetSpacedPoints(divisions int) []math.Vector3 {
	return []math.Vector3(lc.Call("getSpacedPoints", divisions))
}
func (lc *LineCurve3) GetTangent(t int) *math.Vector3 {
	return &math.Vector3{Value: lc.Call("getTangent", t)}
}
func (lc *LineCurve3) GetTangentAt(u int) *math.Vector3 {
	return &math.Vector3{Value: lc.Call("getTangentAt", u)}
}
func (lc *LineCurve3) GetUtoTmapping(u int, distance int) int {
	return lc.Call("getUtoTmapping", u, distance).Int()
}
func (lc *LineCurve3) UpdateArcLengths() {
	lc.Call("updateArcLengths")
}
func (lc *LineCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return lc.Call("create", constructorFunc, getPointFunc)
}

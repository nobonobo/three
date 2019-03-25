package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type LineCurve struct {
	js.Value
}

func NewLineCurve(v1 *math.Vector2, v2 *math.Vector2) *LineCurve {
	return &LineCurve{Value: get("LineCurve").New(v1, v2)}
}
func (lc *LineCurve) ArcLengthDivisions() int {
	return lc.Get("arcLengthDivisions").Int()
}
func (lc *LineCurve) SetArcLengthDivisions(v int) {
	lc.Set("arcLengthDivisions", v)
}
func (lc *LineCurve) V1() *math.Vector2 {
	return &math.Vector2{Value: lc.Get("v1")}
}
func (lc *LineCurve) SetV1(v *math.Vector2) {
	lc.Set("v1", v)
}
func (lc *LineCurve) V2() *math.Vector2 {
	return &math.Vector2{Value: lc.Get("v2")}
}
func (lc *LineCurve) SetV2(v *math.Vector2) {
	lc.Set("v2", v)
}
func (lc *LineCurve) GetLength() int {
	return lc.Call("getLength").Int()
}
func (lc *LineCurve) GetLengths(divisions int) []int {
	return []int(lc.Call("getLengths", divisions))
}
func (lc *LineCurve) GetPoint(t int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: lc.Call("getPoint", t, optionalTarget)}
}
func (lc *LineCurve) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: lc.Call("getPointAt", u, optionalTarget)}
}
func (lc *LineCurve) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(lc.Call("getPoints", divisions))
}
func (lc *LineCurve) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(lc.Call("getSpacedPoints", divisions))
}
func (lc *LineCurve) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: lc.Call("getTangent", t)}
}
func (lc *LineCurve) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: lc.Call("getTangentAt", u)}
}
func (lc *LineCurve) GetUtoTmapping(u int, distance int) int {
	return lc.Call("getUtoTmapping", u, distance).Int()
}
func (lc *LineCurve) UpdateArcLengths() {
	lc.Call("updateArcLengths")
}
func (lc *LineCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return lc.Call("create", constructorFunc, getPointFunc)
}

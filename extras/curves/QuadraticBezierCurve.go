package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type QuadraticBezierCurve struct {
	js.Value
}

func NewQuadraticBezierCurve(v0 *math.Vector2, v1 *math.Vector2, v2 *math.Vector2) *QuadraticBezierCurve {
	return &QuadraticBezierCurve{Value: get("QuadraticBezierCurve").New(v0, v1, v2)}
}
func (qbc *QuadraticBezierCurve) ArcLengthDivisions() int {
	return qbc.Get("arcLengthDivisions").Int()
}
func (qbc *QuadraticBezierCurve) SetArcLengthDivisions(v int) {
	qbc.Set("arcLengthDivisions", v)
}
func (qbc *QuadraticBezierCurve) V0() *math.Vector2 {
	return &math.Vector2{Value: qbc.Get("v0")}
}
func (qbc *QuadraticBezierCurve) SetV0(v *math.Vector2) {
	qbc.Set("v0", v)
}
func (qbc *QuadraticBezierCurve) V1() *math.Vector2 {
	return &math.Vector2{Value: qbc.Get("v1")}
}
func (qbc *QuadraticBezierCurve) SetV1(v *math.Vector2) {
	qbc.Set("v1", v)
}
func (qbc *QuadraticBezierCurve) V2() *math.Vector2 {
	return &math.Vector2{Value: qbc.Get("v2")}
}
func (qbc *QuadraticBezierCurve) SetV2(v *math.Vector2) {
	qbc.Set("v2", v)
}
func (qbc *QuadraticBezierCurve) GetLength() int {
	return qbc.Call("getLength").Int()
}
func (qbc *QuadraticBezierCurve) GetLengths(divisions int) []int {
	return []int(qbc.Call("getLengths", divisions))
}
func (qbc *QuadraticBezierCurve) GetPoint(t int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: qbc.Call("getPoint", t, optionalTarget)}
}
func (qbc *QuadraticBezierCurve) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: qbc.Call("getPointAt", u, optionalTarget)}
}
func (qbc *QuadraticBezierCurve) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(qbc.Call("getPoints", divisions))
}
func (qbc *QuadraticBezierCurve) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(qbc.Call("getSpacedPoints", divisions))
}
func (qbc *QuadraticBezierCurve) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: qbc.Call("getTangent", t)}
}
func (qbc *QuadraticBezierCurve) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: qbc.Call("getTangentAt", u)}
}
func (qbc *QuadraticBezierCurve) GetUtoTmapping(u int, distance int) int {
	return qbc.Call("getUtoTmapping", u, distance).Int()
}
func (qbc *QuadraticBezierCurve) UpdateArcLengths() {
	qbc.Call("updateArcLengths")
}
func (qbc *QuadraticBezierCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return qbc.Call("create", constructorFunc, getPointFunc)
}

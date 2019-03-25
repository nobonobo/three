package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type QuadraticBezierCurve3 struct {
	js.Value
}

func NewQuadraticBezierCurve3(v0 *math.Vector3, v1 *math.Vector3, v2 *math.Vector3) *QuadraticBezierCurve3 {
	return &QuadraticBezierCurve3{Value: get("QuadraticBezierCurve3").New(v0, v1, v2)}
}
func (qbc *QuadraticBezierCurve3) ArcLengthDivisions() int {
	return qbc.Get("arcLengthDivisions").Int()
}
func (qbc *QuadraticBezierCurve3) SetArcLengthDivisions(v int) {
	qbc.Set("arcLengthDivisions", v)
}
func (qbc *QuadraticBezierCurve3) V0() *math.Vector3 {
	return &math.Vector3{Value: qbc.Get("v0")}
}
func (qbc *QuadraticBezierCurve3) SetV0(v *math.Vector3) {
	qbc.Set("v0", v)
}
func (qbc *QuadraticBezierCurve3) V1() *math.Vector3 {
	return &math.Vector3{Value: qbc.Get("v1")}
}
func (qbc *QuadraticBezierCurve3) SetV1(v *math.Vector3) {
	qbc.Set("v1", v)
}
func (qbc *QuadraticBezierCurve3) V2() *math.Vector3 {
	return &math.Vector3{Value: qbc.Get("v2")}
}
func (qbc *QuadraticBezierCurve3) SetV2(v *math.Vector3) {
	qbc.Set("v2", v)
}
func (qbc *QuadraticBezierCurve3) GetLength() int {
	return qbc.Call("getLength").Int()
}
func (qbc *QuadraticBezierCurve3) GetLengths(divisions int) []int {
	return []int(qbc.Call("getLengths", divisions))
}
func (qbc *QuadraticBezierCurve3) GetPoint(t int) *math.Vector3 {
	return &math.Vector3{Value: qbc.Call("getPoint", t)}
}
func (qbc *QuadraticBezierCurve3) GetPointAt(u int, optionalTarget *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: qbc.Call("getPointAt", u, optionalTarget)}
}
func (qbc *QuadraticBezierCurve3) GetPoints(divisions int) []math.Vector3 {
	return []math.Vector3(qbc.Call("getPoints", divisions))
}
func (qbc *QuadraticBezierCurve3) GetSpacedPoints(divisions int) []math.Vector3 {
	return []math.Vector3(qbc.Call("getSpacedPoints", divisions))
}
func (qbc *QuadraticBezierCurve3) GetTangent(t int) *math.Vector3 {
	return &math.Vector3{Value: qbc.Call("getTangent", t)}
}
func (qbc *QuadraticBezierCurve3) GetTangentAt(u int) *math.Vector3 {
	return &math.Vector3{Value: qbc.Call("getTangentAt", u)}
}
func (qbc *QuadraticBezierCurve3) GetUtoTmapping(u int, distance int) int {
	return qbc.Call("getUtoTmapping", u, distance).Int()
}
func (qbc *QuadraticBezierCurve3) UpdateArcLengths() {
	qbc.Call("updateArcLengths")
}
func (qbc *QuadraticBezierCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return qbc.Call("create", constructorFunc, getPointFunc)
}

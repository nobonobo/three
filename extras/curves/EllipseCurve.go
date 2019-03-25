package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type EllipseCurve struct {
	js.Value
}

func NewEllipseCurve(aX int, aY int, xRadius int, yRadius int, aStartAngle int, aEndAngle int, aClockwise bool, aRotation int) *EllipseCurve {
	return &EllipseCurve{Value: get("EllipseCurve").New(aX, aY, xRadius, yRadius, aStartAngle, aEndAngle, aClockwise, aRotation)}
}
func (ec *EllipseCurve) AClockwise() bool {
	return ec.Get("aClockwise").Bool()
}
func (ec *EllipseCurve) SetAClockwise(v bool) {
	ec.Set("aClockwise", v)
}
func (ec *EllipseCurve) AEndAngle() int {
	return ec.Get("aEndAngle").Int()
}
func (ec *EllipseCurve) SetAEndAngle(v int) {
	ec.Set("aEndAngle", v)
}
func (ec *EllipseCurve) ARotation() int {
	return ec.Get("aRotation").Int()
}
func (ec *EllipseCurve) SetARotation(v int) {
	ec.Set("aRotation", v)
}
func (ec *EllipseCurve) AStartAngle() int {
	return ec.Get("aStartAngle").Int()
}
func (ec *EllipseCurve) SetAStartAngle(v int) {
	ec.Set("aStartAngle", v)
}
func (ec *EllipseCurve) AX() int {
	return ec.Get("aX").Int()
}
func (ec *EllipseCurve) SetAX(v int) {
	ec.Set("aX", v)
}
func (ec *EllipseCurve) AY() int {
	return ec.Get("aY").Int()
}
func (ec *EllipseCurve) SetAY(v int) {
	ec.Set("aY", v)
}
func (ec *EllipseCurve) ArcLengthDivisions() int {
	return ec.Get("arcLengthDivisions").Int()
}
func (ec *EllipseCurve) SetArcLengthDivisions(v int) {
	ec.Set("arcLengthDivisions", v)
}
func (ec *EllipseCurve) XRadius() int {
	return ec.Get("xRadius").Int()
}
func (ec *EllipseCurve) SetXRadius(v int) {
	ec.Set("xRadius", v)
}
func (ec *EllipseCurve) YRadius() int {
	return ec.Get("yRadius").Int()
}
func (ec *EllipseCurve) SetYRadius(v int) {
	ec.Set("yRadius", v)
}
func (ec *EllipseCurve) GetLength() int {
	return ec.Call("getLength").Int()
}
func (ec *EllipseCurve) GetLengths(divisions int) []int {
	return []int(ec.Call("getLengths", divisions))
}
func (ec *EllipseCurve) GetPoint(t int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: ec.Call("getPoint", t, optionalTarget)}
}
func (ec *EllipseCurve) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: ec.Call("getPointAt", u, optionalTarget)}
}
func (ec *EllipseCurve) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(ec.Call("getPoints", divisions))
}
func (ec *EllipseCurve) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(ec.Call("getSpacedPoints", divisions))
}
func (ec *EllipseCurve) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: ec.Call("getTangent", t)}
}
func (ec *EllipseCurve) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: ec.Call("getTangentAt", u)}
}
func (ec *EllipseCurve) GetUtoTmapping(u int, distance int) int {
	return ec.Call("getUtoTmapping", u, distance).Int()
}
func (ec *EllipseCurve) UpdateArcLengths() {
	ec.Call("updateArcLengths")
}
func (ec *EllipseCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return ec.Call("create", constructorFunc, getPointFunc)
}

package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type ArcCurve struct {
	js.Value
}

func NewArcCurve(aX int, aY int, aRadius int, aStartAngle int, aEndAngle int, aClockwise bool) *ArcCurve {
	return &ArcCurve{Value: get("ArcCurve").New(aX, aY, aRadius, aStartAngle, aEndAngle, aClockwise)}
}
func (ac *ArcCurve) AClockwise() bool {
	return ac.Get("aClockwise").Bool()
}
func (ac *ArcCurve) SetAClockwise(v bool) {
	ac.Set("aClockwise", v)
}
func (ac *ArcCurve) AEndAngle() int {
	return ac.Get("aEndAngle").Int()
}
func (ac *ArcCurve) SetAEndAngle(v int) {
	ac.Set("aEndAngle", v)
}
func (ac *ArcCurve) ARotation() int {
	return ac.Get("aRotation").Int()
}
func (ac *ArcCurve) SetARotation(v int) {
	ac.Set("aRotation", v)
}
func (ac *ArcCurve) AStartAngle() int {
	return ac.Get("aStartAngle").Int()
}
func (ac *ArcCurve) SetAStartAngle(v int) {
	ac.Set("aStartAngle", v)
}
func (ac *ArcCurve) AX() int {
	return ac.Get("aX").Int()
}
func (ac *ArcCurve) SetAX(v int) {
	ac.Set("aX", v)
}
func (ac *ArcCurve) AY() int {
	return ac.Get("aY").Int()
}
func (ac *ArcCurve) SetAY(v int) {
	ac.Set("aY", v)
}
func (ac *ArcCurve) ArcLengthDivisions() int {
	return ac.Get("arcLengthDivisions").Int()
}
func (ac *ArcCurve) SetArcLengthDivisions(v int) {
	ac.Set("arcLengthDivisions", v)
}
func (ac *ArcCurve) XRadius() int {
	return ac.Get("xRadius").Int()
}
func (ac *ArcCurve) SetXRadius(v int) {
	ac.Set("xRadius", v)
}
func (ac *ArcCurve) YRadius() int {
	return ac.Get("yRadius").Int()
}
func (ac *ArcCurve) SetYRadius(v int) {
	ac.Set("yRadius", v)
}
func (ac *ArcCurve) GetLength() int {
	return ac.Call("getLength").Int()
}
func (ac *ArcCurve) GetLengths(divisions int) []int {
	return []int(ac.Call("getLengths", divisions))
}
func (ac *ArcCurve) GetPoint(t int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: ac.Call("getPoint", t, optionalTarget)}
}
func (ac *ArcCurve) GetPointAt(u int, optionalTarget *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: ac.Call("getPointAt", u, optionalTarget)}
}
func (ac *ArcCurve) GetPoints(divisions int) []math.Vector2 {
	return []math.Vector2(ac.Call("getPoints", divisions))
}
func (ac *ArcCurve) GetSpacedPoints(divisions int) []math.Vector2 {
	return []math.Vector2(ac.Call("getSpacedPoints", divisions))
}
func (ac *ArcCurve) GetTangent(t int) *math.Vector2 {
	return &math.Vector2{Value: ac.Call("getTangent", t)}
}
func (ac *ArcCurve) GetTangentAt(u int) *math.Vector2 {
	return &math.Vector2{Value: ac.Call("getTangentAt", u)}
}
func (ac *ArcCurve) GetUtoTmapping(u int, distance int) int {
	return ac.Call("getUtoTmapping", u, distance).Int()
}
func (ac *ArcCurve) UpdateArcLengths() {
	ac.Call("updateArcLengths")
}
func (ac *ArcCurve) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return ac.Call("create", constructorFunc, getPointFunc)
}

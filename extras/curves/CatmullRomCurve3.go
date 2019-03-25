package curves

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type CatmullRomCurve3 struct {
	js.Value
}

func NewCatmullRomCurve3(points []math.Vector3, closed bool, curveType string, tension int) *CatmullRomCurve3 {
	return &CatmullRomCurve3{Value: get("CatmullRomCurve3").New(points, closed, curveType, tension)}
}
func (crc *CatmullRomCurve3) ArcLengthDivisions() int {
	return crc.Get("arcLengthDivisions").Int()
}
func (crc *CatmullRomCurve3) SetArcLengthDivisions(v int) {
	crc.Set("arcLengthDivisions", v)
}
func (crc *CatmullRomCurve3) Points() []math.Vector3 {
	return []math.Vector3(crc.Get("points"))
}
func (crc *CatmullRomCurve3) SetPoints(v []math.Vector3) {
	crc.Set("points", v)
}
func (crc *CatmullRomCurve3) GetLength() int {
	return crc.Call("getLength").Int()
}
func (crc *CatmullRomCurve3) GetLengths(divisions int) []int {
	return []int(crc.Call("getLengths", divisions))
}
func (crc *CatmullRomCurve3) GetPoint(t int) *math.Vector3 {
	return &math.Vector3{Value: crc.Call("getPoint", t)}
}
func (crc *CatmullRomCurve3) GetPointAt(u int, optionalTarget *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: crc.Call("getPointAt", u, optionalTarget)}
}
func (crc *CatmullRomCurve3) GetPoints(divisions int) []math.Vector3 {
	return []math.Vector3(crc.Call("getPoints", divisions))
}
func (crc *CatmullRomCurve3) GetSpacedPoints(divisions int) []math.Vector3 {
	return []math.Vector3(crc.Call("getSpacedPoints", divisions))
}
func (crc *CatmullRomCurve3) GetTangent(t int) *math.Vector3 {
	return &math.Vector3{Value: crc.Call("getTangent", t)}
}
func (crc *CatmullRomCurve3) GetTangentAt(u int) *math.Vector3 {
	return &math.Vector3{Value: crc.Call("getTangentAt", u)}
}
func (crc *CatmullRomCurve3) GetUtoTmapping(u int, distance int) int {
	return crc.Call("getUtoTmapping", u, distance).Int()
}
func (crc *CatmullRomCurve3) UpdateArcLengths() {
	crc.Call("updateArcLengths")
}
func (crc *CatmullRomCurve3) Create(constructorFunc js.Value, getPointFunc js.Value) js.Value {
	return crc.Call("create", constructorFunc, getPointFunc)
}

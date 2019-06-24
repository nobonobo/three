// Code generated from "math/Box2.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Box2 struct {
	js.Value
}

func NewBox2(min *Vector2, max *Vector2) *Box2 {
	return &Box2{Value: get("Box2").New(min, max)}
}
func (bb *Box2) Max() *Vector2 {
	return &Vector2{Value: bb.Get("max")}
}
func (bb *Box2) SetMax(v *Vector2) {
	bb.Set("max", v)
}
func (bb *Box2) Min() *Vector2 {
	return &Vector2{Value: bb.Get("min")}
}
func (bb *Box2) SetMin(v *Vector2) {
	bb.Set("min", v)
}
func (bb *Box2) ClampPoint(point *Vector2, target *Vector2) *Vector2 {
	return &Vector2{Value: bb.Call("clampPoint", point, target)}
}
func (bb *Box2) Clone() *Box2 {
	return &Box2{Value: bb.Call("clone")}
}
func (bb *Box2) ContainsBox(box *Box2) bool {
	return bb.Call("containsBox", box).Bool()
}
func (bb *Box2) ContainsPoint(point *Vector2) bool {
	return bb.Call("containsPoint", point).Bool()
}
func (bb *Box2) Copy(box *Box2) *Box2 {
	return &Box2{Value: bb.Call("copy", box)}
}
func (bb *Box2) DistanceToPoint(point *Vector2) float64 {
	return bb.Call("distanceToPoint", point).Float()
}
func (bb *Box2) Empty() js.Value {
	return bb.Call("empty")
}
func (bb *Box2) Equals(box *Box2) bool {
	return bb.Call("equals", box).Bool()
}
func (bb *Box2) ExpandByPoint(point *Vector2) *Box2 {
	return &Box2{Value: bb.Call("expandByPoint", point)}
}
func (bb *Box2) ExpandByScalar(scalar float64) *Box2 {
	return &Box2{Value: bb.Call("expandByScalar", scalar)}
}
func (bb *Box2) ExpandByVector(vector *Vector2) *Box2 {
	return &Box2{Value: bb.Call("expandByVector", vector)}
}
func (bb *Box2) GetCenter(target *Vector2) *Vector2 {
	return &Vector2{Value: bb.Call("getCenter", target)}
}
func (bb *Box2) GetParameter(point *Vector2) *Vector2 {
	return &Vector2{Value: bb.Call("getParameter", point)}
}
func (bb *Box2) GetSize(target *Vector2) *Vector2 {
	return &Vector2{Value: bb.Call("getSize", target)}
}
func (bb *Box2) Intersect(box *Box2) *Box2 {
	return &Box2{Value: bb.Call("intersect", box)}
}
func (bb *Box2) IntersectsBox(box *Box2) bool {
	return bb.Call("intersectsBox", box).Bool()
}
func (bb *Box2) IsEmpty() bool {
	return bb.Call("isEmpty").Bool()
}
func (bb *Box2) IsIntersectionBox(b js.Value) js.Value {
	return bb.Call("isIntersectionBox", b)
}
func (bb *Box2) MakeEmpty() *Box2 {
	return &Box2{Value: bb.Call("makeEmpty")}
}
func (bb *Box2) Set2(min *Vector2, max *Vector2) *Box2 {
	return &Box2{Value: bb.Call("set", min, max)}
}
func (bb *Box2) SetFromCenterAndSize(center *Vector2, size *Vector2) *Box2 {
	return &Box2{Value: bb.Call("setFromCenterAndSize", center, size)}
}
func (bb *Box2) SetFromPoints(points js.Value) *Box2 {
	return &Box2{Value: bb.Call("setFromPoints", points)}
}
func (bb *Box2) Translate(offset *Vector2) *Box2 {
	return &Box2{Value: bb.Call("translate", offset)}
}
func (bb *Box2) Union(box *Box2) *Box2 {
	return &Box2{Value: bb.Call("union", box)}
}

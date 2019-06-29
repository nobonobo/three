// Code generated from "math/Box3.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Box3 extend: []
type Box3 struct {
	js.Value
}

func NewBox3(min *Vector3, max *Vector3) *Box3 {
	return &Box3{Value: get("Box3").New(min.JSValue(), max.JSValue())}
}
func (bb *Box3) JSValue() js.Value {
	return bb.Value
}
func (bb *Box3) Max() *Vector3 {
	return &Vector3{Value: bb.Get("max")}
}
func (bb *Box3) SetMax(v *Vector3) {
	bb.Set("max", v.JSValue())
}
func (bb *Box3) Min() *Vector3 {
	return &Vector3{Value: bb.Get("min")}
}
func (bb *Box3) SetMin(v *Vector3) {
	bb.Set("min", v.JSValue())
}
func (bb *Box3) ApplyMatrix4(matrix *Matrix4) *Box3 {
	return &Box3{Value: bb.Call("applyMatrix4", matrix)}
}
func (bb *Box3) ClampPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("clampPoint", point, target)}
}
func (bb *Box3) Clone() *Box3 {
	return &Box3{Value: bb.Call("clone")}
}
func (bb *Box3) ContainsBox(box *Box3) bool {
	return bb.Call("containsBox", box).Bool()
}
func (bb *Box3) ContainsPoint(point *Vector3) bool {
	return bb.Call("containsPoint", point).Bool()
}
func (bb *Box3) Copy(box *Box3) *Box3 {
	return &Box3{Value: bb.Call("copy", box)}
}
func (bb *Box3) DistanceToPoint(point *Vector3) float64 {
	return bb.Call("distanceToPoint", point).Float()
}
func (bb *Box3) Empty() js.Value {
	return bb.Call("empty")
}
func (bb *Box3) Equals(box *Box3) bool {
	return bb.Call("equals", box).Bool()
}
func (bb *Box3) ExpandByObject(object *Object3D) *Box3 {
	return &Box3{Value: bb.Call("expandByObject", object)}
}
func (bb *Box3) ExpandByPoint(point *Vector3) *Box3 {
	return &Box3{Value: bb.Call("expandByPoint", point)}
}
func (bb *Box3) ExpandByScalar(scalar float64) *Box3 {
	return &Box3{Value: bb.Call("expandByScalar", scalar)}
}
func (bb *Box3) ExpandByVector(vector *Vector3) *Box3 {
	return &Box3{Value: bb.Call("expandByVector", vector)}
}
func (bb *Box3) GetBoundingSphere(target *Sphere) *Sphere {
	return &Sphere{Value: bb.Call("getBoundingSphere", target)}
}
func (bb *Box3) GetCenter(target *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("getCenter", target)}
}
func (bb *Box3) GetParameter(point *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("getParameter", point)}
}
func (bb *Box3) GetSize(target *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("getSize", target)}
}
func (bb *Box3) Intersect(box *Box3) *Box3 {
	return &Box3{Value: bb.Call("intersect", box)}
}
func (bb *Box3) IntersectsBox(box *Box3) bool {
	return bb.Call("intersectsBox", box).Bool()
}
func (bb *Box3) IntersectsPlane(plane *Plane) bool {
	return bb.Call("intersectsPlane", plane).Bool()
}
func (bb *Box3) IntersectsSphere(sphere *Sphere) bool {
	return bb.Call("intersectsSphere", sphere).Bool()
}
func (bb *Box3) IsEmpty() bool {
	return bb.Call("isEmpty").Bool()
}
func (bb *Box3) IsIntersectionBox(b js.Value) js.Value {
	return bb.Call("isIntersectionBox", b)
}
func (bb *Box3) IsIntersectionSphere(s js.Value) js.Value {
	return bb.Call("isIntersectionSphere", s)
}
func (bb *Box3) MakeEmpty() *Box3 {
	return &Box3{Value: bb.Call("makeEmpty")}
}
func (bb *Box3) Set2(min *Vector3, max *Vector3) *Box3 {
	return &Box3{Value: bb.Call("set", min, max)}
}
func (bb *Box3) SetFromArray(array js.Value) *Box3 {
	return &Box3{Value: bb.Call("setFromArray", array)}
}
func (bb *Box3) SetFromCenterAndSize(center *Vector3, size *Vector3) *Box3 {
	return &Box3{Value: bb.Call("setFromCenterAndSize", center, size)}
}
func (bb *Box3) SetFromObject(object *Object3D) *Box3 {
	return &Box3{Value: bb.Call("setFromObject", object)}
}
func (bb *Box3) SetFromPoints(points js.Value) *Box3 {
	return &Box3{Value: bb.Call("setFromPoints", points)}
}
func (bb *Box3) Translate(offset *Vector3) *Box3 {
	return &Box3{Value: bb.Call("translate", offset)}
}
func (bb *Box3) Union(box *Box3) *Box3 {
	return &Box3{Value: bb.Call("union", box)}
}

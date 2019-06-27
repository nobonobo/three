// Code generated from "math/Sphere.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Sphere extend: []
type Sphere struct {
	js.Value
}

func NewSphere(center *Vector3, radius float64) *Sphere {
	return &Sphere{Value: get("Sphere").New(center, radius)}
}
func (ss *Sphere) JSValue() js.Value {
	return ss.Value
}
func (ss *Sphere) Center() *Vector3 {
	return &Vector3{Value: ss.Get("center")}
}
func (ss *Sphere) SetCenter(v *Vector3) {
	ss.Set("center", v.Value)
}
func (ss *Sphere) Radius() float64 {
	return ss.Get("radius").Float()
}
func (ss *Sphere) SetRadius(v float64) {
	ss.Set("radius", v)
}
func (ss *Sphere) ApplyMatrix4(matrix *Matrix4) *Sphere {
	return &Sphere{Value: ss.Call("applyMatrix4", matrix)}
}
func (ss *Sphere) ClampPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("clampPoint", point, target)}
}
func (ss *Sphere) Clone() *Sphere {
	return &Sphere{Value: ss.Call("clone")}
}
func (ss *Sphere) ContainsPoint(point *Vector3) bool {
	return ss.Call("containsPoint", point).Bool()
}
func (ss *Sphere) Copy(sphere *Sphere) *Sphere {
	return &Sphere{Value: ss.Call("copy", sphere)}
}
func (ss *Sphere) DistanceToPoint(point *Vector3) float64 {
	return ss.Call("distanceToPoint", point).Float()
}
func (ss *Sphere) Empty() bool {
	return ss.Call("empty").Bool()
}
func (ss *Sphere) Equals(sphere *Sphere) bool {
	return ss.Call("equals", sphere).Bool()
}
func (ss *Sphere) GetBoundingBox(target *Box3) *Box3 {
	return &Box3{Value: ss.Call("getBoundingBox", target)}
}
func (ss *Sphere) IntersectsBox(box *Box3) bool {
	return ss.Call("intersectsBox", box).Bool()
}
func (ss *Sphere) IntersectsPlane(plane *Plane) bool {
	return ss.Call("intersectsPlane", plane).Bool()
}
func (ss *Sphere) IntersectsSphere(sphere *Sphere) bool {
	return ss.Call("intersectsSphere", sphere).Bool()
}
func (ss *Sphere) Set2(center *Vector3, radius float64) *Sphere {
	return &Sphere{Value: ss.Call("set", center, radius)}
}
func (ss *Sphere) SetFromPoints(points js.Value, optionalCenter *Vector3) *Sphere {
	return &Sphere{Value: ss.Call("setFromPoints", points, optionalCenter)}
}
func (ss *Sphere) Translate(offset *Vector3) *Sphere {
	return &Sphere{Value: ss.Call("translate", offset)}
}

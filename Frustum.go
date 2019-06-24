// Code generated from "math/Frustum.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Frustum struct {
	js.Value
}

func NewFrustum(p0 *Plane, p1 *Plane, p2 *Plane, p3 *Plane, p4 *Plane, p5 *Plane) *Frustum {
	return &Frustum{Value: get("Frustum").New(p0, p1, p2, p3, p4, p5)}
}
func (ff *Frustum) Planes() js.Value {
	return ff.Get("planes")
}
func (ff *Frustum) SetPlanes(v js.Value) {
	ff.Set("planes", v)
}
func (ff *Frustum) Clone() *Frustum {
	return &Frustum{Value: ff.Call("clone")}
}
func (ff *Frustum) ContainsPoint(point *Vector3) bool {
	return ff.Call("containsPoint", point).Bool()
}
func (ff *Frustum) Copy(frustum *Frustum) *Frustum {
	return &Frustum{Value: ff.Call("copy", frustum)}
}
func (ff *Frustum) IntersectsBox(box *Box3) bool {
	return ff.Call("intersectsBox", box).Bool()
}
func (ff *Frustum) IntersectsObject(object *Object3D) bool {
	return ff.Call("intersectsObject", object).Bool()
}
func (ff *Frustum) IntersectsObject2(sprite *Sprite) bool {
	return ff.Call("intersectsObject", sprite).Bool()
}
func (ff *Frustum) IntersectsSphere(sphere *Sphere) bool {
	return ff.Call("intersectsSphere", sphere).Bool()
}
func (ff *Frustum) Set2(p0 float64, p1 float64, p2 float64, p3 float64, p4 float64, p5 float64) *Frustum {
	return &Frustum{Value: ff.Call("set", p0, p1, p2, p3, p4, p5)}
}
func (ff *Frustum) SetFromMatrix(m *Matrix4) *Frustum {
	return &Frustum{Value: ff.Call("setFromMatrix", m)}
}

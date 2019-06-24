// Code generated from "math/Ray.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Ray struct {
	js.Value
}

func NewRay(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{Value: get("Ray").New(origin, direction)}
}
func (rr *Ray) Direction() *Vector3 {
	return &Vector3{Value: rr.Get("direction")}
}
func (rr *Ray) SetDirection(v *Vector3) {
	rr.Set("direction", v)
}
func (rr *Ray) Origin() *Vector3 {
	return &Vector3{Value: rr.Get("origin")}
}
func (rr *Ray) SetOrigin(v *Vector3) {
	rr.Set("origin", v)
}
func (rr *Ray) ApplyMatrix4(matrix4 *Matrix4) *Ray {
	return &Ray{Value: rr.Call("applyMatrix4", matrix4)}
}
func (rr *Ray) At(t float64, target *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("at", t, target)}
}
func (rr *Ray) Clone() *Ray {
	return &Ray{Value: rr.Call("clone")}
}
func (rr *Ray) ClosestPointToPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("closestPointToPoint", point, target)}
}
func (rr *Ray) Copy(ray *Ray) *Ray {
	return &Ray{Value: rr.Call("copy", ray)}
}
func (rr *Ray) DistanceSqToPoint(point *Vector3) float64 {
	return rr.Call("distanceSqToPoint", point).Float()
}
func (rr *Ray) DistanceSqToSegment(v0 *Vector3, v1 *Vector3, optionalPointOnRay *Vector3, optionalPointOnSegment *Vector3) float64 {
	return rr.Call("distanceSqToSegment", v0, v1, optionalPointOnRay, optionalPointOnSegment).Float()
}
func (rr *Ray) DistanceToPlane(plane *Plane) float64 {
	return rr.Call("distanceToPlane", plane).Float()
}
func (rr *Ray) DistanceToPoint(point *Vector3) float64 {
	return rr.Call("distanceToPoint", point).Float()
}
func (rr *Ray) Equals(ray *Ray) bool {
	return rr.Call("equals", ray).Bool()
}
func (rr *Ray) IntersectBox(box *Box3, target *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("intersectBox", box, target)}
}
func (rr *Ray) IntersectPlane(plane *Plane, target *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("intersectPlane", plane, target)}
}
func (rr *Ray) IntersectSphere(sphere *Sphere, target *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("intersectSphere", sphere, target)}
}
func (rr *Ray) IntersectTriangle(a *Vector3, b *Vector3, c *Vector3, backfaceCulling bool, target *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("intersectTriangle", a, b, c, backfaceCulling, target)}
}
func (rr *Ray) IntersectsBox(box *Box3) bool {
	return rr.Call("intersectsBox", box).Bool()
}
func (rr *Ray) IntersectsPlane(plane *Plane) bool {
	return rr.Call("intersectsPlane", plane).Bool()
}
func (rr *Ray) IntersectsSphere(sphere *Sphere) bool {
	return rr.Call("intersectsSphere", sphere).Bool()
}
func (rr *Ray) IsIntersectionBox(b js.Value) js.Value {
	return rr.Call("isIntersectionBox", b)
}
func (rr *Ray) IsIntersectionPlane(p js.Value) js.Value {
	return rr.Call("isIntersectionPlane", p)
}
func (rr *Ray) IsIntersectionSphere(s js.Value) js.Value {
	return rr.Call("isIntersectionSphere", s)
}
func (rr *Ray) LookAt(v *Vector3) *Vector3 {
	return &Vector3{Value: rr.Call("lookAt", v)}
}
func (rr *Ray) Recast(t float64) *Ray {
	return &Ray{Value: rr.Call("recast", t)}
}
func (rr *Ray) Set2(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{Value: rr.Call("set", origin, direction)}
}

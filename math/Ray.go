package math

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Ray struct {
	js.Value
}

func NewRay(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{Value: get("Ray").New(origin, direction)}
}
func (r *Ray) Direction() *Vector3 {
	return &Vector3{Value: r.Get("direction")}
}
func (r *Ray) SetDirection(v *Vector3) {
	r.Set("direction", v)
}
func (r *Ray) Origin() *Vector3 {
	return &Vector3{Value: r.Get("origin")}
}
func (r *Ray) SetOrigin(v *Vector3) {
	r.Set("origin", v)
}
func (r *Ray) ApplyMatrix4(matrix4 *Matrix4) *Ray {
	return &Ray{Value: r.Call("applyMatrix4", matrix4)}
}
func (r *Ray) At(t int, target *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("at", t, target)}
}
func (r *Ray) Clone() this {
	return this(r.Call("clone"))
}
func (r *Ray) ClosestPointToPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("closestPointToPoint", point, target)}
}
func (r *Ray) Copy(ray *Ray) this {
	return this(r.Call("copy", ray))
}
func (r *Ray) DistanceSqToPoint(point *Vector3) int {
	return r.Call("distanceSqToPoint", point).Int()
}
func (r *Ray) DistanceSqToSegment(v0 *Vector3, v1 *Vector3, optionalPointOnRay *Vector3, optionalPointOnSegment *Vector3) int {
	return r.Call("distanceSqToSegment", v0, v1, optionalPointOnRay, optionalPointOnSegment).Int()
}
func (r *Ray) DistanceToPlane(plane *Plane) int {
	return r.Call("distanceToPlane", plane).Int()
}
func (r *Ray) DistanceToPoint(point *Vector3) int {
	return r.Call("distanceToPoint", point).Int()
}
func (r *Ray) Equals(ray *Ray) bool {
	return r.Call("equals", ray).Bool()
}
func (r *Ray) IntersectBox(box *Box3, target *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("intersectBox", box, target)}
}
func (r *Ray) IntersectPlane(plane *Plane, target *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("intersectPlane", plane, target)}
}
func (r *Ray) IntersectSphere(sphere *Sphere, target *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("intersectSphere", sphere, target)}
}
func (r *Ray) IntersectTriangle(a *Vector3, b *Vector3, c *Vector3, backfaceCulling bool, target *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("intersectTriangle", a, b, c, backfaceCulling, target)}
}
func (r *Ray) IntersectsBox(box *Box3) bool {
	return r.Call("intersectsBox", box).Bool()
}
func (r *Ray) IntersectsPlane(plane *Plane) bool {
	return r.Call("intersectsPlane", plane).Bool()
}
func (r *Ray) IntersectsSphere(sphere *Sphere) bool {
	return r.Call("intersectsSphere", sphere).Bool()
}
func (r *Ray) IsIntersectionBox(b js.Value) js.Value {
	return r.Call("isIntersectionBox", b)
}
func (r *Ray) IsIntersectionPlane(p js.Value) js.Value {
	return r.Call("isIntersectionPlane", p)
}
func (r *Ray) IsIntersectionSphere(s js.Value) js.Value {
	return r.Call("isIntersectionSphere", s)
}
func (r *Ray) LookAt(v *Vector3) *Vector3 {
	return &Vector3{Value: r.Call("lookAt", v)}
}
func (r *Ray) Recast(t int) *Ray {
	return &Ray{Value: r.Call("recast", t)}
}
func (r *Ray) Set(origin *Vector3, direction *Vector3) *Ray {
	return &Ray{Value: r.Call("set", origin, direction)}
}

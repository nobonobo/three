package math

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
)

type Box3 struct {
	js.Value
}

func NewBox3(min *Vector3, max *Vector3) *Box3 {
	return &Box3{Value: get("Box3").New(min, max)}
}
func (b *Box3) Max() *Vector3 {
	return &Vector3{Value: b.Get("max")}
}
func (b *Box3) SetMax(v *Vector3) {
	b.Set("max", v)
}
func (b *Box3) Min() *Vector3 {
	return &Vector3{Value: b.Get("min")}
}
func (b *Box3) SetMin(v *Vector3) {
	b.Set("min", v)
}
func (b *Box3) ApplyMatrix4(matrix *Matrix4) this {
	return this(b.Call("applyMatrix4", matrix))
}
func (b *Box3) ClampPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: b.Call("clampPoint", point, target)}
}
func (b *Box3) Clone() this {
	return this(b.Call("clone"))
}
func (b *Box3) ContainsBox(box *Box3) bool {
	return b.Call("containsBox", box).Bool()
}
func (b *Box3) ContainsPoint(point *Vector3) bool {
	return b.Call("containsPoint", point).Bool()
}
func (b *Box3) Copy(box *Box3) this {
	return this(b.Call("copy", box))
}
func (b *Box3) DistanceToPoint(point *Vector3) int {
	return b.Call("distanceToPoint", point).Int()
}
func (b *Box3) Empty() js.Value {
	return b.Call("empty")
}
func (b *Box3) Equals(box *Box3) bool {
	return b.Call("equals", box).Bool()
}
func (b *Box3) ExpandByObject(object *core.Object3D) this {
	return this(b.Call("expandByObject", object))
}
func (b *Box3) ExpandByPoint(point *Vector3) this {
	return this(b.Call("expandByPoint", point))
}
func (b *Box3) ExpandByScalar(scalar int) this {
	return this(b.Call("expandByScalar", scalar))
}
func (b *Box3) ExpandByVector(vector *Vector3) this {
	return this(b.Call("expandByVector", vector))
}
func (b *Box3) GetBoundingSphere(target *Sphere) *Sphere {
	return &Sphere{Value: b.Call("getBoundingSphere", target)}
}
func (b *Box3) GetCenter(target *Vector3) *Vector3 {
	return &Vector3{Value: b.Call("getCenter", target)}
}
func (b *Box3) GetParameter(point *Vector3) *Vector3 {
	return &Vector3{Value: b.Call("getParameter", point)}
}
func (b *Box3) GetSize(target *Vector3) *Vector3 {
	return &Vector3{Value: b.Call("getSize", target)}
}
func (b *Box3) Intersect(box *Box3) this {
	return this(b.Call("intersect", box))
}
func (b *Box3) IntersectsBox(box *Box3) bool {
	return b.Call("intersectsBox", box).Bool()
}
func (b *Box3) IntersectsPlane(plane *Plane) bool {
	return b.Call("intersectsPlane", plane).Bool()
}
func (b *Box3) IntersectsSphere(sphere *Sphere) bool {
	return b.Call("intersectsSphere", sphere).Bool()
}
func (b *Box3) IsEmpty() bool {
	return b.Call("isEmpty").Bool()
}
func (b *Box3) IsIntersectionBox(b js.Value) js.Value {
	return b.Call("isIntersectionBox", b)
}
func (b *Box3) IsIntersectionSphere(s js.Value) js.Value {
	return b.Call("isIntersectionSphere", s)
}
func (b *Box3) MakeEmpty() this {
	return this(b.Call("makeEmpty"))
}
func (b *Box3) Set(min *Vector3, max *Vector3) this {
	return this(b.Call("set", min, max))
}
func (b *Box3) SetFromArray(array *ArrayLike) this {
	return this(b.Call("setFromArray", array))
}
func (b *Box3) SetFromCenterAndSize(center *Vector3, size *Vector3) this {
	return this(b.Call("setFromCenterAndSize", center, size))
}
func (b *Box3) SetFromObject(object *core.Object3D) this {
	return this(b.Call("setFromObject", object))
}
func (b *Box3) SetFromPoints(points []Vector3) this {
	return this(b.Call("setFromPoints", points))
}
func (b *Box3) Translate(offset *Vector3) this {
	return this(b.Call("translate", offset))
}
func (b *Box3) Union(box *Box3) this {
	return this(b.Call("union", box))
}

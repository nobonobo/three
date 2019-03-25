package math

import (
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/objects"
)

type Frustum struct {
	js.Value
}

func NewFrustum(p0 *Plane, p1 *Plane, p2 *Plane, p3 *Plane, p4 *Plane, p5 *Plane) *Frustum {
	return &Frustum{Value: get("Frustum").New(p0, p1, p2, p3, p4, p5)}
}
func (f *Frustum) Planes() []Plane {
	return []Plane(f.Get("planes"))
}
func (f *Frustum) SetPlanes(v []Plane) {
	f.Set("planes", v)
}
func (f *Frustum) Clone() this {
	return this(f.Call("clone"))
}
func (f *Frustum) ContainsPoint(point *Vector3) bool {
	return f.Call("containsPoint", point).Bool()
}
func (f *Frustum) Copy(frustum *Frustum) this {
	return this(f.Call("copy", frustum))
}
func (f *Frustum) IntersectsBox(box *Box3) bool {
	return f.Call("intersectsBox", box).Bool()
}
func (f *Frustum) IntersectsObject(object *core.Object3D) bool {
	return f.Call("intersectsObject", object).Bool()
}
func (f *Frustum) IntersectsObject2(sprite *objects.Sprite) bool {
	return f.Call("intersectsObject", sprite).Bool()
}
func (f *Frustum) IntersectsSphere(sphere *Sphere) bool {
	return f.Call("intersectsSphere", sphere).Bool()
}
func (f *Frustum) Set(p0 int, p1 int, p2 int, p3 int, p4 int, p5 int) *Frustum {
	return &Frustum{Value: f.Call("set", p0, p1, p2, p3, p4, p5)}
}
func (f *Frustum) SetFromMatrix(m *Matrix4) *Frustum {
	return &Frustum{Value: f.Call("setFromMatrix", m)}
}

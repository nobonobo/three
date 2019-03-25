package math

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Box2 struct {
	js.Value
}

func NewBox2(min *Vector2, max *Vector2) *Box2 {
	return &Box2{Value: get("Box2").New(min, max)}
}
func (b *Box2) Max() *Vector2 {
	return &Vector2{Value: b.Get("max")}
}
func (b *Box2) SetMax(v *Vector2) {
	b.Set("max", v)
}
func (b *Box2) Min() *Vector2 {
	return &Vector2{Value: b.Get("min")}
}
func (b *Box2) SetMin(v *Vector2) {
	b.Set("min", v)
}
func (b *Box2) ClampPoint(point *Vector2, target *Vector2) *Vector2 {
	return &Vector2{Value: b.Call("clampPoint", point, target)}
}
func (b *Box2) Clone() this {
	return this(b.Call("clone"))
}
func (b *Box2) ContainsBox(box *Box2) bool {
	return b.Call("containsBox", box).Bool()
}
func (b *Box2) ContainsPoint(point *Vector2) bool {
	return b.Call("containsPoint", point).Bool()
}
func (b *Box2) Copy(box *Box2) this {
	return this(b.Call("copy", box))
}
func (b *Box2) DistanceToPoint(point *Vector2) int {
	return b.Call("distanceToPoint", point).Int()
}
func (b *Box2) Empty() js.Value {
	return b.Call("empty")
}
func (b *Box2) Equals(box *Box2) bool {
	return b.Call("equals", box).Bool()
}
func (b *Box2) ExpandByPoint(point *Vector2) *Box2 {
	return &Box2{Value: b.Call("expandByPoint", point)}
}
func (b *Box2) ExpandByScalar(scalar int) *Box2 {
	return &Box2{Value: b.Call("expandByScalar", scalar)}
}
func (b *Box2) ExpandByVector(vector *Vector2) *Box2 {
	return &Box2{Value: b.Call("expandByVector", vector)}
}
func (b *Box2) GetCenter(target *Vector2) *Vector2 {
	return &Vector2{Value: b.Call("getCenter", target)}
}
func (b *Box2) GetParameter(point *Vector2) *Vector2 {
	return &Vector2{Value: b.Call("getParameter", point)}
}
func (b *Box2) GetSize(target *Vector2) *Vector2 {
	return &Vector2{Value: b.Call("getSize", target)}
}
func (b *Box2) Intersect(box *Box2) *Box2 {
	return &Box2{Value: b.Call("intersect", box)}
}
func (b *Box2) IntersectsBox(box *Box2) bool {
	return b.Call("intersectsBox", box).Bool()
}
func (b *Box2) IsEmpty() bool {
	return b.Call("isEmpty").Bool()
}
func (b *Box2) IsIntersectionBox(b js.Value) js.Value {
	return b.Call("isIntersectionBox", b)
}
func (b *Box2) MakeEmpty() *Box2 {
	return &Box2{Value: b.Call("makeEmpty")}
}
func (b *Box2) Set(min *Vector2, max *Vector2) *Box2 {
	return &Box2{Value: b.Call("set", min, max)}
}
func (b *Box2) SetFromCenterAndSize(center *Vector2, size *Vector2) *Box2 {
	return &Box2{Value: b.Call("setFromCenterAndSize", center, size)}
}
func (b *Box2) SetFromPoints(points []Vector2) *Box2 {
	return &Box2{Value: b.Call("setFromPoints", points)}
}
func (b *Box2) Translate(offset *Vector2) *Box2 {
	return &Box2{Value: b.Call("translate", offset)}
}
func (b *Box2) Union(box *Box2) *Box2 {
	return &Box2{Value: b.Call("union", box)}
}

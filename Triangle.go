// Code generated from "math/Triangle.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type SplineControlPoint interface {
}

// Triangle extend: []
type Triangle struct {
	js.Value
}

func NewTriangle(a *Vector3, b *Vector3, c *Vector3) *Triangle {
	return &Triangle{Value: get("Triangle").New(a, b, c)}
}
func (tt *Triangle) JSValue() js.Value {
	return tt.Value
}
func (tt *Triangle) A() *Vector3 {
	return &Vector3{Value: tt.Get("a")}
}
func (tt *Triangle) SetA(v *Vector3) {
	tt.Set("a", v.Value)
}
func (tt *Triangle) B() *Vector3 {
	return &Vector3{Value: tt.Get("b")}
}
func (tt *Triangle) SetB(v *Vector3) {
	tt.Set("b", v.Value)
}
func (tt *Triangle) C() *Vector3 {
	return &Vector3{Value: tt.Get("c")}
}
func (tt *Triangle) SetC(v *Vector3) {
	tt.Set("c", v.Value)
}
func (tt *Triangle) Clone() *Triangle {
	return &Triangle{Value: tt.Call("clone")}
}
func (tt *Triangle) ClosestPointToPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: tt.Call("closestPointToPoint", point, target)}
}
func (tt *Triangle) ContainsPoint(point *Vector3) bool {
	return tt.Call("containsPoint", point).Bool()
}
func (tt *Triangle) Copy(triangle *Triangle) *Triangle {
	return &Triangle{Value: tt.Call("copy", triangle)}
}
func (tt *Triangle) Equals(triangle *Triangle) bool {
	return tt.Call("equals", triangle).Bool()
}
func (tt *Triangle) GetArea() float64 {
	return tt.Call("getArea").Float()
}
func (tt *Triangle) GetBarycoord(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: tt.Call("getBarycoord", point, target)}
}
func (tt *Triangle) GetMidpoint(target *Vector3) *Vector3 {
	return &Vector3{Value: tt.Call("getMidpoint", target)}
}
func (tt *Triangle) GetNormal(target *Vector3) *Vector3 {
	return &Vector3{Value: tt.Call("getNormal", target)}
}
func (tt *Triangle) GetPlane(target *Vector3) *Plane {
	return &Plane{Value: tt.Call("getPlane", target)}
}
func (tt *Triangle) Set2(a *Vector3, b *Vector3, c *Vector3) *Triangle {
	return &Triangle{Value: tt.Call("set", a, b, c)}
}
func (tt *Triangle) SetFromPointsAndIndices(points js.Value, i0 float64, i1 float64, i2 float64) *Triangle {
	return &Triangle{Value: tt.Call("setFromPointsAndIndices", points, i0, i1, i2)}
}
func (tt *Triangle) ContainsPoint2(point *Vector3, a *Vector3, b *Vector3, c *Vector3) bool {
	return tt.Call("containsPoint", point, a, b, c).Bool()
}
func (tt *Triangle) GetBarycoord2(point *Vector3, a *Vector3, b *Vector3, c *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: tt.Call("getBarycoord", point, a, b, c, target)}
}
func (tt *Triangle) GetNormal2(a *Vector3, b *Vector3, c *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: tt.Call("getNormal", a, b, c, target)}
}

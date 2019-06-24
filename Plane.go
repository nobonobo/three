// Code generated from "math/Plane.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Plane struct {
	js.Value
}

func NewPlane(normal *Vector3, constant float64) *Plane {
	return &Plane{Value: get("Plane").New(normal, constant)}
}
func (pp *Plane) Constant() float64 {
	return pp.Get("constant").Float()
}
func (pp *Plane) SetConstant(v float64) {
	pp.Set("constant", v)
}
func (pp *Plane) Normal() *Vector3 {
	return &Vector3{Value: pp.Get("normal")}
}
func (pp *Plane) SetNormal(v *Vector3) {
	pp.Set("normal", v)
}
func (pp *Plane) ApplyMatrix4(matrix *Matrix4, optionalNormalMatrix *Matrix3) *Plane {
	return &Plane{Value: pp.Call("applyMatrix4", matrix, optionalNormalMatrix)}
}
func (pp *Plane) Clone() *Plane {
	return &Plane{Value: pp.Call("clone")}
}
func (pp *Plane) CoplanarPoint(target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("coplanarPoint", target)}
}
func (pp *Plane) Copy(plane *Plane) *Plane {
	return &Plane{Value: pp.Call("copy", plane)}
}
func (pp *Plane) DistanceToPoint(point *Vector3) float64 {
	return pp.Call("distanceToPoint", point).Float()
}
func (pp *Plane) DistanceToSphere(sphere *Sphere) float64 {
	return pp.Call("distanceToSphere", sphere).Float()
}
func (pp *Plane) Equals(plane *Plane) bool {
	return pp.Call("equals", plane).Bool()
}
func (pp *Plane) IntersectLine(line *Line3, target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("intersectLine", line, target)}
}
func (pp *Plane) IntersectsBox(box *Box3) bool {
	return pp.Call("intersectsBox", box).Bool()
}
func (pp *Plane) IntersectsLine(line *Line3) bool {
	return pp.Call("intersectsLine", line).Bool()
}
func (pp *Plane) IsIntersectionLine(l js.Value) js.Value {
	return pp.Call("isIntersectionLine", l)
}
func (pp *Plane) Negate() *Plane {
	return &Plane{Value: pp.Call("negate")}
}
func (pp *Plane) Normalize() *Plane {
	return &Plane{Value: pp.Call("normalize")}
}
func (pp *Plane) OrthoPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("orthoPoint", point, target)}
}
func (pp *Plane) ProjectPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("projectPoint", point, target)}
}
func (pp *Plane) Set2(normal *Vector3, constant float64) *Plane {
	return &Plane{Value: pp.Call("set", normal, constant)}
}
func (pp *Plane) SetComponents(x float64, y float64, z float64, w float64) *Plane {
	return &Plane{Value: pp.Call("setComponents", x, y, z, w)}
}
func (pp *Plane) SetFromCoplanarPoints(a *Vector3, b *Vector3, c *Vector3) *Plane {
	return &Plane{Value: pp.Call("setFromCoplanarPoints", a, b, c)}
}
func (pp *Plane) SetFromNormalAndCoplanarPoint(normal *Vector3, point *Vector3) *Plane {
	return &Plane{Value: pp.Call("setFromNormalAndCoplanarPoint", normal, point)}
}
func (pp *Plane) Translate(offset *Vector3) *Plane {
	return &Plane{Value: pp.Call("translate", offset)}
}

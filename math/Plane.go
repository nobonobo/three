package math

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Plane struct {
	js.Value
}

func NewPlane(normal *Vector3, constant int) *Plane {
	return &Plane{Value: get("Plane").New(normal, constant)}
}
func (p *Plane) Constant() int {
	return p.Get("constant").Int()
}
func (p *Plane) SetConstant(v int) {
	p.Set("constant", v)
}
func (p *Plane) Normal() *Vector3 {
	return &Vector3{Value: p.Get("normal")}
}
func (p *Plane) SetNormal(v *Vector3) {
	p.Set("normal", v)
}
func (p *Plane) ApplyMatrix4(matrix *Matrix4, optionalNormalMatrix *Matrix3) *Plane {
	return &Plane{Value: p.Call("applyMatrix4", matrix, optionalNormalMatrix)}
}
func (p *Plane) Clone() this {
	return this(p.Call("clone"))
}
func (p *Plane) CoplanarPoint(target *Vector3) *Vector3 {
	return &Vector3{Value: p.Call("coplanarPoint", target)}
}
func (p *Plane) Copy(plane *Plane) this {
	return this(p.Call("copy", plane))
}
func (p *Plane) DistanceToPoint(point *Vector3) int {
	return p.Call("distanceToPoint", point).Int()
}
func (p *Plane) DistanceToSphere(sphere *Sphere) int {
	return p.Call("distanceToSphere", sphere).Int()
}
func (p *Plane) Equals(plane *Plane) bool {
	return p.Call("equals", plane).Bool()
}
func (p *Plane) IntersectLine(line *Line3, target *Vector3) *Vector3 {
	return &Vector3{Value: p.Call("intersectLine", line, target)}
}
func (p *Plane) IntersectsBox(box *Box3) bool {
	return p.Call("intersectsBox", box).Bool()
}
func (p *Plane) IntersectsLine(line *Line3) bool {
	return p.Call("intersectsLine", line).Bool()
}
func (p *Plane) IsIntersectionLine(l js.Value) js.Value {
	return p.Call("isIntersectionLine", l)
}
func (p *Plane) Negate() *Plane {
	return &Plane{Value: p.Call("negate")}
}
func (p *Plane) Normalize() *Plane {
	return &Plane{Value: p.Call("normalize")}
}
func (p *Plane) OrthoPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: p.Call("orthoPoint", point, target)}
}
func (p *Plane) ProjectPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: p.Call("projectPoint", point, target)}
}
func (p *Plane) Set(normal *Vector3, constant int) *Plane {
	return &Plane{Value: p.Call("set", normal, constant)}
}
func (p *Plane) SetComponents(x int, y int, z int, w int) *Plane {
	return &Plane{Value: p.Call("setComponents", x, y, z, w)}
}
func (p *Plane) SetFromCoplanarPoints(a *Vector3, b *Vector3, c *Vector3) *Plane {
	return &Plane{Value: p.Call("setFromCoplanarPoints", a, b, c)}
}
func (p *Plane) SetFromNormalAndCoplanarPoint(normal *Vector3, point *Vector3) *Plane {
	return &Plane{Value: p.Call("setFromNormalAndCoplanarPoint", normal, point)}
}
func (p *Plane) Translate(offset *Vector3) *Plane {
	return &Plane{Value: p.Call("translate", offset)}
}

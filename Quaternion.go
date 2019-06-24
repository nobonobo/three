// Code generated from "math/Quaternion.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Quaternion struct {
	js.Value
}

func NewQuaternion(x float64, y float64, z float64, w float64) *Quaternion {
	return &Quaternion{Value: get("Quaternion").New(x, y, z, w)}
}
func (qq *Quaternion) OnChangeCallback() js.Value {
	return qq.Get("onChangeCallback")
}
func (qq *Quaternion) SetOnChangeCallback(v js.Value) {
	qq.Set("onChangeCallback", v)
}
func (qq *Quaternion) W() float64 {
	return qq.Get("w").Float()
}
func (qq *Quaternion) SetW(v float64) {
	qq.Set("w", v)
}
func (qq *Quaternion) X() float64 {
	return qq.Get("x").Float()
}
func (qq *Quaternion) SetX(v float64) {
	qq.Set("x", v)
}
func (qq *Quaternion) Y() float64 {
	return qq.Get("y").Float()
}
func (qq *Quaternion) SetY(v float64) {
	qq.Set("y", v)
}
func (qq *Quaternion) Z() float64 {
	return qq.Get("z").Float()
}
func (qq *Quaternion) SetZ(v float64) {
	qq.Set("z", v)
}
func (qq *Quaternion) AngleTo(q *Quaternion) float64 {
	return qq.Call("angleTo", q).Float()
}
func (qq *Quaternion) Clone() *Quaternion {
	return &Quaternion{Value: qq.Call("clone")}
}
func (qq *Quaternion) Conjugate() *Quaternion {
	return &Quaternion{Value: qq.Call("conjugate")}
}
func (qq *Quaternion) Copy(q *Quaternion) *Quaternion {
	return &Quaternion{Value: qq.Call("copy", q)}
}
func (qq *Quaternion) Dot(v *Quaternion) float64 {
	return qq.Call("dot", v).Float()
}
func (qq *Quaternion) Equals(v *Quaternion) bool {
	return qq.Call("equals", v).Bool()
}
func (qq *Quaternion) FromArray(n js.Value) *Quaternion {
	return &Quaternion{Value: qq.Call("fromArray", n)}
}
func (qq *Quaternion) FromArray2(xyzw js.Value, offset int) *Quaternion {
	return &Quaternion{Value: qq.Call("fromArray", xyzw, offset)}
}
func (qq *Quaternion) Inverse() *Quaternion {
	return &Quaternion{Value: qq.Call("inverse")}
}
func (qq *Quaternion) Length() float64 {
	return qq.Call("length").Float()
}
func (qq *Quaternion) LengthSq() float64 {
	return qq.Call("lengthSq").Float()
}
func (qq *Quaternion) Multiply(q *Quaternion) *Quaternion {
	return &Quaternion{Value: qq.Call("multiply", q)}
}
func (qq *Quaternion) MultiplyQuaternions(a *Quaternion, b *Quaternion) *Quaternion {
	return &Quaternion{Value: qq.Call("multiplyQuaternions", a, b)}
}
func (qq *Quaternion) MultiplyVector3(v js.Value) js.Value {
	return qq.Call("multiplyVector3", v)
}
func (qq *Quaternion) Normalize() *Quaternion {
	return &Quaternion{Value: qq.Call("normalize")}
}
func (qq *Quaternion) OnChange(callback js.Value) *Quaternion {
	return &Quaternion{Value: qq.Call("onChange", callback)}
}
func (qq *Quaternion) Premultiply(q *Quaternion) *Quaternion {
	return &Quaternion{Value: qq.Call("premultiply", q)}
}
func (qq *Quaternion) RotateTowards(q *Quaternion, step int) *Quaternion {
	return &Quaternion{Value: qq.Call("rotateTowards", q, step)}
}
func (qq *Quaternion) Set2(x float64, y float64, z float64, w float64) *Quaternion {
	return &Quaternion{Value: qq.Call("set", x, y, z, w)}
}
func (qq *Quaternion) SetFromAxisAngle(axis *Vector3, angle float64) *Quaternion {
	return &Quaternion{Value: qq.Call("setFromAxisAngle", axis, angle)}
}
func (qq *Quaternion) SetFromEuler(euler *Euler, update bool) *Quaternion {
	return &Quaternion{Value: qq.Call("setFromEuler", euler, update)}
}
func (qq *Quaternion) SetFromRotationMatrix(m *Matrix4) *Quaternion {
	return &Quaternion{Value: qq.Call("setFromRotationMatrix", m)}
}
func (qq *Quaternion) SetFromUnitVectors(vFrom *Vector3, vTo *Vector3) *Quaternion {
	return &Quaternion{Value: qq.Call("setFromUnitVectors", vFrom, vTo)}
}
func (qq *Quaternion) Slerp(qb *Quaternion, t float64) *Quaternion {
	return &Quaternion{Value: qq.Call("slerp", qb, t)}
}
func (qq *Quaternion) ToArray() js.Value {
	return qq.Call("toArray")
}
func (qq *Quaternion) ToArray2(xyzw js.Value, offset int) js.Value {
	return qq.Call("toArray", xyzw, offset)
}
func (qq *Quaternion) Slerp2(qa *Quaternion, qb *Quaternion, qm *Quaternion, t float64) *Quaternion {
	return &Quaternion{Value: qq.Call("slerp", qa, qb, qm, t)}
}
func (qq *Quaternion) SlerpFlat(dst js.Value, dstOffset int, src0 js.Value, srcOffset int, src1 js.Value, stcOffset1 int, t float64) *Quaternion {
	return &Quaternion{Value: qq.Call("slerpFlat", dst, dstOffset, src0, srcOffset, src1, stcOffset1, t)}
}

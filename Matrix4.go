// Code generated from "math/Matrix4.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Matrix4 struct {
	js.Value
}

func NewMatrix4() *Matrix4 {
	return &Matrix4{Value: get("Matrix4").New()}
}
func (mm *Matrix4) Elements() js.Value {
	return mm.Get("elements")
}
func (mm *Matrix4) SetElements(v js.Value) {
	mm.Set("elements", v)
}
func (mm *Matrix4) ApplyToBuffer(buffer *BufferAttribute, offset int, length int) *BufferAttribute {
	return &BufferAttribute{Value: mm.Call("applyToBuffer", buffer, offset, length)}
}
func (mm *Matrix4) ApplyToBufferAttribute(attribute *BufferAttribute) *BufferAttribute {
	return &BufferAttribute{Value: mm.Call("applyToBufferAttribute", attribute)}
}
func (mm *Matrix4) Clone() *Matrix4 {
	return &Matrix4{Value: mm.Call("clone")}
}
func (mm *Matrix4) Compose(translation *Vector3, rotation *Quaternion, scale *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("compose", translation, rotation, scale)}
}
func (mm *Matrix4) Copy(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("copy", m)}
}
func (mm *Matrix4) CopyPosition(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("copyPosition", m)}
}
func (mm *Matrix4) CrossVector(v js.Value) {
	mm.Call("crossVector", v)
}
func (mm *Matrix4) Decompose(translation *Vector3, rotation *Quaternion, scale *Vector3) js.Value {
	return mm.Call("decompose", translation, rotation, scale)
}
func (mm *Matrix4) Determinant() float64 {
	return mm.Call("determinant").Float()
}
func (mm *Matrix4) Equals(matrix *Matrix4) bool {
	return mm.Call("equals", matrix).Bool()
}
func (mm *Matrix4) ExtractBasis(xAxis *Vector3, yAxis *Vector3, zAxis *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("extractBasis", xAxis, yAxis, zAxis)}
}
func (mm *Matrix4) ExtractPosition(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("extractPosition", m)}
}
func (mm *Matrix4) ExtractRotation(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("extractRotation", m)}
}
func (mm *Matrix4) FlattenToArrayOffset(array js.Value, offset int) js.Value {
	return mm.Call("flattenToArrayOffset", array, offset)
}
func (mm *Matrix4) FromArray(array js.Value, offset int) *Matrix4 {
	return &Matrix4{Value: mm.Call("fromArray", array, offset)}
}
func (mm *Matrix4) GetInverse(m *Matrix4, throwOnDegeneratee bool) *Matrix4 {
	return &Matrix4{Value: mm.Call("getInverse", m, throwOnDegeneratee)}
}
func (mm *Matrix4) GetMaxScaleOnAxis() float64 {
	return mm.Call("getMaxScaleOnAxis").Float()
}
func (mm *Matrix4) Identity() *Matrix4 {
	return &Matrix4{Value: mm.Call("identity")}
}
func (mm *Matrix4) LookAt(eye *Vector3, target *Vector3, up *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("lookAt", eye, target, up)}
}
func (mm *Matrix4) MakeBasis(xAxis *Vector3, yAxis *Vector3, zAxis *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeBasis", xAxis, yAxis, zAxis)}
}
func (mm *Matrix4) MakeOrthographic(left float64, right float64, top float64, bottom float64, near float64, far float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeOrthographic", left, right, top, bottom, near, far)}
}
func (mm *Matrix4) MakePerspective(left float64, right float64, bottom float64, top float64, near float64, far float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makePerspective", left, right, bottom, top, near, far)}
}
func (mm *Matrix4) MakePerspective2(fov float64, aspect float64, near float64, far float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makePerspective", fov, aspect, near, far)}
}
func (mm *Matrix4) MakeRotationAxis(axis *Vector3, angle float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationAxis", axis, angle)}
}
func (mm *Matrix4) MakeRotationFromEuler(euler *Euler) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationFromEuler", euler)}
}
func (mm *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationFromQuaternion", q)}
}
func (mm *Matrix4) MakeRotationX(theta float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationX", theta)}
}
func (mm *Matrix4) MakeRotationY(theta float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationY", theta)}
}
func (mm *Matrix4) MakeRotationZ(theta float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationZ", theta)}
}
func (mm *Matrix4) MakeScale(x float64, y float64, z float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeScale", x, y, z)}
}
func (mm *Matrix4) MakeTranslation(x float64, y float64, z float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeTranslation", x, y, z)}
}
func (mm *Matrix4) Multiply(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiply", m)}
}
func (mm *Matrix4) MultiplyMatrices(a *Matrix4, b *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiplyMatrices", a, b)}
}
func (mm *Matrix4) MultiplyScalar(s float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiplyScalar", s)}
}
func (mm *Matrix4) MultiplyToArray(a *Matrix4, b *Matrix4, r js.Value) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiplyToArray", a, b, r)}
}
func (mm *Matrix4) MultiplyVector3(v js.Value) js.Value {
	return mm.Call("multiplyVector3", v)
}
func (mm *Matrix4) MultiplyVector3Array(array js.Value) js.Value {
	return mm.Call("multiplyVector3Array", array)
}
func (mm *Matrix4) MultiplyVector4(v js.Value) js.Value {
	return mm.Call("multiplyVector4", v)
}
func (mm *Matrix4) Premultiply(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("premultiply", m)}
}
func (mm *Matrix4) RotateAxis(v js.Value) {
	mm.Call("rotateAxis", v)
}
func (mm *Matrix4) Scale(v *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("scale", v)}
}
func (mm *Matrix4) Set2(n11 float64, n12 float64, n13 float64, n14 float64, n21 float64, n22 float64, n23 float64, n24 float64, n31 float64, n32 float64, n33 float64, n34 float64, n41 float64, n42 float64, n43 float64, n44 float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("set", n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44)}
}
func (mm *Matrix4) SetPosition(v *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("setPosition", v)}
}
func (mm *Matrix4) SetRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return &Matrix4{Value: mm.Call("setRotationFromQuaternion", q)}
}
func (mm *Matrix4) ToArray() js.Value {
	return mm.Call("toArray")
}
func (mm *Matrix4) Transpose() *Matrix4 {
	return &Matrix4{Value: mm.Call("transpose")}
}

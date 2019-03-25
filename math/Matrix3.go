package math

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
)

type Matrix interface {
	Clone() this
	Copy(m this) this
	Determinant() int
	GetInverse(matrix Matrix, throwOnInvertible bool) Matrix
	Identity() Matrix
	MultiplyScalar(s int) Matrix
	Transpose() Matrix
}
type Matrix3 struct {
	js.Value
}

func NewMatrix3() *Matrix3 {
	return &Matrix3{Value: get("Matrix3").New()}
}
func (m *Matrix3) Elements() []int {
	return []int(m.Get("elements"))
}
func (m *Matrix3) SetElements(v []int) {
	m.Set("elements", v)
}
func (m *Matrix3) ApplyToBuffer(buffer *core.BufferAttribute, offset int, length int) *core.BufferAttribute {
	return &core.BufferAttribute{Value: m.Call("applyToBuffer", buffer, offset, length)}
}
func (m *Matrix3) ApplyToBufferAttribute(attribute *core.BufferAttribute) *core.BufferAttribute {
	return &core.BufferAttribute{Value: m.Call("applyToBufferAttribute", attribute)}
}
func (m *Matrix3) Clone() this {
	return this(m.Call("clone"))
}
func (m *Matrix3) Copy(m *Matrix3) this {
	return this(m.Call("copy", m))
}
func (m *Matrix3) Determinant() int {
	return m.Call("determinant").Int()
}
func (m *Matrix3) FlattenToArrayOffset(array []int, offset int) []int {
	return []int(m.Call("flattenToArrayOffset", array, offset))
}
func (m *Matrix3) FromArray(array []int, offset int) *Matrix3 {
	return &Matrix3{Value: m.Call("fromArray", array, offset)}
}
func (m *Matrix3) GetInverse(matrix *Matrix3, throwOnDegenerate bool) *Matrix3 {
	return &Matrix3{Value: m.Call("getInverse", matrix, throwOnDegenerate)}
}
func (m *Matrix3) GetInverse2(matrix *Matrix4, throwOnDegenerate bool) *Matrix3 {
	return &Matrix3{Value: m.Call("getInverse", matrix, throwOnDegenerate)}
}
func (m *Matrix3) GetNormalMatrix(matrix4 *Matrix4) *Matrix3 {
	return &Matrix3{Value: m.Call("getNormalMatrix", matrix4)}
}
func (m *Matrix3) Identity() *Matrix3 {
	return &Matrix3{Value: m.Call("identity")}
}
func (m *Matrix3) Multiply(m *Matrix3) *Matrix3 {
	return &Matrix3{Value: m.Call("multiply", m)}
}
func (m *Matrix3) MultiplyMatrices(a *Matrix3, b *Matrix3) *Matrix3 {
	return &Matrix3{Value: m.Call("multiplyMatrices", a, b)}
}
func (m *Matrix3) MultiplyScalar(s int) *Matrix3 {
	return &Matrix3{Value: m.Call("multiplyScalar", s)}
}
func (m *Matrix3) MultiplyVector3(vector *Vector3) js.Value {
	return m.Call("multiplyVector3", vector)
}
func (m *Matrix3) MultiplyVector3Array(a js.Value) js.Value {
	return m.Call("multiplyVector3Array", a)
}
func (m *Matrix3) Premultiply(m *Matrix3) *Matrix3 {
	return &Matrix3{Value: m.Call("premultiply", m)}
}
func (m *Matrix3) Set(n11 int, n12 int, n13 int, n21 int, n22 int, n23 int, n31 int, n32 int, n33 int) *Matrix3 {
	return &Matrix3{Value: m.Call("set", n11, n12, n13, n21, n22, n23, n31, n32, n33)}
}
func (m *Matrix3) SetFromMatrix4(m *Matrix4) *Matrix3 {
	return &Matrix3{Value: m.Call("setFromMatrix4", m)}
}
func (m *Matrix3) ToArray() []int {
	return []int(m.Call("toArray"))
}
func (m *Matrix3) Transpose() *Matrix3 {
	return &Matrix3{Value: m.Call("transpose")}
}
func (m *Matrix3) TransposeIntoArray(r []int) []int {
	return []int(m.Call("transposeIntoArray", r))
}

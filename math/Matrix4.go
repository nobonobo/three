package math

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
)

type Matrix4 struct {
	js.Value
}

func NewMatrix4() *Matrix4 {
	return &Matrix4{Value: get("Matrix4").New()}
}
func (m *Matrix4) Elements() []int {
	return []int(m.Get("elements"))
}
func (m *Matrix4) SetElements(v []int) {
	m.Set("elements", v)
}
func (m *Matrix4) ApplyToBuffer(buffer *core.BufferAttribute, offset int, length int) *core.BufferAttribute {
	return &core.BufferAttribute{Value: m.Call("applyToBuffer", buffer, offset, length)}
}
func (m *Matrix4) ApplyToBufferAttribute(attribute *core.BufferAttribute) *core.BufferAttribute {
	return &core.BufferAttribute{Value: m.Call("applyToBufferAttribute", attribute)}
}
func (m *Matrix4) Clone() this {
	return this(m.Call("clone"))
}
func (m *Matrix4) Compose(translation *Vector3, rotation *Quaternion, scale *Vector3) *Matrix4 {
	return &Matrix4{Value: m.Call("compose", translation, rotation, scale)}
}
func (m *Matrix4) Copy(m *Matrix4) this {
	return this(m.Call("copy", m))
}
func (m *Matrix4) CopyPosition(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: m.Call("copyPosition", m)}
}
func (m *Matrix4) CrossVector(v js.Value) {
	m.Call("crossVector", v)
}
func (m *Matrix4) Decompose(translation *Vector3, rotation *Quaternion, scale *Vector3) []Object {
	return []Object(m.Call("decompose", translation, rotation, scale))
}
func (m *Matrix4) Determinant() int {
	return m.Call("determinant").Int()
}
func (m *Matrix4) Equals(matrix *Matrix4) bool {
	return m.Call("equals", matrix).Bool()
}
func (m *Matrix4) ExtractBasis(xAxis *Vector3, yAxis *Vector3, zAxis *Vector3) *Matrix4 {
	return &Matrix4{Value: m.Call("extractBasis", xAxis, yAxis, zAxis)}
}
func (m *Matrix4) ExtractPosition(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: m.Call("extractPosition", m)}
}
func (m *Matrix4) ExtractRotation(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: m.Call("extractRotation", m)}
}
func (m *Matrix4) FlattenToArrayOffset(array []int, offset int) []int {
	return []int(m.Call("flattenToArrayOffset", array, offset))
}
func (m *Matrix4) FromArray(array []int, offset int) *Matrix4 {
	return &Matrix4{Value: m.Call("fromArray", array, offset)}
}
func (m *Matrix4) GetInverse(m *Matrix4, throwOnDegeneratee bool) *Matrix4 {
	return &Matrix4{Value: m.Call("getInverse", m, throwOnDegeneratee)}
}
func (m *Matrix4) GetMaxScaleOnAxis() int {
	return m.Call("getMaxScaleOnAxis").Int()
}
func (m *Matrix4) Identity() *Matrix4 {
	return &Matrix4{Value: m.Call("identity")}
}
func (m *Matrix4) LookAt(eye *Vector3, target *Vector3, up *Vector3) *Matrix4 {
	return &Matrix4{Value: m.Call("lookAt", eye, target, up)}
}
func (m *Matrix4) MakeBasis(xAxis *Vector3, yAxis *Vector3, zAxis *Vector3) *Matrix4 {
	return &Matrix4{Value: m.Call("makeBasis", xAxis, yAxis, zAxis)}
}
func (m *Matrix4) MakeOrthographic(left int, right int, top int, bottom int, near int, far int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeOrthographic", left, right, top, bottom, near, far)}
}
func (m *Matrix4) MakePerspective(left int, right int, bottom int, top int, near int, far int) *Matrix4 {
	return &Matrix4{Value: m.Call("makePerspective", left, right, bottom, top, near, far)}
}
func (m *Matrix4) MakePerspective2(fov int, aspect int, near int, far int) *Matrix4 {
	return &Matrix4{Value: m.Call("makePerspective", fov, aspect, near, far)}
}
func (m *Matrix4) MakeRotationAxis(axis *Vector3, angle int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeRotationAxis", axis, angle)}
}
func (m *Matrix4) MakeRotationFromEuler(euler *Euler) *Matrix4 {
	return &Matrix4{Value: m.Call("makeRotationFromEuler", euler)}
}
func (m *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return &Matrix4{Value: m.Call("makeRotationFromQuaternion", q)}
}
func (m *Matrix4) MakeRotationX(theta int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeRotationX", theta)}
}
func (m *Matrix4) MakeRotationY(theta int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeRotationY", theta)}
}
func (m *Matrix4) MakeRotationZ(theta int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeRotationZ", theta)}
}
func (m *Matrix4) MakeScale(x int, y int, z int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeScale", x, y, z)}
}
func (m *Matrix4) MakeTranslation(x int, y int, z int) *Matrix4 {
	return &Matrix4{Value: m.Call("makeTranslation", x, y, z)}
}
func (m *Matrix4) Multiply(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: m.Call("multiply", m)}
}
func (m *Matrix4) MultiplyMatrices(a *Matrix4, b *Matrix4) *Matrix4 {
	return &Matrix4{Value: m.Call("multiplyMatrices", a, b)}
}
func (m *Matrix4) MultiplyScalar(s int) *Matrix4 {
	return &Matrix4{Value: m.Call("multiplyScalar", s)}
}
func (m *Matrix4) MultiplyToArray(a *Matrix4, b *Matrix4, r []int) *Matrix4 {
	return &Matrix4{Value: m.Call("multiplyToArray", a, b, r)}
}
func (m *Matrix4) MultiplyVector3(v js.Value) js.Value {
	return m.Call("multiplyVector3", v)
}
func (m *Matrix4) MultiplyVector3Array(array []int) []int {
	return []int(m.Call("multiplyVector3Array", array))
}
func (m *Matrix4) MultiplyVector4(v js.Value) js.Value {
	return m.Call("multiplyVector4", v)
}
func (m *Matrix4) Premultiply(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: m.Call("premultiply", m)}
}
func (m *Matrix4) RotateAxis(v js.Value) {
	m.Call("rotateAxis", v)
}
func (m *Matrix4) Scale(v *Vector3) *Matrix4 {
	return &Matrix4{Value: m.Call("scale", v)}
}
func (m *Matrix4) Set(n11 int, n12 int, n13 int, n14 int, n21 int, n22 int, n23 int, n24 int, n31 int, n32 int, n33 int, n34 int, n41 int, n42 int, n43 int, n44 int) *Matrix4 {
	return &Matrix4{Value: m.Call("set", n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44)}
}
func (m *Matrix4) SetPosition(v *Vector3) *Matrix4 {
	return &Matrix4{Value: m.Call("setPosition", v)}
}
func (m *Matrix4) SetRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return &Matrix4{Value: m.Call("setRotationFromQuaternion", q)}
}
func (m *Matrix4) ToArray() []int {
	return []int(m.Call("toArray"))
}
func (m *Matrix4) Transpose() *Matrix4 {
	return &Matrix4{Value: m.Call("transpose")}
}

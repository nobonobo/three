package math

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Quaternion struct {
	js.Value
}

func NewQuaternion(x int, y int, z int, w int) *Quaternion {
	return &Quaternion{Value: get("Quaternion").New(x, y, z, w)}
}
func (q *Quaternion) OnChangeCallback() js.Value {
	return q.Get("onChangeCallback")
}
func (q *Quaternion) SetOnChangeCallback(v js.Value) {
	q.Set("onChangeCallback", v)
}
func (q *Quaternion) W() int {
	return q.Get("w").Int()
}
func (q *Quaternion) SetW(v int) {
	q.Set("w", v)
}
func (q *Quaternion) X() int {
	return q.Get("x").Int()
}
func (q *Quaternion) SetX(v int) {
	q.Set("x", v)
}
func (q *Quaternion) Y() int {
	return q.Get("y").Int()
}
func (q *Quaternion) SetY(v int) {
	q.Set("y", v)
}
func (q *Quaternion) Z() int {
	return q.Get("z").Int()
}
func (q *Quaternion) SetZ(v int) {
	q.Set("z", v)
}
func (q *Quaternion) AngleTo(q *Quaternion) int {
	return q.Call("angleTo", q).Int()
}
func (q *Quaternion) Clone() this {
	return this(q.Call("clone"))
}
func (q *Quaternion) Conjugate() *Quaternion {
	return &Quaternion{Value: q.Call("conjugate")}
}
func (q *Quaternion) Copy(q *Quaternion) this {
	return this(q.Call("copy", q))
}
func (q *Quaternion) Dot(v *Quaternion) int {
	return q.Call("dot", v).Int()
}
func (q *Quaternion) Equals(v *Quaternion) bool {
	return q.Call("equals", v).Bool()
}
func (q *Quaternion) FromArray(n []int) *Quaternion {
	return &Quaternion{Value: q.Call("fromArray", n)}
}
func (q *Quaternion) FromArray2(xyzw []int, offset int) *Quaternion {
	return &Quaternion{Value: q.Call("fromArray", xyzw, offset)}
}
func (q *Quaternion) Inverse() *Quaternion {
	return &Quaternion{Value: q.Call("inverse")}
}
func (q *Quaternion) Length() int {
	return q.Call("length").Int()
}
func (q *Quaternion) LengthSq() int {
	return q.Call("lengthSq").Int()
}
func (q *Quaternion) Multiply(q *Quaternion) *Quaternion {
	return &Quaternion{Value: q.Call("multiply", q)}
}
func (q *Quaternion) MultiplyQuaternions(a *Quaternion, b *Quaternion) *Quaternion {
	return &Quaternion{Value: q.Call("multiplyQuaternions", a, b)}
}
func (q *Quaternion) MultiplyVector3(v js.Value) js.Value {
	return q.Call("multiplyVector3", v)
}
func (q *Quaternion) Normalize() *Quaternion {
	return &Quaternion{Value: q.Call("normalize")}
}
func (q *Quaternion) OnChange(callback js.Value) *Quaternion {
	return &Quaternion{Value: q.Call("onChange", callback)}
}
func (q *Quaternion) Premultiply(q *Quaternion) *Quaternion {
	return &Quaternion{Value: q.Call("premultiply", q)}
}
func (q *Quaternion) RotateTowards(q *Quaternion, step int) *Quaternion {
	return &Quaternion{Value: q.Call("rotateTowards", q, step)}
}
func (q *Quaternion) Set(x int, y int, z int, w int) *Quaternion {
	return &Quaternion{Value: q.Call("set", x, y, z, w)}
}
func (q *Quaternion) SetFromAxisAngle(axis *Vector3, angle int) *Quaternion {
	return &Quaternion{Value: q.Call("setFromAxisAngle", axis, angle)}
}
func (q *Quaternion) SetFromEuler(euler *Euler, update bool) *Quaternion {
	return &Quaternion{Value: q.Call("setFromEuler", euler, update)}
}
func (q *Quaternion) SetFromRotationMatrix(m *Matrix4) *Quaternion {
	return &Quaternion{Value: q.Call("setFromRotationMatrix", m)}
}
func (q *Quaternion) SetFromUnitVectors(vFrom *Vector3, vTo *Vector3) *Quaternion {
	return &Quaternion{Value: q.Call("setFromUnitVectors", vFrom, vTo)}
}
func (q *Quaternion) Slerp(qb *Quaternion, t int) *Quaternion {
	return &Quaternion{Value: q.Call("slerp", qb, t)}
}
func (q *Quaternion) ToArray() []int {
	return []int(q.Call("toArray"))
}
func (q *Quaternion) ToArray2(xyzw []int, offset int) []int {
	return []int(q.Call("toArray", xyzw, offset))
}
func (q *Quaternion) Slerp2(qa *Quaternion, qb *Quaternion, qm *Quaternion, t int) *Quaternion {
	return &Quaternion{Value: q.Call("slerp", qa, qb, qm, t)}
}
func (q *Quaternion) SlerpFlat(dst []int, dstOffset int, src0 []int, srcOffset int, src1 []int, stcOffset1 int, t int) *Quaternion {
	return &Quaternion{Value: q.Call("slerpFlat", dst, dstOffset, src0, srcOffset, src1, stcOffset1, t)}
}

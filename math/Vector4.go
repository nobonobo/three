package math

import (
	"github.com/nobonobo/three/core"
)

type Vector4 struct {
	js.Value
}

func NewVector4(x int, y int, z int, w int) *Vector4 {
	return &Vector4{Value: get("Vector4").New(x, y, z, w)}
}
func (v *Vector4) IsVector4() true {
	return true(v.Get("isVector4"))
}
func (v *Vector4) SetIsVector4(v true) {
	v.Set("isVector4", v)
}
func (v *Vector4) W() int {
	return v.Get("w").Int()
}
func (v *Vector4) SetW(v int) {
	v.Set("w", v)
}
func (v *Vector4) X() int {
	return v.Get("x").Int()
}
func (v *Vector4) SetX(v int) {
	v.Set("x", v)
}
func (v *Vector4) Y() int {
	return v.Get("y").Int()
}
func (v *Vector4) SetY(v int) {
	v.Set("y", v)
}
func (v *Vector4) Z() int {
	return v.Get("z").Int()
}
func (v *Vector4) SetZ(v int) {
	v.Set("z", v)
}
func (v *Vector4) Add(v *Vector4, w *Vector4) this {
	return this(v.Call("add", v, w))
}
func (v *Vector4) AddScalar(scalar int) this {
	return this(v.Call("addScalar", scalar))
}
func (v *Vector4) AddScaledVector(v *Vector4, s int) this {
	return this(v.Call("addScaledVector", v, s))
}
func (v *Vector4) AddVectors(a *Vector4, b *Vector4) this {
	return this(v.Call("addVectors", a, b))
}
func (v *Vector4) ApplyMatrix4(m *Matrix4) this {
	return this(v.Call("applyMatrix4", m))
}
func (v *Vector4) Ceil() this {
	return this(v.Call("ceil"))
}
func (v *Vector4) Clamp(min *Vector4, max *Vector4) this {
	return this(v.Call("clamp", min, max))
}
func (v *Vector4) ClampScalar(min int, max int) this {
	return this(v.Call("clampScalar", min, max))
}
func (v *Vector4) Clone() this {
	return this(v.Call("clone"))
}
func (v *Vector4) Copy(v *Vector4) this {
	return this(v.Call("copy", v))
}
func (v *Vector4) DivideScalar(s int) this {
	return this(v.Call("divideScalar", s))
}
func (v *Vector4) Dot(v *Vector4) int {
	return v.Call("dot", v).Int()
}
func (v *Vector4) Equals(v *Vector4) bool {
	return v.Call("equals", v).Bool()
}
func (v *Vector4) Floor() this {
	return this(v.Call("floor"))
}
func (v *Vector4) FromArray(xyzw []int, offset int) this {
	return this(v.Call("fromArray", xyzw, offset))
}
func (v *Vector4) FromBufferAttribute(attribute *core.BufferAttribute, index int, offset int) this {
	return this(v.Call("fromBufferAttribute", attribute, index, offset))
}
func (v *Vector4) GetComponent(index int) int {
	return v.Call("getComponent", index).Int()
}
func (v *Vector4) Length() int {
	return v.Call("length").Int()
}
func (v *Vector4) LengthSq() int {
	return v.Call("lengthSq").Int()
}
func (v *Vector4) Lerp(v *Vector4, alpha int) this {
	return this(v.Call("lerp", v, alpha))
}
func (v *Vector4) LerpVectors(v1 *Vector4, v2 *Vector4, alpha int) this {
	return this(v.Call("lerpVectors", v1, v2, alpha))
}
func (v *Vector4) ManhattanLength() int {
	return v.Call("manhattanLength").Int()
}
func (v *Vector4) Max(v *Vector4) this {
	return this(v.Call("max", v))
}
func (v *Vector4) Min(v *Vector4) this {
	return this(v.Call("min", v))
}
func (v *Vector4) MultiplyScalar(s int) this {
	return this(v.Call("multiplyScalar", s))
}
func (v *Vector4) Negate() this {
	return this(v.Call("negate"))
}
func (v *Vector4) Normalize() this {
	return this(v.Call("normalize"))
}
func (v *Vector4) Round() this {
	return this(v.Call("round"))
}
func (v *Vector4) RoundToZero() this {
	return this(v.Call("roundToZero"))
}
func (v *Vector4) Set(x int, y int, z int, w int) this {
	return this(v.Call("set", x, y, z, w))
}
func (v *Vector4) SetAxisAngleFromQuaternion(q *Quaternion) this {
	return this(v.Call("setAxisAngleFromQuaternion", q))
}
func (v *Vector4) SetAxisAngleFromRotationMatrix(m *Matrix3) this {
	return this(v.Call("setAxisAngleFromRotationMatrix", m))
}
func (v *Vector4) SetComponent(index int, value int) this {
	return this(v.Call("setComponent", index, value))
}
func (v *Vector4) SetLength(length int) this {
	return this(v.Call("setLength", length))
}
func (v *Vector4) SetScalar(scalar int) this {
	return this(v.Call("setScalar", scalar))
}
func (v *Vector4) SetW(w int) this {
	return this(v.Call("setW", w))
}
func (v *Vector4) SetX(x int) this {
	return this(v.Call("setX", x))
}
func (v *Vector4) SetY(y int) this {
	return this(v.Call("setY", y))
}
func (v *Vector4) SetZ(z int) this {
	return this(v.Call("setZ", z))
}
func (v *Vector4) Sub(v *Vector4) this {
	return this(v.Call("sub", v))
}
func (v *Vector4) SubScalar(s int) this {
	return this(v.Call("subScalar", s))
}
func (v *Vector4) SubVectors(a *Vector4, b *Vector4) this {
	return this(v.Call("subVectors", a, b))
}
func (v *Vector4) ToArray(xyzw []int, offset int) []int {
	return []int(v.Call("toArray", xyzw, offset))
}

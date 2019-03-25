package math

import (
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
)

type Vector3 struct {
	js.Value
}

func NewVector3(x int, y int, z int) *Vector3 {
	return &Vector3{Value: get("Vector3").New(x, y, z)}
}
func (v *Vector3) IsVector3() true {
	return true(v.Get("isVector3"))
}
func (v *Vector3) SetIsVector3(v true) {
	v.Set("isVector3", v)
}
func (v *Vector3) X() int {
	return v.Get("x").Int()
}
func (v *Vector3) SetX(v int) {
	v.Set("x", v)
}
func (v *Vector3) Y() int {
	return v.Get("y").Int()
}
func (v *Vector3) SetY(v int) {
	v.Set("y", v)
}
func (v *Vector3) Z() int {
	return v.Get("z").Int()
}
func (v *Vector3) SetZ(v int) {
	v.Set("z", v)
}
func (v *Vector3) Add(a *Vector3, b *Vector3) this {
	return this(v.Call("add", a, b))
}
func (v *Vector3) AddScalar(s int) this {
	return this(v.Call("addScalar", s))
}
func (v *Vector3) AddScaledVector(v *Vector3, s int) this {
	return this(v.Call("addScaledVector", v, s))
}
func (v *Vector3) AddVectors(a *Vector3, b *Vector3) this {
	return this(v.Call("addVectors", a, b))
}
func (v *Vector3) AngleTo(v *Vector3) int {
	return v.Call("angleTo", v).Int()
}
func (v *Vector3) ApplyAxisAngle(axis *Vector3, angle int) this {
	return this(v.Call("applyAxisAngle", axis, angle))
}
func (v *Vector3) ApplyEuler(euler *Euler) this {
	return this(v.Call("applyEuler", euler))
}
func (v *Vector3) ApplyMatrix3(m *Matrix3) this {
	return this(v.Call("applyMatrix3", m))
}
func (v *Vector3) ApplyMatrix4(m *Matrix4) this {
	return this(v.Call("applyMatrix4", m))
}
func (v *Vector3) ApplyQuaternion(q *Quaternion) this {
	return this(v.Call("applyQuaternion", q))
}
func (v *Vector3) Ceil() this {
	return this(v.Call("ceil"))
}
func (v *Vector3) Clamp(min *Vector3, max *Vector3) this {
	return this(v.Call("clamp", min, max))
}
func (v *Vector3) ClampLength(min int, max int) this {
	return this(v.Call("clampLength", min, max))
}
func (v *Vector3) ClampScalar(min int, max int) this {
	return this(v.Call("clampScalar", min, max))
}
func (v *Vector3) Clone() this {
	return this(v.Call("clone"))
}
func (v *Vector3) Copy(v *Vector3) this {
	return this(v.Call("copy", v))
}
func (v *Vector3) Cross(a *Vector3, w *Vector3) this {
	return this(v.Call("cross", a, w))
}
func (v *Vector3) CrossVectors(a *Vector3, b *Vector3) this {
	return this(v.Call("crossVectors", a, b))
}
func (v *Vector3) DistanceTo(v *Vector3) int {
	return v.Call("distanceTo", v).Int()
}
func (v *Vector3) DistanceToManhattan(v *Vector3) int {
	return v.Call("distanceToManhattan", v).Int()
}
func (v *Vector3) DistanceToSquared(v *Vector3) int {
	return v.Call("distanceToSquared", v).Int()
}
func (v *Vector3) Divide(v *Vector3) this {
	return this(v.Call("divide", v))
}
func (v *Vector3) DivideScalar(s int) this {
	return this(v.Call("divideScalar", s))
}
func (v *Vector3) Dot(v *Vector3) int {
	return v.Call("dot", v).Int()
}
func (v *Vector3) Equals(v *Vector3) bool {
	return v.Call("equals", v).Bool()
}
func (v *Vector3) Floor() this {
	return this(v.Call("floor"))
}
func (v *Vector3) FromArray(xyz []int, offset int) *Vector3 {
	return &Vector3{Value: v.Call("fromArray", xyz, offset)}
}
func (v *Vector3) FromBufferAttribute(attribute *core.BufferAttribute, index int, offset int) this {
	return this(v.Call("fromBufferAttribute", attribute, index, offset))
}
func (v *Vector3) GetComponent(index int) int {
	return v.Call("getComponent", index).Int()
}
func (v *Vector3) Length() int {
	return v.Call("length").Int()
}
func (v *Vector3) LengthManhattan() int {
	return v.Call("lengthManhattan").Int()
}
func (v *Vector3) LengthSq() int {
	return v.Call("lengthSq").Int()
}
func (v *Vector3) Lerp(v *Vector3, alpha int) this {
	return this(v.Call("lerp", v, alpha))
}
func (v *Vector3) LerpVectors(v1 *Vector3, v2 *Vector3, alpha int) this {
	return this(v.Call("lerpVectors", v1, v2, alpha))
}
func (v *Vector3) ManhattanDistanceTo(v *Vector3) int {
	return v.Call("manhattanDistanceTo", v).Int()
}
func (v *Vector3) ManhattanLength() int {
	return v.Call("manhattanLength").Int()
}
func (v *Vector3) Max(v *Vector3) this {
	return this(v.Call("max", v))
}
func (v *Vector3) Min(v *Vector3) this {
	return this(v.Call("min", v))
}
func (v *Vector3) Multiply(v *Vector3) this {
	return this(v.Call("multiply", v))
}
func (v *Vector3) MultiplyScalar(s int) this {
	return this(v.Call("multiplyScalar", s))
}
func (v *Vector3) MultiplyVectors(a *Vector3, b *Vector3) this {
	return this(v.Call("multiplyVectors", a, b))
}
func (v *Vector3) Negate() this {
	return this(v.Call("negate"))
}
func (v *Vector3) Normalize() this {
	return this(v.Call("normalize"))
}
func (v *Vector3) Project(camera *cameras.Camera) this {
	return this(v.Call("project", camera))
}
func (v *Vector3) ProjectOnPlane(planeNormal *Vector3) this {
	return this(v.Call("projectOnPlane", planeNormal))
}
func (v *Vector3) ProjectOnVector(v *Vector3) this {
	return this(v.Call("projectOnVector", v))
}
func (v *Vector3) Reflect(vector *Vector3) this {
	return this(v.Call("reflect", vector))
}
func (v *Vector3) Round() this {
	return this(v.Call("round"))
}
func (v *Vector3) RoundToZero() this {
	return this(v.Call("roundToZero"))
}
func (v *Vector3) Set(x int, y int, z int) this {
	return this(v.Call("set", x, y, z))
}
func (v *Vector3) SetComponent(index int, value int) this {
	return this(v.Call("setComponent", index, value))
}
func (v *Vector3) SetFromCylindrical(s *Cylindrical) this {
	return this(v.Call("setFromCylindrical", s))
}
func (v *Vector3) SetFromMatrixColumn(matrix *Matrix4, index int) this {
	return this(v.Call("setFromMatrixColumn", matrix, index))
}
func (v *Vector3) SetFromMatrixPosition(m *Matrix4) this {
	return this(v.Call("setFromMatrixPosition", m))
}
func (v *Vector3) SetFromMatrixScale(m *Matrix4) this {
	return this(v.Call("setFromMatrixScale", m))
}
func (v *Vector3) SetFromSpherical(s *Spherical) this {
	return this(v.Call("setFromSpherical", s))
}
func (v *Vector3) SetLength(l int) this {
	return this(v.Call("setLength", l))
}
func (v *Vector3) SetScalar(scalar int) this {
	return this(v.Call("setScalar", scalar))
}
func (v *Vector3) SetX(x int) *Vector3 {
	return &Vector3{Value: v.Call("setX", x)}
}
func (v *Vector3) SetY(y int) *Vector3 {
	return &Vector3{Value: v.Call("setY", y)}
}
func (v *Vector3) SetZ(z int) *Vector3 {
	return &Vector3{Value: v.Call("setZ", z)}
}
func (v *Vector3) Sub(a *Vector3) this {
	return this(v.Call("sub", a))
}
func (v *Vector3) SubScalar(s int) this {
	return this(v.Call("subScalar", s))
}
func (v *Vector3) SubVectors(a *Vector3, b *Vector3) this {
	return this(v.Call("subVectors", a, b))
}
func (v *Vector3) ToArray(xyz []int, offset int) []int {
	return []int(v.Call("toArray", xyz, offset))
}
func (v *Vector3) ToArray2(xyz *ArrayLike, offset int) *ArrayLike {
	return &ArrayLike{Value: v.Call("toArray", xyz, offset)}
}
func (v *Vector3) TransformDirection(m *Matrix4) this {
	return this(v.Call("transformDirection", m))
}
func (v *Vector3) Unproject(camera *cameras.Camera) this {
	return this(v.Call("unproject", camera))
}

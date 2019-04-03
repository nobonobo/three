// Code generated from "math/Vector3.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Vector3 struct {
	js.Value
}

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{Value: get("Vector3").New(x, y, z)}
}
func (vv *Vector3) IsVector3() bool {
	return vv.Get("isVector3").Bool()
}
func (vv *Vector3) SetIsVector3(v bool) {
	vv.Set("isVector3", v)
}
func (vv *Vector3) X() float64 {
	return vv.Get("x").Float()
}
func (vv *Vector3) SetX(v float64) {
	vv.Set("x", v)
}
func (vv *Vector3) Y() float64 {
	return vv.Get("y").Float()
}
func (vv *Vector3) SetY(v float64) {
	vv.Set("y", v)
}
func (vv *Vector3) Z() float64 {
	return vv.Get("z").Float()
}
func (vv *Vector3) SetZ(v float64) {
	vv.Set("z", v)
}
func (vv *Vector3) Add(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("add", a, b)}
}
func (vv *Vector3) AddScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("addScalar", s)}
}
func (vv *Vector3) AddScaledVector(v *Vector3, s float64) *Vector3 {
	return &Vector3{Value: vv.Call("addScaledVector", v, s)}
}
func (vv *Vector3) AddVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("addVectors", a, b)}
}
func (vv *Vector3) AngleTo(v *Vector3) float64 {
	return vv.Call("angleTo", v).Float()
}
func (vv *Vector3) ApplyAxisAngle(axis *Vector3, angle float64) *Vector3 {
	return &Vector3{Value: vv.Call("applyAxisAngle", axis, angle)}
}
func (vv *Vector3) ApplyEuler(euler *Euler) *Vector3 {
	return &Vector3{Value: vv.Call("applyEuler", euler)}
}
func (vv *Vector3) ApplyMatrix3(m *Matrix3) *Vector3 {
	return &Vector3{Value: vv.Call("applyMatrix3", m)}
}
func (vv *Vector3) ApplyMatrix4(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("applyMatrix4", m)}
}
func (vv *Vector3) ApplyQuaternion(q *Quaternion) *Vector3 {
	return &Vector3{Value: vv.Call("applyQuaternion", q)}
}
func (vv *Vector3) Ceil() *Vector3 {
	return &Vector3{Value: vv.Call("ceil")}
}
func (vv *Vector3) Clamp(min *Vector3, max *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("clamp", min, max)}
}
func (vv *Vector3) ClampLength(min float64, max float64) *Vector3 {
	return &Vector3{Value: vv.Call("clampLength", min, max)}
}
func (vv *Vector3) ClampScalar(min float64, max float64) *Vector3 {
	return &Vector3{Value: vv.Call("clampScalar", min, max)}
}
func (vv *Vector3) Clone() *Vector3 {
	return &Vector3{Value: vv.Call("clone")}
}
func (vv *Vector3) Copy(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("copy", v)}
}
func (vv *Vector3) Cross(a *Vector3, w *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("cross", a, w)}
}
func (vv *Vector3) CrossVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("crossVectors", a, b)}
}
func (vv *Vector3) DistanceTo(v *Vector3) float64 {
	return vv.Call("distanceTo", v).Float()
}
func (vv *Vector3) DistanceToManhattan(v *Vector3) float64 {
	return vv.Call("distanceToManhattan", v).Float()
}
func (vv *Vector3) DistanceToSquared(v *Vector3) float64 {
	return vv.Call("distanceToSquared", v).Float()
}
func (vv *Vector3) Divide(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("divide", v)}
}
func (vv *Vector3) DivideScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("divideScalar", s)}
}
func (vv *Vector3) Dot(v *Vector3) float64 {
	return vv.Call("dot", v).Float()
}
func (vv *Vector3) Equals(v *Vector3) bool {
	return vv.Call("equals", v).Bool()
}
func (vv *Vector3) Floor() *Vector3 {
	return &Vector3{Value: vv.Call("floor")}
}
func (vv *Vector3) FromArray(xyz js.Value, offset int) *Vector3 {
	return &Vector3{Value: vv.Call("fromArray", xyz, offset)}
}
func (vv *Vector3) FromBufferAttribute(attribute *BufferAttribute, index int, offset int) *Vector3 {
	return &Vector3{Value: vv.Call("fromBufferAttribute", attribute, index, offset)}
}
func (vv *Vector3) GetComponent(index int) float64 {
	return vv.Call("getComponent", index).Float()
}
func (vv *Vector3) Length() float64 {
	return vv.Call("length").Float()
}
func (vv *Vector3) LengthManhattan() float64 {
	return vv.Call("lengthManhattan").Float()
}
func (vv *Vector3) LengthSq() float64 {
	return vv.Call("lengthSq").Float()
}
func (vv *Vector3) Lerp(v *Vector3, alpha float64) *Vector3 {
	return &Vector3{Value: vv.Call("lerp", v, alpha)}
}
func (vv *Vector3) LerpVectors(v1 *Vector3, v2 *Vector3, alpha float64) *Vector3 {
	return &Vector3{Value: vv.Call("lerpVectors", v1, v2, alpha)}
}
func (vv *Vector3) ManhattanDistanceTo(v *Vector3) float64 {
	return vv.Call("manhattanDistanceTo", v).Float()
}
func (vv *Vector3) ManhattanLength() float64 {
	return vv.Call("manhattanLength").Float()
}
func (vv *Vector3) Max(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("max", v)}
}
func (vv *Vector3) Min(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("min", v)}
}
func (vv *Vector3) Multiply(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("multiply", v)}
}
func (vv *Vector3) MultiplyScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("multiplyScalar", s)}
}
func (vv *Vector3) MultiplyVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("multiplyVectors", a, b)}
}
func (vv *Vector3) Negate() *Vector3 {
	return &Vector3{Value: vv.Call("negate")}
}
func (vv *Vector3) Normalize() *Vector3 {
	return &Vector3{Value: vv.Call("normalize")}
}
func (vv *Vector3) Project(camera *Camera) *Vector3 {
	return &Vector3{Value: vv.Call("project", camera)}
}
func (vv *Vector3) ProjectOnPlane(planeNormal *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("projectOnPlane", planeNormal)}
}
func (vv *Vector3) ProjectOnVector(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("projectOnVector", v)}
}
func (vv *Vector3) Reflect(vector *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("reflect", vector)}
}
func (vv *Vector3) Round() *Vector3 {
	return &Vector3{Value: vv.Call("round")}
}
func (vv *Vector3) RoundToZero() *Vector3 {
	return &Vector3{Value: vv.Call("roundToZero")}
}
func (vv *Vector3) Set2(x float64, y float64, z float64) *Vector3 {
	return &Vector3{Value: vv.Call("set", x, y, z)}
}
func (vv *Vector3) SetComponent(index int, value float64) *Vector3 {
	return &Vector3{Value: vv.Call("setComponent", index, value)}
}
func (vv *Vector3) SetFromCylindrical(s *Cylindrical) *Vector3 {
	return &Vector3{Value: vv.Call("setFromCylindrical", s)}
}
func (vv *Vector3) SetFromMatrixColumn(matrix *Matrix4, index int) *Vector3 {
	return &Vector3{Value: vv.Call("setFromMatrixColumn", matrix, index)}
}
func (vv *Vector3) SetFromMatrixPosition(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("setFromMatrixPosition", m)}
}
func (vv *Vector3) SetFromMatrixScale(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("setFromMatrixScale", m)}
}
func (vv *Vector3) SetFromSpherical(s *Spherical) *Vector3 {
	return &Vector3{Value: vv.Call("setFromSpherical", s)}
}
func (vv *Vector3) SetLength(l float64) *Vector3 {
	return &Vector3{Value: vv.Call("setLength", l)}
}
func (vv *Vector3) SetScalar(scalar float64) *Vector3 {
	return &Vector3{Value: vv.Call("setScalar", scalar)}
}
func (vv *Vector3) SetX2(x float64) *Vector3 {
	return &Vector3{Value: vv.Call("setX", x)}
}
func (vv *Vector3) SetY2(y float64) *Vector3 {
	return &Vector3{Value: vv.Call("setY", y)}
}
func (vv *Vector3) SetZ2(z float64) *Vector3 {
	return &Vector3{Value: vv.Call("setZ", z)}
}
func (vv *Vector3) Sub(a *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("sub", a)}
}
func (vv *Vector3) SubScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("subScalar", s)}
}
func (vv *Vector3) SubVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("subVectors", a, b)}
}
func (vv *Vector3) ToArray(xyz js.Value, offset int) js.Value {
	return vv.Call("toArray", xyz, offset)
}
func (vv *Vector3) ToArray2(xyz js.Value, offset int) js.Value {
	return vv.Call("toArray", xyz, offset)
}
func (vv *Vector3) TransformDirection(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("transformDirection", m)}
}
func (vv *Vector3) Unproject(camera *Camera) *Vector3 {
	return &Vector3{Value: vv.Call("unproject", camera)}
}

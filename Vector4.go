// Code generated from "math/Vector4.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Vector4 struct {
	js.Value
}

func NewVector4(x float64, y float64, z float64, w float64) *Vector4 {
	return &Vector4{Value: get("Vector4").New(x, y, z, w)}
}
func (vv *Vector4) IsVector4() bool {
	return vv.Get("isVector4").Bool()
}
func (vv *Vector4) SetIsVector4(v bool) {
	vv.Set("isVector4", v)
}
func (vv *Vector4) W() float64 {
	return vv.Get("w").Float()
}
func (vv *Vector4) SetW(v float64) {
	vv.Set("w", v)
}
func (vv *Vector4) X() float64 {
	return vv.Get("x").Float()
}
func (vv *Vector4) SetX(v float64) {
	vv.Set("x", v)
}
func (vv *Vector4) Y() float64 {
	return vv.Get("y").Float()
}
func (vv *Vector4) SetY(v float64) {
	vv.Set("y", v)
}
func (vv *Vector4) Z() float64 {
	return vv.Get("z").Float()
}
func (vv *Vector4) SetZ(v float64) {
	vv.Set("z", v)
}
func (vv *Vector4) Add(v *Vector4, w *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("add", v, w)}
}
func (vv *Vector4) AddScalar(scalar float64) *Vector4 {
	return &Vector4{Value: vv.Call("addScalar", scalar)}
}
func (vv *Vector4) AddScaledVector(v *Vector4, s float64) *Vector4 {
	return &Vector4{Value: vv.Call("addScaledVector", v, s)}
}
func (vv *Vector4) AddVectors(a *Vector4, b *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("addVectors", a, b)}
}
func (vv *Vector4) ApplyMatrix4(m *Matrix4) *Vector4 {
	return &Vector4{Value: vv.Call("applyMatrix4", m)}
}
func (vv *Vector4) Ceil() *Vector4 {
	return &Vector4{Value: vv.Call("ceil")}
}
func (vv *Vector4) Clamp(min *Vector4, max *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("clamp", min, max)}
}
func (vv *Vector4) ClampScalar(min float64, max float64) *Vector4 {
	return &Vector4{Value: vv.Call("clampScalar", min, max)}
}
func (vv *Vector4) Clone() *Vector4 {
	return &Vector4{Value: vv.Call("clone")}
}
func (vv *Vector4) Copy(v *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("copy", v)}
}
func (vv *Vector4) DivideScalar(s float64) *Vector4 {
	return &Vector4{Value: vv.Call("divideScalar", s)}
}
func (vv *Vector4) Dot(v *Vector4) float64 {
	return vv.Call("dot", v).Float()
}
func (vv *Vector4) Equals(v *Vector4) bool {
	return vv.Call("equals", v).Bool()
}
func (vv *Vector4) Floor() *Vector4 {
	return &Vector4{Value: vv.Call("floor")}
}
func (vv *Vector4) FromArray(xyzw js.Value, offset int) *Vector4 {
	return &Vector4{Value: vv.Call("fromArray", xyzw, offset)}
}
func (vv *Vector4) FromBufferAttribute(attribute *BufferAttribute, index int, offset int) *Vector4 {
	return &Vector4{Value: vv.Call("fromBufferAttribute", attribute, index, offset)}
}
func (vv *Vector4) GetComponent(index int) float64 {
	return vv.Call("getComponent", index).Float()
}
func (vv *Vector4) Length() float64 {
	return vv.Call("length").Float()
}
func (vv *Vector4) LengthSq() float64 {
	return vv.Call("lengthSq").Float()
}
func (vv *Vector4) Lerp(v *Vector4, alpha float64) *Vector4 {
	return &Vector4{Value: vv.Call("lerp", v, alpha)}
}
func (vv *Vector4) LerpVectors(v1 *Vector4, v2 *Vector4, alpha float64) *Vector4 {
	return &Vector4{Value: vv.Call("lerpVectors", v1, v2, alpha)}
}
func (vv *Vector4) ManhattanLength() float64 {
	return vv.Call("manhattanLength").Float()
}
func (vv *Vector4) Max(v *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("max", v)}
}
func (vv *Vector4) Min(v *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("min", v)}
}
func (vv *Vector4) MultiplyScalar(s float64) *Vector4 {
	return &Vector4{Value: vv.Call("multiplyScalar", s)}
}
func (vv *Vector4) Negate() *Vector4 {
	return &Vector4{Value: vv.Call("negate")}
}
func (vv *Vector4) Normalize() *Vector4 {
	return &Vector4{Value: vv.Call("normalize")}
}
func (vv *Vector4) Round() *Vector4 {
	return &Vector4{Value: vv.Call("round")}
}
func (vv *Vector4) RoundToZero() *Vector4 {
	return &Vector4{Value: vv.Call("roundToZero")}
}
func (vv *Vector4) Set2(x float64, y float64, z float64, w float64) *Vector4 {
	return &Vector4{Value: vv.Call("set", x, y, z, w)}
}
func (vv *Vector4) SetAxisAngleFromQuaternion(q *Quaternion) *Vector4 {
	return &Vector4{Value: vv.Call("setAxisAngleFromQuaternion", q)}
}
func (vv *Vector4) SetAxisAngleFromRotationMatrix(m *Matrix3) *Vector4 {
	return &Vector4{Value: vv.Call("setAxisAngleFromRotationMatrix", m)}
}
func (vv *Vector4) SetComponent(index int, value float64) *Vector4 {
	return &Vector4{Value: vv.Call("setComponent", index, value)}
}
func (vv *Vector4) SetLength(length float64) *Vector4 {
	return &Vector4{Value: vv.Call("setLength", length)}
}
func (vv *Vector4) SetScalar(scalar float64) *Vector4 {
	return &Vector4{Value: vv.Call("setScalar", scalar)}
}
func (vv *Vector4) SetW2(w float64) *Vector4 {
	return &Vector4{Value: vv.Call("setW", w)}
}
func (vv *Vector4) SetX2(x float64) *Vector4 {
	return &Vector4{Value: vv.Call("setX", x)}
}
func (vv *Vector4) SetY2(y float64) *Vector4 {
	return &Vector4{Value: vv.Call("setY", y)}
}
func (vv *Vector4) SetZ2(z float64) *Vector4 {
	return &Vector4{Value: vv.Call("setZ", z)}
}
func (vv *Vector4) Sub(v *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("sub", v)}
}
func (vv *Vector4) SubScalar(s float64) *Vector4 {
	return &Vector4{Value: vv.Call("subScalar", s)}
}
func (vv *Vector4) SubVectors(a *Vector4, b *Vector4) *Vector4 {
	return &Vector4{Value: vv.Call("subVectors", a, b)}
}
func (vv *Vector4) ToArray(xyzw js.Value, offset int) js.Value {
	return vv.Call("toArray", xyzw, offset)
}

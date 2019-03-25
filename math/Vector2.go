package math

import (
	"github.com/nobonobo/three/core"
)

type Vector interface {
	Add(v Vector, w Vector) this
	AddScalar(scalar int) this
	AddScaledVector(vector Vector, scale int) this
	AddVectors(a Vector, b Vector) this
	Clone() this
	Copy(v Vector) this
	DistanceTo(v Vector) int
	DistanceToSquared(v Vector) int
	DivideScalar(s int) this
	Dot(v Vector) int
	Equals(v Vector) bool
	GetComponent(index int) int
	Length() int
	LengthSq() int
	Lerp(v Vector, alpha int) this
	MultiplyScalar(s int) this
	Negate() this
	Normalize() this
	Set(args []int) this
	SetComponent(index int, value int) this
	SetLength(l int) this
	SetScalar(scalar int) this
	Sub(v Vector) this
	SubVectors(a Vector, b Vector) this
}
type Vector2 struct {
	js.Value
}

func NewVector2(x int, y int) *Vector2 {
	return &Vector2{Value: get("Vector2").New(x, y)}
}
func (v *Vector2) Height() int {
	return v.Get("height").Int()
}
func (v *Vector2) SetHeight(v int) {
	v.Set("height", v)
}
func (v *Vector2) IsVector2() true {
	return true(v.Get("isVector2"))
}
func (v *Vector2) SetIsVector2(v true) {
	v.Set("isVector2", v)
}
func (v *Vector2) Width() int {
	return v.Get("width").Int()
}
func (v *Vector2) SetWidth(v int) {
	v.Set("width", v)
}
func (v *Vector2) X() int {
	return v.Get("x").Int()
}
func (v *Vector2) SetX(v int) {
	v.Set("x", v)
}
func (v *Vector2) Y() int {
	return v.Get("y").Int()
}
func (v *Vector2) SetY(v int) {
	v.Set("y", v)
}
func (v *Vector2) Add(v *Vector2, w *Vector2) this {
	return this(v.Call("add", v, w))
}
func (v *Vector2) AddScalar(s int) this {
	return this(v.Call("addScalar", s))
}
func (v *Vector2) AddScaledVector(v *Vector2, s int) this {
	return this(v.Call("addScaledVector", v, s))
}
func (v *Vector2) AddVectors(a *Vector2, b *Vector2) this {
	return this(v.Call("addVectors", a, b))
}
func (v *Vector2) Angle() int {
	return v.Call("angle").Int()
}
func (v *Vector2) ApplyMatrix3(m *Matrix3) this {
	return this(v.Call("applyMatrix3", m))
}
func (v *Vector2) Ceil() this {
	return this(v.Call("ceil"))
}
func (v *Vector2) Clamp(min *Vector2, max *Vector2) this {
	return this(v.Call("clamp", min, max))
}
func (v *Vector2) ClampLength(min int, max int) this {
	return this(v.Call("clampLength", min, max))
}
func (v *Vector2) ClampScalar(min int, max int) this {
	return this(v.Call("clampScalar", min, max))
}
func (v *Vector2) Clone() this {
	return this(v.Call("clone"))
}
func (v *Vector2) Copy(v *Vector2) this {
	return this(v.Call("copy", v))
}
func (v *Vector2) DistanceTo(v *Vector2) int {
	return v.Call("distanceTo", v).Int()
}
func (v *Vector2) DistanceToManhattan(v *Vector2) int {
	return v.Call("distanceToManhattan", v).Int()
}
func (v *Vector2) DistanceToSquared(v *Vector2) int {
	return v.Call("distanceToSquared", v).Int()
}
func (v *Vector2) Divide(v *Vector2) this {
	return this(v.Call("divide", v))
}
func (v *Vector2) DivideScalar(s int) this {
	return this(v.Call("divideScalar", s))
}
func (v *Vector2) Dot(v *Vector2) int {
	return v.Call("dot", v).Int()
}
func (v *Vector2) Equals(v *Vector2) bool {
	return v.Call("equals", v).Bool()
}
func (v *Vector2) Floor() this {
	return this(v.Call("floor"))
}
func (v *Vector2) FromArray(array []int, offset int) this {
	return this(v.Call("fromArray", array, offset))
}
func (v *Vector2) FromBufferAttribute(attribute *core.BufferAttribute, index int) this {
	return this(v.Call("fromBufferAttribute", attribute, index))
}
func (v *Vector2) GetComponent(index int) int {
	return v.Call("getComponent", index).Int()
}
func (v *Vector2) Length() int {
	return v.Call("length").Int()
}
func (v *Vector2) LengthManhattan() int {
	return v.Call("lengthManhattan").Int()
}
func (v *Vector2) LengthSq() int {
	return v.Call("lengthSq").Int()
}
func (v *Vector2) Lerp(v *Vector2, alpha int) this {
	return this(v.Call("lerp", v, alpha))
}
func (v *Vector2) LerpVectors(v1 *Vector2, v2 *Vector2, alpha int) this {
	return this(v.Call("lerpVectors", v1, v2, alpha))
}
func (v *Vector2) ManhattanDistanceTo(v *Vector2) int {
	return v.Call("manhattanDistanceTo", v).Int()
}
func (v *Vector2) ManhattanDistanceTo2(v *Vector2) int {
	return v.Call("manhattanDistanceTo", v).Int()
}
func (v *Vector2) ManhattanLength() int {
	return v.Call("manhattanLength").Int()
}
func (v *Vector2) ManhattanLength2() int {
	return v.Call("manhattanLength").Int()
}
func (v *Vector2) Max(v *Vector2) this {
	return this(v.Call("max", v))
}
func (v *Vector2) Min(v *Vector2) this {
	return this(v.Call("min", v))
}
func (v *Vector2) Multiply(v *Vector2) this {
	return this(v.Call("multiply", v))
}
func (v *Vector2) MultiplyScalar(scalar int) this {
	return this(v.Call("multiplyScalar", scalar))
}
func (v *Vector2) Negate() this {
	return this(v.Call("negate"))
}
func (v *Vector2) Normalize() this {
	return this(v.Call("normalize"))
}
func (v *Vector2) RotateAround(center *Vector2, angle int) this {
	return this(v.Call("rotateAround", center, angle))
}
func (v *Vector2) Round() this {
	return this(v.Call("round"))
}
func (v *Vector2) RoundToZero() this {
	return this(v.Call("roundToZero"))
}
func (v *Vector2) Set(x int, y int) this {
	return this(v.Call("set", x, y))
}
func (v *Vector2) SetComponent(index int, value int) this {
	return this(v.Call("setComponent", index, value))
}
func (v *Vector2) SetLength(length int) this {
	return this(v.Call("setLength", length))
}
func (v *Vector2) SetScalar(scalar int) this {
	return this(v.Call("setScalar", scalar))
}
func (v *Vector2) SetX(x int) this {
	return this(v.Call("setX", x))
}
func (v *Vector2) SetY(y int) this {
	return this(v.Call("setY", y))
}
func (v *Vector2) Sub(v *Vector2) this {
	return this(v.Call("sub", v))
}
func (v *Vector2) SubScalar(s int) this {
	return this(v.Call("subScalar", s))
}
func (v *Vector2) SubVectors(a *Vector2, b *Vector2) this {
	return this(v.Call("subVectors", a, b))
}
func (v *Vector2) ToArray(array []int, offset int) []int {
	return []int(v.Call("toArray", array, offset))
}
func (v *Vector2) ToArray2(array *ArrayLike, offset int) *ArrayLike {
	return &ArrayLike{Value: v.Call("toArray", array, offset)}
}

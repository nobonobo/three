package math

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Euler struct {
	js.Value
}

func NewEuler(x int, y int, z int, order string) *Euler {
	return &Euler{Value: get("Euler").New(x, y, z, order)}
}
func (e *Euler) OnChangeCallback() js.Value {
	return e.Get("onChangeCallback")
}
func (e *Euler) SetOnChangeCallback(v js.Value) {
	e.Set("onChangeCallback", v)
}
func (e *Euler) Order() string {
	return e.Get("order").String()
}
func (e *Euler) SetOrder(v string) {
	e.Set("order", v)
}
func (e *Euler) X() int {
	return e.Get("x").Int()
}
func (e *Euler) SetX(v int) {
	e.Set("x", v)
}
func (e *Euler) Y() int {
	return e.Get("y").Int()
}
func (e *Euler) SetY(v int) {
	e.Set("y", v)
}
func (e *Euler) Z() int {
	return e.Get("z").Int()
}
func (e *Euler) SetZ(v int) {
	e.Set("z", v)
}
func (e *Euler) DefaultOrder() string {
	return e.Get("DefaultOrder").String()
}
func (e *Euler) SetDefaultOrder(v string) {
	e.Set("DefaultOrder", v)
}
func (e *Euler) RotationOrders() []string {
	return []string(e.Get("RotationOrders"))
}
func (e *Euler) SetRotationOrders(v []string) {
	e.Set("RotationOrders", v)
}
func (e *Euler) Clone() this {
	return this(e.Call("clone"))
}
func (e *Euler) Copy(euler *Euler) this {
	return this(e.Call("copy", euler))
}
func (e *Euler) Equals(euler *Euler) bool {
	return e.Call("equals", euler).Bool()
}
func (e *Euler) FromArray(xyzo []js.Value) *Euler {
	return &Euler{Value: e.Call("fromArray", xyzo)}
}
func (e *Euler) OnChange(callback js.Value) this {
	return this(e.Call("onChange", callback))
}
func (e *Euler) Reorder(newOrder string) *Euler {
	return &Euler{Value: e.Call("reorder", newOrder)}
}
func (e *Euler) Set(x int, y int, z int, order string) *Euler {
	return &Euler{Value: e.Call("set", x, y, z, order)}
}
func (e *Euler) SetFromQuaternion(q *Quaternion, order string, update bool) *Euler {
	return &Euler{Value: e.Call("setFromQuaternion", q, order, update)}
}
func (e *Euler) SetFromRotationMatrix(m *Matrix4, order string, update bool) *Euler {
	return &Euler{Value: e.Call("setFromRotationMatrix", m, order, update)}
}
func (e *Euler) SetFromVector3(v *Vector3, order string) *Euler {
	return &Euler{Value: e.Call("setFromVector3", v, order)}
}
func (e *Euler) ToArray(array []int, offset int) []int {
	return []int(e.Call("toArray", array, offset))
}
func (e *Euler) ToVector3(optionalResult *Vector3) *Vector3 {
	return &Vector3{Value: e.Call("toVector3", optionalResult)}
}

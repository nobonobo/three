// Code generated from "math/Euler.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Euler struct {
	js.Value
}

func NewEuler(x float64, y float64, z float64, order string) *Euler {
	return &Euler{Value: get("Euler").New(x, y, z, order)}
}
func (ee *Euler) OnChangeCallback() js.Value {
	return ee.Get("onChangeCallback")
}
func (ee *Euler) SetOnChangeCallback(v js.Value) {
	ee.Set("onChangeCallback", v)
}
func (ee *Euler) Order() string {
	return ee.Get("order").String()
}
func (ee *Euler) SetOrder(v string) {
	ee.Set("order", v)
}
func (ee *Euler) X() float64 {
	return ee.Get("x").Float()
}
func (ee *Euler) SetX(v float64) {
	ee.Set("x", v)
}
func (ee *Euler) Y() float64 {
	return ee.Get("y").Float()
}
func (ee *Euler) SetY(v float64) {
	ee.Set("y", v)
}
func (ee *Euler) Z() float64 {
	return ee.Get("z").Float()
}
func (ee *Euler) SetZ(v float64) {
	ee.Set("z", v)
}
func (ee *Euler) DefaultOrder() string {
	return ee.Get("DefaultOrder").String()
}
func (ee *Euler) SetDefaultOrder(v string) {
	ee.Set("DefaultOrder", v)
}
func (ee *Euler) RotationOrders() js.Value {
	return ee.Get("RotationOrders")
}
func (ee *Euler) SetRotationOrders(v js.Value) {
	ee.Set("RotationOrders", v)
}
func (ee *Euler) Clone() *Euler {
	return &Euler{Value: ee.Call("clone")}
}
func (ee *Euler) Copy(euler *Euler) *Euler {
	return &Euler{Value: ee.Call("copy", euler)}
}
func (ee *Euler) Equals(euler *Euler) bool {
	return ee.Call("equals", euler).Bool()
}
func (ee *Euler) FromArray(xyzo js.Value) *Euler {
	return &Euler{Value: ee.Call("fromArray", xyzo)}
}
func (ee *Euler) OnChange(callback js.Value) *Euler {
	return &Euler{Value: ee.Call("onChange", callback)}
}
func (ee *Euler) Reorder(newOrder string) *Euler {
	return &Euler{Value: ee.Call("reorder", newOrder)}
}
func (ee *Euler) Set2(x float64, y float64, z float64, order string) *Euler {
	return &Euler{Value: ee.Call("set", x, y, z, order)}
}
func (ee *Euler) SetFromQuaternion(q *Quaternion, order string, update bool) *Euler {
	return &Euler{Value: ee.Call("setFromQuaternion", q, order, update)}
}
func (ee *Euler) SetFromRotationMatrix(m *Matrix4, order string, update bool) *Euler {
	return &Euler{Value: ee.Call("setFromRotationMatrix", m, order, update)}
}
func (ee *Euler) SetFromVector3(v *Vector3, order string) *Euler {
	return &Euler{Value: ee.Call("setFromVector3", v, order)}
}
func (ee *Euler) ToArray(array js.Value, offset int) js.Value {
	return ee.Call("toArray", array, offset)
}
func (ee *Euler) ToVector3(optionalResult *Vector3) *Vector3 {
	return &Vector3{Value: ee.Call("toVector3", optionalResult)}
}

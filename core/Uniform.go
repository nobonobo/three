package core

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Uniform struct {
	js.Value
}

func NewUniform(value js.Value) *Uniform {
	return &Uniform{Value: get("Uniform").New(value)}
}
func NewUniform2(typ string, value js.Value) *Uniform {
	return &Uniform{Value: get("Uniform").New(typ, value)}
}
func (u *Uniform) Dynamic() bool {
	return u.Get("dynamic").Bool()
}
func (u *Uniform) SetDynamic(v bool) {
	u.Set("dynamic", v)
}
func (u *Uniform) OnUpdateCallback() js.Value {
	return u.Get("onUpdateCallback")
}
func (u *Uniform) SetOnUpdateCallback(v js.Value) {
	u.Set("onUpdateCallback", v)
}
func (u *Uniform) Type() string {
	return u.Get("type").String()
}
func (u *Uniform) SetType(v string) {
	u.Set("type", v)
}
func (u *Uniform) Value() js.Value {
	return u.Get("value")
}
func (u *Uniform) SetValue(v js.Value) {
	u.Set("value", v)
}
func (u *Uniform) OnUpdate(callback js.Value) *Uniform {
	return &Uniform{Value: u.Call("onUpdate", callback)}
}

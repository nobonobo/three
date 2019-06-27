// Code generated from "core/Uniform.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Uniform extend: []
type Uniform struct {
	js.Value
}

func NewUniform(value js.Value) *Uniform {
	return &Uniform{Value: get("Uniform").New(value)}
}
func NewUniform2(typ string, value js.Value) *Uniform {
	return &Uniform{Value: get("Uniform").New(typ, value)}
}
func (uu *Uniform) JSValue() js.Value {
	return uu.Value
}
func (uu *Uniform) Dynamic() bool {
	return uu.Get("dynamic").Bool()
}
func (uu *Uniform) SetDynamic(v bool) {
	uu.Set("dynamic", v)
}
func (uu *Uniform) OnUpdateCallback() js.Value {
	return uu.Get("onUpdateCallback")
}
func (uu *Uniform) SetOnUpdateCallback(v js.Value) {
	uu.Set("onUpdateCallback", v)
}
func (uu *Uniform) Type() string {
	return uu.Get("type").String()
}
func (uu *Uniform) SetType(v string) {
	uu.Set("type", v)
}
func (uu *Uniform) Value2() js.Value {
	return uu.Get("value")
}
func (uu *Uniform) SetValue2(v js.Value) {
	uu.Set("value", v)
}
func (uu *Uniform) OnUpdate(callback js.Value) *Uniform {
	return &Uniform{Value: uu.Call("onUpdate", callback)}
}

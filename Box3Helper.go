// Code generated from "helpers/Box3Helper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Box3Helper struct {
	js.Value
}

func NewBox3Helper(object *Object3D, color *Color) *Box3Helper {
	return &Box3Helper{Value: get("Box3Helper").New(object, color)}
}
func (bh *Box3Helper) Box() *Box3 {
	return &Box3{Value: bh.Get("box")}
}
func (bh *Box3Helper) SetBox(v *Box3) {
	bh.Set("box", v)
}
func (bh *Box3Helper) Type() string {
	return bh.Get("type").String()
}
func (bh *Box3Helper) SetType(v string) {
	bh.Set("type", v)
}

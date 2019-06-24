// Code generated from "animation/AnimationObjectGroup.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type AnimationObjectGroup struct {
	js.Value
}

func NewAnimationObjectGroup(args js.Value) *AnimationObjectGroup {
	return &AnimationObjectGroup{Value: get("AnimationObjectGroup").New(args)}
}
func (aog *AnimationObjectGroup) Stats() js.Value {
	return aog.Get("stats")
}
func (aog *AnimationObjectGroup) SetStats(v js.Value) {
	aog.Set("stats", v)
}
func (aog *AnimationObjectGroup) Uuid() string {
	return aog.Get("uuid").String()
}
func (aog *AnimationObjectGroup) SetUuid(v string) {
	aog.Set("uuid", v)
}
func (aog *AnimationObjectGroup) Add(args js.Value) {
	aog.Call("add", args)
}
func (aog *AnimationObjectGroup) Remove(args js.Value) {
	aog.Call("remove", args)
}
func (aog *AnimationObjectGroup) Uncache(args js.Value) {
	aog.Call("uncache", args)
}

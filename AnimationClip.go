// Code generated from "animation/AnimationClip.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// AnimationClip extend: []
type AnimationClip struct {
	js.Value
}

func NewAnimationClip(name string, duration float64, tracks js.Value) *AnimationClip {
	return &AnimationClip{Value: get("AnimationClip").New(name, duration, tracks)}
}
func (ac *AnimationClip) JSValue() js.Value {
	return ac.Value
}
func (ac *AnimationClip) Duration() float64 {
	return ac.Get("duration").Float()
}
func (ac *AnimationClip) SetDuration(v float64) {
	ac.Set("duration", v)
}
func (ac *AnimationClip) Name() string {
	return ac.Get("name").String()
}
func (ac *AnimationClip) SetName(v string) {
	ac.Set("name", v)
}
func (ac *AnimationClip) Results() js.Value {
	return ac.Get("results")
}
func (ac *AnimationClip) SetResults(v js.Value) {
	ac.Set("results", v)
}
func (ac *AnimationClip) Tracks() js.Value {
	return ac.Get("tracks")
}
func (ac *AnimationClip) SetTracks(v js.Value) {
	ac.Set("tracks", v)
}
func (ac *AnimationClip) Uuid() string {
	return ac.Get("uuid").String()
}
func (ac *AnimationClip) SetUuid(v string) {
	ac.Set("uuid", v)
}
func (ac *AnimationClip) Optimize() *AnimationClip {
	return &AnimationClip{Value: ac.Call("optimize")}
}
func (ac *AnimationClip) ResetDuration() {
	ac.Call("resetDuration")
}
func (ac *AnimationClip) Trim() *AnimationClip {
	return &AnimationClip{Value: ac.Call("trim")}
}
func (ac *AnimationClip) CreateClipsFromMorphTargetSequences(morphTargets js.Value, fps float64, noLoop bool) js.Value {
	return ac.Call("CreateClipsFromMorphTargetSequences", morphTargets, fps, noLoop)
}
func (ac *AnimationClip) CreateFromMorphTargetSequence(name string, morphTargetSequence js.Value, fps float64, noLoop bool) *AnimationClip {
	return &AnimationClip{Value: ac.Call("CreateFromMorphTargetSequence", name, morphTargetSequence, fps, noLoop)}
}
func (ac *AnimationClip) FindByName(clipArray js.Value, name string) *AnimationClip {
	return &AnimationClip{Value: ac.Call("findByName", clipArray, name)}
}
func (ac *AnimationClip) Parse(json js.Value) *AnimationClip {
	return &AnimationClip{Value: ac.Call("parse", json)}
}
func (ac *AnimationClip) ParseAnimation(animation js.Value, bones js.Value, nodeName string) *AnimationClip {
	return &AnimationClip{Value: ac.Call("parseAnimation", animation, bones, nodeName)}
}
func (ac *AnimationClip) ToJSON() js.Value {
	return ac.Call("toJSON")
}

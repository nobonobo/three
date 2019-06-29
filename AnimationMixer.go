// Code generated from "animation/AnimationMixer.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// AnimationMixer extend: [EventDispatcher]
type AnimationMixer struct {
	js.Value
}

func NewAnimationMixer(root js.Value) *AnimationMixer {
	return &AnimationMixer{Value: get("AnimationMixer").New(root)}
}
func (am *AnimationMixer) JSValue() js.Value {
	return am.Value
}
func (am *AnimationMixer) Time() float64 {
	return am.Get("time").Float()
}
func (am *AnimationMixer) SetTime(v float64) {
	am.Set("time", v)
}
func (am *AnimationMixer) TimeScale() float64 {
	return am.Get("timeScale").Float()
}
func (am *AnimationMixer) SetTimeScale(v float64) {
	am.Set("timeScale", v)
}
func (am *AnimationMixer) AddEventListener(typ string, listener js.Value) {
	am.Call("addEventListener", typ, listener)
}
func (am *AnimationMixer) ClipAction(clip *AnimationClip, root js.Value) *AnimationAction {
	return &AnimationAction{Value: am.Call("clipAction", clip, root)}
}
func (am *AnimationMixer) DispatchEvent(event js.Value) {
	am.Call("dispatchEvent", event)
}
func (am *AnimationMixer) ExistingAction(clip *AnimationClip, root js.Value) *AnimationAction {
	return &AnimationAction{Value: am.Call("existingAction", clip, root)}
}
func (am *AnimationMixer) GetRoot() js.Value {
	return am.Call("getRoot")
}
func (am *AnimationMixer) HasEventListener(typ string, listener js.Value) bool {
	return am.Call("hasEventListener", typ, listener).Bool()
}
func (am *AnimationMixer) RemoveEventListener(typ string, listener js.Value) {
	am.Call("removeEventListener", typ, listener)
}
func (am *AnimationMixer) StopAllAction() *AnimationMixer {
	return &AnimationMixer{Value: am.Call("stopAllAction")}
}
func (am *AnimationMixer) UncacheAction(clip *AnimationClip, root js.Value) {
	am.Call("uncacheAction", clip.JSValue(), root)
}
func (am *AnimationMixer) UncacheClip(clip *AnimationClip) {
	am.Call("uncacheClip", clip.JSValue())
}
func (am *AnimationMixer) UncacheRoot(root js.Value) {
	am.Call("uncacheRoot", root)
}
func (am *AnimationMixer) Update(deltaTime float64) *AnimationMixer {
	return &AnimationMixer{Value: am.Call("update", deltaTime)}
}

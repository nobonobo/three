package animation

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
)

type AnimationMixer struct {
	js.Value
}

func NewAnimationMixer(root js.Value) *AnimationMixer {
	return &AnimationMixer{Value: get("AnimationMixer").New(root)}
}
func (am *AnimationMixer) Time() int {
	return am.Get("time").Int()
}
func (am *AnimationMixer) SetTime(v int) {
	am.Set("time", v)
}
func (am *AnimationMixer) TimeScale() int {
	return am.Get("timeScale").Int()
}
func (am *AnimationMixer) SetTimeScale(v int) {
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
	am.Call("uncacheAction", clip, root)
}
func (am *AnimationMixer) UncacheClip(clip *AnimationClip) {
	am.Call("uncacheClip", clip)
}
func (am *AnimationMixer) UncacheRoot(root js.Value) {
	am.Call("uncacheRoot", root)
}
func (am *AnimationMixer) Update(deltaTime int) *AnimationMixer {
	return &AnimationMixer{Value: am.Call("update", deltaTime)}
}

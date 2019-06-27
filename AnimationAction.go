// Code generated from "animation/AnimationAction.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// AnimationAction extend: []
type AnimationAction struct {
	js.Value
}

func (aa *AnimationAction) JSValue() js.Value {
	return aa.Value
}
func (aa *AnimationAction) ClampWhenFinished() bool {
	return aa.Get("clampWhenFinished").Bool()
}
func (aa *AnimationAction) SetClampWhenFinished(v bool) {
	aa.Set("clampWhenFinished", v)
}
func (aa *AnimationAction) Enabled() bool {
	return aa.Get("enabled").Bool()
}
func (aa *AnimationAction) SetEnabled(v bool) {
	aa.Set("enabled", v)
}
func (aa *AnimationAction) Loop() AnimationActionLoopStyles {
	return AnimationActionLoopStyles(aa.Get("loop"))
}
func (aa *AnimationAction) SetLoop(v AnimationActionLoopStyles) {
	aa.Set("loop", v)
}
func (aa *AnimationAction) Paused() bool {
	return aa.Get("paused").Bool()
}
func (aa *AnimationAction) SetPaused(v bool) {
	aa.Set("paused", v)
}
func (aa *AnimationAction) Repetitions() float64 {
	return aa.Get("repetitions").Float()
}
func (aa *AnimationAction) SetRepetitions(v float64) {
	aa.Set("repetitions", v)
}
func (aa *AnimationAction) Time() float64 {
	return aa.Get("time").Float()
}
func (aa *AnimationAction) SetTime(v float64) {
	aa.Set("time", v)
}
func (aa *AnimationAction) TimeScale() float64 {
	return aa.Get("timeScale").Float()
}
func (aa *AnimationAction) SetTimeScale(v float64) {
	aa.Set("timeScale", v)
}
func (aa *AnimationAction) Weight() float64 {
	return aa.Get("weight").Float()
}
func (aa *AnimationAction) SetWeight(v float64) {
	aa.Set("weight", v)
}
func (aa *AnimationAction) ZeroSlopeAtEnd() bool {
	return aa.Get("zeroSlopeAtEnd").Bool()
}
func (aa *AnimationAction) SetZeroSlopeAtEnd(v bool) {
	aa.Set("zeroSlopeAtEnd", v)
}
func (aa *AnimationAction) ZeroSlopeAtStart() bool {
	return aa.Get("zeroSlopeAtStart").Bool()
}
func (aa *AnimationAction) SetZeroSlopeAtStart(v bool) {
	aa.Set("zeroSlopeAtStart", v)
}
func (aa *AnimationAction) CrossFadeFrom(fadeOutAction *AnimationAction, duration float64, warp bool) *AnimationAction {
	return &AnimationAction{Value: aa.Call("crossFadeFrom", fadeOutAction, duration, warp)}
}
func (aa *AnimationAction) CrossFadeTo(fadeInAction *AnimationAction, duration float64, warp bool) *AnimationAction {
	return &AnimationAction{Value: aa.Call("crossFadeTo", fadeInAction, duration, warp)}
}
func (aa *AnimationAction) FadeIn(duration float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("fadeIn", duration)}
}
func (aa *AnimationAction) FadeOut(duration float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("fadeOut", duration)}
}
func (aa *AnimationAction) GetClip() *AnimationClip {
	return &AnimationClip{Value: aa.Call("getClip")}
}
func (aa *AnimationAction) GetEffectiveTimeScale() float64 {
	return aa.Call("getEffectiveTimeScale").Float()
}
func (aa *AnimationAction) GetEffectiveWeight() float64 {
	return aa.Call("getEffectiveWeight").Float()
}
func (aa *AnimationAction) GetMixer() *AnimationMixer {
	return &AnimationMixer{Value: aa.Call("getMixer")}
}
func (aa *AnimationAction) GetRoot() js.Value {
	return aa.Call("getRoot")
}
func (aa *AnimationAction) Halt(duration float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("halt", duration)}
}
func (aa *AnimationAction) IsRunning() bool {
	return aa.Call("isRunning").Bool()
}
func (aa *AnimationAction) IsScheduled() bool {
	return aa.Call("isScheduled").Bool()
}
func (aa *AnimationAction) Play() *AnimationAction {
	return &AnimationAction{Value: aa.Call("play")}
}
func (aa *AnimationAction) Reset() *AnimationAction {
	return &AnimationAction{Value: aa.Call("reset")}
}
func (aa *AnimationAction) SetDuration(duration float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setDuration", duration)}
}
func (aa *AnimationAction) SetEffectiveTimeScale(timeScale float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setEffectiveTimeScale", timeScale)}
}
func (aa *AnimationAction) SetEffectiveWeight(weight float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setEffectiveWeight", weight)}
}
func (aa *AnimationAction) SetLoop2(mode AnimationActionLoopStyles, repetitions float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setLoop", mode, repetitions)}
}
func (aa *AnimationAction) StartAt(time float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("startAt", time)}
}
func (aa *AnimationAction) Stop() *AnimationAction {
	return &AnimationAction{Value: aa.Call("stop")}
}
func (aa *AnimationAction) StopFading() *AnimationAction {
	return &AnimationAction{Value: aa.Call("stopFading")}
}
func (aa *AnimationAction) StopWarping() *AnimationAction {
	return &AnimationAction{Value: aa.Call("stopWarping")}
}
func (aa *AnimationAction) SyncWith(action *AnimationAction) *AnimationAction {
	return &AnimationAction{Value: aa.Call("syncWith", action)}
}
func (aa *AnimationAction) Warp(statTimeScale float64, endTimeScale float64, duration float64) *AnimationAction {
	return &AnimationAction{Value: aa.Call("warp", statTimeScale, endTimeScale, duration)}
}

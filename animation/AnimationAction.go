package animation

import (
	"github.com/gopherjs/gopherwasm/js"
)

type AnimationAction struct {
	js.Value
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
func (aa *AnimationAction) Loop() bool {
	return aa.Get("loop").Bool()
}
func (aa *AnimationAction) SetLoop(v bool) {
	aa.Set("loop", v)
}
func (aa *AnimationAction) Paused() bool {
	return aa.Get("paused").Bool()
}
func (aa *AnimationAction) SetPaused(v bool) {
	aa.Set("paused", v)
}
func (aa *AnimationAction) Repetitions() int {
	return aa.Get("repetitions").Int()
}
func (aa *AnimationAction) SetRepetitions(v int) {
	aa.Set("repetitions", v)
}
func (aa *AnimationAction) Time() int {
	return aa.Get("time").Int()
}
func (aa *AnimationAction) SetTime(v int) {
	aa.Set("time", v)
}
func (aa *AnimationAction) TimeScale() int {
	return aa.Get("timeScale").Int()
}
func (aa *AnimationAction) SetTimeScale(v int) {
	aa.Set("timeScale", v)
}
func (aa *AnimationAction) Weight() int {
	return aa.Get("weight").Int()
}
func (aa *AnimationAction) SetWeight(v int) {
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
func (aa *AnimationAction) CrossFadeFrom(fadeOutAction *AnimationAction, duration int, warp bool) *AnimationAction {
	return &AnimationAction{Value: aa.Call("crossFadeFrom", fadeOutAction, duration, warp)}
}
func (aa *AnimationAction) CrossFadeTo(fadeInAction *AnimationAction, duration int, warp bool) *AnimationAction {
	return &AnimationAction{Value: aa.Call("crossFadeTo", fadeInAction, duration, warp)}
}
func (aa *AnimationAction) FadeIn(duration int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("fadeIn", duration)}
}
func (aa *AnimationAction) FadeOut(duration int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("fadeOut", duration)}
}
func (aa *AnimationAction) GetClip() *AnimationClip {
	return &AnimationClip{Value: aa.Call("getClip")}
}
func (aa *AnimationAction) GetEffectiveTimeScale() int {
	return aa.Call("getEffectiveTimeScale").Int()
}
func (aa *AnimationAction) GetEffectiveWeight() int {
	return aa.Call("getEffectiveWeight").Int()
}
func (aa *AnimationAction) GetMixer() *AnimationMixer {
	return &AnimationMixer{Value: aa.Call("getMixer")}
}
func (aa *AnimationAction) GetRoot() js.Value {
	return aa.Call("getRoot")
}
func (aa *AnimationAction) Halt(duration int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("halt", duration)}
}
func (aa *AnimationAction) IsRunning() bool {
	return aa.Call("isRunning").Bool()
}
func (aa *AnimationAction) Play() *AnimationAction {
	return &AnimationAction{Value: aa.Call("play")}
}
func (aa *AnimationAction) Reset() *AnimationAction {
	return &AnimationAction{Value: aa.Call("reset")}
}
func (aa *AnimationAction) SetDuration(duration int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setDuration", duration)}
}
func (aa *AnimationAction) SetEffectiveTimeScale(timeScale int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setEffectiveTimeScale", timeScale)}
}
func (aa *AnimationAction) SetEffectiveWeight(weight int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setEffectiveWeight", weight)}
}
func (aa *AnimationAction) SetLoop(mode *AnimationActionLoopStyles, repetitions int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("setLoop", mode, repetitions)}
}
func (aa *AnimationAction) StartAt(time int) *AnimationAction {
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
func (aa *AnimationAction) Warp(statTimeScale int, endTimeScale int, duration int) *AnimationAction {
	return &AnimationAction{Value: aa.Call("warp", statTimeScale, endTimeScale, duration)}
}

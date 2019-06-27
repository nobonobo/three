// Code generated from "core/Clock.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Clock extend: []
type Clock struct {
	js.Value
}

func NewClock(autoStart bool) *Clock {
	return &Clock{Value: get("Clock").New(autoStart)}
}
func (cc *Clock) JSValue() js.Value {
	return cc.Value
}
func (cc *Clock) AutoStart() bool {
	return cc.Get("autoStart").Bool()
}
func (cc *Clock) SetAutoStart(v bool) {
	cc.Set("autoStart", v)
}
func (cc *Clock) ElapsedTime() float64 {
	return cc.Get("elapsedTime").Float()
}
func (cc *Clock) SetElapsedTime(v float64) {
	cc.Set("elapsedTime", v)
}
func (cc *Clock) OldTime() float64 {
	return cc.Get("oldTime").Float()
}
func (cc *Clock) SetOldTime(v float64) {
	cc.Set("oldTime", v)
}
func (cc *Clock) Running() bool {
	return cc.Get("running").Bool()
}
func (cc *Clock) SetRunning(v bool) {
	cc.Set("running", v)
}
func (cc *Clock) StartTime() float64 {
	return cc.Get("startTime").Float()
}
func (cc *Clock) SetStartTime(v float64) {
	cc.Set("startTime", v)
}
func (cc *Clock) GetDelta() float64 {
	return cc.Call("getDelta").Float()
}
func (cc *Clock) GetElapsedTime() float64 {
	return cc.Call("getElapsedTime").Float()
}
func (cc *Clock) Start() {
	cc.Call("start")
}
func (cc *Clock) Stop() {
	cc.Call("stop")
}

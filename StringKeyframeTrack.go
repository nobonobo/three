// Code generated from "animation/tracks/StringKeyframeTrack.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// StringKeyframeTrack extend: [KeyframeTrack]
type StringKeyframeTrack struct {
	js.Value
}

func NewStringKeyframeTrack(name string, times js.Value, values js.Value, interpolation InterpolationModes) *StringKeyframeTrack {
	return &StringKeyframeTrack{Value: get("StringKeyframeTrack").New(name, times, values, interpolation)}
}
func (skt *StringKeyframeTrack) JSValue() js.Value {
	return skt.Value
}
func (skt *StringKeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(skt.Get("DefaultInterpolation"))
}
func (skt *StringKeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	skt.Set("DefaultInterpolation", v)
}
func (skt *StringKeyframeTrack) TimeBufferType() js.Value {
	return skt.Get("TimeBufferType")
}
func (skt *StringKeyframeTrack) SetTimeBufferType(v js.Value) {
	skt.Set("TimeBufferType", v)
}
func (skt *StringKeyframeTrack) ValueBufferType() js.Value {
	return skt.Get("ValueBufferType")
}
func (skt *StringKeyframeTrack) SetValueBufferType(v js.Value) {
	skt.Set("ValueBufferType", v)
}
func (skt *StringKeyframeTrack) ValueTypeName() string {
	return skt.Get("ValueTypeName").String()
}
func (skt *StringKeyframeTrack) SetValueTypeName(v string) {
	skt.Set("ValueTypeName", v)
}
func (skt *StringKeyframeTrack) Name() string {
	return skt.Get("name").String()
}
func (skt *StringKeyframeTrack) SetName(v string) {
	skt.Set("name", v)
}
func (skt *StringKeyframeTrack) Times() js.Value {
	return skt.Get("times")
}
func (skt *StringKeyframeTrack) SetTimes(v js.Value) {
	skt.Set("times", v)
}
func (skt *StringKeyframeTrack) Values() js.Value {
	return skt.Get("values")
}
func (skt *StringKeyframeTrack) SetValues(v js.Value) {
	skt.Set("values", v)
}
func (skt *StringKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: skt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (skt *StringKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: skt.Call("InterpolantFactoryMethodLinear", result)}
}
func (skt *StringKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: skt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (skt *StringKeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(skt.Call("getInterpolation"))
}
func (skt *StringKeyframeTrack) GetValuesize() float64 {
	return skt.Call("getValuesize").Float()
}
func (skt *StringKeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: skt.Call("optimize")}
}
func (skt *StringKeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: skt.Call("scale", timeScale)}
}
func (skt *StringKeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	skt.Call("setInterpolation", interpolation)
}
func (skt *StringKeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: skt.Call("shift", timeOffset)}
}
func (skt *StringKeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: skt.Call("trim", startTime, endTime)}
}
func (skt *StringKeyframeTrack) Validate() bool {
	return skt.Call("validate").Bool()
}
func (skt *StringKeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: skt.Call("parse", json)}
}
func (skt *StringKeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return skt.Call("toJSON", track)
}

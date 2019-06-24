// Code generated from "animation/KeyframeTrack.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type KeyframeTrack struct {
	js.Value
}

func NewKeyframeTrack(name string, times js.Value, values js.Value, interpolation InterpolationModes) *KeyframeTrack {
	return &KeyframeTrack{Value: get("KeyframeTrack").New(name, times, values, interpolation)}
}
func (kt *KeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(kt.Get("DefaultInterpolation"))
}
func (kt *KeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	kt.Set("DefaultInterpolation", v)
}
func (kt *KeyframeTrack) TimeBufferType() js.Value {
	return kt.Get("TimeBufferType")
}
func (kt *KeyframeTrack) SetTimeBufferType(v js.Value) {
	kt.Set("TimeBufferType", v)
}
func (kt *KeyframeTrack) ValueBufferType() js.Value {
	return kt.Get("ValueBufferType")
}
func (kt *KeyframeTrack) SetValueBufferType(v js.Value) {
	kt.Set("ValueBufferType", v)
}
func (kt *KeyframeTrack) ValueTypeName() string {
	return kt.Get("ValueTypeName").String()
}
func (kt *KeyframeTrack) SetValueTypeName(v string) {
	kt.Set("ValueTypeName", v)
}
func (kt *KeyframeTrack) Name() string {
	return kt.Get("name").String()
}
func (kt *KeyframeTrack) SetName(v string) {
	kt.Set("name", v)
}
func (kt *KeyframeTrack) Times() js.Value {
	return kt.Get("times")
}
func (kt *KeyframeTrack) SetTimes(v js.Value) {
	kt.Set("times", v)
}
func (kt *KeyframeTrack) Values() js.Value {
	return kt.Get("values")
}
func (kt *KeyframeTrack) SetValues(v js.Value) {
	kt.Set("values", v)
}
func (kt *KeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: kt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (kt *KeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: kt.Call("InterpolantFactoryMethodLinear", result)}
}
func (kt *KeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: kt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (kt *KeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(kt.Call("getInterpolation"))
}
func (kt *KeyframeTrack) GetValuesize() float64 {
	return kt.Call("getValuesize").Float()
}
func (kt *KeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("optimize")}
}
func (kt *KeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("scale", timeScale)}
}
func (kt *KeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	kt.Call("setInterpolation", interpolation)
}
func (kt *KeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("shift", timeOffset)}
}
func (kt *KeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("trim", startTime, endTime)}
}
func (kt *KeyframeTrack) Validate() bool {
	return kt.Call("validate").Bool()
}
func (kt *KeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("parse", json)}
}
func (kt *KeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return kt.Call("toJSON", track)
}

// Code generated from "animation/tracks/ColorKeyframeTrack.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// ColorKeyframeTrack extend: [KeyframeTrack]
type ColorKeyframeTrack struct {
	js.Value
}

func NewColorKeyframeTrack(name string, times js.Value, values js.Value, interpolation InterpolationModes) *ColorKeyframeTrack {
	return &ColorKeyframeTrack{Value: get("ColorKeyframeTrack").New(name, times, values, interpolation)}
}
func (ckt *ColorKeyframeTrack) JSValue() js.Value {
	return ckt.Value
}
func (ckt *ColorKeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(ckt.Get("DefaultInterpolation"))
}
func (ckt *ColorKeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	ckt.Set("DefaultInterpolation", v)
}
func (ckt *ColorKeyframeTrack) TimeBufferType() js.Value {
	return ckt.Get("TimeBufferType")
}
func (ckt *ColorKeyframeTrack) SetTimeBufferType(v js.Value) {
	ckt.Set("TimeBufferType", v)
}
func (ckt *ColorKeyframeTrack) ValueBufferType() js.Value {
	return ckt.Get("ValueBufferType")
}
func (ckt *ColorKeyframeTrack) SetValueBufferType(v js.Value) {
	ckt.Set("ValueBufferType", v)
}
func (ckt *ColorKeyframeTrack) ValueTypeName() string {
	return ckt.Get("ValueTypeName").String()
}
func (ckt *ColorKeyframeTrack) SetValueTypeName(v string) {
	ckt.Set("ValueTypeName", v)
}
func (ckt *ColorKeyframeTrack) Name() string {
	return ckt.Get("name").String()
}
func (ckt *ColorKeyframeTrack) SetName(v string) {
	ckt.Set("name", v)
}
func (ckt *ColorKeyframeTrack) Times() js.Value {
	return ckt.Get("times")
}
func (ckt *ColorKeyframeTrack) SetTimes(v js.Value) {
	ckt.Set("times", v)
}
func (ckt *ColorKeyframeTrack) Values() js.Value {
	return ckt.Get("values")
}
func (ckt *ColorKeyframeTrack) SetValues(v js.Value) {
	ckt.Set("values", v)
}
func (ckt *ColorKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: ckt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (ckt *ColorKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: ckt.Call("InterpolantFactoryMethodLinear", result)}
}
func (ckt *ColorKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: ckt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (ckt *ColorKeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(ckt.Call("getInterpolation"))
}
func (ckt *ColorKeyframeTrack) GetValuesize() float64 {
	return ckt.Call("getValuesize").Float()
}
func (ckt *ColorKeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: ckt.Call("optimize")}
}
func (ckt *ColorKeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: ckt.Call("scale", timeScale)}
}
func (ckt *ColorKeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	ckt.Call("setInterpolation", interpolation)
}
func (ckt *ColorKeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: ckt.Call("shift", timeOffset)}
}
func (ckt *ColorKeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: ckt.Call("trim", startTime, endTime)}
}
func (ckt *ColorKeyframeTrack) Validate() bool {
	return ckt.Call("validate").Bool()
}
func (ckt *ColorKeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: ckt.Call("parse", json)}
}
func (ckt *ColorKeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return ckt.Call("toJSON", track)
}

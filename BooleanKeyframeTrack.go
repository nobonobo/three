// Code generated from "animation/tracks/BooleanKeyframeTrack.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type BooleanKeyframeTrack struct {
	js.Value
}

func NewBooleanKeyframeTrack(name string, times js.Value, values js.Value) *BooleanKeyframeTrack {
	return &BooleanKeyframeTrack{Value: get("BooleanKeyframeTrack").New(name, times, values)}
}
func (bkt *BooleanKeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(bkt.Get("DefaultInterpolation"))
}
func (bkt *BooleanKeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	bkt.Set("DefaultInterpolation", v)
}
func (bkt *BooleanKeyframeTrack) TimeBufferType() js.Value {
	return bkt.Get("TimeBufferType")
}
func (bkt *BooleanKeyframeTrack) SetTimeBufferType(v js.Value) {
	bkt.Set("TimeBufferType", v)
}
func (bkt *BooleanKeyframeTrack) ValueBufferType() js.Value {
	return bkt.Get("ValueBufferType")
}
func (bkt *BooleanKeyframeTrack) SetValueBufferType(v js.Value) {
	bkt.Set("ValueBufferType", v)
}
func (bkt *BooleanKeyframeTrack) ValueTypeName() string {
	return bkt.Get("ValueTypeName").String()
}
func (bkt *BooleanKeyframeTrack) SetValueTypeName(v string) {
	bkt.Set("ValueTypeName", v)
}
func (bkt *BooleanKeyframeTrack) Name() string {
	return bkt.Get("name").String()
}
func (bkt *BooleanKeyframeTrack) SetName(v string) {
	bkt.Set("name", v)
}
func (bkt *BooleanKeyframeTrack) Times() js.Value {
	return bkt.Get("times")
}
func (bkt *BooleanKeyframeTrack) SetTimes(v js.Value) {
	bkt.Set("times", v)
}
func (bkt *BooleanKeyframeTrack) Values() js.Value {
	return bkt.Get("values")
}
func (bkt *BooleanKeyframeTrack) SetValues(v js.Value) {
	bkt.Set("values", v)
}
func (bkt *BooleanKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: bkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (bkt *BooleanKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: bkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (bkt *BooleanKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: bkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (bkt *BooleanKeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(bkt.Call("getInterpolation"))
}
func (bkt *BooleanKeyframeTrack) GetValuesize() float64 {
	return bkt.Call("getValuesize").Float()
}
func (bkt *BooleanKeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: bkt.Call("optimize")}
}
func (bkt *BooleanKeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: bkt.Call("scale", timeScale)}
}
func (bkt *BooleanKeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	bkt.Call("setInterpolation", interpolation)
}
func (bkt *BooleanKeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: bkt.Call("shift", timeOffset)}
}
func (bkt *BooleanKeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: bkt.Call("trim", startTime, endTime)}
}
func (bkt *BooleanKeyframeTrack) Validate() bool {
	return bkt.Call("validate").Bool()
}
func (bkt *BooleanKeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: bkt.Call("parse", json)}
}
func (bkt *BooleanKeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return bkt.Call("toJSON", track)
}

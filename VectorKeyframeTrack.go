// Code generated from "animation/tracks/VectorKeyframeTrack.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// VectorKeyframeTrack extend: [KeyframeTrack]
type VectorKeyframeTrack struct {
	js.Value
}

func NewVectorKeyframeTrack(name string, times js.Value, values js.Value, interpolation InterpolationModes) *VectorKeyframeTrack {
	return &VectorKeyframeTrack{Value: get("VectorKeyframeTrack").New(name, times, values, interpolation)}
}
func (vkt *VectorKeyframeTrack) JSValue() js.Value {
	return vkt.Value
}
func (vkt *VectorKeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(vkt.Get("DefaultInterpolation"))
}
func (vkt *VectorKeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	vkt.Set("DefaultInterpolation", v)
}
func (vkt *VectorKeyframeTrack) TimeBufferType() js.Value {
	return vkt.Get("TimeBufferType")
}
func (vkt *VectorKeyframeTrack) SetTimeBufferType(v js.Value) {
	vkt.Set("TimeBufferType", v)
}
func (vkt *VectorKeyframeTrack) ValueBufferType() js.Value {
	return vkt.Get("ValueBufferType")
}
func (vkt *VectorKeyframeTrack) SetValueBufferType(v js.Value) {
	vkt.Set("ValueBufferType", v)
}
func (vkt *VectorKeyframeTrack) ValueTypeName() string {
	return vkt.Get("ValueTypeName").String()
}
func (vkt *VectorKeyframeTrack) SetValueTypeName(v string) {
	vkt.Set("ValueTypeName", v)
}
func (vkt *VectorKeyframeTrack) Name() string {
	return vkt.Get("name").String()
}
func (vkt *VectorKeyframeTrack) SetName(v string) {
	vkt.Set("name", v)
}
func (vkt *VectorKeyframeTrack) Times() js.Value {
	return vkt.Get("times")
}
func (vkt *VectorKeyframeTrack) SetTimes(v js.Value) {
	vkt.Set("times", v)
}
func (vkt *VectorKeyframeTrack) Values() js.Value {
	return vkt.Get("values")
}
func (vkt *VectorKeyframeTrack) SetValues(v js.Value) {
	vkt.Set("values", v)
}
func (vkt *VectorKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: vkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (vkt *VectorKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: vkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (vkt *VectorKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: vkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (vkt *VectorKeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(vkt.Call("getInterpolation"))
}
func (vkt *VectorKeyframeTrack) GetValuesize() float64 {
	return vkt.Call("getValuesize").Float()
}
func (vkt *VectorKeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: vkt.Call("optimize")}
}
func (vkt *VectorKeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: vkt.Call("scale", timeScale)}
}
func (vkt *VectorKeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	vkt.Call("setInterpolation", interpolation)
}
func (vkt *VectorKeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: vkt.Call("shift", timeOffset)}
}
func (vkt *VectorKeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: vkt.Call("trim", startTime, endTime)}
}
func (vkt *VectorKeyframeTrack) Validate() bool {
	return vkt.Call("validate").Bool()
}
func (vkt *VectorKeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: vkt.Call("parse", json)}
}
func (vkt *VectorKeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return vkt.Call("toJSON", track)
}

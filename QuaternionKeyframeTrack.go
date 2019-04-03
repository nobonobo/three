// Code generated from "animation/tracks/QuaternionKeyframeTrack.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type QuaternionKeyframeTrack struct {
	js.Value
}

func NewQuaternionKeyframeTrack(name string, times js.Value, values js.Value, interpolation InterpolationModes) *QuaternionKeyframeTrack {
	return &QuaternionKeyframeTrack{Value: get("QuaternionKeyframeTrack").New(name, times, values, interpolation)}
}
func (qkt *QuaternionKeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(qkt.Get("DefaultInterpolation"))
}
func (qkt *QuaternionKeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	qkt.Set("DefaultInterpolation", v)
}
func (qkt *QuaternionKeyframeTrack) TimeBufferType() js.Value {
	return qkt.Get("TimeBufferType")
}
func (qkt *QuaternionKeyframeTrack) SetTimeBufferType(v js.Value) {
	qkt.Set("TimeBufferType", v)
}
func (qkt *QuaternionKeyframeTrack) ValueBufferType() js.Value {
	return qkt.Get("ValueBufferType")
}
func (qkt *QuaternionKeyframeTrack) SetValueBufferType(v js.Value) {
	qkt.Set("ValueBufferType", v)
}
func (qkt *QuaternionKeyframeTrack) ValueTypeName() string {
	return qkt.Get("ValueTypeName").String()
}
func (qkt *QuaternionKeyframeTrack) SetValueTypeName(v string) {
	qkt.Set("ValueTypeName", v)
}
func (qkt *QuaternionKeyframeTrack) Name() string {
	return qkt.Get("name").String()
}
func (qkt *QuaternionKeyframeTrack) SetName(v string) {
	qkt.Set("name", v)
}
func (qkt *QuaternionKeyframeTrack) Times() js.Value {
	return qkt.Get("times")
}
func (qkt *QuaternionKeyframeTrack) SetTimes(v js.Value) {
	qkt.Set("times", v)
}
func (qkt *QuaternionKeyframeTrack) Values() js.Value {
	return qkt.Get("values")
}
func (qkt *QuaternionKeyframeTrack) SetValues(v js.Value) {
	qkt.Set("values", v)
}
func (qkt *QuaternionKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: qkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (qkt *QuaternionKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: qkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (qkt *QuaternionKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: qkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (qkt *QuaternionKeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(qkt.Call("getInterpolation"))
}
func (qkt *QuaternionKeyframeTrack) GetValuesize() float64 {
	return qkt.Call("getValuesize").Float()
}
func (qkt *QuaternionKeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: qkt.Call("optimize")}
}
func (qkt *QuaternionKeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: qkt.Call("scale", timeScale)}
}
func (qkt *QuaternionKeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	qkt.Call("setInterpolation", interpolation)
}
func (qkt *QuaternionKeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: qkt.Call("shift", timeOffset)}
}
func (qkt *QuaternionKeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: qkt.Call("trim", startTime, endTime)}
}
func (qkt *QuaternionKeyframeTrack) Validate() bool {
	return qkt.Call("validate").Bool()
}
func (qkt *QuaternionKeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: qkt.Call("parse", json)}
}
func (qkt *QuaternionKeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return qkt.Call("toJSON", track)
}

// Code generated from "animation/tracks/NumberKeyframeTrack.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type NumberKeyframeTrack struct {
	js.Value
}

func NewNumberKeyframeTrack(name string, times js.Value, values js.Value, interpolation InterpolationModes) *NumberKeyframeTrack {
	return &NumberKeyframeTrack{Value: get("NumberKeyframeTrack").New(name, times, values, interpolation)}
}
func (nkt *NumberKeyframeTrack) DefaultInterpolation() InterpolationModes {
	return InterpolationModes(nkt.Get("DefaultInterpolation"))
}
func (nkt *NumberKeyframeTrack) SetDefaultInterpolation(v InterpolationModes) {
	nkt.Set("DefaultInterpolation", v)
}
func (nkt *NumberKeyframeTrack) TimeBufferType() js.Value {
	return nkt.Get("TimeBufferType")
}
func (nkt *NumberKeyframeTrack) SetTimeBufferType(v js.Value) {
	nkt.Set("TimeBufferType", v)
}
func (nkt *NumberKeyframeTrack) ValueBufferType() js.Value {
	return nkt.Get("ValueBufferType")
}
func (nkt *NumberKeyframeTrack) SetValueBufferType(v js.Value) {
	nkt.Set("ValueBufferType", v)
}
func (nkt *NumberKeyframeTrack) ValueTypeName() string {
	return nkt.Get("ValueTypeName").String()
}
func (nkt *NumberKeyframeTrack) SetValueTypeName(v string) {
	nkt.Set("ValueTypeName", v)
}
func (nkt *NumberKeyframeTrack) Name() string {
	return nkt.Get("name").String()
}
func (nkt *NumberKeyframeTrack) SetName(v string) {
	nkt.Set("name", v)
}
func (nkt *NumberKeyframeTrack) Times() js.Value {
	return nkt.Get("times")
}
func (nkt *NumberKeyframeTrack) SetTimes(v js.Value) {
	nkt.Set("times", v)
}
func (nkt *NumberKeyframeTrack) Values() js.Value {
	return nkt.Get("values")
}
func (nkt *NumberKeyframeTrack) SetValues(v js.Value) {
	nkt.Set("values", v)
}
func (nkt *NumberKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: nkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (nkt *NumberKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: nkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (nkt *NumberKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: nkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (nkt *NumberKeyframeTrack) GetInterpolation() InterpolationModes {
	return InterpolationModes(nkt.Call("getInterpolation"))
}
func (nkt *NumberKeyframeTrack) GetValuesize() float64 {
	return nkt.Call("getValuesize").Float()
}
func (nkt *NumberKeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: nkt.Call("optimize")}
}
func (nkt *NumberKeyframeTrack) Scale(timeScale float64) *KeyframeTrack {
	return &KeyframeTrack{Value: nkt.Call("scale", timeScale)}
}
func (nkt *NumberKeyframeTrack) SetInterpolation(interpolation InterpolationModes) {
	nkt.Call("setInterpolation", interpolation)
}
func (nkt *NumberKeyframeTrack) Shift(timeOffset float64) *KeyframeTrack {
	return &KeyframeTrack{Value: nkt.Call("shift", timeOffset)}
}
func (nkt *NumberKeyframeTrack) Trim(startTime float64, endTime float64) *KeyframeTrack {
	return &KeyframeTrack{Value: nkt.Call("trim", startTime, endTime)}
}
func (nkt *NumberKeyframeTrack) Validate() bool {
	return nkt.Call("validate").Bool()
}
func (nkt *NumberKeyframeTrack) Parse(json js.Value) *KeyframeTrack {
	return &KeyframeTrack{Value: nkt.Call("parse", json)}
}
func (nkt *NumberKeyframeTrack) ToJSON(track *KeyframeTrack) js.Value {
	return nkt.Call("toJSON", track)
}

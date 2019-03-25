package tracks

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math/interpolants"
)

type StringKeyframeTrack struct {
	js.Value
}

func NewStringKeyframeTrack(name string, times []js.Value, values []js.Value, interpolation *InterpolationModes) *StringKeyframeTrack {
	return &StringKeyframeTrack{Value: get("StringKeyframeTrack").New(name, times, values, interpolation)}
}
func (skt *StringKeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: skt.Get("DefaultInterpolation")}
}
func (skt *StringKeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	skt.Set("DefaultInterpolation", v)
}
func (skt *StringKeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: skt.Get("TimeBufferType")}
}
func (skt *StringKeyframeTrack) SetTimeBufferType(v *Float32Array) {
	skt.Set("TimeBufferType", v)
}
func (skt *StringKeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: skt.Get("ValueBufferType")}
}
func (skt *StringKeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (skt *StringKeyframeTrack) Times() []js.Value {
	return []js.Value(skt.Get("times"))
}
func (skt *StringKeyframeTrack) SetTimes(v []js.Value) {
	skt.Set("times", v)
}
func (skt *StringKeyframeTrack) Values() []js.Value {
	return []js.Value(skt.Get("values"))
}
func (skt *StringKeyframeTrack) SetValues(v []js.Value) {
	skt.Set("values", v)
}
func (skt *StringKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: skt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (skt *StringKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: skt.Call("InterpolantFactoryMethodLinear", result)}
}
func (skt *StringKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: skt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (skt *StringKeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: skt.Call("getInterpolation")}
}
func (skt *StringKeyframeTrack) GetValuesize() int {
	return skt.Call("getValuesize").Int()
}
func (skt *StringKeyframeTrack) Optimize() *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: skt.Call("optimize")}
}
func (skt *StringKeyframeTrack) Scale(timeScale int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: skt.Call("scale", timeScale)}
}
func (skt *StringKeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	skt.Call("setInterpolation", interpolation)
}
func (skt *StringKeyframeTrack) Shift(timeOffset int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: skt.Call("shift", timeOffset)}
}
func (skt *StringKeyframeTrack) Trim(startTime int, endTime int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: skt.Call("trim", startTime, endTime)}
}
func (skt *StringKeyframeTrack) Validate() bool {
	return skt.Call("validate").Bool()
}
func (skt *StringKeyframeTrack) Parse(json js.Value) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: skt.Call("parse", json)}
}
func (skt *StringKeyframeTrack) ToJSON(track *animation.KeyframeTrack) js.Value {
	return skt.Call("toJSON", track)
}

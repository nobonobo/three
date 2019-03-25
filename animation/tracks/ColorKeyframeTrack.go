package tracks

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math/interpolants"
)

type ColorKeyframeTrack struct {
	js.Value
}

func NewColorKeyframeTrack(name string, times []js.Value, values []js.Value, interpolation *InterpolationModes) *ColorKeyframeTrack {
	return &ColorKeyframeTrack{Value: get("ColorKeyframeTrack").New(name, times, values, interpolation)}
}
func (ckt *ColorKeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: ckt.Get("DefaultInterpolation")}
}
func (ckt *ColorKeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	ckt.Set("DefaultInterpolation", v)
}
func (ckt *ColorKeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: ckt.Get("TimeBufferType")}
}
func (ckt *ColorKeyframeTrack) SetTimeBufferType(v *Float32Array) {
	ckt.Set("TimeBufferType", v)
}
func (ckt *ColorKeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: ckt.Get("ValueBufferType")}
}
func (ckt *ColorKeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (ckt *ColorKeyframeTrack) Times() []js.Value {
	return []js.Value(ckt.Get("times"))
}
func (ckt *ColorKeyframeTrack) SetTimes(v []js.Value) {
	ckt.Set("times", v)
}
func (ckt *ColorKeyframeTrack) Values() []js.Value {
	return []js.Value(ckt.Get("values"))
}
func (ckt *ColorKeyframeTrack) SetValues(v []js.Value) {
	ckt.Set("values", v)
}
func (ckt *ColorKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: ckt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (ckt *ColorKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: ckt.Call("InterpolantFactoryMethodLinear", result)}
}
func (ckt *ColorKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: ckt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (ckt *ColorKeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: ckt.Call("getInterpolation")}
}
func (ckt *ColorKeyframeTrack) GetValuesize() int {
	return ckt.Call("getValuesize").Int()
}
func (ckt *ColorKeyframeTrack) Optimize() *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: ckt.Call("optimize")}
}
func (ckt *ColorKeyframeTrack) Scale(timeScale int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: ckt.Call("scale", timeScale)}
}
func (ckt *ColorKeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	ckt.Call("setInterpolation", interpolation)
}
func (ckt *ColorKeyframeTrack) Shift(timeOffset int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: ckt.Call("shift", timeOffset)}
}
func (ckt *ColorKeyframeTrack) Trim(startTime int, endTime int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: ckt.Call("trim", startTime, endTime)}
}
func (ckt *ColorKeyframeTrack) Validate() bool {
	return ckt.Call("validate").Bool()
}
func (ckt *ColorKeyframeTrack) Parse(json js.Value) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: ckt.Call("parse", json)}
}
func (ckt *ColorKeyframeTrack) ToJSON(track *animation.KeyframeTrack) js.Value {
	return ckt.Call("toJSON", track)
}

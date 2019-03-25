package tracks

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math/interpolants"
)

type NumberKeyframeTrack struct {
	js.Value
}

func NewNumberKeyframeTrack(name string, times []js.Value, values []js.Value, interpolation *InterpolationModes) *NumberKeyframeTrack {
	return &NumberKeyframeTrack{Value: get("NumberKeyframeTrack").New(name, times, values, interpolation)}
}
func (nkt *NumberKeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: nkt.Get("DefaultInterpolation")}
}
func (nkt *NumberKeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	nkt.Set("DefaultInterpolation", v)
}
func (nkt *NumberKeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: nkt.Get("TimeBufferType")}
}
func (nkt *NumberKeyframeTrack) SetTimeBufferType(v *Float32Array) {
	nkt.Set("TimeBufferType", v)
}
func (nkt *NumberKeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: nkt.Get("ValueBufferType")}
}
func (nkt *NumberKeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (nkt *NumberKeyframeTrack) Times() []js.Value {
	return []js.Value(nkt.Get("times"))
}
func (nkt *NumberKeyframeTrack) SetTimes(v []js.Value) {
	nkt.Set("times", v)
}
func (nkt *NumberKeyframeTrack) Values() []js.Value {
	return []js.Value(nkt.Get("values"))
}
func (nkt *NumberKeyframeTrack) SetValues(v []js.Value) {
	nkt.Set("values", v)
}
func (nkt *NumberKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: nkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (nkt *NumberKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: nkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (nkt *NumberKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: nkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (nkt *NumberKeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: nkt.Call("getInterpolation")}
}
func (nkt *NumberKeyframeTrack) GetValuesize() int {
	return nkt.Call("getValuesize").Int()
}
func (nkt *NumberKeyframeTrack) Optimize() *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: nkt.Call("optimize")}
}
func (nkt *NumberKeyframeTrack) Scale(timeScale int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: nkt.Call("scale", timeScale)}
}
func (nkt *NumberKeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	nkt.Call("setInterpolation", interpolation)
}
func (nkt *NumberKeyframeTrack) Shift(timeOffset int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: nkt.Call("shift", timeOffset)}
}
func (nkt *NumberKeyframeTrack) Trim(startTime int, endTime int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: nkt.Call("trim", startTime, endTime)}
}
func (nkt *NumberKeyframeTrack) Validate() bool {
	return nkt.Call("validate").Bool()
}
func (nkt *NumberKeyframeTrack) Parse(json js.Value) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: nkt.Call("parse", json)}
}
func (nkt *NumberKeyframeTrack) ToJSON(track *animation.KeyframeTrack) js.Value {
	return nkt.Call("toJSON", track)
}

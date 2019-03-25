package tracks

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math/interpolants"
)

type BooleanKeyframeTrack struct {
	js.Value
}

func NewBooleanKeyframeTrack(name string, times []js.Value, values []js.Value) *BooleanKeyframeTrack {
	return &BooleanKeyframeTrack{Value: get("BooleanKeyframeTrack").New(name, times, values)}
}
func (bkt *BooleanKeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: bkt.Get("DefaultInterpolation")}
}
func (bkt *BooleanKeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	bkt.Set("DefaultInterpolation", v)
}
func (bkt *BooleanKeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: bkt.Get("TimeBufferType")}
}
func (bkt *BooleanKeyframeTrack) SetTimeBufferType(v *Float32Array) {
	bkt.Set("TimeBufferType", v)
}
func (bkt *BooleanKeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: bkt.Get("ValueBufferType")}
}
func (bkt *BooleanKeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (bkt *BooleanKeyframeTrack) Times() []js.Value {
	return []js.Value(bkt.Get("times"))
}
func (bkt *BooleanKeyframeTrack) SetTimes(v []js.Value) {
	bkt.Set("times", v)
}
func (bkt *BooleanKeyframeTrack) Values() []js.Value {
	return []js.Value(bkt.Get("values"))
}
func (bkt *BooleanKeyframeTrack) SetValues(v []js.Value) {
	bkt.Set("values", v)
}
func (bkt *BooleanKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: bkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (bkt *BooleanKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: bkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (bkt *BooleanKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: bkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (bkt *BooleanKeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: bkt.Call("getInterpolation")}
}
func (bkt *BooleanKeyframeTrack) GetValuesize() int {
	return bkt.Call("getValuesize").Int()
}
func (bkt *BooleanKeyframeTrack) Optimize() *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: bkt.Call("optimize")}
}
func (bkt *BooleanKeyframeTrack) Scale(timeScale int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: bkt.Call("scale", timeScale)}
}
func (bkt *BooleanKeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	bkt.Call("setInterpolation", interpolation)
}
func (bkt *BooleanKeyframeTrack) Shift(timeOffset int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: bkt.Call("shift", timeOffset)}
}
func (bkt *BooleanKeyframeTrack) Trim(startTime int, endTime int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: bkt.Call("trim", startTime, endTime)}
}
func (bkt *BooleanKeyframeTrack) Validate() bool {
	return bkt.Call("validate").Bool()
}
func (bkt *BooleanKeyframeTrack) Parse(json js.Value) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: bkt.Call("parse", json)}
}
func (bkt *BooleanKeyframeTrack) ToJSON(track *animation.KeyframeTrack) js.Value {
	return bkt.Call("toJSON", track)
}

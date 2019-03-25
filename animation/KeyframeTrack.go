package animation

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math/interpolants"
)

type KeyframeTrack struct {
	js.Value
}

func NewKeyframeTrack(name string, times []js.Value, values []js.Value, interpolation *InterpolationModes) *KeyframeTrack {
	return &KeyframeTrack{Value: get("KeyframeTrack").New(name, times, values, interpolation)}
}
func (kt *KeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: kt.Get("DefaultInterpolation")}
}
func (kt *KeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	kt.Set("DefaultInterpolation", v)
}
func (kt *KeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: kt.Get("TimeBufferType")}
}
func (kt *KeyframeTrack) SetTimeBufferType(v *Float32Array) {
	kt.Set("TimeBufferType", v)
}
func (kt *KeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: kt.Get("ValueBufferType")}
}
func (kt *KeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (kt *KeyframeTrack) Times() []js.Value {
	return []js.Value(kt.Get("times"))
}
func (kt *KeyframeTrack) SetTimes(v []js.Value) {
	kt.Set("times", v)
}
func (kt *KeyframeTrack) Values() []js.Value {
	return []js.Value(kt.Get("values"))
}
func (kt *KeyframeTrack) SetValues(v []js.Value) {
	kt.Set("values", v)
}
func (kt *KeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: kt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (kt *KeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: kt.Call("InterpolantFactoryMethodLinear", result)}
}
func (kt *KeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: kt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (kt *KeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: kt.Call("getInterpolation")}
}
func (kt *KeyframeTrack) GetValuesize() int {
	return kt.Call("getValuesize").Int()
}
func (kt *KeyframeTrack) Optimize() *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("optimize")}
}
func (kt *KeyframeTrack) Scale(timeScale int) *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("scale", timeScale)}
}
func (kt *KeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	kt.Call("setInterpolation", interpolation)
}
func (kt *KeyframeTrack) Shift(timeOffset int) *KeyframeTrack {
	return &KeyframeTrack{Value: kt.Call("shift", timeOffset)}
}
func (kt *KeyframeTrack) Trim(startTime int, endTime int) *KeyframeTrack {
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

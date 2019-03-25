package tracks

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math/interpolants"
)

type VectorKeyframeTrack struct {
	js.Value
}

func NewVectorKeyframeTrack(name string, times []js.Value, values []js.Value, interpolation *InterpolationModes) *VectorKeyframeTrack {
	return &VectorKeyframeTrack{Value: get("VectorKeyframeTrack").New(name, times, values, interpolation)}
}
func (vkt *VectorKeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: vkt.Get("DefaultInterpolation")}
}
func (vkt *VectorKeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	vkt.Set("DefaultInterpolation", v)
}
func (vkt *VectorKeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: vkt.Get("TimeBufferType")}
}
func (vkt *VectorKeyframeTrack) SetTimeBufferType(v *Float32Array) {
	vkt.Set("TimeBufferType", v)
}
func (vkt *VectorKeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: vkt.Get("ValueBufferType")}
}
func (vkt *VectorKeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (vkt *VectorKeyframeTrack) Times() []js.Value {
	return []js.Value(vkt.Get("times"))
}
func (vkt *VectorKeyframeTrack) SetTimes(v []js.Value) {
	vkt.Set("times", v)
}
func (vkt *VectorKeyframeTrack) Values() []js.Value {
	return []js.Value(vkt.Get("values"))
}
func (vkt *VectorKeyframeTrack) SetValues(v []js.Value) {
	vkt.Set("values", v)
}
func (vkt *VectorKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: vkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (vkt *VectorKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: vkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (vkt *VectorKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: vkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (vkt *VectorKeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: vkt.Call("getInterpolation")}
}
func (vkt *VectorKeyframeTrack) GetValuesize() int {
	return vkt.Call("getValuesize").Int()
}
func (vkt *VectorKeyframeTrack) Optimize() *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: vkt.Call("optimize")}
}
func (vkt *VectorKeyframeTrack) Scale(timeScale int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: vkt.Call("scale", timeScale)}
}
func (vkt *VectorKeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	vkt.Call("setInterpolation", interpolation)
}
func (vkt *VectorKeyframeTrack) Shift(timeOffset int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: vkt.Call("shift", timeOffset)}
}
func (vkt *VectorKeyframeTrack) Trim(startTime int, endTime int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: vkt.Call("trim", startTime, endTime)}
}
func (vkt *VectorKeyframeTrack) Validate() bool {
	return vkt.Call("validate").Bool()
}
func (vkt *VectorKeyframeTrack) Parse(json js.Value) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: vkt.Call("parse", json)}
}
func (vkt *VectorKeyframeTrack) ToJSON(track *animation.KeyframeTrack) js.Value {
	return vkt.Call("toJSON", track)
}

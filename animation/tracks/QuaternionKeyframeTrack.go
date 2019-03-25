package tracks

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math/interpolants"
)

type QuaternionKeyframeTrack struct {
	js.Value
}

func NewQuaternionKeyframeTrack(name string, times []js.Value, values []js.Value, interpolation *InterpolationModes) *QuaternionKeyframeTrack {
	return &QuaternionKeyframeTrack{Value: get("QuaternionKeyframeTrack").New(name, times, values, interpolation)}
}
func (qkt *QuaternionKeyframeTrack) DefaultInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: qkt.Get("DefaultInterpolation")}
}
func (qkt *QuaternionKeyframeTrack) SetDefaultInterpolation(v *InterpolationModes) {
	qkt.Set("DefaultInterpolation", v)
}
func (qkt *QuaternionKeyframeTrack) TimeBufferType() *Float32Array {
	return &Float32Array{Value: qkt.Get("TimeBufferType")}
}
func (qkt *QuaternionKeyframeTrack) SetTimeBufferType(v *Float32Array) {
	qkt.Set("TimeBufferType", v)
}
func (qkt *QuaternionKeyframeTrack) ValueBufferType() *Float32Array {
	return &Float32Array{Value: qkt.Get("ValueBufferType")}
}
func (qkt *QuaternionKeyframeTrack) SetValueBufferType(v *Float32Array) {
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
func (qkt *QuaternionKeyframeTrack) Times() []js.Value {
	return []js.Value(qkt.Get("times"))
}
func (qkt *QuaternionKeyframeTrack) SetTimes(v []js.Value) {
	qkt.Set("times", v)
}
func (qkt *QuaternionKeyframeTrack) Values() []js.Value {
	return []js.Value(qkt.Get("values"))
}
func (qkt *QuaternionKeyframeTrack) SetValues(v []js.Value) {
	qkt.Set("values", v)
}
func (qkt *QuaternionKeyframeTrack) InterpolantFactoryMethodDiscrete(result js.Value) *interpolants.DiscreteInterpolant {
	return &interpolants.DiscreteInterpolant{Value: qkt.Call("InterpolantFactoryMethodDiscrete", result)}
}
func (qkt *QuaternionKeyframeTrack) InterpolantFactoryMethodLinear(result js.Value) *interpolants.LinearInterpolant {
	return &interpolants.LinearInterpolant{Value: qkt.Call("InterpolantFactoryMethodLinear", result)}
}
func (qkt *QuaternionKeyframeTrack) InterpolantFactoryMethodSmooth(result js.Value) *interpolants.CubicInterpolant {
	return &interpolants.CubicInterpolant{Value: qkt.Call("InterpolantFactoryMethodSmooth", result)}
}
func (qkt *QuaternionKeyframeTrack) GetInterpolation() *InterpolationModes {
	return &InterpolationModes{Value: qkt.Call("getInterpolation")}
}
func (qkt *QuaternionKeyframeTrack) GetValuesize() int {
	return qkt.Call("getValuesize").Int()
}
func (qkt *QuaternionKeyframeTrack) Optimize() *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: qkt.Call("optimize")}
}
func (qkt *QuaternionKeyframeTrack) Scale(timeScale int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: qkt.Call("scale", timeScale)}
}
func (qkt *QuaternionKeyframeTrack) SetInterpolation(interpolation *InterpolationModes) {
	qkt.Call("setInterpolation", interpolation)
}
func (qkt *QuaternionKeyframeTrack) Shift(timeOffset int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: qkt.Call("shift", timeOffset)}
}
func (qkt *QuaternionKeyframeTrack) Trim(startTime int, endTime int) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: qkt.Call("trim", startTime, endTime)}
}
func (qkt *QuaternionKeyframeTrack) Validate() bool {
	return qkt.Call("validate").Bool()
}
func (qkt *QuaternionKeyframeTrack) Parse(json js.Value) *animation.KeyframeTrack {
	return &animation.KeyframeTrack{Value: qkt.Call("parse", json)}
}
func (qkt *QuaternionKeyframeTrack) ToJSON(track *animation.KeyframeTrack) js.Value {
	return qkt.Call("toJSON", track)
}

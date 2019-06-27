// Code generated from "math/interpolants/LinearInterpolant.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// LinearInterpolant extend: [Interpolant]
type LinearInterpolant struct {
	js.Value
}

func NewLinearInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize float64, resultBuffer js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: get("LinearInterpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
}
func (li *LinearInterpolant) JSValue() js.Value {
	return li.Value
}
func (li *LinearInterpolant) ParameterPositions() js.Value {
	return li.Get("parameterPositions")
}
func (li *LinearInterpolant) SetParameterPositions(v js.Value) {
	li.Set("parameterPositions", v)
}
func (li *LinearInterpolant) ResultBuffer() js.Value {
	return li.Get("resultBuffer")
}
func (li *LinearInterpolant) SetResultBuffer(v js.Value) {
	li.Set("resultBuffer", v)
}
func (li *LinearInterpolant) SamplesValues() js.Value {
	return li.Get("samplesValues")
}
func (li *LinearInterpolant) SetSamplesValues(v js.Value) {
	li.Set("samplesValues", v)
}
func (li *LinearInterpolant) ValueSize() float64 {
	return li.Get("valueSize").Float()
}
func (li *LinearInterpolant) SetValueSize(v float64) {
	li.Set("valueSize", v)
}
func (li *LinearInterpolant) Evaluate(time float64) js.Value {
	return li.Call("evaluate", time)
}
func (li *LinearInterpolant) Interpolate_(i1 float64, t0 float64, t float64, t1 float64) js.Value {
	return li.Call("interpolate_", i1, t0, t, t1)
}

// Code generated from "math/interpolants/CubicInterpolant.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type CubicInterpolant struct {
	js.Value
}

func NewCubicInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize float64, resultBuffer js.Value) *CubicInterpolant {
	return &CubicInterpolant{Value: get("CubicInterpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
}
func (ci *CubicInterpolant) ParameterPositions() js.Value {
	return ci.Get("parameterPositions")
}
func (ci *CubicInterpolant) SetParameterPositions(v js.Value) {
	ci.Set("parameterPositions", v)
}
func (ci *CubicInterpolant) ResultBuffer() js.Value {
	return ci.Get("resultBuffer")
}
func (ci *CubicInterpolant) SetResultBuffer(v js.Value) {
	ci.Set("resultBuffer", v)
}
func (ci *CubicInterpolant) SamplesValues() js.Value {
	return ci.Get("samplesValues")
}
func (ci *CubicInterpolant) SetSamplesValues(v js.Value) {
	ci.Set("samplesValues", v)
}
func (ci *CubicInterpolant) ValueSize() float64 {
	return ci.Get("valueSize").Float()
}
func (ci *CubicInterpolant) SetValueSize(v float64) {
	ci.Set("valueSize", v)
}
func (ci *CubicInterpolant) Evaluate(time float64) js.Value {
	return ci.Call("evaluate", time)
}
func (ci *CubicInterpolant) Interpolate_(i1 float64, t0 float64, t float64, t1 float64) js.Value {
	return ci.Call("interpolate_", i1, t0, t, t1)
}

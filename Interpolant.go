// Code generated from "math/Interpolant.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Interpolant struct {
	js.Value
}

func NewInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize float64, resultBuffer js.Value) *Interpolant {
	return &Interpolant{Value: get("Interpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
}
func (ii *Interpolant) ParameterPositions() js.Value {
	return ii.Get("parameterPositions")
}
func (ii *Interpolant) SetParameterPositions(v js.Value) {
	ii.Set("parameterPositions", v)
}
func (ii *Interpolant) ResultBuffer() js.Value {
	return ii.Get("resultBuffer")
}
func (ii *Interpolant) SetResultBuffer(v js.Value) {
	ii.Set("resultBuffer", v)
}
func (ii *Interpolant) SamplesValues() js.Value {
	return ii.Get("samplesValues")
}
func (ii *Interpolant) SetSamplesValues(v js.Value) {
	ii.Set("samplesValues", v)
}
func (ii *Interpolant) ValueSize() float64 {
	return ii.Get("valueSize").Float()
}
func (ii *Interpolant) SetValueSize(v float64) {
	ii.Set("valueSize", v)
}
func (ii *Interpolant) Evaluate(time float64) js.Value {
	return ii.Call("evaluate", time)
}

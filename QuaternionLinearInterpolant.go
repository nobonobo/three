// Code generated from "math/interpolants/QuaternionLinearInterpolant.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type QuaternionLinearInterpolant struct {
	js.Value
}

func NewQuaternionLinearInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize float64, resultBuffer js.Value) *QuaternionLinearInterpolant {
	return &QuaternionLinearInterpolant{Value: get("QuaternionLinearInterpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
}
func (qli *QuaternionLinearInterpolant) ParameterPositions() js.Value {
	return qli.Get("parameterPositions")
}
func (qli *QuaternionLinearInterpolant) SetParameterPositions(v js.Value) {
	qli.Set("parameterPositions", v)
}
func (qli *QuaternionLinearInterpolant) ResultBuffer() js.Value {
	return qli.Get("resultBuffer")
}
func (qli *QuaternionLinearInterpolant) SetResultBuffer(v js.Value) {
	qli.Set("resultBuffer", v)
}
func (qli *QuaternionLinearInterpolant) SamplesValues() js.Value {
	return qli.Get("samplesValues")
}
func (qli *QuaternionLinearInterpolant) SetSamplesValues(v js.Value) {
	qli.Set("samplesValues", v)
}
func (qli *QuaternionLinearInterpolant) ValueSize() float64 {
	return qli.Get("valueSize").Float()
}
func (qli *QuaternionLinearInterpolant) SetValueSize(v float64) {
	qli.Set("valueSize", v)
}
func (qli *QuaternionLinearInterpolant) Evaluate(time float64) js.Value {
	return qli.Call("evaluate", time)
}
func (qli *QuaternionLinearInterpolant) Interpolate_(i1 float64, t0 float64, t float64, t1 float64) js.Value {
	return qli.Call("interpolate_", i1, t0, t, t1)
}

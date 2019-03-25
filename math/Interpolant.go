package math

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Interpolant struct {
	js.Value
}

func NewInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize int, resultBuffer js.Value) *Interpolant {
	return &Interpolant{Value: get("Interpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
}
func (i *Interpolant) ParameterPositions() js.Value {
	return i.Get("parameterPositions")
}
func (i *Interpolant) SetParameterPositions(v js.Value) {
	i.Set("parameterPositions", v)
}
func (i *Interpolant) ResultBuffer() js.Value {
	return i.Get("resultBuffer")
}
func (i *Interpolant) SetResultBuffer(v js.Value) {
	i.Set("resultBuffer", v)
}
func (i *Interpolant) SamplesValues() js.Value {
	return i.Get("samplesValues")
}
func (i *Interpolant) SetSamplesValues(v js.Value) {
	i.Set("samplesValues", v)
}
func (i *Interpolant) ValueSize() int {
	return i.Get("valueSize").Int()
}
func (i *Interpolant) SetValueSize(v int) {
	i.Set("valueSize", v)
}
func (i *Interpolant) Evaluate(time int) js.Value {
	return i.Call("evaluate", time)
}

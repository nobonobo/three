package interpolants

import (
	"github.com/gopherjs/gopherwasm/js"
)

type DiscreteInterpolant struct {
	js.Value
}

func NewDiscreteInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize int, resultBuffer js.Value) *DiscreteInterpolant {
	return &DiscreteInterpolant{Value: get("DiscreteInterpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
}
func (di *DiscreteInterpolant) ParameterPositions() js.Value {
	return di.Get("parameterPositions")
}
func (di *DiscreteInterpolant) SetParameterPositions(v js.Value) {
	di.Set("parameterPositions", v)
}
func (di *DiscreteInterpolant) ResultBuffer() js.Value {
	return di.Get("resultBuffer")
}
func (di *DiscreteInterpolant) SetResultBuffer(v js.Value) {
	di.Set("resultBuffer", v)
}
func (di *DiscreteInterpolant) SamplesValues() js.Value {
	return di.Get("samplesValues")
}
func (di *DiscreteInterpolant) SetSamplesValues(v js.Value) {
	di.Set("samplesValues", v)
}
func (di *DiscreteInterpolant) ValueSize() int {
	return di.Get("valueSize").Int()
}
func (di *DiscreteInterpolant) SetValueSize(v int) {
	di.Set("valueSize", v)
}
func (di *DiscreteInterpolant) Evaluate(time int) js.Value {
	return di.Call("evaluate", time)
}
func (di *DiscreteInterpolant) Interpolate_(i1 int, t0 int, t int, t1 int) js.Value {
	return di.Call("interpolate_", i1, t0, t, t1)
}

package interpolants

import (
	"github.com/gopherjs/gopherwasm/js"
)

type LinearInterpolant struct {
	js.Value
}

func NewLinearInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize int, resultBuffer js.Value) *LinearInterpolant {
	return &LinearInterpolant{Value: get("LinearInterpolant").New(parameterPositions, samplesValues, sampleSize, resultBuffer)}
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
func (li *LinearInterpolant) ValueSize() int {
	return li.Get("valueSize").Int()
}
func (li *LinearInterpolant) SetValueSize(v int) {
	li.Set("valueSize", v)
}
func (li *LinearInterpolant) Evaluate(time int) js.Value {
	return li.Call("evaluate", time)
}
func (li *LinearInterpolant) Interpolate_(i1 int, t0 int, t int, t1 int) js.Value {
	return li.Call("interpolate_", i1, t0, t, t1)
}

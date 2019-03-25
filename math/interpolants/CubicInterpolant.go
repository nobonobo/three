package interpolants

import (
	"github.com/gopherjs/gopherwasm/js"
)

type CubicInterpolant struct {
	js.Value
}

func NewCubicInterpolant(parameterPositions js.Value, samplesValues js.Value, sampleSize int, resultBuffer js.Value) *CubicInterpolant {
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
func (ci *CubicInterpolant) ValueSize() int {
	return ci.Get("valueSize").Int()
}
func (ci *CubicInterpolant) SetValueSize(v int) {
	ci.Set("valueSize", v)
}
func (ci *CubicInterpolant) Evaluate(time int) js.Value {
	return ci.Call("evaluate", time)
}
func (ci *CubicInterpolant) Interpolate_(i1 int, t0 int, t int, t1 int) js.Value {
	return ci.Call("interpolate_", i1, t0, t, t1)
}

package animation

import (
	"github.com/gopherjs/gopherwasm/js"
)

type PropertyMixer struct {
	js.Value
}

func NewPropertyMixer(binding js.Value, typeName string, valueSize int) *PropertyMixer {
	return &PropertyMixer{Value: get("PropertyMixer").New(binding, typeName, valueSize)}
}
func (pm *PropertyMixer) Binding() js.Value {
	return pm.Get("binding")
}
func (pm *PropertyMixer) SetBinding(v js.Value) {
	pm.Set("binding", v)
}
func (pm *PropertyMixer) Buffer() js.Value {
	return pm.Get("buffer")
}
func (pm *PropertyMixer) SetBuffer(v js.Value) {
	pm.Set("buffer", v)
}
func (pm *PropertyMixer) CumulativeWeight() int {
	return pm.Get("cumulativeWeight").Int()
}
func (pm *PropertyMixer) SetCumulativeWeight(v int) {
	pm.Set("cumulativeWeight", v)
}
func (pm *PropertyMixer) ReferenceCount() int {
	return pm.Get("referenceCount").Int()
}
func (pm *PropertyMixer) SetReferenceCount(v int) {
	pm.Set("referenceCount", v)
}
func (pm *PropertyMixer) UseCount() int {
	return pm.Get("useCount").Int()
}
func (pm *PropertyMixer) SetUseCount(v int) {
	pm.Set("useCount", v)
}
func (pm *PropertyMixer) ValueSize() int {
	return pm.Get("valueSize").Int()
}
func (pm *PropertyMixer) SetValueSize(v int) {
	pm.Set("valueSize", v)
}
func (pm *PropertyMixer) Accumulate(accuIndex int, weight int) {
	pm.Call("accumulate", accuIndex, weight)
}
func (pm *PropertyMixer) Apply(accuIndex int) {
	pm.Call("apply", accuIndex)
}
func (pm *PropertyMixer) RestoreOriginalState() {
	pm.Call("restoreOriginalState")
}
func (pm *PropertyMixer) SaveOriginalState() {
	pm.Call("saveOriginalState")
}

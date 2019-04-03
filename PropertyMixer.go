// Code generated from "animation/PropertyMixer.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type PropertyMixer struct {
	js.Value
}

func NewPropertyMixer(binding js.Value, typeName string, valueSize float64) *PropertyMixer {
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
func (pm *PropertyMixer) CumulativeWeight() float64 {
	return pm.Get("cumulativeWeight").Float()
}
func (pm *PropertyMixer) SetCumulativeWeight(v float64) {
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
func (pm *PropertyMixer) ValueSize() float64 {
	return pm.Get("valueSize").Float()
}
func (pm *PropertyMixer) SetValueSize(v float64) {
	pm.Set("valueSize", v)
}
func (pm *PropertyMixer) Accumulate(accuIndex int, weight float64) {
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

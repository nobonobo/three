// Code generated from "core/Layers.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Layers struct {
	js.Value
}

func NewLayers() *Layers {
	return &Layers{Value: get("Layers").New()}
}
func (ll *Layers) Mask() float64 {
	return ll.Get("mask").Float()
}
func (ll *Layers) SetMask(v float64) {
	ll.Set("mask", v)
}
func (ll *Layers) Disable(channel float64) {
	ll.Call("disable", channel)
}
func (ll *Layers) Enable(channel float64) {
	ll.Call("enable", channel)
}
func (ll *Layers) Set2(channel float64) {
	ll.Call("set", channel)
}
func (ll *Layers) Test(layers *Layers) bool {
	return ll.Call("test", layers).Bool()
}
func (ll *Layers) Toggle(channel float64) {
	ll.Call("toggle", channel)
}

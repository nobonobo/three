// Code generated from "math/Spherical.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Spherical extend: []
type Spherical struct {
	js.Value
}

func NewSpherical(radius float64, phi float64, theta float64) *Spherical {
	return &Spherical{Value: get("Spherical").New(radius, phi, theta)}
}
func (ss *Spherical) JSValue() js.Value {
	return ss.Value
}
func (ss *Spherical) Phi() float64 {
	return ss.Get("phi").Float()
}
func (ss *Spherical) SetPhi(v float64) {
	ss.Set("phi", v)
}
func (ss *Spherical) Radius() float64 {
	return ss.Get("radius").Float()
}
func (ss *Spherical) SetRadius(v float64) {
	ss.Set("radius", v)
}
func (ss *Spherical) Theta() float64 {
	return ss.Get("theta").Float()
}
func (ss *Spherical) SetTheta(v float64) {
	ss.Set("theta", v)
}
func (ss *Spherical) Clone() *Spherical {
	return &Spherical{Value: ss.Call("clone")}
}
func (ss *Spherical) Copy(other *Spherical) *Spherical {
	return &Spherical{Value: ss.Call("copy", other)}
}
func (ss *Spherical) MakeSafe() {
	ss.Call("makeSafe")
}
func (ss *Spherical) Set2(radius float64, phi float64, theta float64) *Spherical {
	return &Spherical{Value: ss.Call("set", radius, phi, theta)}
}
func (ss *Spherical) SetFromVector3(vec3 *Vector3) *Spherical {
	return &Spherical{Value: ss.Call("setFromVector3", vec3)}
}

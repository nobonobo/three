// Code generated from "math/Cylindrical.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Cylindrical struct {
	js.Value
}

func NewCylindrical(radius float64, theta float64, y float64) *Cylindrical {
	return &Cylindrical{Value: get("Cylindrical").New(radius, theta, y)}
}
func (cc *Cylindrical) Radius() float64 {
	return cc.Get("radius").Float()
}
func (cc *Cylindrical) SetRadius(v float64) {
	cc.Set("radius", v)
}
func (cc *Cylindrical) Theta() float64 {
	return cc.Get("theta").Float()
}
func (cc *Cylindrical) SetTheta(v float64) {
	cc.Set("theta", v)
}
func (cc *Cylindrical) Y() float64 {
	return cc.Get("y").Float()
}
func (cc *Cylindrical) SetY(v float64) {
	cc.Set("y", v)
}
func (cc *Cylindrical) Clone() *Cylindrical {
	return &Cylindrical{Value: cc.Call("clone")}
}
func (cc *Cylindrical) Copy(other *Cylindrical) *Cylindrical {
	return &Cylindrical{Value: cc.Call("copy", other)}
}
func (cc *Cylindrical) Set2(radius float64, theta float64, y float64) *Cylindrical {
	return &Cylindrical{Value: cc.Call("set", radius, theta, y)}
}
func (cc *Cylindrical) SetFromVector3(vec3 *Vector3) *Cylindrical {
	return &Cylindrical{Value: cc.Call("setFromVector3", vec3)}
}

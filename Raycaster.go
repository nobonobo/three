// Code generated from "core/Raycaster.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Intersection interface {
}
type RaycasterParameters interface {
}

// Raycaster extend: []
type Raycaster struct {
	js.Value
}

func NewRaycaster(origin *Vector3, direction *Vector3, near float64, far float64) *Raycaster {
	return &Raycaster{Value: get("Raycaster").New(origin.JSValue(), direction.JSValue(), near, far)}
}
func (rr *Raycaster) JSValue() js.Value {
	return rr.Value
}
func (rr *Raycaster) Far() float64 {
	return rr.Get("far").Float()
}
func (rr *Raycaster) SetFar(v float64) {
	rr.Set("far", v)
}
func (rr *Raycaster) LinePrecision() float64 {
	return rr.Get("linePrecision").Float()
}
func (rr *Raycaster) SetLinePrecision(v float64) {
	rr.Set("linePrecision", v)
}
func (rr *Raycaster) Near() float64 {
	return rr.Get("near").Float()
}
func (rr *Raycaster) SetNear(v float64) {
	rr.Set("near", v)
}
func (rr *Raycaster) Params() js.Value {
	return rr.Get("params")
}
func (rr *Raycaster) SetParams(v js.Value) {
	rr.Set("params", v)
}
func (rr *Raycaster) Ray() *Ray {
	return &Ray{Value: rr.Get("ray")}
}
func (rr *Raycaster) SetRay(v *Ray) {
	rr.Set("ray", v.JSValue())
}
func (rr *Raycaster) IntersectObject(object *Object3D, recursive bool, optionalTarget js.Value) js.Value {
	return rr.Call("intersectObject", object, recursive, optionalTarget)
}
func (rr *Raycaster) IntersectObjects(objects js.Value, recursive bool, optionalTarget js.Value) js.Value {
	return rr.Call("intersectObjects", objects, recursive, optionalTarget)
}
func (rr *Raycaster) Set2(origin *Vector3, direction *Vector3) {
	rr.Call("set", origin.JSValue(), direction.JSValue())
}
func (rr *Raycaster) SetFromCamera(coords js.Value, camera Camera) {
	rr.Call("setFromCamera", coords, camera.JSValue())
}

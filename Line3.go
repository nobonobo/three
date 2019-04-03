// Code generated from "math/Line3.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Line3 struct {
	js.Value
}

func NewLine3(start *Vector3, end *Vector3) *Line3 {
	return &Line3{Value: get("Line3").New(start, end)}
}
func (ll *Line3) End() *Vector3 {
	return &Vector3{Value: ll.Get("end")}
}
func (ll *Line3) SetEnd(v *Vector3) {
	ll.Set("end", v)
}
func (ll *Line3) Start() *Vector3 {
	return &Vector3{Value: ll.Get("start")}
}
func (ll *Line3) SetStart(v *Vector3) {
	ll.Set("start", v)
}
func (ll *Line3) ApplyMatrix4(matrix *Matrix4) *Line3 {
	return &Line3{Value: ll.Call("applyMatrix4", matrix)}
}
func (ll *Line3) At(t float64, target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("at", t, target)}
}
func (ll *Line3) Clone() *Line3 {
	return &Line3{Value: ll.Call("clone")}
}
func (ll *Line3) ClosestPointToPoint(point *Vector3, clampToLine bool, target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("closestPointToPoint", point, clampToLine, target)}
}
func (ll *Line3) ClosestPointToPointParameter(point *Vector3, clampToLine bool) float64 {
	return ll.Call("closestPointToPointParameter", point, clampToLine).Float()
}
func (ll *Line3) Copy(line *Line3) *Line3 {
	return &Line3{Value: ll.Call("copy", line)}
}
func (ll *Line3) Delta(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("delta", target)}
}
func (ll *Line3) Distance() float64 {
	return ll.Call("distance").Float()
}
func (ll *Line3) DistanceSq() float64 {
	return ll.Call("distanceSq").Float()
}
func (ll *Line3) Equals(line *Line3) bool {
	return ll.Call("equals", line).Bool()
}
func (ll *Line3) GetCenter(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getCenter", target)}
}
func (ll *Line3) Set2(start *Vector3, end *Vector3) *Line3 {
	return &Line3{Value: ll.Call("set", start, end)}
}

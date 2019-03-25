package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/math"
)

type Intersection interface {
}
type RaycasterParameters interface {
}
type Raycaster struct {
	js.Value
}

func NewRaycaster(origin *math.Vector3, direction *math.Vector3, near int, far int) *Raycaster {
	return &Raycaster{Value: get("Raycaster").New(origin, direction, near, far)}
}
func (r *Raycaster) Far() int {
	return r.Get("far").Int()
}
func (r *Raycaster) SetFar(v int) {
	r.Set("far", v)
}
func (r *Raycaster) LinePrecision() int {
	return r.Get("linePrecision").Int()
}
func (r *Raycaster) SetLinePrecision(v int) {
	r.Set("linePrecision", v)
}
func (r *Raycaster) Near() int {
	return r.Get("near").Int()
}
func (r *Raycaster) SetNear(v int) {
	r.Set("near", v)
}
func (r *Raycaster) Params() RaycasterParameters {
	return RaycasterParameters(r.Get("params"))
}
func (r *Raycaster) SetParams(v RaycasterParameters) {
	r.Set("params", v)
}
func (r *Raycaster) Ray() *math.Ray {
	return &math.Ray{Value: r.Get("ray")}
}
func (r *Raycaster) SetRay(v *math.Ray) {
	r.Set("ray", v)
}
func (r *Raycaster) IntersectObject(object *Object3D, recursive bool, optionalTarget []Intersection) []Intersection {
	return []Intersection(r.Call("intersectObject", object, recursive, optionalTarget))
}
func (r *Raycaster) IntersectObjects(objects []Object3D, recursive bool, optionalTarget []Intersection) []Intersection {
	return []Intersection(r.Call("intersectObjects", objects, recursive, optionalTarget))
}
func (r *Raycaster) Set(origin *math.Vector3, direction *math.Vector3) {
	r.Call("set", origin, direction)
}
func (r *Raycaster) SetFromCamera(coords js.Value, camera *cameras.Camera) {
	r.Call("setFromCamera", coords, camera)
}

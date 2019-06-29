// Code generated from "core/DirectGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// DirectGeometry extend: [EventDispatcher]
type DirectGeometry struct {
	js.Value
}

func NewDirectGeometry() *DirectGeometry {
	return &DirectGeometry{Value: get("DirectGeometry").New()}
}
func (dg *DirectGeometry) JSValue() js.Value {
	return dg.Value
}
func (dg *DirectGeometry) BoundingBox() *Box3 {
	return &Box3{Value: dg.Get("boundingBox")}
}
func (dg *DirectGeometry) SetBoundingBox(v *Box3) {
	dg.Set("boundingBox", v.JSValue())
}
func (dg *DirectGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: dg.Get("boundingSphere")}
}
func (dg *DirectGeometry) SetBoundingSphere(v *Sphere) {
	dg.Set("boundingSphere", v.JSValue())
}
func (dg *DirectGeometry) Colors() js.Value {
	return dg.Get("colors")
}
func (dg *DirectGeometry) SetColors(v js.Value) {
	dg.Set("colors", v)
}
func (dg *DirectGeometry) ColorsNeedUpdate() bool {
	return dg.Get("colorsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetColorsNeedUpdate(v bool) {
	dg.Set("colorsNeedUpdate", v)
}
func (dg *DirectGeometry) Groups() js.Value {
	return dg.Get("groups")
}
func (dg *DirectGeometry) SetGroups(v js.Value) {
	dg.Set("groups", v)
}
func (dg *DirectGeometry) GroupsNeedUpdate() bool {
	return dg.Get("groupsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetGroupsNeedUpdate(v bool) {
	dg.Set("groupsNeedUpdate", v)
}
func (dg *DirectGeometry) Id() int {
	return dg.Get("id").Int()
}
func (dg *DirectGeometry) SetId(v int) {
	dg.Set("id", v)
}
func (dg *DirectGeometry) Indices() js.Value {
	return dg.Get("indices")
}
func (dg *DirectGeometry) SetIndices(v js.Value) {
	dg.Set("indices", v)
}
func (dg *DirectGeometry) MorphTargets() js.Value {
	return dg.Get("morphTargets")
}
func (dg *DirectGeometry) SetMorphTargets(v js.Value) {
	dg.Set("morphTargets", v)
}
func (dg *DirectGeometry) Name() string {
	return dg.Get("name").String()
}
func (dg *DirectGeometry) SetName(v string) {
	dg.Set("name", v)
}
func (dg *DirectGeometry) Normals() js.Value {
	return dg.Get("normals")
}
func (dg *DirectGeometry) SetNormals(v js.Value) {
	dg.Set("normals", v)
}
func (dg *DirectGeometry) NormalsNeedUpdate() bool {
	return dg.Get("normalsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetNormalsNeedUpdate(v bool) {
	dg.Set("normalsNeedUpdate", v)
}
func (dg *DirectGeometry) SkinIndices() js.Value {
	return dg.Get("skinIndices")
}
func (dg *DirectGeometry) SetSkinIndices(v js.Value) {
	dg.Set("skinIndices", v)
}
func (dg *DirectGeometry) SkinWeights() js.Value {
	return dg.Get("skinWeights")
}
func (dg *DirectGeometry) SetSkinWeights(v js.Value) {
	dg.Set("skinWeights", v)
}
func (dg *DirectGeometry) Type() string {
	return dg.Get("type").String()
}
func (dg *DirectGeometry) SetType(v string) {
	dg.Set("type", v)
}
func (dg *DirectGeometry) Uuid() string {
	return dg.Get("uuid").String()
}
func (dg *DirectGeometry) SetUuid(v string) {
	dg.Set("uuid", v)
}
func (dg *DirectGeometry) Uvs() js.Value {
	return dg.Get("uvs")
}
func (dg *DirectGeometry) SetUvs(v js.Value) {
	dg.Set("uvs", v)
}
func (dg *DirectGeometry) Uvs2() js.Value {
	return dg.Get("uvs2")
}
func (dg *DirectGeometry) SetUvs2(v js.Value) {
	dg.Set("uvs2", v)
}
func (dg *DirectGeometry) UvsNeedUpdate() bool {
	return dg.Get("uvsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetUvsNeedUpdate(v bool) {
	dg.Set("uvsNeedUpdate", v)
}
func (dg *DirectGeometry) Vertices() js.Value {
	return dg.Get("vertices")
}
func (dg *DirectGeometry) SetVertices(v js.Value) {
	dg.Set("vertices", v)
}
func (dg *DirectGeometry) VerticesNeedUpdate() bool {
	return dg.Get("verticesNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetVerticesNeedUpdate(v bool) {
	dg.Set("verticesNeedUpdate", v)
}
func (dg *DirectGeometry) AddEventListener(typ string, listener js.Value) {
	dg.Call("addEventListener", typ, listener)
}
func (dg *DirectGeometry) ComputeBoundingBox() {
	dg.Call("computeBoundingBox")
}
func (dg *DirectGeometry) ComputeBoundingSphere() {
	dg.Call("computeBoundingSphere")
}
func (dg *DirectGeometry) ComputeGroups(geometry Geometry) {
	dg.Call("computeGroups", geometry.JSValue())
}
func (dg *DirectGeometry) DispatchEvent(event js.Value) {
	dg.Call("dispatchEvent", event)
}
func (dg *DirectGeometry) Dispose() {
	dg.Call("dispose")
}
func (dg *DirectGeometry) FromGeometry(geometry Geometry) *DirectGeometry {
	return &DirectGeometry{Value: dg.Call("fromGeometry", geometry.JSValue())}
}
func (dg *DirectGeometry) HasEventListener(typ string, listener js.Value) bool {
	return dg.Call("hasEventListener", typ, listener).Bool()
}
func (dg *DirectGeometry) RemoveEventListener(typ string, listener js.Value) {
	dg.Call("removeEventListener", typ, listener)
}

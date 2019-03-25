package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type DirectGeometry struct {
	js.Value
}

func NewDirectGeometry() *DirectGeometry {
	return &DirectGeometry{Value: get("DirectGeometry").New()}
}
func (dg *DirectGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: dg.Get("boundingBox")}
}
func (dg *DirectGeometry) SetBoundingBox(v *math.Box3) {
	dg.Set("boundingBox", v)
}
func (dg *DirectGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: dg.Get("boundingSphere")}
}
func (dg *DirectGeometry) SetBoundingSphere(v *math.Sphere) {
	dg.Set("boundingSphere", v)
}
func (dg *DirectGeometry) Colors() []math.Color {
	return []math.Color(dg.Get("colors"))
}
func (dg *DirectGeometry) SetColors(v []math.Color) {
	dg.Set("colors", v)
}
func (dg *DirectGeometry) ColorsNeedUpdate() bool {
	return dg.Get("colorsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetColorsNeedUpdate(v bool) {
	dg.Set("colorsNeedUpdate", v)
}
func (dg *DirectGeometry) Groups() []js.Value {
	return []js.Value(dg.Get("groups"))
}
func (dg *DirectGeometry) SetGroups(v []js.Value) {
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
func (dg *DirectGeometry) Indices() []int {
	return []int(dg.Get("indices"))
}
func (dg *DirectGeometry) SetIndices(v []int) {
	dg.Set("indices", v)
}
func (dg *DirectGeometry) MorphTargets() []MorphTarget {
	return []MorphTarget(dg.Get("morphTargets"))
}
func (dg *DirectGeometry) SetMorphTargets(v []MorphTarget) {
	dg.Set("morphTargets", v)
}
func (dg *DirectGeometry) Name() string {
	return dg.Get("name").String()
}
func (dg *DirectGeometry) SetName(v string) {
	dg.Set("name", v)
}
func (dg *DirectGeometry) Normals() []math.Vector3 {
	return []math.Vector3(dg.Get("normals"))
}
func (dg *DirectGeometry) SetNormals(v []math.Vector3) {
	dg.Set("normals", v)
}
func (dg *DirectGeometry) NormalsNeedUpdate() bool {
	return dg.Get("normalsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetNormalsNeedUpdate(v bool) {
	dg.Set("normalsNeedUpdate", v)
}
func (dg *DirectGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(dg.Get("skinIndices"))
}
func (dg *DirectGeometry) SetSkinIndices(v []math.Vector4) {
	dg.Set("skinIndices", v)
}
func (dg *DirectGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(dg.Get("skinWeights"))
}
func (dg *DirectGeometry) SetSkinWeights(v []math.Vector4) {
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
func (dg *DirectGeometry) Uvs() []math.Vector2 {
	return []math.Vector2(dg.Get("uvs"))
}
func (dg *DirectGeometry) SetUvs(v []math.Vector2) {
	dg.Set("uvs", v)
}
func (dg *DirectGeometry) Uvs2() []math.Vector2 {
	return []math.Vector2(dg.Get("uvs2"))
}
func (dg *DirectGeometry) SetUvs2(v []math.Vector2) {
	dg.Set("uvs2", v)
}
func (dg *DirectGeometry) UvsNeedUpdate() bool {
	return dg.Get("uvsNeedUpdate").Bool()
}
func (dg *DirectGeometry) SetUvsNeedUpdate(v bool) {
	dg.Set("uvsNeedUpdate", v)
}
func (dg *DirectGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(dg.Get("vertices"))
}
func (dg *DirectGeometry) SetVertices(v []math.Vector3) {
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
func (dg *DirectGeometry) ComputeGroups(geometry *Geometry) {
	dg.Call("computeGroups", geometry)
}
func (dg *DirectGeometry) DispatchEvent(event js.Value) {
	dg.Call("dispatchEvent", event)
}
func (dg *DirectGeometry) Dispose() {
	dg.Call("dispose")
}
func (dg *DirectGeometry) FromGeometry(geometry *Geometry) *DirectGeometry {
	return &DirectGeometry{Value: dg.Call("fromGeometry", geometry)}
}
func (dg *DirectGeometry) HasEventListener(typ string, listener js.Value) bool {
	return dg.Call("hasEventListener", typ, listener).Bool()
}
func (dg *DirectGeometry) RemoveEventListener(typ string, listener js.Value) {
	dg.Call("removeEventListener", typ, listener)
}

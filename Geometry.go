// Code generated from "core/Geometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func GeometryIdCount() int {
	return get("GeometryIdCount").Int()
}
func SetGeometryIdCount(v int) {
	set("GeometryIdCount", v)
}

type MorphColor interface {
}
type MorphNormals interface {
}
type MorphTarget interface {
}
type Geometry interface {
	JSValue() js.Value
	Animation() *AnimationClip
	SetAnimation(v *AnimationClip)
	Animations() js.Value
	SetAnimations(v js.Value)
	Bones() js.Value
	SetBones(v js.Value)
	BoundingBox() *Box3
	SetBoundingBox(v *Box3)
	BoundingSphere() *Sphere
	SetBoundingSphere(v *Sphere)
	Colors() js.Value
	SetColors(v js.Value)
	ColorsNeedUpdate() bool
	SetColorsNeedUpdate(v bool)
	ElementsNeedUpdate() bool
	SetElementsNeedUpdate(v bool)
	FaceVertexUvs() js.Value
	SetFaceVertexUvs(v js.Value)
	Faces() js.Value
	SetFaces(v js.Value)
	GroupsNeedUpdate() bool
	SetGroupsNeedUpdate(v bool)
	Id() int
	SetId(v int)
	LineDistances() js.Value
	SetLineDistances(v js.Value)
	LineDistancesNeedUpdate() bool
	SetLineDistancesNeedUpdate(v bool)
	MorphNormals() js.Value
	SetMorphNormals(v js.Value)
	MorphTargets() js.Value
	SetMorphTargets(v js.Value)
	Name() string
	SetName(v string)
	NormalsNeedUpdate() bool
	SetNormalsNeedUpdate(v bool)
	SkinIndices() js.Value
	SetSkinIndices(v js.Value)
	SkinWeights() js.Value
	SetSkinWeights(v js.Value)
	Type() string
	SetType(v string)
	Uuid() string
	SetUuid(v string)
	UvsNeedUpdate() bool
	SetUvsNeedUpdate(v bool)
	Vertices() js.Value
	SetVertices(v js.Value)
	VerticesNeedUpdate() bool
	SetVerticesNeedUpdate(v bool)
	AddEventListener(typ string, listener js.Value)
	ApplyMatrix(matrix *Matrix4) Geometry
	Center() Geometry
	Clone() Geometry
	ComputeBoundingBox()
	ComputeBoundingSphere()
	ComputeFaceNormals()
	ComputeFlatVertexNormals()
	ComputeMorphNormals()
	ComputeVertexNormals(areaWeighted bool)
	Copy(source Geometry) Geometry
	DispatchEvent(event js.Value)
	Dispose()
	FromBufferGeometry(geometry *BufferGeometry) Geometry
	HasEventListener(typ string, listener js.Value) bool
	LookAt(vector *Vector3)
	Merge(geometry Geometry, matrix Matrix, materialIndexOffset int)
	MergeMesh(mesh Mesh)
	MergeVertices() float64
	Normalize() Geometry
	RemoveEventListener(typ string, listener js.Value)
	RotateX(angle float64) Geometry
	RotateY(angle float64) Geometry
	RotateZ(angle float64) Geometry
	Scale(x float64, y float64, z float64) Geometry
	SetFromPoints(points js.Value) Geometry
	SortFacesByMaterialIndex()
	ToJSON() js.Value
	Translate(x float64, y float64, z float64) Geometry
}

// GeometryImpl extend: [EventDispatcher]
type GeometryImpl struct {
	js.Value
}

func NewGeometry() *GeometryImpl {
	return &GeometryImpl{Value: get("Geometry").New()}
}
func (gg *GeometryImpl) JSValue() js.Value {
	return gg.Value
}
func (gg *GeometryImpl) Animation() *AnimationClip {
	return &AnimationClip{Value: gg.Get("animation")}
}
func (gg *GeometryImpl) SetAnimation(v *AnimationClip) {
	gg.Set("animation", v.JSValue())
}
func (gg *GeometryImpl) Animations() js.Value {
	return gg.Get("animations")
}
func (gg *GeometryImpl) SetAnimations(v js.Value) {
	gg.Set("animations", v)
}
func (gg *GeometryImpl) Bones() js.Value {
	return gg.Get("bones")
}
func (gg *GeometryImpl) SetBones(v js.Value) {
	gg.Set("bones", v)
}
func (gg *GeometryImpl) BoundingBox() *Box3 {
	return &Box3{Value: gg.Get("boundingBox")}
}
func (gg *GeometryImpl) SetBoundingBox(v *Box3) {
	gg.Set("boundingBox", v.JSValue())
}
func (gg *GeometryImpl) BoundingSphere() *Sphere {
	return &Sphere{Value: gg.Get("boundingSphere")}
}
func (gg *GeometryImpl) SetBoundingSphere(v *Sphere) {
	gg.Set("boundingSphere", v.JSValue())
}
func (gg *GeometryImpl) Colors() js.Value {
	return gg.Get("colors")
}
func (gg *GeometryImpl) SetColors(v js.Value) {
	gg.Set("colors", v)
}
func (gg *GeometryImpl) ColorsNeedUpdate() bool {
	return gg.Get("colorsNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetColorsNeedUpdate(v bool) {
	gg.Set("colorsNeedUpdate", v)
}
func (gg *GeometryImpl) ElementsNeedUpdate() bool {
	return gg.Get("elementsNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetElementsNeedUpdate(v bool) {
	gg.Set("elementsNeedUpdate", v)
}
func (gg *GeometryImpl) FaceVertexUvs() js.Value {
	return gg.Get("faceVertexUvs")
}
func (gg *GeometryImpl) SetFaceVertexUvs(v js.Value) {
	gg.Set("faceVertexUvs", v)
}
func (gg *GeometryImpl) Faces() js.Value {
	return gg.Get("faces")
}
func (gg *GeometryImpl) SetFaces(v js.Value) {
	gg.Set("faces", v)
}
func (gg *GeometryImpl) GroupsNeedUpdate() bool {
	return gg.Get("groupsNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetGroupsNeedUpdate(v bool) {
	gg.Set("groupsNeedUpdate", v)
}
func (gg *GeometryImpl) Id() int {
	return gg.Get("id").Int()
}
func (gg *GeometryImpl) SetId(v int) {
	gg.Set("id", v)
}
func (gg *GeometryImpl) LineDistances() js.Value {
	return gg.Get("lineDistances")
}
func (gg *GeometryImpl) SetLineDistances(v js.Value) {
	gg.Set("lineDistances", v)
}
func (gg *GeometryImpl) LineDistancesNeedUpdate() bool {
	return gg.Get("lineDistancesNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetLineDistancesNeedUpdate(v bool) {
	gg.Set("lineDistancesNeedUpdate", v)
}
func (gg *GeometryImpl) MorphNormals() js.Value {
	return gg.Get("morphNormals")
}
func (gg *GeometryImpl) SetMorphNormals(v js.Value) {
	gg.Set("morphNormals", v)
}
func (gg *GeometryImpl) MorphTargets() js.Value {
	return gg.Get("morphTargets")
}
func (gg *GeometryImpl) SetMorphTargets(v js.Value) {
	gg.Set("morphTargets", v)
}
func (gg *GeometryImpl) Name() string {
	return gg.Get("name").String()
}
func (gg *GeometryImpl) SetName(v string) {
	gg.Set("name", v)
}
func (gg *GeometryImpl) NormalsNeedUpdate() bool {
	return gg.Get("normalsNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetNormalsNeedUpdate(v bool) {
	gg.Set("normalsNeedUpdate", v)
}
func (gg *GeometryImpl) SkinIndices() js.Value {
	return gg.Get("skinIndices")
}
func (gg *GeometryImpl) SetSkinIndices(v js.Value) {
	gg.Set("skinIndices", v)
}
func (gg *GeometryImpl) SkinWeights() js.Value {
	return gg.Get("skinWeights")
}
func (gg *GeometryImpl) SetSkinWeights(v js.Value) {
	gg.Set("skinWeights", v)
}
func (gg *GeometryImpl) Type() string {
	return gg.Get("type").String()
}
func (gg *GeometryImpl) SetType(v string) {
	gg.Set("type", v)
}
func (gg *GeometryImpl) Uuid() string {
	return gg.Get("uuid").String()
}
func (gg *GeometryImpl) SetUuid(v string) {
	gg.Set("uuid", v)
}
func (gg *GeometryImpl) UvsNeedUpdate() bool {
	return gg.Get("uvsNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetUvsNeedUpdate(v bool) {
	gg.Set("uvsNeedUpdate", v)
}
func (gg *GeometryImpl) Vertices() js.Value {
	return gg.Get("vertices")
}
func (gg *GeometryImpl) SetVertices(v js.Value) {
	gg.Set("vertices", v)
}
func (gg *GeometryImpl) VerticesNeedUpdate() bool {
	return gg.Get("verticesNeedUpdate").Bool()
}
func (gg *GeometryImpl) SetVerticesNeedUpdate(v bool) {
	gg.Set("verticesNeedUpdate", v)
}
func (gg *GeometryImpl) AddEventListener(typ string, listener js.Value) {
	gg.Call("addEventListener", typ, listener)
}
func (gg *GeometryImpl) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: gg.Call("applyMatrix", matrix)}
}
func (gg *GeometryImpl) Center() Geometry {
	return &GeometryImpl{Value: gg.Call("center")}
}
func (gg *GeometryImpl) Clone() Geometry {
	return &GeometryImpl{Value: gg.Call("clone")}
}
func (gg *GeometryImpl) ComputeBoundingBox() {
	gg.Call("computeBoundingBox")
}
func (gg *GeometryImpl) ComputeBoundingSphere() {
	gg.Call("computeBoundingSphere")
}
func (gg *GeometryImpl) ComputeFaceNormals() {
	gg.Call("computeFaceNormals")
}
func (gg *GeometryImpl) ComputeFlatVertexNormals() {
	gg.Call("computeFlatVertexNormals")
}
func (gg *GeometryImpl) ComputeMorphNormals() {
	gg.Call("computeMorphNormals")
}
func (gg *GeometryImpl) ComputeVertexNormals(areaWeighted bool) {
	gg.Call("computeVertexNormals", areaWeighted)
}
func (gg *GeometryImpl) Copy(source Geometry) Geometry {
	return &GeometryImpl{Value: gg.Call("copy", source.JSValue())}
}
func (gg *GeometryImpl) DispatchEvent(event js.Value) {
	gg.Call("dispatchEvent", event)
}
func (gg *GeometryImpl) Dispose() {
	gg.Call("dispose")
}
func (gg *GeometryImpl) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: gg.Call("fromBufferGeometry", geometry)}
}
func (gg *GeometryImpl) HasEventListener(typ string, listener js.Value) bool {
	return gg.Call("hasEventListener", typ, listener).Bool()
}
func (gg *GeometryImpl) LookAt(vector *Vector3) {
	gg.Call("lookAt", vector.JSValue())
}
func (gg *GeometryImpl) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	gg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (gg *GeometryImpl) MergeMesh(mesh Mesh) {
	gg.Call("mergeMesh", mesh.JSValue())
}
func (gg *GeometryImpl) MergeVertices() float64 {
	return gg.Call("mergeVertices").Float()
}
func (gg *GeometryImpl) Normalize() Geometry {
	return &GeometryImpl{Value: gg.Call("normalize")}
}
func (gg *GeometryImpl) RemoveEventListener(typ string, listener js.Value) {
	gg.Call("removeEventListener", typ, listener)
}
func (gg *GeometryImpl) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: gg.Call("rotateX", angle)}
}
func (gg *GeometryImpl) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: gg.Call("rotateY", angle)}
}
func (gg *GeometryImpl) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: gg.Call("rotateZ", angle)}
}
func (gg *GeometryImpl) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: gg.Call("scale", x, y, z)}
}
func (gg *GeometryImpl) SetFromPoints(points js.Value) Geometry {
	return &GeometryImpl{Value: gg.Call("setFromPoints", points)}
}
func (gg *GeometryImpl) SortFacesByMaterialIndex() {
	gg.Call("sortFacesByMaterialIndex")
}
func (gg *GeometryImpl) ToJSON() js.Value {
	return gg.Call("toJSON")
}
func (gg *GeometryImpl) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: gg.Call("translate", x, y, z)}
}

// Code generated from "core/Geometry.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
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
type Geometry struct {
	js.Value
}

func NewGeometry() *Geometry {
	return &Geometry{Value: get("Geometry").New()}
}
func (gg *Geometry) Animation() *AnimationClip {
	return &AnimationClip{Value: gg.Get("animation")}
}
func (gg *Geometry) SetAnimation(v *AnimationClip) {
	gg.Set("animation", v)
}
func (gg *Geometry) Animations() js.Value {
	return gg.Get("animations")
}
func (gg *Geometry) SetAnimations(v js.Value) {
	gg.Set("animations", v)
}
func (gg *Geometry) Bones() js.Value {
	return gg.Get("bones")
}
func (gg *Geometry) SetBones(v js.Value) {
	gg.Set("bones", v)
}
func (gg *Geometry) BoundingBox() *Box3 {
	return &Box3{Value: gg.Get("boundingBox")}
}
func (gg *Geometry) SetBoundingBox(v *Box3) {
	gg.Set("boundingBox", v)
}
func (gg *Geometry) BoundingSphere() *Sphere {
	return &Sphere{Value: gg.Get("boundingSphere")}
}
func (gg *Geometry) SetBoundingSphere(v *Sphere) {
	gg.Set("boundingSphere", v)
}
func (gg *Geometry) Colors() js.Value {
	return gg.Get("colors")
}
func (gg *Geometry) SetColors(v js.Value) {
	gg.Set("colors", v)
}
func (gg *Geometry) ColorsNeedUpdate() bool {
	return gg.Get("colorsNeedUpdate").Bool()
}
func (gg *Geometry) SetColorsNeedUpdate(v bool) {
	gg.Set("colorsNeedUpdate", v)
}
func (gg *Geometry) ElementsNeedUpdate() bool {
	return gg.Get("elementsNeedUpdate").Bool()
}
func (gg *Geometry) SetElementsNeedUpdate(v bool) {
	gg.Set("elementsNeedUpdate", v)
}
func (gg *Geometry) FaceVertexUvs() js.Value {
	return gg.Get("faceVertexUvs")
}
func (gg *Geometry) SetFaceVertexUvs(v js.Value) {
	gg.Set("faceVertexUvs", v)
}
func (gg *Geometry) Faces() js.Value {
	return gg.Get("faces")
}
func (gg *Geometry) SetFaces(v js.Value) {
	gg.Set("faces", v)
}
func (gg *Geometry) GroupsNeedUpdate() bool {
	return gg.Get("groupsNeedUpdate").Bool()
}
func (gg *Geometry) SetGroupsNeedUpdate(v bool) {
	gg.Set("groupsNeedUpdate", v)
}
func (gg *Geometry) Id() int {
	return gg.Get("id").Int()
}
func (gg *Geometry) SetId(v int) {
	gg.Set("id", v)
}
func (gg *Geometry) LineDistances() js.Value {
	return gg.Get("lineDistances")
}
func (gg *Geometry) SetLineDistances(v js.Value) {
	gg.Set("lineDistances", v)
}
func (gg *Geometry) LineDistancesNeedUpdate() bool {
	return gg.Get("lineDistancesNeedUpdate").Bool()
}
func (gg *Geometry) SetLineDistancesNeedUpdate(v bool) {
	gg.Set("lineDistancesNeedUpdate", v)
}
func (gg *Geometry) MorphNormals() js.Value {
	return gg.Get("morphNormals")
}
func (gg *Geometry) SetMorphNormals(v js.Value) {
	gg.Set("morphNormals", v)
}
func (gg *Geometry) MorphTargets() js.Value {
	return gg.Get("morphTargets")
}
func (gg *Geometry) SetMorphTargets(v js.Value) {
	gg.Set("morphTargets", v)
}
func (gg *Geometry) Name() string {
	return gg.Get("name").String()
}
func (gg *Geometry) SetName(v string) {
	gg.Set("name", v)
}
func (gg *Geometry) NormalsNeedUpdate() bool {
	return gg.Get("normalsNeedUpdate").Bool()
}
func (gg *Geometry) SetNormalsNeedUpdate(v bool) {
	gg.Set("normalsNeedUpdate", v)
}
func (gg *Geometry) SkinIndices() js.Value {
	return gg.Get("skinIndices")
}
func (gg *Geometry) SetSkinIndices(v js.Value) {
	gg.Set("skinIndices", v)
}
func (gg *Geometry) SkinWeights() js.Value {
	return gg.Get("skinWeights")
}
func (gg *Geometry) SetSkinWeights(v js.Value) {
	gg.Set("skinWeights", v)
}
func (gg *Geometry) Type() string {
	return gg.Get("type").String()
}
func (gg *Geometry) SetType(v string) {
	gg.Set("type", v)
}
func (gg *Geometry) Uuid() string {
	return gg.Get("uuid").String()
}
func (gg *Geometry) SetUuid(v string) {
	gg.Set("uuid", v)
}
func (gg *Geometry) UvsNeedUpdate() bool {
	return gg.Get("uvsNeedUpdate").Bool()
}
func (gg *Geometry) SetUvsNeedUpdate(v bool) {
	gg.Set("uvsNeedUpdate", v)
}
func (gg *Geometry) Vertices() js.Value {
	return gg.Get("vertices")
}
func (gg *Geometry) SetVertices(v js.Value) {
	gg.Set("vertices", v)
}
func (gg *Geometry) VerticesNeedUpdate() bool {
	return gg.Get("verticesNeedUpdate").Bool()
}
func (gg *Geometry) SetVerticesNeedUpdate(v bool) {
	gg.Set("verticesNeedUpdate", v)
}
func (gg *Geometry) AddEventListener(typ string, listener js.Value) {
	gg.Call("addEventListener", typ, listener)
}
func (gg *Geometry) ApplyMatrix(matrix *Matrix4) *Geometry {
	return &Geometry{Value: gg.Call("applyMatrix", matrix)}
}
func (gg *Geometry) Center() *Geometry {
	return &Geometry{Value: gg.Call("center")}
}
func (gg *Geometry) Clone() *Geometry {
	return &Geometry{Value: gg.Call("clone")}
}
func (gg *Geometry) ComputeBoundingBox() {
	gg.Call("computeBoundingBox")
}
func (gg *Geometry) ComputeBoundingSphere() {
	gg.Call("computeBoundingSphere")
}
func (gg *Geometry) ComputeFaceNormals() {
	gg.Call("computeFaceNormals")
}
func (gg *Geometry) ComputeFlatVertexNormals() {
	gg.Call("computeFlatVertexNormals")
}
func (gg *Geometry) ComputeMorphNormals() {
	gg.Call("computeMorphNormals")
}
func (gg *Geometry) ComputeVertexNormals(areaWeighted bool) {
	gg.Call("computeVertexNormals", areaWeighted)
}
func (gg *Geometry) Copy(source *Geometry) *Geometry {
	return &Geometry{Value: gg.Call("copy", source)}
}
func (gg *Geometry) DispatchEvent(event js.Value) {
	gg.Call("dispatchEvent", event)
}
func (gg *Geometry) Dispose() {
	gg.Call("dispose")
}
func (gg *Geometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: gg.Call("fromBufferGeometry", geometry)}
}
func (gg *Geometry) HasEventListener(typ string, listener js.Value) bool {
	return gg.Call("hasEventListener", typ, listener).Bool()
}
func (gg *Geometry) LookAt(vector *Vector3) {
	gg.Call("lookAt", vector)
}
func (gg *Geometry) Merge(geometry *Geometry, matrix Matrix, materialIndexOffset int) {
	gg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (gg *Geometry) MergeMesh(mesh *Mesh) {
	gg.Call("mergeMesh", mesh)
}
func (gg *Geometry) MergeVertices() float64 {
	return gg.Call("mergeVertices").Float()
}
func (gg *Geometry) Normalize() *Geometry {
	return &Geometry{Value: gg.Call("normalize")}
}
func (gg *Geometry) RemoveEventListener(typ string, listener js.Value) {
	gg.Call("removeEventListener", typ, listener)
}
func (gg *Geometry) RotateX(angle float64) *Geometry {
	return &Geometry{Value: gg.Call("rotateX", angle)}
}
func (gg *Geometry) RotateY(angle float64) *Geometry {
	return &Geometry{Value: gg.Call("rotateY", angle)}
}
func (gg *Geometry) RotateZ(angle float64) *Geometry {
	return &Geometry{Value: gg.Call("rotateZ", angle)}
}
func (gg *Geometry) Scale(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: gg.Call("scale", x, y, z)}
}
func (gg *Geometry) SetFromPoints(points js.Value) *Geometry {
	return &Geometry{Value: gg.Call("setFromPoints", points)}
}
func (gg *Geometry) SortFacesByMaterialIndex() {
	gg.Call("sortFacesByMaterialIndex")
}
func (gg *Geometry) ToJSON() js.Value {
	return gg.Call("toJSON")
}
func (gg *Geometry) Translate(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: gg.Call("translate", x, y, z)}
}

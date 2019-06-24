// Code generated from "geometries/IcosahedronGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type IcosahedronBufferGeometry struct {
	js.Value
}

func NewIcosahedronBufferGeometry(radius float64, detail float64) *IcosahedronBufferGeometry {
	return &IcosahedronBufferGeometry{Value: get("IcosahedronBufferGeometry").New(radius, detail)}
}
func (ibg *IcosahedronBufferGeometry) Attributes() js.Value {
	return ibg.Get("attributes")
}
func (ibg *IcosahedronBufferGeometry) SetAttributes(v js.Value) {
	ibg.Set("attributes", v)
}
func (ibg *IcosahedronBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: ibg.Get("boundingBox")}
}
func (ibg *IcosahedronBufferGeometry) SetBoundingBox(v *Box3) {
	ibg.Set("boundingBox", v)
}
func (ibg *IcosahedronBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: ibg.Get("boundingSphere")}
}
func (ibg *IcosahedronBufferGeometry) SetBoundingSphere(v *Sphere) {
	ibg.Set("boundingSphere", v)
}
func (ibg *IcosahedronBufferGeometry) DrawRange() js.Value {
	return ibg.Get("drawRange")
}
func (ibg *IcosahedronBufferGeometry) SetDrawRange(v js.Value) {
	ibg.Set("drawRange", v)
}
func (ibg *IcosahedronBufferGeometry) Drawcalls() js.Value {
	return ibg.Get("drawcalls")
}
func (ibg *IcosahedronBufferGeometry) SetDrawcalls(v js.Value) {
	ibg.Set("drawcalls", v)
}
func (ibg *IcosahedronBufferGeometry) Groups() js.Value {
	return ibg.Get("groups")
}
func (ibg *IcosahedronBufferGeometry) SetGroups(v js.Value) {
	ibg.Set("groups", v)
}
func (ibg *IcosahedronBufferGeometry) Id() int {
	return ibg.Get("id").Int()
}
func (ibg *IcosahedronBufferGeometry) SetId(v int) {
	ibg.Set("id", v)
}
func (ibg *IcosahedronBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: ibg.Get("index")}
}
func (ibg *IcosahedronBufferGeometry) SetIndex(v *BufferAttribute) {
	ibg.Set("index", v)
}
func (ibg *IcosahedronBufferGeometry) MorphAttributes() js.Value {
	return ibg.Get("morphAttributes")
}
func (ibg *IcosahedronBufferGeometry) SetMorphAttributes(v js.Value) {
	ibg.Set("morphAttributes", v)
}
func (ibg *IcosahedronBufferGeometry) Name() string {
	return ibg.Get("name").String()
}
func (ibg *IcosahedronBufferGeometry) SetName(v string) {
	ibg.Set("name", v)
}
func (ibg *IcosahedronBufferGeometry) Offsets() js.Value {
	return ibg.Get("offsets")
}
func (ibg *IcosahedronBufferGeometry) SetOffsets(v js.Value) {
	ibg.Set("offsets", v)
}
func (ibg *IcosahedronBufferGeometry) Parameters() js.Value {
	return ibg.Get("parameters")
}
func (ibg *IcosahedronBufferGeometry) SetParameters(v js.Value) {
	ibg.Set("parameters", v)
}
func (ibg *IcosahedronBufferGeometry) Type() string {
	return ibg.Get("type").String()
}
func (ibg *IcosahedronBufferGeometry) SetType(v string) {
	ibg.Set("type", v)
}
func (ibg *IcosahedronBufferGeometry) UserData() js.Value {
	return ibg.Get("userData")
}
func (ibg *IcosahedronBufferGeometry) SetUserData(v js.Value) {
	ibg.Set("userData", v)
}
func (ibg *IcosahedronBufferGeometry) Uuid() string {
	return ibg.Get("uuid").String()
}
func (ibg *IcosahedronBufferGeometry) SetUuid(v string) {
	ibg.Set("uuid", v)
}
func (ibg *IcosahedronBufferGeometry) MaxIndex() int {
	return ibg.Get("MaxIndex").Int()
}
func (ibg *IcosahedronBufferGeometry) SetMaxIndex(v int) {
	ibg.Set("MaxIndex", v)
}
func (ibg *IcosahedronBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("addAttribute", name, attribute)}
}
func (ibg *IcosahedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return ibg.Call("addAttribute", name, array, itemSize)
}
func (ibg *IcosahedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	ibg.Call("addDrawCall", start, count, indexOffset)
}
func (ibg *IcosahedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	ibg.Call("addEventListener", typ, listener)
}
func (ibg *IcosahedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	ibg.Call("addGroup", start, count, materialIndex)
}
func (ibg *IcosahedronBufferGeometry) AddIndex(index js.Value) {
	ibg.Call("addIndex", index)
}
func (ibg *IcosahedronBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("applyMatrix", matrix)}
}
func (ibg *IcosahedronBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("center")}
}
func (ibg *IcosahedronBufferGeometry) ClearDrawCalls() {
	ibg.Call("clearDrawCalls")
}
func (ibg *IcosahedronBufferGeometry) ClearGroups() {
	ibg.Call("clearGroups")
}
func (ibg *IcosahedronBufferGeometry) Clone() *IcosahedronBufferGeometry {
	return &IcosahedronBufferGeometry{Value: ibg.Call("clone")}
}
func (ibg *IcosahedronBufferGeometry) ComputeBoundingBox() {
	ibg.Call("computeBoundingBox")
}
func (ibg *IcosahedronBufferGeometry) ComputeBoundingSphere() {
	ibg.Call("computeBoundingSphere")
}
func (ibg *IcosahedronBufferGeometry) ComputeVertexNormals() {
	ibg.Call("computeVertexNormals")
}
func (ibg *IcosahedronBufferGeometry) Copy(source *BufferGeometry) *IcosahedronBufferGeometry {
	return &IcosahedronBufferGeometry{Value: ibg.Call("copy", source)}
}
func (ibg *IcosahedronBufferGeometry) DispatchEvent(event js.Value) {
	ibg.Call("dispatchEvent", event)
}
func (ibg *IcosahedronBufferGeometry) Dispose() {
	ibg.Call("dispose")
}
func (ibg *IcosahedronBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("fromDirectGeometry", geometry)}
}
func (ibg *IcosahedronBufferGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("fromGeometry", geometry, settings)}
}
func (ibg *IcosahedronBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: ibg.Call("getAttribute", name)}
}
func (ibg *IcosahedronBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: ibg.Call("getIndex")}
}
func (ibg *IcosahedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ibg.Call("hasEventListener", typ, listener).Bool()
}
func (ibg *IcosahedronBufferGeometry) LookAt(v *Vector3) {
	ibg.Call("lookAt", v)
}
func (ibg *IcosahedronBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("merge", geometry, offset)}
}
func (ibg *IcosahedronBufferGeometry) NormalizeNormals() {
	ibg.Call("normalizeNormals")
}
func (ibg *IcosahedronBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("removeAttribute", name)}
}
func (ibg *IcosahedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	ibg.Call("removeEventListener", typ, listener)
}
func (ibg *IcosahedronBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("rotateX", angle)}
}
func (ibg *IcosahedronBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("rotateY", angle)}
}
func (ibg *IcosahedronBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("rotateZ", angle)}
}
func (ibg *IcosahedronBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("scale", x, y, z)}
}
func (ibg *IcosahedronBufferGeometry) SetDrawRange2(start int, count int) {
	ibg.Call("setDrawRange", start, count)
}
func (ibg *IcosahedronBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("setFromObject", object)}
}
func (ibg *IcosahedronBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("setFromPoints", points)}
}
func (ibg *IcosahedronBufferGeometry) SetIndex2(index *BufferAttribute) {
	ibg.Call("setIndex", index)
}
func (ibg *IcosahedronBufferGeometry) ToJSON() js.Value {
	return ibg.Call("toJSON")
}
func (ibg *IcosahedronBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("toNonIndexed")}
}
func (ibg *IcosahedronBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("translate", x, y, z)}
}
func (ibg *IcosahedronBufferGeometry) UpdateFromObject(object *Object3D) {
	ibg.Call("updateFromObject", object)
}

type IcosahedronGeometry struct {
	js.Value
}

func NewIcosahedronGeometry(radius float64, detail float64) *IcosahedronGeometry {
	return &IcosahedronGeometry{Value: get("IcosahedronGeometry").New(radius, detail)}
}
func (ig *IcosahedronGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: ig.Get("animation")}
}
func (ig *IcosahedronGeometry) SetAnimation(v *AnimationClip) {
	ig.Set("animation", v)
}
func (ig *IcosahedronGeometry) Animations() js.Value {
	return ig.Get("animations")
}
func (ig *IcosahedronGeometry) SetAnimations(v js.Value) {
	ig.Set("animations", v)
}
func (ig *IcosahedronGeometry) Bones() js.Value {
	return ig.Get("bones")
}
func (ig *IcosahedronGeometry) SetBones(v js.Value) {
	ig.Set("bones", v)
}
func (ig *IcosahedronGeometry) BoundingBox() *Box3 {
	return &Box3{Value: ig.Get("boundingBox")}
}
func (ig *IcosahedronGeometry) SetBoundingBox(v *Box3) {
	ig.Set("boundingBox", v)
}
func (ig *IcosahedronGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: ig.Get("boundingSphere")}
}
func (ig *IcosahedronGeometry) SetBoundingSphere(v *Sphere) {
	ig.Set("boundingSphere", v)
}
func (ig *IcosahedronGeometry) Colors() js.Value {
	return ig.Get("colors")
}
func (ig *IcosahedronGeometry) SetColors(v js.Value) {
	ig.Set("colors", v)
}
func (ig *IcosahedronGeometry) ColorsNeedUpdate() bool {
	return ig.Get("colorsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetColorsNeedUpdate(v bool) {
	ig.Set("colorsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) ElementsNeedUpdate() bool {
	return ig.Get("elementsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetElementsNeedUpdate(v bool) {
	ig.Set("elementsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) FaceVertexUvs() js.Value {
	return ig.Get("faceVertexUvs")
}
func (ig *IcosahedronGeometry) SetFaceVertexUvs(v js.Value) {
	ig.Set("faceVertexUvs", v)
}
func (ig *IcosahedronGeometry) Faces() js.Value {
	return ig.Get("faces")
}
func (ig *IcosahedronGeometry) SetFaces(v js.Value) {
	ig.Set("faces", v)
}
func (ig *IcosahedronGeometry) GroupsNeedUpdate() bool {
	return ig.Get("groupsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetGroupsNeedUpdate(v bool) {
	ig.Set("groupsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) Id() int {
	return ig.Get("id").Int()
}
func (ig *IcosahedronGeometry) SetId(v int) {
	ig.Set("id", v)
}
func (ig *IcosahedronGeometry) LineDistances() js.Value {
	return ig.Get("lineDistances")
}
func (ig *IcosahedronGeometry) SetLineDistances(v js.Value) {
	ig.Set("lineDistances", v)
}
func (ig *IcosahedronGeometry) LineDistancesNeedUpdate() bool {
	return ig.Get("lineDistancesNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	ig.Set("lineDistancesNeedUpdate", v)
}
func (ig *IcosahedronGeometry) MorphNormals() js.Value {
	return ig.Get("morphNormals")
}
func (ig *IcosahedronGeometry) SetMorphNormals(v js.Value) {
	ig.Set("morphNormals", v)
}
func (ig *IcosahedronGeometry) MorphTargets() js.Value {
	return ig.Get("morphTargets")
}
func (ig *IcosahedronGeometry) SetMorphTargets(v js.Value) {
	ig.Set("morphTargets", v)
}
func (ig *IcosahedronGeometry) Name() string {
	return ig.Get("name").String()
}
func (ig *IcosahedronGeometry) SetName(v string) {
	ig.Set("name", v)
}
func (ig *IcosahedronGeometry) NormalsNeedUpdate() bool {
	return ig.Get("normalsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetNormalsNeedUpdate(v bool) {
	ig.Set("normalsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) Parameters() js.Value {
	return ig.Get("parameters")
}
func (ig *IcosahedronGeometry) SetParameters(v js.Value) {
	ig.Set("parameters", v)
}
func (ig *IcosahedronGeometry) SkinIndices() js.Value {
	return ig.Get("skinIndices")
}
func (ig *IcosahedronGeometry) SetSkinIndices(v js.Value) {
	ig.Set("skinIndices", v)
}
func (ig *IcosahedronGeometry) SkinWeights() js.Value {
	return ig.Get("skinWeights")
}
func (ig *IcosahedronGeometry) SetSkinWeights(v js.Value) {
	ig.Set("skinWeights", v)
}
func (ig *IcosahedronGeometry) Type() string {
	return ig.Get("type").String()
}
func (ig *IcosahedronGeometry) SetType(v string) {
	ig.Set("type", v)
}
func (ig *IcosahedronGeometry) Uuid() string {
	return ig.Get("uuid").String()
}
func (ig *IcosahedronGeometry) SetUuid(v string) {
	ig.Set("uuid", v)
}
func (ig *IcosahedronGeometry) UvsNeedUpdate() bool {
	return ig.Get("uvsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetUvsNeedUpdate(v bool) {
	ig.Set("uvsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) Vertices() js.Value {
	return ig.Get("vertices")
}
func (ig *IcosahedronGeometry) SetVertices(v js.Value) {
	ig.Set("vertices", v)
}
func (ig *IcosahedronGeometry) VerticesNeedUpdate() bool {
	return ig.Get("verticesNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetVerticesNeedUpdate(v bool) {
	ig.Set("verticesNeedUpdate", v)
}
func (ig *IcosahedronGeometry) AddEventListener(typ string, listener js.Value) {
	ig.Call("addEventListener", typ, listener)
}
func (ig *IcosahedronGeometry) ApplyMatrix(matrix *Matrix4) *Geometry {
	return &Geometry{Value: ig.Call("applyMatrix", matrix)}
}
func (ig *IcosahedronGeometry) Center() *Geometry {
	return &Geometry{Value: ig.Call("center")}
}
func (ig *IcosahedronGeometry) Clone() *IcosahedronGeometry {
	return &IcosahedronGeometry{Value: ig.Call("clone")}
}
func (ig *IcosahedronGeometry) ComputeBoundingBox() {
	ig.Call("computeBoundingBox")
}
func (ig *IcosahedronGeometry) ComputeBoundingSphere() {
	ig.Call("computeBoundingSphere")
}
func (ig *IcosahedronGeometry) ComputeFaceNormals() {
	ig.Call("computeFaceNormals")
}
func (ig *IcosahedronGeometry) ComputeFlatVertexNormals() {
	ig.Call("computeFlatVertexNormals")
}
func (ig *IcosahedronGeometry) ComputeMorphNormals() {
	ig.Call("computeMorphNormals")
}
func (ig *IcosahedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	ig.Call("computeVertexNormals", areaWeighted)
}
func (ig *IcosahedronGeometry) Copy(source *Geometry) *IcosahedronGeometry {
	return &IcosahedronGeometry{Value: ig.Call("copy", source)}
}
func (ig *IcosahedronGeometry) DispatchEvent(event js.Value) {
	ig.Call("dispatchEvent", event)
}
func (ig *IcosahedronGeometry) Dispose() {
	ig.Call("dispose")
}
func (ig *IcosahedronGeometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: ig.Call("fromBufferGeometry", geometry)}
}
func (ig *IcosahedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ig.Call("hasEventListener", typ, listener).Bool()
}
func (ig *IcosahedronGeometry) LookAt(vector *Vector3) {
	ig.Call("lookAt", vector)
}
func (ig *IcosahedronGeometry) Merge(geometry *Geometry, matrix Matrix, materialIndexOffset int) {
	ig.Call("merge", geometry, matrix, materialIndexOffset)
}
func (ig *IcosahedronGeometry) MergeMesh(mesh *Mesh) {
	ig.Call("mergeMesh", mesh)
}
func (ig *IcosahedronGeometry) MergeVertices() float64 {
	return ig.Call("mergeVertices").Float()
}
func (ig *IcosahedronGeometry) Normalize() *Geometry {
	return &Geometry{Value: ig.Call("normalize")}
}
func (ig *IcosahedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	ig.Call("removeEventListener", typ, listener)
}
func (ig *IcosahedronGeometry) RotateX(angle float64) *Geometry {
	return &Geometry{Value: ig.Call("rotateX", angle)}
}
func (ig *IcosahedronGeometry) RotateY(angle float64) *Geometry {
	return &Geometry{Value: ig.Call("rotateY", angle)}
}
func (ig *IcosahedronGeometry) RotateZ(angle float64) *Geometry {
	return &Geometry{Value: ig.Call("rotateZ", angle)}
}
func (ig *IcosahedronGeometry) Scale(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: ig.Call("scale", x, y, z)}
}
func (ig *IcosahedronGeometry) SetFromPoints(points js.Value) *IcosahedronGeometry {
	return &IcosahedronGeometry{Value: ig.Call("setFromPoints", points)}
}
func (ig *IcosahedronGeometry) SortFacesByMaterialIndex() {
	ig.Call("sortFacesByMaterialIndex")
}
func (ig *IcosahedronGeometry) ToJSON() js.Value {
	return ig.Call("toJSON")
}
func (ig *IcosahedronGeometry) Translate(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: ig.Call("translate", x, y, z)}
}

// Code generated from "geometries/LatheGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// LatheBufferGeometry extend: [BufferGeometry]
type LatheBufferGeometry struct {
	js.Value
}

func NewLatheBufferGeometry(points js.Value, segments int, phiStart float64, phiLength float64) *LatheBufferGeometry {
	return &LatheBufferGeometry{Value: get("LatheBufferGeometry").New(points, segments, phiStart, phiLength)}
}
func (lbg *LatheBufferGeometry) JSValue() js.Value {
	return lbg.Value
}
func (lbg *LatheBufferGeometry) Attributes() js.Value {
	return lbg.Get("attributes")
}
func (lbg *LatheBufferGeometry) SetAttributes(v js.Value) {
	lbg.Set("attributes", v)
}
func (lbg *LatheBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: lbg.Get("boundingBox")}
}
func (lbg *LatheBufferGeometry) SetBoundingBox(v *Box3) {
	lbg.Set("boundingBox", v.JSValue())
}
func (lbg *LatheBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: lbg.Get("boundingSphere")}
}
func (lbg *LatheBufferGeometry) SetBoundingSphere(v *Sphere) {
	lbg.Set("boundingSphere", v.JSValue())
}
func (lbg *LatheBufferGeometry) DrawRange() js.Value {
	return lbg.Get("drawRange")
}
func (lbg *LatheBufferGeometry) SetDrawRange(v js.Value) {
	lbg.Set("drawRange", v)
}
func (lbg *LatheBufferGeometry) Drawcalls() js.Value {
	return lbg.Get("drawcalls")
}
func (lbg *LatheBufferGeometry) SetDrawcalls(v js.Value) {
	lbg.Set("drawcalls", v)
}
func (lbg *LatheBufferGeometry) Groups() js.Value {
	return lbg.Get("groups")
}
func (lbg *LatheBufferGeometry) SetGroups(v js.Value) {
	lbg.Set("groups", v)
}
func (lbg *LatheBufferGeometry) Id() int {
	return lbg.Get("id").Int()
}
func (lbg *LatheBufferGeometry) SetId(v int) {
	lbg.Set("id", v)
}
func (lbg *LatheBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: lbg.Get("index")}
}
func (lbg *LatheBufferGeometry) SetIndex(v *BufferAttribute) {
	lbg.Set("index", v.JSValue())
}
func (lbg *LatheBufferGeometry) MorphAttributes() js.Value {
	return lbg.Get("morphAttributes")
}
func (lbg *LatheBufferGeometry) SetMorphAttributes(v js.Value) {
	lbg.Set("morphAttributes", v)
}
func (lbg *LatheBufferGeometry) Name() string {
	return lbg.Get("name").String()
}
func (lbg *LatheBufferGeometry) SetName(v string) {
	lbg.Set("name", v)
}
func (lbg *LatheBufferGeometry) Offsets() js.Value {
	return lbg.Get("offsets")
}
func (lbg *LatheBufferGeometry) SetOffsets(v js.Value) {
	lbg.Set("offsets", v)
}
func (lbg *LatheBufferGeometry) Parameters() js.Value {
	return lbg.Get("parameters")
}
func (lbg *LatheBufferGeometry) SetParameters(v js.Value) {
	lbg.Set("parameters", v)
}
func (lbg *LatheBufferGeometry) Type() string {
	return lbg.Get("type").String()
}
func (lbg *LatheBufferGeometry) SetType(v string) {
	lbg.Set("type", v)
}
func (lbg *LatheBufferGeometry) UserData() js.Value {
	return lbg.Get("userData")
}
func (lbg *LatheBufferGeometry) SetUserData(v js.Value) {
	lbg.Set("userData", v)
}
func (lbg *LatheBufferGeometry) Uuid() string {
	return lbg.Get("uuid").String()
}
func (lbg *LatheBufferGeometry) SetUuid(v string) {
	lbg.Set("uuid", v)
}
func (lbg *LatheBufferGeometry) MaxIndex() int {
	return lbg.Get("MaxIndex").Int()
}
func (lbg *LatheBufferGeometry) SetMaxIndex(v int) {
	lbg.Set("MaxIndex", v)
}
func (lbg *LatheBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("addAttribute", name, attribute)}
}
func (lbg *LatheBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return lbg.Call("addAttribute", name, array, itemSize)
}
func (lbg *LatheBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	lbg.Call("addDrawCall", start, count, indexOffset)
}
func (lbg *LatheBufferGeometry) AddEventListener(typ string, listener js.Value) {
	lbg.Call("addEventListener", typ, listener)
}
func (lbg *LatheBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	lbg.Call("addGroup", start, count, materialIndex)
}
func (lbg *LatheBufferGeometry) AddIndex(index js.Value) {
	lbg.Call("addIndex", index)
}
func (lbg *LatheBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("applyMatrix", matrix)}
}
func (lbg *LatheBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("center")}
}
func (lbg *LatheBufferGeometry) ClearDrawCalls() {
	lbg.Call("clearDrawCalls")
}
func (lbg *LatheBufferGeometry) ClearGroups() {
	lbg.Call("clearGroups")
}
func (lbg *LatheBufferGeometry) Clone() *LatheBufferGeometry {
	return &LatheBufferGeometry{Value: lbg.Call("clone")}
}
func (lbg *LatheBufferGeometry) ComputeBoundingBox() {
	lbg.Call("computeBoundingBox")
}
func (lbg *LatheBufferGeometry) ComputeBoundingSphere() {
	lbg.Call("computeBoundingSphere")
}
func (lbg *LatheBufferGeometry) ComputeVertexNormals() {
	lbg.Call("computeVertexNormals")
}
func (lbg *LatheBufferGeometry) Copy(source *BufferGeometry) *LatheBufferGeometry {
	return &LatheBufferGeometry{Value: lbg.Call("copy", source)}
}
func (lbg *LatheBufferGeometry) DispatchEvent(event js.Value) {
	lbg.Call("dispatchEvent", event)
}
func (lbg *LatheBufferGeometry) Dispose() {
	lbg.Call("dispose")
}
func (lbg *LatheBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("fromDirectGeometry", geometry)}
}
func (lbg *LatheBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (lbg *LatheBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: lbg.Call("getAttribute", name)}
}
func (lbg *LatheBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: lbg.Call("getIndex")}
}
func (lbg *LatheBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return lbg.Call("hasEventListener", typ, listener).Bool()
}
func (lbg *LatheBufferGeometry) LookAt(v *Vector3) {
	lbg.Call("lookAt", v.JSValue())
}
func (lbg *LatheBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("merge", geometry, offset)}
}
func (lbg *LatheBufferGeometry) NormalizeNormals() {
	lbg.Call("normalizeNormals")
}
func (lbg *LatheBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("removeAttribute", name)}
}
func (lbg *LatheBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	lbg.Call("removeEventListener", typ, listener)
}
func (lbg *LatheBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("rotateX", angle)}
}
func (lbg *LatheBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("rotateY", angle)}
}
func (lbg *LatheBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("rotateZ", angle)}
}
func (lbg *LatheBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("scale", x, y, z)}
}
func (lbg *LatheBufferGeometry) SetDrawRange2(start int, count int) {
	lbg.Call("setDrawRange", start, count)
}
func (lbg *LatheBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("setFromObject", object)}
}
func (lbg *LatheBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("setFromPoints", points)}
}
func (lbg *LatheBufferGeometry) SetIndex2(index *BufferAttribute) {
	lbg.Call("setIndex", index)
}
func (lbg *LatheBufferGeometry) ToJSON() js.Value {
	return lbg.Call("toJSON")
}
func (lbg *LatheBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("toNonIndexed")}
}
func (lbg *LatheBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: lbg.Call("translate", x, y, z)}
}
func (lbg *LatheBufferGeometry) UpdateFromObject(object *Object3D) {
	lbg.Call("updateFromObject", object.JSValue())
}

// LatheGeometry extend: [Geometry]
type LatheGeometry struct {
	js.Value
}

func NewLatheGeometry(points js.Value, segments int, phiStart float64, phiLength float64) *LatheGeometry {
	return &LatheGeometry{Value: get("LatheGeometry").New(points, segments, phiStart, phiLength)}
}
func (lg *LatheGeometry) JSValue() js.Value {
	return lg.Value
}
func (lg *LatheGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: lg.Get("animation")}
}
func (lg *LatheGeometry) SetAnimation(v *AnimationClip) {
	lg.Set("animation", v.JSValue())
}
func (lg *LatheGeometry) Animations() js.Value {
	return lg.Get("animations")
}
func (lg *LatheGeometry) SetAnimations(v js.Value) {
	lg.Set("animations", v)
}
func (lg *LatheGeometry) Bones() js.Value {
	return lg.Get("bones")
}
func (lg *LatheGeometry) SetBones(v js.Value) {
	lg.Set("bones", v)
}
func (lg *LatheGeometry) BoundingBox() *Box3 {
	return &Box3{Value: lg.Get("boundingBox")}
}
func (lg *LatheGeometry) SetBoundingBox(v *Box3) {
	lg.Set("boundingBox", v.JSValue())
}
func (lg *LatheGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: lg.Get("boundingSphere")}
}
func (lg *LatheGeometry) SetBoundingSphere(v *Sphere) {
	lg.Set("boundingSphere", v.JSValue())
}
func (lg *LatheGeometry) Colors() js.Value {
	return lg.Get("colors")
}
func (lg *LatheGeometry) SetColors(v js.Value) {
	lg.Set("colors", v)
}
func (lg *LatheGeometry) ColorsNeedUpdate() bool {
	return lg.Get("colorsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetColorsNeedUpdate(v bool) {
	lg.Set("colorsNeedUpdate", v)
}
func (lg *LatheGeometry) ElementsNeedUpdate() bool {
	return lg.Get("elementsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetElementsNeedUpdate(v bool) {
	lg.Set("elementsNeedUpdate", v)
}
func (lg *LatheGeometry) FaceVertexUvs() js.Value {
	return lg.Get("faceVertexUvs")
}
func (lg *LatheGeometry) SetFaceVertexUvs(v js.Value) {
	lg.Set("faceVertexUvs", v)
}
func (lg *LatheGeometry) Faces() js.Value {
	return lg.Get("faces")
}
func (lg *LatheGeometry) SetFaces(v js.Value) {
	lg.Set("faces", v)
}
func (lg *LatheGeometry) GroupsNeedUpdate() bool {
	return lg.Get("groupsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetGroupsNeedUpdate(v bool) {
	lg.Set("groupsNeedUpdate", v)
}
func (lg *LatheGeometry) Id() int {
	return lg.Get("id").Int()
}
func (lg *LatheGeometry) SetId(v int) {
	lg.Set("id", v)
}
func (lg *LatheGeometry) LineDistances() js.Value {
	return lg.Get("lineDistances")
}
func (lg *LatheGeometry) SetLineDistances(v js.Value) {
	lg.Set("lineDistances", v)
}
func (lg *LatheGeometry) LineDistancesNeedUpdate() bool {
	return lg.Get("lineDistancesNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetLineDistancesNeedUpdate(v bool) {
	lg.Set("lineDistancesNeedUpdate", v)
}
func (lg *LatheGeometry) MorphNormals() js.Value {
	return lg.Get("morphNormals")
}
func (lg *LatheGeometry) SetMorphNormals(v js.Value) {
	lg.Set("morphNormals", v)
}
func (lg *LatheGeometry) MorphTargets() js.Value {
	return lg.Get("morphTargets")
}
func (lg *LatheGeometry) SetMorphTargets(v js.Value) {
	lg.Set("morphTargets", v)
}
func (lg *LatheGeometry) Name() string {
	return lg.Get("name").String()
}
func (lg *LatheGeometry) SetName(v string) {
	lg.Set("name", v)
}
func (lg *LatheGeometry) NormalsNeedUpdate() bool {
	return lg.Get("normalsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetNormalsNeedUpdate(v bool) {
	lg.Set("normalsNeedUpdate", v)
}
func (lg *LatheGeometry) Parameters() js.Value {
	return lg.Get("parameters")
}
func (lg *LatheGeometry) SetParameters(v js.Value) {
	lg.Set("parameters", v)
}
func (lg *LatheGeometry) SkinIndices() js.Value {
	return lg.Get("skinIndices")
}
func (lg *LatheGeometry) SetSkinIndices(v js.Value) {
	lg.Set("skinIndices", v)
}
func (lg *LatheGeometry) SkinWeights() js.Value {
	return lg.Get("skinWeights")
}
func (lg *LatheGeometry) SetSkinWeights(v js.Value) {
	lg.Set("skinWeights", v)
}
func (lg *LatheGeometry) Type() string {
	return lg.Get("type").String()
}
func (lg *LatheGeometry) SetType(v string) {
	lg.Set("type", v)
}
func (lg *LatheGeometry) Uuid() string {
	return lg.Get("uuid").String()
}
func (lg *LatheGeometry) SetUuid(v string) {
	lg.Set("uuid", v)
}
func (lg *LatheGeometry) UvsNeedUpdate() bool {
	return lg.Get("uvsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetUvsNeedUpdate(v bool) {
	lg.Set("uvsNeedUpdate", v)
}
func (lg *LatheGeometry) Vertices() js.Value {
	return lg.Get("vertices")
}
func (lg *LatheGeometry) SetVertices(v js.Value) {
	lg.Set("vertices", v)
}
func (lg *LatheGeometry) VerticesNeedUpdate() bool {
	return lg.Get("verticesNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetVerticesNeedUpdate(v bool) {
	lg.Set("verticesNeedUpdate", v)
}
func (lg *LatheGeometry) AddEventListener(typ string, listener js.Value) {
	lg.Call("addEventListener", typ, listener)
}
func (lg *LatheGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: lg.Call("applyMatrix", matrix)}
}
func (lg *LatheGeometry) Center() Geometry {
	return &GeometryImpl{Value: lg.Call("center")}
}
func (lg *LatheGeometry) Clone() Geometry {
	return &LatheGeometry{Value: lg.Call("clone")}
}
func (lg *LatheGeometry) ComputeBoundingBox() {
	lg.Call("computeBoundingBox")
}
func (lg *LatheGeometry) ComputeBoundingSphere() {
	lg.Call("computeBoundingSphere")
}
func (lg *LatheGeometry) ComputeFaceNormals() {
	lg.Call("computeFaceNormals")
}
func (lg *LatheGeometry) ComputeFlatVertexNormals() {
	lg.Call("computeFlatVertexNormals")
}
func (lg *LatheGeometry) ComputeMorphNormals() {
	lg.Call("computeMorphNormals")
}
func (lg *LatheGeometry) ComputeVertexNormals(areaWeighted bool) {
	lg.Call("computeVertexNormals", areaWeighted)
}
func (lg *LatheGeometry) Copy(source Geometry) Geometry {
	return &LatheGeometry{Value: lg.Call("copy", source.JSValue())}
}
func (lg *LatheGeometry) DispatchEvent(event js.Value) {
	lg.Call("dispatchEvent", event)
}
func (lg *LatheGeometry) Dispose() {
	lg.Call("dispose")
}
func (lg *LatheGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: lg.Call("fromBufferGeometry", geometry)}
}
func (lg *LatheGeometry) HasEventListener(typ string, listener js.Value) bool {
	return lg.Call("hasEventListener", typ, listener).Bool()
}
func (lg *LatheGeometry) LookAt(vector *Vector3) {
	lg.Call("lookAt", vector.JSValue())
}
func (lg *LatheGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	lg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (lg *LatheGeometry) MergeMesh(mesh Mesh) {
	lg.Call("mergeMesh", mesh.JSValue())
}
func (lg *LatheGeometry) MergeVertices() float64 {
	return lg.Call("mergeVertices").Float()
}
func (lg *LatheGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: lg.Call("normalize")}
}
func (lg *LatheGeometry) RemoveEventListener(typ string, listener js.Value) {
	lg.Call("removeEventListener", typ, listener)
}
func (lg *LatheGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: lg.Call("rotateX", angle)}
}
func (lg *LatheGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: lg.Call("rotateY", angle)}
}
func (lg *LatheGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: lg.Call("rotateZ", angle)}
}
func (lg *LatheGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: lg.Call("scale", x, y, z)}
}
func (lg *LatheGeometry) SetFromPoints(points js.Value) Geometry {
	return &LatheGeometry{Value: lg.Call("setFromPoints", points)}
}
func (lg *LatheGeometry) SortFacesByMaterialIndex() {
	lg.Call("sortFacesByMaterialIndex")
}
func (lg *LatheGeometry) ToJSON() js.Value {
	return lg.Call("toJSON")
}
func (lg *LatheGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: lg.Call("translate", x, y, z)}
}

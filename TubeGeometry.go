// Code generated from "geometries/TubeGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// TubeBufferGeometry extend: [BufferGeometry]
type TubeBufferGeometry struct {
	js.Value
}

func NewTubeBufferGeometry(path js.Value, tubularSegments int, radius float64, radiusSegments int, closed bool) *TubeBufferGeometry {
	return &TubeBufferGeometry{Value: get("TubeBufferGeometry").New(path, tubularSegments, radius, radiusSegments, closed)}
}
func (tbg *TubeBufferGeometry) JSValue() js.Value {
	return tbg.Value
}
func (tbg *TubeBufferGeometry) Attributes() js.Value {
	return tbg.Get("attributes")
}
func (tbg *TubeBufferGeometry) SetAttributes(v js.Value) {
	tbg.Set("attributes", v)
}
func (tbg *TubeBufferGeometry) Binormals() js.Value {
	return tbg.Get("binormals")
}
func (tbg *TubeBufferGeometry) SetBinormals(v js.Value) {
	tbg.Set("binormals", v)
}
func (tbg *TubeBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tbg.Get("boundingBox")}
}
func (tbg *TubeBufferGeometry) SetBoundingBox(v *Box3) {
	tbg.Set("boundingBox", v.Value)
}
func (tbg *TubeBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tbg.Get("boundingSphere")}
}
func (tbg *TubeBufferGeometry) SetBoundingSphere(v *Sphere) {
	tbg.Set("boundingSphere", v.Value)
}
func (tbg *TubeBufferGeometry) DrawRange() js.Value {
	return tbg.Get("drawRange")
}
func (tbg *TubeBufferGeometry) SetDrawRange(v js.Value) {
	tbg.Set("drawRange", v)
}
func (tbg *TubeBufferGeometry) Drawcalls() js.Value {
	return tbg.Get("drawcalls")
}
func (tbg *TubeBufferGeometry) SetDrawcalls(v js.Value) {
	tbg.Set("drawcalls", v)
}
func (tbg *TubeBufferGeometry) Groups() js.Value {
	return tbg.Get("groups")
}
func (tbg *TubeBufferGeometry) SetGroups(v js.Value) {
	tbg.Set("groups", v)
}
func (tbg *TubeBufferGeometry) Id() int {
	return tbg.Get("id").Int()
}
func (tbg *TubeBufferGeometry) SetId(v int) {
	tbg.Set("id", v)
}
func (tbg *TubeBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Get("index")}
}
func (tbg *TubeBufferGeometry) SetIndex(v *BufferAttribute) {
	tbg.Set("index", v.Value)
}
func (tbg *TubeBufferGeometry) MorphAttributes() js.Value {
	return tbg.Get("morphAttributes")
}
func (tbg *TubeBufferGeometry) SetMorphAttributes(v js.Value) {
	tbg.Set("morphAttributes", v)
}
func (tbg *TubeBufferGeometry) Name() string {
	return tbg.Get("name").String()
}
func (tbg *TubeBufferGeometry) SetName(v string) {
	tbg.Set("name", v)
}
func (tbg *TubeBufferGeometry) Normals() js.Value {
	return tbg.Get("normals")
}
func (tbg *TubeBufferGeometry) SetNormals(v js.Value) {
	tbg.Set("normals", v)
}
func (tbg *TubeBufferGeometry) Offsets() js.Value {
	return tbg.Get("offsets")
}
func (tbg *TubeBufferGeometry) SetOffsets(v js.Value) {
	tbg.Set("offsets", v)
}
func (tbg *TubeBufferGeometry) Parameters() js.Value {
	return tbg.Get("parameters")
}
func (tbg *TubeBufferGeometry) SetParameters(v js.Value) {
	tbg.Set("parameters", v)
}
func (tbg *TubeBufferGeometry) Tangents() js.Value {
	return tbg.Get("tangents")
}
func (tbg *TubeBufferGeometry) SetTangents(v js.Value) {
	tbg.Set("tangents", v)
}
func (tbg *TubeBufferGeometry) Type() string {
	return tbg.Get("type").String()
}
func (tbg *TubeBufferGeometry) SetType(v string) {
	tbg.Set("type", v)
}
func (tbg *TubeBufferGeometry) UserData() js.Value {
	return tbg.Get("userData")
}
func (tbg *TubeBufferGeometry) SetUserData(v js.Value) {
	tbg.Set("userData", v)
}
func (tbg *TubeBufferGeometry) Uuid() string {
	return tbg.Get("uuid").String()
}
func (tbg *TubeBufferGeometry) SetUuid(v string) {
	tbg.Set("uuid", v)
}
func (tbg *TubeBufferGeometry) MaxIndex() int {
	return tbg.Get("MaxIndex").Int()
}
func (tbg *TubeBufferGeometry) SetMaxIndex(v int) {
	tbg.Set("MaxIndex", v)
}
func (tbg *TubeBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("addAttribute", name, attribute)}
}
func (tbg *TubeBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tbg.Call("addAttribute", name, array, itemSize)
}
func (tbg *TubeBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tbg.Call("addDrawCall", start, count, indexOffset)
}
func (tbg *TubeBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tbg.Call("addEventListener", typ, listener)
}
func (tbg *TubeBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tbg.Call("addGroup", start, count, materialIndex)
}
func (tbg *TubeBufferGeometry) AddIndex(index js.Value) {
	tbg.Call("addIndex", index)
}
func (tbg *TubeBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("applyMatrix", matrix)}
}
func (tbg *TubeBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("center")}
}
func (tbg *TubeBufferGeometry) ClearDrawCalls() {
	tbg.Call("clearDrawCalls")
}
func (tbg *TubeBufferGeometry) ClearGroups() {
	tbg.Call("clearGroups")
}
func (tbg *TubeBufferGeometry) Clone() *TubeBufferGeometry {
	return &TubeBufferGeometry{Value: tbg.Call("clone")}
}
func (tbg *TubeBufferGeometry) ComputeBoundingBox() {
	tbg.Call("computeBoundingBox")
}
func (tbg *TubeBufferGeometry) ComputeBoundingSphere() {
	tbg.Call("computeBoundingSphere")
}
func (tbg *TubeBufferGeometry) ComputeVertexNormals() {
	tbg.Call("computeVertexNormals")
}
func (tbg *TubeBufferGeometry) Copy(source *BufferGeometry) *TubeBufferGeometry {
	return &TubeBufferGeometry{Value: tbg.Call("copy", source)}
}
func (tbg *TubeBufferGeometry) DispatchEvent(event js.Value) {
	tbg.Call("dispatchEvent", event)
}
func (tbg *TubeBufferGeometry) Dispose() {
	tbg.Call("dispose")
}
func (tbg *TubeBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("fromDirectGeometry", geometry)}
}
func (tbg *TubeBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (tbg *TubeBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getAttribute", name)}
}
func (tbg *TubeBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getIndex")}
}
func (tbg *TubeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tbg.Call("hasEventListener", typ, listener).Bool()
}
func (tbg *TubeBufferGeometry) LookAt(v *Vector3) {
	tbg.Call("lookAt", v)
}
func (tbg *TubeBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("merge", geometry, offset)}
}
func (tbg *TubeBufferGeometry) NormalizeNormals() {
	tbg.Call("normalizeNormals")
}
func (tbg *TubeBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("removeAttribute", name)}
}
func (tbg *TubeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tbg.Call("removeEventListener", typ, listener)
}
func (tbg *TubeBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateX", angle)}
}
func (tbg *TubeBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateY", angle)}
}
func (tbg *TubeBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateZ", angle)}
}
func (tbg *TubeBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("scale", x, y, z)}
}
func (tbg *TubeBufferGeometry) SetDrawRange2(start int, count int) {
	tbg.Call("setDrawRange", start, count)
}
func (tbg *TubeBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("setFromObject", object)}
}
func (tbg *TubeBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("setFromPoints", points)}
}
func (tbg *TubeBufferGeometry) SetIndex2(index *BufferAttribute) {
	tbg.Call("setIndex", index)
}
func (tbg *TubeBufferGeometry) ToJSON() js.Value {
	return tbg.Call("toJSON")
}
func (tbg *TubeBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("toNonIndexed")}
}
func (tbg *TubeBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("translate", x, y, z)}
}
func (tbg *TubeBufferGeometry) UpdateFromObject(object *Object3D) {
	tbg.Call("updateFromObject", object)
}

// TubeGeometry extend: [Geometry]
type TubeGeometry struct {
	js.Value
}

func NewTubeGeometry(path js.Value, tubularSegments int, radius float64, radiusSegments int, closed bool) *TubeGeometry {
	return &TubeGeometry{Value: get("TubeGeometry").New(path, tubularSegments, radius, radiusSegments, closed)}
}
func (tg *TubeGeometry) JSValue() js.Value {
	return tg.Value
}
func (tg *TubeGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: tg.Get("animation")}
}
func (tg *TubeGeometry) SetAnimation(v *AnimationClip) {
	tg.Set("animation", v.Value)
}
func (tg *TubeGeometry) Animations() js.Value {
	return tg.Get("animations")
}
func (tg *TubeGeometry) SetAnimations(v js.Value) {
	tg.Set("animations", v)
}
func (tg *TubeGeometry) Binormals() js.Value {
	return tg.Get("binormals")
}
func (tg *TubeGeometry) SetBinormals(v js.Value) {
	tg.Set("binormals", v)
}
func (tg *TubeGeometry) Bones() js.Value {
	return tg.Get("bones")
}
func (tg *TubeGeometry) SetBones(v js.Value) {
	tg.Set("bones", v)
}
func (tg *TubeGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tg.Get("boundingBox")}
}
func (tg *TubeGeometry) SetBoundingBox(v *Box3) {
	tg.Set("boundingBox", v.Value)
}
func (tg *TubeGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tg.Get("boundingSphere")}
}
func (tg *TubeGeometry) SetBoundingSphere(v *Sphere) {
	tg.Set("boundingSphere", v.Value)
}
func (tg *TubeGeometry) Colors() js.Value {
	return tg.Get("colors")
}
func (tg *TubeGeometry) SetColors(v js.Value) {
	tg.Set("colors", v)
}
func (tg *TubeGeometry) ColorsNeedUpdate() bool {
	return tg.Get("colorsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetColorsNeedUpdate(v bool) {
	tg.Set("colorsNeedUpdate", v)
}
func (tg *TubeGeometry) ElementsNeedUpdate() bool {
	return tg.Get("elementsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetElementsNeedUpdate(v bool) {
	tg.Set("elementsNeedUpdate", v)
}
func (tg *TubeGeometry) FaceVertexUvs() js.Value {
	return tg.Get("faceVertexUvs")
}
func (tg *TubeGeometry) SetFaceVertexUvs(v js.Value) {
	tg.Set("faceVertexUvs", v)
}
func (tg *TubeGeometry) Faces() js.Value {
	return tg.Get("faces")
}
func (tg *TubeGeometry) SetFaces(v js.Value) {
	tg.Set("faces", v)
}
func (tg *TubeGeometry) GroupsNeedUpdate() bool {
	return tg.Get("groupsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetGroupsNeedUpdate(v bool) {
	tg.Set("groupsNeedUpdate", v)
}
func (tg *TubeGeometry) Id() int {
	return tg.Get("id").Int()
}
func (tg *TubeGeometry) SetId(v int) {
	tg.Set("id", v)
}
func (tg *TubeGeometry) LineDistances() js.Value {
	return tg.Get("lineDistances")
}
func (tg *TubeGeometry) SetLineDistances(v js.Value) {
	tg.Set("lineDistances", v)
}
func (tg *TubeGeometry) LineDistancesNeedUpdate() bool {
	return tg.Get("lineDistancesNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetLineDistancesNeedUpdate(v bool) {
	tg.Set("lineDistancesNeedUpdate", v)
}
func (tg *TubeGeometry) MorphNormals() js.Value {
	return tg.Get("morphNormals")
}
func (tg *TubeGeometry) SetMorphNormals(v js.Value) {
	tg.Set("morphNormals", v)
}
func (tg *TubeGeometry) MorphTargets() js.Value {
	return tg.Get("morphTargets")
}
func (tg *TubeGeometry) SetMorphTargets(v js.Value) {
	tg.Set("morphTargets", v)
}
func (tg *TubeGeometry) Name() string {
	return tg.Get("name").String()
}
func (tg *TubeGeometry) SetName(v string) {
	tg.Set("name", v)
}
func (tg *TubeGeometry) Normals() js.Value {
	return tg.Get("normals")
}
func (tg *TubeGeometry) SetNormals(v js.Value) {
	tg.Set("normals", v)
}
func (tg *TubeGeometry) NormalsNeedUpdate() bool {
	return tg.Get("normalsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetNormalsNeedUpdate(v bool) {
	tg.Set("normalsNeedUpdate", v)
}
func (tg *TubeGeometry) Parameters() js.Value {
	return tg.Get("parameters")
}
func (tg *TubeGeometry) SetParameters(v js.Value) {
	tg.Set("parameters", v)
}
func (tg *TubeGeometry) SkinIndices() js.Value {
	return tg.Get("skinIndices")
}
func (tg *TubeGeometry) SetSkinIndices(v js.Value) {
	tg.Set("skinIndices", v)
}
func (tg *TubeGeometry) SkinWeights() js.Value {
	return tg.Get("skinWeights")
}
func (tg *TubeGeometry) SetSkinWeights(v js.Value) {
	tg.Set("skinWeights", v)
}
func (tg *TubeGeometry) Tangents() js.Value {
	return tg.Get("tangents")
}
func (tg *TubeGeometry) SetTangents(v js.Value) {
	tg.Set("tangents", v)
}
func (tg *TubeGeometry) Type() string {
	return tg.Get("type").String()
}
func (tg *TubeGeometry) SetType(v string) {
	tg.Set("type", v)
}
func (tg *TubeGeometry) Uuid() string {
	return tg.Get("uuid").String()
}
func (tg *TubeGeometry) SetUuid(v string) {
	tg.Set("uuid", v)
}
func (tg *TubeGeometry) UvsNeedUpdate() bool {
	return tg.Get("uvsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetUvsNeedUpdate(v bool) {
	tg.Set("uvsNeedUpdate", v)
}
func (tg *TubeGeometry) Vertices() js.Value {
	return tg.Get("vertices")
}
func (tg *TubeGeometry) SetVertices(v js.Value) {
	tg.Set("vertices", v)
}
func (tg *TubeGeometry) VerticesNeedUpdate() bool {
	return tg.Get("verticesNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetVerticesNeedUpdate(v bool) {
	tg.Set("verticesNeedUpdate", v)
}
func (tg *TubeGeometry) AddEventListener(typ string, listener js.Value) {
	tg.Call("addEventListener", typ, listener)
}
func (tg *TubeGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: tg.Call("applyMatrix", matrix)}
}
func (tg *TubeGeometry) Center() Geometry {
	return &GeometryImpl{Value: tg.Call("center")}
}
func (tg *TubeGeometry) Clone() Geometry {
	return &TubeGeometry{Value: tg.Call("clone")}
}
func (tg *TubeGeometry) ComputeBoundingBox() {
	tg.Call("computeBoundingBox")
}
func (tg *TubeGeometry) ComputeBoundingSphere() {
	tg.Call("computeBoundingSphere")
}
func (tg *TubeGeometry) ComputeFaceNormals() {
	tg.Call("computeFaceNormals")
}
func (tg *TubeGeometry) ComputeFlatVertexNormals() {
	tg.Call("computeFlatVertexNormals")
}
func (tg *TubeGeometry) ComputeMorphNormals() {
	tg.Call("computeMorphNormals")
}
func (tg *TubeGeometry) ComputeVertexNormals(areaWeighted bool) {
	tg.Call("computeVertexNormals", areaWeighted)
}
func (tg *TubeGeometry) Copy(source Geometry) Geometry {
	return &TubeGeometry{Value: tg.Call("copy", source.JSValue())}
}
func (tg *TubeGeometry) DispatchEvent(event js.Value) {
	tg.Call("dispatchEvent", event)
}
func (tg *TubeGeometry) Dispose() {
	tg.Call("dispose")
}
func (tg *TubeGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: tg.Call("fromBufferGeometry", geometry)}
}
func (tg *TubeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tg.Call("hasEventListener", typ, listener).Bool()
}
func (tg *TubeGeometry) LookAt(vector *Vector3) {
	tg.Call("lookAt", vector)
}
func (tg *TubeGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	tg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (tg *TubeGeometry) MergeMesh(mesh *Mesh) {
	tg.Call("mergeMesh", mesh)
}
func (tg *TubeGeometry) MergeVertices() float64 {
	return tg.Call("mergeVertices").Float()
}
func (tg *TubeGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: tg.Call("normalize")}
}
func (tg *TubeGeometry) RemoveEventListener(typ string, listener js.Value) {
	tg.Call("removeEventListener", typ, listener)
}
func (tg *TubeGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: tg.Call("rotateX", angle)}
}
func (tg *TubeGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: tg.Call("rotateY", angle)}
}
func (tg *TubeGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: tg.Call("rotateZ", angle)}
}
func (tg *TubeGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: tg.Call("scale", x, y, z)}
}
func (tg *TubeGeometry) SetFromPoints(points js.Value) Geometry {
	return &TubeGeometry{Value: tg.Call("setFromPoints", points)}
}
func (tg *TubeGeometry) SortFacesByMaterialIndex() {
	tg.Call("sortFacesByMaterialIndex")
}
func (tg *TubeGeometry) ToJSON() js.Value {
	return tg.Call("toJSON")
}
func (tg *TubeGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: tg.Call("translate", x, y, z)}
}

// Code generated from "geometries/TetrahedronGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// TetrahedronBufferGeometry extend: [PolyhedronBufferGeometry]
type TetrahedronBufferGeometry struct {
	js.Value
}

func NewTetrahedronBufferGeometry(radius float64, detail float64) *TetrahedronBufferGeometry {
	return &TetrahedronBufferGeometry{Value: get("TetrahedronBufferGeometry").New(radius, detail)}
}
func (tbg *TetrahedronBufferGeometry) JSValue() js.Value {
	return tbg.Value
}
func (tbg *TetrahedronBufferGeometry) Attributes() js.Value {
	return tbg.Get("attributes")
}
func (tbg *TetrahedronBufferGeometry) SetAttributes(v js.Value) {
	tbg.Set("attributes", v)
}
func (tbg *TetrahedronBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tbg.Get("boundingBox")}
}
func (tbg *TetrahedronBufferGeometry) SetBoundingBox(v *Box3) {
	tbg.Set("boundingBox", v.JSValue())
}
func (tbg *TetrahedronBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tbg.Get("boundingSphere")}
}
func (tbg *TetrahedronBufferGeometry) SetBoundingSphere(v *Sphere) {
	tbg.Set("boundingSphere", v.JSValue())
}
func (tbg *TetrahedronBufferGeometry) DrawRange() js.Value {
	return tbg.Get("drawRange")
}
func (tbg *TetrahedronBufferGeometry) SetDrawRange(v js.Value) {
	tbg.Set("drawRange", v)
}
func (tbg *TetrahedronBufferGeometry) Drawcalls() js.Value {
	return tbg.Get("drawcalls")
}
func (tbg *TetrahedronBufferGeometry) SetDrawcalls(v js.Value) {
	tbg.Set("drawcalls", v)
}
func (tbg *TetrahedronBufferGeometry) Groups() js.Value {
	return tbg.Get("groups")
}
func (tbg *TetrahedronBufferGeometry) SetGroups(v js.Value) {
	tbg.Set("groups", v)
}
func (tbg *TetrahedronBufferGeometry) Id() int {
	return tbg.Get("id").Int()
}
func (tbg *TetrahedronBufferGeometry) SetId(v int) {
	tbg.Set("id", v)
}
func (tbg *TetrahedronBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Get("index")}
}
func (tbg *TetrahedronBufferGeometry) SetIndex(v *BufferAttribute) {
	tbg.Set("index", v.JSValue())
}
func (tbg *TetrahedronBufferGeometry) MorphAttributes() js.Value {
	return tbg.Get("morphAttributes")
}
func (tbg *TetrahedronBufferGeometry) SetMorphAttributes(v js.Value) {
	tbg.Set("morphAttributes", v)
}
func (tbg *TetrahedronBufferGeometry) Name() string {
	return tbg.Get("name").String()
}
func (tbg *TetrahedronBufferGeometry) SetName(v string) {
	tbg.Set("name", v)
}
func (tbg *TetrahedronBufferGeometry) Offsets() js.Value {
	return tbg.Get("offsets")
}
func (tbg *TetrahedronBufferGeometry) SetOffsets(v js.Value) {
	tbg.Set("offsets", v)
}
func (tbg *TetrahedronBufferGeometry) Parameters() js.Value {
	return tbg.Get("parameters")
}
func (tbg *TetrahedronBufferGeometry) SetParameters(v js.Value) {
	tbg.Set("parameters", v)
}
func (tbg *TetrahedronBufferGeometry) Type() string {
	return tbg.Get("type").String()
}
func (tbg *TetrahedronBufferGeometry) SetType(v string) {
	tbg.Set("type", v)
}
func (tbg *TetrahedronBufferGeometry) UserData() js.Value {
	return tbg.Get("userData")
}
func (tbg *TetrahedronBufferGeometry) SetUserData(v js.Value) {
	tbg.Set("userData", v)
}
func (tbg *TetrahedronBufferGeometry) Uuid() string {
	return tbg.Get("uuid").String()
}
func (tbg *TetrahedronBufferGeometry) SetUuid(v string) {
	tbg.Set("uuid", v)
}
func (tbg *TetrahedronBufferGeometry) MaxIndex() int {
	return tbg.Get("MaxIndex").Int()
}
func (tbg *TetrahedronBufferGeometry) SetMaxIndex(v int) {
	tbg.Set("MaxIndex", v)
}
func (tbg *TetrahedronBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("addAttribute", name, attribute)}
}
func (tbg *TetrahedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tbg.Call("addAttribute", name, array, itemSize)
}
func (tbg *TetrahedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tbg.Call("addDrawCall", start, count, indexOffset)
}
func (tbg *TetrahedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tbg.Call("addEventListener", typ, listener)
}
func (tbg *TetrahedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tbg.Call("addGroup", start, count, materialIndex)
}
func (tbg *TetrahedronBufferGeometry) AddIndex(index js.Value) {
	tbg.Call("addIndex", index)
}
func (tbg *TetrahedronBufferGeometry) ApplyMatrix(matrix *Matrix4) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("applyMatrix", matrix)}
}
func (tbg *TetrahedronBufferGeometry) Center() BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("center")}
}
func (tbg *TetrahedronBufferGeometry) ClearDrawCalls() {
	tbg.Call("clearDrawCalls")
}
func (tbg *TetrahedronBufferGeometry) ClearGroups() {
	tbg.Call("clearGroups")
}
func (tbg *TetrahedronBufferGeometry) Clone() BufferGeometry {
	return &TetrahedronBufferGeometry{Value: tbg.Call("clone")}
}
func (tbg *TetrahedronBufferGeometry) ComputeBoundingBox() {
	tbg.Call("computeBoundingBox")
}
func (tbg *TetrahedronBufferGeometry) ComputeBoundingSphere() {
	tbg.Call("computeBoundingSphere")
}
func (tbg *TetrahedronBufferGeometry) ComputeVertexNormals() {
	tbg.Call("computeVertexNormals")
}
func (tbg *TetrahedronBufferGeometry) Copy(source BufferGeometry) BufferGeometry {
	return &TetrahedronBufferGeometry{Value: tbg.Call("copy", source.JSValue())}
}
func (tbg *TetrahedronBufferGeometry) DispatchEvent(event js.Value) {
	tbg.Call("dispatchEvent", event)
}
func (tbg *TetrahedronBufferGeometry) Dispose() {
	tbg.Call("dispose")
}
func (tbg *TetrahedronBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("fromDirectGeometry", geometry)}
}
func (tbg *TetrahedronBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (tbg *TetrahedronBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getAttribute", name)}
}
func (tbg *TetrahedronBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getIndex")}
}
func (tbg *TetrahedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tbg.Call("hasEventListener", typ, listener).Bool()
}
func (tbg *TetrahedronBufferGeometry) LookAt(v *Vector3) {
	tbg.Call("lookAt", v.JSValue())
}
func (tbg *TetrahedronBufferGeometry) Merge(geometry BufferGeometry, offset int) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("merge", geometry.JSValue(), offset)}
}
func (tbg *TetrahedronBufferGeometry) NormalizeNormals() {
	tbg.Call("normalizeNormals")
}
func (tbg *TetrahedronBufferGeometry) RemoveAttribute(name string) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("removeAttribute", name)}
}
func (tbg *TetrahedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tbg.Call("removeEventListener", typ, listener)
}
func (tbg *TetrahedronBufferGeometry) RotateX(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("rotateX", angle)}
}
func (tbg *TetrahedronBufferGeometry) RotateY(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("rotateY", angle)}
}
func (tbg *TetrahedronBufferGeometry) RotateZ(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("rotateZ", angle)}
}
func (tbg *TetrahedronBufferGeometry) Scale(x float64, y float64, z float64) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("scale", x, y, z)}
}
func (tbg *TetrahedronBufferGeometry) SetDrawRange2(start int, count int) {
	tbg.Call("setDrawRange", start, count)
}
func (tbg *TetrahedronBufferGeometry) SetFromObject(object *Object3D) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("setFromObject", object)}
}
func (tbg *TetrahedronBufferGeometry) SetFromPoints(points js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("setFromPoints", points)}
}
func (tbg *TetrahedronBufferGeometry) SetIndex2(index *BufferAttribute) {
	tbg.Call("setIndex", index)
}
func (tbg *TetrahedronBufferGeometry) ToJSON() js.Value {
	return tbg.Call("toJSON")
}
func (tbg *TetrahedronBufferGeometry) ToNonIndexed() BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("toNonIndexed")}
}
func (tbg *TetrahedronBufferGeometry) Translate(x float64, y float64, z float64) BufferGeometry {
	return &BufferGeometryImpl{Value: tbg.Call("translate", x, y, z)}
}
func (tbg *TetrahedronBufferGeometry) UpdateFromObject(object *Object3D) {
	tbg.Call("updateFromObject", object.JSValue())
}

// TetrahedronGeometry extend: [PolyhedronGeometry]
type TetrahedronGeometry struct {
	js.Value
}

func NewTetrahedronGeometry(radius float64, detail float64) *TetrahedronGeometry {
	return &TetrahedronGeometry{Value: get("TetrahedronGeometry").New(radius, detail)}
}
func (tg *TetrahedronGeometry) JSValue() js.Value {
	return tg.Value
}
func (tg *TetrahedronGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: tg.Get("animation")}
}
func (tg *TetrahedronGeometry) SetAnimation(v *AnimationClip) {
	tg.Set("animation", v.JSValue())
}
func (tg *TetrahedronGeometry) Animations() js.Value {
	return tg.Get("animations")
}
func (tg *TetrahedronGeometry) SetAnimations(v js.Value) {
	tg.Set("animations", v)
}
func (tg *TetrahedronGeometry) Bones() js.Value {
	return tg.Get("bones")
}
func (tg *TetrahedronGeometry) SetBones(v js.Value) {
	tg.Set("bones", v)
}
func (tg *TetrahedronGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tg.Get("boundingBox")}
}
func (tg *TetrahedronGeometry) SetBoundingBox(v *Box3) {
	tg.Set("boundingBox", v.JSValue())
}
func (tg *TetrahedronGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tg.Get("boundingSphere")}
}
func (tg *TetrahedronGeometry) SetBoundingSphere(v *Sphere) {
	tg.Set("boundingSphere", v.JSValue())
}
func (tg *TetrahedronGeometry) Colors() js.Value {
	return tg.Get("colors")
}
func (tg *TetrahedronGeometry) SetColors(v js.Value) {
	tg.Set("colors", v)
}
func (tg *TetrahedronGeometry) ColorsNeedUpdate() bool {
	return tg.Get("colorsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetColorsNeedUpdate(v bool) {
	tg.Set("colorsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) ElementsNeedUpdate() bool {
	return tg.Get("elementsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetElementsNeedUpdate(v bool) {
	tg.Set("elementsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) FaceVertexUvs() js.Value {
	return tg.Get("faceVertexUvs")
}
func (tg *TetrahedronGeometry) SetFaceVertexUvs(v js.Value) {
	tg.Set("faceVertexUvs", v)
}
func (tg *TetrahedronGeometry) Faces() js.Value {
	return tg.Get("faces")
}
func (tg *TetrahedronGeometry) SetFaces(v js.Value) {
	tg.Set("faces", v)
}
func (tg *TetrahedronGeometry) GroupsNeedUpdate() bool {
	return tg.Get("groupsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetGroupsNeedUpdate(v bool) {
	tg.Set("groupsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) Id() int {
	return tg.Get("id").Int()
}
func (tg *TetrahedronGeometry) SetId(v int) {
	tg.Set("id", v)
}
func (tg *TetrahedronGeometry) LineDistances() js.Value {
	return tg.Get("lineDistances")
}
func (tg *TetrahedronGeometry) SetLineDistances(v js.Value) {
	tg.Set("lineDistances", v)
}
func (tg *TetrahedronGeometry) LineDistancesNeedUpdate() bool {
	return tg.Get("lineDistancesNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	tg.Set("lineDistancesNeedUpdate", v)
}
func (tg *TetrahedronGeometry) MorphNormals() js.Value {
	return tg.Get("morphNormals")
}
func (tg *TetrahedronGeometry) SetMorphNormals(v js.Value) {
	tg.Set("morphNormals", v)
}
func (tg *TetrahedronGeometry) MorphTargets() js.Value {
	return tg.Get("morphTargets")
}
func (tg *TetrahedronGeometry) SetMorphTargets(v js.Value) {
	tg.Set("morphTargets", v)
}
func (tg *TetrahedronGeometry) Name() string {
	return tg.Get("name").String()
}
func (tg *TetrahedronGeometry) SetName(v string) {
	tg.Set("name", v)
}
func (tg *TetrahedronGeometry) NormalsNeedUpdate() bool {
	return tg.Get("normalsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetNormalsNeedUpdate(v bool) {
	tg.Set("normalsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) Parameters() js.Value {
	return tg.Get("parameters")
}
func (tg *TetrahedronGeometry) SetParameters(v js.Value) {
	tg.Set("parameters", v)
}
func (tg *TetrahedronGeometry) SkinIndices() js.Value {
	return tg.Get("skinIndices")
}
func (tg *TetrahedronGeometry) SetSkinIndices(v js.Value) {
	tg.Set("skinIndices", v)
}
func (tg *TetrahedronGeometry) SkinWeights() js.Value {
	return tg.Get("skinWeights")
}
func (tg *TetrahedronGeometry) SetSkinWeights(v js.Value) {
	tg.Set("skinWeights", v)
}
func (tg *TetrahedronGeometry) Type() string {
	return tg.Get("type").String()
}
func (tg *TetrahedronGeometry) SetType(v string) {
	tg.Set("type", v)
}
func (tg *TetrahedronGeometry) Uuid() string {
	return tg.Get("uuid").String()
}
func (tg *TetrahedronGeometry) SetUuid(v string) {
	tg.Set("uuid", v)
}
func (tg *TetrahedronGeometry) UvsNeedUpdate() bool {
	return tg.Get("uvsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetUvsNeedUpdate(v bool) {
	tg.Set("uvsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) Vertices() js.Value {
	return tg.Get("vertices")
}
func (tg *TetrahedronGeometry) SetVertices(v js.Value) {
	tg.Set("vertices", v)
}
func (tg *TetrahedronGeometry) VerticesNeedUpdate() bool {
	return tg.Get("verticesNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetVerticesNeedUpdate(v bool) {
	tg.Set("verticesNeedUpdate", v)
}
func (tg *TetrahedronGeometry) AddEventListener(typ string, listener js.Value) {
	tg.Call("addEventListener", typ, listener)
}
func (tg *TetrahedronGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: tg.Call("applyMatrix", matrix)}
}
func (tg *TetrahedronGeometry) Center() Geometry {
	return &GeometryImpl{Value: tg.Call("center")}
}
func (tg *TetrahedronGeometry) Clone() Geometry {
	return &TetrahedronGeometry{Value: tg.Call("clone")}
}
func (tg *TetrahedronGeometry) ComputeBoundingBox() {
	tg.Call("computeBoundingBox")
}
func (tg *TetrahedronGeometry) ComputeBoundingSphere() {
	tg.Call("computeBoundingSphere")
}
func (tg *TetrahedronGeometry) ComputeFaceNormals() {
	tg.Call("computeFaceNormals")
}
func (tg *TetrahedronGeometry) ComputeFlatVertexNormals() {
	tg.Call("computeFlatVertexNormals")
}
func (tg *TetrahedronGeometry) ComputeMorphNormals() {
	tg.Call("computeMorphNormals")
}
func (tg *TetrahedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	tg.Call("computeVertexNormals", areaWeighted)
}
func (tg *TetrahedronGeometry) Copy(source Geometry) Geometry {
	return &TetrahedronGeometry{Value: tg.Call("copy", source.JSValue())}
}
func (tg *TetrahedronGeometry) DispatchEvent(event js.Value) {
	tg.Call("dispatchEvent", event)
}
func (tg *TetrahedronGeometry) Dispose() {
	tg.Call("dispose")
}
func (tg *TetrahedronGeometry) FromBufferGeometry(geometry BufferGeometry) Geometry {
	return &GeometryImpl{Value: tg.Call("fromBufferGeometry", geometry.JSValue())}
}
func (tg *TetrahedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tg.Call("hasEventListener", typ, listener).Bool()
}
func (tg *TetrahedronGeometry) LookAt(vector *Vector3) {
	tg.Call("lookAt", vector.JSValue())
}
func (tg *TetrahedronGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	tg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (tg *TetrahedronGeometry) MergeMesh(mesh Mesh) {
	tg.Call("mergeMesh", mesh.JSValue())
}
func (tg *TetrahedronGeometry) MergeVertices() float64 {
	return tg.Call("mergeVertices").Float()
}
func (tg *TetrahedronGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: tg.Call("normalize")}
}
func (tg *TetrahedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	tg.Call("removeEventListener", typ, listener)
}
func (tg *TetrahedronGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: tg.Call("rotateX", angle)}
}
func (tg *TetrahedronGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: tg.Call("rotateY", angle)}
}
func (tg *TetrahedronGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: tg.Call("rotateZ", angle)}
}
func (tg *TetrahedronGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: tg.Call("scale", x, y, z)}
}
func (tg *TetrahedronGeometry) SetFromPoints(points js.Value) Geometry {
	return &TetrahedronGeometry{Value: tg.Call("setFromPoints", points)}
}
func (tg *TetrahedronGeometry) SortFacesByMaterialIndex() {
	tg.Call("sortFacesByMaterialIndex")
}
func (tg *TetrahedronGeometry) ToJSON() js.Value {
	return tg.Call("toJSON")
}
func (tg *TetrahedronGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: tg.Call("translate", x, y, z)}
}

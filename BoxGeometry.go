// Code generated from "geometries/BoxGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// BoxBufferGeometry extend: [BufferGeometry]
type BoxBufferGeometry struct {
	js.Value
}

func NewBoxBufferGeometry(width float64, height float64, depth float64, widthSegments int, heightSegments int, depthSegments int) *BoxBufferGeometry {
	return &BoxBufferGeometry{Value: get("BoxBufferGeometry").New(width, height, depth, widthSegments, heightSegments, depthSegments)}
}
func (bbg *BoxBufferGeometry) JSValue() js.Value {
	return bbg.Value
}
func (bbg *BoxBufferGeometry) Attributes() js.Value {
	return bbg.Get("attributes")
}
func (bbg *BoxBufferGeometry) SetAttributes(v js.Value) {
	bbg.Set("attributes", v)
}
func (bbg *BoxBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: bbg.Get("boundingBox")}
}
func (bbg *BoxBufferGeometry) SetBoundingBox(v *Box3) {
	bbg.Set("boundingBox", v.Value)
}
func (bbg *BoxBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: bbg.Get("boundingSphere")}
}
func (bbg *BoxBufferGeometry) SetBoundingSphere(v *Sphere) {
	bbg.Set("boundingSphere", v.Value)
}
func (bbg *BoxBufferGeometry) DrawRange() js.Value {
	return bbg.Get("drawRange")
}
func (bbg *BoxBufferGeometry) SetDrawRange(v js.Value) {
	bbg.Set("drawRange", v)
}
func (bbg *BoxBufferGeometry) Drawcalls() js.Value {
	return bbg.Get("drawcalls")
}
func (bbg *BoxBufferGeometry) SetDrawcalls(v js.Value) {
	bbg.Set("drawcalls", v)
}
func (bbg *BoxBufferGeometry) Groups() js.Value {
	return bbg.Get("groups")
}
func (bbg *BoxBufferGeometry) SetGroups(v js.Value) {
	bbg.Set("groups", v)
}
func (bbg *BoxBufferGeometry) Id() int {
	return bbg.Get("id").Int()
}
func (bbg *BoxBufferGeometry) SetId(v int) {
	bbg.Set("id", v)
}
func (bbg *BoxBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: bbg.Get("index")}
}
func (bbg *BoxBufferGeometry) SetIndex(v *BufferAttribute) {
	bbg.Set("index", v.Value)
}
func (bbg *BoxBufferGeometry) MorphAttributes() js.Value {
	return bbg.Get("morphAttributes")
}
func (bbg *BoxBufferGeometry) SetMorphAttributes(v js.Value) {
	bbg.Set("morphAttributes", v)
}
func (bbg *BoxBufferGeometry) Name() string {
	return bbg.Get("name").String()
}
func (bbg *BoxBufferGeometry) SetName(v string) {
	bbg.Set("name", v)
}
func (bbg *BoxBufferGeometry) Offsets() js.Value {
	return bbg.Get("offsets")
}
func (bbg *BoxBufferGeometry) SetOffsets(v js.Value) {
	bbg.Set("offsets", v)
}
func (bbg *BoxBufferGeometry) Parameters() js.Value {
	return bbg.Get("parameters")
}
func (bbg *BoxBufferGeometry) SetParameters(v js.Value) {
	bbg.Set("parameters", v)
}
func (bbg *BoxBufferGeometry) Type() string {
	return bbg.Get("type").String()
}
func (bbg *BoxBufferGeometry) SetType(v string) {
	bbg.Set("type", v)
}
func (bbg *BoxBufferGeometry) UserData() js.Value {
	return bbg.Get("userData")
}
func (bbg *BoxBufferGeometry) SetUserData(v js.Value) {
	bbg.Set("userData", v)
}
func (bbg *BoxBufferGeometry) Uuid() string {
	return bbg.Get("uuid").String()
}
func (bbg *BoxBufferGeometry) SetUuid(v string) {
	bbg.Set("uuid", v)
}
func (bbg *BoxBufferGeometry) MaxIndex() int {
	return bbg.Get("MaxIndex").Int()
}
func (bbg *BoxBufferGeometry) SetMaxIndex(v int) {
	bbg.Set("MaxIndex", v)
}
func (bbg *BoxBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("addAttribute", name, attribute)}
}
func (bbg *BoxBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return bbg.Call("addAttribute", name, array, itemSize)
}
func (bbg *BoxBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	bbg.Call("addDrawCall", start, count, indexOffset)
}
func (bbg *BoxBufferGeometry) AddEventListener(typ string, listener js.Value) {
	bbg.Call("addEventListener", typ, listener)
}
func (bbg *BoxBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	bbg.Call("addGroup", start, count, materialIndex)
}
func (bbg *BoxBufferGeometry) AddIndex(index js.Value) {
	bbg.Call("addIndex", index)
}
func (bbg *BoxBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("applyMatrix", matrix)}
}
func (bbg *BoxBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("center")}
}
func (bbg *BoxBufferGeometry) ClearDrawCalls() {
	bbg.Call("clearDrawCalls")
}
func (bbg *BoxBufferGeometry) ClearGroups() {
	bbg.Call("clearGroups")
}
func (bbg *BoxBufferGeometry) Clone() *BoxBufferGeometry {
	return &BoxBufferGeometry{Value: bbg.Call("clone")}
}
func (bbg *BoxBufferGeometry) ComputeBoundingBox() {
	bbg.Call("computeBoundingBox")
}
func (bbg *BoxBufferGeometry) ComputeBoundingSphere() {
	bbg.Call("computeBoundingSphere")
}
func (bbg *BoxBufferGeometry) ComputeVertexNormals() {
	bbg.Call("computeVertexNormals")
}
func (bbg *BoxBufferGeometry) Copy(source *BufferGeometry) *BoxBufferGeometry {
	return &BoxBufferGeometry{Value: bbg.Call("copy", source)}
}
func (bbg *BoxBufferGeometry) DispatchEvent(event js.Value) {
	bbg.Call("dispatchEvent", event)
}
func (bbg *BoxBufferGeometry) Dispose() {
	bbg.Call("dispose")
}
func (bbg *BoxBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("fromDirectGeometry", geometry)}
}
func (bbg *BoxBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (bbg *BoxBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: bbg.Call("getAttribute", name)}
}
func (bbg *BoxBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: bbg.Call("getIndex")}
}
func (bbg *BoxBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return bbg.Call("hasEventListener", typ, listener).Bool()
}
func (bbg *BoxBufferGeometry) LookAt(v *Vector3) {
	bbg.Call("lookAt", v)
}
func (bbg *BoxBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("merge", geometry, offset)}
}
func (bbg *BoxBufferGeometry) NormalizeNormals() {
	bbg.Call("normalizeNormals")
}
func (bbg *BoxBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("removeAttribute", name)}
}
func (bbg *BoxBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	bbg.Call("removeEventListener", typ, listener)
}
func (bbg *BoxBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("rotateX", angle)}
}
func (bbg *BoxBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("rotateY", angle)}
}
func (bbg *BoxBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("rotateZ", angle)}
}
func (bbg *BoxBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("scale", x, y, z)}
}
func (bbg *BoxBufferGeometry) SetDrawRange2(start int, count int) {
	bbg.Call("setDrawRange", start, count)
}
func (bbg *BoxBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("setFromObject", object)}
}
func (bbg *BoxBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("setFromPoints", points)}
}
func (bbg *BoxBufferGeometry) SetIndex2(index *BufferAttribute) {
	bbg.Call("setIndex", index)
}
func (bbg *BoxBufferGeometry) ToJSON() js.Value {
	return bbg.Call("toJSON")
}
func (bbg *BoxBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("toNonIndexed")}
}
func (bbg *BoxBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: bbg.Call("translate", x, y, z)}
}
func (bbg *BoxBufferGeometry) UpdateFromObject(object *Object3D) {
	bbg.Call("updateFromObject", object)
}

// BoxGeometry extend: [Geometry]
type BoxGeometry struct {
	js.Value
}

func NewBoxGeometry(width float64, height float64, depth float64, widthSegments int, heightSegments int, depthSegments int) *BoxGeometry {
	return &BoxGeometry{Value: get("BoxGeometry").New(width, height, depth, widthSegments, heightSegments, depthSegments)}
}
func (bg *BoxGeometry) JSValue() js.Value {
	return bg.Value
}
func (bg *BoxGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: bg.Get("animation")}
}
func (bg *BoxGeometry) SetAnimation(v *AnimationClip) {
	bg.Set("animation", v.Value)
}
func (bg *BoxGeometry) Animations() js.Value {
	return bg.Get("animations")
}
func (bg *BoxGeometry) SetAnimations(v js.Value) {
	bg.Set("animations", v)
}
func (bg *BoxGeometry) Bones() js.Value {
	return bg.Get("bones")
}
func (bg *BoxGeometry) SetBones(v js.Value) {
	bg.Set("bones", v)
}
func (bg *BoxGeometry) BoundingBox() *Box3 {
	return &Box3{Value: bg.Get("boundingBox")}
}
func (bg *BoxGeometry) SetBoundingBox(v *Box3) {
	bg.Set("boundingBox", v.Value)
}
func (bg *BoxGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: bg.Get("boundingSphere")}
}
func (bg *BoxGeometry) SetBoundingSphere(v *Sphere) {
	bg.Set("boundingSphere", v.Value)
}
func (bg *BoxGeometry) Colors() js.Value {
	return bg.Get("colors")
}
func (bg *BoxGeometry) SetColors(v js.Value) {
	bg.Set("colors", v)
}
func (bg *BoxGeometry) ColorsNeedUpdate() bool {
	return bg.Get("colorsNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetColorsNeedUpdate(v bool) {
	bg.Set("colorsNeedUpdate", v)
}
func (bg *BoxGeometry) ElementsNeedUpdate() bool {
	return bg.Get("elementsNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetElementsNeedUpdate(v bool) {
	bg.Set("elementsNeedUpdate", v)
}
func (bg *BoxGeometry) FaceVertexUvs() js.Value {
	return bg.Get("faceVertexUvs")
}
func (bg *BoxGeometry) SetFaceVertexUvs(v js.Value) {
	bg.Set("faceVertexUvs", v)
}
func (bg *BoxGeometry) Faces() js.Value {
	return bg.Get("faces")
}
func (bg *BoxGeometry) SetFaces(v js.Value) {
	bg.Set("faces", v)
}
func (bg *BoxGeometry) GroupsNeedUpdate() bool {
	return bg.Get("groupsNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetGroupsNeedUpdate(v bool) {
	bg.Set("groupsNeedUpdate", v)
}
func (bg *BoxGeometry) Id() int {
	return bg.Get("id").Int()
}
func (bg *BoxGeometry) SetId(v int) {
	bg.Set("id", v)
}
func (bg *BoxGeometry) LineDistances() js.Value {
	return bg.Get("lineDistances")
}
func (bg *BoxGeometry) SetLineDistances(v js.Value) {
	bg.Set("lineDistances", v)
}
func (bg *BoxGeometry) LineDistancesNeedUpdate() bool {
	return bg.Get("lineDistancesNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetLineDistancesNeedUpdate(v bool) {
	bg.Set("lineDistancesNeedUpdate", v)
}
func (bg *BoxGeometry) MorphNormals() js.Value {
	return bg.Get("morphNormals")
}
func (bg *BoxGeometry) SetMorphNormals(v js.Value) {
	bg.Set("morphNormals", v)
}
func (bg *BoxGeometry) MorphTargets() js.Value {
	return bg.Get("morphTargets")
}
func (bg *BoxGeometry) SetMorphTargets(v js.Value) {
	bg.Set("morphTargets", v)
}
func (bg *BoxGeometry) Name() string {
	return bg.Get("name").String()
}
func (bg *BoxGeometry) SetName(v string) {
	bg.Set("name", v)
}
func (bg *BoxGeometry) NormalsNeedUpdate() bool {
	return bg.Get("normalsNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetNormalsNeedUpdate(v bool) {
	bg.Set("normalsNeedUpdate", v)
}
func (bg *BoxGeometry) Parameters() js.Value {
	return bg.Get("parameters")
}
func (bg *BoxGeometry) SetParameters(v js.Value) {
	bg.Set("parameters", v)
}
func (bg *BoxGeometry) SkinIndices() js.Value {
	return bg.Get("skinIndices")
}
func (bg *BoxGeometry) SetSkinIndices(v js.Value) {
	bg.Set("skinIndices", v)
}
func (bg *BoxGeometry) SkinWeights() js.Value {
	return bg.Get("skinWeights")
}
func (bg *BoxGeometry) SetSkinWeights(v js.Value) {
	bg.Set("skinWeights", v)
}
func (bg *BoxGeometry) Type() string {
	return bg.Get("type").String()
}
func (bg *BoxGeometry) SetType(v string) {
	bg.Set("type", v)
}
func (bg *BoxGeometry) Uuid() string {
	return bg.Get("uuid").String()
}
func (bg *BoxGeometry) SetUuid(v string) {
	bg.Set("uuid", v)
}
func (bg *BoxGeometry) UvsNeedUpdate() bool {
	return bg.Get("uvsNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetUvsNeedUpdate(v bool) {
	bg.Set("uvsNeedUpdate", v)
}
func (bg *BoxGeometry) Vertices() js.Value {
	return bg.Get("vertices")
}
func (bg *BoxGeometry) SetVertices(v js.Value) {
	bg.Set("vertices", v)
}
func (bg *BoxGeometry) VerticesNeedUpdate() bool {
	return bg.Get("verticesNeedUpdate").Bool()
}
func (bg *BoxGeometry) SetVerticesNeedUpdate(v bool) {
	bg.Set("verticesNeedUpdate", v)
}
func (bg *BoxGeometry) AddEventListener(typ string, listener js.Value) {
	bg.Call("addEventListener", typ, listener)
}
func (bg *BoxGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: bg.Call("applyMatrix", matrix)}
}
func (bg *BoxGeometry) Center() Geometry {
	return &GeometryImpl{Value: bg.Call("center")}
}
func (bg *BoxGeometry) Clone() Geometry {
	return &BoxGeometry{Value: bg.Call("clone")}
}
func (bg *BoxGeometry) ComputeBoundingBox() {
	bg.Call("computeBoundingBox")
}
func (bg *BoxGeometry) ComputeBoundingSphere() {
	bg.Call("computeBoundingSphere")
}
func (bg *BoxGeometry) ComputeFaceNormals() {
	bg.Call("computeFaceNormals")
}
func (bg *BoxGeometry) ComputeFlatVertexNormals() {
	bg.Call("computeFlatVertexNormals")
}
func (bg *BoxGeometry) ComputeMorphNormals() {
	bg.Call("computeMorphNormals")
}
func (bg *BoxGeometry) ComputeVertexNormals(areaWeighted bool) {
	bg.Call("computeVertexNormals", areaWeighted)
}
func (bg *BoxGeometry) Copy(source Geometry) Geometry {
	return &BoxGeometry{Value: bg.Call("copy", source.JSValue())}
}
func (bg *BoxGeometry) DispatchEvent(event js.Value) {
	bg.Call("dispatchEvent", event)
}
func (bg *BoxGeometry) Dispose() {
	bg.Call("dispose")
}
func (bg *BoxGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: bg.Call("fromBufferGeometry", geometry)}
}
func (bg *BoxGeometry) HasEventListener(typ string, listener js.Value) bool {
	return bg.Call("hasEventListener", typ, listener).Bool()
}
func (bg *BoxGeometry) LookAt(vector *Vector3) {
	bg.Call("lookAt", vector)
}
func (bg *BoxGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	bg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (bg *BoxGeometry) MergeMesh(mesh *Mesh) {
	bg.Call("mergeMesh", mesh)
}
func (bg *BoxGeometry) MergeVertices() float64 {
	return bg.Call("mergeVertices").Float()
}
func (bg *BoxGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: bg.Call("normalize")}
}
func (bg *BoxGeometry) RemoveEventListener(typ string, listener js.Value) {
	bg.Call("removeEventListener", typ, listener)
}
func (bg *BoxGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: bg.Call("rotateX", angle)}
}
func (bg *BoxGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: bg.Call("rotateY", angle)}
}
func (bg *BoxGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: bg.Call("rotateZ", angle)}
}
func (bg *BoxGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: bg.Call("scale", x, y, z)}
}
func (bg *BoxGeometry) SetFromPoints(points js.Value) Geometry {
	return &BoxGeometry{Value: bg.Call("setFromPoints", points)}
}
func (bg *BoxGeometry) SortFacesByMaterialIndex() {
	bg.Call("sortFacesByMaterialIndex")
}
func (bg *BoxGeometry) ToJSON() js.Value {
	return bg.Call("toJSON")
}
func (bg *BoxGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: bg.Call("translate", x, y, z)}
}

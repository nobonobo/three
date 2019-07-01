// Code generated from "geometries/PlaneGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PlaneBufferGeometry extend: [BufferGeometry]
type PlaneBufferGeometry struct {
	js.Value
}

func NewPlaneBufferGeometry(width float64, height float64, widthSegments int, heightSegments int) *PlaneBufferGeometry {
	return &PlaneBufferGeometry{Value: get("PlaneBufferGeometry").New(width, height, widthSegments, heightSegments)}
}
func (pbg *PlaneBufferGeometry) JSValue() js.Value {
	return pbg.Value
}
func (pbg *PlaneBufferGeometry) Attributes() js.Value {
	return pbg.Get("attributes")
}
func (pbg *PlaneBufferGeometry) SetAttributes(v js.Value) {
	pbg.Set("attributes", v)
}
func (pbg *PlaneBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: pbg.Get("boundingBox")}
}
func (pbg *PlaneBufferGeometry) SetBoundingBox(v *Box3) {
	pbg.Set("boundingBox", v.JSValue())
}
func (pbg *PlaneBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: pbg.Get("boundingSphere")}
}
func (pbg *PlaneBufferGeometry) SetBoundingSphere(v *Sphere) {
	pbg.Set("boundingSphere", v.JSValue())
}
func (pbg *PlaneBufferGeometry) DrawRange() js.Value {
	return pbg.Get("drawRange")
}
func (pbg *PlaneBufferGeometry) SetDrawRange(v js.Value) {
	pbg.Set("drawRange", v)
}
func (pbg *PlaneBufferGeometry) Drawcalls() js.Value {
	return pbg.Get("drawcalls")
}
func (pbg *PlaneBufferGeometry) SetDrawcalls(v js.Value) {
	pbg.Set("drawcalls", v)
}
func (pbg *PlaneBufferGeometry) Groups() js.Value {
	return pbg.Get("groups")
}
func (pbg *PlaneBufferGeometry) SetGroups(v js.Value) {
	pbg.Set("groups", v)
}
func (pbg *PlaneBufferGeometry) Id() int {
	return pbg.Get("id").Int()
}
func (pbg *PlaneBufferGeometry) SetId(v int) {
	pbg.Set("id", v)
}
func (pbg *PlaneBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: pbg.Get("index")}
}
func (pbg *PlaneBufferGeometry) SetIndex(v *BufferAttribute) {
	pbg.Set("index", v.JSValue())
}
func (pbg *PlaneBufferGeometry) MorphAttributes() js.Value {
	return pbg.Get("morphAttributes")
}
func (pbg *PlaneBufferGeometry) SetMorphAttributes(v js.Value) {
	pbg.Set("morphAttributes", v)
}
func (pbg *PlaneBufferGeometry) Name() string {
	return pbg.Get("name").String()
}
func (pbg *PlaneBufferGeometry) SetName(v string) {
	pbg.Set("name", v)
}
func (pbg *PlaneBufferGeometry) Offsets() js.Value {
	return pbg.Get("offsets")
}
func (pbg *PlaneBufferGeometry) SetOffsets(v js.Value) {
	pbg.Set("offsets", v)
}
func (pbg *PlaneBufferGeometry) Parameters() js.Value {
	return pbg.Get("parameters")
}
func (pbg *PlaneBufferGeometry) SetParameters(v js.Value) {
	pbg.Set("parameters", v)
}
func (pbg *PlaneBufferGeometry) Type() string {
	return pbg.Get("type").String()
}
func (pbg *PlaneBufferGeometry) SetType(v string) {
	pbg.Set("type", v)
}
func (pbg *PlaneBufferGeometry) UserData() js.Value {
	return pbg.Get("userData")
}
func (pbg *PlaneBufferGeometry) SetUserData(v js.Value) {
	pbg.Set("userData", v)
}
func (pbg *PlaneBufferGeometry) Uuid() string {
	return pbg.Get("uuid").String()
}
func (pbg *PlaneBufferGeometry) SetUuid(v string) {
	pbg.Set("uuid", v)
}
func (pbg *PlaneBufferGeometry) MaxIndex() int {
	return pbg.Get("MaxIndex").Int()
}
func (pbg *PlaneBufferGeometry) SetMaxIndex(v int) {
	pbg.Set("MaxIndex", v)
}
func (pbg *PlaneBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("addAttribute", name, attribute)}
}
func (pbg *PlaneBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return pbg.Call("addAttribute", name, array, itemSize)
}
func (pbg *PlaneBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	pbg.Call("addDrawCall", start, count, indexOffset)
}
func (pbg *PlaneBufferGeometry) AddEventListener(typ string, listener js.Value) {
	pbg.Call("addEventListener", typ, listener)
}
func (pbg *PlaneBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	pbg.Call("addGroup", start, count, materialIndex)
}
func (pbg *PlaneBufferGeometry) AddIndex(index js.Value) {
	pbg.Call("addIndex", index)
}
func (pbg *PlaneBufferGeometry) ApplyMatrix(matrix *Matrix4) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("applyMatrix", matrix)}
}
func (pbg *PlaneBufferGeometry) Center() BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("center")}
}
func (pbg *PlaneBufferGeometry) ClearDrawCalls() {
	pbg.Call("clearDrawCalls")
}
func (pbg *PlaneBufferGeometry) ClearGroups() {
	pbg.Call("clearGroups")
}
func (pbg *PlaneBufferGeometry) Clone() BufferGeometry {
	return &PlaneBufferGeometry{Value: pbg.Call("clone")}
}
func (pbg *PlaneBufferGeometry) ComputeBoundingBox() {
	pbg.Call("computeBoundingBox")
}
func (pbg *PlaneBufferGeometry) ComputeBoundingSphere() {
	pbg.Call("computeBoundingSphere")
}
func (pbg *PlaneBufferGeometry) ComputeVertexNormals() {
	pbg.Call("computeVertexNormals")
}
func (pbg *PlaneBufferGeometry) Copy(source BufferGeometry) BufferGeometry {
	return &PlaneBufferGeometry{Value: pbg.Call("copy", source.JSValue())}
}
func (pbg *PlaneBufferGeometry) DispatchEvent(event js.Value) {
	pbg.Call("dispatchEvent", event)
}
func (pbg *PlaneBufferGeometry) Dispose() {
	pbg.Call("dispose")
}
func (pbg *PlaneBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("fromDirectGeometry", geometry)}
}
func (pbg *PlaneBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (pbg *PlaneBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: pbg.Call("getAttribute", name)}
}
func (pbg *PlaneBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: pbg.Call("getIndex")}
}
func (pbg *PlaneBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pbg.Call("hasEventListener", typ, listener).Bool()
}
func (pbg *PlaneBufferGeometry) LookAt(v *Vector3) {
	pbg.Call("lookAt", v.JSValue())
}
func (pbg *PlaneBufferGeometry) Merge(geometry BufferGeometry, offset int) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("merge", geometry.JSValue(), offset)}
}
func (pbg *PlaneBufferGeometry) NormalizeNormals() {
	pbg.Call("normalizeNormals")
}
func (pbg *PlaneBufferGeometry) RemoveAttribute(name string) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("removeAttribute", name)}
}
func (pbg *PlaneBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	pbg.Call("removeEventListener", typ, listener)
}
func (pbg *PlaneBufferGeometry) RotateX(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("rotateX", angle)}
}
func (pbg *PlaneBufferGeometry) RotateY(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("rotateY", angle)}
}
func (pbg *PlaneBufferGeometry) RotateZ(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("rotateZ", angle)}
}
func (pbg *PlaneBufferGeometry) Scale(x float64, y float64, z float64) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("scale", x, y, z)}
}
func (pbg *PlaneBufferGeometry) SetDrawRange2(start int, count int) {
	pbg.Call("setDrawRange", start, count)
}
func (pbg *PlaneBufferGeometry) SetFromObject(object *Object3D) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("setFromObject", object)}
}
func (pbg *PlaneBufferGeometry) SetFromPoints(points js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("setFromPoints", points)}
}
func (pbg *PlaneBufferGeometry) SetIndex2(index *BufferAttribute) {
	pbg.Call("setIndex", index)
}
func (pbg *PlaneBufferGeometry) ToJSON() js.Value {
	return pbg.Call("toJSON")
}
func (pbg *PlaneBufferGeometry) ToNonIndexed() BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("toNonIndexed")}
}
func (pbg *PlaneBufferGeometry) Translate(x float64, y float64, z float64) BufferGeometry {
	return &BufferGeometryImpl{Value: pbg.Call("translate", x, y, z)}
}
func (pbg *PlaneBufferGeometry) UpdateFromObject(object *Object3D) {
	pbg.Call("updateFromObject", object.JSValue())
}

// PlaneGeometry extend: [Geometry]
type PlaneGeometry struct {
	js.Value
}

func NewPlaneGeometry(width float64, height float64, widthSegments int, heightSegments int) *PlaneGeometry {
	return &PlaneGeometry{Value: get("PlaneGeometry").New(width, height, widthSegments, heightSegments)}
}
func (pg *PlaneGeometry) JSValue() js.Value {
	return pg.Value
}
func (pg *PlaneGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: pg.Get("animation")}
}
func (pg *PlaneGeometry) SetAnimation(v *AnimationClip) {
	pg.Set("animation", v.JSValue())
}
func (pg *PlaneGeometry) Animations() js.Value {
	return pg.Get("animations")
}
func (pg *PlaneGeometry) SetAnimations(v js.Value) {
	pg.Set("animations", v)
}
func (pg *PlaneGeometry) Bones() js.Value {
	return pg.Get("bones")
}
func (pg *PlaneGeometry) SetBones(v js.Value) {
	pg.Set("bones", v)
}
func (pg *PlaneGeometry) BoundingBox() *Box3 {
	return &Box3{Value: pg.Get("boundingBox")}
}
func (pg *PlaneGeometry) SetBoundingBox(v *Box3) {
	pg.Set("boundingBox", v.JSValue())
}
func (pg *PlaneGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: pg.Get("boundingSphere")}
}
func (pg *PlaneGeometry) SetBoundingSphere(v *Sphere) {
	pg.Set("boundingSphere", v.JSValue())
}
func (pg *PlaneGeometry) Colors() js.Value {
	return pg.Get("colors")
}
func (pg *PlaneGeometry) SetColors(v js.Value) {
	pg.Set("colors", v)
}
func (pg *PlaneGeometry) ColorsNeedUpdate() bool {
	return pg.Get("colorsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetColorsNeedUpdate(v bool) {
	pg.Set("colorsNeedUpdate", v)
}
func (pg *PlaneGeometry) ElementsNeedUpdate() bool {
	return pg.Get("elementsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetElementsNeedUpdate(v bool) {
	pg.Set("elementsNeedUpdate", v)
}
func (pg *PlaneGeometry) FaceVertexUvs() js.Value {
	return pg.Get("faceVertexUvs")
}
func (pg *PlaneGeometry) SetFaceVertexUvs(v js.Value) {
	pg.Set("faceVertexUvs", v)
}
func (pg *PlaneGeometry) Faces() js.Value {
	return pg.Get("faces")
}
func (pg *PlaneGeometry) SetFaces(v js.Value) {
	pg.Set("faces", v)
}
func (pg *PlaneGeometry) GroupsNeedUpdate() bool {
	return pg.Get("groupsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetGroupsNeedUpdate(v bool) {
	pg.Set("groupsNeedUpdate", v)
}
func (pg *PlaneGeometry) Id() int {
	return pg.Get("id").Int()
}
func (pg *PlaneGeometry) SetId(v int) {
	pg.Set("id", v)
}
func (pg *PlaneGeometry) LineDistances() js.Value {
	return pg.Get("lineDistances")
}
func (pg *PlaneGeometry) SetLineDistances(v js.Value) {
	pg.Set("lineDistances", v)
}
func (pg *PlaneGeometry) LineDistancesNeedUpdate() bool {
	return pg.Get("lineDistancesNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetLineDistancesNeedUpdate(v bool) {
	pg.Set("lineDistancesNeedUpdate", v)
}
func (pg *PlaneGeometry) MorphNormals() js.Value {
	return pg.Get("morphNormals")
}
func (pg *PlaneGeometry) SetMorphNormals(v js.Value) {
	pg.Set("morphNormals", v)
}
func (pg *PlaneGeometry) MorphTargets() js.Value {
	return pg.Get("morphTargets")
}
func (pg *PlaneGeometry) SetMorphTargets(v js.Value) {
	pg.Set("morphTargets", v)
}
func (pg *PlaneGeometry) Name() string {
	return pg.Get("name").String()
}
func (pg *PlaneGeometry) SetName(v string) {
	pg.Set("name", v)
}
func (pg *PlaneGeometry) NormalsNeedUpdate() bool {
	return pg.Get("normalsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetNormalsNeedUpdate(v bool) {
	pg.Set("normalsNeedUpdate", v)
}
func (pg *PlaneGeometry) Parameters() js.Value {
	return pg.Get("parameters")
}
func (pg *PlaneGeometry) SetParameters(v js.Value) {
	pg.Set("parameters", v)
}
func (pg *PlaneGeometry) SkinIndices() js.Value {
	return pg.Get("skinIndices")
}
func (pg *PlaneGeometry) SetSkinIndices(v js.Value) {
	pg.Set("skinIndices", v)
}
func (pg *PlaneGeometry) SkinWeights() js.Value {
	return pg.Get("skinWeights")
}
func (pg *PlaneGeometry) SetSkinWeights(v js.Value) {
	pg.Set("skinWeights", v)
}
func (pg *PlaneGeometry) Type() string {
	return pg.Get("type").String()
}
func (pg *PlaneGeometry) SetType(v string) {
	pg.Set("type", v)
}
func (pg *PlaneGeometry) Uuid() string {
	return pg.Get("uuid").String()
}
func (pg *PlaneGeometry) SetUuid(v string) {
	pg.Set("uuid", v)
}
func (pg *PlaneGeometry) UvsNeedUpdate() bool {
	return pg.Get("uvsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetUvsNeedUpdate(v bool) {
	pg.Set("uvsNeedUpdate", v)
}
func (pg *PlaneGeometry) Vertices() js.Value {
	return pg.Get("vertices")
}
func (pg *PlaneGeometry) SetVertices(v js.Value) {
	pg.Set("vertices", v)
}
func (pg *PlaneGeometry) VerticesNeedUpdate() bool {
	return pg.Get("verticesNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetVerticesNeedUpdate(v bool) {
	pg.Set("verticesNeedUpdate", v)
}
func (pg *PlaneGeometry) AddEventListener(typ string, listener js.Value) {
	pg.Call("addEventListener", typ, listener)
}
func (pg *PlaneGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: pg.Call("applyMatrix", matrix)}
}
func (pg *PlaneGeometry) Center() Geometry {
	return &GeometryImpl{Value: pg.Call("center")}
}
func (pg *PlaneGeometry) Clone() Geometry {
	return &PlaneGeometry{Value: pg.Call("clone")}
}
func (pg *PlaneGeometry) ComputeBoundingBox() {
	pg.Call("computeBoundingBox")
}
func (pg *PlaneGeometry) ComputeBoundingSphere() {
	pg.Call("computeBoundingSphere")
}
func (pg *PlaneGeometry) ComputeFaceNormals() {
	pg.Call("computeFaceNormals")
}
func (pg *PlaneGeometry) ComputeFlatVertexNormals() {
	pg.Call("computeFlatVertexNormals")
}
func (pg *PlaneGeometry) ComputeMorphNormals() {
	pg.Call("computeMorphNormals")
}
func (pg *PlaneGeometry) ComputeVertexNormals(areaWeighted bool) {
	pg.Call("computeVertexNormals", areaWeighted)
}
func (pg *PlaneGeometry) Copy(source Geometry) Geometry {
	return &PlaneGeometry{Value: pg.Call("copy", source.JSValue())}
}
func (pg *PlaneGeometry) DispatchEvent(event js.Value) {
	pg.Call("dispatchEvent", event)
}
func (pg *PlaneGeometry) Dispose() {
	pg.Call("dispose")
}
func (pg *PlaneGeometry) FromBufferGeometry(geometry BufferGeometry) Geometry {
	return &GeometryImpl{Value: pg.Call("fromBufferGeometry", geometry.JSValue())}
}
func (pg *PlaneGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pg.Call("hasEventListener", typ, listener).Bool()
}
func (pg *PlaneGeometry) LookAt(vector *Vector3) {
	pg.Call("lookAt", vector.JSValue())
}
func (pg *PlaneGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	pg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (pg *PlaneGeometry) MergeMesh(mesh Mesh) {
	pg.Call("mergeMesh", mesh.JSValue())
}
func (pg *PlaneGeometry) MergeVertices() float64 {
	return pg.Call("mergeVertices").Float()
}
func (pg *PlaneGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: pg.Call("normalize")}
}
func (pg *PlaneGeometry) RemoveEventListener(typ string, listener js.Value) {
	pg.Call("removeEventListener", typ, listener)
}
func (pg *PlaneGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateX", angle)}
}
func (pg *PlaneGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateY", angle)}
}
func (pg *PlaneGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateZ", angle)}
}
func (pg *PlaneGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: pg.Call("scale", x, y, z)}
}
func (pg *PlaneGeometry) SetFromPoints(points js.Value) Geometry {
	return &PlaneGeometry{Value: pg.Call("setFromPoints", points)}
}
func (pg *PlaneGeometry) SortFacesByMaterialIndex() {
	pg.Call("sortFacesByMaterialIndex")
}
func (pg *PlaneGeometry) ToJSON() js.Value {
	return pg.Call("toJSON")
}
func (pg *PlaneGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: pg.Call("translate", x, y, z)}
}

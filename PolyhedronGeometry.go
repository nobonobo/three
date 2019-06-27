// Code generated from "geometries/PolyhedronGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PolyhedronBufferGeometry extend: [BufferGeometry]
type PolyhedronBufferGeometry struct {
	js.Value
}

func NewPolyhedronBufferGeometry(vertices js.Value, indices js.Value, radius float64, detail float64) *PolyhedronBufferGeometry {
	return &PolyhedronBufferGeometry{Value: get("PolyhedronBufferGeometry").New(vertices, indices, radius, detail)}
}
func (pbg *PolyhedronBufferGeometry) JSValue() js.Value {
	return pbg.Value
}
func (pbg *PolyhedronBufferGeometry) Attributes() js.Value {
	return pbg.Get("attributes")
}
func (pbg *PolyhedronBufferGeometry) SetAttributes(v js.Value) {
	pbg.Set("attributes", v)
}
func (pbg *PolyhedronBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: pbg.Get("boundingBox")}
}
func (pbg *PolyhedronBufferGeometry) SetBoundingBox(v *Box3) {
	pbg.Set("boundingBox", v.Value)
}
func (pbg *PolyhedronBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: pbg.Get("boundingSphere")}
}
func (pbg *PolyhedronBufferGeometry) SetBoundingSphere(v *Sphere) {
	pbg.Set("boundingSphere", v.Value)
}
func (pbg *PolyhedronBufferGeometry) DrawRange() js.Value {
	return pbg.Get("drawRange")
}
func (pbg *PolyhedronBufferGeometry) SetDrawRange(v js.Value) {
	pbg.Set("drawRange", v)
}
func (pbg *PolyhedronBufferGeometry) Drawcalls() js.Value {
	return pbg.Get("drawcalls")
}
func (pbg *PolyhedronBufferGeometry) SetDrawcalls(v js.Value) {
	pbg.Set("drawcalls", v)
}
func (pbg *PolyhedronBufferGeometry) Groups() js.Value {
	return pbg.Get("groups")
}
func (pbg *PolyhedronBufferGeometry) SetGroups(v js.Value) {
	pbg.Set("groups", v)
}
func (pbg *PolyhedronBufferGeometry) Id() int {
	return pbg.Get("id").Int()
}
func (pbg *PolyhedronBufferGeometry) SetId(v int) {
	pbg.Set("id", v)
}
func (pbg *PolyhedronBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: pbg.Get("index")}
}
func (pbg *PolyhedronBufferGeometry) SetIndex(v *BufferAttribute) {
	pbg.Set("index", v.Value)
}
func (pbg *PolyhedronBufferGeometry) MorphAttributes() js.Value {
	return pbg.Get("morphAttributes")
}
func (pbg *PolyhedronBufferGeometry) SetMorphAttributes(v js.Value) {
	pbg.Set("morphAttributes", v)
}
func (pbg *PolyhedronBufferGeometry) Name() string {
	return pbg.Get("name").String()
}
func (pbg *PolyhedronBufferGeometry) SetName(v string) {
	pbg.Set("name", v)
}
func (pbg *PolyhedronBufferGeometry) Offsets() js.Value {
	return pbg.Get("offsets")
}
func (pbg *PolyhedronBufferGeometry) SetOffsets(v js.Value) {
	pbg.Set("offsets", v)
}
func (pbg *PolyhedronBufferGeometry) Parameters() js.Value {
	return pbg.Get("parameters")
}
func (pbg *PolyhedronBufferGeometry) SetParameters(v js.Value) {
	pbg.Set("parameters", v)
}
func (pbg *PolyhedronBufferGeometry) Type() string {
	return pbg.Get("type").String()
}
func (pbg *PolyhedronBufferGeometry) SetType(v string) {
	pbg.Set("type", v)
}
func (pbg *PolyhedronBufferGeometry) UserData() js.Value {
	return pbg.Get("userData")
}
func (pbg *PolyhedronBufferGeometry) SetUserData(v js.Value) {
	pbg.Set("userData", v)
}
func (pbg *PolyhedronBufferGeometry) Uuid() string {
	return pbg.Get("uuid").String()
}
func (pbg *PolyhedronBufferGeometry) SetUuid(v string) {
	pbg.Set("uuid", v)
}
func (pbg *PolyhedronBufferGeometry) MaxIndex() int {
	return pbg.Get("MaxIndex").Int()
}
func (pbg *PolyhedronBufferGeometry) SetMaxIndex(v int) {
	pbg.Set("MaxIndex", v)
}
func (pbg *PolyhedronBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("addAttribute", name, attribute)}
}
func (pbg *PolyhedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return pbg.Call("addAttribute", name, array, itemSize)
}
func (pbg *PolyhedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	pbg.Call("addDrawCall", start, count, indexOffset)
}
func (pbg *PolyhedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	pbg.Call("addEventListener", typ, listener)
}
func (pbg *PolyhedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	pbg.Call("addGroup", start, count, materialIndex)
}
func (pbg *PolyhedronBufferGeometry) AddIndex(index js.Value) {
	pbg.Call("addIndex", index)
}
func (pbg *PolyhedronBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("applyMatrix", matrix)}
}
func (pbg *PolyhedronBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("center")}
}
func (pbg *PolyhedronBufferGeometry) ClearDrawCalls() {
	pbg.Call("clearDrawCalls")
}
func (pbg *PolyhedronBufferGeometry) ClearGroups() {
	pbg.Call("clearGroups")
}
func (pbg *PolyhedronBufferGeometry) Clone() *PolyhedronBufferGeometry {
	return &PolyhedronBufferGeometry{Value: pbg.Call("clone")}
}
func (pbg *PolyhedronBufferGeometry) ComputeBoundingBox() {
	pbg.Call("computeBoundingBox")
}
func (pbg *PolyhedronBufferGeometry) ComputeBoundingSphere() {
	pbg.Call("computeBoundingSphere")
}
func (pbg *PolyhedronBufferGeometry) ComputeVertexNormals() {
	pbg.Call("computeVertexNormals")
}
func (pbg *PolyhedronBufferGeometry) Copy(source *BufferGeometry) *PolyhedronBufferGeometry {
	return &PolyhedronBufferGeometry{Value: pbg.Call("copy", source)}
}
func (pbg *PolyhedronBufferGeometry) DispatchEvent(event js.Value) {
	pbg.Call("dispatchEvent", event)
}
func (pbg *PolyhedronBufferGeometry) Dispose() {
	pbg.Call("dispose")
}
func (pbg *PolyhedronBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("fromDirectGeometry", geometry)}
}
func (pbg *PolyhedronBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (pbg *PolyhedronBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: pbg.Call("getAttribute", name)}
}
func (pbg *PolyhedronBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: pbg.Call("getIndex")}
}
func (pbg *PolyhedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pbg.Call("hasEventListener", typ, listener).Bool()
}
func (pbg *PolyhedronBufferGeometry) LookAt(v *Vector3) {
	pbg.Call("lookAt", v)
}
func (pbg *PolyhedronBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("merge", geometry, offset)}
}
func (pbg *PolyhedronBufferGeometry) NormalizeNormals() {
	pbg.Call("normalizeNormals")
}
func (pbg *PolyhedronBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("removeAttribute", name)}
}
func (pbg *PolyhedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	pbg.Call("removeEventListener", typ, listener)
}
func (pbg *PolyhedronBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("rotateX", angle)}
}
func (pbg *PolyhedronBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("rotateY", angle)}
}
func (pbg *PolyhedronBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("rotateZ", angle)}
}
func (pbg *PolyhedronBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("scale", x, y, z)}
}
func (pbg *PolyhedronBufferGeometry) SetDrawRange2(start int, count int) {
	pbg.Call("setDrawRange", start, count)
}
func (pbg *PolyhedronBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("setFromObject", object)}
}
func (pbg *PolyhedronBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("setFromPoints", points)}
}
func (pbg *PolyhedronBufferGeometry) SetIndex2(index *BufferAttribute) {
	pbg.Call("setIndex", index)
}
func (pbg *PolyhedronBufferGeometry) ToJSON() js.Value {
	return pbg.Call("toJSON")
}
func (pbg *PolyhedronBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("toNonIndexed")}
}
func (pbg *PolyhedronBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("translate", x, y, z)}
}
func (pbg *PolyhedronBufferGeometry) UpdateFromObject(object *Object3D) {
	pbg.Call("updateFromObject", object)
}

// PolyhedronGeometry extend: [Geometry]
type PolyhedronGeometry struct {
	js.Value
}

func NewPolyhedronGeometry(vertices js.Value, indices js.Value, radius float64, detail float64) *PolyhedronGeometry {
	return &PolyhedronGeometry{Value: get("PolyhedronGeometry").New(vertices, indices, radius, detail)}
}
func (pg *PolyhedronGeometry) JSValue() js.Value {
	return pg.Value
}
func (pg *PolyhedronGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: pg.Get("animation")}
}
func (pg *PolyhedronGeometry) SetAnimation(v *AnimationClip) {
	pg.Set("animation", v.Value)
}
func (pg *PolyhedronGeometry) Animations() js.Value {
	return pg.Get("animations")
}
func (pg *PolyhedronGeometry) SetAnimations(v js.Value) {
	pg.Set("animations", v)
}
func (pg *PolyhedronGeometry) Bones() js.Value {
	return pg.Get("bones")
}
func (pg *PolyhedronGeometry) SetBones(v js.Value) {
	pg.Set("bones", v)
}
func (pg *PolyhedronGeometry) BoundingBox() *Box3 {
	return &Box3{Value: pg.Get("boundingBox")}
}
func (pg *PolyhedronGeometry) SetBoundingBox(v *Box3) {
	pg.Set("boundingBox", v.Value)
}
func (pg *PolyhedronGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: pg.Get("boundingSphere")}
}
func (pg *PolyhedronGeometry) SetBoundingSphere(v *Sphere) {
	pg.Set("boundingSphere", v.Value)
}
func (pg *PolyhedronGeometry) Colors() js.Value {
	return pg.Get("colors")
}
func (pg *PolyhedronGeometry) SetColors(v js.Value) {
	pg.Set("colors", v)
}
func (pg *PolyhedronGeometry) ColorsNeedUpdate() bool {
	return pg.Get("colorsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetColorsNeedUpdate(v bool) {
	pg.Set("colorsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) ElementsNeedUpdate() bool {
	return pg.Get("elementsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetElementsNeedUpdate(v bool) {
	pg.Set("elementsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) FaceVertexUvs() js.Value {
	return pg.Get("faceVertexUvs")
}
func (pg *PolyhedronGeometry) SetFaceVertexUvs(v js.Value) {
	pg.Set("faceVertexUvs", v)
}
func (pg *PolyhedronGeometry) Faces() js.Value {
	return pg.Get("faces")
}
func (pg *PolyhedronGeometry) SetFaces(v js.Value) {
	pg.Set("faces", v)
}
func (pg *PolyhedronGeometry) GroupsNeedUpdate() bool {
	return pg.Get("groupsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetGroupsNeedUpdate(v bool) {
	pg.Set("groupsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) Id() int {
	return pg.Get("id").Int()
}
func (pg *PolyhedronGeometry) SetId(v int) {
	pg.Set("id", v)
}
func (pg *PolyhedronGeometry) LineDistances() js.Value {
	return pg.Get("lineDistances")
}
func (pg *PolyhedronGeometry) SetLineDistances(v js.Value) {
	pg.Set("lineDistances", v)
}
func (pg *PolyhedronGeometry) LineDistancesNeedUpdate() bool {
	return pg.Get("lineDistancesNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	pg.Set("lineDistancesNeedUpdate", v)
}
func (pg *PolyhedronGeometry) MorphNormals() js.Value {
	return pg.Get("morphNormals")
}
func (pg *PolyhedronGeometry) SetMorphNormals(v js.Value) {
	pg.Set("morphNormals", v)
}
func (pg *PolyhedronGeometry) MorphTargets() js.Value {
	return pg.Get("morphTargets")
}
func (pg *PolyhedronGeometry) SetMorphTargets(v js.Value) {
	pg.Set("morphTargets", v)
}
func (pg *PolyhedronGeometry) Name() string {
	return pg.Get("name").String()
}
func (pg *PolyhedronGeometry) SetName(v string) {
	pg.Set("name", v)
}
func (pg *PolyhedronGeometry) NormalsNeedUpdate() bool {
	return pg.Get("normalsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetNormalsNeedUpdate(v bool) {
	pg.Set("normalsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) Parameters() js.Value {
	return pg.Get("parameters")
}
func (pg *PolyhedronGeometry) SetParameters(v js.Value) {
	pg.Set("parameters", v)
}
func (pg *PolyhedronGeometry) SkinIndices() js.Value {
	return pg.Get("skinIndices")
}
func (pg *PolyhedronGeometry) SetSkinIndices(v js.Value) {
	pg.Set("skinIndices", v)
}
func (pg *PolyhedronGeometry) SkinWeights() js.Value {
	return pg.Get("skinWeights")
}
func (pg *PolyhedronGeometry) SetSkinWeights(v js.Value) {
	pg.Set("skinWeights", v)
}
func (pg *PolyhedronGeometry) Type() string {
	return pg.Get("type").String()
}
func (pg *PolyhedronGeometry) SetType(v string) {
	pg.Set("type", v)
}
func (pg *PolyhedronGeometry) Uuid() string {
	return pg.Get("uuid").String()
}
func (pg *PolyhedronGeometry) SetUuid(v string) {
	pg.Set("uuid", v)
}
func (pg *PolyhedronGeometry) UvsNeedUpdate() bool {
	return pg.Get("uvsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetUvsNeedUpdate(v bool) {
	pg.Set("uvsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) Vertices() js.Value {
	return pg.Get("vertices")
}
func (pg *PolyhedronGeometry) SetVertices(v js.Value) {
	pg.Set("vertices", v)
}
func (pg *PolyhedronGeometry) VerticesNeedUpdate() bool {
	return pg.Get("verticesNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetVerticesNeedUpdate(v bool) {
	pg.Set("verticesNeedUpdate", v)
}
func (pg *PolyhedronGeometry) AddEventListener(typ string, listener js.Value) {
	pg.Call("addEventListener", typ, listener)
}
func (pg *PolyhedronGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: pg.Call("applyMatrix", matrix)}
}
func (pg *PolyhedronGeometry) Center() Geometry {
	return &GeometryImpl{Value: pg.Call("center")}
}
func (pg *PolyhedronGeometry) Clone() Geometry {
	return &PolyhedronGeometry{Value: pg.Call("clone")}
}
func (pg *PolyhedronGeometry) ComputeBoundingBox() {
	pg.Call("computeBoundingBox")
}
func (pg *PolyhedronGeometry) ComputeBoundingSphere() {
	pg.Call("computeBoundingSphere")
}
func (pg *PolyhedronGeometry) ComputeFaceNormals() {
	pg.Call("computeFaceNormals")
}
func (pg *PolyhedronGeometry) ComputeFlatVertexNormals() {
	pg.Call("computeFlatVertexNormals")
}
func (pg *PolyhedronGeometry) ComputeMorphNormals() {
	pg.Call("computeMorphNormals")
}
func (pg *PolyhedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	pg.Call("computeVertexNormals", areaWeighted)
}
func (pg *PolyhedronGeometry) Copy(source Geometry) Geometry {
	return &PolyhedronGeometry{Value: pg.Call("copy", source.JSValue())}
}
func (pg *PolyhedronGeometry) DispatchEvent(event js.Value) {
	pg.Call("dispatchEvent", event)
}
func (pg *PolyhedronGeometry) Dispose() {
	pg.Call("dispose")
}
func (pg *PolyhedronGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: pg.Call("fromBufferGeometry", geometry)}
}
func (pg *PolyhedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pg.Call("hasEventListener", typ, listener).Bool()
}
func (pg *PolyhedronGeometry) LookAt(vector *Vector3) {
	pg.Call("lookAt", vector)
}
func (pg *PolyhedronGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	pg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (pg *PolyhedronGeometry) MergeMesh(mesh *Mesh) {
	pg.Call("mergeMesh", mesh)
}
func (pg *PolyhedronGeometry) MergeVertices() float64 {
	return pg.Call("mergeVertices").Float()
}
func (pg *PolyhedronGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: pg.Call("normalize")}
}
func (pg *PolyhedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	pg.Call("removeEventListener", typ, listener)
}
func (pg *PolyhedronGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateX", angle)}
}
func (pg *PolyhedronGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateY", angle)}
}
func (pg *PolyhedronGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateZ", angle)}
}
func (pg *PolyhedronGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: pg.Call("scale", x, y, z)}
}
func (pg *PolyhedronGeometry) SetFromPoints(points js.Value) Geometry {
	return &PolyhedronGeometry{Value: pg.Call("setFromPoints", points)}
}
func (pg *PolyhedronGeometry) SortFacesByMaterialIndex() {
	pg.Call("sortFacesByMaterialIndex")
}
func (pg *PolyhedronGeometry) ToJSON() js.Value {
	return pg.Call("toJSON")
}
func (pg *PolyhedronGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: pg.Call("translate", x, y, z)}
}

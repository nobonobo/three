// Code generated from "geometries/ParametricGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// ParametricBufferGeometry extend: [BufferGeometry]
type ParametricBufferGeometry struct {
	js.Value
}

func NewParametricBufferGeometry(fn js.Value, slices float64, stacks float64) *ParametricBufferGeometry {
	return &ParametricBufferGeometry{Value: get("ParametricBufferGeometry").New(fn, slices, stacks)}
}
func (pbg *ParametricBufferGeometry) JSValue() js.Value {
	return pbg.Value
}
func (pbg *ParametricBufferGeometry) Attributes() js.Value {
	return pbg.Get("attributes")
}
func (pbg *ParametricBufferGeometry) SetAttributes(v js.Value) {
	pbg.Set("attributes", v)
}
func (pbg *ParametricBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: pbg.Get("boundingBox")}
}
func (pbg *ParametricBufferGeometry) SetBoundingBox(v *Box3) {
	pbg.Set("boundingBox", v.Value)
}
func (pbg *ParametricBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: pbg.Get("boundingSphere")}
}
func (pbg *ParametricBufferGeometry) SetBoundingSphere(v *Sphere) {
	pbg.Set("boundingSphere", v.Value)
}
func (pbg *ParametricBufferGeometry) DrawRange() js.Value {
	return pbg.Get("drawRange")
}
func (pbg *ParametricBufferGeometry) SetDrawRange(v js.Value) {
	pbg.Set("drawRange", v)
}
func (pbg *ParametricBufferGeometry) Drawcalls() js.Value {
	return pbg.Get("drawcalls")
}
func (pbg *ParametricBufferGeometry) SetDrawcalls(v js.Value) {
	pbg.Set("drawcalls", v)
}
func (pbg *ParametricBufferGeometry) Groups() js.Value {
	return pbg.Get("groups")
}
func (pbg *ParametricBufferGeometry) SetGroups(v js.Value) {
	pbg.Set("groups", v)
}
func (pbg *ParametricBufferGeometry) Id() int {
	return pbg.Get("id").Int()
}
func (pbg *ParametricBufferGeometry) SetId(v int) {
	pbg.Set("id", v)
}
func (pbg *ParametricBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: pbg.Get("index")}
}
func (pbg *ParametricBufferGeometry) SetIndex(v *BufferAttribute) {
	pbg.Set("index", v.Value)
}
func (pbg *ParametricBufferGeometry) MorphAttributes() js.Value {
	return pbg.Get("morphAttributes")
}
func (pbg *ParametricBufferGeometry) SetMorphAttributes(v js.Value) {
	pbg.Set("morphAttributes", v)
}
func (pbg *ParametricBufferGeometry) Name() string {
	return pbg.Get("name").String()
}
func (pbg *ParametricBufferGeometry) SetName(v string) {
	pbg.Set("name", v)
}
func (pbg *ParametricBufferGeometry) Offsets() js.Value {
	return pbg.Get("offsets")
}
func (pbg *ParametricBufferGeometry) SetOffsets(v js.Value) {
	pbg.Set("offsets", v)
}
func (pbg *ParametricBufferGeometry) Parameters() js.Value {
	return pbg.Get("parameters")
}
func (pbg *ParametricBufferGeometry) SetParameters(v js.Value) {
	pbg.Set("parameters", v)
}
func (pbg *ParametricBufferGeometry) Type() string {
	return pbg.Get("type").String()
}
func (pbg *ParametricBufferGeometry) SetType(v string) {
	pbg.Set("type", v)
}
func (pbg *ParametricBufferGeometry) UserData() js.Value {
	return pbg.Get("userData")
}
func (pbg *ParametricBufferGeometry) SetUserData(v js.Value) {
	pbg.Set("userData", v)
}
func (pbg *ParametricBufferGeometry) Uuid() string {
	return pbg.Get("uuid").String()
}
func (pbg *ParametricBufferGeometry) SetUuid(v string) {
	pbg.Set("uuid", v)
}
func (pbg *ParametricBufferGeometry) MaxIndex() int {
	return pbg.Get("MaxIndex").Int()
}
func (pbg *ParametricBufferGeometry) SetMaxIndex(v int) {
	pbg.Set("MaxIndex", v)
}
func (pbg *ParametricBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("addAttribute", name, attribute)}
}
func (pbg *ParametricBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return pbg.Call("addAttribute", name, array, itemSize)
}
func (pbg *ParametricBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	pbg.Call("addDrawCall", start, count, indexOffset)
}
func (pbg *ParametricBufferGeometry) AddEventListener(typ string, listener js.Value) {
	pbg.Call("addEventListener", typ, listener)
}
func (pbg *ParametricBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	pbg.Call("addGroup", start, count, materialIndex)
}
func (pbg *ParametricBufferGeometry) AddIndex(index js.Value) {
	pbg.Call("addIndex", index)
}
func (pbg *ParametricBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("applyMatrix", matrix)}
}
func (pbg *ParametricBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("center")}
}
func (pbg *ParametricBufferGeometry) ClearDrawCalls() {
	pbg.Call("clearDrawCalls")
}
func (pbg *ParametricBufferGeometry) ClearGroups() {
	pbg.Call("clearGroups")
}
func (pbg *ParametricBufferGeometry) Clone() *ParametricBufferGeometry {
	return &ParametricBufferGeometry{Value: pbg.Call("clone")}
}
func (pbg *ParametricBufferGeometry) ComputeBoundingBox() {
	pbg.Call("computeBoundingBox")
}
func (pbg *ParametricBufferGeometry) ComputeBoundingSphere() {
	pbg.Call("computeBoundingSphere")
}
func (pbg *ParametricBufferGeometry) ComputeVertexNormals() {
	pbg.Call("computeVertexNormals")
}
func (pbg *ParametricBufferGeometry) Copy(source *BufferGeometry) *ParametricBufferGeometry {
	return &ParametricBufferGeometry{Value: pbg.Call("copy", source)}
}
func (pbg *ParametricBufferGeometry) DispatchEvent(event js.Value) {
	pbg.Call("dispatchEvent", event)
}
func (pbg *ParametricBufferGeometry) Dispose() {
	pbg.Call("dispose")
}
func (pbg *ParametricBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("fromDirectGeometry", geometry)}
}
func (pbg *ParametricBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (pbg *ParametricBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: pbg.Call("getAttribute", name)}
}
func (pbg *ParametricBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: pbg.Call("getIndex")}
}
func (pbg *ParametricBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pbg.Call("hasEventListener", typ, listener).Bool()
}
func (pbg *ParametricBufferGeometry) LookAt(v *Vector3) {
	pbg.Call("lookAt", v)
}
func (pbg *ParametricBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("merge", geometry, offset)}
}
func (pbg *ParametricBufferGeometry) NormalizeNormals() {
	pbg.Call("normalizeNormals")
}
func (pbg *ParametricBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("removeAttribute", name)}
}
func (pbg *ParametricBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	pbg.Call("removeEventListener", typ, listener)
}
func (pbg *ParametricBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("rotateX", angle)}
}
func (pbg *ParametricBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("rotateY", angle)}
}
func (pbg *ParametricBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("rotateZ", angle)}
}
func (pbg *ParametricBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("scale", x, y, z)}
}
func (pbg *ParametricBufferGeometry) SetDrawRange2(start int, count int) {
	pbg.Call("setDrawRange", start, count)
}
func (pbg *ParametricBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("setFromObject", object)}
}
func (pbg *ParametricBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("setFromPoints", points)}
}
func (pbg *ParametricBufferGeometry) SetIndex2(index *BufferAttribute) {
	pbg.Call("setIndex", index)
}
func (pbg *ParametricBufferGeometry) ToJSON() js.Value {
	return pbg.Call("toJSON")
}
func (pbg *ParametricBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("toNonIndexed")}
}
func (pbg *ParametricBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: pbg.Call("translate", x, y, z)}
}
func (pbg *ParametricBufferGeometry) UpdateFromObject(object *Object3D) {
	pbg.Call("updateFromObject", object)
}

// ParametricGeometry extend: [Geometry]
type ParametricGeometry struct {
	js.Value
}

func NewParametricGeometry(fn js.Value, slices float64, stacks float64) *ParametricGeometry {
	return &ParametricGeometry{Value: get("ParametricGeometry").New(fn, slices, stacks)}
}
func (pg *ParametricGeometry) JSValue() js.Value {
	return pg.Value
}
func (pg *ParametricGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: pg.Get("animation")}
}
func (pg *ParametricGeometry) SetAnimation(v *AnimationClip) {
	pg.Set("animation", v.Value)
}
func (pg *ParametricGeometry) Animations() js.Value {
	return pg.Get("animations")
}
func (pg *ParametricGeometry) SetAnimations(v js.Value) {
	pg.Set("animations", v)
}
func (pg *ParametricGeometry) Bones() js.Value {
	return pg.Get("bones")
}
func (pg *ParametricGeometry) SetBones(v js.Value) {
	pg.Set("bones", v)
}
func (pg *ParametricGeometry) BoundingBox() *Box3 {
	return &Box3{Value: pg.Get("boundingBox")}
}
func (pg *ParametricGeometry) SetBoundingBox(v *Box3) {
	pg.Set("boundingBox", v.Value)
}
func (pg *ParametricGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: pg.Get("boundingSphere")}
}
func (pg *ParametricGeometry) SetBoundingSphere(v *Sphere) {
	pg.Set("boundingSphere", v.Value)
}
func (pg *ParametricGeometry) Colors() js.Value {
	return pg.Get("colors")
}
func (pg *ParametricGeometry) SetColors(v js.Value) {
	pg.Set("colors", v)
}
func (pg *ParametricGeometry) ColorsNeedUpdate() bool {
	return pg.Get("colorsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetColorsNeedUpdate(v bool) {
	pg.Set("colorsNeedUpdate", v)
}
func (pg *ParametricGeometry) ElementsNeedUpdate() bool {
	return pg.Get("elementsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetElementsNeedUpdate(v bool) {
	pg.Set("elementsNeedUpdate", v)
}
func (pg *ParametricGeometry) FaceVertexUvs() js.Value {
	return pg.Get("faceVertexUvs")
}
func (pg *ParametricGeometry) SetFaceVertexUvs(v js.Value) {
	pg.Set("faceVertexUvs", v)
}
func (pg *ParametricGeometry) Faces() js.Value {
	return pg.Get("faces")
}
func (pg *ParametricGeometry) SetFaces(v js.Value) {
	pg.Set("faces", v)
}
func (pg *ParametricGeometry) GroupsNeedUpdate() bool {
	return pg.Get("groupsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetGroupsNeedUpdate(v bool) {
	pg.Set("groupsNeedUpdate", v)
}
func (pg *ParametricGeometry) Id() int {
	return pg.Get("id").Int()
}
func (pg *ParametricGeometry) SetId(v int) {
	pg.Set("id", v)
}
func (pg *ParametricGeometry) LineDistances() js.Value {
	return pg.Get("lineDistances")
}
func (pg *ParametricGeometry) SetLineDistances(v js.Value) {
	pg.Set("lineDistances", v)
}
func (pg *ParametricGeometry) LineDistancesNeedUpdate() bool {
	return pg.Get("lineDistancesNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetLineDistancesNeedUpdate(v bool) {
	pg.Set("lineDistancesNeedUpdate", v)
}
func (pg *ParametricGeometry) MorphNormals() js.Value {
	return pg.Get("morphNormals")
}
func (pg *ParametricGeometry) SetMorphNormals(v js.Value) {
	pg.Set("morphNormals", v)
}
func (pg *ParametricGeometry) MorphTargets() js.Value {
	return pg.Get("morphTargets")
}
func (pg *ParametricGeometry) SetMorphTargets(v js.Value) {
	pg.Set("morphTargets", v)
}
func (pg *ParametricGeometry) Name() string {
	return pg.Get("name").String()
}
func (pg *ParametricGeometry) SetName(v string) {
	pg.Set("name", v)
}
func (pg *ParametricGeometry) NormalsNeedUpdate() bool {
	return pg.Get("normalsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetNormalsNeedUpdate(v bool) {
	pg.Set("normalsNeedUpdate", v)
}
func (pg *ParametricGeometry) Parameters() js.Value {
	return pg.Get("parameters")
}
func (pg *ParametricGeometry) SetParameters(v js.Value) {
	pg.Set("parameters", v)
}
func (pg *ParametricGeometry) SkinIndices() js.Value {
	return pg.Get("skinIndices")
}
func (pg *ParametricGeometry) SetSkinIndices(v js.Value) {
	pg.Set("skinIndices", v)
}
func (pg *ParametricGeometry) SkinWeights() js.Value {
	return pg.Get("skinWeights")
}
func (pg *ParametricGeometry) SetSkinWeights(v js.Value) {
	pg.Set("skinWeights", v)
}
func (pg *ParametricGeometry) Type() string {
	return pg.Get("type").String()
}
func (pg *ParametricGeometry) SetType(v string) {
	pg.Set("type", v)
}
func (pg *ParametricGeometry) Uuid() string {
	return pg.Get("uuid").String()
}
func (pg *ParametricGeometry) SetUuid(v string) {
	pg.Set("uuid", v)
}
func (pg *ParametricGeometry) UvsNeedUpdate() bool {
	return pg.Get("uvsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetUvsNeedUpdate(v bool) {
	pg.Set("uvsNeedUpdate", v)
}
func (pg *ParametricGeometry) Vertices() js.Value {
	return pg.Get("vertices")
}
func (pg *ParametricGeometry) SetVertices(v js.Value) {
	pg.Set("vertices", v)
}
func (pg *ParametricGeometry) VerticesNeedUpdate() bool {
	return pg.Get("verticesNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetVerticesNeedUpdate(v bool) {
	pg.Set("verticesNeedUpdate", v)
}
func (pg *ParametricGeometry) AddEventListener(typ string, listener js.Value) {
	pg.Call("addEventListener", typ, listener)
}
func (pg *ParametricGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: pg.Call("applyMatrix", matrix)}
}
func (pg *ParametricGeometry) Center() Geometry {
	return &GeometryImpl{Value: pg.Call("center")}
}
func (pg *ParametricGeometry) Clone() Geometry {
	return &ParametricGeometry{Value: pg.Call("clone")}
}
func (pg *ParametricGeometry) ComputeBoundingBox() {
	pg.Call("computeBoundingBox")
}
func (pg *ParametricGeometry) ComputeBoundingSphere() {
	pg.Call("computeBoundingSphere")
}
func (pg *ParametricGeometry) ComputeFaceNormals() {
	pg.Call("computeFaceNormals")
}
func (pg *ParametricGeometry) ComputeFlatVertexNormals() {
	pg.Call("computeFlatVertexNormals")
}
func (pg *ParametricGeometry) ComputeMorphNormals() {
	pg.Call("computeMorphNormals")
}
func (pg *ParametricGeometry) ComputeVertexNormals(areaWeighted bool) {
	pg.Call("computeVertexNormals", areaWeighted)
}
func (pg *ParametricGeometry) Copy(source Geometry) Geometry {
	return &ParametricGeometry{Value: pg.Call("copy", source.JSValue())}
}
func (pg *ParametricGeometry) DispatchEvent(event js.Value) {
	pg.Call("dispatchEvent", event)
}
func (pg *ParametricGeometry) Dispose() {
	pg.Call("dispose")
}
func (pg *ParametricGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: pg.Call("fromBufferGeometry", geometry)}
}
func (pg *ParametricGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pg.Call("hasEventListener", typ, listener).Bool()
}
func (pg *ParametricGeometry) LookAt(vector *Vector3) {
	pg.Call("lookAt", vector)
}
func (pg *ParametricGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	pg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (pg *ParametricGeometry) MergeMesh(mesh *Mesh) {
	pg.Call("mergeMesh", mesh)
}
func (pg *ParametricGeometry) MergeVertices() float64 {
	return pg.Call("mergeVertices").Float()
}
func (pg *ParametricGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: pg.Call("normalize")}
}
func (pg *ParametricGeometry) RemoveEventListener(typ string, listener js.Value) {
	pg.Call("removeEventListener", typ, listener)
}
func (pg *ParametricGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateX", angle)}
}
func (pg *ParametricGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateY", angle)}
}
func (pg *ParametricGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: pg.Call("rotateZ", angle)}
}
func (pg *ParametricGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: pg.Call("scale", x, y, z)}
}
func (pg *ParametricGeometry) SetFromPoints(points js.Value) Geometry {
	return &ParametricGeometry{Value: pg.Call("setFromPoints", points)}
}
func (pg *ParametricGeometry) SortFacesByMaterialIndex() {
	pg.Call("sortFacesByMaterialIndex")
}
func (pg *ParametricGeometry) ToJSON() js.Value {
	return pg.Call("toJSON")
}
func (pg *ParametricGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: pg.Call("translate", x, y, z)}
}

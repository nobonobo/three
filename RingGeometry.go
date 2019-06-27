// Code generated from "geometries/RingGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// RingBufferGeometry extend: [BufferGeometry]
type RingBufferGeometry struct {
	js.Value
}

func NewRingBufferGeometry(innerRadius float64, outerRadius float64, thetaSegments int, phiSegments int, thetaStart float64, thetaLength float64) *RingBufferGeometry {
	return &RingBufferGeometry{Value: get("RingBufferGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)}
}
func (rbg *RingBufferGeometry) JSValue() js.Value {
	return rbg.Value
}
func (rbg *RingBufferGeometry) Attributes() js.Value {
	return rbg.Get("attributes")
}
func (rbg *RingBufferGeometry) SetAttributes(v js.Value) {
	rbg.Set("attributes", v)
}
func (rbg *RingBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: rbg.Get("boundingBox")}
}
func (rbg *RingBufferGeometry) SetBoundingBox(v *Box3) {
	rbg.Set("boundingBox", v.Value)
}
func (rbg *RingBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: rbg.Get("boundingSphere")}
}
func (rbg *RingBufferGeometry) SetBoundingSphere(v *Sphere) {
	rbg.Set("boundingSphere", v.Value)
}
func (rbg *RingBufferGeometry) DrawRange() js.Value {
	return rbg.Get("drawRange")
}
func (rbg *RingBufferGeometry) SetDrawRange(v js.Value) {
	rbg.Set("drawRange", v)
}
func (rbg *RingBufferGeometry) Drawcalls() js.Value {
	return rbg.Get("drawcalls")
}
func (rbg *RingBufferGeometry) SetDrawcalls(v js.Value) {
	rbg.Set("drawcalls", v)
}
func (rbg *RingBufferGeometry) Groups() js.Value {
	return rbg.Get("groups")
}
func (rbg *RingBufferGeometry) SetGroups(v js.Value) {
	rbg.Set("groups", v)
}
func (rbg *RingBufferGeometry) Id() int {
	return rbg.Get("id").Int()
}
func (rbg *RingBufferGeometry) SetId(v int) {
	rbg.Set("id", v)
}
func (rbg *RingBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: rbg.Get("index")}
}
func (rbg *RingBufferGeometry) SetIndex(v *BufferAttribute) {
	rbg.Set("index", v.Value)
}
func (rbg *RingBufferGeometry) MorphAttributes() js.Value {
	return rbg.Get("morphAttributes")
}
func (rbg *RingBufferGeometry) SetMorphAttributes(v js.Value) {
	rbg.Set("morphAttributes", v)
}
func (rbg *RingBufferGeometry) Name() string {
	return rbg.Get("name").String()
}
func (rbg *RingBufferGeometry) SetName(v string) {
	rbg.Set("name", v)
}
func (rbg *RingBufferGeometry) Offsets() js.Value {
	return rbg.Get("offsets")
}
func (rbg *RingBufferGeometry) SetOffsets(v js.Value) {
	rbg.Set("offsets", v)
}
func (rbg *RingBufferGeometry) Parameters() js.Value {
	return rbg.Get("parameters")
}
func (rbg *RingBufferGeometry) SetParameters(v js.Value) {
	rbg.Set("parameters", v)
}
func (rbg *RingBufferGeometry) Type() string {
	return rbg.Get("type").String()
}
func (rbg *RingBufferGeometry) SetType(v string) {
	rbg.Set("type", v)
}
func (rbg *RingBufferGeometry) UserData() js.Value {
	return rbg.Get("userData")
}
func (rbg *RingBufferGeometry) SetUserData(v js.Value) {
	rbg.Set("userData", v)
}
func (rbg *RingBufferGeometry) Uuid() string {
	return rbg.Get("uuid").String()
}
func (rbg *RingBufferGeometry) SetUuid(v string) {
	rbg.Set("uuid", v)
}
func (rbg *RingBufferGeometry) MaxIndex() int {
	return rbg.Get("MaxIndex").Int()
}
func (rbg *RingBufferGeometry) SetMaxIndex(v int) {
	rbg.Set("MaxIndex", v)
}
func (rbg *RingBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("addAttribute", name, attribute)}
}
func (rbg *RingBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return rbg.Call("addAttribute", name, array, itemSize)
}
func (rbg *RingBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	rbg.Call("addDrawCall", start, count, indexOffset)
}
func (rbg *RingBufferGeometry) AddEventListener(typ string, listener js.Value) {
	rbg.Call("addEventListener", typ, listener)
}
func (rbg *RingBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	rbg.Call("addGroup", start, count, materialIndex)
}
func (rbg *RingBufferGeometry) AddIndex(index js.Value) {
	rbg.Call("addIndex", index)
}
func (rbg *RingBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("applyMatrix", matrix)}
}
func (rbg *RingBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("center")}
}
func (rbg *RingBufferGeometry) ClearDrawCalls() {
	rbg.Call("clearDrawCalls")
}
func (rbg *RingBufferGeometry) ClearGroups() {
	rbg.Call("clearGroups")
}
func (rbg *RingBufferGeometry) Clone() *RingBufferGeometry {
	return &RingBufferGeometry{Value: rbg.Call("clone")}
}
func (rbg *RingBufferGeometry) ComputeBoundingBox() {
	rbg.Call("computeBoundingBox")
}
func (rbg *RingBufferGeometry) ComputeBoundingSphere() {
	rbg.Call("computeBoundingSphere")
}
func (rbg *RingBufferGeometry) ComputeVertexNormals() {
	rbg.Call("computeVertexNormals")
}
func (rbg *RingBufferGeometry) Copy(source *BufferGeometry) *RingBufferGeometry {
	return &RingBufferGeometry{Value: rbg.Call("copy", source)}
}
func (rbg *RingBufferGeometry) DispatchEvent(event js.Value) {
	rbg.Call("dispatchEvent", event)
}
func (rbg *RingBufferGeometry) Dispose() {
	rbg.Call("dispose")
}
func (rbg *RingBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("fromDirectGeometry", geometry)}
}
func (rbg *RingBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (rbg *RingBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: rbg.Call("getAttribute", name)}
}
func (rbg *RingBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: rbg.Call("getIndex")}
}
func (rbg *RingBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return rbg.Call("hasEventListener", typ, listener).Bool()
}
func (rbg *RingBufferGeometry) LookAt(v *Vector3) {
	rbg.Call("lookAt", v)
}
func (rbg *RingBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("merge", geometry, offset)}
}
func (rbg *RingBufferGeometry) NormalizeNormals() {
	rbg.Call("normalizeNormals")
}
func (rbg *RingBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("removeAttribute", name)}
}
func (rbg *RingBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	rbg.Call("removeEventListener", typ, listener)
}
func (rbg *RingBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("rotateX", angle)}
}
func (rbg *RingBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("rotateY", angle)}
}
func (rbg *RingBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("rotateZ", angle)}
}
func (rbg *RingBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("scale", x, y, z)}
}
func (rbg *RingBufferGeometry) SetDrawRange2(start int, count int) {
	rbg.Call("setDrawRange", start, count)
}
func (rbg *RingBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("setFromObject", object)}
}
func (rbg *RingBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("setFromPoints", points)}
}
func (rbg *RingBufferGeometry) SetIndex2(index *BufferAttribute) {
	rbg.Call("setIndex", index)
}
func (rbg *RingBufferGeometry) ToJSON() js.Value {
	return rbg.Call("toJSON")
}
func (rbg *RingBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("toNonIndexed")}
}
func (rbg *RingBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: rbg.Call("translate", x, y, z)}
}
func (rbg *RingBufferGeometry) UpdateFromObject(object *Object3D) {
	rbg.Call("updateFromObject", object)
}

// RingGeometry extend: [Geometry]
type RingGeometry struct {
	js.Value
}

func NewRingGeometry(innerRadius float64, outerRadius float64, thetaSegments int, phiSegments int, thetaStart float64, thetaLength float64) *RingGeometry {
	return &RingGeometry{Value: get("RingGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)}
}
func (rg *RingGeometry) JSValue() js.Value {
	return rg.Value
}
func (rg *RingGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: rg.Get("animation")}
}
func (rg *RingGeometry) SetAnimation(v *AnimationClip) {
	rg.Set("animation", v.Value)
}
func (rg *RingGeometry) Animations() js.Value {
	return rg.Get("animations")
}
func (rg *RingGeometry) SetAnimations(v js.Value) {
	rg.Set("animations", v)
}
func (rg *RingGeometry) Bones() js.Value {
	return rg.Get("bones")
}
func (rg *RingGeometry) SetBones(v js.Value) {
	rg.Set("bones", v)
}
func (rg *RingGeometry) BoundingBox() *Box3 {
	return &Box3{Value: rg.Get("boundingBox")}
}
func (rg *RingGeometry) SetBoundingBox(v *Box3) {
	rg.Set("boundingBox", v.Value)
}
func (rg *RingGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: rg.Get("boundingSphere")}
}
func (rg *RingGeometry) SetBoundingSphere(v *Sphere) {
	rg.Set("boundingSphere", v.Value)
}
func (rg *RingGeometry) Colors() js.Value {
	return rg.Get("colors")
}
func (rg *RingGeometry) SetColors(v js.Value) {
	rg.Set("colors", v)
}
func (rg *RingGeometry) ColorsNeedUpdate() bool {
	return rg.Get("colorsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetColorsNeedUpdate(v bool) {
	rg.Set("colorsNeedUpdate", v)
}
func (rg *RingGeometry) ElementsNeedUpdate() bool {
	return rg.Get("elementsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetElementsNeedUpdate(v bool) {
	rg.Set("elementsNeedUpdate", v)
}
func (rg *RingGeometry) FaceVertexUvs() js.Value {
	return rg.Get("faceVertexUvs")
}
func (rg *RingGeometry) SetFaceVertexUvs(v js.Value) {
	rg.Set("faceVertexUvs", v)
}
func (rg *RingGeometry) Faces() js.Value {
	return rg.Get("faces")
}
func (rg *RingGeometry) SetFaces(v js.Value) {
	rg.Set("faces", v)
}
func (rg *RingGeometry) GroupsNeedUpdate() bool {
	return rg.Get("groupsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetGroupsNeedUpdate(v bool) {
	rg.Set("groupsNeedUpdate", v)
}
func (rg *RingGeometry) Id() int {
	return rg.Get("id").Int()
}
func (rg *RingGeometry) SetId(v int) {
	rg.Set("id", v)
}
func (rg *RingGeometry) LineDistances() js.Value {
	return rg.Get("lineDistances")
}
func (rg *RingGeometry) SetLineDistances(v js.Value) {
	rg.Set("lineDistances", v)
}
func (rg *RingGeometry) LineDistancesNeedUpdate() bool {
	return rg.Get("lineDistancesNeedUpdate").Bool()
}
func (rg *RingGeometry) SetLineDistancesNeedUpdate(v bool) {
	rg.Set("lineDistancesNeedUpdate", v)
}
func (rg *RingGeometry) MorphNormals() js.Value {
	return rg.Get("morphNormals")
}
func (rg *RingGeometry) SetMorphNormals(v js.Value) {
	rg.Set("morphNormals", v)
}
func (rg *RingGeometry) MorphTargets() js.Value {
	return rg.Get("morphTargets")
}
func (rg *RingGeometry) SetMorphTargets(v js.Value) {
	rg.Set("morphTargets", v)
}
func (rg *RingGeometry) Name() string {
	return rg.Get("name").String()
}
func (rg *RingGeometry) SetName(v string) {
	rg.Set("name", v)
}
func (rg *RingGeometry) NormalsNeedUpdate() bool {
	return rg.Get("normalsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetNormalsNeedUpdate(v bool) {
	rg.Set("normalsNeedUpdate", v)
}
func (rg *RingGeometry) Parameters() js.Value {
	return rg.Get("parameters")
}
func (rg *RingGeometry) SetParameters(v js.Value) {
	rg.Set("parameters", v)
}
func (rg *RingGeometry) SkinIndices() js.Value {
	return rg.Get("skinIndices")
}
func (rg *RingGeometry) SetSkinIndices(v js.Value) {
	rg.Set("skinIndices", v)
}
func (rg *RingGeometry) SkinWeights() js.Value {
	return rg.Get("skinWeights")
}
func (rg *RingGeometry) SetSkinWeights(v js.Value) {
	rg.Set("skinWeights", v)
}
func (rg *RingGeometry) Type() string {
	return rg.Get("type").String()
}
func (rg *RingGeometry) SetType(v string) {
	rg.Set("type", v)
}
func (rg *RingGeometry) Uuid() string {
	return rg.Get("uuid").String()
}
func (rg *RingGeometry) SetUuid(v string) {
	rg.Set("uuid", v)
}
func (rg *RingGeometry) UvsNeedUpdate() bool {
	return rg.Get("uvsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetUvsNeedUpdate(v bool) {
	rg.Set("uvsNeedUpdate", v)
}
func (rg *RingGeometry) Vertices() js.Value {
	return rg.Get("vertices")
}
func (rg *RingGeometry) SetVertices(v js.Value) {
	rg.Set("vertices", v)
}
func (rg *RingGeometry) VerticesNeedUpdate() bool {
	return rg.Get("verticesNeedUpdate").Bool()
}
func (rg *RingGeometry) SetVerticesNeedUpdate(v bool) {
	rg.Set("verticesNeedUpdate", v)
}
func (rg *RingGeometry) AddEventListener(typ string, listener js.Value) {
	rg.Call("addEventListener", typ, listener)
}
func (rg *RingGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: rg.Call("applyMatrix", matrix)}
}
func (rg *RingGeometry) Center() Geometry {
	return &GeometryImpl{Value: rg.Call("center")}
}
func (rg *RingGeometry) Clone() Geometry {
	return &RingGeometry{Value: rg.Call("clone")}
}
func (rg *RingGeometry) ComputeBoundingBox() {
	rg.Call("computeBoundingBox")
}
func (rg *RingGeometry) ComputeBoundingSphere() {
	rg.Call("computeBoundingSphere")
}
func (rg *RingGeometry) ComputeFaceNormals() {
	rg.Call("computeFaceNormals")
}
func (rg *RingGeometry) ComputeFlatVertexNormals() {
	rg.Call("computeFlatVertexNormals")
}
func (rg *RingGeometry) ComputeMorphNormals() {
	rg.Call("computeMorphNormals")
}
func (rg *RingGeometry) ComputeVertexNormals(areaWeighted bool) {
	rg.Call("computeVertexNormals", areaWeighted)
}
func (rg *RingGeometry) Copy(source Geometry) Geometry {
	return &RingGeometry{Value: rg.Call("copy", source.JSValue())}
}
func (rg *RingGeometry) DispatchEvent(event js.Value) {
	rg.Call("dispatchEvent", event)
}
func (rg *RingGeometry) Dispose() {
	rg.Call("dispose")
}
func (rg *RingGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: rg.Call("fromBufferGeometry", geometry)}
}
func (rg *RingGeometry) HasEventListener(typ string, listener js.Value) bool {
	return rg.Call("hasEventListener", typ, listener).Bool()
}
func (rg *RingGeometry) LookAt(vector *Vector3) {
	rg.Call("lookAt", vector)
}
func (rg *RingGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	rg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (rg *RingGeometry) MergeMesh(mesh *Mesh) {
	rg.Call("mergeMesh", mesh)
}
func (rg *RingGeometry) MergeVertices() float64 {
	return rg.Call("mergeVertices").Float()
}
func (rg *RingGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: rg.Call("normalize")}
}
func (rg *RingGeometry) RemoveEventListener(typ string, listener js.Value) {
	rg.Call("removeEventListener", typ, listener)
}
func (rg *RingGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: rg.Call("rotateX", angle)}
}
func (rg *RingGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: rg.Call("rotateY", angle)}
}
func (rg *RingGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: rg.Call("rotateZ", angle)}
}
func (rg *RingGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: rg.Call("scale", x, y, z)}
}
func (rg *RingGeometry) SetFromPoints(points js.Value) Geometry {
	return &RingGeometry{Value: rg.Call("setFromPoints", points)}
}
func (rg *RingGeometry) SortFacesByMaterialIndex() {
	rg.Call("sortFacesByMaterialIndex")
}
func (rg *RingGeometry) ToJSON() js.Value {
	return rg.Call("toJSON")
}
func (rg *RingGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: rg.Call("translate", x, y, z)}
}

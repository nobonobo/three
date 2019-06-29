// Code generated from "geometries/ConeGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// ConeBufferGeometry extend: [BufferGeometry]
type ConeBufferGeometry struct {
	js.Value
}

func NewConeBufferGeometry(radius float64, height float64, radialSegment int, heightSegment int, openEnded bool, thetaStart float64, thetaLength float64) *ConeBufferGeometry {
	return &ConeBufferGeometry{Value: get("ConeBufferGeometry").New(radius, height, radialSegment, heightSegment, openEnded, thetaStart, thetaLength)}
}
func (cbg *ConeBufferGeometry) JSValue() js.Value {
	return cbg.Value
}
func (cbg *ConeBufferGeometry) Attributes() js.Value {
	return cbg.Get("attributes")
}
func (cbg *ConeBufferGeometry) SetAttributes(v js.Value) {
	cbg.Set("attributes", v)
}
func (cbg *ConeBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: cbg.Get("boundingBox")}
}
func (cbg *ConeBufferGeometry) SetBoundingBox(v *Box3) {
	cbg.Set("boundingBox", v.JSValue())
}
func (cbg *ConeBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: cbg.Get("boundingSphere")}
}
func (cbg *ConeBufferGeometry) SetBoundingSphere(v *Sphere) {
	cbg.Set("boundingSphere", v.JSValue())
}
func (cbg *ConeBufferGeometry) DrawRange() js.Value {
	return cbg.Get("drawRange")
}
func (cbg *ConeBufferGeometry) SetDrawRange(v js.Value) {
	cbg.Set("drawRange", v)
}
func (cbg *ConeBufferGeometry) Drawcalls() js.Value {
	return cbg.Get("drawcalls")
}
func (cbg *ConeBufferGeometry) SetDrawcalls(v js.Value) {
	cbg.Set("drawcalls", v)
}
func (cbg *ConeBufferGeometry) Groups() js.Value {
	return cbg.Get("groups")
}
func (cbg *ConeBufferGeometry) SetGroups(v js.Value) {
	cbg.Set("groups", v)
}
func (cbg *ConeBufferGeometry) Id() int {
	return cbg.Get("id").Int()
}
func (cbg *ConeBufferGeometry) SetId(v int) {
	cbg.Set("id", v)
}
func (cbg *ConeBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: cbg.Get("index")}
}
func (cbg *ConeBufferGeometry) SetIndex(v *BufferAttribute) {
	cbg.Set("index", v.JSValue())
}
func (cbg *ConeBufferGeometry) MorphAttributes() js.Value {
	return cbg.Get("morphAttributes")
}
func (cbg *ConeBufferGeometry) SetMorphAttributes(v js.Value) {
	cbg.Set("morphAttributes", v)
}
func (cbg *ConeBufferGeometry) Name() string {
	return cbg.Get("name").String()
}
func (cbg *ConeBufferGeometry) SetName(v string) {
	cbg.Set("name", v)
}
func (cbg *ConeBufferGeometry) Offsets() js.Value {
	return cbg.Get("offsets")
}
func (cbg *ConeBufferGeometry) SetOffsets(v js.Value) {
	cbg.Set("offsets", v)
}
func (cbg *ConeBufferGeometry) Type() string {
	return cbg.Get("type").String()
}
func (cbg *ConeBufferGeometry) SetType(v string) {
	cbg.Set("type", v)
}
func (cbg *ConeBufferGeometry) UserData() js.Value {
	return cbg.Get("userData")
}
func (cbg *ConeBufferGeometry) SetUserData(v js.Value) {
	cbg.Set("userData", v)
}
func (cbg *ConeBufferGeometry) Uuid() string {
	return cbg.Get("uuid").String()
}
func (cbg *ConeBufferGeometry) SetUuid(v string) {
	cbg.Set("uuid", v)
}
func (cbg *ConeBufferGeometry) MaxIndex() int {
	return cbg.Get("MaxIndex").Int()
}
func (cbg *ConeBufferGeometry) SetMaxIndex(v int) {
	cbg.Set("MaxIndex", v)
}
func (cbg *ConeBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("addAttribute", name, attribute)}
}
func (cbg *ConeBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return cbg.Call("addAttribute", name, array, itemSize)
}
func (cbg *ConeBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	cbg.Call("addDrawCall", start, count, indexOffset)
}
func (cbg *ConeBufferGeometry) AddEventListener(typ string, listener js.Value) {
	cbg.Call("addEventListener", typ, listener)
}
func (cbg *ConeBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	cbg.Call("addGroup", start, count, materialIndex)
}
func (cbg *ConeBufferGeometry) AddIndex(index js.Value) {
	cbg.Call("addIndex", index)
}
func (cbg *ConeBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("applyMatrix", matrix)}
}
func (cbg *ConeBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("center")}
}
func (cbg *ConeBufferGeometry) ClearDrawCalls() {
	cbg.Call("clearDrawCalls")
}
func (cbg *ConeBufferGeometry) ClearGroups() {
	cbg.Call("clearGroups")
}
func (cbg *ConeBufferGeometry) Clone() *ConeBufferGeometry {
	return &ConeBufferGeometry{Value: cbg.Call("clone")}
}
func (cbg *ConeBufferGeometry) ComputeBoundingBox() {
	cbg.Call("computeBoundingBox")
}
func (cbg *ConeBufferGeometry) ComputeBoundingSphere() {
	cbg.Call("computeBoundingSphere")
}
func (cbg *ConeBufferGeometry) ComputeVertexNormals() {
	cbg.Call("computeVertexNormals")
}
func (cbg *ConeBufferGeometry) Copy(source *BufferGeometry) *ConeBufferGeometry {
	return &ConeBufferGeometry{Value: cbg.Call("copy", source)}
}
func (cbg *ConeBufferGeometry) DispatchEvent(event js.Value) {
	cbg.Call("dispatchEvent", event)
}
func (cbg *ConeBufferGeometry) Dispose() {
	cbg.Call("dispose")
}
func (cbg *ConeBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("fromDirectGeometry", geometry)}
}
func (cbg *ConeBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (cbg *ConeBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: cbg.Call("getAttribute", name)}
}
func (cbg *ConeBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: cbg.Call("getIndex")}
}
func (cbg *ConeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cbg.Call("hasEventListener", typ, listener).Bool()
}
func (cbg *ConeBufferGeometry) LookAt(v *Vector3) {
	cbg.Call("lookAt", v.JSValue())
}
func (cbg *ConeBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("merge", geometry, offset)}
}
func (cbg *ConeBufferGeometry) NormalizeNormals() {
	cbg.Call("normalizeNormals")
}
func (cbg *ConeBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("removeAttribute", name)}
}
func (cbg *ConeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	cbg.Call("removeEventListener", typ, listener)
}
func (cbg *ConeBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateX", angle)}
}
func (cbg *ConeBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateY", angle)}
}
func (cbg *ConeBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateZ", angle)}
}
func (cbg *ConeBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("scale", x, y, z)}
}
func (cbg *ConeBufferGeometry) SetDrawRange2(start int, count int) {
	cbg.Call("setDrawRange", start, count)
}
func (cbg *ConeBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("setFromObject", object)}
}
func (cbg *ConeBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("setFromPoints", points)}
}
func (cbg *ConeBufferGeometry) SetIndex2(index *BufferAttribute) {
	cbg.Call("setIndex", index)
}
func (cbg *ConeBufferGeometry) ToJSON() js.Value {
	return cbg.Call("toJSON")
}
func (cbg *ConeBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("toNonIndexed")}
}
func (cbg *ConeBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("translate", x, y, z)}
}
func (cbg *ConeBufferGeometry) UpdateFromObject(object *Object3D) {
	cbg.Call("updateFromObject", object.JSValue())
}

// ConeGeometry extend: [CylinderGeometry]
type ConeGeometry struct {
	js.Value
}

func NewConeGeometry(radius float64, height float64, radialSegment int, heightSegment int, openEnded bool, thetaStart float64, thetaLength float64) *ConeGeometry {
	return &ConeGeometry{Value: get("ConeGeometry").New(radius, height, radialSegment, heightSegment, openEnded, thetaStart, thetaLength)}
}
func (cg *ConeGeometry) JSValue() js.Value {
	return cg.Value
}
func (cg *ConeGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: cg.Get("animation")}
}
func (cg *ConeGeometry) SetAnimation(v *AnimationClip) {
	cg.Set("animation", v.JSValue())
}
func (cg *ConeGeometry) Animations() js.Value {
	return cg.Get("animations")
}
func (cg *ConeGeometry) SetAnimations(v js.Value) {
	cg.Set("animations", v)
}
func (cg *ConeGeometry) Bones() js.Value {
	return cg.Get("bones")
}
func (cg *ConeGeometry) SetBones(v js.Value) {
	cg.Set("bones", v)
}
func (cg *ConeGeometry) BoundingBox() *Box3 {
	return &Box3{Value: cg.Get("boundingBox")}
}
func (cg *ConeGeometry) SetBoundingBox(v *Box3) {
	cg.Set("boundingBox", v.JSValue())
}
func (cg *ConeGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: cg.Get("boundingSphere")}
}
func (cg *ConeGeometry) SetBoundingSphere(v *Sphere) {
	cg.Set("boundingSphere", v.JSValue())
}
func (cg *ConeGeometry) Colors() js.Value {
	return cg.Get("colors")
}
func (cg *ConeGeometry) SetColors(v js.Value) {
	cg.Set("colors", v)
}
func (cg *ConeGeometry) ColorsNeedUpdate() bool {
	return cg.Get("colorsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetColorsNeedUpdate(v bool) {
	cg.Set("colorsNeedUpdate", v)
}
func (cg *ConeGeometry) ElementsNeedUpdate() bool {
	return cg.Get("elementsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetElementsNeedUpdate(v bool) {
	cg.Set("elementsNeedUpdate", v)
}
func (cg *ConeGeometry) FaceVertexUvs() js.Value {
	return cg.Get("faceVertexUvs")
}
func (cg *ConeGeometry) SetFaceVertexUvs(v js.Value) {
	cg.Set("faceVertexUvs", v)
}
func (cg *ConeGeometry) Faces() js.Value {
	return cg.Get("faces")
}
func (cg *ConeGeometry) SetFaces(v js.Value) {
	cg.Set("faces", v)
}
func (cg *ConeGeometry) GroupsNeedUpdate() bool {
	return cg.Get("groupsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetGroupsNeedUpdate(v bool) {
	cg.Set("groupsNeedUpdate", v)
}
func (cg *ConeGeometry) Id() int {
	return cg.Get("id").Int()
}
func (cg *ConeGeometry) SetId(v int) {
	cg.Set("id", v)
}
func (cg *ConeGeometry) LineDistances() js.Value {
	return cg.Get("lineDistances")
}
func (cg *ConeGeometry) SetLineDistances(v js.Value) {
	cg.Set("lineDistances", v)
}
func (cg *ConeGeometry) LineDistancesNeedUpdate() bool {
	return cg.Get("lineDistancesNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetLineDistancesNeedUpdate(v bool) {
	cg.Set("lineDistancesNeedUpdate", v)
}
func (cg *ConeGeometry) MorphNormals() js.Value {
	return cg.Get("morphNormals")
}
func (cg *ConeGeometry) SetMorphNormals(v js.Value) {
	cg.Set("morphNormals", v)
}
func (cg *ConeGeometry) MorphTargets() js.Value {
	return cg.Get("morphTargets")
}
func (cg *ConeGeometry) SetMorphTargets(v js.Value) {
	cg.Set("morphTargets", v)
}
func (cg *ConeGeometry) Name() string {
	return cg.Get("name").String()
}
func (cg *ConeGeometry) SetName(v string) {
	cg.Set("name", v)
}
func (cg *ConeGeometry) NormalsNeedUpdate() bool {
	return cg.Get("normalsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetNormalsNeedUpdate(v bool) {
	cg.Set("normalsNeedUpdate", v)
}
func (cg *ConeGeometry) Parameters() js.Value {
	return cg.Get("parameters")
}
func (cg *ConeGeometry) SetParameters(v js.Value) {
	cg.Set("parameters", v)
}
func (cg *ConeGeometry) SkinIndices() js.Value {
	return cg.Get("skinIndices")
}
func (cg *ConeGeometry) SetSkinIndices(v js.Value) {
	cg.Set("skinIndices", v)
}
func (cg *ConeGeometry) SkinWeights() js.Value {
	return cg.Get("skinWeights")
}
func (cg *ConeGeometry) SetSkinWeights(v js.Value) {
	cg.Set("skinWeights", v)
}
func (cg *ConeGeometry) Type() string {
	return cg.Get("type").String()
}
func (cg *ConeGeometry) SetType(v string) {
	cg.Set("type", v)
}
func (cg *ConeGeometry) Uuid() string {
	return cg.Get("uuid").String()
}
func (cg *ConeGeometry) SetUuid(v string) {
	cg.Set("uuid", v)
}
func (cg *ConeGeometry) UvsNeedUpdate() bool {
	return cg.Get("uvsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetUvsNeedUpdate(v bool) {
	cg.Set("uvsNeedUpdate", v)
}
func (cg *ConeGeometry) Vertices() js.Value {
	return cg.Get("vertices")
}
func (cg *ConeGeometry) SetVertices(v js.Value) {
	cg.Set("vertices", v)
}
func (cg *ConeGeometry) VerticesNeedUpdate() bool {
	return cg.Get("verticesNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetVerticesNeedUpdate(v bool) {
	cg.Set("verticesNeedUpdate", v)
}
func (cg *ConeGeometry) AddEventListener(typ string, listener js.Value) {
	cg.Call("addEventListener", typ, listener)
}
func (cg *ConeGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: cg.Call("applyMatrix", matrix)}
}
func (cg *ConeGeometry) Center() Geometry {
	return &GeometryImpl{Value: cg.Call("center")}
}
func (cg *ConeGeometry) Clone() Geometry {
	return &ConeGeometry{Value: cg.Call("clone")}
}
func (cg *ConeGeometry) ComputeBoundingBox() {
	cg.Call("computeBoundingBox")
}
func (cg *ConeGeometry) ComputeBoundingSphere() {
	cg.Call("computeBoundingSphere")
}
func (cg *ConeGeometry) ComputeFaceNormals() {
	cg.Call("computeFaceNormals")
}
func (cg *ConeGeometry) ComputeFlatVertexNormals() {
	cg.Call("computeFlatVertexNormals")
}
func (cg *ConeGeometry) ComputeMorphNormals() {
	cg.Call("computeMorphNormals")
}
func (cg *ConeGeometry) ComputeVertexNormals(areaWeighted bool) {
	cg.Call("computeVertexNormals", areaWeighted)
}
func (cg *ConeGeometry) Copy(source Geometry) Geometry {
	return &ConeGeometry{Value: cg.Call("copy", source.JSValue())}
}
func (cg *ConeGeometry) DispatchEvent(event js.Value) {
	cg.Call("dispatchEvent", event)
}
func (cg *ConeGeometry) Dispose() {
	cg.Call("dispose")
}
func (cg *ConeGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: cg.Call("fromBufferGeometry", geometry)}
}
func (cg *ConeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cg.Call("hasEventListener", typ, listener).Bool()
}
func (cg *ConeGeometry) LookAt(vector *Vector3) {
	cg.Call("lookAt", vector.JSValue())
}
func (cg *ConeGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	cg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (cg *ConeGeometry) MergeMesh(mesh Mesh) {
	cg.Call("mergeMesh", mesh.JSValue())
}
func (cg *ConeGeometry) MergeVertices() float64 {
	return cg.Call("mergeVertices").Float()
}
func (cg *ConeGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: cg.Call("normalize")}
}
func (cg *ConeGeometry) RemoveEventListener(typ string, listener js.Value) {
	cg.Call("removeEventListener", typ, listener)
}
func (cg *ConeGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateX", angle)}
}
func (cg *ConeGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateY", angle)}
}
func (cg *ConeGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateZ", angle)}
}
func (cg *ConeGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: cg.Call("scale", x, y, z)}
}
func (cg *ConeGeometry) SetFromPoints(points js.Value) Geometry {
	return &ConeGeometry{Value: cg.Call("setFromPoints", points)}
}
func (cg *ConeGeometry) SortFacesByMaterialIndex() {
	cg.Call("sortFacesByMaterialIndex")
}
func (cg *ConeGeometry) ToJSON() js.Value {
	return cg.Call("toJSON")
}
func (cg *ConeGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: cg.Call("translate", x, y, z)}
}

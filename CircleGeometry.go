// Code generated from "geometries/CircleGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CircleBufferGeometry extend: [BufferGeometry]
type CircleBufferGeometry struct {
	js.Value
}

func NewCircleBufferGeometry(radius float64, segments int, thetaStart float64, thetaLength float64) *CircleBufferGeometry {
	return &CircleBufferGeometry{Value: get("CircleBufferGeometry").New(radius, segments, thetaStart, thetaLength)}
}
func (cbg *CircleBufferGeometry) JSValue() js.Value {
	return cbg.Value
}
func (cbg *CircleBufferGeometry) Attributes() js.Value {
	return cbg.Get("attributes")
}
func (cbg *CircleBufferGeometry) SetAttributes(v js.Value) {
	cbg.Set("attributes", v)
}
func (cbg *CircleBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: cbg.Get("boundingBox")}
}
func (cbg *CircleBufferGeometry) SetBoundingBox(v *Box3) {
	cbg.Set("boundingBox", v.JSValue())
}
func (cbg *CircleBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: cbg.Get("boundingSphere")}
}
func (cbg *CircleBufferGeometry) SetBoundingSphere(v *Sphere) {
	cbg.Set("boundingSphere", v.JSValue())
}
func (cbg *CircleBufferGeometry) DrawRange() js.Value {
	return cbg.Get("drawRange")
}
func (cbg *CircleBufferGeometry) SetDrawRange(v js.Value) {
	cbg.Set("drawRange", v)
}
func (cbg *CircleBufferGeometry) Drawcalls() js.Value {
	return cbg.Get("drawcalls")
}
func (cbg *CircleBufferGeometry) SetDrawcalls(v js.Value) {
	cbg.Set("drawcalls", v)
}
func (cbg *CircleBufferGeometry) Groups() js.Value {
	return cbg.Get("groups")
}
func (cbg *CircleBufferGeometry) SetGroups(v js.Value) {
	cbg.Set("groups", v)
}
func (cbg *CircleBufferGeometry) Id() int {
	return cbg.Get("id").Int()
}
func (cbg *CircleBufferGeometry) SetId(v int) {
	cbg.Set("id", v)
}
func (cbg *CircleBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: cbg.Get("index")}
}
func (cbg *CircleBufferGeometry) SetIndex(v *BufferAttribute) {
	cbg.Set("index", v.JSValue())
}
func (cbg *CircleBufferGeometry) MorphAttributes() js.Value {
	return cbg.Get("morphAttributes")
}
func (cbg *CircleBufferGeometry) SetMorphAttributes(v js.Value) {
	cbg.Set("morphAttributes", v)
}
func (cbg *CircleBufferGeometry) Name() string {
	return cbg.Get("name").String()
}
func (cbg *CircleBufferGeometry) SetName(v string) {
	cbg.Set("name", v)
}
func (cbg *CircleBufferGeometry) Offsets() js.Value {
	return cbg.Get("offsets")
}
func (cbg *CircleBufferGeometry) SetOffsets(v js.Value) {
	cbg.Set("offsets", v)
}
func (cbg *CircleBufferGeometry) Parameters() js.Value {
	return cbg.Get("parameters")
}
func (cbg *CircleBufferGeometry) SetParameters(v js.Value) {
	cbg.Set("parameters", v)
}
func (cbg *CircleBufferGeometry) Type() string {
	return cbg.Get("type").String()
}
func (cbg *CircleBufferGeometry) SetType(v string) {
	cbg.Set("type", v)
}
func (cbg *CircleBufferGeometry) UserData() js.Value {
	return cbg.Get("userData")
}
func (cbg *CircleBufferGeometry) SetUserData(v js.Value) {
	cbg.Set("userData", v)
}
func (cbg *CircleBufferGeometry) Uuid() string {
	return cbg.Get("uuid").String()
}
func (cbg *CircleBufferGeometry) SetUuid(v string) {
	cbg.Set("uuid", v)
}
func (cbg *CircleBufferGeometry) MaxIndex() int {
	return cbg.Get("MaxIndex").Int()
}
func (cbg *CircleBufferGeometry) SetMaxIndex(v int) {
	cbg.Set("MaxIndex", v)
}
func (cbg *CircleBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("addAttribute", name, attribute)}
}
func (cbg *CircleBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return cbg.Call("addAttribute", name, array, itemSize)
}
func (cbg *CircleBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	cbg.Call("addDrawCall", start, count, indexOffset)
}
func (cbg *CircleBufferGeometry) AddEventListener(typ string, listener js.Value) {
	cbg.Call("addEventListener", typ, listener)
}
func (cbg *CircleBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	cbg.Call("addGroup", start, count, materialIndex)
}
func (cbg *CircleBufferGeometry) AddIndex(index js.Value) {
	cbg.Call("addIndex", index)
}
func (cbg *CircleBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("applyMatrix", matrix)}
}
func (cbg *CircleBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("center")}
}
func (cbg *CircleBufferGeometry) ClearDrawCalls() {
	cbg.Call("clearDrawCalls")
}
func (cbg *CircleBufferGeometry) ClearGroups() {
	cbg.Call("clearGroups")
}
func (cbg *CircleBufferGeometry) Clone() *CircleBufferGeometry {
	return &CircleBufferGeometry{Value: cbg.Call("clone")}
}
func (cbg *CircleBufferGeometry) ComputeBoundingBox() {
	cbg.Call("computeBoundingBox")
}
func (cbg *CircleBufferGeometry) ComputeBoundingSphere() {
	cbg.Call("computeBoundingSphere")
}
func (cbg *CircleBufferGeometry) ComputeVertexNormals() {
	cbg.Call("computeVertexNormals")
}
func (cbg *CircleBufferGeometry) Copy(source *BufferGeometry) *CircleBufferGeometry {
	return &CircleBufferGeometry{Value: cbg.Call("copy", source)}
}
func (cbg *CircleBufferGeometry) DispatchEvent(event js.Value) {
	cbg.Call("dispatchEvent", event)
}
func (cbg *CircleBufferGeometry) Dispose() {
	cbg.Call("dispose")
}
func (cbg *CircleBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("fromDirectGeometry", geometry)}
}
func (cbg *CircleBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (cbg *CircleBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: cbg.Call("getAttribute", name)}
}
func (cbg *CircleBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: cbg.Call("getIndex")}
}
func (cbg *CircleBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cbg.Call("hasEventListener", typ, listener).Bool()
}
func (cbg *CircleBufferGeometry) LookAt(v *Vector3) {
	cbg.Call("lookAt", v.JSValue())
}
func (cbg *CircleBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("merge", geometry, offset)}
}
func (cbg *CircleBufferGeometry) NormalizeNormals() {
	cbg.Call("normalizeNormals")
}
func (cbg *CircleBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("removeAttribute", name)}
}
func (cbg *CircleBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	cbg.Call("removeEventListener", typ, listener)
}
func (cbg *CircleBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateX", angle)}
}
func (cbg *CircleBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateY", angle)}
}
func (cbg *CircleBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateZ", angle)}
}
func (cbg *CircleBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("scale", x, y, z)}
}
func (cbg *CircleBufferGeometry) SetDrawRange2(start int, count int) {
	cbg.Call("setDrawRange", start, count)
}
func (cbg *CircleBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("setFromObject", object)}
}
func (cbg *CircleBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("setFromPoints", points)}
}
func (cbg *CircleBufferGeometry) SetIndex2(index *BufferAttribute) {
	cbg.Call("setIndex", index)
}
func (cbg *CircleBufferGeometry) ToJSON() js.Value {
	return cbg.Call("toJSON")
}
func (cbg *CircleBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("toNonIndexed")}
}
func (cbg *CircleBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("translate", x, y, z)}
}
func (cbg *CircleBufferGeometry) UpdateFromObject(object *Object3D) {
	cbg.Call("updateFromObject", object.JSValue())
}

// CircleGeometry extend: [Geometry]
type CircleGeometry struct {
	js.Value
}

func NewCircleGeometry(radius float64, segments int, thetaStart float64, thetaLength float64) *CircleGeometry {
	return &CircleGeometry{Value: get("CircleGeometry").New(radius, segments, thetaStart, thetaLength)}
}
func (cg *CircleGeometry) JSValue() js.Value {
	return cg.Value
}
func (cg *CircleGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: cg.Get("animation")}
}
func (cg *CircleGeometry) SetAnimation(v *AnimationClip) {
	cg.Set("animation", v.JSValue())
}
func (cg *CircleGeometry) Animations() js.Value {
	return cg.Get("animations")
}
func (cg *CircleGeometry) SetAnimations(v js.Value) {
	cg.Set("animations", v)
}
func (cg *CircleGeometry) Bones() js.Value {
	return cg.Get("bones")
}
func (cg *CircleGeometry) SetBones(v js.Value) {
	cg.Set("bones", v)
}
func (cg *CircleGeometry) BoundingBox() *Box3 {
	return &Box3{Value: cg.Get("boundingBox")}
}
func (cg *CircleGeometry) SetBoundingBox(v *Box3) {
	cg.Set("boundingBox", v.JSValue())
}
func (cg *CircleGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: cg.Get("boundingSphere")}
}
func (cg *CircleGeometry) SetBoundingSphere(v *Sphere) {
	cg.Set("boundingSphere", v.JSValue())
}
func (cg *CircleGeometry) Colors() js.Value {
	return cg.Get("colors")
}
func (cg *CircleGeometry) SetColors(v js.Value) {
	cg.Set("colors", v)
}
func (cg *CircleGeometry) ColorsNeedUpdate() bool {
	return cg.Get("colorsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetColorsNeedUpdate(v bool) {
	cg.Set("colorsNeedUpdate", v)
}
func (cg *CircleGeometry) ElementsNeedUpdate() bool {
	return cg.Get("elementsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetElementsNeedUpdate(v bool) {
	cg.Set("elementsNeedUpdate", v)
}
func (cg *CircleGeometry) FaceVertexUvs() js.Value {
	return cg.Get("faceVertexUvs")
}
func (cg *CircleGeometry) SetFaceVertexUvs(v js.Value) {
	cg.Set("faceVertexUvs", v)
}
func (cg *CircleGeometry) Faces() js.Value {
	return cg.Get("faces")
}
func (cg *CircleGeometry) SetFaces(v js.Value) {
	cg.Set("faces", v)
}
func (cg *CircleGeometry) GroupsNeedUpdate() bool {
	return cg.Get("groupsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetGroupsNeedUpdate(v bool) {
	cg.Set("groupsNeedUpdate", v)
}
func (cg *CircleGeometry) Id() int {
	return cg.Get("id").Int()
}
func (cg *CircleGeometry) SetId(v int) {
	cg.Set("id", v)
}
func (cg *CircleGeometry) LineDistances() js.Value {
	return cg.Get("lineDistances")
}
func (cg *CircleGeometry) SetLineDistances(v js.Value) {
	cg.Set("lineDistances", v)
}
func (cg *CircleGeometry) LineDistancesNeedUpdate() bool {
	return cg.Get("lineDistancesNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetLineDistancesNeedUpdate(v bool) {
	cg.Set("lineDistancesNeedUpdate", v)
}
func (cg *CircleGeometry) MorphNormals() js.Value {
	return cg.Get("morphNormals")
}
func (cg *CircleGeometry) SetMorphNormals(v js.Value) {
	cg.Set("morphNormals", v)
}
func (cg *CircleGeometry) MorphTargets() js.Value {
	return cg.Get("morphTargets")
}
func (cg *CircleGeometry) SetMorphTargets(v js.Value) {
	cg.Set("morphTargets", v)
}
func (cg *CircleGeometry) Name() string {
	return cg.Get("name").String()
}
func (cg *CircleGeometry) SetName(v string) {
	cg.Set("name", v)
}
func (cg *CircleGeometry) NormalsNeedUpdate() bool {
	return cg.Get("normalsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetNormalsNeedUpdate(v bool) {
	cg.Set("normalsNeedUpdate", v)
}
func (cg *CircleGeometry) Parameters() js.Value {
	return cg.Get("parameters")
}
func (cg *CircleGeometry) SetParameters(v js.Value) {
	cg.Set("parameters", v)
}
func (cg *CircleGeometry) SkinIndices() js.Value {
	return cg.Get("skinIndices")
}
func (cg *CircleGeometry) SetSkinIndices(v js.Value) {
	cg.Set("skinIndices", v)
}
func (cg *CircleGeometry) SkinWeights() js.Value {
	return cg.Get("skinWeights")
}
func (cg *CircleGeometry) SetSkinWeights(v js.Value) {
	cg.Set("skinWeights", v)
}
func (cg *CircleGeometry) Type() string {
	return cg.Get("type").String()
}
func (cg *CircleGeometry) SetType(v string) {
	cg.Set("type", v)
}
func (cg *CircleGeometry) Uuid() string {
	return cg.Get("uuid").String()
}
func (cg *CircleGeometry) SetUuid(v string) {
	cg.Set("uuid", v)
}
func (cg *CircleGeometry) UvsNeedUpdate() bool {
	return cg.Get("uvsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetUvsNeedUpdate(v bool) {
	cg.Set("uvsNeedUpdate", v)
}
func (cg *CircleGeometry) Vertices() js.Value {
	return cg.Get("vertices")
}
func (cg *CircleGeometry) SetVertices(v js.Value) {
	cg.Set("vertices", v)
}
func (cg *CircleGeometry) VerticesNeedUpdate() bool {
	return cg.Get("verticesNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetVerticesNeedUpdate(v bool) {
	cg.Set("verticesNeedUpdate", v)
}
func (cg *CircleGeometry) AddEventListener(typ string, listener js.Value) {
	cg.Call("addEventListener", typ, listener)
}
func (cg *CircleGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: cg.Call("applyMatrix", matrix)}
}
func (cg *CircleGeometry) Center() Geometry {
	return &GeometryImpl{Value: cg.Call("center")}
}
func (cg *CircleGeometry) Clone() Geometry {
	return &CircleGeometry{Value: cg.Call("clone")}
}
func (cg *CircleGeometry) ComputeBoundingBox() {
	cg.Call("computeBoundingBox")
}
func (cg *CircleGeometry) ComputeBoundingSphere() {
	cg.Call("computeBoundingSphere")
}
func (cg *CircleGeometry) ComputeFaceNormals() {
	cg.Call("computeFaceNormals")
}
func (cg *CircleGeometry) ComputeFlatVertexNormals() {
	cg.Call("computeFlatVertexNormals")
}
func (cg *CircleGeometry) ComputeMorphNormals() {
	cg.Call("computeMorphNormals")
}
func (cg *CircleGeometry) ComputeVertexNormals(areaWeighted bool) {
	cg.Call("computeVertexNormals", areaWeighted)
}
func (cg *CircleGeometry) Copy(source Geometry) Geometry {
	return &CircleGeometry{Value: cg.Call("copy", source.JSValue())}
}
func (cg *CircleGeometry) DispatchEvent(event js.Value) {
	cg.Call("dispatchEvent", event)
}
func (cg *CircleGeometry) Dispose() {
	cg.Call("dispose")
}
func (cg *CircleGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: cg.Call("fromBufferGeometry", geometry)}
}
func (cg *CircleGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cg.Call("hasEventListener", typ, listener).Bool()
}
func (cg *CircleGeometry) LookAt(vector *Vector3) {
	cg.Call("lookAt", vector.JSValue())
}
func (cg *CircleGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	cg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (cg *CircleGeometry) MergeMesh(mesh Mesh) {
	cg.Call("mergeMesh", mesh.JSValue())
}
func (cg *CircleGeometry) MergeVertices() float64 {
	return cg.Call("mergeVertices").Float()
}
func (cg *CircleGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: cg.Call("normalize")}
}
func (cg *CircleGeometry) RemoveEventListener(typ string, listener js.Value) {
	cg.Call("removeEventListener", typ, listener)
}
func (cg *CircleGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateX", angle)}
}
func (cg *CircleGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateY", angle)}
}
func (cg *CircleGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateZ", angle)}
}
func (cg *CircleGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: cg.Call("scale", x, y, z)}
}
func (cg *CircleGeometry) SetFromPoints(points js.Value) Geometry {
	return &CircleGeometry{Value: cg.Call("setFromPoints", points)}
}
func (cg *CircleGeometry) SortFacesByMaterialIndex() {
	cg.Call("sortFacesByMaterialIndex")
}
func (cg *CircleGeometry) ToJSON() js.Value {
	return cg.Call("toJSON")
}
func (cg *CircleGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: cg.Call("translate", x, y, z)}
}

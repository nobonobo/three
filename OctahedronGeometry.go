// Code generated from "geometries/OctahedronGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// OctahedronBufferGeometry extend: [PolyhedronBufferGeometry]
type OctahedronBufferGeometry struct {
	js.Value
}

func NewOctahedronBufferGeometry(radius float64, detail float64) *OctahedronBufferGeometry {
	return &OctahedronBufferGeometry{Value: get("OctahedronBufferGeometry").New(radius, detail)}
}
func (obg *OctahedronBufferGeometry) JSValue() js.Value {
	return obg.Value
}
func (obg *OctahedronBufferGeometry) Attributes() js.Value {
	return obg.Get("attributes")
}
func (obg *OctahedronBufferGeometry) SetAttributes(v js.Value) {
	obg.Set("attributes", v)
}
func (obg *OctahedronBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: obg.Get("boundingBox")}
}
func (obg *OctahedronBufferGeometry) SetBoundingBox(v *Box3) {
	obg.Set("boundingBox", v.JSValue())
}
func (obg *OctahedronBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: obg.Get("boundingSphere")}
}
func (obg *OctahedronBufferGeometry) SetBoundingSphere(v *Sphere) {
	obg.Set("boundingSphere", v.JSValue())
}
func (obg *OctahedronBufferGeometry) DrawRange() js.Value {
	return obg.Get("drawRange")
}
func (obg *OctahedronBufferGeometry) SetDrawRange(v js.Value) {
	obg.Set("drawRange", v)
}
func (obg *OctahedronBufferGeometry) Drawcalls() js.Value {
	return obg.Get("drawcalls")
}
func (obg *OctahedronBufferGeometry) SetDrawcalls(v js.Value) {
	obg.Set("drawcalls", v)
}
func (obg *OctahedronBufferGeometry) Groups() js.Value {
	return obg.Get("groups")
}
func (obg *OctahedronBufferGeometry) SetGroups(v js.Value) {
	obg.Set("groups", v)
}
func (obg *OctahedronBufferGeometry) Id() int {
	return obg.Get("id").Int()
}
func (obg *OctahedronBufferGeometry) SetId(v int) {
	obg.Set("id", v)
}
func (obg *OctahedronBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: obg.Get("index")}
}
func (obg *OctahedronBufferGeometry) SetIndex(v *BufferAttribute) {
	obg.Set("index", v.JSValue())
}
func (obg *OctahedronBufferGeometry) MorphAttributes() js.Value {
	return obg.Get("morphAttributes")
}
func (obg *OctahedronBufferGeometry) SetMorphAttributes(v js.Value) {
	obg.Set("morphAttributes", v)
}
func (obg *OctahedronBufferGeometry) Name() string {
	return obg.Get("name").String()
}
func (obg *OctahedronBufferGeometry) SetName(v string) {
	obg.Set("name", v)
}
func (obg *OctahedronBufferGeometry) Offsets() js.Value {
	return obg.Get("offsets")
}
func (obg *OctahedronBufferGeometry) SetOffsets(v js.Value) {
	obg.Set("offsets", v)
}
func (obg *OctahedronBufferGeometry) Parameters() js.Value {
	return obg.Get("parameters")
}
func (obg *OctahedronBufferGeometry) SetParameters(v js.Value) {
	obg.Set("parameters", v)
}
func (obg *OctahedronBufferGeometry) Type() string {
	return obg.Get("type").String()
}
func (obg *OctahedronBufferGeometry) SetType(v string) {
	obg.Set("type", v)
}
func (obg *OctahedronBufferGeometry) UserData() js.Value {
	return obg.Get("userData")
}
func (obg *OctahedronBufferGeometry) SetUserData(v js.Value) {
	obg.Set("userData", v)
}
func (obg *OctahedronBufferGeometry) Uuid() string {
	return obg.Get("uuid").String()
}
func (obg *OctahedronBufferGeometry) SetUuid(v string) {
	obg.Set("uuid", v)
}
func (obg *OctahedronBufferGeometry) MaxIndex() int {
	return obg.Get("MaxIndex").Int()
}
func (obg *OctahedronBufferGeometry) SetMaxIndex(v int) {
	obg.Set("MaxIndex", v)
}
func (obg *OctahedronBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("addAttribute", name, attribute)}
}
func (obg *OctahedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return obg.Call("addAttribute", name, array, itemSize)
}
func (obg *OctahedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	obg.Call("addDrawCall", start, count, indexOffset)
}
func (obg *OctahedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	obg.Call("addEventListener", typ, listener)
}
func (obg *OctahedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	obg.Call("addGroup", start, count, materialIndex)
}
func (obg *OctahedronBufferGeometry) AddIndex(index js.Value) {
	obg.Call("addIndex", index)
}
func (obg *OctahedronBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("applyMatrix", matrix)}
}
func (obg *OctahedronBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("center")}
}
func (obg *OctahedronBufferGeometry) ClearDrawCalls() {
	obg.Call("clearDrawCalls")
}
func (obg *OctahedronBufferGeometry) ClearGroups() {
	obg.Call("clearGroups")
}
func (obg *OctahedronBufferGeometry) Clone() *OctahedronBufferGeometry {
	return &OctahedronBufferGeometry{Value: obg.Call("clone")}
}
func (obg *OctahedronBufferGeometry) ComputeBoundingBox() {
	obg.Call("computeBoundingBox")
}
func (obg *OctahedronBufferGeometry) ComputeBoundingSphere() {
	obg.Call("computeBoundingSphere")
}
func (obg *OctahedronBufferGeometry) ComputeVertexNormals() {
	obg.Call("computeVertexNormals")
}
func (obg *OctahedronBufferGeometry) Copy(source *BufferGeometry) *OctahedronBufferGeometry {
	return &OctahedronBufferGeometry{Value: obg.Call("copy", source)}
}
func (obg *OctahedronBufferGeometry) DispatchEvent(event js.Value) {
	obg.Call("dispatchEvent", event)
}
func (obg *OctahedronBufferGeometry) Dispose() {
	obg.Call("dispose")
}
func (obg *OctahedronBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("fromDirectGeometry", geometry)}
}
func (obg *OctahedronBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (obg *OctahedronBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: obg.Call("getAttribute", name)}
}
func (obg *OctahedronBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: obg.Call("getIndex")}
}
func (obg *OctahedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return obg.Call("hasEventListener", typ, listener).Bool()
}
func (obg *OctahedronBufferGeometry) LookAt(v *Vector3) {
	obg.Call("lookAt", v.JSValue())
}
func (obg *OctahedronBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("merge", geometry, offset)}
}
func (obg *OctahedronBufferGeometry) NormalizeNormals() {
	obg.Call("normalizeNormals")
}
func (obg *OctahedronBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("removeAttribute", name)}
}
func (obg *OctahedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	obg.Call("removeEventListener", typ, listener)
}
func (obg *OctahedronBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("rotateX", angle)}
}
func (obg *OctahedronBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("rotateY", angle)}
}
func (obg *OctahedronBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("rotateZ", angle)}
}
func (obg *OctahedronBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("scale", x, y, z)}
}
func (obg *OctahedronBufferGeometry) SetDrawRange2(start int, count int) {
	obg.Call("setDrawRange", start, count)
}
func (obg *OctahedronBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("setFromObject", object)}
}
func (obg *OctahedronBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("setFromPoints", points)}
}
func (obg *OctahedronBufferGeometry) SetIndex2(index *BufferAttribute) {
	obg.Call("setIndex", index)
}
func (obg *OctahedronBufferGeometry) ToJSON() js.Value {
	return obg.Call("toJSON")
}
func (obg *OctahedronBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("toNonIndexed")}
}
func (obg *OctahedronBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: obg.Call("translate", x, y, z)}
}
func (obg *OctahedronBufferGeometry) UpdateFromObject(object *Object3D) {
	obg.Call("updateFromObject", object.JSValue())
}

// OctahedronGeometry extend: [PolyhedronGeometry]
type OctahedronGeometry struct {
	js.Value
}

func NewOctahedronGeometry(radius float64, detail float64) *OctahedronGeometry {
	return &OctahedronGeometry{Value: get("OctahedronGeometry").New(radius, detail)}
}
func (og *OctahedronGeometry) JSValue() js.Value {
	return og.Value
}
func (og *OctahedronGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: og.Get("animation")}
}
func (og *OctahedronGeometry) SetAnimation(v *AnimationClip) {
	og.Set("animation", v.JSValue())
}
func (og *OctahedronGeometry) Animations() js.Value {
	return og.Get("animations")
}
func (og *OctahedronGeometry) SetAnimations(v js.Value) {
	og.Set("animations", v)
}
func (og *OctahedronGeometry) Bones() js.Value {
	return og.Get("bones")
}
func (og *OctahedronGeometry) SetBones(v js.Value) {
	og.Set("bones", v)
}
func (og *OctahedronGeometry) BoundingBox() *Box3 {
	return &Box3{Value: og.Get("boundingBox")}
}
func (og *OctahedronGeometry) SetBoundingBox(v *Box3) {
	og.Set("boundingBox", v.JSValue())
}
func (og *OctahedronGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: og.Get("boundingSphere")}
}
func (og *OctahedronGeometry) SetBoundingSphere(v *Sphere) {
	og.Set("boundingSphere", v.JSValue())
}
func (og *OctahedronGeometry) Colors() js.Value {
	return og.Get("colors")
}
func (og *OctahedronGeometry) SetColors(v js.Value) {
	og.Set("colors", v)
}
func (og *OctahedronGeometry) ColorsNeedUpdate() bool {
	return og.Get("colorsNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetColorsNeedUpdate(v bool) {
	og.Set("colorsNeedUpdate", v)
}
func (og *OctahedronGeometry) ElementsNeedUpdate() bool {
	return og.Get("elementsNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetElementsNeedUpdate(v bool) {
	og.Set("elementsNeedUpdate", v)
}
func (og *OctahedronGeometry) FaceVertexUvs() js.Value {
	return og.Get("faceVertexUvs")
}
func (og *OctahedronGeometry) SetFaceVertexUvs(v js.Value) {
	og.Set("faceVertexUvs", v)
}
func (og *OctahedronGeometry) Faces() js.Value {
	return og.Get("faces")
}
func (og *OctahedronGeometry) SetFaces(v js.Value) {
	og.Set("faces", v)
}
func (og *OctahedronGeometry) GroupsNeedUpdate() bool {
	return og.Get("groupsNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetGroupsNeedUpdate(v bool) {
	og.Set("groupsNeedUpdate", v)
}
func (og *OctahedronGeometry) Id() int {
	return og.Get("id").Int()
}
func (og *OctahedronGeometry) SetId(v int) {
	og.Set("id", v)
}
func (og *OctahedronGeometry) LineDistances() js.Value {
	return og.Get("lineDistances")
}
func (og *OctahedronGeometry) SetLineDistances(v js.Value) {
	og.Set("lineDistances", v)
}
func (og *OctahedronGeometry) LineDistancesNeedUpdate() bool {
	return og.Get("lineDistancesNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	og.Set("lineDistancesNeedUpdate", v)
}
func (og *OctahedronGeometry) MorphNormals() js.Value {
	return og.Get("morphNormals")
}
func (og *OctahedronGeometry) SetMorphNormals(v js.Value) {
	og.Set("morphNormals", v)
}
func (og *OctahedronGeometry) MorphTargets() js.Value {
	return og.Get("morphTargets")
}
func (og *OctahedronGeometry) SetMorphTargets(v js.Value) {
	og.Set("morphTargets", v)
}
func (og *OctahedronGeometry) Name() string {
	return og.Get("name").String()
}
func (og *OctahedronGeometry) SetName(v string) {
	og.Set("name", v)
}
func (og *OctahedronGeometry) NormalsNeedUpdate() bool {
	return og.Get("normalsNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetNormalsNeedUpdate(v bool) {
	og.Set("normalsNeedUpdate", v)
}
func (og *OctahedronGeometry) Parameters() js.Value {
	return og.Get("parameters")
}
func (og *OctahedronGeometry) SetParameters(v js.Value) {
	og.Set("parameters", v)
}
func (og *OctahedronGeometry) SkinIndices() js.Value {
	return og.Get("skinIndices")
}
func (og *OctahedronGeometry) SetSkinIndices(v js.Value) {
	og.Set("skinIndices", v)
}
func (og *OctahedronGeometry) SkinWeights() js.Value {
	return og.Get("skinWeights")
}
func (og *OctahedronGeometry) SetSkinWeights(v js.Value) {
	og.Set("skinWeights", v)
}
func (og *OctahedronGeometry) Type() string {
	return og.Get("type").String()
}
func (og *OctahedronGeometry) SetType(v string) {
	og.Set("type", v)
}
func (og *OctahedronGeometry) Uuid() string {
	return og.Get("uuid").String()
}
func (og *OctahedronGeometry) SetUuid(v string) {
	og.Set("uuid", v)
}
func (og *OctahedronGeometry) UvsNeedUpdate() bool {
	return og.Get("uvsNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetUvsNeedUpdate(v bool) {
	og.Set("uvsNeedUpdate", v)
}
func (og *OctahedronGeometry) Vertices() js.Value {
	return og.Get("vertices")
}
func (og *OctahedronGeometry) SetVertices(v js.Value) {
	og.Set("vertices", v)
}
func (og *OctahedronGeometry) VerticesNeedUpdate() bool {
	return og.Get("verticesNeedUpdate").Bool()
}
func (og *OctahedronGeometry) SetVerticesNeedUpdate(v bool) {
	og.Set("verticesNeedUpdate", v)
}
func (og *OctahedronGeometry) AddEventListener(typ string, listener js.Value) {
	og.Call("addEventListener", typ, listener)
}
func (og *OctahedronGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: og.Call("applyMatrix", matrix)}
}
func (og *OctahedronGeometry) Center() Geometry {
	return &GeometryImpl{Value: og.Call("center")}
}
func (og *OctahedronGeometry) Clone() Geometry {
	return &OctahedronGeometry{Value: og.Call("clone")}
}
func (og *OctahedronGeometry) ComputeBoundingBox() {
	og.Call("computeBoundingBox")
}
func (og *OctahedronGeometry) ComputeBoundingSphere() {
	og.Call("computeBoundingSphere")
}
func (og *OctahedronGeometry) ComputeFaceNormals() {
	og.Call("computeFaceNormals")
}
func (og *OctahedronGeometry) ComputeFlatVertexNormals() {
	og.Call("computeFlatVertexNormals")
}
func (og *OctahedronGeometry) ComputeMorphNormals() {
	og.Call("computeMorphNormals")
}
func (og *OctahedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	og.Call("computeVertexNormals", areaWeighted)
}
func (og *OctahedronGeometry) Copy(source Geometry) Geometry {
	return &OctahedronGeometry{Value: og.Call("copy", source.JSValue())}
}
func (og *OctahedronGeometry) DispatchEvent(event js.Value) {
	og.Call("dispatchEvent", event)
}
func (og *OctahedronGeometry) Dispose() {
	og.Call("dispose")
}
func (og *OctahedronGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: og.Call("fromBufferGeometry", geometry)}
}
func (og *OctahedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return og.Call("hasEventListener", typ, listener).Bool()
}
func (og *OctahedronGeometry) LookAt(vector *Vector3) {
	og.Call("lookAt", vector.JSValue())
}
func (og *OctahedronGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	og.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (og *OctahedronGeometry) MergeMesh(mesh Mesh) {
	og.Call("mergeMesh", mesh.JSValue())
}
func (og *OctahedronGeometry) MergeVertices() float64 {
	return og.Call("mergeVertices").Float()
}
func (og *OctahedronGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: og.Call("normalize")}
}
func (og *OctahedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	og.Call("removeEventListener", typ, listener)
}
func (og *OctahedronGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: og.Call("rotateX", angle)}
}
func (og *OctahedronGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: og.Call("rotateY", angle)}
}
func (og *OctahedronGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: og.Call("rotateZ", angle)}
}
func (og *OctahedronGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: og.Call("scale", x, y, z)}
}
func (og *OctahedronGeometry) SetFromPoints(points js.Value) Geometry {
	return &OctahedronGeometry{Value: og.Call("setFromPoints", points)}
}
func (og *OctahedronGeometry) SortFacesByMaterialIndex() {
	og.Call("sortFacesByMaterialIndex")
}
func (og *OctahedronGeometry) ToJSON() js.Value {
	return og.Call("toJSON")
}
func (og *OctahedronGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: og.Call("translate", x, y, z)}
}

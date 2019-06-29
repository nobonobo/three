// Code generated from "geometries/SphereGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// SphereBufferGeometry extend: [BufferGeometry]
type SphereBufferGeometry struct {
	js.Value
}

func NewSphereBufferGeometry(radius float64, widthSegments int, heightSegments int, phiStart float64, phiLength float64, thetaStart float64, thetaLength float64) *SphereBufferGeometry {
	return &SphereBufferGeometry{Value: get("SphereBufferGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)}
}
func (sbg *SphereBufferGeometry) JSValue() js.Value {
	return sbg.Value
}
func (sbg *SphereBufferGeometry) Attributes() js.Value {
	return sbg.Get("attributes")
}
func (sbg *SphereBufferGeometry) SetAttributes(v js.Value) {
	sbg.Set("attributes", v)
}
func (sbg *SphereBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: sbg.Get("boundingBox")}
}
func (sbg *SphereBufferGeometry) SetBoundingBox(v *Box3) {
	sbg.Set("boundingBox", v.JSValue())
}
func (sbg *SphereBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: sbg.Get("boundingSphere")}
}
func (sbg *SphereBufferGeometry) SetBoundingSphere(v *Sphere) {
	sbg.Set("boundingSphere", v.JSValue())
}
func (sbg *SphereBufferGeometry) DrawRange() js.Value {
	return sbg.Get("drawRange")
}
func (sbg *SphereBufferGeometry) SetDrawRange(v js.Value) {
	sbg.Set("drawRange", v)
}
func (sbg *SphereBufferGeometry) Drawcalls() js.Value {
	return sbg.Get("drawcalls")
}
func (sbg *SphereBufferGeometry) SetDrawcalls(v js.Value) {
	sbg.Set("drawcalls", v)
}
func (sbg *SphereBufferGeometry) Groups() js.Value {
	return sbg.Get("groups")
}
func (sbg *SphereBufferGeometry) SetGroups(v js.Value) {
	sbg.Set("groups", v)
}
func (sbg *SphereBufferGeometry) Id() int {
	return sbg.Get("id").Int()
}
func (sbg *SphereBufferGeometry) SetId(v int) {
	sbg.Set("id", v)
}
func (sbg *SphereBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: sbg.Get("index")}
}
func (sbg *SphereBufferGeometry) SetIndex(v *BufferAttribute) {
	sbg.Set("index", v.JSValue())
}
func (sbg *SphereBufferGeometry) MorphAttributes() js.Value {
	return sbg.Get("morphAttributes")
}
func (sbg *SphereBufferGeometry) SetMorphAttributes(v js.Value) {
	sbg.Set("morphAttributes", v)
}
func (sbg *SphereBufferGeometry) Name() string {
	return sbg.Get("name").String()
}
func (sbg *SphereBufferGeometry) SetName(v string) {
	sbg.Set("name", v)
}
func (sbg *SphereBufferGeometry) Offsets() js.Value {
	return sbg.Get("offsets")
}
func (sbg *SphereBufferGeometry) SetOffsets(v js.Value) {
	sbg.Set("offsets", v)
}
func (sbg *SphereBufferGeometry) Parameters() js.Value {
	return sbg.Get("parameters")
}
func (sbg *SphereBufferGeometry) SetParameters(v js.Value) {
	sbg.Set("parameters", v)
}
func (sbg *SphereBufferGeometry) Type() string {
	return sbg.Get("type").String()
}
func (sbg *SphereBufferGeometry) SetType(v string) {
	sbg.Set("type", v)
}
func (sbg *SphereBufferGeometry) UserData() js.Value {
	return sbg.Get("userData")
}
func (sbg *SphereBufferGeometry) SetUserData(v js.Value) {
	sbg.Set("userData", v)
}
func (sbg *SphereBufferGeometry) Uuid() string {
	return sbg.Get("uuid").String()
}
func (sbg *SphereBufferGeometry) SetUuid(v string) {
	sbg.Set("uuid", v)
}
func (sbg *SphereBufferGeometry) MaxIndex() int {
	return sbg.Get("MaxIndex").Int()
}
func (sbg *SphereBufferGeometry) SetMaxIndex(v int) {
	sbg.Set("MaxIndex", v)
}
func (sbg *SphereBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("addAttribute", name, attribute)}
}
func (sbg *SphereBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return sbg.Call("addAttribute", name, array, itemSize)
}
func (sbg *SphereBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	sbg.Call("addDrawCall", start, count, indexOffset)
}
func (sbg *SphereBufferGeometry) AddEventListener(typ string, listener js.Value) {
	sbg.Call("addEventListener", typ, listener)
}
func (sbg *SphereBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	sbg.Call("addGroup", start, count, materialIndex)
}
func (sbg *SphereBufferGeometry) AddIndex(index js.Value) {
	sbg.Call("addIndex", index)
}
func (sbg *SphereBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("applyMatrix", matrix)}
}
func (sbg *SphereBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("center")}
}
func (sbg *SphereBufferGeometry) ClearDrawCalls() {
	sbg.Call("clearDrawCalls")
}
func (sbg *SphereBufferGeometry) ClearGroups() {
	sbg.Call("clearGroups")
}
func (sbg *SphereBufferGeometry) Clone() *SphereBufferGeometry {
	return &SphereBufferGeometry{Value: sbg.Call("clone")}
}
func (sbg *SphereBufferGeometry) ComputeBoundingBox() {
	sbg.Call("computeBoundingBox")
}
func (sbg *SphereBufferGeometry) ComputeBoundingSphere() {
	sbg.Call("computeBoundingSphere")
}
func (sbg *SphereBufferGeometry) ComputeVertexNormals() {
	sbg.Call("computeVertexNormals")
}
func (sbg *SphereBufferGeometry) Copy(source *BufferGeometry) *SphereBufferGeometry {
	return &SphereBufferGeometry{Value: sbg.Call("copy", source)}
}
func (sbg *SphereBufferGeometry) DispatchEvent(event js.Value) {
	sbg.Call("dispatchEvent", event)
}
func (sbg *SphereBufferGeometry) Dispose() {
	sbg.Call("dispose")
}
func (sbg *SphereBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("fromDirectGeometry", geometry)}
}
func (sbg *SphereBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (sbg *SphereBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: sbg.Call("getAttribute", name)}
}
func (sbg *SphereBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: sbg.Call("getIndex")}
}
func (sbg *SphereBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return sbg.Call("hasEventListener", typ, listener).Bool()
}
func (sbg *SphereBufferGeometry) LookAt(v *Vector3) {
	sbg.Call("lookAt", v.JSValue())
}
func (sbg *SphereBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("merge", geometry, offset)}
}
func (sbg *SphereBufferGeometry) NormalizeNormals() {
	sbg.Call("normalizeNormals")
}
func (sbg *SphereBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("removeAttribute", name)}
}
func (sbg *SphereBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	sbg.Call("removeEventListener", typ, listener)
}
func (sbg *SphereBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("rotateX", angle)}
}
func (sbg *SphereBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("rotateY", angle)}
}
func (sbg *SphereBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("rotateZ", angle)}
}
func (sbg *SphereBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("scale", x, y, z)}
}
func (sbg *SphereBufferGeometry) SetDrawRange2(start int, count int) {
	sbg.Call("setDrawRange", start, count)
}
func (sbg *SphereBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("setFromObject", object)}
}
func (sbg *SphereBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("setFromPoints", points)}
}
func (sbg *SphereBufferGeometry) SetIndex2(index *BufferAttribute) {
	sbg.Call("setIndex", index)
}
func (sbg *SphereBufferGeometry) ToJSON() js.Value {
	return sbg.Call("toJSON")
}
func (sbg *SphereBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("toNonIndexed")}
}
func (sbg *SphereBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: sbg.Call("translate", x, y, z)}
}
func (sbg *SphereBufferGeometry) UpdateFromObject(object *Object3D) {
	sbg.Call("updateFromObject", object.JSValue())
}

// SphereGeometry extend: [Geometry]
type SphereGeometry struct {
	js.Value
}

func NewSphereGeometry(radius float64, widthSegments int, heightSegments int, phiStart float64, phiLength float64, thetaStart float64, thetaLength float64) *SphereGeometry {
	return &SphereGeometry{Value: get("SphereGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)}
}
func (sg *SphereGeometry) JSValue() js.Value {
	return sg.Value
}
func (sg *SphereGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: sg.Get("animation")}
}
func (sg *SphereGeometry) SetAnimation(v *AnimationClip) {
	sg.Set("animation", v.JSValue())
}
func (sg *SphereGeometry) Animations() js.Value {
	return sg.Get("animations")
}
func (sg *SphereGeometry) SetAnimations(v js.Value) {
	sg.Set("animations", v)
}
func (sg *SphereGeometry) Bones() js.Value {
	return sg.Get("bones")
}
func (sg *SphereGeometry) SetBones(v js.Value) {
	sg.Set("bones", v)
}
func (sg *SphereGeometry) BoundingBox() *Box3 {
	return &Box3{Value: sg.Get("boundingBox")}
}
func (sg *SphereGeometry) SetBoundingBox(v *Box3) {
	sg.Set("boundingBox", v.JSValue())
}
func (sg *SphereGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: sg.Get("boundingSphere")}
}
func (sg *SphereGeometry) SetBoundingSphere(v *Sphere) {
	sg.Set("boundingSphere", v.JSValue())
}
func (sg *SphereGeometry) Colors() js.Value {
	return sg.Get("colors")
}
func (sg *SphereGeometry) SetColors(v js.Value) {
	sg.Set("colors", v)
}
func (sg *SphereGeometry) ColorsNeedUpdate() bool {
	return sg.Get("colorsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetColorsNeedUpdate(v bool) {
	sg.Set("colorsNeedUpdate", v)
}
func (sg *SphereGeometry) ElementsNeedUpdate() bool {
	return sg.Get("elementsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetElementsNeedUpdate(v bool) {
	sg.Set("elementsNeedUpdate", v)
}
func (sg *SphereGeometry) FaceVertexUvs() js.Value {
	return sg.Get("faceVertexUvs")
}
func (sg *SphereGeometry) SetFaceVertexUvs(v js.Value) {
	sg.Set("faceVertexUvs", v)
}
func (sg *SphereGeometry) Faces() js.Value {
	return sg.Get("faces")
}
func (sg *SphereGeometry) SetFaces(v js.Value) {
	sg.Set("faces", v)
}
func (sg *SphereGeometry) GroupsNeedUpdate() bool {
	return sg.Get("groupsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetGroupsNeedUpdate(v bool) {
	sg.Set("groupsNeedUpdate", v)
}
func (sg *SphereGeometry) Id() int {
	return sg.Get("id").Int()
}
func (sg *SphereGeometry) SetId(v int) {
	sg.Set("id", v)
}
func (sg *SphereGeometry) LineDistances() js.Value {
	return sg.Get("lineDistances")
}
func (sg *SphereGeometry) SetLineDistances(v js.Value) {
	sg.Set("lineDistances", v)
}
func (sg *SphereGeometry) LineDistancesNeedUpdate() bool {
	return sg.Get("lineDistancesNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetLineDistancesNeedUpdate(v bool) {
	sg.Set("lineDistancesNeedUpdate", v)
}
func (sg *SphereGeometry) MorphNormals() js.Value {
	return sg.Get("morphNormals")
}
func (sg *SphereGeometry) SetMorphNormals(v js.Value) {
	sg.Set("morphNormals", v)
}
func (sg *SphereGeometry) MorphTargets() js.Value {
	return sg.Get("morphTargets")
}
func (sg *SphereGeometry) SetMorphTargets(v js.Value) {
	sg.Set("morphTargets", v)
}
func (sg *SphereGeometry) Name() string {
	return sg.Get("name").String()
}
func (sg *SphereGeometry) SetName(v string) {
	sg.Set("name", v)
}
func (sg *SphereGeometry) NormalsNeedUpdate() bool {
	return sg.Get("normalsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetNormalsNeedUpdate(v bool) {
	sg.Set("normalsNeedUpdate", v)
}
func (sg *SphereGeometry) Parameters() js.Value {
	return sg.Get("parameters")
}
func (sg *SphereGeometry) SetParameters(v js.Value) {
	sg.Set("parameters", v)
}
func (sg *SphereGeometry) SkinIndices() js.Value {
	return sg.Get("skinIndices")
}
func (sg *SphereGeometry) SetSkinIndices(v js.Value) {
	sg.Set("skinIndices", v)
}
func (sg *SphereGeometry) SkinWeights() js.Value {
	return sg.Get("skinWeights")
}
func (sg *SphereGeometry) SetSkinWeights(v js.Value) {
	sg.Set("skinWeights", v)
}
func (sg *SphereGeometry) Type() string {
	return sg.Get("type").String()
}
func (sg *SphereGeometry) SetType(v string) {
	sg.Set("type", v)
}
func (sg *SphereGeometry) Uuid() string {
	return sg.Get("uuid").String()
}
func (sg *SphereGeometry) SetUuid(v string) {
	sg.Set("uuid", v)
}
func (sg *SphereGeometry) UvsNeedUpdate() bool {
	return sg.Get("uvsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetUvsNeedUpdate(v bool) {
	sg.Set("uvsNeedUpdate", v)
}
func (sg *SphereGeometry) Vertices() js.Value {
	return sg.Get("vertices")
}
func (sg *SphereGeometry) SetVertices(v js.Value) {
	sg.Set("vertices", v)
}
func (sg *SphereGeometry) VerticesNeedUpdate() bool {
	return sg.Get("verticesNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetVerticesNeedUpdate(v bool) {
	sg.Set("verticesNeedUpdate", v)
}
func (sg *SphereGeometry) AddEventListener(typ string, listener js.Value) {
	sg.Call("addEventListener", typ, listener)
}
func (sg *SphereGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: sg.Call("applyMatrix", matrix)}
}
func (sg *SphereGeometry) Center() Geometry {
	return &GeometryImpl{Value: sg.Call("center")}
}
func (sg *SphereGeometry) Clone() Geometry {
	return &SphereGeometry{Value: sg.Call("clone")}
}
func (sg *SphereGeometry) ComputeBoundingBox() {
	sg.Call("computeBoundingBox")
}
func (sg *SphereGeometry) ComputeBoundingSphere() {
	sg.Call("computeBoundingSphere")
}
func (sg *SphereGeometry) ComputeFaceNormals() {
	sg.Call("computeFaceNormals")
}
func (sg *SphereGeometry) ComputeFlatVertexNormals() {
	sg.Call("computeFlatVertexNormals")
}
func (sg *SphereGeometry) ComputeMorphNormals() {
	sg.Call("computeMorphNormals")
}
func (sg *SphereGeometry) ComputeVertexNormals(areaWeighted bool) {
	sg.Call("computeVertexNormals", areaWeighted)
}
func (sg *SphereGeometry) Copy(source Geometry) Geometry {
	return &SphereGeometry{Value: sg.Call("copy", source.JSValue())}
}
func (sg *SphereGeometry) DispatchEvent(event js.Value) {
	sg.Call("dispatchEvent", event)
}
func (sg *SphereGeometry) Dispose() {
	sg.Call("dispose")
}
func (sg *SphereGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: sg.Call("fromBufferGeometry", geometry)}
}
func (sg *SphereGeometry) HasEventListener(typ string, listener js.Value) bool {
	return sg.Call("hasEventListener", typ, listener).Bool()
}
func (sg *SphereGeometry) LookAt(vector *Vector3) {
	sg.Call("lookAt", vector.JSValue())
}
func (sg *SphereGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	sg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (sg *SphereGeometry) MergeMesh(mesh Mesh) {
	sg.Call("mergeMesh", mesh.JSValue())
}
func (sg *SphereGeometry) MergeVertices() float64 {
	return sg.Call("mergeVertices").Float()
}
func (sg *SphereGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: sg.Call("normalize")}
}
func (sg *SphereGeometry) RemoveEventListener(typ string, listener js.Value) {
	sg.Call("removeEventListener", typ, listener)
}
func (sg *SphereGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: sg.Call("rotateX", angle)}
}
func (sg *SphereGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: sg.Call("rotateY", angle)}
}
func (sg *SphereGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: sg.Call("rotateZ", angle)}
}
func (sg *SphereGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: sg.Call("scale", x, y, z)}
}
func (sg *SphereGeometry) SetFromPoints(points js.Value) Geometry {
	return &SphereGeometry{Value: sg.Call("setFromPoints", points)}
}
func (sg *SphereGeometry) SortFacesByMaterialIndex() {
	sg.Call("sortFacesByMaterialIndex")
}
func (sg *SphereGeometry) ToJSON() js.Value {
	return sg.Call("toJSON")
}
func (sg *SphereGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: sg.Call("translate", x, y, z)}
}

// Code generated from "geometries/CylinderGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CylinderBufferGeometry extend: [BufferGeometry]
type CylinderBufferGeometry struct {
	js.Value
}

func NewCylinderBufferGeometry(radiusTop float64, radiusBottom float64, height float64, radialSegments int, heightSegments int, openEnded bool, thetaStart float64, thetaLength float64) *CylinderBufferGeometry {
	return &CylinderBufferGeometry{Value: get("CylinderBufferGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)}
}
func (cbg *CylinderBufferGeometry) JSValue() js.Value {
	return cbg.Value
}
func (cbg *CylinderBufferGeometry) Attributes() js.Value {
	return cbg.Get("attributes")
}
func (cbg *CylinderBufferGeometry) SetAttributes(v js.Value) {
	cbg.Set("attributes", v)
}
func (cbg *CylinderBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: cbg.Get("boundingBox")}
}
func (cbg *CylinderBufferGeometry) SetBoundingBox(v *Box3) {
	cbg.Set("boundingBox", v.JSValue())
}
func (cbg *CylinderBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: cbg.Get("boundingSphere")}
}
func (cbg *CylinderBufferGeometry) SetBoundingSphere(v *Sphere) {
	cbg.Set("boundingSphere", v.JSValue())
}
func (cbg *CylinderBufferGeometry) DrawRange() js.Value {
	return cbg.Get("drawRange")
}
func (cbg *CylinderBufferGeometry) SetDrawRange(v js.Value) {
	cbg.Set("drawRange", v)
}
func (cbg *CylinderBufferGeometry) Drawcalls() js.Value {
	return cbg.Get("drawcalls")
}
func (cbg *CylinderBufferGeometry) SetDrawcalls(v js.Value) {
	cbg.Set("drawcalls", v)
}
func (cbg *CylinderBufferGeometry) Groups() js.Value {
	return cbg.Get("groups")
}
func (cbg *CylinderBufferGeometry) SetGroups(v js.Value) {
	cbg.Set("groups", v)
}
func (cbg *CylinderBufferGeometry) Id() int {
	return cbg.Get("id").Int()
}
func (cbg *CylinderBufferGeometry) SetId(v int) {
	cbg.Set("id", v)
}
func (cbg *CylinderBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: cbg.Get("index")}
}
func (cbg *CylinderBufferGeometry) SetIndex(v *BufferAttribute) {
	cbg.Set("index", v.JSValue())
}
func (cbg *CylinderBufferGeometry) MorphAttributes() js.Value {
	return cbg.Get("morphAttributes")
}
func (cbg *CylinderBufferGeometry) SetMorphAttributes(v js.Value) {
	cbg.Set("morphAttributes", v)
}
func (cbg *CylinderBufferGeometry) Name() string {
	return cbg.Get("name").String()
}
func (cbg *CylinderBufferGeometry) SetName(v string) {
	cbg.Set("name", v)
}
func (cbg *CylinderBufferGeometry) Offsets() js.Value {
	return cbg.Get("offsets")
}
func (cbg *CylinderBufferGeometry) SetOffsets(v js.Value) {
	cbg.Set("offsets", v)
}
func (cbg *CylinderBufferGeometry) Parameters() js.Value {
	return cbg.Get("parameters")
}
func (cbg *CylinderBufferGeometry) SetParameters(v js.Value) {
	cbg.Set("parameters", v)
}
func (cbg *CylinderBufferGeometry) Type() string {
	return cbg.Get("type").String()
}
func (cbg *CylinderBufferGeometry) SetType(v string) {
	cbg.Set("type", v)
}
func (cbg *CylinderBufferGeometry) UserData() js.Value {
	return cbg.Get("userData")
}
func (cbg *CylinderBufferGeometry) SetUserData(v js.Value) {
	cbg.Set("userData", v)
}
func (cbg *CylinderBufferGeometry) Uuid() string {
	return cbg.Get("uuid").String()
}
func (cbg *CylinderBufferGeometry) SetUuid(v string) {
	cbg.Set("uuid", v)
}
func (cbg *CylinderBufferGeometry) MaxIndex() int {
	return cbg.Get("MaxIndex").Int()
}
func (cbg *CylinderBufferGeometry) SetMaxIndex(v int) {
	cbg.Set("MaxIndex", v)
}
func (cbg *CylinderBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("addAttribute", name, attribute)}
}
func (cbg *CylinderBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return cbg.Call("addAttribute", name, array, itemSize)
}
func (cbg *CylinderBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	cbg.Call("addDrawCall", start, count, indexOffset)
}
func (cbg *CylinderBufferGeometry) AddEventListener(typ string, listener js.Value) {
	cbg.Call("addEventListener", typ, listener)
}
func (cbg *CylinderBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	cbg.Call("addGroup", start, count, materialIndex)
}
func (cbg *CylinderBufferGeometry) AddIndex(index js.Value) {
	cbg.Call("addIndex", index)
}
func (cbg *CylinderBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("applyMatrix", matrix)}
}
func (cbg *CylinderBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("center")}
}
func (cbg *CylinderBufferGeometry) ClearDrawCalls() {
	cbg.Call("clearDrawCalls")
}
func (cbg *CylinderBufferGeometry) ClearGroups() {
	cbg.Call("clearGroups")
}
func (cbg *CylinderBufferGeometry) Clone() *CylinderBufferGeometry {
	return &CylinderBufferGeometry{Value: cbg.Call("clone")}
}
func (cbg *CylinderBufferGeometry) ComputeBoundingBox() {
	cbg.Call("computeBoundingBox")
}
func (cbg *CylinderBufferGeometry) ComputeBoundingSphere() {
	cbg.Call("computeBoundingSphere")
}
func (cbg *CylinderBufferGeometry) ComputeVertexNormals() {
	cbg.Call("computeVertexNormals")
}
func (cbg *CylinderBufferGeometry) Copy(source *BufferGeometry) *CylinderBufferGeometry {
	return &CylinderBufferGeometry{Value: cbg.Call("copy", source)}
}
func (cbg *CylinderBufferGeometry) DispatchEvent(event js.Value) {
	cbg.Call("dispatchEvent", event)
}
func (cbg *CylinderBufferGeometry) Dispose() {
	cbg.Call("dispose")
}
func (cbg *CylinderBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("fromDirectGeometry", geometry)}
}
func (cbg *CylinderBufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (cbg *CylinderBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: cbg.Call("getAttribute", name)}
}
func (cbg *CylinderBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: cbg.Call("getIndex")}
}
func (cbg *CylinderBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cbg.Call("hasEventListener", typ, listener).Bool()
}
func (cbg *CylinderBufferGeometry) LookAt(v *Vector3) {
	cbg.Call("lookAt", v.JSValue())
}
func (cbg *CylinderBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("merge", geometry, offset)}
}
func (cbg *CylinderBufferGeometry) NormalizeNormals() {
	cbg.Call("normalizeNormals")
}
func (cbg *CylinderBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("removeAttribute", name)}
}
func (cbg *CylinderBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	cbg.Call("removeEventListener", typ, listener)
}
func (cbg *CylinderBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateX", angle)}
}
func (cbg *CylinderBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateY", angle)}
}
func (cbg *CylinderBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("rotateZ", angle)}
}
func (cbg *CylinderBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("scale", x, y, z)}
}
func (cbg *CylinderBufferGeometry) SetDrawRange2(start int, count int) {
	cbg.Call("setDrawRange", start, count)
}
func (cbg *CylinderBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("setFromObject", object)}
}
func (cbg *CylinderBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("setFromPoints", points)}
}
func (cbg *CylinderBufferGeometry) SetIndex2(index *BufferAttribute) {
	cbg.Call("setIndex", index)
}
func (cbg *CylinderBufferGeometry) ToJSON() js.Value {
	return cbg.Call("toJSON")
}
func (cbg *CylinderBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("toNonIndexed")}
}
func (cbg *CylinderBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: cbg.Call("translate", x, y, z)}
}
func (cbg *CylinderBufferGeometry) UpdateFromObject(object *Object3D) {
	cbg.Call("updateFromObject", object.JSValue())
}

// CylinderGeometry extend: [Geometry]
type CylinderGeometry struct {
	js.Value
}

func NewCylinderGeometry(radiusTop float64, radiusBottom float64, height float64, radiusSegments int, heightSegments int, openEnded bool, thetaStart float64, thetaLength float64) *CylinderGeometry {
	return &CylinderGeometry{Value: get("CylinderGeometry").New(radiusTop, radiusBottom, height, radiusSegments, heightSegments, openEnded, thetaStart, thetaLength)}
}
func (cg *CylinderGeometry) JSValue() js.Value {
	return cg.Value
}
func (cg *CylinderGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: cg.Get("animation")}
}
func (cg *CylinderGeometry) SetAnimation(v *AnimationClip) {
	cg.Set("animation", v.JSValue())
}
func (cg *CylinderGeometry) Animations() js.Value {
	return cg.Get("animations")
}
func (cg *CylinderGeometry) SetAnimations(v js.Value) {
	cg.Set("animations", v)
}
func (cg *CylinderGeometry) Bones() js.Value {
	return cg.Get("bones")
}
func (cg *CylinderGeometry) SetBones(v js.Value) {
	cg.Set("bones", v)
}
func (cg *CylinderGeometry) BoundingBox() *Box3 {
	return &Box3{Value: cg.Get("boundingBox")}
}
func (cg *CylinderGeometry) SetBoundingBox(v *Box3) {
	cg.Set("boundingBox", v.JSValue())
}
func (cg *CylinderGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: cg.Get("boundingSphere")}
}
func (cg *CylinderGeometry) SetBoundingSphere(v *Sphere) {
	cg.Set("boundingSphere", v.JSValue())
}
func (cg *CylinderGeometry) Colors() js.Value {
	return cg.Get("colors")
}
func (cg *CylinderGeometry) SetColors(v js.Value) {
	cg.Set("colors", v)
}
func (cg *CylinderGeometry) ColorsNeedUpdate() bool {
	return cg.Get("colorsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetColorsNeedUpdate(v bool) {
	cg.Set("colorsNeedUpdate", v)
}
func (cg *CylinderGeometry) ElementsNeedUpdate() bool {
	return cg.Get("elementsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetElementsNeedUpdate(v bool) {
	cg.Set("elementsNeedUpdate", v)
}
func (cg *CylinderGeometry) FaceVertexUvs() js.Value {
	return cg.Get("faceVertexUvs")
}
func (cg *CylinderGeometry) SetFaceVertexUvs(v js.Value) {
	cg.Set("faceVertexUvs", v)
}
func (cg *CylinderGeometry) Faces() js.Value {
	return cg.Get("faces")
}
func (cg *CylinderGeometry) SetFaces(v js.Value) {
	cg.Set("faces", v)
}
func (cg *CylinderGeometry) GroupsNeedUpdate() bool {
	return cg.Get("groupsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetGroupsNeedUpdate(v bool) {
	cg.Set("groupsNeedUpdate", v)
}
func (cg *CylinderGeometry) Id() int {
	return cg.Get("id").Int()
}
func (cg *CylinderGeometry) SetId(v int) {
	cg.Set("id", v)
}
func (cg *CylinderGeometry) LineDistances() js.Value {
	return cg.Get("lineDistances")
}
func (cg *CylinderGeometry) SetLineDistances(v js.Value) {
	cg.Set("lineDistances", v)
}
func (cg *CylinderGeometry) LineDistancesNeedUpdate() bool {
	return cg.Get("lineDistancesNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetLineDistancesNeedUpdate(v bool) {
	cg.Set("lineDistancesNeedUpdate", v)
}
func (cg *CylinderGeometry) MorphNormals() js.Value {
	return cg.Get("morphNormals")
}
func (cg *CylinderGeometry) SetMorphNormals(v js.Value) {
	cg.Set("morphNormals", v)
}
func (cg *CylinderGeometry) MorphTargets() js.Value {
	return cg.Get("morphTargets")
}
func (cg *CylinderGeometry) SetMorphTargets(v js.Value) {
	cg.Set("morphTargets", v)
}
func (cg *CylinderGeometry) Name() string {
	return cg.Get("name").String()
}
func (cg *CylinderGeometry) SetName(v string) {
	cg.Set("name", v)
}
func (cg *CylinderGeometry) NormalsNeedUpdate() bool {
	return cg.Get("normalsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetNormalsNeedUpdate(v bool) {
	cg.Set("normalsNeedUpdate", v)
}
func (cg *CylinderGeometry) Parameters() js.Value {
	return cg.Get("parameters")
}
func (cg *CylinderGeometry) SetParameters(v js.Value) {
	cg.Set("parameters", v)
}
func (cg *CylinderGeometry) SkinIndices() js.Value {
	return cg.Get("skinIndices")
}
func (cg *CylinderGeometry) SetSkinIndices(v js.Value) {
	cg.Set("skinIndices", v)
}
func (cg *CylinderGeometry) SkinWeights() js.Value {
	return cg.Get("skinWeights")
}
func (cg *CylinderGeometry) SetSkinWeights(v js.Value) {
	cg.Set("skinWeights", v)
}
func (cg *CylinderGeometry) Type() string {
	return cg.Get("type").String()
}
func (cg *CylinderGeometry) SetType(v string) {
	cg.Set("type", v)
}
func (cg *CylinderGeometry) Uuid() string {
	return cg.Get("uuid").String()
}
func (cg *CylinderGeometry) SetUuid(v string) {
	cg.Set("uuid", v)
}
func (cg *CylinderGeometry) UvsNeedUpdate() bool {
	return cg.Get("uvsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetUvsNeedUpdate(v bool) {
	cg.Set("uvsNeedUpdate", v)
}
func (cg *CylinderGeometry) Vertices() js.Value {
	return cg.Get("vertices")
}
func (cg *CylinderGeometry) SetVertices(v js.Value) {
	cg.Set("vertices", v)
}
func (cg *CylinderGeometry) VerticesNeedUpdate() bool {
	return cg.Get("verticesNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetVerticesNeedUpdate(v bool) {
	cg.Set("verticesNeedUpdate", v)
}
func (cg *CylinderGeometry) AddEventListener(typ string, listener js.Value) {
	cg.Call("addEventListener", typ, listener)
}
func (cg *CylinderGeometry) ApplyMatrix(matrix *Matrix4) Geometry {
	return &GeometryImpl{Value: cg.Call("applyMatrix", matrix)}
}
func (cg *CylinderGeometry) Center() Geometry {
	return &GeometryImpl{Value: cg.Call("center")}
}
func (cg *CylinderGeometry) Clone() Geometry {
	return &CylinderGeometry{Value: cg.Call("clone")}
}
func (cg *CylinderGeometry) ComputeBoundingBox() {
	cg.Call("computeBoundingBox")
}
func (cg *CylinderGeometry) ComputeBoundingSphere() {
	cg.Call("computeBoundingSphere")
}
func (cg *CylinderGeometry) ComputeFaceNormals() {
	cg.Call("computeFaceNormals")
}
func (cg *CylinderGeometry) ComputeFlatVertexNormals() {
	cg.Call("computeFlatVertexNormals")
}
func (cg *CylinderGeometry) ComputeMorphNormals() {
	cg.Call("computeMorphNormals")
}
func (cg *CylinderGeometry) ComputeVertexNormals(areaWeighted bool) {
	cg.Call("computeVertexNormals", areaWeighted)
}
func (cg *CylinderGeometry) Copy(source Geometry) Geometry {
	return &CylinderGeometry{Value: cg.Call("copy", source.JSValue())}
}
func (cg *CylinderGeometry) DispatchEvent(event js.Value) {
	cg.Call("dispatchEvent", event)
}
func (cg *CylinderGeometry) Dispose() {
	cg.Call("dispose")
}
func (cg *CylinderGeometry) FromBufferGeometry(geometry *BufferGeometry) Geometry {
	return &GeometryImpl{Value: cg.Call("fromBufferGeometry", geometry)}
}
func (cg *CylinderGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cg.Call("hasEventListener", typ, listener).Bool()
}
func (cg *CylinderGeometry) LookAt(vector *Vector3) {
	cg.Call("lookAt", vector.JSValue())
}
func (cg *CylinderGeometry) Merge(geometry Geometry, matrix Matrix, materialIndexOffset int) {
	cg.Call("merge", geometry.JSValue(), matrix, materialIndexOffset)
}
func (cg *CylinderGeometry) MergeMesh(mesh Mesh) {
	cg.Call("mergeMesh", mesh.JSValue())
}
func (cg *CylinderGeometry) MergeVertices() float64 {
	return cg.Call("mergeVertices").Float()
}
func (cg *CylinderGeometry) Normalize() Geometry {
	return &GeometryImpl{Value: cg.Call("normalize")}
}
func (cg *CylinderGeometry) RemoveEventListener(typ string, listener js.Value) {
	cg.Call("removeEventListener", typ, listener)
}
func (cg *CylinderGeometry) RotateX(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateX", angle)}
}
func (cg *CylinderGeometry) RotateY(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateY", angle)}
}
func (cg *CylinderGeometry) RotateZ(angle float64) Geometry {
	return &GeometryImpl{Value: cg.Call("rotateZ", angle)}
}
func (cg *CylinderGeometry) Scale(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: cg.Call("scale", x, y, z)}
}
func (cg *CylinderGeometry) SetFromPoints(points js.Value) Geometry {
	return &CylinderGeometry{Value: cg.Call("setFromPoints", points)}
}
func (cg *CylinderGeometry) SortFacesByMaterialIndex() {
	cg.Call("sortFacesByMaterialIndex")
}
func (cg *CylinderGeometry) ToJSON() js.Value {
	return cg.Call("toJSON")
}
func (cg *CylinderGeometry) Translate(x float64, y float64, z float64) Geometry {
	return &GeometryImpl{Value: cg.Call("translate", x, y, z)}
}

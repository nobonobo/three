// Code generated from "geometries/ExtrudeGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type ExtrudeGeometryOptions interface {
}
type UVGenerator interface {
	GenerateSideWallUV(geometry *ExtrudeBufferGeometry, vertices js.Value, indexA int, indexB int, indexC int, indexD int) js.Value
	GenerateTopUV(geometry *ExtrudeBufferGeometry, vertices js.Value, indexA int, indexB int, indexC int) js.Value
}
type ExtrudeBufferGeometry struct {
	js.Value
}

func NewExtrudeBufferGeometry(shapes *Shape, options ExtrudeGeometryOptions) *ExtrudeBufferGeometry {
	return &ExtrudeBufferGeometry{Value: get("ExtrudeBufferGeometry").New(shapes, options)}
}
func (ebg *ExtrudeBufferGeometry) Attributes() js.Value {
	return ebg.Get("attributes")
}
func (ebg *ExtrudeBufferGeometry) SetAttributes(v js.Value) {
	ebg.Set("attributes", v)
}
func (ebg *ExtrudeBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: ebg.Get("boundingBox")}
}
func (ebg *ExtrudeBufferGeometry) SetBoundingBox(v *Box3) {
	ebg.Set("boundingBox", v)
}
func (ebg *ExtrudeBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: ebg.Get("boundingSphere")}
}
func (ebg *ExtrudeBufferGeometry) SetBoundingSphere(v *Sphere) {
	ebg.Set("boundingSphere", v)
}
func (ebg *ExtrudeBufferGeometry) DrawRange() js.Value {
	return ebg.Get("drawRange")
}
func (ebg *ExtrudeBufferGeometry) SetDrawRange(v js.Value) {
	ebg.Set("drawRange", v)
}
func (ebg *ExtrudeBufferGeometry) Drawcalls() js.Value {
	return ebg.Get("drawcalls")
}
func (ebg *ExtrudeBufferGeometry) SetDrawcalls(v js.Value) {
	ebg.Set("drawcalls", v)
}
func (ebg *ExtrudeBufferGeometry) Groups() js.Value {
	return ebg.Get("groups")
}
func (ebg *ExtrudeBufferGeometry) SetGroups(v js.Value) {
	ebg.Set("groups", v)
}
func (ebg *ExtrudeBufferGeometry) Id() int {
	return ebg.Get("id").Int()
}
func (ebg *ExtrudeBufferGeometry) SetId(v int) {
	ebg.Set("id", v)
}
func (ebg *ExtrudeBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: ebg.Get("index")}
}
func (ebg *ExtrudeBufferGeometry) SetIndex(v *BufferAttribute) {
	ebg.Set("index", v)
}
func (ebg *ExtrudeBufferGeometry) MorphAttributes() js.Value {
	return ebg.Get("morphAttributes")
}
func (ebg *ExtrudeBufferGeometry) SetMorphAttributes(v js.Value) {
	ebg.Set("morphAttributes", v)
}
func (ebg *ExtrudeBufferGeometry) Name() string {
	return ebg.Get("name").String()
}
func (ebg *ExtrudeBufferGeometry) SetName(v string) {
	ebg.Set("name", v)
}
func (ebg *ExtrudeBufferGeometry) Offsets() js.Value {
	return ebg.Get("offsets")
}
func (ebg *ExtrudeBufferGeometry) SetOffsets(v js.Value) {
	ebg.Set("offsets", v)
}
func (ebg *ExtrudeBufferGeometry) Type() string {
	return ebg.Get("type").String()
}
func (ebg *ExtrudeBufferGeometry) SetType(v string) {
	ebg.Set("type", v)
}
func (ebg *ExtrudeBufferGeometry) UserData() js.Value {
	return ebg.Get("userData")
}
func (ebg *ExtrudeBufferGeometry) SetUserData(v js.Value) {
	ebg.Set("userData", v)
}
func (ebg *ExtrudeBufferGeometry) Uuid() string {
	return ebg.Get("uuid").String()
}
func (ebg *ExtrudeBufferGeometry) SetUuid(v string) {
	ebg.Set("uuid", v)
}
func (ebg *ExtrudeBufferGeometry) MaxIndex() int {
	return ebg.Get("MaxIndex").Int()
}
func (ebg *ExtrudeBufferGeometry) SetMaxIndex(v int) {
	ebg.Set("MaxIndex", v)
}
func (ebg *ExtrudeBufferGeometry) WorldUVGenerator() js.Value {
	return ebg.Get("WorldUVGenerator")
}
func (ebg *ExtrudeBufferGeometry) SetWorldUVGenerator(v js.Value) {
	ebg.Set("WorldUVGenerator", v)
}
func (ebg *ExtrudeBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("addAttribute", name, attribute)}
}
func (ebg *ExtrudeBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return ebg.Call("addAttribute", name, array, itemSize)
}
func (ebg *ExtrudeBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	ebg.Call("addDrawCall", start, count, indexOffset)
}
func (ebg *ExtrudeBufferGeometry) AddEventListener(typ string, listener js.Value) {
	ebg.Call("addEventListener", typ, listener)
}
func (ebg *ExtrudeBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	ebg.Call("addGroup", start, count, materialIndex)
}
func (ebg *ExtrudeBufferGeometry) AddIndex(index js.Value) {
	ebg.Call("addIndex", index)
}
func (ebg *ExtrudeBufferGeometry) AddShape(shape *Shape, options js.Value) {
	ebg.Call("addShape", shape, options)
}
func (ebg *ExtrudeBufferGeometry) AddShapeList(shapes js.Value, options js.Value) {
	ebg.Call("addShapeList", shapes, options)
}
func (ebg *ExtrudeBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("applyMatrix", matrix)}
}
func (ebg *ExtrudeBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("center")}
}
func (ebg *ExtrudeBufferGeometry) ClearDrawCalls() {
	ebg.Call("clearDrawCalls")
}
func (ebg *ExtrudeBufferGeometry) ClearGroups() {
	ebg.Call("clearGroups")
}
func (ebg *ExtrudeBufferGeometry) Clone() *ExtrudeBufferGeometry {
	return &ExtrudeBufferGeometry{Value: ebg.Call("clone")}
}
func (ebg *ExtrudeBufferGeometry) ComputeBoundingBox() {
	ebg.Call("computeBoundingBox")
}
func (ebg *ExtrudeBufferGeometry) ComputeBoundingSphere() {
	ebg.Call("computeBoundingSphere")
}
func (ebg *ExtrudeBufferGeometry) ComputeVertexNormals() {
	ebg.Call("computeVertexNormals")
}
func (ebg *ExtrudeBufferGeometry) Copy(source *BufferGeometry) *ExtrudeBufferGeometry {
	return &ExtrudeBufferGeometry{Value: ebg.Call("copy", source)}
}
func (ebg *ExtrudeBufferGeometry) DispatchEvent(event js.Value) {
	ebg.Call("dispatchEvent", event)
}
func (ebg *ExtrudeBufferGeometry) Dispose() {
	ebg.Call("dispose")
}
func (ebg *ExtrudeBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("fromDirectGeometry", geometry)}
}
func (ebg *ExtrudeBufferGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("fromGeometry", geometry, settings)}
}
func (ebg *ExtrudeBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: ebg.Call("getAttribute", name)}
}
func (ebg *ExtrudeBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: ebg.Call("getIndex")}
}
func (ebg *ExtrudeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ebg.Call("hasEventListener", typ, listener).Bool()
}
func (ebg *ExtrudeBufferGeometry) LookAt(v *Vector3) {
	ebg.Call("lookAt", v)
}
func (ebg *ExtrudeBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("merge", geometry, offset)}
}
func (ebg *ExtrudeBufferGeometry) NormalizeNormals() {
	ebg.Call("normalizeNormals")
}
func (ebg *ExtrudeBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("removeAttribute", name)}
}
func (ebg *ExtrudeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	ebg.Call("removeEventListener", typ, listener)
}
func (ebg *ExtrudeBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("rotateX", angle)}
}
func (ebg *ExtrudeBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("rotateY", angle)}
}
func (ebg *ExtrudeBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("rotateZ", angle)}
}
func (ebg *ExtrudeBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("scale", x, y, z)}
}
func (ebg *ExtrudeBufferGeometry) SetDrawRange2(start int, count int) {
	ebg.Call("setDrawRange", start, count)
}
func (ebg *ExtrudeBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("setFromObject", object)}
}
func (ebg *ExtrudeBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("setFromPoints", points)}
}
func (ebg *ExtrudeBufferGeometry) SetIndex2(index *BufferAttribute) {
	ebg.Call("setIndex", index)
}
func (ebg *ExtrudeBufferGeometry) ToJSON() js.Value {
	return ebg.Call("toJSON")
}
func (ebg *ExtrudeBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("toNonIndexed")}
}
func (ebg *ExtrudeBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: ebg.Call("translate", x, y, z)}
}
func (ebg *ExtrudeBufferGeometry) UpdateFromObject(object *Object3D) {
	ebg.Call("updateFromObject", object)
}

type ExtrudeGeometry struct {
	js.Value
}

func NewExtrudeGeometry(shapes *Shape, options ExtrudeGeometryOptions) *ExtrudeGeometry {
	return &ExtrudeGeometry{Value: get("ExtrudeGeometry").New(shapes, options)}
}
func (eg *ExtrudeGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: eg.Get("animation")}
}
func (eg *ExtrudeGeometry) SetAnimation(v *AnimationClip) {
	eg.Set("animation", v)
}
func (eg *ExtrudeGeometry) Animations() js.Value {
	return eg.Get("animations")
}
func (eg *ExtrudeGeometry) SetAnimations(v js.Value) {
	eg.Set("animations", v)
}
func (eg *ExtrudeGeometry) Bones() js.Value {
	return eg.Get("bones")
}
func (eg *ExtrudeGeometry) SetBones(v js.Value) {
	eg.Set("bones", v)
}
func (eg *ExtrudeGeometry) BoundingBox() *Box3 {
	return &Box3{Value: eg.Get("boundingBox")}
}
func (eg *ExtrudeGeometry) SetBoundingBox(v *Box3) {
	eg.Set("boundingBox", v)
}
func (eg *ExtrudeGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: eg.Get("boundingSphere")}
}
func (eg *ExtrudeGeometry) SetBoundingSphere(v *Sphere) {
	eg.Set("boundingSphere", v)
}
func (eg *ExtrudeGeometry) Colors() js.Value {
	return eg.Get("colors")
}
func (eg *ExtrudeGeometry) SetColors(v js.Value) {
	eg.Set("colors", v)
}
func (eg *ExtrudeGeometry) ColorsNeedUpdate() bool {
	return eg.Get("colorsNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetColorsNeedUpdate(v bool) {
	eg.Set("colorsNeedUpdate", v)
}
func (eg *ExtrudeGeometry) ElementsNeedUpdate() bool {
	return eg.Get("elementsNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetElementsNeedUpdate(v bool) {
	eg.Set("elementsNeedUpdate", v)
}
func (eg *ExtrudeGeometry) FaceVertexUvs() js.Value {
	return eg.Get("faceVertexUvs")
}
func (eg *ExtrudeGeometry) SetFaceVertexUvs(v js.Value) {
	eg.Set("faceVertexUvs", v)
}
func (eg *ExtrudeGeometry) Faces() js.Value {
	return eg.Get("faces")
}
func (eg *ExtrudeGeometry) SetFaces(v js.Value) {
	eg.Set("faces", v)
}
func (eg *ExtrudeGeometry) GroupsNeedUpdate() bool {
	return eg.Get("groupsNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetGroupsNeedUpdate(v bool) {
	eg.Set("groupsNeedUpdate", v)
}
func (eg *ExtrudeGeometry) Id() int {
	return eg.Get("id").Int()
}
func (eg *ExtrudeGeometry) SetId(v int) {
	eg.Set("id", v)
}
func (eg *ExtrudeGeometry) LineDistances() js.Value {
	return eg.Get("lineDistances")
}
func (eg *ExtrudeGeometry) SetLineDistances(v js.Value) {
	eg.Set("lineDistances", v)
}
func (eg *ExtrudeGeometry) LineDistancesNeedUpdate() bool {
	return eg.Get("lineDistancesNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetLineDistancesNeedUpdate(v bool) {
	eg.Set("lineDistancesNeedUpdate", v)
}
func (eg *ExtrudeGeometry) MorphNormals() js.Value {
	return eg.Get("morphNormals")
}
func (eg *ExtrudeGeometry) SetMorphNormals(v js.Value) {
	eg.Set("morphNormals", v)
}
func (eg *ExtrudeGeometry) MorphTargets() js.Value {
	return eg.Get("morphTargets")
}
func (eg *ExtrudeGeometry) SetMorphTargets(v js.Value) {
	eg.Set("morphTargets", v)
}
func (eg *ExtrudeGeometry) Name() string {
	return eg.Get("name").String()
}
func (eg *ExtrudeGeometry) SetName(v string) {
	eg.Set("name", v)
}
func (eg *ExtrudeGeometry) NormalsNeedUpdate() bool {
	return eg.Get("normalsNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetNormalsNeedUpdate(v bool) {
	eg.Set("normalsNeedUpdate", v)
}
func (eg *ExtrudeGeometry) SkinIndices() js.Value {
	return eg.Get("skinIndices")
}
func (eg *ExtrudeGeometry) SetSkinIndices(v js.Value) {
	eg.Set("skinIndices", v)
}
func (eg *ExtrudeGeometry) SkinWeights() js.Value {
	return eg.Get("skinWeights")
}
func (eg *ExtrudeGeometry) SetSkinWeights(v js.Value) {
	eg.Set("skinWeights", v)
}
func (eg *ExtrudeGeometry) Type() string {
	return eg.Get("type").String()
}
func (eg *ExtrudeGeometry) SetType(v string) {
	eg.Set("type", v)
}
func (eg *ExtrudeGeometry) Uuid() string {
	return eg.Get("uuid").String()
}
func (eg *ExtrudeGeometry) SetUuid(v string) {
	eg.Set("uuid", v)
}
func (eg *ExtrudeGeometry) UvsNeedUpdate() bool {
	return eg.Get("uvsNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetUvsNeedUpdate(v bool) {
	eg.Set("uvsNeedUpdate", v)
}
func (eg *ExtrudeGeometry) Vertices() js.Value {
	return eg.Get("vertices")
}
func (eg *ExtrudeGeometry) SetVertices(v js.Value) {
	eg.Set("vertices", v)
}
func (eg *ExtrudeGeometry) VerticesNeedUpdate() bool {
	return eg.Get("verticesNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetVerticesNeedUpdate(v bool) {
	eg.Set("verticesNeedUpdate", v)
}
func (eg *ExtrudeGeometry) WorldUVGenerator() js.Value {
	return eg.Get("WorldUVGenerator")
}
func (eg *ExtrudeGeometry) SetWorldUVGenerator(v js.Value) {
	eg.Set("WorldUVGenerator", v)
}
func (eg *ExtrudeGeometry) AddEventListener(typ string, listener js.Value) {
	eg.Call("addEventListener", typ, listener)
}
func (eg *ExtrudeGeometry) AddShape(shape *Shape, options js.Value) {
	eg.Call("addShape", shape, options)
}
func (eg *ExtrudeGeometry) AddShapeList(shapes js.Value, options js.Value) {
	eg.Call("addShapeList", shapes, options)
}
func (eg *ExtrudeGeometry) ApplyMatrix(matrix *Matrix4) *Geometry {
	return &Geometry{Value: eg.Call("applyMatrix", matrix)}
}
func (eg *ExtrudeGeometry) Center() *Geometry {
	return &Geometry{Value: eg.Call("center")}
}
func (eg *ExtrudeGeometry) Clone() *ExtrudeGeometry {
	return &ExtrudeGeometry{Value: eg.Call("clone")}
}
func (eg *ExtrudeGeometry) ComputeBoundingBox() {
	eg.Call("computeBoundingBox")
}
func (eg *ExtrudeGeometry) ComputeBoundingSphere() {
	eg.Call("computeBoundingSphere")
}
func (eg *ExtrudeGeometry) ComputeFaceNormals() {
	eg.Call("computeFaceNormals")
}
func (eg *ExtrudeGeometry) ComputeFlatVertexNormals() {
	eg.Call("computeFlatVertexNormals")
}
func (eg *ExtrudeGeometry) ComputeMorphNormals() {
	eg.Call("computeMorphNormals")
}
func (eg *ExtrudeGeometry) ComputeVertexNormals(areaWeighted bool) {
	eg.Call("computeVertexNormals", areaWeighted)
}
func (eg *ExtrudeGeometry) Copy(source *Geometry) *ExtrudeGeometry {
	return &ExtrudeGeometry{Value: eg.Call("copy", source)}
}
func (eg *ExtrudeGeometry) DispatchEvent(event js.Value) {
	eg.Call("dispatchEvent", event)
}
func (eg *ExtrudeGeometry) Dispose() {
	eg.Call("dispose")
}
func (eg *ExtrudeGeometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: eg.Call("fromBufferGeometry", geometry)}
}
func (eg *ExtrudeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return eg.Call("hasEventListener", typ, listener).Bool()
}
func (eg *ExtrudeGeometry) LookAt(vector *Vector3) {
	eg.Call("lookAt", vector)
}
func (eg *ExtrudeGeometry) Merge(geometry *Geometry, matrix Matrix, materialIndexOffset int) {
	eg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (eg *ExtrudeGeometry) MergeMesh(mesh *Mesh) {
	eg.Call("mergeMesh", mesh)
}
func (eg *ExtrudeGeometry) MergeVertices() float64 {
	return eg.Call("mergeVertices").Float()
}
func (eg *ExtrudeGeometry) Normalize() *Geometry {
	return &Geometry{Value: eg.Call("normalize")}
}
func (eg *ExtrudeGeometry) RemoveEventListener(typ string, listener js.Value) {
	eg.Call("removeEventListener", typ, listener)
}
func (eg *ExtrudeGeometry) RotateX(angle float64) *Geometry {
	return &Geometry{Value: eg.Call("rotateX", angle)}
}
func (eg *ExtrudeGeometry) RotateY(angle float64) *Geometry {
	return &Geometry{Value: eg.Call("rotateY", angle)}
}
func (eg *ExtrudeGeometry) RotateZ(angle float64) *Geometry {
	return &Geometry{Value: eg.Call("rotateZ", angle)}
}
func (eg *ExtrudeGeometry) Scale(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: eg.Call("scale", x, y, z)}
}
func (eg *ExtrudeGeometry) SetFromPoints(points js.Value) *ExtrudeGeometry {
	return &ExtrudeGeometry{Value: eg.Call("setFromPoints", points)}
}
func (eg *ExtrudeGeometry) SortFacesByMaterialIndex() {
	eg.Call("sortFacesByMaterialIndex")
}
func (eg *ExtrudeGeometry) ToJSON() js.Value {
	return eg.Call("toJSON")
}
func (eg *ExtrudeGeometry) Translate(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: eg.Call("translate", x, y, z)}
}

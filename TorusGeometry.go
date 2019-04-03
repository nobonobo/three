// Code generated from "geometries/TorusGeometry.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type TorusBufferGeometry struct {
	js.Value
}

func NewTorusBufferGeometry(radius float64, tube float64, radialSegments float64, tubularSegments float64, arc float64) *TorusBufferGeometry {
	return &TorusBufferGeometry{Value: get("TorusBufferGeometry").New(radius, tube, radialSegments, tubularSegments, arc)}
}
func (tbg *TorusBufferGeometry) Attributes() js.Value {
	return tbg.Get("attributes")
}
func (tbg *TorusBufferGeometry) SetAttributes(v js.Value) {
	tbg.Set("attributes", v)
}
func (tbg *TorusBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tbg.Get("boundingBox")}
}
func (tbg *TorusBufferGeometry) SetBoundingBox(v *Box3) {
	tbg.Set("boundingBox", v)
}
func (tbg *TorusBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tbg.Get("boundingSphere")}
}
func (tbg *TorusBufferGeometry) SetBoundingSphere(v *Sphere) {
	tbg.Set("boundingSphere", v)
}
func (tbg *TorusBufferGeometry) DrawRange() js.Value {
	return tbg.Get("drawRange")
}
func (tbg *TorusBufferGeometry) SetDrawRange(v js.Value) {
	tbg.Set("drawRange", v)
}
func (tbg *TorusBufferGeometry) Drawcalls() js.Value {
	return tbg.Get("drawcalls")
}
func (tbg *TorusBufferGeometry) SetDrawcalls(v js.Value) {
	tbg.Set("drawcalls", v)
}
func (tbg *TorusBufferGeometry) Groups() js.Value {
	return tbg.Get("groups")
}
func (tbg *TorusBufferGeometry) SetGroups(v js.Value) {
	tbg.Set("groups", v)
}
func (tbg *TorusBufferGeometry) Id() int {
	return tbg.Get("id").Int()
}
func (tbg *TorusBufferGeometry) SetId(v int) {
	tbg.Set("id", v)
}
func (tbg *TorusBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Get("index")}
}
func (tbg *TorusBufferGeometry) SetIndex(v *BufferAttribute) {
	tbg.Set("index", v)
}
func (tbg *TorusBufferGeometry) MorphAttributes() js.Value {
	return tbg.Get("morphAttributes")
}
func (tbg *TorusBufferGeometry) SetMorphAttributes(v js.Value) {
	tbg.Set("morphAttributes", v)
}
func (tbg *TorusBufferGeometry) Name() string {
	return tbg.Get("name").String()
}
func (tbg *TorusBufferGeometry) SetName(v string) {
	tbg.Set("name", v)
}
func (tbg *TorusBufferGeometry) Offsets() js.Value {
	return tbg.Get("offsets")
}
func (tbg *TorusBufferGeometry) SetOffsets(v js.Value) {
	tbg.Set("offsets", v)
}
func (tbg *TorusBufferGeometry) Parameters() js.Value {
	return tbg.Get("parameters")
}
func (tbg *TorusBufferGeometry) SetParameters(v js.Value) {
	tbg.Set("parameters", v)
}
func (tbg *TorusBufferGeometry) Type() string {
	return tbg.Get("type").String()
}
func (tbg *TorusBufferGeometry) SetType(v string) {
	tbg.Set("type", v)
}
func (tbg *TorusBufferGeometry) UserData() js.Value {
	return tbg.Get("userData")
}
func (tbg *TorusBufferGeometry) SetUserData(v js.Value) {
	tbg.Set("userData", v)
}
func (tbg *TorusBufferGeometry) Uuid() string {
	return tbg.Get("uuid").String()
}
func (tbg *TorusBufferGeometry) SetUuid(v string) {
	tbg.Set("uuid", v)
}
func (tbg *TorusBufferGeometry) MaxIndex() int {
	return tbg.Get("MaxIndex").Int()
}
func (tbg *TorusBufferGeometry) SetMaxIndex(v int) {
	tbg.Set("MaxIndex", v)
}
func (tbg *TorusBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("addAttribute", name, attribute)}
}
func (tbg *TorusBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tbg.Call("addAttribute", name, array, itemSize)
}
func (tbg *TorusBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tbg.Call("addDrawCall", start, count, indexOffset)
}
func (tbg *TorusBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tbg.Call("addEventListener", typ, listener)
}
func (tbg *TorusBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tbg.Call("addGroup", start, count, materialIndex)
}
func (tbg *TorusBufferGeometry) AddIndex(index js.Value) {
	tbg.Call("addIndex", index)
}
func (tbg *TorusBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("applyMatrix", matrix)}
}
func (tbg *TorusBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("center")}
}
func (tbg *TorusBufferGeometry) ClearDrawCalls() {
	tbg.Call("clearDrawCalls")
}
func (tbg *TorusBufferGeometry) ClearGroups() {
	tbg.Call("clearGroups")
}
func (tbg *TorusBufferGeometry) Clone() *TorusBufferGeometry {
	return &TorusBufferGeometry{Value: tbg.Call("clone")}
}
func (tbg *TorusBufferGeometry) ComputeBoundingBox() {
	tbg.Call("computeBoundingBox")
}
func (tbg *TorusBufferGeometry) ComputeBoundingSphere() {
	tbg.Call("computeBoundingSphere")
}
func (tbg *TorusBufferGeometry) ComputeVertexNormals() {
	tbg.Call("computeVertexNormals")
}
func (tbg *TorusBufferGeometry) Copy(source *BufferGeometry) *TorusBufferGeometry {
	return &TorusBufferGeometry{Value: tbg.Call("copy", source)}
}
func (tbg *TorusBufferGeometry) DispatchEvent(event js.Value) {
	tbg.Call("dispatchEvent", event)
}
func (tbg *TorusBufferGeometry) Dispose() {
	tbg.Call("dispose")
}
func (tbg *TorusBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("fromDirectGeometry", geometry)}
}
func (tbg *TorusBufferGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("fromGeometry", geometry, settings)}
}
func (tbg *TorusBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getAttribute", name)}
}
func (tbg *TorusBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getIndex")}
}
func (tbg *TorusBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tbg.Call("hasEventListener", typ, listener).Bool()
}
func (tbg *TorusBufferGeometry) LookAt(v *Vector3) {
	tbg.Call("lookAt", v)
}
func (tbg *TorusBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("merge", geometry, offset)}
}
func (tbg *TorusBufferGeometry) NormalizeNormals() {
	tbg.Call("normalizeNormals")
}
func (tbg *TorusBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("removeAttribute", name)}
}
func (tbg *TorusBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tbg.Call("removeEventListener", typ, listener)
}
func (tbg *TorusBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateX", angle)}
}
func (tbg *TorusBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateY", angle)}
}
func (tbg *TorusBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateZ", angle)}
}
func (tbg *TorusBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("scale", x, y, z)}
}
func (tbg *TorusBufferGeometry) SetDrawRange2(start int, count int) {
	tbg.Call("setDrawRange", start, count)
}
func (tbg *TorusBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("setFromObject", object)}
}
func (tbg *TorusBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("setFromPoints", points)}
}
func (tbg *TorusBufferGeometry) SetIndex2(index *BufferAttribute) {
	tbg.Call("setIndex", index)
}
func (tbg *TorusBufferGeometry) ToJSON() js.Value {
	return tbg.Call("toJSON")
}
func (tbg *TorusBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("toNonIndexed")}
}
func (tbg *TorusBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("translate", x, y, z)}
}
func (tbg *TorusBufferGeometry) UpdateFromObject(object *Object3D) {
	tbg.Call("updateFromObject", object)
}

type TorusGeometry struct {
	js.Value
}

func NewTorusGeometry(radius float64, tube float64, radialSegments float64, tubularSegments float64, arc float64) *TorusGeometry {
	return &TorusGeometry{Value: get("TorusGeometry").New(radius, tube, radialSegments, tubularSegments, arc)}
}
func (tg *TorusGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: tg.Get("animation")}
}
func (tg *TorusGeometry) SetAnimation(v *AnimationClip) {
	tg.Set("animation", v)
}
func (tg *TorusGeometry) Animations() js.Value {
	return tg.Get("animations")
}
func (tg *TorusGeometry) SetAnimations(v js.Value) {
	tg.Set("animations", v)
}
func (tg *TorusGeometry) Bones() js.Value {
	return tg.Get("bones")
}
func (tg *TorusGeometry) SetBones(v js.Value) {
	tg.Set("bones", v)
}
func (tg *TorusGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tg.Get("boundingBox")}
}
func (tg *TorusGeometry) SetBoundingBox(v *Box3) {
	tg.Set("boundingBox", v)
}
func (tg *TorusGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tg.Get("boundingSphere")}
}
func (tg *TorusGeometry) SetBoundingSphere(v *Sphere) {
	tg.Set("boundingSphere", v)
}
func (tg *TorusGeometry) Colors() js.Value {
	return tg.Get("colors")
}
func (tg *TorusGeometry) SetColors(v js.Value) {
	tg.Set("colors", v)
}
func (tg *TorusGeometry) ColorsNeedUpdate() bool {
	return tg.Get("colorsNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetColorsNeedUpdate(v bool) {
	tg.Set("colorsNeedUpdate", v)
}
func (tg *TorusGeometry) ElementsNeedUpdate() bool {
	return tg.Get("elementsNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetElementsNeedUpdate(v bool) {
	tg.Set("elementsNeedUpdate", v)
}
func (tg *TorusGeometry) FaceVertexUvs() js.Value {
	return tg.Get("faceVertexUvs")
}
func (tg *TorusGeometry) SetFaceVertexUvs(v js.Value) {
	tg.Set("faceVertexUvs", v)
}
func (tg *TorusGeometry) Faces() js.Value {
	return tg.Get("faces")
}
func (tg *TorusGeometry) SetFaces(v js.Value) {
	tg.Set("faces", v)
}
func (tg *TorusGeometry) GroupsNeedUpdate() bool {
	return tg.Get("groupsNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetGroupsNeedUpdate(v bool) {
	tg.Set("groupsNeedUpdate", v)
}
func (tg *TorusGeometry) Id() int {
	return tg.Get("id").Int()
}
func (tg *TorusGeometry) SetId(v int) {
	tg.Set("id", v)
}
func (tg *TorusGeometry) LineDistances() js.Value {
	return tg.Get("lineDistances")
}
func (tg *TorusGeometry) SetLineDistances(v js.Value) {
	tg.Set("lineDistances", v)
}
func (tg *TorusGeometry) LineDistancesNeedUpdate() bool {
	return tg.Get("lineDistancesNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetLineDistancesNeedUpdate(v bool) {
	tg.Set("lineDistancesNeedUpdate", v)
}
func (tg *TorusGeometry) MorphNormals() js.Value {
	return tg.Get("morphNormals")
}
func (tg *TorusGeometry) SetMorphNormals(v js.Value) {
	tg.Set("morphNormals", v)
}
func (tg *TorusGeometry) MorphTargets() js.Value {
	return tg.Get("morphTargets")
}
func (tg *TorusGeometry) SetMorphTargets(v js.Value) {
	tg.Set("morphTargets", v)
}
func (tg *TorusGeometry) Name() string {
	return tg.Get("name").String()
}
func (tg *TorusGeometry) SetName(v string) {
	tg.Set("name", v)
}
func (tg *TorusGeometry) NormalsNeedUpdate() bool {
	return tg.Get("normalsNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetNormalsNeedUpdate(v bool) {
	tg.Set("normalsNeedUpdate", v)
}
func (tg *TorusGeometry) Parameters() js.Value {
	return tg.Get("parameters")
}
func (tg *TorusGeometry) SetParameters(v js.Value) {
	tg.Set("parameters", v)
}
func (tg *TorusGeometry) SkinIndices() js.Value {
	return tg.Get("skinIndices")
}
func (tg *TorusGeometry) SetSkinIndices(v js.Value) {
	tg.Set("skinIndices", v)
}
func (tg *TorusGeometry) SkinWeights() js.Value {
	return tg.Get("skinWeights")
}
func (tg *TorusGeometry) SetSkinWeights(v js.Value) {
	tg.Set("skinWeights", v)
}
func (tg *TorusGeometry) Type() string {
	return tg.Get("type").String()
}
func (tg *TorusGeometry) SetType(v string) {
	tg.Set("type", v)
}
func (tg *TorusGeometry) Uuid() string {
	return tg.Get("uuid").String()
}
func (tg *TorusGeometry) SetUuid(v string) {
	tg.Set("uuid", v)
}
func (tg *TorusGeometry) UvsNeedUpdate() bool {
	return tg.Get("uvsNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetUvsNeedUpdate(v bool) {
	tg.Set("uvsNeedUpdate", v)
}
func (tg *TorusGeometry) Vertices() js.Value {
	return tg.Get("vertices")
}
func (tg *TorusGeometry) SetVertices(v js.Value) {
	tg.Set("vertices", v)
}
func (tg *TorusGeometry) VerticesNeedUpdate() bool {
	return tg.Get("verticesNeedUpdate").Bool()
}
func (tg *TorusGeometry) SetVerticesNeedUpdate(v bool) {
	tg.Set("verticesNeedUpdate", v)
}
func (tg *TorusGeometry) AddEventListener(typ string, listener js.Value) {
	tg.Call("addEventListener", typ, listener)
}
func (tg *TorusGeometry) ApplyMatrix(matrix *Matrix4) *Geometry {
	return &Geometry{Value: tg.Call("applyMatrix", matrix)}
}
func (tg *TorusGeometry) Center() *Geometry {
	return &Geometry{Value: tg.Call("center")}
}
func (tg *TorusGeometry) Clone() *TorusGeometry {
	return &TorusGeometry{Value: tg.Call("clone")}
}
func (tg *TorusGeometry) ComputeBoundingBox() {
	tg.Call("computeBoundingBox")
}
func (tg *TorusGeometry) ComputeBoundingSphere() {
	tg.Call("computeBoundingSphere")
}
func (tg *TorusGeometry) ComputeFaceNormals() {
	tg.Call("computeFaceNormals")
}
func (tg *TorusGeometry) ComputeFlatVertexNormals() {
	tg.Call("computeFlatVertexNormals")
}
func (tg *TorusGeometry) ComputeMorphNormals() {
	tg.Call("computeMorphNormals")
}
func (tg *TorusGeometry) ComputeVertexNormals(areaWeighted bool) {
	tg.Call("computeVertexNormals", areaWeighted)
}
func (tg *TorusGeometry) Copy(source *Geometry) *TorusGeometry {
	return &TorusGeometry{Value: tg.Call("copy", source)}
}
func (tg *TorusGeometry) DispatchEvent(event js.Value) {
	tg.Call("dispatchEvent", event)
}
func (tg *TorusGeometry) Dispose() {
	tg.Call("dispose")
}
func (tg *TorusGeometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: tg.Call("fromBufferGeometry", geometry)}
}
func (tg *TorusGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tg.Call("hasEventListener", typ, listener).Bool()
}
func (tg *TorusGeometry) LookAt(vector *Vector3) {
	tg.Call("lookAt", vector)
}
func (tg *TorusGeometry) Merge(geometry *Geometry, matrix Matrix, materialIndexOffset int) {
	tg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (tg *TorusGeometry) MergeMesh(mesh *Mesh) {
	tg.Call("mergeMesh", mesh)
}
func (tg *TorusGeometry) MergeVertices() float64 {
	return tg.Call("mergeVertices").Float()
}
func (tg *TorusGeometry) Normalize() *Geometry {
	return &Geometry{Value: tg.Call("normalize")}
}
func (tg *TorusGeometry) RemoveEventListener(typ string, listener js.Value) {
	tg.Call("removeEventListener", typ, listener)
}
func (tg *TorusGeometry) RotateX(angle float64) *Geometry {
	return &Geometry{Value: tg.Call("rotateX", angle)}
}
func (tg *TorusGeometry) RotateY(angle float64) *Geometry {
	return &Geometry{Value: tg.Call("rotateY", angle)}
}
func (tg *TorusGeometry) RotateZ(angle float64) *Geometry {
	return &Geometry{Value: tg.Call("rotateZ", angle)}
}
func (tg *TorusGeometry) Scale(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: tg.Call("scale", x, y, z)}
}
func (tg *TorusGeometry) SetFromPoints(points js.Value) *TorusGeometry {
	return &TorusGeometry{Value: tg.Call("setFromPoints", points)}
}
func (tg *TorusGeometry) SortFacesByMaterialIndex() {
	tg.Call("sortFacesByMaterialIndex")
}
func (tg *TorusGeometry) ToJSON() js.Value {
	return tg.Call("toJSON")
}
func (tg *TorusGeometry) Translate(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: tg.Call("translate", x, y, z)}
}

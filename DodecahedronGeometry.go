// Code generated from "geometries/DodecahedronGeometry.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type DodecahedronBufferGeometry struct {
	js.Value
}

func NewDodecahedronBufferGeometry(radius float64, detail float64) *DodecahedronBufferGeometry {
	return &DodecahedronBufferGeometry{Value: get("DodecahedronBufferGeometry").New(radius, detail)}
}
func (dbg *DodecahedronBufferGeometry) Attributes() js.Value {
	return dbg.Get("attributes")
}
func (dbg *DodecahedronBufferGeometry) SetAttributes(v js.Value) {
	dbg.Set("attributes", v)
}
func (dbg *DodecahedronBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: dbg.Get("boundingBox")}
}
func (dbg *DodecahedronBufferGeometry) SetBoundingBox(v *Box3) {
	dbg.Set("boundingBox", v)
}
func (dbg *DodecahedronBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: dbg.Get("boundingSphere")}
}
func (dbg *DodecahedronBufferGeometry) SetBoundingSphere(v *Sphere) {
	dbg.Set("boundingSphere", v)
}
func (dbg *DodecahedronBufferGeometry) DrawRange() js.Value {
	return dbg.Get("drawRange")
}
func (dbg *DodecahedronBufferGeometry) SetDrawRange(v js.Value) {
	dbg.Set("drawRange", v)
}
func (dbg *DodecahedronBufferGeometry) Drawcalls() js.Value {
	return dbg.Get("drawcalls")
}
func (dbg *DodecahedronBufferGeometry) SetDrawcalls(v js.Value) {
	dbg.Set("drawcalls", v)
}
func (dbg *DodecahedronBufferGeometry) Groups() js.Value {
	return dbg.Get("groups")
}
func (dbg *DodecahedronBufferGeometry) SetGroups(v js.Value) {
	dbg.Set("groups", v)
}
func (dbg *DodecahedronBufferGeometry) Id() int {
	return dbg.Get("id").Int()
}
func (dbg *DodecahedronBufferGeometry) SetId(v int) {
	dbg.Set("id", v)
}
func (dbg *DodecahedronBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: dbg.Get("index")}
}
func (dbg *DodecahedronBufferGeometry) SetIndex(v *BufferAttribute) {
	dbg.Set("index", v)
}
func (dbg *DodecahedronBufferGeometry) MorphAttributes() js.Value {
	return dbg.Get("morphAttributes")
}
func (dbg *DodecahedronBufferGeometry) SetMorphAttributes(v js.Value) {
	dbg.Set("morphAttributes", v)
}
func (dbg *DodecahedronBufferGeometry) Name() string {
	return dbg.Get("name").String()
}
func (dbg *DodecahedronBufferGeometry) SetName(v string) {
	dbg.Set("name", v)
}
func (dbg *DodecahedronBufferGeometry) Offsets() js.Value {
	return dbg.Get("offsets")
}
func (dbg *DodecahedronBufferGeometry) SetOffsets(v js.Value) {
	dbg.Set("offsets", v)
}
func (dbg *DodecahedronBufferGeometry) Parameters() js.Value {
	return dbg.Get("parameters")
}
func (dbg *DodecahedronBufferGeometry) SetParameters(v js.Value) {
	dbg.Set("parameters", v)
}
func (dbg *DodecahedronBufferGeometry) Type() string {
	return dbg.Get("type").String()
}
func (dbg *DodecahedronBufferGeometry) SetType(v string) {
	dbg.Set("type", v)
}
func (dbg *DodecahedronBufferGeometry) UserData() js.Value {
	return dbg.Get("userData")
}
func (dbg *DodecahedronBufferGeometry) SetUserData(v js.Value) {
	dbg.Set("userData", v)
}
func (dbg *DodecahedronBufferGeometry) Uuid() string {
	return dbg.Get("uuid").String()
}
func (dbg *DodecahedronBufferGeometry) SetUuid(v string) {
	dbg.Set("uuid", v)
}
func (dbg *DodecahedronBufferGeometry) MaxIndex() int {
	return dbg.Get("MaxIndex").Int()
}
func (dbg *DodecahedronBufferGeometry) SetMaxIndex(v int) {
	dbg.Set("MaxIndex", v)
}
func (dbg *DodecahedronBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("addAttribute", name, attribute)}
}
func (dbg *DodecahedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return dbg.Call("addAttribute", name, array, itemSize)
}
func (dbg *DodecahedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	dbg.Call("addDrawCall", start, count, indexOffset)
}
func (dbg *DodecahedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	dbg.Call("addEventListener", typ, listener)
}
func (dbg *DodecahedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	dbg.Call("addGroup", start, count, materialIndex)
}
func (dbg *DodecahedronBufferGeometry) AddIndex(index js.Value) {
	dbg.Call("addIndex", index)
}
func (dbg *DodecahedronBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("applyMatrix", matrix)}
}
func (dbg *DodecahedronBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("center")}
}
func (dbg *DodecahedronBufferGeometry) ClearDrawCalls() {
	dbg.Call("clearDrawCalls")
}
func (dbg *DodecahedronBufferGeometry) ClearGroups() {
	dbg.Call("clearGroups")
}
func (dbg *DodecahedronBufferGeometry) Clone() *DodecahedronBufferGeometry {
	return &DodecahedronBufferGeometry{Value: dbg.Call("clone")}
}
func (dbg *DodecahedronBufferGeometry) ComputeBoundingBox() {
	dbg.Call("computeBoundingBox")
}
func (dbg *DodecahedronBufferGeometry) ComputeBoundingSphere() {
	dbg.Call("computeBoundingSphere")
}
func (dbg *DodecahedronBufferGeometry) ComputeVertexNormals() {
	dbg.Call("computeVertexNormals")
}
func (dbg *DodecahedronBufferGeometry) Copy(source *BufferGeometry) *DodecahedronBufferGeometry {
	return &DodecahedronBufferGeometry{Value: dbg.Call("copy", source)}
}
func (dbg *DodecahedronBufferGeometry) DispatchEvent(event js.Value) {
	dbg.Call("dispatchEvent", event)
}
func (dbg *DodecahedronBufferGeometry) Dispose() {
	dbg.Call("dispose")
}
func (dbg *DodecahedronBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("fromDirectGeometry", geometry)}
}
func (dbg *DodecahedronBufferGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("fromGeometry", geometry, settings)}
}
func (dbg *DodecahedronBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: dbg.Call("getAttribute", name)}
}
func (dbg *DodecahedronBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: dbg.Call("getIndex")}
}
func (dbg *DodecahedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return dbg.Call("hasEventListener", typ, listener).Bool()
}
func (dbg *DodecahedronBufferGeometry) LookAt(v *Vector3) {
	dbg.Call("lookAt", v)
}
func (dbg *DodecahedronBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("merge", geometry, offset)}
}
func (dbg *DodecahedronBufferGeometry) NormalizeNormals() {
	dbg.Call("normalizeNormals")
}
func (dbg *DodecahedronBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("removeAttribute", name)}
}
func (dbg *DodecahedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	dbg.Call("removeEventListener", typ, listener)
}
func (dbg *DodecahedronBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("rotateX", angle)}
}
func (dbg *DodecahedronBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("rotateY", angle)}
}
func (dbg *DodecahedronBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("rotateZ", angle)}
}
func (dbg *DodecahedronBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("scale", x, y, z)}
}
func (dbg *DodecahedronBufferGeometry) SetDrawRange2(start int, count int) {
	dbg.Call("setDrawRange", start, count)
}
func (dbg *DodecahedronBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("setFromObject", object)}
}
func (dbg *DodecahedronBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("setFromPoints", points)}
}
func (dbg *DodecahedronBufferGeometry) SetIndex2(index *BufferAttribute) {
	dbg.Call("setIndex", index)
}
func (dbg *DodecahedronBufferGeometry) ToJSON() js.Value {
	return dbg.Call("toJSON")
}
func (dbg *DodecahedronBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("toNonIndexed")}
}
func (dbg *DodecahedronBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: dbg.Call("translate", x, y, z)}
}
func (dbg *DodecahedronBufferGeometry) UpdateFromObject(object *Object3D) {
	dbg.Call("updateFromObject", object)
}

type DodecahedronGeometry struct {
	js.Value
}

func NewDodecahedronGeometry(radius float64, detail float64) *DodecahedronGeometry {
	return &DodecahedronGeometry{Value: get("DodecahedronGeometry").New(radius, detail)}
}
func (dg *DodecahedronGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: dg.Get("animation")}
}
func (dg *DodecahedronGeometry) SetAnimation(v *AnimationClip) {
	dg.Set("animation", v)
}
func (dg *DodecahedronGeometry) Animations() js.Value {
	return dg.Get("animations")
}
func (dg *DodecahedronGeometry) SetAnimations(v js.Value) {
	dg.Set("animations", v)
}
func (dg *DodecahedronGeometry) Bones() js.Value {
	return dg.Get("bones")
}
func (dg *DodecahedronGeometry) SetBones(v js.Value) {
	dg.Set("bones", v)
}
func (dg *DodecahedronGeometry) BoundingBox() *Box3 {
	return &Box3{Value: dg.Get("boundingBox")}
}
func (dg *DodecahedronGeometry) SetBoundingBox(v *Box3) {
	dg.Set("boundingBox", v)
}
func (dg *DodecahedronGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: dg.Get("boundingSphere")}
}
func (dg *DodecahedronGeometry) SetBoundingSphere(v *Sphere) {
	dg.Set("boundingSphere", v)
}
func (dg *DodecahedronGeometry) Colors() js.Value {
	return dg.Get("colors")
}
func (dg *DodecahedronGeometry) SetColors(v js.Value) {
	dg.Set("colors", v)
}
func (dg *DodecahedronGeometry) ColorsNeedUpdate() bool {
	return dg.Get("colorsNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetColorsNeedUpdate(v bool) {
	dg.Set("colorsNeedUpdate", v)
}
func (dg *DodecahedronGeometry) ElementsNeedUpdate() bool {
	return dg.Get("elementsNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetElementsNeedUpdate(v bool) {
	dg.Set("elementsNeedUpdate", v)
}
func (dg *DodecahedronGeometry) FaceVertexUvs() js.Value {
	return dg.Get("faceVertexUvs")
}
func (dg *DodecahedronGeometry) SetFaceVertexUvs(v js.Value) {
	dg.Set("faceVertexUvs", v)
}
func (dg *DodecahedronGeometry) Faces() js.Value {
	return dg.Get("faces")
}
func (dg *DodecahedronGeometry) SetFaces(v js.Value) {
	dg.Set("faces", v)
}
func (dg *DodecahedronGeometry) GroupsNeedUpdate() bool {
	return dg.Get("groupsNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetGroupsNeedUpdate(v bool) {
	dg.Set("groupsNeedUpdate", v)
}
func (dg *DodecahedronGeometry) Id() int {
	return dg.Get("id").Int()
}
func (dg *DodecahedronGeometry) SetId(v int) {
	dg.Set("id", v)
}
func (dg *DodecahedronGeometry) LineDistances() js.Value {
	return dg.Get("lineDistances")
}
func (dg *DodecahedronGeometry) SetLineDistances(v js.Value) {
	dg.Set("lineDistances", v)
}
func (dg *DodecahedronGeometry) LineDistancesNeedUpdate() bool {
	return dg.Get("lineDistancesNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	dg.Set("lineDistancesNeedUpdate", v)
}
func (dg *DodecahedronGeometry) MorphNormals() js.Value {
	return dg.Get("morphNormals")
}
func (dg *DodecahedronGeometry) SetMorphNormals(v js.Value) {
	dg.Set("morphNormals", v)
}
func (dg *DodecahedronGeometry) MorphTargets() js.Value {
	return dg.Get("morphTargets")
}
func (dg *DodecahedronGeometry) SetMorphTargets(v js.Value) {
	dg.Set("morphTargets", v)
}
func (dg *DodecahedronGeometry) Name() string {
	return dg.Get("name").String()
}
func (dg *DodecahedronGeometry) SetName(v string) {
	dg.Set("name", v)
}
func (dg *DodecahedronGeometry) NormalsNeedUpdate() bool {
	return dg.Get("normalsNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetNormalsNeedUpdate(v bool) {
	dg.Set("normalsNeedUpdate", v)
}
func (dg *DodecahedronGeometry) Parameters() js.Value {
	return dg.Get("parameters")
}
func (dg *DodecahedronGeometry) SetParameters(v js.Value) {
	dg.Set("parameters", v)
}
func (dg *DodecahedronGeometry) SkinIndices() js.Value {
	return dg.Get("skinIndices")
}
func (dg *DodecahedronGeometry) SetSkinIndices(v js.Value) {
	dg.Set("skinIndices", v)
}
func (dg *DodecahedronGeometry) SkinWeights() js.Value {
	return dg.Get("skinWeights")
}
func (dg *DodecahedronGeometry) SetSkinWeights(v js.Value) {
	dg.Set("skinWeights", v)
}
func (dg *DodecahedronGeometry) Type() string {
	return dg.Get("type").String()
}
func (dg *DodecahedronGeometry) SetType(v string) {
	dg.Set("type", v)
}
func (dg *DodecahedronGeometry) Uuid() string {
	return dg.Get("uuid").String()
}
func (dg *DodecahedronGeometry) SetUuid(v string) {
	dg.Set("uuid", v)
}
func (dg *DodecahedronGeometry) UvsNeedUpdate() bool {
	return dg.Get("uvsNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetUvsNeedUpdate(v bool) {
	dg.Set("uvsNeedUpdate", v)
}
func (dg *DodecahedronGeometry) Vertices() js.Value {
	return dg.Get("vertices")
}
func (dg *DodecahedronGeometry) SetVertices(v js.Value) {
	dg.Set("vertices", v)
}
func (dg *DodecahedronGeometry) VerticesNeedUpdate() bool {
	return dg.Get("verticesNeedUpdate").Bool()
}
func (dg *DodecahedronGeometry) SetVerticesNeedUpdate(v bool) {
	dg.Set("verticesNeedUpdate", v)
}
func (dg *DodecahedronGeometry) AddEventListener(typ string, listener js.Value) {
	dg.Call("addEventListener", typ, listener)
}
func (dg *DodecahedronGeometry) ApplyMatrix(matrix *Matrix4) *Geometry {
	return &Geometry{Value: dg.Call("applyMatrix", matrix)}
}
func (dg *DodecahedronGeometry) Center() *Geometry {
	return &Geometry{Value: dg.Call("center")}
}
func (dg *DodecahedronGeometry) Clone() *DodecahedronGeometry {
	return &DodecahedronGeometry{Value: dg.Call("clone")}
}
func (dg *DodecahedronGeometry) ComputeBoundingBox() {
	dg.Call("computeBoundingBox")
}
func (dg *DodecahedronGeometry) ComputeBoundingSphere() {
	dg.Call("computeBoundingSphere")
}
func (dg *DodecahedronGeometry) ComputeFaceNormals() {
	dg.Call("computeFaceNormals")
}
func (dg *DodecahedronGeometry) ComputeFlatVertexNormals() {
	dg.Call("computeFlatVertexNormals")
}
func (dg *DodecahedronGeometry) ComputeMorphNormals() {
	dg.Call("computeMorphNormals")
}
func (dg *DodecahedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	dg.Call("computeVertexNormals", areaWeighted)
}
func (dg *DodecahedronGeometry) Copy(source *Geometry) *DodecahedronGeometry {
	return &DodecahedronGeometry{Value: dg.Call("copy", source)}
}
func (dg *DodecahedronGeometry) DispatchEvent(event js.Value) {
	dg.Call("dispatchEvent", event)
}
func (dg *DodecahedronGeometry) Dispose() {
	dg.Call("dispose")
}
func (dg *DodecahedronGeometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: dg.Call("fromBufferGeometry", geometry)}
}
func (dg *DodecahedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return dg.Call("hasEventListener", typ, listener).Bool()
}
func (dg *DodecahedronGeometry) LookAt(vector *Vector3) {
	dg.Call("lookAt", vector)
}
func (dg *DodecahedronGeometry) Merge(geometry *Geometry, matrix Matrix, materialIndexOffset int) {
	dg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (dg *DodecahedronGeometry) MergeMesh(mesh *Mesh) {
	dg.Call("mergeMesh", mesh)
}
func (dg *DodecahedronGeometry) MergeVertices() float64 {
	return dg.Call("mergeVertices").Float()
}
func (dg *DodecahedronGeometry) Normalize() *Geometry {
	return &Geometry{Value: dg.Call("normalize")}
}
func (dg *DodecahedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	dg.Call("removeEventListener", typ, listener)
}
func (dg *DodecahedronGeometry) RotateX(angle float64) *Geometry {
	return &Geometry{Value: dg.Call("rotateX", angle)}
}
func (dg *DodecahedronGeometry) RotateY(angle float64) *Geometry {
	return &Geometry{Value: dg.Call("rotateY", angle)}
}
func (dg *DodecahedronGeometry) RotateZ(angle float64) *Geometry {
	return &Geometry{Value: dg.Call("rotateZ", angle)}
}
func (dg *DodecahedronGeometry) Scale(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: dg.Call("scale", x, y, z)}
}
func (dg *DodecahedronGeometry) SetFromPoints(points js.Value) *DodecahedronGeometry {
	return &DodecahedronGeometry{Value: dg.Call("setFromPoints", points)}
}
func (dg *DodecahedronGeometry) SortFacesByMaterialIndex() {
	dg.Call("sortFacesByMaterialIndex")
}
func (dg *DodecahedronGeometry) ToJSON() js.Value {
	return dg.Call("toJSON")
}
func (dg *DodecahedronGeometry) Translate(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: dg.Call("translate", x, y, z)}
}

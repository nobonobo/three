// Code generated from "geometries/TextGeometry.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type TextGeometryParameters interface {
}
type TextBufferGeometry struct {
	js.Value
}

func NewTextBufferGeometry(text string, parameters TextGeometryParameters) *TextBufferGeometry {
	return &TextBufferGeometry{Value: get("TextBufferGeometry").New(text, parameters)}
}
func (tbg *TextBufferGeometry) Attributes() js.Value {
	return tbg.Get("attributes")
}
func (tbg *TextBufferGeometry) SetAttributes(v js.Value) {
	tbg.Set("attributes", v)
}
func (tbg *TextBufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tbg.Get("boundingBox")}
}
func (tbg *TextBufferGeometry) SetBoundingBox(v *Box3) {
	tbg.Set("boundingBox", v)
}
func (tbg *TextBufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tbg.Get("boundingSphere")}
}
func (tbg *TextBufferGeometry) SetBoundingSphere(v *Sphere) {
	tbg.Set("boundingSphere", v)
}
func (tbg *TextBufferGeometry) DrawRange() js.Value {
	return tbg.Get("drawRange")
}
func (tbg *TextBufferGeometry) SetDrawRange(v js.Value) {
	tbg.Set("drawRange", v)
}
func (tbg *TextBufferGeometry) Drawcalls() js.Value {
	return tbg.Get("drawcalls")
}
func (tbg *TextBufferGeometry) SetDrawcalls(v js.Value) {
	tbg.Set("drawcalls", v)
}
func (tbg *TextBufferGeometry) Groups() js.Value {
	return tbg.Get("groups")
}
func (tbg *TextBufferGeometry) SetGroups(v js.Value) {
	tbg.Set("groups", v)
}
func (tbg *TextBufferGeometry) Id() int {
	return tbg.Get("id").Int()
}
func (tbg *TextBufferGeometry) SetId(v int) {
	tbg.Set("id", v)
}
func (tbg *TextBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Get("index")}
}
func (tbg *TextBufferGeometry) SetIndex(v *BufferAttribute) {
	tbg.Set("index", v)
}
func (tbg *TextBufferGeometry) MorphAttributes() js.Value {
	return tbg.Get("morphAttributes")
}
func (tbg *TextBufferGeometry) SetMorphAttributes(v js.Value) {
	tbg.Set("morphAttributes", v)
}
func (tbg *TextBufferGeometry) Name() string {
	return tbg.Get("name").String()
}
func (tbg *TextBufferGeometry) SetName(v string) {
	tbg.Set("name", v)
}
func (tbg *TextBufferGeometry) Offsets() js.Value {
	return tbg.Get("offsets")
}
func (tbg *TextBufferGeometry) SetOffsets(v js.Value) {
	tbg.Set("offsets", v)
}
func (tbg *TextBufferGeometry) Parameters() js.Value {
	return tbg.Get("parameters")
}
func (tbg *TextBufferGeometry) SetParameters(v js.Value) {
	tbg.Set("parameters", v)
}
func (tbg *TextBufferGeometry) Type() string {
	return tbg.Get("type").String()
}
func (tbg *TextBufferGeometry) SetType(v string) {
	tbg.Set("type", v)
}
func (tbg *TextBufferGeometry) UserData() js.Value {
	return tbg.Get("userData")
}
func (tbg *TextBufferGeometry) SetUserData(v js.Value) {
	tbg.Set("userData", v)
}
func (tbg *TextBufferGeometry) Uuid() string {
	return tbg.Get("uuid").String()
}
func (tbg *TextBufferGeometry) SetUuid(v string) {
	tbg.Set("uuid", v)
}
func (tbg *TextBufferGeometry) MaxIndex() int {
	return tbg.Get("MaxIndex").Int()
}
func (tbg *TextBufferGeometry) SetMaxIndex(v int) {
	tbg.Set("MaxIndex", v)
}
func (tbg *TextBufferGeometry) WorldUVGenerator() js.Value {
	return tbg.Get("WorldUVGenerator")
}
func (tbg *TextBufferGeometry) SetWorldUVGenerator(v js.Value) {
	tbg.Set("WorldUVGenerator", v)
}
func (tbg *TextBufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("addAttribute", name, attribute)}
}
func (tbg *TextBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tbg.Call("addAttribute", name, array, itemSize)
}
func (tbg *TextBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tbg.Call("addDrawCall", start, count, indexOffset)
}
func (tbg *TextBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tbg.Call("addEventListener", typ, listener)
}
func (tbg *TextBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tbg.Call("addGroup", start, count, materialIndex)
}
func (tbg *TextBufferGeometry) AddIndex(index js.Value) {
	tbg.Call("addIndex", index)
}
func (tbg *TextBufferGeometry) AddShape(shape *Shape, options js.Value) {
	tbg.Call("addShape", shape, options)
}
func (tbg *TextBufferGeometry) AddShapeList(shapes js.Value, options js.Value) {
	tbg.Call("addShapeList", shapes, options)
}
func (tbg *TextBufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("applyMatrix", matrix)}
}
func (tbg *TextBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("center")}
}
func (tbg *TextBufferGeometry) ClearDrawCalls() {
	tbg.Call("clearDrawCalls")
}
func (tbg *TextBufferGeometry) ClearGroups() {
	tbg.Call("clearGroups")
}
func (tbg *TextBufferGeometry) Clone() *TextBufferGeometry {
	return &TextBufferGeometry{Value: tbg.Call("clone")}
}
func (tbg *TextBufferGeometry) ComputeBoundingBox() {
	tbg.Call("computeBoundingBox")
}
func (tbg *TextBufferGeometry) ComputeBoundingSphere() {
	tbg.Call("computeBoundingSphere")
}
func (tbg *TextBufferGeometry) ComputeVertexNormals() {
	tbg.Call("computeVertexNormals")
}
func (tbg *TextBufferGeometry) Copy(source *BufferGeometry) *TextBufferGeometry {
	return &TextBufferGeometry{Value: tbg.Call("copy", source)}
}
func (tbg *TextBufferGeometry) DispatchEvent(event js.Value) {
	tbg.Call("dispatchEvent", event)
}
func (tbg *TextBufferGeometry) Dispose() {
	tbg.Call("dispose")
}
func (tbg *TextBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("fromDirectGeometry", geometry)}
}
func (tbg *TextBufferGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("fromGeometry", geometry, settings)}
}
func (tbg *TextBufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getAttribute", name)}
}
func (tbg *TextBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: tbg.Call("getIndex")}
}
func (tbg *TextBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tbg.Call("hasEventListener", typ, listener).Bool()
}
func (tbg *TextBufferGeometry) LookAt(v *Vector3) {
	tbg.Call("lookAt", v)
}
func (tbg *TextBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("merge", geometry, offset)}
}
func (tbg *TextBufferGeometry) NormalizeNormals() {
	tbg.Call("normalizeNormals")
}
func (tbg *TextBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("removeAttribute", name)}
}
func (tbg *TextBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tbg.Call("removeEventListener", typ, listener)
}
func (tbg *TextBufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateX", angle)}
}
func (tbg *TextBufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateY", angle)}
}
func (tbg *TextBufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("rotateZ", angle)}
}
func (tbg *TextBufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("scale", x, y, z)}
}
func (tbg *TextBufferGeometry) SetDrawRange2(start int, count int) {
	tbg.Call("setDrawRange", start, count)
}
func (tbg *TextBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("setFromObject", object)}
}
func (tbg *TextBufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("setFromPoints", points)}
}
func (tbg *TextBufferGeometry) SetIndex2(index *BufferAttribute) {
	tbg.Call("setIndex", index)
}
func (tbg *TextBufferGeometry) ToJSON() js.Value {
	return tbg.Call("toJSON")
}
func (tbg *TextBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("toNonIndexed")}
}
func (tbg *TextBufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: tbg.Call("translate", x, y, z)}
}
func (tbg *TextBufferGeometry) UpdateFromObject(object *Object3D) {
	tbg.Call("updateFromObject", object)
}

type TextGeometry struct {
	js.Value
}

func NewTextGeometry(text string, parameters TextGeometryParameters) *TextGeometry {
	return &TextGeometry{Value: get("TextGeometry").New(text, parameters)}
}
func (tg *TextGeometry) Animation() *AnimationClip {
	return &AnimationClip{Value: tg.Get("animation")}
}
func (tg *TextGeometry) SetAnimation(v *AnimationClip) {
	tg.Set("animation", v)
}
func (tg *TextGeometry) Animations() js.Value {
	return tg.Get("animations")
}
func (tg *TextGeometry) SetAnimations(v js.Value) {
	tg.Set("animations", v)
}
func (tg *TextGeometry) Bones() js.Value {
	return tg.Get("bones")
}
func (tg *TextGeometry) SetBones(v js.Value) {
	tg.Set("bones", v)
}
func (tg *TextGeometry) BoundingBox() *Box3 {
	return &Box3{Value: tg.Get("boundingBox")}
}
func (tg *TextGeometry) SetBoundingBox(v *Box3) {
	tg.Set("boundingBox", v)
}
func (tg *TextGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: tg.Get("boundingSphere")}
}
func (tg *TextGeometry) SetBoundingSphere(v *Sphere) {
	tg.Set("boundingSphere", v)
}
func (tg *TextGeometry) Colors() js.Value {
	return tg.Get("colors")
}
func (tg *TextGeometry) SetColors(v js.Value) {
	tg.Set("colors", v)
}
func (tg *TextGeometry) ColorsNeedUpdate() bool {
	return tg.Get("colorsNeedUpdate").Bool()
}
func (tg *TextGeometry) SetColorsNeedUpdate(v bool) {
	tg.Set("colorsNeedUpdate", v)
}
func (tg *TextGeometry) ElementsNeedUpdate() bool {
	return tg.Get("elementsNeedUpdate").Bool()
}
func (tg *TextGeometry) SetElementsNeedUpdate(v bool) {
	tg.Set("elementsNeedUpdate", v)
}
func (tg *TextGeometry) FaceVertexUvs() js.Value {
	return tg.Get("faceVertexUvs")
}
func (tg *TextGeometry) SetFaceVertexUvs(v js.Value) {
	tg.Set("faceVertexUvs", v)
}
func (tg *TextGeometry) Faces() js.Value {
	return tg.Get("faces")
}
func (tg *TextGeometry) SetFaces(v js.Value) {
	tg.Set("faces", v)
}
func (tg *TextGeometry) GroupsNeedUpdate() bool {
	return tg.Get("groupsNeedUpdate").Bool()
}
func (tg *TextGeometry) SetGroupsNeedUpdate(v bool) {
	tg.Set("groupsNeedUpdate", v)
}
func (tg *TextGeometry) Id() int {
	return tg.Get("id").Int()
}
func (tg *TextGeometry) SetId(v int) {
	tg.Set("id", v)
}
func (tg *TextGeometry) LineDistances() js.Value {
	return tg.Get("lineDistances")
}
func (tg *TextGeometry) SetLineDistances(v js.Value) {
	tg.Set("lineDistances", v)
}
func (tg *TextGeometry) LineDistancesNeedUpdate() bool {
	return tg.Get("lineDistancesNeedUpdate").Bool()
}
func (tg *TextGeometry) SetLineDistancesNeedUpdate(v bool) {
	tg.Set("lineDistancesNeedUpdate", v)
}
func (tg *TextGeometry) MorphNormals() js.Value {
	return tg.Get("morphNormals")
}
func (tg *TextGeometry) SetMorphNormals(v js.Value) {
	tg.Set("morphNormals", v)
}
func (tg *TextGeometry) MorphTargets() js.Value {
	return tg.Get("morphTargets")
}
func (tg *TextGeometry) SetMorphTargets(v js.Value) {
	tg.Set("morphTargets", v)
}
func (tg *TextGeometry) Name() string {
	return tg.Get("name").String()
}
func (tg *TextGeometry) SetName(v string) {
	tg.Set("name", v)
}
func (tg *TextGeometry) NormalsNeedUpdate() bool {
	return tg.Get("normalsNeedUpdate").Bool()
}
func (tg *TextGeometry) SetNormalsNeedUpdate(v bool) {
	tg.Set("normalsNeedUpdate", v)
}
func (tg *TextGeometry) Parameters() js.Value {
	return tg.Get("parameters")
}
func (tg *TextGeometry) SetParameters(v js.Value) {
	tg.Set("parameters", v)
}
func (tg *TextGeometry) SkinIndices() js.Value {
	return tg.Get("skinIndices")
}
func (tg *TextGeometry) SetSkinIndices(v js.Value) {
	tg.Set("skinIndices", v)
}
func (tg *TextGeometry) SkinWeights() js.Value {
	return tg.Get("skinWeights")
}
func (tg *TextGeometry) SetSkinWeights(v js.Value) {
	tg.Set("skinWeights", v)
}
func (tg *TextGeometry) Type() string {
	return tg.Get("type").String()
}
func (tg *TextGeometry) SetType(v string) {
	tg.Set("type", v)
}
func (tg *TextGeometry) Uuid() string {
	return tg.Get("uuid").String()
}
func (tg *TextGeometry) SetUuid(v string) {
	tg.Set("uuid", v)
}
func (tg *TextGeometry) UvsNeedUpdate() bool {
	return tg.Get("uvsNeedUpdate").Bool()
}
func (tg *TextGeometry) SetUvsNeedUpdate(v bool) {
	tg.Set("uvsNeedUpdate", v)
}
func (tg *TextGeometry) Vertices() js.Value {
	return tg.Get("vertices")
}
func (tg *TextGeometry) SetVertices(v js.Value) {
	tg.Set("vertices", v)
}
func (tg *TextGeometry) VerticesNeedUpdate() bool {
	return tg.Get("verticesNeedUpdate").Bool()
}
func (tg *TextGeometry) SetVerticesNeedUpdate(v bool) {
	tg.Set("verticesNeedUpdate", v)
}
func (tg *TextGeometry) WorldUVGenerator() js.Value {
	return tg.Get("WorldUVGenerator")
}
func (tg *TextGeometry) SetWorldUVGenerator(v js.Value) {
	tg.Set("WorldUVGenerator", v)
}
func (tg *TextGeometry) AddEventListener(typ string, listener js.Value) {
	tg.Call("addEventListener", typ, listener)
}
func (tg *TextGeometry) AddShape(shape *Shape, options js.Value) {
	tg.Call("addShape", shape, options)
}
func (tg *TextGeometry) AddShapeList(shapes js.Value, options js.Value) {
	tg.Call("addShapeList", shapes, options)
}
func (tg *TextGeometry) ApplyMatrix(matrix *Matrix4) *Geometry {
	return &Geometry{Value: tg.Call("applyMatrix", matrix)}
}
func (tg *TextGeometry) Center() *Geometry {
	return &Geometry{Value: tg.Call("center")}
}
func (tg *TextGeometry) Clone() *TextGeometry {
	return &TextGeometry{Value: tg.Call("clone")}
}
func (tg *TextGeometry) ComputeBoundingBox() {
	tg.Call("computeBoundingBox")
}
func (tg *TextGeometry) ComputeBoundingSphere() {
	tg.Call("computeBoundingSphere")
}
func (tg *TextGeometry) ComputeFaceNormals() {
	tg.Call("computeFaceNormals")
}
func (tg *TextGeometry) ComputeFlatVertexNormals() {
	tg.Call("computeFlatVertexNormals")
}
func (tg *TextGeometry) ComputeMorphNormals() {
	tg.Call("computeMorphNormals")
}
func (tg *TextGeometry) ComputeVertexNormals(areaWeighted bool) {
	tg.Call("computeVertexNormals", areaWeighted)
}
func (tg *TextGeometry) Copy(source *Geometry) *TextGeometry {
	return &TextGeometry{Value: tg.Call("copy", source)}
}
func (tg *TextGeometry) DispatchEvent(event js.Value) {
	tg.Call("dispatchEvent", event)
}
func (tg *TextGeometry) Dispose() {
	tg.Call("dispose")
}
func (tg *TextGeometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: tg.Call("fromBufferGeometry", geometry)}
}
func (tg *TextGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tg.Call("hasEventListener", typ, listener).Bool()
}
func (tg *TextGeometry) LookAt(vector *Vector3) {
	tg.Call("lookAt", vector)
}
func (tg *TextGeometry) Merge(geometry *Geometry, matrix Matrix, materialIndexOffset int) {
	tg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (tg *TextGeometry) MergeMesh(mesh *Mesh) {
	tg.Call("mergeMesh", mesh)
}
func (tg *TextGeometry) MergeVertices() float64 {
	return tg.Call("mergeVertices").Float()
}
func (tg *TextGeometry) Normalize() *Geometry {
	return &Geometry{Value: tg.Call("normalize")}
}
func (tg *TextGeometry) RemoveEventListener(typ string, listener js.Value) {
	tg.Call("removeEventListener", typ, listener)
}
func (tg *TextGeometry) RotateX(angle float64) *Geometry {
	return &Geometry{Value: tg.Call("rotateX", angle)}
}
func (tg *TextGeometry) RotateY(angle float64) *Geometry {
	return &Geometry{Value: tg.Call("rotateY", angle)}
}
func (tg *TextGeometry) RotateZ(angle float64) *Geometry {
	return &Geometry{Value: tg.Call("rotateZ", angle)}
}
func (tg *TextGeometry) Scale(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: tg.Call("scale", x, y, z)}
}
func (tg *TextGeometry) SetFromPoints(points js.Value) *TextGeometry {
	return &TextGeometry{Value: tg.Call("setFromPoints", points)}
}
func (tg *TextGeometry) SortFacesByMaterialIndex() {
	tg.Call("sortFacesByMaterialIndex")
}
func (tg *TextGeometry) ToJSON() js.Value {
	return tg.Call("toJSON")
}
func (tg *TextGeometry) Translate(x float64, y float64, z float64) *Geometry {
	return &Geometry{Value: tg.Call("translate", x, y, z)}
}

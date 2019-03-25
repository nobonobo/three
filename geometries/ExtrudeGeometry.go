package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/extras/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type ExtrudeGeometryOptions interface {
}
type UVGenerator interface {
	GenerateSideWallUV(geometry *ExtrudeBufferGeometry, vertices []int, indexA int, indexB int, indexC int, indexD int) []math.Vector2
	GenerateTopUV(geometry *ExtrudeBufferGeometry, vertices []int, indexA int, indexB int, indexC int) []math.Vector2
}
type ExtrudeBufferGeometry struct {
	js.Value
}

func NewExtrudeBufferGeometry(shapes core.Shape, options ExtrudeGeometryOptions) *ExtrudeBufferGeometry {
	return &ExtrudeBufferGeometry{Value: get("ExtrudeBufferGeometry").New(shapes, options)}
}
func (ebg *ExtrudeBufferGeometry) Attributes() js.Value {
	return ebg.Get("attributes")
}
func (ebg *ExtrudeBufferGeometry) SetAttributes(v js.Value) {
	ebg.Set("attributes", v)
}
func (ebg *ExtrudeBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: ebg.Get("boundingBox")}
}
func (ebg *ExtrudeBufferGeometry) SetBoundingBox(v *math.Box3) {
	ebg.Set("boundingBox", v)
}
func (ebg *ExtrudeBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: ebg.Get("boundingSphere")}
}
func (ebg *ExtrudeBufferGeometry) SetBoundingSphere(v *math.Sphere) {
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
func (ebg *ExtrudeBufferGeometry) Groups() []js.Value {
	return []js.Value(ebg.Get("groups"))
}
func (ebg *ExtrudeBufferGeometry) SetGroups(v []js.Value) {
	ebg.Set("groups", v)
}
func (ebg *ExtrudeBufferGeometry) Id() int {
	return ebg.Get("id").Int()
}
func (ebg *ExtrudeBufferGeometry) SetId(v int) {
	ebg.Set("id", v)
}
func (ebg *ExtrudeBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: ebg.Get("index")}
}
func (ebg *ExtrudeBufferGeometry) SetIndex(v *core.BufferAttribute) {
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
func (ebg *ExtrudeBufferGeometry) WorldUVGenerator() UVGenerator {
	return UVGenerator(ebg.Get("WorldUVGenerator"))
}
func (ebg *ExtrudeBufferGeometry) SetWorldUVGenerator(v UVGenerator) {
	ebg.Set("WorldUVGenerator", v)
}
func (ebg *ExtrudeBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("addAttribute", name, attribute)}
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
func (ebg *ExtrudeBufferGeometry) AddShape(shape *core.Shape, options js.Value) {
	ebg.Call("addShape", shape, options)
}
func (ebg *ExtrudeBufferGeometry) AddShapeList(shapes []core.Shape, options js.Value) {
	ebg.Call("addShapeList", shapes, options)
}
func (ebg *ExtrudeBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("applyMatrix", matrix)}
}
func (ebg *ExtrudeBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("center")}
}
func (ebg *ExtrudeBufferGeometry) ClearDrawCalls() {
	ebg.Call("clearDrawCalls")
}
func (ebg *ExtrudeBufferGeometry) ClearGroups() {
	ebg.Call("clearGroups")
}
func (ebg *ExtrudeBufferGeometry) Clone() this {
	return this(ebg.Call("clone"))
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
func (ebg *ExtrudeBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(ebg.Call("copy", source))
}
func (ebg *ExtrudeBufferGeometry) DispatchEvent(event js.Value) {
	ebg.Call("dispatchEvent", event)
}
func (ebg *ExtrudeBufferGeometry) Dispose() {
	ebg.Call("dispose")
}
func (ebg *ExtrudeBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("fromDirectGeometry", geometry)}
}
func (ebg *ExtrudeBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("fromGeometry", geometry, settings)}
}
func (ebg *ExtrudeBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(ebg.Call("getAttribute", name))
}
func (ebg *ExtrudeBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: ebg.Call("getIndex")}
}
func (ebg *ExtrudeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ebg.Call("hasEventListener", typ, listener).Bool()
}
func (ebg *ExtrudeBufferGeometry) LookAt(v *math.Vector3) {
	ebg.Call("lookAt", v)
}
func (ebg *ExtrudeBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("merge", geometry, offset)}
}
func (ebg *ExtrudeBufferGeometry) NormalizeNormals() {
	ebg.Call("normalizeNormals")
}
func (ebg *ExtrudeBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("removeAttribute", name)}
}
func (ebg *ExtrudeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	ebg.Call("removeEventListener", typ, listener)
}
func (ebg *ExtrudeBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("rotateX", angle)}
}
func (ebg *ExtrudeBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("rotateY", angle)}
}
func (ebg *ExtrudeBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("rotateZ", angle)}
}
func (ebg *ExtrudeBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("scale", x, y, z)}
}
func (ebg *ExtrudeBufferGeometry) SetDrawRange(start int, count int) {
	ebg.Call("setDrawRange", start, count)
}
func (ebg *ExtrudeBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("setFromObject", object)}
}
func (ebg *ExtrudeBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("setFromPoints", points)}
}
func (ebg *ExtrudeBufferGeometry) SetIndex(index core.BufferAttribute) {
	ebg.Call("setIndex", index)
}
func (ebg *ExtrudeBufferGeometry) ToJSON() js.Value {
	return ebg.Call("toJSON")
}
func (ebg *ExtrudeBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("toNonIndexed")}
}
func (ebg *ExtrudeBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ebg.Call("translate", x, y, z)}
}
func (ebg *ExtrudeBufferGeometry) UpdateFromObject(object *core.Object3D) {
	ebg.Call("updateFromObject", object)
}

type ExtrudeGeometry struct {
	js.Value
}

func NewExtrudeGeometry(shapes core.Shape, options ExtrudeGeometryOptions) *ExtrudeGeometry {
	return &ExtrudeGeometry{Value: get("ExtrudeGeometry").New(shapes, options)}
}
func (eg *ExtrudeGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: eg.Get("animation")}
}
func (eg *ExtrudeGeometry) SetAnimation(v *animation.AnimationClip) {
	eg.Set("animation", v)
}
func (eg *ExtrudeGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(eg.Get("animations"))
}
func (eg *ExtrudeGeometry) SetAnimations(v []animation.AnimationClip) {
	eg.Set("animations", v)
}
func (eg *ExtrudeGeometry) Bones() []objects.Bone {
	return []objects.Bone(eg.Get("bones"))
}
func (eg *ExtrudeGeometry) SetBones(v []objects.Bone) {
	eg.Set("bones", v)
}
func (eg *ExtrudeGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: eg.Get("boundingBox")}
}
func (eg *ExtrudeGeometry) SetBoundingBox(v *math.Box3) {
	eg.Set("boundingBox", v)
}
func (eg *ExtrudeGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: eg.Get("boundingSphere")}
}
func (eg *ExtrudeGeometry) SetBoundingSphere(v *math.Sphere) {
	eg.Set("boundingSphere", v)
}
func (eg *ExtrudeGeometry) Colors() []math.Color {
	return []math.Color(eg.Get("colors"))
}
func (eg *ExtrudeGeometry) SetColors(v []math.Color) {
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
func (eg *ExtrudeGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(eg.Get("faceVertexUvs"))
}
func (eg *ExtrudeGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	eg.Set("faceVertexUvs", v)
}
func (eg *ExtrudeGeometry) Faces() []core.Face3 {
	return []core.Face3(eg.Get("faces"))
}
func (eg *ExtrudeGeometry) SetFaces(v []core.Face3) {
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
func (eg *ExtrudeGeometry) LineDistances() []int {
	return []int(eg.Get("lineDistances"))
}
func (eg *ExtrudeGeometry) SetLineDistances(v []int) {
	eg.Set("lineDistances", v)
}
func (eg *ExtrudeGeometry) LineDistancesNeedUpdate() bool {
	return eg.Get("lineDistancesNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetLineDistancesNeedUpdate(v bool) {
	eg.Set("lineDistancesNeedUpdate", v)
}
func (eg *ExtrudeGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(eg.Get("morphNormals"))
}
func (eg *ExtrudeGeometry) SetMorphNormals(v []core.MorphNormals) {
	eg.Set("morphNormals", v)
}
func (eg *ExtrudeGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(eg.Get("morphTargets"))
}
func (eg *ExtrudeGeometry) SetMorphTargets(v []core.MorphTarget) {
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
func (eg *ExtrudeGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(eg.Get("skinIndices"))
}
func (eg *ExtrudeGeometry) SetSkinIndices(v []math.Vector4) {
	eg.Set("skinIndices", v)
}
func (eg *ExtrudeGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(eg.Get("skinWeights"))
}
func (eg *ExtrudeGeometry) SetSkinWeights(v []math.Vector4) {
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
func (eg *ExtrudeGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(eg.Get("vertices"))
}
func (eg *ExtrudeGeometry) SetVertices(v []math.Vector3) {
	eg.Set("vertices", v)
}
func (eg *ExtrudeGeometry) VerticesNeedUpdate() bool {
	return eg.Get("verticesNeedUpdate").Bool()
}
func (eg *ExtrudeGeometry) SetVerticesNeedUpdate(v bool) {
	eg.Set("verticesNeedUpdate", v)
}
func (eg *ExtrudeGeometry) WorldUVGenerator() UVGenerator {
	return UVGenerator(eg.Get("WorldUVGenerator"))
}
func (eg *ExtrudeGeometry) SetWorldUVGenerator(v UVGenerator) {
	eg.Set("WorldUVGenerator", v)
}
func (eg *ExtrudeGeometry) AddEventListener(typ string, listener js.Value) {
	eg.Call("addEventListener", typ, listener)
}
func (eg *ExtrudeGeometry) AddShape(shape *core.Shape, options js.Value) {
	eg.Call("addShape", shape, options)
}
func (eg *ExtrudeGeometry) AddShapeList(shapes []core.Shape, options js.Value) {
	eg.Call("addShapeList", shapes, options)
}
func (eg *ExtrudeGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: eg.Call("applyMatrix", matrix)}
}
func (eg *ExtrudeGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: eg.Call("center")}
}
func (eg *ExtrudeGeometry) Clone() this {
	return this(eg.Call("clone"))
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
func (eg *ExtrudeGeometry) Copy(source *core.Geometry) this {
	return this(eg.Call("copy", source))
}
func (eg *ExtrudeGeometry) DispatchEvent(event js.Value) {
	eg.Call("dispatchEvent", event)
}
func (eg *ExtrudeGeometry) Dispose() {
	eg.Call("dispose")
}
func (eg *ExtrudeGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: eg.Call("fromBufferGeometry", geometry)}
}
func (eg *ExtrudeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return eg.Call("hasEventListener", typ, listener).Bool()
}
func (eg *ExtrudeGeometry) LookAt(vector *math.Vector3) {
	eg.Call("lookAt", vector)
}
func (eg *ExtrudeGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	eg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (eg *ExtrudeGeometry) MergeMesh(mesh *objects.Mesh) {
	eg.Call("mergeMesh", mesh)
}
func (eg *ExtrudeGeometry) MergeVertices() int {
	return eg.Call("mergeVertices").Int()
}
func (eg *ExtrudeGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: eg.Call("normalize")}
}
func (eg *ExtrudeGeometry) RemoveEventListener(typ string, listener js.Value) {
	eg.Call("removeEventListener", typ, listener)
}
func (eg *ExtrudeGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: eg.Call("rotateX", angle)}
}
func (eg *ExtrudeGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: eg.Call("rotateY", angle)}
}
func (eg *ExtrudeGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: eg.Call("rotateZ", angle)}
}
func (eg *ExtrudeGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: eg.Call("scale", x, y, z)}
}
func (eg *ExtrudeGeometry) SetFromPoints(points Array) this {
	return this(eg.Call("setFromPoints", points))
}
func (eg *ExtrudeGeometry) SortFacesByMaterialIndex() {
	eg.Call("sortFacesByMaterialIndex")
}
func (eg *ExtrudeGeometry) ToJSON() js.Value {
	return eg.Call("toJSON")
}
func (eg *ExtrudeGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: eg.Call("translate", x, y, z)}
}

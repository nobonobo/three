package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type IcosahedronBufferGeometry struct {
	js.Value
}

func NewIcosahedronBufferGeometry(radius int, detail int) *IcosahedronBufferGeometry {
	return &IcosahedronBufferGeometry{Value: get("IcosahedronBufferGeometry").New(radius, detail)}
}
func (ibg *IcosahedronBufferGeometry) Attributes() js.Value {
	return ibg.Get("attributes")
}
func (ibg *IcosahedronBufferGeometry) SetAttributes(v js.Value) {
	ibg.Set("attributes", v)
}
func (ibg *IcosahedronBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: ibg.Get("boundingBox")}
}
func (ibg *IcosahedronBufferGeometry) SetBoundingBox(v *math.Box3) {
	ibg.Set("boundingBox", v)
}
func (ibg *IcosahedronBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: ibg.Get("boundingSphere")}
}
func (ibg *IcosahedronBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	ibg.Set("boundingSphere", v)
}
func (ibg *IcosahedronBufferGeometry) DrawRange() js.Value {
	return ibg.Get("drawRange")
}
func (ibg *IcosahedronBufferGeometry) SetDrawRange(v js.Value) {
	ibg.Set("drawRange", v)
}
func (ibg *IcosahedronBufferGeometry) Drawcalls() js.Value {
	return ibg.Get("drawcalls")
}
func (ibg *IcosahedronBufferGeometry) SetDrawcalls(v js.Value) {
	ibg.Set("drawcalls", v)
}
func (ibg *IcosahedronBufferGeometry) Groups() []js.Value {
	return []js.Value(ibg.Get("groups"))
}
func (ibg *IcosahedronBufferGeometry) SetGroups(v []js.Value) {
	ibg.Set("groups", v)
}
func (ibg *IcosahedronBufferGeometry) Id() int {
	return ibg.Get("id").Int()
}
func (ibg *IcosahedronBufferGeometry) SetId(v int) {
	ibg.Set("id", v)
}
func (ibg *IcosahedronBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: ibg.Get("index")}
}
func (ibg *IcosahedronBufferGeometry) SetIndex(v *core.BufferAttribute) {
	ibg.Set("index", v)
}
func (ibg *IcosahedronBufferGeometry) MorphAttributes() js.Value {
	return ibg.Get("morphAttributes")
}
func (ibg *IcosahedronBufferGeometry) SetMorphAttributes(v js.Value) {
	ibg.Set("morphAttributes", v)
}
func (ibg *IcosahedronBufferGeometry) Name() string {
	return ibg.Get("name").String()
}
func (ibg *IcosahedronBufferGeometry) SetName(v string) {
	ibg.Set("name", v)
}
func (ibg *IcosahedronBufferGeometry) Offsets() js.Value {
	return ibg.Get("offsets")
}
func (ibg *IcosahedronBufferGeometry) SetOffsets(v js.Value) {
	ibg.Set("offsets", v)
}
func (ibg *IcosahedronBufferGeometry) Parameters() js.Value {
	return ibg.Get("parameters")
}
func (ibg *IcosahedronBufferGeometry) SetParameters(v js.Value) {
	ibg.Set("parameters", v)
}
func (ibg *IcosahedronBufferGeometry) Type() string {
	return ibg.Get("type").String()
}
func (ibg *IcosahedronBufferGeometry) SetType(v string) {
	ibg.Set("type", v)
}
func (ibg *IcosahedronBufferGeometry) UserData() js.Value {
	return ibg.Get("userData")
}
func (ibg *IcosahedronBufferGeometry) SetUserData(v js.Value) {
	ibg.Set("userData", v)
}
func (ibg *IcosahedronBufferGeometry) Uuid() string {
	return ibg.Get("uuid").String()
}
func (ibg *IcosahedronBufferGeometry) SetUuid(v string) {
	ibg.Set("uuid", v)
}
func (ibg *IcosahedronBufferGeometry) MaxIndex() int {
	return ibg.Get("MaxIndex").Int()
}
func (ibg *IcosahedronBufferGeometry) SetMaxIndex(v int) {
	ibg.Set("MaxIndex", v)
}
func (ibg *IcosahedronBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("addAttribute", name, attribute)}
}
func (ibg *IcosahedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return ibg.Call("addAttribute", name, array, itemSize)
}
func (ibg *IcosahedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	ibg.Call("addDrawCall", start, count, indexOffset)
}
func (ibg *IcosahedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	ibg.Call("addEventListener", typ, listener)
}
func (ibg *IcosahedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	ibg.Call("addGroup", start, count, materialIndex)
}
func (ibg *IcosahedronBufferGeometry) AddIndex(index js.Value) {
	ibg.Call("addIndex", index)
}
func (ibg *IcosahedronBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("applyMatrix", matrix)}
}
func (ibg *IcosahedronBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("center")}
}
func (ibg *IcosahedronBufferGeometry) ClearDrawCalls() {
	ibg.Call("clearDrawCalls")
}
func (ibg *IcosahedronBufferGeometry) ClearGroups() {
	ibg.Call("clearGroups")
}
func (ibg *IcosahedronBufferGeometry) Clone() this {
	return this(ibg.Call("clone"))
}
func (ibg *IcosahedronBufferGeometry) ComputeBoundingBox() {
	ibg.Call("computeBoundingBox")
}
func (ibg *IcosahedronBufferGeometry) ComputeBoundingSphere() {
	ibg.Call("computeBoundingSphere")
}
func (ibg *IcosahedronBufferGeometry) ComputeVertexNormals() {
	ibg.Call("computeVertexNormals")
}
func (ibg *IcosahedronBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(ibg.Call("copy", source))
}
func (ibg *IcosahedronBufferGeometry) DispatchEvent(event js.Value) {
	ibg.Call("dispatchEvent", event)
}
func (ibg *IcosahedronBufferGeometry) Dispose() {
	ibg.Call("dispose")
}
func (ibg *IcosahedronBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("fromDirectGeometry", geometry)}
}
func (ibg *IcosahedronBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("fromGeometry", geometry, settings)}
}
func (ibg *IcosahedronBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(ibg.Call("getAttribute", name))
}
func (ibg *IcosahedronBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: ibg.Call("getIndex")}
}
func (ibg *IcosahedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ibg.Call("hasEventListener", typ, listener).Bool()
}
func (ibg *IcosahedronBufferGeometry) LookAt(v *math.Vector3) {
	ibg.Call("lookAt", v)
}
func (ibg *IcosahedronBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("merge", geometry, offset)}
}
func (ibg *IcosahedronBufferGeometry) NormalizeNormals() {
	ibg.Call("normalizeNormals")
}
func (ibg *IcosahedronBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("removeAttribute", name)}
}
func (ibg *IcosahedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	ibg.Call("removeEventListener", typ, listener)
}
func (ibg *IcosahedronBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("rotateX", angle)}
}
func (ibg *IcosahedronBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("rotateY", angle)}
}
func (ibg *IcosahedronBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("rotateZ", angle)}
}
func (ibg *IcosahedronBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("scale", x, y, z)}
}
func (ibg *IcosahedronBufferGeometry) SetDrawRange(start int, count int) {
	ibg.Call("setDrawRange", start, count)
}
func (ibg *IcosahedronBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("setFromObject", object)}
}
func (ibg *IcosahedronBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("setFromPoints", points)}
}
func (ibg *IcosahedronBufferGeometry) SetIndex(index core.BufferAttribute) {
	ibg.Call("setIndex", index)
}
func (ibg *IcosahedronBufferGeometry) ToJSON() js.Value {
	return ibg.Call("toJSON")
}
func (ibg *IcosahedronBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("toNonIndexed")}
}
func (ibg *IcosahedronBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: ibg.Call("translate", x, y, z)}
}
func (ibg *IcosahedronBufferGeometry) UpdateFromObject(object *core.Object3D) {
	ibg.Call("updateFromObject", object)
}

type IcosahedronGeometry struct {
	js.Value
}

func NewIcosahedronGeometry(radius int, detail int) *IcosahedronGeometry {
	return &IcosahedronGeometry{Value: get("IcosahedronGeometry").New(radius, detail)}
}
func (ig *IcosahedronGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: ig.Get("animation")}
}
func (ig *IcosahedronGeometry) SetAnimation(v *animation.AnimationClip) {
	ig.Set("animation", v)
}
func (ig *IcosahedronGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(ig.Get("animations"))
}
func (ig *IcosahedronGeometry) SetAnimations(v []animation.AnimationClip) {
	ig.Set("animations", v)
}
func (ig *IcosahedronGeometry) Bones() []objects.Bone {
	return []objects.Bone(ig.Get("bones"))
}
func (ig *IcosahedronGeometry) SetBones(v []objects.Bone) {
	ig.Set("bones", v)
}
func (ig *IcosahedronGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: ig.Get("boundingBox")}
}
func (ig *IcosahedronGeometry) SetBoundingBox(v *math.Box3) {
	ig.Set("boundingBox", v)
}
func (ig *IcosahedronGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: ig.Get("boundingSphere")}
}
func (ig *IcosahedronGeometry) SetBoundingSphere(v *math.Sphere) {
	ig.Set("boundingSphere", v)
}
func (ig *IcosahedronGeometry) Colors() []math.Color {
	return []math.Color(ig.Get("colors"))
}
func (ig *IcosahedronGeometry) SetColors(v []math.Color) {
	ig.Set("colors", v)
}
func (ig *IcosahedronGeometry) ColorsNeedUpdate() bool {
	return ig.Get("colorsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetColorsNeedUpdate(v bool) {
	ig.Set("colorsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) ElementsNeedUpdate() bool {
	return ig.Get("elementsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetElementsNeedUpdate(v bool) {
	ig.Set("elementsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(ig.Get("faceVertexUvs"))
}
func (ig *IcosahedronGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	ig.Set("faceVertexUvs", v)
}
func (ig *IcosahedronGeometry) Faces() []core.Face3 {
	return []core.Face3(ig.Get("faces"))
}
func (ig *IcosahedronGeometry) SetFaces(v []core.Face3) {
	ig.Set("faces", v)
}
func (ig *IcosahedronGeometry) GroupsNeedUpdate() bool {
	return ig.Get("groupsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetGroupsNeedUpdate(v bool) {
	ig.Set("groupsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) Id() int {
	return ig.Get("id").Int()
}
func (ig *IcosahedronGeometry) SetId(v int) {
	ig.Set("id", v)
}
func (ig *IcosahedronGeometry) LineDistances() []int {
	return []int(ig.Get("lineDistances"))
}
func (ig *IcosahedronGeometry) SetLineDistances(v []int) {
	ig.Set("lineDistances", v)
}
func (ig *IcosahedronGeometry) LineDistancesNeedUpdate() bool {
	return ig.Get("lineDistancesNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	ig.Set("lineDistancesNeedUpdate", v)
}
func (ig *IcosahedronGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(ig.Get("morphNormals"))
}
func (ig *IcosahedronGeometry) SetMorphNormals(v []core.MorphNormals) {
	ig.Set("morphNormals", v)
}
func (ig *IcosahedronGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(ig.Get("morphTargets"))
}
func (ig *IcosahedronGeometry) SetMorphTargets(v []core.MorphTarget) {
	ig.Set("morphTargets", v)
}
func (ig *IcosahedronGeometry) Name() string {
	return ig.Get("name").String()
}
func (ig *IcosahedronGeometry) SetName(v string) {
	ig.Set("name", v)
}
func (ig *IcosahedronGeometry) NormalsNeedUpdate() bool {
	return ig.Get("normalsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetNormalsNeedUpdate(v bool) {
	ig.Set("normalsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) Parameters() js.Value {
	return ig.Get("parameters")
}
func (ig *IcosahedronGeometry) SetParameters(v js.Value) {
	ig.Set("parameters", v)
}
func (ig *IcosahedronGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(ig.Get("skinIndices"))
}
func (ig *IcosahedronGeometry) SetSkinIndices(v []math.Vector4) {
	ig.Set("skinIndices", v)
}
func (ig *IcosahedronGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(ig.Get("skinWeights"))
}
func (ig *IcosahedronGeometry) SetSkinWeights(v []math.Vector4) {
	ig.Set("skinWeights", v)
}
func (ig *IcosahedronGeometry) Type() string {
	return ig.Get("type").String()
}
func (ig *IcosahedronGeometry) SetType(v string) {
	ig.Set("type", v)
}
func (ig *IcosahedronGeometry) Uuid() string {
	return ig.Get("uuid").String()
}
func (ig *IcosahedronGeometry) SetUuid(v string) {
	ig.Set("uuid", v)
}
func (ig *IcosahedronGeometry) UvsNeedUpdate() bool {
	return ig.Get("uvsNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetUvsNeedUpdate(v bool) {
	ig.Set("uvsNeedUpdate", v)
}
func (ig *IcosahedronGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(ig.Get("vertices"))
}
func (ig *IcosahedronGeometry) SetVertices(v []math.Vector3) {
	ig.Set("vertices", v)
}
func (ig *IcosahedronGeometry) VerticesNeedUpdate() bool {
	return ig.Get("verticesNeedUpdate").Bool()
}
func (ig *IcosahedronGeometry) SetVerticesNeedUpdate(v bool) {
	ig.Set("verticesNeedUpdate", v)
}
func (ig *IcosahedronGeometry) AddEventListener(typ string, listener js.Value) {
	ig.Call("addEventListener", typ, listener)
}
func (ig *IcosahedronGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: ig.Call("applyMatrix", matrix)}
}
func (ig *IcosahedronGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: ig.Call("center")}
}
func (ig *IcosahedronGeometry) Clone() this {
	return this(ig.Call("clone"))
}
func (ig *IcosahedronGeometry) ComputeBoundingBox() {
	ig.Call("computeBoundingBox")
}
func (ig *IcosahedronGeometry) ComputeBoundingSphere() {
	ig.Call("computeBoundingSphere")
}
func (ig *IcosahedronGeometry) ComputeFaceNormals() {
	ig.Call("computeFaceNormals")
}
func (ig *IcosahedronGeometry) ComputeFlatVertexNormals() {
	ig.Call("computeFlatVertexNormals")
}
func (ig *IcosahedronGeometry) ComputeMorphNormals() {
	ig.Call("computeMorphNormals")
}
func (ig *IcosahedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	ig.Call("computeVertexNormals", areaWeighted)
}
func (ig *IcosahedronGeometry) Copy(source *core.Geometry) this {
	return this(ig.Call("copy", source))
}
func (ig *IcosahedronGeometry) DispatchEvent(event js.Value) {
	ig.Call("dispatchEvent", event)
}
func (ig *IcosahedronGeometry) Dispose() {
	ig.Call("dispose")
}
func (ig *IcosahedronGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: ig.Call("fromBufferGeometry", geometry)}
}
func (ig *IcosahedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ig.Call("hasEventListener", typ, listener).Bool()
}
func (ig *IcosahedronGeometry) LookAt(vector *math.Vector3) {
	ig.Call("lookAt", vector)
}
func (ig *IcosahedronGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	ig.Call("merge", geometry, matrix, materialIndexOffset)
}
func (ig *IcosahedronGeometry) MergeMesh(mesh *objects.Mesh) {
	ig.Call("mergeMesh", mesh)
}
func (ig *IcosahedronGeometry) MergeVertices() int {
	return ig.Call("mergeVertices").Int()
}
func (ig *IcosahedronGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: ig.Call("normalize")}
}
func (ig *IcosahedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	ig.Call("removeEventListener", typ, listener)
}
func (ig *IcosahedronGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: ig.Call("rotateX", angle)}
}
func (ig *IcosahedronGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: ig.Call("rotateY", angle)}
}
func (ig *IcosahedronGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: ig.Call("rotateZ", angle)}
}
func (ig *IcosahedronGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: ig.Call("scale", x, y, z)}
}
func (ig *IcosahedronGeometry) SetFromPoints(points Array) this {
	return this(ig.Call("setFromPoints", points))
}
func (ig *IcosahedronGeometry) SortFacesByMaterialIndex() {
	ig.Call("sortFacesByMaterialIndex")
}
func (ig *IcosahedronGeometry) ToJSON() js.Value {
	return ig.Call("toJSON")
}
func (ig *IcosahedronGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: ig.Call("translate", x, y, z)}
}

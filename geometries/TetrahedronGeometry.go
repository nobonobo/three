package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type TetrahedronBufferGeometry struct {
	js.Value
}

func NewTetrahedronBufferGeometry(radius int, detail int) *TetrahedronBufferGeometry {
	return &TetrahedronBufferGeometry{Value: get("TetrahedronBufferGeometry").New(radius, detail)}
}
func (tbg *TetrahedronBufferGeometry) Attributes() js.Value {
	return tbg.Get("attributes")
}
func (tbg *TetrahedronBufferGeometry) SetAttributes(v js.Value) {
	tbg.Set("attributes", v)
}
func (tbg *TetrahedronBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: tbg.Get("boundingBox")}
}
func (tbg *TetrahedronBufferGeometry) SetBoundingBox(v *math.Box3) {
	tbg.Set("boundingBox", v)
}
func (tbg *TetrahedronBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: tbg.Get("boundingSphere")}
}
func (tbg *TetrahedronBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	tbg.Set("boundingSphere", v)
}
func (tbg *TetrahedronBufferGeometry) DrawRange() js.Value {
	return tbg.Get("drawRange")
}
func (tbg *TetrahedronBufferGeometry) SetDrawRange(v js.Value) {
	tbg.Set("drawRange", v)
}
func (tbg *TetrahedronBufferGeometry) Drawcalls() js.Value {
	return tbg.Get("drawcalls")
}
func (tbg *TetrahedronBufferGeometry) SetDrawcalls(v js.Value) {
	tbg.Set("drawcalls", v)
}
func (tbg *TetrahedronBufferGeometry) Groups() []js.Value {
	return []js.Value(tbg.Get("groups"))
}
func (tbg *TetrahedronBufferGeometry) SetGroups(v []js.Value) {
	tbg.Set("groups", v)
}
func (tbg *TetrahedronBufferGeometry) Id() int {
	return tbg.Get("id").Int()
}
func (tbg *TetrahedronBufferGeometry) SetId(v int) {
	tbg.Set("id", v)
}
func (tbg *TetrahedronBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: tbg.Get("index")}
}
func (tbg *TetrahedronBufferGeometry) SetIndex(v *core.BufferAttribute) {
	tbg.Set("index", v)
}
func (tbg *TetrahedronBufferGeometry) MorphAttributes() js.Value {
	return tbg.Get("morphAttributes")
}
func (tbg *TetrahedronBufferGeometry) SetMorphAttributes(v js.Value) {
	tbg.Set("morphAttributes", v)
}
func (tbg *TetrahedronBufferGeometry) Name() string {
	return tbg.Get("name").String()
}
func (tbg *TetrahedronBufferGeometry) SetName(v string) {
	tbg.Set("name", v)
}
func (tbg *TetrahedronBufferGeometry) Offsets() js.Value {
	return tbg.Get("offsets")
}
func (tbg *TetrahedronBufferGeometry) SetOffsets(v js.Value) {
	tbg.Set("offsets", v)
}
func (tbg *TetrahedronBufferGeometry) Parameters() js.Value {
	return tbg.Get("parameters")
}
func (tbg *TetrahedronBufferGeometry) SetParameters(v js.Value) {
	tbg.Set("parameters", v)
}
func (tbg *TetrahedronBufferGeometry) Type() string {
	return tbg.Get("type").String()
}
func (tbg *TetrahedronBufferGeometry) SetType(v string) {
	tbg.Set("type", v)
}
func (tbg *TetrahedronBufferGeometry) UserData() js.Value {
	return tbg.Get("userData")
}
func (tbg *TetrahedronBufferGeometry) SetUserData(v js.Value) {
	tbg.Set("userData", v)
}
func (tbg *TetrahedronBufferGeometry) Uuid() string {
	return tbg.Get("uuid").String()
}
func (tbg *TetrahedronBufferGeometry) SetUuid(v string) {
	tbg.Set("uuid", v)
}
func (tbg *TetrahedronBufferGeometry) MaxIndex() int {
	return tbg.Get("MaxIndex").Int()
}
func (tbg *TetrahedronBufferGeometry) SetMaxIndex(v int) {
	tbg.Set("MaxIndex", v)
}
func (tbg *TetrahedronBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("addAttribute", name, attribute)}
}
func (tbg *TetrahedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tbg.Call("addAttribute", name, array, itemSize)
}
func (tbg *TetrahedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tbg.Call("addDrawCall", start, count, indexOffset)
}
func (tbg *TetrahedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tbg.Call("addEventListener", typ, listener)
}
func (tbg *TetrahedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tbg.Call("addGroup", start, count, materialIndex)
}
func (tbg *TetrahedronBufferGeometry) AddIndex(index js.Value) {
	tbg.Call("addIndex", index)
}
func (tbg *TetrahedronBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("applyMatrix", matrix)}
}
func (tbg *TetrahedronBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("center")}
}
func (tbg *TetrahedronBufferGeometry) ClearDrawCalls() {
	tbg.Call("clearDrawCalls")
}
func (tbg *TetrahedronBufferGeometry) ClearGroups() {
	tbg.Call("clearGroups")
}
func (tbg *TetrahedronBufferGeometry) Clone() this {
	return this(tbg.Call("clone"))
}
func (tbg *TetrahedronBufferGeometry) ComputeBoundingBox() {
	tbg.Call("computeBoundingBox")
}
func (tbg *TetrahedronBufferGeometry) ComputeBoundingSphere() {
	tbg.Call("computeBoundingSphere")
}
func (tbg *TetrahedronBufferGeometry) ComputeVertexNormals() {
	tbg.Call("computeVertexNormals")
}
func (tbg *TetrahedronBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(tbg.Call("copy", source))
}
func (tbg *TetrahedronBufferGeometry) DispatchEvent(event js.Value) {
	tbg.Call("dispatchEvent", event)
}
func (tbg *TetrahedronBufferGeometry) Dispose() {
	tbg.Call("dispose")
}
func (tbg *TetrahedronBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("fromDirectGeometry", geometry)}
}
func (tbg *TetrahedronBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("fromGeometry", geometry, settings)}
}
func (tbg *TetrahedronBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(tbg.Call("getAttribute", name))
}
func (tbg *TetrahedronBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: tbg.Call("getIndex")}
}
func (tbg *TetrahedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tbg.Call("hasEventListener", typ, listener).Bool()
}
func (tbg *TetrahedronBufferGeometry) LookAt(v *math.Vector3) {
	tbg.Call("lookAt", v)
}
func (tbg *TetrahedronBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("merge", geometry, offset)}
}
func (tbg *TetrahedronBufferGeometry) NormalizeNormals() {
	tbg.Call("normalizeNormals")
}
func (tbg *TetrahedronBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("removeAttribute", name)}
}
func (tbg *TetrahedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tbg.Call("removeEventListener", typ, listener)
}
func (tbg *TetrahedronBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("rotateX", angle)}
}
func (tbg *TetrahedronBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("rotateY", angle)}
}
func (tbg *TetrahedronBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("rotateZ", angle)}
}
func (tbg *TetrahedronBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("scale", x, y, z)}
}
func (tbg *TetrahedronBufferGeometry) SetDrawRange(start int, count int) {
	tbg.Call("setDrawRange", start, count)
}
func (tbg *TetrahedronBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("setFromObject", object)}
}
func (tbg *TetrahedronBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("setFromPoints", points)}
}
func (tbg *TetrahedronBufferGeometry) SetIndex(index core.BufferAttribute) {
	tbg.Call("setIndex", index)
}
func (tbg *TetrahedronBufferGeometry) ToJSON() js.Value {
	return tbg.Call("toJSON")
}
func (tbg *TetrahedronBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("toNonIndexed")}
}
func (tbg *TetrahedronBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("translate", x, y, z)}
}
func (tbg *TetrahedronBufferGeometry) UpdateFromObject(object *core.Object3D) {
	tbg.Call("updateFromObject", object)
}

type TetrahedronGeometry struct {
	js.Value
}

func NewTetrahedronGeometry(radius int, detail int) *TetrahedronGeometry {
	return &TetrahedronGeometry{Value: get("TetrahedronGeometry").New(radius, detail)}
}
func (tg *TetrahedronGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: tg.Get("animation")}
}
func (tg *TetrahedronGeometry) SetAnimation(v *animation.AnimationClip) {
	tg.Set("animation", v)
}
func (tg *TetrahedronGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(tg.Get("animations"))
}
func (tg *TetrahedronGeometry) SetAnimations(v []animation.AnimationClip) {
	tg.Set("animations", v)
}
func (tg *TetrahedronGeometry) Bones() []objects.Bone {
	return []objects.Bone(tg.Get("bones"))
}
func (tg *TetrahedronGeometry) SetBones(v []objects.Bone) {
	tg.Set("bones", v)
}
func (tg *TetrahedronGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: tg.Get("boundingBox")}
}
func (tg *TetrahedronGeometry) SetBoundingBox(v *math.Box3) {
	tg.Set("boundingBox", v)
}
func (tg *TetrahedronGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: tg.Get("boundingSphere")}
}
func (tg *TetrahedronGeometry) SetBoundingSphere(v *math.Sphere) {
	tg.Set("boundingSphere", v)
}
func (tg *TetrahedronGeometry) Colors() []math.Color {
	return []math.Color(tg.Get("colors"))
}
func (tg *TetrahedronGeometry) SetColors(v []math.Color) {
	tg.Set("colors", v)
}
func (tg *TetrahedronGeometry) ColorsNeedUpdate() bool {
	return tg.Get("colorsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetColorsNeedUpdate(v bool) {
	tg.Set("colorsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) ElementsNeedUpdate() bool {
	return tg.Get("elementsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetElementsNeedUpdate(v bool) {
	tg.Set("elementsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(tg.Get("faceVertexUvs"))
}
func (tg *TetrahedronGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	tg.Set("faceVertexUvs", v)
}
func (tg *TetrahedronGeometry) Faces() []core.Face3 {
	return []core.Face3(tg.Get("faces"))
}
func (tg *TetrahedronGeometry) SetFaces(v []core.Face3) {
	tg.Set("faces", v)
}
func (tg *TetrahedronGeometry) GroupsNeedUpdate() bool {
	return tg.Get("groupsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetGroupsNeedUpdate(v bool) {
	tg.Set("groupsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) Id() int {
	return tg.Get("id").Int()
}
func (tg *TetrahedronGeometry) SetId(v int) {
	tg.Set("id", v)
}
func (tg *TetrahedronGeometry) LineDistances() []int {
	return []int(tg.Get("lineDistances"))
}
func (tg *TetrahedronGeometry) SetLineDistances(v []int) {
	tg.Set("lineDistances", v)
}
func (tg *TetrahedronGeometry) LineDistancesNeedUpdate() bool {
	return tg.Get("lineDistancesNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	tg.Set("lineDistancesNeedUpdate", v)
}
func (tg *TetrahedronGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(tg.Get("morphNormals"))
}
func (tg *TetrahedronGeometry) SetMorphNormals(v []core.MorphNormals) {
	tg.Set("morphNormals", v)
}
func (tg *TetrahedronGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(tg.Get("morphTargets"))
}
func (tg *TetrahedronGeometry) SetMorphTargets(v []core.MorphTarget) {
	tg.Set("morphTargets", v)
}
func (tg *TetrahedronGeometry) Name() string {
	return tg.Get("name").String()
}
func (tg *TetrahedronGeometry) SetName(v string) {
	tg.Set("name", v)
}
func (tg *TetrahedronGeometry) NormalsNeedUpdate() bool {
	return tg.Get("normalsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetNormalsNeedUpdate(v bool) {
	tg.Set("normalsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) Parameters() js.Value {
	return tg.Get("parameters")
}
func (tg *TetrahedronGeometry) SetParameters(v js.Value) {
	tg.Set("parameters", v)
}
func (tg *TetrahedronGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(tg.Get("skinIndices"))
}
func (tg *TetrahedronGeometry) SetSkinIndices(v []math.Vector4) {
	tg.Set("skinIndices", v)
}
func (tg *TetrahedronGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(tg.Get("skinWeights"))
}
func (tg *TetrahedronGeometry) SetSkinWeights(v []math.Vector4) {
	tg.Set("skinWeights", v)
}
func (tg *TetrahedronGeometry) Type() string {
	return tg.Get("type").String()
}
func (tg *TetrahedronGeometry) SetType(v string) {
	tg.Set("type", v)
}
func (tg *TetrahedronGeometry) Uuid() string {
	return tg.Get("uuid").String()
}
func (tg *TetrahedronGeometry) SetUuid(v string) {
	tg.Set("uuid", v)
}
func (tg *TetrahedronGeometry) UvsNeedUpdate() bool {
	return tg.Get("uvsNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetUvsNeedUpdate(v bool) {
	tg.Set("uvsNeedUpdate", v)
}
func (tg *TetrahedronGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(tg.Get("vertices"))
}
func (tg *TetrahedronGeometry) SetVertices(v []math.Vector3) {
	tg.Set("vertices", v)
}
func (tg *TetrahedronGeometry) VerticesNeedUpdate() bool {
	return tg.Get("verticesNeedUpdate").Bool()
}
func (tg *TetrahedronGeometry) SetVerticesNeedUpdate(v bool) {
	tg.Set("verticesNeedUpdate", v)
}
func (tg *TetrahedronGeometry) AddEventListener(typ string, listener js.Value) {
	tg.Call("addEventListener", typ, listener)
}
func (tg *TetrahedronGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: tg.Call("applyMatrix", matrix)}
}
func (tg *TetrahedronGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: tg.Call("center")}
}
func (tg *TetrahedronGeometry) Clone() this {
	return this(tg.Call("clone"))
}
func (tg *TetrahedronGeometry) ComputeBoundingBox() {
	tg.Call("computeBoundingBox")
}
func (tg *TetrahedronGeometry) ComputeBoundingSphere() {
	tg.Call("computeBoundingSphere")
}
func (tg *TetrahedronGeometry) ComputeFaceNormals() {
	tg.Call("computeFaceNormals")
}
func (tg *TetrahedronGeometry) ComputeFlatVertexNormals() {
	tg.Call("computeFlatVertexNormals")
}
func (tg *TetrahedronGeometry) ComputeMorphNormals() {
	tg.Call("computeMorphNormals")
}
func (tg *TetrahedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	tg.Call("computeVertexNormals", areaWeighted)
}
func (tg *TetrahedronGeometry) Copy(source *core.Geometry) this {
	return this(tg.Call("copy", source))
}
func (tg *TetrahedronGeometry) DispatchEvent(event js.Value) {
	tg.Call("dispatchEvent", event)
}
func (tg *TetrahedronGeometry) Dispose() {
	tg.Call("dispose")
}
func (tg *TetrahedronGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: tg.Call("fromBufferGeometry", geometry)}
}
func (tg *TetrahedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tg.Call("hasEventListener", typ, listener).Bool()
}
func (tg *TetrahedronGeometry) LookAt(vector *math.Vector3) {
	tg.Call("lookAt", vector)
}
func (tg *TetrahedronGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	tg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (tg *TetrahedronGeometry) MergeMesh(mesh *objects.Mesh) {
	tg.Call("mergeMesh", mesh)
}
func (tg *TetrahedronGeometry) MergeVertices() int {
	return tg.Call("mergeVertices").Int()
}
func (tg *TetrahedronGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: tg.Call("normalize")}
}
func (tg *TetrahedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	tg.Call("removeEventListener", typ, listener)
}
func (tg *TetrahedronGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("rotateX", angle)}
}
func (tg *TetrahedronGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("rotateY", angle)}
}
func (tg *TetrahedronGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("rotateZ", angle)}
}
func (tg *TetrahedronGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("scale", x, y, z)}
}
func (tg *TetrahedronGeometry) SetFromPoints(points Array) this {
	return this(tg.Call("setFromPoints", points))
}
func (tg *TetrahedronGeometry) SortFacesByMaterialIndex() {
	tg.Call("sortFacesByMaterialIndex")
}
func (tg *TetrahedronGeometry) ToJSON() js.Value {
	return tg.Call("toJSON")
}
func (tg *TetrahedronGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("translate", x, y, z)}
}

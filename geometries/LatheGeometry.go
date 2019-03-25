package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type LatheBufferGeometry struct {
	js.Value
}

func NewLatheBufferGeometry(points []math.Vector2, segments int, phiStart int, phiLength int) *LatheBufferGeometry {
	return &LatheBufferGeometry{Value: get("LatheBufferGeometry").New(points, segments, phiStart, phiLength)}
}
func (lbg *LatheBufferGeometry) Attributes() js.Value {
	return lbg.Get("attributes")
}
func (lbg *LatheBufferGeometry) SetAttributes(v js.Value) {
	lbg.Set("attributes", v)
}
func (lbg *LatheBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: lbg.Get("boundingBox")}
}
func (lbg *LatheBufferGeometry) SetBoundingBox(v *math.Box3) {
	lbg.Set("boundingBox", v)
}
func (lbg *LatheBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: lbg.Get("boundingSphere")}
}
func (lbg *LatheBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	lbg.Set("boundingSphere", v)
}
func (lbg *LatheBufferGeometry) DrawRange() js.Value {
	return lbg.Get("drawRange")
}
func (lbg *LatheBufferGeometry) SetDrawRange(v js.Value) {
	lbg.Set("drawRange", v)
}
func (lbg *LatheBufferGeometry) Drawcalls() js.Value {
	return lbg.Get("drawcalls")
}
func (lbg *LatheBufferGeometry) SetDrawcalls(v js.Value) {
	lbg.Set("drawcalls", v)
}
func (lbg *LatheBufferGeometry) Groups() []js.Value {
	return []js.Value(lbg.Get("groups"))
}
func (lbg *LatheBufferGeometry) SetGroups(v []js.Value) {
	lbg.Set("groups", v)
}
func (lbg *LatheBufferGeometry) Id() int {
	return lbg.Get("id").Int()
}
func (lbg *LatheBufferGeometry) SetId(v int) {
	lbg.Set("id", v)
}
func (lbg *LatheBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: lbg.Get("index")}
}
func (lbg *LatheBufferGeometry) SetIndex(v *core.BufferAttribute) {
	lbg.Set("index", v)
}
func (lbg *LatheBufferGeometry) MorphAttributes() js.Value {
	return lbg.Get("morphAttributes")
}
func (lbg *LatheBufferGeometry) SetMorphAttributes(v js.Value) {
	lbg.Set("morphAttributes", v)
}
func (lbg *LatheBufferGeometry) Name() string {
	return lbg.Get("name").String()
}
func (lbg *LatheBufferGeometry) SetName(v string) {
	lbg.Set("name", v)
}
func (lbg *LatheBufferGeometry) Offsets() js.Value {
	return lbg.Get("offsets")
}
func (lbg *LatheBufferGeometry) SetOffsets(v js.Value) {
	lbg.Set("offsets", v)
}
func (lbg *LatheBufferGeometry) Parameters() js.Value {
	return lbg.Get("parameters")
}
func (lbg *LatheBufferGeometry) SetParameters(v js.Value) {
	lbg.Set("parameters", v)
}
func (lbg *LatheBufferGeometry) Type() string {
	return lbg.Get("type").String()
}
func (lbg *LatheBufferGeometry) SetType(v string) {
	lbg.Set("type", v)
}
func (lbg *LatheBufferGeometry) UserData() js.Value {
	return lbg.Get("userData")
}
func (lbg *LatheBufferGeometry) SetUserData(v js.Value) {
	lbg.Set("userData", v)
}
func (lbg *LatheBufferGeometry) Uuid() string {
	return lbg.Get("uuid").String()
}
func (lbg *LatheBufferGeometry) SetUuid(v string) {
	lbg.Set("uuid", v)
}
func (lbg *LatheBufferGeometry) MaxIndex() int {
	return lbg.Get("MaxIndex").Int()
}
func (lbg *LatheBufferGeometry) SetMaxIndex(v int) {
	lbg.Set("MaxIndex", v)
}
func (lbg *LatheBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("addAttribute", name, attribute)}
}
func (lbg *LatheBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return lbg.Call("addAttribute", name, array, itemSize)
}
func (lbg *LatheBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	lbg.Call("addDrawCall", start, count, indexOffset)
}
func (lbg *LatheBufferGeometry) AddEventListener(typ string, listener js.Value) {
	lbg.Call("addEventListener", typ, listener)
}
func (lbg *LatheBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	lbg.Call("addGroup", start, count, materialIndex)
}
func (lbg *LatheBufferGeometry) AddIndex(index js.Value) {
	lbg.Call("addIndex", index)
}
func (lbg *LatheBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("applyMatrix", matrix)}
}
func (lbg *LatheBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("center")}
}
func (lbg *LatheBufferGeometry) ClearDrawCalls() {
	lbg.Call("clearDrawCalls")
}
func (lbg *LatheBufferGeometry) ClearGroups() {
	lbg.Call("clearGroups")
}
func (lbg *LatheBufferGeometry) Clone() this {
	return this(lbg.Call("clone"))
}
func (lbg *LatheBufferGeometry) ComputeBoundingBox() {
	lbg.Call("computeBoundingBox")
}
func (lbg *LatheBufferGeometry) ComputeBoundingSphere() {
	lbg.Call("computeBoundingSphere")
}
func (lbg *LatheBufferGeometry) ComputeVertexNormals() {
	lbg.Call("computeVertexNormals")
}
func (lbg *LatheBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(lbg.Call("copy", source))
}
func (lbg *LatheBufferGeometry) DispatchEvent(event js.Value) {
	lbg.Call("dispatchEvent", event)
}
func (lbg *LatheBufferGeometry) Dispose() {
	lbg.Call("dispose")
}
func (lbg *LatheBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("fromDirectGeometry", geometry)}
}
func (lbg *LatheBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("fromGeometry", geometry, settings)}
}
func (lbg *LatheBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(lbg.Call("getAttribute", name))
}
func (lbg *LatheBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: lbg.Call("getIndex")}
}
func (lbg *LatheBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return lbg.Call("hasEventListener", typ, listener).Bool()
}
func (lbg *LatheBufferGeometry) LookAt(v *math.Vector3) {
	lbg.Call("lookAt", v)
}
func (lbg *LatheBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("merge", geometry, offset)}
}
func (lbg *LatheBufferGeometry) NormalizeNormals() {
	lbg.Call("normalizeNormals")
}
func (lbg *LatheBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("removeAttribute", name)}
}
func (lbg *LatheBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	lbg.Call("removeEventListener", typ, listener)
}
func (lbg *LatheBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("rotateX", angle)}
}
func (lbg *LatheBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("rotateY", angle)}
}
func (lbg *LatheBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("rotateZ", angle)}
}
func (lbg *LatheBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("scale", x, y, z)}
}
func (lbg *LatheBufferGeometry) SetDrawRange(start int, count int) {
	lbg.Call("setDrawRange", start, count)
}
func (lbg *LatheBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("setFromObject", object)}
}
func (lbg *LatheBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("setFromPoints", points)}
}
func (lbg *LatheBufferGeometry) SetIndex(index core.BufferAttribute) {
	lbg.Call("setIndex", index)
}
func (lbg *LatheBufferGeometry) ToJSON() js.Value {
	return lbg.Call("toJSON")
}
func (lbg *LatheBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("toNonIndexed")}
}
func (lbg *LatheBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: lbg.Call("translate", x, y, z)}
}
func (lbg *LatheBufferGeometry) UpdateFromObject(object *core.Object3D) {
	lbg.Call("updateFromObject", object)
}

type LatheGeometry struct {
	js.Value
}

func NewLatheGeometry(points []math.Vector2, segments int, phiStart int, phiLength int) *LatheGeometry {
	return &LatheGeometry{Value: get("LatheGeometry").New(points, segments, phiStart, phiLength)}
}
func (lg *LatheGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: lg.Get("animation")}
}
func (lg *LatheGeometry) SetAnimation(v *animation.AnimationClip) {
	lg.Set("animation", v)
}
func (lg *LatheGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(lg.Get("animations"))
}
func (lg *LatheGeometry) SetAnimations(v []animation.AnimationClip) {
	lg.Set("animations", v)
}
func (lg *LatheGeometry) Bones() []objects.Bone {
	return []objects.Bone(lg.Get("bones"))
}
func (lg *LatheGeometry) SetBones(v []objects.Bone) {
	lg.Set("bones", v)
}
func (lg *LatheGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: lg.Get("boundingBox")}
}
func (lg *LatheGeometry) SetBoundingBox(v *math.Box3) {
	lg.Set("boundingBox", v)
}
func (lg *LatheGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: lg.Get("boundingSphere")}
}
func (lg *LatheGeometry) SetBoundingSphere(v *math.Sphere) {
	lg.Set("boundingSphere", v)
}
func (lg *LatheGeometry) Colors() []math.Color {
	return []math.Color(lg.Get("colors"))
}
func (lg *LatheGeometry) SetColors(v []math.Color) {
	lg.Set("colors", v)
}
func (lg *LatheGeometry) ColorsNeedUpdate() bool {
	return lg.Get("colorsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetColorsNeedUpdate(v bool) {
	lg.Set("colorsNeedUpdate", v)
}
func (lg *LatheGeometry) ElementsNeedUpdate() bool {
	return lg.Get("elementsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetElementsNeedUpdate(v bool) {
	lg.Set("elementsNeedUpdate", v)
}
func (lg *LatheGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(lg.Get("faceVertexUvs"))
}
func (lg *LatheGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	lg.Set("faceVertexUvs", v)
}
func (lg *LatheGeometry) Faces() []core.Face3 {
	return []core.Face3(lg.Get("faces"))
}
func (lg *LatheGeometry) SetFaces(v []core.Face3) {
	lg.Set("faces", v)
}
func (lg *LatheGeometry) GroupsNeedUpdate() bool {
	return lg.Get("groupsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetGroupsNeedUpdate(v bool) {
	lg.Set("groupsNeedUpdate", v)
}
func (lg *LatheGeometry) Id() int {
	return lg.Get("id").Int()
}
func (lg *LatheGeometry) SetId(v int) {
	lg.Set("id", v)
}
func (lg *LatheGeometry) LineDistances() []int {
	return []int(lg.Get("lineDistances"))
}
func (lg *LatheGeometry) SetLineDistances(v []int) {
	lg.Set("lineDistances", v)
}
func (lg *LatheGeometry) LineDistancesNeedUpdate() bool {
	return lg.Get("lineDistancesNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetLineDistancesNeedUpdate(v bool) {
	lg.Set("lineDistancesNeedUpdate", v)
}
func (lg *LatheGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(lg.Get("morphNormals"))
}
func (lg *LatheGeometry) SetMorphNormals(v []core.MorphNormals) {
	lg.Set("morphNormals", v)
}
func (lg *LatheGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(lg.Get("morphTargets"))
}
func (lg *LatheGeometry) SetMorphTargets(v []core.MorphTarget) {
	lg.Set("morphTargets", v)
}
func (lg *LatheGeometry) Name() string {
	return lg.Get("name").String()
}
func (lg *LatheGeometry) SetName(v string) {
	lg.Set("name", v)
}
func (lg *LatheGeometry) NormalsNeedUpdate() bool {
	return lg.Get("normalsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetNormalsNeedUpdate(v bool) {
	lg.Set("normalsNeedUpdate", v)
}
func (lg *LatheGeometry) Parameters() js.Value {
	return lg.Get("parameters")
}
func (lg *LatheGeometry) SetParameters(v js.Value) {
	lg.Set("parameters", v)
}
func (lg *LatheGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(lg.Get("skinIndices"))
}
func (lg *LatheGeometry) SetSkinIndices(v []math.Vector4) {
	lg.Set("skinIndices", v)
}
func (lg *LatheGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(lg.Get("skinWeights"))
}
func (lg *LatheGeometry) SetSkinWeights(v []math.Vector4) {
	lg.Set("skinWeights", v)
}
func (lg *LatheGeometry) Type() string {
	return lg.Get("type").String()
}
func (lg *LatheGeometry) SetType(v string) {
	lg.Set("type", v)
}
func (lg *LatheGeometry) Uuid() string {
	return lg.Get("uuid").String()
}
func (lg *LatheGeometry) SetUuid(v string) {
	lg.Set("uuid", v)
}
func (lg *LatheGeometry) UvsNeedUpdate() bool {
	return lg.Get("uvsNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetUvsNeedUpdate(v bool) {
	lg.Set("uvsNeedUpdate", v)
}
func (lg *LatheGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(lg.Get("vertices"))
}
func (lg *LatheGeometry) SetVertices(v []math.Vector3) {
	lg.Set("vertices", v)
}
func (lg *LatheGeometry) VerticesNeedUpdate() bool {
	return lg.Get("verticesNeedUpdate").Bool()
}
func (lg *LatheGeometry) SetVerticesNeedUpdate(v bool) {
	lg.Set("verticesNeedUpdate", v)
}
func (lg *LatheGeometry) AddEventListener(typ string, listener js.Value) {
	lg.Call("addEventListener", typ, listener)
}
func (lg *LatheGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: lg.Call("applyMatrix", matrix)}
}
func (lg *LatheGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: lg.Call("center")}
}
func (lg *LatheGeometry) Clone() this {
	return this(lg.Call("clone"))
}
func (lg *LatheGeometry) ComputeBoundingBox() {
	lg.Call("computeBoundingBox")
}
func (lg *LatheGeometry) ComputeBoundingSphere() {
	lg.Call("computeBoundingSphere")
}
func (lg *LatheGeometry) ComputeFaceNormals() {
	lg.Call("computeFaceNormals")
}
func (lg *LatheGeometry) ComputeFlatVertexNormals() {
	lg.Call("computeFlatVertexNormals")
}
func (lg *LatheGeometry) ComputeMorphNormals() {
	lg.Call("computeMorphNormals")
}
func (lg *LatheGeometry) ComputeVertexNormals(areaWeighted bool) {
	lg.Call("computeVertexNormals", areaWeighted)
}
func (lg *LatheGeometry) Copy(source *core.Geometry) this {
	return this(lg.Call("copy", source))
}
func (lg *LatheGeometry) DispatchEvent(event js.Value) {
	lg.Call("dispatchEvent", event)
}
func (lg *LatheGeometry) Dispose() {
	lg.Call("dispose")
}
func (lg *LatheGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: lg.Call("fromBufferGeometry", geometry)}
}
func (lg *LatheGeometry) HasEventListener(typ string, listener js.Value) bool {
	return lg.Call("hasEventListener", typ, listener).Bool()
}
func (lg *LatheGeometry) LookAt(vector *math.Vector3) {
	lg.Call("lookAt", vector)
}
func (lg *LatheGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	lg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (lg *LatheGeometry) MergeMesh(mesh *objects.Mesh) {
	lg.Call("mergeMesh", mesh)
}
func (lg *LatheGeometry) MergeVertices() int {
	return lg.Call("mergeVertices").Int()
}
func (lg *LatheGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: lg.Call("normalize")}
}
func (lg *LatheGeometry) RemoveEventListener(typ string, listener js.Value) {
	lg.Call("removeEventListener", typ, listener)
}
func (lg *LatheGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: lg.Call("rotateX", angle)}
}
func (lg *LatheGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: lg.Call("rotateY", angle)}
}
func (lg *LatheGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: lg.Call("rotateZ", angle)}
}
func (lg *LatheGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: lg.Call("scale", x, y, z)}
}
func (lg *LatheGeometry) SetFromPoints(points Array) this {
	return this(lg.Call("setFromPoints", points))
}
func (lg *LatheGeometry) SortFacesByMaterialIndex() {
	lg.Call("sortFacesByMaterialIndex")
}
func (lg *LatheGeometry) ToJSON() js.Value {
	return lg.Call("toJSON")
}
func (lg *LatheGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: lg.Call("translate", x, y, z)}
}

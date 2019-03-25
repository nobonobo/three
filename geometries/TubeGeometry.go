package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type TubeBufferGeometry struct {
	js.Value
}

func NewTubeBufferGeometry(path *core.Curve, tubularSegments int, radius int, radiusSegments int, closed bool) *TubeBufferGeometry {
	return &TubeBufferGeometry{Value: get("TubeBufferGeometry").New(path, tubularSegments, radius, radiusSegments, closed)}
}
func (tbg *TubeBufferGeometry) Attributes() js.Value {
	return tbg.Get("attributes")
}
func (tbg *TubeBufferGeometry) SetAttributes(v js.Value) {
	tbg.Set("attributes", v)
}
func (tbg *TubeBufferGeometry) Binormals() []math.Vector3 {
	return []math.Vector3(tbg.Get("binormals"))
}
func (tbg *TubeBufferGeometry) SetBinormals(v []math.Vector3) {
	tbg.Set("binormals", v)
}
func (tbg *TubeBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: tbg.Get("boundingBox")}
}
func (tbg *TubeBufferGeometry) SetBoundingBox(v *math.Box3) {
	tbg.Set("boundingBox", v)
}
func (tbg *TubeBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: tbg.Get("boundingSphere")}
}
func (tbg *TubeBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	tbg.Set("boundingSphere", v)
}
func (tbg *TubeBufferGeometry) DrawRange() js.Value {
	return tbg.Get("drawRange")
}
func (tbg *TubeBufferGeometry) SetDrawRange(v js.Value) {
	tbg.Set("drawRange", v)
}
func (tbg *TubeBufferGeometry) Drawcalls() js.Value {
	return tbg.Get("drawcalls")
}
func (tbg *TubeBufferGeometry) SetDrawcalls(v js.Value) {
	tbg.Set("drawcalls", v)
}
func (tbg *TubeBufferGeometry) Groups() []js.Value {
	return []js.Value(tbg.Get("groups"))
}
func (tbg *TubeBufferGeometry) SetGroups(v []js.Value) {
	tbg.Set("groups", v)
}
func (tbg *TubeBufferGeometry) Id() int {
	return tbg.Get("id").Int()
}
func (tbg *TubeBufferGeometry) SetId(v int) {
	tbg.Set("id", v)
}
func (tbg *TubeBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: tbg.Get("index")}
}
func (tbg *TubeBufferGeometry) SetIndex(v *core.BufferAttribute) {
	tbg.Set("index", v)
}
func (tbg *TubeBufferGeometry) MorphAttributes() js.Value {
	return tbg.Get("morphAttributes")
}
func (tbg *TubeBufferGeometry) SetMorphAttributes(v js.Value) {
	tbg.Set("morphAttributes", v)
}
func (tbg *TubeBufferGeometry) Name() string {
	return tbg.Get("name").String()
}
func (tbg *TubeBufferGeometry) SetName(v string) {
	tbg.Set("name", v)
}
func (tbg *TubeBufferGeometry) Normals() []math.Vector3 {
	return []math.Vector3(tbg.Get("normals"))
}
func (tbg *TubeBufferGeometry) SetNormals(v []math.Vector3) {
	tbg.Set("normals", v)
}
func (tbg *TubeBufferGeometry) Offsets() js.Value {
	return tbg.Get("offsets")
}
func (tbg *TubeBufferGeometry) SetOffsets(v js.Value) {
	tbg.Set("offsets", v)
}
func (tbg *TubeBufferGeometry) Parameters() js.Value {
	return tbg.Get("parameters")
}
func (tbg *TubeBufferGeometry) SetParameters(v js.Value) {
	tbg.Set("parameters", v)
}
func (tbg *TubeBufferGeometry) Tangents() []math.Vector3 {
	return []math.Vector3(tbg.Get("tangents"))
}
func (tbg *TubeBufferGeometry) SetTangents(v []math.Vector3) {
	tbg.Set("tangents", v)
}
func (tbg *TubeBufferGeometry) Type() string {
	return tbg.Get("type").String()
}
func (tbg *TubeBufferGeometry) SetType(v string) {
	tbg.Set("type", v)
}
func (tbg *TubeBufferGeometry) UserData() js.Value {
	return tbg.Get("userData")
}
func (tbg *TubeBufferGeometry) SetUserData(v js.Value) {
	tbg.Set("userData", v)
}
func (tbg *TubeBufferGeometry) Uuid() string {
	return tbg.Get("uuid").String()
}
func (tbg *TubeBufferGeometry) SetUuid(v string) {
	tbg.Set("uuid", v)
}
func (tbg *TubeBufferGeometry) MaxIndex() int {
	return tbg.Get("MaxIndex").Int()
}
func (tbg *TubeBufferGeometry) SetMaxIndex(v int) {
	tbg.Set("MaxIndex", v)
}
func (tbg *TubeBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("addAttribute", name, attribute)}
}
func (tbg *TubeBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tbg.Call("addAttribute", name, array, itemSize)
}
func (tbg *TubeBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tbg.Call("addDrawCall", start, count, indexOffset)
}
func (tbg *TubeBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tbg.Call("addEventListener", typ, listener)
}
func (tbg *TubeBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tbg.Call("addGroup", start, count, materialIndex)
}
func (tbg *TubeBufferGeometry) AddIndex(index js.Value) {
	tbg.Call("addIndex", index)
}
func (tbg *TubeBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("applyMatrix", matrix)}
}
func (tbg *TubeBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("center")}
}
func (tbg *TubeBufferGeometry) ClearDrawCalls() {
	tbg.Call("clearDrawCalls")
}
func (tbg *TubeBufferGeometry) ClearGroups() {
	tbg.Call("clearGroups")
}
func (tbg *TubeBufferGeometry) Clone() this {
	return this(tbg.Call("clone"))
}
func (tbg *TubeBufferGeometry) ComputeBoundingBox() {
	tbg.Call("computeBoundingBox")
}
func (tbg *TubeBufferGeometry) ComputeBoundingSphere() {
	tbg.Call("computeBoundingSphere")
}
func (tbg *TubeBufferGeometry) ComputeVertexNormals() {
	tbg.Call("computeVertexNormals")
}
func (tbg *TubeBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(tbg.Call("copy", source))
}
func (tbg *TubeBufferGeometry) DispatchEvent(event js.Value) {
	tbg.Call("dispatchEvent", event)
}
func (tbg *TubeBufferGeometry) Dispose() {
	tbg.Call("dispose")
}
func (tbg *TubeBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("fromDirectGeometry", geometry)}
}
func (tbg *TubeBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("fromGeometry", geometry, settings)}
}
func (tbg *TubeBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(tbg.Call("getAttribute", name))
}
func (tbg *TubeBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: tbg.Call("getIndex")}
}
func (tbg *TubeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tbg.Call("hasEventListener", typ, listener).Bool()
}
func (tbg *TubeBufferGeometry) LookAt(v *math.Vector3) {
	tbg.Call("lookAt", v)
}
func (tbg *TubeBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("merge", geometry, offset)}
}
func (tbg *TubeBufferGeometry) NormalizeNormals() {
	tbg.Call("normalizeNormals")
}
func (tbg *TubeBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("removeAttribute", name)}
}
func (tbg *TubeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tbg.Call("removeEventListener", typ, listener)
}
func (tbg *TubeBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("rotateX", angle)}
}
func (tbg *TubeBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("rotateY", angle)}
}
func (tbg *TubeBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("rotateZ", angle)}
}
func (tbg *TubeBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("scale", x, y, z)}
}
func (tbg *TubeBufferGeometry) SetDrawRange(start int, count int) {
	tbg.Call("setDrawRange", start, count)
}
func (tbg *TubeBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("setFromObject", object)}
}
func (tbg *TubeBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("setFromPoints", points)}
}
func (tbg *TubeBufferGeometry) SetIndex(index core.BufferAttribute) {
	tbg.Call("setIndex", index)
}
func (tbg *TubeBufferGeometry) ToJSON() js.Value {
	return tbg.Call("toJSON")
}
func (tbg *TubeBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("toNonIndexed")}
}
func (tbg *TubeBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tbg.Call("translate", x, y, z)}
}
func (tbg *TubeBufferGeometry) UpdateFromObject(object *core.Object3D) {
	tbg.Call("updateFromObject", object)
}

type TubeGeometry struct {
	js.Value
}

func NewTubeGeometry(path *core.Curve, tubularSegments int, radius int, radiusSegments int, closed bool) *TubeGeometry {
	return &TubeGeometry{Value: get("TubeGeometry").New(path, tubularSegments, radius, radiusSegments, closed)}
}
func (tg *TubeGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: tg.Get("animation")}
}
func (tg *TubeGeometry) SetAnimation(v *animation.AnimationClip) {
	tg.Set("animation", v)
}
func (tg *TubeGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(tg.Get("animations"))
}
func (tg *TubeGeometry) SetAnimations(v []animation.AnimationClip) {
	tg.Set("animations", v)
}
func (tg *TubeGeometry) Binormals() []math.Vector3 {
	return []math.Vector3(tg.Get("binormals"))
}
func (tg *TubeGeometry) SetBinormals(v []math.Vector3) {
	tg.Set("binormals", v)
}
func (tg *TubeGeometry) Bones() []objects.Bone {
	return []objects.Bone(tg.Get("bones"))
}
func (tg *TubeGeometry) SetBones(v []objects.Bone) {
	tg.Set("bones", v)
}
func (tg *TubeGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: tg.Get("boundingBox")}
}
func (tg *TubeGeometry) SetBoundingBox(v *math.Box3) {
	tg.Set("boundingBox", v)
}
func (tg *TubeGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: tg.Get("boundingSphere")}
}
func (tg *TubeGeometry) SetBoundingSphere(v *math.Sphere) {
	tg.Set("boundingSphere", v)
}
func (tg *TubeGeometry) Colors() []math.Color {
	return []math.Color(tg.Get("colors"))
}
func (tg *TubeGeometry) SetColors(v []math.Color) {
	tg.Set("colors", v)
}
func (tg *TubeGeometry) ColorsNeedUpdate() bool {
	return tg.Get("colorsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetColorsNeedUpdate(v bool) {
	tg.Set("colorsNeedUpdate", v)
}
func (tg *TubeGeometry) ElementsNeedUpdate() bool {
	return tg.Get("elementsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetElementsNeedUpdate(v bool) {
	tg.Set("elementsNeedUpdate", v)
}
func (tg *TubeGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(tg.Get("faceVertexUvs"))
}
func (tg *TubeGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	tg.Set("faceVertexUvs", v)
}
func (tg *TubeGeometry) Faces() []core.Face3 {
	return []core.Face3(tg.Get("faces"))
}
func (tg *TubeGeometry) SetFaces(v []core.Face3) {
	tg.Set("faces", v)
}
func (tg *TubeGeometry) GroupsNeedUpdate() bool {
	return tg.Get("groupsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetGroupsNeedUpdate(v bool) {
	tg.Set("groupsNeedUpdate", v)
}
func (tg *TubeGeometry) Id() int {
	return tg.Get("id").Int()
}
func (tg *TubeGeometry) SetId(v int) {
	tg.Set("id", v)
}
func (tg *TubeGeometry) LineDistances() []int {
	return []int(tg.Get("lineDistances"))
}
func (tg *TubeGeometry) SetLineDistances(v []int) {
	tg.Set("lineDistances", v)
}
func (tg *TubeGeometry) LineDistancesNeedUpdate() bool {
	return tg.Get("lineDistancesNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetLineDistancesNeedUpdate(v bool) {
	tg.Set("lineDistancesNeedUpdate", v)
}
func (tg *TubeGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(tg.Get("morphNormals"))
}
func (tg *TubeGeometry) SetMorphNormals(v []core.MorphNormals) {
	tg.Set("morphNormals", v)
}
func (tg *TubeGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(tg.Get("morphTargets"))
}
func (tg *TubeGeometry) SetMorphTargets(v []core.MorphTarget) {
	tg.Set("morphTargets", v)
}
func (tg *TubeGeometry) Name() string {
	return tg.Get("name").String()
}
func (tg *TubeGeometry) SetName(v string) {
	tg.Set("name", v)
}
func (tg *TubeGeometry) Normals() []math.Vector3 {
	return []math.Vector3(tg.Get("normals"))
}
func (tg *TubeGeometry) SetNormals(v []math.Vector3) {
	tg.Set("normals", v)
}
func (tg *TubeGeometry) NormalsNeedUpdate() bool {
	return tg.Get("normalsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetNormalsNeedUpdate(v bool) {
	tg.Set("normalsNeedUpdate", v)
}
func (tg *TubeGeometry) Parameters() js.Value {
	return tg.Get("parameters")
}
func (tg *TubeGeometry) SetParameters(v js.Value) {
	tg.Set("parameters", v)
}
func (tg *TubeGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(tg.Get("skinIndices"))
}
func (tg *TubeGeometry) SetSkinIndices(v []math.Vector4) {
	tg.Set("skinIndices", v)
}
func (tg *TubeGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(tg.Get("skinWeights"))
}
func (tg *TubeGeometry) SetSkinWeights(v []math.Vector4) {
	tg.Set("skinWeights", v)
}
func (tg *TubeGeometry) Tangents() []math.Vector3 {
	return []math.Vector3(tg.Get("tangents"))
}
func (tg *TubeGeometry) SetTangents(v []math.Vector3) {
	tg.Set("tangents", v)
}
func (tg *TubeGeometry) Type() string {
	return tg.Get("type").String()
}
func (tg *TubeGeometry) SetType(v string) {
	tg.Set("type", v)
}
func (tg *TubeGeometry) Uuid() string {
	return tg.Get("uuid").String()
}
func (tg *TubeGeometry) SetUuid(v string) {
	tg.Set("uuid", v)
}
func (tg *TubeGeometry) UvsNeedUpdate() bool {
	return tg.Get("uvsNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetUvsNeedUpdate(v bool) {
	tg.Set("uvsNeedUpdate", v)
}
func (tg *TubeGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(tg.Get("vertices"))
}
func (tg *TubeGeometry) SetVertices(v []math.Vector3) {
	tg.Set("vertices", v)
}
func (tg *TubeGeometry) VerticesNeedUpdate() bool {
	return tg.Get("verticesNeedUpdate").Bool()
}
func (tg *TubeGeometry) SetVerticesNeedUpdate(v bool) {
	tg.Set("verticesNeedUpdate", v)
}
func (tg *TubeGeometry) AddEventListener(typ string, listener js.Value) {
	tg.Call("addEventListener", typ, listener)
}
func (tg *TubeGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: tg.Call("applyMatrix", matrix)}
}
func (tg *TubeGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: tg.Call("center")}
}
func (tg *TubeGeometry) Clone() this {
	return this(tg.Call("clone"))
}
func (tg *TubeGeometry) ComputeBoundingBox() {
	tg.Call("computeBoundingBox")
}
func (tg *TubeGeometry) ComputeBoundingSphere() {
	tg.Call("computeBoundingSphere")
}
func (tg *TubeGeometry) ComputeFaceNormals() {
	tg.Call("computeFaceNormals")
}
func (tg *TubeGeometry) ComputeFlatVertexNormals() {
	tg.Call("computeFlatVertexNormals")
}
func (tg *TubeGeometry) ComputeMorphNormals() {
	tg.Call("computeMorphNormals")
}
func (tg *TubeGeometry) ComputeVertexNormals(areaWeighted bool) {
	tg.Call("computeVertexNormals", areaWeighted)
}
func (tg *TubeGeometry) Copy(source *core.Geometry) this {
	return this(tg.Call("copy", source))
}
func (tg *TubeGeometry) DispatchEvent(event js.Value) {
	tg.Call("dispatchEvent", event)
}
func (tg *TubeGeometry) Dispose() {
	tg.Call("dispose")
}
func (tg *TubeGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: tg.Call("fromBufferGeometry", geometry)}
}
func (tg *TubeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tg.Call("hasEventListener", typ, listener).Bool()
}
func (tg *TubeGeometry) LookAt(vector *math.Vector3) {
	tg.Call("lookAt", vector)
}
func (tg *TubeGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	tg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (tg *TubeGeometry) MergeMesh(mesh *objects.Mesh) {
	tg.Call("mergeMesh", mesh)
}
func (tg *TubeGeometry) MergeVertices() int {
	return tg.Call("mergeVertices").Int()
}
func (tg *TubeGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: tg.Call("normalize")}
}
func (tg *TubeGeometry) RemoveEventListener(typ string, listener js.Value) {
	tg.Call("removeEventListener", typ, listener)
}
func (tg *TubeGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("rotateX", angle)}
}
func (tg *TubeGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("rotateY", angle)}
}
func (tg *TubeGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("rotateZ", angle)}
}
func (tg *TubeGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("scale", x, y, z)}
}
func (tg *TubeGeometry) SetFromPoints(points Array) this {
	return this(tg.Call("setFromPoints", points))
}
func (tg *TubeGeometry) SortFacesByMaterialIndex() {
	tg.Call("sortFacesByMaterialIndex")
}
func (tg *TubeGeometry) ToJSON() js.Value {
	return tg.Call("toJSON")
}
func (tg *TubeGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: tg.Call("translate", x, y, z)}
}

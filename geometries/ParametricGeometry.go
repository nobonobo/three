package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type ParametricBufferGeometry struct {
	js.Value
}

func NewParametricBufferGeometry(fn js.Value, slices int, stacks int) *ParametricBufferGeometry {
	return &ParametricBufferGeometry{Value: get("ParametricBufferGeometry").New(fn, slices, stacks)}
}
func (pbg *ParametricBufferGeometry) Attributes() js.Value {
	return pbg.Get("attributes")
}
func (pbg *ParametricBufferGeometry) SetAttributes(v js.Value) {
	pbg.Set("attributes", v)
}
func (pbg *ParametricBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: pbg.Get("boundingBox")}
}
func (pbg *ParametricBufferGeometry) SetBoundingBox(v *math.Box3) {
	pbg.Set("boundingBox", v)
}
func (pbg *ParametricBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: pbg.Get("boundingSphere")}
}
func (pbg *ParametricBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	pbg.Set("boundingSphere", v)
}
func (pbg *ParametricBufferGeometry) DrawRange() js.Value {
	return pbg.Get("drawRange")
}
func (pbg *ParametricBufferGeometry) SetDrawRange(v js.Value) {
	pbg.Set("drawRange", v)
}
func (pbg *ParametricBufferGeometry) Drawcalls() js.Value {
	return pbg.Get("drawcalls")
}
func (pbg *ParametricBufferGeometry) SetDrawcalls(v js.Value) {
	pbg.Set("drawcalls", v)
}
func (pbg *ParametricBufferGeometry) Groups() []js.Value {
	return []js.Value(pbg.Get("groups"))
}
func (pbg *ParametricBufferGeometry) SetGroups(v []js.Value) {
	pbg.Set("groups", v)
}
func (pbg *ParametricBufferGeometry) Id() int {
	return pbg.Get("id").Int()
}
func (pbg *ParametricBufferGeometry) SetId(v int) {
	pbg.Set("id", v)
}
func (pbg *ParametricBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: pbg.Get("index")}
}
func (pbg *ParametricBufferGeometry) SetIndex(v *core.BufferAttribute) {
	pbg.Set("index", v)
}
func (pbg *ParametricBufferGeometry) MorphAttributes() js.Value {
	return pbg.Get("morphAttributes")
}
func (pbg *ParametricBufferGeometry) SetMorphAttributes(v js.Value) {
	pbg.Set("morphAttributes", v)
}
func (pbg *ParametricBufferGeometry) Name() string {
	return pbg.Get("name").String()
}
func (pbg *ParametricBufferGeometry) SetName(v string) {
	pbg.Set("name", v)
}
func (pbg *ParametricBufferGeometry) Offsets() js.Value {
	return pbg.Get("offsets")
}
func (pbg *ParametricBufferGeometry) SetOffsets(v js.Value) {
	pbg.Set("offsets", v)
}
func (pbg *ParametricBufferGeometry) Parameters() js.Value {
	return pbg.Get("parameters")
}
func (pbg *ParametricBufferGeometry) SetParameters(v js.Value) {
	pbg.Set("parameters", v)
}
func (pbg *ParametricBufferGeometry) Type() string {
	return pbg.Get("type").String()
}
func (pbg *ParametricBufferGeometry) SetType(v string) {
	pbg.Set("type", v)
}
func (pbg *ParametricBufferGeometry) UserData() js.Value {
	return pbg.Get("userData")
}
func (pbg *ParametricBufferGeometry) SetUserData(v js.Value) {
	pbg.Set("userData", v)
}
func (pbg *ParametricBufferGeometry) Uuid() string {
	return pbg.Get("uuid").String()
}
func (pbg *ParametricBufferGeometry) SetUuid(v string) {
	pbg.Set("uuid", v)
}
func (pbg *ParametricBufferGeometry) MaxIndex() int {
	return pbg.Get("MaxIndex").Int()
}
func (pbg *ParametricBufferGeometry) SetMaxIndex(v int) {
	pbg.Set("MaxIndex", v)
}
func (pbg *ParametricBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("addAttribute", name, attribute)}
}
func (pbg *ParametricBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return pbg.Call("addAttribute", name, array, itemSize)
}
func (pbg *ParametricBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	pbg.Call("addDrawCall", start, count, indexOffset)
}
func (pbg *ParametricBufferGeometry) AddEventListener(typ string, listener js.Value) {
	pbg.Call("addEventListener", typ, listener)
}
func (pbg *ParametricBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	pbg.Call("addGroup", start, count, materialIndex)
}
func (pbg *ParametricBufferGeometry) AddIndex(index js.Value) {
	pbg.Call("addIndex", index)
}
func (pbg *ParametricBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("applyMatrix", matrix)}
}
func (pbg *ParametricBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("center")}
}
func (pbg *ParametricBufferGeometry) ClearDrawCalls() {
	pbg.Call("clearDrawCalls")
}
func (pbg *ParametricBufferGeometry) ClearGroups() {
	pbg.Call("clearGroups")
}
func (pbg *ParametricBufferGeometry) Clone() this {
	return this(pbg.Call("clone"))
}
func (pbg *ParametricBufferGeometry) ComputeBoundingBox() {
	pbg.Call("computeBoundingBox")
}
func (pbg *ParametricBufferGeometry) ComputeBoundingSphere() {
	pbg.Call("computeBoundingSphere")
}
func (pbg *ParametricBufferGeometry) ComputeVertexNormals() {
	pbg.Call("computeVertexNormals")
}
func (pbg *ParametricBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(pbg.Call("copy", source))
}
func (pbg *ParametricBufferGeometry) DispatchEvent(event js.Value) {
	pbg.Call("dispatchEvent", event)
}
func (pbg *ParametricBufferGeometry) Dispose() {
	pbg.Call("dispose")
}
func (pbg *ParametricBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("fromDirectGeometry", geometry)}
}
func (pbg *ParametricBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("fromGeometry", geometry, settings)}
}
func (pbg *ParametricBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(pbg.Call("getAttribute", name))
}
func (pbg *ParametricBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: pbg.Call("getIndex")}
}
func (pbg *ParametricBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pbg.Call("hasEventListener", typ, listener).Bool()
}
func (pbg *ParametricBufferGeometry) LookAt(v *math.Vector3) {
	pbg.Call("lookAt", v)
}
func (pbg *ParametricBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("merge", geometry, offset)}
}
func (pbg *ParametricBufferGeometry) NormalizeNormals() {
	pbg.Call("normalizeNormals")
}
func (pbg *ParametricBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("removeAttribute", name)}
}
func (pbg *ParametricBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	pbg.Call("removeEventListener", typ, listener)
}
func (pbg *ParametricBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateX", angle)}
}
func (pbg *ParametricBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateY", angle)}
}
func (pbg *ParametricBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateZ", angle)}
}
func (pbg *ParametricBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("scale", x, y, z)}
}
func (pbg *ParametricBufferGeometry) SetDrawRange(start int, count int) {
	pbg.Call("setDrawRange", start, count)
}
func (pbg *ParametricBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("setFromObject", object)}
}
func (pbg *ParametricBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("setFromPoints", points)}
}
func (pbg *ParametricBufferGeometry) SetIndex(index core.BufferAttribute) {
	pbg.Call("setIndex", index)
}
func (pbg *ParametricBufferGeometry) ToJSON() js.Value {
	return pbg.Call("toJSON")
}
func (pbg *ParametricBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("toNonIndexed")}
}
func (pbg *ParametricBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("translate", x, y, z)}
}
func (pbg *ParametricBufferGeometry) UpdateFromObject(object *core.Object3D) {
	pbg.Call("updateFromObject", object)
}

type ParametricGeometry struct {
	js.Value
}

func NewParametricGeometry(fn js.Value, slices int, stacks int) *ParametricGeometry {
	return &ParametricGeometry{Value: get("ParametricGeometry").New(fn, slices, stacks)}
}
func (pg *ParametricGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: pg.Get("animation")}
}
func (pg *ParametricGeometry) SetAnimation(v *animation.AnimationClip) {
	pg.Set("animation", v)
}
func (pg *ParametricGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(pg.Get("animations"))
}
func (pg *ParametricGeometry) SetAnimations(v []animation.AnimationClip) {
	pg.Set("animations", v)
}
func (pg *ParametricGeometry) Bones() []objects.Bone {
	return []objects.Bone(pg.Get("bones"))
}
func (pg *ParametricGeometry) SetBones(v []objects.Bone) {
	pg.Set("bones", v)
}
func (pg *ParametricGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: pg.Get("boundingBox")}
}
func (pg *ParametricGeometry) SetBoundingBox(v *math.Box3) {
	pg.Set("boundingBox", v)
}
func (pg *ParametricGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: pg.Get("boundingSphere")}
}
func (pg *ParametricGeometry) SetBoundingSphere(v *math.Sphere) {
	pg.Set("boundingSphere", v)
}
func (pg *ParametricGeometry) Colors() []math.Color {
	return []math.Color(pg.Get("colors"))
}
func (pg *ParametricGeometry) SetColors(v []math.Color) {
	pg.Set("colors", v)
}
func (pg *ParametricGeometry) ColorsNeedUpdate() bool {
	return pg.Get("colorsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetColorsNeedUpdate(v bool) {
	pg.Set("colorsNeedUpdate", v)
}
func (pg *ParametricGeometry) ElementsNeedUpdate() bool {
	return pg.Get("elementsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetElementsNeedUpdate(v bool) {
	pg.Set("elementsNeedUpdate", v)
}
func (pg *ParametricGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(pg.Get("faceVertexUvs"))
}
func (pg *ParametricGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	pg.Set("faceVertexUvs", v)
}
func (pg *ParametricGeometry) Faces() []core.Face3 {
	return []core.Face3(pg.Get("faces"))
}
func (pg *ParametricGeometry) SetFaces(v []core.Face3) {
	pg.Set("faces", v)
}
func (pg *ParametricGeometry) GroupsNeedUpdate() bool {
	return pg.Get("groupsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetGroupsNeedUpdate(v bool) {
	pg.Set("groupsNeedUpdate", v)
}
func (pg *ParametricGeometry) Id() int {
	return pg.Get("id").Int()
}
func (pg *ParametricGeometry) SetId(v int) {
	pg.Set("id", v)
}
func (pg *ParametricGeometry) LineDistances() []int {
	return []int(pg.Get("lineDistances"))
}
func (pg *ParametricGeometry) SetLineDistances(v []int) {
	pg.Set("lineDistances", v)
}
func (pg *ParametricGeometry) LineDistancesNeedUpdate() bool {
	return pg.Get("lineDistancesNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetLineDistancesNeedUpdate(v bool) {
	pg.Set("lineDistancesNeedUpdate", v)
}
func (pg *ParametricGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(pg.Get("morphNormals"))
}
func (pg *ParametricGeometry) SetMorphNormals(v []core.MorphNormals) {
	pg.Set("morphNormals", v)
}
func (pg *ParametricGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(pg.Get("morphTargets"))
}
func (pg *ParametricGeometry) SetMorphTargets(v []core.MorphTarget) {
	pg.Set("morphTargets", v)
}
func (pg *ParametricGeometry) Name() string {
	return pg.Get("name").String()
}
func (pg *ParametricGeometry) SetName(v string) {
	pg.Set("name", v)
}
func (pg *ParametricGeometry) NormalsNeedUpdate() bool {
	return pg.Get("normalsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetNormalsNeedUpdate(v bool) {
	pg.Set("normalsNeedUpdate", v)
}
func (pg *ParametricGeometry) Parameters() js.Value {
	return pg.Get("parameters")
}
func (pg *ParametricGeometry) SetParameters(v js.Value) {
	pg.Set("parameters", v)
}
func (pg *ParametricGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(pg.Get("skinIndices"))
}
func (pg *ParametricGeometry) SetSkinIndices(v []math.Vector4) {
	pg.Set("skinIndices", v)
}
func (pg *ParametricGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(pg.Get("skinWeights"))
}
func (pg *ParametricGeometry) SetSkinWeights(v []math.Vector4) {
	pg.Set("skinWeights", v)
}
func (pg *ParametricGeometry) Type() string {
	return pg.Get("type").String()
}
func (pg *ParametricGeometry) SetType(v string) {
	pg.Set("type", v)
}
func (pg *ParametricGeometry) Uuid() string {
	return pg.Get("uuid").String()
}
func (pg *ParametricGeometry) SetUuid(v string) {
	pg.Set("uuid", v)
}
func (pg *ParametricGeometry) UvsNeedUpdate() bool {
	return pg.Get("uvsNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetUvsNeedUpdate(v bool) {
	pg.Set("uvsNeedUpdate", v)
}
func (pg *ParametricGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(pg.Get("vertices"))
}
func (pg *ParametricGeometry) SetVertices(v []math.Vector3) {
	pg.Set("vertices", v)
}
func (pg *ParametricGeometry) VerticesNeedUpdate() bool {
	return pg.Get("verticesNeedUpdate").Bool()
}
func (pg *ParametricGeometry) SetVerticesNeedUpdate(v bool) {
	pg.Set("verticesNeedUpdate", v)
}
func (pg *ParametricGeometry) AddEventListener(typ string, listener js.Value) {
	pg.Call("addEventListener", typ, listener)
}
func (pg *ParametricGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: pg.Call("applyMatrix", matrix)}
}
func (pg *ParametricGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: pg.Call("center")}
}
func (pg *ParametricGeometry) Clone() this {
	return this(pg.Call("clone"))
}
func (pg *ParametricGeometry) ComputeBoundingBox() {
	pg.Call("computeBoundingBox")
}
func (pg *ParametricGeometry) ComputeBoundingSphere() {
	pg.Call("computeBoundingSphere")
}
func (pg *ParametricGeometry) ComputeFaceNormals() {
	pg.Call("computeFaceNormals")
}
func (pg *ParametricGeometry) ComputeFlatVertexNormals() {
	pg.Call("computeFlatVertexNormals")
}
func (pg *ParametricGeometry) ComputeMorphNormals() {
	pg.Call("computeMorphNormals")
}
func (pg *ParametricGeometry) ComputeVertexNormals(areaWeighted bool) {
	pg.Call("computeVertexNormals", areaWeighted)
}
func (pg *ParametricGeometry) Copy(source *core.Geometry) this {
	return this(pg.Call("copy", source))
}
func (pg *ParametricGeometry) DispatchEvent(event js.Value) {
	pg.Call("dispatchEvent", event)
}
func (pg *ParametricGeometry) Dispose() {
	pg.Call("dispose")
}
func (pg *ParametricGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: pg.Call("fromBufferGeometry", geometry)}
}
func (pg *ParametricGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pg.Call("hasEventListener", typ, listener).Bool()
}
func (pg *ParametricGeometry) LookAt(vector *math.Vector3) {
	pg.Call("lookAt", vector)
}
func (pg *ParametricGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	pg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (pg *ParametricGeometry) MergeMesh(mesh *objects.Mesh) {
	pg.Call("mergeMesh", mesh)
}
func (pg *ParametricGeometry) MergeVertices() int {
	return pg.Call("mergeVertices").Int()
}
func (pg *ParametricGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: pg.Call("normalize")}
}
func (pg *ParametricGeometry) RemoveEventListener(typ string, listener js.Value) {
	pg.Call("removeEventListener", typ, listener)
}
func (pg *ParametricGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateX", angle)}
}
func (pg *ParametricGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateY", angle)}
}
func (pg *ParametricGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateZ", angle)}
}
func (pg *ParametricGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("scale", x, y, z)}
}
func (pg *ParametricGeometry) SetFromPoints(points Array) this {
	return this(pg.Call("setFromPoints", points))
}
func (pg *ParametricGeometry) SortFacesByMaterialIndex() {
	pg.Call("sortFacesByMaterialIndex")
}
func (pg *ParametricGeometry) ToJSON() js.Value {
	return pg.Call("toJSON")
}
func (pg *ParametricGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("translate", x, y, z)}
}

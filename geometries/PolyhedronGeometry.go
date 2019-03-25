package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type PolyhedronBufferGeometry struct {
	js.Value
}

func NewPolyhedronBufferGeometry(vertices []int, indices []int, radius int, detail int) *PolyhedronBufferGeometry {
	return &PolyhedronBufferGeometry{Value: get("PolyhedronBufferGeometry").New(vertices, indices, radius, detail)}
}
func (pbg *PolyhedronBufferGeometry) Attributes() js.Value {
	return pbg.Get("attributes")
}
func (pbg *PolyhedronBufferGeometry) SetAttributes(v js.Value) {
	pbg.Set("attributes", v)
}
func (pbg *PolyhedronBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: pbg.Get("boundingBox")}
}
func (pbg *PolyhedronBufferGeometry) SetBoundingBox(v *math.Box3) {
	pbg.Set("boundingBox", v)
}
func (pbg *PolyhedronBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: pbg.Get("boundingSphere")}
}
func (pbg *PolyhedronBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	pbg.Set("boundingSphere", v)
}
func (pbg *PolyhedronBufferGeometry) DrawRange() js.Value {
	return pbg.Get("drawRange")
}
func (pbg *PolyhedronBufferGeometry) SetDrawRange(v js.Value) {
	pbg.Set("drawRange", v)
}
func (pbg *PolyhedronBufferGeometry) Drawcalls() js.Value {
	return pbg.Get("drawcalls")
}
func (pbg *PolyhedronBufferGeometry) SetDrawcalls(v js.Value) {
	pbg.Set("drawcalls", v)
}
func (pbg *PolyhedronBufferGeometry) Groups() []js.Value {
	return []js.Value(pbg.Get("groups"))
}
func (pbg *PolyhedronBufferGeometry) SetGroups(v []js.Value) {
	pbg.Set("groups", v)
}
func (pbg *PolyhedronBufferGeometry) Id() int {
	return pbg.Get("id").Int()
}
func (pbg *PolyhedronBufferGeometry) SetId(v int) {
	pbg.Set("id", v)
}
func (pbg *PolyhedronBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: pbg.Get("index")}
}
func (pbg *PolyhedronBufferGeometry) SetIndex(v *core.BufferAttribute) {
	pbg.Set("index", v)
}
func (pbg *PolyhedronBufferGeometry) MorphAttributes() js.Value {
	return pbg.Get("morphAttributes")
}
func (pbg *PolyhedronBufferGeometry) SetMorphAttributes(v js.Value) {
	pbg.Set("morphAttributes", v)
}
func (pbg *PolyhedronBufferGeometry) Name() string {
	return pbg.Get("name").String()
}
func (pbg *PolyhedronBufferGeometry) SetName(v string) {
	pbg.Set("name", v)
}
func (pbg *PolyhedronBufferGeometry) Offsets() js.Value {
	return pbg.Get("offsets")
}
func (pbg *PolyhedronBufferGeometry) SetOffsets(v js.Value) {
	pbg.Set("offsets", v)
}
func (pbg *PolyhedronBufferGeometry) Parameters() js.Value {
	return pbg.Get("parameters")
}
func (pbg *PolyhedronBufferGeometry) SetParameters(v js.Value) {
	pbg.Set("parameters", v)
}
func (pbg *PolyhedronBufferGeometry) Type() string {
	return pbg.Get("type").String()
}
func (pbg *PolyhedronBufferGeometry) SetType(v string) {
	pbg.Set("type", v)
}
func (pbg *PolyhedronBufferGeometry) UserData() js.Value {
	return pbg.Get("userData")
}
func (pbg *PolyhedronBufferGeometry) SetUserData(v js.Value) {
	pbg.Set("userData", v)
}
func (pbg *PolyhedronBufferGeometry) Uuid() string {
	return pbg.Get("uuid").String()
}
func (pbg *PolyhedronBufferGeometry) SetUuid(v string) {
	pbg.Set("uuid", v)
}
func (pbg *PolyhedronBufferGeometry) MaxIndex() int {
	return pbg.Get("MaxIndex").Int()
}
func (pbg *PolyhedronBufferGeometry) SetMaxIndex(v int) {
	pbg.Set("MaxIndex", v)
}
func (pbg *PolyhedronBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("addAttribute", name, attribute)}
}
func (pbg *PolyhedronBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return pbg.Call("addAttribute", name, array, itemSize)
}
func (pbg *PolyhedronBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	pbg.Call("addDrawCall", start, count, indexOffset)
}
func (pbg *PolyhedronBufferGeometry) AddEventListener(typ string, listener js.Value) {
	pbg.Call("addEventListener", typ, listener)
}
func (pbg *PolyhedronBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	pbg.Call("addGroup", start, count, materialIndex)
}
func (pbg *PolyhedronBufferGeometry) AddIndex(index js.Value) {
	pbg.Call("addIndex", index)
}
func (pbg *PolyhedronBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("applyMatrix", matrix)}
}
func (pbg *PolyhedronBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("center")}
}
func (pbg *PolyhedronBufferGeometry) ClearDrawCalls() {
	pbg.Call("clearDrawCalls")
}
func (pbg *PolyhedronBufferGeometry) ClearGroups() {
	pbg.Call("clearGroups")
}
func (pbg *PolyhedronBufferGeometry) Clone() this {
	return this(pbg.Call("clone"))
}
func (pbg *PolyhedronBufferGeometry) ComputeBoundingBox() {
	pbg.Call("computeBoundingBox")
}
func (pbg *PolyhedronBufferGeometry) ComputeBoundingSphere() {
	pbg.Call("computeBoundingSphere")
}
func (pbg *PolyhedronBufferGeometry) ComputeVertexNormals() {
	pbg.Call("computeVertexNormals")
}
func (pbg *PolyhedronBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(pbg.Call("copy", source))
}
func (pbg *PolyhedronBufferGeometry) DispatchEvent(event js.Value) {
	pbg.Call("dispatchEvent", event)
}
func (pbg *PolyhedronBufferGeometry) Dispose() {
	pbg.Call("dispose")
}
func (pbg *PolyhedronBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("fromDirectGeometry", geometry)}
}
func (pbg *PolyhedronBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("fromGeometry", geometry, settings)}
}
func (pbg *PolyhedronBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(pbg.Call("getAttribute", name))
}
func (pbg *PolyhedronBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: pbg.Call("getIndex")}
}
func (pbg *PolyhedronBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pbg.Call("hasEventListener", typ, listener).Bool()
}
func (pbg *PolyhedronBufferGeometry) LookAt(v *math.Vector3) {
	pbg.Call("lookAt", v)
}
func (pbg *PolyhedronBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("merge", geometry, offset)}
}
func (pbg *PolyhedronBufferGeometry) NormalizeNormals() {
	pbg.Call("normalizeNormals")
}
func (pbg *PolyhedronBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("removeAttribute", name)}
}
func (pbg *PolyhedronBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	pbg.Call("removeEventListener", typ, listener)
}
func (pbg *PolyhedronBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateX", angle)}
}
func (pbg *PolyhedronBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateY", angle)}
}
func (pbg *PolyhedronBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateZ", angle)}
}
func (pbg *PolyhedronBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("scale", x, y, z)}
}
func (pbg *PolyhedronBufferGeometry) SetDrawRange(start int, count int) {
	pbg.Call("setDrawRange", start, count)
}
func (pbg *PolyhedronBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("setFromObject", object)}
}
func (pbg *PolyhedronBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("setFromPoints", points)}
}
func (pbg *PolyhedronBufferGeometry) SetIndex(index core.BufferAttribute) {
	pbg.Call("setIndex", index)
}
func (pbg *PolyhedronBufferGeometry) ToJSON() js.Value {
	return pbg.Call("toJSON")
}
func (pbg *PolyhedronBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("toNonIndexed")}
}
func (pbg *PolyhedronBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("translate", x, y, z)}
}
func (pbg *PolyhedronBufferGeometry) UpdateFromObject(object *core.Object3D) {
	pbg.Call("updateFromObject", object)
}

type PolyhedronGeometry struct {
	js.Value
}

func NewPolyhedronGeometry(vertices []int, indices []int, radius int, detail int) *PolyhedronGeometry {
	return &PolyhedronGeometry{Value: get("PolyhedronGeometry").New(vertices, indices, radius, detail)}
}
func (pg *PolyhedronGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: pg.Get("animation")}
}
func (pg *PolyhedronGeometry) SetAnimation(v *animation.AnimationClip) {
	pg.Set("animation", v)
}
func (pg *PolyhedronGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(pg.Get("animations"))
}
func (pg *PolyhedronGeometry) SetAnimations(v []animation.AnimationClip) {
	pg.Set("animations", v)
}
func (pg *PolyhedronGeometry) Bones() []objects.Bone {
	return []objects.Bone(pg.Get("bones"))
}
func (pg *PolyhedronGeometry) SetBones(v []objects.Bone) {
	pg.Set("bones", v)
}
func (pg *PolyhedronGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: pg.Get("boundingBox")}
}
func (pg *PolyhedronGeometry) SetBoundingBox(v *math.Box3) {
	pg.Set("boundingBox", v)
}
func (pg *PolyhedronGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: pg.Get("boundingSphere")}
}
func (pg *PolyhedronGeometry) SetBoundingSphere(v *math.Sphere) {
	pg.Set("boundingSphere", v)
}
func (pg *PolyhedronGeometry) Colors() []math.Color {
	return []math.Color(pg.Get("colors"))
}
func (pg *PolyhedronGeometry) SetColors(v []math.Color) {
	pg.Set("colors", v)
}
func (pg *PolyhedronGeometry) ColorsNeedUpdate() bool {
	return pg.Get("colorsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetColorsNeedUpdate(v bool) {
	pg.Set("colorsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) ElementsNeedUpdate() bool {
	return pg.Get("elementsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetElementsNeedUpdate(v bool) {
	pg.Set("elementsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(pg.Get("faceVertexUvs"))
}
func (pg *PolyhedronGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	pg.Set("faceVertexUvs", v)
}
func (pg *PolyhedronGeometry) Faces() []core.Face3 {
	return []core.Face3(pg.Get("faces"))
}
func (pg *PolyhedronGeometry) SetFaces(v []core.Face3) {
	pg.Set("faces", v)
}
func (pg *PolyhedronGeometry) GroupsNeedUpdate() bool {
	return pg.Get("groupsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetGroupsNeedUpdate(v bool) {
	pg.Set("groupsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) Id() int {
	return pg.Get("id").Int()
}
func (pg *PolyhedronGeometry) SetId(v int) {
	pg.Set("id", v)
}
func (pg *PolyhedronGeometry) LineDistances() []int {
	return []int(pg.Get("lineDistances"))
}
func (pg *PolyhedronGeometry) SetLineDistances(v []int) {
	pg.Set("lineDistances", v)
}
func (pg *PolyhedronGeometry) LineDistancesNeedUpdate() bool {
	return pg.Get("lineDistancesNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetLineDistancesNeedUpdate(v bool) {
	pg.Set("lineDistancesNeedUpdate", v)
}
func (pg *PolyhedronGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(pg.Get("morphNormals"))
}
func (pg *PolyhedronGeometry) SetMorphNormals(v []core.MorphNormals) {
	pg.Set("morphNormals", v)
}
func (pg *PolyhedronGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(pg.Get("morphTargets"))
}
func (pg *PolyhedronGeometry) SetMorphTargets(v []core.MorphTarget) {
	pg.Set("morphTargets", v)
}
func (pg *PolyhedronGeometry) Name() string {
	return pg.Get("name").String()
}
func (pg *PolyhedronGeometry) SetName(v string) {
	pg.Set("name", v)
}
func (pg *PolyhedronGeometry) NormalsNeedUpdate() bool {
	return pg.Get("normalsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetNormalsNeedUpdate(v bool) {
	pg.Set("normalsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) Parameters() js.Value {
	return pg.Get("parameters")
}
func (pg *PolyhedronGeometry) SetParameters(v js.Value) {
	pg.Set("parameters", v)
}
func (pg *PolyhedronGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(pg.Get("skinIndices"))
}
func (pg *PolyhedronGeometry) SetSkinIndices(v []math.Vector4) {
	pg.Set("skinIndices", v)
}
func (pg *PolyhedronGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(pg.Get("skinWeights"))
}
func (pg *PolyhedronGeometry) SetSkinWeights(v []math.Vector4) {
	pg.Set("skinWeights", v)
}
func (pg *PolyhedronGeometry) Type() string {
	return pg.Get("type").String()
}
func (pg *PolyhedronGeometry) SetType(v string) {
	pg.Set("type", v)
}
func (pg *PolyhedronGeometry) Uuid() string {
	return pg.Get("uuid").String()
}
func (pg *PolyhedronGeometry) SetUuid(v string) {
	pg.Set("uuid", v)
}
func (pg *PolyhedronGeometry) UvsNeedUpdate() bool {
	return pg.Get("uvsNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetUvsNeedUpdate(v bool) {
	pg.Set("uvsNeedUpdate", v)
}
func (pg *PolyhedronGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(pg.Get("vertices"))
}
func (pg *PolyhedronGeometry) SetVertices(v []math.Vector3) {
	pg.Set("vertices", v)
}
func (pg *PolyhedronGeometry) VerticesNeedUpdate() bool {
	return pg.Get("verticesNeedUpdate").Bool()
}
func (pg *PolyhedronGeometry) SetVerticesNeedUpdate(v bool) {
	pg.Set("verticesNeedUpdate", v)
}
func (pg *PolyhedronGeometry) AddEventListener(typ string, listener js.Value) {
	pg.Call("addEventListener", typ, listener)
}
func (pg *PolyhedronGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: pg.Call("applyMatrix", matrix)}
}
func (pg *PolyhedronGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: pg.Call("center")}
}
func (pg *PolyhedronGeometry) Clone() this {
	return this(pg.Call("clone"))
}
func (pg *PolyhedronGeometry) ComputeBoundingBox() {
	pg.Call("computeBoundingBox")
}
func (pg *PolyhedronGeometry) ComputeBoundingSphere() {
	pg.Call("computeBoundingSphere")
}
func (pg *PolyhedronGeometry) ComputeFaceNormals() {
	pg.Call("computeFaceNormals")
}
func (pg *PolyhedronGeometry) ComputeFlatVertexNormals() {
	pg.Call("computeFlatVertexNormals")
}
func (pg *PolyhedronGeometry) ComputeMorphNormals() {
	pg.Call("computeMorphNormals")
}
func (pg *PolyhedronGeometry) ComputeVertexNormals(areaWeighted bool) {
	pg.Call("computeVertexNormals", areaWeighted)
}
func (pg *PolyhedronGeometry) Copy(source *core.Geometry) this {
	return this(pg.Call("copy", source))
}
func (pg *PolyhedronGeometry) DispatchEvent(event js.Value) {
	pg.Call("dispatchEvent", event)
}
func (pg *PolyhedronGeometry) Dispose() {
	pg.Call("dispose")
}
func (pg *PolyhedronGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: pg.Call("fromBufferGeometry", geometry)}
}
func (pg *PolyhedronGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pg.Call("hasEventListener", typ, listener).Bool()
}
func (pg *PolyhedronGeometry) LookAt(vector *math.Vector3) {
	pg.Call("lookAt", vector)
}
func (pg *PolyhedronGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	pg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (pg *PolyhedronGeometry) MergeMesh(mesh *objects.Mesh) {
	pg.Call("mergeMesh", mesh)
}
func (pg *PolyhedronGeometry) MergeVertices() int {
	return pg.Call("mergeVertices").Int()
}
func (pg *PolyhedronGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: pg.Call("normalize")}
}
func (pg *PolyhedronGeometry) RemoveEventListener(typ string, listener js.Value) {
	pg.Call("removeEventListener", typ, listener)
}
func (pg *PolyhedronGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateX", angle)}
}
func (pg *PolyhedronGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateY", angle)}
}
func (pg *PolyhedronGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateZ", angle)}
}
func (pg *PolyhedronGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("scale", x, y, z)}
}
func (pg *PolyhedronGeometry) SetFromPoints(points Array) this {
	return this(pg.Call("setFromPoints", points))
}
func (pg *PolyhedronGeometry) SortFacesByMaterialIndex() {
	pg.Call("sortFacesByMaterialIndex")
}
func (pg *PolyhedronGeometry) ToJSON() js.Value {
	return pg.Call("toJSON")
}
func (pg *PolyhedronGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("translate", x, y, z)}
}

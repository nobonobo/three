package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type PlaneBufferGeometry struct {
	js.Value
}

func NewPlaneBufferGeometry(width int, height int, widthSegments int, heightSegments int) *PlaneBufferGeometry {
	return &PlaneBufferGeometry{Value: get("PlaneBufferGeometry").New(width, height, widthSegments, heightSegments)}
}
func (pbg *PlaneBufferGeometry) Attributes() js.Value {
	return pbg.Get("attributes")
}
func (pbg *PlaneBufferGeometry) SetAttributes(v js.Value) {
	pbg.Set("attributes", v)
}
func (pbg *PlaneBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: pbg.Get("boundingBox")}
}
func (pbg *PlaneBufferGeometry) SetBoundingBox(v *math.Box3) {
	pbg.Set("boundingBox", v)
}
func (pbg *PlaneBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: pbg.Get("boundingSphere")}
}
func (pbg *PlaneBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	pbg.Set("boundingSphere", v)
}
func (pbg *PlaneBufferGeometry) DrawRange() js.Value {
	return pbg.Get("drawRange")
}
func (pbg *PlaneBufferGeometry) SetDrawRange(v js.Value) {
	pbg.Set("drawRange", v)
}
func (pbg *PlaneBufferGeometry) Drawcalls() js.Value {
	return pbg.Get("drawcalls")
}
func (pbg *PlaneBufferGeometry) SetDrawcalls(v js.Value) {
	pbg.Set("drawcalls", v)
}
func (pbg *PlaneBufferGeometry) Groups() []js.Value {
	return []js.Value(pbg.Get("groups"))
}
func (pbg *PlaneBufferGeometry) SetGroups(v []js.Value) {
	pbg.Set("groups", v)
}
func (pbg *PlaneBufferGeometry) Id() int {
	return pbg.Get("id").Int()
}
func (pbg *PlaneBufferGeometry) SetId(v int) {
	pbg.Set("id", v)
}
func (pbg *PlaneBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: pbg.Get("index")}
}
func (pbg *PlaneBufferGeometry) SetIndex(v *core.BufferAttribute) {
	pbg.Set("index", v)
}
func (pbg *PlaneBufferGeometry) MorphAttributes() js.Value {
	return pbg.Get("morphAttributes")
}
func (pbg *PlaneBufferGeometry) SetMorphAttributes(v js.Value) {
	pbg.Set("morphAttributes", v)
}
func (pbg *PlaneBufferGeometry) Name() string {
	return pbg.Get("name").String()
}
func (pbg *PlaneBufferGeometry) SetName(v string) {
	pbg.Set("name", v)
}
func (pbg *PlaneBufferGeometry) Offsets() js.Value {
	return pbg.Get("offsets")
}
func (pbg *PlaneBufferGeometry) SetOffsets(v js.Value) {
	pbg.Set("offsets", v)
}
func (pbg *PlaneBufferGeometry) Parameters() js.Value {
	return pbg.Get("parameters")
}
func (pbg *PlaneBufferGeometry) SetParameters(v js.Value) {
	pbg.Set("parameters", v)
}
func (pbg *PlaneBufferGeometry) Type() string {
	return pbg.Get("type").String()
}
func (pbg *PlaneBufferGeometry) SetType(v string) {
	pbg.Set("type", v)
}
func (pbg *PlaneBufferGeometry) UserData() js.Value {
	return pbg.Get("userData")
}
func (pbg *PlaneBufferGeometry) SetUserData(v js.Value) {
	pbg.Set("userData", v)
}
func (pbg *PlaneBufferGeometry) Uuid() string {
	return pbg.Get("uuid").String()
}
func (pbg *PlaneBufferGeometry) SetUuid(v string) {
	pbg.Set("uuid", v)
}
func (pbg *PlaneBufferGeometry) MaxIndex() int {
	return pbg.Get("MaxIndex").Int()
}
func (pbg *PlaneBufferGeometry) SetMaxIndex(v int) {
	pbg.Set("MaxIndex", v)
}
func (pbg *PlaneBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("addAttribute", name, attribute)}
}
func (pbg *PlaneBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return pbg.Call("addAttribute", name, array, itemSize)
}
func (pbg *PlaneBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	pbg.Call("addDrawCall", start, count, indexOffset)
}
func (pbg *PlaneBufferGeometry) AddEventListener(typ string, listener js.Value) {
	pbg.Call("addEventListener", typ, listener)
}
func (pbg *PlaneBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	pbg.Call("addGroup", start, count, materialIndex)
}
func (pbg *PlaneBufferGeometry) AddIndex(index js.Value) {
	pbg.Call("addIndex", index)
}
func (pbg *PlaneBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("applyMatrix", matrix)}
}
func (pbg *PlaneBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("center")}
}
func (pbg *PlaneBufferGeometry) ClearDrawCalls() {
	pbg.Call("clearDrawCalls")
}
func (pbg *PlaneBufferGeometry) ClearGroups() {
	pbg.Call("clearGroups")
}
func (pbg *PlaneBufferGeometry) Clone() this {
	return this(pbg.Call("clone"))
}
func (pbg *PlaneBufferGeometry) ComputeBoundingBox() {
	pbg.Call("computeBoundingBox")
}
func (pbg *PlaneBufferGeometry) ComputeBoundingSphere() {
	pbg.Call("computeBoundingSphere")
}
func (pbg *PlaneBufferGeometry) ComputeVertexNormals() {
	pbg.Call("computeVertexNormals")
}
func (pbg *PlaneBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(pbg.Call("copy", source))
}
func (pbg *PlaneBufferGeometry) DispatchEvent(event js.Value) {
	pbg.Call("dispatchEvent", event)
}
func (pbg *PlaneBufferGeometry) Dispose() {
	pbg.Call("dispose")
}
func (pbg *PlaneBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("fromDirectGeometry", geometry)}
}
func (pbg *PlaneBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("fromGeometry", geometry, settings)}
}
func (pbg *PlaneBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(pbg.Call("getAttribute", name))
}
func (pbg *PlaneBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: pbg.Call("getIndex")}
}
func (pbg *PlaneBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pbg.Call("hasEventListener", typ, listener).Bool()
}
func (pbg *PlaneBufferGeometry) LookAt(v *math.Vector3) {
	pbg.Call("lookAt", v)
}
func (pbg *PlaneBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("merge", geometry, offset)}
}
func (pbg *PlaneBufferGeometry) NormalizeNormals() {
	pbg.Call("normalizeNormals")
}
func (pbg *PlaneBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("removeAttribute", name)}
}
func (pbg *PlaneBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	pbg.Call("removeEventListener", typ, listener)
}
func (pbg *PlaneBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateX", angle)}
}
func (pbg *PlaneBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateY", angle)}
}
func (pbg *PlaneBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("rotateZ", angle)}
}
func (pbg *PlaneBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("scale", x, y, z)}
}
func (pbg *PlaneBufferGeometry) SetDrawRange(start int, count int) {
	pbg.Call("setDrawRange", start, count)
}
func (pbg *PlaneBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("setFromObject", object)}
}
func (pbg *PlaneBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("setFromPoints", points)}
}
func (pbg *PlaneBufferGeometry) SetIndex(index core.BufferAttribute) {
	pbg.Call("setIndex", index)
}
func (pbg *PlaneBufferGeometry) ToJSON() js.Value {
	return pbg.Call("toJSON")
}
func (pbg *PlaneBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("toNonIndexed")}
}
func (pbg *PlaneBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: pbg.Call("translate", x, y, z)}
}
func (pbg *PlaneBufferGeometry) UpdateFromObject(object *core.Object3D) {
	pbg.Call("updateFromObject", object)
}

type PlaneGeometry struct {
	js.Value
}

func NewPlaneGeometry(width int, height int, widthSegments int, heightSegments int) *PlaneGeometry {
	return &PlaneGeometry{Value: get("PlaneGeometry").New(width, height, widthSegments, heightSegments)}
}
func (pg *PlaneGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: pg.Get("animation")}
}
func (pg *PlaneGeometry) SetAnimation(v *animation.AnimationClip) {
	pg.Set("animation", v)
}
func (pg *PlaneGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(pg.Get("animations"))
}
func (pg *PlaneGeometry) SetAnimations(v []animation.AnimationClip) {
	pg.Set("animations", v)
}
func (pg *PlaneGeometry) Bones() []objects.Bone {
	return []objects.Bone(pg.Get("bones"))
}
func (pg *PlaneGeometry) SetBones(v []objects.Bone) {
	pg.Set("bones", v)
}
func (pg *PlaneGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: pg.Get("boundingBox")}
}
func (pg *PlaneGeometry) SetBoundingBox(v *math.Box3) {
	pg.Set("boundingBox", v)
}
func (pg *PlaneGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: pg.Get("boundingSphere")}
}
func (pg *PlaneGeometry) SetBoundingSphere(v *math.Sphere) {
	pg.Set("boundingSphere", v)
}
func (pg *PlaneGeometry) Colors() []math.Color {
	return []math.Color(pg.Get("colors"))
}
func (pg *PlaneGeometry) SetColors(v []math.Color) {
	pg.Set("colors", v)
}
func (pg *PlaneGeometry) ColorsNeedUpdate() bool {
	return pg.Get("colorsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetColorsNeedUpdate(v bool) {
	pg.Set("colorsNeedUpdate", v)
}
func (pg *PlaneGeometry) ElementsNeedUpdate() bool {
	return pg.Get("elementsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetElementsNeedUpdate(v bool) {
	pg.Set("elementsNeedUpdate", v)
}
func (pg *PlaneGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(pg.Get("faceVertexUvs"))
}
func (pg *PlaneGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	pg.Set("faceVertexUvs", v)
}
func (pg *PlaneGeometry) Faces() []core.Face3 {
	return []core.Face3(pg.Get("faces"))
}
func (pg *PlaneGeometry) SetFaces(v []core.Face3) {
	pg.Set("faces", v)
}
func (pg *PlaneGeometry) GroupsNeedUpdate() bool {
	return pg.Get("groupsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetGroupsNeedUpdate(v bool) {
	pg.Set("groupsNeedUpdate", v)
}
func (pg *PlaneGeometry) Id() int {
	return pg.Get("id").Int()
}
func (pg *PlaneGeometry) SetId(v int) {
	pg.Set("id", v)
}
func (pg *PlaneGeometry) LineDistances() []int {
	return []int(pg.Get("lineDistances"))
}
func (pg *PlaneGeometry) SetLineDistances(v []int) {
	pg.Set("lineDistances", v)
}
func (pg *PlaneGeometry) LineDistancesNeedUpdate() bool {
	return pg.Get("lineDistancesNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetLineDistancesNeedUpdate(v bool) {
	pg.Set("lineDistancesNeedUpdate", v)
}
func (pg *PlaneGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(pg.Get("morphNormals"))
}
func (pg *PlaneGeometry) SetMorphNormals(v []core.MorphNormals) {
	pg.Set("morphNormals", v)
}
func (pg *PlaneGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(pg.Get("morphTargets"))
}
func (pg *PlaneGeometry) SetMorphTargets(v []core.MorphTarget) {
	pg.Set("morphTargets", v)
}
func (pg *PlaneGeometry) Name() string {
	return pg.Get("name").String()
}
func (pg *PlaneGeometry) SetName(v string) {
	pg.Set("name", v)
}
func (pg *PlaneGeometry) NormalsNeedUpdate() bool {
	return pg.Get("normalsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetNormalsNeedUpdate(v bool) {
	pg.Set("normalsNeedUpdate", v)
}
func (pg *PlaneGeometry) Parameters() js.Value {
	return pg.Get("parameters")
}
func (pg *PlaneGeometry) SetParameters(v js.Value) {
	pg.Set("parameters", v)
}
func (pg *PlaneGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(pg.Get("skinIndices"))
}
func (pg *PlaneGeometry) SetSkinIndices(v []math.Vector4) {
	pg.Set("skinIndices", v)
}
func (pg *PlaneGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(pg.Get("skinWeights"))
}
func (pg *PlaneGeometry) SetSkinWeights(v []math.Vector4) {
	pg.Set("skinWeights", v)
}
func (pg *PlaneGeometry) Type() string {
	return pg.Get("type").String()
}
func (pg *PlaneGeometry) SetType(v string) {
	pg.Set("type", v)
}
func (pg *PlaneGeometry) Uuid() string {
	return pg.Get("uuid").String()
}
func (pg *PlaneGeometry) SetUuid(v string) {
	pg.Set("uuid", v)
}
func (pg *PlaneGeometry) UvsNeedUpdate() bool {
	return pg.Get("uvsNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetUvsNeedUpdate(v bool) {
	pg.Set("uvsNeedUpdate", v)
}
func (pg *PlaneGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(pg.Get("vertices"))
}
func (pg *PlaneGeometry) SetVertices(v []math.Vector3) {
	pg.Set("vertices", v)
}
func (pg *PlaneGeometry) VerticesNeedUpdate() bool {
	return pg.Get("verticesNeedUpdate").Bool()
}
func (pg *PlaneGeometry) SetVerticesNeedUpdate(v bool) {
	pg.Set("verticesNeedUpdate", v)
}
func (pg *PlaneGeometry) AddEventListener(typ string, listener js.Value) {
	pg.Call("addEventListener", typ, listener)
}
func (pg *PlaneGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: pg.Call("applyMatrix", matrix)}
}
func (pg *PlaneGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: pg.Call("center")}
}
func (pg *PlaneGeometry) Clone() this {
	return this(pg.Call("clone"))
}
func (pg *PlaneGeometry) ComputeBoundingBox() {
	pg.Call("computeBoundingBox")
}
func (pg *PlaneGeometry) ComputeBoundingSphere() {
	pg.Call("computeBoundingSphere")
}
func (pg *PlaneGeometry) ComputeFaceNormals() {
	pg.Call("computeFaceNormals")
}
func (pg *PlaneGeometry) ComputeFlatVertexNormals() {
	pg.Call("computeFlatVertexNormals")
}
func (pg *PlaneGeometry) ComputeMorphNormals() {
	pg.Call("computeMorphNormals")
}
func (pg *PlaneGeometry) ComputeVertexNormals(areaWeighted bool) {
	pg.Call("computeVertexNormals", areaWeighted)
}
func (pg *PlaneGeometry) Copy(source *core.Geometry) this {
	return this(pg.Call("copy", source))
}
func (pg *PlaneGeometry) DispatchEvent(event js.Value) {
	pg.Call("dispatchEvent", event)
}
func (pg *PlaneGeometry) Dispose() {
	pg.Call("dispose")
}
func (pg *PlaneGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: pg.Call("fromBufferGeometry", geometry)}
}
func (pg *PlaneGeometry) HasEventListener(typ string, listener js.Value) bool {
	return pg.Call("hasEventListener", typ, listener).Bool()
}
func (pg *PlaneGeometry) LookAt(vector *math.Vector3) {
	pg.Call("lookAt", vector)
}
func (pg *PlaneGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	pg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (pg *PlaneGeometry) MergeMesh(mesh *objects.Mesh) {
	pg.Call("mergeMesh", mesh)
}
func (pg *PlaneGeometry) MergeVertices() int {
	return pg.Call("mergeVertices").Int()
}
func (pg *PlaneGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: pg.Call("normalize")}
}
func (pg *PlaneGeometry) RemoveEventListener(typ string, listener js.Value) {
	pg.Call("removeEventListener", typ, listener)
}
func (pg *PlaneGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateX", angle)}
}
func (pg *PlaneGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateY", angle)}
}
func (pg *PlaneGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("rotateZ", angle)}
}
func (pg *PlaneGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("scale", x, y, z)}
}
func (pg *PlaneGeometry) SetFromPoints(points Array) this {
	return this(pg.Call("setFromPoints", points))
}
func (pg *PlaneGeometry) SortFacesByMaterialIndex() {
	pg.Call("sortFacesByMaterialIndex")
}
func (pg *PlaneGeometry) ToJSON() js.Value {
	return pg.Call("toJSON")
}
func (pg *PlaneGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: pg.Call("translate", x, y, z)}
}

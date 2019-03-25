package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type ConeBufferGeometry struct {
	js.Value
}

func NewConeBufferGeometry(radius int, height int, radialSegment int, heightSegment int, openEnded bool, thetaStart int, thetaLength int) *ConeBufferGeometry {
	return &ConeBufferGeometry{Value: get("ConeBufferGeometry").New(radius, height, radialSegment, heightSegment, openEnded, thetaStart, thetaLength)}
}
func (cbg *ConeBufferGeometry) Attributes() js.Value {
	return cbg.Get("attributes")
}
func (cbg *ConeBufferGeometry) SetAttributes(v js.Value) {
	cbg.Set("attributes", v)
}
func (cbg *ConeBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: cbg.Get("boundingBox")}
}
func (cbg *ConeBufferGeometry) SetBoundingBox(v *math.Box3) {
	cbg.Set("boundingBox", v)
}
func (cbg *ConeBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: cbg.Get("boundingSphere")}
}
func (cbg *ConeBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	cbg.Set("boundingSphere", v)
}
func (cbg *ConeBufferGeometry) DrawRange() js.Value {
	return cbg.Get("drawRange")
}
func (cbg *ConeBufferGeometry) SetDrawRange(v js.Value) {
	cbg.Set("drawRange", v)
}
func (cbg *ConeBufferGeometry) Drawcalls() js.Value {
	return cbg.Get("drawcalls")
}
func (cbg *ConeBufferGeometry) SetDrawcalls(v js.Value) {
	cbg.Set("drawcalls", v)
}
func (cbg *ConeBufferGeometry) Groups() []js.Value {
	return []js.Value(cbg.Get("groups"))
}
func (cbg *ConeBufferGeometry) SetGroups(v []js.Value) {
	cbg.Set("groups", v)
}
func (cbg *ConeBufferGeometry) Id() int {
	return cbg.Get("id").Int()
}
func (cbg *ConeBufferGeometry) SetId(v int) {
	cbg.Set("id", v)
}
func (cbg *ConeBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: cbg.Get("index")}
}
func (cbg *ConeBufferGeometry) SetIndex(v *core.BufferAttribute) {
	cbg.Set("index", v)
}
func (cbg *ConeBufferGeometry) MorphAttributes() js.Value {
	return cbg.Get("morphAttributes")
}
func (cbg *ConeBufferGeometry) SetMorphAttributes(v js.Value) {
	cbg.Set("morphAttributes", v)
}
func (cbg *ConeBufferGeometry) Name() string {
	return cbg.Get("name").String()
}
func (cbg *ConeBufferGeometry) SetName(v string) {
	cbg.Set("name", v)
}
func (cbg *ConeBufferGeometry) Offsets() js.Value {
	return cbg.Get("offsets")
}
func (cbg *ConeBufferGeometry) SetOffsets(v js.Value) {
	cbg.Set("offsets", v)
}
func (cbg *ConeBufferGeometry) Type() string {
	return cbg.Get("type").String()
}
func (cbg *ConeBufferGeometry) SetType(v string) {
	cbg.Set("type", v)
}
func (cbg *ConeBufferGeometry) UserData() js.Value {
	return cbg.Get("userData")
}
func (cbg *ConeBufferGeometry) SetUserData(v js.Value) {
	cbg.Set("userData", v)
}
func (cbg *ConeBufferGeometry) Uuid() string {
	return cbg.Get("uuid").String()
}
func (cbg *ConeBufferGeometry) SetUuid(v string) {
	cbg.Set("uuid", v)
}
func (cbg *ConeBufferGeometry) MaxIndex() int {
	return cbg.Get("MaxIndex").Int()
}
func (cbg *ConeBufferGeometry) SetMaxIndex(v int) {
	cbg.Set("MaxIndex", v)
}
func (cbg *ConeBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("addAttribute", name, attribute)}
}
func (cbg *ConeBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return cbg.Call("addAttribute", name, array, itemSize)
}
func (cbg *ConeBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	cbg.Call("addDrawCall", start, count, indexOffset)
}
func (cbg *ConeBufferGeometry) AddEventListener(typ string, listener js.Value) {
	cbg.Call("addEventListener", typ, listener)
}
func (cbg *ConeBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	cbg.Call("addGroup", start, count, materialIndex)
}
func (cbg *ConeBufferGeometry) AddIndex(index js.Value) {
	cbg.Call("addIndex", index)
}
func (cbg *ConeBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("applyMatrix", matrix)}
}
func (cbg *ConeBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("center")}
}
func (cbg *ConeBufferGeometry) ClearDrawCalls() {
	cbg.Call("clearDrawCalls")
}
func (cbg *ConeBufferGeometry) ClearGroups() {
	cbg.Call("clearGroups")
}
func (cbg *ConeBufferGeometry) Clone() this {
	return this(cbg.Call("clone"))
}
func (cbg *ConeBufferGeometry) ComputeBoundingBox() {
	cbg.Call("computeBoundingBox")
}
func (cbg *ConeBufferGeometry) ComputeBoundingSphere() {
	cbg.Call("computeBoundingSphere")
}
func (cbg *ConeBufferGeometry) ComputeVertexNormals() {
	cbg.Call("computeVertexNormals")
}
func (cbg *ConeBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(cbg.Call("copy", source))
}
func (cbg *ConeBufferGeometry) DispatchEvent(event js.Value) {
	cbg.Call("dispatchEvent", event)
}
func (cbg *ConeBufferGeometry) Dispose() {
	cbg.Call("dispose")
}
func (cbg *ConeBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("fromDirectGeometry", geometry)}
}
func (cbg *ConeBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("fromGeometry", geometry, settings)}
}
func (cbg *ConeBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(cbg.Call("getAttribute", name))
}
func (cbg *ConeBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: cbg.Call("getIndex")}
}
func (cbg *ConeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cbg.Call("hasEventListener", typ, listener).Bool()
}
func (cbg *ConeBufferGeometry) LookAt(v *math.Vector3) {
	cbg.Call("lookAt", v)
}
func (cbg *ConeBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("merge", geometry, offset)}
}
func (cbg *ConeBufferGeometry) NormalizeNormals() {
	cbg.Call("normalizeNormals")
}
func (cbg *ConeBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("removeAttribute", name)}
}
func (cbg *ConeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	cbg.Call("removeEventListener", typ, listener)
}
func (cbg *ConeBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateX", angle)}
}
func (cbg *ConeBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateY", angle)}
}
func (cbg *ConeBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateZ", angle)}
}
func (cbg *ConeBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("scale", x, y, z)}
}
func (cbg *ConeBufferGeometry) SetDrawRange(start int, count int) {
	cbg.Call("setDrawRange", start, count)
}
func (cbg *ConeBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("setFromObject", object)}
}
func (cbg *ConeBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("setFromPoints", points)}
}
func (cbg *ConeBufferGeometry) SetIndex(index core.BufferAttribute) {
	cbg.Call("setIndex", index)
}
func (cbg *ConeBufferGeometry) ToJSON() js.Value {
	return cbg.Call("toJSON")
}
func (cbg *ConeBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("toNonIndexed")}
}
func (cbg *ConeBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("translate", x, y, z)}
}
func (cbg *ConeBufferGeometry) UpdateFromObject(object *core.Object3D) {
	cbg.Call("updateFromObject", object)
}

type ConeGeometry struct {
	js.Value
}

func NewConeGeometry(radius int, height int, radialSegment int, heightSegment int, openEnded bool, thetaStart int, thetaLength int) *ConeGeometry {
	return &ConeGeometry{Value: get("ConeGeometry").New(radius, height, radialSegment, heightSegment, openEnded, thetaStart, thetaLength)}
}
func (cg *ConeGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: cg.Get("animation")}
}
func (cg *ConeGeometry) SetAnimation(v *animation.AnimationClip) {
	cg.Set("animation", v)
}
func (cg *ConeGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(cg.Get("animations"))
}
func (cg *ConeGeometry) SetAnimations(v []animation.AnimationClip) {
	cg.Set("animations", v)
}
func (cg *ConeGeometry) Bones() []objects.Bone {
	return []objects.Bone(cg.Get("bones"))
}
func (cg *ConeGeometry) SetBones(v []objects.Bone) {
	cg.Set("bones", v)
}
func (cg *ConeGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: cg.Get("boundingBox")}
}
func (cg *ConeGeometry) SetBoundingBox(v *math.Box3) {
	cg.Set("boundingBox", v)
}
func (cg *ConeGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: cg.Get("boundingSphere")}
}
func (cg *ConeGeometry) SetBoundingSphere(v *math.Sphere) {
	cg.Set("boundingSphere", v)
}
func (cg *ConeGeometry) Colors() []math.Color {
	return []math.Color(cg.Get("colors"))
}
func (cg *ConeGeometry) SetColors(v []math.Color) {
	cg.Set("colors", v)
}
func (cg *ConeGeometry) ColorsNeedUpdate() bool {
	return cg.Get("colorsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetColorsNeedUpdate(v bool) {
	cg.Set("colorsNeedUpdate", v)
}
func (cg *ConeGeometry) ElementsNeedUpdate() bool {
	return cg.Get("elementsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetElementsNeedUpdate(v bool) {
	cg.Set("elementsNeedUpdate", v)
}
func (cg *ConeGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(cg.Get("faceVertexUvs"))
}
func (cg *ConeGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	cg.Set("faceVertexUvs", v)
}
func (cg *ConeGeometry) Faces() []core.Face3 {
	return []core.Face3(cg.Get("faces"))
}
func (cg *ConeGeometry) SetFaces(v []core.Face3) {
	cg.Set("faces", v)
}
func (cg *ConeGeometry) GroupsNeedUpdate() bool {
	return cg.Get("groupsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetGroupsNeedUpdate(v bool) {
	cg.Set("groupsNeedUpdate", v)
}
func (cg *ConeGeometry) Id() int {
	return cg.Get("id").Int()
}
func (cg *ConeGeometry) SetId(v int) {
	cg.Set("id", v)
}
func (cg *ConeGeometry) LineDistances() []int {
	return []int(cg.Get("lineDistances"))
}
func (cg *ConeGeometry) SetLineDistances(v []int) {
	cg.Set("lineDistances", v)
}
func (cg *ConeGeometry) LineDistancesNeedUpdate() bool {
	return cg.Get("lineDistancesNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetLineDistancesNeedUpdate(v bool) {
	cg.Set("lineDistancesNeedUpdate", v)
}
func (cg *ConeGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(cg.Get("morphNormals"))
}
func (cg *ConeGeometry) SetMorphNormals(v []core.MorphNormals) {
	cg.Set("morphNormals", v)
}
func (cg *ConeGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(cg.Get("morphTargets"))
}
func (cg *ConeGeometry) SetMorphTargets(v []core.MorphTarget) {
	cg.Set("morphTargets", v)
}
func (cg *ConeGeometry) Name() string {
	return cg.Get("name").String()
}
func (cg *ConeGeometry) SetName(v string) {
	cg.Set("name", v)
}
func (cg *ConeGeometry) NormalsNeedUpdate() bool {
	return cg.Get("normalsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetNormalsNeedUpdate(v bool) {
	cg.Set("normalsNeedUpdate", v)
}
func (cg *ConeGeometry) Parameters() js.Value {
	return cg.Get("parameters")
}
func (cg *ConeGeometry) SetParameters(v js.Value) {
	cg.Set("parameters", v)
}
func (cg *ConeGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(cg.Get("skinIndices"))
}
func (cg *ConeGeometry) SetSkinIndices(v []math.Vector4) {
	cg.Set("skinIndices", v)
}
func (cg *ConeGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(cg.Get("skinWeights"))
}
func (cg *ConeGeometry) SetSkinWeights(v []math.Vector4) {
	cg.Set("skinWeights", v)
}
func (cg *ConeGeometry) Type() string {
	return cg.Get("type").String()
}
func (cg *ConeGeometry) SetType(v string) {
	cg.Set("type", v)
}
func (cg *ConeGeometry) Uuid() string {
	return cg.Get("uuid").String()
}
func (cg *ConeGeometry) SetUuid(v string) {
	cg.Set("uuid", v)
}
func (cg *ConeGeometry) UvsNeedUpdate() bool {
	return cg.Get("uvsNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetUvsNeedUpdate(v bool) {
	cg.Set("uvsNeedUpdate", v)
}
func (cg *ConeGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(cg.Get("vertices"))
}
func (cg *ConeGeometry) SetVertices(v []math.Vector3) {
	cg.Set("vertices", v)
}
func (cg *ConeGeometry) VerticesNeedUpdate() bool {
	return cg.Get("verticesNeedUpdate").Bool()
}
func (cg *ConeGeometry) SetVerticesNeedUpdate(v bool) {
	cg.Set("verticesNeedUpdate", v)
}
func (cg *ConeGeometry) AddEventListener(typ string, listener js.Value) {
	cg.Call("addEventListener", typ, listener)
}
func (cg *ConeGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: cg.Call("applyMatrix", matrix)}
}
func (cg *ConeGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: cg.Call("center")}
}
func (cg *ConeGeometry) Clone() this {
	return this(cg.Call("clone"))
}
func (cg *ConeGeometry) ComputeBoundingBox() {
	cg.Call("computeBoundingBox")
}
func (cg *ConeGeometry) ComputeBoundingSphere() {
	cg.Call("computeBoundingSphere")
}
func (cg *ConeGeometry) ComputeFaceNormals() {
	cg.Call("computeFaceNormals")
}
func (cg *ConeGeometry) ComputeFlatVertexNormals() {
	cg.Call("computeFlatVertexNormals")
}
func (cg *ConeGeometry) ComputeMorphNormals() {
	cg.Call("computeMorphNormals")
}
func (cg *ConeGeometry) ComputeVertexNormals(areaWeighted bool) {
	cg.Call("computeVertexNormals", areaWeighted)
}
func (cg *ConeGeometry) Copy(source *core.Geometry) this {
	return this(cg.Call("copy", source))
}
func (cg *ConeGeometry) DispatchEvent(event js.Value) {
	cg.Call("dispatchEvent", event)
}
func (cg *ConeGeometry) Dispose() {
	cg.Call("dispose")
}
func (cg *ConeGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: cg.Call("fromBufferGeometry", geometry)}
}
func (cg *ConeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cg.Call("hasEventListener", typ, listener).Bool()
}
func (cg *ConeGeometry) LookAt(vector *math.Vector3) {
	cg.Call("lookAt", vector)
}
func (cg *ConeGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	cg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (cg *ConeGeometry) MergeMesh(mesh *objects.Mesh) {
	cg.Call("mergeMesh", mesh)
}
func (cg *ConeGeometry) MergeVertices() int {
	return cg.Call("mergeVertices").Int()
}
func (cg *ConeGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: cg.Call("normalize")}
}
func (cg *ConeGeometry) RemoveEventListener(typ string, listener js.Value) {
	cg.Call("removeEventListener", typ, listener)
}
func (cg *ConeGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateX", angle)}
}
func (cg *ConeGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateY", angle)}
}
func (cg *ConeGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateZ", angle)}
}
func (cg *ConeGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("scale", x, y, z)}
}
func (cg *ConeGeometry) SetFromPoints(points Array) this {
	return this(cg.Call("setFromPoints", points))
}
func (cg *ConeGeometry) SortFacesByMaterialIndex() {
	cg.Call("sortFacesByMaterialIndex")
}
func (cg *ConeGeometry) ToJSON() js.Value {
	return cg.Call("toJSON")
}
func (cg *ConeGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("translate", x, y, z)}
}

package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type CircleBufferGeometry struct {
	js.Value
}

func NewCircleBufferGeometry(radius int, segments int, thetaStart int, thetaLength int) *CircleBufferGeometry {
	return &CircleBufferGeometry{Value: get("CircleBufferGeometry").New(radius, segments, thetaStart, thetaLength)}
}
func (cbg *CircleBufferGeometry) Attributes() js.Value {
	return cbg.Get("attributes")
}
func (cbg *CircleBufferGeometry) SetAttributes(v js.Value) {
	cbg.Set("attributes", v)
}
func (cbg *CircleBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: cbg.Get("boundingBox")}
}
func (cbg *CircleBufferGeometry) SetBoundingBox(v *math.Box3) {
	cbg.Set("boundingBox", v)
}
func (cbg *CircleBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: cbg.Get("boundingSphere")}
}
func (cbg *CircleBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	cbg.Set("boundingSphere", v)
}
func (cbg *CircleBufferGeometry) DrawRange() js.Value {
	return cbg.Get("drawRange")
}
func (cbg *CircleBufferGeometry) SetDrawRange(v js.Value) {
	cbg.Set("drawRange", v)
}
func (cbg *CircleBufferGeometry) Drawcalls() js.Value {
	return cbg.Get("drawcalls")
}
func (cbg *CircleBufferGeometry) SetDrawcalls(v js.Value) {
	cbg.Set("drawcalls", v)
}
func (cbg *CircleBufferGeometry) Groups() []js.Value {
	return []js.Value(cbg.Get("groups"))
}
func (cbg *CircleBufferGeometry) SetGroups(v []js.Value) {
	cbg.Set("groups", v)
}
func (cbg *CircleBufferGeometry) Id() int {
	return cbg.Get("id").Int()
}
func (cbg *CircleBufferGeometry) SetId(v int) {
	cbg.Set("id", v)
}
func (cbg *CircleBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: cbg.Get("index")}
}
func (cbg *CircleBufferGeometry) SetIndex(v *core.BufferAttribute) {
	cbg.Set("index", v)
}
func (cbg *CircleBufferGeometry) MorphAttributes() js.Value {
	return cbg.Get("morphAttributes")
}
func (cbg *CircleBufferGeometry) SetMorphAttributes(v js.Value) {
	cbg.Set("morphAttributes", v)
}
func (cbg *CircleBufferGeometry) Name() string {
	return cbg.Get("name").String()
}
func (cbg *CircleBufferGeometry) SetName(v string) {
	cbg.Set("name", v)
}
func (cbg *CircleBufferGeometry) Offsets() js.Value {
	return cbg.Get("offsets")
}
func (cbg *CircleBufferGeometry) SetOffsets(v js.Value) {
	cbg.Set("offsets", v)
}
func (cbg *CircleBufferGeometry) Parameters() js.Value {
	return cbg.Get("parameters")
}
func (cbg *CircleBufferGeometry) SetParameters(v js.Value) {
	cbg.Set("parameters", v)
}
func (cbg *CircleBufferGeometry) Type() string {
	return cbg.Get("type").String()
}
func (cbg *CircleBufferGeometry) SetType(v string) {
	cbg.Set("type", v)
}
func (cbg *CircleBufferGeometry) UserData() js.Value {
	return cbg.Get("userData")
}
func (cbg *CircleBufferGeometry) SetUserData(v js.Value) {
	cbg.Set("userData", v)
}
func (cbg *CircleBufferGeometry) Uuid() string {
	return cbg.Get("uuid").String()
}
func (cbg *CircleBufferGeometry) SetUuid(v string) {
	cbg.Set("uuid", v)
}
func (cbg *CircleBufferGeometry) MaxIndex() int {
	return cbg.Get("MaxIndex").Int()
}
func (cbg *CircleBufferGeometry) SetMaxIndex(v int) {
	cbg.Set("MaxIndex", v)
}
func (cbg *CircleBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("addAttribute", name, attribute)}
}
func (cbg *CircleBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return cbg.Call("addAttribute", name, array, itemSize)
}
func (cbg *CircleBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	cbg.Call("addDrawCall", start, count, indexOffset)
}
func (cbg *CircleBufferGeometry) AddEventListener(typ string, listener js.Value) {
	cbg.Call("addEventListener", typ, listener)
}
func (cbg *CircleBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	cbg.Call("addGroup", start, count, materialIndex)
}
func (cbg *CircleBufferGeometry) AddIndex(index js.Value) {
	cbg.Call("addIndex", index)
}
func (cbg *CircleBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("applyMatrix", matrix)}
}
func (cbg *CircleBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("center")}
}
func (cbg *CircleBufferGeometry) ClearDrawCalls() {
	cbg.Call("clearDrawCalls")
}
func (cbg *CircleBufferGeometry) ClearGroups() {
	cbg.Call("clearGroups")
}
func (cbg *CircleBufferGeometry) Clone() this {
	return this(cbg.Call("clone"))
}
func (cbg *CircleBufferGeometry) ComputeBoundingBox() {
	cbg.Call("computeBoundingBox")
}
func (cbg *CircleBufferGeometry) ComputeBoundingSphere() {
	cbg.Call("computeBoundingSphere")
}
func (cbg *CircleBufferGeometry) ComputeVertexNormals() {
	cbg.Call("computeVertexNormals")
}
func (cbg *CircleBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(cbg.Call("copy", source))
}
func (cbg *CircleBufferGeometry) DispatchEvent(event js.Value) {
	cbg.Call("dispatchEvent", event)
}
func (cbg *CircleBufferGeometry) Dispose() {
	cbg.Call("dispose")
}
func (cbg *CircleBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("fromDirectGeometry", geometry)}
}
func (cbg *CircleBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("fromGeometry", geometry, settings)}
}
func (cbg *CircleBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(cbg.Call("getAttribute", name))
}
func (cbg *CircleBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: cbg.Call("getIndex")}
}
func (cbg *CircleBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cbg.Call("hasEventListener", typ, listener).Bool()
}
func (cbg *CircleBufferGeometry) LookAt(v *math.Vector3) {
	cbg.Call("lookAt", v)
}
func (cbg *CircleBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("merge", geometry, offset)}
}
func (cbg *CircleBufferGeometry) NormalizeNormals() {
	cbg.Call("normalizeNormals")
}
func (cbg *CircleBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("removeAttribute", name)}
}
func (cbg *CircleBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	cbg.Call("removeEventListener", typ, listener)
}
func (cbg *CircleBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateX", angle)}
}
func (cbg *CircleBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateY", angle)}
}
func (cbg *CircleBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateZ", angle)}
}
func (cbg *CircleBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("scale", x, y, z)}
}
func (cbg *CircleBufferGeometry) SetDrawRange(start int, count int) {
	cbg.Call("setDrawRange", start, count)
}
func (cbg *CircleBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("setFromObject", object)}
}
func (cbg *CircleBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("setFromPoints", points)}
}
func (cbg *CircleBufferGeometry) SetIndex(index core.BufferAttribute) {
	cbg.Call("setIndex", index)
}
func (cbg *CircleBufferGeometry) ToJSON() js.Value {
	return cbg.Call("toJSON")
}
func (cbg *CircleBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("toNonIndexed")}
}
func (cbg *CircleBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("translate", x, y, z)}
}
func (cbg *CircleBufferGeometry) UpdateFromObject(object *core.Object3D) {
	cbg.Call("updateFromObject", object)
}

type CircleGeometry struct {
	js.Value
}

func NewCircleGeometry(radius int, segments int, thetaStart int, thetaLength int) *CircleGeometry {
	return &CircleGeometry{Value: get("CircleGeometry").New(radius, segments, thetaStart, thetaLength)}
}
func (cg *CircleGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: cg.Get("animation")}
}
func (cg *CircleGeometry) SetAnimation(v *animation.AnimationClip) {
	cg.Set("animation", v)
}
func (cg *CircleGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(cg.Get("animations"))
}
func (cg *CircleGeometry) SetAnimations(v []animation.AnimationClip) {
	cg.Set("animations", v)
}
func (cg *CircleGeometry) Bones() []objects.Bone {
	return []objects.Bone(cg.Get("bones"))
}
func (cg *CircleGeometry) SetBones(v []objects.Bone) {
	cg.Set("bones", v)
}
func (cg *CircleGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: cg.Get("boundingBox")}
}
func (cg *CircleGeometry) SetBoundingBox(v *math.Box3) {
	cg.Set("boundingBox", v)
}
func (cg *CircleGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: cg.Get("boundingSphere")}
}
func (cg *CircleGeometry) SetBoundingSphere(v *math.Sphere) {
	cg.Set("boundingSphere", v)
}
func (cg *CircleGeometry) Colors() []math.Color {
	return []math.Color(cg.Get("colors"))
}
func (cg *CircleGeometry) SetColors(v []math.Color) {
	cg.Set("colors", v)
}
func (cg *CircleGeometry) ColorsNeedUpdate() bool {
	return cg.Get("colorsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetColorsNeedUpdate(v bool) {
	cg.Set("colorsNeedUpdate", v)
}
func (cg *CircleGeometry) ElementsNeedUpdate() bool {
	return cg.Get("elementsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetElementsNeedUpdate(v bool) {
	cg.Set("elementsNeedUpdate", v)
}
func (cg *CircleGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(cg.Get("faceVertexUvs"))
}
func (cg *CircleGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	cg.Set("faceVertexUvs", v)
}
func (cg *CircleGeometry) Faces() []core.Face3 {
	return []core.Face3(cg.Get("faces"))
}
func (cg *CircleGeometry) SetFaces(v []core.Face3) {
	cg.Set("faces", v)
}
func (cg *CircleGeometry) GroupsNeedUpdate() bool {
	return cg.Get("groupsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetGroupsNeedUpdate(v bool) {
	cg.Set("groupsNeedUpdate", v)
}
func (cg *CircleGeometry) Id() int {
	return cg.Get("id").Int()
}
func (cg *CircleGeometry) SetId(v int) {
	cg.Set("id", v)
}
func (cg *CircleGeometry) LineDistances() []int {
	return []int(cg.Get("lineDistances"))
}
func (cg *CircleGeometry) SetLineDistances(v []int) {
	cg.Set("lineDistances", v)
}
func (cg *CircleGeometry) LineDistancesNeedUpdate() bool {
	return cg.Get("lineDistancesNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetLineDistancesNeedUpdate(v bool) {
	cg.Set("lineDistancesNeedUpdate", v)
}
func (cg *CircleGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(cg.Get("morphNormals"))
}
func (cg *CircleGeometry) SetMorphNormals(v []core.MorphNormals) {
	cg.Set("morphNormals", v)
}
func (cg *CircleGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(cg.Get("morphTargets"))
}
func (cg *CircleGeometry) SetMorphTargets(v []core.MorphTarget) {
	cg.Set("morphTargets", v)
}
func (cg *CircleGeometry) Name() string {
	return cg.Get("name").String()
}
func (cg *CircleGeometry) SetName(v string) {
	cg.Set("name", v)
}
func (cg *CircleGeometry) NormalsNeedUpdate() bool {
	return cg.Get("normalsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetNormalsNeedUpdate(v bool) {
	cg.Set("normalsNeedUpdate", v)
}
func (cg *CircleGeometry) Parameters() js.Value {
	return cg.Get("parameters")
}
func (cg *CircleGeometry) SetParameters(v js.Value) {
	cg.Set("parameters", v)
}
func (cg *CircleGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(cg.Get("skinIndices"))
}
func (cg *CircleGeometry) SetSkinIndices(v []math.Vector4) {
	cg.Set("skinIndices", v)
}
func (cg *CircleGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(cg.Get("skinWeights"))
}
func (cg *CircleGeometry) SetSkinWeights(v []math.Vector4) {
	cg.Set("skinWeights", v)
}
func (cg *CircleGeometry) Type() string {
	return cg.Get("type").String()
}
func (cg *CircleGeometry) SetType(v string) {
	cg.Set("type", v)
}
func (cg *CircleGeometry) Uuid() string {
	return cg.Get("uuid").String()
}
func (cg *CircleGeometry) SetUuid(v string) {
	cg.Set("uuid", v)
}
func (cg *CircleGeometry) UvsNeedUpdate() bool {
	return cg.Get("uvsNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetUvsNeedUpdate(v bool) {
	cg.Set("uvsNeedUpdate", v)
}
func (cg *CircleGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(cg.Get("vertices"))
}
func (cg *CircleGeometry) SetVertices(v []math.Vector3) {
	cg.Set("vertices", v)
}
func (cg *CircleGeometry) VerticesNeedUpdate() bool {
	return cg.Get("verticesNeedUpdate").Bool()
}
func (cg *CircleGeometry) SetVerticesNeedUpdate(v bool) {
	cg.Set("verticesNeedUpdate", v)
}
func (cg *CircleGeometry) AddEventListener(typ string, listener js.Value) {
	cg.Call("addEventListener", typ, listener)
}
func (cg *CircleGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: cg.Call("applyMatrix", matrix)}
}
func (cg *CircleGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: cg.Call("center")}
}
func (cg *CircleGeometry) Clone() this {
	return this(cg.Call("clone"))
}
func (cg *CircleGeometry) ComputeBoundingBox() {
	cg.Call("computeBoundingBox")
}
func (cg *CircleGeometry) ComputeBoundingSphere() {
	cg.Call("computeBoundingSphere")
}
func (cg *CircleGeometry) ComputeFaceNormals() {
	cg.Call("computeFaceNormals")
}
func (cg *CircleGeometry) ComputeFlatVertexNormals() {
	cg.Call("computeFlatVertexNormals")
}
func (cg *CircleGeometry) ComputeMorphNormals() {
	cg.Call("computeMorphNormals")
}
func (cg *CircleGeometry) ComputeVertexNormals(areaWeighted bool) {
	cg.Call("computeVertexNormals", areaWeighted)
}
func (cg *CircleGeometry) Copy(source *core.Geometry) this {
	return this(cg.Call("copy", source))
}
func (cg *CircleGeometry) DispatchEvent(event js.Value) {
	cg.Call("dispatchEvent", event)
}
func (cg *CircleGeometry) Dispose() {
	cg.Call("dispose")
}
func (cg *CircleGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: cg.Call("fromBufferGeometry", geometry)}
}
func (cg *CircleGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cg.Call("hasEventListener", typ, listener).Bool()
}
func (cg *CircleGeometry) LookAt(vector *math.Vector3) {
	cg.Call("lookAt", vector)
}
func (cg *CircleGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	cg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (cg *CircleGeometry) MergeMesh(mesh *objects.Mesh) {
	cg.Call("mergeMesh", mesh)
}
func (cg *CircleGeometry) MergeVertices() int {
	return cg.Call("mergeVertices").Int()
}
func (cg *CircleGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: cg.Call("normalize")}
}
func (cg *CircleGeometry) RemoveEventListener(typ string, listener js.Value) {
	cg.Call("removeEventListener", typ, listener)
}
func (cg *CircleGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateX", angle)}
}
func (cg *CircleGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateY", angle)}
}
func (cg *CircleGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateZ", angle)}
}
func (cg *CircleGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("scale", x, y, z)}
}
func (cg *CircleGeometry) SetFromPoints(points Array) this {
	return this(cg.Call("setFromPoints", points))
}
func (cg *CircleGeometry) SortFacesByMaterialIndex() {
	cg.Call("sortFacesByMaterialIndex")
}
func (cg *CircleGeometry) ToJSON() js.Value {
	return cg.Call("toJSON")
}
func (cg *CircleGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("translate", x, y, z)}
}

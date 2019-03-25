package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type RingBufferGeometry struct {
	js.Value
}

func NewRingBufferGeometry(innerRadius int, outerRadius int, thetaSegments int, phiSegments int, thetaStart int, thetaLength int) *RingBufferGeometry {
	return &RingBufferGeometry{Value: get("RingBufferGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)}
}
func (rbg *RingBufferGeometry) Attributes() js.Value {
	return rbg.Get("attributes")
}
func (rbg *RingBufferGeometry) SetAttributes(v js.Value) {
	rbg.Set("attributes", v)
}
func (rbg *RingBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: rbg.Get("boundingBox")}
}
func (rbg *RingBufferGeometry) SetBoundingBox(v *math.Box3) {
	rbg.Set("boundingBox", v)
}
func (rbg *RingBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: rbg.Get("boundingSphere")}
}
func (rbg *RingBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	rbg.Set("boundingSphere", v)
}
func (rbg *RingBufferGeometry) DrawRange() js.Value {
	return rbg.Get("drawRange")
}
func (rbg *RingBufferGeometry) SetDrawRange(v js.Value) {
	rbg.Set("drawRange", v)
}
func (rbg *RingBufferGeometry) Drawcalls() js.Value {
	return rbg.Get("drawcalls")
}
func (rbg *RingBufferGeometry) SetDrawcalls(v js.Value) {
	rbg.Set("drawcalls", v)
}
func (rbg *RingBufferGeometry) Groups() []js.Value {
	return []js.Value(rbg.Get("groups"))
}
func (rbg *RingBufferGeometry) SetGroups(v []js.Value) {
	rbg.Set("groups", v)
}
func (rbg *RingBufferGeometry) Id() int {
	return rbg.Get("id").Int()
}
func (rbg *RingBufferGeometry) SetId(v int) {
	rbg.Set("id", v)
}
func (rbg *RingBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: rbg.Get("index")}
}
func (rbg *RingBufferGeometry) SetIndex(v *core.BufferAttribute) {
	rbg.Set("index", v)
}
func (rbg *RingBufferGeometry) MorphAttributes() js.Value {
	return rbg.Get("morphAttributes")
}
func (rbg *RingBufferGeometry) SetMorphAttributes(v js.Value) {
	rbg.Set("morphAttributes", v)
}
func (rbg *RingBufferGeometry) Name() string {
	return rbg.Get("name").String()
}
func (rbg *RingBufferGeometry) SetName(v string) {
	rbg.Set("name", v)
}
func (rbg *RingBufferGeometry) Offsets() js.Value {
	return rbg.Get("offsets")
}
func (rbg *RingBufferGeometry) SetOffsets(v js.Value) {
	rbg.Set("offsets", v)
}
func (rbg *RingBufferGeometry) Parameters() js.Value {
	return rbg.Get("parameters")
}
func (rbg *RingBufferGeometry) SetParameters(v js.Value) {
	rbg.Set("parameters", v)
}
func (rbg *RingBufferGeometry) Type() string {
	return rbg.Get("type").String()
}
func (rbg *RingBufferGeometry) SetType(v string) {
	rbg.Set("type", v)
}
func (rbg *RingBufferGeometry) UserData() js.Value {
	return rbg.Get("userData")
}
func (rbg *RingBufferGeometry) SetUserData(v js.Value) {
	rbg.Set("userData", v)
}
func (rbg *RingBufferGeometry) Uuid() string {
	return rbg.Get("uuid").String()
}
func (rbg *RingBufferGeometry) SetUuid(v string) {
	rbg.Set("uuid", v)
}
func (rbg *RingBufferGeometry) MaxIndex() int {
	return rbg.Get("MaxIndex").Int()
}
func (rbg *RingBufferGeometry) SetMaxIndex(v int) {
	rbg.Set("MaxIndex", v)
}
func (rbg *RingBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("addAttribute", name, attribute)}
}
func (rbg *RingBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return rbg.Call("addAttribute", name, array, itemSize)
}
func (rbg *RingBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	rbg.Call("addDrawCall", start, count, indexOffset)
}
func (rbg *RingBufferGeometry) AddEventListener(typ string, listener js.Value) {
	rbg.Call("addEventListener", typ, listener)
}
func (rbg *RingBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	rbg.Call("addGroup", start, count, materialIndex)
}
func (rbg *RingBufferGeometry) AddIndex(index js.Value) {
	rbg.Call("addIndex", index)
}
func (rbg *RingBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("applyMatrix", matrix)}
}
func (rbg *RingBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("center")}
}
func (rbg *RingBufferGeometry) ClearDrawCalls() {
	rbg.Call("clearDrawCalls")
}
func (rbg *RingBufferGeometry) ClearGroups() {
	rbg.Call("clearGroups")
}
func (rbg *RingBufferGeometry) Clone() this {
	return this(rbg.Call("clone"))
}
func (rbg *RingBufferGeometry) ComputeBoundingBox() {
	rbg.Call("computeBoundingBox")
}
func (rbg *RingBufferGeometry) ComputeBoundingSphere() {
	rbg.Call("computeBoundingSphere")
}
func (rbg *RingBufferGeometry) ComputeVertexNormals() {
	rbg.Call("computeVertexNormals")
}
func (rbg *RingBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(rbg.Call("copy", source))
}
func (rbg *RingBufferGeometry) DispatchEvent(event js.Value) {
	rbg.Call("dispatchEvent", event)
}
func (rbg *RingBufferGeometry) Dispose() {
	rbg.Call("dispose")
}
func (rbg *RingBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("fromDirectGeometry", geometry)}
}
func (rbg *RingBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("fromGeometry", geometry, settings)}
}
func (rbg *RingBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(rbg.Call("getAttribute", name))
}
func (rbg *RingBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: rbg.Call("getIndex")}
}
func (rbg *RingBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return rbg.Call("hasEventListener", typ, listener).Bool()
}
func (rbg *RingBufferGeometry) LookAt(v *math.Vector3) {
	rbg.Call("lookAt", v)
}
func (rbg *RingBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("merge", geometry, offset)}
}
func (rbg *RingBufferGeometry) NormalizeNormals() {
	rbg.Call("normalizeNormals")
}
func (rbg *RingBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("removeAttribute", name)}
}
func (rbg *RingBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	rbg.Call("removeEventListener", typ, listener)
}
func (rbg *RingBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("rotateX", angle)}
}
func (rbg *RingBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("rotateY", angle)}
}
func (rbg *RingBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("rotateZ", angle)}
}
func (rbg *RingBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("scale", x, y, z)}
}
func (rbg *RingBufferGeometry) SetDrawRange(start int, count int) {
	rbg.Call("setDrawRange", start, count)
}
func (rbg *RingBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("setFromObject", object)}
}
func (rbg *RingBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("setFromPoints", points)}
}
func (rbg *RingBufferGeometry) SetIndex(index core.BufferAttribute) {
	rbg.Call("setIndex", index)
}
func (rbg *RingBufferGeometry) ToJSON() js.Value {
	return rbg.Call("toJSON")
}
func (rbg *RingBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("toNonIndexed")}
}
func (rbg *RingBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: rbg.Call("translate", x, y, z)}
}
func (rbg *RingBufferGeometry) UpdateFromObject(object *core.Object3D) {
	rbg.Call("updateFromObject", object)
}

type RingGeometry struct {
	js.Value
}

func NewRingGeometry(innerRadius int, outerRadius int, thetaSegments int, phiSegments int, thetaStart int, thetaLength int) *RingGeometry {
	return &RingGeometry{Value: get("RingGeometry").New(innerRadius, outerRadius, thetaSegments, phiSegments, thetaStart, thetaLength)}
}
func (rg *RingGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: rg.Get("animation")}
}
func (rg *RingGeometry) SetAnimation(v *animation.AnimationClip) {
	rg.Set("animation", v)
}
func (rg *RingGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(rg.Get("animations"))
}
func (rg *RingGeometry) SetAnimations(v []animation.AnimationClip) {
	rg.Set("animations", v)
}
func (rg *RingGeometry) Bones() []objects.Bone {
	return []objects.Bone(rg.Get("bones"))
}
func (rg *RingGeometry) SetBones(v []objects.Bone) {
	rg.Set("bones", v)
}
func (rg *RingGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: rg.Get("boundingBox")}
}
func (rg *RingGeometry) SetBoundingBox(v *math.Box3) {
	rg.Set("boundingBox", v)
}
func (rg *RingGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: rg.Get("boundingSphere")}
}
func (rg *RingGeometry) SetBoundingSphere(v *math.Sphere) {
	rg.Set("boundingSphere", v)
}
func (rg *RingGeometry) Colors() []math.Color {
	return []math.Color(rg.Get("colors"))
}
func (rg *RingGeometry) SetColors(v []math.Color) {
	rg.Set("colors", v)
}
func (rg *RingGeometry) ColorsNeedUpdate() bool {
	return rg.Get("colorsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetColorsNeedUpdate(v bool) {
	rg.Set("colorsNeedUpdate", v)
}
func (rg *RingGeometry) ElementsNeedUpdate() bool {
	return rg.Get("elementsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetElementsNeedUpdate(v bool) {
	rg.Set("elementsNeedUpdate", v)
}
func (rg *RingGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(rg.Get("faceVertexUvs"))
}
func (rg *RingGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	rg.Set("faceVertexUvs", v)
}
func (rg *RingGeometry) Faces() []core.Face3 {
	return []core.Face3(rg.Get("faces"))
}
func (rg *RingGeometry) SetFaces(v []core.Face3) {
	rg.Set("faces", v)
}
func (rg *RingGeometry) GroupsNeedUpdate() bool {
	return rg.Get("groupsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetGroupsNeedUpdate(v bool) {
	rg.Set("groupsNeedUpdate", v)
}
func (rg *RingGeometry) Id() int {
	return rg.Get("id").Int()
}
func (rg *RingGeometry) SetId(v int) {
	rg.Set("id", v)
}
func (rg *RingGeometry) LineDistances() []int {
	return []int(rg.Get("lineDistances"))
}
func (rg *RingGeometry) SetLineDistances(v []int) {
	rg.Set("lineDistances", v)
}
func (rg *RingGeometry) LineDistancesNeedUpdate() bool {
	return rg.Get("lineDistancesNeedUpdate").Bool()
}
func (rg *RingGeometry) SetLineDistancesNeedUpdate(v bool) {
	rg.Set("lineDistancesNeedUpdate", v)
}
func (rg *RingGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(rg.Get("morphNormals"))
}
func (rg *RingGeometry) SetMorphNormals(v []core.MorphNormals) {
	rg.Set("morphNormals", v)
}
func (rg *RingGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(rg.Get("morphTargets"))
}
func (rg *RingGeometry) SetMorphTargets(v []core.MorphTarget) {
	rg.Set("morphTargets", v)
}
func (rg *RingGeometry) Name() string {
	return rg.Get("name").String()
}
func (rg *RingGeometry) SetName(v string) {
	rg.Set("name", v)
}
func (rg *RingGeometry) NormalsNeedUpdate() bool {
	return rg.Get("normalsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetNormalsNeedUpdate(v bool) {
	rg.Set("normalsNeedUpdate", v)
}
func (rg *RingGeometry) Parameters() js.Value {
	return rg.Get("parameters")
}
func (rg *RingGeometry) SetParameters(v js.Value) {
	rg.Set("parameters", v)
}
func (rg *RingGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(rg.Get("skinIndices"))
}
func (rg *RingGeometry) SetSkinIndices(v []math.Vector4) {
	rg.Set("skinIndices", v)
}
func (rg *RingGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(rg.Get("skinWeights"))
}
func (rg *RingGeometry) SetSkinWeights(v []math.Vector4) {
	rg.Set("skinWeights", v)
}
func (rg *RingGeometry) Type() string {
	return rg.Get("type").String()
}
func (rg *RingGeometry) SetType(v string) {
	rg.Set("type", v)
}
func (rg *RingGeometry) Uuid() string {
	return rg.Get("uuid").String()
}
func (rg *RingGeometry) SetUuid(v string) {
	rg.Set("uuid", v)
}
func (rg *RingGeometry) UvsNeedUpdate() bool {
	return rg.Get("uvsNeedUpdate").Bool()
}
func (rg *RingGeometry) SetUvsNeedUpdate(v bool) {
	rg.Set("uvsNeedUpdate", v)
}
func (rg *RingGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(rg.Get("vertices"))
}
func (rg *RingGeometry) SetVertices(v []math.Vector3) {
	rg.Set("vertices", v)
}
func (rg *RingGeometry) VerticesNeedUpdate() bool {
	return rg.Get("verticesNeedUpdate").Bool()
}
func (rg *RingGeometry) SetVerticesNeedUpdate(v bool) {
	rg.Set("verticesNeedUpdate", v)
}
func (rg *RingGeometry) AddEventListener(typ string, listener js.Value) {
	rg.Call("addEventListener", typ, listener)
}
func (rg *RingGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: rg.Call("applyMatrix", matrix)}
}
func (rg *RingGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: rg.Call("center")}
}
func (rg *RingGeometry) Clone() this {
	return this(rg.Call("clone"))
}
func (rg *RingGeometry) ComputeBoundingBox() {
	rg.Call("computeBoundingBox")
}
func (rg *RingGeometry) ComputeBoundingSphere() {
	rg.Call("computeBoundingSphere")
}
func (rg *RingGeometry) ComputeFaceNormals() {
	rg.Call("computeFaceNormals")
}
func (rg *RingGeometry) ComputeFlatVertexNormals() {
	rg.Call("computeFlatVertexNormals")
}
func (rg *RingGeometry) ComputeMorphNormals() {
	rg.Call("computeMorphNormals")
}
func (rg *RingGeometry) ComputeVertexNormals(areaWeighted bool) {
	rg.Call("computeVertexNormals", areaWeighted)
}
func (rg *RingGeometry) Copy(source *core.Geometry) this {
	return this(rg.Call("copy", source))
}
func (rg *RingGeometry) DispatchEvent(event js.Value) {
	rg.Call("dispatchEvent", event)
}
func (rg *RingGeometry) Dispose() {
	rg.Call("dispose")
}
func (rg *RingGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: rg.Call("fromBufferGeometry", geometry)}
}
func (rg *RingGeometry) HasEventListener(typ string, listener js.Value) bool {
	return rg.Call("hasEventListener", typ, listener).Bool()
}
func (rg *RingGeometry) LookAt(vector *math.Vector3) {
	rg.Call("lookAt", vector)
}
func (rg *RingGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	rg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (rg *RingGeometry) MergeMesh(mesh *objects.Mesh) {
	rg.Call("mergeMesh", mesh)
}
func (rg *RingGeometry) MergeVertices() int {
	return rg.Call("mergeVertices").Int()
}
func (rg *RingGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: rg.Call("normalize")}
}
func (rg *RingGeometry) RemoveEventListener(typ string, listener js.Value) {
	rg.Call("removeEventListener", typ, listener)
}
func (rg *RingGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: rg.Call("rotateX", angle)}
}
func (rg *RingGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: rg.Call("rotateY", angle)}
}
func (rg *RingGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: rg.Call("rotateZ", angle)}
}
func (rg *RingGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: rg.Call("scale", x, y, z)}
}
func (rg *RingGeometry) SetFromPoints(points Array) this {
	return this(rg.Call("setFromPoints", points))
}
func (rg *RingGeometry) SortFacesByMaterialIndex() {
	rg.Call("sortFacesByMaterialIndex")
}
func (rg *RingGeometry) ToJSON() js.Value {
	return rg.Call("toJSON")
}
func (rg *RingGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: rg.Call("translate", x, y, z)}
}

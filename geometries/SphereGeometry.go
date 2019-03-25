package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type SphereBufferGeometry struct {
	js.Value
}

func NewSphereBufferGeometry(radius int, widthSegments int, heightSegments int, phiStart int, phiLength int, thetaStart int, thetaLength int) *SphereBufferGeometry {
	return &SphereBufferGeometry{Value: get("SphereBufferGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)}
}
func (sbg *SphereBufferGeometry) Attributes() js.Value {
	return sbg.Get("attributes")
}
func (sbg *SphereBufferGeometry) SetAttributes(v js.Value) {
	sbg.Set("attributes", v)
}
func (sbg *SphereBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: sbg.Get("boundingBox")}
}
func (sbg *SphereBufferGeometry) SetBoundingBox(v *math.Box3) {
	sbg.Set("boundingBox", v)
}
func (sbg *SphereBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: sbg.Get("boundingSphere")}
}
func (sbg *SphereBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	sbg.Set("boundingSphere", v)
}
func (sbg *SphereBufferGeometry) DrawRange() js.Value {
	return sbg.Get("drawRange")
}
func (sbg *SphereBufferGeometry) SetDrawRange(v js.Value) {
	sbg.Set("drawRange", v)
}
func (sbg *SphereBufferGeometry) Drawcalls() js.Value {
	return sbg.Get("drawcalls")
}
func (sbg *SphereBufferGeometry) SetDrawcalls(v js.Value) {
	sbg.Set("drawcalls", v)
}
func (sbg *SphereBufferGeometry) Groups() []js.Value {
	return []js.Value(sbg.Get("groups"))
}
func (sbg *SphereBufferGeometry) SetGroups(v []js.Value) {
	sbg.Set("groups", v)
}
func (sbg *SphereBufferGeometry) Id() int {
	return sbg.Get("id").Int()
}
func (sbg *SphereBufferGeometry) SetId(v int) {
	sbg.Set("id", v)
}
func (sbg *SphereBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: sbg.Get("index")}
}
func (sbg *SphereBufferGeometry) SetIndex(v *core.BufferAttribute) {
	sbg.Set("index", v)
}
func (sbg *SphereBufferGeometry) MorphAttributes() js.Value {
	return sbg.Get("morphAttributes")
}
func (sbg *SphereBufferGeometry) SetMorphAttributes(v js.Value) {
	sbg.Set("morphAttributes", v)
}
func (sbg *SphereBufferGeometry) Name() string {
	return sbg.Get("name").String()
}
func (sbg *SphereBufferGeometry) SetName(v string) {
	sbg.Set("name", v)
}
func (sbg *SphereBufferGeometry) Offsets() js.Value {
	return sbg.Get("offsets")
}
func (sbg *SphereBufferGeometry) SetOffsets(v js.Value) {
	sbg.Set("offsets", v)
}
func (sbg *SphereBufferGeometry) Parameters() js.Value {
	return sbg.Get("parameters")
}
func (sbg *SphereBufferGeometry) SetParameters(v js.Value) {
	sbg.Set("parameters", v)
}
func (sbg *SphereBufferGeometry) Type() string {
	return sbg.Get("type").String()
}
func (sbg *SphereBufferGeometry) SetType(v string) {
	sbg.Set("type", v)
}
func (sbg *SphereBufferGeometry) UserData() js.Value {
	return sbg.Get("userData")
}
func (sbg *SphereBufferGeometry) SetUserData(v js.Value) {
	sbg.Set("userData", v)
}
func (sbg *SphereBufferGeometry) Uuid() string {
	return sbg.Get("uuid").String()
}
func (sbg *SphereBufferGeometry) SetUuid(v string) {
	sbg.Set("uuid", v)
}
func (sbg *SphereBufferGeometry) MaxIndex() int {
	return sbg.Get("MaxIndex").Int()
}
func (sbg *SphereBufferGeometry) SetMaxIndex(v int) {
	sbg.Set("MaxIndex", v)
}
func (sbg *SphereBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("addAttribute", name, attribute)}
}
func (sbg *SphereBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return sbg.Call("addAttribute", name, array, itemSize)
}
func (sbg *SphereBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	sbg.Call("addDrawCall", start, count, indexOffset)
}
func (sbg *SphereBufferGeometry) AddEventListener(typ string, listener js.Value) {
	sbg.Call("addEventListener", typ, listener)
}
func (sbg *SphereBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	sbg.Call("addGroup", start, count, materialIndex)
}
func (sbg *SphereBufferGeometry) AddIndex(index js.Value) {
	sbg.Call("addIndex", index)
}
func (sbg *SphereBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("applyMatrix", matrix)}
}
func (sbg *SphereBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("center")}
}
func (sbg *SphereBufferGeometry) ClearDrawCalls() {
	sbg.Call("clearDrawCalls")
}
func (sbg *SphereBufferGeometry) ClearGroups() {
	sbg.Call("clearGroups")
}
func (sbg *SphereBufferGeometry) Clone() this {
	return this(sbg.Call("clone"))
}
func (sbg *SphereBufferGeometry) ComputeBoundingBox() {
	sbg.Call("computeBoundingBox")
}
func (sbg *SphereBufferGeometry) ComputeBoundingSphere() {
	sbg.Call("computeBoundingSphere")
}
func (sbg *SphereBufferGeometry) ComputeVertexNormals() {
	sbg.Call("computeVertexNormals")
}
func (sbg *SphereBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(sbg.Call("copy", source))
}
func (sbg *SphereBufferGeometry) DispatchEvent(event js.Value) {
	sbg.Call("dispatchEvent", event)
}
func (sbg *SphereBufferGeometry) Dispose() {
	sbg.Call("dispose")
}
func (sbg *SphereBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("fromDirectGeometry", geometry)}
}
func (sbg *SphereBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("fromGeometry", geometry, settings)}
}
func (sbg *SphereBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(sbg.Call("getAttribute", name))
}
func (sbg *SphereBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: sbg.Call("getIndex")}
}
func (sbg *SphereBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return sbg.Call("hasEventListener", typ, listener).Bool()
}
func (sbg *SphereBufferGeometry) LookAt(v *math.Vector3) {
	sbg.Call("lookAt", v)
}
func (sbg *SphereBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("merge", geometry, offset)}
}
func (sbg *SphereBufferGeometry) NormalizeNormals() {
	sbg.Call("normalizeNormals")
}
func (sbg *SphereBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("removeAttribute", name)}
}
func (sbg *SphereBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	sbg.Call("removeEventListener", typ, listener)
}
func (sbg *SphereBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("rotateX", angle)}
}
func (sbg *SphereBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("rotateY", angle)}
}
func (sbg *SphereBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("rotateZ", angle)}
}
func (sbg *SphereBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("scale", x, y, z)}
}
func (sbg *SphereBufferGeometry) SetDrawRange(start int, count int) {
	sbg.Call("setDrawRange", start, count)
}
func (sbg *SphereBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("setFromObject", object)}
}
func (sbg *SphereBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("setFromPoints", points)}
}
func (sbg *SphereBufferGeometry) SetIndex(index core.BufferAttribute) {
	sbg.Call("setIndex", index)
}
func (sbg *SphereBufferGeometry) ToJSON() js.Value {
	return sbg.Call("toJSON")
}
func (sbg *SphereBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("toNonIndexed")}
}
func (sbg *SphereBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("translate", x, y, z)}
}
func (sbg *SphereBufferGeometry) UpdateFromObject(object *core.Object3D) {
	sbg.Call("updateFromObject", object)
}

type SphereGeometry struct {
	js.Value
}

func NewSphereGeometry(radius int, widthSegments int, heightSegments int, phiStart int, phiLength int, thetaStart int, thetaLength int) *SphereGeometry {
	return &SphereGeometry{Value: get("SphereGeometry").New(radius, widthSegments, heightSegments, phiStart, phiLength, thetaStart, thetaLength)}
}
func (sg *SphereGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: sg.Get("animation")}
}
func (sg *SphereGeometry) SetAnimation(v *animation.AnimationClip) {
	sg.Set("animation", v)
}
func (sg *SphereGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(sg.Get("animations"))
}
func (sg *SphereGeometry) SetAnimations(v []animation.AnimationClip) {
	sg.Set("animations", v)
}
func (sg *SphereGeometry) Bones() []objects.Bone {
	return []objects.Bone(sg.Get("bones"))
}
func (sg *SphereGeometry) SetBones(v []objects.Bone) {
	sg.Set("bones", v)
}
func (sg *SphereGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: sg.Get("boundingBox")}
}
func (sg *SphereGeometry) SetBoundingBox(v *math.Box3) {
	sg.Set("boundingBox", v)
}
func (sg *SphereGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: sg.Get("boundingSphere")}
}
func (sg *SphereGeometry) SetBoundingSphere(v *math.Sphere) {
	sg.Set("boundingSphere", v)
}
func (sg *SphereGeometry) Colors() []math.Color {
	return []math.Color(sg.Get("colors"))
}
func (sg *SphereGeometry) SetColors(v []math.Color) {
	sg.Set("colors", v)
}
func (sg *SphereGeometry) ColorsNeedUpdate() bool {
	return sg.Get("colorsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetColorsNeedUpdate(v bool) {
	sg.Set("colorsNeedUpdate", v)
}
func (sg *SphereGeometry) ElementsNeedUpdate() bool {
	return sg.Get("elementsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetElementsNeedUpdate(v bool) {
	sg.Set("elementsNeedUpdate", v)
}
func (sg *SphereGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(sg.Get("faceVertexUvs"))
}
func (sg *SphereGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	sg.Set("faceVertexUvs", v)
}
func (sg *SphereGeometry) Faces() []core.Face3 {
	return []core.Face3(sg.Get("faces"))
}
func (sg *SphereGeometry) SetFaces(v []core.Face3) {
	sg.Set("faces", v)
}
func (sg *SphereGeometry) GroupsNeedUpdate() bool {
	return sg.Get("groupsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetGroupsNeedUpdate(v bool) {
	sg.Set("groupsNeedUpdate", v)
}
func (sg *SphereGeometry) Id() int {
	return sg.Get("id").Int()
}
func (sg *SphereGeometry) SetId(v int) {
	sg.Set("id", v)
}
func (sg *SphereGeometry) LineDistances() []int {
	return []int(sg.Get("lineDistances"))
}
func (sg *SphereGeometry) SetLineDistances(v []int) {
	sg.Set("lineDistances", v)
}
func (sg *SphereGeometry) LineDistancesNeedUpdate() bool {
	return sg.Get("lineDistancesNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetLineDistancesNeedUpdate(v bool) {
	sg.Set("lineDistancesNeedUpdate", v)
}
func (sg *SphereGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(sg.Get("morphNormals"))
}
func (sg *SphereGeometry) SetMorphNormals(v []core.MorphNormals) {
	sg.Set("morphNormals", v)
}
func (sg *SphereGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(sg.Get("morphTargets"))
}
func (sg *SphereGeometry) SetMorphTargets(v []core.MorphTarget) {
	sg.Set("morphTargets", v)
}
func (sg *SphereGeometry) Name() string {
	return sg.Get("name").String()
}
func (sg *SphereGeometry) SetName(v string) {
	sg.Set("name", v)
}
func (sg *SphereGeometry) NormalsNeedUpdate() bool {
	return sg.Get("normalsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetNormalsNeedUpdate(v bool) {
	sg.Set("normalsNeedUpdate", v)
}
func (sg *SphereGeometry) Parameters() js.Value {
	return sg.Get("parameters")
}
func (sg *SphereGeometry) SetParameters(v js.Value) {
	sg.Set("parameters", v)
}
func (sg *SphereGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(sg.Get("skinIndices"))
}
func (sg *SphereGeometry) SetSkinIndices(v []math.Vector4) {
	sg.Set("skinIndices", v)
}
func (sg *SphereGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(sg.Get("skinWeights"))
}
func (sg *SphereGeometry) SetSkinWeights(v []math.Vector4) {
	sg.Set("skinWeights", v)
}
func (sg *SphereGeometry) Type() string {
	return sg.Get("type").String()
}
func (sg *SphereGeometry) SetType(v string) {
	sg.Set("type", v)
}
func (sg *SphereGeometry) Uuid() string {
	return sg.Get("uuid").String()
}
func (sg *SphereGeometry) SetUuid(v string) {
	sg.Set("uuid", v)
}
func (sg *SphereGeometry) UvsNeedUpdate() bool {
	return sg.Get("uvsNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetUvsNeedUpdate(v bool) {
	sg.Set("uvsNeedUpdate", v)
}
func (sg *SphereGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(sg.Get("vertices"))
}
func (sg *SphereGeometry) SetVertices(v []math.Vector3) {
	sg.Set("vertices", v)
}
func (sg *SphereGeometry) VerticesNeedUpdate() bool {
	return sg.Get("verticesNeedUpdate").Bool()
}
func (sg *SphereGeometry) SetVerticesNeedUpdate(v bool) {
	sg.Set("verticesNeedUpdate", v)
}
func (sg *SphereGeometry) AddEventListener(typ string, listener js.Value) {
	sg.Call("addEventListener", typ, listener)
}
func (sg *SphereGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: sg.Call("applyMatrix", matrix)}
}
func (sg *SphereGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: sg.Call("center")}
}
func (sg *SphereGeometry) Clone() this {
	return this(sg.Call("clone"))
}
func (sg *SphereGeometry) ComputeBoundingBox() {
	sg.Call("computeBoundingBox")
}
func (sg *SphereGeometry) ComputeBoundingSphere() {
	sg.Call("computeBoundingSphere")
}
func (sg *SphereGeometry) ComputeFaceNormals() {
	sg.Call("computeFaceNormals")
}
func (sg *SphereGeometry) ComputeFlatVertexNormals() {
	sg.Call("computeFlatVertexNormals")
}
func (sg *SphereGeometry) ComputeMorphNormals() {
	sg.Call("computeMorphNormals")
}
func (sg *SphereGeometry) ComputeVertexNormals(areaWeighted bool) {
	sg.Call("computeVertexNormals", areaWeighted)
}
func (sg *SphereGeometry) Copy(source *core.Geometry) this {
	return this(sg.Call("copy", source))
}
func (sg *SphereGeometry) DispatchEvent(event js.Value) {
	sg.Call("dispatchEvent", event)
}
func (sg *SphereGeometry) Dispose() {
	sg.Call("dispose")
}
func (sg *SphereGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: sg.Call("fromBufferGeometry", geometry)}
}
func (sg *SphereGeometry) HasEventListener(typ string, listener js.Value) bool {
	return sg.Call("hasEventListener", typ, listener).Bool()
}
func (sg *SphereGeometry) LookAt(vector *math.Vector3) {
	sg.Call("lookAt", vector)
}
func (sg *SphereGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	sg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (sg *SphereGeometry) MergeMesh(mesh *objects.Mesh) {
	sg.Call("mergeMesh", mesh)
}
func (sg *SphereGeometry) MergeVertices() int {
	return sg.Call("mergeVertices").Int()
}
func (sg *SphereGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: sg.Call("normalize")}
}
func (sg *SphereGeometry) RemoveEventListener(typ string, listener js.Value) {
	sg.Call("removeEventListener", typ, listener)
}
func (sg *SphereGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("rotateX", angle)}
}
func (sg *SphereGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("rotateY", angle)}
}
func (sg *SphereGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("rotateZ", angle)}
}
func (sg *SphereGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("scale", x, y, z)}
}
func (sg *SphereGeometry) SetFromPoints(points Array) this {
	return this(sg.Call("setFromPoints", points))
}
func (sg *SphereGeometry) SortFacesByMaterialIndex() {
	sg.Call("sortFacesByMaterialIndex")
}
func (sg *SphereGeometry) ToJSON() js.Value {
	return sg.Call("toJSON")
}
func (sg *SphereGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("translate", x, y, z)}
}

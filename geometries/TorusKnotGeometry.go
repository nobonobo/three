package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type TorusKnotBufferGeometry struct {
	js.Value
}

func NewTorusKnotBufferGeometry(radius int, tube int, tubularSegments int, radialSegments int, p int, q int) *TorusKnotBufferGeometry {
	return &TorusKnotBufferGeometry{Value: get("TorusKnotBufferGeometry").New(radius, tube, tubularSegments, radialSegments, p, q)}
}
func (tkbg *TorusKnotBufferGeometry) Attributes() js.Value {
	return tkbg.Get("attributes")
}
func (tkbg *TorusKnotBufferGeometry) SetAttributes(v js.Value) {
	tkbg.Set("attributes", v)
}
func (tkbg *TorusKnotBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: tkbg.Get("boundingBox")}
}
func (tkbg *TorusKnotBufferGeometry) SetBoundingBox(v *math.Box3) {
	tkbg.Set("boundingBox", v)
}
func (tkbg *TorusKnotBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: tkbg.Get("boundingSphere")}
}
func (tkbg *TorusKnotBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	tkbg.Set("boundingSphere", v)
}
func (tkbg *TorusKnotBufferGeometry) DrawRange() js.Value {
	return tkbg.Get("drawRange")
}
func (tkbg *TorusKnotBufferGeometry) SetDrawRange(v js.Value) {
	tkbg.Set("drawRange", v)
}
func (tkbg *TorusKnotBufferGeometry) Drawcalls() js.Value {
	return tkbg.Get("drawcalls")
}
func (tkbg *TorusKnotBufferGeometry) SetDrawcalls(v js.Value) {
	tkbg.Set("drawcalls", v)
}
func (tkbg *TorusKnotBufferGeometry) Groups() []js.Value {
	return []js.Value(tkbg.Get("groups"))
}
func (tkbg *TorusKnotBufferGeometry) SetGroups(v []js.Value) {
	tkbg.Set("groups", v)
}
func (tkbg *TorusKnotBufferGeometry) Id() int {
	return tkbg.Get("id").Int()
}
func (tkbg *TorusKnotBufferGeometry) SetId(v int) {
	tkbg.Set("id", v)
}
func (tkbg *TorusKnotBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: tkbg.Get("index")}
}
func (tkbg *TorusKnotBufferGeometry) SetIndex(v *core.BufferAttribute) {
	tkbg.Set("index", v)
}
func (tkbg *TorusKnotBufferGeometry) MorphAttributes() js.Value {
	return tkbg.Get("morphAttributes")
}
func (tkbg *TorusKnotBufferGeometry) SetMorphAttributes(v js.Value) {
	tkbg.Set("morphAttributes", v)
}
func (tkbg *TorusKnotBufferGeometry) Name() string {
	return tkbg.Get("name").String()
}
func (tkbg *TorusKnotBufferGeometry) SetName(v string) {
	tkbg.Set("name", v)
}
func (tkbg *TorusKnotBufferGeometry) Offsets() js.Value {
	return tkbg.Get("offsets")
}
func (tkbg *TorusKnotBufferGeometry) SetOffsets(v js.Value) {
	tkbg.Set("offsets", v)
}
func (tkbg *TorusKnotBufferGeometry) Parameters() js.Value {
	return tkbg.Get("parameters")
}
func (tkbg *TorusKnotBufferGeometry) SetParameters(v js.Value) {
	tkbg.Set("parameters", v)
}
func (tkbg *TorusKnotBufferGeometry) Type() string {
	return tkbg.Get("type").String()
}
func (tkbg *TorusKnotBufferGeometry) SetType(v string) {
	tkbg.Set("type", v)
}
func (tkbg *TorusKnotBufferGeometry) UserData() js.Value {
	return tkbg.Get("userData")
}
func (tkbg *TorusKnotBufferGeometry) SetUserData(v js.Value) {
	tkbg.Set("userData", v)
}
func (tkbg *TorusKnotBufferGeometry) Uuid() string {
	return tkbg.Get("uuid").String()
}
func (tkbg *TorusKnotBufferGeometry) SetUuid(v string) {
	tkbg.Set("uuid", v)
}
func (tkbg *TorusKnotBufferGeometry) MaxIndex() int {
	return tkbg.Get("MaxIndex").Int()
}
func (tkbg *TorusKnotBufferGeometry) SetMaxIndex(v int) {
	tkbg.Set("MaxIndex", v)
}
func (tkbg *TorusKnotBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("addAttribute", name, attribute)}
}
func (tkbg *TorusKnotBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return tkbg.Call("addAttribute", name, array, itemSize)
}
func (tkbg *TorusKnotBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	tkbg.Call("addDrawCall", start, count, indexOffset)
}
func (tkbg *TorusKnotBufferGeometry) AddEventListener(typ string, listener js.Value) {
	tkbg.Call("addEventListener", typ, listener)
}
func (tkbg *TorusKnotBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	tkbg.Call("addGroup", start, count, materialIndex)
}
func (tkbg *TorusKnotBufferGeometry) AddIndex(index js.Value) {
	tkbg.Call("addIndex", index)
}
func (tkbg *TorusKnotBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("applyMatrix", matrix)}
}
func (tkbg *TorusKnotBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("center")}
}
func (tkbg *TorusKnotBufferGeometry) ClearDrawCalls() {
	tkbg.Call("clearDrawCalls")
}
func (tkbg *TorusKnotBufferGeometry) ClearGroups() {
	tkbg.Call("clearGroups")
}
func (tkbg *TorusKnotBufferGeometry) Clone() this {
	return this(tkbg.Call("clone"))
}
func (tkbg *TorusKnotBufferGeometry) ComputeBoundingBox() {
	tkbg.Call("computeBoundingBox")
}
func (tkbg *TorusKnotBufferGeometry) ComputeBoundingSphere() {
	tkbg.Call("computeBoundingSphere")
}
func (tkbg *TorusKnotBufferGeometry) ComputeVertexNormals() {
	tkbg.Call("computeVertexNormals")
}
func (tkbg *TorusKnotBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(tkbg.Call("copy", source))
}
func (tkbg *TorusKnotBufferGeometry) DispatchEvent(event js.Value) {
	tkbg.Call("dispatchEvent", event)
}
func (tkbg *TorusKnotBufferGeometry) Dispose() {
	tkbg.Call("dispose")
}
func (tkbg *TorusKnotBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("fromDirectGeometry", geometry)}
}
func (tkbg *TorusKnotBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("fromGeometry", geometry, settings)}
}
func (tkbg *TorusKnotBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(tkbg.Call("getAttribute", name))
}
func (tkbg *TorusKnotBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: tkbg.Call("getIndex")}
}
func (tkbg *TorusKnotBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tkbg.Call("hasEventListener", typ, listener).Bool()
}
func (tkbg *TorusKnotBufferGeometry) LookAt(v *math.Vector3) {
	tkbg.Call("lookAt", v)
}
func (tkbg *TorusKnotBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("merge", geometry, offset)}
}
func (tkbg *TorusKnotBufferGeometry) NormalizeNormals() {
	tkbg.Call("normalizeNormals")
}
func (tkbg *TorusKnotBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("removeAttribute", name)}
}
func (tkbg *TorusKnotBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	tkbg.Call("removeEventListener", typ, listener)
}
func (tkbg *TorusKnotBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("rotateX", angle)}
}
func (tkbg *TorusKnotBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("rotateY", angle)}
}
func (tkbg *TorusKnotBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("rotateZ", angle)}
}
func (tkbg *TorusKnotBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("scale", x, y, z)}
}
func (tkbg *TorusKnotBufferGeometry) SetDrawRange(start int, count int) {
	tkbg.Call("setDrawRange", start, count)
}
func (tkbg *TorusKnotBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("setFromObject", object)}
}
func (tkbg *TorusKnotBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("setFromPoints", points)}
}
func (tkbg *TorusKnotBufferGeometry) SetIndex(index core.BufferAttribute) {
	tkbg.Call("setIndex", index)
}
func (tkbg *TorusKnotBufferGeometry) ToJSON() js.Value {
	return tkbg.Call("toJSON")
}
func (tkbg *TorusKnotBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("toNonIndexed")}
}
func (tkbg *TorusKnotBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: tkbg.Call("translate", x, y, z)}
}
func (tkbg *TorusKnotBufferGeometry) UpdateFromObject(object *core.Object3D) {
	tkbg.Call("updateFromObject", object)
}

type TorusKnotGeometry struct {
	js.Value
}

func NewTorusKnotGeometry(radius int, tube int, tubularSegments int, radialSegments int, p int, q int) *TorusKnotGeometry {
	return &TorusKnotGeometry{Value: get("TorusKnotGeometry").New(radius, tube, tubularSegments, radialSegments, p, q)}
}
func (tkg *TorusKnotGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: tkg.Get("animation")}
}
func (tkg *TorusKnotGeometry) SetAnimation(v *animation.AnimationClip) {
	tkg.Set("animation", v)
}
func (tkg *TorusKnotGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(tkg.Get("animations"))
}
func (tkg *TorusKnotGeometry) SetAnimations(v []animation.AnimationClip) {
	tkg.Set("animations", v)
}
func (tkg *TorusKnotGeometry) Bones() []objects.Bone {
	return []objects.Bone(tkg.Get("bones"))
}
func (tkg *TorusKnotGeometry) SetBones(v []objects.Bone) {
	tkg.Set("bones", v)
}
func (tkg *TorusKnotGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: tkg.Get("boundingBox")}
}
func (tkg *TorusKnotGeometry) SetBoundingBox(v *math.Box3) {
	tkg.Set("boundingBox", v)
}
func (tkg *TorusKnotGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: tkg.Get("boundingSphere")}
}
func (tkg *TorusKnotGeometry) SetBoundingSphere(v *math.Sphere) {
	tkg.Set("boundingSphere", v)
}
func (tkg *TorusKnotGeometry) Colors() []math.Color {
	return []math.Color(tkg.Get("colors"))
}
func (tkg *TorusKnotGeometry) SetColors(v []math.Color) {
	tkg.Set("colors", v)
}
func (tkg *TorusKnotGeometry) ColorsNeedUpdate() bool {
	return tkg.Get("colorsNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetColorsNeedUpdate(v bool) {
	tkg.Set("colorsNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) ElementsNeedUpdate() bool {
	return tkg.Get("elementsNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetElementsNeedUpdate(v bool) {
	tkg.Set("elementsNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(tkg.Get("faceVertexUvs"))
}
func (tkg *TorusKnotGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	tkg.Set("faceVertexUvs", v)
}
func (tkg *TorusKnotGeometry) Faces() []core.Face3 {
	return []core.Face3(tkg.Get("faces"))
}
func (tkg *TorusKnotGeometry) SetFaces(v []core.Face3) {
	tkg.Set("faces", v)
}
func (tkg *TorusKnotGeometry) GroupsNeedUpdate() bool {
	return tkg.Get("groupsNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetGroupsNeedUpdate(v bool) {
	tkg.Set("groupsNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) Id() int {
	return tkg.Get("id").Int()
}
func (tkg *TorusKnotGeometry) SetId(v int) {
	tkg.Set("id", v)
}
func (tkg *TorusKnotGeometry) LineDistances() []int {
	return []int(tkg.Get("lineDistances"))
}
func (tkg *TorusKnotGeometry) SetLineDistances(v []int) {
	tkg.Set("lineDistances", v)
}
func (tkg *TorusKnotGeometry) LineDistancesNeedUpdate() bool {
	return tkg.Get("lineDistancesNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetLineDistancesNeedUpdate(v bool) {
	tkg.Set("lineDistancesNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(tkg.Get("morphNormals"))
}
func (tkg *TorusKnotGeometry) SetMorphNormals(v []core.MorphNormals) {
	tkg.Set("morphNormals", v)
}
func (tkg *TorusKnotGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(tkg.Get("morphTargets"))
}
func (tkg *TorusKnotGeometry) SetMorphTargets(v []core.MorphTarget) {
	tkg.Set("morphTargets", v)
}
func (tkg *TorusKnotGeometry) Name() string {
	return tkg.Get("name").String()
}
func (tkg *TorusKnotGeometry) SetName(v string) {
	tkg.Set("name", v)
}
func (tkg *TorusKnotGeometry) NormalsNeedUpdate() bool {
	return tkg.Get("normalsNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetNormalsNeedUpdate(v bool) {
	tkg.Set("normalsNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) Parameters() js.Value {
	return tkg.Get("parameters")
}
func (tkg *TorusKnotGeometry) SetParameters(v js.Value) {
	tkg.Set("parameters", v)
}
func (tkg *TorusKnotGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(tkg.Get("skinIndices"))
}
func (tkg *TorusKnotGeometry) SetSkinIndices(v []math.Vector4) {
	tkg.Set("skinIndices", v)
}
func (tkg *TorusKnotGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(tkg.Get("skinWeights"))
}
func (tkg *TorusKnotGeometry) SetSkinWeights(v []math.Vector4) {
	tkg.Set("skinWeights", v)
}
func (tkg *TorusKnotGeometry) Type() string {
	return tkg.Get("type").String()
}
func (tkg *TorusKnotGeometry) SetType(v string) {
	tkg.Set("type", v)
}
func (tkg *TorusKnotGeometry) Uuid() string {
	return tkg.Get("uuid").String()
}
func (tkg *TorusKnotGeometry) SetUuid(v string) {
	tkg.Set("uuid", v)
}
func (tkg *TorusKnotGeometry) UvsNeedUpdate() bool {
	return tkg.Get("uvsNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetUvsNeedUpdate(v bool) {
	tkg.Set("uvsNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(tkg.Get("vertices"))
}
func (tkg *TorusKnotGeometry) SetVertices(v []math.Vector3) {
	tkg.Set("vertices", v)
}
func (tkg *TorusKnotGeometry) VerticesNeedUpdate() bool {
	return tkg.Get("verticesNeedUpdate").Bool()
}
func (tkg *TorusKnotGeometry) SetVerticesNeedUpdate(v bool) {
	tkg.Set("verticesNeedUpdate", v)
}
func (tkg *TorusKnotGeometry) AddEventListener(typ string, listener js.Value) {
	tkg.Call("addEventListener", typ, listener)
}
func (tkg *TorusKnotGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("applyMatrix", matrix)}
}
func (tkg *TorusKnotGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: tkg.Call("center")}
}
func (tkg *TorusKnotGeometry) Clone() this {
	return this(tkg.Call("clone"))
}
func (tkg *TorusKnotGeometry) ComputeBoundingBox() {
	tkg.Call("computeBoundingBox")
}
func (tkg *TorusKnotGeometry) ComputeBoundingSphere() {
	tkg.Call("computeBoundingSphere")
}
func (tkg *TorusKnotGeometry) ComputeFaceNormals() {
	tkg.Call("computeFaceNormals")
}
func (tkg *TorusKnotGeometry) ComputeFlatVertexNormals() {
	tkg.Call("computeFlatVertexNormals")
}
func (tkg *TorusKnotGeometry) ComputeMorphNormals() {
	tkg.Call("computeMorphNormals")
}
func (tkg *TorusKnotGeometry) ComputeVertexNormals(areaWeighted bool) {
	tkg.Call("computeVertexNormals", areaWeighted)
}
func (tkg *TorusKnotGeometry) Copy(source *core.Geometry) this {
	return this(tkg.Call("copy", source))
}
func (tkg *TorusKnotGeometry) DispatchEvent(event js.Value) {
	tkg.Call("dispatchEvent", event)
}
func (tkg *TorusKnotGeometry) Dispose() {
	tkg.Call("dispose")
}
func (tkg *TorusKnotGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("fromBufferGeometry", geometry)}
}
func (tkg *TorusKnotGeometry) HasEventListener(typ string, listener js.Value) bool {
	return tkg.Call("hasEventListener", typ, listener).Bool()
}
func (tkg *TorusKnotGeometry) LookAt(vector *math.Vector3) {
	tkg.Call("lookAt", vector)
}
func (tkg *TorusKnotGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	tkg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (tkg *TorusKnotGeometry) MergeMesh(mesh *objects.Mesh) {
	tkg.Call("mergeMesh", mesh)
}
func (tkg *TorusKnotGeometry) MergeVertices() int {
	return tkg.Call("mergeVertices").Int()
}
func (tkg *TorusKnotGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: tkg.Call("normalize")}
}
func (tkg *TorusKnotGeometry) RemoveEventListener(typ string, listener js.Value) {
	tkg.Call("removeEventListener", typ, listener)
}
func (tkg *TorusKnotGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("rotateX", angle)}
}
func (tkg *TorusKnotGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("rotateY", angle)}
}
func (tkg *TorusKnotGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("rotateZ", angle)}
}
func (tkg *TorusKnotGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("scale", x, y, z)}
}
func (tkg *TorusKnotGeometry) SetFromPoints(points Array) this {
	return this(tkg.Call("setFromPoints", points))
}
func (tkg *TorusKnotGeometry) SortFacesByMaterialIndex() {
	tkg.Call("sortFacesByMaterialIndex")
}
func (tkg *TorusKnotGeometry) ToJSON() js.Value {
	return tkg.Call("toJSON")
}
func (tkg *TorusKnotGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: tkg.Call("translate", x, y, z)}
}

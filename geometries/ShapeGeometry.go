package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/extras/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type ShapeBufferGeometry struct {
	js.Value
}

func NewShapeBufferGeometry(shapes core.Shape, curveSegments int) *ShapeBufferGeometry {
	return &ShapeBufferGeometry{Value: get("ShapeBufferGeometry").New(shapes, curveSegments)}
}
func (sbg *ShapeBufferGeometry) Attributes() js.Value {
	return sbg.Get("attributes")
}
func (sbg *ShapeBufferGeometry) SetAttributes(v js.Value) {
	sbg.Set("attributes", v)
}
func (sbg *ShapeBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: sbg.Get("boundingBox")}
}
func (sbg *ShapeBufferGeometry) SetBoundingBox(v *math.Box3) {
	sbg.Set("boundingBox", v)
}
func (sbg *ShapeBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: sbg.Get("boundingSphere")}
}
func (sbg *ShapeBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	sbg.Set("boundingSphere", v)
}
func (sbg *ShapeBufferGeometry) DrawRange() js.Value {
	return sbg.Get("drawRange")
}
func (sbg *ShapeBufferGeometry) SetDrawRange(v js.Value) {
	sbg.Set("drawRange", v)
}
func (sbg *ShapeBufferGeometry) Drawcalls() js.Value {
	return sbg.Get("drawcalls")
}
func (sbg *ShapeBufferGeometry) SetDrawcalls(v js.Value) {
	sbg.Set("drawcalls", v)
}
func (sbg *ShapeBufferGeometry) Groups() []js.Value {
	return []js.Value(sbg.Get("groups"))
}
func (sbg *ShapeBufferGeometry) SetGroups(v []js.Value) {
	sbg.Set("groups", v)
}
func (sbg *ShapeBufferGeometry) Id() int {
	return sbg.Get("id").Int()
}
func (sbg *ShapeBufferGeometry) SetId(v int) {
	sbg.Set("id", v)
}
func (sbg *ShapeBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: sbg.Get("index")}
}
func (sbg *ShapeBufferGeometry) SetIndex(v *core.BufferAttribute) {
	sbg.Set("index", v)
}
func (sbg *ShapeBufferGeometry) MorphAttributes() js.Value {
	return sbg.Get("morphAttributes")
}
func (sbg *ShapeBufferGeometry) SetMorphAttributes(v js.Value) {
	sbg.Set("morphAttributes", v)
}
func (sbg *ShapeBufferGeometry) Name() string {
	return sbg.Get("name").String()
}
func (sbg *ShapeBufferGeometry) SetName(v string) {
	sbg.Set("name", v)
}
func (sbg *ShapeBufferGeometry) Offsets() js.Value {
	return sbg.Get("offsets")
}
func (sbg *ShapeBufferGeometry) SetOffsets(v js.Value) {
	sbg.Set("offsets", v)
}
func (sbg *ShapeBufferGeometry) Type() string {
	return sbg.Get("type").String()
}
func (sbg *ShapeBufferGeometry) SetType(v string) {
	sbg.Set("type", v)
}
func (sbg *ShapeBufferGeometry) UserData() js.Value {
	return sbg.Get("userData")
}
func (sbg *ShapeBufferGeometry) SetUserData(v js.Value) {
	sbg.Set("userData", v)
}
func (sbg *ShapeBufferGeometry) Uuid() string {
	return sbg.Get("uuid").String()
}
func (sbg *ShapeBufferGeometry) SetUuid(v string) {
	sbg.Set("uuid", v)
}
func (sbg *ShapeBufferGeometry) MaxIndex() int {
	return sbg.Get("MaxIndex").Int()
}
func (sbg *ShapeBufferGeometry) SetMaxIndex(v int) {
	sbg.Set("MaxIndex", v)
}
func (sbg *ShapeBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("addAttribute", name, attribute)}
}
func (sbg *ShapeBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return sbg.Call("addAttribute", name, array, itemSize)
}
func (sbg *ShapeBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	sbg.Call("addDrawCall", start, count, indexOffset)
}
func (sbg *ShapeBufferGeometry) AddEventListener(typ string, listener js.Value) {
	sbg.Call("addEventListener", typ, listener)
}
func (sbg *ShapeBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	sbg.Call("addGroup", start, count, materialIndex)
}
func (sbg *ShapeBufferGeometry) AddIndex(index js.Value) {
	sbg.Call("addIndex", index)
}
func (sbg *ShapeBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("applyMatrix", matrix)}
}
func (sbg *ShapeBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("center")}
}
func (sbg *ShapeBufferGeometry) ClearDrawCalls() {
	sbg.Call("clearDrawCalls")
}
func (sbg *ShapeBufferGeometry) ClearGroups() {
	sbg.Call("clearGroups")
}
func (sbg *ShapeBufferGeometry) Clone() this {
	return this(sbg.Call("clone"))
}
func (sbg *ShapeBufferGeometry) ComputeBoundingBox() {
	sbg.Call("computeBoundingBox")
}
func (sbg *ShapeBufferGeometry) ComputeBoundingSphere() {
	sbg.Call("computeBoundingSphere")
}
func (sbg *ShapeBufferGeometry) ComputeVertexNormals() {
	sbg.Call("computeVertexNormals")
}
func (sbg *ShapeBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(sbg.Call("copy", source))
}
func (sbg *ShapeBufferGeometry) DispatchEvent(event js.Value) {
	sbg.Call("dispatchEvent", event)
}
func (sbg *ShapeBufferGeometry) Dispose() {
	sbg.Call("dispose")
}
func (sbg *ShapeBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("fromDirectGeometry", geometry)}
}
func (sbg *ShapeBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("fromGeometry", geometry, settings)}
}
func (sbg *ShapeBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(sbg.Call("getAttribute", name))
}
func (sbg *ShapeBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: sbg.Call("getIndex")}
}
func (sbg *ShapeBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return sbg.Call("hasEventListener", typ, listener).Bool()
}
func (sbg *ShapeBufferGeometry) LookAt(v *math.Vector3) {
	sbg.Call("lookAt", v)
}
func (sbg *ShapeBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("merge", geometry, offset)}
}
func (sbg *ShapeBufferGeometry) NormalizeNormals() {
	sbg.Call("normalizeNormals")
}
func (sbg *ShapeBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("removeAttribute", name)}
}
func (sbg *ShapeBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	sbg.Call("removeEventListener", typ, listener)
}
func (sbg *ShapeBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("rotateX", angle)}
}
func (sbg *ShapeBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("rotateY", angle)}
}
func (sbg *ShapeBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("rotateZ", angle)}
}
func (sbg *ShapeBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("scale", x, y, z)}
}
func (sbg *ShapeBufferGeometry) SetDrawRange(start int, count int) {
	sbg.Call("setDrawRange", start, count)
}
func (sbg *ShapeBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("setFromObject", object)}
}
func (sbg *ShapeBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("setFromPoints", points)}
}
func (sbg *ShapeBufferGeometry) SetIndex(index core.BufferAttribute) {
	sbg.Call("setIndex", index)
}
func (sbg *ShapeBufferGeometry) ToJSON() js.Value {
	return sbg.Call("toJSON")
}
func (sbg *ShapeBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("toNonIndexed")}
}
func (sbg *ShapeBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: sbg.Call("translate", x, y, z)}
}
func (sbg *ShapeBufferGeometry) UpdateFromObject(object *core.Object3D) {
	sbg.Call("updateFromObject", object)
}

type ShapeGeometry struct {
	js.Value
}

func NewShapeGeometry(shapes core.Shape, curveSegments int) *ShapeGeometry {
	return &ShapeGeometry{Value: get("ShapeGeometry").New(shapes, curveSegments)}
}
func (sg *ShapeGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: sg.Get("animation")}
}
func (sg *ShapeGeometry) SetAnimation(v *animation.AnimationClip) {
	sg.Set("animation", v)
}
func (sg *ShapeGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(sg.Get("animations"))
}
func (sg *ShapeGeometry) SetAnimations(v []animation.AnimationClip) {
	sg.Set("animations", v)
}
func (sg *ShapeGeometry) Bones() []objects.Bone {
	return []objects.Bone(sg.Get("bones"))
}
func (sg *ShapeGeometry) SetBones(v []objects.Bone) {
	sg.Set("bones", v)
}
func (sg *ShapeGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: sg.Get("boundingBox")}
}
func (sg *ShapeGeometry) SetBoundingBox(v *math.Box3) {
	sg.Set("boundingBox", v)
}
func (sg *ShapeGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: sg.Get("boundingSphere")}
}
func (sg *ShapeGeometry) SetBoundingSphere(v *math.Sphere) {
	sg.Set("boundingSphere", v)
}
func (sg *ShapeGeometry) Colors() []math.Color {
	return []math.Color(sg.Get("colors"))
}
func (sg *ShapeGeometry) SetColors(v []math.Color) {
	sg.Set("colors", v)
}
func (sg *ShapeGeometry) ColorsNeedUpdate() bool {
	return sg.Get("colorsNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetColorsNeedUpdate(v bool) {
	sg.Set("colorsNeedUpdate", v)
}
func (sg *ShapeGeometry) ElementsNeedUpdate() bool {
	return sg.Get("elementsNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetElementsNeedUpdate(v bool) {
	sg.Set("elementsNeedUpdate", v)
}
func (sg *ShapeGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(sg.Get("faceVertexUvs"))
}
func (sg *ShapeGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	sg.Set("faceVertexUvs", v)
}
func (sg *ShapeGeometry) Faces() []core.Face3 {
	return []core.Face3(sg.Get("faces"))
}
func (sg *ShapeGeometry) SetFaces(v []core.Face3) {
	sg.Set("faces", v)
}
func (sg *ShapeGeometry) GroupsNeedUpdate() bool {
	return sg.Get("groupsNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetGroupsNeedUpdate(v bool) {
	sg.Set("groupsNeedUpdate", v)
}
func (sg *ShapeGeometry) Id() int {
	return sg.Get("id").Int()
}
func (sg *ShapeGeometry) SetId(v int) {
	sg.Set("id", v)
}
func (sg *ShapeGeometry) LineDistances() []int {
	return []int(sg.Get("lineDistances"))
}
func (sg *ShapeGeometry) SetLineDistances(v []int) {
	sg.Set("lineDistances", v)
}
func (sg *ShapeGeometry) LineDistancesNeedUpdate() bool {
	return sg.Get("lineDistancesNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetLineDistancesNeedUpdate(v bool) {
	sg.Set("lineDistancesNeedUpdate", v)
}
func (sg *ShapeGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(sg.Get("morphNormals"))
}
func (sg *ShapeGeometry) SetMorphNormals(v []core.MorphNormals) {
	sg.Set("morphNormals", v)
}
func (sg *ShapeGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(sg.Get("morphTargets"))
}
func (sg *ShapeGeometry) SetMorphTargets(v []core.MorphTarget) {
	sg.Set("morphTargets", v)
}
func (sg *ShapeGeometry) Name() string {
	return sg.Get("name").String()
}
func (sg *ShapeGeometry) SetName(v string) {
	sg.Set("name", v)
}
func (sg *ShapeGeometry) NormalsNeedUpdate() bool {
	return sg.Get("normalsNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetNormalsNeedUpdate(v bool) {
	sg.Set("normalsNeedUpdate", v)
}
func (sg *ShapeGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(sg.Get("skinIndices"))
}
func (sg *ShapeGeometry) SetSkinIndices(v []math.Vector4) {
	sg.Set("skinIndices", v)
}
func (sg *ShapeGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(sg.Get("skinWeights"))
}
func (sg *ShapeGeometry) SetSkinWeights(v []math.Vector4) {
	sg.Set("skinWeights", v)
}
func (sg *ShapeGeometry) Type() string {
	return sg.Get("type").String()
}
func (sg *ShapeGeometry) SetType(v string) {
	sg.Set("type", v)
}
func (sg *ShapeGeometry) Uuid() string {
	return sg.Get("uuid").String()
}
func (sg *ShapeGeometry) SetUuid(v string) {
	sg.Set("uuid", v)
}
func (sg *ShapeGeometry) UvsNeedUpdate() bool {
	return sg.Get("uvsNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetUvsNeedUpdate(v bool) {
	sg.Set("uvsNeedUpdate", v)
}
func (sg *ShapeGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(sg.Get("vertices"))
}
func (sg *ShapeGeometry) SetVertices(v []math.Vector3) {
	sg.Set("vertices", v)
}
func (sg *ShapeGeometry) VerticesNeedUpdate() bool {
	return sg.Get("verticesNeedUpdate").Bool()
}
func (sg *ShapeGeometry) SetVerticesNeedUpdate(v bool) {
	sg.Set("verticesNeedUpdate", v)
}
func (sg *ShapeGeometry) AddEventListener(typ string, listener js.Value) {
	sg.Call("addEventListener", typ, listener)
}
func (sg *ShapeGeometry) AddShape(shape *core.Shape, options js.Value) {
	sg.Call("addShape", shape, options)
}
func (sg *ShapeGeometry) AddShapeList(shapes []core.Shape, options js.Value) *ShapeGeometry {
	return &ShapeGeometry{Value: sg.Call("addShapeList", shapes, options)}
}
func (sg *ShapeGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: sg.Call("applyMatrix", matrix)}
}
func (sg *ShapeGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: sg.Call("center")}
}
func (sg *ShapeGeometry) Clone() this {
	return this(sg.Call("clone"))
}
func (sg *ShapeGeometry) ComputeBoundingBox() {
	sg.Call("computeBoundingBox")
}
func (sg *ShapeGeometry) ComputeBoundingSphere() {
	sg.Call("computeBoundingSphere")
}
func (sg *ShapeGeometry) ComputeFaceNormals() {
	sg.Call("computeFaceNormals")
}
func (sg *ShapeGeometry) ComputeFlatVertexNormals() {
	sg.Call("computeFlatVertexNormals")
}
func (sg *ShapeGeometry) ComputeMorphNormals() {
	sg.Call("computeMorphNormals")
}
func (sg *ShapeGeometry) ComputeVertexNormals(areaWeighted bool) {
	sg.Call("computeVertexNormals", areaWeighted)
}
func (sg *ShapeGeometry) Copy(source *core.Geometry) this {
	return this(sg.Call("copy", source))
}
func (sg *ShapeGeometry) DispatchEvent(event js.Value) {
	sg.Call("dispatchEvent", event)
}
func (sg *ShapeGeometry) Dispose() {
	sg.Call("dispose")
}
func (sg *ShapeGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: sg.Call("fromBufferGeometry", geometry)}
}
func (sg *ShapeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return sg.Call("hasEventListener", typ, listener).Bool()
}
func (sg *ShapeGeometry) LookAt(vector *math.Vector3) {
	sg.Call("lookAt", vector)
}
func (sg *ShapeGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	sg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (sg *ShapeGeometry) MergeMesh(mesh *objects.Mesh) {
	sg.Call("mergeMesh", mesh)
}
func (sg *ShapeGeometry) MergeVertices() int {
	return sg.Call("mergeVertices").Int()
}
func (sg *ShapeGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: sg.Call("normalize")}
}
func (sg *ShapeGeometry) RemoveEventListener(typ string, listener js.Value) {
	sg.Call("removeEventListener", typ, listener)
}
func (sg *ShapeGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("rotateX", angle)}
}
func (sg *ShapeGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("rotateY", angle)}
}
func (sg *ShapeGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("rotateZ", angle)}
}
func (sg *ShapeGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("scale", x, y, z)}
}
func (sg *ShapeGeometry) SetFromPoints(points Array) this {
	return this(sg.Call("setFromPoints", points))
}
func (sg *ShapeGeometry) SortFacesByMaterialIndex() {
	sg.Call("sortFacesByMaterialIndex")
}
func (sg *ShapeGeometry) ToJSON() js.Value {
	return sg.Call("toJSON")
}
func (sg *ShapeGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: sg.Call("translate", x, y, z)}
}

package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

type CylinderBufferGeometry struct {
	js.Value
}

func NewCylinderBufferGeometry(radiusTop int, radiusBottom int, height int, radialSegments int, heightSegments int, openEnded bool, thetaStart int, thetaLength int) *CylinderBufferGeometry {
	return &CylinderBufferGeometry{Value: get("CylinderBufferGeometry").New(radiusTop, radiusBottom, height, radialSegments, heightSegments, openEnded, thetaStart, thetaLength)}
}
func (cbg *CylinderBufferGeometry) Attributes() js.Value {
	return cbg.Get("attributes")
}
func (cbg *CylinderBufferGeometry) SetAttributes(v js.Value) {
	cbg.Set("attributes", v)
}
func (cbg *CylinderBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: cbg.Get("boundingBox")}
}
func (cbg *CylinderBufferGeometry) SetBoundingBox(v *math.Box3) {
	cbg.Set("boundingBox", v)
}
func (cbg *CylinderBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: cbg.Get("boundingSphere")}
}
func (cbg *CylinderBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	cbg.Set("boundingSphere", v)
}
func (cbg *CylinderBufferGeometry) DrawRange() js.Value {
	return cbg.Get("drawRange")
}
func (cbg *CylinderBufferGeometry) SetDrawRange(v js.Value) {
	cbg.Set("drawRange", v)
}
func (cbg *CylinderBufferGeometry) Drawcalls() js.Value {
	return cbg.Get("drawcalls")
}
func (cbg *CylinderBufferGeometry) SetDrawcalls(v js.Value) {
	cbg.Set("drawcalls", v)
}
func (cbg *CylinderBufferGeometry) Groups() []js.Value {
	return []js.Value(cbg.Get("groups"))
}
func (cbg *CylinderBufferGeometry) SetGroups(v []js.Value) {
	cbg.Set("groups", v)
}
func (cbg *CylinderBufferGeometry) Id() int {
	return cbg.Get("id").Int()
}
func (cbg *CylinderBufferGeometry) SetId(v int) {
	cbg.Set("id", v)
}
func (cbg *CylinderBufferGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: cbg.Get("index")}
}
func (cbg *CylinderBufferGeometry) SetIndex(v *core.BufferAttribute) {
	cbg.Set("index", v)
}
func (cbg *CylinderBufferGeometry) MorphAttributes() js.Value {
	return cbg.Get("morphAttributes")
}
func (cbg *CylinderBufferGeometry) SetMorphAttributes(v js.Value) {
	cbg.Set("morphAttributes", v)
}
func (cbg *CylinderBufferGeometry) Name() string {
	return cbg.Get("name").String()
}
func (cbg *CylinderBufferGeometry) SetName(v string) {
	cbg.Set("name", v)
}
func (cbg *CylinderBufferGeometry) Offsets() js.Value {
	return cbg.Get("offsets")
}
func (cbg *CylinderBufferGeometry) SetOffsets(v js.Value) {
	cbg.Set("offsets", v)
}
func (cbg *CylinderBufferGeometry) Parameters() js.Value {
	return cbg.Get("parameters")
}
func (cbg *CylinderBufferGeometry) SetParameters(v js.Value) {
	cbg.Set("parameters", v)
}
func (cbg *CylinderBufferGeometry) Type() string {
	return cbg.Get("type").String()
}
func (cbg *CylinderBufferGeometry) SetType(v string) {
	cbg.Set("type", v)
}
func (cbg *CylinderBufferGeometry) UserData() js.Value {
	return cbg.Get("userData")
}
func (cbg *CylinderBufferGeometry) SetUserData(v js.Value) {
	cbg.Set("userData", v)
}
func (cbg *CylinderBufferGeometry) Uuid() string {
	return cbg.Get("uuid").String()
}
func (cbg *CylinderBufferGeometry) SetUuid(v string) {
	cbg.Set("uuid", v)
}
func (cbg *CylinderBufferGeometry) MaxIndex() int {
	return cbg.Get("MaxIndex").Int()
}
func (cbg *CylinderBufferGeometry) SetMaxIndex(v int) {
	cbg.Set("MaxIndex", v)
}
func (cbg *CylinderBufferGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("addAttribute", name, attribute)}
}
func (cbg *CylinderBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return cbg.Call("addAttribute", name, array, itemSize)
}
func (cbg *CylinderBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	cbg.Call("addDrawCall", start, count, indexOffset)
}
func (cbg *CylinderBufferGeometry) AddEventListener(typ string, listener js.Value) {
	cbg.Call("addEventListener", typ, listener)
}
func (cbg *CylinderBufferGeometry) AddGroup(start int, count int, materialIndex int) {
	cbg.Call("addGroup", start, count, materialIndex)
}
func (cbg *CylinderBufferGeometry) AddIndex(index js.Value) {
	cbg.Call("addIndex", index)
}
func (cbg *CylinderBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("applyMatrix", matrix)}
}
func (cbg *CylinderBufferGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("center")}
}
func (cbg *CylinderBufferGeometry) ClearDrawCalls() {
	cbg.Call("clearDrawCalls")
}
func (cbg *CylinderBufferGeometry) ClearGroups() {
	cbg.Call("clearGroups")
}
func (cbg *CylinderBufferGeometry) Clone() this {
	return this(cbg.Call("clone"))
}
func (cbg *CylinderBufferGeometry) ComputeBoundingBox() {
	cbg.Call("computeBoundingBox")
}
func (cbg *CylinderBufferGeometry) ComputeBoundingSphere() {
	cbg.Call("computeBoundingSphere")
}
func (cbg *CylinderBufferGeometry) ComputeVertexNormals() {
	cbg.Call("computeVertexNormals")
}
func (cbg *CylinderBufferGeometry) Copy(source *core.BufferGeometry) this {
	return this(cbg.Call("copy", source))
}
func (cbg *CylinderBufferGeometry) DispatchEvent(event js.Value) {
	cbg.Call("dispatchEvent", event)
}
func (cbg *CylinderBufferGeometry) Dispose() {
	cbg.Call("dispose")
}
func (cbg *CylinderBufferGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("fromDirectGeometry", geometry)}
}
func (cbg *CylinderBufferGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("fromGeometry", geometry, settings)}
}
func (cbg *CylinderBufferGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(cbg.Call("getAttribute", name))
}
func (cbg *CylinderBufferGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: cbg.Call("getIndex")}
}
func (cbg *CylinderBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cbg.Call("hasEventListener", typ, listener).Bool()
}
func (cbg *CylinderBufferGeometry) LookAt(v *math.Vector3) {
	cbg.Call("lookAt", v)
}
func (cbg *CylinderBufferGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("merge", geometry, offset)}
}
func (cbg *CylinderBufferGeometry) NormalizeNormals() {
	cbg.Call("normalizeNormals")
}
func (cbg *CylinderBufferGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("removeAttribute", name)}
}
func (cbg *CylinderBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	cbg.Call("removeEventListener", typ, listener)
}
func (cbg *CylinderBufferGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateX", angle)}
}
func (cbg *CylinderBufferGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateY", angle)}
}
func (cbg *CylinderBufferGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("rotateZ", angle)}
}
func (cbg *CylinderBufferGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("scale", x, y, z)}
}
func (cbg *CylinderBufferGeometry) SetDrawRange(start int, count int) {
	cbg.Call("setDrawRange", start, count)
}
func (cbg *CylinderBufferGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("setFromObject", object)}
}
func (cbg *CylinderBufferGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("setFromPoints", points)}
}
func (cbg *CylinderBufferGeometry) SetIndex(index core.BufferAttribute) {
	cbg.Call("setIndex", index)
}
func (cbg *CylinderBufferGeometry) ToJSON() js.Value {
	return cbg.Call("toJSON")
}
func (cbg *CylinderBufferGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("toNonIndexed")}
}
func (cbg *CylinderBufferGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: cbg.Call("translate", x, y, z)}
}
func (cbg *CylinderBufferGeometry) UpdateFromObject(object *core.Object3D) {
	cbg.Call("updateFromObject", object)
}

type CylinderGeometry struct {
	js.Value
}

func NewCylinderGeometry(radiusTop int, radiusBottom int, height int, radiusSegments int, heightSegments int, openEnded bool, thetaStart int, thetaLength int) *CylinderGeometry {
	return &CylinderGeometry{Value: get("CylinderGeometry").New(radiusTop, radiusBottom, height, radiusSegments, heightSegments, openEnded, thetaStart, thetaLength)}
}
func (cg *CylinderGeometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: cg.Get("animation")}
}
func (cg *CylinderGeometry) SetAnimation(v *animation.AnimationClip) {
	cg.Set("animation", v)
}
func (cg *CylinderGeometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(cg.Get("animations"))
}
func (cg *CylinderGeometry) SetAnimations(v []animation.AnimationClip) {
	cg.Set("animations", v)
}
func (cg *CylinderGeometry) Bones() []objects.Bone {
	return []objects.Bone(cg.Get("bones"))
}
func (cg *CylinderGeometry) SetBones(v []objects.Bone) {
	cg.Set("bones", v)
}
func (cg *CylinderGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: cg.Get("boundingBox")}
}
func (cg *CylinderGeometry) SetBoundingBox(v *math.Box3) {
	cg.Set("boundingBox", v)
}
func (cg *CylinderGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: cg.Get("boundingSphere")}
}
func (cg *CylinderGeometry) SetBoundingSphere(v *math.Sphere) {
	cg.Set("boundingSphere", v)
}
func (cg *CylinderGeometry) Colors() []math.Color {
	return []math.Color(cg.Get("colors"))
}
func (cg *CylinderGeometry) SetColors(v []math.Color) {
	cg.Set("colors", v)
}
func (cg *CylinderGeometry) ColorsNeedUpdate() bool {
	return cg.Get("colorsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetColorsNeedUpdate(v bool) {
	cg.Set("colorsNeedUpdate", v)
}
func (cg *CylinderGeometry) ElementsNeedUpdate() bool {
	return cg.Get("elementsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetElementsNeedUpdate(v bool) {
	cg.Set("elementsNeedUpdate", v)
}
func (cg *CylinderGeometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(cg.Get("faceVertexUvs"))
}
func (cg *CylinderGeometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	cg.Set("faceVertexUvs", v)
}
func (cg *CylinderGeometry) Faces() []core.Face3 {
	return []core.Face3(cg.Get("faces"))
}
func (cg *CylinderGeometry) SetFaces(v []core.Face3) {
	cg.Set("faces", v)
}
func (cg *CylinderGeometry) GroupsNeedUpdate() bool {
	return cg.Get("groupsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetGroupsNeedUpdate(v bool) {
	cg.Set("groupsNeedUpdate", v)
}
func (cg *CylinderGeometry) Id() int {
	return cg.Get("id").Int()
}
func (cg *CylinderGeometry) SetId(v int) {
	cg.Set("id", v)
}
func (cg *CylinderGeometry) LineDistances() []int {
	return []int(cg.Get("lineDistances"))
}
func (cg *CylinderGeometry) SetLineDistances(v []int) {
	cg.Set("lineDistances", v)
}
func (cg *CylinderGeometry) LineDistancesNeedUpdate() bool {
	return cg.Get("lineDistancesNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetLineDistancesNeedUpdate(v bool) {
	cg.Set("lineDistancesNeedUpdate", v)
}
func (cg *CylinderGeometry) MorphNormals() []core.MorphNormals {
	return []core.MorphNormals(cg.Get("morphNormals"))
}
func (cg *CylinderGeometry) SetMorphNormals(v []core.MorphNormals) {
	cg.Set("morphNormals", v)
}
func (cg *CylinderGeometry) MorphTargets() []core.MorphTarget {
	return []core.MorphTarget(cg.Get("morphTargets"))
}
func (cg *CylinderGeometry) SetMorphTargets(v []core.MorphTarget) {
	cg.Set("morphTargets", v)
}
func (cg *CylinderGeometry) Name() string {
	return cg.Get("name").String()
}
func (cg *CylinderGeometry) SetName(v string) {
	cg.Set("name", v)
}
func (cg *CylinderGeometry) NormalsNeedUpdate() bool {
	return cg.Get("normalsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetNormalsNeedUpdate(v bool) {
	cg.Set("normalsNeedUpdate", v)
}
func (cg *CylinderGeometry) Parameters() js.Value {
	return cg.Get("parameters")
}
func (cg *CylinderGeometry) SetParameters(v js.Value) {
	cg.Set("parameters", v)
}
func (cg *CylinderGeometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(cg.Get("skinIndices"))
}
func (cg *CylinderGeometry) SetSkinIndices(v []math.Vector4) {
	cg.Set("skinIndices", v)
}
func (cg *CylinderGeometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(cg.Get("skinWeights"))
}
func (cg *CylinderGeometry) SetSkinWeights(v []math.Vector4) {
	cg.Set("skinWeights", v)
}
func (cg *CylinderGeometry) Type() string {
	return cg.Get("type").String()
}
func (cg *CylinderGeometry) SetType(v string) {
	cg.Set("type", v)
}
func (cg *CylinderGeometry) Uuid() string {
	return cg.Get("uuid").String()
}
func (cg *CylinderGeometry) SetUuid(v string) {
	cg.Set("uuid", v)
}
func (cg *CylinderGeometry) UvsNeedUpdate() bool {
	return cg.Get("uvsNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetUvsNeedUpdate(v bool) {
	cg.Set("uvsNeedUpdate", v)
}
func (cg *CylinderGeometry) Vertices() []math.Vector3 {
	return []math.Vector3(cg.Get("vertices"))
}
func (cg *CylinderGeometry) SetVertices(v []math.Vector3) {
	cg.Set("vertices", v)
}
func (cg *CylinderGeometry) VerticesNeedUpdate() bool {
	return cg.Get("verticesNeedUpdate").Bool()
}
func (cg *CylinderGeometry) SetVerticesNeedUpdate(v bool) {
	cg.Set("verticesNeedUpdate", v)
}
func (cg *CylinderGeometry) AddEventListener(typ string, listener js.Value) {
	cg.Call("addEventListener", typ, listener)
}
func (cg *CylinderGeometry) ApplyMatrix(matrix *math.Matrix4) *core.Geometry {
	return &core.Geometry{Value: cg.Call("applyMatrix", matrix)}
}
func (cg *CylinderGeometry) Center() *core.Geometry {
	return &core.Geometry{Value: cg.Call("center")}
}
func (cg *CylinderGeometry) Clone() this {
	return this(cg.Call("clone"))
}
func (cg *CylinderGeometry) ComputeBoundingBox() {
	cg.Call("computeBoundingBox")
}
func (cg *CylinderGeometry) ComputeBoundingSphere() {
	cg.Call("computeBoundingSphere")
}
func (cg *CylinderGeometry) ComputeFaceNormals() {
	cg.Call("computeFaceNormals")
}
func (cg *CylinderGeometry) ComputeFlatVertexNormals() {
	cg.Call("computeFlatVertexNormals")
}
func (cg *CylinderGeometry) ComputeMorphNormals() {
	cg.Call("computeMorphNormals")
}
func (cg *CylinderGeometry) ComputeVertexNormals(areaWeighted bool) {
	cg.Call("computeVertexNormals", areaWeighted)
}
func (cg *CylinderGeometry) Copy(source *core.Geometry) this {
	return this(cg.Call("copy", source))
}
func (cg *CylinderGeometry) DispatchEvent(event js.Value) {
	cg.Call("dispatchEvent", event)
}
func (cg *CylinderGeometry) Dispose() {
	cg.Call("dispose")
}
func (cg *CylinderGeometry) FromBufferGeometry(geometry *core.BufferGeometry) *core.Geometry {
	return &core.Geometry{Value: cg.Call("fromBufferGeometry", geometry)}
}
func (cg *CylinderGeometry) HasEventListener(typ string, listener js.Value) bool {
	return cg.Call("hasEventListener", typ, listener).Bool()
}
func (cg *CylinderGeometry) LookAt(vector *math.Vector3) {
	cg.Call("lookAt", vector)
}
func (cg *CylinderGeometry) Merge(geometry *core.Geometry, matrix math.Matrix, materialIndexOffset int) {
	cg.Call("merge", geometry, matrix, materialIndexOffset)
}
func (cg *CylinderGeometry) MergeMesh(mesh *objects.Mesh) {
	cg.Call("mergeMesh", mesh)
}
func (cg *CylinderGeometry) MergeVertices() int {
	return cg.Call("mergeVertices").Int()
}
func (cg *CylinderGeometry) Normalize() *core.Geometry {
	return &core.Geometry{Value: cg.Call("normalize")}
}
func (cg *CylinderGeometry) RemoveEventListener(typ string, listener js.Value) {
	cg.Call("removeEventListener", typ, listener)
}
func (cg *CylinderGeometry) RotateX(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateX", angle)}
}
func (cg *CylinderGeometry) RotateY(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateY", angle)}
}
func (cg *CylinderGeometry) RotateZ(angle int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("rotateZ", angle)}
}
func (cg *CylinderGeometry) Scale(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("scale", x, y, z)}
}
func (cg *CylinderGeometry) SetFromPoints(points Array) this {
	return this(cg.Call("setFromPoints", points))
}
func (cg *CylinderGeometry) SortFacesByMaterialIndex() {
	cg.Call("sortFacesByMaterialIndex")
}
func (cg *CylinderGeometry) ToJSON() js.Value {
	return cg.Call("toJSON")
}
func (cg *CylinderGeometry) Translate(x int, y int, z int) *core.Geometry {
	return &core.Geometry{Value: cg.Call("translate", x, y, z)}
}

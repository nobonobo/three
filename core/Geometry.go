package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/animation"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
)

func GeometryIdCount() int {
	return get("GeometryIdCount").Int()
}
func SetGeometryIdCount(v int) {
	set("GeometryIdCount", v)
}

type MorphColor interface {
}
type MorphNormals interface {
}
type MorphTarget interface {
}
type Geometry struct {
	js.Value
}

func NewGeometry() *Geometry {
	return &Geometry{Value: get("Geometry").New()}
}
func (g *Geometry) Animation() *animation.AnimationClip {
	return &animation.AnimationClip{Value: g.Get("animation")}
}
func (g *Geometry) SetAnimation(v *animation.AnimationClip) {
	g.Set("animation", v)
}
func (g *Geometry) Animations() []animation.AnimationClip {
	return []animation.AnimationClip(g.Get("animations"))
}
func (g *Geometry) SetAnimations(v []animation.AnimationClip) {
	g.Set("animations", v)
}
func (g *Geometry) Bones() []objects.Bone {
	return []objects.Bone(g.Get("bones"))
}
func (g *Geometry) SetBones(v []objects.Bone) {
	g.Set("bones", v)
}
func (g *Geometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: g.Get("boundingBox")}
}
func (g *Geometry) SetBoundingBox(v *math.Box3) {
	g.Set("boundingBox", v)
}
func (g *Geometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: g.Get("boundingSphere")}
}
func (g *Geometry) SetBoundingSphere(v *math.Sphere) {
	g.Set("boundingSphere", v)
}
func (g *Geometry) Colors() []math.Color {
	return []math.Color(g.Get("colors"))
}
func (g *Geometry) SetColors(v []math.Color) {
	g.Set("colors", v)
}
func (g *Geometry) ColorsNeedUpdate() bool {
	return g.Get("colorsNeedUpdate").Bool()
}
func (g *Geometry) SetColorsNeedUpdate(v bool) {
	g.Set("colorsNeedUpdate", v)
}
func (g *Geometry) ElementsNeedUpdate() bool {
	return g.Get("elementsNeedUpdate").Bool()
}
func (g *Geometry) SetElementsNeedUpdate(v bool) {
	g.Set("elementsNeedUpdate", v)
}
func (g *Geometry) FaceVertexUvs() [][][]math.Vector2 {
	return [][][]math.Vector2(g.Get("faceVertexUvs"))
}
func (g *Geometry) SetFaceVertexUvs(v [][][]math.Vector2) {
	g.Set("faceVertexUvs", v)
}
func (g *Geometry) Faces() []Face3 {
	return []Face3(g.Get("faces"))
}
func (g *Geometry) SetFaces(v []Face3) {
	g.Set("faces", v)
}
func (g *Geometry) GroupsNeedUpdate() bool {
	return g.Get("groupsNeedUpdate").Bool()
}
func (g *Geometry) SetGroupsNeedUpdate(v bool) {
	g.Set("groupsNeedUpdate", v)
}
func (g *Geometry) Id() int {
	return g.Get("id").Int()
}
func (g *Geometry) SetId(v int) {
	g.Set("id", v)
}
func (g *Geometry) LineDistances() []int {
	return []int(g.Get("lineDistances"))
}
func (g *Geometry) SetLineDistances(v []int) {
	g.Set("lineDistances", v)
}
func (g *Geometry) LineDistancesNeedUpdate() bool {
	return g.Get("lineDistancesNeedUpdate").Bool()
}
func (g *Geometry) SetLineDistancesNeedUpdate(v bool) {
	g.Set("lineDistancesNeedUpdate", v)
}
func (g *Geometry) MorphNormals() []MorphNormals {
	return []MorphNormals(g.Get("morphNormals"))
}
func (g *Geometry) SetMorphNormals(v []MorphNormals) {
	g.Set("morphNormals", v)
}
func (g *Geometry) MorphTargets() []MorphTarget {
	return []MorphTarget(g.Get("morphTargets"))
}
func (g *Geometry) SetMorphTargets(v []MorphTarget) {
	g.Set("morphTargets", v)
}
func (g *Geometry) Name() string {
	return g.Get("name").String()
}
func (g *Geometry) SetName(v string) {
	g.Set("name", v)
}
func (g *Geometry) NormalsNeedUpdate() bool {
	return g.Get("normalsNeedUpdate").Bool()
}
func (g *Geometry) SetNormalsNeedUpdate(v bool) {
	g.Set("normalsNeedUpdate", v)
}
func (g *Geometry) SkinIndices() []math.Vector4 {
	return []math.Vector4(g.Get("skinIndices"))
}
func (g *Geometry) SetSkinIndices(v []math.Vector4) {
	g.Set("skinIndices", v)
}
func (g *Geometry) SkinWeights() []math.Vector4 {
	return []math.Vector4(g.Get("skinWeights"))
}
func (g *Geometry) SetSkinWeights(v []math.Vector4) {
	g.Set("skinWeights", v)
}
func (g *Geometry) Type() string {
	return g.Get("type").String()
}
func (g *Geometry) SetType(v string) {
	g.Set("type", v)
}
func (g *Geometry) Uuid() string {
	return g.Get("uuid").String()
}
func (g *Geometry) SetUuid(v string) {
	g.Set("uuid", v)
}
func (g *Geometry) UvsNeedUpdate() bool {
	return g.Get("uvsNeedUpdate").Bool()
}
func (g *Geometry) SetUvsNeedUpdate(v bool) {
	g.Set("uvsNeedUpdate", v)
}
func (g *Geometry) Vertices() []math.Vector3 {
	return []math.Vector3(g.Get("vertices"))
}
func (g *Geometry) SetVertices(v []math.Vector3) {
	g.Set("vertices", v)
}
func (g *Geometry) VerticesNeedUpdate() bool {
	return g.Get("verticesNeedUpdate").Bool()
}
func (g *Geometry) SetVerticesNeedUpdate(v bool) {
	g.Set("verticesNeedUpdate", v)
}
func (g *Geometry) AddEventListener(typ string, listener js.Value) {
	g.Call("addEventListener", typ, listener)
}
func (g *Geometry) ApplyMatrix(matrix *math.Matrix4) *Geometry {
	return &Geometry{Value: g.Call("applyMatrix", matrix)}
}
func (g *Geometry) Center() *Geometry {
	return &Geometry{Value: g.Call("center")}
}
func (g *Geometry) Clone() this {
	return this(g.Call("clone"))
}
func (g *Geometry) ComputeBoundingBox() {
	g.Call("computeBoundingBox")
}
func (g *Geometry) ComputeBoundingSphere() {
	g.Call("computeBoundingSphere")
}
func (g *Geometry) ComputeFaceNormals() {
	g.Call("computeFaceNormals")
}
func (g *Geometry) ComputeFlatVertexNormals() {
	g.Call("computeFlatVertexNormals")
}
func (g *Geometry) ComputeMorphNormals() {
	g.Call("computeMorphNormals")
}
func (g *Geometry) ComputeVertexNormals(areaWeighted bool) {
	g.Call("computeVertexNormals", areaWeighted)
}
func (g *Geometry) Copy(source *Geometry) this {
	return this(g.Call("copy", source))
}
func (g *Geometry) DispatchEvent(event js.Value) {
	g.Call("dispatchEvent", event)
}
func (g *Geometry) Dispose() {
	g.Call("dispose")
}
func (g *Geometry) FromBufferGeometry(geometry *BufferGeometry) *Geometry {
	return &Geometry{Value: g.Call("fromBufferGeometry", geometry)}
}
func (g *Geometry) HasEventListener(typ string, listener js.Value) bool {
	return g.Call("hasEventListener", typ, listener).Bool()
}
func (g *Geometry) LookAt(vector *math.Vector3) {
	g.Call("lookAt", vector)
}
func (g *Geometry) Merge(geometry *Geometry, matrix math.Matrix, materialIndexOffset int) {
	g.Call("merge", geometry, matrix, materialIndexOffset)
}
func (g *Geometry) MergeMesh(mesh *objects.Mesh) {
	g.Call("mergeMesh", mesh)
}
func (g *Geometry) MergeVertices() int {
	return g.Call("mergeVertices").Int()
}
func (g *Geometry) Normalize() *Geometry {
	return &Geometry{Value: g.Call("normalize")}
}
func (g *Geometry) RemoveEventListener(typ string, listener js.Value) {
	g.Call("removeEventListener", typ, listener)
}
func (g *Geometry) RotateX(angle int) *Geometry {
	return &Geometry{Value: g.Call("rotateX", angle)}
}
func (g *Geometry) RotateY(angle int) *Geometry {
	return &Geometry{Value: g.Call("rotateY", angle)}
}
func (g *Geometry) RotateZ(angle int) *Geometry {
	return &Geometry{Value: g.Call("rotateZ", angle)}
}
func (g *Geometry) Scale(x int, y int, z int) *Geometry {
	return &Geometry{Value: g.Call("scale", x, y, z)}
}
func (g *Geometry) SetFromPoints(points Array) this {
	return this(g.Call("setFromPoints", points))
}
func (g *Geometry) SortFacesByMaterialIndex() {
	g.Call("sortFacesByMaterialIndex")
}
func (g *Geometry) ToJSON() js.Value {
	return g.Call("toJSON")
}
func (g *Geometry) Translate(x int, y int, z int) *Geometry {
	return &Geometry{Value: g.Call("translate", x, y, z)}
}

package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type EdgesGeometry struct {
	js.Value
}

func NewEdgesGeometry(geometry core.BufferGeometry, thresholdAngle int) *EdgesGeometry {
	return &EdgesGeometry{Value: get("EdgesGeometry").New(geometry, thresholdAngle)}
}
func (eg *EdgesGeometry) Attributes() js.Value {
	return eg.Get("attributes")
}
func (eg *EdgesGeometry) SetAttributes(v js.Value) {
	eg.Set("attributes", v)
}
func (eg *EdgesGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: eg.Get("boundingBox")}
}
func (eg *EdgesGeometry) SetBoundingBox(v *math.Box3) {
	eg.Set("boundingBox", v)
}
func (eg *EdgesGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: eg.Get("boundingSphere")}
}
func (eg *EdgesGeometry) SetBoundingSphere(v *math.Sphere) {
	eg.Set("boundingSphere", v)
}
func (eg *EdgesGeometry) DrawRange() js.Value {
	return eg.Get("drawRange")
}
func (eg *EdgesGeometry) SetDrawRange(v js.Value) {
	eg.Set("drawRange", v)
}
func (eg *EdgesGeometry) Drawcalls() js.Value {
	return eg.Get("drawcalls")
}
func (eg *EdgesGeometry) SetDrawcalls(v js.Value) {
	eg.Set("drawcalls", v)
}
func (eg *EdgesGeometry) Groups() []js.Value {
	return []js.Value(eg.Get("groups"))
}
func (eg *EdgesGeometry) SetGroups(v []js.Value) {
	eg.Set("groups", v)
}
func (eg *EdgesGeometry) Id() int {
	return eg.Get("id").Int()
}
func (eg *EdgesGeometry) SetId(v int) {
	eg.Set("id", v)
}
func (eg *EdgesGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: eg.Get("index")}
}
func (eg *EdgesGeometry) SetIndex(v *core.BufferAttribute) {
	eg.Set("index", v)
}
func (eg *EdgesGeometry) MorphAttributes() js.Value {
	return eg.Get("morphAttributes")
}
func (eg *EdgesGeometry) SetMorphAttributes(v js.Value) {
	eg.Set("morphAttributes", v)
}
func (eg *EdgesGeometry) Name() string {
	return eg.Get("name").String()
}
func (eg *EdgesGeometry) SetName(v string) {
	eg.Set("name", v)
}
func (eg *EdgesGeometry) Offsets() js.Value {
	return eg.Get("offsets")
}
func (eg *EdgesGeometry) SetOffsets(v js.Value) {
	eg.Set("offsets", v)
}
func (eg *EdgesGeometry) Type() string {
	return eg.Get("type").String()
}
func (eg *EdgesGeometry) SetType(v string) {
	eg.Set("type", v)
}
func (eg *EdgesGeometry) UserData() js.Value {
	return eg.Get("userData")
}
func (eg *EdgesGeometry) SetUserData(v js.Value) {
	eg.Set("userData", v)
}
func (eg *EdgesGeometry) Uuid() string {
	return eg.Get("uuid").String()
}
func (eg *EdgesGeometry) SetUuid(v string) {
	eg.Set("uuid", v)
}
func (eg *EdgesGeometry) MaxIndex() int {
	return eg.Get("MaxIndex").Int()
}
func (eg *EdgesGeometry) SetMaxIndex(v int) {
	eg.Set("MaxIndex", v)
}
func (eg *EdgesGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("addAttribute", name, attribute)}
}
func (eg *EdgesGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return eg.Call("addAttribute", name, array, itemSize)
}
func (eg *EdgesGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	eg.Call("addDrawCall", start, count, indexOffset)
}
func (eg *EdgesGeometry) AddEventListener(typ string, listener js.Value) {
	eg.Call("addEventListener", typ, listener)
}
func (eg *EdgesGeometry) AddGroup(start int, count int, materialIndex int) {
	eg.Call("addGroup", start, count, materialIndex)
}
func (eg *EdgesGeometry) AddIndex(index js.Value) {
	eg.Call("addIndex", index)
}
func (eg *EdgesGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("applyMatrix", matrix)}
}
func (eg *EdgesGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("center")}
}
func (eg *EdgesGeometry) ClearDrawCalls() {
	eg.Call("clearDrawCalls")
}
func (eg *EdgesGeometry) ClearGroups() {
	eg.Call("clearGroups")
}
func (eg *EdgesGeometry) Clone() this {
	return this(eg.Call("clone"))
}
func (eg *EdgesGeometry) ComputeBoundingBox() {
	eg.Call("computeBoundingBox")
}
func (eg *EdgesGeometry) ComputeBoundingSphere() {
	eg.Call("computeBoundingSphere")
}
func (eg *EdgesGeometry) ComputeVertexNormals() {
	eg.Call("computeVertexNormals")
}
func (eg *EdgesGeometry) Copy(source *core.BufferGeometry) this {
	return this(eg.Call("copy", source))
}
func (eg *EdgesGeometry) DispatchEvent(event js.Value) {
	eg.Call("dispatchEvent", event)
}
func (eg *EdgesGeometry) Dispose() {
	eg.Call("dispose")
}
func (eg *EdgesGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("fromDirectGeometry", geometry)}
}
func (eg *EdgesGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("fromGeometry", geometry, settings)}
}
func (eg *EdgesGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(eg.Call("getAttribute", name))
}
func (eg *EdgesGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: eg.Call("getIndex")}
}
func (eg *EdgesGeometry) HasEventListener(typ string, listener js.Value) bool {
	return eg.Call("hasEventListener", typ, listener).Bool()
}
func (eg *EdgesGeometry) LookAt(v *math.Vector3) {
	eg.Call("lookAt", v)
}
func (eg *EdgesGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("merge", geometry, offset)}
}
func (eg *EdgesGeometry) NormalizeNormals() {
	eg.Call("normalizeNormals")
}
func (eg *EdgesGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("removeAttribute", name)}
}
func (eg *EdgesGeometry) RemoveEventListener(typ string, listener js.Value) {
	eg.Call("removeEventListener", typ, listener)
}
func (eg *EdgesGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("rotateX", angle)}
}
func (eg *EdgesGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("rotateY", angle)}
}
func (eg *EdgesGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("rotateZ", angle)}
}
func (eg *EdgesGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("scale", x, y, z)}
}
func (eg *EdgesGeometry) SetDrawRange(start int, count int) {
	eg.Call("setDrawRange", start, count)
}
func (eg *EdgesGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("setFromObject", object)}
}
func (eg *EdgesGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("setFromPoints", points)}
}
func (eg *EdgesGeometry) SetIndex(index core.BufferAttribute) {
	eg.Call("setIndex", index)
}
func (eg *EdgesGeometry) ToJSON() js.Value {
	return eg.Call("toJSON")
}
func (eg *EdgesGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("toNonIndexed")}
}
func (eg *EdgesGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: eg.Call("translate", x, y, z)}
}
func (eg *EdgesGeometry) UpdateFromObject(object *core.Object3D) {
	eg.Call("updateFromObject", object)
}

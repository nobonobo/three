package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type InstancedBufferGeometry struct {
	js.Value
}

func NewInstancedBufferGeometry() *InstancedBufferGeometry {
	return &InstancedBufferGeometry{Value: get("InstancedBufferGeometry").New()}
}
func (ibg *InstancedBufferGeometry) Attributes() js.Value {
	return ibg.Get("attributes")
}
func (ibg *InstancedBufferGeometry) SetAttributes(v js.Value) {
	ibg.Set("attributes", v)
}
func (ibg *InstancedBufferGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: ibg.Get("boundingBox")}
}
func (ibg *InstancedBufferGeometry) SetBoundingBox(v *math.Box3) {
	ibg.Set("boundingBox", v)
}
func (ibg *InstancedBufferGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: ibg.Get("boundingSphere")}
}
func (ibg *InstancedBufferGeometry) SetBoundingSphere(v *math.Sphere) {
	ibg.Set("boundingSphere", v)
}
func (ibg *InstancedBufferGeometry) DrawRange() js.Value {
	return ibg.Get("drawRange")
}
func (ibg *InstancedBufferGeometry) SetDrawRange(v js.Value) {
	ibg.Set("drawRange", v)
}
func (ibg *InstancedBufferGeometry) Drawcalls() js.Value {
	return ibg.Get("drawcalls")
}
func (ibg *InstancedBufferGeometry) SetDrawcalls(v js.Value) {
	ibg.Set("drawcalls", v)
}
func (ibg *InstancedBufferGeometry) Groups() []js.Value {
	return []js.Value(ibg.Get("groups"))
}
func (ibg *InstancedBufferGeometry) SetGroups(v []js.Value) {
	ibg.Set("groups", v)
}
func (ibg *InstancedBufferGeometry) Id() int {
	return ibg.Get("id").Int()
}
func (ibg *InstancedBufferGeometry) SetId(v int) {
	ibg.Set("id", v)
}
func (ibg *InstancedBufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: ibg.Get("index")}
}
func (ibg *InstancedBufferGeometry) SetIndex(v *BufferAttribute) {
	ibg.Set("index", v)
}
func (ibg *InstancedBufferGeometry) MaxInstancedCount() int {
	return ibg.Get("maxInstancedCount").Int()
}
func (ibg *InstancedBufferGeometry) SetMaxInstancedCount(v int) {
	ibg.Set("maxInstancedCount", v)
}
func (ibg *InstancedBufferGeometry) MorphAttributes() js.Value {
	return ibg.Get("morphAttributes")
}
func (ibg *InstancedBufferGeometry) SetMorphAttributes(v js.Value) {
	ibg.Set("morphAttributes", v)
}
func (ibg *InstancedBufferGeometry) Name() string {
	return ibg.Get("name").String()
}
func (ibg *InstancedBufferGeometry) SetName(v string) {
	ibg.Set("name", v)
}
func (ibg *InstancedBufferGeometry) Offsets() js.Value {
	return ibg.Get("offsets")
}
func (ibg *InstancedBufferGeometry) SetOffsets(v js.Value) {
	ibg.Set("offsets", v)
}
func (ibg *InstancedBufferGeometry) Type() string {
	return ibg.Get("type").String()
}
func (ibg *InstancedBufferGeometry) SetType(v string) {
	ibg.Set("type", v)
}
func (ibg *InstancedBufferGeometry) UserData() js.Value {
	return ibg.Get("userData")
}
func (ibg *InstancedBufferGeometry) SetUserData(v js.Value) {
	ibg.Set("userData", v)
}
func (ibg *InstancedBufferGeometry) Uuid() string {
	return ibg.Get("uuid").String()
}
func (ibg *InstancedBufferGeometry) SetUuid(v string) {
	ibg.Set("uuid", v)
}
func (ibg *InstancedBufferGeometry) MaxIndex() int {
	return ibg.Get("MaxIndex").Int()
}
func (ibg *InstancedBufferGeometry) SetMaxIndex(v int) {
	ibg.Set("MaxIndex", v)
}
func (ibg *InstancedBufferGeometry) AddAttribute(name string, attribute BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("addAttribute", name, attribute)}
}
func (ibg *InstancedBufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return ibg.Call("addAttribute", name, array, itemSize)
}
func (ibg *InstancedBufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	ibg.Call("addDrawCall", start, count, indexOffset)
}
func (ibg *InstancedBufferGeometry) AddEventListener(typ string, listener js.Value) {
	ibg.Call("addEventListener", typ, listener)
}
func (ibg *InstancedBufferGeometry) AddGroup(start int, count int, instances int) {
	ibg.Call("addGroup", start, count, instances)
}
func (ibg *InstancedBufferGeometry) AddIndex(index js.Value) {
	ibg.Call("addIndex", index)
}
func (ibg *InstancedBufferGeometry) ApplyMatrix(matrix *math.Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("applyMatrix", matrix)}
}
func (ibg *InstancedBufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("center")}
}
func (ibg *InstancedBufferGeometry) ClearDrawCalls() {
	ibg.Call("clearDrawCalls")
}
func (ibg *InstancedBufferGeometry) ClearGroups() {
	ibg.Call("clearGroups")
}
func (ibg *InstancedBufferGeometry) Clone() this {
	return this(ibg.Call("clone"))
}
func (ibg *InstancedBufferGeometry) ComputeBoundingBox() {
	ibg.Call("computeBoundingBox")
}
func (ibg *InstancedBufferGeometry) ComputeBoundingSphere() {
	ibg.Call("computeBoundingSphere")
}
func (ibg *InstancedBufferGeometry) ComputeVertexNormals() {
	ibg.Call("computeVertexNormals")
}
func (ibg *InstancedBufferGeometry) Copy(source *BufferGeometry) this {
	return this(ibg.Call("copy", source))
}
func (ibg *InstancedBufferGeometry) DispatchEvent(event js.Value) {
	ibg.Call("dispatchEvent", event)
}
func (ibg *InstancedBufferGeometry) Dispose() {
	ibg.Call("dispose")
}
func (ibg *InstancedBufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("fromDirectGeometry", geometry)}
}
func (ibg *InstancedBufferGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("fromGeometry", geometry, settings)}
}
func (ibg *InstancedBufferGeometry) GetAttribute(name string) BufferAttribute {
	return BufferAttribute(ibg.Call("getAttribute", name))
}
func (ibg *InstancedBufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: ibg.Call("getIndex")}
}
func (ibg *InstancedBufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return ibg.Call("hasEventListener", typ, listener).Bool()
}
func (ibg *InstancedBufferGeometry) LookAt(v *math.Vector3) {
	ibg.Call("lookAt", v)
}
func (ibg *InstancedBufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("merge", geometry, offset)}
}
func (ibg *InstancedBufferGeometry) NormalizeNormals() {
	ibg.Call("normalizeNormals")
}
func (ibg *InstancedBufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("removeAttribute", name)}
}
func (ibg *InstancedBufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	ibg.Call("removeEventListener", typ, listener)
}
func (ibg *InstancedBufferGeometry) RotateX(angle int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("rotateX", angle)}
}
func (ibg *InstancedBufferGeometry) RotateY(angle int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("rotateY", angle)}
}
func (ibg *InstancedBufferGeometry) RotateZ(angle int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("rotateZ", angle)}
}
func (ibg *InstancedBufferGeometry) Scale(x int, y int, z int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("scale", x, y, z)}
}
func (ibg *InstancedBufferGeometry) SetDrawRange(start int, count int) {
	ibg.Call("setDrawRange", start, count)
}
func (ibg *InstancedBufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("setFromObject", object)}
}
func (ibg *InstancedBufferGeometry) SetFromPoints(points []math.Vector3) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("setFromPoints", points)}
}
func (ibg *InstancedBufferGeometry) SetIndex(index BufferAttribute) {
	ibg.Call("setIndex", index)
}
func (ibg *InstancedBufferGeometry) ToJSON() js.Value {
	return ibg.Call("toJSON")
}
func (ibg *InstancedBufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("toNonIndexed")}
}
func (ibg *InstancedBufferGeometry) Translate(x int, y int, z int) *BufferGeometry {
	return &BufferGeometry{Value: ibg.Call("translate", x, y, z)}
}
func (ibg *InstancedBufferGeometry) UpdateFromObject(object *Object3D) {
	ibg.Call("updateFromObject", object)
}

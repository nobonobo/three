// Code generated from "core/BufferGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// BufferGeometry extend: [EventDispatcher]
type BufferGeometry struct {
	js.Value
}

func NewBufferGeometry() *BufferGeometry {
	return &BufferGeometry{Value: get("BufferGeometry").New()}
}
func (bg *BufferGeometry) JSValue() js.Value {
	return bg.Value
}
func (bg *BufferGeometry) Attributes() js.Value {
	return bg.Get("attributes")
}
func (bg *BufferGeometry) SetAttributes(v js.Value) {
	bg.Set("attributes", v)
}
func (bg *BufferGeometry) BoundingBox() *Box3 {
	return &Box3{Value: bg.Get("boundingBox")}
}
func (bg *BufferGeometry) SetBoundingBox(v *Box3) {
	bg.Set("boundingBox", v.JSValue())
}
func (bg *BufferGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: bg.Get("boundingSphere")}
}
func (bg *BufferGeometry) SetBoundingSphere(v *Sphere) {
	bg.Set("boundingSphere", v.JSValue())
}
func (bg *BufferGeometry) DrawRange() js.Value {
	return bg.Get("drawRange")
}
func (bg *BufferGeometry) SetDrawRange(v js.Value) {
	bg.Set("drawRange", v)
}
func (bg *BufferGeometry) Drawcalls() js.Value {
	return bg.Get("drawcalls")
}
func (bg *BufferGeometry) SetDrawcalls(v js.Value) {
	bg.Set("drawcalls", v)
}
func (bg *BufferGeometry) Groups() js.Value {
	return bg.Get("groups")
}
func (bg *BufferGeometry) SetGroups(v js.Value) {
	bg.Set("groups", v)
}
func (bg *BufferGeometry) Id() int {
	return bg.Get("id").Int()
}
func (bg *BufferGeometry) SetId(v int) {
	bg.Set("id", v)
}
func (bg *BufferGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: bg.Get("index")}
}
func (bg *BufferGeometry) SetIndex(v *BufferAttribute) {
	bg.Set("index", v.JSValue())
}
func (bg *BufferGeometry) MorphAttributes() js.Value {
	return bg.Get("morphAttributes")
}
func (bg *BufferGeometry) SetMorphAttributes(v js.Value) {
	bg.Set("morphAttributes", v)
}
func (bg *BufferGeometry) Name() string {
	return bg.Get("name").String()
}
func (bg *BufferGeometry) SetName(v string) {
	bg.Set("name", v)
}
func (bg *BufferGeometry) Offsets() js.Value {
	return bg.Get("offsets")
}
func (bg *BufferGeometry) SetOffsets(v js.Value) {
	bg.Set("offsets", v)
}
func (bg *BufferGeometry) Type() string {
	return bg.Get("type").String()
}
func (bg *BufferGeometry) SetType(v string) {
	bg.Set("type", v)
}
func (bg *BufferGeometry) UserData() js.Value {
	return bg.Get("userData")
}
func (bg *BufferGeometry) SetUserData(v js.Value) {
	bg.Set("userData", v)
}
func (bg *BufferGeometry) Uuid() string {
	return bg.Get("uuid").String()
}
func (bg *BufferGeometry) SetUuid(v string) {
	bg.Set("uuid", v)
}
func (bg *BufferGeometry) MaxIndex() int {
	return bg.Get("MaxIndex").Int()
}
func (bg *BufferGeometry) SetMaxIndex(v int) {
	bg.Set("MaxIndex", v)
}
func (bg *BufferGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("addAttribute", name, attribute)}
}
func (bg *BufferGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return bg.Call("addAttribute", name, array, itemSize)
}
func (bg *BufferGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	bg.Call("addDrawCall", start, count, indexOffset)
}
func (bg *BufferGeometry) AddEventListener(typ string, listener js.Value) {
	bg.Call("addEventListener", typ, listener)
}
func (bg *BufferGeometry) AddGroup(start int, count int, materialIndex int) {
	bg.Call("addGroup", start, count, materialIndex)
}
func (bg *BufferGeometry) AddIndex(index js.Value) {
	bg.Call("addIndex", index)
}
func (bg *BufferGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("applyMatrix", matrix)}
}
func (bg *BufferGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("center")}
}
func (bg *BufferGeometry) ClearDrawCalls() {
	bg.Call("clearDrawCalls")
}
func (bg *BufferGeometry) ClearGroups() {
	bg.Call("clearGroups")
}
func (bg *BufferGeometry) Clone() *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("clone")}
}
func (bg *BufferGeometry) ComputeBoundingBox() {
	bg.Call("computeBoundingBox")
}
func (bg *BufferGeometry) ComputeBoundingSphere() {
	bg.Call("computeBoundingSphere")
}
func (bg *BufferGeometry) ComputeVertexNormals() {
	bg.Call("computeVertexNormals")
}
func (bg *BufferGeometry) Copy(source *BufferGeometry) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("copy", source)}
}
func (bg *BufferGeometry) DispatchEvent(event js.Value) {
	bg.Call("dispatchEvent", event)
}
func (bg *BufferGeometry) Dispose() {
	bg.Call("dispose")
}
func (bg *BufferGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("fromDirectGeometry", geometry)}
}
func (bg *BufferGeometry) FromGeometry(geometry Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (bg *BufferGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: bg.Call("getAttribute", name)}
}
func (bg *BufferGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: bg.Call("getIndex")}
}
func (bg *BufferGeometry) HasEventListener(typ string, listener js.Value) bool {
	return bg.Call("hasEventListener", typ, listener).Bool()
}
func (bg *BufferGeometry) LookAt(v *Vector3) {
	bg.Call("lookAt", v.JSValue())
}
func (bg *BufferGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("merge", geometry, offset)}
}
func (bg *BufferGeometry) NormalizeNormals() {
	bg.Call("normalizeNormals")
}
func (bg *BufferGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("removeAttribute", name)}
}
func (bg *BufferGeometry) RemoveEventListener(typ string, listener js.Value) {
	bg.Call("removeEventListener", typ, listener)
}
func (bg *BufferGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("rotateX", angle)}
}
func (bg *BufferGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("rotateY", angle)}
}
func (bg *BufferGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("rotateZ", angle)}
}
func (bg *BufferGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("scale", x, y, z)}
}
func (bg *BufferGeometry) SetDrawRange2(start int, count int) {
	bg.Call("setDrawRange", start, count)
}
func (bg *BufferGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("setFromObject", object)}
}
func (bg *BufferGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("setFromPoints", points)}
}
func (bg *BufferGeometry) SetIndex2(index *BufferAttribute) {
	bg.Call("setIndex", index)
}
func (bg *BufferGeometry) ToJSON() js.Value {
	return bg.Call("toJSON")
}
func (bg *BufferGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("toNonIndexed")}
}
func (bg *BufferGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: bg.Call("translate", x, y, z)}
}
func (bg *BufferGeometry) UpdateFromObject(object *Object3D) {
	bg.Call("updateFromObject", object.JSValue())
}

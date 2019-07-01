// Code generated from "core/BufferGeometry.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type BufferGeometry interface {
	JSValue() js.Value
	Attributes() js.Value
	SetAttributes(v js.Value)
	BoundingBox() *Box3
	SetBoundingBox(v *Box3)
	BoundingSphere() *Sphere
	SetBoundingSphere(v *Sphere)
	DrawRange() js.Value
	SetDrawRange(v js.Value)
	Drawcalls() js.Value
	SetDrawcalls(v js.Value)
	Groups() js.Value
	SetGroups(v js.Value)
	Id() int
	SetId(v int)
	Index() *BufferAttribute
	SetIndex(v *BufferAttribute)
	MorphAttributes() js.Value
	SetMorphAttributes(v js.Value)
	Name() string
	SetName(v string)
	Offsets() js.Value
	SetOffsets(v js.Value)
	Type() string
	SetType(v string)
	UserData() js.Value
	SetUserData(v js.Value)
	Uuid() string
	SetUuid(v string)
	MaxIndex() int
	SetMaxIndex(v int)
	AddAttribute(name string, attribute *BufferAttribute) BufferGeometry
	AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value
	AddDrawCall(start js.Value, count js.Value, indexOffset js.Value)
	AddEventListener(typ string, listener js.Value)
	AddGroup(start int, count int, materialIndex int)
	AddIndex(index js.Value)
	ApplyMatrix(matrix *Matrix4) BufferGeometry
	Center() BufferGeometry
	ClearDrawCalls()
	ClearGroups()
	Clone() BufferGeometry
	ComputeBoundingBox()
	ComputeBoundingSphere()
	ComputeVertexNormals()
	Copy(source BufferGeometry) BufferGeometry
	DispatchEvent(event js.Value)
	Dispose()
	FromDirectGeometry(geometry *DirectGeometry) BufferGeometry
	FromGeometry(geometry Geometry, settings js.Value) BufferGeometry
	GetAttribute(name string) *BufferAttribute
	GetIndex() *BufferAttribute
	HasEventListener(typ string, listener js.Value) bool
	LookAt(v *Vector3)
	Merge(geometry BufferGeometry, offset int) BufferGeometry
	NormalizeNormals()
	RemoveAttribute(name string) BufferGeometry
	RemoveEventListener(typ string, listener js.Value)
	RotateX(angle float64) BufferGeometry
	RotateY(angle float64) BufferGeometry
	RotateZ(angle float64) BufferGeometry
	Scale(x float64, y float64, z float64) BufferGeometry
	SetDrawRange2(start int, count int)
	SetFromObject(object *Object3D) BufferGeometry
	SetFromPoints(points js.Value) BufferGeometry
	SetIndex2(index *BufferAttribute)
	ToJSON() js.Value
	ToNonIndexed() BufferGeometry
	Translate(x float64, y float64, z float64) BufferGeometry
	UpdateFromObject(object *Object3D)
}

// BufferGeometryImpl extend: [EventDispatcher]
type BufferGeometryImpl struct {
	js.Value
}

func NewBufferGeometry() *BufferGeometryImpl {
	return &BufferGeometryImpl{Value: get("BufferGeometry").New()}
}
func (bg *BufferGeometryImpl) JSValue() js.Value {
	return bg.Value
}
func (bg *BufferGeometryImpl) Attributes() js.Value {
	return bg.Get("attributes")
}
func (bg *BufferGeometryImpl) SetAttributes(v js.Value) {
	bg.Set("attributes", v)
}
func (bg *BufferGeometryImpl) BoundingBox() *Box3 {
	return &Box3{Value: bg.Get("boundingBox")}
}
func (bg *BufferGeometryImpl) SetBoundingBox(v *Box3) {
	bg.Set("boundingBox", v.JSValue())
}
func (bg *BufferGeometryImpl) BoundingSphere() *Sphere {
	return &Sphere{Value: bg.Get("boundingSphere")}
}
func (bg *BufferGeometryImpl) SetBoundingSphere(v *Sphere) {
	bg.Set("boundingSphere", v.JSValue())
}
func (bg *BufferGeometryImpl) DrawRange() js.Value {
	return bg.Get("drawRange")
}
func (bg *BufferGeometryImpl) SetDrawRange(v js.Value) {
	bg.Set("drawRange", v)
}
func (bg *BufferGeometryImpl) Drawcalls() js.Value {
	return bg.Get("drawcalls")
}
func (bg *BufferGeometryImpl) SetDrawcalls(v js.Value) {
	bg.Set("drawcalls", v)
}
func (bg *BufferGeometryImpl) Groups() js.Value {
	return bg.Get("groups")
}
func (bg *BufferGeometryImpl) SetGroups(v js.Value) {
	bg.Set("groups", v)
}
func (bg *BufferGeometryImpl) Id() int {
	return bg.Get("id").Int()
}
func (bg *BufferGeometryImpl) SetId(v int) {
	bg.Set("id", v)
}
func (bg *BufferGeometryImpl) Index() *BufferAttribute {
	return &BufferAttribute{Value: bg.Get("index")}
}
func (bg *BufferGeometryImpl) SetIndex(v *BufferAttribute) {
	bg.Set("index", v.JSValue())
}
func (bg *BufferGeometryImpl) MorphAttributes() js.Value {
	return bg.Get("morphAttributes")
}
func (bg *BufferGeometryImpl) SetMorphAttributes(v js.Value) {
	bg.Set("morphAttributes", v)
}
func (bg *BufferGeometryImpl) Name() string {
	return bg.Get("name").String()
}
func (bg *BufferGeometryImpl) SetName(v string) {
	bg.Set("name", v)
}
func (bg *BufferGeometryImpl) Offsets() js.Value {
	return bg.Get("offsets")
}
func (bg *BufferGeometryImpl) SetOffsets(v js.Value) {
	bg.Set("offsets", v)
}
func (bg *BufferGeometryImpl) Type() string {
	return bg.Get("type").String()
}
func (bg *BufferGeometryImpl) SetType(v string) {
	bg.Set("type", v)
}
func (bg *BufferGeometryImpl) UserData() js.Value {
	return bg.Get("userData")
}
func (bg *BufferGeometryImpl) SetUserData(v js.Value) {
	bg.Set("userData", v)
}
func (bg *BufferGeometryImpl) Uuid() string {
	return bg.Get("uuid").String()
}
func (bg *BufferGeometryImpl) SetUuid(v string) {
	bg.Set("uuid", v)
}
func (bg *BufferGeometryImpl) MaxIndex() int {
	return bg.Get("MaxIndex").Int()
}
func (bg *BufferGeometryImpl) SetMaxIndex(v int) {
	bg.Set("MaxIndex", v)
}
func (bg *BufferGeometryImpl) AddAttribute(name string, attribute *BufferAttribute) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("addAttribute", name, attribute)}
}
func (bg *BufferGeometryImpl) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return bg.Call("addAttribute", name, array, itemSize)
}
func (bg *BufferGeometryImpl) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	bg.Call("addDrawCall", start, count, indexOffset)
}
func (bg *BufferGeometryImpl) AddEventListener(typ string, listener js.Value) {
	bg.Call("addEventListener", typ, listener)
}
func (bg *BufferGeometryImpl) AddGroup(start int, count int, materialIndex int) {
	bg.Call("addGroup", start, count, materialIndex)
}
func (bg *BufferGeometryImpl) AddIndex(index js.Value) {
	bg.Call("addIndex", index)
}
func (bg *BufferGeometryImpl) ApplyMatrix(matrix *Matrix4) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("applyMatrix", matrix)}
}
func (bg *BufferGeometryImpl) Center() BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("center")}
}
func (bg *BufferGeometryImpl) ClearDrawCalls() {
	bg.Call("clearDrawCalls")
}
func (bg *BufferGeometryImpl) ClearGroups() {
	bg.Call("clearGroups")
}
func (bg *BufferGeometryImpl) Clone() BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("clone")}
}
func (bg *BufferGeometryImpl) ComputeBoundingBox() {
	bg.Call("computeBoundingBox")
}
func (bg *BufferGeometryImpl) ComputeBoundingSphere() {
	bg.Call("computeBoundingSphere")
}
func (bg *BufferGeometryImpl) ComputeVertexNormals() {
	bg.Call("computeVertexNormals")
}
func (bg *BufferGeometryImpl) Copy(source BufferGeometry) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("copy", source.JSValue())}
}
func (bg *BufferGeometryImpl) DispatchEvent(event js.Value) {
	bg.Call("dispatchEvent", event)
}
func (bg *BufferGeometryImpl) Dispose() {
	bg.Call("dispose")
}
func (bg *BufferGeometryImpl) FromDirectGeometry(geometry *DirectGeometry) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("fromDirectGeometry", geometry)}
}
func (bg *BufferGeometryImpl) FromGeometry(geometry Geometry, settings js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("fromGeometry", geometry.JSValue(), settings)}
}
func (bg *BufferGeometryImpl) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: bg.Call("getAttribute", name)}
}
func (bg *BufferGeometryImpl) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: bg.Call("getIndex")}
}
func (bg *BufferGeometryImpl) HasEventListener(typ string, listener js.Value) bool {
	return bg.Call("hasEventListener", typ, listener).Bool()
}
func (bg *BufferGeometryImpl) LookAt(v *Vector3) {
	bg.Call("lookAt", v.JSValue())
}
func (bg *BufferGeometryImpl) Merge(geometry BufferGeometry, offset int) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("merge", geometry.JSValue(), offset)}
}
func (bg *BufferGeometryImpl) NormalizeNormals() {
	bg.Call("normalizeNormals")
}
func (bg *BufferGeometryImpl) RemoveAttribute(name string) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("removeAttribute", name)}
}
func (bg *BufferGeometryImpl) RemoveEventListener(typ string, listener js.Value) {
	bg.Call("removeEventListener", typ, listener)
}
func (bg *BufferGeometryImpl) RotateX(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("rotateX", angle)}
}
func (bg *BufferGeometryImpl) RotateY(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("rotateY", angle)}
}
func (bg *BufferGeometryImpl) RotateZ(angle float64) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("rotateZ", angle)}
}
func (bg *BufferGeometryImpl) Scale(x float64, y float64, z float64) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("scale", x, y, z)}
}
func (bg *BufferGeometryImpl) SetDrawRange2(start int, count int) {
	bg.Call("setDrawRange", start, count)
}
func (bg *BufferGeometryImpl) SetFromObject(object *Object3D) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("setFromObject", object)}
}
func (bg *BufferGeometryImpl) SetFromPoints(points js.Value) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("setFromPoints", points)}
}
func (bg *BufferGeometryImpl) SetIndex2(index *BufferAttribute) {
	bg.Call("setIndex", index)
}
func (bg *BufferGeometryImpl) ToJSON() js.Value {
	return bg.Call("toJSON")
}
func (bg *BufferGeometryImpl) ToNonIndexed() BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("toNonIndexed")}
}
func (bg *BufferGeometryImpl) Translate(x float64, y float64, z float64) BufferGeometry {
	return &BufferGeometryImpl{Value: bg.Call("translate", x, y, z)}
}
func (bg *BufferGeometryImpl) UpdateFromObject(object *Object3D) {
	bg.Call("updateFromObject", object.JSValue())
}

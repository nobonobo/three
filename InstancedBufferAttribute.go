// Code generated from "core/InstancedBufferAttribute.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// InstancedBufferAttribute extend: [BufferAttribute]
type InstancedBufferAttribute struct {
	js.Value
}

func NewInstancedBufferAttribute(array js.Value, itemSize int, normalized bool, meshPerAttribute float64) *InstancedBufferAttribute {
	return &InstancedBufferAttribute{Value: get("InstancedBufferAttribute").New(array, itemSize, normalized, meshPerAttribute)}
}
func (iba *InstancedBufferAttribute) JSValue() js.Value {
	return iba.Value
}
func (iba *InstancedBufferAttribute) Array() js.Value {
	return iba.Get("array")
}
func (iba *InstancedBufferAttribute) SetArray(v js.Value) {
	iba.Set("array", v)
}
func (iba *InstancedBufferAttribute) Count() int {
	return iba.Get("count").Int()
}
func (iba *InstancedBufferAttribute) SetCount(v int) {
	iba.Set("count", v)
}
func (iba *InstancedBufferAttribute) Dynamic() bool {
	return iba.Get("dynamic").Bool()
}
func (iba *InstancedBufferAttribute) SetDynamic(v bool) {
	iba.Set("dynamic", v)
}
func (iba *InstancedBufferAttribute) ItemSize() int {
	return iba.Get("itemSize").Int()
}
func (iba *InstancedBufferAttribute) SetItemSize(v int) {
	iba.Set("itemSize", v)
}
func (iba *InstancedBufferAttribute) Length() int {
	return iba.Get("length").Int()
}
func (iba *InstancedBufferAttribute) SetLength(v int) {
	iba.Set("length", v)
}
func (iba *InstancedBufferAttribute) MeshPerAttribute() float64 {
	return iba.Get("meshPerAttribute").Float()
}
func (iba *InstancedBufferAttribute) SetMeshPerAttribute(v float64) {
	iba.Set("meshPerAttribute", v)
}
func (iba *InstancedBufferAttribute) NeedsUpdate() bool {
	return iba.Get("needsUpdate").Bool()
}
func (iba *InstancedBufferAttribute) SetNeedsUpdate(v bool) {
	iba.Set("needsUpdate", v)
}
func (iba *InstancedBufferAttribute) Normalized() bool {
	return iba.Get("normalized").Bool()
}
func (iba *InstancedBufferAttribute) SetNormalized(v bool) {
	iba.Set("normalized", v)
}
func (iba *InstancedBufferAttribute) OnUpload() js.Value {
	return iba.Get("onUpload")
}
func (iba *InstancedBufferAttribute) SetOnUpload(v js.Value) {
	iba.Set("onUpload", v)
}
func (iba *InstancedBufferAttribute) UpdateRange() js.Value {
	return iba.Get("updateRange")
}
func (iba *InstancedBufferAttribute) SetUpdateRange(v js.Value) {
	iba.Set("updateRange", v)
}
func (iba *InstancedBufferAttribute) Uuid() string {
	return iba.Get("uuid").String()
}
func (iba *InstancedBufferAttribute) SetUuid(v string) {
	iba.Set("uuid", v)
}
func (iba *InstancedBufferAttribute) Version() float64 {
	return iba.Get("version").Float()
}
func (iba *InstancedBufferAttribute) SetVersion(v float64) {
	iba.Set("version", v)
}
func (iba *InstancedBufferAttribute) Clone() *InstancedBufferAttribute {
	return &InstancedBufferAttribute{Value: iba.Call("clone")}
}
func (iba *InstancedBufferAttribute) Copy(source *BufferAttribute) *InstancedBufferAttribute {
	return &InstancedBufferAttribute{Value: iba.Call("copy", source)}
}
func (iba *InstancedBufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyArray", array)}
}
func (iba *InstancedBufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyAt", index1, attribute, index2)}
}
func (iba *InstancedBufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyColorsArray", colors)}
}
func (iba *InstancedBufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector2sArray", vectors)}
}
func (iba *InstancedBufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector3sArray", vectors)}
}
func (iba *InstancedBufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector4sArray", vectors)}
}
func (iba *InstancedBufferAttribute) GetW(index int) float64 {
	return iba.Call("getW", index).Float()
}
func (iba *InstancedBufferAttribute) GetX(index int) float64 {
	return iba.Call("getX", index).Float()
}
func (iba *InstancedBufferAttribute) GetY(index int) float64 {
	return iba.Call("getY", index).Float()
}
func (iba *InstancedBufferAttribute) GetZ(index int) float64 {
	return iba.Call("getZ", index).Float()
}
func (iba *InstancedBufferAttribute) Set2(value js.Value, offset float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("set", value, offset)}
}
func (iba *InstancedBufferAttribute) SetArray2(array js.Value) {
	iba.Call("setArray", array)
}
func (iba *InstancedBufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setDynamic", dynamic)}
}
func (iba *InstancedBufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *InstancedBufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *InstancedBufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *InstancedBufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *InstancedBufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *InstancedBufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *InstancedBufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setZ", index, z)}
}

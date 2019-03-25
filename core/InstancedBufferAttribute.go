package core

import (
	"github.com/gopherjs/gopherwasm/js"
)

type InstancedBufferAttribute struct {
	js.Value
}

func NewInstancedBufferAttribute(array *ArrayLike, itemSize int, normalized bool, meshPerAttribute int) *InstancedBufferAttribute {
	return &InstancedBufferAttribute{Value: get("InstancedBufferAttribute").New(array, itemSize, normalized, meshPerAttribute)}
}
func (iba *InstancedBufferAttribute) Array() *ArrayLike {
	return &ArrayLike{Value: iba.Get("array")}
}
func (iba *InstancedBufferAttribute) SetArray(v *ArrayLike) {
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
func (iba *InstancedBufferAttribute) MeshPerAttribute() int {
	return iba.Get("meshPerAttribute").Int()
}
func (iba *InstancedBufferAttribute) SetMeshPerAttribute(v int) {
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
func (iba *InstancedBufferAttribute) Version() int {
	return iba.Get("version").Int()
}
func (iba *InstancedBufferAttribute) SetVersion(v int) {
	iba.Set("version", v)
}
func (iba *InstancedBufferAttribute) Clone() this {
	return this(iba.Call("clone"))
}
func (iba *InstancedBufferAttribute) Copy(source *BufferAttribute) this {
	return this(iba.Call("copy", source))
}
func (iba *InstancedBufferAttribute) CopyArray(array *ArrayLike) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyArray", array)}
}
func (iba *InstancedBufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyAt", index1, attribute, index2)}
}
func (iba *InstancedBufferAttribute) CopyColorsArray(colors []js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyColorsArray", colors)}
}
func (iba *InstancedBufferAttribute) CopyVector2sArray(vectors []js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector2sArray", vectors)}
}
func (iba *InstancedBufferAttribute) CopyVector3sArray(vectors []js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector3sArray", vectors)}
}
func (iba *InstancedBufferAttribute) CopyVector4sArray(vectors []js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector4sArray", vectors)}
}
func (iba *InstancedBufferAttribute) GetW(index int) int {
	return iba.Call("getW", index).Int()
}
func (iba *InstancedBufferAttribute) GetX(index int) int {
	return iba.Call("getX", index).Int()
}
func (iba *InstancedBufferAttribute) GetY(index int) int {
	return iba.Call("getY", index).Int()
}
func (iba *InstancedBufferAttribute) GetZ(index int) int {
	return iba.Call("getZ", index).Int()
}
func (iba *InstancedBufferAttribute) Set(value ArrayLike, offset int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("set", value, offset)}
}
func (iba *InstancedBufferAttribute) SetArray(array *ArrayBufferView) {
	iba.Call("setArray", array)
}
func (iba *InstancedBufferAttribute) SetDynamic(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setDynamic", dynamic)}
}
func (iba *InstancedBufferAttribute) SetW(index int, z int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *InstancedBufferAttribute) SetX(index int, x int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *InstancedBufferAttribute) SetXY(index int, x int, y int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *InstancedBufferAttribute) SetXYZ(index int, x int, y int, z int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *InstancedBufferAttribute) SetXYZW(index int, x int, y int, z int, w int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *InstancedBufferAttribute) SetY(index int, y int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *InstancedBufferAttribute) SetZ(index int, z int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setZ", index, z)}
}

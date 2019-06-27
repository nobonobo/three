// Code generated from "core/InstancedInterleavedBuffer.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// InstancedInterleavedBuffer extend: [InterleavedBuffer]
type InstancedInterleavedBuffer struct {
	js.Value
}

func NewInstancedInterleavedBuffer(array js.Value, stride int, meshPerAttribute float64) *InstancedInterleavedBuffer {
	return &InstancedInterleavedBuffer{Value: get("InstancedInterleavedBuffer").New(array, stride, meshPerAttribute)}
}
func (iib *InstancedInterleavedBuffer) JSValue() js.Value {
	return iib.Value
}
func (iib *InstancedInterleavedBuffer) Array() js.Value {
	return iib.Get("array")
}
func (iib *InstancedInterleavedBuffer) SetArray(v js.Value) {
	iib.Set("array", v)
}
func (iib *InstancedInterleavedBuffer) Count() int {
	return iib.Get("count").Int()
}
func (iib *InstancedInterleavedBuffer) SetCount(v int) {
	iib.Set("count", v)
}
func (iib *InstancedInterleavedBuffer) Dynamic() bool {
	return iib.Get("dynamic").Bool()
}
func (iib *InstancedInterleavedBuffer) SetDynamic(v bool) {
	iib.Set("dynamic", v)
}
func (iib *InstancedInterleavedBuffer) Length() int {
	return iib.Get("length").Int()
}
func (iib *InstancedInterleavedBuffer) SetLength(v int) {
	iib.Set("length", v)
}
func (iib *InstancedInterleavedBuffer) MeshPerAttribute() float64 {
	return iib.Get("meshPerAttribute").Float()
}
func (iib *InstancedInterleavedBuffer) SetMeshPerAttribute(v float64) {
	iib.Set("meshPerAttribute", v)
}
func (iib *InstancedInterleavedBuffer) NeedsUpdate() bool {
	return iib.Get("needsUpdate").Bool()
}
func (iib *InstancedInterleavedBuffer) SetNeedsUpdate(v bool) {
	iib.Set("needsUpdate", v)
}
func (iib *InstancedInterleavedBuffer) Stride() int {
	return iib.Get("stride").Int()
}
func (iib *InstancedInterleavedBuffer) SetStride(v int) {
	iib.Set("stride", v)
}
func (iib *InstancedInterleavedBuffer) UpdateRange() js.Value {
	return iib.Get("updateRange")
}
func (iib *InstancedInterleavedBuffer) SetUpdateRange(v js.Value) {
	iib.Set("updateRange", v)
}
func (iib *InstancedInterleavedBuffer) Version() float64 {
	return iib.Get("version").Float()
}
func (iib *InstancedInterleavedBuffer) SetVersion(v float64) {
	iib.Set("version", v)
}
func (iib *InstancedInterleavedBuffer) Clone() *InstancedInterleavedBuffer {
	return &InstancedInterleavedBuffer{Value: iib.Call("clone")}
}
func (iib *InstancedInterleavedBuffer) Copy(source *InterleavedBuffer) *InstancedInterleavedBuffer {
	return &InstancedInterleavedBuffer{Value: iib.Call("copy", source)}
}
func (iib *InstancedInterleavedBuffer) CopyAt(index1 int, attribute *InterleavedBufferAttribute, index2 int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: iib.Call("copyAt", index1, attribute, index2)}
}
func (iib *InstancedInterleavedBuffer) Set2(value js.Value, index int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: iib.Call("set", value, index)}
}
func (iib *InstancedInterleavedBuffer) SetArray2(array js.Value) {
	iib.Call("setArray", array)
}
func (iib *InstancedInterleavedBuffer) SetDynamic2(dynamic bool) *InterleavedBuffer {
	return &InterleavedBuffer{Value: iib.Call("setDynamic", dynamic)}
}

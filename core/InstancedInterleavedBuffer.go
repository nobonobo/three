package core

import (
	"github.com/gopherjs/gopherwasm/js"
)

type InstancedInterleavedBuffer struct {
	js.Value
}

func NewInstancedInterleavedBuffer(array *ArrayLike, stride int, meshPerAttribute int) *InstancedInterleavedBuffer {
	return &InstancedInterleavedBuffer{Value: get("InstancedInterleavedBuffer").New(array, stride, meshPerAttribute)}
}
func (iib *InstancedInterleavedBuffer) Array() *ArrayLike {
	return &ArrayLike{Value: iib.Get("array")}
}
func (iib *InstancedInterleavedBuffer) SetArray(v *ArrayLike) {
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
func (iib *InstancedInterleavedBuffer) MeshPerAttribute() int {
	return iib.Get("meshPerAttribute").Int()
}
func (iib *InstancedInterleavedBuffer) SetMeshPerAttribute(v int) {
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
func (iib *InstancedInterleavedBuffer) Version() int {
	return iib.Get("version").Int()
}
func (iib *InstancedInterleavedBuffer) SetVersion(v int) {
	iib.Set("version", v)
}
func (iib *InstancedInterleavedBuffer) Clone() this {
	return this(iib.Call("clone"))
}
func (iib *InstancedInterleavedBuffer) Copy(source *InterleavedBuffer) this {
	return this(iib.Call("copy", source))
}
func (iib *InstancedInterleavedBuffer) CopyAt(index1 int, attribute *InterleavedBufferAttribute, index2 int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: iib.Call("copyAt", index1, attribute, index2)}
}
func (iib *InstancedInterleavedBuffer) Set(value *ArrayLike, index int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: iib.Call("set", value, index)}
}
func (iib *InstancedInterleavedBuffer) SetArray(array *ArrayBufferView) {
	iib.Call("setArray", array)
}
func (iib *InstancedInterleavedBuffer) SetDynamic(dynamic bool) *InterleavedBuffer {
	return &InterleavedBuffer{Value: iib.Call("setDynamic", dynamic)}
}

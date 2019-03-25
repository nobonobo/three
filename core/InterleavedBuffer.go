package core

import (
	"github.com/gopherjs/gopherwasm/js"
)

type InterleavedBuffer struct {
	js.Value
}

func NewInterleavedBuffer(array *ArrayLike, stride int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: get("InterleavedBuffer").New(array, stride)}
}
func (ib *InterleavedBuffer) Array() *ArrayLike {
	return &ArrayLike{Value: ib.Get("array")}
}
func (ib *InterleavedBuffer) SetArray(v *ArrayLike) {
	ib.Set("array", v)
}
func (ib *InterleavedBuffer) Count() int {
	return ib.Get("count").Int()
}
func (ib *InterleavedBuffer) SetCount(v int) {
	ib.Set("count", v)
}
func (ib *InterleavedBuffer) Dynamic() bool {
	return ib.Get("dynamic").Bool()
}
func (ib *InterleavedBuffer) SetDynamic(v bool) {
	ib.Set("dynamic", v)
}
func (ib *InterleavedBuffer) Length() int {
	return ib.Get("length").Int()
}
func (ib *InterleavedBuffer) SetLength(v int) {
	ib.Set("length", v)
}
func (ib *InterleavedBuffer) NeedsUpdate() bool {
	return ib.Get("needsUpdate").Bool()
}
func (ib *InterleavedBuffer) SetNeedsUpdate(v bool) {
	ib.Set("needsUpdate", v)
}
func (ib *InterleavedBuffer) Stride() int {
	return ib.Get("stride").Int()
}
func (ib *InterleavedBuffer) SetStride(v int) {
	ib.Set("stride", v)
}
func (ib *InterleavedBuffer) UpdateRange() js.Value {
	return ib.Get("updateRange")
}
func (ib *InterleavedBuffer) SetUpdateRange(v js.Value) {
	ib.Set("updateRange", v)
}
func (ib *InterleavedBuffer) Version() int {
	return ib.Get("version").Int()
}
func (ib *InterleavedBuffer) SetVersion(v int) {
	ib.Set("version", v)
}
func (ib *InterleavedBuffer) Clone() this {
	return this(ib.Call("clone"))
}
func (ib *InterleavedBuffer) Copy(source *InterleavedBuffer) this {
	return this(ib.Call("copy", source))
}
func (ib *InterleavedBuffer) CopyAt(index1 int, attribute *InterleavedBufferAttribute, index2 int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: ib.Call("copyAt", index1, attribute, index2)}
}
func (ib *InterleavedBuffer) Set(value *ArrayLike, index int) *InterleavedBuffer {
	return &InterleavedBuffer{Value: ib.Call("set", value, index)}
}
func (ib *InterleavedBuffer) SetArray(array *ArrayBufferView) {
	ib.Call("setArray", array)
}
func (ib *InterleavedBuffer) SetDynamic(dynamic bool) *InterleavedBuffer {
	return &InterleavedBuffer{Value: ib.Call("setDynamic", dynamic)}
}

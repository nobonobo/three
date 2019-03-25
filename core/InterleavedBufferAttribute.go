package core

import (
	"github.com/gopherjs/gopherwasm/js"
)

type InterleavedBufferAttribute struct {
	js.Value
}

func NewInterleavedBufferAttribute(interleavedBuffer *InterleavedBuffer, itemSize int, offset int, normalized bool) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: get("InterleavedBufferAttribute").New(interleavedBuffer, itemSize, offset, normalized)}
}
func (iba *InterleavedBufferAttribute) Array() []js.Value {
	return []js.Value(iba.Get("array"))
}
func (iba *InterleavedBufferAttribute) SetArray(v []js.Value) {
	iba.Set("array", v)
}
func (iba *InterleavedBufferAttribute) Count() int {
	return iba.Get("count").Int()
}
func (iba *InterleavedBufferAttribute) SetCount(v int) {
	iba.Set("count", v)
}
func (iba *InterleavedBufferAttribute) Data() *InterleavedBuffer {
	return &InterleavedBuffer{Value: iba.Get("data")}
}
func (iba *InterleavedBufferAttribute) SetData(v *InterleavedBuffer) {
	iba.Set("data", v)
}
func (iba *InterleavedBufferAttribute) ItemSize() int {
	return iba.Get("itemSize").Int()
}
func (iba *InterleavedBufferAttribute) SetItemSize(v int) {
	iba.Set("itemSize", v)
}
func (iba *InterleavedBufferAttribute) Length() int {
	return iba.Get("length").Int()
}
func (iba *InterleavedBufferAttribute) SetLength(v int) {
	iba.Set("length", v)
}
func (iba *InterleavedBufferAttribute) Normalized() bool {
	return iba.Get("normalized").Bool()
}
func (iba *InterleavedBufferAttribute) SetNormalized(v bool) {
	iba.Set("normalized", v)
}
func (iba *InterleavedBufferAttribute) Offset() int {
	return iba.Get("offset").Int()
}
func (iba *InterleavedBufferAttribute) SetOffset(v int) {
	iba.Set("offset", v)
}
func (iba *InterleavedBufferAttribute) Uuid() string {
	return iba.Get("uuid").String()
}
func (iba *InterleavedBufferAttribute) SetUuid(v string) {
	iba.Set("uuid", v)
}
func (iba *InterleavedBufferAttribute) GetW(index int) int {
	return iba.Call("getW", index).Int()
}
func (iba *InterleavedBufferAttribute) GetX(index int) int {
	return iba.Call("getX", index).Int()
}
func (iba *InterleavedBufferAttribute) GetY(index int) int {
	return iba.Call("getY", index).Int()
}
func (iba *InterleavedBufferAttribute) GetZ(index int) int {
	return iba.Call("getZ", index).Int()
}
func (iba *InterleavedBufferAttribute) SetW(index int, z int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *InterleavedBufferAttribute) SetX(index int, x int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *InterleavedBufferAttribute) SetXY(index int, x int, y int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *InterleavedBufferAttribute) SetXYZ(index int, x int, y int, z int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *InterleavedBufferAttribute) SetXYZW(index int, x int, y int, z int, w int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *InterleavedBufferAttribute) SetY(index int, y int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *InterleavedBufferAttribute) SetZ(index int, z int) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setZ", index, z)}
}

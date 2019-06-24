// Code generated from "core/InterleavedBufferAttribute.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type InterleavedBufferAttribute struct {
	js.Value
}

func NewInterleavedBufferAttribute(interleavedBuffer *InterleavedBuffer, itemSize int, offset int, normalized bool) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: get("InterleavedBufferAttribute").New(interleavedBuffer, itemSize, offset, normalized)}
}
func (iba *InterleavedBufferAttribute) Array() js.Value {
	return iba.Get("array")
}
func (iba *InterleavedBufferAttribute) SetArray(v js.Value) {
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
func (iba *InterleavedBufferAttribute) GetW(index int) float64 {
	return iba.Call("getW", index).Float()
}
func (iba *InterleavedBufferAttribute) GetX(index int) float64 {
	return iba.Call("getX", index).Float()
}
func (iba *InterleavedBufferAttribute) GetY(index int) float64 {
	return iba.Call("getY", index).Float()
}
func (iba *InterleavedBufferAttribute) GetZ(index int) float64 {
	return iba.Call("getZ", index).Float()
}
func (iba *InterleavedBufferAttribute) SetW(index int, z float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *InterleavedBufferAttribute) SetX(index int, x float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *InterleavedBufferAttribute) SetXY(index int, x float64, y float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *InterleavedBufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *InterleavedBufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *InterleavedBufferAttribute) SetY(index int, y float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *InterleavedBufferAttribute) SetZ(index int, z float64) *InterleavedBufferAttribute {
	return &InterleavedBufferAttribute{Value: iba.Call("setZ", index, z)}
}

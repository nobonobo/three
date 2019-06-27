// Code generated from "core/BufferAttribute.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// BufferAttribute extend: []
type BufferAttribute struct {
	js.Value
}

func NewBufferAttribute(array js.Value, itemSize int, normalized bool) *BufferAttribute {
	return &BufferAttribute{Value: get("BufferAttribute").New(array, itemSize, normalized)}
}
func (ba *BufferAttribute) JSValue() js.Value {
	return ba.Value
}
func (ba *BufferAttribute) Array() js.Value {
	return ba.Get("array")
}
func (ba *BufferAttribute) SetArray(v js.Value) {
	ba.Set("array", v)
}
func (ba *BufferAttribute) Count() int {
	return ba.Get("count").Int()
}
func (ba *BufferAttribute) SetCount(v int) {
	ba.Set("count", v)
}
func (ba *BufferAttribute) Dynamic() bool {
	return ba.Get("dynamic").Bool()
}
func (ba *BufferAttribute) SetDynamic(v bool) {
	ba.Set("dynamic", v)
}
func (ba *BufferAttribute) ItemSize() int {
	return ba.Get("itemSize").Int()
}
func (ba *BufferAttribute) SetItemSize(v int) {
	ba.Set("itemSize", v)
}
func (ba *BufferAttribute) Length() int {
	return ba.Get("length").Int()
}
func (ba *BufferAttribute) SetLength(v int) {
	ba.Set("length", v)
}
func (ba *BufferAttribute) NeedsUpdate() bool {
	return ba.Get("needsUpdate").Bool()
}
func (ba *BufferAttribute) SetNeedsUpdate(v bool) {
	ba.Set("needsUpdate", v)
}
func (ba *BufferAttribute) Normalized() bool {
	return ba.Get("normalized").Bool()
}
func (ba *BufferAttribute) SetNormalized(v bool) {
	ba.Set("normalized", v)
}
func (ba *BufferAttribute) OnUpload() js.Value {
	return ba.Get("onUpload")
}
func (ba *BufferAttribute) SetOnUpload(v js.Value) {
	ba.Set("onUpload", v)
}
func (ba *BufferAttribute) UpdateRange() js.Value {
	return ba.Get("updateRange")
}
func (ba *BufferAttribute) SetUpdateRange(v js.Value) {
	ba.Set("updateRange", v)
}
func (ba *BufferAttribute) Uuid() string {
	return ba.Get("uuid").String()
}
func (ba *BufferAttribute) SetUuid(v string) {
	ba.Set("uuid", v)
}
func (ba *BufferAttribute) Version() float64 {
	return ba.Get("version").Float()
}
func (ba *BufferAttribute) SetVersion(v float64) {
	ba.Set("version", v)
}
func (ba *BufferAttribute) Clone() *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("clone")}
}
func (ba *BufferAttribute) Copy(source *BufferAttribute) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copy", source)}
}
func (ba *BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copyArray", array)}
}
func (ba *BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copyAt", index1, attribute, index2)}
}
func (ba *BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copyColorsArray", colors)}
}
func (ba *BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copyVector2sArray", vectors)}
}
func (ba *BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copyVector3sArray", vectors)}
}
func (ba *BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("copyVector4sArray", vectors)}
}
func (ba *BufferAttribute) GetW(index int) float64 {
	return ba.Call("getW", index).Float()
}
func (ba *BufferAttribute) GetX(index int) float64 {
	return ba.Call("getX", index).Float()
}
func (ba *BufferAttribute) GetY(index int) float64 {
	return ba.Call("getY", index).Float()
}
func (ba *BufferAttribute) GetZ(index int) float64 {
	return ba.Call("getZ", index).Float()
}
func (ba *BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("set", value, offset)}
}
func (ba *BufferAttribute) SetArray2(array js.Value) {
	ba.Call("setArray", array)
}
func (ba *BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setDynamic", dynamic)}
}
func (ba *BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setW", index, z)}
}
func (ba *BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setX", index, x)}
}
func (ba *BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setXY", index, x, y)}
}
func (ba *BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setXYZ", index, x, y, z)}
}
func (ba *BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setXYZW", index, x, y, z, w)}
}
func (ba *BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setY", index, y)}
}
func (ba *BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ba.Call("setZ", index, z)}
}

// Float32Attribute extend: [BufferAttribute]
type Float32Attribute struct {
	js.Value
}

func NewFloat32Attribute(array js.Value, itemSize int) *Float32Attribute {
	return &Float32Attribute{Value: get("Float32Attribute").New(array, itemSize)}
}
func (fa *Float32Attribute) JSValue() js.Value {
	return fa.Value
}
func (fa *Float32Attribute) Array() js.Value {
	return fa.Get("array")
}
func (fa *Float32Attribute) SetArray(v js.Value) {
	fa.Set("array", v)
}
func (fa *Float32Attribute) Count() int {
	return fa.Get("count").Int()
}
func (fa *Float32Attribute) SetCount(v int) {
	fa.Set("count", v)
}
func (fa *Float32Attribute) Dynamic() bool {
	return fa.Get("dynamic").Bool()
}
func (fa *Float32Attribute) SetDynamic(v bool) {
	fa.Set("dynamic", v)
}
func (fa *Float32Attribute) ItemSize() int {
	return fa.Get("itemSize").Int()
}
func (fa *Float32Attribute) SetItemSize(v int) {
	fa.Set("itemSize", v)
}
func (fa *Float32Attribute) Length() int {
	return fa.Get("length").Int()
}
func (fa *Float32Attribute) SetLength(v int) {
	fa.Set("length", v)
}
func (fa *Float32Attribute) NeedsUpdate() bool {
	return fa.Get("needsUpdate").Bool()
}
func (fa *Float32Attribute) SetNeedsUpdate(v bool) {
	fa.Set("needsUpdate", v)
}
func (fa *Float32Attribute) Normalized() bool {
	return fa.Get("normalized").Bool()
}
func (fa *Float32Attribute) SetNormalized(v bool) {
	fa.Set("normalized", v)
}
func (fa *Float32Attribute) OnUpload() js.Value {
	return fa.Get("onUpload")
}
func (fa *Float32Attribute) SetOnUpload(v js.Value) {
	fa.Set("onUpload", v)
}
func (fa *Float32Attribute) UpdateRange() js.Value {
	return fa.Get("updateRange")
}
func (fa *Float32Attribute) SetUpdateRange(v js.Value) {
	fa.Set("updateRange", v)
}
func (fa *Float32Attribute) Uuid() string {
	return fa.Get("uuid").String()
}
func (fa *Float32Attribute) SetUuid(v string) {
	fa.Set("uuid", v)
}
func (fa *Float32Attribute) Version() float64 {
	return fa.Get("version").Float()
}
func (fa *Float32Attribute) SetVersion(v float64) {
	fa.Set("version", v)
}
func (fa *Float32Attribute) Clone() *Float32Attribute {
	return &Float32Attribute{Value: fa.Call("clone")}
}
func (fa *Float32Attribute) Copy(source *BufferAttribute) *Float32Attribute {
	return &Float32Attribute{Value: fa.Call("copy", source)}
}
func (fa *Float32Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyArray", array)}
}
func (fa *Float32Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyAt", index1, attribute, index2)}
}
func (fa *Float32Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyColorsArray", colors)}
}
func (fa *Float32Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyVector2sArray", vectors)}
}
func (fa *Float32Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyVector3sArray", vectors)}
}
func (fa *Float32Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyVector4sArray", vectors)}
}
func (fa *Float32Attribute) GetW(index int) float64 {
	return fa.Call("getW", index).Float()
}
func (fa *Float32Attribute) GetX(index int) float64 {
	return fa.Call("getX", index).Float()
}
func (fa *Float32Attribute) GetY(index int) float64 {
	return fa.Call("getY", index).Float()
}
func (fa *Float32Attribute) GetZ(index int) float64 {
	return fa.Call("getZ", index).Float()
}
func (fa *Float32Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("set", value, offset)}
}
func (fa *Float32Attribute) SetArray2(array js.Value) {
	fa.Call("setArray", array)
}
func (fa *Float32Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setDynamic", dynamic)}
}
func (fa *Float32Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setW", index, z)}
}
func (fa *Float32Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setX", index, x)}
}
func (fa *Float32Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setXY", index, x, y)}
}
func (fa *Float32Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setXYZ", index, x, y, z)}
}
func (fa *Float32Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setXYZW", index, x, y, z, w)}
}
func (fa *Float32Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setY", index, y)}
}
func (fa *Float32Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setZ", index, z)}
}

// Float32BufferAttribute extend: [BufferAttribute]
type Float32BufferAttribute struct {
	js.Value
}

func NewFloat32BufferAttribute(array js.Value, itemSize int, normalized bool) *Float32BufferAttribute {
	return &Float32BufferAttribute{Value: get("Float32BufferAttribute").New(array, itemSize, normalized)}
}
func (fba *Float32BufferAttribute) JSValue() js.Value {
	return fba.Value
}
func (fba *Float32BufferAttribute) Array() js.Value {
	return fba.Get("array")
}
func (fba *Float32BufferAttribute) SetArray(v js.Value) {
	fba.Set("array", v)
}
func (fba *Float32BufferAttribute) Count() int {
	return fba.Get("count").Int()
}
func (fba *Float32BufferAttribute) SetCount(v int) {
	fba.Set("count", v)
}
func (fba *Float32BufferAttribute) Dynamic() bool {
	return fba.Get("dynamic").Bool()
}
func (fba *Float32BufferAttribute) SetDynamic(v bool) {
	fba.Set("dynamic", v)
}
func (fba *Float32BufferAttribute) ItemSize() int {
	return fba.Get("itemSize").Int()
}
func (fba *Float32BufferAttribute) SetItemSize(v int) {
	fba.Set("itemSize", v)
}
func (fba *Float32BufferAttribute) Length() int {
	return fba.Get("length").Int()
}
func (fba *Float32BufferAttribute) SetLength(v int) {
	fba.Set("length", v)
}
func (fba *Float32BufferAttribute) NeedsUpdate() bool {
	return fba.Get("needsUpdate").Bool()
}
func (fba *Float32BufferAttribute) SetNeedsUpdate(v bool) {
	fba.Set("needsUpdate", v)
}
func (fba *Float32BufferAttribute) Normalized() bool {
	return fba.Get("normalized").Bool()
}
func (fba *Float32BufferAttribute) SetNormalized(v bool) {
	fba.Set("normalized", v)
}
func (fba *Float32BufferAttribute) OnUpload() js.Value {
	return fba.Get("onUpload")
}
func (fba *Float32BufferAttribute) SetOnUpload(v js.Value) {
	fba.Set("onUpload", v)
}
func (fba *Float32BufferAttribute) UpdateRange() js.Value {
	return fba.Get("updateRange")
}
func (fba *Float32BufferAttribute) SetUpdateRange(v js.Value) {
	fba.Set("updateRange", v)
}
func (fba *Float32BufferAttribute) Uuid() string {
	return fba.Get("uuid").String()
}
func (fba *Float32BufferAttribute) SetUuid(v string) {
	fba.Set("uuid", v)
}
func (fba *Float32BufferAttribute) Version() float64 {
	return fba.Get("version").Float()
}
func (fba *Float32BufferAttribute) SetVersion(v float64) {
	fba.Set("version", v)
}
func (fba *Float32BufferAttribute) Clone() *Float32BufferAttribute {
	return &Float32BufferAttribute{Value: fba.Call("clone")}
}
func (fba *Float32BufferAttribute) Copy(source *BufferAttribute) *Float32BufferAttribute {
	return &Float32BufferAttribute{Value: fba.Call("copy", source)}
}
func (fba *Float32BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyArray", array)}
}
func (fba *Float32BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyAt", index1, attribute, index2)}
}
func (fba *Float32BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyColorsArray", colors)}
}
func (fba *Float32BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyVector2sArray", vectors)}
}
func (fba *Float32BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyVector3sArray", vectors)}
}
func (fba *Float32BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyVector4sArray", vectors)}
}
func (fba *Float32BufferAttribute) GetW(index int) float64 {
	return fba.Call("getW", index).Float()
}
func (fba *Float32BufferAttribute) GetX(index int) float64 {
	return fba.Call("getX", index).Float()
}
func (fba *Float32BufferAttribute) GetY(index int) float64 {
	return fba.Call("getY", index).Float()
}
func (fba *Float32BufferAttribute) GetZ(index int) float64 {
	return fba.Call("getZ", index).Float()
}
func (fba *Float32BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("set", value, offset)}
}
func (fba *Float32BufferAttribute) SetArray2(array js.Value) {
	fba.Call("setArray", array)
}
func (fba *Float32BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setDynamic", dynamic)}
}
func (fba *Float32BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setW", index, z)}
}
func (fba *Float32BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setX", index, x)}
}
func (fba *Float32BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setXY", index, x, y)}
}
func (fba *Float32BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setXYZ", index, x, y, z)}
}
func (fba *Float32BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setXYZW", index, x, y, z, w)}
}
func (fba *Float32BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setY", index, y)}
}
func (fba *Float32BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setZ", index, z)}
}

// Float64Attribute extend: [BufferAttribute]
type Float64Attribute struct {
	js.Value
}

func NewFloat64Attribute(array js.Value, itemSize int) *Float64Attribute {
	return &Float64Attribute{Value: get("Float64Attribute").New(array, itemSize)}
}
func (fa *Float64Attribute) JSValue() js.Value {
	return fa.Value
}
func (fa *Float64Attribute) Array() js.Value {
	return fa.Get("array")
}
func (fa *Float64Attribute) SetArray(v js.Value) {
	fa.Set("array", v)
}
func (fa *Float64Attribute) Count() int {
	return fa.Get("count").Int()
}
func (fa *Float64Attribute) SetCount(v int) {
	fa.Set("count", v)
}
func (fa *Float64Attribute) Dynamic() bool {
	return fa.Get("dynamic").Bool()
}
func (fa *Float64Attribute) SetDynamic(v bool) {
	fa.Set("dynamic", v)
}
func (fa *Float64Attribute) ItemSize() int {
	return fa.Get("itemSize").Int()
}
func (fa *Float64Attribute) SetItemSize(v int) {
	fa.Set("itemSize", v)
}
func (fa *Float64Attribute) Length() float64 {
	return fa.Get("length").Float()
}
func (fa *Float64Attribute) SetLength(v float64) {
	fa.Set("length", v)
}
func (fa *Float64Attribute) NeedsUpdate() bool {
	return fa.Get("needsUpdate").Bool()
}
func (fa *Float64Attribute) SetNeedsUpdate(v bool) {
	fa.Set("needsUpdate", v)
}
func (fa *Float64Attribute) Normalized() bool {
	return fa.Get("normalized").Bool()
}
func (fa *Float64Attribute) SetNormalized(v bool) {
	fa.Set("normalized", v)
}
func (fa *Float64Attribute) OnUpload() js.Value {
	return fa.Get("onUpload")
}
func (fa *Float64Attribute) SetOnUpload(v js.Value) {
	fa.Set("onUpload", v)
}
func (fa *Float64Attribute) UpdateRange() js.Value {
	return fa.Get("updateRange")
}
func (fa *Float64Attribute) SetUpdateRange(v js.Value) {
	fa.Set("updateRange", v)
}
func (fa *Float64Attribute) Uuid() string {
	return fa.Get("uuid").String()
}
func (fa *Float64Attribute) SetUuid(v string) {
	fa.Set("uuid", v)
}
func (fa *Float64Attribute) Version() float64 {
	return fa.Get("version").Float()
}
func (fa *Float64Attribute) SetVersion(v float64) {
	fa.Set("version", v)
}
func (fa *Float64Attribute) Clone() *Float64Attribute {
	return &Float64Attribute{Value: fa.Call("clone")}
}
func (fa *Float64Attribute) Copy(source *BufferAttribute) *Float64Attribute {
	return &Float64Attribute{Value: fa.Call("copy", source)}
}
func (fa *Float64Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyArray", array)}
}
func (fa *Float64Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyAt", index1, attribute, index2)}
}
func (fa *Float64Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyColorsArray", colors)}
}
func (fa *Float64Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyVector2sArray", vectors)}
}
func (fa *Float64Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyVector3sArray", vectors)}
}
func (fa *Float64Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("copyVector4sArray", vectors)}
}
func (fa *Float64Attribute) GetW(index int) float64 {
	return fa.Call("getW", index).Float()
}
func (fa *Float64Attribute) GetX(index int) float64 {
	return fa.Call("getX", index).Float()
}
func (fa *Float64Attribute) GetY(index int) float64 {
	return fa.Call("getY", index).Float()
}
func (fa *Float64Attribute) GetZ(index int) float64 {
	return fa.Call("getZ", index).Float()
}
func (fa *Float64Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("set", value, offset)}
}
func (fa *Float64Attribute) SetArray2(array js.Value) {
	fa.Call("setArray", array)
}
func (fa *Float64Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setDynamic", dynamic)}
}
func (fa *Float64Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setW", index, z)}
}
func (fa *Float64Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setX", index, x)}
}
func (fa *Float64Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setXY", index, x, y)}
}
func (fa *Float64Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setXYZ", index, x, y, z)}
}
func (fa *Float64Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setXYZW", index, x, y, z, w)}
}
func (fa *Float64Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setY", index, y)}
}
func (fa *Float64Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fa.Call("setZ", index, z)}
}

// Float64BufferAttribute extend: [BufferAttribute]
type Float64BufferAttribute struct {
	js.Value
}

func NewFloat64BufferAttribute(array js.Value, itemSize int, normalized bool) *Float64BufferAttribute {
	return &Float64BufferAttribute{Value: get("Float64BufferAttribute").New(array, itemSize, normalized)}
}
func (fba *Float64BufferAttribute) JSValue() js.Value {
	return fba.Value
}
func (fba *Float64BufferAttribute) Array() js.Value {
	return fba.Get("array")
}
func (fba *Float64BufferAttribute) SetArray(v js.Value) {
	fba.Set("array", v)
}
func (fba *Float64BufferAttribute) Count() int {
	return fba.Get("count").Int()
}
func (fba *Float64BufferAttribute) SetCount(v int) {
	fba.Set("count", v)
}
func (fba *Float64BufferAttribute) Dynamic() bool {
	return fba.Get("dynamic").Bool()
}
func (fba *Float64BufferAttribute) SetDynamic(v bool) {
	fba.Set("dynamic", v)
}
func (fba *Float64BufferAttribute) ItemSize() int {
	return fba.Get("itemSize").Int()
}
func (fba *Float64BufferAttribute) SetItemSize(v int) {
	fba.Set("itemSize", v)
}
func (fba *Float64BufferAttribute) Length() int {
	return fba.Get("length").Int()
}
func (fba *Float64BufferAttribute) SetLength(v int) {
	fba.Set("length", v)
}
func (fba *Float64BufferAttribute) NeedsUpdate() bool {
	return fba.Get("needsUpdate").Bool()
}
func (fba *Float64BufferAttribute) SetNeedsUpdate(v bool) {
	fba.Set("needsUpdate", v)
}
func (fba *Float64BufferAttribute) Normalized() bool {
	return fba.Get("normalized").Bool()
}
func (fba *Float64BufferAttribute) SetNormalized(v bool) {
	fba.Set("normalized", v)
}
func (fba *Float64BufferAttribute) OnUpload() js.Value {
	return fba.Get("onUpload")
}
func (fba *Float64BufferAttribute) SetOnUpload(v js.Value) {
	fba.Set("onUpload", v)
}
func (fba *Float64BufferAttribute) UpdateRange() js.Value {
	return fba.Get("updateRange")
}
func (fba *Float64BufferAttribute) SetUpdateRange(v js.Value) {
	fba.Set("updateRange", v)
}
func (fba *Float64BufferAttribute) Uuid() string {
	return fba.Get("uuid").String()
}
func (fba *Float64BufferAttribute) SetUuid(v string) {
	fba.Set("uuid", v)
}
func (fba *Float64BufferAttribute) Version() float64 {
	return fba.Get("version").Float()
}
func (fba *Float64BufferAttribute) SetVersion(v float64) {
	fba.Set("version", v)
}
func (fba *Float64BufferAttribute) Clone() *Float64BufferAttribute {
	return &Float64BufferAttribute{Value: fba.Call("clone")}
}
func (fba *Float64BufferAttribute) Copy(source *BufferAttribute) *Float64BufferAttribute {
	return &Float64BufferAttribute{Value: fba.Call("copy", source)}
}
func (fba *Float64BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyArray", array)}
}
func (fba *Float64BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyAt", index1, attribute, index2)}
}
func (fba *Float64BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyColorsArray", colors)}
}
func (fba *Float64BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyVector2sArray", vectors)}
}
func (fba *Float64BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyVector3sArray", vectors)}
}
func (fba *Float64BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("copyVector4sArray", vectors)}
}
func (fba *Float64BufferAttribute) GetW(index int) float64 {
	return fba.Call("getW", index).Float()
}
func (fba *Float64BufferAttribute) GetX(index int) float64 {
	return fba.Call("getX", index).Float()
}
func (fba *Float64BufferAttribute) GetY(index int) float64 {
	return fba.Call("getY", index).Float()
}
func (fba *Float64BufferAttribute) GetZ(index int) float64 {
	return fba.Call("getZ", index).Float()
}
func (fba *Float64BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("set", value, offset)}
}
func (fba *Float64BufferAttribute) SetArray2(array js.Value) {
	fba.Call("setArray", array)
}
func (fba *Float64BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setDynamic", dynamic)}
}
func (fba *Float64BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setW", index, z)}
}
func (fba *Float64BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setX", index, x)}
}
func (fba *Float64BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setXY", index, x, y)}
}
func (fba *Float64BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setXYZ", index, x, y, z)}
}
func (fba *Float64BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setXYZW", index, x, y, z, w)}
}
func (fba *Float64BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setY", index, y)}
}
func (fba *Float64BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: fba.Call("setZ", index, z)}
}

// Int16Attribute extend: [BufferAttribute]
type Int16Attribute struct {
	js.Value
}

func NewInt16Attribute(array js.Value, itemSize int) *Int16Attribute {
	return &Int16Attribute{Value: get("Int16Attribute").New(array, itemSize)}
}
func (ia *Int16Attribute) JSValue() js.Value {
	return ia.Value
}
func (ia *Int16Attribute) Array() js.Value {
	return ia.Get("array")
}
func (ia *Int16Attribute) SetArray(v js.Value) {
	ia.Set("array", v)
}
func (ia *Int16Attribute) Count() int {
	return ia.Get("count").Int()
}
func (ia *Int16Attribute) SetCount(v int) {
	ia.Set("count", v)
}
func (ia *Int16Attribute) Dynamic() bool {
	return ia.Get("dynamic").Bool()
}
func (ia *Int16Attribute) SetDynamic(v bool) {
	ia.Set("dynamic", v)
}
func (ia *Int16Attribute) ItemSize() int {
	return ia.Get("itemSize").Int()
}
func (ia *Int16Attribute) SetItemSize(v int) {
	ia.Set("itemSize", v)
}
func (ia *Int16Attribute) Length() int {
	return ia.Get("length").Int()
}
func (ia *Int16Attribute) SetLength(v int) {
	ia.Set("length", v)
}
func (ia *Int16Attribute) NeedsUpdate() bool {
	return ia.Get("needsUpdate").Bool()
}
func (ia *Int16Attribute) SetNeedsUpdate(v bool) {
	ia.Set("needsUpdate", v)
}
func (ia *Int16Attribute) Normalized() bool {
	return ia.Get("normalized").Bool()
}
func (ia *Int16Attribute) SetNormalized(v bool) {
	ia.Set("normalized", v)
}
func (ia *Int16Attribute) OnUpload() js.Value {
	return ia.Get("onUpload")
}
func (ia *Int16Attribute) SetOnUpload(v js.Value) {
	ia.Set("onUpload", v)
}
func (ia *Int16Attribute) UpdateRange() js.Value {
	return ia.Get("updateRange")
}
func (ia *Int16Attribute) SetUpdateRange(v js.Value) {
	ia.Set("updateRange", v)
}
func (ia *Int16Attribute) Uuid() string {
	return ia.Get("uuid").String()
}
func (ia *Int16Attribute) SetUuid(v string) {
	ia.Set("uuid", v)
}
func (ia *Int16Attribute) Version() float64 {
	return ia.Get("version").Float()
}
func (ia *Int16Attribute) SetVersion(v float64) {
	ia.Set("version", v)
}
func (ia *Int16Attribute) Clone() *Int16Attribute {
	return &Int16Attribute{Value: ia.Call("clone")}
}
func (ia *Int16Attribute) Copy(source *BufferAttribute) *Int16Attribute {
	return &Int16Attribute{Value: ia.Call("copy", source)}
}
func (ia *Int16Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyArray", array)}
}
func (ia *Int16Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyAt", index1, attribute, index2)}
}
func (ia *Int16Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyColorsArray", colors)}
}
func (ia *Int16Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector2sArray", vectors)}
}
func (ia *Int16Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector3sArray", vectors)}
}
func (ia *Int16Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector4sArray", vectors)}
}
func (ia *Int16Attribute) GetW(index int) float64 {
	return ia.Call("getW", index).Float()
}
func (ia *Int16Attribute) GetX(index int) float64 {
	return ia.Call("getX", index).Float()
}
func (ia *Int16Attribute) GetY(index int) float64 {
	return ia.Call("getY", index).Float()
}
func (ia *Int16Attribute) GetZ(index int) float64 {
	return ia.Call("getZ", index).Float()
}
func (ia *Int16Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("set", value, offset)}
}
func (ia *Int16Attribute) SetArray2(array js.Value) {
	ia.Call("setArray", array)
}
func (ia *Int16Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setDynamic", dynamic)}
}
func (ia *Int16Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setW", index, z)}
}
func (ia *Int16Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setX", index, x)}
}
func (ia *Int16Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXY", index, x, y)}
}
func (ia *Int16Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXYZ", index, x, y, z)}
}
func (ia *Int16Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXYZW", index, x, y, z, w)}
}
func (ia *Int16Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setY", index, y)}
}
func (ia *Int16Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setZ", index, z)}
}

// Int16BufferAttribute extend: [BufferAttribute]
type Int16BufferAttribute struct {
	js.Value
}

func NewInt16BufferAttribute(array js.Value, itemSize int, normalized bool) *Int16BufferAttribute {
	return &Int16BufferAttribute{Value: get("Int16BufferAttribute").New(array, itemSize, normalized)}
}
func (iba *Int16BufferAttribute) JSValue() js.Value {
	return iba.Value
}
func (iba *Int16BufferAttribute) Array() js.Value {
	return iba.Get("array")
}
func (iba *Int16BufferAttribute) SetArray(v js.Value) {
	iba.Set("array", v)
}
func (iba *Int16BufferAttribute) Count() int {
	return iba.Get("count").Int()
}
func (iba *Int16BufferAttribute) SetCount(v int) {
	iba.Set("count", v)
}
func (iba *Int16BufferAttribute) Dynamic() bool {
	return iba.Get("dynamic").Bool()
}
func (iba *Int16BufferAttribute) SetDynamic(v bool) {
	iba.Set("dynamic", v)
}
func (iba *Int16BufferAttribute) ItemSize() int {
	return iba.Get("itemSize").Int()
}
func (iba *Int16BufferAttribute) SetItemSize(v int) {
	iba.Set("itemSize", v)
}
func (iba *Int16BufferAttribute) Length() int {
	return iba.Get("length").Int()
}
func (iba *Int16BufferAttribute) SetLength(v int) {
	iba.Set("length", v)
}
func (iba *Int16BufferAttribute) NeedsUpdate() bool {
	return iba.Get("needsUpdate").Bool()
}
func (iba *Int16BufferAttribute) SetNeedsUpdate(v bool) {
	iba.Set("needsUpdate", v)
}
func (iba *Int16BufferAttribute) Normalized() bool {
	return iba.Get("normalized").Bool()
}
func (iba *Int16BufferAttribute) SetNormalized(v bool) {
	iba.Set("normalized", v)
}
func (iba *Int16BufferAttribute) OnUpload() js.Value {
	return iba.Get("onUpload")
}
func (iba *Int16BufferAttribute) SetOnUpload(v js.Value) {
	iba.Set("onUpload", v)
}
func (iba *Int16BufferAttribute) UpdateRange() js.Value {
	return iba.Get("updateRange")
}
func (iba *Int16BufferAttribute) SetUpdateRange(v js.Value) {
	iba.Set("updateRange", v)
}
func (iba *Int16BufferAttribute) Uuid() string {
	return iba.Get("uuid").String()
}
func (iba *Int16BufferAttribute) SetUuid(v string) {
	iba.Set("uuid", v)
}
func (iba *Int16BufferAttribute) Version() float64 {
	return iba.Get("version").Float()
}
func (iba *Int16BufferAttribute) SetVersion(v float64) {
	iba.Set("version", v)
}
func (iba *Int16BufferAttribute) Clone() *Int16BufferAttribute {
	return &Int16BufferAttribute{Value: iba.Call("clone")}
}
func (iba *Int16BufferAttribute) Copy(source *BufferAttribute) *Int16BufferAttribute {
	return &Int16BufferAttribute{Value: iba.Call("copy", source)}
}
func (iba *Int16BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyArray", array)}
}
func (iba *Int16BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyAt", index1, attribute, index2)}
}
func (iba *Int16BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyColorsArray", colors)}
}
func (iba *Int16BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector2sArray", vectors)}
}
func (iba *Int16BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector3sArray", vectors)}
}
func (iba *Int16BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector4sArray", vectors)}
}
func (iba *Int16BufferAttribute) GetW(index int) float64 {
	return iba.Call("getW", index).Float()
}
func (iba *Int16BufferAttribute) GetX(index int) float64 {
	return iba.Call("getX", index).Float()
}
func (iba *Int16BufferAttribute) GetY(index int) float64 {
	return iba.Call("getY", index).Float()
}
func (iba *Int16BufferAttribute) GetZ(index int) float64 {
	return iba.Call("getZ", index).Float()
}
func (iba *Int16BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("set", value, offset)}
}
func (iba *Int16BufferAttribute) SetArray2(array js.Value) {
	iba.Call("setArray", array)
}
func (iba *Int16BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setDynamic", dynamic)}
}
func (iba *Int16BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *Int16BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *Int16BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *Int16BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *Int16BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *Int16BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *Int16BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setZ", index, z)}
}

// Int32Attribute extend: [BufferAttribute]
type Int32Attribute struct {
	js.Value
}

func NewInt32Attribute(array js.Value, itemSize int) *Int32Attribute {
	return &Int32Attribute{Value: get("Int32Attribute").New(array, itemSize)}
}
func (ia *Int32Attribute) JSValue() js.Value {
	return ia.Value
}
func (ia *Int32Attribute) Array() js.Value {
	return ia.Get("array")
}
func (ia *Int32Attribute) SetArray(v js.Value) {
	ia.Set("array", v)
}
func (ia *Int32Attribute) Count() int {
	return ia.Get("count").Int()
}
func (ia *Int32Attribute) SetCount(v int) {
	ia.Set("count", v)
}
func (ia *Int32Attribute) Dynamic() bool {
	return ia.Get("dynamic").Bool()
}
func (ia *Int32Attribute) SetDynamic(v bool) {
	ia.Set("dynamic", v)
}
func (ia *Int32Attribute) ItemSize() int {
	return ia.Get("itemSize").Int()
}
func (ia *Int32Attribute) SetItemSize(v int) {
	ia.Set("itemSize", v)
}
func (ia *Int32Attribute) Length() int {
	return ia.Get("length").Int()
}
func (ia *Int32Attribute) SetLength(v int) {
	ia.Set("length", v)
}
func (ia *Int32Attribute) NeedsUpdate() bool {
	return ia.Get("needsUpdate").Bool()
}
func (ia *Int32Attribute) SetNeedsUpdate(v bool) {
	ia.Set("needsUpdate", v)
}
func (ia *Int32Attribute) Normalized() bool {
	return ia.Get("normalized").Bool()
}
func (ia *Int32Attribute) SetNormalized(v bool) {
	ia.Set("normalized", v)
}
func (ia *Int32Attribute) OnUpload() js.Value {
	return ia.Get("onUpload")
}
func (ia *Int32Attribute) SetOnUpload(v js.Value) {
	ia.Set("onUpload", v)
}
func (ia *Int32Attribute) UpdateRange() js.Value {
	return ia.Get("updateRange")
}
func (ia *Int32Attribute) SetUpdateRange(v js.Value) {
	ia.Set("updateRange", v)
}
func (ia *Int32Attribute) Uuid() string {
	return ia.Get("uuid").String()
}
func (ia *Int32Attribute) SetUuid(v string) {
	ia.Set("uuid", v)
}
func (ia *Int32Attribute) Version() float64 {
	return ia.Get("version").Float()
}
func (ia *Int32Attribute) SetVersion(v float64) {
	ia.Set("version", v)
}
func (ia *Int32Attribute) Clone() *Int32Attribute {
	return &Int32Attribute{Value: ia.Call("clone")}
}
func (ia *Int32Attribute) Copy(source *BufferAttribute) *Int32Attribute {
	return &Int32Attribute{Value: ia.Call("copy", source)}
}
func (ia *Int32Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyArray", array)}
}
func (ia *Int32Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyAt", index1, attribute, index2)}
}
func (ia *Int32Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyColorsArray", colors)}
}
func (ia *Int32Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector2sArray", vectors)}
}
func (ia *Int32Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector3sArray", vectors)}
}
func (ia *Int32Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector4sArray", vectors)}
}
func (ia *Int32Attribute) GetW(index int) float64 {
	return ia.Call("getW", index).Float()
}
func (ia *Int32Attribute) GetX(index int) float64 {
	return ia.Call("getX", index).Float()
}
func (ia *Int32Attribute) GetY(index int) float64 {
	return ia.Call("getY", index).Float()
}
func (ia *Int32Attribute) GetZ(index int) float64 {
	return ia.Call("getZ", index).Float()
}
func (ia *Int32Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("set", value, offset)}
}
func (ia *Int32Attribute) SetArray2(array js.Value) {
	ia.Call("setArray", array)
}
func (ia *Int32Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setDynamic", dynamic)}
}
func (ia *Int32Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setW", index, z)}
}
func (ia *Int32Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setX", index, x)}
}
func (ia *Int32Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXY", index, x, y)}
}
func (ia *Int32Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXYZ", index, x, y, z)}
}
func (ia *Int32Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXYZW", index, x, y, z, w)}
}
func (ia *Int32Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setY", index, y)}
}
func (ia *Int32Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setZ", index, z)}
}

// Int32BufferAttribute extend: [BufferAttribute]
type Int32BufferAttribute struct {
	js.Value
}

func NewInt32BufferAttribute(array js.Value, itemSize int, normalized bool) *Int32BufferAttribute {
	return &Int32BufferAttribute{Value: get("Int32BufferAttribute").New(array, itemSize, normalized)}
}
func (iba *Int32BufferAttribute) JSValue() js.Value {
	return iba.Value
}
func (iba *Int32BufferAttribute) Array() js.Value {
	return iba.Get("array")
}
func (iba *Int32BufferAttribute) SetArray(v js.Value) {
	iba.Set("array", v)
}
func (iba *Int32BufferAttribute) Count() int {
	return iba.Get("count").Int()
}
func (iba *Int32BufferAttribute) SetCount(v int) {
	iba.Set("count", v)
}
func (iba *Int32BufferAttribute) Dynamic() bool {
	return iba.Get("dynamic").Bool()
}
func (iba *Int32BufferAttribute) SetDynamic(v bool) {
	iba.Set("dynamic", v)
}
func (iba *Int32BufferAttribute) ItemSize() int {
	return iba.Get("itemSize").Int()
}
func (iba *Int32BufferAttribute) SetItemSize(v int) {
	iba.Set("itemSize", v)
}
func (iba *Int32BufferAttribute) Length() float64 {
	return iba.Get("length").Float()
}
func (iba *Int32BufferAttribute) SetLength(v float64) {
	iba.Set("length", v)
}
func (iba *Int32BufferAttribute) NeedsUpdate() bool {
	return iba.Get("needsUpdate").Bool()
}
func (iba *Int32BufferAttribute) SetNeedsUpdate(v bool) {
	iba.Set("needsUpdate", v)
}
func (iba *Int32BufferAttribute) Normalized() bool {
	return iba.Get("normalized").Bool()
}
func (iba *Int32BufferAttribute) SetNormalized(v bool) {
	iba.Set("normalized", v)
}
func (iba *Int32BufferAttribute) OnUpload() js.Value {
	return iba.Get("onUpload")
}
func (iba *Int32BufferAttribute) SetOnUpload(v js.Value) {
	iba.Set("onUpload", v)
}
func (iba *Int32BufferAttribute) UpdateRange() js.Value {
	return iba.Get("updateRange")
}
func (iba *Int32BufferAttribute) SetUpdateRange(v js.Value) {
	iba.Set("updateRange", v)
}
func (iba *Int32BufferAttribute) Uuid() string {
	return iba.Get("uuid").String()
}
func (iba *Int32BufferAttribute) SetUuid(v string) {
	iba.Set("uuid", v)
}
func (iba *Int32BufferAttribute) Version() float64 {
	return iba.Get("version").Float()
}
func (iba *Int32BufferAttribute) SetVersion(v float64) {
	iba.Set("version", v)
}
func (iba *Int32BufferAttribute) Clone() *Int32BufferAttribute {
	return &Int32BufferAttribute{Value: iba.Call("clone")}
}
func (iba *Int32BufferAttribute) Copy(source *BufferAttribute) *Int32BufferAttribute {
	return &Int32BufferAttribute{Value: iba.Call("copy", source)}
}
func (iba *Int32BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyArray", array)}
}
func (iba *Int32BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyAt", index1, attribute, index2)}
}
func (iba *Int32BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyColorsArray", colors)}
}
func (iba *Int32BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector2sArray", vectors)}
}
func (iba *Int32BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector3sArray", vectors)}
}
func (iba *Int32BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector4sArray", vectors)}
}
func (iba *Int32BufferAttribute) GetW(index int) float64 {
	return iba.Call("getW", index).Float()
}
func (iba *Int32BufferAttribute) GetX(index int) float64 {
	return iba.Call("getX", index).Float()
}
func (iba *Int32BufferAttribute) GetY(index int) float64 {
	return iba.Call("getY", index).Float()
}
func (iba *Int32BufferAttribute) GetZ(index int) float64 {
	return iba.Call("getZ", index).Float()
}
func (iba *Int32BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("set", value, offset)}
}
func (iba *Int32BufferAttribute) SetArray2(array js.Value) {
	iba.Call("setArray", array)
}
func (iba *Int32BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setDynamic", dynamic)}
}
func (iba *Int32BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *Int32BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *Int32BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *Int32BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *Int32BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *Int32BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *Int32BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setZ", index, z)}
}

// Int8Attribute extend: [BufferAttribute]
type Int8Attribute struct {
	js.Value
}

func NewInt8Attribute(array js.Value, itemSize int) *Int8Attribute {
	return &Int8Attribute{Value: get("Int8Attribute").New(array, itemSize)}
}
func (ia *Int8Attribute) JSValue() js.Value {
	return ia.Value
}
func (ia *Int8Attribute) Array() js.Value {
	return ia.Get("array")
}
func (ia *Int8Attribute) SetArray(v js.Value) {
	ia.Set("array", v)
}
func (ia *Int8Attribute) Count() int {
	return ia.Get("count").Int()
}
func (ia *Int8Attribute) SetCount(v int) {
	ia.Set("count", v)
}
func (ia *Int8Attribute) Dynamic() bool {
	return ia.Get("dynamic").Bool()
}
func (ia *Int8Attribute) SetDynamic(v bool) {
	ia.Set("dynamic", v)
}
func (ia *Int8Attribute) ItemSize() int {
	return ia.Get("itemSize").Int()
}
func (ia *Int8Attribute) SetItemSize(v int) {
	ia.Set("itemSize", v)
}
func (ia *Int8Attribute) Length() int {
	return ia.Get("length").Int()
}
func (ia *Int8Attribute) SetLength(v int) {
	ia.Set("length", v)
}
func (ia *Int8Attribute) NeedsUpdate() bool {
	return ia.Get("needsUpdate").Bool()
}
func (ia *Int8Attribute) SetNeedsUpdate(v bool) {
	ia.Set("needsUpdate", v)
}
func (ia *Int8Attribute) Normalized() bool {
	return ia.Get("normalized").Bool()
}
func (ia *Int8Attribute) SetNormalized(v bool) {
	ia.Set("normalized", v)
}
func (ia *Int8Attribute) OnUpload() js.Value {
	return ia.Get("onUpload")
}
func (ia *Int8Attribute) SetOnUpload(v js.Value) {
	ia.Set("onUpload", v)
}
func (ia *Int8Attribute) UpdateRange() js.Value {
	return ia.Get("updateRange")
}
func (ia *Int8Attribute) SetUpdateRange(v js.Value) {
	ia.Set("updateRange", v)
}
func (ia *Int8Attribute) Uuid() string {
	return ia.Get("uuid").String()
}
func (ia *Int8Attribute) SetUuid(v string) {
	ia.Set("uuid", v)
}
func (ia *Int8Attribute) Version() float64 {
	return ia.Get("version").Float()
}
func (ia *Int8Attribute) SetVersion(v float64) {
	ia.Set("version", v)
}
func (ia *Int8Attribute) Clone() *Int8Attribute {
	return &Int8Attribute{Value: ia.Call("clone")}
}
func (ia *Int8Attribute) Copy(source *BufferAttribute) *Int8Attribute {
	return &Int8Attribute{Value: ia.Call("copy", source)}
}
func (ia *Int8Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyArray", array)}
}
func (ia *Int8Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyAt", index1, attribute, index2)}
}
func (ia *Int8Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyColorsArray", colors)}
}
func (ia *Int8Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector2sArray", vectors)}
}
func (ia *Int8Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector3sArray", vectors)}
}
func (ia *Int8Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("copyVector4sArray", vectors)}
}
func (ia *Int8Attribute) GetW(index int) float64 {
	return ia.Call("getW", index).Float()
}
func (ia *Int8Attribute) GetX(index int) float64 {
	return ia.Call("getX", index).Float()
}
func (ia *Int8Attribute) GetY(index int) float64 {
	return ia.Call("getY", index).Float()
}
func (ia *Int8Attribute) GetZ(index int) float64 {
	return ia.Call("getZ", index).Float()
}
func (ia *Int8Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("set", value, offset)}
}
func (ia *Int8Attribute) SetArray2(array js.Value) {
	ia.Call("setArray", array)
}
func (ia *Int8Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setDynamic", dynamic)}
}
func (ia *Int8Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setW", index, z)}
}
func (ia *Int8Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setX", index, x)}
}
func (ia *Int8Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXY", index, x, y)}
}
func (ia *Int8Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXYZ", index, x, y, z)}
}
func (ia *Int8Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setXYZW", index, x, y, z, w)}
}
func (ia *Int8Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setY", index, y)}
}
func (ia *Int8Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ia.Call("setZ", index, z)}
}

// Int8BufferAttribute extend: [BufferAttribute]
type Int8BufferAttribute struct {
	js.Value
}

func NewInt8BufferAttribute(array js.Value, itemSize int, normalized bool) *Int8BufferAttribute {
	return &Int8BufferAttribute{Value: get("Int8BufferAttribute").New(array, itemSize, normalized)}
}
func (iba *Int8BufferAttribute) JSValue() js.Value {
	return iba.Value
}
func (iba *Int8BufferAttribute) Array() js.Value {
	return iba.Get("array")
}
func (iba *Int8BufferAttribute) SetArray(v js.Value) {
	iba.Set("array", v)
}
func (iba *Int8BufferAttribute) Count() int {
	return iba.Get("count").Int()
}
func (iba *Int8BufferAttribute) SetCount(v int) {
	iba.Set("count", v)
}
func (iba *Int8BufferAttribute) Dynamic() bool {
	return iba.Get("dynamic").Bool()
}
func (iba *Int8BufferAttribute) SetDynamic(v bool) {
	iba.Set("dynamic", v)
}
func (iba *Int8BufferAttribute) ItemSize() int {
	return iba.Get("itemSize").Int()
}
func (iba *Int8BufferAttribute) SetItemSize(v int) {
	iba.Set("itemSize", v)
}
func (iba *Int8BufferAttribute) Length() int {
	return iba.Get("length").Int()
}
func (iba *Int8BufferAttribute) SetLength(v int) {
	iba.Set("length", v)
}
func (iba *Int8BufferAttribute) NeedsUpdate() bool {
	return iba.Get("needsUpdate").Bool()
}
func (iba *Int8BufferAttribute) SetNeedsUpdate(v bool) {
	iba.Set("needsUpdate", v)
}
func (iba *Int8BufferAttribute) Normalized() bool {
	return iba.Get("normalized").Bool()
}
func (iba *Int8BufferAttribute) SetNormalized(v bool) {
	iba.Set("normalized", v)
}
func (iba *Int8BufferAttribute) OnUpload() js.Value {
	return iba.Get("onUpload")
}
func (iba *Int8BufferAttribute) SetOnUpload(v js.Value) {
	iba.Set("onUpload", v)
}
func (iba *Int8BufferAttribute) UpdateRange() js.Value {
	return iba.Get("updateRange")
}
func (iba *Int8BufferAttribute) SetUpdateRange(v js.Value) {
	iba.Set("updateRange", v)
}
func (iba *Int8BufferAttribute) Uuid() string {
	return iba.Get("uuid").String()
}
func (iba *Int8BufferAttribute) SetUuid(v string) {
	iba.Set("uuid", v)
}
func (iba *Int8BufferAttribute) Version() float64 {
	return iba.Get("version").Float()
}
func (iba *Int8BufferAttribute) SetVersion(v float64) {
	iba.Set("version", v)
}
func (iba *Int8BufferAttribute) Clone() *Int8BufferAttribute {
	return &Int8BufferAttribute{Value: iba.Call("clone")}
}
func (iba *Int8BufferAttribute) Copy(source *BufferAttribute) *Int8BufferAttribute {
	return &Int8BufferAttribute{Value: iba.Call("copy", source)}
}
func (iba *Int8BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyArray", array)}
}
func (iba *Int8BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyAt", index1, attribute, index2)}
}
func (iba *Int8BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyColorsArray", colors)}
}
func (iba *Int8BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector2sArray", vectors)}
}
func (iba *Int8BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector3sArray", vectors)}
}
func (iba *Int8BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("copyVector4sArray", vectors)}
}
func (iba *Int8BufferAttribute) GetW(index int) float64 {
	return iba.Call("getW", index).Float()
}
func (iba *Int8BufferAttribute) GetX(index int) float64 {
	return iba.Call("getX", index).Float()
}
func (iba *Int8BufferAttribute) GetY(index int) float64 {
	return iba.Call("getY", index).Float()
}
func (iba *Int8BufferAttribute) GetZ(index int) float64 {
	return iba.Call("getZ", index).Float()
}
func (iba *Int8BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("set", value, offset)}
}
func (iba *Int8BufferAttribute) SetArray2(array js.Value) {
	iba.Call("setArray", array)
}
func (iba *Int8BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setDynamic", dynamic)}
}
func (iba *Int8BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setW", index, z)}
}
func (iba *Int8BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setX", index, x)}
}
func (iba *Int8BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXY", index, x, y)}
}
func (iba *Int8BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZ", index, x, y, z)}
}
func (iba *Int8BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setXYZW", index, x, y, z, w)}
}
func (iba *Int8BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setY", index, y)}
}
func (iba *Int8BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: iba.Call("setZ", index, z)}
}

// Uint16Attribute extend: [BufferAttribute]
type Uint16Attribute struct {
	js.Value
}

func NewUint16Attribute(array js.Value, itemSize int) *Uint16Attribute {
	return &Uint16Attribute{Value: get("Uint16Attribute").New(array, itemSize)}
}
func (ua *Uint16Attribute) JSValue() js.Value {
	return ua.Value
}
func (ua *Uint16Attribute) Array() js.Value {
	return ua.Get("array")
}
func (ua *Uint16Attribute) SetArray(v js.Value) {
	ua.Set("array", v)
}
func (ua *Uint16Attribute) Count() int {
	return ua.Get("count").Int()
}
func (ua *Uint16Attribute) SetCount(v int) {
	ua.Set("count", v)
}
func (ua *Uint16Attribute) Dynamic() bool {
	return ua.Get("dynamic").Bool()
}
func (ua *Uint16Attribute) SetDynamic(v bool) {
	ua.Set("dynamic", v)
}
func (ua *Uint16Attribute) ItemSize() int {
	return ua.Get("itemSize").Int()
}
func (ua *Uint16Attribute) SetItemSize(v int) {
	ua.Set("itemSize", v)
}
func (ua *Uint16Attribute) Length() int {
	return ua.Get("length").Int()
}
func (ua *Uint16Attribute) SetLength(v int) {
	ua.Set("length", v)
}
func (ua *Uint16Attribute) NeedsUpdate() bool {
	return ua.Get("needsUpdate").Bool()
}
func (ua *Uint16Attribute) SetNeedsUpdate(v bool) {
	ua.Set("needsUpdate", v)
}
func (ua *Uint16Attribute) Normalized() bool {
	return ua.Get("normalized").Bool()
}
func (ua *Uint16Attribute) SetNormalized(v bool) {
	ua.Set("normalized", v)
}
func (ua *Uint16Attribute) OnUpload() js.Value {
	return ua.Get("onUpload")
}
func (ua *Uint16Attribute) SetOnUpload(v js.Value) {
	ua.Set("onUpload", v)
}
func (ua *Uint16Attribute) UpdateRange() js.Value {
	return ua.Get("updateRange")
}
func (ua *Uint16Attribute) SetUpdateRange(v js.Value) {
	ua.Set("updateRange", v)
}
func (ua *Uint16Attribute) Uuid() string {
	return ua.Get("uuid").String()
}
func (ua *Uint16Attribute) SetUuid(v string) {
	ua.Set("uuid", v)
}
func (ua *Uint16Attribute) Version() float64 {
	return ua.Get("version").Float()
}
func (ua *Uint16Attribute) SetVersion(v float64) {
	ua.Set("version", v)
}
func (ua *Uint16Attribute) Clone() *Uint16Attribute {
	return &Uint16Attribute{Value: ua.Call("clone")}
}
func (ua *Uint16Attribute) Copy(source *BufferAttribute) *Uint16Attribute {
	return &Uint16Attribute{Value: ua.Call("copy", source)}
}
func (ua *Uint16Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyArray", array)}
}
func (ua *Uint16Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyAt", index1, attribute, index2)}
}
func (ua *Uint16Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyColorsArray", colors)}
}
func (ua *Uint16Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector2sArray", vectors)}
}
func (ua *Uint16Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector3sArray", vectors)}
}
func (ua *Uint16Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector4sArray", vectors)}
}
func (ua *Uint16Attribute) GetW(index int) float64 {
	return ua.Call("getW", index).Float()
}
func (ua *Uint16Attribute) GetX(index int) float64 {
	return ua.Call("getX", index).Float()
}
func (ua *Uint16Attribute) GetY(index int) float64 {
	return ua.Call("getY", index).Float()
}
func (ua *Uint16Attribute) GetZ(index int) float64 {
	return ua.Call("getZ", index).Float()
}
func (ua *Uint16Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("set", value, offset)}
}
func (ua *Uint16Attribute) SetArray2(array js.Value) {
	ua.Call("setArray", array)
}
func (ua *Uint16Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setDynamic", dynamic)}
}
func (ua *Uint16Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setW", index, z)}
}
func (ua *Uint16Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setX", index, x)}
}
func (ua *Uint16Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXY", index, x, y)}
}
func (ua *Uint16Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXYZ", index, x, y, z)}
}
func (ua *Uint16Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXYZW", index, x, y, z, w)}
}
func (ua *Uint16Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setY", index, y)}
}
func (ua *Uint16Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setZ", index, z)}
}

// Uint16BufferAttribute extend: [BufferAttribute]
type Uint16BufferAttribute struct {
	js.Value
}

func NewUint16BufferAttribute(array js.Value, itemSize int, normalized bool) *Uint16BufferAttribute {
	return &Uint16BufferAttribute{Value: get("Uint16BufferAttribute").New(array, itemSize, normalized)}
}
func (uba *Uint16BufferAttribute) JSValue() js.Value {
	return uba.Value
}
func (uba *Uint16BufferAttribute) Array() js.Value {
	return uba.Get("array")
}
func (uba *Uint16BufferAttribute) SetArray(v js.Value) {
	uba.Set("array", v)
}
func (uba *Uint16BufferAttribute) Count() int {
	return uba.Get("count").Int()
}
func (uba *Uint16BufferAttribute) SetCount(v int) {
	uba.Set("count", v)
}
func (uba *Uint16BufferAttribute) Dynamic() bool {
	return uba.Get("dynamic").Bool()
}
func (uba *Uint16BufferAttribute) SetDynamic(v bool) {
	uba.Set("dynamic", v)
}
func (uba *Uint16BufferAttribute) ItemSize() int {
	return uba.Get("itemSize").Int()
}
func (uba *Uint16BufferAttribute) SetItemSize(v int) {
	uba.Set("itemSize", v)
}
func (uba *Uint16BufferAttribute) Length() int {
	return uba.Get("length").Int()
}
func (uba *Uint16BufferAttribute) SetLength(v int) {
	uba.Set("length", v)
}
func (uba *Uint16BufferAttribute) NeedsUpdate() bool {
	return uba.Get("needsUpdate").Bool()
}
func (uba *Uint16BufferAttribute) SetNeedsUpdate(v bool) {
	uba.Set("needsUpdate", v)
}
func (uba *Uint16BufferAttribute) Normalized() bool {
	return uba.Get("normalized").Bool()
}
func (uba *Uint16BufferAttribute) SetNormalized(v bool) {
	uba.Set("normalized", v)
}
func (uba *Uint16BufferAttribute) OnUpload() js.Value {
	return uba.Get("onUpload")
}
func (uba *Uint16BufferAttribute) SetOnUpload(v js.Value) {
	uba.Set("onUpload", v)
}
func (uba *Uint16BufferAttribute) UpdateRange() js.Value {
	return uba.Get("updateRange")
}
func (uba *Uint16BufferAttribute) SetUpdateRange(v js.Value) {
	uba.Set("updateRange", v)
}
func (uba *Uint16BufferAttribute) Uuid() string {
	return uba.Get("uuid").String()
}
func (uba *Uint16BufferAttribute) SetUuid(v string) {
	uba.Set("uuid", v)
}
func (uba *Uint16BufferAttribute) Version() float64 {
	return uba.Get("version").Float()
}
func (uba *Uint16BufferAttribute) SetVersion(v float64) {
	uba.Set("version", v)
}
func (uba *Uint16BufferAttribute) Clone() *Uint16BufferAttribute {
	return &Uint16BufferAttribute{Value: uba.Call("clone")}
}
func (uba *Uint16BufferAttribute) Copy(source *BufferAttribute) *Uint16BufferAttribute {
	return &Uint16BufferAttribute{Value: uba.Call("copy", source)}
}
func (uba *Uint16BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyArray", array)}
}
func (uba *Uint16BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyAt", index1, attribute, index2)}
}
func (uba *Uint16BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyColorsArray", colors)}
}
func (uba *Uint16BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector2sArray", vectors)}
}
func (uba *Uint16BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector3sArray", vectors)}
}
func (uba *Uint16BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector4sArray", vectors)}
}
func (uba *Uint16BufferAttribute) GetW(index int) float64 {
	return uba.Call("getW", index).Float()
}
func (uba *Uint16BufferAttribute) GetX(index int) float64 {
	return uba.Call("getX", index).Float()
}
func (uba *Uint16BufferAttribute) GetY(index int) float64 {
	return uba.Call("getY", index).Float()
}
func (uba *Uint16BufferAttribute) GetZ(index int) float64 {
	return uba.Call("getZ", index).Float()
}
func (uba *Uint16BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("set", value, offset)}
}
func (uba *Uint16BufferAttribute) SetArray2(array js.Value) {
	uba.Call("setArray", array)
}
func (uba *Uint16BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setDynamic", dynamic)}
}
func (uba *Uint16BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setW", index, z)}
}
func (uba *Uint16BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setX", index, x)}
}
func (uba *Uint16BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXY", index, x, y)}
}
func (uba *Uint16BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXYZ", index, x, y, z)}
}
func (uba *Uint16BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXYZW", index, x, y, z, w)}
}
func (uba *Uint16BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setY", index, y)}
}
func (uba *Uint16BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setZ", index, z)}
}

// Uint32Attribute extend: [BufferAttribute]
type Uint32Attribute struct {
	js.Value
}

func NewUint32Attribute(array js.Value, itemSize int) *Uint32Attribute {
	return &Uint32Attribute{Value: get("Uint32Attribute").New(array, itemSize)}
}
func (ua *Uint32Attribute) JSValue() js.Value {
	return ua.Value
}
func (ua *Uint32Attribute) Array() js.Value {
	return ua.Get("array")
}
func (ua *Uint32Attribute) SetArray(v js.Value) {
	ua.Set("array", v)
}
func (ua *Uint32Attribute) Count() int {
	return ua.Get("count").Int()
}
func (ua *Uint32Attribute) SetCount(v int) {
	ua.Set("count", v)
}
func (ua *Uint32Attribute) Dynamic() bool {
	return ua.Get("dynamic").Bool()
}
func (ua *Uint32Attribute) SetDynamic(v bool) {
	ua.Set("dynamic", v)
}
func (ua *Uint32Attribute) ItemSize() int {
	return ua.Get("itemSize").Int()
}
func (ua *Uint32Attribute) SetItemSize(v int) {
	ua.Set("itemSize", v)
}
func (ua *Uint32Attribute) Length() int {
	return ua.Get("length").Int()
}
func (ua *Uint32Attribute) SetLength(v int) {
	ua.Set("length", v)
}
func (ua *Uint32Attribute) NeedsUpdate() bool {
	return ua.Get("needsUpdate").Bool()
}
func (ua *Uint32Attribute) SetNeedsUpdate(v bool) {
	ua.Set("needsUpdate", v)
}
func (ua *Uint32Attribute) Normalized() bool {
	return ua.Get("normalized").Bool()
}
func (ua *Uint32Attribute) SetNormalized(v bool) {
	ua.Set("normalized", v)
}
func (ua *Uint32Attribute) OnUpload() js.Value {
	return ua.Get("onUpload")
}
func (ua *Uint32Attribute) SetOnUpload(v js.Value) {
	ua.Set("onUpload", v)
}
func (ua *Uint32Attribute) UpdateRange() js.Value {
	return ua.Get("updateRange")
}
func (ua *Uint32Attribute) SetUpdateRange(v js.Value) {
	ua.Set("updateRange", v)
}
func (ua *Uint32Attribute) Uuid() string {
	return ua.Get("uuid").String()
}
func (ua *Uint32Attribute) SetUuid(v string) {
	ua.Set("uuid", v)
}
func (ua *Uint32Attribute) Version() float64 {
	return ua.Get("version").Float()
}
func (ua *Uint32Attribute) SetVersion(v float64) {
	ua.Set("version", v)
}
func (ua *Uint32Attribute) Clone() *Uint32Attribute {
	return &Uint32Attribute{Value: ua.Call("clone")}
}
func (ua *Uint32Attribute) Copy(source *BufferAttribute) *Uint32Attribute {
	return &Uint32Attribute{Value: ua.Call("copy", source)}
}
func (ua *Uint32Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyArray", array)}
}
func (ua *Uint32Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyAt", index1, attribute, index2)}
}
func (ua *Uint32Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyColorsArray", colors)}
}
func (ua *Uint32Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector2sArray", vectors)}
}
func (ua *Uint32Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector3sArray", vectors)}
}
func (ua *Uint32Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector4sArray", vectors)}
}
func (ua *Uint32Attribute) GetW(index int) float64 {
	return ua.Call("getW", index).Float()
}
func (ua *Uint32Attribute) GetX(index int) float64 {
	return ua.Call("getX", index).Float()
}
func (ua *Uint32Attribute) GetY(index int) float64 {
	return ua.Call("getY", index).Float()
}
func (ua *Uint32Attribute) GetZ(index int) float64 {
	return ua.Call("getZ", index).Float()
}
func (ua *Uint32Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("set", value, offset)}
}
func (ua *Uint32Attribute) SetArray2(array js.Value) {
	ua.Call("setArray", array)
}
func (ua *Uint32Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setDynamic", dynamic)}
}
func (ua *Uint32Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setW", index, z)}
}
func (ua *Uint32Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setX", index, x)}
}
func (ua *Uint32Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXY", index, x, y)}
}
func (ua *Uint32Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXYZ", index, x, y, z)}
}
func (ua *Uint32Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXYZW", index, x, y, z, w)}
}
func (ua *Uint32Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setY", index, y)}
}
func (ua *Uint32Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setZ", index, z)}
}

// Uint32BufferAttribute extend: [BufferAttribute]
type Uint32BufferAttribute struct {
	js.Value
}

func NewUint32BufferAttribute(array js.Value, itemSize int, normalized bool) *Uint32BufferAttribute {
	return &Uint32BufferAttribute{Value: get("Uint32BufferAttribute").New(array, itemSize, normalized)}
}
func (uba *Uint32BufferAttribute) JSValue() js.Value {
	return uba.Value
}
func (uba *Uint32BufferAttribute) Array() js.Value {
	return uba.Get("array")
}
func (uba *Uint32BufferAttribute) SetArray(v js.Value) {
	uba.Set("array", v)
}
func (uba *Uint32BufferAttribute) Count() int {
	return uba.Get("count").Int()
}
func (uba *Uint32BufferAttribute) SetCount(v int) {
	uba.Set("count", v)
}
func (uba *Uint32BufferAttribute) Dynamic() bool {
	return uba.Get("dynamic").Bool()
}
func (uba *Uint32BufferAttribute) SetDynamic(v bool) {
	uba.Set("dynamic", v)
}
func (uba *Uint32BufferAttribute) ItemSize() int {
	return uba.Get("itemSize").Int()
}
func (uba *Uint32BufferAttribute) SetItemSize(v int) {
	uba.Set("itemSize", v)
}
func (uba *Uint32BufferAttribute) Length() int {
	return uba.Get("length").Int()
}
func (uba *Uint32BufferAttribute) SetLength(v int) {
	uba.Set("length", v)
}
func (uba *Uint32BufferAttribute) NeedsUpdate() bool {
	return uba.Get("needsUpdate").Bool()
}
func (uba *Uint32BufferAttribute) SetNeedsUpdate(v bool) {
	uba.Set("needsUpdate", v)
}
func (uba *Uint32BufferAttribute) Normalized() bool {
	return uba.Get("normalized").Bool()
}
func (uba *Uint32BufferAttribute) SetNormalized(v bool) {
	uba.Set("normalized", v)
}
func (uba *Uint32BufferAttribute) OnUpload() js.Value {
	return uba.Get("onUpload")
}
func (uba *Uint32BufferAttribute) SetOnUpload(v js.Value) {
	uba.Set("onUpload", v)
}
func (uba *Uint32BufferAttribute) UpdateRange() js.Value {
	return uba.Get("updateRange")
}
func (uba *Uint32BufferAttribute) SetUpdateRange(v js.Value) {
	uba.Set("updateRange", v)
}
func (uba *Uint32BufferAttribute) Uuid() string {
	return uba.Get("uuid").String()
}
func (uba *Uint32BufferAttribute) SetUuid(v string) {
	uba.Set("uuid", v)
}
func (uba *Uint32BufferAttribute) Version() float64 {
	return uba.Get("version").Float()
}
func (uba *Uint32BufferAttribute) SetVersion(v float64) {
	uba.Set("version", v)
}
func (uba *Uint32BufferAttribute) Clone() *Uint32BufferAttribute {
	return &Uint32BufferAttribute{Value: uba.Call("clone")}
}
func (uba *Uint32BufferAttribute) Copy(source *BufferAttribute) *Uint32BufferAttribute {
	return &Uint32BufferAttribute{Value: uba.Call("copy", source)}
}
func (uba *Uint32BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyArray", array)}
}
func (uba *Uint32BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyAt", index1, attribute, index2)}
}
func (uba *Uint32BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyColorsArray", colors)}
}
func (uba *Uint32BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector2sArray", vectors)}
}
func (uba *Uint32BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector3sArray", vectors)}
}
func (uba *Uint32BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector4sArray", vectors)}
}
func (uba *Uint32BufferAttribute) GetW(index int) float64 {
	return uba.Call("getW", index).Float()
}
func (uba *Uint32BufferAttribute) GetX(index int) float64 {
	return uba.Call("getX", index).Float()
}
func (uba *Uint32BufferAttribute) GetY(index int) float64 {
	return uba.Call("getY", index).Float()
}
func (uba *Uint32BufferAttribute) GetZ(index int) float64 {
	return uba.Call("getZ", index).Float()
}
func (uba *Uint32BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("set", value, offset)}
}
func (uba *Uint32BufferAttribute) SetArray2(array js.Value) {
	uba.Call("setArray", array)
}
func (uba *Uint32BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setDynamic", dynamic)}
}
func (uba *Uint32BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setW", index, z)}
}
func (uba *Uint32BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setX", index, x)}
}
func (uba *Uint32BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXY", index, x, y)}
}
func (uba *Uint32BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXYZ", index, x, y, z)}
}
func (uba *Uint32BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXYZW", index, x, y, z, w)}
}
func (uba *Uint32BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setY", index, y)}
}
func (uba *Uint32BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setZ", index, z)}
}

// Uint8Attribute extend: [BufferAttribute]
type Uint8Attribute struct {
	js.Value
}

func NewUint8Attribute(array js.Value, itemSize int) *Uint8Attribute {
	return &Uint8Attribute{Value: get("Uint8Attribute").New(array, itemSize)}
}
func (ua *Uint8Attribute) JSValue() js.Value {
	return ua.Value
}
func (ua *Uint8Attribute) Array() js.Value {
	return ua.Get("array")
}
func (ua *Uint8Attribute) SetArray(v js.Value) {
	ua.Set("array", v)
}
func (ua *Uint8Attribute) Count() int {
	return ua.Get("count").Int()
}
func (ua *Uint8Attribute) SetCount(v int) {
	ua.Set("count", v)
}
func (ua *Uint8Attribute) Dynamic() bool {
	return ua.Get("dynamic").Bool()
}
func (ua *Uint8Attribute) SetDynamic(v bool) {
	ua.Set("dynamic", v)
}
func (ua *Uint8Attribute) ItemSize() int {
	return ua.Get("itemSize").Int()
}
func (ua *Uint8Attribute) SetItemSize(v int) {
	ua.Set("itemSize", v)
}
func (ua *Uint8Attribute) Length() int {
	return ua.Get("length").Int()
}
func (ua *Uint8Attribute) SetLength(v int) {
	ua.Set("length", v)
}
func (ua *Uint8Attribute) NeedsUpdate() bool {
	return ua.Get("needsUpdate").Bool()
}
func (ua *Uint8Attribute) SetNeedsUpdate(v bool) {
	ua.Set("needsUpdate", v)
}
func (ua *Uint8Attribute) Normalized() bool {
	return ua.Get("normalized").Bool()
}
func (ua *Uint8Attribute) SetNormalized(v bool) {
	ua.Set("normalized", v)
}
func (ua *Uint8Attribute) OnUpload() js.Value {
	return ua.Get("onUpload")
}
func (ua *Uint8Attribute) SetOnUpload(v js.Value) {
	ua.Set("onUpload", v)
}
func (ua *Uint8Attribute) UpdateRange() js.Value {
	return ua.Get("updateRange")
}
func (ua *Uint8Attribute) SetUpdateRange(v js.Value) {
	ua.Set("updateRange", v)
}
func (ua *Uint8Attribute) Uuid() string {
	return ua.Get("uuid").String()
}
func (ua *Uint8Attribute) SetUuid(v string) {
	ua.Set("uuid", v)
}
func (ua *Uint8Attribute) Version() float64 {
	return ua.Get("version").Float()
}
func (ua *Uint8Attribute) SetVersion(v float64) {
	ua.Set("version", v)
}
func (ua *Uint8Attribute) Clone() *Uint8Attribute {
	return &Uint8Attribute{Value: ua.Call("clone")}
}
func (ua *Uint8Attribute) Copy(source *BufferAttribute) *Uint8Attribute {
	return &Uint8Attribute{Value: ua.Call("copy", source)}
}
func (ua *Uint8Attribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyArray", array)}
}
func (ua *Uint8Attribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyAt", index1, attribute, index2)}
}
func (ua *Uint8Attribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyColorsArray", colors)}
}
func (ua *Uint8Attribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector2sArray", vectors)}
}
func (ua *Uint8Attribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector3sArray", vectors)}
}
func (ua *Uint8Attribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("copyVector4sArray", vectors)}
}
func (ua *Uint8Attribute) GetW(index int) float64 {
	return ua.Call("getW", index).Float()
}
func (ua *Uint8Attribute) GetX(index int) float64 {
	return ua.Call("getX", index).Float()
}
func (ua *Uint8Attribute) GetY(index int) float64 {
	return ua.Call("getY", index).Float()
}
func (ua *Uint8Attribute) GetZ(index int) float64 {
	return ua.Call("getZ", index).Float()
}
func (ua *Uint8Attribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("set", value, offset)}
}
func (ua *Uint8Attribute) SetArray2(array js.Value) {
	ua.Call("setArray", array)
}
func (ua *Uint8Attribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setDynamic", dynamic)}
}
func (ua *Uint8Attribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setW", index, z)}
}
func (ua *Uint8Attribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setX", index, x)}
}
func (ua *Uint8Attribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXY", index, x, y)}
}
func (ua *Uint8Attribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXYZ", index, x, y, z)}
}
func (ua *Uint8Attribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setXYZW", index, x, y, z, w)}
}
func (ua *Uint8Attribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setY", index, y)}
}
func (ua *Uint8Attribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ua.Call("setZ", index, z)}
}

// Uint8BufferAttribute extend: [BufferAttribute]
type Uint8BufferAttribute struct {
	js.Value
}

func NewUint8BufferAttribute(array js.Value, itemSize int, normalized bool) *Uint8BufferAttribute {
	return &Uint8BufferAttribute{Value: get("Uint8BufferAttribute").New(array, itemSize, normalized)}
}
func (uba *Uint8BufferAttribute) JSValue() js.Value {
	return uba.Value
}
func (uba *Uint8BufferAttribute) Array() js.Value {
	return uba.Get("array")
}
func (uba *Uint8BufferAttribute) SetArray(v js.Value) {
	uba.Set("array", v)
}
func (uba *Uint8BufferAttribute) Count() int {
	return uba.Get("count").Int()
}
func (uba *Uint8BufferAttribute) SetCount(v int) {
	uba.Set("count", v)
}
func (uba *Uint8BufferAttribute) Dynamic() bool {
	return uba.Get("dynamic").Bool()
}
func (uba *Uint8BufferAttribute) SetDynamic(v bool) {
	uba.Set("dynamic", v)
}
func (uba *Uint8BufferAttribute) ItemSize() int {
	return uba.Get("itemSize").Int()
}
func (uba *Uint8BufferAttribute) SetItemSize(v int) {
	uba.Set("itemSize", v)
}
func (uba *Uint8BufferAttribute) Length() int {
	return uba.Get("length").Int()
}
func (uba *Uint8BufferAttribute) SetLength(v int) {
	uba.Set("length", v)
}
func (uba *Uint8BufferAttribute) NeedsUpdate() bool {
	return uba.Get("needsUpdate").Bool()
}
func (uba *Uint8BufferAttribute) SetNeedsUpdate(v bool) {
	uba.Set("needsUpdate", v)
}
func (uba *Uint8BufferAttribute) Normalized() bool {
	return uba.Get("normalized").Bool()
}
func (uba *Uint8BufferAttribute) SetNormalized(v bool) {
	uba.Set("normalized", v)
}
func (uba *Uint8BufferAttribute) OnUpload() js.Value {
	return uba.Get("onUpload")
}
func (uba *Uint8BufferAttribute) SetOnUpload(v js.Value) {
	uba.Set("onUpload", v)
}
func (uba *Uint8BufferAttribute) UpdateRange() js.Value {
	return uba.Get("updateRange")
}
func (uba *Uint8BufferAttribute) SetUpdateRange(v js.Value) {
	uba.Set("updateRange", v)
}
func (uba *Uint8BufferAttribute) Uuid() string {
	return uba.Get("uuid").String()
}
func (uba *Uint8BufferAttribute) SetUuid(v string) {
	uba.Set("uuid", v)
}
func (uba *Uint8BufferAttribute) Version() float64 {
	return uba.Get("version").Float()
}
func (uba *Uint8BufferAttribute) SetVersion(v float64) {
	uba.Set("version", v)
}
func (uba *Uint8BufferAttribute) Clone() *Uint8BufferAttribute {
	return &Uint8BufferAttribute{Value: uba.Call("clone")}
}
func (uba *Uint8BufferAttribute) Copy(source *BufferAttribute) *Uint8BufferAttribute {
	return &Uint8BufferAttribute{Value: uba.Call("copy", source)}
}
func (uba *Uint8BufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyArray", array)}
}
func (uba *Uint8BufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyAt", index1, attribute, index2)}
}
func (uba *Uint8BufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyColorsArray", colors)}
}
func (uba *Uint8BufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector2sArray", vectors)}
}
func (uba *Uint8BufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector3sArray", vectors)}
}
func (uba *Uint8BufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("copyVector4sArray", vectors)}
}
func (uba *Uint8BufferAttribute) GetW(index int) float64 {
	return uba.Call("getW", index).Float()
}
func (uba *Uint8BufferAttribute) GetX(index int) float64 {
	return uba.Call("getX", index).Float()
}
func (uba *Uint8BufferAttribute) GetY(index int) float64 {
	return uba.Call("getY", index).Float()
}
func (uba *Uint8BufferAttribute) GetZ(index int) float64 {
	return uba.Call("getZ", index).Float()
}
func (uba *Uint8BufferAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("set", value, offset)}
}
func (uba *Uint8BufferAttribute) SetArray2(array js.Value) {
	uba.Call("setArray", array)
}
func (uba *Uint8BufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setDynamic", dynamic)}
}
func (uba *Uint8BufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setW", index, z)}
}
func (uba *Uint8BufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setX", index, x)}
}
func (uba *Uint8BufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXY", index, x, y)}
}
func (uba *Uint8BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXYZ", index, x, y, z)}
}
func (uba *Uint8BufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setXYZW", index, x, y, z, w)}
}
func (uba *Uint8BufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setY", index, y)}
}
func (uba *Uint8BufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uba.Call("setZ", index, z)}
}

// Uint8ClampedAttribute extend: [BufferAttribute]
type Uint8ClampedAttribute struct {
	js.Value
}

func NewUint8ClampedAttribute(array js.Value, itemSize int) *Uint8ClampedAttribute {
	return &Uint8ClampedAttribute{Value: get("Uint8ClampedAttribute").New(array, itemSize)}
}
func (uca *Uint8ClampedAttribute) JSValue() js.Value {
	return uca.Value
}
func (uca *Uint8ClampedAttribute) Array() js.Value {
	return uca.Get("array")
}
func (uca *Uint8ClampedAttribute) SetArray(v js.Value) {
	uca.Set("array", v)
}
func (uca *Uint8ClampedAttribute) Count() int {
	return uca.Get("count").Int()
}
func (uca *Uint8ClampedAttribute) SetCount(v int) {
	uca.Set("count", v)
}
func (uca *Uint8ClampedAttribute) Dynamic() bool {
	return uca.Get("dynamic").Bool()
}
func (uca *Uint8ClampedAttribute) SetDynamic(v bool) {
	uca.Set("dynamic", v)
}
func (uca *Uint8ClampedAttribute) ItemSize() int {
	return uca.Get("itemSize").Int()
}
func (uca *Uint8ClampedAttribute) SetItemSize(v int) {
	uca.Set("itemSize", v)
}
func (uca *Uint8ClampedAttribute) Length() int {
	return uca.Get("length").Int()
}
func (uca *Uint8ClampedAttribute) SetLength(v int) {
	uca.Set("length", v)
}
func (uca *Uint8ClampedAttribute) NeedsUpdate() bool {
	return uca.Get("needsUpdate").Bool()
}
func (uca *Uint8ClampedAttribute) SetNeedsUpdate(v bool) {
	uca.Set("needsUpdate", v)
}
func (uca *Uint8ClampedAttribute) Normalized() bool {
	return uca.Get("normalized").Bool()
}
func (uca *Uint8ClampedAttribute) SetNormalized(v bool) {
	uca.Set("normalized", v)
}
func (uca *Uint8ClampedAttribute) OnUpload() js.Value {
	return uca.Get("onUpload")
}
func (uca *Uint8ClampedAttribute) SetOnUpload(v js.Value) {
	uca.Set("onUpload", v)
}
func (uca *Uint8ClampedAttribute) UpdateRange() js.Value {
	return uca.Get("updateRange")
}
func (uca *Uint8ClampedAttribute) SetUpdateRange(v js.Value) {
	uca.Set("updateRange", v)
}
func (uca *Uint8ClampedAttribute) Uuid() string {
	return uca.Get("uuid").String()
}
func (uca *Uint8ClampedAttribute) SetUuid(v string) {
	uca.Set("uuid", v)
}
func (uca *Uint8ClampedAttribute) Version() float64 {
	return uca.Get("version").Float()
}
func (uca *Uint8ClampedAttribute) SetVersion(v float64) {
	uca.Set("version", v)
}
func (uca *Uint8ClampedAttribute) Clone() *Uint8ClampedAttribute {
	return &Uint8ClampedAttribute{Value: uca.Call("clone")}
}
func (uca *Uint8ClampedAttribute) Copy(source *BufferAttribute) *Uint8ClampedAttribute {
	return &Uint8ClampedAttribute{Value: uca.Call("copy", source)}
}
func (uca *Uint8ClampedAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("copyArray", array)}
}
func (uca *Uint8ClampedAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("copyAt", index1, attribute, index2)}
}
func (uca *Uint8ClampedAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("copyColorsArray", colors)}
}
func (uca *Uint8ClampedAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("copyVector2sArray", vectors)}
}
func (uca *Uint8ClampedAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("copyVector3sArray", vectors)}
}
func (uca *Uint8ClampedAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("copyVector4sArray", vectors)}
}
func (uca *Uint8ClampedAttribute) GetW(index int) float64 {
	return uca.Call("getW", index).Float()
}
func (uca *Uint8ClampedAttribute) GetX(index int) float64 {
	return uca.Call("getX", index).Float()
}
func (uca *Uint8ClampedAttribute) GetY(index int) float64 {
	return uca.Call("getY", index).Float()
}
func (uca *Uint8ClampedAttribute) GetZ(index int) float64 {
	return uca.Call("getZ", index).Float()
}
func (uca *Uint8ClampedAttribute) Set2(value js.Value, offset int) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("set", value, offset)}
}
func (uca *Uint8ClampedAttribute) SetArray2(array js.Value) {
	uca.Call("setArray", array)
}
func (uca *Uint8ClampedAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setDynamic", dynamic)}
}
func (uca *Uint8ClampedAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setW", index, z)}
}
func (uca *Uint8ClampedAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setX", index, x)}
}
func (uca *Uint8ClampedAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setXY", index, x, y)}
}
func (uca *Uint8ClampedAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setXYZ", index, x, y, z)}
}
func (uca *Uint8ClampedAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setXYZW", index, x, y, z, w)}
}
func (uca *Uint8ClampedAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setY", index, y)}
}
func (uca *Uint8ClampedAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: uca.Call("setZ", index, z)}
}

// Uint8ClampedBufferAttribute extend: [BufferAttribute]
type Uint8ClampedBufferAttribute struct {
	js.Value
}

func NewUint8ClampedBufferAttribute(array js.Value, itemSize int, normalized bool) *Uint8ClampedBufferAttribute {
	return &Uint8ClampedBufferAttribute{Value: get("Uint8ClampedBufferAttribute").New(array, itemSize, normalized)}
}
func (ucba *Uint8ClampedBufferAttribute) JSValue() js.Value {
	return ucba.Value
}
func (ucba *Uint8ClampedBufferAttribute) Array() js.Value {
	return ucba.Get("array")
}
func (ucba *Uint8ClampedBufferAttribute) SetArray(v js.Value) {
	ucba.Set("array", v)
}
func (ucba *Uint8ClampedBufferAttribute) Count() int {
	return ucba.Get("count").Int()
}
func (ucba *Uint8ClampedBufferAttribute) SetCount(v int) {
	ucba.Set("count", v)
}
func (ucba *Uint8ClampedBufferAttribute) Dynamic() bool {
	return ucba.Get("dynamic").Bool()
}
func (ucba *Uint8ClampedBufferAttribute) SetDynamic(v bool) {
	ucba.Set("dynamic", v)
}
func (ucba *Uint8ClampedBufferAttribute) ItemSize() int {
	return ucba.Get("itemSize").Int()
}
func (ucba *Uint8ClampedBufferAttribute) SetItemSize(v int) {
	ucba.Set("itemSize", v)
}
func (ucba *Uint8ClampedBufferAttribute) Length() int {
	return ucba.Get("length").Int()
}
func (ucba *Uint8ClampedBufferAttribute) SetLength(v int) {
	ucba.Set("length", v)
}
func (ucba *Uint8ClampedBufferAttribute) NeedsUpdate() bool {
	return ucba.Get("needsUpdate").Bool()
}
func (ucba *Uint8ClampedBufferAttribute) SetNeedsUpdate(v bool) {
	ucba.Set("needsUpdate", v)
}
func (ucba *Uint8ClampedBufferAttribute) Normalized() bool {
	return ucba.Get("normalized").Bool()
}
func (ucba *Uint8ClampedBufferAttribute) SetNormalized(v bool) {
	ucba.Set("normalized", v)
}
func (ucba *Uint8ClampedBufferAttribute) OnUpload() js.Value {
	return ucba.Get("onUpload")
}
func (ucba *Uint8ClampedBufferAttribute) SetOnUpload(v js.Value) {
	ucba.Set("onUpload", v)
}
func (ucba *Uint8ClampedBufferAttribute) UpdateRange() js.Value {
	return ucba.Get("updateRange")
}
func (ucba *Uint8ClampedBufferAttribute) SetUpdateRange(v js.Value) {
	ucba.Set("updateRange", v)
}
func (ucba *Uint8ClampedBufferAttribute) Uuid() string {
	return ucba.Get("uuid").String()
}
func (ucba *Uint8ClampedBufferAttribute) SetUuid(v string) {
	ucba.Set("uuid", v)
}
func (ucba *Uint8ClampedBufferAttribute) Version() float64 {
	return ucba.Get("version").Float()
}
func (ucba *Uint8ClampedBufferAttribute) SetVersion(v float64) {
	ucba.Set("version", v)
}
func (ucba *Uint8ClampedBufferAttribute) Clone() *Uint8ClampedBufferAttribute {
	return &Uint8ClampedBufferAttribute{Value: ucba.Call("clone")}
}
func (ucba *Uint8ClampedBufferAttribute) Copy(source *BufferAttribute) *Uint8ClampedBufferAttribute {
	return &Uint8ClampedBufferAttribute{Value: ucba.Call("copy", source)}
}
func (ucba *Uint8ClampedBufferAttribute) CopyArray(array js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("copyArray", array)}
}
func (ucba *Uint8ClampedBufferAttribute) CopyAt(index1 int, attribute *BufferAttribute, index2 int) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("copyAt", index1, attribute, index2)}
}
func (ucba *Uint8ClampedBufferAttribute) CopyColorsArray(colors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("copyColorsArray", colors)}
}
func (ucba *Uint8ClampedBufferAttribute) CopyVector2sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("copyVector2sArray", vectors)}
}
func (ucba *Uint8ClampedBufferAttribute) CopyVector3sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("copyVector3sArray", vectors)}
}
func (ucba *Uint8ClampedBufferAttribute) CopyVector4sArray(vectors js.Value) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("copyVector4sArray", vectors)}
}
func (ucba *Uint8ClampedBufferAttribute) GetW(index int) float64 {
	return ucba.Call("getW", index).Float()
}
func (ucba *Uint8ClampedBufferAttribute) GetX(index int) float64 {
	return ucba.Call("getX", index).Float()
}
func (ucba *Uint8ClampedBufferAttribute) GetY(index int) float64 {
	return ucba.Call("getY", index).Float()
}
func (ucba *Uint8ClampedBufferAttribute) GetZ(index int) float64 {
	return ucba.Call("getZ", index).Float()
}
func (ucba *Uint8ClampedBufferAttribute) Set2(value js.Value, offset float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("set", value, offset)}
}
func (ucba *Uint8ClampedBufferAttribute) SetArray2(array js.Value) {
	ucba.Call("setArray", array)
}
func (ucba *Uint8ClampedBufferAttribute) SetDynamic2(dynamic bool) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setDynamic", dynamic)}
}
func (ucba *Uint8ClampedBufferAttribute) SetW(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setW", index, z)}
}
func (ucba *Uint8ClampedBufferAttribute) SetX(index int, x float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setX", index, x)}
}
func (ucba *Uint8ClampedBufferAttribute) SetXY(index int, x float64, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setXY", index, x, y)}
}
func (ucba *Uint8ClampedBufferAttribute) SetXYZ(index int, x float64, y float64, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setXYZ", index, x, y, z)}
}
func (ucba *Uint8ClampedBufferAttribute) SetXYZW(index int, x float64, y float64, z float64, w float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setXYZW", index, x, y, z, w)}
}
func (ucba *Uint8ClampedBufferAttribute) SetY(index int, y float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setY", index, y)}
}
func (ucba *Uint8ClampedBufferAttribute) SetZ(index int, z float64) *BufferAttribute {
	return &BufferAttribute{Value: ucba.Call("setZ", index, z)}
}

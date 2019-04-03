// Code generated from "textures/DataTexture.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type DataTexture struct {
	js.Value
}

func NewDataTexture(data js.Value, width float64, height float64, format PixelFormat, typ TextureDataType, mapping Mapping, wrapS Wrapping, wrapT Wrapping, magFilter TextureFilter, minFilter TextureFilter, anisotropy float64, encoding TextureEncoding) *DataTexture {
	return &DataTexture{Value: get("DataTexture").New(data, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy, encoding)}
}
func (dt *DataTexture) Anisotropy() float64 {
	return dt.Get("anisotropy").Float()
}
func (dt *DataTexture) SetAnisotropy(v float64) {
	dt.Set("anisotropy", v)
}
func (dt *DataTexture) Center() *Vector2 {
	return &Vector2{Value: dt.Get("center")}
}
func (dt *DataTexture) SetCenter(v *Vector2) {
	dt.Set("center", v)
}
func (dt *DataTexture) Encoding() TextureEncoding {
	return TextureEncoding(dt.Get("encoding"))
}
func (dt *DataTexture) SetEncoding(v TextureEncoding) {
	dt.Set("encoding", v)
}
func (dt *DataTexture) FlipY() bool {
	return dt.Get("flipY").Bool()
}
func (dt *DataTexture) SetFlipY(v bool) {
	dt.Set("flipY", v)
}
func (dt *DataTexture) Format() PixelFormat {
	return PixelFormat(dt.Get("format"))
}
func (dt *DataTexture) SetFormat(v PixelFormat) {
	dt.Set("format", v)
}
func (dt *DataTexture) GenerateMipmaps() bool {
	return dt.Get("generateMipmaps").Bool()
}
func (dt *DataTexture) SetGenerateMipmaps(v bool) {
	dt.Set("generateMipmaps", v)
}
func (dt *DataTexture) Id() int {
	return dt.Get("id").Int()
}
func (dt *DataTexture) SetId(v int) {
	dt.Set("id", v)
}
func (dt *DataTexture) Image() js.Value {
	return dt.Get("image")
}
func (dt *DataTexture) SetImage(v js.Value) {
	dt.Set("image", v)
}
func (dt *DataTexture) MagFilter() TextureFilter {
	return TextureFilter(dt.Get("magFilter"))
}
func (dt *DataTexture) SetMagFilter(v TextureFilter) {
	dt.Set("magFilter", v)
}
func (dt *DataTexture) Mapping() Mapping {
	return Mapping(dt.Get("mapping"))
}
func (dt *DataTexture) SetMapping(v Mapping) {
	dt.Set("mapping", v)
}
func (dt *DataTexture) MinFilter() TextureFilter {
	return TextureFilter(dt.Get("minFilter"))
}
func (dt *DataTexture) SetMinFilter(v TextureFilter) {
	dt.Set("minFilter", v)
}
func (dt *DataTexture) Mipmaps() js.Value {
	return dt.Get("mipmaps")
}
func (dt *DataTexture) SetMipmaps(v js.Value) {
	dt.Set("mipmaps", v)
}
func (dt *DataTexture) Name() string {
	return dt.Get("name").String()
}
func (dt *DataTexture) SetName(v string) {
	dt.Set("name", v)
}
func (dt *DataTexture) NeedsUpdate() bool {
	return dt.Get("needsUpdate").Bool()
}
func (dt *DataTexture) SetNeedsUpdate(v bool) {
	dt.Set("needsUpdate", v)
}
func (dt *DataTexture) Offset() *Vector2 {
	return &Vector2{Value: dt.Get("offset")}
}
func (dt *DataTexture) SetOffset(v *Vector2) {
	dt.Set("offset", v)
}
func (dt *DataTexture) OnUpdate() js.Value {
	return dt.Get("onUpdate")
}
func (dt *DataTexture) SetOnUpdate(v js.Value) {
	dt.Set("onUpdate", v)
}
func (dt *DataTexture) PremultiplyAlpha() bool {
	return dt.Get("premultiplyAlpha").Bool()
}
func (dt *DataTexture) SetPremultiplyAlpha(v bool) {
	dt.Set("premultiplyAlpha", v)
}
func (dt *DataTexture) Repeat() *Vector2 {
	return &Vector2{Value: dt.Get("repeat")}
}
func (dt *DataTexture) SetRepeat(v *Vector2) {
	dt.Set("repeat", v)
}
func (dt *DataTexture) Rotation() float64 {
	return dt.Get("rotation").Float()
}
func (dt *DataTexture) SetRotation(v float64) {
	dt.Set("rotation", v)
}
func (dt *DataTexture) SourceFile() string {
	return dt.Get("sourceFile").String()
}
func (dt *DataTexture) SetSourceFile(v string) {
	dt.Set("sourceFile", v)
}
func (dt *DataTexture) Type() TextureDataType {
	return TextureDataType(dt.Get("type"))
}
func (dt *DataTexture) SetType(v TextureDataType) {
	dt.Set("type", v)
}
func (dt *DataTexture) UnpackAlignment() float64 {
	return dt.Get("unpackAlignment").Float()
}
func (dt *DataTexture) SetUnpackAlignment(v float64) {
	dt.Set("unpackAlignment", v)
}
func (dt *DataTexture) Uuid() string {
	return dt.Get("uuid").String()
}
func (dt *DataTexture) SetUuid(v string) {
	dt.Set("uuid", v)
}
func (dt *DataTexture) Version() float64 {
	return dt.Get("version").Float()
}
func (dt *DataTexture) SetVersion(v float64) {
	dt.Set("version", v)
}
func (dt *DataTexture) WrapS() Wrapping {
	return Wrapping(dt.Get("wrapS"))
}
func (dt *DataTexture) SetWrapS(v Wrapping) {
	dt.Set("wrapS", v)
}
func (dt *DataTexture) WrapT() Wrapping {
	return Wrapping(dt.Get("wrapT"))
}
func (dt *DataTexture) SetWrapT(v Wrapping) {
	dt.Set("wrapT", v)
}
func (dt *DataTexture) DEFAULT_IMAGE() js.Value {
	return dt.Get("DEFAULT_IMAGE")
}
func (dt *DataTexture) SetDEFAULT_IMAGE(v js.Value) {
	dt.Set("DEFAULT_IMAGE", v)
}
func (dt *DataTexture) DEFAULT_MAPPING() js.Value {
	return dt.Get("DEFAULT_MAPPING")
}
func (dt *DataTexture) SetDEFAULT_MAPPING(v js.Value) {
	dt.Set("DEFAULT_MAPPING", v)
}
func (dt *DataTexture) AddEventListener(typ string, listener js.Value) {
	dt.Call("addEventListener", typ, listener)
}
func (dt *DataTexture) Clone() *DataTexture {
	return &DataTexture{Value: dt.Call("clone")}
}
func (dt *DataTexture) Copy(source *Texture) *DataTexture {
	return &DataTexture{Value: dt.Call("copy", source)}
}
func (dt *DataTexture) DispatchEvent(event js.Value) {
	dt.Call("dispatchEvent", event)
}
func (dt *DataTexture) Dispose() {
	dt.Call("dispose")
}
func (dt *DataTexture) HasEventListener(typ string, listener js.Value) bool {
	return dt.Call("hasEventListener", typ, listener).Bool()
}
func (dt *DataTexture) RemoveEventListener(typ string, listener js.Value) {
	dt.Call("removeEventListener", typ, listener)
}
func (dt *DataTexture) ToJSON(meta js.Value) js.Value {
	return dt.Call("toJSON", meta)
}
func (dt *DataTexture) TransformUv(uv Vector) {
	dt.Call("transformUv", uv)
}

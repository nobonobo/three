// Code generated from "textures/CompressedTexture.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type CompressedTexture struct {
	js.Value
}

func NewCompressedTexture(mipmaps js.Value, width float64, height float64, format PixelFormat, typ TextureDataType, mapping Mapping, wrapS Wrapping, wrapT Wrapping, magFilter TextureFilter, minFilter TextureFilter, anisotropy float64, encoding TextureEncoding) *CompressedTexture {
	return &CompressedTexture{Value: get("CompressedTexture").New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy, encoding)}
}
func (ct *CompressedTexture) Anisotropy() float64 {
	return ct.Get("anisotropy").Float()
}
func (ct *CompressedTexture) SetAnisotropy(v float64) {
	ct.Set("anisotropy", v)
}
func (ct *CompressedTexture) Center() *Vector2 {
	return &Vector2{Value: ct.Get("center")}
}
func (ct *CompressedTexture) SetCenter(v *Vector2) {
	ct.Set("center", v)
}
func (ct *CompressedTexture) Encoding() TextureEncoding {
	return TextureEncoding(ct.Get("encoding"))
}
func (ct *CompressedTexture) SetEncoding(v TextureEncoding) {
	ct.Set("encoding", v)
}
func (ct *CompressedTexture) FlipY() bool {
	return ct.Get("flipY").Bool()
}
func (ct *CompressedTexture) SetFlipY(v bool) {
	ct.Set("flipY", v)
}
func (ct *CompressedTexture) Format() PixelFormat {
	return PixelFormat(ct.Get("format"))
}
func (ct *CompressedTexture) SetFormat(v PixelFormat) {
	ct.Set("format", v)
}
func (ct *CompressedTexture) GenerateMipmaps() bool {
	return ct.Get("generateMipmaps").Bool()
}
func (ct *CompressedTexture) SetGenerateMipmaps(v bool) {
	ct.Set("generateMipmaps", v)
}
func (ct *CompressedTexture) Id() int {
	return ct.Get("id").Int()
}
func (ct *CompressedTexture) SetId(v int) {
	ct.Set("id", v)
}
func (ct *CompressedTexture) Image() js.Value {
	return ct.Get("image")
}
func (ct *CompressedTexture) SetImage(v js.Value) {
	ct.Set("image", v)
}
func (ct *CompressedTexture) MagFilter() TextureFilter {
	return TextureFilter(ct.Get("magFilter"))
}
func (ct *CompressedTexture) SetMagFilter(v TextureFilter) {
	ct.Set("magFilter", v)
}
func (ct *CompressedTexture) Mapping() Mapping {
	return Mapping(ct.Get("mapping"))
}
func (ct *CompressedTexture) SetMapping(v Mapping) {
	ct.Set("mapping", v)
}
func (ct *CompressedTexture) MinFilter() TextureFilter {
	return TextureFilter(ct.Get("minFilter"))
}
func (ct *CompressedTexture) SetMinFilter(v TextureFilter) {
	ct.Set("minFilter", v)
}
func (ct *CompressedTexture) Mipmaps() js.Value {
	return ct.Get("mipmaps")
}
func (ct *CompressedTexture) SetMipmaps(v js.Value) {
	ct.Set("mipmaps", v)
}
func (ct *CompressedTexture) Name() string {
	return ct.Get("name").String()
}
func (ct *CompressedTexture) SetName(v string) {
	ct.Set("name", v)
}
func (ct *CompressedTexture) NeedsUpdate() bool {
	return ct.Get("needsUpdate").Bool()
}
func (ct *CompressedTexture) SetNeedsUpdate(v bool) {
	ct.Set("needsUpdate", v)
}
func (ct *CompressedTexture) Offset() *Vector2 {
	return &Vector2{Value: ct.Get("offset")}
}
func (ct *CompressedTexture) SetOffset(v *Vector2) {
	ct.Set("offset", v)
}
func (ct *CompressedTexture) OnUpdate() js.Value {
	return ct.Get("onUpdate")
}
func (ct *CompressedTexture) SetOnUpdate(v js.Value) {
	ct.Set("onUpdate", v)
}
func (ct *CompressedTexture) PremultiplyAlpha() bool {
	return ct.Get("premultiplyAlpha").Bool()
}
func (ct *CompressedTexture) SetPremultiplyAlpha(v bool) {
	ct.Set("premultiplyAlpha", v)
}
func (ct *CompressedTexture) Repeat() *Vector2 {
	return &Vector2{Value: ct.Get("repeat")}
}
func (ct *CompressedTexture) SetRepeat(v *Vector2) {
	ct.Set("repeat", v)
}
func (ct *CompressedTexture) Rotation() float64 {
	return ct.Get("rotation").Float()
}
func (ct *CompressedTexture) SetRotation(v float64) {
	ct.Set("rotation", v)
}
func (ct *CompressedTexture) SourceFile() string {
	return ct.Get("sourceFile").String()
}
func (ct *CompressedTexture) SetSourceFile(v string) {
	ct.Set("sourceFile", v)
}
func (ct *CompressedTexture) Type() TextureDataType {
	return TextureDataType(ct.Get("type"))
}
func (ct *CompressedTexture) SetType(v TextureDataType) {
	ct.Set("type", v)
}
func (ct *CompressedTexture) UnpackAlignment() float64 {
	return ct.Get("unpackAlignment").Float()
}
func (ct *CompressedTexture) SetUnpackAlignment(v float64) {
	ct.Set("unpackAlignment", v)
}
func (ct *CompressedTexture) Uuid() string {
	return ct.Get("uuid").String()
}
func (ct *CompressedTexture) SetUuid(v string) {
	ct.Set("uuid", v)
}
func (ct *CompressedTexture) Version() float64 {
	return ct.Get("version").Float()
}
func (ct *CompressedTexture) SetVersion(v float64) {
	ct.Set("version", v)
}
func (ct *CompressedTexture) WrapS() Wrapping {
	return Wrapping(ct.Get("wrapS"))
}
func (ct *CompressedTexture) SetWrapS(v Wrapping) {
	ct.Set("wrapS", v)
}
func (ct *CompressedTexture) WrapT() Wrapping {
	return Wrapping(ct.Get("wrapT"))
}
func (ct *CompressedTexture) SetWrapT(v Wrapping) {
	ct.Set("wrapT", v)
}
func (ct *CompressedTexture) DEFAULT_IMAGE() js.Value {
	return ct.Get("DEFAULT_IMAGE")
}
func (ct *CompressedTexture) SetDEFAULT_IMAGE(v js.Value) {
	ct.Set("DEFAULT_IMAGE", v)
}
func (ct *CompressedTexture) DEFAULT_MAPPING() js.Value {
	return ct.Get("DEFAULT_MAPPING")
}
func (ct *CompressedTexture) SetDEFAULT_MAPPING(v js.Value) {
	ct.Set("DEFAULT_MAPPING", v)
}
func (ct *CompressedTexture) AddEventListener(typ string, listener js.Value) {
	ct.Call("addEventListener", typ, listener)
}
func (ct *CompressedTexture) Clone() *CompressedTexture {
	return &CompressedTexture{Value: ct.Call("clone")}
}
func (ct *CompressedTexture) Copy(source *Texture) *CompressedTexture {
	return &CompressedTexture{Value: ct.Call("copy", source)}
}
func (ct *CompressedTexture) DispatchEvent(event js.Value) {
	ct.Call("dispatchEvent", event)
}
func (ct *CompressedTexture) Dispose() {
	ct.Call("dispose")
}
func (ct *CompressedTexture) HasEventListener(typ string, listener js.Value) bool {
	return ct.Call("hasEventListener", typ, listener).Bool()
}
func (ct *CompressedTexture) RemoveEventListener(typ string, listener js.Value) {
	ct.Call("removeEventListener", typ, listener)
}
func (ct *CompressedTexture) ToJSON(meta js.Value) js.Value {
	return ct.Call("toJSON", meta)
}
func (ct *CompressedTexture) TransformUv(uv Vector) {
	ct.Call("transformUv", uv)
}

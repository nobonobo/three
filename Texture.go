// Code generated from "textures/Texture.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func TextureIdCount() int {
	return get("TextureIdCount").Int()
}
func SetTextureIdCount(v int) {
	set("TextureIdCount", v)
}

// Texture extend: [EventDispatcher]
type Texture struct {
	js.Value
}

func NewTexture(image js.Value, mapping Mapping, wrapS Wrapping, wrapT Wrapping, magFilter TextureFilter, minFilter TextureFilter, format PixelFormat, typ TextureDataType, anisotropy float64, encoding TextureEncoding) *Texture {
	return &Texture{Value: get("Texture").New(image, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy, encoding)}
}
func (tt *Texture) JSValue() js.Value {
	return tt.Value
}
func (tt *Texture) Anisotropy() float64 {
	return tt.Get("anisotropy").Float()
}
func (tt *Texture) SetAnisotropy(v float64) {
	tt.Set("anisotropy", v)
}
func (tt *Texture) Center() *Vector2 {
	return &Vector2{Value: tt.Get("center")}
}
func (tt *Texture) SetCenter(v *Vector2) {
	tt.Set("center", v.JSValue())
}
func (tt *Texture) Encoding() TextureEncoding {
	return TextureEncoding(tt.Get("encoding"))
}
func (tt *Texture) SetEncoding(v TextureEncoding) {
	tt.Set("encoding", v)
}
func (tt *Texture) FlipY() bool {
	return tt.Get("flipY").Bool()
}
func (tt *Texture) SetFlipY(v bool) {
	tt.Set("flipY", v)
}
func (tt *Texture) Format() PixelFormat {
	return PixelFormat(tt.Get("format"))
}
func (tt *Texture) SetFormat(v PixelFormat) {
	tt.Set("format", v)
}
func (tt *Texture) GenerateMipmaps() bool {
	return tt.Get("generateMipmaps").Bool()
}
func (tt *Texture) SetGenerateMipmaps(v bool) {
	tt.Set("generateMipmaps", v)
}
func (tt *Texture) Id() int {
	return tt.Get("id").Int()
}
func (tt *Texture) SetId(v int) {
	tt.Set("id", v)
}
func (tt *Texture) Image() js.Value {
	return tt.Get("image")
}
func (tt *Texture) SetImage(v js.Value) {
	tt.Set("image", v)
}
func (tt *Texture) MagFilter() TextureFilter {
	return TextureFilter(tt.Get("magFilter"))
}
func (tt *Texture) SetMagFilter(v TextureFilter) {
	tt.Set("magFilter", v)
}
func (tt *Texture) Mapping() Mapping {
	return Mapping(tt.Get("mapping"))
}
func (tt *Texture) SetMapping(v Mapping) {
	tt.Set("mapping", v)
}
func (tt *Texture) MinFilter() TextureFilter {
	return TextureFilter(tt.Get("minFilter"))
}
func (tt *Texture) SetMinFilter(v TextureFilter) {
	tt.Set("minFilter", v)
}
func (tt *Texture) Mipmaps() js.Value {
	return tt.Get("mipmaps")
}
func (tt *Texture) SetMipmaps(v js.Value) {
	tt.Set("mipmaps", v)
}
func (tt *Texture) Name() string {
	return tt.Get("name").String()
}
func (tt *Texture) SetName(v string) {
	tt.Set("name", v)
}
func (tt *Texture) NeedsUpdate() bool {
	return tt.Get("needsUpdate").Bool()
}
func (tt *Texture) SetNeedsUpdate(v bool) {
	tt.Set("needsUpdate", v)
}
func (tt *Texture) Offset() *Vector2 {
	return &Vector2{Value: tt.Get("offset")}
}
func (tt *Texture) SetOffset(v *Vector2) {
	tt.Set("offset", v.JSValue())
}
func (tt *Texture) OnUpdate() js.Value {
	return tt.Get("onUpdate")
}
func (tt *Texture) SetOnUpdate(v js.Value) {
	tt.Set("onUpdate", v)
}
func (tt *Texture) PremultiplyAlpha() bool {
	return tt.Get("premultiplyAlpha").Bool()
}
func (tt *Texture) SetPremultiplyAlpha(v bool) {
	tt.Set("premultiplyAlpha", v)
}
func (tt *Texture) Repeat() *Vector2 {
	return &Vector2{Value: tt.Get("repeat")}
}
func (tt *Texture) SetRepeat(v *Vector2) {
	tt.Set("repeat", v.JSValue())
}
func (tt *Texture) Rotation() float64 {
	return tt.Get("rotation").Float()
}
func (tt *Texture) SetRotation(v float64) {
	tt.Set("rotation", v)
}
func (tt *Texture) SourceFile() string {
	return tt.Get("sourceFile").String()
}
func (tt *Texture) SetSourceFile(v string) {
	tt.Set("sourceFile", v)
}
func (tt *Texture) Type() TextureDataType {
	return TextureDataType(tt.Get("type"))
}
func (tt *Texture) SetType(v TextureDataType) {
	tt.Set("type", v)
}
func (tt *Texture) UnpackAlignment() float64 {
	return tt.Get("unpackAlignment").Float()
}
func (tt *Texture) SetUnpackAlignment(v float64) {
	tt.Set("unpackAlignment", v)
}
func (tt *Texture) Uuid() string {
	return tt.Get("uuid").String()
}
func (tt *Texture) SetUuid(v string) {
	tt.Set("uuid", v)
}
func (tt *Texture) Version() float64 {
	return tt.Get("version").Float()
}
func (tt *Texture) SetVersion(v float64) {
	tt.Set("version", v)
}
func (tt *Texture) WrapS() Wrapping {
	return Wrapping(tt.Get("wrapS"))
}
func (tt *Texture) SetWrapS(v Wrapping) {
	tt.Set("wrapS", v)
}
func (tt *Texture) WrapT() Wrapping {
	return Wrapping(tt.Get("wrapT"))
}
func (tt *Texture) SetWrapT(v Wrapping) {
	tt.Set("wrapT", v)
}
func (tt *Texture) DEFAULT_IMAGE() js.Value {
	return tt.Get("DEFAULT_IMAGE")
}
func (tt *Texture) SetDEFAULT_IMAGE(v js.Value) {
	tt.Set("DEFAULT_IMAGE", v)
}
func (tt *Texture) DEFAULT_MAPPING() js.Value {
	return tt.Get("DEFAULT_MAPPING")
}
func (tt *Texture) SetDEFAULT_MAPPING(v js.Value) {
	tt.Set("DEFAULT_MAPPING", v)
}
func (tt *Texture) AddEventListener(typ string, listener js.Value) {
	tt.Call("addEventListener", typ, listener)
}
func (tt *Texture) Clone() *Texture {
	return &Texture{Value: tt.Call("clone")}
}
func (tt *Texture) Copy(source *Texture) *Texture {
	return &Texture{Value: tt.Call("copy", source)}
}
func (tt *Texture) DispatchEvent(event js.Value) {
	tt.Call("dispatchEvent", event)
}
func (tt *Texture) Dispose() {
	tt.Call("dispose")
}
func (tt *Texture) HasEventListener(typ string, listener js.Value) bool {
	return tt.Call("hasEventListener", typ, listener).Bool()
}
func (tt *Texture) RemoveEventListener(typ string, listener js.Value) {
	tt.Call("removeEventListener", typ, listener)
}
func (tt *Texture) ToJSON(meta js.Value) js.Value {
	return tt.Call("toJSON", meta)
}
func (tt *Texture) TransformUv(uv Vector) {
	tt.Call("transformUv", uv)
}

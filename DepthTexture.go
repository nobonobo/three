// Code generated from "textures/DepthTexture.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// DepthTexture extend: [Texture]
type DepthTexture struct {
	js.Value
}

func NewDepthTexture(width float64, heighht float64, typ TextureDataType, mapping Mapping, wrapS Wrapping, wrapT Wrapping, magFilter TextureFilter, minFilter TextureFilter, anisotropy float64) *DepthTexture {
	return &DepthTexture{Value: get("DepthTexture").New(width, heighht, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy)}
}
func (dt *DepthTexture) JSValue() js.Value {
	return dt.Value
}
func (dt *DepthTexture) Anisotropy() float64 {
	return dt.Get("anisotropy").Float()
}
func (dt *DepthTexture) SetAnisotropy(v float64) {
	dt.Set("anisotropy", v)
}
func (dt *DepthTexture) Center() *Vector2 {
	return &Vector2{Value: dt.Get("center")}
}
func (dt *DepthTexture) SetCenter(v *Vector2) {
	dt.Set("center", v.JSValue())
}
func (dt *DepthTexture) Encoding() TextureEncoding {
	return TextureEncoding(dt.Get("encoding"))
}
func (dt *DepthTexture) SetEncoding(v TextureEncoding) {
	dt.Set("encoding", v)
}
func (dt *DepthTexture) FlipY() bool {
	return dt.Get("flipY").Bool()
}
func (dt *DepthTexture) SetFlipY(v bool) {
	dt.Set("flipY", v)
}
func (dt *DepthTexture) Format() PixelFormat {
	return PixelFormat(dt.Get("format"))
}
func (dt *DepthTexture) SetFormat(v PixelFormat) {
	dt.Set("format", v)
}
func (dt *DepthTexture) GenerateMipmaps() bool {
	return dt.Get("generateMipmaps").Bool()
}
func (dt *DepthTexture) SetGenerateMipmaps(v bool) {
	dt.Set("generateMipmaps", v)
}
func (dt *DepthTexture) Id() int {
	return dt.Get("id").Int()
}
func (dt *DepthTexture) SetId(v int) {
	dt.Set("id", v)
}
func (dt *DepthTexture) Image() js.Value {
	return dt.Get("image")
}
func (dt *DepthTexture) SetImage(v js.Value) {
	dt.Set("image", v)
}
func (dt *DepthTexture) MagFilter() TextureFilter {
	return TextureFilter(dt.Get("magFilter"))
}
func (dt *DepthTexture) SetMagFilter(v TextureFilter) {
	dt.Set("magFilter", v)
}
func (dt *DepthTexture) Mapping() Mapping {
	return Mapping(dt.Get("mapping"))
}
func (dt *DepthTexture) SetMapping(v Mapping) {
	dt.Set("mapping", v)
}
func (dt *DepthTexture) MinFilter() TextureFilter {
	return TextureFilter(dt.Get("minFilter"))
}
func (dt *DepthTexture) SetMinFilter(v TextureFilter) {
	dt.Set("minFilter", v)
}
func (dt *DepthTexture) Mipmaps() js.Value {
	return dt.Get("mipmaps")
}
func (dt *DepthTexture) SetMipmaps(v js.Value) {
	dt.Set("mipmaps", v)
}
func (dt *DepthTexture) Name() string {
	return dt.Get("name").String()
}
func (dt *DepthTexture) SetName(v string) {
	dt.Set("name", v)
}
func (dt *DepthTexture) NeedsUpdate() bool {
	return dt.Get("needsUpdate").Bool()
}
func (dt *DepthTexture) SetNeedsUpdate(v bool) {
	dt.Set("needsUpdate", v)
}
func (dt *DepthTexture) Offset() *Vector2 {
	return &Vector2{Value: dt.Get("offset")}
}
func (dt *DepthTexture) SetOffset(v *Vector2) {
	dt.Set("offset", v.JSValue())
}
func (dt *DepthTexture) OnUpdate() js.Value {
	return dt.Get("onUpdate")
}
func (dt *DepthTexture) SetOnUpdate(v js.Value) {
	dt.Set("onUpdate", v)
}
func (dt *DepthTexture) PremultiplyAlpha() bool {
	return dt.Get("premultiplyAlpha").Bool()
}
func (dt *DepthTexture) SetPremultiplyAlpha(v bool) {
	dt.Set("premultiplyAlpha", v)
}
func (dt *DepthTexture) Repeat() *Vector2 {
	return &Vector2{Value: dt.Get("repeat")}
}
func (dt *DepthTexture) SetRepeat(v *Vector2) {
	dt.Set("repeat", v.JSValue())
}
func (dt *DepthTexture) Rotation() float64 {
	return dt.Get("rotation").Float()
}
func (dt *DepthTexture) SetRotation(v float64) {
	dt.Set("rotation", v)
}
func (dt *DepthTexture) SourceFile() string {
	return dt.Get("sourceFile").String()
}
func (dt *DepthTexture) SetSourceFile(v string) {
	dt.Set("sourceFile", v)
}
func (dt *DepthTexture) Type() TextureDataType {
	return TextureDataType(dt.Get("type"))
}
func (dt *DepthTexture) SetType(v TextureDataType) {
	dt.Set("type", v)
}
func (dt *DepthTexture) UnpackAlignment() float64 {
	return dt.Get("unpackAlignment").Float()
}
func (dt *DepthTexture) SetUnpackAlignment(v float64) {
	dt.Set("unpackAlignment", v)
}
func (dt *DepthTexture) Uuid() string {
	return dt.Get("uuid").String()
}
func (dt *DepthTexture) SetUuid(v string) {
	dt.Set("uuid", v)
}
func (dt *DepthTexture) Version() float64 {
	return dt.Get("version").Float()
}
func (dt *DepthTexture) SetVersion(v float64) {
	dt.Set("version", v)
}
func (dt *DepthTexture) WrapS() Wrapping {
	return Wrapping(dt.Get("wrapS"))
}
func (dt *DepthTexture) SetWrapS(v Wrapping) {
	dt.Set("wrapS", v)
}
func (dt *DepthTexture) WrapT() Wrapping {
	return Wrapping(dt.Get("wrapT"))
}
func (dt *DepthTexture) SetWrapT(v Wrapping) {
	dt.Set("wrapT", v)
}
func (dt *DepthTexture) DEFAULT_IMAGE() js.Value {
	return dt.Get("DEFAULT_IMAGE")
}
func (dt *DepthTexture) SetDEFAULT_IMAGE(v js.Value) {
	dt.Set("DEFAULT_IMAGE", v)
}
func (dt *DepthTexture) DEFAULT_MAPPING() js.Value {
	return dt.Get("DEFAULT_MAPPING")
}
func (dt *DepthTexture) SetDEFAULT_MAPPING(v js.Value) {
	dt.Set("DEFAULT_MAPPING", v)
}
func (dt *DepthTexture) AddEventListener(typ string, listener js.Value) {
	dt.Call("addEventListener", typ, listener)
}
func (dt *DepthTexture) Clone() *DepthTexture {
	return &DepthTexture{Value: dt.Call("clone")}
}
func (dt *DepthTexture) Copy(source *Texture) *DepthTexture {
	return &DepthTexture{Value: dt.Call("copy", source)}
}
func (dt *DepthTexture) DispatchEvent(event js.Value) {
	dt.Call("dispatchEvent", event)
}
func (dt *DepthTexture) Dispose() {
	dt.Call("dispose")
}
func (dt *DepthTexture) HasEventListener(typ string, listener js.Value) bool {
	return dt.Call("hasEventListener", typ, listener).Bool()
}
func (dt *DepthTexture) RemoveEventListener(typ string, listener js.Value) {
	dt.Call("removeEventListener", typ, listener)
}
func (dt *DepthTexture) ToJSON(meta js.Value) js.Value {
	return dt.Call("toJSON", meta)
}
func (dt *DepthTexture) TransformUv(uv Vector) {
	dt.Call("transformUv", uv)
}

// Code generated from "textures/CubeTexture.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CubeTexture extend: [Texture]
type CubeTexture struct {
	js.Value
}

func NewCubeTexture(images js.Value, mapping Mapping, wrapS Wrapping, wrapT Wrapping, magFilter TextureFilter, minFilter TextureFilter, format PixelFormat, typ TextureDataType, anisotropy float64, encoding TextureEncoding) *CubeTexture {
	return &CubeTexture{Value: get("CubeTexture").New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy, encoding)}
}
func (ct *CubeTexture) JSValue() js.Value {
	return ct.Value
}
func (ct *CubeTexture) Anisotropy() float64 {
	return ct.Get("anisotropy").Float()
}
func (ct *CubeTexture) SetAnisotropy(v float64) {
	ct.Set("anisotropy", v)
}
func (ct *CubeTexture) Center() *Vector2 {
	return &Vector2{Value: ct.Get("center")}
}
func (ct *CubeTexture) SetCenter(v *Vector2) {
	ct.Set("center", v.JSValue())
}
func (ct *CubeTexture) Encoding() TextureEncoding {
	return TextureEncoding(ct.Get("encoding"))
}
func (ct *CubeTexture) SetEncoding(v TextureEncoding) {
	ct.Set("encoding", v)
}
func (ct *CubeTexture) FlipY() bool {
	return ct.Get("flipY").Bool()
}
func (ct *CubeTexture) SetFlipY(v bool) {
	ct.Set("flipY", v)
}
func (ct *CubeTexture) Format() PixelFormat {
	return PixelFormat(ct.Get("format"))
}
func (ct *CubeTexture) SetFormat(v PixelFormat) {
	ct.Set("format", v)
}
func (ct *CubeTexture) GenerateMipmaps() bool {
	return ct.Get("generateMipmaps").Bool()
}
func (ct *CubeTexture) SetGenerateMipmaps(v bool) {
	ct.Set("generateMipmaps", v)
}
func (ct *CubeTexture) Id() int {
	return ct.Get("id").Int()
}
func (ct *CubeTexture) SetId(v int) {
	ct.Set("id", v)
}
func (ct *CubeTexture) Image() js.Value {
	return ct.Get("image")
}
func (ct *CubeTexture) SetImage(v js.Value) {
	ct.Set("image", v)
}
func (ct *CubeTexture) Images() js.Value {
	return ct.Get("images")
}
func (ct *CubeTexture) SetImages(v js.Value) {
	ct.Set("images", v)
}
func (ct *CubeTexture) MagFilter() TextureFilter {
	return TextureFilter(ct.Get("magFilter"))
}
func (ct *CubeTexture) SetMagFilter(v TextureFilter) {
	ct.Set("magFilter", v)
}
func (ct *CubeTexture) Mapping() Mapping {
	return Mapping(ct.Get("mapping"))
}
func (ct *CubeTexture) SetMapping(v Mapping) {
	ct.Set("mapping", v)
}
func (ct *CubeTexture) MinFilter() TextureFilter {
	return TextureFilter(ct.Get("minFilter"))
}
func (ct *CubeTexture) SetMinFilter(v TextureFilter) {
	ct.Set("minFilter", v)
}
func (ct *CubeTexture) Mipmaps() js.Value {
	return ct.Get("mipmaps")
}
func (ct *CubeTexture) SetMipmaps(v js.Value) {
	ct.Set("mipmaps", v)
}
func (ct *CubeTexture) Name() string {
	return ct.Get("name").String()
}
func (ct *CubeTexture) SetName(v string) {
	ct.Set("name", v)
}
func (ct *CubeTexture) NeedsUpdate() bool {
	return ct.Get("needsUpdate").Bool()
}
func (ct *CubeTexture) SetNeedsUpdate(v bool) {
	ct.Set("needsUpdate", v)
}
func (ct *CubeTexture) Offset() *Vector2 {
	return &Vector2{Value: ct.Get("offset")}
}
func (ct *CubeTexture) SetOffset(v *Vector2) {
	ct.Set("offset", v.JSValue())
}
func (ct *CubeTexture) OnUpdate() js.Value {
	return ct.Get("onUpdate")
}
func (ct *CubeTexture) SetOnUpdate(v js.Value) {
	ct.Set("onUpdate", v)
}
func (ct *CubeTexture) PremultiplyAlpha() bool {
	return ct.Get("premultiplyAlpha").Bool()
}
func (ct *CubeTexture) SetPremultiplyAlpha(v bool) {
	ct.Set("premultiplyAlpha", v)
}
func (ct *CubeTexture) Repeat() *Vector2 {
	return &Vector2{Value: ct.Get("repeat")}
}
func (ct *CubeTexture) SetRepeat(v *Vector2) {
	ct.Set("repeat", v.JSValue())
}
func (ct *CubeTexture) Rotation() float64 {
	return ct.Get("rotation").Float()
}
func (ct *CubeTexture) SetRotation(v float64) {
	ct.Set("rotation", v)
}
func (ct *CubeTexture) SourceFile() string {
	return ct.Get("sourceFile").String()
}
func (ct *CubeTexture) SetSourceFile(v string) {
	ct.Set("sourceFile", v)
}
func (ct *CubeTexture) Type() TextureDataType {
	return TextureDataType(ct.Get("type"))
}
func (ct *CubeTexture) SetType(v TextureDataType) {
	ct.Set("type", v)
}
func (ct *CubeTexture) UnpackAlignment() float64 {
	return ct.Get("unpackAlignment").Float()
}
func (ct *CubeTexture) SetUnpackAlignment(v float64) {
	ct.Set("unpackAlignment", v)
}
func (ct *CubeTexture) Uuid() string {
	return ct.Get("uuid").String()
}
func (ct *CubeTexture) SetUuid(v string) {
	ct.Set("uuid", v)
}
func (ct *CubeTexture) Version() float64 {
	return ct.Get("version").Float()
}
func (ct *CubeTexture) SetVersion(v float64) {
	ct.Set("version", v)
}
func (ct *CubeTexture) WrapS() Wrapping {
	return Wrapping(ct.Get("wrapS"))
}
func (ct *CubeTexture) SetWrapS(v Wrapping) {
	ct.Set("wrapS", v)
}
func (ct *CubeTexture) WrapT() Wrapping {
	return Wrapping(ct.Get("wrapT"))
}
func (ct *CubeTexture) SetWrapT(v Wrapping) {
	ct.Set("wrapT", v)
}
func (ct *CubeTexture) DEFAULT_IMAGE() js.Value {
	return ct.Get("DEFAULT_IMAGE")
}
func (ct *CubeTexture) SetDEFAULT_IMAGE(v js.Value) {
	ct.Set("DEFAULT_IMAGE", v)
}
func (ct *CubeTexture) DEFAULT_MAPPING() js.Value {
	return ct.Get("DEFAULT_MAPPING")
}
func (ct *CubeTexture) SetDEFAULT_MAPPING(v js.Value) {
	ct.Set("DEFAULT_MAPPING", v)
}
func (ct *CubeTexture) AddEventListener(typ string, listener js.Value) {
	ct.Call("addEventListener", typ, listener)
}
func (ct *CubeTexture) Clone() *CubeTexture {
	return &CubeTexture{Value: ct.Call("clone")}
}
func (ct *CubeTexture) Copy(source *Texture) *CubeTexture {
	return &CubeTexture{Value: ct.Call("copy", source)}
}
func (ct *CubeTexture) DispatchEvent(event js.Value) {
	ct.Call("dispatchEvent", event)
}
func (ct *CubeTexture) Dispose() {
	ct.Call("dispose")
}
func (ct *CubeTexture) HasEventListener(typ string, listener js.Value) bool {
	return ct.Call("hasEventListener", typ, listener).Bool()
}
func (ct *CubeTexture) RemoveEventListener(typ string, listener js.Value) {
	ct.Call("removeEventListener", typ, listener)
}
func (ct *CubeTexture) ToJSON(meta js.Value) js.Value {
	return ct.Call("toJSON", meta)
}
func (ct *CubeTexture) TransformUv(uv Vector) {
	ct.Call("transformUv", uv)
}

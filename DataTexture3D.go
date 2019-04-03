// Code generated from "textures/DataTexture3D.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type DataTexture3D struct {
	js.Value
}

func NewDataTexture3D(data js.Value, width float64, height float64, depth float64) *DataTexture3D {
	return &DataTexture3D{Value: get("DataTexture3D").New(data, width, height, depth)}
}
func (dtd *DataTexture3D) Anisotropy() float64 {
	return dtd.Get("anisotropy").Float()
}
func (dtd *DataTexture3D) SetAnisotropy(v float64) {
	dtd.Set("anisotropy", v)
}
func (dtd *DataTexture3D) Center() *Vector2 {
	return &Vector2{Value: dtd.Get("center")}
}
func (dtd *DataTexture3D) SetCenter(v *Vector2) {
	dtd.Set("center", v)
}
func (dtd *DataTexture3D) Encoding() TextureEncoding {
	return TextureEncoding(dtd.Get("encoding"))
}
func (dtd *DataTexture3D) SetEncoding(v TextureEncoding) {
	dtd.Set("encoding", v)
}
func (dtd *DataTexture3D) FlipY() bool {
	return dtd.Get("flipY").Bool()
}
func (dtd *DataTexture3D) SetFlipY(v bool) {
	dtd.Set("flipY", v)
}
func (dtd *DataTexture3D) Format() PixelFormat {
	return PixelFormat(dtd.Get("format"))
}
func (dtd *DataTexture3D) SetFormat(v PixelFormat) {
	dtd.Set("format", v)
}
func (dtd *DataTexture3D) GenerateMipmaps() bool {
	return dtd.Get("generateMipmaps").Bool()
}
func (dtd *DataTexture3D) SetGenerateMipmaps(v bool) {
	dtd.Set("generateMipmaps", v)
}
func (dtd *DataTexture3D) Id() int {
	return dtd.Get("id").Int()
}
func (dtd *DataTexture3D) SetId(v int) {
	dtd.Set("id", v)
}
func (dtd *DataTexture3D) Image() js.Value {
	return dtd.Get("image")
}
func (dtd *DataTexture3D) SetImage(v js.Value) {
	dtd.Set("image", v)
}
func (dtd *DataTexture3D) MagFilter() TextureFilter {
	return TextureFilter(dtd.Get("magFilter"))
}
func (dtd *DataTexture3D) SetMagFilter(v TextureFilter) {
	dtd.Set("magFilter", v)
}
func (dtd *DataTexture3D) Mapping() Mapping {
	return Mapping(dtd.Get("mapping"))
}
func (dtd *DataTexture3D) SetMapping(v Mapping) {
	dtd.Set("mapping", v)
}
func (dtd *DataTexture3D) MinFilter() TextureFilter {
	return TextureFilter(dtd.Get("minFilter"))
}
func (dtd *DataTexture3D) SetMinFilter(v TextureFilter) {
	dtd.Set("minFilter", v)
}
func (dtd *DataTexture3D) Mipmaps() js.Value {
	return dtd.Get("mipmaps")
}
func (dtd *DataTexture3D) SetMipmaps(v js.Value) {
	dtd.Set("mipmaps", v)
}
func (dtd *DataTexture3D) Name() string {
	return dtd.Get("name").String()
}
func (dtd *DataTexture3D) SetName(v string) {
	dtd.Set("name", v)
}
func (dtd *DataTexture3D) NeedsUpdate() bool {
	return dtd.Get("needsUpdate").Bool()
}
func (dtd *DataTexture3D) SetNeedsUpdate(v bool) {
	dtd.Set("needsUpdate", v)
}
func (dtd *DataTexture3D) Offset() *Vector2 {
	return &Vector2{Value: dtd.Get("offset")}
}
func (dtd *DataTexture3D) SetOffset(v *Vector2) {
	dtd.Set("offset", v)
}
func (dtd *DataTexture3D) OnUpdate() js.Value {
	return dtd.Get("onUpdate")
}
func (dtd *DataTexture3D) SetOnUpdate(v js.Value) {
	dtd.Set("onUpdate", v)
}
func (dtd *DataTexture3D) PremultiplyAlpha() bool {
	return dtd.Get("premultiplyAlpha").Bool()
}
func (dtd *DataTexture3D) SetPremultiplyAlpha(v bool) {
	dtd.Set("premultiplyAlpha", v)
}
func (dtd *DataTexture3D) Repeat() *Vector2 {
	return &Vector2{Value: dtd.Get("repeat")}
}
func (dtd *DataTexture3D) SetRepeat(v *Vector2) {
	dtd.Set("repeat", v)
}
func (dtd *DataTexture3D) Rotation() float64 {
	return dtd.Get("rotation").Float()
}
func (dtd *DataTexture3D) SetRotation(v float64) {
	dtd.Set("rotation", v)
}
func (dtd *DataTexture3D) SourceFile() string {
	return dtd.Get("sourceFile").String()
}
func (dtd *DataTexture3D) SetSourceFile(v string) {
	dtd.Set("sourceFile", v)
}
func (dtd *DataTexture3D) Type() TextureDataType {
	return TextureDataType(dtd.Get("type"))
}
func (dtd *DataTexture3D) SetType(v TextureDataType) {
	dtd.Set("type", v)
}
func (dtd *DataTexture3D) UnpackAlignment() float64 {
	return dtd.Get("unpackAlignment").Float()
}
func (dtd *DataTexture3D) SetUnpackAlignment(v float64) {
	dtd.Set("unpackAlignment", v)
}
func (dtd *DataTexture3D) Uuid() string {
	return dtd.Get("uuid").String()
}
func (dtd *DataTexture3D) SetUuid(v string) {
	dtd.Set("uuid", v)
}
func (dtd *DataTexture3D) Version() float64 {
	return dtd.Get("version").Float()
}
func (dtd *DataTexture3D) SetVersion(v float64) {
	dtd.Set("version", v)
}
func (dtd *DataTexture3D) WrapS() Wrapping {
	return Wrapping(dtd.Get("wrapS"))
}
func (dtd *DataTexture3D) SetWrapS(v Wrapping) {
	dtd.Set("wrapS", v)
}
func (dtd *DataTexture3D) WrapT() Wrapping {
	return Wrapping(dtd.Get("wrapT"))
}
func (dtd *DataTexture3D) SetWrapT(v Wrapping) {
	dtd.Set("wrapT", v)
}
func (dtd *DataTexture3D) DEFAULT_IMAGE() js.Value {
	return dtd.Get("DEFAULT_IMAGE")
}
func (dtd *DataTexture3D) SetDEFAULT_IMAGE(v js.Value) {
	dtd.Set("DEFAULT_IMAGE", v)
}
func (dtd *DataTexture3D) DEFAULT_MAPPING() js.Value {
	return dtd.Get("DEFAULT_MAPPING")
}
func (dtd *DataTexture3D) SetDEFAULT_MAPPING(v js.Value) {
	dtd.Set("DEFAULT_MAPPING", v)
}
func (dtd *DataTexture3D) AddEventListener(typ string, listener js.Value) {
	dtd.Call("addEventListener", typ, listener)
}
func (dtd *DataTexture3D) Clone() *DataTexture3D {
	return &DataTexture3D{Value: dtd.Call("clone")}
}
func (dtd *DataTexture3D) Copy(source *Texture) *DataTexture3D {
	return &DataTexture3D{Value: dtd.Call("copy", source)}
}
func (dtd *DataTexture3D) DispatchEvent(event js.Value) {
	dtd.Call("dispatchEvent", event)
}
func (dtd *DataTexture3D) Dispose() {
	dtd.Call("dispose")
}
func (dtd *DataTexture3D) HasEventListener(typ string, listener js.Value) bool {
	return dtd.Call("hasEventListener", typ, listener).Bool()
}
func (dtd *DataTexture3D) RemoveEventListener(typ string, listener js.Value) {
	dtd.Call("removeEventListener", typ, listener)
}
func (dtd *DataTexture3D) ToJSON(meta js.Value) js.Value {
	return dtd.Call("toJSON", meta)
}
func (dtd *DataTexture3D) TransformUv(uv Vector) {
	dtd.Call("transformUv", uv)
}

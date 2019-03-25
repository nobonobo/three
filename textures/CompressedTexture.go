package textures

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type CompressedTexture struct {
	js.Value
}

func NewCompressedTexture(mipmaps []ImageData, width int, height int, format *PixelFormat, typ *TextureDataType, mapping *Mapping, wrapS *Wrapping, wrapT *Wrapping, magFilter *TextureFilter, minFilter *TextureFilter, anisotropy int, encoding *TextureEncoding) *CompressedTexture {
	return &CompressedTexture{Value: get("CompressedTexture").New(mipmaps, width, height, format, typ, mapping, wrapS, wrapT, magFilter, minFilter, anisotropy, encoding)}
}
func (ct *CompressedTexture) Anisotropy() int {
	return ct.Get("anisotropy").Int()
}
func (ct *CompressedTexture) SetAnisotropy(v int) {
	ct.Set("anisotropy", v)
}
func (ct *CompressedTexture) Center() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("center")}
}
func (ct *CompressedTexture) SetCenter(v *math.Vector2) {
	ct.Set("center", v)
}
func (ct *CompressedTexture) Encoding() *TextureEncoding {
	return &TextureEncoding{Value: ct.Get("encoding")}
}
func (ct *CompressedTexture) SetEncoding(v *TextureEncoding) {
	ct.Set("encoding", v)
}
func (ct *CompressedTexture) FlipY() bool {
	return ct.Get("flipY").Bool()
}
func (ct *CompressedTexture) SetFlipY(v bool) {
	ct.Set("flipY", v)
}
func (ct *CompressedTexture) Format() *PixelFormat {
	return &PixelFormat{Value: ct.Get("format")}
}
func (ct *CompressedTexture) SetFormat(v *PixelFormat) {
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
func (ct *CompressedTexture) MagFilter() *TextureFilter {
	return &TextureFilter{Value: ct.Get("magFilter")}
}
func (ct *CompressedTexture) SetMagFilter(v *TextureFilter) {
	ct.Set("magFilter", v)
}
func (ct *CompressedTexture) Mapping() *Mapping {
	return &Mapping{Value: ct.Get("mapping")}
}
func (ct *CompressedTexture) SetMapping(v *Mapping) {
	ct.Set("mapping", v)
}
func (ct *CompressedTexture) MinFilter() *TextureFilter {
	return &TextureFilter{Value: ct.Get("minFilter")}
}
func (ct *CompressedTexture) SetMinFilter(v *TextureFilter) {
	ct.Set("minFilter", v)
}
func (ct *CompressedTexture) Mipmaps() []ImageData {
	return []ImageData(ct.Get("mipmaps"))
}
func (ct *CompressedTexture) SetMipmaps(v []ImageData) {
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
func (ct *CompressedTexture) Offset() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("offset")}
}
func (ct *CompressedTexture) SetOffset(v *math.Vector2) {
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
func (ct *CompressedTexture) Repeat() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("repeat")}
}
func (ct *CompressedTexture) SetRepeat(v *math.Vector2) {
	ct.Set("repeat", v)
}
func (ct *CompressedTexture) Rotation() int {
	return ct.Get("rotation").Int()
}
func (ct *CompressedTexture) SetRotation(v int) {
	ct.Set("rotation", v)
}
func (ct *CompressedTexture) SourceFile() string {
	return ct.Get("sourceFile").String()
}
func (ct *CompressedTexture) SetSourceFile(v string) {
	ct.Set("sourceFile", v)
}
func (ct *CompressedTexture) Type() *TextureDataType {
	return &TextureDataType{Value: ct.Get("type")}
}
func (ct *CompressedTexture) SetType(v *TextureDataType) {
	ct.Set("type", v)
}
func (ct *CompressedTexture) UnpackAlignment() int {
	return ct.Get("unpackAlignment").Int()
}
func (ct *CompressedTexture) SetUnpackAlignment(v int) {
	ct.Set("unpackAlignment", v)
}
func (ct *CompressedTexture) Uuid() string {
	return ct.Get("uuid").String()
}
func (ct *CompressedTexture) SetUuid(v string) {
	ct.Set("uuid", v)
}
func (ct *CompressedTexture) Version() int {
	return ct.Get("version").Int()
}
func (ct *CompressedTexture) SetVersion(v int) {
	ct.Set("version", v)
}
func (ct *CompressedTexture) WrapS() *Wrapping {
	return &Wrapping{Value: ct.Get("wrapS")}
}
func (ct *CompressedTexture) SetWrapS(v *Wrapping) {
	ct.Set("wrapS", v)
}
func (ct *CompressedTexture) WrapT() *Wrapping {
	return &Wrapping{Value: ct.Get("wrapT")}
}
func (ct *CompressedTexture) SetWrapT(v *Wrapping) {
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
func (ct *CompressedTexture) Clone() this {
	return this(ct.Call("clone"))
}
func (ct *CompressedTexture) Copy(source *Texture) this {
	return this(ct.Call("copy", source))
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
func (ct *CompressedTexture) TransformUv(uv math.Vector) {
	ct.Call("transformUv", uv)
}

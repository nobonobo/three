package textures

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type CanvasTexture struct {
	js.Value
}

func NewCanvasTexture(canvas HTMLImageElement, mapping *Mapping, wrapS *Wrapping, wrapT *Wrapping, magFilter *TextureFilter, minFilter *TextureFilter, format *PixelFormat, typ *TextureDataType, anisotropy int) *CanvasTexture {
	return &CanvasTexture{Value: get("CanvasTexture").New(canvas, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)}
}
func (ct *CanvasTexture) Anisotropy() int {
	return ct.Get("anisotropy").Int()
}
func (ct *CanvasTexture) SetAnisotropy(v int) {
	ct.Set("anisotropy", v)
}
func (ct *CanvasTexture) Center() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("center")}
}
func (ct *CanvasTexture) SetCenter(v *math.Vector2) {
	ct.Set("center", v)
}
func (ct *CanvasTexture) Encoding() *TextureEncoding {
	return &TextureEncoding{Value: ct.Get("encoding")}
}
func (ct *CanvasTexture) SetEncoding(v *TextureEncoding) {
	ct.Set("encoding", v)
}
func (ct *CanvasTexture) FlipY() bool {
	return ct.Get("flipY").Bool()
}
func (ct *CanvasTexture) SetFlipY(v bool) {
	ct.Set("flipY", v)
}
func (ct *CanvasTexture) Format() *PixelFormat {
	return &PixelFormat{Value: ct.Get("format")}
}
func (ct *CanvasTexture) SetFormat(v *PixelFormat) {
	ct.Set("format", v)
}
func (ct *CanvasTexture) GenerateMipmaps() bool {
	return ct.Get("generateMipmaps").Bool()
}
func (ct *CanvasTexture) SetGenerateMipmaps(v bool) {
	ct.Set("generateMipmaps", v)
}
func (ct *CanvasTexture) Id() int {
	return ct.Get("id").Int()
}
func (ct *CanvasTexture) SetId(v int) {
	ct.Set("id", v)
}
func (ct *CanvasTexture) Image() js.Value {
	return ct.Get("image")
}
func (ct *CanvasTexture) SetImage(v js.Value) {
	ct.Set("image", v)
}
func (ct *CanvasTexture) MagFilter() *TextureFilter {
	return &TextureFilter{Value: ct.Get("magFilter")}
}
func (ct *CanvasTexture) SetMagFilter(v *TextureFilter) {
	ct.Set("magFilter", v)
}
func (ct *CanvasTexture) Mapping() *Mapping {
	return &Mapping{Value: ct.Get("mapping")}
}
func (ct *CanvasTexture) SetMapping(v *Mapping) {
	ct.Set("mapping", v)
}
func (ct *CanvasTexture) MinFilter() *TextureFilter {
	return &TextureFilter{Value: ct.Get("minFilter")}
}
func (ct *CanvasTexture) SetMinFilter(v *TextureFilter) {
	ct.Set("minFilter", v)
}
func (ct *CanvasTexture) Mipmaps() []ImageData {
	return []ImageData(ct.Get("mipmaps"))
}
func (ct *CanvasTexture) SetMipmaps(v []ImageData) {
	ct.Set("mipmaps", v)
}
func (ct *CanvasTexture) Name() string {
	return ct.Get("name").String()
}
func (ct *CanvasTexture) SetName(v string) {
	ct.Set("name", v)
}
func (ct *CanvasTexture) NeedsUpdate() bool {
	return ct.Get("needsUpdate").Bool()
}
func (ct *CanvasTexture) SetNeedsUpdate(v bool) {
	ct.Set("needsUpdate", v)
}
func (ct *CanvasTexture) Offset() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("offset")}
}
func (ct *CanvasTexture) SetOffset(v *math.Vector2) {
	ct.Set("offset", v)
}
func (ct *CanvasTexture) OnUpdate() js.Value {
	return ct.Get("onUpdate")
}
func (ct *CanvasTexture) SetOnUpdate(v js.Value) {
	ct.Set("onUpdate", v)
}
func (ct *CanvasTexture) PremultiplyAlpha() bool {
	return ct.Get("premultiplyAlpha").Bool()
}
func (ct *CanvasTexture) SetPremultiplyAlpha(v bool) {
	ct.Set("premultiplyAlpha", v)
}
func (ct *CanvasTexture) Repeat() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("repeat")}
}
func (ct *CanvasTexture) SetRepeat(v *math.Vector2) {
	ct.Set("repeat", v)
}
func (ct *CanvasTexture) Rotation() int {
	return ct.Get("rotation").Int()
}
func (ct *CanvasTexture) SetRotation(v int) {
	ct.Set("rotation", v)
}
func (ct *CanvasTexture) SourceFile() string {
	return ct.Get("sourceFile").String()
}
func (ct *CanvasTexture) SetSourceFile(v string) {
	ct.Set("sourceFile", v)
}
func (ct *CanvasTexture) Type() *TextureDataType {
	return &TextureDataType{Value: ct.Get("type")}
}
func (ct *CanvasTexture) SetType(v *TextureDataType) {
	ct.Set("type", v)
}
func (ct *CanvasTexture) UnpackAlignment() int {
	return ct.Get("unpackAlignment").Int()
}
func (ct *CanvasTexture) SetUnpackAlignment(v int) {
	ct.Set("unpackAlignment", v)
}
func (ct *CanvasTexture) Uuid() string {
	return ct.Get("uuid").String()
}
func (ct *CanvasTexture) SetUuid(v string) {
	ct.Set("uuid", v)
}
func (ct *CanvasTexture) Version() int {
	return ct.Get("version").Int()
}
func (ct *CanvasTexture) SetVersion(v int) {
	ct.Set("version", v)
}
func (ct *CanvasTexture) WrapS() *Wrapping {
	return &Wrapping{Value: ct.Get("wrapS")}
}
func (ct *CanvasTexture) SetWrapS(v *Wrapping) {
	ct.Set("wrapS", v)
}
func (ct *CanvasTexture) WrapT() *Wrapping {
	return &Wrapping{Value: ct.Get("wrapT")}
}
func (ct *CanvasTexture) SetWrapT(v *Wrapping) {
	ct.Set("wrapT", v)
}
func (ct *CanvasTexture) DEFAULT_IMAGE() js.Value {
	return ct.Get("DEFAULT_IMAGE")
}
func (ct *CanvasTexture) SetDEFAULT_IMAGE(v js.Value) {
	ct.Set("DEFAULT_IMAGE", v)
}
func (ct *CanvasTexture) DEFAULT_MAPPING() js.Value {
	return ct.Get("DEFAULT_MAPPING")
}
func (ct *CanvasTexture) SetDEFAULT_MAPPING(v js.Value) {
	ct.Set("DEFAULT_MAPPING", v)
}
func (ct *CanvasTexture) AddEventListener(typ string, listener js.Value) {
	ct.Call("addEventListener", typ, listener)
}
func (ct *CanvasTexture) Clone() this {
	return this(ct.Call("clone"))
}
func (ct *CanvasTexture) Copy(source *Texture) this {
	return this(ct.Call("copy", source))
}
func (ct *CanvasTexture) DispatchEvent(event js.Value) {
	ct.Call("dispatchEvent", event)
}
func (ct *CanvasTexture) Dispose() {
	ct.Call("dispose")
}
func (ct *CanvasTexture) HasEventListener(typ string, listener js.Value) bool {
	return ct.Call("hasEventListener", typ, listener).Bool()
}
func (ct *CanvasTexture) RemoveEventListener(typ string, listener js.Value) {
	ct.Call("removeEventListener", typ, listener)
}
func (ct *CanvasTexture) ToJSON(meta js.Value) js.Value {
	return ct.Call("toJSON", meta)
}
func (ct *CanvasTexture) TransformUv(uv math.Vector) {
	ct.Call("transformUv", uv)
}

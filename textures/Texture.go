package textures

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

func TextureIdCount() int {
	return get("TextureIdCount").Int()
}
func SetTextureIdCount(v int) {
	set("TextureIdCount", v)
}

type Texture struct {
	js.Value
}

func NewTexture(image HTMLImageElement, mapping *Mapping, wrapS *Wrapping, wrapT *Wrapping, magFilter *TextureFilter, minFilter *TextureFilter, format *PixelFormat, typ *TextureDataType, anisotropy int, encoding *TextureEncoding) *Texture {
	return &Texture{Value: get("Texture").New(image, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy, encoding)}
}
func (t *Texture) Anisotropy() int {
	return t.Get("anisotropy").Int()
}
func (t *Texture) SetAnisotropy(v int) {
	t.Set("anisotropy", v)
}
func (t *Texture) Center() *math.Vector2 {
	return &math.Vector2{Value: t.Get("center")}
}
func (t *Texture) SetCenter(v *math.Vector2) {
	t.Set("center", v)
}
func (t *Texture) Encoding() *TextureEncoding {
	return &TextureEncoding{Value: t.Get("encoding")}
}
func (t *Texture) SetEncoding(v *TextureEncoding) {
	t.Set("encoding", v)
}
func (t *Texture) FlipY() bool {
	return t.Get("flipY").Bool()
}
func (t *Texture) SetFlipY(v bool) {
	t.Set("flipY", v)
}
func (t *Texture) Format() *PixelFormat {
	return &PixelFormat{Value: t.Get("format")}
}
func (t *Texture) SetFormat(v *PixelFormat) {
	t.Set("format", v)
}
func (t *Texture) GenerateMipmaps() bool {
	return t.Get("generateMipmaps").Bool()
}
func (t *Texture) SetGenerateMipmaps(v bool) {
	t.Set("generateMipmaps", v)
}
func (t *Texture) Id() int {
	return t.Get("id").Int()
}
func (t *Texture) SetId(v int) {
	t.Set("id", v)
}
func (t *Texture) Image() js.Value {
	return t.Get("image")
}
func (t *Texture) SetImage(v js.Value) {
	t.Set("image", v)
}
func (t *Texture) MagFilter() *TextureFilter {
	return &TextureFilter{Value: t.Get("magFilter")}
}
func (t *Texture) SetMagFilter(v *TextureFilter) {
	t.Set("magFilter", v)
}
func (t *Texture) Mapping() *Mapping {
	return &Mapping{Value: t.Get("mapping")}
}
func (t *Texture) SetMapping(v *Mapping) {
	t.Set("mapping", v)
}
func (t *Texture) MinFilter() *TextureFilter {
	return &TextureFilter{Value: t.Get("minFilter")}
}
func (t *Texture) SetMinFilter(v *TextureFilter) {
	t.Set("minFilter", v)
}
func (t *Texture) Mipmaps() []ImageData {
	return []ImageData(t.Get("mipmaps"))
}
func (t *Texture) SetMipmaps(v []ImageData) {
	t.Set("mipmaps", v)
}
func (t *Texture) Name() string {
	return t.Get("name").String()
}
func (t *Texture) SetName(v string) {
	t.Set("name", v)
}
func (t *Texture) NeedsUpdate() bool {
	return t.Get("needsUpdate").Bool()
}
func (t *Texture) SetNeedsUpdate(v bool) {
	t.Set("needsUpdate", v)
}
func (t *Texture) Offset() *math.Vector2 {
	return &math.Vector2{Value: t.Get("offset")}
}
func (t *Texture) SetOffset(v *math.Vector2) {
	t.Set("offset", v)
}
func (t *Texture) OnUpdate() js.Value {
	return t.Get("onUpdate")
}
func (t *Texture) SetOnUpdate(v js.Value) {
	t.Set("onUpdate", v)
}
func (t *Texture) PremultiplyAlpha() bool {
	return t.Get("premultiplyAlpha").Bool()
}
func (t *Texture) SetPremultiplyAlpha(v bool) {
	t.Set("premultiplyAlpha", v)
}
func (t *Texture) Repeat() *math.Vector2 {
	return &math.Vector2{Value: t.Get("repeat")}
}
func (t *Texture) SetRepeat(v *math.Vector2) {
	t.Set("repeat", v)
}
func (t *Texture) Rotation() int {
	return t.Get("rotation").Int()
}
func (t *Texture) SetRotation(v int) {
	t.Set("rotation", v)
}
func (t *Texture) SourceFile() string {
	return t.Get("sourceFile").String()
}
func (t *Texture) SetSourceFile(v string) {
	t.Set("sourceFile", v)
}
func (t *Texture) Type() *TextureDataType {
	return &TextureDataType{Value: t.Get("type")}
}
func (t *Texture) SetType(v *TextureDataType) {
	t.Set("type", v)
}
func (t *Texture) UnpackAlignment() int {
	return t.Get("unpackAlignment").Int()
}
func (t *Texture) SetUnpackAlignment(v int) {
	t.Set("unpackAlignment", v)
}
func (t *Texture) Uuid() string {
	return t.Get("uuid").String()
}
func (t *Texture) SetUuid(v string) {
	t.Set("uuid", v)
}
func (t *Texture) Version() int {
	return t.Get("version").Int()
}
func (t *Texture) SetVersion(v int) {
	t.Set("version", v)
}
func (t *Texture) WrapS() *Wrapping {
	return &Wrapping{Value: t.Get("wrapS")}
}
func (t *Texture) SetWrapS(v *Wrapping) {
	t.Set("wrapS", v)
}
func (t *Texture) WrapT() *Wrapping {
	return &Wrapping{Value: t.Get("wrapT")}
}
func (t *Texture) SetWrapT(v *Wrapping) {
	t.Set("wrapT", v)
}
func (t *Texture) DEFAULT_IMAGE() js.Value {
	return t.Get("DEFAULT_IMAGE")
}
func (t *Texture) SetDEFAULT_IMAGE(v js.Value) {
	t.Set("DEFAULT_IMAGE", v)
}
func (t *Texture) DEFAULT_MAPPING() js.Value {
	return t.Get("DEFAULT_MAPPING")
}
func (t *Texture) SetDEFAULT_MAPPING(v js.Value) {
	t.Set("DEFAULT_MAPPING", v)
}
func (t *Texture) AddEventListener(typ string, listener js.Value) {
	t.Call("addEventListener", typ, listener)
}
func (t *Texture) Clone() this {
	return this(t.Call("clone"))
}
func (t *Texture) Copy(source *Texture) this {
	return this(t.Call("copy", source))
}
func (t *Texture) DispatchEvent(event js.Value) {
	t.Call("dispatchEvent", event)
}
func (t *Texture) Dispose() {
	t.Call("dispose")
}
func (t *Texture) HasEventListener(typ string, listener js.Value) bool {
	return t.Call("hasEventListener", typ, listener).Bool()
}
func (t *Texture) RemoveEventListener(typ string, listener js.Value) {
	t.Call("removeEventListener", typ, listener)
}
func (t *Texture) ToJSON(meta js.Value) js.Value {
	return t.Call("toJSON", meta)
}
func (t *Texture) TransformUv(uv math.Vector) {
	t.Call("transformUv", uv)
}

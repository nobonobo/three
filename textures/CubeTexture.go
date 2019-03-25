package textures

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type CubeTexture struct {
	js.Value
}

func NewCubeTexture(images []js.Value, mapping *Mapping, wrapS *Wrapping, wrapT *Wrapping, magFilter *TextureFilter, minFilter *TextureFilter, format *PixelFormat, typ *TextureDataType, anisotropy int, encoding *TextureEncoding) *CubeTexture {
	return &CubeTexture{Value: get("CubeTexture").New(images, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy, encoding)}
}
func (ct *CubeTexture) Anisotropy() int {
	return ct.Get("anisotropy").Int()
}
func (ct *CubeTexture) SetAnisotropy(v int) {
	ct.Set("anisotropy", v)
}
func (ct *CubeTexture) Center() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("center")}
}
func (ct *CubeTexture) SetCenter(v *math.Vector2) {
	ct.Set("center", v)
}
func (ct *CubeTexture) Encoding() *TextureEncoding {
	return &TextureEncoding{Value: ct.Get("encoding")}
}
func (ct *CubeTexture) SetEncoding(v *TextureEncoding) {
	ct.Set("encoding", v)
}
func (ct *CubeTexture) FlipY() bool {
	return ct.Get("flipY").Bool()
}
func (ct *CubeTexture) SetFlipY(v bool) {
	ct.Set("flipY", v)
}
func (ct *CubeTexture) Format() *PixelFormat {
	return &PixelFormat{Value: ct.Get("format")}
}
func (ct *CubeTexture) SetFormat(v *PixelFormat) {
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
func (ct *CubeTexture) MagFilter() *TextureFilter {
	return &TextureFilter{Value: ct.Get("magFilter")}
}
func (ct *CubeTexture) SetMagFilter(v *TextureFilter) {
	ct.Set("magFilter", v)
}
func (ct *CubeTexture) Mapping() *Mapping {
	return &Mapping{Value: ct.Get("mapping")}
}
func (ct *CubeTexture) SetMapping(v *Mapping) {
	ct.Set("mapping", v)
}
func (ct *CubeTexture) MinFilter() *TextureFilter {
	return &TextureFilter{Value: ct.Get("minFilter")}
}
func (ct *CubeTexture) SetMinFilter(v *TextureFilter) {
	ct.Set("minFilter", v)
}
func (ct *CubeTexture) Mipmaps() []ImageData {
	return []ImageData(ct.Get("mipmaps"))
}
func (ct *CubeTexture) SetMipmaps(v []ImageData) {
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
func (ct *CubeTexture) Offset() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("offset")}
}
func (ct *CubeTexture) SetOffset(v *math.Vector2) {
	ct.Set("offset", v)
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
func (ct *CubeTexture) Repeat() *math.Vector2 {
	return &math.Vector2{Value: ct.Get("repeat")}
}
func (ct *CubeTexture) SetRepeat(v *math.Vector2) {
	ct.Set("repeat", v)
}
func (ct *CubeTexture) Rotation() int {
	return ct.Get("rotation").Int()
}
func (ct *CubeTexture) SetRotation(v int) {
	ct.Set("rotation", v)
}
func (ct *CubeTexture) SourceFile() string {
	return ct.Get("sourceFile").String()
}
func (ct *CubeTexture) SetSourceFile(v string) {
	ct.Set("sourceFile", v)
}
func (ct *CubeTexture) Type() *TextureDataType {
	return &TextureDataType{Value: ct.Get("type")}
}
func (ct *CubeTexture) SetType(v *TextureDataType) {
	ct.Set("type", v)
}
func (ct *CubeTexture) UnpackAlignment() int {
	return ct.Get("unpackAlignment").Int()
}
func (ct *CubeTexture) SetUnpackAlignment(v int) {
	ct.Set("unpackAlignment", v)
}
func (ct *CubeTexture) Uuid() string {
	return ct.Get("uuid").String()
}
func (ct *CubeTexture) SetUuid(v string) {
	ct.Set("uuid", v)
}
func (ct *CubeTexture) Version() int {
	return ct.Get("version").Int()
}
func (ct *CubeTexture) SetVersion(v int) {
	ct.Set("version", v)
}
func (ct *CubeTexture) WrapS() *Wrapping {
	return &Wrapping{Value: ct.Get("wrapS")}
}
func (ct *CubeTexture) SetWrapS(v *Wrapping) {
	ct.Set("wrapS", v)
}
func (ct *CubeTexture) WrapT() *Wrapping {
	return &Wrapping{Value: ct.Get("wrapT")}
}
func (ct *CubeTexture) SetWrapT(v *Wrapping) {
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
func (ct *CubeTexture) Clone() this {
	return this(ct.Call("clone"))
}
func (ct *CubeTexture) Copy(source *Texture) this {
	return this(ct.Call("copy", source))
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
func (ct *CubeTexture) TransformUv(uv math.Vector) {
	ct.Call("transformUv", uv)
}

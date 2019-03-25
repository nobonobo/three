package textures

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type VideoTexture struct {
	js.Value
}

func NewVideoTexture(video *HTMLVideoElement, mapping *Mapping, wrapS *Wrapping, wrapT *Wrapping, magFilter *TextureFilter, minFilter *TextureFilter, format *PixelFormat, typ *TextureDataType, anisotropy int) *VideoTexture {
	return &VideoTexture{Value: get("VideoTexture").New(video, mapping, wrapS, wrapT, magFilter, minFilter, format, typ, anisotropy)}
}
func (vt *VideoTexture) Anisotropy() int {
	return vt.Get("anisotropy").Int()
}
func (vt *VideoTexture) SetAnisotropy(v int) {
	vt.Set("anisotropy", v)
}
func (vt *VideoTexture) Center() *math.Vector2 {
	return &math.Vector2{Value: vt.Get("center")}
}
func (vt *VideoTexture) SetCenter(v *math.Vector2) {
	vt.Set("center", v)
}
func (vt *VideoTexture) Encoding() *TextureEncoding {
	return &TextureEncoding{Value: vt.Get("encoding")}
}
func (vt *VideoTexture) SetEncoding(v *TextureEncoding) {
	vt.Set("encoding", v)
}
func (vt *VideoTexture) FlipY() bool {
	return vt.Get("flipY").Bool()
}
func (vt *VideoTexture) SetFlipY(v bool) {
	vt.Set("flipY", v)
}
func (vt *VideoTexture) Format() *PixelFormat {
	return &PixelFormat{Value: vt.Get("format")}
}
func (vt *VideoTexture) SetFormat(v *PixelFormat) {
	vt.Set("format", v)
}
func (vt *VideoTexture) GenerateMipmaps() bool {
	return vt.Get("generateMipmaps").Bool()
}
func (vt *VideoTexture) SetGenerateMipmaps(v bool) {
	vt.Set("generateMipmaps", v)
}
func (vt *VideoTexture) Id() int {
	return vt.Get("id").Int()
}
func (vt *VideoTexture) SetId(v int) {
	vt.Set("id", v)
}
func (vt *VideoTexture) Image() js.Value {
	return vt.Get("image")
}
func (vt *VideoTexture) SetImage(v js.Value) {
	vt.Set("image", v)
}
func (vt *VideoTexture) MagFilter() *TextureFilter {
	return &TextureFilter{Value: vt.Get("magFilter")}
}
func (vt *VideoTexture) SetMagFilter(v *TextureFilter) {
	vt.Set("magFilter", v)
}
func (vt *VideoTexture) Mapping() *Mapping {
	return &Mapping{Value: vt.Get("mapping")}
}
func (vt *VideoTexture) SetMapping(v *Mapping) {
	vt.Set("mapping", v)
}
func (vt *VideoTexture) MinFilter() *TextureFilter {
	return &TextureFilter{Value: vt.Get("minFilter")}
}
func (vt *VideoTexture) SetMinFilter(v *TextureFilter) {
	vt.Set("minFilter", v)
}
func (vt *VideoTexture) Mipmaps() []ImageData {
	return []ImageData(vt.Get("mipmaps"))
}
func (vt *VideoTexture) SetMipmaps(v []ImageData) {
	vt.Set("mipmaps", v)
}
func (vt *VideoTexture) Name() string {
	return vt.Get("name").String()
}
func (vt *VideoTexture) SetName(v string) {
	vt.Set("name", v)
}
func (vt *VideoTexture) NeedsUpdate() bool {
	return vt.Get("needsUpdate").Bool()
}
func (vt *VideoTexture) SetNeedsUpdate(v bool) {
	vt.Set("needsUpdate", v)
}
func (vt *VideoTexture) Offset() *math.Vector2 {
	return &math.Vector2{Value: vt.Get("offset")}
}
func (vt *VideoTexture) SetOffset(v *math.Vector2) {
	vt.Set("offset", v)
}
func (vt *VideoTexture) OnUpdate() js.Value {
	return vt.Get("onUpdate")
}
func (vt *VideoTexture) SetOnUpdate(v js.Value) {
	vt.Set("onUpdate", v)
}
func (vt *VideoTexture) PremultiplyAlpha() bool {
	return vt.Get("premultiplyAlpha").Bool()
}
func (vt *VideoTexture) SetPremultiplyAlpha(v bool) {
	vt.Set("premultiplyAlpha", v)
}
func (vt *VideoTexture) Repeat() *math.Vector2 {
	return &math.Vector2{Value: vt.Get("repeat")}
}
func (vt *VideoTexture) SetRepeat(v *math.Vector2) {
	vt.Set("repeat", v)
}
func (vt *VideoTexture) Rotation() int {
	return vt.Get("rotation").Int()
}
func (vt *VideoTexture) SetRotation(v int) {
	vt.Set("rotation", v)
}
func (vt *VideoTexture) SourceFile() string {
	return vt.Get("sourceFile").String()
}
func (vt *VideoTexture) SetSourceFile(v string) {
	vt.Set("sourceFile", v)
}
func (vt *VideoTexture) Type() *TextureDataType {
	return &TextureDataType{Value: vt.Get("type")}
}
func (vt *VideoTexture) SetType(v *TextureDataType) {
	vt.Set("type", v)
}
func (vt *VideoTexture) UnpackAlignment() int {
	return vt.Get("unpackAlignment").Int()
}
func (vt *VideoTexture) SetUnpackAlignment(v int) {
	vt.Set("unpackAlignment", v)
}
func (vt *VideoTexture) Uuid() string {
	return vt.Get("uuid").String()
}
func (vt *VideoTexture) SetUuid(v string) {
	vt.Set("uuid", v)
}
func (vt *VideoTexture) Version() int {
	return vt.Get("version").Int()
}
func (vt *VideoTexture) SetVersion(v int) {
	vt.Set("version", v)
}
func (vt *VideoTexture) WrapS() *Wrapping {
	return &Wrapping{Value: vt.Get("wrapS")}
}
func (vt *VideoTexture) SetWrapS(v *Wrapping) {
	vt.Set("wrapS", v)
}
func (vt *VideoTexture) WrapT() *Wrapping {
	return &Wrapping{Value: vt.Get("wrapT")}
}
func (vt *VideoTexture) SetWrapT(v *Wrapping) {
	vt.Set("wrapT", v)
}
func (vt *VideoTexture) DEFAULT_IMAGE() js.Value {
	return vt.Get("DEFAULT_IMAGE")
}
func (vt *VideoTexture) SetDEFAULT_IMAGE(v js.Value) {
	vt.Set("DEFAULT_IMAGE", v)
}
func (vt *VideoTexture) DEFAULT_MAPPING() js.Value {
	return vt.Get("DEFAULT_MAPPING")
}
func (vt *VideoTexture) SetDEFAULT_MAPPING(v js.Value) {
	vt.Set("DEFAULT_MAPPING", v)
}
func (vt *VideoTexture) AddEventListener(typ string, listener js.Value) {
	vt.Call("addEventListener", typ, listener)
}
func (vt *VideoTexture) Clone() this {
	return this(vt.Call("clone"))
}
func (vt *VideoTexture) Copy(source *Texture) this {
	return this(vt.Call("copy", source))
}
func (vt *VideoTexture) DispatchEvent(event js.Value) {
	vt.Call("dispatchEvent", event)
}
func (vt *VideoTexture) Dispose() {
	vt.Call("dispose")
}
func (vt *VideoTexture) HasEventListener(typ string, listener js.Value) bool {
	return vt.Call("hasEventListener", typ, listener).Bool()
}
func (vt *VideoTexture) RemoveEventListener(typ string, listener js.Value) {
	vt.Call("removeEventListener", typ, listener)
}
func (vt *VideoTexture) ToJSON(meta js.Value) js.Value {
	return vt.Call("toJSON", meta)
}
func (vt *VideoTexture) TransformUv(uv math.Vector) {
	vt.Call("transformUv", uv)
}

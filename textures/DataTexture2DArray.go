package textures

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type DataTexture2DArray struct {
	js.Value
}

func NewDataTexture2DArray(data ArrayBuffer, width int, height int, depth int) *DataTexture2DArray {
	return &DataTexture2DArray{Value: get("DataTexture2DArray").New(data, width, height, depth)}
}
func (dtda *DataTexture2DArray) Anisotropy() int {
	return dtda.Get("anisotropy").Int()
}
func (dtda *DataTexture2DArray) SetAnisotropy(v int) {
	dtda.Set("anisotropy", v)
}
func (dtda *DataTexture2DArray) Center() *math.Vector2 {
	return &math.Vector2{Value: dtda.Get("center")}
}
func (dtda *DataTexture2DArray) SetCenter(v *math.Vector2) {
	dtda.Set("center", v)
}
func (dtda *DataTexture2DArray) Encoding() *TextureEncoding {
	return &TextureEncoding{Value: dtda.Get("encoding")}
}
func (dtda *DataTexture2DArray) SetEncoding(v *TextureEncoding) {
	dtda.Set("encoding", v)
}
func (dtda *DataTexture2DArray) FlipY() bool {
	return dtda.Get("flipY").Bool()
}
func (dtda *DataTexture2DArray) SetFlipY(v bool) {
	dtda.Set("flipY", v)
}
func (dtda *DataTexture2DArray) Format() *PixelFormat {
	return &PixelFormat{Value: dtda.Get("format")}
}
func (dtda *DataTexture2DArray) SetFormat(v *PixelFormat) {
	dtda.Set("format", v)
}
func (dtda *DataTexture2DArray) GenerateMipmaps() bool {
	return dtda.Get("generateMipmaps").Bool()
}
func (dtda *DataTexture2DArray) SetGenerateMipmaps(v bool) {
	dtda.Set("generateMipmaps", v)
}
func (dtda *DataTexture2DArray) Id() int {
	return dtda.Get("id").Int()
}
func (dtda *DataTexture2DArray) SetId(v int) {
	dtda.Set("id", v)
}
func (dtda *DataTexture2DArray) Image() js.Value {
	return dtda.Get("image")
}
func (dtda *DataTexture2DArray) SetImage(v js.Value) {
	dtda.Set("image", v)
}
func (dtda *DataTexture2DArray) MagFilter() *TextureFilter {
	return &TextureFilter{Value: dtda.Get("magFilter")}
}
func (dtda *DataTexture2DArray) SetMagFilter(v *TextureFilter) {
	dtda.Set("magFilter", v)
}
func (dtda *DataTexture2DArray) Mapping() *Mapping {
	return &Mapping{Value: dtda.Get("mapping")}
}
func (dtda *DataTexture2DArray) SetMapping(v *Mapping) {
	dtda.Set("mapping", v)
}
func (dtda *DataTexture2DArray) MinFilter() *TextureFilter {
	return &TextureFilter{Value: dtda.Get("minFilter")}
}
func (dtda *DataTexture2DArray) SetMinFilter(v *TextureFilter) {
	dtda.Set("minFilter", v)
}
func (dtda *DataTexture2DArray) Mipmaps() []ImageData {
	return []ImageData(dtda.Get("mipmaps"))
}
func (dtda *DataTexture2DArray) SetMipmaps(v []ImageData) {
	dtda.Set("mipmaps", v)
}
func (dtda *DataTexture2DArray) Name() string {
	return dtda.Get("name").String()
}
func (dtda *DataTexture2DArray) SetName(v string) {
	dtda.Set("name", v)
}
func (dtda *DataTexture2DArray) NeedsUpdate() bool {
	return dtda.Get("needsUpdate").Bool()
}
func (dtda *DataTexture2DArray) SetNeedsUpdate(v bool) {
	dtda.Set("needsUpdate", v)
}
func (dtda *DataTexture2DArray) Offset() *math.Vector2 {
	return &math.Vector2{Value: dtda.Get("offset")}
}
func (dtda *DataTexture2DArray) SetOffset(v *math.Vector2) {
	dtda.Set("offset", v)
}
func (dtda *DataTexture2DArray) OnUpdate() js.Value {
	return dtda.Get("onUpdate")
}
func (dtda *DataTexture2DArray) SetOnUpdate(v js.Value) {
	dtda.Set("onUpdate", v)
}
func (dtda *DataTexture2DArray) PremultiplyAlpha() bool {
	return dtda.Get("premultiplyAlpha").Bool()
}
func (dtda *DataTexture2DArray) SetPremultiplyAlpha(v bool) {
	dtda.Set("premultiplyAlpha", v)
}
func (dtda *DataTexture2DArray) Repeat() *math.Vector2 {
	return &math.Vector2{Value: dtda.Get("repeat")}
}
func (dtda *DataTexture2DArray) SetRepeat(v *math.Vector2) {
	dtda.Set("repeat", v)
}
func (dtda *DataTexture2DArray) Rotation() int {
	return dtda.Get("rotation").Int()
}
func (dtda *DataTexture2DArray) SetRotation(v int) {
	dtda.Set("rotation", v)
}
func (dtda *DataTexture2DArray) SourceFile() string {
	return dtda.Get("sourceFile").String()
}
func (dtda *DataTexture2DArray) SetSourceFile(v string) {
	dtda.Set("sourceFile", v)
}
func (dtda *DataTexture2DArray) Type() *TextureDataType {
	return &TextureDataType{Value: dtda.Get("type")}
}
func (dtda *DataTexture2DArray) SetType(v *TextureDataType) {
	dtda.Set("type", v)
}
func (dtda *DataTexture2DArray) UnpackAlignment() int {
	return dtda.Get("unpackAlignment").Int()
}
func (dtda *DataTexture2DArray) SetUnpackAlignment(v int) {
	dtda.Set("unpackAlignment", v)
}
func (dtda *DataTexture2DArray) Uuid() string {
	return dtda.Get("uuid").String()
}
func (dtda *DataTexture2DArray) SetUuid(v string) {
	dtda.Set("uuid", v)
}
func (dtda *DataTexture2DArray) Version() int {
	return dtda.Get("version").Int()
}
func (dtda *DataTexture2DArray) SetVersion(v int) {
	dtda.Set("version", v)
}
func (dtda *DataTexture2DArray) WrapS() *Wrapping {
	return &Wrapping{Value: dtda.Get("wrapS")}
}
func (dtda *DataTexture2DArray) SetWrapS(v *Wrapping) {
	dtda.Set("wrapS", v)
}
func (dtda *DataTexture2DArray) WrapT() *Wrapping {
	return &Wrapping{Value: dtda.Get("wrapT")}
}
func (dtda *DataTexture2DArray) SetWrapT(v *Wrapping) {
	dtda.Set("wrapT", v)
}
func (dtda *DataTexture2DArray) DEFAULT_IMAGE() js.Value {
	return dtda.Get("DEFAULT_IMAGE")
}
func (dtda *DataTexture2DArray) SetDEFAULT_IMAGE(v js.Value) {
	dtda.Set("DEFAULT_IMAGE", v)
}
func (dtda *DataTexture2DArray) DEFAULT_MAPPING() js.Value {
	return dtda.Get("DEFAULT_MAPPING")
}
func (dtda *DataTexture2DArray) SetDEFAULT_MAPPING(v js.Value) {
	dtda.Set("DEFAULT_MAPPING", v)
}
func (dtda *DataTexture2DArray) AddEventListener(typ string, listener js.Value) {
	dtda.Call("addEventListener", typ, listener)
}
func (dtda *DataTexture2DArray) Clone() this {
	return this(dtda.Call("clone"))
}
func (dtda *DataTexture2DArray) Copy(source *Texture) this {
	return this(dtda.Call("copy", source))
}
func (dtda *DataTexture2DArray) DispatchEvent(event js.Value) {
	dtda.Call("dispatchEvent", event)
}
func (dtda *DataTexture2DArray) Dispose() {
	dtda.Call("dispose")
}
func (dtda *DataTexture2DArray) HasEventListener(typ string, listener js.Value) bool {
	return dtda.Call("hasEventListener", typ, listener).Bool()
}
func (dtda *DataTexture2DArray) RemoveEventListener(typ string, listener js.Value) {
	dtda.Call("removeEventListener", typ, listener)
}
func (dtda *DataTexture2DArray) ToJSON(meta js.Value) js.Value {
	return dtda.Call("toJSON", meta)
}
func (dtda *DataTexture2DArray) TransformUv(uv math.Vector) {
	dtda.Call("transformUv", uv)
}

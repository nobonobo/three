// Code generated from "materials/Material.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

func MaterialIdCount() int {
	return get("MaterialIdCount").Int()
}
func SetMaterialIdCount(v int) {
	set("MaterialIdCount", v)
}

type MaterialParameters interface {
}
type Material struct {
	js.Value
}

func NewMaterial() *Material {
	return &Material{Value: get("Material").New()}
}
func (mm *Material) AlphaTest() float64 {
	return mm.Get("alphaTest").Float()
}
func (mm *Material) SetAlphaTest(v float64) {
	mm.Set("alphaTest", v)
}
func (mm *Material) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mm.Get("blendDst"))
}
func (mm *Material) SetBlendDst(v BlendingDstFactor) {
	mm.Set("blendDst", v)
}
func (mm *Material) BlendDstAlpha() float64 {
	return mm.Get("blendDstAlpha").Float()
}
func (mm *Material) SetBlendDstAlpha(v float64) {
	mm.Set("blendDstAlpha", v)
}
func (mm *Material) BlendEquation() BlendingEquation {
	return BlendingEquation(mm.Get("blendEquation"))
}
func (mm *Material) SetBlendEquation(v BlendingEquation) {
	mm.Set("blendEquation", v)
}
func (mm *Material) BlendEquationAlpha() float64 {
	return mm.Get("blendEquationAlpha").Float()
}
func (mm *Material) SetBlendEquationAlpha(v float64) {
	mm.Set("blendEquationAlpha", v)
}
func (mm *Material) BlendSrc() js.Value {
	return mm.Get("blendSrc")
}
func (mm *Material) SetBlendSrc(v js.Value) {
	mm.Set("blendSrc", v)
}
func (mm *Material) BlendSrcAlpha() float64 {
	return mm.Get("blendSrcAlpha").Float()
}
func (mm *Material) SetBlendSrcAlpha(v float64) {
	mm.Set("blendSrcAlpha", v)
}
func (mm *Material) Blending() Blending {
	return Blending(mm.Get("blending"))
}
func (mm *Material) SetBlending(v Blending) {
	mm.Set("blending", v)
}
func (mm *Material) ClipIntersection() bool {
	return mm.Get("clipIntersection").Bool()
}
func (mm *Material) SetClipIntersection(v bool) {
	mm.Set("clipIntersection", v)
}
func (mm *Material) ClipShadows() bool {
	return mm.Get("clipShadows").Bool()
}
func (mm *Material) SetClipShadows(v bool) {
	mm.Set("clipShadows", v)
}
func (mm *Material) ClippingPlanes() js.Value {
	return mm.Get("clippingPlanes")
}
func (mm *Material) SetClippingPlanes(v js.Value) {
	mm.Set("clippingPlanes", v)
}
func (mm *Material) ColorWrite() bool {
	return mm.Get("colorWrite").Bool()
}
func (mm *Material) SetColorWrite(v bool) {
	mm.Set("colorWrite", v)
}
func (mm *Material) DepthFunc() DepthModes {
	return DepthModes(mm.Get("depthFunc"))
}
func (mm *Material) SetDepthFunc(v DepthModes) {
	mm.Set("depthFunc", v)
}
func (mm *Material) DepthTest() bool {
	return mm.Get("depthTest").Bool()
}
func (mm *Material) SetDepthTest(v bool) {
	mm.Set("depthTest", v)
}
func (mm *Material) DepthWrite() bool {
	return mm.Get("depthWrite").Bool()
}
func (mm *Material) SetDepthWrite(v bool) {
	mm.Set("depthWrite", v)
}
func (mm *Material) Dithering() bool {
	return mm.Get("dithering").Bool()
}
func (mm *Material) SetDithering(v bool) {
	mm.Set("dithering", v)
}
func (mm *Material) FlatShading() bool {
	return mm.Get("flatShading").Bool()
}
func (mm *Material) SetFlatShading(v bool) {
	mm.Set("flatShading", v)
}
func (mm *Material) Fog() bool {
	return mm.Get("fog").Bool()
}
func (mm *Material) SetFog(v bool) {
	mm.Set("fog", v)
}
func (mm *Material) Id() int {
	return mm.Get("id").Int()
}
func (mm *Material) SetId(v int) {
	mm.Set("id", v)
}
func (mm *Material) IsMaterial() bool {
	return mm.Get("isMaterial").Bool()
}
func (mm *Material) SetIsMaterial(v bool) {
	mm.Set("isMaterial", v)
}
func (mm *Material) Lights() bool {
	return mm.Get("lights").Bool()
}
func (mm *Material) SetLights(v bool) {
	mm.Set("lights", v)
}
func (mm *Material) Name() string {
	return mm.Get("name").String()
}
func (mm *Material) SetName(v string) {
	mm.Set("name", v)
}
func (mm *Material) NeedsUpdate() bool {
	return mm.Get("needsUpdate").Bool()
}
func (mm *Material) SetNeedsUpdate(v bool) {
	mm.Set("needsUpdate", v)
}
func (mm *Material) Opacity() float64 {
	return mm.Get("opacity").Float()
}
func (mm *Material) SetOpacity(v float64) {
	mm.Set("opacity", v)
}
func (mm *Material) Overdraw() float64 {
	return mm.Get("overdraw").Float()
}
func (mm *Material) SetOverdraw(v float64) {
	mm.Set("overdraw", v)
}
func (mm *Material) PolygonOffset() bool {
	return mm.Get("polygonOffset").Bool()
}
func (mm *Material) SetPolygonOffset(v bool) {
	mm.Set("polygonOffset", v)
}
func (mm *Material) PolygonOffsetFactor() float64 {
	return mm.Get("polygonOffsetFactor").Float()
}
func (mm *Material) SetPolygonOffsetFactor(v float64) {
	mm.Set("polygonOffsetFactor", v)
}
func (mm *Material) PolygonOffsetUnits() float64 {
	return mm.Get("polygonOffsetUnits").Float()
}
func (mm *Material) SetPolygonOffsetUnits(v float64) {
	mm.Set("polygonOffsetUnits", v)
}
func (mm *Material) Precision() string {
	return mm.Get("precision").String()
}
func (mm *Material) SetPrecision(v string) {
	mm.Set("precision", v)
}
func (mm *Material) PremultipliedAlpha() bool {
	return mm.Get("premultipliedAlpha").Bool()
}
func (mm *Material) SetPremultipliedAlpha(v bool) {
	mm.Set("premultipliedAlpha", v)
}
func (mm *Material) Side() Side {
	return Side(mm.Get("side"))
}
func (mm *Material) SetSide(v Side) {
	mm.Set("side", v)
}
func (mm *Material) Transparent() bool {
	return mm.Get("transparent").Bool()
}
func (mm *Material) SetTransparent(v bool) {
	mm.Set("transparent", v)
}
func (mm *Material) Type() string {
	return mm.Get("type").String()
}
func (mm *Material) SetType(v string) {
	mm.Set("type", v)
}
func (mm *Material) UserData() js.Value {
	return mm.Get("userData")
}
func (mm *Material) SetUserData(v js.Value) {
	mm.Set("userData", v)
}
func (mm *Material) Uuid() string {
	return mm.Get("uuid").String()
}
func (mm *Material) SetUuid(v string) {
	mm.Set("uuid", v)
}
func (mm *Material) VertexColors() Colors {
	return Colors(mm.Get("vertexColors"))
}
func (mm *Material) SetVertexColors(v Colors) {
	mm.Set("vertexColors", v)
}
func (mm *Material) VertexTangents() bool {
	return mm.Get("vertexTangents").Bool()
}
func (mm *Material) SetVertexTangents(v bool) {
	mm.Set("vertexTangents", v)
}
func (mm *Material) Visible() bool {
	return mm.Get("visible").Bool()
}
func (mm *Material) SetVisible(v bool) {
	mm.Set("visible", v)
}
func (mm *Material) AddEventListener(typ string, listener js.Value) {
	mm.Call("addEventListener", typ, listener)
}
func (mm *Material) Clone() *Material {
	return &Material{Value: mm.Call("clone")}
}
func (mm *Material) Copy(material *Material) *Material {
	return &Material{Value: mm.Call("copy", material)}
}
func (mm *Material) DispatchEvent(event js.Value) {
	mm.Call("dispatchEvent", event)
}
func (mm *Material) Dispose() {
	mm.Call("dispose")
}
func (mm *Material) HasEventListener(typ string, listener js.Value) bool {
	return mm.Call("hasEventListener", typ, listener).Bool()
}
func (mm *Material) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mm.Call("onBeforeCompile", shader, renderer)
}
func (mm *Material) RemoveEventListener(typ string, listener js.Value) {
	mm.Call("removeEventListener", typ, listener)
}
func (mm *Material) SetValues(values MaterialParameters) {
	mm.Call("setValues", values)
}
func (mm *Material) ToJSON(meta js.Value) js.Value {
	return mm.Call("toJSON", meta)
}
func (mm *Material) Update() {
	mm.Call("update")
}

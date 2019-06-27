// Code generated from "materials/LineBasicMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type LineBasicMaterialParameters interface {
}

// LineBasicMaterial extend: [Material]
type LineBasicMaterial struct {
	js.Value
}

func NewLineBasicMaterial(parameters LineBasicMaterialParameters) *LineBasicMaterial {
	return &LineBasicMaterial{Value: get("LineBasicMaterial").New(parameters)}
}
func (lbm *LineBasicMaterial) JSValue() js.Value {
	return lbm.Value
}
func (lbm *LineBasicMaterial) AlphaTest() float64 {
	return lbm.Get("alphaTest").Float()
}
func (lbm *LineBasicMaterial) SetAlphaTest(v float64) {
	lbm.Set("alphaTest", v)
}
func (lbm *LineBasicMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(lbm.Get("blendDst"))
}
func (lbm *LineBasicMaterial) SetBlendDst(v BlendingDstFactor) {
	lbm.Set("blendDst", v)
}
func (lbm *LineBasicMaterial) BlendDstAlpha() float64 {
	return lbm.Get("blendDstAlpha").Float()
}
func (lbm *LineBasicMaterial) SetBlendDstAlpha(v float64) {
	lbm.Set("blendDstAlpha", v)
}
func (lbm *LineBasicMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(lbm.Get("blendEquation"))
}
func (lbm *LineBasicMaterial) SetBlendEquation(v BlendingEquation) {
	lbm.Set("blendEquation", v)
}
func (lbm *LineBasicMaterial) BlendEquationAlpha() float64 {
	return lbm.Get("blendEquationAlpha").Float()
}
func (lbm *LineBasicMaterial) SetBlendEquationAlpha(v float64) {
	lbm.Set("blendEquationAlpha", v)
}
func (lbm *LineBasicMaterial) BlendSrc() js.Value {
	return lbm.Get("blendSrc")
}
func (lbm *LineBasicMaterial) SetBlendSrc(v js.Value) {
	lbm.Set("blendSrc", v)
}
func (lbm *LineBasicMaterial) BlendSrcAlpha() float64 {
	return lbm.Get("blendSrcAlpha").Float()
}
func (lbm *LineBasicMaterial) SetBlendSrcAlpha(v float64) {
	lbm.Set("blendSrcAlpha", v)
}
func (lbm *LineBasicMaterial) Blending() Blending {
	return Blending(lbm.Get("blending"))
}
func (lbm *LineBasicMaterial) SetBlending(v Blending) {
	lbm.Set("blending", v)
}
func (lbm *LineBasicMaterial) ClipIntersection() bool {
	return lbm.Get("clipIntersection").Bool()
}
func (lbm *LineBasicMaterial) SetClipIntersection(v bool) {
	lbm.Set("clipIntersection", v)
}
func (lbm *LineBasicMaterial) ClipShadows() bool {
	return lbm.Get("clipShadows").Bool()
}
func (lbm *LineBasicMaterial) SetClipShadows(v bool) {
	lbm.Set("clipShadows", v)
}
func (lbm *LineBasicMaterial) ClippingPlanes() js.Value {
	return lbm.Get("clippingPlanes")
}
func (lbm *LineBasicMaterial) SetClippingPlanes(v js.Value) {
	lbm.Set("clippingPlanes", v)
}
func (lbm *LineBasicMaterial) Color() *Color {
	return &Color{Value: lbm.Get("color")}
}
func (lbm *LineBasicMaterial) SetColor(v *Color) {
	lbm.Set("color", v.Value)
}
func (lbm *LineBasicMaterial) ColorWrite() bool {
	return lbm.Get("colorWrite").Bool()
}
func (lbm *LineBasicMaterial) SetColorWrite(v bool) {
	lbm.Set("colorWrite", v)
}
func (lbm *LineBasicMaterial) DepthFunc() DepthModes {
	return DepthModes(lbm.Get("depthFunc"))
}
func (lbm *LineBasicMaterial) SetDepthFunc(v DepthModes) {
	lbm.Set("depthFunc", v)
}
func (lbm *LineBasicMaterial) DepthTest() bool {
	return lbm.Get("depthTest").Bool()
}
func (lbm *LineBasicMaterial) SetDepthTest(v bool) {
	lbm.Set("depthTest", v)
}
func (lbm *LineBasicMaterial) DepthWrite() bool {
	return lbm.Get("depthWrite").Bool()
}
func (lbm *LineBasicMaterial) SetDepthWrite(v bool) {
	lbm.Set("depthWrite", v)
}
func (lbm *LineBasicMaterial) Dithering() bool {
	return lbm.Get("dithering").Bool()
}
func (lbm *LineBasicMaterial) SetDithering(v bool) {
	lbm.Set("dithering", v)
}
func (lbm *LineBasicMaterial) FlatShading() bool {
	return lbm.Get("flatShading").Bool()
}
func (lbm *LineBasicMaterial) SetFlatShading(v bool) {
	lbm.Set("flatShading", v)
}
func (lbm *LineBasicMaterial) Fog() bool {
	return lbm.Get("fog").Bool()
}
func (lbm *LineBasicMaterial) SetFog(v bool) {
	lbm.Set("fog", v)
}
func (lbm *LineBasicMaterial) Id() int {
	return lbm.Get("id").Int()
}
func (lbm *LineBasicMaterial) SetId(v int) {
	lbm.Set("id", v)
}
func (lbm *LineBasicMaterial) IsMaterial() bool {
	return lbm.Get("isMaterial").Bool()
}
func (lbm *LineBasicMaterial) SetIsMaterial(v bool) {
	lbm.Set("isMaterial", v)
}
func (lbm *LineBasicMaterial) Lights() bool {
	return lbm.Get("lights").Bool()
}
func (lbm *LineBasicMaterial) SetLights(v bool) {
	lbm.Set("lights", v)
}
func (lbm *LineBasicMaterial) Linecap() string {
	return lbm.Get("linecap").String()
}
func (lbm *LineBasicMaterial) SetLinecap(v string) {
	lbm.Set("linecap", v)
}
func (lbm *LineBasicMaterial) Linejoin() string {
	return lbm.Get("linejoin").String()
}
func (lbm *LineBasicMaterial) SetLinejoin(v string) {
	lbm.Set("linejoin", v)
}
func (lbm *LineBasicMaterial) Linewidth() float64 {
	return lbm.Get("linewidth").Float()
}
func (lbm *LineBasicMaterial) SetLinewidth(v float64) {
	lbm.Set("linewidth", v)
}
func (lbm *LineBasicMaterial) Name() string {
	return lbm.Get("name").String()
}
func (lbm *LineBasicMaterial) SetName(v string) {
	lbm.Set("name", v)
}
func (lbm *LineBasicMaterial) NeedsUpdate() bool {
	return lbm.Get("needsUpdate").Bool()
}
func (lbm *LineBasicMaterial) SetNeedsUpdate(v bool) {
	lbm.Set("needsUpdate", v)
}
func (lbm *LineBasicMaterial) Opacity() float64 {
	return lbm.Get("opacity").Float()
}
func (lbm *LineBasicMaterial) SetOpacity(v float64) {
	lbm.Set("opacity", v)
}
func (lbm *LineBasicMaterial) Overdraw() float64 {
	return lbm.Get("overdraw").Float()
}
func (lbm *LineBasicMaterial) SetOverdraw(v float64) {
	lbm.Set("overdraw", v)
}
func (lbm *LineBasicMaterial) PolygonOffset() bool {
	return lbm.Get("polygonOffset").Bool()
}
func (lbm *LineBasicMaterial) SetPolygonOffset(v bool) {
	lbm.Set("polygonOffset", v)
}
func (lbm *LineBasicMaterial) PolygonOffsetFactor() float64 {
	return lbm.Get("polygonOffsetFactor").Float()
}
func (lbm *LineBasicMaterial) SetPolygonOffsetFactor(v float64) {
	lbm.Set("polygonOffsetFactor", v)
}
func (lbm *LineBasicMaterial) PolygonOffsetUnits() float64 {
	return lbm.Get("polygonOffsetUnits").Float()
}
func (lbm *LineBasicMaterial) SetPolygonOffsetUnits(v float64) {
	lbm.Set("polygonOffsetUnits", v)
}
func (lbm *LineBasicMaterial) Precision() string {
	return lbm.Get("precision").String()
}
func (lbm *LineBasicMaterial) SetPrecision(v string) {
	lbm.Set("precision", v)
}
func (lbm *LineBasicMaterial) PremultipliedAlpha() bool {
	return lbm.Get("premultipliedAlpha").Bool()
}
func (lbm *LineBasicMaterial) SetPremultipliedAlpha(v bool) {
	lbm.Set("premultipliedAlpha", v)
}
func (lbm *LineBasicMaterial) Side() Side {
	return Side(lbm.Get("side"))
}
func (lbm *LineBasicMaterial) SetSide(v Side) {
	lbm.Set("side", v)
}
func (lbm *LineBasicMaterial) Transparent() bool {
	return lbm.Get("transparent").Bool()
}
func (lbm *LineBasicMaterial) SetTransparent(v bool) {
	lbm.Set("transparent", v)
}
func (lbm *LineBasicMaterial) Type() string {
	return lbm.Get("type").String()
}
func (lbm *LineBasicMaterial) SetType(v string) {
	lbm.Set("type", v)
}
func (lbm *LineBasicMaterial) UserData() js.Value {
	return lbm.Get("userData")
}
func (lbm *LineBasicMaterial) SetUserData(v js.Value) {
	lbm.Set("userData", v)
}
func (lbm *LineBasicMaterial) Uuid() string {
	return lbm.Get("uuid").String()
}
func (lbm *LineBasicMaterial) SetUuid(v string) {
	lbm.Set("uuid", v)
}
func (lbm *LineBasicMaterial) VertexColors() Colors {
	return Colors(lbm.Get("vertexColors"))
}
func (lbm *LineBasicMaterial) SetVertexColors(v Colors) {
	lbm.Set("vertexColors", v)
}
func (lbm *LineBasicMaterial) VertexTangents() bool {
	return lbm.Get("vertexTangents").Bool()
}
func (lbm *LineBasicMaterial) SetVertexTangents(v bool) {
	lbm.Set("vertexTangents", v)
}
func (lbm *LineBasicMaterial) Visible() bool {
	return lbm.Get("visible").Bool()
}
func (lbm *LineBasicMaterial) SetVisible(v bool) {
	lbm.Set("visible", v)
}
func (lbm *LineBasicMaterial) AddEventListener(typ string, listener js.Value) {
	lbm.Call("addEventListener", typ, listener)
}
func (lbm *LineBasicMaterial) Clone() *LineBasicMaterial {
	return &LineBasicMaterial{Value: lbm.Call("clone")}
}
func (lbm *LineBasicMaterial) Copy(material Material) *LineBasicMaterial {
	return &LineBasicMaterial{Value: lbm.Call("copy", material.JSValue())}
}
func (lbm *LineBasicMaterial) DispatchEvent(event js.Value) {
	lbm.Call("dispatchEvent", event)
}
func (lbm *LineBasicMaterial) Dispose() {
	lbm.Call("dispose")
}
func (lbm *LineBasicMaterial) HasEventListener(typ string, listener js.Value) bool {
	return lbm.Call("hasEventListener", typ, listener).Bool()
}
func (lbm *LineBasicMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	lbm.Call("onBeforeCompile", shader, renderer)
}
func (lbm *LineBasicMaterial) RemoveEventListener(typ string, listener js.Value) {
	lbm.Call("removeEventListener", typ, listener)
}
func (lbm *LineBasicMaterial) SetValues(parameters LineBasicMaterialParameters) {
	lbm.Call("setValues", parameters)
}
func (lbm *LineBasicMaterial) ToJSON(meta js.Value) js.Value {
	return lbm.Call("toJSON", meta)
}
func (lbm *LineBasicMaterial) Update() {
	lbm.Call("update")
}

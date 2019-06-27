// Code generated from "materials/Material.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func MaterialIdCount() int {
	return get("MaterialIdCount").Int()
}
func SetMaterialIdCount(v int) {
	set("MaterialIdCount", v)
}

type MaterialParameters interface {
}
type Material interface {
	JSValue() js.Value
	AlphaTest() float64
	SetAlphaTest(v float64)
	BlendDst() BlendingDstFactor
	SetBlendDst(v BlendingDstFactor)
	BlendDstAlpha() float64
	SetBlendDstAlpha(v float64)
	BlendEquation() BlendingEquation
	SetBlendEquation(v BlendingEquation)
	BlendEquationAlpha() float64
	SetBlendEquationAlpha(v float64)
	BlendSrc() js.Value
	SetBlendSrc(v js.Value)
	BlendSrcAlpha() float64
	SetBlendSrcAlpha(v float64)
	Blending() Blending
	SetBlending(v Blending)
	ClipIntersection() bool
	SetClipIntersection(v bool)
	ClipShadows() bool
	SetClipShadows(v bool)
	ClippingPlanes() js.Value
	SetClippingPlanes(v js.Value)
	ColorWrite() bool
	SetColorWrite(v bool)
	DepthFunc() DepthModes
	SetDepthFunc(v DepthModes)
	DepthTest() bool
	SetDepthTest(v bool)
	DepthWrite() bool
	SetDepthWrite(v bool)
	Dithering() bool
	SetDithering(v bool)
	FlatShading() bool
	SetFlatShading(v bool)
	Fog() bool
	SetFog(v bool)
	Id() int
	SetId(v int)
	IsMaterial() bool
	SetIsMaterial(v bool)
	Lights() bool
	SetLights(v bool)
	Name() string
	SetName(v string)
	NeedsUpdate() bool
	SetNeedsUpdate(v bool)
	Opacity() float64
	SetOpacity(v float64)
	Overdraw() float64
	SetOverdraw(v float64)
	PolygonOffset() bool
	SetPolygonOffset(v bool)
	PolygonOffsetFactor() float64
	SetPolygonOffsetFactor(v float64)
	PolygonOffsetUnits() float64
	SetPolygonOffsetUnits(v float64)
	Precision() string
	SetPrecision(v string)
	PremultipliedAlpha() bool
	SetPremultipliedAlpha(v bool)
	Side() Side
	SetSide(v Side)
	Transparent() bool
	SetTransparent(v bool)
	Type() string
	SetType(v string)
	UserData() js.Value
	SetUserData(v js.Value)
	Uuid() string
	SetUuid(v string)
	VertexColors() Colors
	SetVertexColors(v Colors)
	VertexTangents() bool
	SetVertexTangents(v bool)
	Visible() bool
	SetVisible(v bool)
	AddEventListener(typ string, listener js.Value)
	Clone() Material
	Copy(material Material) Material
	DispatchEvent(event js.Value)
	Dispose()
	HasEventListener(typ string, listener js.Value) bool
	OnBeforeCompile(shader js.Value, renderer *WebGLRenderer)
	RemoveEventListener(typ string, listener js.Value)
	SetValues(values js.Value)
	ToJSON(meta js.Value) js.Value
	Update()
}

// MaterialImpl extend: [EventDispatcher]
type MaterialImpl struct {
	js.Value
}

func NewMaterial() *MaterialImpl {
	return &MaterialImpl{Value: get("Material").New()}
}
func (mm *MaterialImpl) JSValue() js.Value {
	return mm.Value
}
func (mm *MaterialImpl) AlphaTest() float64 {
	return mm.Get("alphaTest").Float()
}
func (mm *MaterialImpl) SetAlphaTest(v float64) {
	mm.Set("alphaTest", v)
}
func (mm *MaterialImpl) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mm.Get("blendDst"))
}
func (mm *MaterialImpl) SetBlendDst(v BlendingDstFactor) {
	mm.Set("blendDst", v)
}
func (mm *MaterialImpl) BlendDstAlpha() float64 {
	return mm.Get("blendDstAlpha").Float()
}
func (mm *MaterialImpl) SetBlendDstAlpha(v float64) {
	mm.Set("blendDstAlpha", v)
}
func (mm *MaterialImpl) BlendEquation() BlendingEquation {
	return BlendingEquation(mm.Get("blendEquation"))
}
func (mm *MaterialImpl) SetBlendEquation(v BlendingEquation) {
	mm.Set("blendEquation", v)
}
func (mm *MaterialImpl) BlendEquationAlpha() float64 {
	return mm.Get("blendEquationAlpha").Float()
}
func (mm *MaterialImpl) SetBlendEquationAlpha(v float64) {
	mm.Set("blendEquationAlpha", v)
}
func (mm *MaterialImpl) BlendSrc() js.Value {
	return mm.Get("blendSrc")
}
func (mm *MaterialImpl) SetBlendSrc(v js.Value) {
	mm.Set("blendSrc", v)
}
func (mm *MaterialImpl) BlendSrcAlpha() float64 {
	return mm.Get("blendSrcAlpha").Float()
}
func (mm *MaterialImpl) SetBlendSrcAlpha(v float64) {
	mm.Set("blendSrcAlpha", v)
}
func (mm *MaterialImpl) Blending() Blending {
	return Blending(mm.Get("blending"))
}
func (mm *MaterialImpl) SetBlending(v Blending) {
	mm.Set("blending", v)
}
func (mm *MaterialImpl) ClipIntersection() bool {
	return mm.Get("clipIntersection").Bool()
}
func (mm *MaterialImpl) SetClipIntersection(v bool) {
	mm.Set("clipIntersection", v)
}
func (mm *MaterialImpl) ClipShadows() bool {
	return mm.Get("clipShadows").Bool()
}
func (mm *MaterialImpl) SetClipShadows(v bool) {
	mm.Set("clipShadows", v)
}
func (mm *MaterialImpl) ClippingPlanes() js.Value {
	return mm.Get("clippingPlanes")
}
func (mm *MaterialImpl) SetClippingPlanes(v js.Value) {
	mm.Set("clippingPlanes", v)
}
func (mm *MaterialImpl) ColorWrite() bool {
	return mm.Get("colorWrite").Bool()
}
func (mm *MaterialImpl) SetColorWrite(v bool) {
	mm.Set("colorWrite", v)
}
func (mm *MaterialImpl) DepthFunc() DepthModes {
	return DepthModes(mm.Get("depthFunc"))
}
func (mm *MaterialImpl) SetDepthFunc(v DepthModes) {
	mm.Set("depthFunc", v)
}
func (mm *MaterialImpl) DepthTest() bool {
	return mm.Get("depthTest").Bool()
}
func (mm *MaterialImpl) SetDepthTest(v bool) {
	mm.Set("depthTest", v)
}
func (mm *MaterialImpl) DepthWrite() bool {
	return mm.Get("depthWrite").Bool()
}
func (mm *MaterialImpl) SetDepthWrite(v bool) {
	mm.Set("depthWrite", v)
}
func (mm *MaterialImpl) Dithering() bool {
	return mm.Get("dithering").Bool()
}
func (mm *MaterialImpl) SetDithering(v bool) {
	mm.Set("dithering", v)
}
func (mm *MaterialImpl) FlatShading() bool {
	return mm.Get("flatShading").Bool()
}
func (mm *MaterialImpl) SetFlatShading(v bool) {
	mm.Set("flatShading", v)
}
func (mm *MaterialImpl) Fog() bool {
	return mm.Get("fog").Bool()
}
func (mm *MaterialImpl) SetFog(v bool) {
	mm.Set("fog", v)
}
func (mm *MaterialImpl) Id() int {
	return mm.Get("id").Int()
}
func (mm *MaterialImpl) SetId(v int) {
	mm.Set("id", v)
}
func (mm *MaterialImpl) IsMaterial() bool {
	return mm.Get("isMaterial").Bool()
}
func (mm *MaterialImpl) SetIsMaterial(v bool) {
	mm.Set("isMaterial", v)
}
func (mm *MaterialImpl) Lights() bool {
	return mm.Get("lights").Bool()
}
func (mm *MaterialImpl) SetLights(v bool) {
	mm.Set("lights", v)
}
func (mm *MaterialImpl) Name() string {
	return mm.Get("name").String()
}
func (mm *MaterialImpl) SetName(v string) {
	mm.Set("name", v)
}
func (mm *MaterialImpl) NeedsUpdate() bool {
	return mm.Get("needsUpdate").Bool()
}
func (mm *MaterialImpl) SetNeedsUpdate(v bool) {
	mm.Set("needsUpdate", v)
}
func (mm *MaterialImpl) Opacity() float64 {
	return mm.Get("opacity").Float()
}
func (mm *MaterialImpl) SetOpacity(v float64) {
	mm.Set("opacity", v)
}
func (mm *MaterialImpl) Overdraw() float64 {
	return mm.Get("overdraw").Float()
}
func (mm *MaterialImpl) SetOverdraw(v float64) {
	mm.Set("overdraw", v)
}
func (mm *MaterialImpl) PolygonOffset() bool {
	return mm.Get("polygonOffset").Bool()
}
func (mm *MaterialImpl) SetPolygonOffset(v bool) {
	mm.Set("polygonOffset", v)
}
func (mm *MaterialImpl) PolygonOffsetFactor() float64 {
	return mm.Get("polygonOffsetFactor").Float()
}
func (mm *MaterialImpl) SetPolygonOffsetFactor(v float64) {
	mm.Set("polygonOffsetFactor", v)
}
func (mm *MaterialImpl) PolygonOffsetUnits() float64 {
	return mm.Get("polygonOffsetUnits").Float()
}
func (mm *MaterialImpl) SetPolygonOffsetUnits(v float64) {
	mm.Set("polygonOffsetUnits", v)
}
func (mm *MaterialImpl) Precision() string {
	return mm.Get("precision").String()
}
func (mm *MaterialImpl) SetPrecision(v string) {
	mm.Set("precision", v)
}
func (mm *MaterialImpl) PremultipliedAlpha() bool {
	return mm.Get("premultipliedAlpha").Bool()
}
func (mm *MaterialImpl) SetPremultipliedAlpha(v bool) {
	mm.Set("premultipliedAlpha", v)
}
func (mm *MaterialImpl) Side() Side {
	return Side(mm.Get("side"))
}
func (mm *MaterialImpl) SetSide(v Side) {
	mm.Set("side", v)
}
func (mm *MaterialImpl) Transparent() bool {
	return mm.Get("transparent").Bool()
}
func (mm *MaterialImpl) SetTransparent(v bool) {
	mm.Set("transparent", v)
}
func (mm *MaterialImpl) Type() string {
	return mm.Get("type").String()
}
func (mm *MaterialImpl) SetType(v string) {
	mm.Set("type", v)
}
func (mm *MaterialImpl) UserData() js.Value {
	return mm.Get("userData")
}
func (mm *MaterialImpl) SetUserData(v js.Value) {
	mm.Set("userData", v)
}
func (mm *MaterialImpl) Uuid() string {
	return mm.Get("uuid").String()
}
func (mm *MaterialImpl) SetUuid(v string) {
	mm.Set("uuid", v)
}
func (mm *MaterialImpl) VertexColors() Colors {
	return Colors(mm.Get("vertexColors"))
}
func (mm *MaterialImpl) SetVertexColors(v Colors) {
	mm.Set("vertexColors", v)
}
func (mm *MaterialImpl) VertexTangents() bool {
	return mm.Get("vertexTangents").Bool()
}
func (mm *MaterialImpl) SetVertexTangents(v bool) {
	mm.Set("vertexTangents", v)
}
func (mm *MaterialImpl) Visible() bool {
	return mm.Get("visible").Bool()
}
func (mm *MaterialImpl) SetVisible(v bool) {
	mm.Set("visible", v)
}
func (mm *MaterialImpl) AddEventListener(typ string, listener js.Value) {
	mm.Call("addEventListener", typ, listener)
}
func (mm *MaterialImpl) Clone() Material {
	return &MaterialImpl{Value: mm.Call("clone")}
}
func (mm *MaterialImpl) Copy(material Material) Material {
	return &MaterialImpl{Value: mm.Call("copy", material.JSValue())}
}
func (mm *MaterialImpl) DispatchEvent(event js.Value) {
	mm.Call("dispatchEvent", event)
}
func (mm *MaterialImpl) Dispose() {
	mm.Call("dispose")
}
func (mm *MaterialImpl) HasEventListener(typ string, listener js.Value) bool {
	return mm.Call("hasEventListener", typ, listener).Bool()
}
func (mm *MaterialImpl) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mm.Call("onBeforeCompile", shader, renderer)
}
func (mm *MaterialImpl) RemoveEventListener(typ string, listener js.Value) {
	mm.Call("removeEventListener", typ, listener)
}
func (mm *MaterialImpl) SetValues(values js.Value) {
	mm.Call("setValues", values)
}
func (mm *MaterialImpl) ToJSON(meta js.Value) js.Value {
	return mm.Call("toJSON", meta)
}
func (mm *MaterialImpl) Update() {
	mm.Call("update")
}

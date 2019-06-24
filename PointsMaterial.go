// Code generated from "materials/PointsMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type PointsMaterialParameters interface {
}
type MultiMaterial struct {
	js.Value
}

func NewMultiMaterial(materials js.Value) *MultiMaterial {
	return &MultiMaterial{Value: get("MultiMaterial").New(materials)}
}
func (mm *MultiMaterial) AlphaTest() float64 {
	return mm.Get("alphaTest").Float()
}
func (mm *MultiMaterial) SetAlphaTest(v float64) {
	mm.Set("alphaTest", v)
}
func (mm *MultiMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mm.Get("blendDst"))
}
func (mm *MultiMaterial) SetBlendDst(v BlendingDstFactor) {
	mm.Set("blendDst", v)
}
func (mm *MultiMaterial) BlendDstAlpha() float64 {
	return mm.Get("blendDstAlpha").Float()
}
func (mm *MultiMaterial) SetBlendDstAlpha(v float64) {
	mm.Set("blendDstAlpha", v)
}
func (mm *MultiMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mm.Get("blendEquation"))
}
func (mm *MultiMaterial) SetBlendEquation(v BlendingEquation) {
	mm.Set("blendEquation", v)
}
func (mm *MultiMaterial) BlendEquationAlpha() float64 {
	return mm.Get("blendEquationAlpha").Float()
}
func (mm *MultiMaterial) SetBlendEquationAlpha(v float64) {
	mm.Set("blendEquationAlpha", v)
}
func (mm *MultiMaterial) BlendSrc() js.Value {
	return mm.Get("blendSrc")
}
func (mm *MultiMaterial) SetBlendSrc(v js.Value) {
	mm.Set("blendSrc", v)
}
func (mm *MultiMaterial) BlendSrcAlpha() float64 {
	return mm.Get("blendSrcAlpha").Float()
}
func (mm *MultiMaterial) SetBlendSrcAlpha(v float64) {
	mm.Set("blendSrcAlpha", v)
}
func (mm *MultiMaterial) Blending() Blending {
	return Blending(mm.Get("blending"))
}
func (mm *MultiMaterial) SetBlending(v Blending) {
	mm.Set("blending", v)
}
func (mm *MultiMaterial) ClipIntersection() bool {
	return mm.Get("clipIntersection").Bool()
}
func (mm *MultiMaterial) SetClipIntersection(v bool) {
	mm.Set("clipIntersection", v)
}
func (mm *MultiMaterial) ClipShadows() bool {
	return mm.Get("clipShadows").Bool()
}
func (mm *MultiMaterial) SetClipShadows(v bool) {
	mm.Set("clipShadows", v)
}
func (mm *MultiMaterial) ClippingPlanes() js.Value {
	return mm.Get("clippingPlanes")
}
func (mm *MultiMaterial) SetClippingPlanes(v js.Value) {
	mm.Set("clippingPlanes", v)
}
func (mm *MultiMaterial) ColorWrite() bool {
	return mm.Get("colorWrite").Bool()
}
func (mm *MultiMaterial) SetColorWrite(v bool) {
	mm.Set("colorWrite", v)
}
func (mm *MultiMaterial) DepthFunc() DepthModes {
	return DepthModes(mm.Get("depthFunc"))
}
func (mm *MultiMaterial) SetDepthFunc(v DepthModes) {
	mm.Set("depthFunc", v)
}
func (mm *MultiMaterial) DepthTest() bool {
	return mm.Get("depthTest").Bool()
}
func (mm *MultiMaterial) SetDepthTest(v bool) {
	mm.Set("depthTest", v)
}
func (mm *MultiMaterial) DepthWrite() bool {
	return mm.Get("depthWrite").Bool()
}
func (mm *MultiMaterial) SetDepthWrite(v bool) {
	mm.Set("depthWrite", v)
}
func (mm *MultiMaterial) Dithering() bool {
	return mm.Get("dithering").Bool()
}
func (mm *MultiMaterial) SetDithering(v bool) {
	mm.Set("dithering", v)
}
func (mm *MultiMaterial) FlatShading() bool {
	return mm.Get("flatShading").Bool()
}
func (mm *MultiMaterial) SetFlatShading(v bool) {
	mm.Set("flatShading", v)
}
func (mm *MultiMaterial) Fog() bool {
	return mm.Get("fog").Bool()
}
func (mm *MultiMaterial) SetFog(v bool) {
	mm.Set("fog", v)
}
func (mm *MultiMaterial) Id() int {
	return mm.Get("id").Int()
}
func (mm *MultiMaterial) SetId(v int) {
	mm.Set("id", v)
}
func (mm *MultiMaterial) IsMaterial() bool {
	return mm.Get("isMaterial").Bool()
}
func (mm *MultiMaterial) SetIsMaterial(v bool) {
	mm.Set("isMaterial", v)
}
func (mm *MultiMaterial) IsMultiMaterial() bool {
	return mm.Get("isMultiMaterial").Bool()
}
func (mm *MultiMaterial) SetIsMultiMaterial(v bool) {
	mm.Set("isMultiMaterial", v)
}
func (mm *MultiMaterial) Lights() bool {
	return mm.Get("lights").Bool()
}
func (mm *MultiMaterial) SetLights(v bool) {
	mm.Set("lights", v)
}
func (mm *MultiMaterial) Materials() js.Value {
	return mm.Get("materials")
}
func (mm *MultiMaterial) SetMaterials(v js.Value) {
	mm.Set("materials", v)
}
func (mm *MultiMaterial) Name() string {
	return mm.Get("name").String()
}
func (mm *MultiMaterial) SetName(v string) {
	mm.Set("name", v)
}
func (mm *MultiMaterial) NeedsUpdate() bool {
	return mm.Get("needsUpdate").Bool()
}
func (mm *MultiMaterial) SetNeedsUpdate(v bool) {
	mm.Set("needsUpdate", v)
}
func (mm *MultiMaterial) Opacity() float64 {
	return mm.Get("opacity").Float()
}
func (mm *MultiMaterial) SetOpacity(v float64) {
	mm.Set("opacity", v)
}
func (mm *MultiMaterial) Overdraw() float64 {
	return mm.Get("overdraw").Float()
}
func (mm *MultiMaterial) SetOverdraw(v float64) {
	mm.Set("overdraw", v)
}
func (mm *MultiMaterial) PolygonOffset() bool {
	return mm.Get("polygonOffset").Bool()
}
func (mm *MultiMaterial) SetPolygonOffset(v bool) {
	mm.Set("polygonOffset", v)
}
func (mm *MultiMaterial) PolygonOffsetFactor() float64 {
	return mm.Get("polygonOffsetFactor").Float()
}
func (mm *MultiMaterial) SetPolygonOffsetFactor(v float64) {
	mm.Set("polygonOffsetFactor", v)
}
func (mm *MultiMaterial) PolygonOffsetUnits() float64 {
	return mm.Get("polygonOffsetUnits").Float()
}
func (mm *MultiMaterial) SetPolygonOffsetUnits(v float64) {
	mm.Set("polygonOffsetUnits", v)
}
func (mm *MultiMaterial) Precision() string {
	return mm.Get("precision").String()
}
func (mm *MultiMaterial) SetPrecision(v string) {
	mm.Set("precision", v)
}
func (mm *MultiMaterial) PremultipliedAlpha() bool {
	return mm.Get("premultipliedAlpha").Bool()
}
func (mm *MultiMaterial) SetPremultipliedAlpha(v bool) {
	mm.Set("premultipliedAlpha", v)
}
func (mm *MultiMaterial) Side() Side {
	return Side(mm.Get("side"))
}
func (mm *MultiMaterial) SetSide(v Side) {
	mm.Set("side", v)
}
func (mm *MultiMaterial) Transparent() bool {
	return mm.Get("transparent").Bool()
}
func (mm *MultiMaterial) SetTransparent(v bool) {
	mm.Set("transparent", v)
}
func (mm *MultiMaterial) Type() string {
	return mm.Get("type").String()
}
func (mm *MultiMaterial) SetType(v string) {
	mm.Set("type", v)
}
func (mm *MultiMaterial) UserData() js.Value {
	return mm.Get("userData")
}
func (mm *MultiMaterial) SetUserData(v js.Value) {
	mm.Set("userData", v)
}
func (mm *MultiMaterial) Uuid() string {
	return mm.Get("uuid").String()
}
func (mm *MultiMaterial) SetUuid(v string) {
	mm.Set("uuid", v)
}
func (mm *MultiMaterial) VertexColors() Colors {
	return Colors(mm.Get("vertexColors"))
}
func (mm *MultiMaterial) SetVertexColors(v Colors) {
	mm.Set("vertexColors", v)
}
func (mm *MultiMaterial) VertexTangents() bool {
	return mm.Get("vertexTangents").Bool()
}
func (mm *MultiMaterial) SetVertexTangents(v bool) {
	mm.Set("vertexTangents", v)
}
func (mm *MultiMaterial) Visible() bool {
	return mm.Get("visible").Bool()
}
func (mm *MultiMaterial) SetVisible(v bool) {
	mm.Set("visible", v)
}
func (mm *MultiMaterial) AddEventListener(typ string, listener js.Value) {
	mm.Call("addEventListener", typ, listener)
}
func (mm *MultiMaterial) Clone() *MultiMaterial {
	return &MultiMaterial{Value: mm.Call("clone")}
}
func (mm *MultiMaterial) Copy(material *Material) *MultiMaterial {
	return &MultiMaterial{Value: mm.Call("copy", material)}
}
func (mm *MultiMaterial) DispatchEvent(event js.Value) {
	mm.Call("dispatchEvent", event)
}
func (mm *MultiMaterial) Dispose() {
	mm.Call("dispose")
}
func (mm *MultiMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mm.Call("hasEventListener", typ, listener).Bool()
}
func (mm *MultiMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mm.Call("onBeforeCompile", shader, renderer)
}
func (mm *MultiMaterial) RemoveEventListener(typ string, listener js.Value) {
	mm.Call("removeEventListener", typ, listener)
}
func (mm *MultiMaterial) SetValues(values MaterialParameters) {
	mm.Call("setValues", values)
}
func (mm *MultiMaterial) ToJSON(meta js.Value) js.Value {
	return mm.Call("toJSON", meta)
}
func (mm *MultiMaterial) Update() {
	mm.Call("update")
}

type PointsMaterial struct {
	js.Value
}

func NewPointsMaterial(parameters PointsMaterialParameters) *PointsMaterial {
	return &PointsMaterial{Value: get("PointsMaterial").New(parameters)}
}
func (pm *PointsMaterial) AlphaTest() float64 {
	return pm.Get("alphaTest").Float()
}
func (pm *PointsMaterial) SetAlphaTest(v float64) {
	pm.Set("alphaTest", v)
}
func (pm *PointsMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(pm.Get("blendDst"))
}
func (pm *PointsMaterial) SetBlendDst(v BlendingDstFactor) {
	pm.Set("blendDst", v)
}
func (pm *PointsMaterial) BlendDstAlpha() float64 {
	return pm.Get("blendDstAlpha").Float()
}
func (pm *PointsMaterial) SetBlendDstAlpha(v float64) {
	pm.Set("blendDstAlpha", v)
}
func (pm *PointsMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(pm.Get("blendEquation"))
}
func (pm *PointsMaterial) SetBlendEquation(v BlendingEquation) {
	pm.Set("blendEquation", v)
}
func (pm *PointsMaterial) BlendEquationAlpha() float64 {
	return pm.Get("blendEquationAlpha").Float()
}
func (pm *PointsMaterial) SetBlendEquationAlpha(v float64) {
	pm.Set("blendEquationAlpha", v)
}
func (pm *PointsMaterial) BlendSrc() js.Value {
	return pm.Get("blendSrc")
}
func (pm *PointsMaterial) SetBlendSrc(v js.Value) {
	pm.Set("blendSrc", v)
}
func (pm *PointsMaterial) BlendSrcAlpha() float64 {
	return pm.Get("blendSrcAlpha").Float()
}
func (pm *PointsMaterial) SetBlendSrcAlpha(v float64) {
	pm.Set("blendSrcAlpha", v)
}
func (pm *PointsMaterial) Blending() Blending {
	return Blending(pm.Get("blending"))
}
func (pm *PointsMaterial) SetBlending(v Blending) {
	pm.Set("blending", v)
}
func (pm *PointsMaterial) ClipIntersection() bool {
	return pm.Get("clipIntersection").Bool()
}
func (pm *PointsMaterial) SetClipIntersection(v bool) {
	pm.Set("clipIntersection", v)
}
func (pm *PointsMaterial) ClipShadows() bool {
	return pm.Get("clipShadows").Bool()
}
func (pm *PointsMaterial) SetClipShadows(v bool) {
	pm.Set("clipShadows", v)
}
func (pm *PointsMaterial) ClippingPlanes() js.Value {
	return pm.Get("clippingPlanes")
}
func (pm *PointsMaterial) SetClippingPlanes(v js.Value) {
	pm.Set("clippingPlanes", v)
}
func (pm *PointsMaterial) Color() *Color {
	return &Color{Value: pm.Get("color")}
}
func (pm *PointsMaterial) SetColor(v *Color) {
	pm.Set("color", v)
}
func (pm *PointsMaterial) ColorWrite() bool {
	return pm.Get("colorWrite").Bool()
}
func (pm *PointsMaterial) SetColorWrite(v bool) {
	pm.Set("colorWrite", v)
}
func (pm *PointsMaterial) DepthFunc() DepthModes {
	return DepthModes(pm.Get("depthFunc"))
}
func (pm *PointsMaterial) SetDepthFunc(v DepthModes) {
	pm.Set("depthFunc", v)
}
func (pm *PointsMaterial) DepthTest() bool {
	return pm.Get("depthTest").Bool()
}
func (pm *PointsMaterial) SetDepthTest(v bool) {
	pm.Set("depthTest", v)
}
func (pm *PointsMaterial) DepthWrite() bool {
	return pm.Get("depthWrite").Bool()
}
func (pm *PointsMaterial) SetDepthWrite(v bool) {
	pm.Set("depthWrite", v)
}
func (pm *PointsMaterial) Dithering() bool {
	return pm.Get("dithering").Bool()
}
func (pm *PointsMaterial) SetDithering(v bool) {
	pm.Set("dithering", v)
}
func (pm *PointsMaterial) FlatShading() bool {
	return pm.Get("flatShading").Bool()
}
func (pm *PointsMaterial) SetFlatShading(v bool) {
	pm.Set("flatShading", v)
}
func (pm *PointsMaterial) Fog() bool {
	return pm.Get("fog").Bool()
}
func (pm *PointsMaterial) SetFog(v bool) {
	pm.Set("fog", v)
}
func (pm *PointsMaterial) Id() int {
	return pm.Get("id").Int()
}
func (pm *PointsMaterial) SetId(v int) {
	pm.Set("id", v)
}
func (pm *PointsMaterial) IsMaterial() bool {
	return pm.Get("isMaterial").Bool()
}
func (pm *PointsMaterial) SetIsMaterial(v bool) {
	pm.Set("isMaterial", v)
}
func (pm *PointsMaterial) Lights() bool {
	return pm.Get("lights").Bool()
}
func (pm *PointsMaterial) SetLights(v bool) {
	pm.Set("lights", v)
}
func (pm *PointsMaterial) Map() *Texture {
	return &Texture{Value: pm.Get("map")}
}
func (pm *PointsMaterial) SetMap(v *Texture) {
	pm.Set("map", v)
}
func (pm *PointsMaterial) Name() string {
	return pm.Get("name").String()
}
func (pm *PointsMaterial) SetName(v string) {
	pm.Set("name", v)
}
func (pm *PointsMaterial) NeedsUpdate() bool {
	return pm.Get("needsUpdate").Bool()
}
func (pm *PointsMaterial) SetNeedsUpdate(v bool) {
	pm.Set("needsUpdate", v)
}
func (pm *PointsMaterial) Opacity() float64 {
	return pm.Get("opacity").Float()
}
func (pm *PointsMaterial) SetOpacity(v float64) {
	pm.Set("opacity", v)
}
func (pm *PointsMaterial) Overdraw() float64 {
	return pm.Get("overdraw").Float()
}
func (pm *PointsMaterial) SetOverdraw(v float64) {
	pm.Set("overdraw", v)
}
func (pm *PointsMaterial) PolygonOffset() bool {
	return pm.Get("polygonOffset").Bool()
}
func (pm *PointsMaterial) SetPolygonOffset(v bool) {
	pm.Set("polygonOffset", v)
}
func (pm *PointsMaterial) PolygonOffsetFactor() float64 {
	return pm.Get("polygonOffsetFactor").Float()
}
func (pm *PointsMaterial) SetPolygonOffsetFactor(v float64) {
	pm.Set("polygonOffsetFactor", v)
}
func (pm *PointsMaterial) PolygonOffsetUnits() float64 {
	return pm.Get("polygonOffsetUnits").Float()
}
func (pm *PointsMaterial) SetPolygonOffsetUnits(v float64) {
	pm.Set("polygonOffsetUnits", v)
}
func (pm *PointsMaterial) Precision() string {
	return pm.Get("precision").String()
}
func (pm *PointsMaterial) SetPrecision(v string) {
	pm.Set("precision", v)
}
func (pm *PointsMaterial) PremultipliedAlpha() bool {
	return pm.Get("premultipliedAlpha").Bool()
}
func (pm *PointsMaterial) SetPremultipliedAlpha(v bool) {
	pm.Set("premultipliedAlpha", v)
}
func (pm *PointsMaterial) Side() Side {
	return Side(pm.Get("side"))
}
func (pm *PointsMaterial) SetSide(v Side) {
	pm.Set("side", v)
}
func (pm *PointsMaterial) Size() float64 {
	return pm.Get("size").Float()
}
func (pm *PointsMaterial) SetSize(v float64) {
	pm.Set("size", v)
}
func (pm *PointsMaterial) SizeAttenuation() bool {
	return pm.Get("sizeAttenuation").Bool()
}
func (pm *PointsMaterial) SetSizeAttenuation(v bool) {
	pm.Set("sizeAttenuation", v)
}
func (pm *PointsMaterial) Transparent() bool {
	return pm.Get("transparent").Bool()
}
func (pm *PointsMaterial) SetTransparent(v bool) {
	pm.Set("transparent", v)
}
func (pm *PointsMaterial) Type() string {
	return pm.Get("type").String()
}
func (pm *PointsMaterial) SetType(v string) {
	pm.Set("type", v)
}
func (pm *PointsMaterial) UserData() js.Value {
	return pm.Get("userData")
}
func (pm *PointsMaterial) SetUserData(v js.Value) {
	pm.Set("userData", v)
}
func (pm *PointsMaterial) Uuid() string {
	return pm.Get("uuid").String()
}
func (pm *PointsMaterial) SetUuid(v string) {
	pm.Set("uuid", v)
}
func (pm *PointsMaterial) VertexColors() Colors {
	return Colors(pm.Get("vertexColors"))
}
func (pm *PointsMaterial) SetVertexColors(v Colors) {
	pm.Set("vertexColors", v)
}
func (pm *PointsMaterial) VertexTangents() bool {
	return pm.Get("vertexTangents").Bool()
}
func (pm *PointsMaterial) SetVertexTangents(v bool) {
	pm.Set("vertexTangents", v)
}
func (pm *PointsMaterial) Visible() bool {
	return pm.Get("visible").Bool()
}
func (pm *PointsMaterial) SetVisible(v bool) {
	pm.Set("visible", v)
}
func (pm *PointsMaterial) AddEventListener(typ string, listener js.Value) {
	pm.Call("addEventListener", typ, listener)
}
func (pm *PointsMaterial) Clone() *PointsMaterial {
	return &PointsMaterial{Value: pm.Call("clone")}
}
func (pm *PointsMaterial) Copy(material *Material) *PointsMaterial {
	return &PointsMaterial{Value: pm.Call("copy", material)}
}
func (pm *PointsMaterial) DispatchEvent(event js.Value) {
	pm.Call("dispatchEvent", event)
}
func (pm *PointsMaterial) Dispose() {
	pm.Call("dispose")
}
func (pm *PointsMaterial) HasEventListener(typ string, listener js.Value) bool {
	return pm.Call("hasEventListener", typ, listener).Bool()
}
func (pm *PointsMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	pm.Call("onBeforeCompile", shader, renderer)
}
func (pm *PointsMaterial) RemoveEventListener(typ string, listener js.Value) {
	pm.Call("removeEventListener", typ, listener)
}
func (pm *PointsMaterial) SetValues(parameters PointsMaterialParameters) {
	pm.Call("setValues", parameters)
}
func (pm *PointsMaterial) ToJSON(meta js.Value) js.Value {
	return pm.Call("toJSON", meta)
}
func (pm *PointsMaterial) Update() {
	pm.Call("update")
}

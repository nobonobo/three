// Code generated from "materials/ShadowMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// ShadowMaterial extend: [ShaderMaterial]
type ShadowMaterial struct {
	js.Value
}

func NewShadowMaterial(parameters ShaderMaterialParameters) *ShadowMaterial {
	return &ShadowMaterial{Value: get("ShadowMaterial").New(parameters)}
}
func (sm *ShadowMaterial) JSValue() js.Value {
	return sm.Value
}
func (sm *ShadowMaterial) AlphaTest() float64 {
	return sm.Get("alphaTest").Float()
}
func (sm *ShadowMaterial) SetAlphaTest(v float64) {
	sm.Set("alphaTest", v)
}
func (sm *ShadowMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(sm.Get("blendDst"))
}
func (sm *ShadowMaterial) SetBlendDst(v BlendingDstFactor) {
	sm.Set("blendDst", v)
}
func (sm *ShadowMaterial) BlendDstAlpha() float64 {
	return sm.Get("blendDstAlpha").Float()
}
func (sm *ShadowMaterial) SetBlendDstAlpha(v float64) {
	sm.Set("blendDstAlpha", v)
}
func (sm *ShadowMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(sm.Get("blendEquation"))
}
func (sm *ShadowMaterial) SetBlendEquation(v BlendingEquation) {
	sm.Set("blendEquation", v)
}
func (sm *ShadowMaterial) BlendEquationAlpha() float64 {
	return sm.Get("blendEquationAlpha").Float()
}
func (sm *ShadowMaterial) SetBlendEquationAlpha(v float64) {
	sm.Set("blendEquationAlpha", v)
}
func (sm *ShadowMaterial) BlendSrc() js.Value {
	return sm.Get("blendSrc")
}
func (sm *ShadowMaterial) SetBlendSrc(v js.Value) {
	sm.Set("blendSrc", v)
}
func (sm *ShadowMaterial) BlendSrcAlpha() float64 {
	return sm.Get("blendSrcAlpha").Float()
}
func (sm *ShadowMaterial) SetBlendSrcAlpha(v float64) {
	sm.Set("blendSrcAlpha", v)
}
func (sm *ShadowMaterial) Blending() Blending {
	return Blending(sm.Get("blending"))
}
func (sm *ShadowMaterial) SetBlending(v Blending) {
	sm.Set("blending", v)
}
func (sm *ShadowMaterial) ClipIntersection() bool {
	return sm.Get("clipIntersection").Bool()
}
func (sm *ShadowMaterial) SetClipIntersection(v bool) {
	sm.Set("clipIntersection", v)
}
func (sm *ShadowMaterial) ClipShadows() bool {
	return sm.Get("clipShadows").Bool()
}
func (sm *ShadowMaterial) SetClipShadows(v bool) {
	sm.Set("clipShadows", v)
}
func (sm *ShadowMaterial) Clipping() bool {
	return sm.Get("clipping").Bool()
}
func (sm *ShadowMaterial) SetClipping(v bool) {
	sm.Set("clipping", v)
}
func (sm *ShadowMaterial) ClippingPlanes() js.Value {
	return sm.Get("clippingPlanes")
}
func (sm *ShadowMaterial) SetClippingPlanes(v js.Value) {
	sm.Set("clippingPlanes", v)
}
func (sm *ShadowMaterial) ColorWrite() bool {
	return sm.Get("colorWrite").Bool()
}
func (sm *ShadowMaterial) SetColorWrite(v bool) {
	sm.Set("colorWrite", v)
}
func (sm *ShadowMaterial) DefaultAttributeValues() js.Value {
	return sm.Get("defaultAttributeValues")
}
func (sm *ShadowMaterial) SetDefaultAttributeValues(v js.Value) {
	sm.Set("defaultAttributeValues", v)
}
func (sm *ShadowMaterial) Defines() js.Value {
	return sm.Get("defines")
}
func (sm *ShadowMaterial) SetDefines(v js.Value) {
	sm.Set("defines", v)
}
func (sm *ShadowMaterial) DepthFunc() DepthModes {
	return DepthModes(sm.Get("depthFunc"))
}
func (sm *ShadowMaterial) SetDepthFunc(v DepthModes) {
	sm.Set("depthFunc", v)
}
func (sm *ShadowMaterial) DepthTest() bool {
	return sm.Get("depthTest").Bool()
}
func (sm *ShadowMaterial) SetDepthTest(v bool) {
	sm.Set("depthTest", v)
}
func (sm *ShadowMaterial) DepthWrite() bool {
	return sm.Get("depthWrite").Bool()
}
func (sm *ShadowMaterial) SetDepthWrite(v bool) {
	sm.Set("depthWrite", v)
}
func (sm *ShadowMaterial) Derivatives() js.Value {
	return sm.Get("derivatives")
}
func (sm *ShadowMaterial) SetDerivatives(v js.Value) {
	sm.Set("derivatives", v)
}
func (sm *ShadowMaterial) Dithering() bool {
	return sm.Get("dithering").Bool()
}
func (sm *ShadowMaterial) SetDithering(v bool) {
	sm.Set("dithering", v)
}
func (sm *ShadowMaterial) Extensions() js.Value {
	return sm.Get("extensions")
}
func (sm *ShadowMaterial) SetExtensions(v js.Value) {
	sm.Set("extensions", v)
}
func (sm *ShadowMaterial) FlatShading() bool {
	return sm.Get("flatShading").Bool()
}
func (sm *ShadowMaterial) SetFlatShading(v bool) {
	sm.Set("flatShading", v)
}
func (sm *ShadowMaterial) Fog() bool {
	return sm.Get("fog").Bool()
}
func (sm *ShadowMaterial) SetFog(v bool) {
	sm.Set("fog", v)
}
func (sm *ShadowMaterial) FragmentShader() string {
	return sm.Get("fragmentShader").String()
}
func (sm *ShadowMaterial) SetFragmentShader(v string) {
	sm.Set("fragmentShader", v)
}
func (sm *ShadowMaterial) Id() int {
	return sm.Get("id").Int()
}
func (sm *ShadowMaterial) SetId(v int) {
	sm.Set("id", v)
}
func (sm *ShadowMaterial) Index0AttributeName() string {
	return sm.Get("index0AttributeName").String()
}
func (sm *ShadowMaterial) SetIndex0AttributeName(v string) {
	sm.Set("index0AttributeName", v)
}
func (sm *ShadowMaterial) IsMaterial() bool {
	return sm.Get("isMaterial").Bool()
}
func (sm *ShadowMaterial) SetIsMaterial(v bool) {
	sm.Set("isMaterial", v)
}
func (sm *ShadowMaterial) Lights() bool {
	return sm.Get("lights").Bool()
}
func (sm *ShadowMaterial) SetLights(v bool) {
	sm.Set("lights", v)
}
func (sm *ShadowMaterial) Linewidth() float64 {
	return sm.Get("linewidth").Float()
}
func (sm *ShadowMaterial) SetLinewidth(v float64) {
	sm.Set("linewidth", v)
}
func (sm *ShadowMaterial) MorphNormals() bool {
	return sm.Get("morphNormals").Bool()
}
func (sm *ShadowMaterial) SetMorphNormals(v bool) {
	sm.Set("morphNormals", v)
}
func (sm *ShadowMaterial) MorphTargets() bool {
	return sm.Get("morphTargets").Bool()
}
func (sm *ShadowMaterial) SetMorphTargets(v bool) {
	sm.Set("morphTargets", v)
}
func (sm *ShadowMaterial) Name() string {
	return sm.Get("name").String()
}
func (sm *ShadowMaterial) SetName(v string) {
	sm.Set("name", v)
}
func (sm *ShadowMaterial) NeedsUpdate() bool {
	return sm.Get("needsUpdate").Bool()
}
func (sm *ShadowMaterial) SetNeedsUpdate(v bool) {
	sm.Set("needsUpdate", v)
}
func (sm *ShadowMaterial) Opacity() float64 {
	return sm.Get("opacity").Float()
}
func (sm *ShadowMaterial) SetOpacity(v float64) {
	sm.Set("opacity", v)
}
func (sm *ShadowMaterial) Overdraw() float64 {
	return sm.Get("overdraw").Float()
}
func (sm *ShadowMaterial) SetOverdraw(v float64) {
	sm.Set("overdraw", v)
}
func (sm *ShadowMaterial) PolygonOffset() bool {
	return sm.Get("polygonOffset").Bool()
}
func (sm *ShadowMaterial) SetPolygonOffset(v bool) {
	sm.Set("polygonOffset", v)
}
func (sm *ShadowMaterial) PolygonOffsetFactor() float64 {
	return sm.Get("polygonOffsetFactor").Float()
}
func (sm *ShadowMaterial) SetPolygonOffsetFactor(v float64) {
	sm.Set("polygonOffsetFactor", v)
}
func (sm *ShadowMaterial) PolygonOffsetUnits() float64 {
	return sm.Get("polygonOffsetUnits").Float()
}
func (sm *ShadowMaterial) SetPolygonOffsetUnits(v float64) {
	sm.Set("polygonOffsetUnits", v)
}
func (sm *ShadowMaterial) Precision() string {
	return sm.Get("precision").String()
}
func (sm *ShadowMaterial) SetPrecision(v string) {
	sm.Set("precision", v)
}
func (sm *ShadowMaterial) PremultipliedAlpha() bool {
	return sm.Get("premultipliedAlpha").Bool()
}
func (sm *ShadowMaterial) SetPremultipliedAlpha(v bool) {
	sm.Set("premultipliedAlpha", v)
}
func (sm *ShadowMaterial) Side() Side {
	return Side(sm.Get("side"))
}
func (sm *ShadowMaterial) SetSide(v Side) {
	sm.Set("side", v)
}
func (sm *ShadowMaterial) Skinning() bool {
	return sm.Get("skinning").Bool()
}
func (sm *ShadowMaterial) SetSkinning(v bool) {
	sm.Set("skinning", v)
}
func (sm *ShadowMaterial) Transparent() bool {
	return sm.Get("transparent").Bool()
}
func (sm *ShadowMaterial) SetTransparent(v bool) {
	sm.Set("transparent", v)
}
func (sm *ShadowMaterial) Type() string {
	return sm.Get("type").String()
}
func (sm *ShadowMaterial) SetType(v string) {
	sm.Set("type", v)
}
func (sm *ShadowMaterial) Uniforms() js.Value {
	return sm.Get("uniforms")
}
func (sm *ShadowMaterial) SetUniforms(v js.Value) {
	sm.Set("uniforms", v)
}
func (sm *ShadowMaterial) UserData() js.Value {
	return sm.Get("userData")
}
func (sm *ShadowMaterial) SetUserData(v js.Value) {
	sm.Set("userData", v)
}
func (sm *ShadowMaterial) Uuid() string {
	return sm.Get("uuid").String()
}
func (sm *ShadowMaterial) SetUuid(v string) {
	sm.Set("uuid", v)
}
func (sm *ShadowMaterial) VertexColors() Colors {
	return Colors(sm.Get("vertexColors"))
}
func (sm *ShadowMaterial) SetVertexColors(v Colors) {
	sm.Set("vertexColors", v)
}
func (sm *ShadowMaterial) VertexShader() string {
	return sm.Get("vertexShader").String()
}
func (sm *ShadowMaterial) SetVertexShader(v string) {
	sm.Set("vertexShader", v)
}
func (sm *ShadowMaterial) VertexTangents() bool {
	return sm.Get("vertexTangents").Bool()
}
func (sm *ShadowMaterial) SetVertexTangents(v bool) {
	sm.Set("vertexTangents", v)
}
func (sm *ShadowMaterial) Visible() bool {
	return sm.Get("visible").Bool()
}
func (sm *ShadowMaterial) SetVisible(v bool) {
	sm.Set("visible", v)
}
func (sm *ShadowMaterial) Wireframe() bool {
	return sm.Get("wireframe").Bool()
}
func (sm *ShadowMaterial) SetWireframe(v bool) {
	sm.Set("wireframe", v)
}
func (sm *ShadowMaterial) WireframeLinewidth() float64 {
	return sm.Get("wireframeLinewidth").Float()
}
func (sm *ShadowMaterial) SetWireframeLinewidth(v float64) {
	sm.Set("wireframeLinewidth", v)
}
func (sm *ShadowMaterial) AddEventListener(typ string, listener js.Value) {
	sm.Call("addEventListener", typ, listener)
}
func (sm *ShadowMaterial) Clone() *ShadowMaterial {
	return &ShadowMaterial{Value: sm.Call("clone")}
}
func (sm *ShadowMaterial) Copy(material Material) *ShadowMaterial {
	return &ShadowMaterial{Value: sm.Call("copy", material.JSValue())}
}
func (sm *ShadowMaterial) DispatchEvent(event js.Value) {
	sm.Call("dispatchEvent", event)
}
func (sm *ShadowMaterial) Dispose() {
	sm.Call("dispose")
}
func (sm *ShadowMaterial) HasEventListener(typ string, listener js.Value) bool {
	return sm.Call("hasEventListener", typ, listener).Bool()
}
func (sm *ShadowMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	sm.Call("onBeforeCompile", shader, renderer)
}
func (sm *ShadowMaterial) RemoveEventListener(typ string, listener js.Value) {
	sm.Call("removeEventListener", typ, listener)
}
func (sm *ShadowMaterial) SetValues(parameters ShaderMaterialParameters) {
	sm.Call("setValues", parameters)
}
func (sm *ShadowMaterial) ToJSON(meta js.Value) js.Value {
	return sm.Call("toJSON", meta)
}
func (sm *ShadowMaterial) Update() {
	sm.Call("update")
}

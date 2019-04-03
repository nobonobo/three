// Code generated from "materials/RawShaderMaterial.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type RawShaderMaterial struct {
	js.Value
}

func NewRawShaderMaterial(parameters ShaderMaterialParameters) *RawShaderMaterial {
	return &RawShaderMaterial{Value: get("RawShaderMaterial").New(parameters)}
}
func (rsm *RawShaderMaterial) AlphaTest() float64 {
	return rsm.Get("alphaTest").Float()
}
func (rsm *RawShaderMaterial) SetAlphaTest(v float64) {
	rsm.Set("alphaTest", v)
}
func (rsm *RawShaderMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(rsm.Get("blendDst"))
}
func (rsm *RawShaderMaterial) SetBlendDst(v BlendingDstFactor) {
	rsm.Set("blendDst", v)
}
func (rsm *RawShaderMaterial) BlendDstAlpha() float64 {
	return rsm.Get("blendDstAlpha").Float()
}
func (rsm *RawShaderMaterial) SetBlendDstAlpha(v float64) {
	rsm.Set("blendDstAlpha", v)
}
func (rsm *RawShaderMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(rsm.Get("blendEquation"))
}
func (rsm *RawShaderMaterial) SetBlendEquation(v BlendingEquation) {
	rsm.Set("blendEquation", v)
}
func (rsm *RawShaderMaterial) BlendEquationAlpha() float64 {
	return rsm.Get("blendEquationAlpha").Float()
}
func (rsm *RawShaderMaterial) SetBlendEquationAlpha(v float64) {
	rsm.Set("blendEquationAlpha", v)
}
func (rsm *RawShaderMaterial) BlendSrc() js.Value {
	return rsm.Get("blendSrc")
}
func (rsm *RawShaderMaterial) SetBlendSrc(v js.Value) {
	rsm.Set("blendSrc", v)
}
func (rsm *RawShaderMaterial) BlendSrcAlpha() float64 {
	return rsm.Get("blendSrcAlpha").Float()
}
func (rsm *RawShaderMaterial) SetBlendSrcAlpha(v float64) {
	rsm.Set("blendSrcAlpha", v)
}
func (rsm *RawShaderMaterial) Blending() Blending {
	return Blending(rsm.Get("blending"))
}
func (rsm *RawShaderMaterial) SetBlending(v Blending) {
	rsm.Set("blending", v)
}
func (rsm *RawShaderMaterial) ClipIntersection() bool {
	return rsm.Get("clipIntersection").Bool()
}
func (rsm *RawShaderMaterial) SetClipIntersection(v bool) {
	rsm.Set("clipIntersection", v)
}
func (rsm *RawShaderMaterial) ClipShadows() bool {
	return rsm.Get("clipShadows").Bool()
}
func (rsm *RawShaderMaterial) SetClipShadows(v bool) {
	rsm.Set("clipShadows", v)
}
func (rsm *RawShaderMaterial) Clipping() bool {
	return rsm.Get("clipping").Bool()
}
func (rsm *RawShaderMaterial) SetClipping(v bool) {
	rsm.Set("clipping", v)
}
func (rsm *RawShaderMaterial) ClippingPlanes() js.Value {
	return rsm.Get("clippingPlanes")
}
func (rsm *RawShaderMaterial) SetClippingPlanes(v js.Value) {
	rsm.Set("clippingPlanes", v)
}
func (rsm *RawShaderMaterial) ColorWrite() bool {
	return rsm.Get("colorWrite").Bool()
}
func (rsm *RawShaderMaterial) SetColorWrite(v bool) {
	rsm.Set("colorWrite", v)
}
func (rsm *RawShaderMaterial) DefaultAttributeValues() js.Value {
	return rsm.Get("defaultAttributeValues")
}
func (rsm *RawShaderMaterial) SetDefaultAttributeValues(v js.Value) {
	rsm.Set("defaultAttributeValues", v)
}
func (rsm *RawShaderMaterial) Defines() js.Value {
	return rsm.Get("defines")
}
func (rsm *RawShaderMaterial) SetDefines(v js.Value) {
	rsm.Set("defines", v)
}
func (rsm *RawShaderMaterial) DepthFunc() DepthModes {
	return DepthModes(rsm.Get("depthFunc"))
}
func (rsm *RawShaderMaterial) SetDepthFunc(v DepthModes) {
	rsm.Set("depthFunc", v)
}
func (rsm *RawShaderMaterial) DepthTest() bool {
	return rsm.Get("depthTest").Bool()
}
func (rsm *RawShaderMaterial) SetDepthTest(v bool) {
	rsm.Set("depthTest", v)
}
func (rsm *RawShaderMaterial) DepthWrite() bool {
	return rsm.Get("depthWrite").Bool()
}
func (rsm *RawShaderMaterial) SetDepthWrite(v bool) {
	rsm.Set("depthWrite", v)
}
func (rsm *RawShaderMaterial) Derivatives() js.Value {
	return rsm.Get("derivatives")
}
func (rsm *RawShaderMaterial) SetDerivatives(v js.Value) {
	rsm.Set("derivatives", v)
}
func (rsm *RawShaderMaterial) Dithering() bool {
	return rsm.Get("dithering").Bool()
}
func (rsm *RawShaderMaterial) SetDithering(v bool) {
	rsm.Set("dithering", v)
}
func (rsm *RawShaderMaterial) Extensions() js.Value {
	return rsm.Get("extensions")
}
func (rsm *RawShaderMaterial) SetExtensions(v js.Value) {
	rsm.Set("extensions", v)
}
func (rsm *RawShaderMaterial) FlatShading() bool {
	return rsm.Get("flatShading").Bool()
}
func (rsm *RawShaderMaterial) SetFlatShading(v bool) {
	rsm.Set("flatShading", v)
}
func (rsm *RawShaderMaterial) Fog() bool {
	return rsm.Get("fog").Bool()
}
func (rsm *RawShaderMaterial) SetFog(v bool) {
	rsm.Set("fog", v)
}
func (rsm *RawShaderMaterial) FragmentShader() string {
	return rsm.Get("fragmentShader").String()
}
func (rsm *RawShaderMaterial) SetFragmentShader(v string) {
	rsm.Set("fragmentShader", v)
}
func (rsm *RawShaderMaterial) Id() int {
	return rsm.Get("id").Int()
}
func (rsm *RawShaderMaterial) SetId(v int) {
	rsm.Set("id", v)
}
func (rsm *RawShaderMaterial) Index0AttributeName() string {
	return rsm.Get("index0AttributeName").String()
}
func (rsm *RawShaderMaterial) SetIndex0AttributeName(v string) {
	rsm.Set("index0AttributeName", v)
}
func (rsm *RawShaderMaterial) IsMaterial() bool {
	return rsm.Get("isMaterial").Bool()
}
func (rsm *RawShaderMaterial) SetIsMaterial(v bool) {
	rsm.Set("isMaterial", v)
}
func (rsm *RawShaderMaterial) Lights() bool {
	return rsm.Get("lights").Bool()
}
func (rsm *RawShaderMaterial) SetLights(v bool) {
	rsm.Set("lights", v)
}
func (rsm *RawShaderMaterial) Linewidth() float64 {
	return rsm.Get("linewidth").Float()
}
func (rsm *RawShaderMaterial) SetLinewidth(v float64) {
	rsm.Set("linewidth", v)
}
func (rsm *RawShaderMaterial) MorphNormals() bool {
	return rsm.Get("morphNormals").Bool()
}
func (rsm *RawShaderMaterial) SetMorphNormals(v bool) {
	rsm.Set("morphNormals", v)
}
func (rsm *RawShaderMaterial) MorphTargets() bool {
	return rsm.Get("morphTargets").Bool()
}
func (rsm *RawShaderMaterial) SetMorphTargets(v bool) {
	rsm.Set("morphTargets", v)
}
func (rsm *RawShaderMaterial) Name() string {
	return rsm.Get("name").String()
}
func (rsm *RawShaderMaterial) SetName(v string) {
	rsm.Set("name", v)
}
func (rsm *RawShaderMaterial) NeedsUpdate() bool {
	return rsm.Get("needsUpdate").Bool()
}
func (rsm *RawShaderMaterial) SetNeedsUpdate(v bool) {
	rsm.Set("needsUpdate", v)
}
func (rsm *RawShaderMaterial) Opacity() float64 {
	return rsm.Get("opacity").Float()
}
func (rsm *RawShaderMaterial) SetOpacity(v float64) {
	rsm.Set("opacity", v)
}
func (rsm *RawShaderMaterial) Overdraw() float64 {
	return rsm.Get("overdraw").Float()
}
func (rsm *RawShaderMaterial) SetOverdraw(v float64) {
	rsm.Set("overdraw", v)
}
func (rsm *RawShaderMaterial) PolygonOffset() bool {
	return rsm.Get("polygonOffset").Bool()
}
func (rsm *RawShaderMaterial) SetPolygonOffset(v bool) {
	rsm.Set("polygonOffset", v)
}
func (rsm *RawShaderMaterial) PolygonOffsetFactor() float64 {
	return rsm.Get("polygonOffsetFactor").Float()
}
func (rsm *RawShaderMaterial) SetPolygonOffsetFactor(v float64) {
	rsm.Set("polygonOffsetFactor", v)
}
func (rsm *RawShaderMaterial) PolygonOffsetUnits() float64 {
	return rsm.Get("polygonOffsetUnits").Float()
}
func (rsm *RawShaderMaterial) SetPolygonOffsetUnits(v float64) {
	rsm.Set("polygonOffsetUnits", v)
}
func (rsm *RawShaderMaterial) Precision() string {
	return rsm.Get("precision").String()
}
func (rsm *RawShaderMaterial) SetPrecision(v string) {
	rsm.Set("precision", v)
}
func (rsm *RawShaderMaterial) PremultipliedAlpha() bool {
	return rsm.Get("premultipliedAlpha").Bool()
}
func (rsm *RawShaderMaterial) SetPremultipliedAlpha(v bool) {
	rsm.Set("premultipliedAlpha", v)
}
func (rsm *RawShaderMaterial) Side() Side {
	return Side(rsm.Get("side"))
}
func (rsm *RawShaderMaterial) SetSide(v Side) {
	rsm.Set("side", v)
}
func (rsm *RawShaderMaterial) Skinning() bool {
	return rsm.Get("skinning").Bool()
}
func (rsm *RawShaderMaterial) SetSkinning(v bool) {
	rsm.Set("skinning", v)
}
func (rsm *RawShaderMaterial) Transparent() bool {
	return rsm.Get("transparent").Bool()
}
func (rsm *RawShaderMaterial) SetTransparent(v bool) {
	rsm.Set("transparent", v)
}
func (rsm *RawShaderMaterial) Type() string {
	return rsm.Get("type").String()
}
func (rsm *RawShaderMaterial) SetType(v string) {
	rsm.Set("type", v)
}
func (rsm *RawShaderMaterial) Uniforms() js.Value {
	return rsm.Get("uniforms")
}
func (rsm *RawShaderMaterial) SetUniforms(v js.Value) {
	rsm.Set("uniforms", v)
}
func (rsm *RawShaderMaterial) UserData() js.Value {
	return rsm.Get("userData")
}
func (rsm *RawShaderMaterial) SetUserData(v js.Value) {
	rsm.Set("userData", v)
}
func (rsm *RawShaderMaterial) Uuid() string {
	return rsm.Get("uuid").String()
}
func (rsm *RawShaderMaterial) SetUuid(v string) {
	rsm.Set("uuid", v)
}
func (rsm *RawShaderMaterial) VertexColors() Colors {
	return Colors(rsm.Get("vertexColors"))
}
func (rsm *RawShaderMaterial) SetVertexColors(v Colors) {
	rsm.Set("vertexColors", v)
}
func (rsm *RawShaderMaterial) VertexShader() string {
	return rsm.Get("vertexShader").String()
}
func (rsm *RawShaderMaterial) SetVertexShader(v string) {
	rsm.Set("vertexShader", v)
}
func (rsm *RawShaderMaterial) VertexTangents() bool {
	return rsm.Get("vertexTangents").Bool()
}
func (rsm *RawShaderMaterial) SetVertexTangents(v bool) {
	rsm.Set("vertexTangents", v)
}
func (rsm *RawShaderMaterial) Visible() bool {
	return rsm.Get("visible").Bool()
}
func (rsm *RawShaderMaterial) SetVisible(v bool) {
	rsm.Set("visible", v)
}
func (rsm *RawShaderMaterial) Wireframe() bool {
	return rsm.Get("wireframe").Bool()
}
func (rsm *RawShaderMaterial) SetWireframe(v bool) {
	rsm.Set("wireframe", v)
}
func (rsm *RawShaderMaterial) WireframeLinewidth() float64 {
	return rsm.Get("wireframeLinewidth").Float()
}
func (rsm *RawShaderMaterial) SetWireframeLinewidth(v float64) {
	rsm.Set("wireframeLinewidth", v)
}
func (rsm *RawShaderMaterial) AddEventListener(typ string, listener js.Value) {
	rsm.Call("addEventListener", typ, listener)
}
func (rsm *RawShaderMaterial) Clone() *RawShaderMaterial {
	return &RawShaderMaterial{Value: rsm.Call("clone")}
}
func (rsm *RawShaderMaterial) Copy(material *Material) *RawShaderMaterial {
	return &RawShaderMaterial{Value: rsm.Call("copy", material)}
}
func (rsm *RawShaderMaterial) DispatchEvent(event js.Value) {
	rsm.Call("dispatchEvent", event)
}
func (rsm *RawShaderMaterial) Dispose() {
	rsm.Call("dispose")
}
func (rsm *RawShaderMaterial) HasEventListener(typ string, listener js.Value) bool {
	return rsm.Call("hasEventListener", typ, listener).Bool()
}
func (rsm *RawShaderMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	rsm.Call("onBeforeCompile", shader, renderer)
}
func (rsm *RawShaderMaterial) RemoveEventListener(typ string, listener js.Value) {
	rsm.Call("removeEventListener", typ, listener)
}
func (rsm *RawShaderMaterial) SetValues(parameters ShaderMaterialParameters) {
	rsm.Call("setValues", parameters)
}
func (rsm *RawShaderMaterial) ToJSON(meta js.Value) js.Value {
	return rsm.Call("toJSON", meta)
}
func (rsm *RawShaderMaterial) Update() {
	rsm.Call("update")
}

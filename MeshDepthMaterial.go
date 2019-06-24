// Code generated from "materials/MeshDepthMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type MeshDepthMaterialParameters interface {
}
type MeshDepthMaterial struct {
	js.Value
}

func NewMeshDepthMaterial(parameters MeshDepthMaterialParameters) *MeshDepthMaterial {
	return &MeshDepthMaterial{Value: get("MeshDepthMaterial").New(parameters)}
}
func (mdm *MeshDepthMaterial) AlphaTest() float64 {
	return mdm.Get("alphaTest").Float()
}
func (mdm *MeshDepthMaterial) SetAlphaTest(v float64) {
	mdm.Set("alphaTest", v)
}
func (mdm *MeshDepthMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mdm.Get("blendDst"))
}
func (mdm *MeshDepthMaterial) SetBlendDst(v BlendingDstFactor) {
	mdm.Set("blendDst", v)
}
func (mdm *MeshDepthMaterial) BlendDstAlpha() float64 {
	return mdm.Get("blendDstAlpha").Float()
}
func (mdm *MeshDepthMaterial) SetBlendDstAlpha(v float64) {
	mdm.Set("blendDstAlpha", v)
}
func (mdm *MeshDepthMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mdm.Get("blendEquation"))
}
func (mdm *MeshDepthMaterial) SetBlendEquation(v BlendingEquation) {
	mdm.Set("blendEquation", v)
}
func (mdm *MeshDepthMaterial) BlendEquationAlpha() float64 {
	return mdm.Get("blendEquationAlpha").Float()
}
func (mdm *MeshDepthMaterial) SetBlendEquationAlpha(v float64) {
	mdm.Set("blendEquationAlpha", v)
}
func (mdm *MeshDepthMaterial) BlendSrc() js.Value {
	return mdm.Get("blendSrc")
}
func (mdm *MeshDepthMaterial) SetBlendSrc(v js.Value) {
	mdm.Set("blendSrc", v)
}
func (mdm *MeshDepthMaterial) BlendSrcAlpha() float64 {
	return mdm.Get("blendSrcAlpha").Float()
}
func (mdm *MeshDepthMaterial) SetBlendSrcAlpha(v float64) {
	mdm.Set("blendSrcAlpha", v)
}
func (mdm *MeshDepthMaterial) Blending() Blending {
	return Blending(mdm.Get("blending"))
}
func (mdm *MeshDepthMaterial) SetBlending(v Blending) {
	mdm.Set("blending", v)
}
func (mdm *MeshDepthMaterial) ClipIntersection() bool {
	return mdm.Get("clipIntersection").Bool()
}
func (mdm *MeshDepthMaterial) SetClipIntersection(v bool) {
	mdm.Set("clipIntersection", v)
}
func (mdm *MeshDepthMaterial) ClipShadows() bool {
	return mdm.Get("clipShadows").Bool()
}
func (mdm *MeshDepthMaterial) SetClipShadows(v bool) {
	mdm.Set("clipShadows", v)
}
func (mdm *MeshDepthMaterial) ClippingPlanes() js.Value {
	return mdm.Get("clippingPlanes")
}
func (mdm *MeshDepthMaterial) SetClippingPlanes(v js.Value) {
	mdm.Set("clippingPlanes", v)
}
func (mdm *MeshDepthMaterial) ColorWrite() bool {
	return mdm.Get("colorWrite").Bool()
}
func (mdm *MeshDepthMaterial) SetColorWrite(v bool) {
	mdm.Set("colorWrite", v)
}
func (mdm *MeshDepthMaterial) DepthFunc() DepthModes {
	return DepthModes(mdm.Get("depthFunc"))
}
func (mdm *MeshDepthMaterial) SetDepthFunc(v DepthModes) {
	mdm.Set("depthFunc", v)
}
func (mdm *MeshDepthMaterial) DepthPacking() DepthPackingStrategies {
	return DepthPackingStrategies(mdm.Get("depthPacking"))
}
func (mdm *MeshDepthMaterial) SetDepthPacking(v DepthPackingStrategies) {
	mdm.Set("depthPacking", v)
}
func (mdm *MeshDepthMaterial) DepthTest() bool {
	return mdm.Get("depthTest").Bool()
}
func (mdm *MeshDepthMaterial) SetDepthTest(v bool) {
	mdm.Set("depthTest", v)
}
func (mdm *MeshDepthMaterial) DepthWrite() bool {
	return mdm.Get("depthWrite").Bool()
}
func (mdm *MeshDepthMaterial) SetDepthWrite(v bool) {
	mdm.Set("depthWrite", v)
}
func (mdm *MeshDepthMaterial) DisplacementBias() int {
	return mdm.Get("displacementBias").Int()
}
func (mdm *MeshDepthMaterial) SetDisplacementBias(v int) {
	mdm.Set("displacementBias", v)
}
func (mdm *MeshDepthMaterial) DisplacementMap() *Texture {
	return &Texture{Value: mdm.Get("displacementMap")}
}
func (mdm *MeshDepthMaterial) SetDisplacementMap(v *Texture) {
	mdm.Set("displacementMap", v)
}
func (mdm *MeshDepthMaterial) DisplacementScale() int {
	return mdm.Get("displacementScale").Int()
}
func (mdm *MeshDepthMaterial) SetDisplacementScale(v int) {
	mdm.Set("displacementScale", v)
}
func (mdm *MeshDepthMaterial) Dithering() bool {
	return mdm.Get("dithering").Bool()
}
func (mdm *MeshDepthMaterial) SetDithering(v bool) {
	mdm.Set("dithering", v)
}
func (mdm *MeshDepthMaterial) FlatShading() bool {
	return mdm.Get("flatShading").Bool()
}
func (mdm *MeshDepthMaterial) SetFlatShading(v bool) {
	mdm.Set("flatShading", v)
}
func (mdm *MeshDepthMaterial) Fog() bool {
	return mdm.Get("fog").Bool()
}
func (mdm *MeshDepthMaterial) SetFog(v bool) {
	mdm.Set("fog", v)
}
func (mdm *MeshDepthMaterial) Id() int {
	return mdm.Get("id").Int()
}
func (mdm *MeshDepthMaterial) SetId(v int) {
	mdm.Set("id", v)
}
func (mdm *MeshDepthMaterial) IsMaterial() bool {
	return mdm.Get("isMaterial").Bool()
}
func (mdm *MeshDepthMaterial) SetIsMaterial(v bool) {
	mdm.Set("isMaterial", v)
}
func (mdm *MeshDepthMaterial) Lights() bool {
	return mdm.Get("lights").Bool()
}
func (mdm *MeshDepthMaterial) SetLights(v bool) {
	mdm.Set("lights", v)
}
func (mdm *MeshDepthMaterial) Name() string {
	return mdm.Get("name").String()
}
func (mdm *MeshDepthMaterial) SetName(v string) {
	mdm.Set("name", v)
}
func (mdm *MeshDepthMaterial) NeedsUpdate() bool {
	return mdm.Get("needsUpdate").Bool()
}
func (mdm *MeshDepthMaterial) SetNeedsUpdate(v bool) {
	mdm.Set("needsUpdate", v)
}
func (mdm *MeshDepthMaterial) Opacity() float64 {
	return mdm.Get("opacity").Float()
}
func (mdm *MeshDepthMaterial) SetOpacity(v float64) {
	mdm.Set("opacity", v)
}
func (mdm *MeshDepthMaterial) Overdraw() float64 {
	return mdm.Get("overdraw").Float()
}
func (mdm *MeshDepthMaterial) SetOverdraw(v float64) {
	mdm.Set("overdraw", v)
}
func (mdm *MeshDepthMaterial) PolygonOffset() bool {
	return mdm.Get("polygonOffset").Bool()
}
func (mdm *MeshDepthMaterial) SetPolygonOffset(v bool) {
	mdm.Set("polygonOffset", v)
}
func (mdm *MeshDepthMaterial) PolygonOffsetFactor() float64 {
	return mdm.Get("polygonOffsetFactor").Float()
}
func (mdm *MeshDepthMaterial) SetPolygonOffsetFactor(v float64) {
	mdm.Set("polygonOffsetFactor", v)
}
func (mdm *MeshDepthMaterial) PolygonOffsetUnits() float64 {
	return mdm.Get("polygonOffsetUnits").Float()
}
func (mdm *MeshDepthMaterial) SetPolygonOffsetUnits(v float64) {
	mdm.Set("polygonOffsetUnits", v)
}
func (mdm *MeshDepthMaterial) Precision() string {
	return mdm.Get("precision").String()
}
func (mdm *MeshDepthMaterial) SetPrecision(v string) {
	mdm.Set("precision", v)
}
func (mdm *MeshDepthMaterial) PremultipliedAlpha() bool {
	return mdm.Get("premultipliedAlpha").Bool()
}
func (mdm *MeshDepthMaterial) SetPremultipliedAlpha(v bool) {
	mdm.Set("premultipliedAlpha", v)
}
func (mdm *MeshDepthMaterial) Side() Side {
	return Side(mdm.Get("side"))
}
func (mdm *MeshDepthMaterial) SetSide(v Side) {
	mdm.Set("side", v)
}
func (mdm *MeshDepthMaterial) Transparent() bool {
	return mdm.Get("transparent").Bool()
}
func (mdm *MeshDepthMaterial) SetTransparent(v bool) {
	mdm.Set("transparent", v)
}
func (mdm *MeshDepthMaterial) Type() string {
	return mdm.Get("type").String()
}
func (mdm *MeshDepthMaterial) SetType(v string) {
	mdm.Set("type", v)
}
func (mdm *MeshDepthMaterial) UserData() js.Value {
	return mdm.Get("userData")
}
func (mdm *MeshDepthMaterial) SetUserData(v js.Value) {
	mdm.Set("userData", v)
}
func (mdm *MeshDepthMaterial) Uuid() string {
	return mdm.Get("uuid").String()
}
func (mdm *MeshDepthMaterial) SetUuid(v string) {
	mdm.Set("uuid", v)
}
func (mdm *MeshDepthMaterial) VertexColors() Colors {
	return Colors(mdm.Get("vertexColors"))
}
func (mdm *MeshDepthMaterial) SetVertexColors(v Colors) {
	mdm.Set("vertexColors", v)
}
func (mdm *MeshDepthMaterial) VertexTangents() bool {
	return mdm.Get("vertexTangents").Bool()
}
func (mdm *MeshDepthMaterial) SetVertexTangents(v bool) {
	mdm.Set("vertexTangents", v)
}
func (mdm *MeshDepthMaterial) Visible() bool {
	return mdm.Get("visible").Bool()
}
func (mdm *MeshDepthMaterial) SetVisible(v bool) {
	mdm.Set("visible", v)
}
func (mdm *MeshDepthMaterial) Wireframe() bool {
	return mdm.Get("wireframe").Bool()
}
func (mdm *MeshDepthMaterial) SetWireframe(v bool) {
	mdm.Set("wireframe", v)
}
func (mdm *MeshDepthMaterial) WireframeLinewidth() float64 {
	return mdm.Get("wireframeLinewidth").Float()
}
func (mdm *MeshDepthMaterial) SetWireframeLinewidth(v float64) {
	mdm.Set("wireframeLinewidth", v)
}
func (mdm *MeshDepthMaterial) AddEventListener(typ string, listener js.Value) {
	mdm.Call("addEventListener", typ, listener)
}
func (mdm *MeshDepthMaterial) Clone() *MeshDepthMaterial {
	return &MeshDepthMaterial{Value: mdm.Call("clone")}
}
func (mdm *MeshDepthMaterial) Copy(material *Material) *MeshDepthMaterial {
	return &MeshDepthMaterial{Value: mdm.Call("copy", material)}
}
func (mdm *MeshDepthMaterial) DispatchEvent(event js.Value) {
	mdm.Call("dispatchEvent", event)
}
func (mdm *MeshDepthMaterial) Dispose() {
	mdm.Call("dispose")
}
func (mdm *MeshDepthMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mdm.Call("hasEventListener", typ, listener).Bool()
}
func (mdm *MeshDepthMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mdm.Call("onBeforeCompile", shader, renderer)
}
func (mdm *MeshDepthMaterial) RemoveEventListener(typ string, listener js.Value) {
	mdm.Call("removeEventListener", typ, listener)
}
func (mdm *MeshDepthMaterial) SetValues(parameters MeshDepthMaterialParameters) {
	mdm.Call("setValues", parameters)
}
func (mdm *MeshDepthMaterial) ToJSON(meta js.Value) js.Value {
	return mdm.Call("toJSON", meta)
}
func (mdm *MeshDepthMaterial) Update() {
	mdm.Call("update")
}

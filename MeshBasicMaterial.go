// Code generated from "materials/MeshBasicMaterial.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type MeshBasicMaterialParameters interface {
}
type MeshBasicMaterial struct {
	js.Value
}

func NewMeshBasicMaterial(parameters MeshBasicMaterialParameters) *MeshBasicMaterial {
	return &MeshBasicMaterial{Value: get("MeshBasicMaterial").New(parameters)}
}
func (mbm *MeshBasicMaterial) AlphaMap() *Texture {
	return &Texture{Value: mbm.Get("alphaMap")}
}
func (mbm *MeshBasicMaterial) SetAlphaMap(v *Texture) {
	mbm.Set("alphaMap", v)
}
func (mbm *MeshBasicMaterial) AlphaTest() float64 {
	return mbm.Get("alphaTest").Float()
}
func (mbm *MeshBasicMaterial) SetAlphaTest(v float64) {
	mbm.Set("alphaTest", v)
}
func (mbm *MeshBasicMaterial) AoMap() *Texture {
	return &Texture{Value: mbm.Get("aoMap")}
}
func (mbm *MeshBasicMaterial) SetAoMap(v *Texture) {
	mbm.Set("aoMap", v)
}
func (mbm *MeshBasicMaterial) AoMapIntensity() float64 {
	return mbm.Get("aoMapIntensity").Float()
}
func (mbm *MeshBasicMaterial) SetAoMapIntensity(v float64) {
	mbm.Set("aoMapIntensity", v)
}
func (mbm *MeshBasicMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mbm.Get("blendDst"))
}
func (mbm *MeshBasicMaterial) SetBlendDst(v BlendingDstFactor) {
	mbm.Set("blendDst", v)
}
func (mbm *MeshBasicMaterial) BlendDstAlpha() float64 {
	return mbm.Get("blendDstAlpha").Float()
}
func (mbm *MeshBasicMaterial) SetBlendDstAlpha(v float64) {
	mbm.Set("blendDstAlpha", v)
}
func (mbm *MeshBasicMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mbm.Get("blendEquation"))
}
func (mbm *MeshBasicMaterial) SetBlendEquation(v BlendingEquation) {
	mbm.Set("blendEquation", v)
}
func (mbm *MeshBasicMaterial) BlendEquationAlpha() float64 {
	return mbm.Get("blendEquationAlpha").Float()
}
func (mbm *MeshBasicMaterial) SetBlendEquationAlpha(v float64) {
	mbm.Set("blendEquationAlpha", v)
}
func (mbm *MeshBasicMaterial) BlendSrc() js.Value {
	return mbm.Get("blendSrc")
}
func (mbm *MeshBasicMaterial) SetBlendSrc(v js.Value) {
	mbm.Set("blendSrc", v)
}
func (mbm *MeshBasicMaterial) BlendSrcAlpha() float64 {
	return mbm.Get("blendSrcAlpha").Float()
}
func (mbm *MeshBasicMaterial) SetBlendSrcAlpha(v float64) {
	mbm.Set("blendSrcAlpha", v)
}
func (mbm *MeshBasicMaterial) Blending() Blending {
	return Blending(mbm.Get("blending"))
}
func (mbm *MeshBasicMaterial) SetBlending(v Blending) {
	mbm.Set("blending", v)
}
func (mbm *MeshBasicMaterial) ClipIntersection() bool {
	return mbm.Get("clipIntersection").Bool()
}
func (mbm *MeshBasicMaterial) SetClipIntersection(v bool) {
	mbm.Set("clipIntersection", v)
}
func (mbm *MeshBasicMaterial) ClipShadows() bool {
	return mbm.Get("clipShadows").Bool()
}
func (mbm *MeshBasicMaterial) SetClipShadows(v bool) {
	mbm.Set("clipShadows", v)
}
func (mbm *MeshBasicMaterial) ClippingPlanes() js.Value {
	return mbm.Get("clippingPlanes")
}
func (mbm *MeshBasicMaterial) SetClippingPlanes(v js.Value) {
	mbm.Set("clippingPlanes", v)
}
func (mbm *MeshBasicMaterial) Color() *Color {
	return &Color{Value: mbm.Get("color")}
}
func (mbm *MeshBasicMaterial) SetColor(v *Color) {
	mbm.Set("color", v)
}
func (mbm *MeshBasicMaterial) ColorWrite() bool {
	return mbm.Get("colorWrite").Bool()
}
func (mbm *MeshBasicMaterial) SetColorWrite(v bool) {
	mbm.Set("colorWrite", v)
}
func (mbm *MeshBasicMaterial) Combine() Combine {
	return Combine(mbm.Get("combine"))
}
func (mbm *MeshBasicMaterial) SetCombine(v Combine) {
	mbm.Set("combine", v)
}
func (mbm *MeshBasicMaterial) DepthFunc() DepthModes {
	return DepthModes(mbm.Get("depthFunc"))
}
func (mbm *MeshBasicMaterial) SetDepthFunc(v DepthModes) {
	mbm.Set("depthFunc", v)
}
func (mbm *MeshBasicMaterial) DepthTest() bool {
	return mbm.Get("depthTest").Bool()
}
func (mbm *MeshBasicMaterial) SetDepthTest(v bool) {
	mbm.Set("depthTest", v)
}
func (mbm *MeshBasicMaterial) DepthWrite() bool {
	return mbm.Get("depthWrite").Bool()
}
func (mbm *MeshBasicMaterial) SetDepthWrite(v bool) {
	mbm.Set("depthWrite", v)
}
func (mbm *MeshBasicMaterial) Dithering() bool {
	return mbm.Get("dithering").Bool()
}
func (mbm *MeshBasicMaterial) SetDithering(v bool) {
	mbm.Set("dithering", v)
}
func (mbm *MeshBasicMaterial) EnvMap() *Texture {
	return &Texture{Value: mbm.Get("envMap")}
}
func (mbm *MeshBasicMaterial) SetEnvMap(v *Texture) {
	mbm.Set("envMap", v)
}
func (mbm *MeshBasicMaterial) FlatShading() bool {
	return mbm.Get("flatShading").Bool()
}
func (mbm *MeshBasicMaterial) SetFlatShading(v bool) {
	mbm.Set("flatShading", v)
}
func (mbm *MeshBasicMaterial) Fog() bool {
	return mbm.Get("fog").Bool()
}
func (mbm *MeshBasicMaterial) SetFog(v bool) {
	mbm.Set("fog", v)
}
func (mbm *MeshBasicMaterial) Id() int {
	return mbm.Get("id").Int()
}
func (mbm *MeshBasicMaterial) SetId(v int) {
	mbm.Set("id", v)
}
func (mbm *MeshBasicMaterial) IsMaterial() bool {
	return mbm.Get("isMaterial").Bool()
}
func (mbm *MeshBasicMaterial) SetIsMaterial(v bool) {
	mbm.Set("isMaterial", v)
}
func (mbm *MeshBasicMaterial) Lights() bool {
	return mbm.Get("lights").Bool()
}
func (mbm *MeshBasicMaterial) SetLights(v bool) {
	mbm.Set("lights", v)
}
func (mbm *MeshBasicMaterial) Map() *Texture {
	return &Texture{Value: mbm.Get("map")}
}
func (mbm *MeshBasicMaterial) SetMap(v *Texture) {
	mbm.Set("map", v)
}
func (mbm *MeshBasicMaterial) MorphTargets() bool {
	return mbm.Get("morphTargets").Bool()
}
func (mbm *MeshBasicMaterial) SetMorphTargets(v bool) {
	mbm.Set("morphTargets", v)
}
func (mbm *MeshBasicMaterial) Name() string {
	return mbm.Get("name").String()
}
func (mbm *MeshBasicMaterial) SetName(v string) {
	mbm.Set("name", v)
}
func (mbm *MeshBasicMaterial) NeedsUpdate() bool {
	return mbm.Get("needsUpdate").Bool()
}
func (mbm *MeshBasicMaterial) SetNeedsUpdate(v bool) {
	mbm.Set("needsUpdate", v)
}
func (mbm *MeshBasicMaterial) Opacity() float64 {
	return mbm.Get("opacity").Float()
}
func (mbm *MeshBasicMaterial) SetOpacity(v float64) {
	mbm.Set("opacity", v)
}
func (mbm *MeshBasicMaterial) Overdraw() float64 {
	return mbm.Get("overdraw").Float()
}
func (mbm *MeshBasicMaterial) SetOverdraw(v float64) {
	mbm.Set("overdraw", v)
}
func (mbm *MeshBasicMaterial) PolygonOffset() bool {
	return mbm.Get("polygonOffset").Bool()
}
func (mbm *MeshBasicMaterial) SetPolygonOffset(v bool) {
	mbm.Set("polygonOffset", v)
}
func (mbm *MeshBasicMaterial) PolygonOffsetFactor() float64 {
	return mbm.Get("polygonOffsetFactor").Float()
}
func (mbm *MeshBasicMaterial) SetPolygonOffsetFactor(v float64) {
	mbm.Set("polygonOffsetFactor", v)
}
func (mbm *MeshBasicMaterial) PolygonOffsetUnits() float64 {
	return mbm.Get("polygonOffsetUnits").Float()
}
func (mbm *MeshBasicMaterial) SetPolygonOffsetUnits(v float64) {
	mbm.Set("polygonOffsetUnits", v)
}
func (mbm *MeshBasicMaterial) Precision() string {
	return mbm.Get("precision").String()
}
func (mbm *MeshBasicMaterial) SetPrecision(v string) {
	mbm.Set("precision", v)
}
func (mbm *MeshBasicMaterial) PremultipliedAlpha() bool {
	return mbm.Get("premultipliedAlpha").Bool()
}
func (mbm *MeshBasicMaterial) SetPremultipliedAlpha(v bool) {
	mbm.Set("premultipliedAlpha", v)
}
func (mbm *MeshBasicMaterial) Reflectivity() float64 {
	return mbm.Get("reflectivity").Float()
}
func (mbm *MeshBasicMaterial) SetReflectivity(v float64) {
	mbm.Set("reflectivity", v)
}
func (mbm *MeshBasicMaterial) RefractionRatio() float64 {
	return mbm.Get("refractionRatio").Float()
}
func (mbm *MeshBasicMaterial) SetRefractionRatio(v float64) {
	mbm.Set("refractionRatio", v)
}
func (mbm *MeshBasicMaterial) Side() Side {
	return Side(mbm.Get("side"))
}
func (mbm *MeshBasicMaterial) SetSide(v Side) {
	mbm.Set("side", v)
}
func (mbm *MeshBasicMaterial) Skinning() bool {
	return mbm.Get("skinning").Bool()
}
func (mbm *MeshBasicMaterial) SetSkinning(v bool) {
	mbm.Set("skinning", v)
}
func (mbm *MeshBasicMaterial) SpecularMap() *Texture {
	return &Texture{Value: mbm.Get("specularMap")}
}
func (mbm *MeshBasicMaterial) SetSpecularMap(v *Texture) {
	mbm.Set("specularMap", v)
}
func (mbm *MeshBasicMaterial) Transparent() bool {
	return mbm.Get("transparent").Bool()
}
func (mbm *MeshBasicMaterial) SetTransparent(v bool) {
	mbm.Set("transparent", v)
}
func (mbm *MeshBasicMaterial) Type() string {
	return mbm.Get("type").String()
}
func (mbm *MeshBasicMaterial) SetType(v string) {
	mbm.Set("type", v)
}
func (mbm *MeshBasicMaterial) UserData() js.Value {
	return mbm.Get("userData")
}
func (mbm *MeshBasicMaterial) SetUserData(v js.Value) {
	mbm.Set("userData", v)
}
func (mbm *MeshBasicMaterial) Uuid() string {
	return mbm.Get("uuid").String()
}
func (mbm *MeshBasicMaterial) SetUuid(v string) {
	mbm.Set("uuid", v)
}
func (mbm *MeshBasicMaterial) VertexColors() Colors {
	return Colors(mbm.Get("vertexColors"))
}
func (mbm *MeshBasicMaterial) SetVertexColors(v Colors) {
	mbm.Set("vertexColors", v)
}
func (mbm *MeshBasicMaterial) VertexTangents() bool {
	return mbm.Get("vertexTangents").Bool()
}
func (mbm *MeshBasicMaterial) SetVertexTangents(v bool) {
	mbm.Set("vertexTangents", v)
}
func (mbm *MeshBasicMaterial) Visible() bool {
	return mbm.Get("visible").Bool()
}
func (mbm *MeshBasicMaterial) SetVisible(v bool) {
	mbm.Set("visible", v)
}
func (mbm *MeshBasicMaterial) Wireframe() bool {
	return mbm.Get("wireframe").Bool()
}
func (mbm *MeshBasicMaterial) SetWireframe(v bool) {
	mbm.Set("wireframe", v)
}
func (mbm *MeshBasicMaterial) WireframeLinecap() string {
	return mbm.Get("wireframeLinecap").String()
}
func (mbm *MeshBasicMaterial) SetWireframeLinecap(v string) {
	mbm.Set("wireframeLinecap", v)
}
func (mbm *MeshBasicMaterial) WireframeLinejoin() string {
	return mbm.Get("wireframeLinejoin").String()
}
func (mbm *MeshBasicMaterial) SetWireframeLinejoin(v string) {
	mbm.Set("wireframeLinejoin", v)
}
func (mbm *MeshBasicMaterial) WireframeLinewidth() float64 {
	return mbm.Get("wireframeLinewidth").Float()
}
func (mbm *MeshBasicMaterial) SetWireframeLinewidth(v float64) {
	mbm.Set("wireframeLinewidth", v)
}
func (mbm *MeshBasicMaterial) AddEventListener(typ string, listener js.Value) {
	mbm.Call("addEventListener", typ, listener)
}
func (mbm *MeshBasicMaterial) Clone() *MeshBasicMaterial {
	return &MeshBasicMaterial{Value: mbm.Call("clone")}
}
func (mbm *MeshBasicMaterial) Copy(material *Material) *MeshBasicMaterial {
	return &MeshBasicMaterial{Value: mbm.Call("copy", material)}
}
func (mbm *MeshBasicMaterial) DispatchEvent(event js.Value) {
	mbm.Call("dispatchEvent", event)
}
func (mbm *MeshBasicMaterial) Dispose() {
	mbm.Call("dispose")
}
func (mbm *MeshBasicMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mbm.Call("hasEventListener", typ, listener).Bool()
}
func (mbm *MeshBasicMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mbm.Call("onBeforeCompile", shader, renderer)
}
func (mbm *MeshBasicMaterial) RemoveEventListener(typ string, listener js.Value) {
	mbm.Call("removeEventListener", typ, listener)
}
func (mbm *MeshBasicMaterial) SetValues(parameters MeshBasicMaterialParameters) {
	mbm.Call("setValues", parameters)
}
func (mbm *MeshBasicMaterial) ToJSON(meta js.Value) js.Value {
	return mbm.Call("toJSON", meta)
}
func (mbm *MeshBasicMaterial) Update() {
	mbm.Call("update")
}

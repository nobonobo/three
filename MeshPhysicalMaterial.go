// Code generated from "materials/MeshPhysicalMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type MeshPhysicalMaterialParameters interface {
}

// MeshPhysicalMaterial extend: [MeshStandardMaterial]
type MeshPhysicalMaterial struct {
	js.Value
}

func NewMeshPhysicalMaterial(parameters MeshPhysicalMaterialParameters) *MeshPhysicalMaterial {
	return &MeshPhysicalMaterial{Value: get("MeshPhysicalMaterial").New(parameters)}
}
func (mpm *MeshPhysicalMaterial) JSValue() js.Value {
	return mpm.Value
}
func (mpm *MeshPhysicalMaterial) AlphaMap() *Texture {
	return &Texture{Value: mpm.Get("alphaMap")}
}
func (mpm *MeshPhysicalMaterial) SetAlphaMap(v *Texture) {
	mpm.Set("alphaMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) AlphaTest() float64 {
	return mpm.Get("alphaTest").Float()
}
func (mpm *MeshPhysicalMaterial) SetAlphaTest(v float64) {
	mpm.Set("alphaTest", v)
}
func (mpm *MeshPhysicalMaterial) AoMap() *Texture {
	return &Texture{Value: mpm.Get("aoMap")}
}
func (mpm *MeshPhysicalMaterial) SetAoMap(v *Texture) {
	mpm.Set("aoMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) AoMapIntensity() float64 {
	return mpm.Get("aoMapIntensity").Float()
}
func (mpm *MeshPhysicalMaterial) SetAoMapIntensity(v float64) {
	mpm.Set("aoMapIntensity", v)
}
func (mpm *MeshPhysicalMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mpm.Get("blendDst"))
}
func (mpm *MeshPhysicalMaterial) SetBlendDst(v BlendingDstFactor) {
	mpm.Set("blendDst", v)
}
func (mpm *MeshPhysicalMaterial) BlendDstAlpha() float64 {
	return mpm.Get("blendDstAlpha").Float()
}
func (mpm *MeshPhysicalMaterial) SetBlendDstAlpha(v float64) {
	mpm.Set("blendDstAlpha", v)
}
func (mpm *MeshPhysicalMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mpm.Get("blendEquation"))
}
func (mpm *MeshPhysicalMaterial) SetBlendEquation(v BlendingEquation) {
	mpm.Set("blendEquation", v)
}
func (mpm *MeshPhysicalMaterial) BlendEquationAlpha() float64 {
	return mpm.Get("blendEquationAlpha").Float()
}
func (mpm *MeshPhysicalMaterial) SetBlendEquationAlpha(v float64) {
	mpm.Set("blendEquationAlpha", v)
}
func (mpm *MeshPhysicalMaterial) BlendSrc() js.Value {
	return mpm.Get("blendSrc")
}
func (mpm *MeshPhysicalMaterial) SetBlendSrc(v js.Value) {
	mpm.Set("blendSrc", v)
}
func (mpm *MeshPhysicalMaterial) BlendSrcAlpha() float64 {
	return mpm.Get("blendSrcAlpha").Float()
}
func (mpm *MeshPhysicalMaterial) SetBlendSrcAlpha(v float64) {
	mpm.Set("blendSrcAlpha", v)
}
func (mpm *MeshPhysicalMaterial) Blending() Blending {
	return Blending(mpm.Get("blending"))
}
func (mpm *MeshPhysicalMaterial) SetBlending(v Blending) {
	mpm.Set("blending", v)
}
func (mpm *MeshPhysicalMaterial) BumpMap() *Texture {
	return &Texture{Value: mpm.Get("bumpMap")}
}
func (mpm *MeshPhysicalMaterial) SetBumpMap(v *Texture) {
	mpm.Set("bumpMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) BumpScale() float64 {
	return mpm.Get("bumpScale").Float()
}
func (mpm *MeshPhysicalMaterial) SetBumpScale(v float64) {
	mpm.Set("bumpScale", v)
}
func (mpm *MeshPhysicalMaterial) ClearCoat() float64 {
	return mpm.Get("clearCoat").Float()
}
func (mpm *MeshPhysicalMaterial) SetClearCoat(v float64) {
	mpm.Set("clearCoat", v)
}
func (mpm *MeshPhysicalMaterial) ClearCoatRoughness() float64 {
	return mpm.Get("clearCoatRoughness").Float()
}
func (mpm *MeshPhysicalMaterial) SetClearCoatRoughness(v float64) {
	mpm.Set("clearCoatRoughness", v)
}
func (mpm *MeshPhysicalMaterial) ClipIntersection() bool {
	return mpm.Get("clipIntersection").Bool()
}
func (mpm *MeshPhysicalMaterial) SetClipIntersection(v bool) {
	mpm.Set("clipIntersection", v)
}
func (mpm *MeshPhysicalMaterial) ClipShadows() bool {
	return mpm.Get("clipShadows").Bool()
}
func (mpm *MeshPhysicalMaterial) SetClipShadows(v bool) {
	mpm.Set("clipShadows", v)
}
func (mpm *MeshPhysicalMaterial) ClippingPlanes() js.Value {
	return mpm.Get("clippingPlanes")
}
func (mpm *MeshPhysicalMaterial) SetClippingPlanes(v js.Value) {
	mpm.Set("clippingPlanes", v)
}
func (mpm *MeshPhysicalMaterial) Color() *Color {
	return &Color{Value: mpm.Get("color")}
}
func (mpm *MeshPhysicalMaterial) SetColor(v *Color) {
	mpm.Set("color", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) ColorWrite() bool {
	return mpm.Get("colorWrite").Bool()
}
func (mpm *MeshPhysicalMaterial) SetColorWrite(v bool) {
	mpm.Set("colorWrite", v)
}
func (mpm *MeshPhysicalMaterial) Defines() js.Value {
	return mpm.Get("defines")
}
func (mpm *MeshPhysicalMaterial) SetDefines(v js.Value) {
	mpm.Set("defines", v)
}
func (mpm *MeshPhysicalMaterial) DepthFunc() DepthModes {
	return DepthModes(mpm.Get("depthFunc"))
}
func (mpm *MeshPhysicalMaterial) SetDepthFunc(v DepthModes) {
	mpm.Set("depthFunc", v)
}
func (mpm *MeshPhysicalMaterial) DepthTest() bool {
	return mpm.Get("depthTest").Bool()
}
func (mpm *MeshPhysicalMaterial) SetDepthTest(v bool) {
	mpm.Set("depthTest", v)
}
func (mpm *MeshPhysicalMaterial) DepthWrite() bool {
	return mpm.Get("depthWrite").Bool()
}
func (mpm *MeshPhysicalMaterial) SetDepthWrite(v bool) {
	mpm.Set("depthWrite", v)
}
func (mpm *MeshPhysicalMaterial) DisplacementBias() float64 {
	return mpm.Get("displacementBias").Float()
}
func (mpm *MeshPhysicalMaterial) SetDisplacementBias(v float64) {
	mpm.Set("displacementBias", v)
}
func (mpm *MeshPhysicalMaterial) DisplacementMap() *Texture {
	return &Texture{Value: mpm.Get("displacementMap")}
}
func (mpm *MeshPhysicalMaterial) SetDisplacementMap(v *Texture) {
	mpm.Set("displacementMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) DisplacementScale() float64 {
	return mpm.Get("displacementScale").Float()
}
func (mpm *MeshPhysicalMaterial) SetDisplacementScale(v float64) {
	mpm.Set("displacementScale", v)
}
func (mpm *MeshPhysicalMaterial) Dithering() bool {
	return mpm.Get("dithering").Bool()
}
func (mpm *MeshPhysicalMaterial) SetDithering(v bool) {
	mpm.Set("dithering", v)
}
func (mpm *MeshPhysicalMaterial) Emissive() *Color {
	return &Color{Value: mpm.Get("emissive")}
}
func (mpm *MeshPhysicalMaterial) SetEmissive(v *Color) {
	mpm.Set("emissive", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) EmissiveIntensity() float64 {
	return mpm.Get("emissiveIntensity").Float()
}
func (mpm *MeshPhysicalMaterial) SetEmissiveIntensity(v float64) {
	mpm.Set("emissiveIntensity", v)
}
func (mpm *MeshPhysicalMaterial) EmissiveMap() *Texture {
	return &Texture{Value: mpm.Get("emissiveMap")}
}
func (mpm *MeshPhysicalMaterial) SetEmissiveMap(v *Texture) {
	mpm.Set("emissiveMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) EnvMap() *Texture {
	return &Texture{Value: mpm.Get("envMap")}
}
func (mpm *MeshPhysicalMaterial) SetEnvMap(v *Texture) {
	mpm.Set("envMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) EnvMapIntensity() float64 {
	return mpm.Get("envMapIntensity").Float()
}
func (mpm *MeshPhysicalMaterial) SetEnvMapIntensity(v float64) {
	mpm.Set("envMapIntensity", v)
}
func (mpm *MeshPhysicalMaterial) FlatShading() bool {
	return mpm.Get("flatShading").Bool()
}
func (mpm *MeshPhysicalMaterial) SetFlatShading(v bool) {
	mpm.Set("flatShading", v)
}
func (mpm *MeshPhysicalMaterial) Fog() bool {
	return mpm.Get("fog").Bool()
}
func (mpm *MeshPhysicalMaterial) SetFog(v bool) {
	mpm.Set("fog", v)
}
func (mpm *MeshPhysicalMaterial) Id() int {
	return mpm.Get("id").Int()
}
func (mpm *MeshPhysicalMaterial) SetId(v int) {
	mpm.Set("id", v)
}
func (mpm *MeshPhysicalMaterial) IsMaterial() bool {
	return mpm.Get("isMaterial").Bool()
}
func (mpm *MeshPhysicalMaterial) SetIsMaterial(v bool) {
	mpm.Set("isMaterial", v)
}
func (mpm *MeshPhysicalMaterial) LightMap() *Texture {
	return &Texture{Value: mpm.Get("lightMap")}
}
func (mpm *MeshPhysicalMaterial) SetLightMap(v *Texture) {
	mpm.Set("lightMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) LightMapIntensity() float64 {
	return mpm.Get("lightMapIntensity").Float()
}
func (mpm *MeshPhysicalMaterial) SetLightMapIntensity(v float64) {
	mpm.Set("lightMapIntensity", v)
}
func (mpm *MeshPhysicalMaterial) Lights() bool {
	return mpm.Get("lights").Bool()
}
func (mpm *MeshPhysicalMaterial) SetLights(v bool) {
	mpm.Set("lights", v)
}
func (mpm *MeshPhysicalMaterial) Map() *Texture {
	return &Texture{Value: mpm.Get("map")}
}
func (mpm *MeshPhysicalMaterial) SetMap(v *Texture) {
	mpm.Set("map", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) Metalness() float64 {
	return mpm.Get("metalness").Float()
}
func (mpm *MeshPhysicalMaterial) SetMetalness(v float64) {
	mpm.Set("metalness", v)
}
func (mpm *MeshPhysicalMaterial) MetalnessMap() *Texture {
	return &Texture{Value: mpm.Get("metalnessMap")}
}
func (mpm *MeshPhysicalMaterial) SetMetalnessMap(v *Texture) {
	mpm.Set("metalnessMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) MorphNormals() bool {
	return mpm.Get("morphNormals").Bool()
}
func (mpm *MeshPhysicalMaterial) SetMorphNormals(v bool) {
	mpm.Set("morphNormals", v)
}
func (mpm *MeshPhysicalMaterial) MorphTargets() bool {
	return mpm.Get("morphTargets").Bool()
}
func (mpm *MeshPhysicalMaterial) SetMorphTargets(v bool) {
	mpm.Set("morphTargets", v)
}
func (mpm *MeshPhysicalMaterial) Name() string {
	return mpm.Get("name").String()
}
func (mpm *MeshPhysicalMaterial) SetName(v string) {
	mpm.Set("name", v)
}
func (mpm *MeshPhysicalMaterial) NeedsUpdate() bool {
	return mpm.Get("needsUpdate").Bool()
}
func (mpm *MeshPhysicalMaterial) SetNeedsUpdate(v bool) {
	mpm.Set("needsUpdate", v)
}
func (mpm *MeshPhysicalMaterial) NormalMap() *Texture {
	return &Texture{Value: mpm.Get("normalMap")}
}
func (mpm *MeshPhysicalMaterial) SetNormalMap(v *Texture) {
	mpm.Set("normalMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) NormalMapType() NormalMapTypes {
	return NormalMapTypes(mpm.Get("normalMapType"))
}
func (mpm *MeshPhysicalMaterial) SetNormalMapType(v NormalMapTypes) {
	mpm.Set("normalMapType", v)
}
func (mpm *MeshPhysicalMaterial) NormalScale() float64 {
	return mpm.Get("normalScale").Float()
}
func (mpm *MeshPhysicalMaterial) SetNormalScale(v float64) {
	mpm.Set("normalScale", v)
}
func (mpm *MeshPhysicalMaterial) Opacity() float64 {
	return mpm.Get("opacity").Float()
}
func (mpm *MeshPhysicalMaterial) SetOpacity(v float64) {
	mpm.Set("opacity", v)
}
func (mpm *MeshPhysicalMaterial) Overdraw() float64 {
	return mpm.Get("overdraw").Float()
}
func (mpm *MeshPhysicalMaterial) SetOverdraw(v float64) {
	mpm.Set("overdraw", v)
}
func (mpm *MeshPhysicalMaterial) PolygonOffset() bool {
	return mpm.Get("polygonOffset").Bool()
}
func (mpm *MeshPhysicalMaterial) SetPolygonOffset(v bool) {
	mpm.Set("polygonOffset", v)
}
func (mpm *MeshPhysicalMaterial) PolygonOffsetFactor() float64 {
	return mpm.Get("polygonOffsetFactor").Float()
}
func (mpm *MeshPhysicalMaterial) SetPolygonOffsetFactor(v float64) {
	mpm.Set("polygonOffsetFactor", v)
}
func (mpm *MeshPhysicalMaterial) PolygonOffsetUnits() float64 {
	return mpm.Get("polygonOffsetUnits").Float()
}
func (mpm *MeshPhysicalMaterial) SetPolygonOffsetUnits(v float64) {
	mpm.Set("polygonOffsetUnits", v)
}
func (mpm *MeshPhysicalMaterial) Precision() string {
	return mpm.Get("precision").String()
}
func (mpm *MeshPhysicalMaterial) SetPrecision(v string) {
	mpm.Set("precision", v)
}
func (mpm *MeshPhysicalMaterial) PremultipliedAlpha() bool {
	return mpm.Get("premultipliedAlpha").Bool()
}
func (mpm *MeshPhysicalMaterial) SetPremultipliedAlpha(v bool) {
	mpm.Set("premultipliedAlpha", v)
}
func (mpm *MeshPhysicalMaterial) Reflectivity() float64 {
	return mpm.Get("reflectivity").Float()
}
func (mpm *MeshPhysicalMaterial) SetReflectivity(v float64) {
	mpm.Set("reflectivity", v)
}
func (mpm *MeshPhysicalMaterial) RefractionRatio() float64 {
	return mpm.Get("refractionRatio").Float()
}
func (mpm *MeshPhysicalMaterial) SetRefractionRatio(v float64) {
	mpm.Set("refractionRatio", v)
}
func (mpm *MeshPhysicalMaterial) Roughness() float64 {
	return mpm.Get("roughness").Float()
}
func (mpm *MeshPhysicalMaterial) SetRoughness(v float64) {
	mpm.Set("roughness", v)
}
func (mpm *MeshPhysicalMaterial) RoughnessMap() *Texture {
	return &Texture{Value: mpm.Get("roughnessMap")}
}
func (mpm *MeshPhysicalMaterial) SetRoughnessMap(v *Texture) {
	mpm.Set("roughnessMap", v.JSValue())
}
func (mpm *MeshPhysicalMaterial) Side() Side {
	return Side(mpm.Get("side"))
}
func (mpm *MeshPhysicalMaterial) SetSide(v Side) {
	mpm.Set("side", v)
}
func (mpm *MeshPhysicalMaterial) Skinning() bool {
	return mpm.Get("skinning").Bool()
}
func (mpm *MeshPhysicalMaterial) SetSkinning(v bool) {
	mpm.Set("skinning", v)
}
func (mpm *MeshPhysicalMaterial) Transparent() bool {
	return mpm.Get("transparent").Bool()
}
func (mpm *MeshPhysicalMaterial) SetTransparent(v bool) {
	mpm.Set("transparent", v)
}
func (mpm *MeshPhysicalMaterial) Type() string {
	return mpm.Get("type").String()
}
func (mpm *MeshPhysicalMaterial) SetType(v string) {
	mpm.Set("type", v)
}
func (mpm *MeshPhysicalMaterial) UserData() js.Value {
	return mpm.Get("userData")
}
func (mpm *MeshPhysicalMaterial) SetUserData(v js.Value) {
	mpm.Set("userData", v)
}
func (mpm *MeshPhysicalMaterial) Uuid() string {
	return mpm.Get("uuid").String()
}
func (mpm *MeshPhysicalMaterial) SetUuid(v string) {
	mpm.Set("uuid", v)
}
func (mpm *MeshPhysicalMaterial) VertexColors() Colors {
	return Colors(mpm.Get("vertexColors"))
}
func (mpm *MeshPhysicalMaterial) SetVertexColors(v Colors) {
	mpm.Set("vertexColors", v)
}
func (mpm *MeshPhysicalMaterial) VertexTangents() bool {
	return mpm.Get("vertexTangents").Bool()
}
func (mpm *MeshPhysicalMaterial) SetVertexTangents(v bool) {
	mpm.Set("vertexTangents", v)
}
func (mpm *MeshPhysicalMaterial) Visible() bool {
	return mpm.Get("visible").Bool()
}
func (mpm *MeshPhysicalMaterial) SetVisible(v bool) {
	mpm.Set("visible", v)
}
func (mpm *MeshPhysicalMaterial) Wireframe() bool {
	return mpm.Get("wireframe").Bool()
}
func (mpm *MeshPhysicalMaterial) SetWireframe(v bool) {
	mpm.Set("wireframe", v)
}
func (mpm *MeshPhysicalMaterial) WireframeLinewidth() float64 {
	return mpm.Get("wireframeLinewidth").Float()
}
func (mpm *MeshPhysicalMaterial) SetWireframeLinewidth(v float64) {
	mpm.Set("wireframeLinewidth", v)
}
func (mpm *MeshPhysicalMaterial) AddEventListener(typ string, listener js.Value) {
	mpm.Call("addEventListener", typ, listener)
}
func (mpm *MeshPhysicalMaterial) Clone() *MeshPhysicalMaterial {
	return &MeshPhysicalMaterial{Value: mpm.Call("clone")}
}
func (mpm *MeshPhysicalMaterial) Copy(material Material) *MeshPhysicalMaterial {
	return &MeshPhysicalMaterial{Value: mpm.Call("copy", material.JSValue())}
}
func (mpm *MeshPhysicalMaterial) DispatchEvent(event js.Value) {
	mpm.Call("dispatchEvent", event)
}
func (mpm *MeshPhysicalMaterial) Dispose() {
	mpm.Call("dispose")
}
func (mpm *MeshPhysicalMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mpm.Call("hasEventListener", typ, listener).Bool()
}
func (mpm *MeshPhysicalMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mpm.Call("onBeforeCompile", shader, renderer.JSValue())
}
func (mpm *MeshPhysicalMaterial) RemoveEventListener(typ string, listener js.Value) {
	mpm.Call("removeEventListener", typ, listener)
}
func (mpm *MeshPhysicalMaterial) SetValues(parameters MeshStandardMaterialParameters) {
	mpm.Call("setValues", parameters)
}
func (mpm *MeshPhysicalMaterial) ToJSON(meta js.Value) js.Value {
	return mpm.Call("toJSON", meta)
}
func (mpm *MeshPhysicalMaterial) Update() {
	mpm.Call("update")
}

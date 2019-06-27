// Code generated from "materials/MeshPhongMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type MeshPhongMaterialParameters interface {
}

// MeshPhongMaterial extend: [Material]
type MeshPhongMaterial struct {
	js.Value
}

func NewMeshPhongMaterial(parameters MeshPhongMaterialParameters) *MeshPhongMaterial {
	return &MeshPhongMaterial{Value: get("MeshPhongMaterial").New(parameters)}
}
func (mpm *MeshPhongMaterial) JSValue() js.Value {
	return mpm.Value
}
func (mpm *MeshPhongMaterial) AlphaMap() *Texture {
	return &Texture{Value: mpm.Get("alphaMap")}
}
func (mpm *MeshPhongMaterial) SetAlphaMap(v *Texture) {
	mpm.Set("alphaMap", v.Value)
}
func (mpm *MeshPhongMaterial) AlphaTest() float64 {
	return mpm.Get("alphaTest").Float()
}
func (mpm *MeshPhongMaterial) SetAlphaTest(v float64) {
	mpm.Set("alphaTest", v)
}
func (mpm *MeshPhongMaterial) AoMap() *Texture {
	return &Texture{Value: mpm.Get("aoMap")}
}
func (mpm *MeshPhongMaterial) SetAoMap(v *Texture) {
	mpm.Set("aoMap", v.Value)
}
func (mpm *MeshPhongMaterial) AoMapIntensity() float64 {
	return mpm.Get("aoMapIntensity").Float()
}
func (mpm *MeshPhongMaterial) SetAoMapIntensity(v float64) {
	mpm.Set("aoMapIntensity", v)
}
func (mpm *MeshPhongMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mpm.Get("blendDst"))
}
func (mpm *MeshPhongMaterial) SetBlendDst(v BlendingDstFactor) {
	mpm.Set("blendDst", v)
}
func (mpm *MeshPhongMaterial) BlendDstAlpha() float64 {
	return mpm.Get("blendDstAlpha").Float()
}
func (mpm *MeshPhongMaterial) SetBlendDstAlpha(v float64) {
	mpm.Set("blendDstAlpha", v)
}
func (mpm *MeshPhongMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mpm.Get("blendEquation"))
}
func (mpm *MeshPhongMaterial) SetBlendEquation(v BlendingEquation) {
	mpm.Set("blendEquation", v)
}
func (mpm *MeshPhongMaterial) BlendEquationAlpha() float64 {
	return mpm.Get("blendEquationAlpha").Float()
}
func (mpm *MeshPhongMaterial) SetBlendEquationAlpha(v float64) {
	mpm.Set("blendEquationAlpha", v)
}
func (mpm *MeshPhongMaterial) BlendSrc() js.Value {
	return mpm.Get("blendSrc")
}
func (mpm *MeshPhongMaterial) SetBlendSrc(v js.Value) {
	mpm.Set("blendSrc", v)
}
func (mpm *MeshPhongMaterial) BlendSrcAlpha() float64 {
	return mpm.Get("blendSrcAlpha").Float()
}
func (mpm *MeshPhongMaterial) SetBlendSrcAlpha(v float64) {
	mpm.Set("blendSrcAlpha", v)
}
func (mpm *MeshPhongMaterial) Blending() Blending {
	return Blending(mpm.Get("blending"))
}
func (mpm *MeshPhongMaterial) SetBlending(v Blending) {
	mpm.Set("blending", v)
}
func (mpm *MeshPhongMaterial) BumpMap() *Texture {
	return &Texture{Value: mpm.Get("bumpMap")}
}
func (mpm *MeshPhongMaterial) SetBumpMap(v *Texture) {
	mpm.Set("bumpMap", v.Value)
}
func (mpm *MeshPhongMaterial) BumpScale() float64 {
	return mpm.Get("bumpScale").Float()
}
func (mpm *MeshPhongMaterial) SetBumpScale(v float64) {
	mpm.Set("bumpScale", v)
}
func (mpm *MeshPhongMaterial) ClipIntersection() bool {
	return mpm.Get("clipIntersection").Bool()
}
func (mpm *MeshPhongMaterial) SetClipIntersection(v bool) {
	mpm.Set("clipIntersection", v)
}
func (mpm *MeshPhongMaterial) ClipShadows() bool {
	return mpm.Get("clipShadows").Bool()
}
func (mpm *MeshPhongMaterial) SetClipShadows(v bool) {
	mpm.Set("clipShadows", v)
}
func (mpm *MeshPhongMaterial) ClippingPlanes() js.Value {
	return mpm.Get("clippingPlanes")
}
func (mpm *MeshPhongMaterial) SetClippingPlanes(v js.Value) {
	mpm.Set("clippingPlanes", v)
}
func (mpm *MeshPhongMaterial) Color() *Color {
	return &Color{Value: mpm.Get("color")}
}
func (mpm *MeshPhongMaterial) SetColor(v *Color) {
	mpm.Set("color", v.Value)
}
func (mpm *MeshPhongMaterial) ColorWrite() bool {
	return mpm.Get("colorWrite").Bool()
}
func (mpm *MeshPhongMaterial) SetColorWrite(v bool) {
	mpm.Set("colorWrite", v)
}
func (mpm *MeshPhongMaterial) Combine() Combine {
	return Combine(mpm.Get("combine"))
}
func (mpm *MeshPhongMaterial) SetCombine(v Combine) {
	mpm.Set("combine", v)
}
func (mpm *MeshPhongMaterial) DepthFunc() DepthModes {
	return DepthModes(mpm.Get("depthFunc"))
}
func (mpm *MeshPhongMaterial) SetDepthFunc(v DepthModes) {
	mpm.Set("depthFunc", v)
}
func (mpm *MeshPhongMaterial) DepthTest() bool {
	return mpm.Get("depthTest").Bool()
}
func (mpm *MeshPhongMaterial) SetDepthTest(v bool) {
	mpm.Set("depthTest", v)
}
func (mpm *MeshPhongMaterial) DepthWrite() bool {
	return mpm.Get("depthWrite").Bool()
}
func (mpm *MeshPhongMaterial) SetDepthWrite(v bool) {
	mpm.Set("depthWrite", v)
}
func (mpm *MeshPhongMaterial) DisplacementBias() float64 {
	return mpm.Get("displacementBias").Float()
}
func (mpm *MeshPhongMaterial) SetDisplacementBias(v float64) {
	mpm.Set("displacementBias", v)
}
func (mpm *MeshPhongMaterial) DisplacementMap() *Texture {
	return &Texture{Value: mpm.Get("displacementMap")}
}
func (mpm *MeshPhongMaterial) SetDisplacementMap(v *Texture) {
	mpm.Set("displacementMap", v.Value)
}
func (mpm *MeshPhongMaterial) DisplacementScale() float64 {
	return mpm.Get("displacementScale").Float()
}
func (mpm *MeshPhongMaterial) SetDisplacementScale(v float64) {
	mpm.Set("displacementScale", v)
}
func (mpm *MeshPhongMaterial) Dithering() bool {
	return mpm.Get("dithering").Bool()
}
func (mpm *MeshPhongMaterial) SetDithering(v bool) {
	mpm.Set("dithering", v)
}
func (mpm *MeshPhongMaterial) Emissive() *Color {
	return &Color{Value: mpm.Get("emissive")}
}
func (mpm *MeshPhongMaterial) SetEmissive(v *Color) {
	mpm.Set("emissive", v.Value)
}
func (mpm *MeshPhongMaterial) EmissiveIntensity() float64 {
	return mpm.Get("emissiveIntensity").Float()
}
func (mpm *MeshPhongMaterial) SetEmissiveIntensity(v float64) {
	mpm.Set("emissiveIntensity", v)
}
func (mpm *MeshPhongMaterial) EmissiveMap() *Texture {
	return &Texture{Value: mpm.Get("emissiveMap")}
}
func (mpm *MeshPhongMaterial) SetEmissiveMap(v *Texture) {
	mpm.Set("emissiveMap", v.Value)
}
func (mpm *MeshPhongMaterial) EnvMap() *Texture {
	return &Texture{Value: mpm.Get("envMap")}
}
func (mpm *MeshPhongMaterial) SetEnvMap(v *Texture) {
	mpm.Set("envMap", v.Value)
}
func (mpm *MeshPhongMaterial) FlatShading() bool {
	return mpm.Get("flatShading").Bool()
}
func (mpm *MeshPhongMaterial) SetFlatShading(v bool) {
	mpm.Set("flatShading", v)
}
func (mpm *MeshPhongMaterial) Fog() bool {
	return mpm.Get("fog").Bool()
}
func (mpm *MeshPhongMaterial) SetFog(v bool) {
	mpm.Set("fog", v)
}
func (mpm *MeshPhongMaterial) Id() int {
	return mpm.Get("id").Int()
}
func (mpm *MeshPhongMaterial) SetId(v int) {
	mpm.Set("id", v)
}
func (mpm *MeshPhongMaterial) IsMaterial() bool {
	return mpm.Get("isMaterial").Bool()
}
func (mpm *MeshPhongMaterial) SetIsMaterial(v bool) {
	mpm.Set("isMaterial", v)
}
func (mpm *MeshPhongMaterial) LightMap() *Texture {
	return &Texture{Value: mpm.Get("lightMap")}
}
func (mpm *MeshPhongMaterial) SetLightMap(v *Texture) {
	mpm.Set("lightMap", v.Value)
}
func (mpm *MeshPhongMaterial) LightMapIntensity() float64 {
	return mpm.Get("lightMapIntensity").Float()
}
func (mpm *MeshPhongMaterial) SetLightMapIntensity(v float64) {
	mpm.Set("lightMapIntensity", v)
}
func (mpm *MeshPhongMaterial) Lights() bool {
	return mpm.Get("lights").Bool()
}
func (mpm *MeshPhongMaterial) SetLights(v bool) {
	mpm.Set("lights", v)
}
func (mpm *MeshPhongMaterial) Map() *Texture {
	return &Texture{Value: mpm.Get("map")}
}
func (mpm *MeshPhongMaterial) SetMap(v *Texture) {
	mpm.Set("map", v.Value)
}
func (mpm *MeshPhongMaterial) Metal() bool {
	return mpm.Get("metal").Bool()
}
func (mpm *MeshPhongMaterial) SetMetal(v bool) {
	mpm.Set("metal", v)
}
func (mpm *MeshPhongMaterial) MorphNormals() bool {
	return mpm.Get("morphNormals").Bool()
}
func (mpm *MeshPhongMaterial) SetMorphNormals(v bool) {
	mpm.Set("morphNormals", v)
}
func (mpm *MeshPhongMaterial) MorphTargets() bool {
	return mpm.Get("morphTargets").Bool()
}
func (mpm *MeshPhongMaterial) SetMorphTargets(v bool) {
	mpm.Set("morphTargets", v)
}
func (mpm *MeshPhongMaterial) Name() string {
	return mpm.Get("name").String()
}
func (mpm *MeshPhongMaterial) SetName(v string) {
	mpm.Set("name", v)
}
func (mpm *MeshPhongMaterial) NeedsUpdate() bool {
	return mpm.Get("needsUpdate").Bool()
}
func (mpm *MeshPhongMaterial) SetNeedsUpdate(v bool) {
	mpm.Set("needsUpdate", v)
}
func (mpm *MeshPhongMaterial) NormalMap() *Texture {
	return &Texture{Value: mpm.Get("normalMap")}
}
func (mpm *MeshPhongMaterial) SetNormalMap(v *Texture) {
	mpm.Set("normalMap", v.Value)
}
func (mpm *MeshPhongMaterial) NormalMapType() NormalMapTypes {
	return NormalMapTypes(mpm.Get("normalMapType"))
}
func (mpm *MeshPhongMaterial) SetNormalMapType(v NormalMapTypes) {
	mpm.Set("normalMapType", v)
}
func (mpm *MeshPhongMaterial) NormalScale() *Vector2 {
	return &Vector2{Value: mpm.Get("normalScale")}
}
func (mpm *MeshPhongMaterial) SetNormalScale(v *Vector2) {
	mpm.Set("normalScale", v.Value)
}
func (mpm *MeshPhongMaterial) Opacity() float64 {
	return mpm.Get("opacity").Float()
}
func (mpm *MeshPhongMaterial) SetOpacity(v float64) {
	mpm.Set("opacity", v)
}
func (mpm *MeshPhongMaterial) Overdraw() float64 {
	return mpm.Get("overdraw").Float()
}
func (mpm *MeshPhongMaterial) SetOverdraw(v float64) {
	mpm.Set("overdraw", v)
}
func (mpm *MeshPhongMaterial) PolygonOffset() bool {
	return mpm.Get("polygonOffset").Bool()
}
func (mpm *MeshPhongMaterial) SetPolygonOffset(v bool) {
	mpm.Set("polygonOffset", v)
}
func (mpm *MeshPhongMaterial) PolygonOffsetFactor() float64 {
	return mpm.Get("polygonOffsetFactor").Float()
}
func (mpm *MeshPhongMaterial) SetPolygonOffsetFactor(v float64) {
	mpm.Set("polygonOffsetFactor", v)
}
func (mpm *MeshPhongMaterial) PolygonOffsetUnits() float64 {
	return mpm.Get("polygonOffsetUnits").Float()
}
func (mpm *MeshPhongMaterial) SetPolygonOffsetUnits(v float64) {
	mpm.Set("polygonOffsetUnits", v)
}
func (mpm *MeshPhongMaterial) Precision() string {
	return mpm.Get("precision").String()
}
func (mpm *MeshPhongMaterial) SetPrecision(v string) {
	mpm.Set("precision", v)
}
func (mpm *MeshPhongMaterial) PremultipliedAlpha() bool {
	return mpm.Get("premultipliedAlpha").Bool()
}
func (mpm *MeshPhongMaterial) SetPremultipliedAlpha(v bool) {
	mpm.Set("premultipliedAlpha", v)
}
func (mpm *MeshPhongMaterial) Reflectivity() float64 {
	return mpm.Get("reflectivity").Float()
}
func (mpm *MeshPhongMaterial) SetReflectivity(v float64) {
	mpm.Set("reflectivity", v)
}
func (mpm *MeshPhongMaterial) RefractionRatio() float64 {
	return mpm.Get("refractionRatio").Float()
}
func (mpm *MeshPhongMaterial) SetRefractionRatio(v float64) {
	mpm.Set("refractionRatio", v)
}
func (mpm *MeshPhongMaterial) Shininess() float64 {
	return mpm.Get("shininess").Float()
}
func (mpm *MeshPhongMaterial) SetShininess(v float64) {
	mpm.Set("shininess", v)
}
func (mpm *MeshPhongMaterial) Side() Side {
	return Side(mpm.Get("side"))
}
func (mpm *MeshPhongMaterial) SetSide(v Side) {
	mpm.Set("side", v)
}
func (mpm *MeshPhongMaterial) Skinning() bool {
	return mpm.Get("skinning").Bool()
}
func (mpm *MeshPhongMaterial) SetSkinning(v bool) {
	mpm.Set("skinning", v)
}
func (mpm *MeshPhongMaterial) Specular() *Color {
	return &Color{Value: mpm.Get("specular")}
}
func (mpm *MeshPhongMaterial) SetSpecular(v *Color) {
	mpm.Set("specular", v.Value)
}
func (mpm *MeshPhongMaterial) SpecularMap() *Texture {
	return &Texture{Value: mpm.Get("specularMap")}
}
func (mpm *MeshPhongMaterial) SetSpecularMap(v *Texture) {
	mpm.Set("specularMap", v.Value)
}
func (mpm *MeshPhongMaterial) Transparent() bool {
	return mpm.Get("transparent").Bool()
}
func (mpm *MeshPhongMaterial) SetTransparent(v bool) {
	mpm.Set("transparent", v)
}
func (mpm *MeshPhongMaterial) Type() string {
	return mpm.Get("type").String()
}
func (mpm *MeshPhongMaterial) SetType(v string) {
	mpm.Set("type", v)
}
func (mpm *MeshPhongMaterial) UserData() js.Value {
	return mpm.Get("userData")
}
func (mpm *MeshPhongMaterial) SetUserData(v js.Value) {
	mpm.Set("userData", v)
}
func (mpm *MeshPhongMaterial) Uuid() string {
	return mpm.Get("uuid").String()
}
func (mpm *MeshPhongMaterial) SetUuid(v string) {
	mpm.Set("uuid", v)
}
func (mpm *MeshPhongMaterial) VertexColors() Colors {
	return Colors(mpm.Get("vertexColors"))
}
func (mpm *MeshPhongMaterial) SetVertexColors(v Colors) {
	mpm.Set("vertexColors", v)
}
func (mpm *MeshPhongMaterial) VertexTangents() bool {
	return mpm.Get("vertexTangents").Bool()
}
func (mpm *MeshPhongMaterial) SetVertexTangents(v bool) {
	mpm.Set("vertexTangents", v)
}
func (mpm *MeshPhongMaterial) Visible() bool {
	return mpm.Get("visible").Bool()
}
func (mpm *MeshPhongMaterial) SetVisible(v bool) {
	mpm.Set("visible", v)
}
func (mpm *MeshPhongMaterial) Wireframe() bool {
	return mpm.Get("wireframe").Bool()
}
func (mpm *MeshPhongMaterial) SetWireframe(v bool) {
	mpm.Set("wireframe", v)
}
func (mpm *MeshPhongMaterial) WireframeLinecap() string {
	return mpm.Get("wireframeLinecap").String()
}
func (mpm *MeshPhongMaterial) SetWireframeLinecap(v string) {
	mpm.Set("wireframeLinecap", v)
}
func (mpm *MeshPhongMaterial) WireframeLinejoin() string {
	return mpm.Get("wireframeLinejoin").String()
}
func (mpm *MeshPhongMaterial) SetWireframeLinejoin(v string) {
	mpm.Set("wireframeLinejoin", v)
}
func (mpm *MeshPhongMaterial) WireframeLinewidth() float64 {
	return mpm.Get("wireframeLinewidth").Float()
}
func (mpm *MeshPhongMaterial) SetWireframeLinewidth(v float64) {
	mpm.Set("wireframeLinewidth", v)
}
func (mpm *MeshPhongMaterial) AddEventListener(typ string, listener js.Value) {
	mpm.Call("addEventListener", typ, listener)
}
func (mpm *MeshPhongMaterial) Clone() *MeshPhongMaterial {
	return &MeshPhongMaterial{Value: mpm.Call("clone")}
}
func (mpm *MeshPhongMaterial) Copy(material Material) *MeshPhongMaterial {
	return &MeshPhongMaterial{Value: mpm.Call("copy", material.JSValue())}
}
func (mpm *MeshPhongMaterial) DispatchEvent(event js.Value) {
	mpm.Call("dispatchEvent", event)
}
func (mpm *MeshPhongMaterial) Dispose() {
	mpm.Call("dispose")
}
func (mpm *MeshPhongMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mpm.Call("hasEventListener", typ, listener).Bool()
}
func (mpm *MeshPhongMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mpm.Call("onBeforeCompile", shader, renderer)
}
func (mpm *MeshPhongMaterial) RemoveEventListener(typ string, listener js.Value) {
	mpm.Call("removeEventListener", typ, listener)
}
func (mpm *MeshPhongMaterial) SetValues(parameters MeshPhongMaterialParameters) {
	mpm.Call("setValues", parameters)
}
func (mpm *MeshPhongMaterial) ToJSON(meta js.Value) js.Value {
	return mpm.Call("toJSON", meta)
}
func (mpm *MeshPhongMaterial) Update() {
	mpm.Call("update")
}

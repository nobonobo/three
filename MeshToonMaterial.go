// Code generated from "materials/MeshToonMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type MeshToonMaterialParameters interface {
}

// MeshToonMaterial extend: [MeshPhongMaterial]
type MeshToonMaterial struct {
	js.Value
}

func NewMeshToonMaterial(parameters MeshToonMaterialParameters) *MeshToonMaterial {
	return &MeshToonMaterial{Value: get("MeshToonMaterial").New(parameters)}
}
func (mtm *MeshToonMaterial) JSValue() js.Value {
	return mtm.Value
}
func (mtm *MeshToonMaterial) AlphaMap() *Texture {
	return &Texture{Value: mtm.Get("alphaMap")}
}
func (mtm *MeshToonMaterial) SetAlphaMap(v *Texture) {
	mtm.Set("alphaMap", v.Value)
}
func (mtm *MeshToonMaterial) AlphaTest() int {
	return mtm.Get("alphaTest").Int()
}
func (mtm *MeshToonMaterial) SetAlphaTest(v int) {
	mtm.Set("alphaTest", v)
}
func (mtm *MeshToonMaterial) AoMap() *Texture {
	return &Texture{Value: mtm.Get("aoMap")}
}
func (mtm *MeshToonMaterial) SetAoMap(v *Texture) {
	mtm.Set("aoMap", v.Value)
}
func (mtm *MeshToonMaterial) AoMapIntensity() int {
	return mtm.Get("aoMapIntensity").Int()
}
func (mtm *MeshToonMaterial) SetAoMapIntensity(v int) {
	mtm.Set("aoMapIntensity", v)
}
func (mtm *MeshToonMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mtm.Get("blendDst"))
}
func (mtm *MeshToonMaterial) SetBlendDst(v BlendingDstFactor) {
	mtm.Set("blendDst", v)
}
func (mtm *MeshToonMaterial) BlendDstAlpha() int {
	return mtm.Get("blendDstAlpha").Int()
}
func (mtm *MeshToonMaterial) SetBlendDstAlpha(v int) {
	mtm.Set("blendDstAlpha", v)
}
func (mtm *MeshToonMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mtm.Get("blendEquation"))
}
func (mtm *MeshToonMaterial) SetBlendEquation(v BlendingEquation) {
	mtm.Set("blendEquation", v)
}
func (mtm *MeshToonMaterial) BlendEquationAlpha() int {
	return mtm.Get("blendEquationAlpha").Int()
}
func (mtm *MeshToonMaterial) SetBlendEquationAlpha(v int) {
	mtm.Set("blendEquationAlpha", v)
}
func (mtm *MeshToonMaterial) BlendSrc() js.Value {
	return mtm.Get("blendSrc")
}
func (mtm *MeshToonMaterial) SetBlendSrc(v js.Value) {
	mtm.Set("blendSrc", v)
}
func (mtm *MeshToonMaterial) BlendSrcAlpha() int {
	return mtm.Get("blendSrcAlpha").Int()
}
func (mtm *MeshToonMaterial) SetBlendSrcAlpha(v int) {
	mtm.Set("blendSrcAlpha", v)
}
func (mtm *MeshToonMaterial) Blending() Blending {
	return Blending(mtm.Get("blending"))
}
func (mtm *MeshToonMaterial) SetBlending(v Blending) {
	mtm.Set("blending", v)
}
func (mtm *MeshToonMaterial) BumpMap() *Texture {
	return &Texture{Value: mtm.Get("bumpMap")}
}
func (mtm *MeshToonMaterial) SetBumpMap(v *Texture) {
	mtm.Set("bumpMap", v.Value)
}
func (mtm *MeshToonMaterial) BumpScale() int {
	return mtm.Get("bumpScale").Int()
}
func (mtm *MeshToonMaterial) SetBumpScale(v int) {
	mtm.Set("bumpScale", v)
}
func (mtm *MeshToonMaterial) ClipIntersection() bool {
	return mtm.Get("clipIntersection").Bool()
}
func (mtm *MeshToonMaterial) SetClipIntersection(v bool) {
	mtm.Set("clipIntersection", v)
}
func (mtm *MeshToonMaterial) ClipShadows() bool {
	return mtm.Get("clipShadows").Bool()
}
func (mtm *MeshToonMaterial) SetClipShadows(v bool) {
	mtm.Set("clipShadows", v)
}
func (mtm *MeshToonMaterial) ClippingPlanes() js.Value {
	return mtm.Get("clippingPlanes")
}
func (mtm *MeshToonMaterial) SetClippingPlanes(v js.Value) {
	mtm.Set("clippingPlanes", v)
}
func (mtm *MeshToonMaterial) Color() *Color {
	return &Color{Value: mtm.Get("color")}
}
func (mtm *MeshToonMaterial) SetColor(v *Color) {
	mtm.Set("color", v.Value)
}
func (mtm *MeshToonMaterial) ColorWrite() bool {
	return mtm.Get("colorWrite").Bool()
}
func (mtm *MeshToonMaterial) SetColorWrite(v bool) {
	mtm.Set("colorWrite", v)
}
func (mtm *MeshToonMaterial) Combine() Combine {
	return Combine(mtm.Get("combine"))
}
func (mtm *MeshToonMaterial) SetCombine(v Combine) {
	mtm.Set("combine", v)
}
func (mtm *MeshToonMaterial) DepthFunc() DepthModes {
	return DepthModes(mtm.Get("depthFunc"))
}
func (mtm *MeshToonMaterial) SetDepthFunc(v DepthModes) {
	mtm.Set("depthFunc", v)
}
func (mtm *MeshToonMaterial) DepthTest() bool {
	return mtm.Get("depthTest").Bool()
}
func (mtm *MeshToonMaterial) SetDepthTest(v bool) {
	mtm.Set("depthTest", v)
}
func (mtm *MeshToonMaterial) DepthWrite() bool {
	return mtm.Get("depthWrite").Bool()
}
func (mtm *MeshToonMaterial) SetDepthWrite(v bool) {
	mtm.Set("depthWrite", v)
}
func (mtm *MeshToonMaterial) DisplacementBias() int {
	return mtm.Get("displacementBias").Int()
}
func (mtm *MeshToonMaterial) SetDisplacementBias(v int) {
	mtm.Set("displacementBias", v)
}
func (mtm *MeshToonMaterial) DisplacementMap() *Texture {
	return &Texture{Value: mtm.Get("displacementMap")}
}
func (mtm *MeshToonMaterial) SetDisplacementMap(v *Texture) {
	mtm.Set("displacementMap", v.Value)
}
func (mtm *MeshToonMaterial) DisplacementScale() int {
	return mtm.Get("displacementScale").Int()
}
func (mtm *MeshToonMaterial) SetDisplacementScale(v int) {
	mtm.Set("displacementScale", v)
}
func (mtm *MeshToonMaterial) Dithering() bool {
	return mtm.Get("dithering").Bool()
}
func (mtm *MeshToonMaterial) SetDithering(v bool) {
	mtm.Set("dithering", v)
}
func (mtm *MeshToonMaterial) Emissive() *Color {
	return &Color{Value: mtm.Get("emissive")}
}
func (mtm *MeshToonMaterial) SetEmissive(v *Color) {
	mtm.Set("emissive", v.Value)
}
func (mtm *MeshToonMaterial) EmissiveIntensity() int {
	return mtm.Get("emissiveIntensity").Int()
}
func (mtm *MeshToonMaterial) SetEmissiveIntensity(v int) {
	mtm.Set("emissiveIntensity", v)
}
func (mtm *MeshToonMaterial) EmissiveMap() *Texture {
	return &Texture{Value: mtm.Get("emissiveMap")}
}
func (mtm *MeshToonMaterial) SetEmissiveMap(v *Texture) {
	mtm.Set("emissiveMap", v.Value)
}
func (mtm *MeshToonMaterial) EnvMap() *Texture {
	return &Texture{Value: mtm.Get("envMap")}
}
func (mtm *MeshToonMaterial) SetEnvMap(v *Texture) {
	mtm.Set("envMap", v.Value)
}
func (mtm *MeshToonMaterial) FlatShading() bool {
	return mtm.Get("flatShading").Bool()
}
func (mtm *MeshToonMaterial) SetFlatShading(v bool) {
	mtm.Set("flatShading", v)
}
func (mtm *MeshToonMaterial) Fog() bool {
	return mtm.Get("fog").Bool()
}
func (mtm *MeshToonMaterial) SetFog(v bool) {
	mtm.Set("fog", v)
}
func (mtm *MeshToonMaterial) GradientMap() *Texture {
	return &Texture{Value: mtm.Get("gradientMap")}
}
func (mtm *MeshToonMaterial) SetGradientMap(v *Texture) {
	mtm.Set("gradientMap", v.Value)
}
func (mtm *MeshToonMaterial) Id() int {
	return mtm.Get("id").Int()
}
func (mtm *MeshToonMaterial) SetId(v int) {
	mtm.Set("id", v)
}
func (mtm *MeshToonMaterial) IsMaterial() bool {
	return mtm.Get("isMaterial").Bool()
}
func (mtm *MeshToonMaterial) SetIsMaterial(v bool) {
	mtm.Set("isMaterial", v)
}
func (mtm *MeshToonMaterial) LightMap() *Texture {
	return &Texture{Value: mtm.Get("lightMap")}
}
func (mtm *MeshToonMaterial) SetLightMap(v *Texture) {
	mtm.Set("lightMap", v.Value)
}
func (mtm *MeshToonMaterial) LightMapIntensity() int {
	return mtm.Get("lightMapIntensity").Int()
}
func (mtm *MeshToonMaterial) SetLightMapIntensity(v int) {
	mtm.Set("lightMapIntensity", v)
}
func (mtm *MeshToonMaterial) Lights() bool {
	return mtm.Get("lights").Bool()
}
func (mtm *MeshToonMaterial) SetLights(v bool) {
	mtm.Set("lights", v)
}
func (mtm *MeshToonMaterial) Map() *Texture {
	return &Texture{Value: mtm.Get("map")}
}
func (mtm *MeshToonMaterial) SetMap(v *Texture) {
	mtm.Set("map", v.Value)
}
func (mtm *MeshToonMaterial) Metal() bool {
	return mtm.Get("metal").Bool()
}
func (mtm *MeshToonMaterial) SetMetal(v bool) {
	mtm.Set("metal", v)
}
func (mtm *MeshToonMaterial) MorphNormals() bool {
	return mtm.Get("morphNormals").Bool()
}
func (mtm *MeshToonMaterial) SetMorphNormals(v bool) {
	mtm.Set("morphNormals", v)
}
func (mtm *MeshToonMaterial) MorphTargets() bool {
	return mtm.Get("morphTargets").Bool()
}
func (mtm *MeshToonMaterial) SetMorphTargets(v bool) {
	mtm.Set("morphTargets", v)
}
func (mtm *MeshToonMaterial) Name() string {
	return mtm.Get("name").String()
}
func (mtm *MeshToonMaterial) SetName(v string) {
	mtm.Set("name", v)
}
func (mtm *MeshToonMaterial) NeedsUpdate() bool {
	return mtm.Get("needsUpdate").Bool()
}
func (mtm *MeshToonMaterial) SetNeedsUpdate(v bool) {
	mtm.Set("needsUpdate", v)
}
func (mtm *MeshToonMaterial) NormalMap() *Texture {
	return &Texture{Value: mtm.Get("normalMap")}
}
func (mtm *MeshToonMaterial) SetNormalMap(v *Texture) {
	mtm.Set("normalMap", v.Value)
}
func (mtm *MeshToonMaterial) NormalMapType() NormalMapTypes {
	return NormalMapTypes(mtm.Get("normalMapType"))
}
func (mtm *MeshToonMaterial) SetNormalMapType(v NormalMapTypes) {
	mtm.Set("normalMapType", v)
}
func (mtm *MeshToonMaterial) NormalScale() *Vector2 {
	return &Vector2{Value: mtm.Get("normalScale")}
}
func (mtm *MeshToonMaterial) SetNormalScale(v *Vector2) {
	mtm.Set("normalScale", v.Value)
}
func (mtm *MeshToonMaterial) Opacity() int {
	return mtm.Get("opacity").Int()
}
func (mtm *MeshToonMaterial) SetOpacity(v int) {
	mtm.Set("opacity", v)
}
func (mtm *MeshToonMaterial) Overdraw() int {
	return mtm.Get("overdraw").Int()
}
func (mtm *MeshToonMaterial) SetOverdraw(v int) {
	mtm.Set("overdraw", v)
}
func (mtm *MeshToonMaterial) PolygonOffset() bool {
	return mtm.Get("polygonOffset").Bool()
}
func (mtm *MeshToonMaterial) SetPolygonOffset(v bool) {
	mtm.Set("polygonOffset", v)
}
func (mtm *MeshToonMaterial) PolygonOffsetFactor() int {
	return mtm.Get("polygonOffsetFactor").Int()
}
func (mtm *MeshToonMaterial) SetPolygonOffsetFactor(v int) {
	mtm.Set("polygonOffsetFactor", v)
}
func (mtm *MeshToonMaterial) PolygonOffsetUnits() int {
	return mtm.Get("polygonOffsetUnits").Int()
}
func (mtm *MeshToonMaterial) SetPolygonOffsetUnits(v int) {
	mtm.Set("polygonOffsetUnits", v)
}
func (mtm *MeshToonMaterial) Precision() string {
	return mtm.Get("precision").String()
}
func (mtm *MeshToonMaterial) SetPrecision(v string) {
	mtm.Set("precision", v)
}
func (mtm *MeshToonMaterial) PremultipliedAlpha() bool {
	return mtm.Get("premultipliedAlpha").Bool()
}
func (mtm *MeshToonMaterial) SetPremultipliedAlpha(v bool) {
	mtm.Set("premultipliedAlpha", v)
}
func (mtm *MeshToonMaterial) Reflectivity() int {
	return mtm.Get("reflectivity").Int()
}
func (mtm *MeshToonMaterial) SetReflectivity(v int) {
	mtm.Set("reflectivity", v)
}
func (mtm *MeshToonMaterial) RefractionRatio() int {
	return mtm.Get("refractionRatio").Int()
}
func (mtm *MeshToonMaterial) SetRefractionRatio(v int) {
	mtm.Set("refractionRatio", v)
}
func (mtm *MeshToonMaterial) Shininess() int {
	return mtm.Get("shininess").Int()
}
func (mtm *MeshToonMaterial) SetShininess(v int) {
	mtm.Set("shininess", v)
}
func (mtm *MeshToonMaterial) Side() Side {
	return Side(mtm.Get("side"))
}
func (mtm *MeshToonMaterial) SetSide(v Side) {
	mtm.Set("side", v)
}
func (mtm *MeshToonMaterial) Skinning() bool {
	return mtm.Get("skinning").Bool()
}
func (mtm *MeshToonMaterial) SetSkinning(v bool) {
	mtm.Set("skinning", v)
}
func (mtm *MeshToonMaterial) Specular() *Color {
	return &Color{Value: mtm.Get("specular")}
}
func (mtm *MeshToonMaterial) SetSpecular(v *Color) {
	mtm.Set("specular", v.Value)
}
func (mtm *MeshToonMaterial) SpecularMap() *Texture {
	return &Texture{Value: mtm.Get("specularMap")}
}
func (mtm *MeshToonMaterial) SetSpecularMap(v *Texture) {
	mtm.Set("specularMap", v.Value)
}
func (mtm *MeshToonMaterial) Transparent() bool {
	return mtm.Get("transparent").Bool()
}
func (mtm *MeshToonMaterial) SetTransparent(v bool) {
	mtm.Set("transparent", v)
}
func (mtm *MeshToonMaterial) Type() string {
	return mtm.Get("type").String()
}
func (mtm *MeshToonMaterial) SetType(v string) {
	mtm.Set("type", v)
}
func (mtm *MeshToonMaterial) UserData() js.Value {
	return mtm.Get("userData")
}
func (mtm *MeshToonMaterial) SetUserData(v js.Value) {
	mtm.Set("userData", v)
}
func (mtm *MeshToonMaterial) Uuid() string {
	return mtm.Get("uuid").String()
}
func (mtm *MeshToonMaterial) SetUuid(v string) {
	mtm.Set("uuid", v)
}
func (mtm *MeshToonMaterial) VertexColors() Colors {
	return Colors(mtm.Get("vertexColors"))
}
func (mtm *MeshToonMaterial) SetVertexColors(v Colors) {
	mtm.Set("vertexColors", v)
}
func (mtm *MeshToonMaterial) VertexTangents() bool {
	return mtm.Get("vertexTangents").Bool()
}
func (mtm *MeshToonMaterial) SetVertexTangents(v bool) {
	mtm.Set("vertexTangents", v)
}
func (mtm *MeshToonMaterial) Visible() bool {
	return mtm.Get("visible").Bool()
}
func (mtm *MeshToonMaterial) SetVisible(v bool) {
	mtm.Set("visible", v)
}
func (mtm *MeshToonMaterial) Wireframe() bool {
	return mtm.Get("wireframe").Bool()
}
func (mtm *MeshToonMaterial) SetWireframe(v bool) {
	mtm.Set("wireframe", v)
}
func (mtm *MeshToonMaterial) WireframeLinecap() string {
	return mtm.Get("wireframeLinecap").String()
}
func (mtm *MeshToonMaterial) SetWireframeLinecap(v string) {
	mtm.Set("wireframeLinecap", v)
}
func (mtm *MeshToonMaterial) WireframeLinejoin() string {
	return mtm.Get("wireframeLinejoin").String()
}
func (mtm *MeshToonMaterial) SetWireframeLinejoin(v string) {
	mtm.Set("wireframeLinejoin", v)
}
func (mtm *MeshToonMaterial) WireframeLinewidth() int {
	return mtm.Get("wireframeLinewidth").Int()
}
func (mtm *MeshToonMaterial) SetWireframeLinewidth(v int) {
	mtm.Set("wireframeLinewidth", v)
}
func (mtm *MeshToonMaterial) AddEventListener(typ string, listener js.Value) {
	mtm.Call("addEventListener", typ, listener)
}
func (mtm *MeshToonMaterial) Clone() *MeshToonMaterial {
	return &MeshToonMaterial{Value: mtm.Call("clone")}
}
func (mtm *MeshToonMaterial) Copy(material Material) *MeshToonMaterial {
	return &MeshToonMaterial{Value: mtm.Call("copy", material.JSValue())}
}
func (mtm *MeshToonMaterial) DispatchEvent(event js.Value) {
	mtm.Call("dispatchEvent", event)
}
func (mtm *MeshToonMaterial) Dispose() {
	mtm.Call("dispose")
}
func (mtm *MeshToonMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mtm.Call("hasEventListener", typ, listener).Bool()
}
func (mtm *MeshToonMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mtm.Call("onBeforeCompile", shader, renderer)
}
func (mtm *MeshToonMaterial) RemoveEventListener(typ string, listener js.Value) {
	mtm.Call("removeEventListener", typ, listener)
}
func (mtm *MeshToonMaterial) SetValues(parameters MeshToonMaterialParameters) {
	mtm.Call("setValues", parameters)
}
func (mtm *MeshToonMaterial) ToJSON(meta js.Value) js.Value {
	return mtm.Call("toJSON", meta)
}
func (mtm *MeshToonMaterial) Update() {
	mtm.Call("update")
}

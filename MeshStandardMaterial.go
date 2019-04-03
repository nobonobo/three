// Code generated from "materials/MeshStandardMaterial.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type MeshStandardMaterialParameters interface {
}
type MeshStandardMaterial struct {
	js.Value
}

func NewMeshStandardMaterial(parameters MeshStandardMaterialParameters) *MeshStandardMaterial {
	return &MeshStandardMaterial{Value: get("MeshStandardMaterial").New(parameters)}
}
func (msm *MeshStandardMaterial) AlphaMap() *Texture {
	return &Texture{Value: msm.Get("alphaMap")}
}
func (msm *MeshStandardMaterial) SetAlphaMap(v *Texture) {
	msm.Set("alphaMap", v)
}
func (msm *MeshStandardMaterial) AlphaTest() float64 {
	return msm.Get("alphaTest").Float()
}
func (msm *MeshStandardMaterial) SetAlphaTest(v float64) {
	msm.Set("alphaTest", v)
}
func (msm *MeshStandardMaterial) AoMap() *Texture {
	return &Texture{Value: msm.Get("aoMap")}
}
func (msm *MeshStandardMaterial) SetAoMap(v *Texture) {
	msm.Set("aoMap", v)
}
func (msm *MeshStandardMaterial) AoMapIntensity() float64 {
	return msm.Get("aoMapIntensity").Float()
}
func (msm *MeshStandardMaterial) SetAoMapIntensity(v float64) {
	msm.Set("aoMapIntensity", v)
}
func (msm *MeshStandardMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(msm.Get("blendDst"))
}
func (msm *MeshStandardMaterial) SetBlendDst(v BlendingDstFactor) {
	msm.Set("blendDst", v)
}
func (msm *MeshStandardMaterial) BlendDstAlpha() float64 {
	return msm.Get("blendDstAlpha").Float()
}
func (msm *MeshStandardMaterial) SetBlendDstAlpha(v float64) {
	msm.Set("blendDstAlpha", v)
}
func (msm *MeshStandardMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(msm.Get("blendEquation"))
}
func (msm *MeshStandardMaterial) SetBlendEquation(v BlendingEquation) {
	msm.Set("blendEquation", v)
}
func (msm *MeshStandardMaterial) BlendEquationAlpha() float64 {
	return msm.Get("blendEquationAlpha").Float()
}
func (msm *MeshStandardMaterial) SetBlendEquationAlpha(v float64) {
	msm.Set("blendEquationAlpha", v)
}
func (msm *MeshStandardMaterial) BlendSrc() js.Value {
	return msm.Get("blendSrc")
}
func (msm *MeshStandardMaterial) SetBlendSrc(v js.Value) {
	msm.Set("blendSrc", v)
}
func (msm *MeshStandardMaterial) BlendSrcAlpha() float64 {
	return msm.Get("blendSrcAlpha").Float()
}
func (msm *MeshStandardMaterial) SetBlendSrcAlpha(v float64) {
	msm.Set("blendSrcAlpha", v)
}
func (msm *MeshStandardMaterial) Blending() Blending {
	return Blending(msm.Get("blending"))
}
func (msm *MeshStandardMaterial) SetBlending(v Blending) {
	msm.Set("blending", v)
}
func (msm *MeshStandardMaterial) BumpMap() *Texture {
	return &Texture{Value: msm.Get("bumpMap")}
}
func (msm *MeshStandardMaterial) SetBumpMap(v *Texture) {
	msm.Set("bumpMap", v)
}
func (msm *MeshStandardMaterial) BumpScale() float64 {
	return msm.Get("bumpScale").Float()
}
func (msm *MeshStandardMaterial) SetBumpScale(v float64) {
	msm.Set("bumpScale", v)
}
func (msm *MeshStandardMaterial) ClipIntersection() bool {
	return msm.Get("clipIntersection").Bool()
}
func (msm *MeshStandardMaterial) SetClipIntersection(v bool) {
	msm.Set("clipIntersection", v)
}
func (msm *MeshStandardMaterial) ClipShadows() bool {
	return msm.Get("clipShadows").Bool()
}
func (msm *MeshStandardMaterial) SetClipShadows(v bool) {
	msm.Set("clipShadows", v)
}
func (msm *MeshStandardMaterial) ClippingPlanes() js.Value {
	return msm.Get("clippingPlanes")
}
func (msm *MeshStandardMaterial) SetClippingPlanes(v js.Value) {
	msm.Set("clippingPlanes", v)
}
func (msm *MeshStandardMaterial) Color() *Color {
	return &Color{Value: msm.Get("color")}
}
func (msm *MeshStandardMaterial) SetColor(v *Color) {
	msm.Set("color", v)
}
func (msm *MeshStandardMaterial) ColorWrite() bool {
	return msm.Get("colorWrite").Bool()
}
func (msm *MeshStandardMaterial) SetColorWrite(v bool) {
	msm.Set("colorWrite", v)
}
func (msm *MeshStandardMaterial) Defines() js.Value {
	return msm.Get("defines")
}
func (msm *MeshStandardMaterial) SetDefines(v js.Value) {
	msm.Set("defines", v)
}
func (msm *MeshStandardMaterial) DepthFunc() DepthModes {
	return DepthModes(msm.Get("depthFunc"))
}
func (msm *MeshStandardMaterial) SetDepthFunc(v DepthModes) {
	msm.Set("depthFunc", v)
}
func (msm *MeshStandardMaterial) DepthTest() bool {
	return msm.Get("depthTest").Bool()
}
func (msm *MeshStandardMaterial) SetDepthTest(v bool) {
	msm.Set("depthTest", v)
}
func (msm *MeshStandardMaterial) DepthWrite() bool {
	return msm.Get("depthWrite").Bool()
}
func (msm *MeshStandardMaterial) SetDepthWrite(v bool) {
	msm.Set("depthWrite", v)
}
func (msm *MeshStandardMaterial) DisplacementBias() float64 {
	return msm.Get("displacementBias").Float()
}
func (msm *MeshStandardMaterial) SetDisplacementBias(v float64) {
	msm.Set("displacementBias", v)
}
func (msm *MeshStandardMaterial) DisplacementMap() *Texture {
	return &Texture{Value: msm.Get("displacementMap")}
}
func (msm *MeshStandardMaterial) SetDisplacementMap(v *Texture) {
	msm.Set("displacementMap", v)
}
func (msm *MeshStandardMaterial) DisplacementScale() float64 {
	return msm.Get("displacementScale").Float()
}
func (msm *MeshStandardMaterial) SetDisplacementScale(v float64) {
	msm.Set("displacementScale", v)
}
func (msm *MeshStandardMaterial) Dithering() bool {
	return msm.Get("dithering").Bool()
}
func (msm *MeshStandardMaterial) SetDithering(v bool) {
	msm.Set("dithering", v)
}
func (msm *MeshStandardMaterial) Emissive() *Color {
	return &Color{Value: msm.Get("emissive")}
}
func (msm *MeshStandardMaterial) SetEmissive(v *Color) {
	msm.Set("emissive", v)
}
func (msm *MeshStandardMaterial) EmissiveIntensity() float64 {
	return msm.Get("emissiveIntensity").Float()
}
func (msm *MeshStandardMaterial) SetEmissiveIntensity(v float64) {
	msm.Set("emissiveIntensity", v)
}
func (msm *MeshStandardMaterial) EmissiveMap() *Texture {
	return &Texture{Value: msm.Get("emissiveMap")}
}
func (msm *MeshStandardMaterial) SetEmissiveMap(v *Texture) {
	msm.Set("emissiveMap", v)
}
func (msm *MeshStandardMaterial) EnvMap() *Texture {
	return &Texture{Value: msm.Get("envMap")}
}
func (msm *MeshStandardMaterial) SetEnvMap(v *Texture) {
	msm.Set("envMap", v)
}
func (msm *MeshStandardMaterial) EnvMapIntensity() float64 {
	return msm.Get("envMapIntensity").Float()
}
func (msm *MeshStandardMaterial) SetEnvMapIntensity(v float64) {
	msm.Set("envMapIntensity", v)
}
func (msm *MeshStandardMaterial) FlatShading() bool {
	return msm.Get("flatShading").Bool()
}
func (msm *MeshStandardMaterial) SetFlatShading(v bool) {
	msm.Set("flatShading", v)
}
func (msm *MeshStandardMaterial) Fog() bool {
	return msm.Get("fog").Bool()
}
func (msm *MeshStandardMaterial) SetFog(v bool) {
	msm.Set("fog", v)
}
func (msm *MeshStandardMaterial) Id() int {
	return msm.Get("id").Int()
}
func (msm *MeshStandardMaterial) SetId(v int) {
	msm.Set("id", v)
}
func (msm *MeshStandardMaterial) IsMaterial() bool {
	return msm.Get("isMaterial").Bool()
}
func (msm *MeshStandardMaterial) SetIsMaterial(v bool) {
	msm.Set("isMaterial", v)
}
func (msm *MeshStandardMaterial) LightMap() *Texture {
	return &Texture{Value: msm.Get("lightMap")}
}
func (msm *MeshStandardMaterial) SetLightMap(v *Texture) {
	msm.Set("lightMap", v)
}
func (msm *MeshStandardMaterial) LightMapIntensity() float64 {
	return msm.Get("lightMapIntensity").Float()
}
func (msm *MeshStandardMaterial) SetLightMapIntensity(v float64) {
	msm.Set("lightMapIntensity", v)
}
func (msm *MeshStandardMaterial) Lights() bool {
	return msm.Get("lights").Bool()
}
func (msm *MeshStandardMaterial) SetLights(v bool) {
	msm.Set("lights", v)
}
func (msm *MeshStandardMaterial) Map() *Texture {
	return &Texture{Value: msm.Get("map")}
}
func (msm *MeshStandardMaterial) SetMap(v *Texture) {
	msm.Set("map", v)
}
func (msm *MeshStandardMaterial) Metalness() float64 {
	return msm.Get("metalness").Float()
}
func (msm *MeshStandardMaterial) SetMetalness(v float64) {
	msm.Set("metalness", v)
}
func (msm *MeshStandardMaterial) MetalnessMap() *Texture {
	return &Texture{Value: msm.Get("metalnessMap")}
}
func (msm *MeshStandardMaterial) SetMetalnessMap(v *Texture) {
	msm.Set("metalnessMap", v)
}
func (msm *MeshStandardMaterial) MorphNormals() bool {
	return msm.Get("morphNormals").Bool()
}
func (msm *MeshStandardMaterial) SetMorphNormals(v bool) {
	msm.Set("morphNormals", v)
}
func (msm *MeshStandardMaterial) MorphTargets() bool {
	return msm.Get("morphTargets").Bool()
}
func (msm *MeshStandardMaterial) SetMorphTargets(v bool) {
	msm.Set("morphTargets", v)
}
func (msm *MeshStandardMaterial) Name() string {
	return msm.Get("name").String()
}
func (msm *MeshStandardMaterial) SetName(v string) {
	msm.Set("name", v)
}
func (msm *MeshStandardMaterial) NeedsUpdate() bool {
	return msm.Get("needsUpdate").Bool()
}
func (msm *MeshStandardMaterial) SetNeedsUpdate(v bool) {
	msm.Set("needsUpdate", v)
}
func (msm *MeshStandardMaterial) NormalMap() *Texture {
	return &Texture{Value: msm.Get("normalMap")}
}
func (msm *MeshStandardMaterial) SetNormalMap(v *Texture) {
	msm.Set("normalMap", v)
}
func (msm *MeshStandardMaterial) NormalMapType() NormalMapTypes {
	return NormalMapTypes(msm.Get("normalMapType"))
}
func (msm *MeshStandardMaterial) SetNormalMapType(v NormalMapTypes) {
	msm.Set("normalMapType", v)
}
func (msm *MeshStandardMaterial) NormalScale() float64 {
	return msm.Get("normalScale").Float()
}
func (msm *MeshStandardMaterial) SetNormalScale(v float64) {
	msm.Set("normalScale", v)
}
func (msm *MeshStandardMaterial) Opacity() float64 {
	return msm.Get("opacity").Float()
}
func (msm *MeshStandardMaterial) SetOpacity(v float64) {
	msm.Set("opacity", v)
}
func (msm *MeshStandardMaterial) Overdraw() float64 {
	return msm.Get("overdraw").Float()
}
func (msm *MeshStandardMaterial) SetOverdraw(v float64) {
	msm.Set("overdraw", v)
}
func (msm *MeshStandardMaterial) PolygonOffset() bool {
	return msm.Get("polygonOffset").Bool()
}
func (msm *MeshStandardMaterial) SetPolygonOffset(v bool) {
	msm.Set("polygonOffset", v)
}
func (msm *MeshStandardMaterial) PolygonOffsetFactor() float64 {
	return msm.Get("polygonOffsetFactor").Float()
}
func (msm *MeshStandardMaterial) SetPolygonOffsetFactor(v float64) {
	msm.Set("polygonOffsetFactor", v)
}
func (msm *MeshStandardMaterial) PolygonOffsetUnits() float64 {
	return msm.Get("polygonOffsetUnits").Float()
}
func (msm *MeshStandardMaterial) SetPolygonOffsetUnits(v float64) {
	msm.Set("polygonOffsetUnits", v)
}
func (msm *MeshStandardMaterial) Precision() string {
	return msm.Get("precision").String()
}
func (msm *MeshStandardMaterial) SetPrecision(v string) {
	msm.Set("precision", v)
}
func (msm *MeshStandardMaterial) PremultipliedAlpha() bool {
	return msm.Get("premultipliedAlpha").Bool()
}
func (msm *MeshStandardMaterial) SetPremultipliedAlpha(v bool) {
	msm.Set("premultipliedAlpha", v)
}
func (msm *MeshStandardMaterial) RefractionRatio() float64 {
	return msm.Get("refractionRatio").Float()
}
func (msm *MeshStandardMaterial) SetRefractionRatio(v float64) {
	msm.Set("refractionRatio", v)
}
func (msm *MeshStandardMaterial) Roughness() float64 {
	return msm.Get("roughness").Float()
}
func (msm *MeshStandardMaterial) SetRoughness(v float64) {
	msm.Set("roughness", v)
}
func (msm *MeshStandardMaterial) RoughnessMap() *Texture {
	return &Texture{Value: msm.Get("roughnessMap")}
}
func (msm *MeshStandardMaterial) SetRoughnessMap(v *Texture) {
	msm.Set("roughnessMap", v)
}
func (msm *MeshStandardMaterial) Side() Side {
	return Side(msm.Get("side"))
}
func (msm *MeshStandardMaterial) SetSide(v Side) {
	msm.Set("side", v)
}
func (msm *MeshStandardMaterial) Skinning() bool {
	return msm.Get("skinning").Bool()
}
func (msm *MeshStandardMaterial) SetSkinning(v bool) {
	msm.Set("skinning", v)
}
func (msm *MeshStandardMaterial) Transparent() bool {
	return msm.Get("transparent").Bool()
}
func (msm *MeshStandardMaterial) SetTransparent(v bool) {
	msm.Set("transparent", v)
}
func (msm *MeshStandardMaterial) Type() string {
	return msm.Get("type").String()
}
func (msm *MeshStandardMaterial) SetType(v string) {
	msm.Set("type", v)
}
func (msm *MeshStandardMaterial) UserData() js.Value {
	return msm.Get("userData")
}
func (msm *MeshStandardMaterial) SetUserData(v js.Value) {
	msm.Set("userData", v)
}
func (msm *MeshStandardMaterial) Uuid() string {
	return msm.Get("uuid").String()
}
func (msm *MeshStandardMaterial) SetUuid(v string) {
	msm.Set("uuid", v)
}
func (msm *MeshStandardMaterial) VertexColors() Colors {
	return Colors(msm.Get("vertexColors"))
}
func (msm *MeshStandardMaterial) SetVertexColors(v Colors) {
	msm.Set("vertexColors", v)
}
func (msm *MeshStandardMaterial) VertexTangents() bool {
	return msm.Get("vertexTangents").Bool()
}
func (msm *MeshStandardMaterial) SetVertexTangents(v bool) {
	msm.Set("vertexTangents", v)
}
func (msm *MeshStandardMaterial) Visible() bool {
	return msm.Get("visible").Bool()
}
func (msm *MeshStandardMaterial) SetVisible(v bool) {
	msm.Set("visible", v)
}
func (msm *MeshStandardMaterial) Wireframe() bool {
	return msm.Get("wireframe").Bool()
}
func (msm *MeshStandardMaterial) SetWireframe(v bool) {
	msm.Set("wireframe", v)
}
func (msm *MeshStandardMaterial) WireframeLinewidth() float64 {
	return msm.Get("wireframeLinewidth").Float()
}
func (msm *MeshStandardMaterial) SetWireframeLinewidth(v float64) {
	msm.Set("wireframeLinewidth", v)
}
func (msm *MeshStandardMaterial) AddEventListener(typ string, listener js.Value) {
	msm.Call("addEventListener", typ, listener)
}
func (msm *MeshStandardMaterial) Clone() *MeshStandardMaterial {
	return &MeshStandardMaterial{Value: msm.Call("clone")}
}
func (msm *MeshStandardMaterial) Copy(material *Material) *MeshStandardMaterial {
	return &MeshStandardMaterial{Value: msm.Call("copy", material)}
}
func (msm *MeshStandardMaterial) DispatchEvent(event js.Value) {
	msm.Call("dispatchEvent", event)
}
func (msm *MeshStandardMaterial) Dispose() {
	msm.Call("dispose")
}
func (msm *MeshStandardMaterial) HasEventListener(typ string, listener js.Value) bool {
	return msm.Call("hasEventListener", typ, listener).Bool()
}
func (msm *MeshStandardMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	msm.Call("onBeforeCompile", shader, renderer)
}
func (msm *MeshStandardMaterial) RemoveEventListener(typ string, listener js.Value) {
	msm.Call("removeEventListener", typ, listener)
}
func (msm *MeshStandardMaterial) SetValues(parameters MeshStandardMaterialParameters) {
	msm.Call("setValues", parameters)
}
func (msm *MeshStandardMaterial) ToJSON(meta js.Value) js.Value {
	return msm.Call("toJSON", meta)
}
func (msm *MeshStandardMaterial) Update() {
	msm.Call("update")
}

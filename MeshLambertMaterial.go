// Code generated from "materials/MeshLambertMaterial.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type MeshLambertMaterialParameters interface {
}
type MeshLambertMaterial struct {
	js.Value
}

func NewMeshLambertMaterial(parameters MeshLambertMaterialParameters) *MeshLambertMaterial {
	return &MeshLambertMaterial{Value: get("MeshLambertMaterial").New(parameters)}
}
func (mlm *MeshLambertMaterial) AlphaMap() *Texture {
	return &Texture{Value: mlm.Get("alphaMap")}
}
func (mlm *MeshLambertMaterial) SetAlphaMap(v *Texture) {
	mlm.Set("alphaMap", v)
}
func (mlm *MeshLambertMaterial) AlphaTest() float64 {
	return mlm.Get("alphaTest").Float()
}
func (mlm *MeshLambertMaterial) SetAlphaTest(v float64) {
	mlm.Set("alphaTest", v)
}
func (mlm *MeshLambertMaterial) AoMap() *Texture {
	return &Texture{Value: mlm.Get("aoMap")}
}
func (mlm *MeshLambertMaterial) SetAoMap(v *Texture) {
	mlm.Set("aoMap", v)
}
func (mlm *MeshLambertMaterial) AoMapIntensity() float64 {
	return mlm.Get("aoMapIntensity").Float()
}
func (mlm *MeshLambertMaterial) SetAoMapIntensity(v float64) {
	mlm.Set("aoMapIntensity", v)
}
func (mlm *MeshLambertMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mlm.Get("blendDst"))
}
func (mlm *MeshLambertMaterial) SetBlendDst(v BlendingDstFactor) {
	mlm.Set("blendDst", v)
}
func (mlm *MeshLambertMaterial) BlendDstAlpha() float64 {
	return mlm.Get("blendDstAlpha").Float()
}
func (mlm *MeshLambertMaterial) SetBlendDstAlpha(v float64) {
	mlm.Set("blendDstAlpha", v)
}
func (mlm *MeshLambertMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mlm.Get("blendEquation"))
}
func (mlm *MeshLambertMaterial) SetBlendEquation(v BlendingEquation) {
	mlm.Set("blendEquation", v)
}
func (mlm *MeshLambertMaterial) BlendEquationAlpha() float64 {
	return mlm.Get("blendEquationAlpha").Float()
}
func (mlm *MeshLambertMaterial) SetBlendEquationAlpha(v float64) {
	mlm.Set("blendEquationAlpha", v)
}
func (mlm *MeshLambertMaterial) BlendSrc() js.Value {
	return mlm.Get("blendSrc")
}
func (mlm *MeshLambertMaterial) SetBlendSrc(v js.Value) {
	mlm.Set("blendSrc", v)
}
func (mlm *MeshLambertMaterial) BlendSrcAlpha() float64 {
	return mlm.Get("blendSrcAlpha").Float()
}
func (mlm *MeshLambertMaterial) SetBlendSrcAlpha(v float64) {
	mlm.Set("blendSrcAlpha", v)
}
func (mlm *MeshLambertMaterial) Blending() Blending {
	return Blending(mlm.Get("blending"))
}
func (mlm *MeshLambertMaterial) SetBlending(v Blending) {
	mlm.Set("blending", v)
}
func (mlm *MeshLambertMaterial) ClipIntersection() bool {
	return mlm.Get("clipIntersection").Bool()
}
func (mlm *MeshLambertMaterial) SetClipIntersection(v bool) {
	mlm.Set("clipIntersection", v)
}
func (mlm *MeshLambertMaterial) ClipShadows() bool {
	return mlm.Get("clipShadows").Bool()
}
func (mlm *MeshLambertMaterial) SetClipShadows(v bool) {
	mlm.Set("clipShadows", v)
}
func (mlm *MeshLambertMaterial) ClippingPlanes() js.Value {
	return mlm.Get("clippingPlanes")
}
func (mlm *MeshLambertMaterial) SetClippingPlanes(v js.Value) {
	mlm.Set("clippingPlanes", v)
}
func (mlm *MeshLambertMaterial) Color() *Color {
	return &Color{Value: mlm.Get("color")}
}
func (mlm *MeshLambertMaterial) SetColor(v *Color) {
	mlm.Set("color", v)
}
func (mlm *MeshLambertMaterial) ColorWrite() bool {
	return mlm.Get("colorWrite").Bool()
}
func (mlm *MeshLambertMaterial) SetColorWrite(v bool) {
	mlm.Set("colorWrite", v)
}
func (mlm *MeshLambertMaterial) Combine() Combine {
	return Combine(mlm.Get("combine"))
}
func (mlm *MeshLambertMaterial) SetCombine(v Combine) {
	mlm.Set("combine", v)
}
func (mlm *MeshLambertMaterial) DepthFunc() DepthModes {
	return DepthModes(mlm.Get("depthFunc"))
}
func (mlm *MeshLambertMaterial) SetDepthFunc(v DepthModes) {
	mlm.Set("depthFunc", v)
}
func (mlm *MeshLambertMaterial) DepthTest() bool {
	return mlm.Get("depthTest").Bool()
}
func (mlm *MeshLambertMaterial) SetDepthTest(v bool) {
	mlm.Set("depthTest", v)
}
func (mlm *MeshLambertMaterial) DepthWrite() bool {
	return mlm.Get("depthWrite").Bool()
}
func (mlm *MeshLambertMaterial) SetDepthWrite(v bool) {
	mlm.Set("depthWrite", v)
}
func (mlm *MeshLambertMaterial) Dithering() bool {
	return mlm.Get("dithering").Bool()
}
func (mlm *MeshLambertMaterial) SetDithering(v bool) {
	mlm.Set("dithering", v)
}
func (mlm *MeshLambertMaterial) Emissive() *Color {
	return &Color{Value: mlm.Get("emissive")}
}
func (mlm *MeshLambertMaterial) SetEmissive(v *Color) {
	mlm.Set("emissive", v)
}
func (mlm *MeshLambertMaterial) EmissiveIntensity() float64 {
	return mlm.Get("emissiveIntensity").Float()
}
func (mlm *MeshLambertMaterial) SetEmissiveIntensity(v float64) {
	mlm.Set("emissiveIntensity", v)
}
func (mlm *MeshLambertMaterial) EmissiveMap() *Texture {
	return &Texture{Value: mlm.Get("emissiveMap")}
}
func (mlm *MeshLambertMaterial) SetEmissiveMap(v *Texture) {
	mlm.Set("emissiveMap", v)
}
func (mlm *MeshLambertMaterial) EnvMap() *Texture {
	return &Texture{Value: mlm.Get("envMap")}
}
func (mlm *MeshLambertMaterial) SetEnvMap(v *Texture) {
	mlm.Set("envMap", v)
}
func (mlm *MeshLambertMaterial) FlatShading() bool {
	return mlm.Get("flatShading").Bool()
}
func (mlm *MeshLambertMaterial) SetFlatShading(v bool) {
	mlm.Set("flatShading", v)
}
func (mlm *MeshLambertMaterial) Fog() bool {
	return mlm.Get("fog").Bool()
}
func (mlm *MeshLambertMaterial) SetFog(v bool) {
	mlm.Set("fog", v)
}
func (mlm *MeshLambertMaterial) Id() int {
	return mlm.Get("id").Int()
}
func (mlm *MeshLambertMaterial) SetId(v int) {
	mlm.Set("id", v)
}
func (mlm *MeshLambertMaterial) IsMaterial() bool {
	return mlm.Get("isMaterial").Bool()
}
func (mlm *MeshLambertMaterial) SetIsMaterial(v bool) {
	mlm.Set("isMaterial", v)
}
func (mlm *MeshLambertMaterial) LightMap() *Texture {
	return &Texture{Value: mlm.Get("lightMap")}
}
func (mlm *MeshLambertMaterial) SetLightMap(v *Texture) {
	mlm.Set("lightMap", v)
}
func (mlm *MeshLambertMaterial) LightMapIntensity() float64 {
	return mlm.Get("lightMapIntensity").Float()
}
func (mlm *MeshLambertMaterial) SetLightMapIntensity(v float64) {
	mlm.Set("lightMapIntensity", v)
}
func (mlm *MeshLambertMaterial) Lights() bool {
	return mlm.Get("lights").Bool()
}
func (mlm *MeshLambertMaterial) SetLights(v bool) {
	mlm.Set("lights", v)
}
func (mlm *MeshLambertMaterial) Map() *Texture {
	return &Texture{Value: mlm.Get("map")}
}
func (mlm *MeshLambertMaterial) SetMap(v *Texture) {
	mlm.Set("map", v)
}
func (mlm *MeshLambertMaterial) MorphNormals() bool {
	return mlm.Get("morphNormals").Bool()
}
func (mlm *MeshLambertMaterial) SetMorphNormals(v bool) {
	mlm.Set("morphNormals", v)
}
func (mlm *MeshLambertMaterial) MorphTargets() bool {
	return mlm.Get("morphTargets").Bool()
}
func (mlm *MeshLambertMaterial) SetMorphTargets(v bool) {
	mlm.Set("morphTargets", v)
}
func (mlm *MeshLambertMaterial) Name() string {
	return mlm.Get("name").String()
}
func (mlm *MeshLambertMaterial) SetName(v string) {
	mlm.Set("name", v)
}
func (mlm *MeshLambertMaterial) NeedsUpdate() bool {
	return mlm.Get("needsUpdate").Bool()
}
func (mlm *MeshLambertMaterial) SetNeedsUpdate(v bool) {
	mlm.Set("needsUpdate", v)
}
func (mlm *MeshLambertMaterial) Opacity() float64 {
	return mlm.Get("opacity").Float()
}
func (mlm *MeshLambertMaterial) SetOpacity(v float64) {
	mlm.Set("opacity", v)
}
func (mlm *MeshLambertMaterial) Overdraw() float64 {
	return mlm.Get("overdraw").Float()
}
func (mlm *MeshLambertMaterial) SetOverdraw(v float64) {
	mlm.Set("overdraw", v)
}
func (mlm *MeshLambertMaterial) PolygonOffset() bool {
	return mlm.Get("polygonOffset").Bool()
}
func (mlm *MeshLambertMaterial) SetPolygonOffset(v bool) {
	mlm.Set("polygonOffset", v)
}
func (mlm *MeshLambertMaterial) PolygonOffsetFactor() float64 {
	return mlm.Get("polygonOffsetFactor").Float()
}
func (mlm *MeshLambertMaterial) SetPolygonOffsetFactor(v float64) {
	mlm.Set("polygonOffsetFactor", v)
}
func (mlm *MeshLambertMaterial) PolygonOffsetUnits() float64 {
	return mlm.Get("polygonOffsetUnits").Float()
}
func (mlm *MeshLambertMaterial) SetPolygonOffsetUnits(v float64) {
	mlm.Set("polygonOffsetUnits", v)
}
func (mlm *MeshLambertMaterial) Precision() string {
	return mlm.Get("precision").String()
}
func (mlm *MeshLambertMaterial) SetPrecision(v string) {
	mlm.Set("precision", v)
}
func (mlm *MeshLambertMaterial) PremultipliedAlpha() bool {
	return mlm.Get("premultipliedAlpha").Bool()
}
func (mlm *MeshLambertMaterial) SetPremultipliedAlpha(v bool) {
	mlm.Set("premultipliedAlpha", v)
}
func (mlm *MeshLambertMaterial) Reflectivity() float64 {
	return mlm.Get("reflectivity").Float()
}
func (mlm *MeshLambertMaterial) SetReflectivity(v float64) {
	mlm.Set("reflectivity", v)
}
func (mlm *MeshLambertMaterial) RefractionRatio() float64 {
	return mlm.Get("refractionRatio").Float()
}
func (mlm *MeshLambertMaterial) SetRefractionRatio(v float64) {
	mlm.Set("refractionRatio", v)
}
func (mlm *MeshLambertMaterial) Side() Side {
	return Side(mlm.Get("side"))
}
func (mlm *MeshLambertMaterial) SetSide(v Side) {
	mlm.Set("side", v)
}
func (mlm *MeshLambertMaterial) Skinning() bool {
	return mlm.Get("skinning").Bool()
}
func (mlm *MeshLambertMaterial) SetSkinning(v bool) {
	mlm.Set("skinning", v)
}
func (mlm *MeshLambertMaterial) SpecularMap() *Texture {
	return &Texture{Value: mlm.Get("specularMap")}
}
func (mlm *MeshLambertMaterial) SetSpecularMap(v *Texture) {
	mlm.Set("specularMap", v)
}
func (mlm *MeshLambertMaterial) Transparent() bool {
	return mlm.Get("transparent").Bool()
}
func (mlm *MeshLambertMaterial) SetTransparent(v bool) {
	mlm.Set("transparent", v)
}
func (mlm *MeshLambertMaterial) Type() string {
	return mlm.Get("type").String()
}
func (mlm *MeshLambertMaterial) SetType(v string) {
	mlm.Set("type", v)
}
func (mlm *MeshLambertMaterial) UserData() js.Value {
	return mlm.Get("userData")
}
func (mlm *MeshLambertMaterial) SetUserData(v js.Value) {
	mlm.Set("userData", v)
}
func (mlm *MeshLambertMaterial) Uuid() string {
	return mlm.Get("uuid").String()
}
func (mlm *MeshLambertMaterial) SetUuid(v string) {
	mlm.Set("uuid", v)
}
func (mlm *MeshLambertMaterial) VertexColors() Colors {
	return Colors(mlm.Get("vertexColors"))
}
func (mlm *MeshLambertMaterial) SetVertexColors(v Colors) {
	mlm.Set("vertexColors", v)
}
func (mlm *MeshLambertMaterial) VertexTangents() bool {
	return mlm.Get("vertexTangents").Bool()
}
func (mlm *MeshLambertMaterial) SetVertexTangents(v bool) {
	mlm.Set("vertexTangents", v)
}
func (mlm *MeshLambertMaterial) Visible() bool {
	return mlm.Get("visible").Bool()
}
func (mlm *MeshLambertMaterial) SetVisible(v bool) {
	mlm.Set("visible", v)
}
func (mlm *MeshLambertMaterial) Wireframe() bool {
	return mlm.Get("wireframe").Bool()
}
func (mlm *MeshLambertMaterial) SetWireframe(v bool) {
	mlm.Set("wireframe", v)
}
func (mlm *MeshLambertMaterial) WireframeLinecap() string {
	return mlm.Get("wireframeLinecap").String()
}
func (mlm *MeshLambertMaterial) SetWireframeLinecap(v string) {
	mlm.Set("wireframeLinecap", v)
}
func (mlm *MeshLambertMaterial) WireframeLinejoin() string {
	return mlm.Get("wireframeLinejoin").String()
}
func (mlm *MeshLambertMaterial) SetWireframeLinejoin(v string) {
	mlm.Set("wireframeLinejoin", v)
}
func (mlm *MeshLambertMaterial) WireframeLinewidth() float64 {
	return mlm.Get("wireframeLinewidth").Float()
}
func (mlm *MeshLambertMaterial) SetWireframeLinewidth(v float64) {
	mlm.Set("wireframeLinewidth", v)
}
func (mlm *MeshLambertMaterial) AddEventListener(typ string, listener js.Value) {
	mlm.Call("addEventListener", typ, listener)
}
func (mlm *MeshLambertMaterial) Clone() *MeshLambertMaterial {
	return &MeshLambertMaterial{Value: mlm.Call("clone")}
}
func (mlm *MeshLambertMaterial) Copy(material *Material) *MeshLambertMaterial {
	return &MeshLambertMaterial{Value: mlm.Call("copy", material)}
}
func (mlm *MeshLambertMaterial) DispatchEvent(event js.Value) {
	mlm.Call("dispatchEvent", event)
}
func (mlm *MeshLambertMaterial) Dispose() {
	mlm.Call("dispose")
}
func (mlm *MeshLambertMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mlm.Call("hasEventListener", typ, listener).Bool()
}
func (mlm *MeshLambertMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mlm.Call("onBeforeCompile", shader, renderer)
}
func (mlm *MeshLambertMaterial) RemoveEventListener(typ string, listener js.Value) {
	mlm.Call("removeEventListener", typ, listener)
}
func (mlm *MeshLambertMaterial) SetValues(parameters MeshLambertMaterialParameters) {
	mlm.Call("setValues", parameters)
}
func (mlm *MeshLambertMaterial) ToJSON(meta js.Value) js.Value {
	return mlm.Call("toJSON", meta)
}
func (mlm *MeshLambertMaterial) Update() {
	mlm.Call("update")
}

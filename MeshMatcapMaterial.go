// Code generated from "materials/MeshMatcapMaterial.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type MeshMatcapMaterialParameters interface {
}
type MeshMatcapMaterial struct {
	js.Value
}

func NewMeshMatcapMaterial(parameters MeshMatcapMaterialParameters) *MeshMatcapMaterial {
	return &MeshMatcapMaterial{Value: get("MeshMatcapMaterial").New(parameters)}
}
func (mmm *MeshMatcapMaterial) AlphaMap() *Texture {
	return &Texture{Value: mmm.Get("alphaMap")}
}
func (mmm *MeshMatcapMaterial) SetAlphaMap(v *Texture) {
	mmm.Set("alphaMap", v)
}
func (mmm *MeshMatcapMaterial) AlphaTest() int {
	return mmm.Get("alphaTest").Int()
}
func (mmm *MeshMatcapMaterial) SetAlphaTest(v int) {
	mmm.Set("alphaTest", v)
}
func (mmm *MeshMatcapMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mmm.Get("blendDst"))
}
func (mmm *MeshMatcapMaterial) SetBlendDst(v BlendingDstFactor) {
	mmm.Set("blendDst", v)
}
func (mmm *MeshMatcapMaterial) BlendDstAlpha() int {
	return mmm.Get("blendDstAlpha").Int()
}
func (mmm *MeshMatcapMaterial) SetBlendDstAlpha(v int) {
	mmm.Set("blendDstAlpha", v)
}
func (mmm *MeshMatcapMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mmm.Get("blendEquation"))
}
func (mmm *MeshMatcapMaterial) SetBlendEquation(v BlendingEquation) {
	mmm.Set("blendEquation", v)
}
func (mmm *MeshMatcapMaterial) BlendEquationAlpha() int {
	return mmm.Get("blendEquationAlpha").Int()
}
func (mmm *MeshMatcapMaterial) SetBlendEquationAlpha(v int) {
	mmm.Set("blendEquationAlpha", v)
}
func (mmm *MeshMatcapMaterial) BlendSrc() js.Value {
	return mmm.Get("blendSrc")
}
func (mmm *MeshMatcapMaterial) SetBlendSrc(v js.Value) {
	mmm.Set("blendSrc", v)
}
func (mmm *MeshMatcapMaterial) BlendSrcAlpha() int {
	return mmm.Get("blendSrcAlpha").Int()
}
func (mmm *MeshMatcapMaterial) SetBlendSrcAlpha(v int) {
	mmm.Set("blendSrcAlpha", v)
}
func (mmm *MeshMatcapMaterial) Blending() Blending {
	return Blending(mmm.Get("blending"))
}
func (mmm *MeshMatcapMaterial) SetBlending(v Blending) {
	mmm.Set("blending", v)
}
func (mmm *MeshMatcapMaterial) BumpMap() *Texture {
	return &Texture{Value: mmm.Get("bumpMap")}
}
func (mmm *MeshMatcapMaterial) SetBumpMap(v *Texture) {
	mmm.Set("bumpMap", v)
}
func (mmm *MeshMatcapMaterial) BumpScale() int {
	return mmm.Get("bumpScale").Int()
}
func (mmm *MeshMatcapMaterial) SetBumpScale(v int) {
	mmm.Set("bumpScale", v)
}
func (mmm *MeshMatcapMaterial) ClipIntersection() bool {
	return mmm.Get("clipIntersection").Bool()
}
func (mmm *MeshMatcapMaterial) SetClipIntersection(v bool) {
	mmm.Set("clipIntersection", v)
}
func (mmm *MeshMatcapMaterial) ClipShadows() bool {
	return mmm.Get("clipShadows").Bool()
}
func (mmm *MeshMatcapMaterial) SetClipShadows(v bool) {
	mmm.Set("clipShadows", v)
}
func (mmm *MeshMatcapMaterial) ClippingPlanes() js.Value {
	return mmm.Get("clippingPlanes")
}
func (mmm *MeshMatcapMaterial) SetClippingPlanes(v js.Value) {
	mmm.Set("clippingPlanes", v)
}
func (mmm *MeshMatcapMaterial) Color() *Color {
	return &Color{Value: mmm.Get("color")}
}
func (mmm *MeshMatcapMaterial) SetColor(v *Color) {
	mmm.Set("color", v)
}
func (mmm *MeshMatcapMaterial) ColorWrite() bool {
	return mmm.Get("colorWrite").Bool()
}
func (mmm *MeshMatcapMaterial) SetColorWrite(v bool) {
	mmm.Set("colorWrite", v)
}
func (mmm *MeshMatcapMaterial) DepthFunc() DepthModes {
	return DepthModes(mmm.Get("depthFunc"))
}
func (mmm *MeshMatcapMaterial) SetDepthFunc(v DepthModes) {
	mmm.Set("depthFunc", v)
}
func (mmm *MeshMatcapMaterial) DepthTest() bool {
	return mmm.Get("depthTest").Bool()
}
func (mmm *MeshMatcapMaterial) SetDepthTest(v bool) {
	mmm.Set("depthTest", v)
}
func (mmm *MeshMatcapMaterial) DepthWrite() bool {
	return mmm.Get("depthWrite").Bool()
}
func (mmm *MeshMatcapMaterial) SetDepthWrite(v bool) {
	mmm.Set("depthWrite", v)
}
func (mmm *MeshMatcapMaterial) DisplacementBias() int {
	return mmm.Get("displacementBias").Int()
}
func (mmm *MeshMatcapMaterial) SetDisplacementBias(v int) {
	mmm.Set("displacementBias", v)
}
func (mmm *MeshMatcapMaterial) DisplacementMap() *Texture {
	return &Texture{Value: mmm.Get("displacementMap")}
}
func (mmm *MeshMatcapMaterial) SetDisplacementMap(v *Texture) {
	mmm.Set("displacementMap", v)
}
func (mmm *MeshMatcapMaterial) DisplacementScale() int {
	return mmm.Get("displacementScale").Int()
}
func (mmm *MeshMatcapMaterial) SetDisplacementScale(v int) {
	mmm.Set("displacementScale", v)
}
func (mmm *MeshMatcapMaterial) Dithering() bool {
	return mmm.Get("dithering").Bool()
}
func (mmm *MeshMatcapMaterial) SetDithering(v bool) {
	mmm.Set("dithering", v)
}
func (mmm *MeshMatcapMaterial) FlatShading() bool {
	return mmm.Get("flatShading").Bool()
}
func (mmm *MeshMatcapMaterial) SetFlatShading(v bool) {
	mmm.Set("flatShading", v)
}
func (mmm *MeshMatcapMaterial) Fog() bool {
	return mmm.Get("fog").Bool()
}
func (mmm *MeshMatcapMaterial) SetFog(v bool) {
	mmm.Set("fog", v)
}
func (mmm *MeshMatcapMaterial) Id() int {
	return mmm.Get("id").Int()
}
func (mmm *MeshMatcapMaterial) SetId(v int) {
	mmm.Set("id", v)
}
func (mmm *MeshMatcapMaterial) IsMaterial() bool {
	return mmm.Get("isMaterial").Bool()
}
func (mmm *MeshMatcapMaterial) SetIsMaterial(v bool) {
	mmm.Set("isMaterial", v)
}
func (mmm *MeshMatcapMaterial) Lights() bool {
	return mmm.Get("lights").Bool()
}
func (mmm *MeshMatcapMaterial) SetLights(v bool) {
	mmm.Set("lights", v)
}
func (mmm *MeshMatcapMaterial) Map() *Texture {
	return &Texture{Value: mmm.Get("map")}
}
func (mmm *MeshMatcapMaterial) SetMap(v *Texture) {
	mmm.Set("map", v)
}
func (mmm *MeshMatcapMaterial) MatMap() *Texture {
	return &Texture{Value: mmm.Get("matMap")}
}
func (mmm *MeshMatcapMaterial) SetMatMap(v *Texture) {
	mmm.Set("matMap", v)
}
func (mmm *MeshMatcapMaterial) MorphNormals() bool {
	return mmm.Get("morphNormals").Bool()
}
func (mmm *MeshMatcapMaterial) SetMorphNormals(v bool) {
	mmm.Set("morphNormals", v)
}
func (mmm *MeshMatcapMaterial) MorphTargets() bool {
	return mmm.Get("morphTargets").Bool()
}
func (mmm *MeshMatcapMaterial) SetMorphTargets(v bool) {
	mmm.Set("morphTargets", v)
}
func (mmm *MeshMatcapMaterial) Name() string {
	return mmm.Get("name").String()
}
func (mmm *MeshMatcapMaterial) SetName(v string) {
	mmm.Set("name", v)
}
func (mmm *MeshMatcapMaterial) NeedsUpdate() bool {
	return mmm.Get("needsUpdate").Bool()
}
func (mmm *MeshMatcapMaterial) SetNeedsUpdate(v bool) {
	mmm.Set("needsUpdate", v)
}
func (mmm *MeshMatcapMaterial) NormalMap() *Texture {
	return &Texture{Value: mmm.Get("normalMap")}
}
func (mmm *MeshMatcapMaterial) SetNormalMap(v *Texture) {
	mmm.Set("normalMap", v)
}
func (mmm *MeshMatcapMaterial) NormalMapType() NormalMapTypes {
	return NormalMapTypes(mmm.Get("normalMapType"))
}
func (mmm *MeshMatcapMaterial) SetNormalMapType(v NormalMapTypes) {
	mmm.Set("normalMapType", v)
}
func (mmm *MeshMatcapMaterial) NormalScale() *Vector2 {
	return &Vector2{Value: mmm.Get("normalScale")}
}
func (mmm *MeshMatcapMaterial) SetNormalScale(v *Vector2) {
	mmm.Set("normalScale", v)
}
func (mmm *MeshMatcapMaterial) Opacity() int {
	return mmm.Get("opacity").Int()
}
func (mmm *MeshMatcapMaterial) SetOpacity(v int) {
	mmm.Set("opacity", v)
}
func (mmm *MeshMatcapMaterial) Overdraw() int {
	return mmm.Get("overdraw").Int()
}
func (mmm *MeshMatcapMaterial) SetOverdraw(v int) {
	mmm.Set("overdraw", v)
}
func (mmm *MeshMatcapMaterial) PolygonOffset() bool {
	return mmm.Get("polygonOffset").Bool()
}
func (mmm *MeshMatcapMaterial) SetPolygonOffset(v bool) {
	mmm.Set("polygonOffset", v)
}
func (mmm *MeshMatcapMaterial) PolygonOffsetFactor() int {
	return mmm.Get("polygonOffsetFactor").Int()
}
func (mmm *MeshMatcapMaterial) SetPolygonOffsetFactor(v int) {
	mmm.Set("polygonOffsetFactor", v)
}
func (mmm *MeshMatcapMaterial) PolygonOffsetUnits() int {
	return mmm.Get("polygonOffsetUnits").Int()
}
func (mmm *MeshMatcapMaterial) SetPolygonOffsetUnits(v int) {
	mmm.Set("polygonOffsetUnits", v)
}
func (mmm *MeshMatcapMaterial) Precision() string {
	return mmm.Get("precision").String()
}
func (mmm *MeshMatcapMaterial) SetPrecision(v string) {
	mmm.Set("precision", v)
}
func (mmm *MeshMatcapMaterial) PremultipliedAlpha() bool {
	return mmm.Get("premultipliedAlpha").Bool()
}
func (mmm *MeshMatcapMaterial) SetPremultipliedAlpha(v bool) {
	mmm.Set("premultipliedAlpha", v)
}
func (mmm *MeshMatcapMaterial) Side() Side {
	return Side(mmm.Get("side"))
}
func (mmm *MeshMatcapMaterial) SetSide(v Side) {
	mmm.Set("side", v)
}
func (mmm *MeshMatcapMaterial) Skinning() bool {
	return mmm.Get("skinning").Bool()
}
func (mmm *MeshMatcapMaterial) SetSkinning(v bool) {
	mmm.Set("skinning", v)
}
func (mmm *MeshMatcapMaterial) Transparent() bool {
	return mmm.Get("transparent").Bool()
}
func (mmm *MeshMatcapMaterial) SetTransparent(v bool) {
	mmm.Set("transparent", v)
}
func (mmm *MeshMatcapMaterial) Type() string {
	return mmm.Get("type").String()
}
func (mmm *MeshMatcapMaterial) SetType(v string) {
	mmm.Set("type", v)
}
func (mmm *MeshMatcapMaterial) UserData() js.Value {
	return mmm.Get("userData")
}
func (mmm *MeshMatcapMaterial) SetUserData(v js.Value) {
	mmm.Set("userData", v)
}
func (mmm *MeshMatcapMaterial) Uuid() string {
	return mmm.Get("uuid").String()
}
func (mmm *MeshMatcapMaterial) SetUuid(v string) {
	mmm.Set("uuid", v)
}
func (mmm *MeshMatcapMaterial) VertexColors() Colors {
	return Colors(mmm.Get("vertexColors"))
}
func (mmm *MeshMatcapMaterial) SetVertexColors(v Colors) {
	mmm.Set("vertexColors", v)
}
func (mmm *MeshMatcapMaterial) VertexTangents() bool {
	return mmm.Get("vertexTangents").Bool()
}
func (mmm *MeshMatcapMaterial) SetVertexTangents(v bool) {
	mmm.Set("vertexTangents", v)
}
func (mmm *MeshMatcapMaterial) Visible() bool {
	return mmm.Get("visible").Bool()
}
func (mmm *MeshMatcapMaterial) SetVisible(v bool) {
	mmm.Set("visible", v)
}
func (mmm *MeshMatcapMaterial) AddEventListener(typ string, listener js.Value) {
	mmm.Call("addEventListener", typ, listener)
}
func (mmm *MeshMatcapMaterial) Clone() *MeshMatcapMaterial {
	return &MeshMatcapMaterial{Value: mmm.Call("clone")}
}
func (mmm *MeshMatcapMaterial) Copy(material *Material) *MeshMatcapMaterial {
	return &MeshMatcapMaterial{Value: mmm.Call("copy", material)}
}
func (mmm *MeshMatcapMaterial) DispatchEvent(event js.Value) {
	mmm.Call("dispatchEvent", event)
}
func (mmm *MeshMatcapMaterial) Dispose() {
	mmm.Call("dispose")
}
func (mmm *MeshMatcapMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mmm.Call("hasEventListener", typ, listener).Bool()
}
func (mmm *MeshMatcapMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mmm.Call("onBeforeCompile", shader, renderer)
}
func (mmm *MeshMatcapMaterial) RemoveEventListener(typ string, listener js.Value) {
	mmm.Call("removeEventListener", typ, listener)
}
func (mmm *MeshMatcapMaterial) SetValues(parameters MeshMatcapMaterialParameters) {
	mmm.Call("setValues", parameters)
}
func (mmm *MeshMatcapMaterial) ToJSON(meta js.Value) js.Value {
	return mmm.Call("toJSON", meta)
}
func (mmm *MeshMatcapMaterial) Update() {
	mmm.Call("update")
}

// Code generated from "materials/MeshDistanceMaterial.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type MeshDistanceMaterialParameters interface {
}
type MeshDistanceMaterial struct {
	js.Value
}

func NewMeshDistanceMaterial(parameters MeshDistanceMaterialParameters) *MeshDistanceMaterial {
	return &MeshDistanceMaterial{Value: get("MeshDistanceMaterial").New(parameters)}
}
func (mdm *MeshDistanceMaterial) AlphaTest() int {
	return mdm.Get("alphaTest").Int()
}
func (mdm *MeshDistanceMaterial) SetAlphaTest(v int) {
	mdm.Set("alphaTest", v)
}
func (mdm *MeshDistanceMaterial) BlendDst() BlendingDstFactor {
	return BlendingDstFactor(mdm.Get("blendDst"))
}
func (mdm *MeshDistanceMaterial) SetBlendDst(v BlendingDstFactor) {
	mdm.Set("blendDst", v)
}
func (mdm *MeshDistanceMaterial) BlendDstAlpha() int {
	return mdm.Get("blendDstAlpha").Int()
}
func (mdm *MeshDistanceMaterial) SetBlendDstAlpha(v int) {
	mdm.Set("blendDstAlpha", v)
}
func (mdm *MeshDistanceMaterial) BlendEquation() BlendingEquation {
	return BlendingEquation(mdm.Get("blendEquation"))
}
func (mdm *MeshDistanceMaterial) SetBlendEquation(v BlendingEquation) {
	mdm.Set("blendEquation", v)
}
func (mdm *MeshDistanceMaterial) BlendEquationAlpha() int {
	return mdm.Get("blendEquationAlpha").Int()
}
func (mdm *MeshDistanceMaterial) SetBlendEquationAlpha(v int) {
	mdm.Set("blendEquationAlpha", v)
}
func (mdm *MeshDistanceMaterial) BlendSrc() js.Value {
	return mdm.Get("blendSrc")
}
func (mdm *MeshDistanceMaterial) SetBlendSrc(v js.Value) {
	mdm.Set("blendSrc", v)
}
func (mdm *MeshDistanceMaterial) BlendSrcAlpha() int {
	return mdm.Get("blendSrcAlpha").Int()
}
func (mdm *MeshDistanceMaterial) SetBlendSrcAlpha(v int) {
	mdm.Set("blendSrcAlpha", v)
}
func (mdm *MeshDistanceMaterial) Blending() Blending {
	return Blending(mdm.Get("blending"))
}
func (mdm *MeshDistanceMaterial) SetBlending(v Blending) {
	mdm.Set("blending", v)
}
func (mdm *MeshDistanceMaterial) ClipIntersection() bool {
	return mdm.Get("clipIntersection").Bool()
}
func (mdm *MeshDistanceMaterial) SetClipIntersection(v bool) {
	mdm.Set("clipIntersection", v)
}
func (mdm *MeshDistanceMaterial) ClipShadows() bool {
	return mdm.Get("clipShadows").Bool()
}
func (mdm *MeshDistanceMaterial) SetClipShadows(v bool) {
	mdm.Set("clipShadows", v)
}
func (mdm *MeshDistanceMaterial) ClippingPlanes() js.Value {
	return mdm.Get("clippingPlanes")
}
func (mdm *MeshDistanceMaterial) SetClippingPlanes(v js.Value) {
	mdm.Set("clippingPlanes", v)
}
func (mdm *MeshDistanceMaterial) ColorWrite() bool {
	return mdm.Get("colorWrite").Bool()
}
func (mdm *MeshDistanceMaterial) SetColorWrite(v bool) {
	mdm.Set("colorWrite", v)
}
func (mdm *MeshDistanceMaterial) DepthFunc() DepthModes {
	return DepthModes(mdm.Get("depthFunc"))
}
func (mdm *MeshDistanceMaterial) SetDepthFunc(v DepthModes) {
	mdm.Set("depthFunc", v)
}
func (mdm *MeshDistanceMaterial) DepthTest() bool {
	return mdm.Get("depthTest").Bool()
}
func (mdm *MeshDistanceMaterial) SetDepthTest(v bool) {
	mdm.Set("depthTest", v)
}
func (mdm *MeshDistanceMaterial) DepthWrite() bool {
	return mdm.Get("depthWrite").Bool()
}
func (mdm *MeshDistanceMaterial) SetDepthWrite(v bool) {
	mdm.Set("depthWrite", v)
}
func (mdm *MeshDistanceMaterial) DisplacementBias() int {
	return mdm.Get("displacementBias").Int()
}
func (mdm *MeshDistanceMaterial) SetDisplacementBias(v int) {
	mdm.Set("displacementBias", v)
}
func (mdm *MeshDistanceMaterial) DisplacementMap() *Texture {
	return &Texture{Value: mdm.Get("displacementMap")}
}
func (mdm *MeshDistanceMaterial) SetDisplacementMap(v *Texture) {
	mdm.Set("displacementMap", v)
}
func (mdm *MeshDistanceMaterial) DisplacementScale() int {
	return mdm.Get("displacementScale").Int()
}
func (mdm *MeshDistanceMaterial) SetDisplacementScale(v int) {
	mdm.Set("displacementScale", v)
}
func (mdm *MeshDistanceMaterial) Dithering() bool {
	return mdm.Get("dithering").Bool()
}
func (mdm *MeshDistanceMaterial) SetDithering(v bool) {
	mdm.Set("dithering", v)
}
func (mdm *MeshDistanceMaterial) FarDistance() int {
	return mdm.Get("farDistance").Int()
}
func (mdm *MeshDistanceMaterial) SetFarDistance(v int) {
	mdm.Set("farDistance", v)
}
func (mdm *MeshDistanceMaterial) FlatShading() bool {
	return mdm.Get("flatShading").Bool()
}
func (mdm *MeshDistanceMaterial) SetFlatShading(v bool) {
	mdm.Set("flatShading", v)
}
func (mdm *MeshDistanceMaterial) Fog() bool {
	return mdm.Get("fog").Bool()
}
func (mdm *MeshDistanceMaterial) SetFog(v bool) {
	mdm.Set("fog", v)
}
func (mdm *MeshDistanceMaterial) Id() int {
	return mdm.Get("id").Int()
}
func (mdm *MeshDistanceMaterial) SetId(v int) {
	mdm.Set("id", v)
}
func (mdm *MeshDistanceMaterial) IsMaterial() bool {
	return mdm.Get("isMaterial").Bool()
}
func (mdm *MeshDistanceMaterial) SetIsMaterial(v bool) {
	mdm.Set("isMaterial", v)
}
func (mdm *MeshDistanceMaterial) Lights() bool {
	return mdm.Get("lights").Bool()
}
func (mdm *MeshDistanceMaterial) SetLights(v bool) {
	mdm.Set("lights", v)
}
func (mdm *MeshDistanceMaterial) Name() string {
	return mdm.Get("name").String()
}
func (mdm *MeshDistanceMaterial) SetName(v string) {
	mdm.Set("name", v)
}
func (mdm *MeshDistanceMaterial) NearDistance() int {
	return mdm.Get("nearDistance").Int()
}
func (mdm *MeshDistanceMaterial) SetNearDistance(v int) {
	mdm.Set("nearDistance", v)
}
func (mdm *MeshDistanceMaterial) NeedsUpdate() bool {
	return mdm.Get("needsUpdate").Bool()
}
func (mdm *MeshDistanceMaterial) SetNeedsUpdate(v bool) {
	mdm.Set("needsUpdate", v)
}
func (mdm *MeshDistanceMaterial) Opacity() int {
	return mdm.Get("opacity").Int()
}
func (mdm *MeshDistanceMaterial) SetOpacity(v int) {
	mdm.Set("opacity", v)
}
func (mdm *MeshDistanceMaterial) Overdraw() int {
	return mdm.Get("overdraw").Int()
}
func (mdm *MeshDistanceMaterial) SetOverdraw(v int) {
	mdm.Set("overdraw", v)
}
func (mdm *MeshDistanceMaterial) PolygonOffset() bool {
	return mdm.Get("polygonOffset").Bool()
}
func (mdm *MeshDistanceMaterial) SetPolygonOffset(v bool) {
	mdm.Set("polygonOffset", v)
}
func (mdm *MeshDistanceMaterial) PolygonOffsetFactor() int {
	return mdm.Get("polygonOffsetFactor").Int()
}
func (mdm *MeshDistanceMaterial) SetPolygonOffsetFactor(v int) {
	mdm.Set("polygonOffsetFactor", v)
}
func (mdm *MeshDistanceMaterial) PolygonOffsetUnits() int {
	return mdm.Get("polygonOffsetUnits").Int()
}
func (mdm *MeshDistanceMaterial) SetPolygonOffsetUnits(v int) {
	mdm.Set("polygonOffsetUnits", v)
}
func (mdm *MeshDistanceMaterial) Precision() string {
	return mdm.Get("precision").String()
}
func (mdm *MeshDistanceMaterial) SetPrecision(v string) {
	mdm.Set("precision", v)
}
func (mdm *MeshDistanceMaterial) PremultipliedAlpha() bool {
	return mdm.Get("premultipliedAlpha").Bool()
}
func (mdm *MeshDistanceMaterial) SetPremultipliedAlpha(v bool) {
	mdm.Set("premultipliedAlpha", v)
}
func (mdm *MeshDistanceMaterial) ReferencePosition() *Vector3 {
	return &Vector3{Value: mdm.Get("referencePosition")}
}
func (mdm *MeshDistanceMaterial) SetReferencePosition(v *Vector3) {
	mdm.Set("referencePosition", v)
}
func (mdm *MeshDistanceMaterial) Side() Side {
	return Side(mdm.Get("side"))
}
func (mdm *MeshDistanceMaterial) SetSide(v Side) {
	mdm.Set("side", v)
}
func (mdm *MeshDistanceMaterial) Transparent() bool {
	return mdm.Get("transparent").Bool()
}
func (mdm *MeshDistanceMaterial) SetTransparent(v bool) {
	mdm.Set("transparent", v)
}
func (mdm *MeshDistanceMaterial) Type() string {
	return mdm.Get("type").String()
}
func (mdm *MeshDistanceMaterial) SetType(v string) {
	mdm.Set("type", v)
}
func (mdm *MeshDistanceMaterial) UserData() js.Value {
	return mdm.Get("userData")
}
func (mdm *MeshDistanceMaterial) SetUserData(v js.Value) {
	mdm.Set("userData", v)
}
func (mdm *MeshDistanceMaterial) Uuid() string {
	return mdm.Get("uuid").String()
}
func (mdm *MeshDistanceMaterial) SetUuid(v string) {
	mdm.Set("uuid", v)
}
func (mdm *MeshDistanceMaterial) VertexColors() Colors {
	return Colors(mdm.Get("vertexColors"))
}
func (mdm *MeshDistanceMaterial) SetVertexColors(v Colors) {
	mdm.Set("vertexColors", v)
}
func (mdm *MeshDistanceMaterial) VertexTangents() bool {
	return mdm.Get("vertexTangents").Bool()
}
func (mdm *MeshDistanceMaterial) SetVertexTangents(v bool) {
	mdm.Set("vertexTangents", v)
}
func (mdm *MeshDistanceMaterial) Visible() bool {
	return mdm.Get("visible").Bool()
}
func (mdm *MeshDistanceMaterial) SetVisible(v bool) {
	mdm.Set("visible", v)
}
func (mdm *MeshDistanceMaterial) AddEventListener(typ string, listener js.Value) {
	mdm.Call("addEventListener", typ, listener)
}
func (mdm *MeshDistanceMaterial) Clone() *MeshDistanceMaterial {
	return &MeshDistanceMaterial{Value: mdm.Call("clone")}
}
func (mdm *MeshDistanceMaterial) Copy(material *Material) *MeshDistanceMaterial {
	return &MeshDistanceMaterial{Value: mdm.Call("copy", material)}
}
func (mdm *MeshDistanceMaterial) DispatchEvent(event js.Value) {
	mdm.Call("dispatchEvent", event)
}
func (mdm *MeshDistanceMaterial) Dispose() {
	mdm.Call("dispose")
}
func (mdm *MeshDistanceMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mdm.Call("hasEventListener", typ, listener).Bool()
}
func (mdm *MeshDistanceMaterial) OnBeforeCompile(shader js.Value, renderer *WebGLRenderer) {
	mdm.Call("onBeforeCompile", shader, renderer)
}
func (mdm *MeshDistanceMaterial) RemoveEventListener(typ string, listener js.Value) {
	mdm.Call("removeEventListener", typ, listener)
}
func (mdm *MeshDistanceMaterial) SetValues(parameters MeshDistanceMaterialParameters) {
	mdm.Call("setValues", parameters)
}
func (mdm *MeshDistanceMaterial) ToJSON(meta js.Value) js.Value {
	return mdm.Call("toJSON", meta)
}
func (mdm *MeshDistanceMaterial) Update() {
	mdm.Call("update")
}

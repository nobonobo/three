package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
)

type MeshNormalMaterialParameters interface {
}
type MeshNormalMaterial struct {
	js.Value
}

func NewMeshNormalMaterial(parameters MeshNormalMaterialParameters) *MeshNormalMaterial {
	return &MeshNormalMaterial{Value: get("MeshNormalMaterial").New(parameters)}
}
func (mnm *MeshNormalMaterial) AlphaTest() int {
	return mnm.Get("alphaTest").Int()
}
func (mnm *MeshNormalMaterial) SetAlphaTest(v int) {
	mnm.Set("alphaTest", v)
}
func (mnm *MeshNormalMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: mnm.Get("blendDst")}
}
func (mnm *MeshNormalMaterial) SetBlendDst(v *BlendingDstFactor) {
	mnm.Set("blendDst", v)
}
func (mnm *MeshNormalMaterial) BlendDstAlpha() int {
	return mnm.Get("blendDstAlpha").Int()
}
func (mnm *MeshNormalMaterial) SetBlendDstAlpha(v int) {
	mnm.Set("blendDstAlpha", v)
}
func (mnm *MeshNormalMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: mnm.Get("blendEquation")}
}
func (mnm *MeshNormalMaterial) SetBlendEquation(v *BlendingEquation) {
	mnm.Set("blendEquation", v)
}
func (mnm *MeshNormalMaterial) BlendEquationAlpha() int {
	return mnm.Get("blendEquationAlpha").Int()
}
func (mnm *MeshNormalMaterial) SetBlendEquationAlpha(v int) {
	mnm.Set("blendEquationAlpha", v)
}
func (mnm *MeshNormalMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(mnm.Get("blendSrc"))
}
func (mnm *MeshNormalMaterial) SetBlendSrc(v BlendingSrcFactor) {
	mnm.Set("blendSrc", v)
}
func (mnm *MeshNormalMaterial) BlendSrcAlpha() int {
	return mnm.Get("blendSrcAlpha").Int()
}
func (mnm *MeshNormalMaterial) SetBlendSrcAlpha(v int) {
	mnm.Set("blendSrcAlpha", v)
}
func (mnm *MeshNormalMaterial) Blending() *Blending {
	return &Blending{Value: mnm.Get("blending")}
}
func (mnm *MeshNormalMaterial) SetBlending(v *Blending) {
	mnm.Set("blending", v)
}
func (mnm *MeshNormalMaterial) ClipIntersection() bool {
	return mnm.Get("clipIntersection").Bool()
}
func (mnm *MeshNormalMaterial) SetClipIntersection(v bool) {
	mnm.Set("clipIntersection", v)
}
func (mnm *MeshNormalMaterial) ClipShadows() bool {
	return mnm.Get("clipShadows").Bool()
}
func (mnm *MeshNormalMaterial) SetClipShadows(v bool) {
	mnm.Set("clipShadows", v)
}
func (mnm *MeshNormalMaterial) ClippingPlanes() js.Value {
	return mnm.Get("clippingPlanes")
}
func (mnm *MeshNormalMaterial) SetClippingPlanes(v js.Value) {
	mnm.Set("clippingPlanes", v)
}
func (mnm *MeshNormalMaterial) ColorWrite() bool {
	return mnm.Get("colorWrite").Bool()
}
func (mnm *MeshNormalMaterial) SetColorWrite(v bool) {
	mnm.Set("colorWrite", v)
}
func (mnm *MeshNormalMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: mnm.Get("depthFunc")}
}
func (mnm *MeshNormalMaterial) SetDepthFunc(v *DepthModes) {
	mnm.Set("depthFunc", v)
}
func (mnm *MeshNormalMaterial) DepthTest() bool {
	return mnm.Get("depthTest").Bool()
}
func (mnm *MeshNormalMaterial) SetDepthTest(v bool) {
	mnm.Set("depthTest", v)
}
func (mnm *MeshNormalMaterial) DepthWrite() bool {
	return mnm.Get("depthWrite").Bool()
}
func (mnm *MeshNormalMaterial) SetDepthWrite(v bool) {
	mnm.Set("depthWrite", v)
}
func (mnm *MeshNormalMaterial) Dithering() bool {
	return mnm.Get("dithering").Bool()
}
func (mnm *MeshNormalMaterial) SetDithering(v bool) {
	mnm.Set("dithering", v)
}
func (mnm *MeshNormalMaterial) FlatShading() bool {
	return mnm.Get("flatShading").Bool()
}
func (mnm *MeshNormalMaterial) SetFlatShading(v bool) {
	mnm.Set("flatShading", v)
}
func (mnm *MeshNormalMaterial) Fog() bool {
	return mnm.Get("fog").Bool()
}
func (mnm *MeshNormalMaterial) SetFog(v bool) {
	mnm.Set("fog", v)
}
func (mnm *MeshNormalMaterial) Id() int {
	return mnm.Get("id").Int()
}
func (mnm *MeshNormalMaterial) SetId(v int) {
	mnm.Set("id", v)
}
func (mnm *MeshNormalMaterial) IsMaterial() bool {
	return mnm.Get("isMaterial").Bool()
}
func (mnm *MeshNormalMaterial) SetIsMaterial(v bool) {
	mnm.Set("isMaterial", v)
}
func (mnm *MeshNormalMaterial) Lights() bool {
	return mnm.Get("lights").Bool()
}
func (mnm *MeshNormalMaterial) SetLights(v bool) {
	mnm.Set("lights", v)
}
func (mnm *MeshNormalMaterial) MorphTargets() bool {
	return mnm.Get("morphTargets").Bool()
}
func (mnm *MeshNormalMaterial) SetMorphTargets(v bool) {
	mnm.Set("morphTargets", v)
}
func (mnm *MeshNormalMaterial) Name() string {
	return mnm.Get("name").String()
}
func (mnm *MeshNormalMaterial) SetName(v string) {
	mnm.Set("name", v)
}
func (mnm *MeshNormalMaterial) NeedsUpdate() bool {
	return mnm.Get("needsUpdate").Bool()
}
func (mnm *MeshNormalMaterial) SetNeedsUpdate(v bool) {
	mnm.Set("needsUpdate", v)
}
func (mnm *MeshNormalMaterial) Opacity() int {
	return mnm.Get("opacity").Int()
}
func (mnm *MeshNormalMaterial) SetOpacity(v int) {
	mnm.Set("opacity", v)
}
func (mnm *MeshNormalMaterial) Overdraw() int {
	return mnm.Get("overdraw").Int()
}
func (mnm *MeshNormalMaterial) SetOverdraw(v int) {
	mnm.Set("overdraw", v)
}
func (mnm *MeshNormalMaterial) PolygonOffset() bool {
	return mnm.Get("polygonOffset").Bool()
}
func (mnm *MeshNormalMaterial) SetPolygonOffset(v bool) {
	mnm.Set("polygonOffset", v)
}
func (mnm *MeshNormalMaterial) PolygonOffsetFactor() int {
	return mnm.Get("polygonOffsetFactor").Int()
}
func (mnm *MeshNormalMaterial) SetPolygonOffsetFactor(v int) {
	mnm.Set("polygonOffsetFactor", v)
}
func (mnm *MeshNormalMaterial) PolygonOffsetUnits() int {
	return mnm.Get("polygonOffsetUnits").Int()
}
func (mnm *MeshNormalMaterial) SetPolygonOffsetUnits(v int) {
	mnm.Set("polygonOffsetUnits", v)
}
func (mnm *MeshNormalMaterial) Precision() string {
	return mnm.Get("precision").String()
}
func (mnm *MeshNormalMaterial) SetPrecision(v string) {
	mnm.Set("precision", v)
}
func (mnm *MeshNormalMaterial) PremultipliedAlpha() bool {
	return mnm.Get("premultipliedAlpha").Bool()
}
func (mnm *MeshNormalMaterial) SetPremultipliedAlpha(v bool) {
	mnm.Set("premultipliedAlpha", v)
}
func (mnm *MeshNormalMaterial) Side() *Side {
	return &Side{Value: mnm.Get("side")}
}
func (mnm *MeshNormalMaterial) SetSide(v *Side) {
	mnm.Set("side", v)
}
func (mnm *MeshNormalMaterial) Transparent() bool {
	return mnm.Get("transparent").Bool()
}
func (mnm *MeshNormalMaterial) SetTransparent(v bool) {
	mnm.Set("transparent", v)
}
func (mnm *MeshNormalMaterial) Type() string {
	return mnm.Get("type").String()
}
func (mnm *MeshNormalMaterial) SetType(v string) {
	mnm.Set("type", v)
}
func (mnm *MeshNormalMaterial) UserData() js.Value {
	return mnm.Get("userData")
}
func (mnm *MeshNormalMaterial) SetUserData(v js.Value) {
	mnm.Set("userData", v)
}
func (mnm *MeshNormalMaterial) Uuid() string {
	return mnm.Get("uuid").String()
}
func (mnm *MeshNormalMaterial) SetUuid(v string) {
	mnm.Set("uuid", v)
}
func (mnm *MeshNormalMaterial) VertexColors() *Colors {
	return &Colors{Value: mnm.Get("vertexColors")}
}
func (mnm *MeshNormalMaterial) SetVertexColors(v *Colors) {
	mnm.Set("vertexColors", v)
}
func (mnm *MeshNormalMaterial) VertexTangents() bool {
	return mnm.Get("vertexTangents").Bool()
}
func (mnm *MeshNormalMaterial) SetVertexTangents(v bool) {
	mnm.Set("vertexTangents", v)
}
func (mnm *MeshNormalMaterial) Visible() bool {
	return mnm.Get("visible").Bool()
}
func (mnm *MeshNormalMaterial) SetVisible(v bool) {
	mnm.Set("visible", v)
}
func (mnm *MeshNormalMaterial) Wireframe() bool {
	return mnm.Get("wireframe").Bool()
}
func (mnm *MeshNormalMaterial) SetWireframe(v bool) {
	mnm.Set("wireframe", v)
}
func (mnm *MeshNormalMaterial) WireframeLinewidth() int {
	return mnm.Get("wireframeLinewidth").Int()
}
func (mnm *MeshNormalMaterial) SetWireframeLinewidth(v int) {
	mnm.Set("wireframeLinewidth", v)
}
func (mnm *MeshNormalMaterial) AddEventListener(typ string, listener js.Value) {
	mnm.Call("addEventListener", typ, listener)
}
func (mnm *MeshNormalMaterial) Clone() this {
	return this(mnm.Call("clone"))
}
func (mnm *MeshNormalMaterial) Copy(material *Material) this {
	return this(mnm.Call("copy", material))
}
func (mnm *MeshNormalMaterial) DispatchEvent(event js.Value) {
	mnm.Call("dispatchEvent", event)
}
func (mnm *MeshNormalMaterial) Dispose() {
	mnm.Call("dispose")
}
func (mnm *MeshNormalMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mnm.Call("hasEventListener", typ, listener).Bool()
}
func (mnm *MeshNormalMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	mnm.Call("onBeforeCompile", shader, renderer)
}
func (mnm *MeshNormalMaterial) RemoveEventListener(typ string, listener js.Value) {
	mnm.Call("removeEventListener", typ, listener)
}
func (mnm *MeshNormalMaterial) SetValues(parameters MeshNormalMaterialParameters) {
	mnm.Call("setValues", parameters)
}
func (mnm *MeshNormalMaterial) ToJSON(meta js.Value) js.Value {
	return mnm.Call("toJSON", meta)
}
func (mnm *MeshNormalMaterial) Update() {
	mnm.Call("update")
}

package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
)

type LineDashedMaterialParameters interface {
}
type LineDashedMaterial struct {
	js.Value
}

func NewLineDashedMaterial(parameters LineDashedMaterialParameters) *LineDashedMaterial {
	return &LineDashedMaterial{Value: get("LineDashedMaterial").New(parameters)}
}
func (ldm *LineDashedMaterial) AlphaTest() int {
	return ldm.Get("alphaTest").Int()
}
func (ldm *LineDashedMaterial) SetAlphaTest(v int) {
	ldm.Set("alphaTest", v)
}
func (ldm *LineDashedMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: ldm.Get("blendDst")}
}
func (ldm *LineDashedMaterial) SetBlendDst(v *BlendingDstFactor) {
	ldm.Set("blendDst", v)
}
func (ldm *LineDashedMaterial) BlendDstAlpha() int {
	return ldm.Get("blendDstAlpha").Int()
}
func (ldm *LineDashedMaterial) SetBlendDstAlpha(v int) {
	ldm.Set("blendDstAlpha", v)
}
func (ldm *LineDashedMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: ldm.Get("blendEquation")}
}
func (ldm *LineDashedMaterial) SetBlendEquation(v *BlendingEquation) {
	ldm.Set("blendEquation", v)
}
func (ldm *LineDashedMaterial) BlendEquationAlpha() int {
	return ldm.Get("blendEquationAlpha").Int()
}
func (ldm *LineDashedMaterial) SetBlendEquationAlpha(v int) {
	ldm.Set("blendEquationAlpha", v)
}
func (ldm *LineDashedMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(ldm.Get("blendSrc"))
}
func (ldm *LineDashedMaterial) SetBlendSrc(v BlendingSrcFactor) {
	ldm.Set("blendSrc", v)
}
func (ldm *LineDashedMaterial) BlendSrcAlpha() int {
	return ldm.Get("blendSrcAlpha").Int()
}
func (ldm *LineDashedMaterial) SetBlendSrcAlpha(v int) {
	ldm.Set("blendSrcAlpha", v)
}
func (ldm *LineDashedMaterial) Blending() *Blending {
	return &Blending{Value: ldm.Get("blending")}
}
func (ldm *LineDashedMaterial) SetBlending(v *Blending) {
	ldm.Set("blending", v)
}
func (ldm *LineDashedMaterial) ClipIntersection() bool {
	return ldm.Get("clipIntersection").Bool()
}
func (ldm *LineDashedMaterial) SetClipIntersection(v bool) {
	ldm.Set("clipIntersection", v)
}
func (ldm *LineDashedMaterial) ClipShadows() bool {
	return ldm.Get("clipShadows").Bool()
}
func (ldm *LineDashedMaterial) SetClipShadows(v bool) {
	ldm.Set("clipShadows", v)
}
func (ldm *LineDashedMaterial) ClippingPlanes() js.Value {
	return ldm.Get("clippingPlanes")
}
func (ldm *LineDashedMaterial) SetClippingPlanes(v js.Value) {
	ldm.Set("clippingPlanes", v)
}
func (ldm *LineDashedMaterial) Color() *math.Color {
	return &math.Color{Value: ldm.Get("color")}
}
func (ldm *LineDashedMaterial) SetColor(v *math.Color) {
	ldm.Set("color", v)
}
func (ldm *LineDashedMaterial) ColorWrite() bool {
	return ldm.Get("colorWrite").Bool()
}
func (ldm *LineDashedMaterial) SetColorWrite(v bool) {
	ldm.Set("colorWrite", v)
}
func (ldm *LineDashedMaterial) DashSize() int {
	return ldm.Get("dashSize").Int()
}
func (ldm *LineDashedMaterial) SetDashSize(v int) {
	ldm.Set("dashSize", v)
}
func (ldm *LineDashedMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: ldm.Get("depthFunc")}
}
func (ldm *LineDashedMaterial) SetDepthFunc(v *DepthModes) {
	ldm.Set("depthFunc", v)
}
func (ldm *LineDashedMaterial) DepthTest() bool {
	return ldm.Get("depthTest").Bool()
}
func (ldm *LineDashedMaterial) SetDepthTest(v bool) {
	ldm.Set("depthTest", v)
}
func (ldm *LineDashedMaterial) DepthWrite() bool {
	return ldm.Get("depthWrite").Bool()
}
func (ldm *LineDashedMaterial) SetDepthWrite(v bool) {
	ldm.Set("depthWrite", v)
}
func (ldm *LineDashedMaterial) Dithering() bool {
	return ldm.Get("dithering").Bool()
}
func (ldm *LineDashedMaterial) SetDithering(v bool) {
	ldm.Set("dithering", v)
}
func (ldm *LineDashedMaterial) FlatShading() bool {
	return ldm.Get("flatShading").Bool()
}
func (ldm *LineDashedMaterial) SetFlatShading(v bool) {
	ldm.Set("flatShading", v)
}
func (ldm *LineDashedMaterial) Fog() bool {
	return ldm.Get("fog").Bool()
}
func (ldm *LineDashedMaterial) SetFog(v bool) {
	ldm.Set("fog", v)
}
func (ldm *LineDashedMaterial) GapSize() int {
	return ldm.Get("gapSize").Int()
}
func (ldm *LineDashedMaterial) SetGapSize(v int) {
	ldm.Set("gapSize", v)
}
func (ldm *LineDashedMaterial) Id() int {
	return ldm.Get("id").Int()
}
func (ldm *LineDashedMaterial) SetId(v int) {
	ldm.Set("id", v)
}
func (ldm *LineDashedMaterial) IsLineDashedMaterial() bool {
	return ldm.Get("isLineDashedMaterial").Bool()
}
func (ldm *LineDashedMaterial) SetIsLineDashedMaterial(v bool) {
	ldm.Set("isLineDashedMaterial", v)
}
func (ldm *LineDashedMaterial) IsMaterial() bool {
	return ldm.Get("isMaterial").Bool()
}
func (ldm *LineDashedMaterial) SetIsMaterial(v bool) {
	ldm.Set("isMaterial", v)
}
func (ldm *LineDashedMaterial) Lights() bool {
	return ldm.Get("lights").Bool()
}
func (ldm *LineDashedMaterial) SetLights(v bool) {
	ldm.Set("lights", v)
}
func (ldm *LineDashedMaterial) Linecap() string {
	return ldm.Get("linecap").String()
}
func (ldm *LineDashedMaterial) SetLinecap(v string) {
	ldm.Set("linecap", v)
}
func (ldm *LineDashedMaterial) Linejoin() string {
	return ldm.Get("linejoin").String()
}
func (ldm *LineDashedMaterial) SetLinejoin(v string) {
	ldm.Set("linejoin", v)
}
func (ldm *LineDashedMaterial) Linewidth() int {
	return ldm.Get("linewidth").Int()
}
func (ldm *LineDashedMaterial) SetLinewidth(v int) {
	ldm.Set("linewidth", v)
}
func (ldm *LineDashedMaterial) Name() string {
	return ldm.Get("name").String()
}
func (ldm *LineDashedMaterial) SetName(v string) {
	ldm.Set("name", v)
}
func (ldm *LineDashedMaterial) NeedsUpdate() bool {
	return ldm.Get("needsUpdate").Bool()
}
func (ldm *LineDashedMaterial) SetNeedsUpdate(v bool) {
	ldm.Set("needsUpdate", v)
}
func (ldm *LineDashedMaterial) Opacity() int {
	return ldm.Get("opacity").Int()
}
func (ldm *LineDashedMaterial) SetOpacity(v int) {
	ldm.Set("opacity", v)
}
func (ldm *LineDashedMaterial) Overdraw() int {
	return ldm.Get("overdraw").Int()
}
func (ldm *LineDashedMaterial) SetOverdraw(v int) {
	ldm.Set("overdraw", v)
}
func (ldm *LineDashedMaterial) PolygonOffset() bool {
	return ldm.Get("polygonOffset").Bool()
}
func (ldm *LineDashedMaterial) SetPolygonOffset(v bool) {
	ldm.Set("polygonOffset", v)
}
func (ldm *LineDashedMaterial) PolygonOffsetFactor() int {
	return ldm.Get("polygonOffsetFactor").Int()
}
func (ldm *LineDashedMaterial) SetPolygonOffsetFactor(v int) {
	ldm.Set("polygonOffsetFactor", v)
}
func (ldm *LineDashedMaterial) PolygonOffsetUnits() int {
	return ldm.Get("polygonOffsetUnits").Int()
}
func (ldm *LineDashedMaterial) SetPolygonOffsetUnits(v int) {
	ldm.Set("polygonOffsetUnits", v)
}
func (ldm *LineDashedMaterial) Precision() string {
	return ldm.Get("precision").String()
}
func (ldm *LineDashedMaterial) SetPrecision(v string) {
	ldm.Set("precision", v)
}
func (ldm *LineDashedMaterial) PremultipliedAlpha() bool {
	return ldm.Get("premultipliedAlpha").Bool()
}
func (ldm *LineDashedMaterial) SetPremultipliedAlpha(v bool) {
	ldm.Set("premultipliedAlpha", v)
}
func (ldm *LineDashedMaterial) Scale() int {
	return ldm.Get("scale").Int()
}
func (ldm *LineDashedMaterial) SetScale(v int) {
	ldm.Set("scale", v)
}
func (ldm *LineDashedMaterial) Side() *Side {
	return &Side{Value: ldm.Get("side")}
}
func (ldm *LineDashedMaterial) SetSide(v *Side) {
	ldm.Set("side", v)
}
func (ldm *LineDashedMaterial) Transparent() bool {
	return ldm.Get("transparent").Bool()
}
func (ldm *LineDashedMaterial) SetTransparent(v bool) {
	ldm.Set("transparent", v)
}
func (ldm *LineDashedMaterial) Type() string {
	return ldm.Get("type").String()
}
func (ldm *LineDashedMaterial) SetType(v string) {
	ldm.Set("type", v)
}
func (ldm *LineDashedMaterial) UserData() js.Value {
	return ldm.Get("userData")
}
func (ldm *LineDashedMaterial) SetUserData(v js.Value) {
	ldm.Set("userData", v)
}
func (ldm *LineDashedMaterial) Uuid() string {
	return ldm.Get("uuid").String()
}
func (ldm *LineDashedMaterial) SetUuid(v string) {
	ldm.Set("uuid", v)
}
func (ldm *LineDashedMaterial) VertexColors() *Colors {
	return &Colors{Value: ldm.Get("vertexColors")}
}
func (ldm *LineDashedMaterial) SetVertexColors(v *Colors) {
	ldm.Set("vertexColors", v)
}
func (ldm *LineDashedMaterial) VertexTangents() bool {
	return ldm.Get("vertexTangents").Bool()
}
func (ldm *LineDashedMaterial) SetVertexTangents(v bool) {
	ldm.Set("vertexTangents", v)
}
func (ldm *LineDashedMaterial) Visible() bool {
	return ldm.Get("visible").Bool()
}
func (ldm *LineDashedMaterial) SetVisible(v bool) {
	ldm.Set("visible", v)
}
func (ldm *LineDashedMaterial) AddEventListener(typ string, listener js.Value) {
	ldm.Call("addEventListener", typ, listener)
}
func (ldm *LineDashedMaterial) Clone() this {
	return this(ldm.Call("clone"))
}
func (ldm *LineDashedMaterial) Copy(material *Material) this {
	return this(ldm.Call("copy", material))
}
func (ldm *LineDashedMaterial) DispatchEvent(event js.Value) {
	ldm.Call("dispatchEvent", event)
}
func (ldm *LineDashedMaterial) Dispose() {
	ldm.Call("dispose")
}
func (ldm *LineDashedMaterial) HasEventListener(typ string, listener js.Value) bool {
	return ldm.Call("hasEventListener", typ, listener).Bool()
}
func (ldm *LineDashedMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	ldm.Call("onBeforeCompile", shader, renderer)
}
func (ldm *LineDashedMaterial) RemoveEventListener(typ string, listener js.Value) {
	ldm.Call("removeEventListener", typ, listener)
}
func (ldm *LineDashedMaterial) SetValues(parameters LineDashedMaterialParameters) {
	ldm.Call("setValues", parameters)
}
func (ldm *LineDashedMaterial) ToJSON(meta js.Value) js.Value {
	return ldm.Call("toJSON", meta)
}
func (ldm *LineDashedMaterial) Update() {
	ldm.Call("update")
}

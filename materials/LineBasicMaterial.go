package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
)

type LineBasicMaterialParameters interface {
}
type LineBasicMaterial struct {
	js.Value
}

func NewLineBasicMaterial(parameters LineBasicMaterialParameters) *LineBasicMaterial {
	return &LineBasicMaterial{Value: get("LineBasicMaterial").New(parameters)}
}
func (lbm *LineBasicMaterial) AlphaTest() int {
	return lbm.Get("alphaTest").Int()
}
func (lbm *LineBasicMaterial) SetAlphaTest(v int) {
	lbm.Set("alphaTest", v)
}
func (lbm *LineBasicMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: lbm.Get("blendDst")}
}
func (lbm *LineBasicMaterial) SetBlendDst(v *BlendingDstFactor) {
	lbm.Set("blendDst", v)
}
func (lbm *LineBasicMaterial) BlendDstAlpha() int {
	return lbm.Get("blendDstAlpha").Int()
}
func (lbm *LineBasicMaterial) SetBlendDstAlpha(v int) {
	lbm.Set("blendDstAlpha", v)
}
func (lbm *LineBasicMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: lbm.Get("blendEquation")}
}
func (lbm *LineBasicMaterial) SetBlendEquation(v *BlendingEquation) {
	lbm.Set("blendEquation", v)
}
func (lbm *LineBasicMaterial) BlendEquationAlpha() int {
	return lbm.Get("blendEquationAlpha").Int()
}
func (lbm *LineBasicMaterial) SetBlendEquationAlpha(v int) {
	lbm.Set("blendEquationAlpha", v)
}
func (lbm *LineBasicMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(lbm.Get("blendSrc"))
}
func (lbm *LineBasicMaterial) SetBlendSrc(v BlendingSrcFactor) {
	lbm.Set("blendSrc", v)
}
func (lbm *LineBasicMaterial) BlendSrcAlpha() int {
	return lbm.Get("blendSrcAlpha").Int()
}
func (lbm *LineBasicMaterial) SetBlendSrcAlpha(v int) {
	lbm.Set("blendSrcAlpha", v)
}
func (lbm *LineBasicMaterial) Blending() *Blending {
	return &Blending{Value: lbm.Get("blending")}
}
func (lbm *LineBasicMaterial) SetBlending(v *Blending) {
	lbm.Set("blending", v)
}
func (lbm *LineBasicMaterial) ClipIntersection() bool {
	return lbm.Get("clipIntersection").Bool()
}
func (lbm *LineBasicMaterial) SetClipIntersection(v bool) {
	lbm.Set("clipIntersection", v)
}
func (lbm *LineBasicMaterial) ClipShadows() bool {
	return lbm.Get("clipShadows").Bool()
}
func (lbm *LineBasicMaterial) SetClipShadows(v bool) {
	lbm.Set("clipShadows", v)
}
func (lbm *LineBasicMaterial) ClippingPlanes() js.Value {
	return lbm.Get("clippingPlanes")
}
func (lbm *LineBasicMaterial) SetClippingPlanes(v js.Value) {
	lbm.Set("clippingPlanes", v)
}
func (lbm *LineBasicMaterial) Color() *math.Color {
	return &math.Color{Value: lbm.Get("color")}
}
func (lbm *LineBasicMaterial) SetColor(v *math.Color) {
	lbm.Set("color", v)
}
func (lbm *LineBasicMaterial) ColorWrite() bool {
	return lbm.Get("colorWrite").Bool()
}
func (lbm *LineBasicMaterial) SetColorWrite(v bool) {
	lbm.Set("colorWrite", v)
}
func (lbm *LineBasicMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: lbm.Get("depthFunc")}
}
func (lbm *LineBasicMaterial) SetDepthFunc(v *DepthModes) {
	lbm.Set("depthFunc", v)
}
func (lbm *LineBasicMaterial) DepthTest() bool {
	return lbm.Get("depthTest").Bool()
}
func (lbm *LineBasicMaterial) SetDepthTest(v bool) {
	lbm.Set("depthTest", v)
}
func (lbm *LineBasicMaterial) DepthWrite() bool {
	return lbm.Get("depthWrite").Bool()
}
func (lbm *LineBasicMaterial) SetDepthWrite(v bool) {
	lbm.Set("depthWrite", v)
}
func (lbm *LineBasicMaterial) Dithering() bool {
	return lbm.Get("dithering").Bool()
}
func (lbm *LineBasicMaterial) SetDithering(v bool) {
	lbm.Set("dithering", v)
}
func (lbm *LineBasicMaterial) FlatShading() bool {
	return lbm.Get("flatShading").Bool()
}
func (lbm *LineBasicMaterial) SetFlatShading(v bool) {
	lbm.Set("flatShading", v)
}
func (lbm *LineBasicMaterial) Fog() bool {
	return lbm.Get("fog").Bool()
}
func (lbm *LineBasicMaterial) SetFog(v bool) {
	lbm.Set("fog", v)
}
func (lbm *LineBasicMaterial) Id() int {
	return lbm.Get("id").Int()
}
func (lbm *LineBasicMaterial) SetId(v int) {
	lbm.Set("id", v)
}
func (lbm *LineBasicMaterial) IsMaterial() bool {
	return lbm.Get("isMaterial").Bool()
}
func (lbm *LineBasicMaterial) SetIsMaterial(v bool) {
	lbm.Set("isMaterial", v)
}
func (lbm *LineBasicMaterial) Lights() bool {
	return lbm.Get("lights").Bool()
}
func (lbm *LineBasicMaterial) SetLights(v bool) {
	lbm.Set("lights", v)
}
func (lbm *LineBasicMaterial) Linecap() string {
	return lbm.Get("linecap").String()
}
func (lbm *LineBasicMaterial) SetLinecap(v string) {
	lbm.Set("linecap", v)
}
func (lbm *LineBasicMaterial) Linejoin() string {
	return lbm.Get("linejoin").String()
}
func (lbm *LineBasicMaterial) SetLinejoin(v string) {
	lbm.Set("linejoin", v)
}
func (lbm *LineBasicMaterial) Linewidth() int {
	return lbm.Get("linewidth").Int()
}
func (lbm *LineBasicMaterial) SetLinewidth(v int) {
	lbm.Set("linewidth", v)
}
func (lbm *LineBasicMaterial) Name() string {
	return lbm.Get("name").String()
}
func (lbm *LineBasicMaterial) SetName(v string) {
	lbm.Set("name", v)
}
func (lbm *LineBasicMaterial) NeedsUpdate() bool {
	return lbm.Get("needsUpdate").Bool()
}
func (lbm *LineBasicMaterial) SetNeedsUpdate(v bool) {
	lbm.Set("needsUpdate", v)
}
func (lbm *LineBasicMaterial) Opacity() int {
	return lbm.Get("opacity").Int()
}
func (lbm *LineBasicMaterial) SetOpacity(v int) {
	lbm.Set("opacity", v)
}
func (lbm *LineBasicMaterial) Overdraw() int {
	return lbm.Get("overdraw").Int()
}
func (lbm *LineBasicMaterial) SetOverdraw(v int) {
	lbm.Set("overdraw", v)
}
func (lbm *LineBasicMaterial) PolygonOffset() bool {
	return lbm.Get("polygonOffset").Bool()
}
func (lbm *LineBasicMaterial) SetPolygonOffset(v bool) {
	lbm.Set("polygonOffset", v)
}
func (lbm *LineBasicMaterial) PolygonOffsetFactor() int {
	return lbm.Get("polygonOffsetFactor").Int()
}
func (lbm *LineBasicMaterial) SetPolygonOffsetFactor(v int) {
	lbm.Set("polygonOffsetFactor", v)
}
func (lbm *LineBasicMaterial) PolygonOffsetUnits() int {
	return lbm.Get("polygonOffsetUnits").Int()
}
func (lbm *LineBasicMaterial) SetPolygonOffsetUnits(v int) {
	lbm.Set("polygonOffsetUnits", v)
}
func (lbm *LineBasicMaterial) Precision() string {
	return lbm.Get("precision").String()
}
func (lbm *LineBasicMaterial) SetPrecision(v string) {
	lbm.Set("precision", v)
}
func (lbm *LineBasicMaterial) PremultipliedAlpha() bool {
	return lbm.Get("premultipliedAlpha").Bool()
}
func (lbm *LineBasicMaterial) SetPremultipliedAlpha(v bool) {
	lbm.Set("premultipliedAlpha", v)
}
func (lbm *LineBasicMaterial) Side() *Side {
	return &Side{Value: lbm.Get("side")}
}
func (lbm *LineBasicMaterial) SetSide(v *Side) {
	lbm.Set("side", v)
}
func (lbm *LineBasicMaterial) Transparent() bool {
	return lbm.Get("transparent").Bool()
}
func (lbm *LineBasicMaterial) SetTransparent(v bool) {
	lbm.Set("transparent", v)
}
func (lbm *LineBasicMaterial) Type() string {
	return lbm.Get("type").String()
}
func (lbm *LineBasicMaterial) SetType(v string) {
	lbm.Set("type", v)
}
func (lbm *LineBasicMaterial) UserData() js.Value {
	return lbm.Get("userData")
}
func (lbm *LineBasicMaterial) SetUserData(v js.Value) {
	lbm.Set("userData", v)
}
func (lbm *LineBasicMaterial) Uuid() string {
	return lbm.Get("uuid").String()
}
func (lbm *LineBasicMaterial) SetUuid(v string) {
	lbm.Set("uuid", v)
}
func (lbm *LineBasicMaterial) VertexColors() *Colors {
	return &Colors{Value: lbm.Get("vertexColors")}
}
func (lbm *LineBasicMaterial) SetVertexColors(v *Colors) {
	lbm.Set("vertexColors", v)
}
func (lbm *LineBasicMaterial) VertexTangents() bool {
	return lbm.Get("vertexTangents").Bool()
}
func (lbm *LineBasicMaterial) SetVertexTangents(v bool) {
	lbm.Set("vertexTangents", v)
}
func (lbm *LineBasicMaterial) Visible() bool {
	return lbm.Get("visible").Bool()
}
func (lbm *LineBasicMaterial) SetVisible(v bool) {
	lbm.Set("visible", v)
}
func (lbm *LineBasicMaterial) AddEventListener(typ string, listener js.Value) {
	lbm.Call("addEventListener", typ, listener)
}
func (lbm *LineBasicMaterial) Clone() this {
	return this(lbm.Call("clone"))
}
func (lbm *LineBasicMaterial) Copy(material *Material) this {
	return this(lbm.Call("copy", material))
}
func (lbm *LineBasicMaterial) DispatchEvent(event js.Value) {
	lbm.Call("dispatchEvent", event)
}
func (lbm *LineBasicMaterial) Dispose() {
	lbm.Call("dispose")
}
func (lbm *LineBasicMaterial) HasEventListener(typ string, listener js.Value) bool {
	return lbm.Call("hasEventListener", typ, listener).Bool()
}
func (lbm *LineBasicMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	lbm.Call("onBeforeCompile", shader, renderer)
}
func (lbm *LineBasicMaterial) RemoveEventListener(typ string, listener js.Value) {
	lbm.Call("removeEventListener", typ, listener)
}
func (lbm *LineBasicMaterial) SetValues(parameters LineBasicMaterialParameters) {
	lbm.Call("setValues", parameters)
}
func (lbm *LineBasicMaterial) ToJSON(meta js.Value) js.Value {
	return lbm.Call("toJSON", meta)
}
func (lbm *LineBasicMaterial) Update() {
	lbm.Call("update")
}

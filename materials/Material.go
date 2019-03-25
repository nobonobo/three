package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
)

func MaterialIdCount() int {
	return get("MaterialIdCount").Int()
}
func SetMaterialIdCount(v int) {
	set("MaterialIdCount", v)
}

type MaterialParameters interface {
}
type Material struct {
	js.Value
}

func NewMaterial() *Material {
	return &Material{Value: get("Material").New()}
}
func (m *Material) AlphaTest() int {
	return m.Get("alphaTest").Int()
}
func (m *Material) SetAlphaTest(v int) {
	m.Set("alphaTest", v)
}
func (m *Material) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: m.Get("blendDst")}
}
func (m *Material) SetBlendDst(v *BlendingDstFactor) {
	m.Set("blendDst", v)
}
func (m *Material) BlendDstAlpha() int {
	return m.Get("blendDstAlpha").Int()
}
func (m *Material) SetBlendDstAlpha(v int) {
	m.Set("blendDstAlpha", v)
}
func (m *Material) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: m.Get("blendEquation")}
}
func (m *Material) SetBlendEquation(v *BlendingEquation) {
	m.Set("blendEquation", v)
}
func (m *Material) BlendEquationAlpha() int {
	return m.Get("blendEquationAlpha").Int()
}
func (m *Material) SetBlendEquationAlpha(v int) {
	m.Set("blendEquationAlpha", v)
}
func (m *Material) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(m.Get("blendSrc"))
}
func (m *Material) SetBlendSrc(v BlendingSrcFactor) {
	m.Set("blendSrc", v)
}
func (m *Material) BlendSrcAlpha() int {
	return m.Get("blendSrcAlpha").Int()
}
func (m *Material) SetBlendSrcAlpha(v int) {
	m.Set("blendSrcAlpha", v)
}
func (m *Material) Blending() *Blending {
	return &Blending{Value: m.Get("blending")}
}
func (m *Material) SetBlending(v *Blending) {
	m.Set("blending", v)
}
func (m *Material) ClipIntersection() bool {
	return m.Get("clipIntersection").Bool()
}
func (m *Material) SetClipIntersection(v bool) {
	m.Set("clipIntersection", v)
}
func (m *Material) ClipShadows() bool {
	return m.Get("clipShadows").Bool()
}
func (m *Material) SetClipShadows(v bool) {
	m.Set("clipShadows", v)
}
func (m *Material) ClippingPlanes() js.Value {
	return m.Get("clippingPlanes")
}
func (m *Material) SetClippingPlanes(v js.Value) {
	m.Set("clippingPlanes", v)
}
func (m *Material) ColorWrite() bool {
	return m.Get("colorWrite").Bool()
}
func (m *Material) SetColorWrite(v bool) {
	m.Set("colorWrite", v)
}
func (m *Material) DepthFunc() *DepthModes {
	return &DepthModes{Value: m.Get("depthFunc")}
}
func (m *Material) SetDepthFunc(v *DepthModes) {
	m.Set("depthFunc", v)
}
func (m *Material) DepthTest() bool {
	return m.Get("depthTest").Bool()
}
func (m *Material) SetDepthTest(v bool) {
	m.Set("depthTest", v)
}
func (m *Material) DepthWrite() bool {
	return m.Get("depthWrite").Bool()
}
func (m *Material) SetDepthWrite(v bool) {
	m.Set("depthWrite", v)
}
func (m *Material) Dithering() bool {
	return m.Get("dithering").Bool()
}
func (m *Material) SetDithering(v bool) {
	m.Set("dithering", v)
}
func (m *Material) FlatShading() bool {
	return m.Get("flatShading").Bool()
}
func (m *Material) SetFlatShading(v bool) {
	m.Set("flatShading", v)
}
func (m *Material) Fog() bool {
	return m.Get("fog").Bool()
}
func (m *Material) SetFog(v bool) {
	m.Set("fog", v)
}
func (m *Material) Id() int {
	return m.Get("id").Int()
}
func (m *Material) SetId(v int) {
	m.Set("id", v)
}
func (m *Material) IsMaterial() bool {
	return m.Get("isMaterial").Bool()
}
func (m *Material) SetIsMaterial(v bool) {
	m.Set("isMaterial", v)
}
func (m *Material) Lights() bool {
	return m.Get("lights").Bool()
}
func (m *Material) SetLights(v bool) {
	m.Set("lights", v)
}
func (m *Material) Name() string {
	return m.Get("name").String()
}
func (m *Material) SetName(v string) {
	m.Set("name", v)
}
func (m *Material) NeedsUpdate() bool {
	return m.Get("needsUpdate").Bool()
}
func (m *Material) SetNeedsUpdate(v bool) {
	m.Set("needsUpdate", v)
}
func (m *Material) Opacity() int {
	return m.Get("opacity").Int()
}
func (m *Material) SetOpacity(v int) {
	m.Set("opacity", v)
}
func (m *Material) Overdraw() int {
	return m.Get("overdraw").Int()
}
func (m *Material) SetOverdraw(v int) {
	m.Set("overdraw", v)
}
func (m *Material) PolygonOffset() bool {
	return m.Get("polygonOffset").Bool()
}
func (m *Material) SetPolygonOffset(v bool) {
	m.Set("polygonOffset", v)
}
func (m *Material) PolygonOffsetFactor() int {
	return m.Get("polygonOffsetFactor").Int()
}
func (m *Material) SetPolygonOffsetFactor(v int) {
	m.Set("polygonOffsetFactor", v)
}
func (m *Material) PolygonOffsetUnits() int {
	return m.Get("polygonOffsetUnits").Int()
}
func (m *Material) SetPolygonOffsetUnits(v int) {
	m.Set("polygonOffsetUnits", v)
}
func (m *Material) Precision() string {
	return m.Get("precision").String()
}
func (m *Material) SetPrecision(v string) {
	m.Set("precision", v)
}
func (m *Material) PremultipliedAlpha() bool {
	return m.Get("premultipliedAlpha").Bool()
}
func (m *Material) SetPremultipliedAlpha(v bool) {
	m.Set("premultipliedAlpha", v)
}
func (m *Material) Side() *Side {
	return &Side{Value: m.Get("side")}
}
func (m *Material) SetSide(v *Side) {
	m.Set("side", v)
}
func (m *Material) Transparent() bool {
	return m.Get("transparent").Bool()
}
func (m *Material) SetTransparent(v bool) {
	m.Set("transparent", v)
}
func (m *Material) Type() string {
	return m.Get("type").String()
}
func (m *Material) SetType(v string) {
	m.Set("type", v)
}
func (m *Material) UserData() js.Value {
	return m.Get("userData")
}
func (m *Material) SetUserData(v js.Value) {
	m.Set("userData", v)
}
func (m *Material) Uuid() string {
	return m.Get("uuid").String()
}
func (m *Material) SetUuid(v string) {
	m.Set("uuid", v)
}
func (m *Material) VertexColors() *Colors {
	return &Colors{Value: m.Get("vertexColors")}
}
func (m *Material) SetVertexColors(v *Colors) {
	m.Set("vertexColors", v)
}
func (m *Material) VertexTangents() bool {
	return m.Get("vertexTangents").Bool()
}
func (m *Material) SetVertexTangents(v bool) {
	m.Set("vertexTangents", v)
}
func (m *Material) Visible() bool {
	return m.Get("visible").Bool()
}
func (m *Material) SetVisible(v bool) {
	m.Set("visible", v)
}
func (m *Material) AddEventListener(typ string, listener js.Value) {
	m.Call("addEventListener", typ, listener)
}
func (m *Material) Clone() this {
	return this(m.Call("clone"))
}
func (m *Material) Copy(material *Material) this {
	return this(m.Call("copy", material))
}
func (m *Material) DispatchEvent(event js.Value) {
	m.Call("dispatchEvent", event)
}
func (m *Material) Dispose() {
	m.Call("dispose")
}
func (m *Material) HasEventListener(typ string, listener js.Value) bool {
	return m.Call("hasEventListener", typ, listener).Bool()
}
func (m *Material) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	m.Call("onBeforeCompile", shader, renderer)
}
func (m *Material) RemoveEventListener(typ string, listener js.Value) {
	m.Call("removeEventListener", typ, listener)
}
func (m *Material) SetValues(values MaterialParameters) {
	m.Call("setValues", values)
}
func (m *Material) ToJSON(meta js.Value) js.Value {
	return m.Call("toJSON", meta)
}
func (m *Material) Update() {
	m.Call("update")
}

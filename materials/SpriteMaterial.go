package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
	"github.com/nobonobo/three/textures"
)

type SpriteMaterialParameters interface {
}
type SpriteMaterial struct {
	js.Value
}

func NewSpriteMaterial(parameters SpriteMaterialParameters) *SpriteMaterial {
	return &SpriteMaterial{Value: get("SpriteMaterial").New(parameters)}
}
func (sm *SpriteMaterial) AlphaTest() int {
	return sm.Get("alphaTest").Int()
}
func (sm *SpriteMaterial) SetAlphaTest(v int) {
	sm.Set("alphaTest", v)
}
func (sm *SpriteMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: sm.Get("blendDst")}
}
func (sm *SpriteMaterial) SetBlendDst(v *BlendingDstFactor) {
	sm.Set("blendDst", v)
}
func (sm *SpriteMaterial) BlendDstAlpha() int {
	return sm.Get("blendDstAlpha").Int()
}
func (sm *SpriteMaterial) SetBlendDstAlpha(v int) {
	sm.Set("blendDstAlpha", v)
}
func (sm *SpriteMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: sm.Get("blendEquation")}
}
func (sm *SpriteMaterial) SetBlendEquation(v *BlendingEquation) {
	sm.Set("blendEquation", v)
}
func (sm *SpriteMaterial) BlendEquationAlpha() int {
	return sm.Get("blendEquationAlpha").Int()
}
func (sm *SpriteMaterial) SetBlendEquationAlpha(v int) {
	sm.Set("blendEquationAlpha", v)
}
func (sm *SpriteMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(sm.Get("blendSrc"))
}
func (sm *SpriteMaterial) SetBlendSrc(v BlendingSrcFactor) {
	sm.Set("blendSrc", v)
}
func (sm *SpriteMaterial) BlendSrcAlpha() int {
	return sm.Get("blendSrcAlpha").Int()
}
func (sm *SpriteMaterial) SetBlendSrcAlpha(v int) {
	sm.Set("blendSrcAlpha", v)
}
func (sm *SpriteMaterial) Blending() *Blending {
	return &Blending{Value: sm.Get("blending")}
}
func (sm *SpriteMaterial) SetBlending(v *Blending) {
	sm.Set("blending", v)
}
func (sm *SpriteMaterial) ClipIntersection() bool {
	return sm.Get("clipIntersection").Bool()
}
func (sm *SpriteMaterial) SetClipIntersection(v bool) {
	sm.Set("clipIntersection", v)
}
func (sm *SpriteMaterial) ClipShadows() bool {
	return sm.Get("clipShadows").Bool()
}
func (sm *SpriteMaterial) SetClipShadows(v bool) {
	sm.Set("clipShadows", v)
}
func (sm *SpriteMaterial) ClippingPlanes() js.Value {
	return sm.Get("clippingPlanes")
}
func (sm *SpriteMaterial) SetClippingPlanes(v js.Value) {
	sm.Set("clippingPlanes", v)
}
func (sm *SpriteMaterial) Color() *math.Color {
	return &math.Color{Value: sm.Get("color")}
}
func (sm *SpriteMaterial) SetColor(v *math.Color) {
	sm.Set("color", v)
}
func (sm *SpriteMaterial) ColorWrite() bool {
	return sm.Get("colorWrite").Bool()
}
func (sm *SpriteMaterial) SetColorWrite(v bool) {
	sm.Set("colorWrite", v)
}
func (sm *SpriteMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: sm.Get("depthFunc")}
}
func (sm *SpriteMaterial) SetDepthFunc(v *DepthModes) {
	sm.Set("depthFunc", v)
}
func (sm *SpriteMaterial) DepthTest() bool {
	return sm.Get("depthTest").Bool()
}
func (sm *SpriteMaterial) SetDepthTest(v bool) {
	sm.Set("depthTest", v)
}
func (sm *SpriteMaterial) DepthWrite() bool {
	return sm.Get("depthWrite").Bool()
}
func (sm *SpriteMaterial) SetDepthWrite(v bool) {
	sm.Set("depthWrite", v)
}
func (sm *SpriteMaterial) Dithering() bool {
	return sm.Get("dithering").Bool()
}
func (sm *SpriteMaterial) SetDithering(v bool) {
	sm.Set("dithering", v)
}
func (sm *SpriteMaterial) FlatShading() bool {
	return sm.Get("flatShading").Bool()
}
func (sm *SpriteMaterial) SetFlatShading(v bool) {
	sm.Set("flatShading", v)
}
func (sm *SpriteMaterial) Fog() bool {
	return sm.Get("fog").Bool()
}
func (sm *SpriteMaterial) SetFog(v bool) {
	sm.Set("fog", v)
}
func (sm *SpriteMaterial) Id() int {
	return sm.Get("id").Int()
}
func (sm *SpriteMaterial) SetId(v int) {
	sm.Set("id", v)
}
func (sm *SpriteMaterial) IsMaterial() bool {
	return sm.Get("isMaterial").Bool()
}
func (sm *SpriteMaterial) SetIsMaterial(v bool) {
	sm.Set("isMaterial", v)
}
func (sm *SpriteMaterial) IsSpriteMaterial() true {
	return true(sm.Get("isSpriteMaterial"))
}
func (sm *SpriteMaterial) SetIsSpriteMaterial(v true) {
	sm.Set("isSpriteMaterial", v)
}
func (sm *SpriteMaterial) Lights() bool {
	return sm.Get("lights").Bool()
}
func (sm *SpriteMaterial) SetLights(v bool) {
	sm.Set("lights", v)
}
func (sm *SpriteMaterial) Map() textures.Texture {
	return textures.Texture(sm.Get("map"))
}
func (sm *SpriteMaterial) SetMap(v textures.Texture) {
	sm.Set("map", v)
}
func (sm *SpriteMaterial) Name() string {
	return sm.Get("name").String()
}
func (sm *SpriteMaterial) SetName(v string) {
	sm.Set("name", v)
}
func (sm *SpriteMaterial) NeedsUpdate() bool {
	return sm.Get("needsUpdate").Bool()
}
func (sm *SpriteMaterial) SetNeedsUpdate(v bool) {
	sm.Set("needsUpdate", v)
}
func (sm *SpriteMaterial) Opacity() int {
	return sm.Get("opacity").Int()
}
func (sm *SpriteMaterial) SetOpacity(v int) {
	sm.Set("opacity", v)
}
func (sm *SpriteMaterial) Overdraw() int {
	return sm.Get("overdraw").Int()
}
func (sm *SpriteMaterial) SetOverdraw(v int) {
	sm.Set("overdraw", v)
}
func (sm *SpriteMaterial) PolygonOffset() bool {
	return sm.Get("polygonOffset").Bool()
}
func (sm *SpriteMaterial) SetPolygonOffset(v bool) {
	sm.Set("polygonOffset", v)
}
func (sm *SpriteMaterial) PolygonOffsetFactor() int {
	return sm.Get("polygonOffsetFactor").Int()
}
func (sm *SpriteMaterial) SetPolygonOffsetFactor(v int) {
	sm.Set("polygonOffsetFactor", v)
}
func (sm *SpriteMaterial) PolygonOffsetUnits() int {
	return sm.Get("polygonOffsetUnits").Int()
}
func (sm *SpriteMaterial) SetPolygonOffsetUnits(v int) {
	sm.Set("polygonOffsetUnits", v)
}
func (sm *SpriteMaterial) Precision() string {
	return sm.Get("precision").String()
}
func (sm *SpriteMaterial) SetPrecision(v string) {
	sm.Set("precision", v)
}
func (sm *SpriteMaterial) PremultipliedAlpha() bool {
	return sm.Get("premultipliedAlpha").Bool()
}
func (sm *SpriteMaterial) SetPremultipliedAlpha(v bool) {
	sm.Set("premultipliedAlpha", v)
}
func (sm *SpriteMaterial) Rotation() int {
	return sm.Get("rotation").Int()
}
func (sm *SpriteMaterial) SetRotation(v int) {
	sm.Set("rotation", v)
}
func (sm *SpriteMaterial) Side() *Side {
	return &Side{Value: sm.Get("side")}
}
func (sm *SpriteMaterial) SetSide(v *Side) {
	sm.Set("side", v)
}
func (sm *SpriteMaterial) Transparent() bool {
	return sm.Get("transparent").Bool()
}
func (sm *SpriteMaterial) SetTransparent(v bool) {
	sm.Set("transparent", v)
}
func (sm *SpriteMaterial) Type() string {
	return sm.Get("type").String()
}
func (sm *SpriteMaterial) SetType(v string) {
	sm.Set("type", v)
}
func (sm *SpriteMaterial) UserData() js.Value {
	return sm.Get("userData")
}
func (sm *SpriteMaterial) SetUserData(v js.Value) {
	sm.Set("userData", v)
}
func (sm *SpriteMaterial) Uuid() string {
	return sm.Get("uuid").String()
}
func (sm *SpriteMaterial) SetUuid(v string) {
	sm.Set("uuid", v)
}
func (sm *SpriteMaterial) VertexColors() *Colors {
	return &Colors{Value: sm.Get("vertexColors")}
}
func (sm *SpriteMaterial) SetVertexColors(v *Colors) {
	sm.Set("vertexColors", v)
}
func (sm *SpriteMaterial) VertexTangents() bool {
	return sm.Get("vertexTangents").Bool()
}
func (sm *SpriteMaterial) SetVertexTangents(v bool) {
	sm.Set("vertexTangents", v)
}
func (sm *SpriteMaterial) Visible() bool {
	return sm.Get("visible").Bool()
}
func (sm *SpriteMaterial) SetVisible(v bool) {
	sm.Set("visible", v)
}
func (sm *SpriteMaterial) AddEventListener(typ string, listener js.Value) {
	sm.Call("addEventListener", typ, listener)
}
func (sm *SpriteMaterial) Clone() this {
	return this(sm.Call("clone"))
}
func (sm *SpriteMaterial) Copy(source *SpriteMaterial) this {
	return this(sm.Call("copy", source))
}
func (sm *SpriteMaterial) DispatchEvent(event js.Value) {
	sm.Call("dispatchEvent", event)
}
func (sm *SpriteMaterial) Dispose() {
	sm.Call("dispose")
}
func (sm *SpriteMaterial) HasEventListener(typ string, listener js.Value) bool {
	return sm.Call("hasEventListener", typ, listener).Bool()
}
func (sm *SpriteMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	sm.Call("onBeforeCompile", shader, renderer)
}
func (sm *SpriteMaterial) RemoveEventListener(typ string, listener js.Value) {
	sm.Call("removeEventListener", typ, listener)
}
func (sm *SpriteMaterial) SetValues(parameters SpriteMaterialParameters) {
	sm.Call("setValues", parameters)
}
func (sm *SpriteMaterial) ToJSON(meta js.Value) js.Value {
	return sm.Call("toJSON", meta)
}
func (sm *SpriteMaterial) Update() {
	sm.Call("update")
}

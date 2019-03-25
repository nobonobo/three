package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
)

type ShaderMaterialParameters interface {
}
type ShaderMaterial struct {
	js.Value
}

func NewShaderMaterial(parameters ShaderMaterialParameters) *ShaderMaterial {
	return &ShaderMaterial{Value: get("ShaderMaterial").New(parameters)}
}
func (sm *ShaderMaterial) AlphaTest() int {
	return sm.Get("alphaTest").Int()
}
func (sm *ShaderMaterial) SetAlphaTest(v int) {
	sm.Set("alphaTest", v)
}
func (sm *ShaderMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: sm.Get("blendDst")}
}
func (sm *ShaderMaterial) SetBlendDst(v *BlendingDstFactor) {
	sm.Set("blendDst", v)
}
func (sm *ShaderMaterial) BlendDstAlpha() int {
	return sm.Get("blendDstAlpha").Int()
}
func (sm *ShaderMaterial) SetBlendDstAlpha(v int) {
	sm.Set("blendDstAlpha", v)
}
func (sm *ShaderMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: sm.Get("blendEquation")}
}
func (sm *ShaderMaterial) SetBlendEquation(v *BlendingEquation) {
	sm.Set("blendEquation", v)
}
func (sm *ShaderMaterial) BlendEquationAlpha() int {
	return sm.Get("blendEquationAlpha").Int()
}
func (sm *ShaderMaterial) SetBlendEquationAlpha(v int) {
	sm.Set("blendEquationAlpha", v)
}
func (sm *ShaderMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(sm.Get("blendSrc"))
}
func (sm *ShaderMaterial) SetBlendSrc(v BlendingSrcFactor) {
	sm.Set("blendSrc", v)
}
func (sm *ShaderMaterial) BlendSrcAlpha() int {
	return sm.Get("blendSrcAlpha").Int()
}
func (sm *ShaderMaterial) SetBlendSrcAlpha(v int) {
	sm.Set("blendSrcAlpha", v)
}
func (sm *ShaderMaterial) Blending() *Blending {
	return &Blending{Value: sm.Get("blending")}
}
func (sm *ShaderMaterial) SetBlending(v *Blending) {
	sm.Set("blending", v)
}
func (sm *ShaderMaterial) ClipIntersection() bool {
	return sm.Get("clipIntersection").Bool()
}
func (sm *ShaderMaterial) SetClipIntersection(v bool) {
	sm.Set("clipIntersection", v)
}
func (sm *ShaderMaterial) ClipShadows() bool {
	return sm.Get("clipShadows").Bool()
}
func (sm *ShaderMaterial) SetClipShadows(v bool) {
	sm.Set("clipShadows", v)
}
func (sm *ShaderMaterial) Clipping() bool {
	return sm.Get("clipping").Bool()
}
func (sm *ShaderMaterial) SetClipping(v bool) {
	sm.Set("clipping", v)
}
func (sm *ShaderMaterial) ClippingPlanes() js.Value {
	return sm.Get("clippingPlanes")
}
func (sm *ShaderMaterial) SetClippingPlanes(v js.Value) {
	sm.Set("clippingPlanes", v)
}
func (sm *ShaderMaterial) ColorWrite() bool {
	return sm.Get("colorWrite").Bool()
}
func (sm *ShaderMaterial) SetColorWrite(v bool) {
	sm.Set("colorWrite", v)
}
func (sm *ShaderMaterial) DefaultAttributeValues() js.Value {
	return sm.Get("defaultAttributeValues")
}
func (sm *ShaderMaterial) SetDefaultAttributeValues(v js.Value) {
	sm.Set("defaultAttributeValues", v)
}
func (sm *ShaderMaterial) Defines() js.Value {
	return sm.Get("defines")
}
func (sm *ShaderMaterial) SetDefines(v js.Value) {
	sm.Set("defines", v)
}
func (sm *ShaderMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: sm.Get("depthFunc")}
}
func (sm *ShaderMaterial) SetDepthFunc(v *DepthModes) {
	sm.Set("depthFunc", v)
}
func (sm *ShaderMaterial) DepthTest() bool {
	return sm.Get("depthTest").Bool()
}
func (sm *ShaderMaterial) SetDepthTest(v bool) {
	sm.Set("depthTest", v)
}
func (sm *ShaderMaterial) DepthWrite() bool {
	return sm.Get("depthWrite").Bool()
}
func (sm *ShaderMaterial) SetDepthWrite(v bool) {
	sm.Set("depthWrite", v)
}
func (sm *ShaderMaterial) Derivatives() js.Value {
	return sm.Get("derivatives")
}
func (sm *ShaderMaterial) SetDerivatives(v js.Value) {
	sm.Set("derivatives", v)
}
func (sm *ShaderMaterial) Dithering() bool {
	return sm.Get("dithering").Bool()
}
func (sm *ShaderMaterial) SetDithering(v bool) {
	sm.Set("dithering", v)
}
func (sm *ShaderMaterial) Extensions() js.Value {
	return sm.Get("extensions")
}
func (sm *ShaderMaterial) SetExtensions(v js.Value) {
	sm.Set("extensions", v)
}
func (sm *ShaderMaterial) FlatShading() bool {
	return sm.Get("flatShading").Bool()
}
func (sm *ShaderMaterial) SetFlatShading(v bool) {
	sm.Set("flatShading", v)
}
func (sm *ShaderMaterial) Fog() bool {
	return sm.Get("fog").Bool()
}
func (sm *ShaderMaterial) SetFog(v bool) {
	sm.Set("fog", v)
}
func (sm *ShaderMaterial) FragmentShader() string {
	return sm.Get("fragmentShader").String()
}
func (sm *ShaderMaterial) SetFragmentShader(v string) {
	sm.Set("fragmentShader", v)
}
func (sm *ShaderMaterial) Id() int {
	return sm.Get("id").Int()
}
func (sm *ShaderMaterial) SetId(v int) {
	sm.Set("id", v)
}
func (sm *ShaderMaterial) Index0AttributeName() string {
	return sm.Get("index0AttributeName").String()
}
func (sm *ShaderMaterial) SetIndex0AttributeName(v string) {
	sm.Set("index0AttributeName", v)
}
func (sm *ShaderMaterial) IsMaterial() bool {
	return sm.Get("isMaterial").Bool()
}
func (sm *ShaderMaterial) SetIsMaterial(v bool) {
	sm.Set("isMaterial", v)
}
func (sm *ShaderMaterial) Lights() bool {
	return sm.Get("lights").Bool()
}
func (sm *ShaderMaterial) SetLights(v bool) {
	sm.Set("lights", v)
}
func (sm *ShaderMaterial) Linewidth() int {
	return sm.Get("linewidth").Int()
}
func (sm *ShaderMaterial) SetLinewidth(v int) {
	sm.Set("linewidth", v)
}
func (sm *ShaderMaterial) MorphNormals() bool {
	return sm.Get("morphNormals").Bool()
}
func (sm *ShaderMaterial) SetMorphNormals(v bool) {
	sm.Set("morphNormals", v)
}
func (sm *ShaderMaterial) MorphTargets() bool {
	return sm.Get("morphTargets").Bool()
}
func (sm *ShaderMaterial) SetMorphTargets(v bool) {
	sm.Set("morphTargets", v)
}
func (sm *ShaderMaterial) Name() string {
	return sm.Get("name").String()
}
func (sm *ShaderMaterial) SetName(v string) {
	sm.Set("name", v)
}
func (sm *ShaderMaterial) NeedsUpdate() bool {
	return sm.Get("needsUpdate").Bool()
}
func (sm *ShaderMaterial) SetNeedsUpdate(v bool) {
	sm.Set("needsUpdate", v)
}
func (sm *ShaderMaterial) Opacity() int {
	return sm.Get("opacity").Int()
}
func (sm *ShaderMaterial) SetOpacity(v int) {
	sm.Set("opacity", v)
}
func (sm *ShaderMaterial) Overdraw() int {
	return sm.Get("overdraw").Int()
}
func (sm *ShaderMaterial) SetOverdraw(v int) {
	sm.Set("overdraw", v)
}
func (sm *ShaderMaterial) PolygonOffset() bool {
	return sm.Get("polygonOffset").Bool()
}
func (sm *ShaderMaterial) SetPolygonOffset(v bool) {
	sm.Set("polygonOffset", v)
}
func (sm *ShaderMaterial) PolygonOffsetFactor() int {
	return sm.Get("polygonOffsetFactor").Int()
}
func (sm *ShaderMaterial) SetPolygonOffsetFactor(v int) {
	sm.Set("polygonOffsetFactor", v)
}
func (sm *ShaderMaterial) PolygonOffsetUnits() int {
	return sm.Get("polygonOffsetUnits").Int()
}
func (sm *ShaderMaterial) SetPolygonOffsetUnits(v int) {
	sm.Set("polygonOffsetUnits", v)
}
func (sm *ShaderMaterial) Precision() string {
	return sm.Get("precision").String()
}
func (sm *ShaderMaterial) SetPrecision(v string) {
	sm.Set("precision", v)
}
func (sm *ShaderMaterial) PremultipliedAlpha() bool {
	return sm.Get("premultipliedAlpha").Bool()
}
func (sm *ShaderMaterial) SetPremultipliedAlpha(v bool) {
	sm.Set("premultipliedAlpha", v)
}
func (sm *ShaderMaterial) Side() *Side {
	return &Side{Value: sm.Get("side")}
}
func (sm *ShaderMaterial) SetSide(v *Side) {
	sm.Set("side", v)
}
func (sm *ShaderMaterial) Skinning() bool {
	return sm.Get("skinning").Bool()
}
func (sm *ShaderMaterial) SetSkinning(v bool) {
	sm.Set("skinning", v)
}
func (sm *ShaderMaterial) Transparent() bool {
	return sm.Get("transparent").Bool()
}
func (sm *ShaderMaterial) SetTransparent(v bool) {
	sm.Set("transparent", v)
}
func (sm *ShaderMaterial) Type() string {
	return sm.Get("type").String()
}
func (sm *ShaderMaterial) SetType(v string) {
	sm.Set("type", v)
}
func (sm *ShaderMaterial) Uniforms() js.Value {
	return sm.Get("uniforms")
}
func (sm *ShaderMaterial) SetUniforms(v js.Value) {
	sm.Set("uniforms", v)
}
func (sm *ShaderMaterial) UserData() js.Value {
	return sm.Get("userData")
}
func (sm *ShaderMaterial) SetUserData(v js.Value) {
	sm.Set("userData", v)
}
func (sm *ShaderMaterial) Uuid() string {
	return sm.Get("uuid").String()
}
func (sm *ShaderMaterial) SetUuid(v string) {
	sm.Set("uuid", v)
}
func (sm *ShaderMaterial) VertexColors() *Colors {
	return &Colors{Value: sm.Get("vertexColors")}
}
func (sm *ShaderMaterial) SetVertexColors(v *Colors) {
	sm.Set("vertexColors", v)
}
func (sm *ShaderMaterial) VertexShader() string {
	return sm.Get("vertexShader").String()
}
func (sm *ShaderMaterial) SetVertexShader(v string) {
	sm.Set("vertexShader", v)
}
func (sm *ShaderMaterial) VertexTangents() bool {
	return sm.Get("vertexTangents").Bool()
}
func (sm *ShaderMaterial) SetVertexTangents(v bool) {
	sm.Set("vertexTangents", v)
}
func (sm *ShaderMaterial) Visible() bool {
	return sm.Get("visible").Bool()
}
func (sm *ShaderMaterial) SetVisible(v bool) {
	sm.Set("visible", v)
}
func (sm *ShaderMaterial) Wireframe() bool {
	return sm.Get("wireframe").Bool()
}
func (sm *ShaderMaterial) SetWireframe(v bool) {
	sm.Set("wireframe", v)
}
func (sm *ShaderMaterial) WireframeLinewidth() int {
	return sm.Get("wireframeLinewidth").Int()
}
func (sm *ShaderMaterial) SetWireframeLinewidth(v int) {
	sm.Set("wireframeLinewidth", v)
}
func (sm *ShaderMaterial) AddEventListener(typ string, listener js.Value) {
	sm.Call("addEventListener", typ, listener)
}
func (sm *ShaderMaterial) Clone() this {
	return this(sm.Call("clone"))
}
func (sm *ShaderMaterial) Copy(material *Material) this {
	return this(sm.Call("copy", material))
}
func (sm *ShaderMaterial) DispatchEvent(event js.Value) {
	sm.Call("dispatchEvent", event)
}
func (sm *ShaderMaterial) Dispose() {
	sm.Call("dispose")
}
func (sm *ShaderMaterial) HasEventListener(typ string, listener js.Value) bool {
	return sm.Call("hasEventListener", typ, listener).Bool()
}
func (sm *ShaderMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	sm.Call("onBeforeCompile", shader, renderer)
}
func (sm *ShaderMaterial) RemoveEventListener(typ string, listener js.Value) {
	sm.Call("removeEventListener", typ, listener)
}
func (sm *ShaderMaterial) SetValues(parameters ShaderMaterialParameters) {
	sm.Call("setValues", parameters)
}
func (sm *ShaderMaterial) ToJSON(meta js.Value) js.Value {
	return sm.Call("toJSON", meta)
}
func (sm *ShaderMaterial) Update() {
	sm.Call("update")
}

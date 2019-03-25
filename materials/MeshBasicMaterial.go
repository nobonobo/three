package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
	"github.com/nobonobo/three/textures"
)

type MeshBasicMaterialParameters interface {
}
type MeshBasicMaterial struct {
	js.Value
}

func NewMeshBasicMaterial(parameters MeshBasicMaterialParameters) *MeshBasicMaterial {
	return &MeshBasicMaterial{Value: get("MeshBasicMaterial").New(parameters)}
}
func (mbm *MeshBasicMaterial) AlphaMap() textures.Texture {
	return textures.Texture(mbm.Get("alphaMap"))
}
func (mbm *MeshBasicMaterial) SetAlphaMap(v textures.Texture) {
	mbm.Set("alphaMap", v)
}
func (mbm *MeshBasicMaterial) AlphaTest() int {
	return mbm.Get("alphaTest").Int()
}
func (mbm *MeshBasicMaterial) SetAlphaTest(v int) {
	mbm.Set("alphaTest", v)
}
func (mbm *MeshBasicMaterial) AoMap() textures.Texture {
	return textures.Texture(mbm.Get("aoMap"))
}
func (mbm *MeshBasicMaterial) SetAoMap(v textures.Texture) {
	mbm.Set("aoMap", v)
}
func (mbm *MeshBasicMaterial) AoMapIntensity() int {
	return mbm.Get("aoMapIntensity").Int()
}
func (mbm *MeshBasicMaterial) SetAoMapIntensity(v int) {
	mbm.Set("aoMapIntensity", v)
}
func (mbm *MeshBasicMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: mbm.Get("blendDst")}
}
func (mbm *MeshBasicMaterial) SetBlendDst(v *BlendingDstFactor) {
	mbm.Set("blendDst", v)
}
func (mbm *MeshBasicMaterial) BlendDstAlpha() int {
	return mbm.Get("blendDstAlpha").Int()
}
func (mbm *MeshBasicMaterial) SetBlendDstAlpha(v int) {
	mbm.Set("blendDstAlpha", v)
}
func (mbm *MeshBasicMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: mbm.Get("blendEquation")}
}
func (mbm *MeshBasicMaterial) SetBlendEquation(v *BlendingEquation) {
	mbm.Set("blendEquation", v)
}
func (mbm *MeshBasicMaterial) BlendEquationAlpha() int {
	return mbm.Get("blendEquationAlpha").Int()
}
func (mbm *MeshBasicMaterial) SetBlendEquationAlpha(v int) {
	mbm.Set("blendEquationAlpha", v)
}
func (mbm *MeshBasicMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(mbm.Get("blendSrc"))
}
func (mbm *MeshBasicMaterial) SetBlendSrc(v BlendingSrcFactor) {
	mbm.Set("blendSrc", v)
}
func (mbm *MeshBasicMaterial) BlendSrcAlpha() int {
	return mbm.Get("blendSrcAlpha").Int()
}
func (mbm *MeshBasicMaterial) SetBlendSrcAlpha(v int) {
	mbm.Set("blendSrcAlpha", v)
}
func (mbm *MeshBasicMaterial) Blending() *Blending {
	return &Blending{Value: mbm.Get("blending")}
}
func (mbm *MeshBasicMaterial) SetBlending(v *Blending) {
	mbm.Set("blending", v)
}
func (mbm *MeshBasicMaterial) ClipIntersection() bool {
	return mbm.Get("clipIntersection").Bool()
}
func (mbm *MeshBasicMaterial) SetClipIntersection(v bool) {
	mbm.Set("clipIntersection", v)
}
func (mbm *MeshBasicMaterial) ClipShadows() bool {
	return mbm.Get("clipShadows").Bool()
}
func (mbm *MeshBasicMaterial) SetClipShadows(v bool) {
	mbm.Set("clipShadows", v)
}
func (mbm *MeshBasicMaterial) ClippingPlanes() js.Value {
	return mbm.Get("clippingPlanes")
}
func (mbm *MeshBasicMaterial) SetClippingPlanes(v js.Value) {
	mbm.Set("clippingPlanes", v)
}
func (mbm *MeshBasicMaterial) Color() *math.Color {
	return &math.Color{Value: mbm.Get("color")}
}
func (mbm *MeshBasicMaterial) SetColor(v *math.Color) {
	mbm.Set("color", v)
}
func (mbm *MeshBasicMaterial) ColorWrite() bool {
	return mbm.Get("colorWrite").Bool()
}
func (mbm *MeshBasicMaterial) SetColorWrite(v bool) {
	mbm.Set("colorWrite", v)
}
func (mbm *MeshBasicMaterial) Combine() *Combine {
	return &Combine{Value: mbm.Get("combine")}
}
func (mbm *MeshBasicMaterial) SetCombine(v *Combine) {
	mbm.Set("combine", v)
}
func (mbm *MeshBasicMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: mbm.Get("depthFunc")}
}
func (mbm *MeshBasicMaterial) SetDepthFunc(v *DepthModes) {
	mbm.Set("depthFunc", v)
}
func (mbm *MeshBasicMaterial) DepthTest() bool {
	return mbm.Get("depthTest").Bool()
}
func (mbm *MeshBasicMaterial) SetDepthTest(v bool) {
	mbm.Set("depthTest", v)
}
func (mbm *MeshBasicMaterial) DepthWrite() bool {
	return mbm.Get("depthWrite").Bool()
}
func (mbm *MeshBasicMaterial) SetDepthWrite(v bool) {
	mbm.Set("depthWrite", v)
}
func (mbm *MeshBasicMaterial) Dithering() bool {
	return mbm.Get("dithering").Bool()
}
func (mbm *MeshBasicMaterial) SetDithering(v bool) {
	mbm.Set("dithering", v)
}
func (mbm *MeshBasicMaterial) EnvMap() textures.Texture {
	return textures.Texture(mbm.Get("envMap"))
}
func (mbm *MeshBasicMaterial) SetEnvMap(v textures.Texture) {
	mbm.Set("envMap", v)
}
func (mbm *MeshBasicMaterial) FlatShading() bool {
	return mbm.Get("flatShading").Bool()
}
func (mbm *MeshBasicMaterial) SetFlatShading(v bool) {
	mbm.Set("flatShading", v)
}
func (mbm *MeshBasicMaterial) Fog() bool {
	return mbm.Get("fog").Bool()
}
func (mbm *MeshBasicMaterial) SetFog(v bool) {
	mbm.Set("fog", v)
}
func (mbm *MeshBasicMaterial) Id() int {
	return mbm.Get("id").Int()
}
func (mbm *MeshBasicMaterial) SetId(v int) {
	mbm.Set("id", v)
}
func (mbm *MeshBasicMaterial) IsMaterial() bool {
	return mbm.Get("isMaterial").Bool()
}
func (mbm *MeshBasicMaterial) SetIsMaterial(v bool) {
	mbm.Set("isMaterial", v)
}
func (mbm *MeshBasicMaterial) Lights() bool {
	return mbm.Get("lights").Bool()
}
func (mbm *MeshBasicMaterial) SetLights(v bool) {
	mbm.Set("lights", v)
}
func (mbm *MeshBasicMaterial) Map() textures.Texture {
	return textures.Texture(mbm.Get("map"))
}
func (mbm *MeshBasicMaterial) SetMap(v textures.Texture) {
	mbm.Set("map", v)
}
func (mbm *MeshBasicMaterial) MorphTargets() bool {
	return mbm.Get("morphTargets").Bool()
}
func (mbm *MeshBasicMaterial) SetMorphTargets(v bool) {
	mbm.Set("morphTargets", v)
}
func (mbm *MeshBasicMaterial) Name() string {
	return mbm.Get("name").String()
}
func (mbm *MeshBasicMaterial) SetName(v string) {
	mbm.Set("name", v)
}
func (mbm *MeshBasicMaterial) NeedsUpdate() bool {
	return mbm.Get("needsUpdate").Bool()
}
func (mbm *MeshBasicMaterial) SetNeedsUpdate(v bool) {
	mbm.Set("needsUpdate", v)
}
func (mbm *MeshBasicMaterial) Opacity() int {
	return mbm.Get("opacity").Int()
}
func (mbm *MeshBasicMaterial) SetOpacity(v int) {
	mbm.Set("opacity", v)
}
func (mbm *MeshBasicMaterial) Overdraw() int {
	return mbm.Get("overdraw").Int()
}
func (mbm *MeshBasicMaterial) SetOverdraw(v int) {
	mbm.Set("overdraw", v)
}
func (mbm *MeshBasicMaterial) PolygonOffset() bool {
	return mbm.Get("polygonOffset").Bool()
}
func (mbm *MeshBasicMaterial) SetPolygonOffset(v bool) {
	mbm.Set("polygonOffset", v)
}
func (mbm *MeshBasicMaterial) PolygonOffsetFactor() int {
	return mbm.Get("polygonOffsetFactor").Int()
}
func (mbm *MeshBasicMaterial) SetPolygonOffsetFactor(v int) {
	mbm.Set("polygonOffsetFactor", v)
}
func (mbm *MeshBasicMaterial) PolygonOffsetUnits() int {
	return mbm.Get("polygonOffsetUnits").Int()
}
func (mbm *MeshBasicMaterial) SetPolygonOffsetUnits(v int) {
	mbm.Set("polygonOffsetUnits", v)
}
func (mbm *MeshBasicMaterial) Precision() string {
	return mbm.Get("precision").String()
}
func (mbm *MeshBasicMaterial) SetPrecision(v string) {
	mbm.Set("precision", v)
}
func (mbm *MeshBasicMaterial) PremultipliedAlpha() bool {
	return mbm.Get("premultipliedAlpha").Bool()
}
func (mbm *MeshBasicMaterial) SetPremultipliedAlpha(v bool) {
	mbm.Set("premultipliedAlpha", v)
}
func (mbm *MeshBasicMaterial) Reflectivity() int {
	return mbm.Get("reflectivity").Int()
}
func (mbm *MeshBasicMaterial) SetReflectivity(v int) {
	mbm.Set("reflectivity", v)
}
func (mbm *MeshBasicMaterial) RefractionRatio() int {
	return mbm.Get("refractionRatio").Int()
}
func (mbm *MeshBasicMaterial) SetRefractionRatio(v int) {
	mbm.Set("refractionRatio", v)
}
func (mbm *MeshBasicMaterial) Side() *Side {
	return &Side{Value: mbm.Get("side")}
}
func (mbm *MeshBasicMaterial) SetSide(v *Side) {
	mbm.Set("side", v)
}
func (mbm *MeshBasicMaterial) Skinning() bool {
	return mbm.Get("skinning").Bool()
}
func (mbm *MeshBasicMaterial) SetSkinning(v bool) {
	mbm.Set("skinning", v)
}
func (mbm *MeshBasicMaterial) SpecularMap() textures.Texture {
	return textures.Texture(mbm.Get("specularMap"))
}
func (mbm *MeshBasicMaterial) SetSpecularMap(v textures.Texture) {
	mbm.Set("specularMap", v)
}
func (mbm *MeshBasicMaterial) Transparent() bool {
	return mbm.Get("transparent").Bool()
}
func (mbm *MeshBasicMaterial) SetTransparent(v bool) {
	mbm.Set("transparent", v)
}
func (mbm *MeshBasicMaterial) Type() string {
	return mbm.Get("type").String()
}
func (mbm *MeshBasicMaterial) SetType(v string) {
	mbm.Set("type", v)
}
func (mbm *MeshBasicMaterial) UserData() js.Value {
	return mbm.Get("userData")
}
func (mbm *MeshBasicMaterial) SetUserData(v js.Value) {
	mbm.Set("userData", v)
}
func (mbm *MeshBasicMaterial) Uuid() string {
	return mbm.Get("uuid").String()
}
func (mbm *MeshBasicMaterial) SetUuid(v string) {
	mbm.Set("uuid", v)
}
func (mbm *MeshBasicMaterial) VertexColors() *Colors {
	return &Colors{Value: mbm.Get("vertexColors")}
}
func (mbm *MeshBasicMaterial) SetVertexColors(v *Colors) {
	mbm.Set("vertexColors", v)
}
func (mbm *MeshBasicMaterial) VertexTangents() bool {
	return mbm.Get("vertexTangents").Bool()
}
func (mbm *MeshBasicMaterial) SetVertexTangents(v bool) {
	mbm.Set("vertexTangents", v)
}
func (mbm *MeshBasicMaterial) Visible() bool {
	return mbm.Get("visible").Bool()
}
func (mbm *MeshBasicMaterial) SetVisible(v bool) {
	mbm.Set("visible", v)
}
func (mbm *MeshBasicMaterial) Wireframe() bool {
	return mbm.Get("wireframe").Bool()
}
func (mbm *MeshBasicMaterial) SetWireframe(v bool) {
	mbm.Set("wireframe", v)
}
func (mbm *MeshBasicMaterial) WireframeLinecap() string {
	return mbm.Get("wireframeLinecap").String()
}
func (mbm *MeshBasicMaterial) SetWireframeLinecap(v string) {
	mbm.Set("wireframeLinecap", v)
}
func (mbm *MeshBasicMaterial) WireframeLinejoin() string {
	return mbm.Get("wireframeLinejoin").String()
}
func (mbm *MeshBasicMaterial) SetWireframeLinejoin(v string) {
	mbm.Set("wireframeLinejoin", v)
}
func (mbm *MeshBasicMaterial) WireframeLinewidth() int {
	return mbm.Get("wireframeLinewidth").Int()
}
func (mbm *MeshBasicMaterial) SetWireframeLinewidth(v int) {
	mbm.Set("wireframeLinewidth", v)
}
func (mbm *MeshBasicMaterial) AddEventListener(typ string, listener js.Value) {
	mbm.Call("addEventListener", typ, listener)
}
func (mbm *MeshBasicMaterial) Clone() this {
	return this(mbm.Call("clone"))
}
func (mbm *MeshBasicMaterial) Copy(material *Material) this {
	return this(mbm.Call("copy", material))
}
func (mbm *MeshBasicMaterial) DispatchEvent(event js.Value) {
	mbm.Call("dispatchEvent", event)
}
func (mbm *MeshBasicMaterial) Dispose() {
	mbm.Call("dispose")
}
func (mbm *MeshBasicMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mbm.Call("hasEventListener", typ, listener).Bool()
}
func (mbm *MeshBasicMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	mbm.Call("onBeforeCompile", shader, renderer)
}
func (mbm *MeshBasicMaterial) RemoveEventListener(typ string, listener js.Value) {
	mbm.Call("removeEventListener", typ, listener)
}
func (mbm *MeshBasicMaterial) SetValues(parameters MeshBasicMaterialParameters) {
	mbm.Call("setValues", parameters)
}
func (mbm *MeshBasicMaterial) ToJSON(meta js.Value) js.Value {
	return mbm.Call("toJSON", meta)
}
func (mbm *MeshBasicMaterial) Update() {
	mbm.Call("update")
}

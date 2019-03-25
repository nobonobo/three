package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
	"github.com/nobonobo/three/textures"
)

type MeshStandardMaterialParameters interface {
}
type MeshStandardMaterial struct {
	js.Value
}

func NewMeshStandardMaterial(parameters MeshStandardMaterialParameters) *MeshStandardMaterial {
	return &MeshStandardMaterial{Value: get("MeshStandardMaterial").New(parameters)}
}
func (msm *MeshStandardMaterial) AlphaMap() textures.Texture {
	return textures.Texture(msm.Get("alphaMap"))
}
func (msm *MeshStandardMaterial) SetAlphaMap(v textures.Texture) {
	msm.Set("alphaMap", v)
}
func (msm *MeshStandardMaterial) AlphaTest() int {
	return msm.Get("alphaTest").Int()
}
func (msm *MeshStandardMaterial) SetAlphaTest(v int) {
	msm.Set("alphaTest", v)
}
func (msm *MeshStandardMaterial) AoMap() textures.Texture {
	return textures.Texture(msm.Get("aoMap"))
}
func (msm *MeshStandardMaterial) SetAoMap(v textures.Texture) {
	msm.Set("aoMap", v)
}
func (msm *MeshStandardMaterial) AoMapIntensity() int {
	return msm.Get("aoMapIntensity").Int()
}
func (msm *MeshStandardMaterial) SetAoMapIntensity(v int) {
	msm.Set("aoMapIntensity", v)
}
func (msm *MeshStandardMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: msm.Get("blendDst")}
}
func (msm *MeshStandardMaterial) SetBlendDst(v *BlendingDstFactor) {
	msm.Set("blendDst", v)
}
func (msm *MeshStandardMaterial) BlendDstAlpha() int {
	return msm.Get("blendDstAlpha").Int()
}
func (msm *MeshStandardMaterial) SetBlendDstAlpha(v int) {
	msm.Set("blendDstAlpha", v)
}
func (msm *MeshStandardMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: msm.Get("blendEquation")}
}
func (msm *MeshStandardMaterial) SetBlendEquation(v *BlendingEquation) {
	msm.Set("blendEquation", v)
}
func (msm *MeshStandardMaterial) BlendEquationAlpha() int {
	return msm.Get("blendEquationAlpha").Int()
}
func (msm *MeshStandardMaterial) SetBlendEquationAlpha(v int) {
	msm.Set("blendEquationAlpha", v)
}
func (msm *MeshStandardMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(msm.Get("blendSrc"))
}
func (msm *MeshStandardMaterial) SetBlendSrc(v BlendingSrcFactor) {
	msm.Set("blendSrc", v)
}
func (msm *MeshStandardMaterial) BlendSrcAlpha() int {
	return msm.Get("blendSrcAlpha").Int()
}
func (msm *MeshStandardMaterial) SetBlendSrcAlpha(v int) {
	msm.Set("blendSrcAlpha", v)
}
func (msm *MeshStandardMaterial) Blending() *Blending {
	return &Blending{Value: msm.Get("blending")}
}
func (msm *MeshStandardMaterial) SetBlending(v *Blending) {
	msm.Set("blending", v)
}
func (msm *MeshStandardMaterial) BumpMap() textures.Texture {
	return textures.Texture(msm.Get("bumpMap"))
}
func (msm *MeshStandardMaterial) SetBumpMap(v textures.Texture) {
	msm.Set("bumpMap", v)
}
func (msm *MeshStandardMaterial) BumpScale() int {
	return msm.Get("bumpScale").Int()
}
func (msm *MeshStandardMaterial) SetBumpScale(v int) {
	msm.Set("bumpScale", v)
}
func (msm *MeshStandardMaterial) ClipIntersection() bool {
	return msm.Get("clipIntersection").Bool()
}
func (msm *MeshStandardMaterial) SetClipIntersection(v bool) {
	msm.Set("clipIntersection", v)
}
func (msm *MeshStandardMaterial) ClipShadows() bool {
	return msm.Get("clipShadows").Bool()
}
func (msm *MeshStandardMaterial) SetClipShadows(v bool) {
	msm.Set("clipShadows", v)
}
func (msm *MeshStandardMaterial) ClippingPlanes() js.Value {
	return msm.Get("clippingPlanes")
}
func (msm *MeshStandardMaterial) SetClippingPlanes(v js.Value) {
	msm.Set("clippingPlanes", v)
}
func (msm *MeshStandardMaterial) Color() *math.Color {
	return &math.Color{Value: msm.Get("color")}
}
func (msm *MeshStandardMaterial) SetColor(v *math.Color) {
	msm.Set("color", v)
}
func (msm *MeshStandardMaterial) ColorWrite() bool {
	return msm.Get("colorWrite").Bool()
}
func (msm *MeshStandardMaterial) SetColorWrite(v bool) {
	msm.Set("colorWrite", v)
}
func (msm *MeshStandardMaterial) Defines() js.Value {
	return msm.Get("defines")
}
func (msm *MeshStandardMaterial) SetDefines(v js.Value) {
	msm.Set("defines", v)
}
func (msm *MeshStandardMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: msm.Get("depthFunc")}
}
func (msm *MeshStandardMaterial) SetDepthFunc(v *DepthModes) {
	msm.Set("depthFunc", v)
}
func (msm *MeshStandardMaterial) DepthTest() bool {
	return msm.Get("depthTest").Bool()
}
func (msm *MeshStandardMaterial) SetDepthTest(v bool) {
	msm.Set("depthTest", v)
}
func (msm *MeshStandardMaterial) DepthWrite() bool {
	return msm.Get("depthWrite").Bool()
}
func (msm *MeshStandardMaterial) SetDepthWrite(v bool) {
	msm.Set("depthWrite", v)
}
func (msm *MeshStandardMaterial) DisplacementBias() int {
	return msm.Get("displacementBias").Int()
}
func (msm *MeshStandardMaterial) SetDisplacementBias(v int) {
	msm.Set("displacementBias", v)
}
func (msm *MeshStandardMaterial) DisplacementMap() textures.Texture {
	return textures.Texture(msm.Get("displacementMap"))
}
func (msm *MeshStandardMaterial) SetDisplacementMap(v textures.Texture) {
	msm.Set("displacementMap", v)
}
func (msm *MeshStandardMaterial) DisplacementScale() int {
	return msm.Get("displacementScale").Int()
}
func (msm *MeshStandardMaterial) SetDisplacementScale(v int) {
	msm.Set("displacementScale", v)
}
func (msm *MeshStandardMaterial) Dithering() bool {
	return msm.Get("dithering").Bool()
}
func (msm *MeshStandardMaterial) SetDithering(v bool) {
	msm.Set("dithering", v)
}
func (msm *MeshStandardMaterial) Emissive() *math.Color {
	return &math.Color{Value: msm.Get("emissive")}
}
func (msm *MeshStandardMaterial) SetEmissive(v *math.Color) {
	msm.Set("emissive", v)
}
func (msm *MeshStandardMaterial) EmissiveIntensity() int {
	return msm.Get("emissiveIntensity").Int()
}
func (msm *MeshStandardMaterial) SetEmissiveIntensity(v int) {
	msm.Set("emissiveIntensity", v)
}
func (msm *MeshStandardMaterial) EmissiveMap() textures.Texture {
	return textures.Texture(msm.Get("emissiveMap"))
}
func (msm *MeshStandardMaterial) SetEmissiveMap(v textures.Texture) {
	msm.Set("emissiveMap", v)
}
func (msm *MeshStandardMaterial) EnvMap() textures.Texture {
	return textures.Texture(msm.Get("envMap"))
}
func (msm *MeshStandardMaterial) SetEnvMap(v textures.Texture) {
	msm.Set("envMap", v)
}
func (msm *MeshStandardMaterial) EnvMapIntensity() int {
	return msm.Get("envMapIntensity").Int()
}
func (msm *MeshStandardMaterial) SetEnvMapIntensity(v int) {
	msm.Set("envMapIntensity", v)
}
func (msm *MeshStandardMaterial) FlatShading() bool {
	return msm.Get("flatShading").Bool()
}
func (msm *MeshStandardMaterial) SetFlatShading(v bool) {
	msm.Set("flatShading", v)
}
func (msm *MeshStandardMaterial) Fog() bool {
	return msm.Get("fog").Bool()
}
func (msm *MeshStandardMaterial) SetFog(v bool) {
	msm.Set("fog", v)
}
func (msm *MeshStandardMaterial) Id() int {
	return msm.Get("id").Int()
}
func (msm *MeshStandardMaterial) SetId(v int) {
	msm.Set("id", v)
}
func (msm *MeshStandardMaterial) IsMaterial() bool {
	return msm.Get("isMaterial").Bool()
}
func (msm *MeshStandardMaterial) SetIsMaterial(v bool) {
	msm.Set("isMaterial", v)
}
func (msm *MeshStandardMaterial) LightMap() textures.Texture {
	return textures.Texture(msm.Get("lightMap"))
}
func (msm *MeshStandardMaterial) SetLightMap(v textures.Texture) {
	msm.Set("lightMap", v)
}
func (msm *MeshStandardMaterial) LightMapIntensity() int {
	return msm.Get("lightMapIntensity").Int()
}
func (msm *MeshStandardMaterial) SetLightMapIntensity(v int) {
	msm.Set("lightMapIntensity", v)
}
func (msm *MeshStandardMaterial) Lights() bool {
	return msm.Get("lights").Bool()
}
func (msm *MeshStandardMaterial) SetLights(v bool) {
	msm.Set("lights", v)
}
func (msm *MeshStandardMaterial) Map() textures.Texture {
	return textures.Texture(msm.Get("map"))
}
func (msm *MeshStandardMaterial) SetMap(v textures.Texture) {
	msm.Set("map", v)
}
func (msm *MeshStandardMaterial) Metalness() int {
	return msm.Get("metalness").Int()
}
func (msm *MeshStandardMaterial) SetMetalness(v int) {
	msm.Set("metalness", v)
}
func (msm *MeshStandardMaterial) MetalnessMap() textures.Texture {
	return textures.Texture(msm.Get("metalnessMap"))
}
func (msm *MeshStandardMaterial) SetMetalnessMap(v textures.Texture) {
	msm.Set("metalnessMap", v)
}
func (msm *MeshStandardMaterial) MorphNormals() bool {
	return msm.Get("morphNormals").Bool()
}
func (msm *MeshStandardMaterial) SetMorphNormals(v bool) {
	msm.Set("morphNormals", v)
}
func (msm *MeshStandardMaterial) MorphTargets() bool {
	return msm.Get("morphTargets").Bool()
}
func (msm *MeshStandardMaterial) SetMorphTargets(v bool) {
	msm.Set("morphTargets", v)
}
func (msm *MeshStandardMaterial) Name() string {
	return msm.Get("name").String()
}
func (msm *MeshStandardMaterial) SetName(v string) {
	msm.Set("name", v)
}
func (msm *MeshStandardMaterial) NeedsUpdate() bool {
	return msm.Get("needsUpdate").Bool()
}
func (msm *MeshStandardMaterial) SetNeedsUpdate(v bool) {
	msm.Set("needsUpdate", v)
}
func (msm *MeshStandardMaterial) NormalMap() textures.Texture {
	return textures.Texture(msm.Get("normalMap"))
}
func (msm *MeshStandardMaterial) SetNormalMap(v textures.Texture) {
	msm.Set("normalMap", v)
}
func (msm *MeshStandardMaterial) NormalScale() int {
	return msm.Get("normalScale").Int()
}
func (msm *MeshStandardMaterial) SetNormalScale(v int) {
	msm.Set("normalScale", v)
}
func (msm *MeshStandardMaterial) Opacity() int {
	return msm.Get("opacity").Int()
}
func (msm *MeshStandardMaterial) SetOpacity(v int) {
	msm.Set("opacity", v)
}
func (msm *MeshStandardMaterial) Overdraw() int {
	return msm.Get("overdraw").Int()
}
func (msm *MeshStandardMaterial) SetOverdraw(v int) {
	msm.Set("overdraw", v)
}
func (msm *MeshStandardMaterial) PolygonOffset() bool {
	return msm.Get("polygonOffset").Bool()
}
func (msm *MeshStandardMaterial) SetPolygonOffset(v bool) {
	msm.Set("polygonOffset", v)
}
func (msm *MeshStandardMaterial) PolygonOffsetFactor() int {
	return msm.Get("polygonOffsetFactor").Int()
}
func (msm *MeshStandardMaterial) SetPolygonOffsetFactor(v int) {
	msm.Set("polygonOffsetFactor", v)
}
func (msm *MeshStandardMaterial) PolygonOffsetUnits() int {
	return msm.Get("polygonOffsetUnits").Int()
}
func (msm *MeshStandardMaterial) SetPolygonOffsetUnits(v int) {
	msm.Set("polygonOffsetUnits", v)
}
func (msm *MeshStandardMaterial) Precision() string {
	return msm.Get("precision").String()
}
func (msm *MeshStandardMaterial) SetPrecision(v string) {
	msm.Set("precision", v)
}
func (msm *MeshStandardMaterial) PremultipliedAlpha() bool {
	return msm.Get("premultipliedAlpha").Bool()
}
func (msm *MeshStandardMaterial) SetPremultipliedAlpha(v bool) {
	msm.Set("premultipliedAlpha", v)
}
func (msm *MeshStandardMaterial) RefractionRatio() int {
	return msm.Get("refractionRatio").Int()
}
func (msm *MeshStandardMaterial) SetRefractionRatio(v int) {
	msm.Set("refractionRatio", v)
}
func (msm *MeshStandardMaterial) Roughness() int {
	return msm.Get("roughness").Int()
}
func (msm *MeshStandardMaterial) SetRoughness(v int) {
	msm.Set("roughness", v)
}
func (msm *MeshStandardMaterial) RoughnessMap() textures.Texture {
	return textures.Texture(msm.Get("roughnessMap"))
}
func (msm *MeshStandardMaterial) SetRoughnessMap(v textures.Texture) {
	msm.Set("roughnessMap", v)
}
func (msm *MeshStandardMaterial) Side() *Side {
	return &Side{Value: msm.Get("side")}
}
func (msm *MeshStandardMaterial) SetSide(v *Side) {
	msm.Set("side", v)
}
func (msm *MeshStandardMaterial) Skinning() bool {
	return msm.Get("skinning").Bool()
}
func (msm *MeshStandardMaterial) SetSkinning(v bool) {
	msm.Set("skinning", v)
}
func (msm *MeshStandardMaterial) Transparent() bool {
	return msm.Get("transparent").Bool()
}
func (msm *MeshStandardMaterial) SetTransparent(v bool) {
	msm.Set("transparent", v)
}
func (msm *MeshStandardMaterial) Type() string {
	return msm.Get("type").String()
}
func (msm *MeshStandardMaterial) SetType(v string) {
	msm.Set("type", v)
}
func (msm *MeshStandardMaterial) UserData() js.Value {
	return msm.Get("userData")
}
func (msm *MeshStandardMaterial) SetUserData(v js.Value) {
	msm.Set("userData", v)
}
func (msm *MeshStandardMaterial) Uuid() string {
	return msm.Get("uuid").String()
}
func (msm *MeshStandardMaterial) SetUuid(v string) {
	msm.Set("uuid", v)
}
func (msm *MeshStandardMaterial) VertexColors() *Colors {
	return &Colors{Value: msm.Get("vertexColors")}
}
func (msm *MeshStandardMaterial) SetVertexColors(v *Colors) {
	msm.Set("vertexColors", v)
}
func (msm *MeshStandardMaterial) VertexTangents() bool {
	return msm.Get("vertexTangents").Bool()
}
func (msm *MeshStandardMaterial) SetVertexTangents(v bool) {
	msm.Set("vertexTangents", v)
}
func (msm *MeshStandardMaterial) Visible() bool {
	return msm.Get("visible").Bool()
}
func (msm *MeshStandardMaterial) SetVisible(v bool) {
	msm.Set("visible", v)
}
func (msm *MeshStandardMaterial) Wireframe() bool {
	return msm.Get("wireframe").Bool()
}
func (msm *MeshStandardMaterial) SetWireframe(v bool) {
	msm.Set("wireframe", v)
}
func (msm *MeshStandardMaterial) WireframeLinewidth() int {
	return msm.Get("wireframeLinewidth").Int()
}
func (msm *MeshStandardMaterial) SetWireframeLinewidth(v int) {
	msm.Set("wireframeLinewidth", v)
}
func (msm *MeshStandardMaterial) AddEventListener(typ string, listener js.Value) {
	msm.Call("addEventListener", typ, listener)
}
func (msm *MeshStandardMaterial) Clone() this {
	return this(msm.Call("clone"))
}
func (msm *MeshStandardMaterial) Copy(material *Material) this {
	return this(msm.Call("copy", material))
}
func (msm *MeshStandardMaterial) DispatchEvent(event js.Value) {
	msm.Call("dispatchEvent", event)
}
func (msm *MeshStandardMaterial) Dispose() {
	msm.Call("dispose")
}
func (msm *MeshStandardMaterial) HasEventListener(typ string, listener js.Value) bool {
	return msm.Call("hasEventListener", typ, listener).Bool()
}
func (msm *MeshStandardMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	msm.Call("onBeforeCompile", shader, renderer)
}
func (msm *MeshStandardMaterial) RemoveEventListener(typ string, listener js.Value) {
	msm.Call("removeEventListener", typ, listener)
}
func (msm *MeshStandardMaterial) SetValues(parameters MeshStandardMaterialParameters) {
	msm.Call("setValues", parameters)
}
func (msm *MeshStandardMaterial) ToJSON(meta js.Value) js.Value {
	return msm.Call("toJSON", meta)
}
func (msm *MeshStandardMaterial) Update() {
	msm.Call("update")
}

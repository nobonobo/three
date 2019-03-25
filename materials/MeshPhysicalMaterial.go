package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
	"github.com/nobonobo/three/textures"
)

type MeshPhysicalMaterialParameters interface {
}
type MeshPhysicalMaterial struct {
	js.Value
}

func NewMeshPhysicalMaterial(parameters MeshPhysicalMaterialParameters) *MeshPhysicalMaterial {
	return &MeshPhysicalMaterial{Value: get("MeshPhysicalMaterial").New(parameters)}
}
func (mpm *MeshPhysicalMaterial) AlphaMap() textures.Texture {
	return textures.Texture(mpm.Get("alphaMap"))
}
func (mpm *MeshPhysicalMaterial) SetAlphaMap(v textures.Texture) {
	mpm.Set("alphaMap", v)
}
func (mpm *MeshPhysicalMaterial) AlphaTest() int {
	return mpm.Get("alphaTest").Int()
}
func (mpm *MeshPhysicalMaterial) SetAlphaTest(v int) {
	mpm.Set("alphaTest", v)
}
func (mpm *MeshPhysicalMaterial) AoMap() textures.Texture {
	return textures.Texture(mpm.Get("aoMap"))
}
func (mpm *MeshPhysicalMaterial) SetAoMap(v textures.Texture) {
	mpm.Set("aoMap", v)
}
func (mpm *MeshPhysicalMaterial) AoMapIntensity() int {
	return mpm.Get("aoMapIntensity").Int()
}
func (mpm *MeshPhysicalMaterial) SetAoMapIntensity(v int) {
	mpm.Set("aoMapIntensity", v)
}
func (mpm *MeshPhysicalMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: mpm.Get("blendDst")}
}
func (mpm *MeshPhysicalMaterial) SetBlendDst(v *BlendingDstFactor) {
	mpm.Set("blendDst", v)
}
func (mpm *MeshPhysicalMaterial) BlendDstAlpha() int {
	return mpm.Get("blendDstAlpha").Int()
}
func (mpm *MeshPhysicalMaterial) SetBlendDstAlpha(v int) {
	mpm.Set("blendDstAlpha", v)
}
func (mpm *MeshPhysicalMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: mpm.Get("blendEquation")}
}
func (mpm *MeshPhysicalMaterial) SetBlendEquation(v *BlendingEquation) {
	mpm.Set("blendEquation", v)
}
func (mpm *MeshPhysicalMaterial) BlendEquationAlpha() int {
	return mpm.Get("blendEquationAlpha").Int()
}
func (mpm *MeshPhysicalMaterial) SetBlendEquationAlpha(v int) {
	mpm.Set("blendEquationAlpha", v)
}
func (mpm *MeshPhysicalMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(mpm.Get("blendSrc"))
}
func (mpm *MeshPhysicalMaterial) SetBlendSrc(v BlendingSrcFactor) {
	mpm.Set("blendSrc", v)
}
func (mpm *MeshPhysicalMaterial) BlendSrcAlpha() int {
	return mpm.Get("blendSrcAlpha").Int()
}
func (mpm *MeshPhysicalMaterial) SetBlendSrcAlpha(v int) {
	mpm.Set("blendSrcAlpha", v)
}
func (mpm *MeshPhysicalMaterial) Blending() *Blending {
	return &Blending{Value: mpm.Get("blending")}
}
func (mpm *MeshPhysicalMaterial) SetBlending(v *Blending) {
	mpm.Set("blending", v)
}
func (mpm *MeshPhysicalMaterial) BumpMap() textures.Texture {
	return textures.Texture(mpm.Get("bumpMap"))
}
func (mpm *MeshPhysicalMaterial) SetBumpMap(v textures.Texture) {
	mpm.Set("bumpMap", v)
}
func (mpm *MeshPhysicalMaterial) BumpScale() int {
	return mpm.Get("bumpScale").Int()
}
func (mpm *MeshPhysicalMaterial) SetBumpScale(v int) {
	mpm.Set("bumpScale", v)
}
func (mpm *MeshPhysicalMaterial) ClearCoat() int {
	return mpm.Get("clearCoat").Int()
}
func (mpm *MeshPhysicalMaterial) SetClearCoat(v int) {
	mpm.Set("clearCoat", v)
}
func (mpm *MeshPhysicalMaterial) ClearCoatRoughness() int {
	return mpm.Get("clearCoatRoughness").Int()
}
func (mpm *MeshPhysicalMaterial) SetClearCoatRoughness(v int) {
	mpm.Set("clearCoatRoughness", v)
}
func (mpm *MeshPhysicalMaterial) ClipIntersection() bool {
	return mpm.Get("clipIntersection").Bool()
}
func (mpm *MeshPhysicalMaterial) SetClipIntersection(v bool) {
	mpm.Set("clipIntersection", v)
}
func (mpm *MeshPhysicalMaterial) ClipShadows() bool {
	return mpm.Get("clipShadows").Bool()
}
func (mpm *MeshPhysicalMaterial) SetClipShadows(v bool) {
	mpm.Set("clipShadows", v)
}
func (mpm *MeshPhysicalMaterial) ClippingPlanes() js.Value {
	return mpm.Get("clippingPlanes")
}
func (mpm *MeshPhysicalMaterial) SetClippingPlanes(v js.Value) {
	mpm.Set("clippingPlanes", v)
}
func (mpm *MeshPhysicalMaterial) Color() *math.Color {
	return &math.Color{Value: mpm.Get("color")}
}
func (mpm *MeshPhysicalMaterial) SetColor(v *math.Color) {
	mpm.Set("color", v)
}
func (mpm *MeshPhysicalMaterial) ColorWrite() bool {
	return mpm.Get("colorWrite").Bool()
}
func (mpm *MeshPhysicalMaterial) SetColorWrite(v bool) {
	mpm.Set("colorWrite", v)
}
func (mpm *MeshPhysicalMaterial) Defines() js.Value {
	return mpm.Get("defines")
}
func (mpm *MeshPhysicalMaterial) SetDefines(v js.Value) {
	mpm.Set("defines", v)
}
func (mpm *MeshPhysicalMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: mpm.Get("depthFunc")}
}
func (mpm *MeshPhysicalMaterial) SetDepthFunc(v *DepthModes) {
	mpm.Set("depthFunc", v)
}
func (mpm *MeshPhysicalMaterial) DepthTest() bool {
	return mpm.Get("depthTest").Bool()
}
func (mpm *MeshPhysicalMaterial) SetDepthTest(v bool) {
	mpm.Set("depthTest", v)
}
func (mpm *MeshPhysicalMaterial) DepthWrite() bool {
	return mpm.Get("depthWrite").Bool()
}
func (mpm *MeshPhysicalMaterial) SetDepthWrite(v bool) {
	mpm.Set("depthWrite", v)
}
func (mpm *MeshPhysicalMaterial) DisplacementBias() int {
	return mpm.Get("displacementBias").Int()
}
func (mpm *MeshPhysicalMaterial) SetDisplacementBias(v int) {
	mpm.Set("displacementBias", v)
}
func (mpm *MeshPhysicalMaterial) DisplacementMap() textures.Texture {
	return textures.Texture(mpm.Get("displacementMap"))
}
func (mpm *MeshPhysicalMaterial) SetDisplacementMap(v textures.Texture) {
	mpm.Set("displacementMap", v)
}
func (mpm *MeshPhysicalMaterial) DisplacementScale() int {
	return mpm.Get("displacementScale").Int()
}
func (mpm *MeshPhysicalMaterial) SetDisplacementScale(v int) {
	mpm.Set("displacementScale", v)
}
func (mpm *MeshPhysicalMaterial) Dithering() bool {
	return mpm.Get("dithering").Bool()
}
func (mpm *MeshPhysicalMaterial) SetDithering(v bool) {
	mpm.Set("dithering", v)
}
func (mpm *MeshPhysicalMaterial) Emissive() *math.Color {
	return &math.Color{Value: mpm.Get("emissive")}
}
func (mpm *MeshPhysicalMaterial) SetEmissive(v *math.Color) {
	mpm.Set("emissive", v)
}
func (mpm *MeshPhysicalMaterial) EmissiveIntensity() int {
	return mpm.Get("emissiveIntensity").Int()
}
func (mpm *MeshPhysicalMaterial) SetEmissiveIntensity(v int) {
	mpm.Set("emissiveIntensity", v)
}
func (mpm *MeshPhysicalMaterial) EmissiveMap() textures.Texture {
	return textures.Texture(mpm.Get("emissiveMap"))
}
func (mpm *MeshPhysicalMaterial) SetEmissiveMap(v textures.Texture) {
	mpm.Set("emissiveMap", v)
}
func (mpm *MeshPhysicalMaterial) EnvMap() textures.Texture {
	return textures.Texture(mpm.Get("envMap"))
}
func (mpm *MeshPhysicalMaterial) SetEnvMap(v textures.Texture) {
	mpm.Set("envMap", v)
}
func (mpm *MeshPhysicalMaterial) EnvMapIntensity() int {
	return mpm.Get("envMapIntensity").Int()
}
func (mpm *MeshPhysicalMaterial) SetEnvMapIntensity(v int) {
	mpm.Set("envMapIntensity", v)
}
func (mpm *MeshPhysicalMaterial) FlatShading() bool {
	return mpm.Get("flatShading").Bool()
}
func (mpm *MeshPhysicalMaterial) SetFlatShading(v bool) {
	mpm.Set("flatShading", v)
}
func (mpm *MeshPhysicalMaterial) Fog() bool {
	return mpm.Get("fog").Bool()
}
func (mpm *MeshPhysicalMaterial) SetFog(v bool) {
	mpm.Set("fog", v)
}
func (mpm *MeshPhysicalMaterial) Id() int {
	return mpm.Get("id").Int()
}
func (mpm *MeshPhysicalMaterial) SetId(v int) {
	mpm.Set("id", v)
}
func (mpm *MeshPhysicalMaterial) IsMaterial() bool {
	return mpm.Get("isMaterial").Bool()
}
func (mpm *MeshPhysicalMaterial) SetIsMaterial(v bool) {
	mpm.Set("isMaterial", v)
}
func (mpm *MeshPhysicalMaterial) LightMap() textures.Texture {
	return textures.Texture(mpm.Get("lightMap"))
}
func (mpm *MeshPhysicalMaterial) SetLightMap(v textures.Texture) {
	mpm.Set("lightMap", v)
}
func (mpm *MeshPhysicalMaterial) LightMapIntensity() int {
	return mpm.Get("lightMapIntensity").Int()
}
func (mpm *MeshPhysicalMaterial) SetLightMapIntensity(v int) {
	mpm.Set("lightMapIntensity", v)
}
func (mpm *MeshPhysicalMaterial) Lights() bool {
	return mpm.Get("lights").Bool()
}
func (mpm *MeshPhysicalMaterial) SetLights(v bool) {
	mpm.Set("lights", v)
}
func (mpm *MeshPhysicalMaterial) Map() textures.Texture {
	return textures.Texture(mpm.Get("map"))
}
func (mpm *MeshPhysicalMaterial) SetMap(v textures.Texture) {
	mpm.Set("map", v)
}
func (mpm *MeshPhysicalMaterial) Metalness() int {
	return mpm.Get("metalness").Int()
}
func (mpm *MeshPhysicalMaterial) SetMetalness(v int) {
	mpm.Set("metalness", v)
}
func (mpm *MeshPhysicalMaterial) MetalnessMap() textures.Texture {
	return textures.Texture(mpm.Get("metalnessMap"))
}
func (mpm *MeshPhysicalMaterial) SetMetalnessMap(v textures.Texture) {
	mpm.Set("metalnessMap", v)
}
func (mpm *MeshPhysicalMaterial) MorphNormals() bool {
	return mpm.Get("morphNormals").Bool()
}
func (mpm *MeshPhysicalMaterial) SetMorphNormals(v bool) {
	mpm.Set("morphNormals", v)
}
func (mpm *MeshPhysicalMaterial) MorphTargets() bool {
	return mpm.Get("morphTargets").Bool()
}
func (mpm *MeshPhysicalMaterial) SetMorphTargets(v bool) {
	mpm.Set("morphTargets", v)
}
func (mpm *MeshPhysicalMaterial) Name() string {
	return mpm.Get("name").String()
}
func (mpm *MeshPhysicalMaterial) SetName(v string) {
	mpm.Set("name", v)
}
func (mpm *MeshPhysicalMaterial) NeedsUpdate() bool {
	return mpm.Get("needsUpdate").Bool()
}
func (mpm *MeshPhysicalMaterial) SetNeedsUpdate(v bool) {
	mpm.Set("needsUpdate", v)
}
func (mpm *MeshPhysicalMaterial) NormalMap() textures.Texture {
	return textures.Texture(mpm.Get("normalMap"))
}
func (mpm *MeshPhysicalMaterial) SetNormalMap(v textures.Texture) {
	mpm.Set("normalMap", v)
}
func (mpm *MeshPhysicalMaterial) NormalScale() int {
	return mpm.Get("normalScale").Int()
}
func (mpm *MeshPhysicalMaterial) SetNormalScale(v int) {
	mpm.Set("normalScale", v)
}
func (mpm *MeshPhysicalMaterial) Opacity() int {
	return mpm.Get("opacity").Int()
}
func (mpm *MeshPhysicalMaterial) SetOpacity(v int) {
	mpm.Set("opacity", v)
}
func (mpm *MeshPhysicalMaterial) Overdraw() int {
	return mpm.Get("overdraw").Int()
}
func (mpm *MeshPhysicalMaterial) SetOverdraw(v int) {
	mpm.Set("overdraw", v)
}
func (mpm *MeshPhysicalMaterial) PolygonOffset() bool {
	return mpm.Get("polygonOffset").Bool()
}
func (mpm *MeshPhysicalMaterial) SetPolygonOffset(v bool) {
	mpm.Set("polygonOffset", v)
}
func (mpm *MeshPhysicalMaterial) PolygonOffsetFactor() int {
	return mpm.Get("polygonOffsetFactor").Int()
}
func (mpm *MeshPhysicalMaterial) SetPolygonOffsetFactor(v int) {
	mpm.Set("polygonOffsetFactor", v)
}
func (mpm *MeshPhysicalMaterial) PolygonOffsetUnits() int {
	return mpm.Get("polygonOffsetUnits").Int()
}
func (mpm *MeshPhysicalMaterial) SetPolygonOffsetUnits(v int) {
	mpm.Set("polygonOffsetUnits", v)
}
func (mpm *MeshPhysicalMaterial) Precision() string {
	return mpm.Get("precision").String()
}
func (mpm *MeshPhysicalMaterial) SetPrecision(v string) {
	mpm.Set("precision", v)
}
func (mpm *MeshPhysicalMaterial) PremultipliedAlpha() bool {
	return mpm.Get("premultipliedAlpha").Bool()
}
func (mpm *MeshPhysicalMaterial) SetPremultipliedAlpha(v bool) {
	mpm.Set("premultipliedAlpha", v)
}
func (mpm *MeshPhysicalMaterial) Reflectivity() int {
	return mpm.Get("reflectivity").Int()
}
func (mpm *MeshPhysicalMaterial) SetReflectivity(v int) {
	mpm.Set("reflectivity", v)
}
func (mpm *MeshPhysicalMaterial) RefractionRatio() int {
	return mpm.Get("refractionRatio").Int()
}
func (mpm *MeshPhysicalMaterial) SetRefractionRatio(v int) {
	mpm.Set("refractionRatio", v)
}
func (mpm *MeshPhysicalMaterial) Roughness() int {
	return mpm.Get("roughness").Int()
}
func (mpm *MeshPhysicalMaterial) SetRoughness(v int) {
	mpm.Set("roughness", v)
}
func (mpm *MeshPhysicalMaterial) RoughnessMap() textures.Texture {
	return textures.Texture(mpm.Get("roughnessMap"))
}
func (mpm *MeshPhysicalMaterial) SetRoughnessMap(v textures.Texture) {
	mpm.Set("roughnessMap", v)
}
func (mpm *MeshPhysicalMaterial) Side() *Side {
	return &Side{Value: mpm.Get("side")}
}
func (mpm *MeshPhysicalMaterial) SetSide(v *Side) {
	mpm.Set("side", v)
}
func (mpm *MeshPhysicalMaterial) Skinning() bool {
	return mpm.Get("skinning").Bool()
}
func (mpm *MeshPhysicalMaterial) SetSkinning(v bool) {
	mpm.Set("skinning", v)
}
func (mpm *MeshPhysicalMaterial) Transparent() bool {
	return mpm.Get("transparent").Bool()
}
func (mpm *MeshPhysicalMaterial) SetTransparent(v bool) {
	mpm.Set("transparent", v)
}
func (mpm *MeshPhysicalMaterial) Type() string {
	return mpm.Get("type").String()
}
func (mpm *MeshPhysicalMaterial) SetType(v string) {
	mpm.Set("type", v)
}
func (mpm *MeshPhysicalMaterial) UserData() js.Value {
	return mpm.Get("userData")
}
func (mpm *MeshPhysicalMaterial) SetUserData(v js.Value) {
	mpm.Set("userData", v)
}
func (mpm *MeshPhysicalMaterial) Uuid() string {
	return mpm.Get("uuid").String()
}
func (mpm *MeshPhysicalMaterial) SetUuid(v string) {
	mpm.Set("uuid", v)
}
func (mpm *MeshPhysicalMaterial) VertexColors() *Colors {
	return &Colors{Value: mpm.Get("vertexColors")}
}
func (mpm *MeshPhysicalMaterial) SetVertexColors(v *Colors) {
	mpm.Set("vertexColors", v)
}
func (mpm *MeshPhysicalMaterial) VertexTangents() bool {
	return mpm.Get("vertexTangents").Bool()
}
func (mpm *MeshPhysicalMaterial) SetVertexTangents(v bool) {
	mpm.Set("vertexTangents", v)
}
func (mpm *MeshPhysicalMaterial) Visible() bool {
	return mpm.Get("visible").Bool()
}
func (mpm *MeshPhysicalMaterial) SetVisible(v bool) {
	mpm.Set("visible", v)
}
func (mpm *MeshPhysicalMaterial) Wireframe() bool {
	return mpm.Get("wireframe").Bool()
}
func (mpm *MeshPhysicalMaterial) SetWireframe(v bool) {
	mpm.Set("wireframe", v)
}
func (mpm *MeshPhysicalMaterial) WireframeLinewidth() int {
	return mpm.Get("wireframeLinewidth").Int()
}
func (mpm *MeshPhysicalMaterial) SetWireframeLinewidth(v int) {
	mpm.Set("wireframeLinewidth", v)
}
func (mpm *MeshPhysicalMaterial) AddEventListener(typ string, listener js.Value) {
	mpm.Call("addEventListener", typ, listener)
}
func (mpm *MeshPhysicalMaterial) Clone() this {
	return this(mpm.Call("clone"))
}
func (mpm *MeshPhysicalMaterial) Copy(material *Material) this {
	return this(mpm.Call("copy", material))
}
func (mpm *MeshPhysicalMaterial) DispatchEvent(event js.Value) {
	mpm.Call("dispatchEvent", event)
}
func (mpm *MeshPhysicalMaterial) Dispose() {
	mpm.Call("dispose")
}
func (mpm *MeshPhysicalMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mpm.Call("hasEventListener", typ, listener).Bool()
}
func (mpm *MeshPhysicalMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	mpm.Call("onBeforeCompile", shader, renderer)
}
func (mpm *MeshPhysicalMaterial) RemoveEventListener(typ string, listener js.Value) {
	mpm.Call("removeEventListener", typ, listener)
}
func (mpm *MeshPhysicalMaterial) SetValues(parameters MeshStandardMaterialParameters) {
	mpm.Call("setValues", parameters)
}
func (mpm *MeshPhysicalMaterial) ToJSON(meta js.Value) js.Value {
	return mpm.Call("toJSON", meta)
}
func (mpm *MeshPhysicalMaterial) Update() {
	mpm.Call("update")
}

package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
	"github.com/nobonobo/three/textures"
)

type MeshPhongMaterialParameters interface {
}
type MeshPhongMaterial struct {
	js.Value
}

func NewMeshPhongMaterial(parameters MeshPhongMaterialParameters) *MeshPhongMaterial {
	return &MeshPhongMaterial{Value: get("MeshPhongMaterial").New(parameters)}
}
func (mpm *MeshPhongMaterial) AlphaMap() textures.Texture {
	return textures.Texture(mpm.Get("alphaMap"))
}
func (mpm *MeshPhongMaterial) SetAlphaMap(v textures.Texture) {
	mpm.Set("alphaMap", v)
}
func (mpm *MeshPhongMaterial) AlphaTest() int {
	return mpm.Get("alphaTest").Int()
}
func (mpm *MeshPhongMaterial) SetAlphaTest(v int) {
	mpm.Set("alphaTest", v)
}
func (mpm *MeshPhongMaterial) AoMap() textures.Texture {
	return textures.Texture(mpm.Get("aoMap"))
}
func (mpm *MeshPhongMaterial) SetAoMap(v textures.Texture) {
	mpm.Set("aoMap", v)
}
func (mpm *MeshPhongMaterial) AoMapIntensity() int {
	return mpm.Get("aoMapIntensity").Int()
}
func (mpm *MeshPhongMaterial) SetAoMapIntensity(v int) {
	mpm.Set("aoMapIntensity", v)
}
func (mpm *MeshPhongMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: mpm.Get("blendDst")}
}
func (mpm *MeshPhongMaterial) SetBlendDst(v *BlendingDstFactor) {
	mpm.Set("blendDst", v)
}
func (mpm *MeshPhongMaterial) BlendDstAlpha() int {
	return mpm.Get("blendDstAlpha").Int()
}
func (mpm *MeshPhongMaterial) SetBlendDstAlpha(v int) {
	mpm.Set("blendDstAlpha", v)
}
func (mpm *MeshPhongMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: mpm.Get("blendEquation")}
}
func (mpm *MeshPhongMaterial) SetBlendEquation(v *BlendingEquation) {
	mpm.Set("blendEquation", v)
}
func (mpm *MeshPhongMaterial) BlendEquationAlpha() int {
	return mpm.Get("blendEquationAlpha").Int()
}
func (mpm *MeshPhongMaterial) SetBlendEquationAlpha(v int) {
	mpm.Set("blendEquationAlpha", v)
}
func (mpm *MeshPhongMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(mpm.Get("blendSrc"))
}
func (mpm *MeshPhongMaterial) SetBlendSrc(v BlendingSrcFactor) {
	mpm.Set("blendSrc", v)
}
func (mpm *MeshPhongMaterial) BlendSrcAlpha() int {
	return mpm.Get("blendSrcAlpha").Int()
}
func (mpm *MeshPhongMaterial) SetBlendSrcAlpha(v int) {
	mpm.Set("blendSrcAlpha", v)
}
func (mpm *MeshPhongMaterial) Blending() *Blending {
	return &Blending{Value: mpm.Get("blending")}
}
func (mpm *MeshPhongMaterial) SetBlending(v *Blending) {
	mpm.Set("blending", v)
}
func (mpm *MeshPhongMaterial) BumpMap() textures.Texture {
	return textures.Texture(mpm.Get("bumpMap"))
}
func (mpm *MeshPhongMaterial) SetBumpMap(v textures.Texture) {
	mpm.Set("bumpMap", v)
}
func (mpm *MeshPhongMaterial) BumpScale() int {
	return mpm.Get("bumpScale").Int()
}
func (mpm *MeshPhongMaterial) SetBumpScale(v int) {
	mpm.Set("bumpScale", v)
}
func (mpm *MeshPhongMaterial) ClipIntersection() bool {
	return mpm.Get("clipIntersection").Bool()
}
func (mpm *MeshPhongMaterial) SetClipIntersection(v bool) {
	mpm.Set("clipIntersection", v)
}
func (mpm *MeshPhongMaterial) ClipShadows() bool {
	return mpm.Get("clipShadows").Bool()
}
func (mpm *MeshPhongMaterial) SetClipShadows(v bool) {
	mpm.Set("clipShadows", v)
}
func (mpm *MeshPhongMaterial) ClippingPlanes() js.Value {
	return mpm.Get("clippingPlanes")
}
func (mpm *MeshPhongMaterial) SetClippingPlanes(v js.Value) {
	mpm.Set("clippingPlanes", v)
}
func (mpm *MeshPhongMaterial) Color() *math.Color {
	return &math.Color{Value: mpm.Get("color")}
}
func (mpm *MeshPhongMaterial) SetColor(v *math.Color) {
	mpm.Set("color", v)
}
func (mpm *MeshPhongMaterial) ColorWrite() bool {
	return mpm.Get("colorWrite").Bool()
}
func (mpm *MeshPhongMaterial) SetColorWrite(v bool) {
	mpm.Set("colorWrite", v)
}
func (mpm *MeshPhongMaterial) Combine() *Combine {
	return &Combine{Value: mpm.Get("combine")}
}
func (mpm *MeshPhongMaterial) SetCombine(v *Combine) {
	mpm.Set("combine", v)
}
func (mpm *MeshPhongMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: mpm.Get("depthFunc")}
}
func (mpm *MeshPhongMaterial) SetDepthFunc(v *DepthModes) {
	mpm.Set("depthFunc", v)
}
func (mpm *MeshPhongMaterial) DepthTest() bool {
	return mpm.Get("depthTest").Bool()
}
func (mpm *MeshPhongMaterial) SetDepthTest(v bool) {
	mpm.Set("depthTest", v)
}
func (mpm *MeshPhongMaterial) DepthWrite() bool {
	return mpm.Get("depthWrite").Bool()
}
func (mpm *MeshPhongMaterial) SetDepthWrite(v bool) {
	mpm.Set("depthWrite", v)
}
func (mpm *MeshPhongMaterial) DisplacementBias() int {
	return mpm.Get("displacementBias").Int()
}
func (mpm *MeshPhongMaterial) SetDisplacementBias(v int) {
	mpm.Set("displacementBias", v)
}
func (mpm *MeshPhongMaterial) DisplacementMap() textures.Texture {
	return textures.Texture(mpm.Get("displacementMap"))
}
func (mpm *MeshPhongMaterial) SetDisplacementMap(v textures.Texture) {
	mpm.Set("displacementMap", v)
}
func (mpm *MeshPhongMaterial) DisplacementScale() int {
	return mpm.Get("displacementScale").Int()
}
func (mpm *MeshPhongMaterial) SetDisplacementScale(v int) {
	mpm.Set("displacementScale", v)
}
func (mpm *MeshPhongMaterial) Dithering() bool {
	return mpm.Get("dithering").Bool()
}
func (mpm *MeshPhongMaterial) SetDithering(v bool) {
	mpm.Set("dithering", v)
}
func (mpm *MeshPhongMaterial) Emissive() *math.Color {
	return &math.Color{Value: mpm.Get("emissive")}
}
func (mpm *MeshPhongMaterial) SetEmissive(v *math.Color) {
	mpm.Set("emissive", v)
}
func (mpm *MeshPhongMaterial) EmissiveIntensity() int {
	return mpm.Get("emissiveIntensity").Int()
}
func (mpm *MeshPhongMaterial) SetEmissiveIntensity(v int) {
	mpm.Set("emissiveIntensity", v)
}
func (mpm *MeshPhongMaterial) EmissiveMap() textures.Texture {
	return textures.Texture(mpm.Get("emissiveMap"))
}
func (mpm *MeshPhongMaterial) SetEmissiveMap(v textures.Texture) {
	mpm.Set("emissiveMap", v)
}
func (mpm *MeshPhongMaterial) EnvMap() textures.Texture {
	return textures.Texture(mpm.Get("envMap"))
}
func (mpm *MeshPhongMaterial) SetEnvMap(v textures.Texture) {
	mpm.Set("envMap", v)
}
func (mpm *MeshPhongMaterial) FlatShading() bool {
	return mpm.Get("flatShading").Bool()
}
func (mpm *MeshPhongMaterial) SetFlatShading(v bool) {
	mpm.Set("flatShading", v)
}
func (mpm *MeshPhongMaterial) Fog() bool {
	return mpm.Get("fog").Bool()
}
func (mpm *MeshPhongMaterial) SetFog(v bool) {
	mpm.Set("fog", v)
}
func (mpm *MeshPhongMaterial) Id() int {
	return mpm.Get("id").Int()
}
func (mpm *MeshPhongMaterial) SetId(v int) {
	mpm.Set("id", v)
}
func (mpm *MeshPhongMaterial) IsMaterial() bool {
	return mpm.Get("isMaterial").Bool()
}
func (mpm *MeshPhongMaterial) SetIsMaterial(v bool) {
	mpm.Set("isMaterial", v)
}
func (mpm *MeshPhongMaterial) LightMap() textures.Texture {
	return textures.Texture(mpm.Get("lightMap"))
}
func (mpm *MeshPhongMaterial) SetLightMap(v textures.Texture) {
	mpm.Set("lightMap", v)
}
func (mpm *MeshPhongMaterial) LightMapIntensity() int {
	return mpm.Get("lightMapIntensity").Int()
}
func (mpm *MeshPhongMaterial) SetLightMapIntensity(v int) {
	mpm.Set("lightMapIntensity", v)
}
func (mpm *MeshPhongMaterial) Lights() bool {
	return mpm.Get("lights").Bool()
}
func (mpm *MeshPhongMaterial) SetLights(v bool) {
	mpm.Set("lights", v)
}
func (mpm *MeshPhongMaterial) Map() textures.Texture {
	return textures.Texture(mpm.Get("map"))
}
func (mpm *MeshPhongMaterial) SetMap(v textures.Texture) {
	mpm.Set("map", v)
}
func (mpm *MeshPhongMaterial) Metal() bool {
	return mpm.Get("metal").Bool()
}
func (mpm *MeshPhongMaterial) SetMetal(v bool) {
	mpm.Set("metal", v)
}
func (mpm *MeshPhongMaterial) MorphNormals() bool {
	return mpm.Get("morphNormals").Bool()
}
func (mpm *MeshPhongMaterial) SetMorphNormals(v bool) {
	mpm.Set("morphNormals", v)
}
func (mpm *MeshPhongMaterial) MorphTargets() bool {
	return mpm.Get("morphTargets").Bool()
}
func (mpm *MeshPhongMaterial) SetMorphTargets(v bool) {
	mpm.Set("morphTargets", v)
}
func (mpm *MeshPhongMaterial) Name() string {
	return mpm.Get("name").String()
}
func (mpm *MeshPhongMaterial) SetName(v string) {
	mpm.Set("name", v)
}
func (mpm *MeshPhongMaterial) NeedsUpdate() bool {
	return mpm.Get("needsUpdate").Bool()
}
func (mpm *MeshPhongMaterial) SetNeedsUpdate(v bool) {
	mpm.Set("needsUpdate", v)
}
func (mpm *MeshPhongMaterial) NormalMap() textures.Texture {
	return textures.Texture(mpm.Get("normalMap"))
}
func (mpm *MeshPhongMaterial) SetNormalMap(v textures.Texture) {
	mpm.Set("normalMap", v)
}
func (mpm *MeshPhongMaterial) NormalScale() *math.Vector2 {
	return &math.Vector2{Value: mpm.Get("normalScale")}
}
func (mpm *MeshPhongMaterial) SetNormalScale(v *math.Vector2) {
	mpm.Set("normalScale", v)
}
func (mpm *MeshPhongMaterial) Opacity() int {
	return mpm.Get("opacity").Int()
}
func (mpm *MeshPhongMaterial) SetOpacity(v int) {
	mpm.Set("opacity", v)
}
func (mpm *MeshPhongMaterial) Overdraw() int {
	return mpm.Get("overdraw").Int()
}
func (mpm *MeshPhongMaterial) SetOverdraw(v int) {
	mpm.Set("overdraw", v)
}
func (mpm *MeshPhongMaterial) PolygonOffset() bool {
	return mpm.Get("polygonOffset").Bool()
}
func (mpm *MeshPhongMaterial) SetPolygonOffset(v bool) {
	mpm.Set("polygonOffset", v)
}
func (mpm *MeshPhongMaterial) PolygonOffsetFactor() int {
	return mpm.Get("polygonOffsetFactor").Int()
}
func (mpm *MeshPhongMaterial) SetPolygonOffsetFactor(v int) {
	mpm.Set("polygonOffsetFactor", v)
}
func (mpm *MeshPhongMaterial) PolygonOffsetUnits() int {
	return mpm.Get("polygonOffsetUnits").Int()
}
func (mpm *MeshPhongMaterial) SetPolygonOffsetUnits(v int) {
	mpm.Set("polygonOffsetUnits", v)
}
func (mpm *MeshPhongMaterial) Precision() string {
	return mpm.Get("precision").String()
}
func (mpm *MeshPhongMaterial) SetPrecision(v string) {
	mpm.Set("precision", v)
}
func (mpm *MeshPhongMaterial) PremultipliedAlpha() bool {
	return mpm.Get("premultipliedAlpha").Bool()
}
func (mpm *MeshPhongMaterial) SetPremultipliedAlpha(v bool) {
	mpm.Set("premultipliedAlpha", v)
}
func (mpm *MeshPhongMaterial) Reflectivity() int {
	return mpm.Get("reflectivity").Int()
}
func (mpm *MeshPhongMaterial) SetReflectivity(v int) {
	mpm.Set("reflectivity", v)
}
func (mpm *MeshPhongMaterial) RefractionRatio() int {
	return mpm.Get("refractionRatio").Int()
}
func (mpm *MeshPhongMaterial) SetRefractionRatio(v int) {
	mpm.Set("refractionRatio", v)
}
func (mpm *MeshPhongMaterial) Shininess() int {
	return mpm.Get("shininess").Int()
}
func (mpm *MeshPhongMaterial) SetShininess(v int) {
	mpm.Set("shininess", v)
}
func (mpm *MeshPhongMaterial) Side() *Side {
	return &Side{Value: mpm.Get("side")}
}
func (mpm *MeshPhongMaterial) SetSide(v *Side) {
	mpm.Set("side", v)
}
func (mpm *MeshPhongMaterial) Skinning() bool {
	return mpm.Get("skinning").Bool()
}
func (mpm *MeshPhongMaterial) SetSkinning(v bool) {
	mpm.Set("skinning", v)
}
func (mpm *MeshPhongMaterial) Specular() *math.Color {
	return &math.Color{Value: mpm.Get("specular")}
}
func (mpm *MeshPhongMaterial) SetSpecular(v *math.Color) {
	mpm.Set("specular", v)
}
func (mpm *MeshPhongMaterial) SpecularMap() textures.Texture {
	return textures.Texture(mpm.Get("specularMap"))
}
func (mpm *MeshPhongMaterial) SetSpecularMap(v textures.Texture) {
	mpm.Set("specularMap", v)
}
func (mpm *MeshPhongMaterial) Transparent() bool {
	return mpm.Get("transparent").Bool()
}
func (mpm *MeshPhongMaterial) SetTransparent(v bool) {
	mpm.Set("transparent", v)
}
func (mpm *MeshPhongMaterial) Type() string {
	return mpm.Get("type").String()
}
func (mpm *MeshPhongMaterial) SetType(v string) {
	mpm.Set("type", v)
}
func (mpm *MeshPhongMaterial) UserData() js.Value {
	return mpm.Get("userData")
}
func (mpm *MeshPhongMaterial) SetUserData(v js.Value) {
	mpm.Set("userData", v)
}
func (mpm *MeshPhongMaterial) Uuid() string {
	return mpm.Get("uuid").String()
}
func (mpm *MeshPhongMaterial) SetUuid(v string) {
	mpm.Set("uuid", v)
}
func (mpm *MeshPhongMaterial) VertexColors() *Colors {
	return &Colors{Value: mpm.Get("vertexColors")}
}
func (mpm *MeshPhongMaterial) SetVertexColors(v *Colors) {
	mpm.Set("vertexColors", v)
}
func (mpm *MeshPhongMaterial) VertexTangents() bool {
	return mpm.Get("vertexTangents").Bool()
}
func (mpm *MeshPhongMaterial) SetVertexTangents(v bool) {
	mpm.Set("vertexTangents", v)
}
func (mpm *MeshPhongMaterial) Visible() bool {
	return mpm.Get("visible").Bool()
}
func (mpm *MeshPhongMaterial) SetVisible(v bool) {
	mpm.Set("visible", v)
}
func (mpm *MeshPhongMaterial) Wireframe() bool {
	return mpm.Get("wireframe").Bool()
}
func (mpm *MeshPhongMaterial) SetWireframe(v bool) {
	mpm.Set("wireframe", v)
}
func (mpm *MeshPhongMaterial) WireframeLinecap() string {
	return mpm.Get("wireframeLinecap").String()
}
func (mpm *MeshPhongMaterial) SetWireframeLinecap(v string) {
	mpm.Set("wireframeLinecap", v)
}
func (mpm *MeshPhongMaterial) WireframeLinejoin() string {
	return mpm.Get("wireframeLinejoin").String()
}
func (mpm *MeshPhongMaterial) SetWireframeLinejoin(v string) {
	mpm.Set("wireframeLinejoin", v)
}
func (mpm *MeshPhongMaterial) WireframeLinewidth() int {
	return mpm.Get("wireframeLinewidth").Int()
}
func (mpm *MeshPhongMaterial) SetWireframeLinewidth(v int) {
	mpm.Set("wireframeLinewidth", v)
}
func (mpm *MeshPhongMaterial) AddEventListener(typ string, listener js.Value) {
	mpm.Call("addEventListener", typ, listener)
}
func (mpm *MeshPhongMaterial) Clone() this {
	return this(mpm.Call("clone"))
}
func (mpm *MeshPhongMaterial) Copy(material *Material) this {
	return this(mpm.Call("copy", material))
}
func (mpm *MeshPhongMaterial) DispatchEvent(event js.Value) {
	mpm.Call("dispatchEvent", event)
}
func (mpm *MeshPhongMaterial) Dispose() {
	mpm.Call("dispose")
}
func (mpm *MeshPhongMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mpm.Call("hasEventListener", typ, listener).Bool()
}
func (mpm *MeshPhongMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	mpm.Call("onBeforeCompile", shader, renderer)
}
func (mpm *MeshPhongMaterial) RemoveEventListener(typ string, listener js.Value) {
	mpm.Call("removeEventListener", typ, listener)
}
func (mpm *MeshPhongMaterial) SetValues(parameters MeshPhongMaterialParameters) {
	mpm.Call("setValues", parameters)
}
func (mpm *MeshPhongMaterial) ToJSON(meta js.Value) js.Value {
	return mpm.Call("toJSON", meta)
}
func (mpm *MeshPhongMaterial) Update() {
	mpm.Call("update")
}

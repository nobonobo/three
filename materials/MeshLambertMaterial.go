package materials

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/shaders"
	"github.com/nobonobo/three/textures"
)

type MeshLambertMaterialParameters interface {
}
type MeshLambertMaterial struct {
	js.Value
}

func NewMeshLambertMaterial(parameters MeshLambertMaterialParameters) *MeshLambertMaterial {
	return &MeshLambertMaterial{Value: get("MeshLambertMaterial").New(parameters)}
}
func (mlm *MeshLambertMaterial) AlphaMap() textures.Texture {
	return textures.Texture(mlm.Get("alphaMap"))
}
func (mlm *MeshLambertMaterial) SetAlphaMap(v textures.Texture) {
	mlm.Set("alphaMap", v)
}
func (mlm *MeshLambertMaterial) AlphaTest() int {
	return mlm.Get("alphaTest").Int()
}
func (mlm *MeshLambertMaterial) SetAlphaTest(v int) {
	mlm.Set("alphaTest", v)
}
func (mlm *MeshLambertMaterial) AoMap() textures.Texture {
	return textures.Texture(mlm.Get("aoMap"))
}
func (mlm *MeshLambertMaterial) SetAoMap(v textures.Texture) {
	mlm.Set("aoMap", v)
}
func (mlm *MeshLambertMaterial) AoMapIntensity() int {
	return mlm.Get("aoMapIntensity").Int()
}
func (mlm *MeshLambertMaterial) SetAoMapIntensity(v int) {
	mlm.Set("aoMapIntensity", v)
}
func (mlm *MeshLambertMaterial) BlendDst() *BlendingDstFactor {
	return &BlendingDstFactor{Value: mlm.Get("blendDst")}
}
func (mlm *MeshLambertMaterial) SetBlendDst(v *BlendingDstFactor) {
	mlm.Set("blendDst", v)
}
func (mlm *MeshLambertMaterial) BlendDstAlpha() int {
	return mlm.Get("blendDstAlpha").Int()
}
func (mlm *MeshLambertMaterial) SetBlendDstAlpha(v int) {
	mlm.Set("blendDstAlpha", v)
}
func (mlm *MeshLambertMaterial) BlendEquation() *BlendingEquation {
	return &BlendingEquation{Value: mlm.Get("blendEquation")}
}
func (mlm *MeshLambertMaterial) SetBlendEquation(v *BlendingEquation) {
	mlm.Set("blendEquation", v)
}
func (mlm *MeshLambertMaterial) BlendEquationAlpha() int {
	return mlm.Get("blendEquationAlpha").Int()
}
func (mlm *MeshLambertMaterial) SetBlendEquationAlpha(v int) {
	mlm.Set("blendEquationAlpha", v)
}
func (mlm *MeshLambertMaterial) BlendSrc() BlendingSrcFactor {
	return BlendingSrcFactor(mlm.Get("blendSrc"))
}
func (mlm *MeshLambertMaterial) SetBlendSrc(v BlendingSrcFactor) {
	mlm.Set("blendSrc", v)
}
func (mlm *MeshLambertMaterial) BlendSrcAlpha() int {
	return mlm.Get("blendSrcAlpha").Int()
}
func (mlm *MeshLambertMaterial) SetBlendSrcAlpha(v int) {
	mlm.Set("blendSrcAlpha", v)
}
func (mlm *MeshLambertMaterial) Blending() *Blending {
	return &Blending{Value: mlm.Get("blending")}
}
func (mlm *MeshLambertMaterial) SetBlending(v *Blending) {
	mlm.Set("blending", v)
}
func (mlm *MeshLambertMaterial) ClipIntersection() bool {
	return mlm.Get("clipIntersection").Bool()
}
func (mlm *MeshLambertMaterial) SetClipIntersection(v bool) {
	mlm.Set("clipIntersection", v)
}
func (mlm *MeshLambertMaterial) ClipShadows() bool {
	return mlm.Get("clipShadows").Bool()
}
func (mlm *MeshLambertMaterial) SetClipShadows(v bool) {
	mlm.Set("clipShadows", v)
}
func (mlm *MeshLambertMaterial) ClippingPlanes() js.Value {
	return mlm.Get("clippingPlanes")
}
func (mlm *MeshLambertMaterial) SetClippingPlanes(v js.Value) {
	mlm.Set("clippingPlanes", v)
}
func (mlm *MeshLambertMaterial) Color() *math.Color {
	return &math.Color{Value: mlm.Get("color")}
}
func (mlm *MeshLambertMaterial) SetColor(v *math.Color) {
	mlm.Set("color", v)
}
func (mlm *MeshLambertMaterial) ColorWrite() bool {
	return mlm.Get("colorWrite").Bool()
}
func (mlm *MeshLambertMaterial) SetColorWrite(v bool) {
	mlm.Set("colorWrite", v)
}
func (mlm *MeshLambertMaterial) Combine() *Combine {
	return &Combine{Value: mlm.Get("combine")}
}
func (mlm *MeshLambertMaterial) SetCombine(v *Combine) {
	mlm.Set("combine", v)
}
func (mlm *MeshLambertMaterial) DepthFunc() *DepthModes {
	return &DepthModes{Value: mlm.Get("depthFunc")}
}
func (mlm *MeshLambertMaterial) SetDepthFunc(v *DepthModes) {
	mlm.Set("depthFunc", v)
}
func (mlm *MeshLambertMaterial) DepthTest() bool {
	return mlm.Get("depthTest").Bool()
}
func (mlm *MeshLambertMaterial) SetDepthTest(v bool) {
	mlm.Set("depthTest", v)
}
func (mlm *MeshLambertMaterial) DepthWrite() bool {
	return mlm.Get("depthWrite").Bool()
}
func (mlm *MeshLambertMaterial) SetDepthWrite(v bool) {
	mlm.Set("depthWrite", v)
}
func (mlm *MeshLambertMaterial) Dithering() bool {
	return mlm.Get("dithering").Bool()
}
func (mlm *MeshLambertMaterial) SetDithering(v bool) {
	mlm.Set("dithering", v)
}
func (mlm *MeshLambertMaterial) Emissive() *math.Color {
	return &math.Color{Value: mlm.Get("emissive")}
}
func (mlm *MeshLambertMaterial) SetEmissive(v *math.Color) {
	mlm.Set("emissive", v)
}
func (mlm *MeshLambertMaterial) EmissiveIntensity() int {
	return mlm.Get("emissiveIntensity").Int()
}
func (mlm *MeshLambertMaterial) SetEmissiveIntensity(v int) {
	mlm.Set("emissiveIntensity", v)
}
func (mlm *MeshLambertMaterial) EmissiveMap() textures.Texture {
	return textures.Texture(mlm.Get("emissiveMap"))
}
func (mlm *MeshLambertMaterial) SetEmissiveMap(v textures.Texture) {
	mlm.Set("emissiveMap", v)
}
func (mlm *MeshLambertMaterial) EnvMap() textures.Texture {
	return textures.Texture(mlm.Get("envMap"))
}
func (mlm *MeshLambertMaterial) SetEnvMap(v textures.Texture) {
	mlm.Set("envMap", v)
}
func (mlm *MeshLambertMaterial) FlatShading() bool {
	return mlm.Get("flatShading").Bool()
}
func (mlm *MeshLambertMaterial) SetFlatShading(v bool) {
	mlm.Set("flatShading", v)
}
func (mlm *MeshLambertMaterial) Fog() bool {
	return mlm.Get("fog").Bool()
}
func (mlm *MeshLambertMaterial) SetFog(v bool) {
	mlm.Set("fog", v)
}
func (mlm *MeshLambertMaterial) Id() int {
	return mlm.Get("id").Int()
}
func (mlm *MeshLambertMaterial) SetId(v int) {
	mlm.Set("id", v)
}
func (mlm *MeshLambertMaterial) IsMaterial() bool {
	return mlm.Get("isMaterial").Bool()
}
func (mlm *MeshLambertMaterial) SetIsMaterial(v bool) {
	mlm.Set("isMaterial", v)
}
func (mlm *MeshLambertMaterial) LightMap() textures.Texture {
	return textures.Texture(mlm.Get("lightMap"))
}
func (mlm *MeshLambertMaterial) SetLightMap(v textures.Texture) {
	mlm.Set("lightMap", v)
}
func (mlm *MeshLambertMaterial) LightMapIntensity() int {
	return mlm.Get("lightMapIntensity").Int()
}
func (mlm *MeshLambertMaterial) SetLightMapIntensity(v int) {
	mlm.Set("lightMapIntensity", v)
}
func (mlm *MeshLambertMaterial) Lights() bool {
	return mlm.Get("lights").Bool()
}
func (mlm *MeshLambertMaterial) SetLights(v bool) {
	mlm.Set("lights", v)
}
func (mlm *MeshLambertMaterial) Map() textures.Texture {
	return textures.Texture(mlm.Get("map"))
}
func (mlm *MeshLambertMaterial) SetMap(v textures.Texture) {
	mlm.Set("map", v)
}
func (mlm *MeshLambertMaterial) MorphNormals() bool {
	return mlm.Get("morphNormals").Bool()
}
func (mlm *MeshLambertMaterial) SetMorphNormals(v bool) {
	mlm.Set("morphNormals", v)
}
func (mlm *MeshLambertMaterial) MorphTargets() bool {
	return mlm.Get("morphTargets").Bool()
}
func (mlm *MeshLambertMaterial) SetMorphTargets(v bool) {
	mlm.Set("morphTargets", v)
}
func (mlm *MeshLambertMaterial) Name() string {
	return mlm.Get("name").String()
}
func (mlm *MeshLambertMaterial) SetName(v string) {
	mlm.Set("name", v)
}
func (mlm *MeshLambertMaterial) NeedsUpdate() bool {
	return mlm.Get("needsUpdate").Bool()
}
func (mlm *MeshLambertMaterial) SetNeedsUpdate(v bool) {
	mlm.Set("needsUpdate", v)
}
func (mlm *MeshLambertMaterial) Opacity() int {
	return mlm.Get("opacity").Int()
}
func (mlm *MeshLambertMaterial) SetOpacity(v int) {
	mlm.Set("opacity", v)
}
func (mlm *MeshLambertMaterial) Overdraw() int {
	return mlm.Get("overdraw").Int()
}
func (mlm *MeshLambertMaterial) SetOverdraw(v int) {
	mlm.Set("overdraw", v)
}
func (mlm *MeshLambertMaterial) PolygonOffset() bool {
	return mlm.Get("polygonOffset").Bool()
}
func (mlm *MeshLambertMaterial) SetPolygonOffset(v bool) {
	mlm.Set("polygonOffset", v)
}
func (mlm *MeshLambertMaterial) PolygonOffsetFactor() int {
	return mlm.Get("polygonOffsetFactor").Int()
}
func (mlm *MeshLambertMaterial) SetPolygonOffsetFactor(v int) {
	mlm.Set("polygonOffsetFactor", v)
}
func (mlm *MeshLambertMaterial) PolygonOffsetUnits() int {
	return mlm.Get("polygonOffsetUnits").Int()
}
func (mlm *MeshLambertMaterial) SetPolygonOffsetUnits(v int) {
	mlm.Set("polygonOffsetUnits", v)
}
func (mlm *MeshLambertMaterial) Precision() string {
	return mlm.Get("precision").String()
}
func (mlm *MeshLambertMaterial) SetPrecision(v string) {
	mlm.Set("precision", v)
}
func (mlm *MeshLambertMaterial) PremultipliedAlpha() bool {
	return mlm.Get("premultipliedAlpha").Bool()
}
func (mlm *MeshLambertMaterial) SetPremultipliedAlpha(v bool) {
	mlm.Set("premultipliedAlpha", v)
}
func (mlm *MeshLambertMaterial) Reflectivity() int {
	return mlm.Get("reflectivity").Int()
}
func (mlm *MeshLambertMaterial) SetReflectivity(v int) {
	mlm.Set("reflectivity", v)
}
func (mlm *MeshLambertMaterial) RefractionRatio() int {
	return mlm.Get("refractionRatio").Int()
}
func (mlm *MeshLambertMaterial) SetRefractionRatio(v int) {
	mlm.Set("refractionRatio", v)
}
func (mlm *MeshLambertMaterial) Side() *Side {
	return &Side{Value: mlm.Get("side")}
}
func (mlm *MeshLambertMaterial) SetSide(v *Side) {
	mlm.Set("side", v)
}
func (mlm *MeshLambertMaterial) Skinning() bool {
	return mlm.Get("skinning").Bool()
}
func (mlm *MeshLambertMaterial) SetSkinning(v bool) {
	mlm.Set("skinning", v)
}
func (mlm *MeshLambertMaterial) SpecularMap() textures.Texture {
	return textures.Texture(mlm.Get("specularMap"))
}
func (mlm *MeshLambertMaterial) SetSpecularMap(v textures.Texture) {
	mlm.Set("specularMap", v)
}
func (mlm *MeshLambertMaterial) Transparent() bool {
	return mlm.Get("transparent").Bool()
}
func (mlm *MeshLambertMaterial) SetTransparent(v bool) {
	mlm.Set("transparent", v)
}
func (mlm *MeshLambertMaterial) Type() string {
	return mlm.Get("type").String()
}
func (mlm *MeshLambertMaterial) SetType(v string) {
	mlm.Set("type", v)
}
func (mlm *MeshLambertMaterial) UserData() js.Value {
	return mlm.Get("userData")
}
func (mlm *MeshLambertMaterial) SetUserData(v js.Value) {
	mlm.Set("userData", v)
}
func (mlm *MeshLambertMaterial) Uuid() string {
	return mlm.Get("uuid").String()
}
func (mlm *MeshLambertMaterial) SetUuid(v string) {
	mlm.Set("uuid", v)
}
func (mlm *MeshLambertMaterial) VertexColors() *Colors {
	return &Colors{Value: mlm.Get("vertexColors")}
}
func (mlm *MeshLambertMaterial) SetVertexColors(v *Colors) {
	mlm.Set("vertexColors", v)
}
func (mlm *MeshLambertMaterial) VertexTangents() bool {
	return mlm.Get("vertexTangents").Bool()
}
func (mlm *MeshLambertMaterial) SetVertexTangents(v bool) {
	mlm.Set("vertexTangents", v)
}
func (mlm *MeshLambertMaterial) Visible() bool {
	return mlm.Get("visible").Bool()
}
func (mlm *MeshLambertMaterial) SetVisible(v bool) {
	mlm.Set("visible", v)
}
func (mlm *MeshLambertMaterial) Wireframe() bool {
	return mlm.Get("wireframe").Bool()
}
func (mlm *MeshLambertMaterial) SetWireframe(v bool) {
	mlm.Set("wireframe", v)
}
func (mlm *MeshLambertMaterial) WireframeLinecap() string {
	return mlm.Get("wireframeLinecap").String()
}
func (mlm *MeshLambertMaterial) SetWireframeLinecap(v string) {
	mlm.Set("wireframeLinecap", v)
}
func (mlm *MeshLambertMaterial) WireframeLinejoin() string {
	return mlm.Get("wireframeLinejoin").String()
}
func (mlm *MeshLambertMaterial) SetWireframeLinejoin(v string) {
	mlm.Set("wireframeLinejoin", v)
}
func (mlm *MeshLambertMaterial) WireframeLinewidth() int {
	return mlm.Get("wireframeLinewidth").Int()
}
func (mlm *MeshLambertMaterial) SetWireframeLinewidth(v int) {
	mlm.Set("wireframeLinewidth", v)
}
func (mlm *MeshLambertMaterial) AddEventListener(typ string, listener js.Value) {
	mlm.Call("addEventListener", typ, listener)
}
func (mlm *MeshLambertMaterial) Clone() this {
	return this(mlm.Call("clone"))
}
func (mlm *MeshLambertMaterial) Copy(material *Material) this {
	return this(mlm.Call("copy", material))
}
func (mlm *MeshLambertMaterial) DispatchEvent(event js.Value) {
	mlm.Call("dispatchEvent", event)
}
func (mlm *MeshLambertMaterial) Dispose() {
	mlm.Call("dispose")
}
func (mlm *MeshLambertMaterial) HasEventListener(typ string, listener js.Value) bool {
	return mlm.Call("hasEventListener", typ, listener).Bool()
}
func (mlm *MeshLambertMaterial) OnBeforeCompile(shader shaders.Shader, renderer *renderers.WebGLRenderer) {
	mlm.Call("onBeforeCompile", shader, renderer)
}
func (mlm *MeshLambertMaterial) RemoveEventListener(typ string, listener js.Value) {
	mlm.Call("removeEventListener", typ, listener)
}
func (mlm *MeshLambertMaterial) SetValues(parameters MeshLambertMaterialParameters) {
	mlm.Call("setValues", parameters)
}
func (mlm *MeshLambertMaterial) ToJSON(meta js.Value) js.Value {
	return mlm.Call("toJSON", meta)
}
func (mlm *MeshLambertMaterial) Update() {
	mlm.Call("update")
}

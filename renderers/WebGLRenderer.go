package renderers

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers/webgl"
	"github.com/nobonobo/three/renderers/webvr"
	"github.com/nobonobo/three/scenes"
	"github.com/nobonobo/three/textures"
)

type Renderer interface {
	Render(scene *scenes.Scene, camera *cameras.Camera)
	SetSize(width int, height int, updateStyle bool)
}
type WebGLRendererParameters interface {
}
type WebGLRenderer struct {
	js.Value
}

func NewWebGLRenderer(parameters WebGLRendererParameters) *WebGLRenderer {
	return &WebGLRenderer{Value: get("WebGLRenderer").New(parameters)}
}
func (wglr *WebGLRenderer) AllocTextureUnit() js.Value {
	return wglr.Get("allocTextureUnit")
}
func (wglr *WebGLRenderer) SetAllocTextureUnit(v js.Value) {
	wglr.Set("allocTextureUnit", v)
}
func (wglr *WebGLRenderer) AutoClear() bool {
	return wglr.Get("autoClear").Bool()
}
func (wglr *WebGLRenderer) SetAutoClear(v bool) {
	wglr.Set("autoClear", v)
}
func (wglr *WebGLRenderer) AutoClearColor() bool {
	return wglr.Get("autoClearColor").Bool()
}
func (wglr *WebGLRenderer) SetAutoClearColor(v bool) {
	wglr.Set("autoClearColor", v)
}
func (wglr *WebGLRenderer) AutoClearDepth() bool {
	return wglr.Get("autoClearDepth").Bool()
}
func (wglr *WebGLRenderer) SetAutoClearDepth(v bool) {
	wglr.Set("autoClearDepth", v)
}
func (wglr *WebGLRenderer) AutoClearStencil() bool {
	return wglr.Get("autoClearStencil").Bool()
}
func (wglr *WebGLRenderer) SetAutoClearStencil(v bool) {
	wglr.Set("autoClearStencil", v)
}
func (wglr *WebGLRenderer) Capabilities() *webgl.WebGLCapabilities {
	return &webgl.WebGLCapabilities{Value: wglr.Get("capabilities")}
}
func (wglr *WebGLRenderer) SetCapabilities(v *webgl.WebGLCapabilities) {
	wglr.Set("capabilities", v)
}
func (wglr *WebGLRenderer) ClippingPlanes() []js.Value {
	return []js.Value(wglr.Get("clippingPlanes"))
}
func (wglr *WebGLRenderer) SetClippingPlanes(v []js.Value) {
	wglr.Set("clippingPlanes", v)
}
func (wglr *WebGLRenderer) Context() *WebGLRenderingContext {
	return &WebGLRenderingContext{Value: wglr.Get("context")}
}
func (wglr *WebGLRenderer) SetContext(v *WebGLRenderingContext) {
	wglr.Set("context", v)
}
func (wglr *WebGLRenderer) DomElement() *HTMLCanvasElement {
	return &HTMLCanvasElement{Value: wglr.Get("domElement")}
}
func (wglr *WebGLRenderer) SetDomElement(v *HTMLCanvasElement) {
	wglr.Set("domElement", v)
}
func (wglr *WebGLRenderer) Extensions() *webgl.WebGLExtensions {
	return &webgl.WebGLExtensions{Value: wglr.Get("extensions")}
}
func (wglr *WebGLRenderer) SetExtensions(v *webgl.WebGLExtensions) {
	wglr.Set("extensions", v)
}
func (wglr *WebGLRenderer) GammaFactor() int {
	return wglr.Get("gammaFactor").Int()
}
func (wglr *WebGLRenderer) SetGammaFactor(v int) {
	wglr.Set("gammaFactor", v)
}
func (wglr *WebGLRenderer) GammaInput() bool {
	return wglr.Get("gammaInput").Bool()
}
func (wglr *WebGLRenderer) SetGammaInput(v bool) {
	wglr.Set("gammaInput", v)
}
func (wglr *WebGLRenderer) GammaOutput() bool {
	return wglr.Get("gammaOutput").Bool()
}
func (wglr *WebGLRenderer) SetGammaOutput(v bool) {
	wglr.Set("gammaOutput", v)
}
func (wglr *WebGLRenderer) Info() *webgl.WebGLInfo {
	return &webgl.WebGLInfo{Value: wglr.Get("info")}
}
func (wglr *WebGLRenderer) SetInfo(v *webgl.WebGLInfo) {
	wglr.Set("info", v)
}
func (wglr *WebGLRenderer) LocalClippingEnabled() bool {
	return wglr.Get("localClippingEnabled").Bool()
}
func (wglr *WebGLRenderer) SetLocalClippingEnabled(v bool) {
	wglr.Set("localClippingEnabled", v)
}
func (wglr *WebGLRenderer) MaxMorphNormals() int {
	return wglr.Get("maxMorphNormals").Int()
}
func (wglr *WebGLRenderer) SetMaxMorphNormals(v int) {
	wglr.Set("maxMorphNormals", v)
}
func (wglr *WebGLRenderer) MaxMorphTargets() int {
	return wglr.Get("maxMorphTargets").Int()
}
func (wglr *WebGLRenderer) SetMaxMorphTargets(v int) {
	wglr.Set("maxMorphTargets", v)
}
func (wglr *WebGLRenderer) PhysicallyCorrectLights() bool {
	return wglr.Get("physicallyCorrectLights").Bool()
}
func (wglr *WebGLRenderer) SetPhysicallyCorrectLights(v bool) {
	wglr.Set("physicallyCorrectLights", v)
}
func (wglr *WebGLRenderer) PixelRation() int {
	return wglr.Get("pixelRation").Int()
}
func (wglr *WebGLRenderer) SetPixelRation(v int) {
	wglr.Set("pixelRation", v)
}
func (wglr *WebGLRenderer) Properties() *webgl.WebGLProperties {
	return &webgl.WebGLProperties{Value: wglr.Get("properties")}
}
func (wglr *WebGLRenderer) SetProperties(v *webgl.WebGLProperties) {
	wglr.Set("properties", v)
}
func (wglr *WebGLRenderer) RenderLists() *webgl.WebGLRenderLists {
	return &webgl.WebGLRenderLists{Value: wglr.Get("renderLists")}
}
func (wglr *WebGLRenderer) SetRenderLists(v *webgl.WebGLRenderLists) {
	wglr.Set("renderLists", v)
}
func (wglr *WebGLRenderer) ShadowMap() *webgl.WebGLShadowMap {
	return &webgl.WebGLShadowMap{Value: wglr.Get("shadowMap")}
}
func (wglr *WebGLRenderer) SetShadowMap(v *webgl.WebGLShadowMap) {
	wglr.Set("shadowMap", v)
}
func (wglr *WebGLRenderer) ShadowMapCullFace() *CullFace {
	return &CullFace{Value: wglr.Get("shadowMapCullFace")}
}
func (wglr *WebGLRenderer) SetShadowMapCullFace(v *CullFace) {
	wglr.Set("shadowMapCullFace", v)
}
func (wglr *WebGLRenderer) ShadowMapDebug() bool {
	return wglr.Get("shadowMapDebug").Bool()
}
func (wglr *WebGLRenderer) SetShadowMapDebug(v bool) {
	wglr.Set("shadowMapDebug", v)
}
func (wglr *WebGLRenderer) ShadowMapEnabled() bool {
	return wglr.Get("shadowMapEnabled").Bool()
}
func (wglr *WebGLRenderer) SetShadowMapEnabled(v bool) {
	wglr.Set("shadowMapEnabled", v)
}
func (wglr *WebGLRenderer) ShadowMapType() *ShadowMapType {
	return &ShadowMapType{Value: wglr.Get("shadowMapType")}
}
func (wglr *WebGLRenderer) SetShadowMapType(v *ShadowMapType) {
	wglr.Set("shadowMapType", v)
}
func (wglr *WebGLRenderer) SortObjects() bool {
	return wglr.Get("sortObjects").Bool()
}
func (wglr *WebGLRenderer) SetSortObjects(v bool) {
	wglr.Set("sortObjects", v)
}
func (wglr *WebGLRenderer) State() *webgl.WebGLState {
	return &webgl.WebGLState{Value: wglr.Get("state")}
}
func (wglr *WebGLRenderer) SetState(v *webgl.WebGLState) {
	wglr.Set("state", v)
}
func (wglr *WebGLRenderer) ToneMapping() *ToneMapping {
	return &ToneMapping{Value: wglr.Get("toneMapping")}
}
func (wglr *WebGLRenderer) SetToneMapping(v *ToneMapping) {
	wglr.Set("toneMapping", v)
}
func (wglr *WebGLRenderer) ToneMappingExposure() int {
	return wglr.Get("toneMappingExposure").Int()
}
func (wglr *WebGLRenderer) SetToneMappingExposure(v int) {
	wglr.Set("toneMappingExposure", v)
}
func (wglr *WebGLRenderer) ToneMappingWhitePoint() int {
	return wglr.Get("toneMappingWhitePoint").Int()
}
func (wglr *WebGLRenderer) SetToneMappingWhitePoint(v int) {
	wglr.Set("toneMappingWhitePoint", v)
}
func (wglr *WebGLRenderer) Vr() webvr.WebVRManager {
	return webvr.WebVRManager(wglr.Get("vr"))
}
func (wglr *WebGLRenderer) SetVr(v webvr.WebVRManager) {
	wglr.Set("vr", v)
}
func (wglr *WebGLRenderer) Animate(callback js.Value) {
	wglr.Call("animate", callback)
}
func (wglr *WebGLRenderer) Clear(color bool, depth bool, stencil bool) {
	wglr.Call("clear", color, depth, stencil)
}
func (wglr *WebGLRenderer) ClearColor() {
	wglr.Call("clearColor")
}
func (wglr *WebGLRenderer) ClearDepth() {
	wglr.Call("clearDepth")
}
func (wglr *WebGLRenderer) ClearStencil() {
	wglr.Call("clearStencil")
}
func (wglr *WebGLRenderer) ClearTarget(renderTarget *WebGLRenderTarget, color bool, depth bool, stencil bool) {
	wglr.Call("clearTarget", renderTarget, color, depth, stencil)
}
func (wglr *WebGLRenderer) Dispose() {
	wglr.Call("dispose")
}
func (wglr *WebGLRenderer) EnableScissorTest(boolean js.Value) js.Value {
	return wglr.Call("enableScissorTest", boolean)
}
func (wglr *WebGLRenderer) ForceContextLoss() {
	wglr.Call("forceContextLoss")
}
func (wglr *WebGLRenderer) GetClearAlpha() int {
	return wglr.Call("getClearAlpha").Int()
}
func (wglr *WebGLRenderer) GetClearColor() *math.Color {
	return &math.Color{Value: wglr.Call("getClearColor")}
}
func (wglr *WebGLRenderer) GetContext() *WebGLRenderingContext {
	return &WebGLRenderingContext{Value: wglr.Call("getContext")}
}
func (wglr *WebGLRenderer) GetContextAttributes() js.Value {
	return wglr.Call("getContextAttributes")
}
func (wglr *WebGLRenderer) GetCurrentRenderTarget() webgl.RenderTarget {
	return webgl.RenderTarget(wglr.Call("getCurrentRenderTarget"))
}
func (wglr *WebGLRenderer) GetCurrentViewport(target *math.Vector4) *math.Vector4 {
	return &math.Vector4{Value: wglr.Call("getCurrentViewport", target)}
}
func (wglr *WebGLRenderer) GetDrawingBufferSize() js.Value {
	return wglr.Call("getDrawingBufferSize")
}
func (wglr *WebGLRenderer) GetMaxAnisotropy() int {
	return wglr.Call("getMaxAnisotropy").Int()
}
func (wglr *WebGLRenderer) GetPixelRatio() int {
	return wglr.Call("getPixelRatio").Int()
}
func (wglr *WebGLRenderer) GetPrecision() string {
	return wglr.Call("getPrecision").String()
}
func (wglr *WebGLRenderer) GetRenderTarget() webgl.RenderTarget {
	return webgl.RenderTarget(wglr.Call("getRenderTarget"))
}
func (wglr *WebGLRenderer) GetScissor(target *math.Vector4) *math.Vector4 {
	return &math.Vector4{Value: wglr.Call("getScissor", target)}
}
func (wglr *WebGLRenderer) GetScissorTest() bool {
	return wglr.Call("getScissorTest").Bool()
}
func (wglr *WebGLRenderer) GetSize(target *math.Vector2) *math.Vector2 {
	return &math.Vector2{Value: wglr.Call("getSize", target)}
}
func (wglr *WebGLRenderer) GetViewport(target *math.Vector4) *math.Vector4 {
	return &math.Vector4{Value: wglr.Call("getViewport", target)}
}
func (wglr *WebGLRenderer) ReadRenderTargetPixels(renderTarget webgl.RenderTarget, x int, y int, width int, height int, buffer js.Value) {
	wglr.Call("readRenderTargetPixels", renderTarget, x, y, width, height, buffer)
}
func (wglr *WebGLRenderer) Render(scene *scenes.Scene, camera *cameras.Camera) {
	wglr.Call("render", scene, camera)
}
func (wglr *WebGLRenderer) RenderBufferDirect(camera *cameras.Camera, fog *scenes.Fog, material *materials.Material, geometryGroup js.Value, object *core.Object3D) {
	wglr.Call("renderBufferDirect", camera, fog, material, geometryGroup, object)
}
func (wglr *WebGLRenderer) RenderBufferImmediate(object *core.Object3D, program *Object, material *materials.Material) {
	wglr.Call("renderBufferImmediate", object, program, material)
}
func (wglr *WebGLRenderer) ResetGLState() {
	wglr.Call("resetGLState")
}
func (wglr *WebGLRenderer) SetAnimationLoop(callback js.Value) {
	wglr.Call("setAnimationLoop", callback)
}
func (wglr *WebGLRenderer) SetClearAlpha(alpha int) {
	wglr.Call("setClearAlpha", alpha)
}
func (wglr *WebGLRenderer) SetClearColor(color *math.Color, alpha int) {
	wglr.Call("setClearColor", color, alpha)
}
func (wglr *WebGLRenderer) SetClearColor2(color string, alpha int) {
	wglr.Call("setClearColor", color, alpha)
}
func (wglr *WebGLRenderer) SetClearColor3(color int, alpha int) {
	wglr.Call("setClearColor", color, alpha)
}
func (wglr *WebGLRenderer) SetDrawingBufferSize(width int, height int, pixelRatio int) {
	wglr.Call("setDrawingBufferSize", width, height, pixelRatio)
}
func (wglr *WebGLRenderer) SetPixelRatio(value int) {
	wglr.Call("setPixelRatio", value)
}
func (wglr *WebGLRenderer) SetRenderTarget(renderTarget webgl.RenderTarget, activeCubeFace int, activeMipMapLevel int) {
	wglr.Call("setRenderTarget", renderTarget, activeCubeFace, activeMipMapLevel)
}
func (wglr *WebGLRenderer) SetScissor(x math.Vector4, y int, width int, height int) {
	wglr.Call("setScissor", x, y, width, height)
}
func (wglr *WebGLRenderer) SetScissorTest(enable bool) {
	wglr.Call("setScissorTest", enable)
}
func (wglr *WebGLRenderer) SetSize(width int, height int, updateStyle bool) {
	wglr.Call("setSize", width, height, updateStyle)
}
func (wglr *WebGLRenderer) SetTexture(texture *textures.Texture, slot int) {
	wglr.Call("setTexture", texture, slot)
}
func (wglr *WebGLRenderer) SetTexture2D(texture *textures.Texture, slot int) {
	wglr.Call("setTexture2D", texture, slot)
}
func (wglr *WebGLRenderer) SetTextureCube(texture *textures.Texture, slot int) {
	wglr.Call("setTextureCube", texture, slot)
}
func (wglr *WebGLRenderer) SetViewport(x math.Vector4, y int, width int, height int) {
	wglr.Call("setViewport", x, y, width, height)
}
func (wglr *WebGLRenderer) SupportsBlendMinMax() js.Value {
	return wglr.Call("supportsBlendMinMax")
}
func (wglr *WebGLRenderer) SupportsCompressedTexturePVRTC() js.Value {
	return wglr.Call("supportsCompressedTexturePVRTC")
}
func (wglr *WebGLRenderer) SupportsCompressedTextureS3TC() js.Value {
	return wglr.Call("supportsCompressedTextureS3TC")
}
func (wglr *WebGLRenderer) SupportsFloatTextures() js.Value {
	return wglr.Call("supportsFloatTextures")
}
func (wglr *WebGLRenderer) SupportsHalfFloatTextures() js.Value {
	return wglr.Call("supportsHalfFloatTextures")
}
func (wglr *WebGLRenderer) SupportsInstancedArrays() js.Value {
	return wglr.Call("supportsInstancedArrays")
}
func (wglr *WebGLRenderer) SupportsStandardDerivatives() js.Value {
	return wglr.Call("supportsStandardDerivatives")
}
func (wglr *WebGLRenderer) SupportsVertexTextures() js.Value {
	return wglr.Call("supportsVertexTextures")
}

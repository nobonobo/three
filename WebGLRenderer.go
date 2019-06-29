// Code generated from "renderers/WebGLRenderer.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Renderer interface {
	Render(scene Scene, camera Camera)
	SetSize(width float64, height float64, updateStyle bool)
}
type WebGLDebug interface {
}
type WebGLRendererParameters interface {
}

// WebGLRenderer extend: []
type WebGLRenderer struct {
	js.Value
}

func NewWebGLRenderer(parameters WebGLRendererParameters) *WebGLRenderer {
	return &WebGLRenderer{Value: get("WebGLRenderer").New(parameters)}
}
func (wglr *WebGLRenderer) JSValue() js.Value {
	return wglr.Value
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
func (wglr *WebGLRenderer) Capabilities() *WebGLCapabilities {
	return &WebGLCapabilities{Value: wglr.Get("capabilities")}
}
func (wglr *WebGLRenderer) SetCapabilities(v *WebGLCapabilities) {
	wglr.Set("capabilities", v.JSValue())
}
func (wglr *WebGLRenderer) ClippingPlanes() js.Value {
	return wglr.Get("clippingPlanes")
}
func (wglr *WebGLRenderer) SetClippingPlanes(v js.Value) {
	wglr.Set("clippingPlanes", v)
}
func (wglr *WebGLRenderer) Context() js.Value {
	return wglr.Get("context")
}
func (wglr *WebGLRenderer) SetContext(v js.Value) {
	wglr.Set("context", v)
}
func (wglr *WebGLRenderer) Debug() js.Value {
	return wglr.Get("debug")
}
func (wglr *WebGLRenderer) SetDebug(v js.Value) {
	wglr.Set("debug", v)
}
func (wglr *WebGLRenderer) DomElement() js.Value {
	return wglr.Get("domElement")
}
func (wglr *WebGLRenderer) SetDomElement(v js.Value) {
	wglr.Set("domElement", v)
}
func (wglr *WebGLRenderer) Extensions() *WebGLExtensions {
	return &WebGLExtensions{Value: wglr.Get("extensions")}
}
func (wglr *WebGLRenderer) SetExtensions(v *WebGLExtensions) {
	wglr.Set("extensions", v.JSValue())
}
func (wglr *WebGLRenderer) GammaFactor() float64 {
	return wglr.Get("gammaFactor").Float()
}
func (wglr *WebGLRenderer) SetGammaFactor(v float64) {
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
func (wglr *WebGLRenderer) Info() *WebGLInfo {
	return &WebGLInfo{Value: wglr.Get("info")}
}
func (wglr *WebGLRenderer) SetInfo(v *WebGLInfo) {
	wglr.Set("info", v.JSValue())
}
func (wglr *WebGLRenderer) LocalClippingEnabled() bool {
	return wglr.Get("localClippingEnabled").Bool()
}
func (wglr *WebGLRenderer) SetLocalClippingEnabled(v bool) {
	wglr.Set("localClippingEnabled", v)
}
func (wglr *WebGLRenderer) MaxMorphNormals() float64 {
	return wglr.Get("maxMorphNormals").Float()
}
func (wglr *WebGLRenderer) SetMaxMorphNormals(v float64) {
	wglr.Set("maxMorphNormals", v)
}
func (wglr *WebGLRenderer) MaxMorphTargets() float64 {
	return wglr.Get("maxMorphTargets").Float()
}
func (wglr *WebGLRenderer) SetMaxMorphTargets(v float64) {
	wglr.Set("maxMorphTargets", v)
}
func (wglr *WebGLRenderer) PhysicallyCorrectLights() bool {
	return wglr.Get("physicallyCorrectLights").Bool()
}
func (wglr *WebGLRenderer) SetPhysicallyCorrectLights(v bool) {
	wglr.Set("physicallyCorrectLights", v)
}
func (wglr *WebGLRenderer) PixelRation() float64 {
	return wglr.Get("pixelRation").Float()
}
func (wglr *WebGLRenderer) SetPixelRation(v float64) {
	wglr.Set("pixelRation", v)
}
func (wglr *WebGLRenderer) Properties() *WebGLProperties {
	return &WebGLProperties{Value: wglr.Get("properties")}
}
func (wglr *WebGLRenderer) SetProperties(v *WebGLProperties) {
	wglr.Set("properties", v.JSValue())
}
func (wglr *WebGLRenderer) RenderLists() *WebGLRenderLists {
	return &WebGLRenderLists{Value: wglr.Get("renderLists")}
}
func (wglr *WebGLRenderer) SetRenderLists(v *WebGLRenderLists) {
	wglr.Set("renderLists", v.JSValue())
}
func (wglr *WebGLRenderer) ShadowMap() *WebGLShadowMap {
	return &WebGLShadowMap{Value: wglr.Get("shadowMap")}
}
func (wglr *WebGLRenderer) SetShadowMap(v *WebGLShadowMap) {
	wglr.Set("shadowMap", v.JSValue())
}
func (wglr *WebGLRenderer) ShadowMapCullFace() CullFace {
	return CullFace(wglr.Get("shadowMapCullFace"))
}
func (wglr *WebGLRenderer) SetShadowMapCullFace(v CullFace) {
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
func (wglr *WebGLRenderer) ShadowMapType() ShadowMapType {
	return ShadowMapType(wglr.Get("shadowMapType"))
}
func (wglr *WebGLRenderer) SetShadowMapType(v ShadowMapType) {
	wglr.Set("shadowMapType", v)
}
func (wglr *WebGLRenderer) SortObjects() bool {
	return wglr.Get("sortObjects").Bool()
}
func (wglr *WebGLRenderer) SetSortObjects(v bool) {
	wglr.Set("sortObjects", v)
}
func (wglr *WebGLRenderer) State() *WebGLState {
	return &WebGLState{Value: wglr.Get("state")}
}
func (wglr *WebGLRenderer) SetState(v *WebGLState) {
	wglr.Set("state", v.JSValue())
}
func (wglr *WebGLRenderer) ToneMapping() ToneMapping {
	return ToneMapping(wglr.Get("toneMapping"))
}
func (wglr *WebGLRenderer) SetToneMapping(v ToneMapping) {
	wglr.Set("toneMapping", v)
}
func (wglr *WebGLRenderer) ToneMappingExposure() float64 {
	return wglr.Get("toneMappingExposure").Float()
}
func (wglr *WebGLRenderer) SetToneMappingExposure(v float64) {
	wglr.Set("toneMappingExposure", v)
}
func (wglr *WebGLRenderer) ToneMappingWhitePoint() float64 {
	return wglr.Get("toneMappingWhitePoint").Float()
}
func (wglr *WebGLRenderer) SetToneMappingWhitePoint(v float64) {
	wglr.Set("toneMappingWhitePoint", v)
}
func (wglr *WebGLRenderer) Vr() js.Value {
	return wglr.Get("vr")
}
func (wglr *WebGLRenderer) SetVr(v js.Value) {
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
	wglr.Call("clearTarget", renderTarget.JSValue(), color, depth, stencil)
}
func (wglr *WebGLRenderer) Compile(scene Scene, camera Camera) {
	wglr.Call("compile", scene.JSValue(), camera.JSValue())
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
func (wglr *WebGLRenderer) GetClearAlpha() float64 {
	return wglr.Call("getClearAlpha").Float()
}
func (wglr *WebGLRenderer) GetClearColor() *Color {
	return &Color{Value: wglr.Call("getClearColor")}
}
func (wglr *WebGLRenderer) GetContext() js.Value {
	return wglr.Call("getContext")
}
func (wglr *WebGLRenderer) GetContextAttributes() js.Value {
	return wglr.Call("getContextAttributes")
}
func (wglr *WebGLRenderer) GetCurrentRenderTarget() js.Value {
	return wglr.Call("getCurrentRenderTarget")
}
func (wglr *WebGLRenderer) GetCurrentViewport(target *Vector4) *Vector4 {
	return &Vector4{Value: wglr.Call("getCurrentViewport", target)}
}
func (wglr *WebGLRenderer) GetDrawingBufferSize(target *Vector2) *Vector2 {
	return &Vector2{Value: wglr.Call("getDrawingBufferSize", target)}
}
func (wglr *WebGLRenderer) GetMaxAnisotropy() float64 {
	return wglr.Call("getMaxAnisotropy").Float()
}
func (wglr *WebGLRenderer) GetPixelRatio() float64 {
	return wglr.Call("getPixelRatio").Float()
}
func (wglr *WebGLRenderer) GetPrecision() string {
	return wglr.Call("getPrecision").String()
}
func (wglr *WebGLRenderer) GetRenderTarget() js.Value {
	return wglr.Call("getRenderTarget")
}
func (wglr *WebGLRenderer) GetScissor(target *Vector4) *Vector4 {
	return &Vector4{Value: wglr.Call("getScissor", target)}
}
func (wglr *WebGLRenderer) GetScissorTest() bool {
	return wglr.Call("getScissorTest").Bool()
}
func (wglr *WebGLRenderer) GetSize(target *Vector2) *Vector2 {
	return &Vector2{Value: wglr.Call("getSize", target)}
}
func (wglr *WebGLRenderer) GetViewport(target *Vector4) *Vector4 {
	return &Vector4{Value: wglr.Call("getViewport", target)}
}
func (wglr *WebGLRenderer) ReadRenderTargetPixels(renderTarget js.Value, x float64, y float64, width float64, height float64, buffer js.Value) {
	wglr.Call("readRenderTargetPixels", renderTarget, x, y, width, height, buffer)
}
func (wglr *WebGLRenderer) Render(scene Scene, camera Camera) {
	wglr.Call("render", scene.JSValue(), camera.JSValue())
}
func (wglr *WebGLRenderer) RenderBufferDirect(camera Camera, fog *Fog, material Material, geometryGroup js.Value, object *Object3D) {
	wglr.Call("renderBufferDirect", camera.JSValue(), fog.JSValue(), material.JSValue(), geometryGroup, object.JSValue())
}
func (wglr *WebGLRenderer) RenderBufferImmediate(object *Object3D, program js.Value, material Material) {
	wglr.Call("renderBufferImmediate", object.JSValue(), program, material.JSValue())
}
func (wglr *WebGLRenderer) ResetGLState() {
	wglr.Call("resetGLState")
}
func (wglr *WebGLRenderer) SetAnimationLoop(callback js.Value) {
	wglr.Call("setAnimationLoop", callback)
}
func (wglr *WebGLRenderer) SetClearAlpha(alpha float64) {
	wglr.Call("setClearAlpha", alpha)
}
func (wglr *WebGLRenderer) SetClearColor(color *Color, alpha float64) {
	wglr.Call("setClearColor", color.JSValue(), alpha)
}
func (wglr *WebGLRenderer) SetClearColor2(color string, alpha float64) {
	wglr.Call("setClearColor", color, alpha)
}
func (wglr *WebGLRenderer) SetClearColor3(color int, alpha float64) {
	wglr.Call("setClearColor", color, alpha)
}
func (wglr *WebGLRenderer) SetDrawingBufferSize(width float64, height float64, pixelRatio float64) {
	wglr.Call("setDrawingBufferSize", width, height, pixelRatio)
}
func (wglr *WebGLRenderer) SetPixelRatio(value float64) {
	wglr.Call("setPixelRatio", value)
}
func (wglr *WebGLRenderer) SetRenderTarget(renderTarget js.Value, activeCubeFace float64, activeMipMapLevel float64) {
	wglr.Call("setRenderTarget", renderTarget, activeCubeFace, activeMipMapLevel)
}
func (wglr *WebGLRenderer) SetScissor(x *Vector4, y float64, width float64, height float64) {
	wglr.Call("setScissor", x, y, width, height)
}
func (wglr *WebGLRenderer) SetScissorTest(enable bool) {
	wglr.Call("setScissorTest", enable)
}
func (wglr *WebGLRenderer) SetSize(width float64, height float64, updateStyle bool) {
	wglr.Call("setSize", width, height, updateStyle)
}
func (wglr *WebGLRenderer) SetViewport(x *Vector4, y float64, width float64, height float64) {
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

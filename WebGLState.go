// Code generated from "renderers/webgl/WebGLState.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLColorBuffer struct {
	js.Value
}

func NewWebGLColorBuffer() *WebGLColorBuffer {
	return &WebGLColorBuffer{Value: get("WebGLColorBuffer").New()}
}
func (wglcb *WebGLColorBuffer) Reset() {
	wglcb.Call("reset")
}
func (wglcb *WebGLColorBuffer) SetClear(r float64, g float64, b float64, a float64, premultipliedAlpha bool) {
	wglcb.Call("setClear", r, g, b, a, premultipliedAlpha)
}
func (wglcb *WebGLColorBuffer) SetLocked(lock bool) {
	wglcb.Call("setLocked", lock)
}
func (wglcb *WebGLColorBuffer) SetMask(colorMask bool) {
	wglcb.Call("setMask", colorMask)
}

type WebGLDepthBuffer struct {
	js.Value
}

func NewWebGLDepthBuffer() *WebGLDepthBuffer {
	return &WebGLDepthBuffer{Value: get("WebGLDepthBuffer").New()}
}
func (wgldb *WebGLDepthBuffer) Reset() {
	wgldb.Call("reset")
}
func (wgldb *WebGLDepthBuffer) SetClear(depth int) {
	wgldb.Call("setClear", depth)
}
func (wgldb *WebGLDepthBuffer) SetFunc(depthFunc DepthModes) {
	wgldb.Call("setFunc", depthFunc)
}
func (wgldb *WebGLDepthBuffer) SetLocked(lock bool) {
	wgldb.Call("setLocked", lock)
}
func (wgldb *WebGLDepthBuffer) SetMask(depthMask bool) {
	wgldb.Call("setMask", depthMask)
}
func (wgldb *WebGLDepthBuffer) SetTest(depthTest bool) {
	wgldb.Call("setTest", depthTest)
}

type WebGLState struct {
	js.Value
}

func NewWebGLState(gl js.Value, extensions *WebGLExtensions, utils js.Value, capabilities *WebGLCapabilities) *WebGLState {
	return &WebGLState{Value: get("WebGLState").New(gl, extensions, utils, capabilities)}
}
func (wgls *WebGLState) Buffers() js.Value {
	return wgls.Get("buffers")
}
func (wgls *WebGLState) SetBuffers(v js.Value) {
	wgls.Set("buffers", v)
}
func (wgls *WebGLState) ActiveTexture(webglSlot int) {
	wgls.Call("activeTexture", webglSlot)
}
func (wgls *WebGLState) BindTexture(webglType int, webglTexture js.Value) {
	wgls.Call("bindTexture", webglType, webglTexture)
}
func (wgls *WebGLState) CompressedTexImage2D(target int, level int, internalformat int, width int, height int, border int, data js.Value) {
	wgls.Call("compressedTexImage2D", target, level, internalformat, width, height, border, data)
}
func (wgls *WebGLState) Disable(id int) {
	wgls.Call("disable", id)
}
func (wgls *WebGLState) DisableUnusedAttributes() {
	wgls.Call("disableUnusedAttributes")
}
func (wgls *WebGLState) Enable(id int) {
	wgls.Call("enable", id)
}
func (wgls *WebGLState) EnableAttribute(attribute int) {
	wgls.Call("enableAttribute", attribute)
}
func (wgls *WebGLState) EnableAttributeAndDivisor(attribute int, meshPerAttribute int) {
	wgls.Call("enableAttributeAndDivisor", attribute, meshPerAttribute)
}
func (wgls *WebGLState) GetCompressedTextureFormats() js.Value {
	return wgls.Call("getCompressedTextureFormats")
}
func (wgls *WebGLState) InitAttributes() {
	wgls.Call("initAttributes")
}
func (wgls *WebGLState) Reset() {
	wgls.Call("reset")
}
func (wgls *WebGLState) Scissor(scissor *Vector4) {
	wgls.Call("scissor", scissor)
}
func (wgls *WebGLState) SetBlending(blending Blending, blendEquation BlendingEquation, blendSrc js.Value, blendDst BlendingDstFactor, blendEquationAlpha BlendingEquation, blendSrcAlpha js.Value, blendDstAlpha BlendingDstFactor, premultiplyAlpha bool) {
	wgls.Call("setBlending", blending, blendEquation, blendSrc, blendDst, blendEquationAlpha, blendSrcAlpha, blendDstAlpha, premultiplyAlpha)
}
func (wgls *WebGLState) SetCullFace(cullFace CullFace) {
	wgls.Call("setCullFace", cullFace)
}
func (wgls *WebGLState) SetFlipSided(flipSided bool) {
	wgls.Call("setFlipSided", flipSided)
}
func (wgls *WebGLState) SetLineWidth(width float64) {
	wgls.Call("setLineWidth", width)
}
func (wgls *WebGLState) SetMaterial(material *Material, frontFaceCW bool) {
	wgls.Call("setMaterial", material, frontFaceCW)
}
func (wgls *WebGLState) SetPolygonOffset(polygonoffset bool, factor float64, units float64) {
	wgls.Call("setPolygonOffset", polygonoffset, factor, units)
}
func (wgls *WebGLState) SetScissorTest(scissorTest bool) {
	wgls.Call("setScissorTest", scissorTest)
}
func (wgls *WebGLState) TexImage2D(target int, level int, internalformat int, width int, height int, border int, format int, typ int, pixels js.Value) {
	wgls.Call("texImage2D", target, level, internalformat, width, height, border, format, typ, pixels)
}
func (wgls *WebGLState) TexImage2D2(target int, level int, internalformat int, format int, typ int, source js.Value) {
	wgls.Call("texImage2D", target, level, internalformat, format, typ, source)
}
func (wgls *WebGLState) TexImage3D(target int, level int, internalformat int, width int, height int, depth int, border int, format int, typ int, pixels js.Value) {
	wgls.Call("texImage3D", target, level, internalformat, width, height, depth, border, format, typ, pixels)
}
func (wgls *WebGLState) UseProgram(program js.Value) bool {
	return wgls.Call("useProgram", program).Bool()
}
func (wgls *WebGLState) Viewport(viewport *Vector4) {
	wgls.Call("viewport", viewport)
}

type WebGLStencilBuffer struct {
	js.Value
}

func NewWebGLStencilBuffer() *WebGLStencilBuffer {
	return &WebGLStencilBuffer{Value: get("WebGLStencilBuffer").New()}
}
func (wglsb *WebGLStencilBuffer) Reset() {
	wglsb.Call("reset")
}
func (wglsb *WebGLStencilBuffer) SetClear(stencil int) {
	wglsb.Call("setClear", stencil)
}
func (wglsb *WebGLStencilBuffer) SetFunc(stencilFunc float64, stencilRef int, stencilMask float64) {
	wglsb.Call("setFunc", stencilFunc, stencilRef, stencilMask)
}
func (wglsb *WebGLStencilBuffer) SetLocked(lock bool) {
	wglsb.Call("setLocked", lock)
}
func (wglsb *WebGLStencilBuffer) SetMask(stencilMask float64) {
	wglsb.Call("setMask", stencilMask)
}
func (wglsb *WebGLStencilBuffer) SetOp(stencilFail int, stencilZFail int, stencilZPass int) {
	wglsb.Call("setOp", stencilFail, stencilZFail, stencilZPass)
}
func (wglsb *WebGLStencilBuffer) SetTest(stencilTest bool) {
	wglsb.Call("setTest", stencilTest)
}

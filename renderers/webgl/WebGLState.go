package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLColorBuffer struct {
	js.Value
}

func NewWebGLColorBuffer(gl js.Value, state js.Value) *WebGLColorBuffer {
	return &WebGLColorBuffer{Value: get("WebGLColorBuffer").New(gl, state)}
}
func (wglcb *WebGLColorBuffer) Reset() {
	wglcb.Call("reset")
}
func (wglcb *WebGLColorBuffer) SetClear(r int, g int, b int, a int) {
	wglcb.Call("setClear", r, g, b, a)
}
func (wglcb *WebGLColorBuffer) SetLocked(lock bool) {
	wglcb.Call("setLocked", lock)
}
func (wglcb *WebGLColorBuffer) SetMask(colorMask int) {
	wglcb.Call("setMask", colorMask)
}

type WebGLDepthBuffer struct {
	js.Value
}

func NewWebGLDepthBuffer(gl js.Value, state js.Value) *WebGLDepthBuffer {
	return &WebGLDepthBuffer{Value: get("WebGLDepthBuffer").New(gl, state)}
}
func (wgldb *WebGLDepthBuffer) Reset() {
	wgldb.Call("reset")
}
func (wgldb *WebGLDepthBuffer) SetClear(depth js.Value) {
	wgldb.Call("setClear", depth)
}
func (wgldb *WebGLDepthBuffer) SetFunc(depthFunc int) {
	wgldb.Call("setFunc", depthFunc)
}
func (wgldb *WebGLDepthBuffer) SetLocked(lock bool) {
	wgldb.Call("setLocked", lock)
}
func (wgldb *WebGLDepthBuffer) SetMask(depthMask int) {
	wgldb.Call("setMask", depthMask)
}
func (wgldb *WebGLDepthBuffer) SetTest(depthTest bool) {
	wgldb.Call("setTest", depthTest)
}

type WebGLState struct {
	js.Value
}

func NewWebGLState(gl js.Value, extensions js.Value, paramThreeToGL js.Value) *WebGLState {
	return &WebGLState{Value: get("WebGLState").New(gl, extensions, paramThreeToGL)}
}
func (wgls *WebGLState) Buffers() js.Value {
	return wgls.Get("buffers")
}
func (wgls *WebGLState) SetBuffers(v js.Value) {
	wgls.Set("buffers", v)
}
func (wgls *WebGLState) ActiveTexture(webglSlot js.Value) {
	wgls.Call("activeTexture", webglSlot)
}
func (wgls *WebGLState) BindTexture(webglType js.Value, webglTexture js.Value) {
	wgls.Call("bindTexture", webglType, webglTexture)
}
func (wgls *WebGLState) ClearColor(r int, g int, b int, a int) {
	wgls.Call("clearColor", r, g, b, a)
}
func (wgls *WebGLState) ClearDepth(depth int) {
	wgls.Call("clearDepth", depth)
}
func (wgls *WebGLState) ClearStencil(stencil js.Value) {
	wgls.Call("clearStencil", stencil)
}
func (wgls *WebGLState) CompressedTexImage2D() {
	wgls.Call("compressedTexImage2D")
}
func (wgls *WebGLState) Disable(id string) {
	wgls.Call("disable", id)
}
func (wgls *WebGLState) DisableUnusedAttributes() {
	wgls.Call("disableUnusedAttributes")
}
func (wgls *WebGLState) Enable(id string) {
	wgls.Call("enable", id)
}
func (wgls *WebGLState) EnableAttribute(attribute string) {
	wgls.Call("enableAttribute", attribute)
}
func (wgls *WebGLState) EnableAttributeAndDivisor(attribute string, meshPerAttribute js.Value, extension js.Value) {
	wgls.Call("enableAttributeAndDivisor", attribute, meshPerAttribute, extension)
}
func (wgls *WebGLState) GetCompressedTextureFormats() []js.Value {
	return []js.Value(wgls.Call("getCompressedTextureFormats"))
}
func (wgls *WebGLState) GetScissorTest() bool {
	return wgls.Call("getScissorTest").Bool()
}
func (wgls *WebGLState) Init() {
	wgls.Call("init")
}
func (wgls *WebGLState) InitAttributes() {
	wgls.Call("initAttributes")
}
func (wgls *WebGLState) Reset() {
	wgls.Call("reset")
}
func (wgls *WebGLState) Scissor(scissor js.Value) {
	wgls.Call("scissor", scissor)
}
func (wgls *WebGLState) SetBlending(blending int, blendEquation int, blendSrc int, blendDst int, blendEquationAlpha int, blendSrcAlpha int, blendDstAlpha int, premultiplyAlpha bool) {
	wgls.Call("setBlending", blending, blendEquation, blendSrc, blendDst, blendEquationAlpha, blendSrcAlpha, blendDstAlpha, premultiplyAlpha)
}
func (wgls *WebGLState) SetColorWrite(colorWrite int) {
	wgls.Call("setColorWrite", colorWrite)
}
func (wgls *WebGLState) SetCullFace(cullFace *CullFace) {
	wgls.Call("setCullFace", cullFace)
}
func (wgls *WebGLState) SetDepthFunc(depthFunc js.Value) {
	wgls.Call("setDepthFunc", depthFunc)
}
func (wgls *WebGLState) SetDepthTest(depthTest int) {
	wgls.Call("setDepthTest", depthTest)
}
func (wgls *WebGLState) SetDepthWrite(depthWrite int) {
	wgls.Call("setDepthWrite", depthWrite)
}
func (wgls *WebGLState) SetFlipSided(flipSided int) {
	wgls.Call("setFlipSided", flipSided)
}
func (wgls *WebGLState) SetLineWidth(width int) {
	wgls.Call("setLineWidth", width)
}
func (wgls *WebGLState) SetPolygonOffset(polygonoffset int, factor int, units int) {
	wgls.Call("setPolygonOffset", polygonoffset, factor, units)
}
func (wgls *WebGLState) SetScissorTest(scissorTest bool) {
	wgls.Call("setScissorTest", scissorTest)
}
func (wgls *WebGLState) SetStencilFunc(stencilFunc js.Value, stencilRef js.Value, stencilMask int) {
	wgls.Call("setStencilFunc", stencilFunc, stencilRef, stencilMask)
}
func (wgls *WebGLState) SetStencilOp(stencilFail js.Value, stencilZFail js.Value, stencilZPass js.Value) {
	wgls.Call("setStencilOp", stencilFail, stencilZFail, stencilZPass)
}
func (wgls *WebGLState) SetStencilTest(stencilTest bool) {
	wgls.Call("setStencilTest", stencilTest)
}
func (wgls *WebGLState) SetStencilWrite(stencilWrite js.Value) {
	wgls.Call("setStencilWrite", stencilWrite)
}
func (wgls *WebGLState) TexImage2D() {
	wgls.Call("texImage2D")
}
func (wgls *WebGLState) Viewport(viewport js.Value) {
	wgls.Call("viewport", viewport)
}

type WebGLStencilBuffer struct {
	js.Value
}

func NewWebGLStencilBuffer(gl js.Value, state js.Value) *WebGLStencilBuffer {
	return &WebGLStencilBuffer{Value: get("WebGLStencilBuffer").New(gl, state)}
}
func (wglsb *WebGLStencilBuffer) Reset() {
	wglsb.Call("reset")
}
func (wglsb *WebGLStencilBuffer) SetClear(stencil js.Value) {
	wglsb.Call("setClear", stencil)
}
func (wglsb *WebGLStencilBuffer) SetFunc(stencilFunc int, stencilRef js.Value, stencilMask int) {
	wglsb.Call("setFunc", stencilFunc, stencilRef, stencilMask)
}
func (wglsb *WebGLStencilBuffer) SetLocked(lock bool) {
	wglsb.Call("setLocked", lock)
}
func (wglsb *WebGLStencilBuffer) SetMask(stencilMask int) {
	wglsb.Call("setMask", stencilMask)
}
func (wglsb *WebGLStencilBuffer) SetOp(stencilFail js.Value, stencilZFail js.Value, stencilZPass js.Value) {
	wglsb.Call("setOp", stencilFail, stencilZFail, stencilZPass)
}
func (wglsb *WebGLStencilBuffer) SetTest(stencilTest bool) {
	wglsb.Call("setTest", stencilTest)
}

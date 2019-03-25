package renderers

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/textures"
)

type WebGLRenderTargetCube struct {
	js.Value
}

func NewWebGLRenderTargetCube(width int, height int, options WebGLRenderTargetOptions) *WebGLRenderTargetCube {
	return &WebGLRenderTargetCube{Value: get("WebGLRenderTargetCube").New(width, height, options)}
}
func (wglrtc *WebGLRenderTargetCube) Anisotropy() js.Value {
	return wglrtc.Get("anisotropy")
}
func (wglrtc *WebGLRenderTargetCube) SetAnisotropy(v js.Value) {
	wglrtc.Set("anisotropy", v)
}
func (wglrtc *WebGLRenderTargetCube) DepthBuffer() bool {
	return wglrtc.Get("depthBuffer").Bool()
}
func (wglrtc *WebGLRenderTargetCube) SetDepthBuffer(v bool) {
	wglrtc.Set("depthBuffer", v)
}
func (wglrtc *WebGLRenderTargetCube) DepthTexture() *textures.Texture {
	return &textures.Texture{Value: wglrtc.Get("depthTexture")}
}
func (wglrtc *WebGLRenderTargetCube) SetDepthTexture(v *textures.Texture) {
	wglrtc.Set("depthTexture", v)
}
func (wglrtc *WebGLRenderTargetCube) Format() js.Value {
	return wglrtc.Get("format")
}
func (wglrtc *WebGLRenderTargetCube) SetFormat(v js.Value) {
	wglrtc.Set("format", v)
}
func (wglrtc *WebGLRenderTargetCube) GenerateMipmaps() js.Value {
	return wglrtc.Get("generateMipmaps")
}
func (wglrtc *WebGLRenderTargetCube) SetGenerateMipmaps(v js.Value) {
	wglrtc.Set("generateMipmaps", v)
}
func (wglrtc *WebGLRenderTargetCube) Height() int {
	return wglrtc.Get("height").Int()
}
func (wglrtc *WebGLRenderTargetCube) SetHeight(v int) {
	wglrtc.Set("height", v)
}
func (wglrtc *WebGLRenderTargetCube) MagFilter() js.Value {
	return wglrtc.Get("magFilter")
}
func (wglrtc *WebGLRenderTargetCube) SetMagFilter(v js.Value) {
	wglrtc.Set("magFilter", v)
}
func (wglrtc *WebGLRenderTargetCube) MinFilter() js.Value {
	return wglrtc.Get("minFilter")
}
func (wglrtc *WebGLRenderTargetCube) SetMinFilter(v js.Value) {
	wglrtc.Set("minFilter", v)
}
func (wglrtc *WebGLRenderTargetCube) Offset() js.Value {
	return wglrtc.Get("offset")
}
func (wglrtc *WebGLRenderTargetCube) SetOffset(v js.Value) {
	wglrtc.Set("offset", v)
}
func (wglrtc *WebGLRenderTargetCube) Repeat() js.Value {
	return wglrtc.Get("repeat")
}
func (wglrtc *WebGLRenderTargetCube) SetRepeat(v js.Value) {
	wglrtc.Set("repeat", v)
}
func (wglrtc *WebGLRenderTargetCube) Scissor() *math.Vector4 {
	return &math.Vector4{Value: wglrtc.Get("scissor")}
}
func (wglrtc *WebGLRenderTargetCube) SetScissor(v *math.Vector4) {
	wglrtc.Set("scissor", v)
}
func (wglrtc *WebGLRenderTargetCube) ScissorTest() bool {
	return wglrtc.Get("scissorTest").Bool()
}
func (wglrtc *WebGLRenderTargetCube) SetScissorTest(v bool) {
	wglrtc.Set("scissorTest", v)
}
func (wglrtc *WebGLRenderTargetCube) StencilBuffer() bool {
	return wglrtc.Get("stencilBuffer").Bool()
}
func (wglrtc *WebGLRenderTargetCube) SetStencilBuffer(v bool) {
	wglrtc.Set("stencilBuffer", v)
}
func (wglrtc *WebGLRenderTargetCube) Texture() *textures.Texture {
	return &textures.Texture{Value: wglrtc.Get("texture")}
}
func (wglrtc *WebGLRenderTargetCube) SetTexture(v *textures.Texture) {
	wglrtc.Set("texture", v)
}
func (wglrtc *WebGLRenderTargetCube) Type() js.Value {
	return wglrtc.Get("type")
}
func (wglrtc *WebGLRenderTargetCube) SetType(v js.Value) {
	wglrtc.Set("type", v)
}
func (wglrtc *WebGLRenderTargetCube) Uuid() string {
	return wglrtc.Get("uuid").String()
}
func (wglrtc *WebGLRenderTargetCube) SetUuid(v string) {
	wglrtc.Set("uuid", v)
}
func (wglrtc *WebGLRenderTargetCube) Viewport() *math.Vector4 {
	return &math.Vector4{Value: wglrtc.Get("viewport")}
}
func (wglrtc *WebGLRenderTargetCube) SetViewport(v *math.Vector4) {
	wglrtc.Set("viewport", v)
}
func (wglrtc *WebGLRenderTargetCube) Width() int {
	return wglrtc.Get("width").Int()
}
func (wglrtc *WebGLRenderTargetCube) SetWidth(v int) {
	wglrtc.Set("width", v)
}
func (wglrtc *WebGLRenderTargetCube) WrapS() js.Value {
	return wglrtc.Get("wrapS")
}
func (wglrtc *WebGLRenderTargetCube) SetWrapS(v js.Value) {
	wglrtc.Set("wrapS", v)
}
func (wglrtc *WebGLRenderTargetCube) WrapT() js.Value {
	return wglrtc.Get("wrapT")
}
func (wglrtc *WebGLRenderTargetCube) SetWrapT(v js.Value) {
	wglrtc.Set("wrapT", v)
}
func (wglrtc *WebGLRenderTargetCube) AddEventListener(typ string, listener js.Value) {
	wglrtc.Call("addEventListener", typ, listener)
}
func (wglrtc *WebGLRenderTargetCube) Clone() this {
	return this(wglrtc.Call("clone"))
}
func (wglrtc *WebGLRenderTargetCube) Copy(source *WebGLRenderTarget) this {
	return this(wglrtc.Call("copy", source))
}
func (wglrtc *WebGLRenderTargetCube) DispatchEvent(event js.Value) {
	wglrtc.Call("dispatchEvent", event)
}
func (wglrtc *WebGLRenderTargetCube) Dispose() {
	wglrtc.Call("dispose")
}
func (wglrtc *WebGLRenderTargetCube) HasEventListener(typ string, listener js.Value) bool {
	return wglrtc.Call("hasEventListener", typ, listener).Bool()
}
func (wglrtc *WebGLRenderTargetCube) RemoveEventListener(typ string, listener js.Value) {
	wglrtc.Call("removeEventListener", typ, listener)
}
func (wglrtc *WebGLRenderTargetCube) SetSize(width int, height int) {
	wglrtc.Call("setSize", width, height)
}

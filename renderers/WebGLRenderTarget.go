package renderers

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/textures"
)

type WebGLRenderTargetOptions interface {
}
type WebGLRenderTarget struct {
	js.Value
}

func NewWebGLRenderTarget(width int, height int, options WebGLRenderTargetOptions) *WebGLRenderTarget {
	return &WebGLRenderTarget{Value: get("WebGLRenderTarget").New(width, height, options)}
}
func (wglrt *WebGLRenderTarget) Anisotropy() js.Value {
	return wglrt.Get("anisotropy")
}
func (wglrt *WebGLRenderTarget) SetAnisotropy(v js.Value) {
	wglrt.Set("anisotropy", v)
}
func (wglrt *WebGLRenderTarget) DepthBuffer() bool {
	return wglrt.Get("depthBuffer").Bool()
}
func (wglrt *WebGLRenderTarget) SetDepthBuffer(v bool) {
	wglrt.Set("depthBuffer", v)
}
func (wglrt *WebGLRenderTarget) DepthTexture() *textures.Texture {
	return &textures.Texture{Value: wglrt.Get("depthTexture")}
}
func (wglrt *WebGLRenderTarget) SetDepthTexture(v *textures.Texture) {
	wglrt.Set("depthTexture", v)
}
func (wglrt *WebGLRenderTarget) Format() js.Value {
	return wglrt.Get("format")
}
func (wglrt *WebGLRenderTarget) SetFormat(v js.Value) {
	wglrt.Set("format", v)
}
func (wglrt *WebGLRenderTarget) GenerateMipmaps() js.Value {
	return wglrt.Get("generateMipmaps")
}
func (wglrt *WebGLRenderTarget) SetGenerateMipmaps(v js.Value) {
	wglrt.Set("generateMipmaps", v)
}
func (wglrt *WebGLRenderTarget) Height() int {
	return wglrt.Get("height").Int()
}
func (wglrt *WebGLRenderTarget) SetHeight(v int) {
	wglrt.Set("height", v)
}
func (wglrt *WebGLRenderTarget) MagFilter() js.Value {
	return wglrt.Get("magFilter")
}
func (wglrt *WebGLRenderTarget) SetMagFilter(v js.Value) {
	wglrt.Set("magFilter", v)
}
func (wglrt *WebGLRenderTarget) MinFilter() js.Value {
	return wglrt.Get("minFilter")
}
func (wglrt *WebGLRenderTarget) SetMinFilter(v js.Value) {
	wglrt.Set("minFilter", v)
}
func (wglrt *WebGLRenderTarget) Offset() js.Value {
	return wglrt.Get("offset")
}
func (wglrt *WebGLRenderTarget) SetOffset(v js.Value) {
	wglrt.Set("offset", v)
}
func (wglrt *WebGLRenderTarget) Repeat() js.Value {
	return wglrt.Get("repeat")
}
func (wglrt *WebGLRenderTarget) SetRepeat(v js.Value) {
	wglrt.Set("repeat", v)
}
func (wglrt *WebGLRenderTarget) Scissor() *math.Vector4 {
	return &math.Vector4{Value: wglrt.Get("scissor")}
}
func (wglrt *WebGLRenderTarget) SetScissor(v *math.Vector4) {
	wglrt.Set("scissor", v)
}
func (wglrt *WebGLRenderTarget) ScissorTest() bool {
	return wglrt.Get("scissorTest").Bool()
}
func (wglrt *WebGLRenderTarget) SetScissorTest(v bool) {
	wglrt.Set("scissorTest", v)
}
func (wglrt *WebGLRenderTarget) StencilBuffer() bool {
	return wglrt.Get("stencilBuffer").Bool()
}
func (wglrt *WebGLRenderTarget) SetStencilBuffer(v bool) {
	wglrt.Set("stencilBuffer", v)
}
func (wglrt *WebGLRenderTarget) Texture() *textures.Texture {
	return &textures.Texture{Value: wglrt.Get("texture")}
}
func (wglrt *WebGLRenderTarget) SetTexture(v *textures.Texture) {
	wglrt.Set("texture", v)
}
func (wglrt *WebGLRenderTarget) Type() js.Value {
	return wglrt.Get("type")
}
func (wglrt *WebGLRenderTarget) SetType(v js.Value) {
	wglrt.Set("type", v)
}
func (wglrt *WebGLRenderTarget) Uuid() string {
	return wglrt.Get("uuid").String()
}
func (wglrt *WebGLRenderTarget) SetUuid(v string) {
	wglrt.Set("uuid", v)
}
func (wglrt *WebGLRenderTarget) Viewport() *math.Vector4 {
	return &math.Vector4{Value: wglrt.Get("viewport")}
}
func (wglrt *WebGLRenderTarget) SetViewport(v *math.Vector4) {
	wglrt.Set("viewport", v)
}
func (wglrt *WebGLRenderTarget) Width() int {
	return wglrt.Get("width").Int()
}
func (wglrt *WebGLRenderTarget) SetWidth(v int) {
	wglrt.Set("width", v)
}
func (wglrt *WebGLRenderTarget) WrapS() js.Value {
	return wglrt.Get("wrapS")
}
func (wglrt *WebGLRenderTarget) SetWrapS(v js.Value) {
	wglrt.Set("wrapS", v)
}
func (wglrt *WebGLRenderTarget) WrapT() js.Value {
	return wglrt.Get("wrapT")
}
func (wglrt *WebGLRenderTarget) SetWrapT(v js.Value) {
	wglrt.Set("wrapT", v)
}
func (wglrt *WebGLRenderTarget) AddEventListener(typ string, listener js.Value) {
	wglrt.Call("addEventListener", typ, listener)
}
func (wglrt *WebGLRenderTarget) Clone() this {
	return this(wglrt.Call("clone"))
}
func (wglrt *WebGLRenderTarget) Copy(source *WebGLRenderTarget) this {
	return this(wglrt.Call("copy", source))
}
func (wglrt *WebGLRenderTarget) DispatchEvent(event js.Value) {
	wglrt.Call("dispatchEvent", event)
}
func (wglrt *WebGLRenderTarget) Dispose() {
	wglrt.Call("dispose")
}
func (wglrt *WebGLRenderTarget) HasEventListener(typ string, listener js.Value) bool {
	return wglrt.Call("hasEventListener", typ, listener).Bool()
}
func (wglrt *WebGLRenderTarget) RemoveEventListener(typ string, listener js.Value) {
	wglrt.Call("removeEventListener", typ, listener)
}
func (wglrt *WebGLRenderTarget) SetSize(width int, height int) {
	wglrt.Call("setSize", width, height)
}

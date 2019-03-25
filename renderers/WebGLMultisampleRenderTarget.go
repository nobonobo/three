package renderers

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/textures"
)

type WebGLMultisampleRenderTarget struct {
	js.Value
}

func NewWebGLMultisampleRenderTarget(width int, height int, options WebGLRenderTargetOptions) *WebGLMultisampleRenderTarget {
	return &WebGLMultisampleRenderTarget{Value: get("WebGLMultisampleRenderTarget").New(width, height, options)}
}
func (wglmrt *WebGLMultisampleRenderTarget) Anisotropy() js.Value {
	return wglmrt.Get("anisotropy")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetAnisotropy(v js.Value) {
	wglmrt.Set("anisotropy", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) DepthBuffer() bool {
	return wglmrt.Get("depthBuffer").Bool()
}
func (wglmrt *WebGLMultisampleRenderTarget) SetDepthBuffer(v bool) {
	wglmrt.Set("depthBuffer", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) DepthTexture() *textures.Texture {
	return &textures.Texture{Value: wglmrt.Get("depthTexture")}
}
func (wglmrt *WebGLMultisampleRenderTarget) SetDepthTexture(v *textures.Texture) {
	wglmrt.Set("depthTexture", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Format() js.Value {
	return wglmrt.Get("format")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetFormat(v js.Value) {
	wglmrt.Set("format", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) GenerateMipmaps() js.Value {
	return wglmrt.Get("generateMipmaps")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetGenerateMipmaps(v js.Value) {
	wglmrt.Set("generateMipmaps", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Height() int {
	return wglmrt.Get("height").Int()
}
func (wglmrt *WebGLMultisampleRenderTarget) SetHeight(v int) {
	wglmrt.Set("height", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) MagFilter() js.Value {
	return wglmrt.Get("magFilter")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetMagFilter(v js.Value) {
	wglmrt.Set("magFilter", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) MinFilter() js.Value {
	return wglmrt.Get("minFilter")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetMinFilter(v js.Value) {
	wglmrt.Set("minFilter", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Offset() js.Value {
	return wglmrt.Get("offset")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetOffset(v js.Value) {
	wglmrt.Set("offset", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Repeat() js.Value {
	return wglmrt.Get("repeat")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetRepeat(v js.Value) {
	wglmrt.Set("repeat", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Scissor() *math.Vector4 {
	return &math.Vector4{Value: wglmrt.Get("scissor")}
}
func (wglmrt *WebGLMultisampleRenderTarget) SetScissor(v *math.Vector4) {
	wglmrt.Set("scissor", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) ScissorTest() bool {
	return wglmrt.Get("scissorTest").Bool()
}
func (wglmrt *WebGLMultisampleRenderTarget) SetScissorTest(v bool) {
	wglmrt.Set("scissorTest", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) StencilBuffer() bool {
	return wglmrt.Get("stencilBuffer").Bool()
}
func (wglmrt *WebGLMultisampleRenderTarget) SetStencilBuffer(v bool) {
	wglmrt.Set("stencilBuffer", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Texture() *textures.Texture {
	return &textures.Texture{Value: wglmrt.Get("texture")}
}
func (wglmrt *WebGLMultisampleRenderTarget) SetTexture(v *textures.Texture) {
	wglmrt.Set("texture", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Type() js.Value {
	return wglmrt.Get("type")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetType(v js.Value) {
	wglmrt.Set("type", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Uuid() string {
	return wglmrt.Get("uuid").String()
}
func (wglmrt *WebGLMultisampleRenderTarget) SetUuid(v string) {
	wglmrt.Set("uuid", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Viewport() *math.Vector4 {
	return &math.Vector4{Value: wglmrt.Get("viewport")}
}
func (wglmrt *WebGLMultisampleRenderTarget) SetViewport(v *math.Vector4) {
	wglmrt.Set("viewport", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) Width() int {
	return wglmrt.Get("width").Int()
}
func (wglmrt *WebGLMultisampleRenderTarget) SetWidth(v int) {
	wglmrt.Set("width", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) WrapS() js.Value {
	return wglmrt.Get("wrapS")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetWrapS(v js.Value) {
	wglmrt.Set("wrapS", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) WrapT() js.Value {
	return wglmrt.Get("wrapT")
}
func (wglmrt *WebGLMultisampleRenderTarget) SetWrapT(v js.Value) {
	wglmrt.Set("wrapT", v)
}
func (wglmrt *WebGLMultisampleRenderTarget) AddEventListener(typ string, listener js.Value) {
	wglmrt.Call("addEventListener", typ, listener)
}
func (wglmrt *WebGLMultisampleRenderTarget) Clone() this {
	return this(wglmrt.Call("clone"))
}
func (wglmrt *WebGLMultisampleRenderTarget) Copy(source *WebGLRenderTarget) this {
	return this(wglmrt.Call("copy", source))
}
func (wglmrt *WebGLMultisampleRenderTarget) DispatchEvent(event js.Value) {
	wglmrt.Call("dispatchEvent", event)
}
func (wglmrt *WebGLMultisampleRenderTarget) Dispose() {
	wglmrt.Call("dispose")
}
func (wglmrt *WebGLMultisampleRenderTarget) HasEventListener(typ string, listener js.Value) bool {
	return wglmrt.Call("hasEventListener", typ, listener).Bool()
}
func (wglmrt *WebGLMultisampleRenderTarget) RemoveEventListener(typ string, listener js.Value) {
	wglmrt.Call("removeEventListener", typ, listener)
}
func (wglmrt *WebGLMultisampleRenderTarget) SetSize(width int, height int) {
	wglmrt.Call("setSize", width, height)
}

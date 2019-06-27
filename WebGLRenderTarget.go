// Code generated from "renderers/WebGLRenderTarget.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type WebGLRenderTargetOptions interface {
}

// WebGLRenderTarget extend: [EventDispatcher]
type WebGLRenderTarget struct {
	js.Value
}

func NewWebGLRenderTarget(width float64, height float64, options WebGLRenderTargetOptions) *WebGLRenderTarget {
	return &WebGLRenderTarget{Value: get("WebGLRenderTarget").New(width, height, options)}
}
func (wglrt *WebGLRenderTarget) JSValue() js.Value {
	return wglrt.Value
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
func (wglrt *WebGLRenderTarget) DepthTexture() *Texture {
	return &Texture{Value: wglrt.Get("depthTexture")}
}
func (wglrt *WebGLRenderTarget) SetDepthTexture(v *Texture) {
	wglrt.Set("depthTexture", v.Value)
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
func (wglrt *WebGLRenderTarget) Height() float64 {
	return wglrt.Get("height").Float()
}
func (wglrt *WebGLRenderTarget) SetHeight(v float64) {
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
func (wglrt *WebGLRenderTarget) Scissor() *Vector4 {
	return &Vector4{Value: wglrt.Get("scissor")}
}
func (wglrt *WebGLRenderTarget) SetScissor(v *Vector4) {
	wglrt.Set("scissor", v.Value)
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
func (wglrt *WebGLRenderTarget) Texture() *Texture {
	return &Texture{Value: wglrt.Get("texture")}
}
func (wglrt *WebGLRenderTarget) SetTexture(v *Texture) {
	wglrt.Set("texture", v.Value)
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
func (wglrt *WebGLRenderTarget) Viewport() *Vector4 {
	return &Vector4{Value: wglrt.Get("viewport")}
}
func (wglrt *WebGLRenderTarget) SetViewport(v *Vector4) {
	wglrt.Set("viewport", v.Value)
}
func (wglrt *WebGLRenderTarget) Width() float64 {
	return wglrt.Get("width").Float()
}
func (wglrt *WebGLRenderTarget) SetWidth(v float64) {
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
func (wglrt *WebGLRenderTarget) Clone() *WebGLRenderTarget {
	return &WebGLRenderTarget{Value: wglrt.Call("clone")}
}
func (wglrt *WebGLRenderTarget) Copy(source *WebGLRenderTarget) *WebGLRenderTarget {
	return &WebGLRenderTarget{Value: wglrt.Call("copy", source)}
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
func (wglrt *WebGLRenderTarget) SetSize(width float64, height float64) {
	wglrt.Call("setSize", width, height)
}

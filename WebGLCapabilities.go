// Code generated from "renderers/webgl/WebGLCapabilities.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLCapabilitiesParameters interface {
}
type WebGLCapabilities struct {
	js.Value
}

func NewWebGLCapabilities(gl js.Value, extensions js.Value, parameters WebGLCapabilitiesParameters) *WebGLCapabilities {
	return &WebGLCapabilities{Value: get("WebGLCapabilities").New(gl, extensions, parameters)}
}
func (wglc *WebGLCapabilities) FloatFragmentTextures() js.Value {
	return wglc.Get("floatFragmentTextures")
}
func (wglc *WebGLCapabilities) SetFloatFragmentTextures(v js.Value) {
	wglc.Set("floatFragmentTextures", v)
}
func (wglc *WebGLCapabilities) FloatVertexTextures() js.Value {
	return wglc.Get("floatVertexTextures")
}
func (wglc *WebGLCapabilities) SetFloatVertexTextures(v js.Value) {
	wglc.Set("floatVertexTextures", v)
}
func (wglc *WebGLCapabilities) LogarithmicDepthBuffer() js.Value {
	return wglc.Get("logarithmicDepthBuffer")
}
func (wglc *WebGLCapabilities) SetLogarithmicDepthBuffer(v js.Value) {
	wglc.Set("logarithmicDepthBuffer", v)
}
func (wglc *WebGLCapabilities) MaxAttributes() js.Value {
	return wglc.Get("maxAttributes")
}
func (wglc *WebGLCapabilities) SetMaxAttributes(v js.Value) {
	wglc.Set("maxAttributes", v)
}
func (wglc *WebGLCapabilities) MaxCubemapSize() js.Value {
	return wglc.Get("maxCubemapSize")
}
func (wglc *WebGLCapabilities) SetMaxCubemapSize(v js.Value) {
	wglc.Set("maxCubemapSize", v)
}
func (wglc *WebGLCapabilities) MaxFragmentUniforms() js.Value {
	return wglc.Get("maxFragmentUniforms")
}
func (wglc *WebGLCapabilities) SetMaxFragmentUniforms(v js.Value) {
	wglc.Set("maxFragmentUniforms", v)
}
func (wglc *WebGLCapabilities) MaxTextureSize() js.Value {
	return wglc.Get("maxTextureSize")
}
func (wglc *WebGLCapabilities) SetMaxTextureSize(v js.Value) {
	wglc.Set("maxTextureSize", v)
}
func (wglc *WebGLCapabilities) MaxTextures() js.Value {
	return wglc.Get("maxTextures")
}
func (wglc *WebGLCapabilities) SetMaxTextures(v js.Value) {
	wglc.Set("maxTextures", v)
}
func (wglc *WebGLCapabilities) MaxVaryings() js.Value {
	return wglc.Get("maxVaryings")
}
func (wglc *WebGLCapabilities) SetMaxVaryings(v js.Value) {
	wglc.Set("maxVaryings", v)
}
func (wglc *WebGLCapabilities) MaxVertexTextures() js.Value {
	return wglc.Get("maxVertexTextures")
}
func (wglc *WebGLCapabilities) SetMaxVertexTextures(v js.Value) {
	wglc.Set("maxVertexTextures", v)
}
func (wglc *WebGLCapabilities) MaxVertexUniforms() js.Value {
	return wglc.Get("maxVertexUniforms")
}
func (wglc *WebGLCapabilities) SetMaxVertexUniforms(v js.Value) {
	wglc.Set("maxVertexUniforms", v)
}
func (wglc *WebGLCapabilities) Precision() js.Value {
	return wglc.Get("precision")
}
func (wglc *WebGLCapabilities) SetPrecision(v js.Value) {
	wglc.Set("precision", v)
}
func (wglc *WebGLCapabilities) VertexTextures() js.Value {
	return wglc.Get("vertexTextures")
}
func (wglc *WebGLCapabilities) SetVertexTextures(v js.Value) {
	wglc.Set("vertexTextures", v)
}
func (wglc *WebGLCapabilities) GetMaxAnisotropy() float64 {
	return wglc.Call("getMaxAnisotropy").Float()
}
func (wglc *WebGLCapabilities) GetMaxPrecision(precision string) string {
	return wglc.Call("getMaxPrecision", precision).String()
}

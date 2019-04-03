// Code generated from "renderers/webgl/WebGLTextures.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebGLTextures struct {
	js.Value
}

func NewWebGLTextures(gl js.Value, extensions js.Value, state js.Value, properties js.Value, capabilities js.Value, paramThreeToGL js.Value, info js.Value) *WebGLTextures {
	return &WebGLTextures{Value: get("WebGLTextures").New(gl, extensions, state, properties, capabilities, paramThreeToGL, info)}
}
func (wglt *WebGLTextures) SetTexture2D(texture js.Value, slot float64) {
	wglt.Call("setTexture2D", texture, slot)
}
func (wglt *WebGLTextures) SetTextureCube(texture js.Value, slot float64) {
	wglt.Call("setTextureCube", texture, slot)
}
func (wglt *WebGLTextures) SetTextureCubeDynamic(texture js.Value, slot float64) {
	wglt.Call("setTextureCubeDynamic", texture, slot)
}
func (wglt *WebGLTextures) SetupRenderTarget(renderTarget js.Value) {
	wglt.Call("setupRenderTarget", renderTarget)
}
func (wglt *WebGLTextures) UpdateRenderTargetMipmap(renderTarget js.Value) {
	wglt.Call("updateRenderTargetMipmap", renderTarget)
}

package lights

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers/webgl"
)

type LightShadow struct {
	js.Value
}

func NewLightShadow(camera *cameras.Camera) *LightShadow {
	return &LightShadow{Value: get("LightShadow").New(camera)}
}
func (ls *LightShadow) Bias() int {
	return ls.Get("bias").Int()
}
func (ls *LightShadow) SetBias(v int) {
	ls.Set("bias", v)
}
func (ls *LightShadow) Camera() *cameras.Camera {
	return &cameras.Camera{Value: ls.Get("camera")}
}
func (ls *LightShadow) SetCamera(v *cameras.Camera) {
	ls.Set("camera", v)
}
func (ls *LightShadow) Map() webgl.RenderTarget {
	return webgl.RenderTarget(ls.Get("map"))
}
func (ls *LightShadow) SetMap(v webgl.RenderTarget) {
	ls.Set("map", v)
}
func (ls *LightShadow) MapSize() *math.Vector2 {
	return &math.Vector2{Value: ls.Get("mapSize")}
}
func (ls *LightShadow) SetMapSize(v *math.Vector2) {
	ls.Set("mapSize", v)
}
func (ls *LightShadow) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ls.Get("matrix")}
}
func (ls *LightShadow) SetMatrix(v *math.Matrix4) {
	ls.Set("matrix", v)
}
func (ls *LightShadow) Radius() int {
	return ls.Get("radius").Int()
}
func (ls *LightShadow) SetRadius(v int) {
	ls.Set("radius", v)
}
func (ls *LightShadow) Clone(recursive bool) this {
	return this(ls.Call("clone", recursive))
}
func (ls *LightShadow) Copy(source *LightShadow) this {
	return this(ls.Call("copy", source))
}
func (ls *LightShadow) ToJSON() js.Value {
	return ls.Call("toJSON")
}

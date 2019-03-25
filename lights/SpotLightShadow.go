package lights

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers/webgl"
)

type SpotLightShadow struct {
	js.Value
}

func NewSpotLightShadow(camera *cameras.Camera) *SpotLightShadow {
	return &SpotLightShadow{Value: get("SpotLightShadow").New(camera)}
}
func (sls *SpotLightShadow) Bias() int {
	return sls.Get("bias").Int()
}
func (sls *SpotLightShadow) SetBias(v int) {
	sls.Set("bias", v)
}
func (sls *SpotLightShadow) Camera() *cameras.PerspectiveCamera {
	return &cameras.PerspectiveCamera{Value: sls.Get("camera")}
}
func (sls *SpotLightShadow) SetCamera(v *cameras.PerspectiveCamera) {
	sls.Set("camera", v)
}
func (sls *SpotLightShadow) Map() webgl.RenderTarget {
	return webgl.RenderTarget(sls.Get("map"))
}
func (sls *SpotLightShadow) SetMap(v webgl.RenderTarget) {
	sls.Set("map", v)
}
func (sls *SpotLightShadow) MapSize() *math.Vector2 {
	return &math.Vector2{Value: sls.Get("mapSize")}
}
func (sls *SpotLightShadow) SetMapSize(v *math.Vector2) {
	sls.Set("mapSize", v)
}
func (sls *SpotLightShadow) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: sls.Get("matrix")}
}
func (sls *SpotLightShadow) SetMatrix(v *math.Matrix4) {
	sls.Set("matrix", v)
}
func (sls *SpotLightShadow) Radius() int {
	return sls.Get("radius").Int()
}
func (sls *SpotLightShadow) SetRadius(v int) {
	sls.Set("radius", v)
}
func (sls *SpotLightShadow) Clone(recursive bool) this {
	return this(sls.Call("clone", recursive))
}
func (sls *SpotLightShadow) Copy(source *LightShadow) this {
	return this(sls.Call("copy", source))
}
func (sls *SpotLightShadow) ToJSON() js.Value {
	return sls.Call("toJSON")
}
func (sls *SpotLightShadow) Update(light *Light) {
	sls.Call("update", light)
}

package lights

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers/webgl"
)

type DirectionalLightShadow struct {
	js.Value
}

func NewDirectionalLightShadow(camera *cameras.Camera) *DirectionalLightShadow {
	return &DirectionalLightShadow{Value: get("DirectionalLightShadow").New(camera)}
}
func (dls *DirectionalLightShadow) Bias() int {
	return dls.Get("bias").Int()
}
func (dls *DirectionalLightShadow) SetBias(v int) {
	dls.Set("bias", v)
}
func (dls *DirectionalLightShadow) Camera() *cameras.OrthographicCamera {
	return &cameras.OrthographicCamera{Value: dls.Get("camera")}
}
func (dls *DirectionalLightShadow) SetCamera(v *cameras.OrthographicCamera) {
	dls.Set("camera", v)
}
func (dls *DirectionalLightShadow) Map() webgl.RenderTarget {
	return webgl.RenderTarget(dls.Get("map"))
}
func (dls *DirectionalLightShadow) SetMap(v webgl.RenderTarget) {
	dls.Set("map", v)
}
func (dls *DirectionalLightShadow) MapSize() *math.Vector2 {
	return &math.Vector2{Value: dls.Get("mapSize")}
}
func (dls *DirectionalLightShadow) SetMapSize(v *math.Vector2) {
	dls.Set("mapSize", v)
}
func (dls *DirectionalLightShadow) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: dls.Get("matrix")}
}
func (dls *DirectionalLightShadow) SetMatrix(v *math.Matrix4) {
	dls.Set("matrix", v)
}
func (dls *DirectionalLightShadow) Radius() int {
	return dls.Get("radius").Int()
}
func (dls *DirectionalLightShadow) SetRadius(v int) {
	dls.Set("radius", v)
}
func (dls *DirectionalLightShadow) Clone(recursive bool) this {
	return this(dls.Call("clone", recursive))
}
func (dls *DirectionalLightShadow) Copy(source *LightShadow) this {
	return this(dls.Call("copy", source))
}
func (dls *DirectionalLightShadow) ToJSON() js.Value {
	return dls.Call("toJSON")
}

package scenes

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type FogExp2 struct {
	js.Value
}

func NewFogExp2(hex int, density int) *FogExp2 {
	return &FogExp2{Value: get("FogExp2").New(hex, density)}
}
func (fe *FogExp2) Color() *math.Color {
	return &math.Color{Value: fe.Get("color")}
}
func (fe *FogExp2) SetColor(v *math.Color) {
	fe.Set("color", v)
}
func (fe *FogExp2) Density() int {
	return fe.Get("density").Int()
}
func (fe *FogExp2) SetDensity(v int) {
	fe.Set("density", v)
}
func (fe *FogExp2) Name() string {
	return fe.Get("name").String()
}
func (fe *FogExp2) SetName(v string) {
	fe.Set("name", v)
}
func (fe *FogExp2) Clone() this {
	return this(fe.Call("clone"))
}
func (fe *FogExp2) ToJSON() js.Value {
	return fe.Call("toJSON")
}

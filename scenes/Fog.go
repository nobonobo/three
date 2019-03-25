package scenes

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type IFog interface {
	Clone() this
	ToJSON() js.Value
}
type Fog struct {
	js.Value
}

func NewFog(hex int, near int, far int) *Fog {
	return &Fog{Value: get("Fog").New(hex, near, far)}
}
func (f *Fog) Color() *math.Color {
	return &math.Color{Value: f.Get("color")}
}
func (f *Fog) SetColor(v *math.Color) {
	f.Set("color", v)
}
func (f *Fog) Far() int {
	return f.Get("far").Int()
}
func (f *Fog) SetFar(v int) {
	f.Set("far", v)
}
func (f *Fog) Name() string {
	return f.Get("name").String()
}
func (f *Fog) SetName(v string) {
	f.Set("name", v)
}
func (f *Fog) Near() int {
	return f.Get("near").Int()
}
func (f *Fog) SetNear(v int) {
	f.Set("near", v)
}
func (f *Fog) Clone() this {
	return this(f.Call("clone"))
}
func (f *Fog) ToJSON() js.Value {
	return f.Call("toJSON")
}

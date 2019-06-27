// Code generated from "math/Color.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type HSL interface {
}

// Color extend: []
type Color struct {
	js.Value
}

func NewColor(color *Color) *Color {
	return &Color{Value: get("Color").New(color)}
}
func NewColor2(r float64, g float64, b float64) *Color {
	return &Color{Value: get("Color").New(r, g, b)}
}
func (cc *Color) JSValue() js.Value {
	return cc.Value
}
func (cc *Color) B() float64 {
	return cc.Get("b").Float()
}
func (cc *Color) SetB(v float64) {
	cc.Set("b", v)
}
func (cc *Color) G() float64 {
	return cc.Get("g").Float()
}
func (cc *Color) SetG(v float64) {
	cc.Set("g", v)
}
func (cc *Color) R() float64 {
	return cc.Get("r").Float()
}
func (cc *Color) SetR(v float64) {
	cc.Set("r", v)
}
func (cc *Color) Add(color *Color) *Color {
	return &Color{Value: cc.Call("add", color)}
}
func (cc *Color) AddColors(color1 *Color, color2 *Color) *Color {
	return &Color{Value: cc.Call("addColors", color1, color2)}
}
func (cc *Color) AddScalar(s float64) *Color {
	return &Color{Value: cc.Call("addScalar", s)}
}
func (cc *Color) Clone() *Color {
	return &Color{Value: cc.Call("clone")}
}
func (cc *Color) ConvertGammaToLinear() *Color {
	return &Color{Value: cc.Call("convertGammaToLinear")}
}
func (cc *Color) ConvertLinearToGamma() *Color {
	return &Color{Value: cc.Call("convertLinearToGamma")}
}
func (cc *Color) Copy(color *Color) *Color {
	return &Color{Value: cc.Call("copy", color)}
}
func (cc *Color) CopyGammaToLinear(color *Color, gammaFactor float64) *Color {
	return &Color{Value: cc.Call("copyGammaToLinear", color, gammaFactor)}
}
func (cc *Color) CopyLinearToGamma(color *Color, gammaFactor float64) *Color {
	return &Color{Value: cc.Call("copyLinearToGamma", color, gammaFactor)}
}
func (cc *Color) Equals(color *Color) bool {
	return cc.Call("equals", color).Bool()
}
func (cc *Color) FromArray(rgb js.Value, offset int) *Color {
	return &Color{Value: cc.Call("fromArray", rgb, offset)}
}
func (cc *Color) GetHSL(target HSL) HSL {
	return HSL(cc.Call("getHSL", target))
}
func (cc *Color) GetHex() int {
	return cc.Call("getHex").Int()
}
func (cc *Color) GetHexString() string {
	return cc.Call("getHexString").String()
}
func (cc *Color) GetStyle() string {
	return cc.Call("getStyle").String()
}
func (cc *Color) Lerp(color *Color, alpha float64) *Color {
	return &Color{Value: cc.Call("lerp", color, alpha)}
}
func (cc *Color) LerpHSL(color *Color, alpha float64) *Color {
	return &Color{Value: cc.Call("lerpHSL", color, alpha)}
}
func (cc *Color) Multiply(color *Color) *Color {
	return &Color{Value: cc.Call("multiply", color)}
}
func (cc *Color) MultiplyScalar(s float64) *Color {
	return &Color{Value: cc.Call("multiplyScalar", s)}
}
func (cc *Color) OffsetHSL(h float64, s float64, l float64) *Color {
	return &Color{Value: cc.Call("offsetHSL", h, s, l)}
}
func (cc *Color) Set2(color *Color) *Color {
	return &Color{Value: cc.Call("set", color)}
}
func (cc *Color) Set3(color int) *Color {
	return &Color{Value: cc.Call("set", color)}
}
func (cc *Color) Set4(color string) *Color {
	return &Color{Value: cc.Call("set", color)}
}
func (cc *Color) SetHSL(h float64, s float64, l float64) *Color {
	return &Color{Value: cc.Call("setHSL", h, s, l)}
}
func (cc *Color) SetHex(hex int) *Color {
	return &Color{Value: cc.Call("setHex", hex)}
}
func (cc *Color) SetRGB(r float64, g float64, b float64) *Color {
	return &Color{Value: cc.Call("setRGB", r, g, b)}
}
func (cc *Color) SetScalar(scalar float64) *Color {
	return &Color{Value: cc.Call("setScalar", scalar)}
}
func (cc *Color) SetStyle(style string) *Color {
	return &Color{Value: cc.Call("setStyle", style)}
}
func (cc *Color) Sub(color *Color) *Color {
	return &Color{Value: cc.Call("sub", color)}
}
func (cc *Color) ToArray(array js.Value, offset int) js.Value {
	return cc.Call("toArray", array, offset)
}
func (cc *Color) ToArray2(xyz js.Value, offset int) js.Value {
	return cc.Call("toArray", xyz, offset)
}

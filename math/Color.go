package math

type HSL interface {
}
type Color struct {
	js.Value
}

func NewColor(color Color) *Color {
	return &Color{Value: get("Color").New(color)}
}
func NewColor2(r int, g int, b int) *Color {
	return &Color{Value: get("Color").New(r, g, b)}
}
func (c *Color) B() int {
	return c.Get("b").Int()
}
func (c *Color) SetB(v int) {
	c.Set("b", v)
}
func (c *Color) G() int {
	return c.Get("g").Int()
}
func (c *Color) SetG(v int) {
	c.Set("g", v)
}
func (c *Color) R() int {
	return c.Get("r").Int()
}
func (c *Color) SetR(v int) {
	c.Set("r", v)
}
func (c *Color) Add(color *Color) this {
	return this(c.Call("add", color))
}
func (c *Color) AddColors(color1 *Color, color2 *Color) this {
	return this(c.Call("addColors", color1, color2))
}
func (c *Color) AddScalar(s int) this {
	return this(c.Call("addScalar", s))
}
func (c *Color) Clone() this {
	return this(c.Call("clone"))
}
func (c *Color) ConvertGammaToLinear() *Color {
	return &Color{Value: c.Call("convertGammaToLinear")}
}
func (c *Color) ConvertLinearToGamma() *Color {
	return &Color{Value: c.Call("convertLinearToGamma")}
}
func (c *Color) Copy(color *Color) this {
	return this(c.Call("copy", color))
}
func (c *Color) CopyGammaToLinear(color *Color, gammaFactor int) *Color {
	return &Color{Value: c.Call("copyGammaToLinear", color, gammaFactor)}
}
func (c *Color) CopyLinearToGamma(color *Color, gammaFactor int) *Color {
	return &Color{Value: c.Call("copyLinearToGamma", color, gammaFactor)}
}
func (c *Color) Equals(color *Color) bool {
	return c.Call("equals", color).Bool()
}
func (c *Color) FromArray(rgb []int, offset int) this {
	return this(c.Call("fromArray", rgb, offset))
}
func (c *Color) GetHSL(target HSL) HSL {
	return HSL(c.Call("getHSL", target))
}
func (c *Color) GetHex() int {
	return c.Call("getHex").Int()
}
func (c *Color) GetHexString() string {
	return c.Call("getHexString").String()
}
func (c *Color) GetStyle() string {
	return c.Call("getStyle").String()
}
func (c *Color) Lerp(color *Color, alpha int) this {
	return this(c.Call("lerp", color, alpha))
}
func (c *Color) LerpHSL(color *Color, alpha int) this {
	return this(c.Call("lerpHSL", color, alpha))
}
func (c *Color) Multiply(color *Color) this {
	return this(c.Call("multiply", color))
}
func (c *Color) MultiplyScalar(s int) this {
	return this(c.Call("multiplyScalar", s))
}
func (c *Color) OffsetHSL(h int, s int, l int) this {
	return this(c.Call("offsetHSL", h, s, l))
}
func (c *Color) Set(color *Color) *Color {
	return &Color{Value: c.Call("set", color)}
}
func (c *Color) Set2(color int) *Color {
	return &Color{Value: c.Call("set", color)}
}
func (c *Color) Set3(color string) *Color {
	return &Color{Value: c.Call("set", color)}
}
func (c *Color) SetHSL(h int, s int, l int) *Color {
	return &Color{Value: c.Call("setHSL", h, s, l)}
}
func (c *Color) SetHex(hex int) *Color {
	return &Color{Value: c.Call("setHex", hex)}
}
func (c *Color) SetRGB(r int, g int, b int) *Color {
	return &Color{Value: c.Call("setRGB", r, g, b)}
}
func (c *Color) SetScalar(scalar int) *Color {
	return &Color{Value: c.Call("setScalar", scalar)}
}
func (c *Color) SetStyle(style string) *Color {
	return &Color{Value: c.Call("setStyle", style)}
}
func (c *Color) Sub(color *Color) this {
	return this(c.Call("sub", color))
}
func (c *Color) ToArray(array []int, offset int) []int {
	return []int(c.Call("toArray", array, offset))
}
func (c *Color) ToArray2(xyz *ArrayLike, offset int) *ArrayLike {
	return &ArrayLike{Value: c.Call("toArray", xyz, offset)}
}

package math

type SplineControlPoint interface {
}
type Triangle struct {
	js.Value
}

func NewTriangle(a *Vector3, b *Vector3, c *Vector3) *Triangle {
	return &Triangle{Value: get("Triangle").New(a, b, c)}
}
func (t *Triangle) A() *Vector3 {
	return &Vector3{Value: t.Get("a")}
}
func (t *Triangle) SetA(v *Vector3) {
	t.Set("a", v)
}
func (t *Triangle) B() *Vector3 {
	return &Vector3{Value: t.Get("b")}
}
func (t *Triangle) SetB(v *Vector3) {
	t.Set("b", v)
}
func (t *Triangle) C() *Vector3 {
	return &Vector3{Value: t.Get("c")}
}
func (t *Triangle) SetC(v *Vector3) {
	t.Set("c", v)
}
func (t *Triangle) Clone() this {
	return this(t.Call("clone"))
}
func (t *Triangle) ClosestPointToPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: t.Call("closestPointToPoint", point, target)}
}
func (t *Triangle) ContainsPoint(point *Vector3) bool {
	return t.Call("containsPoint", point).Bool()
}
func (t *Triangle) Copy(triangle *Triangle) this {
	return this(t.Call("copy", triangle))
}
func (t *Triangle) Equals(triangle *Triangle) bool {
	return t.Call("equals", triangle).Bool()
}
func (t *Triangle) GetArea() int {
	return t.Call("getArea").Int()
}
func (t *Triangle) GetBarycoord(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: t.Call("getBarycoord", point, target)}
}
func (t *Triangle) GetMidpoint(target *Vector3) *Vector3 {
	return &Vector3{Value: t.Call("getMidpoint", target)}
}
func (t *Triangle) GetNormal(target *Vector3) *Vector3 {
	return &Vector3{Value: t.Call("getNormal", target)}
}
func (t *Triangle) GetPlane(target *Vector3) *Plane {
	return &Plane{Value: t.Call("getPlane", target)}
}
func (t *Triangle) Set(a *Vector3, b *Vector3, c *Vector3) *Triangle {
	return &Triangle{Value: t.Call("set", a, b, c)}
}
func (t *Triangle) SetFromPointsAndIndices(points []Vector3, i0 int, i1 int, i2 int) *Triangle {
	return &Triangle{Value: t.Call("setFromPointsAndIndices", points, i0, i1, i2)}
}
func (t *Triangle) ContainsPoint2(point *Vector3, a *Vector3, b *Vector3, c *Vector3) bool {
	return t.Call("containsPoint", point, a, b, c).Bool()
}
func (t *Triangle) GetBarycoord2(point *Vector3, a *Vector3, b *Vector3, c *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: t.Call("getBarycoord", point, a, b, c, target)}
}
func (t *Triangle) GetNormal2(a *Vector3, b *Vector3, c *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: t.Call("getNormal", a, b, c, target)}
}

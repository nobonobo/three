package math

type Cylindrical struct {
	js.Value
}

func NewCylindrical(radius int, theta int, y int) *Cylindrical {
	return &Cylindrical{Value: get("Cylindrical").New(radius, theta, y)}
}
func (c *Cylindrical) Radius() int {
	return c.Get("radius").Int()
}
func (c *Cylindrical) SetRadius(v int) {
	c.Set("radius", v)
}
func (c *Cylindrical) Theta() int {
	return c.Get("theta").Int()
}
func (c *Cylindrical) SetTheta(v int) {
	c.Set("theta", v)
}
func (c *Cylindrical) Y() int {
	return c.Get("y").Int()
}
func (c *Cylindrical) SetY(v int) {
	c.Set("y", v)
}
func (c *Cylindrical) Clone() this {
	return this(c.Call("clone"))
}
func (c *Cylindrical) Copy(other *Cylindrical) this {
	return this(c.Call("copy", other))
}
func (c *Cylindrical) Set(radius int, theta int, y int) this {
	return this(c.Call("set", radius, theta, y))
}
func (c *Cylindrical) SetFromVector3(vec3 *Vector3) this {
	return this(c.Call("setFromVector3", vec3))
}

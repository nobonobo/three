package math

type Spherical struct {
	js.Value
}

func NewSpherical(radius int, phi int, theta int) *Spherical {
	return &Spherical{Value: get("Spherical").New(radius, phi, theta)}
}
func (s *Spherical) Phi() int {
	return s.Get("phi").Int()
}
func (s *Spherical) SetPhi(v int) {
	s.Set("phi", v)
}
func (s *Spherical) Radius() int {
	return s.Get("radius").Int()
}
func (s *Spherical) SetRadius(v int) {
	s.Set("radius", v)
}
func (s *Spherical) Theta() int {
	return s.Get("theta").Int()
}
func (s *Spherical) SetTheta(v int) {
	s.Set("theta", v)
}
func (s *Spherical) Clone() this {
	return this(s.Call("clone"))
}
func (s *Spherical) Copy(other *Spherical) this {
	return this(s.Call("copy", other))
}
func (s *Spherical) MakeSafe() {
	s.Call("makeSafe")
}
func (s *Spherical) Set(radius int, phi int, theta int) *Spherical {
	return &Spherical{Value: s.Call("set", radius, phi, theta)}
}
func (s *Spherical) SetFromVector3(vec3 *Vector3) *Spherical {
	return &Spherical{Value: s.Call("setFromVector3", vec3)}
}

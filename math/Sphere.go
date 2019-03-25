package math

type Sphere struct {
	js.Value
}

func NewSphere(center *Vector3, radius int) *Sphere {
	return &Sphere{Value: get("Sphere").New(center, radius)}
}
func (s *Sphere) Center() *Vector3 {
	return &Vector3{Value: s.Get("center")}
}
func (s *Sphere) SetCenter(v *Vector3) {
	s.Set("center", v)
}
func (s *Sphere) Radius() int {
	return s.Get("radius").Int()
}
func (s *Sphere) SetRadius(v int) {
	s.Set("radius", v)
}
func (s *Sphere) ApplyMatrix4(matrix *Matrix4) *Sphere {
	return &Sphere{Value: s.Call("applyMatrix4", matrix)}
}
func (s *Sphere) ClampPoint(point *Vector3, target *Vector3) *Vector3 {
	return &Vector3{Value: s.Call("clampPoint", point, target)}
}
func (s *Sphere) Clone() this {
	return this(s.Call("clone"))
}
func (s *Sphere) ContainsPoint(point *Vector3) bool {
	return s.Call("containsPoint", point).Bool()
}
func (s *Sphere) Copy(sphere *Sphere) this {
	return this(s.Call("copy", sphere))
}
func (s *Sphere) DistanceToPoint(point *Vector3) int {
	return s.Call("distanceToPoint", point).Int()
}
func (s *Sphere) Empty() bool {
	return s.Call("empty").Bool()
}
func (s *Sphere) Equals(sphere *Sphere) bool {
	return s.Call("equals", sphere).Bool()
}
func (s *Sphere) GetBoundingBox(target *Box3) *Box3 {
	return &Box3{Value: s.Call("getBoundingBox", target)}
}
func (s *Sphere) IntersectsBox(box *Box3) bool {
	return s.Call("intersectsBox", box).Bool()
}
func (s *Sphere) IntersectsPlane(plane *Plane) bool {
	return s.Call("intersectsPlane", plane).Bool()
}
func (s *Sphere) IntersectsSphere(sphere *Sphere) bool {
	return s.Call("intersectsSphere", sphere).Bool()
}
func (s *Sphere) Set(center *Vector3, radius int) *Sphere {
	return &Sphere{Value: s.Call("set", center, radius)}
}
func (s *Sphere) SetFromPoints(points []Vector3, optionalCenter *Vector3) *Sphere {
	return &Sphere{Value: s.Call("setFromPoints", points, optionalCenter)}
}
func (s *Sphere) Translate(offset *Vector3) *Sphere {
	return &Sphere{Value: s.Call("translate", offset)}
}

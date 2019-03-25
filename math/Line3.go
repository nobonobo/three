package math

type Line3 struct {
	js.Value
}

func NewLine3(start *Vector3, end *Vector3) *Line3 {
	return &Line3{Value: get("Line3").New(start, end)}
}
func (l *Line3) End() *Vector3 {
	return &Vector3{Value: l.Get("end")}
}
func (l *Line3) SetEnd(v *Vector3) {
	l.Set("end", v)
}
func (l *Line3) Start() *Vector3 {
	return &Vector3{Value: l.Get("start")}
}
func (l *Line3) SetStart(v *Vector3) {
	l.Set("start", v)
}
func (l *Line3) ApplyMatrix4(matrix *Matrix4) *Line3 {
	return &Line3{Value: l.Call("applyMatrix4", matrix)}
}
func (l *Line3) At(t int, target *Vector3) *Vector3 {
	return &Vector3{Value: l.Call("at", t, target)}
}
func (l *Line3) Clone() this {
	return this(l.Call("clone"))
}
func (l *Line3) ClosestPointToPoint(point *Vector3, clampToLine bool, target *Vector3) *Vector3 {
	return &Vector3{Value: l.Call("closestPointToPoint", point, clampToLine, target)}
}
func (l *Line3) ClosestPointToPointParameter(point *Vector3, clampToLine bool) int {
	return l.Call("closestPointToPointParameter", point, clampToLine).Int()
}
func (l *Line3) Copy(line *Line3) this {
	return this(l.Call("copy", line))
}
func (l *Line3) Delta(target *Vector3) *Vector3 {
	return &Vector3{Value: l.Call("delta", target)}
}
func (l *Line3) Distance() int {
	return l.Call("distance").Int()
}
func (l *Line3) DistanceSq() int {
	return l.Call("distanceSq").Int()
}
func (l *Line3) Equals(line *Line3) bool {
	return l.Call("equals", line).Bool()
}
func (l *Line3) GetCenter(target *Vector3) *Vector3 {
	return &Vector3{Value: l.Call("getCenter", target)}
}
func (l *Line3) Set(start *Vector3, end *Vector3) *Line3 {
	return &Line3{Value: l.Call("set", start, end)}
}

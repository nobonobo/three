package curveutils

func Interpolate(p0 int, p1 int, p2 int, p3 int, t int) int {
	return _Global.Call("interpolate", p0, p1, p2, p3, t).Int()
}
func TangentCubicBezier(t int, p0 int, p1 int, p2 int, p3 int) int {
	return _Global.Call("tangentCubicBezier", t, p0, p1, p2, p3).Int()
}
func TangentQuadraticBezier(t int, p0 int, p1 int, p2 int) int {
	return _Global.Call("tangentQuadraticBezier", t, p0, p1, p2).Int()
}
func TangentSpline(t int, p0 int, p1 int, p2 int, p3 int) int {
	return _Global.Call("tangentSpline", t, p0, p1, p2, p3).Int()
}

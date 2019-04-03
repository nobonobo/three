// Code generated from "CurveUtils/curveutils.d.ts"; DO NOT EDIT.

package CurveUtils

func Interpolate(p0 float64, p1 float64, p2 float64, p3 float64, t float64) float64 {
	return _Global.Call("interpolate", p0, p1, p2, p3, t).Float()
}
func TangentCubicBezier(t float64, p0 float64, p1 float64, p2 float64, p3 float64) float64 {
	return _Global.Call("tangentCubicBezier", t, p0, p1, p2, p3).Float()
}
func TangentQuadraticBezier(t float64, p0 float64, p1 float64, p2 float64) float64 {
	return _Global.Call("tangentQuadraticBezier", t, p0, p1, p2).Float()
}
func TangentSpline(t float64, p0 float64, p1 float64, p2 float64, p3 float64) float64 {
	return _Global.Call("tangentSpline", t, p0, p1, p2, p3).Float()
}

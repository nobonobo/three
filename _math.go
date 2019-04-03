// Code generated from "_Math/_math.d.ts"; DO NOT EDIT.

package three

func DEG2RAD() float64 {
	return get("DEG2RAD").Float()
}
func SetDEG2RAD(v float64) {
	set("DEG2RAD", v)
}
func RAD2DEG() float64 {
	return get("RAD2DEG").Float()
}
func SetRAD2DEG(v float64) {
	set("RAD2DEG", v)
}
func CeilPowerOfTwo(value float64) float64 {
	return _Global.Call("ceilPowerOfTwo", value).Float()
}
func Clamp(value float64, min float64, max float64) float64 {
	return _Global.Call("clamp", value, min, max).Float()
}
func DegToRad(degrees float64) float64 {
	return _Global.Call("degToRad", degrees).Float()
}
func EuclideanModulo(n float64, m float64) float64 {
	return _Global.Call("euclideanModulo", n, m).Float()
}
func FloorPowerOfTwo(value float64) float64 {
	return _Global.Call("floorPowerOfTwo", value).Float()
}
func GenerateUUID() string {
	return _Global.Call("generateUUID").String()
}
func IsPowerOfTwo(value float64) bool {
	return _Global.Call("isPowerOfTwo", value).Bool()
}
func Lerp(x float64, y float64, t float64) float64 {
	return _Global.Call("lerp", x, y, t).Float()
}
func MapLinear(x float64, a1 float64, a2 float64, b1 float64, b2 float64) float64 {
	return _Global.Call("mapLinear", x, a1, a2, b1, b2).Float()
}
func NearestPowerOfTwo(value float64) float64 {
	return _Global.Call("nearestPowerOfTwo", value).Float()
}
func NextPowerOfTwo(value float64) float64 {
	return _Global.Call("nextPowerOfTwo", value).Float()
}
func RadToDeg(radians float64) float64 {
	return _Global.Call("radToDeg", radians).Float()
}
func RandFloat(low float64, high float64) float64 {
	return _Global.Call("randFloat", low, high).Float()
}
func RandFloatSpread(rng float64) float64 {
	return _Global.Call("randFloatSpread", rng).Float()
}
func RandInt(low float64, high float64) float64 {
	return _Global.Call("randInt", low, high).Float()
}
func Random16() float64 {
	return _Global.Call("random16").Float()
}
func Smootherstep(x float64, min float64, max float64) float64 {
	return _Global.Call("smootherstep", x, min, max).Float()
}
func Smoothstep(x float64, min float64, max float64) float64 {
	return _Global.Call("smoothstep", x, min, max).Float()
}

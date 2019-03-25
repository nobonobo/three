package _math

func DEG2RAD() int {
	return get("DEG2RAD").Int()
}
func SetDEG2RAD(v int) {
	set("DEG2RAD", v)
}
func RAD2DEG() int {
	return get("RAD2DEG").Int()
}
func SetRAD2DEG(v int) {
	set("RAD2DEG", v)
}
func CeilPowerOfTwo(value int) int {
	return _Global.Call("ceilPowerOfTwo", value).Int()
}
func Clamp(value int, min int, max int) int {
	return _Global.Call("clamp", value, min, max).Int()
}
func DegToRad(degrees int) int {
	return _Global.Call("degToRad", degrees).Int()
}
func EuclideanModulo(n int, m int) int {
	return _Global.Call("euclideanModulo", n, m).Int()
}
func FloorPowerOfTwo(value int) int {
	return _Global.Call("floorPowerOfTwo", value).Int()
}
func GenerateUUID() string {
	return _Global.Call("generateUUID").String()
}
func IsPowerOfTwo(value int) bool {
	return _Global.Call("isPowerOfTwo", value).Bool()
}
func Lerp(x int, y int, t int) int {
	return _Global.Call("lerp", x, y, t).Int()
}
func MapLinear(x int, a1 int, a2 int, b1 int, b2 int) int {
	return _Global.Call("mapLinear", x, a1, a2, b1, b2).Int()
}
func NearestPowerOfTwo(value int) int {
	return _Global.Call("nearestPowerOfTwo", value).Int()
}
func NextPowerOfTwo(value int) int {
	return _Global.Call("nextPowerOfTwo", value).Int()
}
func RadToDeg(radians int) int {
	return _Global.Call("radToDeg", radians).Int()
}
func RandFloat(low int, high int) int {
	return _Global.Call("randFloat", low, high).Int()
}
func RandFloatSpread(rng int) int {
	return _Global.Call("randFloatSpread", rng).Int()
}
func RandInt(low int, high int) int {
	return _Global.Call("randInt", low, high).Int()
}
func Random16() int {
	return _Global.Call("random16").Int()
}
func Smootherstep(x int, min int, max int) int {
	return _Global.Call("smootherstep", x, min, max).Int()
}
func Smoothstep(x int, min int, max int) int {
	return _Global.Call("smoothstep", x, min, max).Int()
}

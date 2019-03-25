package animationutils

import (
	"github.com/gopherjs/gopherwasm/js"
)

func ArraySlice(array js.Value, from int, to int) js.Value {
	return _Global.Call("arraySlice", array, from, to)
}
func ConvertArray(array js.Value, typ js.Value, forceClone bool) js.Value {
	return _Global.Call("convertArray", array, typ, forceClone)
}
func FlattenJSON(jsonKeys []string, times []js.Value, values []js.Value, valuePropertyName string) {
	_Global.Call("flattenJSON", jsonKeys, times, values, valuePropertyName)
}
func GetKeyFrameOrder(times int) []int {
	return []int(_Global.Call("getKeyFrameOrder", times))
}
func IsTypedArray(object js.Value) bool {
	return _Global.Call("isTypedArray", object).Bool()
}
func SortedArray(values []js.Value, stride int, order []int) []js.Value {
	return []js.Value(_Global.Call("sortedArray", values, stride, order))
}

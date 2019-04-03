// Code generated from "AnimationUtils/animationutils.d.ts"; DO NOT EDIT.

package AnimationUtils

import (
	"github.com/gopherjs/gopherwasm/js"
)

func ArraySlice(array js.Value, from float64, to float64) js.Value {
	return _Global.Call("arraySlice", array, from, to)
}
func ConvertArray(array js.Value, typ js.Value, forceClone bool) js.Value {
	return _Global.Call("convertArray", array, typ, forceClone)
}
func FlattenJSON(jsonKeys js.Value, times js.Value, values js.Value, valuePropertyName string) {
	_Global.Call("flattenJSON", jsonKeys, times, values, valuePropertyName)
}
func GetKeyFrameOrder(times float64) js.Value {
	return _Global.Call("getKeyFrameOrder", times)
}
func IsTypedArray(object js.Value) bool {
	return _Global.Call("isTypedArray", object).Bool()
}
func SortedArray(values js.Value, stride float64, order js.Value) js.Value {
	return _Global.Call("sortedArray", values, stride, order)
}

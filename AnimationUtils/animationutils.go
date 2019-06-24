// Code generated from "AnimationUtils/animationutils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package AnimationUtils

import (
	"syscall/js"
)

func ArraySlice(array js.Value, from float64, to float64) js.Value {
	return global().Call("arraySlice", array, from, to)
}
func ConvertArray(array js.Value, typ js.Value, forceClone bool) js.Value {
	return global().Call("convertArray", array, typ, forceClone)
}
func FlattenJSON(jsonKeys js.Value, times js.Value, values js.Value, valuePropertyName string) {
	global().Call("flattenJSON", jsonKeys, times, values, valuePropertyName)
}
func GetKeyFrameOrder(times float64) js.Value {
	return global().Call("getKeyFrameOrder", times)
}
func IsTypedArray(object js.Value) bool {
	return global().Call("isTypedArray", object).Bool()
}
func SortedArray(values js.Value, stride float64, order js.Value) js.Value {
	return global().Call("sortedArray", values, stride, order)
}

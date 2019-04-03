// Code generated from "UniformsUtils/uniformsutils.d.ts"; DO NOT EDIT.

package UniformsUtils

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("UniformsUtils")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v interface{}) {
	_Global.Set(key, v)
}

package uniformsutils

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("shaders").Get("UniformsUtils")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v js.Value) {
	return _Global.Set(key, v)
}

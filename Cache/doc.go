// Code generated from "Cache/cache.d.ts"; DO NOT EDIT.

package Cache

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("Cache")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v interface{}) {
	_Global.Set(key, v)
}

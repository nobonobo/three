// Code generated from "ImageUtils/imageutils.d.ts"; DO NOT EDIT.

package ImageUtils

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("ImageUtils")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v interface{}) {
	_Global.Set(key, v)
}

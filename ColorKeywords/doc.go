// Code generated from "ColorKeywords/colorkeywords.d.ts"; DO NOT EDIT.

package ColorKeywords

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("ColorKeywords")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v interface{}) {
	_Global.Set(key, v)
}
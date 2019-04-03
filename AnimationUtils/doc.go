// Code generated from "AnimationUtils/animationutils.d.ts"; DO NOT EDIT.

package AnimationUtils

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("AnimationUtils")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v interface{}) {
	_Global.Set(key, v)
}

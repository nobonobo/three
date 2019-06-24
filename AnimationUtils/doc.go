// Code generated from "AnimationUtils/animationutils.d.ts"; DO NOT EDIT.

package AnimationUtils

import (
	"syscall/js"
)

func global() js.Value {
	return js.Global().Get("THREE")
}

func get(key string) js.Value {
	return global().Get(key)
}

func set(key string, v interface{}) {
	global().Set(key, v)
}

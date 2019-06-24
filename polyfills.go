// Code generated from "polyfills.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func Error(message js.Value, optionalParams js.Value) {
	global().Call("error", message, optionalParams)
}
func Log(message js.Value, optionalParams js.Value) {
	global().Call("log", message, optionalParams)
}
func Warn(message js.Value, optionalParams js.Value) {
	global().Call("warn", message, optionalParams)
}

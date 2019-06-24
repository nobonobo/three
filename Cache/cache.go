// Code generated from "Cache/cache.d.ts"; DO NOT EDIT.

// +build go1.12 js

package Cache

import (
	"syscall/js"
)

func Enabled() bool {
	return get("enabled").Bool()
}
func SetEnabled(v bool) {
	set("enabled", v)
}
func Files() js.Value {
	return get("files")
}
func SetFiles(v js.Value) {
	set("files", v)
}
func Add(key string, file js.Value) {
	global().Call("add", key, file)
}
func Clear() {
	global().Call("clear")
}
func Get(key string) js.Value {
	return global().Call("get", key)
}
func Remove(key string) {
	global().Call("remove", key)
}

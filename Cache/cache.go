// Code generated from "Cache/cache.d.ts"; DO NOT EDIT.

package Cache

import (
	"github.com/gopherjs/gopherwasm/js"
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
	_Global.Call("add", key, file)
}
func Clear() {
	_Global.Call("clear")
}
func Get(key string) js.Value {
	return _Global.Call("get", key)
}
func Remove(key string) {
	_Global.Call("remove", key)
}

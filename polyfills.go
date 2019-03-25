package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

func Error(message js.Value, optionalParams []js.Value) {
	_Global.Call("error", message, optionalParams)
}
func Log(message js.Value, optionalParams []js.Value) {
	_Global.Call("log", message, optionalParams)
}
func Warn(message js.Value, optionalParams []js.Value) {
	_Global.Call("warn", message, optionalParams)
}

// Code generated from "SceneUtils/sceneutils.d.ts"; DO NOT EDIT.

package SceneUtils

import "github.com/gopherjs/gopherwasm/js"

var _Global = js.Global().Get("THREE").Get("SceneUtils")

func get(key string) js.Value {
	return _Global.Get(key)
}

func set(key string, v interface{}) {
	_Global.Set(key, v)
}

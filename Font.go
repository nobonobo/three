// Code generated from "extras/core/Font.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Font extend: []
type Font struct {
	js.Value
}

func NewFont(jsondata js.Value) *Font {
	return &Font{Value: get("Font").New(jsondata)}
}
func (ff *Font) JSValue() js.Value {
	return ff.Value
}
func (ff *Font) Data() string {
	return ff.Get("data").String()
}
func (ff *Font) SetData(v string) {
	ff.Set("data", v)
}
func (ff *Font) GenerateShapes(text string, size float64, divisions float64) js.Value {
	return ff.Call("generateShapes", text, size, divisions)
}

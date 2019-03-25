package core

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Font struct {
	js.Value
}

func NewFont(jsondata js.Value) *Font {
	return &Font{Value: get("Font").New(jsondata)}
}
func (f *Font) Data() string {
	return f.Get("data").String()
}
func (f *Font) SetData(v string) {
	f.Set("data", v)
}
func (f *Font) GenerateShapes(text string, size int, divisions int) []js.Value {
	return []js.Value(f.Call("generateShapes", text, size, divisions))
}

// Code generated from "loaders/LoaderUtils.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type LoaderUtils struct {
	js.Value
}

func (lu *LoaderUtils) DecodeText(array js.Value) string {
	return lu.Call("decodeText", array).String()
}
func (lu *LoaderUtils) ExtractUrlBase(url string) string {
	return lu.Call("extractUrlBase", url).String()
}

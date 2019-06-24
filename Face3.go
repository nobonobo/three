// Code generated from "core/Face3.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Event interface {
}
type Face3 struct {
	js.Value
}

func NewFace3(a float64, b float64, c float64, normal *Vector3, color *Color, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, normal, color, materialIndex)}
}
func NewFace32(a float64, b float64, c float64, normal *Vector3, vertexColors js.Value, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, normal, vertexColors, materialIndex)}
}
func NewFace33(a float64, b float64, c float64, vertexNormals js.Value, color *Color, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, vertexNormals, color, materialIndex)}
}
func NewFace34(a float64, b float64, c float64, vertexNormals js.Value, vertexColors js.Value, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, vertexNormals, vertexColors, materialIndex)}
}
func (ff *Face3) A() float64 {
	return ff.Get("a").Float()
}
func (ff *Face3) SetA(v float64) {
	ff.Set("a", v)
}
func (ff *Face3) B() float64 {
	return ff.Get("b").Float()
}
func (ff *Face3) SetB(v float64) {
	ff.Set("b", v)
}
func (ff *Face3) C() float64 {
	return ff.Get("c").Float()
}
func (ff *Face3) SetC(v float64) {
	ff.Set("c", v)
}
func (ff *Face3) Color() *Color {
	return &Color{Value: ff.Get("color")}
}
func (ff *Face3) SetColor(v *Color) {
	ff.Set("color", v)
}
func (ff *Face3) MaterialIndex() int {
	return ff.Get("materialIndex").Int()
}
func (ff *Face3) SetMaterialIndex(v int) {
	ff.Set("materialIndex", v)
}
func (ff *Face3) Normal() *Vector3 {
	return &Vector3{Value: ff.Get("normal")}
}
func (ff *Face3) SetNormal(v *Vector3) {
	ff.Set("normal", v)
}
func (ff *Face3) VertexColors() js.Value {
	return ff.Get("vertexColors")
}
func (ff *Face3) SetVertexColors(v js.Value) {
	ff.Set("vertexColors", v)
}
func (ff *Face3) VertexNormals() js.Value {
	return ff.Get("vertexNormals")
}
func (ff *Face3) SetVertexNormals(v js.Value) {
	ff.Set("vertexNormals", v)
}
func (ff *Face3) Clone() *Face3 {
	return &Face3{Value: ff.Call("clone")}
}
func (ff *Face3) Copy(source *Face3) *Face3 {
	return &Face3{Value: ff.Call("copy", source)}
}

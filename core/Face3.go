package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/math"
)

type Event interface {
}
type Face3 struct {
	js.Value
}

func NewFace3(a int, b int, c int, normal *math.Vector3, color *math.Color, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, normal, color, materialIndex)}
}
func NewFace32(a int, b int, c int, normal *math.Vector3, vertexColors []math.Color, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, normal, vertexColors, materialIndex)}
}
func NewFace33(a int, b int, c int, vertexNormals []math.Vector3, color *math.Color, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, vertexNormals, color, materialIndex)}
}
func NewFace34(a int, b int, c int, vertexNormals []math.Vector3, vertexColors []math.Color, materialIndex int) *Face3 {
	return &Face3{Value: get("Face3").New(a, b, c, vertexNormals, vertexColors, materialIndex)}
}
func (f *Face3) A() int {
	return f.Get("a").Int()
}
func (f *Face3) SetA(v int) {
	f.Set("a", v)
}
func (f *Face3) B() int {
	return f.Get("b").Int()
}
func (f *Face3) SetB(v int) {
	f.Set("b", v)
}
func (f *Face3) C() int {
	return f.Get("c").Int()
}
func (f *Face3) SetC(v int) {
	f.Set("c", v)
}
func (f *Face3) Color() *math.Color {
	return &math.Color{Value: f.Get("color")}
}
func (f *Face3) SetColor(v *math.Color) {
	f.Set("color", v)
}
func (f *Face3) MaterialIndex() int {
	return f.Get("materialIndex").Int()
}
func (f *Face3) SetMaterialIndex(v int) {
	f.Set("materialIndex", v)
}
func (f *Face3) Normal() *math.Vector3 {
	return &math.Vector3{Value: f.Get("normal")}
}
func (f *Face3) SetNormal(v *math.Vector3) {
	f.Set("normal", v)
}
func (f *Face3) VertexColors() []math.Color {
	return []math.Color(f.Get("vertexColors"))
}
func (f *Face3) SetVertexColors(v []math.Color) {
	f.Set("vertexColors", v)
}
func (f *Face3) VertexNormals() []math.Vector3 {
	return []math.Vector3(f.Get("vertexNormals"))
}
func (f *Face3) SetVertexNormals(v []math.Vector3) {
	f.Set("vertexNormals", v)
}
func (f *Face3) Clone() this {
	return this(f.Call("clone"))
}
func (f *Face3) Copy(source *Face3) this {
	return this(f.Call("copy", source))
}

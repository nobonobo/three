package loaders

import (
	"github.com/nobonobo/three/materials"
)

type Loader struct {
	js.Value
}

func NewLoader() *Loader {
	return &Loader{Value: get("Loader").New()}
}
func (l *Loader) CrossOrigin() string {
	return l.Get("crossOrigin").String()
}
func (l *Loader) SetCrossOrigin(v string) {
	l.Set("crossOrigin", v)
}
func (l *Loader) OnLoadComplete() js.Value {
	return l.Get("onLoadComplete")
}
func (l *Loader) SetOnLoadComplete(v js.Value) {
	l.Set("onLoadComplete", v)
}
func (l *Loader) OnLoadProgress() js.Value {
	return l.Get("onLoadProgress")
}
func (l *Loader) SetOnLoadProgress(v js.Value) {
	l.Set("onLoadProgress", v)
}
func (l *Loader) OnLoadStart() js.Value {
	return l.Get("onLoadStart")
}
func (l *Loader) SetOnLoadStart(v js.Value) {
	l.Set("onLoadStart", v)
}
func (l *Loader) Handlers() LoaderHandler {
	return LoaderHandler(l.Get("Handlers"))
}
func (l *Loader) SetHandlers(v LoaderHandler) {
	l.Set("Handlers", v)
}
func (l *Loader) CreateMaterial(m *materials.Material, texturePath string, crossOrigin string) bool {
	return l.Call("createMaterial", m, texturePath, crossOrigin).Bool()
}
func (l *Loader) ExtractUrlBase(url string) string {
	return l.Call("extractUrlBase", url).String()
}
func (l *Loader) InitMaterials(materials []materials.Material, texturePath string) []materials.Material {
	return []materials.Material(l.Call("initMaterials", materials, texturePath))
}

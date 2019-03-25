package core

type Layers struct {
	js.Value
}

func NewLayers() *Layers {
	return &Layers{Value: get("Layers").New()}
}
func (l *Layers) Mask() int {
	return l.Get("mask").Int()
}
func (l *Layers) SetMask(v int) {
	l.Set("mask", v)
}
func (l *Layers) Disable(channel int) {
	l.Call("disable", channel)
}
func (l *Layers) Enable(channel int) {
	l.Call("enable", channel)
}
func (l *Layers) Set(channel int) {
	l.Call("set", channel)
}
func (l *Layers) Test(layers *Layers) bool {
	return l.Call("test", layers).Bool()
}
func (l *Layers) Toggle(channel int) {
	l.Call("toggle", channel)
}

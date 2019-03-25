package audio

import (
	"github.com/gopherjs/gopherwasm/js"
)

type AudioAnalyser struct {
	js.Value
}

func NewAudioAnalyser(audio js.Value, fftSize int) *AudioAnalyser {
	return &AudioAnalyser{Value: get("AudioAnalyser").New(audio, fftSize)}
}
func (aa *AudioAnalyser) Analyser() js.Value {
	return aa.Get("analyser")
}
func (aa *AudioAnalyser) SetAnalyser(v js.Value) {
	aa.Set("analyser", v)
}
func (aa *AudioAnalyser) Data() *Uint8Array {
	return &Uint8Array{Value: aa.Get("data")}
}
func (aa *AudioAnalyser) SetData(v *Uint8Array) {
	aa.Set("data", v)
}
func (aa *AudioAnalyser) GetAverageFrequency() int {
	return aa.Call("getAverageFrequency").Int()
}
func (aa *AudioAnalyser) GetData(file js.Value) js.Value {
	return aa.Call("getData", file)
}
func (aa *AudioAnalyser) GetFrequencyData() *Uint8Array {
	return &Uint8Array{Value: aa.Call("getFrequencyData")}
}

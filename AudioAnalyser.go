// Code generated from "audio/AudioAnalyser.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type AudioAnalyser struct {
	js.Value
}

func NewAudioAnalyser(audio js.Value, fftSize float64) *AudioAnalyser {
	return &AudioAnalyser{Value: get("AudioAnalyser").New(audio, fftSize)}
}
func (aa *AudioAnalyser) Analyser() js.Value {
	return aa.Get("analyser")
}
func (aa *AudioAnalyser) SetAnalyser(v js.Value) {
	aa.Set("analyser", v)
}
func (aa *AudioAnalyser) Data() js.Value {
	return aa.Get("data")
}
func (aa *AudioAnalyser) SetData(v js.Value) {
	aa.Set("data", v)
}
func (aa *AudioAnalyser) GetAverageFrequency() float64 {
	return aa.Call("getAverageFrequency").Float()
}
func (aa *AudioAnalyser) GetData(file js.Value) js.Value {
	return aa.Call("getData", file)
}
func (aa *AudioAnalyser) GetFrequencyData() js.Value {
	return aa.Call("getFrequencyData")
}

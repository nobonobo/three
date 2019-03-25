package audio

func AudioContext() *AudioContext {
	return &AudioContext{Value: get("AudioContext")}
}
func SetAudioContext(v *AudioContext) {
	set("AudioContext", v)
}

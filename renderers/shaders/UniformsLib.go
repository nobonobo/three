package shaders

import (
	"github.com/gopherjs/gopherwasm/js"
)

func UniformsLib() js.Value {
	return get("UniformsLib")
}
func SetUniformsLib(v js.Value) {
	set("UniformsLib", v)
}

type IUniform interface {
}

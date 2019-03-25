package shaders

func ShaderLib() js.Value {
	return get("ShaderLib")
}
func SetShaderLib(v js.Value) {
	set("ShaderLib", v)
}

type Shader interface {
}

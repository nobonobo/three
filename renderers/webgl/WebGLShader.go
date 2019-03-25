package webgl

type WebGLShader struct {
	js.Value
}

func NewWebGLShader(gl js.Value, typ string, string string) *WebGLShader {
	return &WebGLShader{Value: get("WebGLShader").New(gl, typ, string)}
}

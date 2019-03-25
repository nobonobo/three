package helpers

type PolarGridHelper struct {
	js.Value
}

func NewPolarGridHelper(radius int, radials int, circles int, divisions int, color1 math.Color, color2 math.Color) *PolarGridHelper {
	return &PolarGridHelper{Value: get("PolarGridHelper").New(radius, radials, circles, divisions, color1, color2)}
}

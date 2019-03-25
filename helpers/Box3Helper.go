package helpers

import (
	"github.com/nobonobo/three/math"
)

type Box3Helper struct {
	js.Value
}

func NewBox3Helper(object *core.Object3D, color *math.Color) *Box3Helper {
	return &Box3Helper{Value: get("Box3Helper").New(object, color)}
}
func (bh *Box3Helper) Box() *math.Box3 {
	return &math.Box3{Value: bh.Get("box")}
}
func (bh *Box3Helper) SetBox(v *math.Box3) {
	bh.Set("box", v)
}
func (bh *Box3Helper) Type() string {
	return bh.Get("type").String()
}
func (bh *Box3Helper) SetType(v string) {
	bh.Set("type", v)
}

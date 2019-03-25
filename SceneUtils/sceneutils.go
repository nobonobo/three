package sceneutils

import (
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/scenes"
)

func Attach(child *core.Object3D, scene *scenes.Scene, parent *core.Object3D) {
	_Global.Call("attach", child, scene, parent)
}
func CreateMultiMaterialObject(geometry *core.Geometry, materials []materials.Material) *core.Object3D {
	return &core.Object3D{Value: _Global.Call("createMultiMaterialObject", geometry, materials)}
}
func Detach(child *core.Object3D, parent *core.Object3D, scene *scenes.Scene) {
	_Global.Call("detach", child, parent, scene)
}

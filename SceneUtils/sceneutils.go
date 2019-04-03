// Code generated from "SceneUtils/sceneutils.d.ts"; DO NOT EDIT.

package SceneUtils

import (
	"github.com/nobonobo/three"
)

func Attach(child *three.Object3D, scene *three.Scene, parent *three.Object3D) {
	_Global.Call("attach", child, scene, parent)
}
func CreateMultiMaterialObject(geometry *three.Geometry, materials []three.Material) *three.Object3D {
	return &three.Object3D{Value: _Global.Call("createMultiMaterialObject", geometry, materials)}
}
func Detach(child *three.Object3D, parent *three.Object3D, scene *three.Scene) {
	_Global.Call("detach", child, parent, scene)
}

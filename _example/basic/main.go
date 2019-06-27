package main

import (
	"syscall/js"

	"github.com/nobonobo/three"
)

const width = 640
const height = 480

func main() {

	renderer := three.NewWebGLRenderer(map[string]interface{}{
		"canvas": js.Global().Get("document").Call("querySelector", "#myCanvas"),
	})
	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(45, width/height, 10, 10000)
	camera.Position().Set2(0, 0, 1000)
	geometry := three.NewBoxGeometry(400, 400, 400, 1, 1, 1)
	material := three.NewMeshNormalMaterial(js.Null())
	box := three.NewMesh(geometry, material)

	scene.Add(box.JSValue())
	var tick func(this js.Value, args []js.Value) interface{}
	tick = func(this js.Value, args []js.Value) interface{} {
		box.Rotation().SetY(box.Rotation().Y() + 0.01)
		renderer.Render(scene, camera)
		js.Global().Call("requestAnimationFrame", js.FuncOf(tick))
		return nil
	}
	tick(js.Null(), nil)
}

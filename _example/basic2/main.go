package main

import (
	"syscall/js"

	"github.com/nobonobo/three"
)

var (
	renderer *three.WebGLRenderer
	scene three.Scene
	camera three.Camera
	mesh three.Mesh
)

func onWindowResize(this js.Value, args []js.Value) interface{} {
	width := js.Global().Get("innerWidth").Float()
	height := js.Global().Get("innerHeight").Float()
	renderer.SetPixelRatio(js.Global().Get("devicePixelRatio").Float())
	renderer.SetSize(width, height, true)
	camera.(*three.PerspectiveCamera).SetAspect(width / height)
	return nil
}

func animate(this js.Value, args []js.Value) interface{} {
	js.Global().Call("requestAnimationFrame", js.FuncOf(animate))
	mesh.Rotation().SetX(mesh.Rotation().X()+0.005)
	mesh.Rotation().SetY(mesh.Rotation().Y()+0.01)
	renderer.Render(scene, camera);
	return nil
}

func init() {
	camera = three.NewPerspectiveCamera(70, 4/3, 1, 1000)
	camera.Position().SetZ(400)
	scene = three.NewScene()
	texture := three.NewTextureLoader(three.DefaultLoadingManager()).Load(
		"../textures/crate.gif",
		js.Null(),
		js.Null(),
		js.Null(),
	)
	geometry := three.NewBoxBufferGeometry(200, 200, 200, 1, 1, 1)
	material := three.NewMeshBasicMaterial(js.ValueOf(map[string]interface{}{
		"map": texture,
	}))
	mesh = three.NewMesh(geometry, material)
	scene.Add(mesh.JSValue())
	renderer = three.NewWebGLRenderer(map[string]interface{}{"antialias": true})
	renderer.SetPixelRatio( js.Global().Get("devicePixelRatio").Float())
	width := js.Global().Get("innerWidth").Float()
	height := js.Global().Get("innerHeight").Float()
	renderer.SetSize(width, height, false)
	js.Global().Get("document").Get("body").Call("appendChild", renderer.Get("domElement"))
	js.Global().Call("addEventListener", "resize", js.FuncOf(onWindowResize));
	onWindowResize(js.Null(), nil)
}

func main() {
	animate(js.Null(), nil)
}

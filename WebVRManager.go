// Code generated from "renderers/webvr/WebVRManager.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WebVRManager interface {
	Dispose()
	GetCamera(camera *PerspectiveCamera) *PerspectiveCamera
	GetDevice() js.Value
	SetDevice(device js.Value)
	SetPoseTarget(object *Object3D)
	SubmitFrame()
}

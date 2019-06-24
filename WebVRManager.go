// Code generated from "renderers/webvr/WebVRManager.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type WebVRManager interface {
	Dispose()
	GetCamera(camera *PerspectiveCamera) *PerspectiveCamera
	GetDevice() js.Value
	SetDevice(device js.Value)
	SetPoseTarget(object *Object3D)
	SubmitFrame()
}

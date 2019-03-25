package webvr

import (
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
)

type WebVRManager interface {
	Dispose()
	GetCamera(camera *cameras.PerspectiveCamera) cameras.PerspectiveCamera
	GetDevice() VRDisplay
	SetDevice(device VRDisplay)
	SetPoseTarget(object core.Object3D)
	SubmitFrame()
}

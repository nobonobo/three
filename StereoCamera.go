// Code generated from "cameras/StereoCamera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type StereoCamera struct {
	js.Value
}

func NewStereoCamera() *StereoCamera {
	return &StereoCamera{Value: get("StereoCamera").New()}
}
func (sc *StereoCamera) Aspect() float64 {
	return sc.Get("aspect").Float()
}
func (sc *StereoCamera) SetAspect(v float64) {
	sc.Set("aspect", v)
}
func (sc *StereoCamera) CameraL() *PerspectiveCamera {
	return &PerspectiveCamera{Value: sc.Get("cameraL")}
}
func (sc *StereoCamera) SetCameraL(v *PerspectiveCamera) {
	sc.Set("cameraL", v)
}
func (sc *StereoCamera) CameraR() *PerspectiveCamera {
	return &PerspectiveCamera{Value: sc.Get("cameraR")}
}
func (sc *StereoCamera) SetCameraR(v *PerspectiveCamera) {
	sc.Set("cameraR", v)
}
func (sc *StereoCamera) CastShadow() bool {
	return sc.Get("castShadow").Bool()
}
func (sc *StereoCamera) SetCastShadow(v bool) {
	sc.Set("castShadow", v)
}
func (sc *StereoCamera) Children() js.Value {
	return sc.Get("children")
}
func (sc *StereoCamera) SetChildren(v js.Value) {
	sc.Set("children", v)
}
func (sc *StereoCamera) EyeSep() float64 {
	return sc.Get("eyeSep").Float()
}
func (sc *StereoCamera) SetEyeSep(v float64) {
	sc.Set("eyeSep", v)
}
func (sc *StereoCamera) FrustumCulled() bool {
	return sc.Get("frustumCulled").Bool()
}
func (sc *StereoCamera) SetFrustumCulled(v bool) {
	sc.Set("frustumCulled", v)
}
func (sc *StereoCamera) Id() int {
	return sc.Get("id").Int()
}
func (sc *StereoCamera) SetId(v int) {
	sc.Set("id", v)
}
func (sc *StereoCamera) IsCamera() bool {
	return sc.Get("isCamera").Bool()
}
func (sc *StereoCamera) SetIsCamera(v bool) {
	sc.Set("isCamera", v)
}
func (sc *StereoCamera) IsObject3D() bool {
	return sc.Get("isObject3D").Bool()
}
func (sc *StereoCamera) SetIsObject3D(v bool) {
	sc.Set("isObject3D", v)
}
func (sc *StereoCamera) Layers() *Layers {
	return &Layers{Value: sc.Get("layers")}
}
func (sc *StereoCamera) SetLayers(v *Layers) {
	sc.Set("layers", v)
}
func (sc *StereoCamera) Matrix() *Matrix4 {
	return &Matrix4{Value: sc.Get("matrix")}
}
func (sc *StereoCamera) SetMatrix(v *Matrix4) {
	sc.Set("matrix", v)
}
func (sc *StereoCamera) MatrixAutoUpdate() bool {
	return sc.Get("matrixAutoUpdate").Bool()
}
func (sc *StereoCamera) SetMatrixAutoUpdate(v bool) {
	sc.Set("matrixAutoUpdate", v)
}
func (sc *StereoCamera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: sc.Get("matrixWorld")}
}
func (sc *StereoCamera) SetMatrixWorld(v *Matrix4) {
	sc.Set("matrixWorld", v)
}
func (sc *StereoCamera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: sc.Get("matrixWorldInverse")}
}
func (sc *StereoCamera) SetMatrixWorldInverse(v *Matrix4) {
	sc.Set("matrixWorldInverse", v)
}
func (sc *StereoCamera) MatrixWorldNeedsUpdate() bool {
	return sc.Get("matrixWorldNeedsUpdate").Bool()
}
func (sc *StereoCamera) SetMatrixWorldNeedsUpdate(v bool) {
	sc.Set("matrixWorldNeedsUpdate", v)
}
func (sc *StereoCamera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: sc.Get("modelViewMatrix")}
}
func (sc *StereoCamera) SetModelViewMatrix(v *Matrix4) {
	sc.Set("modelViewMatrix", v)
}
func (sc *StereoCamera) Name() string {
	return sc.Get("name").String()
}
func (sc *StereoCamera) SetName(v string) {
	sc.Set("name", v)
}
func (sc *StereoCamera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: sc.Get("normalMatrix")}
}
func (sc *StereoCamera) SetNormalMatrix(v *Matrix3) {
	sc.Set("normalMatrix", v)
}
func (sc *StereoCamera) OnAfterRender() js.Value {
	return sc.Get("onAfterRender")
}
func (sc *StereoCamera) SetOnAfterRender(v js.Value) {
	sc.Set("onAfterRender", v)
}
func (sc *StereoCamera) OnBeforeRender() js.Value {
	return sc.Get("onBeforeRender")
}
func (sc *StereoCamera) SetOnBeforeRender(v js.Value) {
	sc.Set("onBeforeRender", v)
}
func (sc *StereoCamera) Parent() *Object3D {
	return &Object3D{Value: sc.Get("parent")}
}
func (sc *StereoCamera) SetParent(v *Object3D) {
	sc.Set("parent", v)
}
func (sc *StereoCamera) Position() *Vector3 {
	return &Vector3{Value: sc.Get("position")}
}
func (sc *StereoCamera) SetPosition(v *Vector3) {
	sc.Set("position", v)
}
func (sc *StereoCamera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: sc.Get("projectionMatrix")}
}
func (sc *StereoCamera) SetProjectionMatrix(v *Matrix4) {
	sc.Set("projectionMatrix", v)
}
func (sc *StereoCamera) Quaternion() *Quaternion {
	return &Quaternion{Value: sc.Get("quaternion")}
}
func (sc *StereoCamera) SetQuaternion(v *Quaternion) {
	sc.Set("quaternion", v)
}
func (sc *StereoCamera) ReceiveShadow() bool {
	return sc.Get("receiveShadow").Bool()
}
func (sc *StereoCamera) SetReceiveShadow(v bool) {
	sc.Set("receiveShadow", v)
}
func (sc *StereoCamera) RenderOrder() float64 {
	return sc.Get("renderOrder").Float()
}
func (sc *StereoCamera) SetRenderOrder(v float64) {
	sc.Set("renderOrder", v)
}
func (sc *StereoCamera) Rotation() *Euler {
	return &Euler{Value: sc.Get("rotation")}
}
func (sc *StereoCamera) SetRotation(v *Euler) {
	sc.Set("rotation", v)
}
func (sc *StereoCamera) Scale() *Vector3 {
	return &Vector3{Value: sc.Get("scale")}
}
func (sc *StereoCamera) SetScale(v *Vector3) {
	sc.Set("scale", v)
}
func (sc *StereoCamera) Type() string {
	return sc.Get("type").String()
}
func (sc *StereoCamera) SetType(v string) {
	sc.Set("type", v)
}
func (sc *StereoCamera) Up() *Vector3 {
	return &Vector3{Value: sc.Get("up")}
}
func (sc *StereoCamera) SetUp(v *Vector3) {
	sc.Set("up", v)
}
func (sc *StereoCamera) UserData() js.Value {
	return sc.Get("userData")
}
func (sc *StereoCamera) SetUserData(v js.Value) {
	sc.Set("userData", v)
}
func (sc *StereoCamera) Uuid() string {
	return sc.Get("uuid").String()
}
func (sc *StereoCamera) SetUuid(v string) {
	sc.Set("uuid", v)
}
func (sc *StereoCamera) Visible() bool {
	return sc.Get("visible").Bool()
}
func (sc *StereoCamera) SetVisible(v bool) {
	sc.Set("visible", v)
}
func (sc *StereoCamera) DefaultMatrixAutoUpdate() bool {
	return sc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sc *StereoCamera) SetDefaultMatrixAutoUpdate(v bool) {
	sc.Set("DefaultMatrixAutoUpdate", v)
}
func (sc *StereoCamera) DefaultUp() *Vector3 {
	return &Vector3{Value: sc.Get("DefaultUp")}
}
func (sc *StereoCamera) SetDefaultUp(v *Vector3) {
	sc.Set("DefaultUp", v)
}
func (sc *StereoCamera) Add(object js.Value) *StereoCamera {
	return &StereoCamera{Value: sc.Call("add", object)}
}
func (sc *StereoCamera) AddEventListener(typ string, listener js.Value) {
	sc.Call("addEventListener", typ, listener)
}
func (sc *StereoCamera) ApplyMatrix(matrix *Matrix4) {
	sc.Call("applyMatrix", matrix)
}
func (sc *StereoCamera) ApplyQuaternion(quaternion *Quaternion) *StereoCamera {
	return &StereoCamera{Value: sc.Call("applyQuaternion", quaternion)}
}
func (sc *StereoCamera) Clone(recursive bool) *StereoCamera {
	return &StereoCamera{Value: sc.Call("clone", recursive)}
}
func (sc *StereoCamera) Copy(source *Camera, recursive bool) *StereoCamera {
	return &StereoCamera{Value: sc.Call("copy", source, recursive)}
}
func (sc *StereoCamera) DispatchEvent(event js.Value) {
	sc.Call("dispatchEvent", event)
}
func (sc *StereoCamera) GetObjectById(id int) *Object3D {
	return &Object3D{Value: sc.Call("getObjectById", id)}
}
func (sc *StereoCamera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: sc.Call("getObjectByName", name)}
}
func (sc *StereoCamera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: sc.Call("getObjectByProperty", name, value)}
}
func (sc *StereoCamera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: sc.Call("getWorldDirection", target)}
}
func (sc *StereoCamera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: sc.Call("getWorldPosition", target)}
}
func (sc *StereoCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: sc.Call("getWorldQuaternion", target)}
}
func (sc *StereoCamera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: sc.Call("getWorldScale", target)}
}
func (sc *StereoCamera) HasEventListener(typ string, listener js.Value) bool {
	return sc.Call("hasEventListener", typ, listener).Bool()
}
func (sc *StereoCamera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: sc.Call("localToWorld", vector)}
}
func (sc *StereoCamera) LookAt(vector *Vector3, y float64, z float64) {
	sc.Call("lookAt", vector, y, z)
}
func (sc *StereoCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
	sc.Call("raycast", raycaster, intersects)
}
func (sc *StereoCamera) Remove(object js.Value) *StereoCamera {
	return &StereoCamera{Value: sc.Call("remove", object)}
}
func (sc *StereoCamera) RemoveEventListener(typ string, listener js.Value) {
	sc.Call("removeEventListener", typ, listener)
}
func (sc *StereoCamera) RotateOnAxis(axis *Vector3, angle float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("rotateOnAxis", axis, angle)}
}
func (sc *StereoCamera) RotateOnWorldAxis(axis *Vector3, angle float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("rotateOnWorldAxis", axis, angle)}
}
func (sc *StereoCamera) RotateX(angle float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("rotateX", angle)}
}
func (sc *StereoCamera) RotateY(angle float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("rotateY", angle)}
}
func (sc *StereoCamera) RotateZ(angle float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("rotateZ", angle)}
}
func (sc *StereoCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	sc.Call("setRotationFromAxisAngle", axis, angle)
}
func (sc *StereoCamera) SetRotationFromEuler(euler *Euler) {
	sc.Call("setRotationFromEuler", euler)
}
func (sc *StereoCamera) SetRotationFromMatrix(m *Matrix4) {
	sc.Call("setRotationFromMatrix", m)
}
func (sc *StereoCamera) SetRotationFromQuaternion(q *Quaternion) {
	sc.Call("setRotationFromQuaternion", q)
}
func (sc *StereoCamera) ToJSON(meta js.Value) js.Value {
	return sc.Call("toJSON", meta)
}
func (sc *StereoCamera) TranslateOnAxis(axis *Vector3, distance float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("translateOnAxis", axis, distance)}
}
func (sc *StereoCamera) TranslateX(distance float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("translateX", distance)}
}
func (sc *StereoCamera) TranslateY(distance float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("translateY", distance)}
}
func (sc *StereoCamera) TranslateZ(distance float64) *StereoCamera {
	return &StereoCamera{Value: sc.Call("translateZ", distance)}
}
func (sc *StereoCamera) Traverse(callback js.Value) {
	sc.Call("traverse", callback)
}
func (sc *StereoCamera) TraverseAncestors(callback js.Value) {
	sc.Call("traverseAncestors", callback)
}
func (sc *StereoCamera) TraverseVisible(callback js.Value) {
	sc.Call("traverseVisible", callback)
}
func (sc *StereoCamera) Update(camera *PerspectiveCamera) {
	sc.Call("update", camera)
}
func (sc *StereoCamera) UpdateMatrix() {
	sc.Call("updateMatrix")
}
func (sc *StereoCamera) UpdateMatrixWorld(force bool) {
	sc.Call("updateMatrixWorld", force)
}
func (sc *StereoCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sc *StereoCamera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: sc.Call("worldToLocal", vector)}
}

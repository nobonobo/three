// Code generated from "cameras/ArrayCamera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// ArrayCamera extend: [PerspectiveCamera]
type ArrayCamera struct {
	js.Value
}

func NewArrayCamera(cameras js.Value) *ArrayCamera {
	return &ArrayCamera{Value: get("ArrayCamera").New(cameras)}
}
func (ac *ArrayCamera) JSValue() js.Value {
	return ac.Value
}
func (ac *ArrayCamera) Aspect() float64 {
	return ac.Get("aspect").Float()
}
func (ac *ArrayCamera) SetAspect(v float64) {
	ac.Set("aspect", v)
}
func (ac *ArrayCamera) Cameras() js.Value {
	return ac.Get("cameras")
}
func (ac *ArrayCamera) SetCameras(v js.Value) {
	ac.Set("cameras", v)
}
func (ac *ArrayCamera) CastShadow() bool {
	return ac.Get("castShadow").Bool()
}
func (ac *ArrayCamera) SetCastShadow(v bool) {
	ac.Set("castShadow", v)
}
func (ac *ArrayCamera) Children() js.Value {
	return ac.Get("children")
}
func (ac *ArrayCamera) SetChildren(v js.Value) {
	ac.Set("children", v)
}
func (ac *ArrayCamera) Far() float64 {
	return ac.Get("far").Float()
}
func (ac *ArrayCamera) SetFar(v float64) {
	ac.Set("far", v)
}
func (ac *ArrayCamera) FilmGauge() float64 {
	return ac.Get("filmGauge").Float()
}
func (ac *ArrayCamera) SetFilmGauge(v float64) {
	ac.Set("filmGauge", v)
}
func (ac *ArrayCamera) FilmOffset() float64 {
	return ac.Get("filmOffset").Float()
}
func (ac *ArrayCamera) SetFilmOffset(v float64) {
	ac.Set("filmOffset", v)
}
func (ac *ArrayCamera) Focus() float64 {
	return ac.Get("focus").Float()
}
func (ac *ArrayCamera) SetFocus(v float64) {
	ac.Set("focus", v)
}
func (ac *ArrayCamera) Fov() float64 {
	return ac.Get("fov").Float()
}
func (ac *ArrayCamera) SetFov(v float64) {
	ac.Set("fov", v)
}
func (ac *ArrayCamera) FrustumCulled() bool {
	return ac.Get("frustumCulled").Bool()
}
func (ac *ArrayCamera) SetFrustumCulled(v bool) {
	ac.Set("frustumCulled", v)
}
func (ac *ArrayCamera) Id() int {
	return ac.Get("id").Int()
}
func (ac *ArrayCamera) SetId(v int) {
	ac.Set("id", v)
}
func (ac *ArrayCamera) IsArrayCamera() bool {
	return ac.Get("isArrayCamera").Bool()
}
func (ac *ArrayCamera) SetIsArrayCamera(v bool) {
	ac.Set("isArrayCamera", v)
}
func (ac *ArrayCamera) IsCamera() bool {
	return ac.Get("isCamera").Bool()
}
func (ac *ArrayCamera) SetIsCamera(v bool) {
	ac.Set("isCamera", v)
}
func (ac *ArrayCamera) IsObject3D() bool {
	return ac.Get("isObject3D").Bool()
}
func (ac *ArrayCamera) SetIsObject3D(v bool) {
	ac.Set("isObject3D", v)
}
func (ac *ArrayCamera) IsPerspectiveCamera() bool {
	return ac.Get("isPerspectiveCamera").Bool()
}
func (ac *ArrayCamera) SetIsPerspectiveCamera(v bool) {
	ac.Set("isPerspectiveCamera", v)
}
func (ac *ArrayCamera) Layers() *Layers {
	return &Layers{Value: ac.Get("layers")}
}
func (ac *ArrayCamera) SetLayers(v *Layers) {
	ac.Set("layers", v.JSValue())
}
func (ac *ArrayCamera) Matrix() *Matrix4 {
	return &Matrix4{Value: ac.Get("matrix")}
}
func (ac *ArrayCamera) SetMatrix(v *Matrix4) {
	ac.Set("matrix", v.JSValue())
}
func (ac *ArrayCamera) MatrixAutoUpdate() bool {
	return ac.Get("matrixAutoUpdate").Bool()
}
func (ac *ArrayCamera) SetMatrixAutoUpdate(v bool) {
	ac.Set("matrixAutoUpdate", v)
}
func (ac *ArrayCamera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ac.Get("matrixWorld")}
}
func (ac *ArrayCamera) SetMatrixWorld(v *Matrix4) {
	ac.Set("matrixWorld", v.JSValue())
}
func (ac *ArrayCamera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: ac.Get("matrixWorldInverse")}
}
func (ac *ArrayCamera) SetMatrixWorldInverse(v *Matrix4) {
	ac.Set("matrixWorldInverse", v.JSValue())
}
func (ac *ArrayCamera) MatrixWorldNeedsUpdate() bool {
	return ac.Get("matrixWorldNeedsUpdate").Bool()
}
func (ac *ArrayCamera) SetMatrixWorldNeedsUpdate(v bool) {
	ac.Set("matrixWorldNeedsUpdate", v)
}
func (ac *ArrayCamera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ac.Get("modelViewMatrix")}
}
func (ac *ArrayCamera) SetModelViewMatrix(v *Matrix4) {
	ac.Set("modelViewMatrix", v.JSValue())
}
func (ac *ArrayCamera) Name() string {
	return ac.Get("name").String()
}
func (ac *ArrayCamera) SetName(v string) {
	ac.Set("name", v)
}
func (ac *ArrayCamera) Near() float64 {
	return ac.Get("near").Float()
}
func (ac *ArrayCamera) SetNear(v float64) {
	ac.Set("near", v)
}
func (ac *ArrayCamera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ac.Get("normalMatrix")}
}
func (ac *ArrayCamera) SetNormalMatrix(v *Matrix3) {
	ac.Set("normalMatrix", v.JSValue())
}
func (ac *ArrayCamera) OnAfterRender() js.Value {
	return ac.Get("onAfterRender")
}
func (ac *ArrayCamera) SetOnAfterRender(v js.Value) {
	ac.Set("onAfterRender", v)
}
func (ac *ArrayCamera) OnBeforeRender() js.Value {
	return ac.Get("onBeforeRender")
}
func (ac *ArrayCamera) SetOnBeforeRender(v js.Value) {
	ac.Set("onBeforeRender", v)
}
func (ac *ArrayCamera) Parent() *Object3D {
	return &Object3D{Value: ac.Get("parent")}
}
func (ac *ArrayCamera) SetParent(v *Object3D) {
	ac.Set("parent", v.JSValue())
}
func (ac *ArrayCamera) Position() *Vector3 {
	return &Vector3{Value: ac.Get("position")}
}
func (ac *ArrayCamera) SetPosition(v *Vector3) {
	ac.Set("position", v.JSValue())
}
func (ac *ArrayCamera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: ac.Get("projectionMatrix")}
}
func (ac *ArrayCamera) SetProjectionMatrix(v *Matrix4) {
	ac.Set("projectionMatrix", v.JSValue())
}
func (ac *ArrayCamera) Quaternion() *Quaternion {
	return &Quaternion{Value: ac.Get("quaternion")}
}
func (ac *ArrayCamera) SetQuaternion(v *Quaternion) {
	ac.Set("quaternion", v.JSValue())
}
func (ac *ArrayCamera) ReceiveShadow() bool {
	return ac.Get("receiveShadow").Bool()
}
func (ac *ArrayCamera) SetReceiveShadow(v bool) {
	ac.Set("receiveShadow", v)
}
func (ac *ArrayCamera) RenderOrder() float64 {
	return ac.Get("renderOrder").Float()
}
func (ac *ArrayCamera) SetRenderOrder(v float64) {
	ac.Set("renderOrder", v)
}
func (ac *ArrayCamera) Rotation() *Euler {
	return &Euler{Value: ac.Get("rotation")}
}
func (ac *ArrayCamera) SetRotation(v *Euler) {
	ac.Set("rotation", v.JSValue())
}
func (ac *ArrayCamera) Scale() *Vector3 {
	return &Vector3{Value: ac.Get("scale")}
}
func (ac *ArrayCamera) SetScale(v *Vector3) {
	ac.Set("scale", v.JSValue())
}
func (ac *ArrayCamera) Type() string {
	return ac.Get("type").String()
}
func (ac *ArrayCamera) SetType(v string) {
	ac.Set("type", v)
}
func (ac *ArrayCamera) Up() *Vector3 {
	return &Vector3{Value: ac.Get("up")}
}
func (ac *ArrayCamera) SetUp(v *Vector3) {
	ac.Set("up", v.JSValue())
}
func (ac *ArrayCamera) UserData() js.Value {
	return ac.Get("userData")
}
func (ac *ArrayCamera) SetUserData(v js.Value) {
	ac.Set("userData", v)
}
func (ac *ArrayCamera) Uuid() string {
	return ac.Get("uuid").String()
}
func (ac *ArrayCamera) SetUuid(v string) {
	ac.Set("uuid", v)
}
func (ac *ArrayCamera) View() js.Value {
	return ac.Get("view")
}
func (ac *ArrayCamera) SetView(v js.Value) {
	ac.Set("view", v)
}
func (ac *ArrayCamera) Visible() bool {
	return ac.Get("visible").Bool()
}
func (ac *ArrayCamera) SetVisible(v bool) {
	ac.Set("visible", v)
}
func (ac *ArrayCamera) Zoom() float64 {
	return ac.Get("zoom").Float()
}
func (ac *ArrayCamera) SetZoom(v float64) {
	ac.Set("zoom", v)
}
func (ac *ArrayCamera) DefaultMatrixAutoUpdate() bool {
	return ac.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ac *ArrayCamera) SetDefaultMatrixAutoUpdate(v bool) {
	ac.Set("DefaultMatrixAutoUpdate", v)
}
func (ac *ArrayCamera) DefaultUp() *Vector3 {
	return &Vector3{Value: ac.Get("DefaultUp")}
}
func (ac *ArrayCamera) SetDefaultUp(v *Vector3) {
	ac.Set("DefaultUp", v.JSValue())
}
func (ac *ArrayCamera) Add(object js.Value) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("add", object)}
}
func (ac *ArrayCamera) AddEventListener(typ string, listener js.Value) {
	ac.Call("addEventListener", typ, listener)
}
func (ac *ArrayCamera) ApplyMatrix(matrix *Matrix4) {
	ac.Call("applyMatrix", matrix.JSValue())
}
func (ac *ArrayCamera) ApplyQuaternion(quaternion *Quaternion) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("applyQuaternion", quaternion)}
}
func (ac *ArrayCamera) ClearViewOffset() {
	ac.Call("clearViewOffset")
}
func (ac *ArrayCamera) Clone(recursive bool) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("clone", recursive)}
}
func (ac *ArrayCamera) Copy(source Camera, recursive bool) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("copy", source.JSValue(), recursive)}
}
func (ac *ArrayCamera) DispatchEvent(event js.Value) {
	ac.Call("dispatchEvent", event)
}
func (ac *ArrayCamera) GetEffectiveFOV() float64 {
	return ac.Call("getEffectiveFOV").Float()
}
func (ac *ArrayCamera) GetFilmHeight() float64 {
	return ac.Call("getFilmHeight").Float()
}
func (ac *ArrayCamera) GetFilmWidth() float64 {
	return ac.Call("getFilmWidth").Float()
}
func (ac *ArrayCamera) GetFocalLength() float64 {
	return ac.Call("getFocalLength").Float()
}
func (ac *ArrayCamera) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ac.Call("getObjectById", id)}
}
func (ac *ArrayCamera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ac.Call("getObjectByName", name)}
}
func (ac *ArrayCamera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ac.Call("getObjectByProperty", name, value)}
}
func (ac *ArrayCamera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ac.Call("getWorldDirection", target)}
}
func (ac *ArrayCamera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ac.Call("getWorldPosition", target)}
}
func (ac *ArrayCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ac.Call("getWorldQuaternion", target)}
}
func (ac *ArrayCamera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ac.Call("getWorldScale", target)}
}
func (ac *ArrayCamera) HasEventListener(typ string, listener js.Value) bool {
	return ac.Call("hasEventListener", typ, listener).Bool()
}
func (ac *ArrayCamera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ac.Call("localToWorld", vector)}
}
func (ac *ArrayCamera) LookAt(vector *Vector3, y float64, z float64) {
	ac.Call("lookAt", vector, y, z)
}
func (ac *ArrayCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
	ac.Call("raycast", raycaster.JSValue(), intersects)
}
func (ac *ArrayCamera) Remove(object js.Value) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("remove", object)}
}
func (ac *ArrayCamera) RemoveEventListener(typ string, listener js.Value) {
	ac.Call("removeEventListener", typ, listener)
}
func (ac *ArrayCamera) RotateOnAxis(axis *Vector3, angle float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("rotateOnAxis", axis, angle)}
}
func (ac *ArrayCamera) RotateOnWorldAxis(axis *Vector3, angle float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("rotateOnWorldAxis", axis, angle)}
}
func (ac *ArrayCamera) RotateX(angle float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("rotateX", angle)}
}
func (ac *ArrayCamera) RotateY(angle float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("rotateY", angle)}
}
func (ac *ArrayCamera) RotateZ(angle float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("rotateZ", angle)}
}
func (ac *ArrayCamera) SetFocalLength(focalLength float64) {
	ac.Call("setFocalLength", focalLength)
}
func (ac *ArrayCamera) SetLens(focalLength float64, frameHeight float64) {
	ac.Call("setLens", focalLength, frameHeight)
}
func (ac *ArrayCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ac.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (ac *ArrayCamera) SetRotationFromEuler(euler *Euler) {
	ac.Call("setRotationFromEuler", euler.JSValue())
}
func (ac *ArrayCamera) SetRotationFromMatrix(m *Matrix4) {
	ac.Call("setRotationFromMatrix", m.JSValue())
}
func (ac *ArrayCamera) SetRotationFromQuaternion(q *Quaternion) {
	ac.Call("setRotationFromQuaternion", q.JSValue())
}
func (ac *ArrayCamera) SetViewOffset(fullWidth float64, fullHeight float64, x float64, y float64, width float64, height float64) {
	ac.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}
func (ac *ArrayCamera) ToJSON(meta js.Value) js.Value {
	return ac.Call("toJSON", meta)
}
func (ac *ArrayCamera) TranslateOnAxis(axis *Vector3, distance float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("translateOnAxis", axis, distance)}
}
func (ac *ArrayCamera) TranslateX(distance float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("translateX", distance)}
}
func (ac *ArrayCamera) TranslateY(distance float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("translateY", distance)}
}
func (ac *ArrayCamera) TranslateZ(distance float64) *ArrayCamera {
	return &ArrayCamera{Value: ac.Call("translateZ", distance)}
}
func (ac *ArrayCamera) Traverse(callback js.Value) {
	ac.Call("traverse", callback)
}
func (ac *ArrayCamera) TraverseAncestors(callback js.Value) {
	ac.Call("traverseAncestors", callback)
}
func (ac *ArrayCamera) TraverseVisible(callback js.Value) {
	ac.Call("traverseVisible", callback)
}
func (ac *ArrayCamera) UpdateMatrix() {
	ac.Call("updateMatrix")
}
func (ac *ArrayCamera) UpdateMatrixWorld(force bool) {
	ac.Call("updateMatrixWorld", force)
}
func (ac *ArrayCamera) UpdateProjectionMatrix() {
	ac.Call("updateProjectionMatrix")
}
func (ac *ArrayCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ac.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ac *ArrayCamera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ac.Call("worldToLocal", vector)}
}

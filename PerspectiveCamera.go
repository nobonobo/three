// Code generated from "cameras/PerspectiveCamera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PerspectiveCamera extend: [Camera]
type PerspectiveCamera struct {
	js.Value
}

func NewPerspectiveCamera(fov float64, aspect float64, near float64, far float64) *PerspectiveCamera {
	return &PerspectiveCamera{Value: get("PerspectiveCamera").New(fov, aspect, near, far)}
}
func (pc *PerspectiveCamera) JSValue() js.Value {
	return pc.Value
}
func (pc *PerspectiveCamera) Aspect() float64 {
	return pc.Get("aspect").Float()
}
func (pc *PerspectiveCamera) SetAspect(v float64) {
	pc.Set("aspect", v)
}
func (pc *PerspectiveCamera) CastShadow() bool {
	return pc.Get("castShadow").Bool()
}
func (pc *PerspectiveCamera) SetCastShadow(v bool) {
	pc.Set("castShadow", v)
}
func (pc *PerspectiveCamera) Children() js.Value {
	return pc.Get("children")
}
func (pc *PerspectiveCamera) SetChildren(v js.Value) {
	pc.Set("children", v)
}
func (pc *PerspectiveCamera) Far() float64 {
	return pc.Get("far").Float()
}
func (pc *PerspectiveCamera) SetFar(v float64) {
	pc.Set("far", v)
}
func (pc *PerspectiveCamera) FilmGauge() float64 {
	return pc.Get("filmGauge").Float()
}
func (pc *PerspectiveCamera) SetFilmGauge(v float64) {
	pc.Set("filmGauge", v)
}
func (pc *PerspectiveCamera) FilmOffset() float64 {
	return pc.Get("filmOffset").Float()
}
func (pc *PerspectiveCamera) SetFilmOffset(v float64) {
	pc.Set("filmOffset", v)
}
func (pc *PerspectiveCamera) Focus() float64 {
	return pc.Get("focus").Float()
}
func (pc *PerspectiveCamera) SetFocus(v float64) {
	pc.Set("focus", v)
}
func (pc *PerspectiveCamera) Fov() float64 {
	return pc.Get("fov").Float()
}
func (pc *PerspectiveCamera) SetFov(v float64) {
	pc.Set("fov", v)
}
func (pc *PerspectiveCamera) FrustumCulled() bool {
	return pc.Get("frustumCulled").Bool()
}
func (pc *PerspectiveCamera) SetFrustumCulled(v bool) {
	pc.Set("frustumCulled", v)
}
func (pc *PerspectiveCamera) Id() int {
	return pc.Get("id").Int()
}
func (pc *PerspectiveCamera) SetId(v int) {
	pc.Set("id", v)
}
func (pc *PerspectiveCamera) IsCamera() bool {
	return pc.Get("isCamera").Bool()
}
func (pc *PerspectiveCamera) SetIsCamera(v bool) {
	pc.Set("isCamera", v)
}
func (pc *PerspectiveCamera) IsObject3D() bool {
	return pc.Get("isObject3D").Bool()
}
func (pc *PerspectiveCamera) SetIsObject3D(v bool) {
	pc.Set("isObject3D", v)
}
func (pc *PerspectiveCamera) IsPerspectiveCamera() bool {
	return pc.Get("isPerspectiveCamera").Bool()
}
func (pc *PerspectiveCamera) SetIsPerspectiveCamera(v bool) {
	pc.Set("isPerspectiveCamera", v)
}
func (pc *PerspectiveCamera) Layers() *Layers {
	return &Layers{Value: pc.Get("layers")}
}
func (pc *PerspectiveCamera) SetLayers(v *Layers) {
	pc.Set("layers", v.JSValue())
}
func (pc *PerspectiveCamera) Matrix() *Matrix4 {
	return &Matrix4{Value: pc.Get("matrix")}
}
func (pc *PerspectiveCamera) SetMatrix(v *Matrix4) {
	pc.Set("matrix", v.JSValue())
}
func (pc *PerspectiveCamera) MatrixAutoUpdate() bool {
	return pc.Get("matrixAutoUpdate").Bool()
}
func (pc *PerspectiveCamera) SetMatrixAutoUpdate(v bool) {
	pc.Set("matrixAutoUpdate", v)
}
func (pc *PerspectiveCamera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: pc.Get("matrixWorld")}
}
func (pc *PerspectiveCamera) SetMatrixWorld(v *Matrix4) {
	pc.Set("matrixWorld", v.JSValue())
}
func (pc *PerspectiveCamera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: pc.Get("matrixWorldInverse")}
}
func (pc *PerspectiveCamera) SetMatrixWorldInverse(v *Matrix4) {
	pc.Set("matrixWorldInverse", v.JSValue())
}
func (pc *PerspectiveCamera) MatrixWorldNeedsUpdate() bool {
	return pc.Get("matrixWorldNeedsUpdate").Bool()
}
func (pc *PerspectiveCamera) SetMatrixWorldNeedsUpdate(v bool) {
	pc.Set("matrixWorldNeedsUpdate", v)
}
func (pc *PerspectiveCamera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: pc.Get("modelViewMatrix")}
}
func (pc *PerspectiveCamera) SetModelViewMatrix(v *Matrix4) {
	pc.Set("modelViewMatrix", v.JSValue())
}
func (pc *PerspectiveCamera) Name() string {
	return pc.Get("name").String()
}
func (pc *PerspectiveCamera) SetName(v string) {
	pc.Set("name", v)
}
func (pc *PerspectiveCamera) Near() float64 {
	return pc.Get("near").Float()
}
func (pc *PerspectiveCamera) SetNear(v float64) {
	pc.Set("near", v)
}
func (pc *PerspectiveCamera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: pc.Get("normalMatrix")}
}
func (pc *PerspectiveCamera) SetNormalMatrix(v *Matrix3) {
	pc.Set("normalMatrix", v.JSValue())
}
func (pc *PerspectiveCamera) OnAfterRender() js.Value {
	return pc.Get("onAfterRender")
}
func (pc *PerspectiveCamera) SetOnAfterRender(v js.Value) {
	pc.Set("onAfterRender", v)
}
func (pc *PerspectiveCamera) OnBeforeRender() js.Value {
	return pc.Get("onBeforeRender")
}
func (pc *PerspectiveCamera) SetOnBeforeRender(v js.Value) {
	pc.Set("onBeforeRender", v)
}
func (pc *PerspectiveCamera) Parent() *Object3D {
	return &Object3D{Value: pc.Get("parent")}
}
func (pc *PerspectiveCamera) SetParent(v *Object3D) {
	pc.Set("parent", v.JSValue())
}
func (pc *PerspectiveCamera) Position() *Vector3 {
	return &Vector3{Value: pc.Get("position")}
}
func (pc *PerspectiveCamera) SetPosition(v *Vector3) {
	pc.Set("position", v.JSValue())
}
func (pc *PerspectiveCamera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: pc.Get("projectionMatrix")}
}
func (pc *PerspectiveCamera) SetProjectionMatrix(v *Matrix4) {
	pc.Set("projectionMatrix", v.JSValue())
}
func (pc *PerspectiveCamera) Quaternion() *Quaternion {
	return &Quaternion{Value: pc.Get("quaternion")}
}
func (pc *PerspectiveCamera) SetQuaternion(v *Quaternion) {
	pc.Set("quaternion", v.JSValue())
}
func (pc *PerspectiveCamera) ReceiveShadow() bool {
	return pc.Get("receiveShadow").Bool()
}
func (pc *PerspectiveCamera) SetReceiveShadow(v bool) {
	pc.Set("receiveShadow", v)
}
func (pc *PerspectiveCamera) RenderOrder() float64 {
	return pc.Get("renderOrder").Float()
}
func (pc *PerspectiveCamera) SetRenderOrder(v float64) {
	pc.Set("renderOrder", v)
}
func (pc *PerspectiveCamera) Rotation() *Euler {
	return &Euler{Value: pc.Get("rotation")}
}
func (pc *PerspectiveCamera) SetRotation(v *Euler) {
	pc.Set("rotation", v.JSValue())
}
func (pc *PerspectiveCamera) Scale() *Vector3 {
	return &Vector3{Value: pc.Get("scale")}
}
func (pc *PerspectiveCamera) SetScale(v *Vector3) {
	pc.Set("scale", v.JSValue())
}
func (pc *PerspectiveCamera) Type() string {
	return pc.Get("type").String()
}
func (pc *PerspectiveCamera) SetType(v string) {
	pc.Set("type", v)
}
func (pc *PerspectiveCamera) Up() *Vector3 {
	return &Vector3{Value: pc.Get("up")}
}
func (pc *PerspectiveCamera) SetUp(v *Vector3) {
	pc.Set("up", v.JSValue())
}
func (pc *PerspectiveCamera) UserData() js.Value {
	return pc.Get("userData")
}
func (pc *PerspectiveCamera) SetUserData(v js.Value) {
	pc.Set("userData", v)
}
func (pc *PerspectiveCamera) Uuid() string {
	return pc.Get("uuid").String()
}
func (pc *PerspectiveCamera) SetUuid(v string) {
	pc.Set("uuid", v)
}
func (pc *PerspectiveCamera) View() js.Value {
	return pc.Get("view")
}
func (pc *PerspectiveCamera) SetView(v js.Value) {
	pc.Set("view", v)
}
func (pc *PerspectiveCamera) Visible() bool {
	return pc.Get("visible").Bool()
}
func (pc *PerspectiveCamera) SetVisible(v bool) {
	pc.Set("visible", v)
}
func (pc *PerspectiveCamera) Zoom() float64 {
	return pc.Get("zoom").Float()
}
func (pc *PerspectiveCamera) SetZoom(v float64) {
	pc.Set("zoom", v)
}
func (pc *PerspectiveCamera) DefaultMatrixAutoUpdate() bool {
	return pc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pc *PerspectiveCamera) SetDefaultMatrixAutoUpdate(v bool) {
	pc.Set("DefaultMatrixAutoUpdate", v)
}
func (pc *PerspectiveCamera) DefaultUp() *Vector3 {
	return &Vector3{Value: pc.Get("DefaultUp")}
}
func (pc *PerspectiveCamera) SetDefaultUp(v *Vector3) {
	pc.Set("DefaultUp", v.JSValue())
}
func (pc *PerspectiveCamera) Add(object js.Value) Camera {
	return &PerspectiveCamera{Value: pc.Call("add", object)}
}
func (pc *PerspectiveCamera) AddEventListener(typ string, listener js.Value) {
	pc.Call("addEventListener", typ, listener)
}
func (pc *PerspectiveCamera) ApplyMatrix(matrix *Matrix4) {
	pc.Call("applyMatrix", matrix.JSValue())
}
func (pc *PerspectiveCamera) ApplyQuaternion(quaternion *Quaternion) Camera {
	return &PerspectiveCamera{Value: pc.Call("applyQuaternion", quaternion)}
}
func (pc *PerspectiveCamera) ClearViewOffset() {
	pc.Call("clearViewOffset")
}
func (pc *PerspectiveCamera) Clone(recursive bool) Camera {
	return &PerspectiveCamera{Value: pc.Call("clone", recursive)}
}
func (pc *PerspectiveCamera) Copy(source Camera, recursive bool) Camera {
	return &PerspectiveCamera{Value: pc.Call("copy", source.JSValue(), recursive)}
}
func (pc *PerspectiveCamera) DispatchEvent(event js.Value) {
	pc.Call("dispatchEvent", event)
}
func (pc *PerspectiveCamera) GetEffectiveFOV() float64 {
	return pc.Call("getEffectiveFOV").Float()
}
func (pc *PerspectiveCamera) GetFilmHeight() float64 {
	return pc.Call("getFilmHeight").Float()
}
func (pc *PerspectiveCamera) GetFilmWidth() float64 {
	return pc.Call("getFilmWidth").Float()
}
func (pc *PerspectiveCamera) GetFocalLength() float64 {
	return pc.Call("getFocalLength").Float()
}
func (pc *PerspectiveCamera) GetObjectById(id int) *Object3D {
	return &Object3D{Value: pc.Call("getObjectById", id)}
}
func (pc *PerspectiveCamera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: pc.Call("getObjectByName", name)}
}
func (pc *PerspectiveCamera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: pc.Call("getObjectByProperty", name, value)}
}
func (pc *PerspectiveCamera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("getWorldDirection", target)}
}
func (pc *PerspectiveCamera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("getWorldPosition", target)}
}
func (pc *PerspectiveCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: pc.Call("getWorldQuaternion", target)}
}
func (pc *PerspectiveCamera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("getWorldScale", target)}
}
func (pc *PerspectiveCamera) HasEventListener(typ string, listener js.Value) bool {
	return pc.Call("hasEventListener", typ, listener).Bool()
}
func (pc *PerspectiveCamera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("localToWorld", vector)}
}
func (pc *PerspectiveCamera) LookAt(vector *Vector3, y float64, z float64) {
	pc.Call("lookAt", vector, y, z)
}
func (pc *PerspectiveCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
	pc.Call("raycast", raycaster.JSValue(), intersects)
}
func (pc *PerspectiveCamera) Remove(object js.Value) Camera {
	return &PerspectiveCamera{Value: pc.Call("remove", object)}
}
func (pc *PerspectiveCamera) RemoveEventListener(typ string, listener js.Value) {
	pc.Call("removeEventListener", typ, listener)
}
func (pc *PerspectiveCamera) RotateOnAxis(axis *Vector3, angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateOnAxis", axis, angle)}
}
func (pc *PerspectiveCamera) RotateOnWorldAxis(axis *Vector3, angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateOnWorldAxis", axis, angle)}
}
func (pc *PerspectiveCamera) RotateX(angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateX", angle)}
}
func (pc *PerspectiveCamera) RotateY(angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateY", angle)}
}
func (pc *PerspectiveCamera) RotateZ(angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateZ", angle)}
}
func (pc *PerspectiveCamera) SetFocalLength(focalLength float64) {
	pc.Call("setFocalLength", focalLength)
}
func (pc *PerspectiveCamera) SetLens(focalLength float64, frameHeight float64) {
	pc.Call("setLens", focalLength, frameHeight)
}
func (pc *PerspectiveCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	pc.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (pc *PerspectiveCamera) SetRotationFromEuler(euler *Euler) {
	pc.Call("setRotationFromEuler", euler.JSValue())
}
func (pc *PerspectiveCamera) SetRotationFromMatrix(m *Matrix4) {
	pc.Call("setRotationFromMatrix", m.JSValue())
}
func (pc *PerspectiveCamera) SetRotationFromQuaternion(q *Quaternion) {
	pc.Call("setRotationFromQuaternion", q.JSValue())
}
func (pc *PerspectiveCamera) SetViewOffset(fullWidth float64, fullHeight float64, x float64, y float64, width float64, height float64) {
	pc.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}
func (pc *PerspectiveCamera) ToJSON(meta js.Value) js.Value {
	return pc.Call("toJSON", meta)
}
func (pc *PerspectiveCamera) TranslateOnAxis(axis *Vector3, distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateOnAxis", axis, distance)}
}
func (pc *PerspectiveCamera) TranslateX(distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateX", distance)}
}
func (pc *PerspectiveCamera) TranslateY(distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateY", distance)}
}
func (pc *PerspectiveCamera) TranslateZ(distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateZ", distance)}
}
func (pc *PerspectiveCamera) Traverse(callback js.Value) {
	pc.Call("traverse", callback)
}
func (pc *PerspectiveCamera) TraverseAncestors(callback js.Value) {
	pc.Call("traverseAncestors", callback)
}
func (pc *PerspectiveCamera) TraverseVisible(callback js.Value) {
	pc.Call("traverseVisible", callback)
}
func (pc *PerspectiveCamera) UpdateMatrix() {
	pc.Call("updateMatrix")
}
func (pc *PerspectiveCamera) UpdateMatrixWorld(force bool) {
	pc.Call("updateMatrixWorld", force)
}
func (pc *PerspectiveCamera) UpdateProjectionMatrix() {
	pc.Call("updateProjectionMatrix")
}
func (pc *PerspectiveCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pc *PerspectiveCamera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("worldToLocal", vector)}
}

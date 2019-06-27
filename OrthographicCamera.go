// Code generated from "cameras/OrthographicCamera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// OrthographicCamera extend: [Camera]
type OrthographicCamera struct {
	js.Value
}

func NewOrthographicCamera(left float64, right float64, top float64, bottom float64, near float64, far float64) *OrthographicCamera {
	return &OrthographicCamera{Value: get("OrthographicCamera").New(left, right, top, bottom, near, far)}
}
func (oc *OrthographicCamera) JSValue() js.Value {
	return oc.Value
}
func (oc *OrthographicCamera) Bottom() float64 {
	return oc.Get("bottom").Float()
}
func (oc *OrthographicCamera) SetBottom(v float64) {
	oc.Set("bottom", v)
}
func (oc *OrthographicCamera) CastShadow() bool {
	return oc.Get("castShadow").Bool()
}
func (oc *OrthographicCamera) SetCastShadow(v bool) {
	oc.Set("castShadow", v)
}
func (oc *OrthographicCamera) Children() js.Value {
	return oc.Get("children")
}
func (oc *OrthographicCamera) SetChildren(v js.Value) {
	oc.Set("children", v)
}
func (oc *OrthographicCamera) Far() float64 {
	return oc.Get("far").Float()
}
func (oc *OrthographicCamera) SetFar(v float64) {
	oc.Set("far", v)
}
func (oc *OrthographicCamera) FrustumCulled() bool {
	return oc.Get("frustumCulled").Bool()
}
func (oc *OrthographicCamera) SetFrustumCulled(v bool) {
	oc.Set("frustumCulled", v)
}
func (oc *OrthographicCamera) Id() int {
	return oc.Get("id").Int()
}
func (oc *OrthographicCamera) SetId(v int) {
	oc.Set("id", v)
}
func (oc *OrthographicCamera) IsCamera() bool {
	return oc.Get("isCamera").Bool()
}
func (oc *OrthographicCamera) SetIsCamera(v bool) {
	oc.Set("isCamera", v)
}
func (oc *OrthographicCamera) IsObject3D() bool {
	return oc.Get("isObject3D").Bool()
}
func (oc *OrthographicCamera) SetIsObject3D(v bool) {
	oc.Set("isObject3D", v)
}
func (oc *OrthographicCamera) IsOrthographicCamera() bool {
	return oc.Get("isOrthographicCamera").Bool()
}
func (oc *OrthographicCamera) SetIsOrthographicCamera(v bool) {
	oc.Set("isOrthographicCamera", v)
}
func (oc *OrthographicCamera) Layers() *Layers {
	return &Layers{Value: oc.Get("layers")}
}
func (oc *OrthographicCamera) SetLayers(v *Layers) {
	oc.Set("layers", v.Value)
}
func (oc *OrthographicCamera) Left() float64 {
	return oc.Get("left").Float()
}
func (oc *OrthographicCamera) SetLeft(v float64) {
	oc.Set("left", v)
}
func (oc *OrthographicCamera) Matrix() *Matrix4 {
	return &Matrix4{Value: oc.Get("matrix")}
}
func (oc *OrthographicCamera) SetMatrix(v *Matrix4) {
	oc.Set("matrix", v.Value)
}
func (oc *OrthographicCamera) MatrixAutoUpdate() bool {
	return oc.Get("matrixAutoUpdate").Bool()
}
func (oc *OrthographicCamera) SetMatrixAutoUpdate(v bool) {
	oc.Set("matrixAutoUpdate", v)
}
func (oc *OrthographicCamera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: oc.Get("matrixWorld")}
}
func (oc *OrthographicCamera) SetMatrixWorld(v *Matrix4) {
	oc.Set("matrixWorld", v.Value)
}
func (oc *OrthographicCamera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: oc.Get("matrixWorldInverse")}
}
func (oc *OrthographicCamera) SetMatrixWorldInverse(v *Matrix4) {
	oc.Set("matrixWorldInverse", v.Value)
}
func (oc *OrthographicCamera) MatrixWorldNeedsUpdate() bool {
	return oc.Get("matrixWorldNeedsUpdate").Bool()
}
func (oc *OrthographicCamera) SetMatrixWorldNeedsUpdate(v bool) {
	oc.Set("matrixWorldNeedsUpdate", v)
}
func (oc *OrthographicCamera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: oc.Get("modelViewMatrix")}
}
func (oc *OrthographicCamera) SetModelViewMatrix(v *Matrix4) {
	oc.Set("modelViewMatrix", v.Value)
}
func (oc *OrthographicCamera) Name() string {
	return oc.Get("name").String()
}
func (oc *OrthographicCamera) SetName(v string) {
	oc.Set("name", v)
}
func (oc *OrthographicCamera) Near() float64 {
	return oc.Get("near").Float()
}
func (oc *OrthographicCamera) SetNear(v float64) {
	oc.Set("near", v)
}
func (oc *OrthographicCamera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: oc.Get("normalMatrix")}
}
func (oc *OrthographicCamera) SetNormalMatrix(v *Matrix3) {
	oc.Set("normalMatrix", v.Value)
}
func (oc *OrthographicCamera) OnAfterRender() js.Value {
	return oc.Get("onAfterRender")
}
func (oc *OrthographicCamera) SetOnAfterRender(v js.Value) {
	oc.Set("onAfterRender", v)
}
func (oc *OrthographicCamera) OnBeforeRender() js.Value {
	return oc.Get("onBeforeRender")
}
func (oc *OrthographicCamera) SetOnBeforeRender(v js.Value) {
	oc.Set("onBeforeRender", v)
}
func (oc *OrthographicCamera) Parent() *Object3D {
	return &Object3D{Value: oc.Get("parent")}
}
func (oc *OrthographicCamera) SetParent(v *Object3D) {
	oc.Set("parent", v.Value)
}
func (oc *OrthographicCamera) Position() *Vector3 {
	return &Vector3{Value: oc.Get("position")}
}
func (oc *OrthographicCamera) SetPosition(v *Vector3) {
	oc.Set("position", v.Value)
}
func (oc *OrthographicCamera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: oc.Get("projectionMatrix")}
}
func (oc *OrthographicCamera) SetProjectionMatrix(v *Matrix4) {
	oc.Set("projectionMatrix", v.Value)
}
func (oc *OrthographicCamera) Quaternion() *Quaternion {
	return &Quaternion{Value: oc.Get("quaternion")}
}
func (oc *OrthographicCamera) SetQuaternion(v *Quaternion) {
	oc.Set("quaternion", v.Value)
}
func (oc *OrthographicCamera) ReceiveShadow() bool {
	return oc.Get("receiveShadow").Bool()
}
func (oc *OrthographicCamera) SetReceiveShadow(v bool) {
	oc.Set("receiveShadow", v)
}
func (oc *OrthographicCamera) RenderOrder() float64 {
	return oc.Get("renderOrder").Float()
}
func (oc *OrthographicCamera) SetRenderOrder(v float64) {
	oc.Set("renderOrder", v)
}
func (oc *OrthographicCamera) Right() float64 {
	return oc.Get("right").Float()
}
func (oc *OrthographicCamera) SetRight(v float64) {
	oc.Set("right", v)
}
func (oc *OrthographicCamera) Rotation() *Euler {
	return &Euler{Value: oc.Get("rotation")}
}
func (oc *OrthographicCamera) SetRotation(v *Euler) {
	oc.Set("rotation", v.Value)
}
func (oc *OrthographicCamera) Scale() *Vector3 {
	return &Vector3{Value: oc.Get("scale")}
}
func (oc *OrthographicCamera) SetScale(v *Vector3) {
	oc.Set("scale", v.Value)
}
func (oc *OrthographicCamera) Top() float64 {
	return oc.Get("top").Float()
}
func (oc *OrthographicCamera) SetTop(v float64) {
	oc.Set("top", v)
}
func (oc *OrthographicCamera) Type() string {
	return oc.Get("type").String()
}
func (oc *OrthographicCamera) SetType(v string) {
	oc.Set("type", v)
}
func (oc *OrthographicCamera) Up() *Vector3 {
	return &Vector3{Value: oc.Get("up")}
}
func (oc *OrthographicCamera) SetUp(v *Vector3) {
	oc.Set("up", v.Value)
}
func (oc *OrthographicCamera) UserData() js.Value {
	return oc.Get("userData")
}
func (oc *OrthographicCamera) SetUserData(v js.Value) {
	oc.Set("userData", v)
}
func (oc *OrthographicCamera) Uuid() string {
	return oc.Get("uuid").String()
}
func (oc *OrthographicCamera) SetUuid(v string) {
	oc.Set("uuid", v)
}
func (oc *OrthographicCamera) View() js.Value {
	return oc.Get("view")
}
func (oc *OrthographicCamera) SetView(v js.Value) {
	oc.Set("view", v)
}
func (oc *OrthographicCamera) Visible() bool {
	return oc.Get("visible").Bool()
}
func (oc *OrthographicCamera) SetVisible(v bool) {
	oc.Set("visible", v)
}
func (oc *OrthographicCamera) Zoom() float64 {
	return oc.Get("zoom").Float()
}
func (oc *OrthographicCamera) SetZoom(v float64) {
	oc.Set("zoom", v)
}
func (oc *OrthographicCamera) DefaultMatrixAutoUpdate() bool {
	return oc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (oc *OrthographicCamera) SetDefaultMatrixAutoUpdate(v bool) {
	oc.Set("DefaultMatrixAutoUpdate", v)
}
func (oc *OrthographicCamera) DefaultUp() *Vector3 {
	return &Vector3{Value: oc.Get("DefaultUp")}
}
func (oc *OrthographicCamera) SetDefaultUp(v *Vector3) {
	oc.Set("DefaultUp", v.Value)
}
func (oc *OrthographicCamera) Add(object js.Value) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("add", object)}
}
func (oc *OrthographicCamera) AddEventListener(typ string, listener js.Value) {
	oc.Call("addEventListener", typ, listener)
}
func (oc *OrthographicCamera) ApplyMatrix(matrix *Matrix4) {
	oc.Call("applyMatrix", matrix)
}
func (oc *OrthographicCamera) ApplyQuaternion(quaternion *Quaternion) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("applyQuaternion", quaternion)}
}
func (oc *OrthographicCamera) ClearViewOffset() {
	oc.Call("clearViewOffset")
}
func (oc *OrthographicCamera) Clone(recursive bool) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("clone", recursive)}
}
func (oc *OrthographicCamera) Copy(source Camera, recursive bool) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("copy", source.JSValue(), recursive)}
}
func (oc *OrthographicCamera) DispatchEvent(event js.Value) {
	oc.Call("dispatchEvent", event)
}
func (oc *OrthographicCamera) GetObjectById(id int) *Object3D {
	return &Object3D{Value: oc.Call("getObjectById", id)}
}
func (oc *OrthographicCamera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: oc.Call("getObjectByName", name)}
}
func (oc *OrthographicCamera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: oc.Call("getObjectByProperty", name, value)}
}
func (oc *OrthographicCamera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: oc.Call("getWorldDirection", target)}
}
func (oc *OrthographicCamera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: oc.Call("getWorldPosition", target)}
}
func (oc *OrthographicCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: oc.Call("getWorldQuaternion", target)}
}
func (oc *OrthographicCamera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: oc.Call("getWorldScale", target)}
}
func (oc *OrthographicCamera) HasEventListener(typ string, listener js.Value) bool {
	return oc.Call("hasEventListener", typ, listener).Bool()
}
func (oc *OrthographicCamera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: oc.Call("localToWorld", vector)}
}
func (oc *OrthographicCamera) LookAt(vector *Vector3, y float64, z float64) {
	oc.Call("lookAt", vector, y, z)
}
func (oc *OrthographicCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
	oc.Call("raycast", raycaster, intersects)
}
func (oc *OrthographicCamera) Remove(object js.Value) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("remove", object)}
}
func (oc *OrthographicCamera) RemoveEventListener(typ string, listener js.Value) {
	oc.Call("removeEventListener", typ, listener)
}
func (oc *OrthographicCamera) RotateOnAxis(axis *Vector3, angle float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("rotateOnAxis", axis, angle)}
}
func (oc *OrthographicCamera) RotateOnWorldAxis(axis *Vector3, angle float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("rotateOnWorldAxis", axis, angle)}
}
func (oc *OrthographicCamera) RotateX(angle float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("rotateX", angle)}
}
func (oc *OrthographicCamera) RotateY(angle float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("rotateY", angle)}
}
func (oc *OrthographicCamera) RotateZ(angle float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("rotateZ", angle)}
}
func (oc *OrthographicCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	oc.Call("setRotationFromAxisAngle", axis, angle)
}
func (oc *OrthographicCamera) SetRotationFromEuler(euler *Euler) {
	oc.Call("setRotationFromEuler", euler)
}
func (oc *OrthographicCamera) SetRotationFromMatrix(m *Matrix4) {
	oc.Call("setRotationFromMatrix", m)
}
func (oc *OrthographicCamera) SetRotationFromQuaternion(q *Quaternion) {
	oc.Call("setRotationFromQuaternion", q)
}
func (oc *OrthographicCamera) SetViewOffset(fullWidth float64, fullHeight float64, offsetX float64, offsetY float64, width float64, height float64) {
	oc.Call("setViewOffset", fullWidth, fullHeight, offsetX, offsetY, width, height)
}
func (oc *OrthographicCamera) ToJSON(meta js.Value) js.Value {
	return oc.Call("toJSON", meta)
}
func (oc *OrthographicCamera) TranslateOnAxis(axis *Vector3, distance float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("translateOnAxis", axis, distance)}
}
func (oc *OrthographicCamera) TranslateX(distance float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("translateX", distance)}
}
func (oc *OrthographicCamera) TranslateY(distance float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("translateY", distance)}
}
func (oc *OrthographicCamera) TranslateZ(distance float64) *OrthographicCamera {
	return &OrthographicCamera{Value: oc.Call("translateZ", distance)}
}
func (oc *OrthographicCamera) Traverse(callback js.Value) {
	oc.Call("traverse", callback)
}
func (oc *OrthographicCamera) TraverseAncestors(callback js.Value) {
	oc.Call("traverseAncestors", callback)
}
func (oc *OrthographicCamera) TraverseVisible(callback js.Value) {
	oc.Call("traverseVisible", callback)
}
func (oc *OrthographicCamera) UpdateMatrix() {
	oc.Call("updateMatrix")
}
func (oc *OrthographicCamera) UpdateMatrixWorld(force bool) {
	oc.Call("updateMatrixWorld", force)
}
func (oc *OrthographicCamera) UpdateProjectionMatrix() {
	oc.Call("updateProjectionMatrix")
}
func (oc *OrthographicCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	oc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (oc *OrthographicCamera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: oc.Call("worldToLocal", vector)}
}

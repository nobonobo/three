// Code generated from "cameras/Camera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Camera struct {
	js.Value
}

func NewCamera() *Camera {
	return &Camera{Value: get("Camera").New()}
}
func (cc *Camera) CastShadow() bool {
	return cc.Get("castShadow").Bool()
}
func (cc *Camera) SetCastShadow(v bool) {
	cc.Set("castShadow", v)
}
func (cc *Camera) Children() js.Value {
	return cc.Get("children")
}
func (cc *Camera) SetChildren(v js.Value) {
	cc.Set("children", v)
}
func (cc *Camera) FrustumCulled() bool {
	return cc.Get("frustumCulled").Bool()
}
func (cc *Camera) SetFrustumCulled(v bool) {
	cc.Set("frustumCulled", v)
}
func (cc *Camera) Id() int {
	return cc.Get("id").Int()
}
func (cc *Camera) SetId(v int) {
	cc.Set("id", v)
}
func (cc *Camera) IsCamera() bool {
	return cc.Get("isCamera").Bool()
}
func (cc *Camera) SetIsCamera(v bool) {
	cc.Set("isCamera", v)
}
func (cc *Camera) IsObject3D() bool {
	return cc.Get("isObject3D").Bool()
}
func (cc *Camera) SetIsObject3D(v bool) {
	cc.Set("isObject3D", v)
}
func (cc *Camera) Layers() *Layers {
	return &Layers{Value: cc.Get("layers")}
}
func (cc *Camera) SetLayers(v *Layers) {
	cc.Set("layers", v)
}
func (cc *Camera) Matrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrix")}
}
func (cc *Camera) SetMatrix(v *Matrix4) {
	cc.Set("matrix", v)
}
func (cc *Camera) MatrixAutoUpdate() bool {
	return cc.Get("matrixAutoUpdate").Bool()
}
func (cc *Camera) SetMatrixAutoUpdate(v bool) {
	cc.Set("matrixAutoUpdate", v)
}
func (cc *Camera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorld")}
}
func (cc *Camera) SetMatrixWorld(v *Matrix4) {
	cc.Set("matrixWorld", v)
}
func (cc *Camera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorldInverse")}
}
func (cc *Camera) SetMatrixWorldInverse(v *Matrix4) {
	cc.Set("matrixWorldInverse", v)
}
func (cc *Camera) MatrixWorldNeedsUpdate() bool {
	return cc.Get("matrixWorldNeedsUpdate").Bool()
}
func (cc *Camera) SetMatrixWorldNeedsUpdate(v bool) {
	cc.Set("matrixWorldNeedsUpdate", v)
}
func (cc *Camera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("modelViewMatrix")}
}
func (cc *Camera) SetModelViewMatrix(v *Matrix4) {
	cc.Set("modelViewMatrix", v)
}
func (cc *Camera) Name() string {
	return cc.Get("name").String()
}
func (cc *Camera) SetName(v string) {
	cc.Set("name", v)
}
func (cc *Camera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: cc.Get("normalMatrix")}
}
func (cc *Camera) SetNormalMatrix(v *Matrix3) {
	cc.Set("normalMatrix", v)
}
func (cc *Camera) OnAfterRender() js.Value {
	return cc.Get("onAfterRender")
}
func (cc *Camera) SetOnAfterRender(v js.Value) {
	cc.Set("onAfterRender", v)
}
func (cc *Camera) OnBeforeRender() js.Value {
	return cc.Get("onBeforeRender")
}
func (cc *Camera) SetOnBeforeRender(v js.Value) {
	cc.Set("onBeforeRender", v)
}
func (cc *Camera) Parent() *Object3D {
	return &Object3D{Value: cc.Get("parent")}
}
func (cc *Camera) SetParent(v *Object3D) {
	cc.Set("parent", v)
}
func (cc *Camera) Position() *Vector3 {
	return &Vector3{Value: cc.Get("position")}
}
func (cc *Camera) SetPosition(v *Vector3) {
	cc.Set("position", v)
}
func (cc *Camera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("projectionMatrix")}
}
func (cc *Camera) SetProjectionMatrix(v *Matrix4) {
	cc.Set("projectionMatrix", v)
}
func (cc *Camera) Quaternion() *Quaternion {
	return &Quaternion{Value: cc.Get("quaternion")}
}
func (cc *Camera) SetQuaternion(v *Quaternion) {
	cc.Set("quaternion", v)
}
func (cc *Camera) ReceiveShadow() bool {
	return cc.Get("receiveShadow").Bool()
}
func (cc *Camera) SetReceiveShadow(v bool) {
	cc.Set("receiveShadow", v)
}
func (cc *Camera) RenderOrder() float64 {
	return cc.Get("renderOrder").Float()
}
func (cc *Camera) SetRenderOrder(v float64) {
	cc.Set("renderOrder", v)
}
func (cc *Camera) Rotation() *Euler {
	return &Euler{Value: cc.Get("rotation")}
}
func (cc *Camera) SetRotation(v *Euler) {
	cc.Set("rotation", v)
}
func (cc *Camera) Scale() *Vector3 {
	return &Vector3{Value: cc.Get("scale")}
}
func (cc *Camera) SetScale(v *Vector3) {
	cc.Set("scale", v)
}
func (cc *Camera) Type() string {
	return cc.Get("type").String()
}
func (cc *Camera) SetType(v string) {
	cc.Set("type", v)
}
func (cc *Camera) Up() *Vector3 {
	return &Vector3{Value: cc.Get("up")}
}
func (cc *Camera) SetUp(v *Vector3) {
	cc.Set("up", v)
}
func (cc *Camera) UserData() js.Value {
	return cc.Get("userData")
}
func (cc *Camera) SetUserData(v js.Value) {
	cc.Set("userData", v)
}
func (cc *Camera) Uuid() string {
	return cc.Get("uuid").String()
}
func (cc *Camera) SetUuid(v string) {
	cc.Set("uuid", v)
}
func (cc *Camera) Visible() bool {
	return cc.Get("visible").Bool()
}
func (cc *Camera) SetVisible(v bool) {
	cc.Set("visible", v)
}
func (cc *Camera) DefaultMatrixAutoUpdate() bool {
	return cc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (cc *Camera) SetDefaultMatrixAutoUpdate(v bool) {
	cc.Set("DefaultMatrixAutoUpdate", v)
}
func (cc *Camera) DefaultUp() *Vector3 {
	return &Vector3{Value: cc.Get("DefaultUp")}
}
func (cc *Camera) SetDefaultUp(v *Vector3) {
	cc.Set("DefaultUp", v)
}
func (cc *Camera) Add(object js.Value) *Camera {
	return &Camera{Value: cc.Call("add", object)}
}
func (cc *Camera) AddEventListener(typ string, listener js.Value) {
	cc.Call("addEventListener", typ, listener)
}
func (cc *Camera) ApplyMatrix(matrix *Matrix4) {
	cc.Call("applyMatrix", matrix)
}
func (cc *Camera) ApplyQuaternion(quaternion *Quaternion) *Camera {
	return &Camera{Value: cc.Call("applyQuaternion", quaternion)}
}
func (cc *Camera) Clone(recursive bool) *Camera {
	return &Camera{Value: cc.Call("clone", recursive)}
}
func (cc *Camera) Copy(source *Camera, recursive bool) *Camera {
	return &Camera{Value: cc.Call("copy", source, recursive)}
}
func (cc *Camera) DispatchEvent(event js.Value) {
	cc.Call("dispatchEvent", event)
}
func (cc *Camera) GetObjectById(id int) *Object3D {
	return &Object3D{Value: cc.Call("getObjectById", id)}
}
func (cc *Camera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByName", name)}
}
func (cc *Camera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByProperty", name, value)}
}
func (cc *Camera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldDirection", target)}
}
func (cc *Camera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldPosition", target)}
}
func (cc *Camera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: cc.Call("getWorldQuaternion", target)}
}
func (cc *Camera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldScale", target)}
}
func (cc *Camera) HasEventListener(typ string, listener js.Value) bool {
	return cc.Call("hasEventListener", typ, listener).Bool()
}
func (cc *Camera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("localToWorld", vector)}
}
func (cc *Camera) LookAt(vector *Vector3, y float64, z float64) {
	cc.Call("lookAt", vector, y, z)
}
func (cc *Camera) Raycast(raycaster *Raycaster, intersects js.Value) {
	cc.Call("raycast", raycaster, intersects)
}
func (cc *Camera) Remove(object js.Value) *Camera {
	return &Camera{Value: cc.Call("remove", object)}
}
func (cc *Camera) RemoveEventListener(typ string, listener js.Value) {
	cc.Call("removeEventListener", typ, listener)
}
func (cc *Camera) RotateOnAxis(axis *Vector3, angle float64) *Camera {
	return &Camera{Value: cc.Call("rotateOnAxis", axis, angle)}
}
func (cc *Camera) RotateOnWorldAxis(axis *Vector3, angle float64) *Camera {
	return &Camera{Value: cc.Call("rotateOnWorldAxis", axis, angle)}
}
func (cc *Camera) RotateX(angle float64) *Camera {
	return &Camera{Value: cc.Call("rotateX", angle)}
}
func (cc *Camera) RotateY(angle float64) *Camera {
	return &Camera{Value: cc.Call("rotateY", angle)}
}
func (cc *Camera) RotateZ(angle float64) *Camera {
	return &Camera{Value: cc.Call("rotateZ", angle)}
}
func (cc *Camera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	cc.Call("setRotationFromAxisAngle", axis, angle)
}
func (cc *Camera) SetRotationFromEuler(euler *Euler) {
	cc.Call("setRotationFromEuler", euler)
}
func (cc *Camera) SetRotationFromMatrix(m *Matrix4) {
	cc.Call("setRotationFromMatrix", m)
}
func (cc *Camera) SetRotationFromQuaternion(q *Quaternion) {
	cc.Call("setRotationFromQuaternion", q)
}
func (cc *Camera) ToJSON(meta js.Value) js.Value {
	return cc.Call("toJSON", meta)
}
func (cc *Camera) TranslateOnAxis(axis *Vector3, distance float64) *Camera {
	return &Camera{Value: cc.Call("translateOnAxis", axis, distance)}
}
func (cc *Camera) TranslateX(distance float64) *Camera {
	return &Camera{Value: cc.Call("translateX", distance)}
}
func (cc *Camera) TranslateY(distance float64) *Camera {
	return &Camera{Value: cc.Call("translateY", distance)}
}
func (cc *Camera) TranslateZ(distance float64) *Camera {
	return &Camera{Value: cc.Call("translateZ", distance)}
}
func (cc *Camera) Traverse(callback js.Value) {
	cc.Call("traverse", callback)
}
func (cc *Camera) TraverseAncestors(callback js.Value) {
	cc.Call("traverseAncestors", callback)
}
func (cc *Camera) TraverseVisible(callback js.Value) {
	cc.Call("traverseVisible", callback)
}
func (cc *Camera) UpdateMatrix() {
	cc.Call("updateMatrix")
}
func (cc *Camera) UpdateMatrixWorld(force bool) {
	cc.Call("updateMatrixWorld", force)
}
func (cc *Camera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	cc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (cc *Camera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("worldToLocal", vector)}
}

// Code generated from "scenes/Scene.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Scene struct {
	js.Value
}

func NewScene() *Scene {
	return &Scene{Value: get("Scene").New()}
}
func (ss *Scene) AutoUpdate() bool {
	return ss.Get("autoUpdate").Bool()
}
func (ss *Scene) SetAutoUpdate(v bool) {
	ss.Set("autoUpdate", v)
}
func (ss *Scene) Background() js.Value {
	return ss.Get("background")
}
func (ss *Scene) SetBackground(v js.Value) {
	ss.Set("background", v)
}
func (ss *Scene) CastShadow() bool {
	return ss.Get("castShadow").Bool()
}
func (ss *Scene) SetCastShadow(v bool) {
	ss.Set("castShadow", v)
}
func (ss *Scene) Children() js.Value {
	return ss.Get("children")
}
func (ss *Scene) SetChildren(v js.Value) {
	ss.Set("children", v)
}
func (ss *Scene) Fog() js.Value {
	return ss.Get("fog")
}
func (ss *Scene) SetFog(v js.Value) {
	ss.Set("fog", v)
}
func (ss *Scene) FrustumCulled() bool {
	return ss.Get("frustumCulled").Bool()
}
func (ss *Scene) SetFrustumCulled(v bool) {
	ss.Set("frustumCulled", v)
}
func (ss *Scene) Id() int {
	return ss.Get("id").Int()
}
func (ss *Scene) SetId(v int) {
	ss.Set("id", v)
}
func (ss *Scene) IsObject3D() bool {
	return ss.Get("isObject3D").Bool()
}
func (ss *Scene) SetIsObject3D(v bool) {
	ss.Set("isObject3D", v)
}
func (ss *Scene) Layers() *Layers {
	return &Layers{Value: ss.Get("layers")}
}
func (ss *Scene) SetLayers(v *Layers) {
	ss.Set("layers", v)
}
func (ss *Scene) Matrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrix")}
}
func (ss *Scene) SetMatrix(v *Matrix4) {
	ss.Set("matrix", v)
}
func (ss *Scene) MatrixAutoUpdate() bool {
	return ss.Get("matrixAutoUpdate").Bool()
}
func (ss *Scene) SetMatrixAutoUpdate(v bool) {
	ss.Set("matrixAutoUpdate", v)
}
func (ss *Scene) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrixWorld")}
}
func (ss *Scene) SetMatrixWorld(v *Matrix4) {
	ss.Set("matrixWorld", v)
}
func (ss *Scene) MatrixWorldNeedsUpdate() bool {
	return ss.Get("matrixWorldNeedsUpdate").Bool()
}
func (ss *Scene) SetMatrixWorldNeedsUpdate(v bool) {
	ss.Set("matrixWorldNeedsUpdate", v)
}
func (ss *Scene) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("modelViewMatrix")}
}
func (ss *Scene) SetModelViewMatrix(v *Matrix4) {
	ss.Set("modelViewMatrix", v)
}
func (ss *Scene) Name() string {
	return ss.Get("name").String()
}
func (ss *Scene) SetName(v string) {
	ss.Set("name", v)
}
func (ss *Scene) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ss.Get("normalMatrix")}
}
func (ss *Scene) SetNormalMatrix(v *Matrix3) {
	ss.Set("normalMatrix", v)
}
func (ss *Scene) OnAfterRender() js.Value {
	return ss.Get("onAfterRender")
}
func (ss *Scene) SetOnAfterRender(v js.Value) {
	ss.Set("onAfterRender", v)
}
func (ss *Scene) OnBeforeRender() js.Value {
	return ss.Get("onBeforeRender")
}
func (ss *Scene) SetOnBeforeRender(v js.Value) {
	ss.Set("onBeforeRender", v)
}
func (ss *Scene) OverrideMaterial() *Material {
	return &Material{Value: ss.Get("overrideMaterial")}
}
func (ss *Scene) SetOverrideMaterial(v *Material) {
	ss.Set("overrideMaterial", v)
}
func (ss *Scene) Parent() *Object3D {
	return &Object3D{Value: ss.Get("parent")}
}
func (ss *Scene) SetParent(v *Object3D) {
	ss.Set("parent", v)
}
func (ss *Scene) Position() *Vector3 {
	return &Vector3{Value: ss.Get("position")}
}
func (ss *Scene) SetPosition(v *Vector3) {
	ss.Set("position", v)
}
func (ss *Scene) Quaternion() *Quaternion {
	return &Quaternion{Value: ss.Get("quaternion")}
}
func (ss *Scene) SetQuaternion(v *Quaternion) {
	ss.Set("quaternion", v)
}
func (ss *Scene) ReceiveShadow() bool {
	return ss.Get("receiveShadow").Bool()
}
func (ss *Scene) SetReceiveShadow(v bool) {
	ss.Set("receiveShadow", v)
}
func (ss *Scene) RenderOrder() float64 {
	return ss.Get("renderOrder").Float()
}
func (ss *Scene) SetRenderOrder(v float64) {
	ss.Set("renderOrder", v)
}
func (ss *Scene) Rotation() *Euler {
	return &Euler{Value: ss.Get("rotation")}
}
func (ss *Scene) SetRotation(v *Euler) {
	ss.Set("rotation", v)
}
func (ss *Scene) Scale() *Vector3 {
	return &Vector3{Value: ss.Get("scale")}
}
func (ss *Scene) SetScale(v *Vector3) {
	ss.Set("scale", v)
}
func (ss *Scene) Type() string {
	return ss.Get("type").String()
}
func (ss *Scene) SetType(v string) {
	ss.Set("type", v)
}
func (ss *Scene) Up() *Vector3 {
	return &Vector3{Value: ss.Get("up")}
}
func (ss *Scene) SetUp(v *Vector3) {
	ss.Set("up", v)
}
func (ss *Scene) UserData() js.Value {
	return ss.Get("userData")
}
func (ss *Scene) SetUserData(v js.Value) {
	ss.Set("userData", v)
}
func (ss *Scene) Uuid() string {
	return ss.Get("uuid").String()
}
func (ss *Scene) SetUuid(v string) {
	ss.Set("uuid", v)
}
func (ss *Scene) Visible() bool {
	return ss.Get("visible").Bool()
}
func (ss *Scene) SetVisible(v bool) {
	ss.Set("visible", v)
}
func (ss *Scene) DefaultMatrixAutoUpdate() bool {
	return ss.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ss *Scene) SetDefaultMatrixAutoUpdate(v bool) {
	ss.Set("DefaultMatrixAutoUpdate", v)
}
func (ss *Scene) DefaultUp() *Vector3 {
	return &Vector3{Value: ss.Get("DefaultUp")}
}
func (ss *Scene) SetDefaultUp(v *Vector3) {
	ss.Set("DefaultUp", v)
}
func (ss *Scene) Add(object js.Value) *Scene {
	return &Scene{Value: ss.Call("add", object)}
}
func (ss *Scene) AddEventListener(typ string, listener js.Value) {
	ss.Call("addEventListener", typ, listener)
}
func (ss *Scene) ApplyMatrix(matrix *Matrix4) {
	ss.Call("applyMatrix", matrix)
}
func (ss *Scene) ApplyQuaternion(quaternion *Quaternion) *Scene {
	return &Scene{Value: ss.Call("applyQuaternion", quaternion)}
}
func (ss *Scene) Clone(recursive bool) *Scene {
	return &Scene{Value: ss.Call("clone", recursive)}
}
func (ss *Scene) Copy(source *Scene, recursive bool) *Scene {
	return &Scene{Value: ss.Call("copy", source, recursive)}
}
func (ss *Scene) DispatchEvent(event js.Value) {
	ss.Call("dispatchEvent", event)
}
func (ss *Scene) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ss.Call("getObjectById", id)}
}
func (ss *Scene) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByName", name)}
}
func (ss *Scene) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByProperty", name, value)}
}
func (ss *Scene) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldDirection", target)}
}
func (ss *Scene) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldPosition", target)}
}
func (ss *Scene) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ss.Call("getWorldQuaternion", target)}
}
func (ss *Scene) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldScale", target)}
}
func (ss *Scene) HasEventListener(typ string, listener js.Value) bool {
	return ss.Call("hasEventListener", typ, listener).Bool()
}
func (ss *Scene) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("localToWorld", vector)}
}
func (ss *Scene) LookAt(vector *Vector3, y float64, z float64) {
	ss.Call("lookAt", vector, y, z)
}
func (ss *Scene) Raycast(raycaster *Raycaster, intersects js.Value) {
	ss.Call("raycast", raycaster, intersects)
}
func (ss *Scene) Remove(object js.Value) *Scene {
	return &Scene{Value: ss.Call("remove", object)}
}
func (ss *Scene) RemoveEventListener(typ string, listener js.Value) {
	ss.Call("removeEventListener", typ, listener)
}
func (ss *Scene) RotateOnAxis(axis *Vector3, angle float64) *Scene {
	return &Scene{Value: ss.Call("rotateOnAxis", axis, angle)}
}
func (ss *Scene) RotateOnWorldAxis(axis *Vector3, angle float64) *Scene {
	return &Scene{Value: ss.Call("rotateOnWorldAxis", axis, angle)}
}
func (ss *Scene) RotateX(angle float64) *Scene {
	return &Scene{Value: ss.Call("rotateX", angle)}
}
func (ss *Scene) RotateY(angle float64) *Scene {
	return &Scene{Value: ss.Call("rotateY", angle)}
}
func (ss *Scene) RotateZ(angle float64) *Scene {
	return &Scene{Value: ss.Call("rotateZ", angle)}
}
func (ss *Scene) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ss.Call("setRotationFromAxisAngle", axis, angle)
}
func (ss *Scene) SetRotationFromEuler(euler *Euler) {
	ss.Call("setRotationFromEuler", euler)
}
func (ss *Scene) SetRotationFromMatrix(m *Matrix4) {
	ss.Call("setRotationFromMatrix", m)
}
func (ss *Scene) SetRotationFromQuaternion(q *Quaternion) {
	ss.Call("setRotationFromQuaternion", q)
}
func (ss *Scene) ToJSON(meta js.Value) js.Value {
	return ss.Call("toJSON", meta)
}
func (ss *Scene) TranslateOnAxis(axis *Vector3, distance float64) *Scene {
	return &Scene{Value: ss.Call("translateOnAxis", axis, distance)}
}
func (ss *Scene) TranslateX(distance float64) *Scene {
	return &Scene{Value: ss.Call("translateX", distance)}
}
func (ss *Scene) TranslateY(distance float64) *Scene {
	return &Scene{Value: ss.Call("translateY", distance)}
}
func (ss *Scene) TranslateZ(distance float64) *Scene {
	return &Scene{Value: ss.Call("translateZ", distance)}
}
func (ss *Scene) Traverse(callback js.Value) {
	ss.Call("traverse", callback)
}
func (ss *Scene) TraverseAncestors(callback js.Value) {
	ss.Call("traverseAncestors", callback)
}
func (ss *Scene) TraverseVisible(callback js.Value) {
	ss.Call("traverseVisible", callback)
}
func (ss *Scene) UpdateMatrix() {
	ss.Call("updateMatrix")
}
func (ss *Scene) UpdateMatrixWorld(force bool) {
	ss.Call("updateMatrixWorld", force)
}
func (ss *Scene) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ss.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ss *Scene) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("worldToLocal", vector)}
}

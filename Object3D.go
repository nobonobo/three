// Code generated from "core/Object3D.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

func Object3DIdCount() int {
	return get("Object3DIdCount").Int()
}
func SetObject3DIdCount(v int) {
	set("Object3DIdCount", v)
}

// Object3D extend: [EventDispatcher]
type Object3D struct {
	js.Value
}

func NewObject3D() *Object3D {
	return &Object3D{Value: get("Object3D").New()}
}
func (od *Object3D) JSValue() js.Value {
	return od.Value
}
func (od *Object3D) CastShadow() bool {
	return od.Get("castShadow").Bool()
}
func (od *Object3D) SetCastShadow(v bool) {
	od.Set("castShadow", v)
}
func (od *Object3D) Children() js.Value {
	return od.Get("children")
}
func (od *Object3D) SetChildren(v js.Value) {
	od.Set("children", v)
}
func (od *Object3D) FrustumCulled() bool {
	return od.Get("frustumCulled").Bool()
}
func (od *Object3D) SetFrustumCulled(v bool) {
	od.Set("frustumCulled", v)
}
func (od *Object3D) Id() int {
	return od.Get("id").Int()
}
func (od *Object3D) SetId(v int) {
	od.Set("id", v)
}
func (od *Object3D) IsObject3D() bool {
	return od.Get("isObject3D").Bool()
}
func (od *Object3D) SetIsObject3D(v bool) {
	od.Set("isObject3D", v)
}
func (od *Object3D) Layers() *Layers {
	return &Layers{Value: od.Get("layers")}
}
func (od *Object3D) SetLayers(v *Layers) {
	od.Set("layers", v.JSValue())
}
func (od *Object3D) Matrix() *Matrix4 {
	return &Matrix4{Value: od.Get("matrix")}
}
func (od *Object3D) SetMatrix(v *Matrix4) {
	od.Set("matrix", v.JSValue())
}
func (od *Object3D) MatrixAutoUpdate() bool {
	return od.Get("matrixAutoUpdate").Bool()
}
func (od *Object3D) SetMatrixAutoUpdate(v bool) {
	od.Set("matrixAutoUpdate", v)
}
func (od *Object3D) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: od.Get("matrixWorld")}
}
func (od *Object3D) SetMatrixWorld(v *Matrix4) {
	od.Set("matrixWorld", v.JSValue())
}
func (od *Object3D) MatrixWorldNeedsUpdate() bool {
	return od.Get("matrixWorldNeedsUpdate").Bool()
}
func (od *Object3D) SetMatrixWorldNeedsUpdate(v bool) {
	od.Set("matrixWorldNeedsUpdate", v)
}
func (od *Object3D) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: od.Get("modelViewMatrix")}
}
func (od *Object3D) SetModelViewMatrix(v *Matrix4) {
	od.Set("modelViewMatrix", v.JSValue())
}
func (od *Object3D) Name() string {
	return od.Get("name").String()
}
func (od *Object3D) SetName(v string) {
	od.Set("name", v)
}
func (od *Object3D) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: od.Get("normalMatrix")}
}
func (od *Object3D) SetNormalMatrix(v *Matrix3) {
	od.Set("normalMatrix", v.JSValue())
}
func (od *Object3D) OnAfterRender() js.Value {
	return od.Get("onAfterRender")
}
func (od *Object3D) SetOnAfterRender(v js.Value) {
	od.Set("onAfterRender", v)
}
func (od *Object3D) OnBeforeRender() js.Value {
	return od.Get("onBeforeRender")
}
func (od *Object3D) SetOnBeforeRender(v js.Value) {
	od.Set("onBeforeRender", v)
}
func (od *Object3D) Parent() *Object3D {
	return &Object3D{Value: od.Get("parent")}
}
func (od *Object3D) SetParent(v *Object3D) {
	od.Set("parent", v.JSValue())
}
func (od *Object3D) Position() *Vector3 {
	return &Vector3{Value: od.Get("position")}
}
func (od *Object3D) SetPosition(v *Vector3) {
	od.Set("position", v.JSValue())
}
func (od *Object3D) Quaternion() *Quaternion {
	return &Quaternion{Value: od.Get("quaternion")}
}
func (od *Object3D) SetQuaternion(v *Quaternion) {
	od.Set("quaternion", v.JSValue())
}
func (od *Object3D) ReceiveShadow() bool {
	return od.Get("receiveShadow").Bool()
}
func (od *Object3D) SetReceiveShadow(v bool) {
	od.Set("receiveShadow", v)
}
func (od *Object3D) RenderOrder() float64 {
	return od.Get("renderOrder").Float()
}
func (od *Object3D) SetRenderOrder(v float64) {
	od.Set("renderOrder", v)
}
func (od *Object3D) Rotation() *Euler {
	return &Euler{Value: od.Get("rotation")}
}
func (od *Object3D) SetRotation(v *Euler) {
	od.Set("rotation", v.JSValue())
}
func (od *Object3D) Scale() *Vector3 {
	return &Vector3{Value: od.Get("scale")}
}
func (od *Object3D) SetScale(v *Vector3) {
	od.Set("scale", v.JSValue())
}
func (od *Object3D) Type() string {
	return od.Get("type").String()
}
func (od *Object3D) SetType(v string) {
	od.Set("type", v)
}
func (od *Object3D) Up() *Vector3 {
	return &Vector3{Value: od.Get("up")}
}
func (od *Object3D) SetUp(v *Vector3) {
	od.Set("up", v.JSValue())
}
func (od *Object3D) UserData() js.Value {
	return od.Get("userData")
}
func (od *Object3D) SetUserData(v js.Value) {
	od.Set("userData", v)
}
func (od *Object3D) Uuid() string {
	return od.Get("uuid").String()
}
func (od *Object3D) SetUuid(v string) {
	od.Set("uuid", v)
}
func (od *Object3D) Visible() bool {
	return od.Get("visible").Bool()
}
func (od *Object3D) SetVisible(v bool) {
	od.Set("visible", v)
}
func (od *Object3D) DefaultMatrixAutoUpdate() bool {
	return od.Get("DefaultMatrixAutoUpdate").Bool()
}
func (od *Object3D) SetDefaultMatrixAutoUpdate(v bool) {
	od.Set("DefaultMatrixAutoUpdate", v)
}
func (od *Object3D) DefaultUp() *Vector3 {
	return &Vector3{Value: od.Get("DefaultUp")}
}
func (od *Object3D) SetDefaultUp(v *Vector3) {
	od.Set("DefaultUp", v.JSValue())
}
func (od *Object3D) Add(object js.Value) *Object3D {
	return &Object3D{Value: od.Call("add", object)}
}
func (od *Object3D) AddEventListener(typ string, listener js.Value) {
	od.Call("addEventListener", typ, listener)
}
func (od *Object3D) ApplyMatrix(matrix *Matrix4) {
	od.Call("applyMatrix", matrix.JSValue())
}
func (od *Object3D) ApplyQuaternion(quaternion *Quaternion) *Object3D {
	return &Object3D{Value: od.Call("applyQuaternion", quaternion)}
}
func (od *Object3D) Clone(recursive bool) *Object3D {
	return &Object3D{Value: od.Call("clone", recursive)}
}
func (od *Object3D) Copy(source *Object3D, recursive bool) *Object3D {
	return &Object3D{Value: od.Call("copy", source, recursive)}
}
func (od *Object3D) DispatchEvent(event js.Value) {
	od.Call("dispatchEvent", event)
}
func (od *Object3D) GetObjectById(id int) *Object3D {
	return &Object3D{Value: od.Call("getObjectById", id)}
}
func (od *Object3D) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: od.Call("getObjectByName", name)}
}
func (od *Object3D) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: od.Call("getObjectByProperty", name, value)}
}
func (od *Object3D) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldDirection", target)}
}
func (od *Object3D) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldPosition", target)}
}
func (od *Object3D) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: od.Call("getWorldQuaternion", target)}
}
func (od *Object3D) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldScale", target)}
}
func (od *Object3D) HasEventListener(typ string, listener js.Value) bool {
	return od.Call("hasEventListener", typ, listener).Bool()
}
func (od *Object3D) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("localToWorld", vector)}
}
func (od *Object3D) LookAt(vector *Vector3, y float64, z float64) {
	od.Call("lookAt", vector, y, z)
}
func (od *Object3D) Raycast(raycaster *Raycaster, intersects js.Value) {
	od.Call("raycast", raycaster.JSValue(), intersects)
}
func (od *Object3D) Remove(object js.Value) *Object3D {
	return &Object3D{Value: od.Call("remove", object)}
}
func (od *Object3D) RemoveEventListener(typ string, listener js.Value) {
	od.Call("removeEventListener", typ, listener)
}
func (od *Object3D) RotateOnAxis(axis *Vector3, angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateOnAxis", axis, angle)}
}
func (od *Object3D) RotateOnWorldAxis(axis *Vector3, angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateOnWorldAxis", axis, angle)}
}
func (od *Object3D) RotateX(angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateX", angle)}
}
func (od *Object3D) RotateY(angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateY", angle)}
}
func (od *Object3D) RotateZ(angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateZ", angle)}
}
func (od *Object3D) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	od.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (od *Object3D) SetRotationFromEuler(euler *Euler) {
	od.Call("setRotationFromEuler", euler.JSValue())
}
func (od *Object3D) SetRotationFromMatrix(m *Matrix4) {
	od.Call("setRotationFromMatrix", m.JSValue())
}
func (od *Object3D) SetRotationFromQuaternion(q *Quaternion) {
	od.Call("setRotationFromQuaternion", q.JSValue())
}
func (od *Object3D) ToJSON(meta js.Value) js.Value {
	return od.Call("toJSON", meta)
}
func (od *Object3D) TranslateOnAxis(axis *Vector3, distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateOnAxis", axis, distance)}
}
func (od *Object3D) TranslateX(distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateX", distance)}
}
func (od *Object3D) TranslateY(distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateY", distance)}
}
func (od *Object3D) TranslateZ(distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateZ", distance)}
}
func (od *Object3D) Traverse(callback js.Value) {
	od.Call("traverse", callback)
}
func (od *Object3D) TraverseAncestors(callback js.Value) {
	od.Call("traverseAncestors", callback)
}
func (od *Object3D) TraverseVisible(callback js.Value) {
	od.Call("traverseVisible", callback)
}
func (od *Object3D) UpdateMatrix() {
	od.Call("updateMatrix")
}
func (od *Object3D) UpdateMatrixWorld(force bool) {
	od.Call("updateMatrixWorld", force)
}
func (od *Object3D) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	od.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (od *Object3D) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("worldToLocal", vector)}
}

// Code generated from "helpers/PointLightHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PointLightHelper extend: [Object3D]
type PointLightHelper struct {
	js.Value
}

func NewPointLightHelper(light *PointLight, sphereSize float64, color *Color) *PointLightHelper {
	return &PointLightHelper{Value: get("PointLightHelper").New(light.JSValue(), sphereSize, color)}
}
func (plh *PointLightHelper) JSValue() js.Value {
	return plh.Value
}
func (plh *PointLightHelper) CastShadow() bool {
	return plh.Get("castShadow").Bool()
}
func (plh *PointLightHelper) SetCastShadow(v bool) {
	plh.Set("castShadow", v)
}
func (plh *PointLightHelper) Children() js.Value {
	return plh.Get("children")
}
func (plh *PointLightHelper) SetChildren(v js.Value) {
	plh.Set("children", v)
}
func (plh *PointLightHelper) Color() *Color {
	return &Color{Value: plh.Get("color")}
}
func (plh *PointLightHelper) SetColor(v *Color) {
	plh.Set("color", v.JSValue())
}
func (plh *PointLightHelper) FrustumCulled() bool {
	return plh.Get("frustumCulled").Bool()
}
func (plh *PointLightHelper) SetFrustumCulled(v bool) {
	plh.Set("frustumCulled", v)
}
func (plh *PointLightHelper) Id() int {
	return plh.Get("id").Int()
}
func (plh *PointLightHelper) SetId(v int) {
	plh.Set("id", v)
}
func (plh *PointLightHelper) IsObject3D() bool {
	return plh.Get("isObject3D").Bool()
}
func (plh *PointLightHelper) SetIsObject3D(v bool) {
	plh.Set("isObject3D", v)
}
func (plh *PointLightHelper) Layers() *Layers {
	return &Layers{Value: plh.Get("layers")}
}
func (plh *PointLightHelper) SetLayers(v *Layers) {
	plh.Set("layers", v.JSValue())
}
func (plh *PointLightHelper) Light() *PointLight {
	return &PointLight{Value: plh.Get("light")}
}
func (plh *PointLightHelper) SetLight(v *PointLight) {
	plh.Set("light", v.JSValue())
}
func (plh *PointLightHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: plh.Get("matrix")}
}
func (plh *PointLightHelper) SetMatrix(v *Matrix4) {
	plh.Set("matrix", v.JSValue())
}
func (plh *PointLightHelper) MatrixAutoUpdate() bool {
	return plh.Get("matrixAutoUpdate").Bool()
}
func (plh *PointLightHelper) SetMatrixAutoUpdate(v bool) {
	plh.Set("matrixAutoUpdate", v)
}
func (plh *PointLightHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: plh.Get("matrixWorld")}
}
func (plh *PointLightHelper) SetMatrixWorld(v *Matrix4) {
	plh.Set("matrixWorld", v.JSValue())
}
func (plh *PointLightHelper) MatrixWorldNeedsUpdate() bool {
	return plh.Get("matrixWorldNeedsUpdate").Bool()
}
func (plh *PointLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	plh.Set("matrixWorldNeedsUpdate", v)
}
func (plh *PointLightHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: plh.Get("modelViewMatrix")}
}
func (plh *PointLightHelper) SetModelViewMatrix(v *Matrix4) {
	plh.Set("modelViewMatrix", v.JSValue())
}
func (plh *PointLightHelper) Name() string {
	return plh.Get("name").String()
}
func (plh *PointLightHelper) SetName(v string) {
	plh.Set("name", v)
}
func (plh *PointLightHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: plh.Get("normalMatrix")}
}
func (plh *PointLightHelper) SetNormalMatrix(v *Matrix3) {
	plh.Set("normalMatrix", v.JSValue())
}
func (plh *PointLightHelper) OnAfterRender() js.Value {
	return plh.Get("onAfterRender")
}
func (plh *PointLightHelper) SetOnAfterRender(v js.Value) {
	plh.Set("onAfterRender", v)
}
func (plh *PointLightHelper) OnBeforeRender() js.Value {
	return plh.Get("onBeforeRender")
}
func (plh *PointLightHelper) SetOnBeforeRender(v js.Value) {
	plh.Set("onBeforeRender", v)
}
func (plh *PointLightHelper) Parent() *Object3D {
	return &Object3D{Value: plh.Get("parent")}
}
func (plh *PointLightHelper) SetParent(v *Object3D) {
	plh.Set("parent", v.JSValue())
}
func (plh *PointLightHelper) Position() *Vector3 {
	return &Vector3{Value: plh.Get("position")}
}
func (plh *PointLightHelper) SetPosition(v *Vector3) {
	plh.Set("position", v.JSValue())
}
func (plh *PointLightHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: plh.Get("quaternion")}
}
func (plh *PointLightHelper) SetQuaternion(v *Quaternion) {
	plh.Set("quaternion", v.JSValue())
}
func (plh *PointLightHelper) ReceiveShadow() bool {
	return plh.Get("receiveShadow").Bool()
}
func (plh *PointLightHelper) SetReceiveShadow(v bool) {
	plh.Set("receiveShadow", v)
}
func (plh *PointLightHelper) RenderOrder() float64 {
	return plh.Get("renderOrder").Float()
}
func (plh *PointLightHelper) SetRenderOrder(v float64) {
	plh.Set("renderOrder", v)
}
func (plh *PointLightHelper) Rotation() *Euler {
	return &Euler{Value: plh.Get("rotation")}
}
func (plh *PointLightHelper) SetRotation(v *Euler) {
	plh.Set("rotation", v.JSValue())
}
func (plh *PointLightHelper) Scale() *Vector3 {
	return &Vector3{Value: plh.Get("scale")}
}
func (plh *PointLightHelper) SetScale(v *Vector3) {
	plh.Set("scale", v.JSValue())
}
func (plh *PointLightHelper) Type() string {
	return plh.Get("type").String()
}
func (plh *PointLightHelper) SetType(v string) {
	plh.Set("type", v)
}
func (plh *PointLightHelper) Up() *Vector3 {
	return &Vector3{Value: plh.Get("up")}
}
func (plh *PointLightHelper) SetUp(v *Vector3) {
	plh.Set("up", v.JSValue())
}
func (plh *PointLightHelper) UserData() js.Value {
	return plh.Get("userData")
}
func (plh *PointLightHelper) SetUserData(v js.Value) {
	plh.Set("userData", v)
}
func (plh *PointLightHelper) Uuid() string {
	return plh.Get("uuid").String()
}
func (plh *PointLightHelper) SetUuid(v string) {
	plh.Set("uuid", v)
}
func (plh *PointLightHelper) Visible() bool {
	return plh.Get("visible").Bool()
}
func (plh *PointLightHelper) SetVisible(v bool) {
	plh.Set("visible", v)
}
func (plh *PointLightHelper) DefaultMatrixAutoUpdate() bool {
	return plh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (plh *PointLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	plh.Set("DefaultMatrixAutoUpdate", v)
}
func (plh *PointLightHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: plh.Get("DefaultUp")}
}
func (plh *PointLightHelper) SetDefaultUp(v *Vector3) {
	plh.Set("DefaultUp", v.JSValue())
}
func (plh *PointLightHelper) Add(object js.Value) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("add", object)}
}
func (plh *PointLightHelper) AddEventListener(typ string, listener js.Value) {
	plh.Call("addEventListener", typ, listener)
}
func (plh *PointLightHelper) ApplyMatrix(matrix *Matrix4) {
	plh.Call("applyMatrix", matrix.JSValue())
}
func (plh *PointLightHelper) ApplyQuaternion(quaternion *Quaternion) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("applyQuaternion", quaternion)}
}
func (plh *PointLightHelper) Clone(recursive bool) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("clone", recursive)}
}
func (plh *PointLightHelper) Copy(source *Object3D, recursive bool) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("copy", source, recursive)}
}
func (plh *PointLightHelper) DispatchEvent(event js.Value) {
	plh.Call("dispatchEvent", event)
}
func (plh *PointLightHelper) Dispose() {
	plh.Call("dispose")
}
func (plh *PointLightHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: plh.Call("getObjectById", id)}
}
func (plh *PointLightHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: plh.Call("getObjectByName", name)}
}
func (plh *PointLightHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: plh.Call("getObjectByProperty", name, value)}
}
func (plh *PointLightHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: plh.Call("getWorldDirection", target)}
}
func (plh *PointLightHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: plh.Call("getWorldPosition", target)}
}
func (plh *PointLightHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: plh.Call("getWorldQuaternion", target)}
}
func (plh *PointLightHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: plh.Call("getWorldScale", target)}
}
func (plh *PointLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return plh.Call("hasEventListener", typ, listener).Bool()
}
func (plh *PointLightHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: plh.Call("localToWorld", vector)}
}
func (plh *PointLightHelper) LookAt(vector *Vector3, y float64, z float64) {
	plh.Call("lookAt", vector, y, z)
}
func (plh *PointLightHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	plh.Call("raycast", raycaster.JSValue(), intersects)
}
func (plh *PointLightHelper) Remove(object js.Value) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("remove", object)}
}
func (plh *PointLightHelper) RemoveEventListener(typ string, listener js.Value) {
	plh.Call("removeEventListener", typ, listener)
}
func (plh *PointLightHelper) RotateOnAxis(axis *Vector3, angle float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("rotateOnAxis", axis, angle)}
}
func (plh *PointLightHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("rotateOnWorldAxis", axis, angle)}
}
func (plh *PointLightHelper) RotateX(angle float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("rotateX", angle)}
}
func (plh *PointLightHelper) RotateY(angle float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("rotateY", angle)}
}
func (plh *PointLightHelper) RotateZ(angle float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("rotateZ", angle)}
}
func (plh *PointLightHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	plh.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (plh *PointLightHelper) SetRotationFromEuler(euler *Euler) {
	plh.Call("setRotationFromEuler", euler.JSValue())
}
func (plh *PointLightHelper) SetRotationFromMatrix(m *Matrix4) {
	plh.Call("setRotationFromMatrix", m.JSValue())
}
func (plh *PointLightHelper) SetRotationFromQuaternion(q *Quaternion) {
	plh.Call("setRotationFromQuaternion", q.JSValue())
}
func (plh *PointLightHelper) ToJSON(meta js.Value) js.Value {
	return plh.Call("toJSON", meta)
}
func (plh *PointLightHelper) TranslateOnAxis(axis *Vector3, distance float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("translateOnAxis", axis, distance)}
}
func (plh *PointLightHelper) TranslateX(distance float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("translateX", distance)}
}
func (plh *PointLightHelper) TranslateY(distance float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("translateY", distance)}
}
func (plh *PointLightHelper) TranslateZ(distance float64) *PointLightHelper {
	return &PointLightHelper{Value: plh.Call("translateZ", distance)}
}
func (plh *PointLightHelper) Traverse(callback js.Value) {
	plh.Call("traverse", callback)
}
func (plh *PointLightHelper) TraverseAncestors(callback js.Value) {
	plh.Call("traverseAncestors", callback)
}
func (plh *PointLightHelper) TraverseVisible(callback js.Value) {
	plh.Call("traverseVisible", callback)
}
func (plh *PointLightHelper) Update() {
	plh.Call("update")
}
func (plh *PointLightHelper) UpdateMatrix() {
	plh.Call("updateMatrix")
}
func (plh *PointLightHelper) UpdateMatrixWorld(force bool) {
	plh.Call("updateMatrixWorld", force)
}
func (plh *PointLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	plh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (plh *PointLightHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: plh.Call("worldToLocal", vector)}
}

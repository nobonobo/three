// Code generated from "helpers/DirectionalLightHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type DirectionalLightHelper struct {
	js.Value
}

func NewDirectionalLightHelper(light *DirectionalLight, size float64, color *Color) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: get("DirectionalLightHelper").New(light, size, color)}
}
func (dlh *DirectionalLightHelper) CastShadow() bool {
	return dlh.Get("castShadow").Bool()
}
func (dlh *DirectionalLightHelper) SetCastShadow(v bool) {
	dlh.Set("castShadow", v)
}
func (dlh *DirectionalLightHelper) Children() js.Value {
	return dlh.Get("children")
}
func (dlh *DirectionalLightHelper) SetChildren(v js.Value) {
	dlh.Set("children", v)
}
func (dlh *DirectionalLightHelper) Color() *Color {
	return &Color{Value: dlh.Get("color")}
}
func (dlh *DirectionalLightHelper) SetColor(v *Color) {
	dlh.Set("color", v)
}
func (dlh *DirectionalLightHelper) FrustumCulled() bool {
	return dlh.Get("frustumCulled").Bool()
}
func (dlh *DirectionalLightHelper) SetFrustumCulled(v bool) {
	dlh.Set("frustumCulled", v)
}
func (dlh *DirectionalLightHelper) Id() int {
	return dlh.Get("id").Int()
}
func (dlh *DirectionalLightHelper) SetId(v int) {
	dlh.Set("id", v)
}
func (dlh *DirectionalLightHelper) IsObject3D() bool {
	return dlh.Get("isObject3D").Bool()
}
func (dlh *DirectionalLightHelper) SetIsObject3D(v bool) {
	dlh.Set("isObject3D", v)
}
func (dlh *DirectionalLightHelper) Layers() *Layers {
	return &Layers{Value: dlh.Get("layers")}
}
func (dlh *DirectionalLightHelper) SetLayers(v *Layers) {
	dlh.Set("layers", v)
}
func (dlh *DirectionalLightHelper) Light() *DirectionalLight {
	return &DirectionalLight{Value: dlh.Get("light")}
}
func (dlh *DirectionalLightHelper) SetLight(v *DirectionalLight) {
	dlh.Set("light", v)
}
func (dlh *DirectionalLightHelper) LightPlane() *Line {
	return &Line{Value: dlh.Get("lightPlane")}
}
func (dlh *DirectionalLightHelper) SetLightPlane(v *Line) {
	dlh.Set("lightPlane", v)
}
func (dlh *DirectionalLightHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: dlh.Get("matrix")}
}
func (dlh *DirectionalLightHelper) SetMatrix(v *Matrix4) {
	dlh.Set("matrix", v)
}
func (dlh *DirectionalLightHelper) MatrixAutoUpdate() bool {
	return dlh.Get("matrixAutoUpdate").Bool()
}
func (dlh *DirectionalLightHelper) SetMatrixAutoUpdate(v bool) {
	dlh.Set("matrixAutoUpdate", v)
}
func (dlh *DirectionalLightHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: dlh.Get("matrixWorld")}
}
func (dlh *DirectionalLightHelper) SetMatrixWorld(v *Matrix4) {
	dlh.Set("matrixWorld", v)
}
func (dlh *DirectionalLightHelper) MatrixWorldNeedsUpdate() bool {
	return dlh.Get("matrixWorldNeedsUpdate").Bool()
}
func (dlh *DirectionalLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	dlh.Set("matrixWorldNeedsUpdate", v)
}
func (dlh *DirectionalLightHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: dlh.Get("modelViewMatrix")}
}
func (dlh *DirectionalLightHelper) SetModelViewMatrix(v *Matrix4) {
	dlh.Set("modelViewMatrix", v)
}
func (dlh *DirectionalLightHelper) Name() string {
	return dlh.Get("name").String()
}
func (dlh *DirectionalLightHelper) SetName(v string) {
	dlh.Set("name", v)
}
func (dlh *DirectionalLightHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: dlh.Get("normalMatrix")}
}
func (dlh *DirectionalLightHelper) SetNormalMatrix(v *Matrix3) {
	dlh.Set("normalMatrix", v)
}
func (dlh *DirectionalLightHelper) OnAfterRender() js.Value {
	return dlh.Get("onAfterRender")
}
func (dlh *DirectionalLightHelper) SetOnAfterRender(v js.Value) {
	dlh.Set("onAfterRender", v)
}
func (dlh *DirectionalLightHelper) OnBeforeRender() js.Value {
	return dlh.Get("onBeforeRender")
}
func (dlh *DirectionalLightHelper) SetOnBeforeRender(v js.Value) {
	dlh.Set("onBeforeRender", v)
}
func (dlh *DirectionalLightHelper) Parent() *Object3D {
	return &Object3D{Value: dlh.Get("parent")}
}
func (dlh *DirectionalLightHelper) SetParent(v *Object3D) {
	dlh.Set("parent", v)
}
func (dlh *DirectionalLightHelper) Position() *Vector3 {
	return &Vector3{Value: dlh.Get("position")}
}
func (dlh *DirectionalLightHelper) SetPosition(v *Vector3) {
	dlh.Set("position", v)
}
func (dlh *DirectionalLightHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: dlh.Get("quaternion")}
}
func (dlh *DirectionalLightHelper) SetQuaternion(v *Quaternion) {
	dlh.Set("quaternion", v)
}
func (dlh *DirectionalLightHelper) ReceiveShadow() bool {
	return dlh.Get("receiveShadow").Bool()
}
func (dlh *DirectionalLightHelper) SetReceiveShadow(v bool) {
	dlh.Set("receiveShadow", v)
}
func (dlh *DirectionalLightHelper) RenderOrder() float64 {
	return dlh.Get("renderOrder").Float()
}
func (dlh *DirectionalLightHelper) SetRenderOrder(v float64) {
	dlh.Set("renderOrder", v)
}
func (dlh *DirectionalLightHelper) Rotation() *Euler {
	return &Euler{Value: dlh.Get("rotation")}
}
func (dlh *DirectionalLightHelper) SetRotation(v *Euler) {
	dlh.Set("rotation", v)
}
func (dlh *DirectionalLightHelper) Scale() *Vector3 {
	return &Vector3{Value: dlh.Get("scale")}
}
func (dlh *DirectionalLightHelper) SetScale(v *Vector3) {
	dlh.Set("scale", v)
}
func (dlh *DirectionalLightHelper) TargetPlane() *Line {
	return &Line{Value: dlh.Get("targetPlane")}
}
func (dlh *DirectionalLightHelper) SetTargetPlane(v *Line) {
	dlh.Set("targetPlane", v)
}
func (dlh *DirectionalLightHelper) Type() string {
	return dlh.Get("type").String()
}
func (dlh *DirectionalLightHelper) SetType(v string) {
	dlh.Set("type", v)
}
func (dlh *DirectionalLightHelper) Up() *Vector3 {
	return &Vector3{Value: dlh.Get("up")}
}
func (dlh *DirectionalLightHelper) SetUp(v *Vector3) {
	dlh.Set("up", v)
}
func (dlh *DirectionalLightHelper) UserData() js.Value {
	return dlh.Get("userData")
}
func (dlh *DirectionalLightHelper) SetUserData(v js.Value) {
	dlh.Set("userData", v)
}
func (dlh *DirectionalLightHelper) Uuid() string {
	return dlh.Get("uuid").String()
}
func (dlh *DirectionalLightHelper) SetUuid(v string) {
	dlh.Set("uuid", v)
}
func (dlh *DirectionalLightHelper) Visible() bool {
	return dlh.Get("visible").Bool()
}
func (dlh *DirectionalLightHelper) SetVisible(v bool) {
	dlh.Set("visible", v)
}
func (dlh *DirectionalLightHelper) DefaultMatrixAutoUpdate() bool {
	return dlh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (dlh *DirectionalLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	dlh.Set("DefaultMatrixAutoUpdate", v)
}
func (dlh *DirectionalLightHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: dlh.Get("DefaultUp")}
}
func (dlh *DirectionalLightHelper) SetDefaultUp(v *Vector3) {
	dlh.Set("DefaultUp", v)
}
func (dlh *DirectionalLightHelper) Add(object js.Value) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("add", object)}
}
func (dlh *DirectionalLightHelper) AddEventListener(typ string, listener js.Value) {
	dlh.Call("addEventListener", typ, listener)
}
func (dlh *DirectionalLightHelper) ApplyMatrix(matrix *Matrix4) {
	dlh.Call("applyMatrix", matrix)
}
func (dlh *DirectionalLightHelper) ApplyQuaternion(quaternion *Quaternion) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("applyQuaternion", quaternion)}
}
func (dlh *DirectionalLightHelper) Clone(recursive bool) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("clone", recursive)}
}
func (dlh *DirectionalLightHelper) Copy(source *Object3D, recursive bool) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("copy", source, recursive)}
}
func (dlh *DirectionalLightHelper) DispatchEvent(event js.Value) {
	dlh.Call("dispatchEvent", event)
}
func (dlh *DirectionalLightHelper) Dispose() {
	dlh.Call("dispose")
}
func (dlh *DirectionalLightHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: dlh.Call("getObjectById", id)}
}
func (dlh *DirectionalLightHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: dlh.Call("getObjectByName", name)}
}
func (dlh *DirectionalLightHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: dlh.Call("getObjectByProperty", name, value)}
}
func (dlh *DirectionalLightHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: dlh.Call("getWorldDirection", target)}
}
func (dlh *DirectionalLightHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: dlh.Call("getWorldPosition", target)}
}
func (dlh *DirectionalLightHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: dlh.Call("getWorldQuaternion", target)}
}
func (dlh *DirectionalLightHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: dlh.Call("getWorldScale", target)}
}
func (dlh *DirectionalLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return dlh.Call("hasEventListener", typ, listener).Bool()
}
func (dlh *DirectionalLightHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: dlh.Call("localToWorld", vector)}
}
func (dlh *DirectionalLightHelper) LookAt(vector *Vector3, y float64, z float64) {
	dlh.Call("lookAt", vector, y, z)
}
func (dlh *DirectionalLightHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	dlh.Call("raycast", raycaster, intersects)
}
func (dlh *DirectionalLightHelper) Remove(object js.Value) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("remove", object)}
}
func (dlh *DirectionalLightHelper) RemoveEventListener(typ string, listener js.Value) {
	dlh.Call("removeEventListener", typ, listener)
}
func (dlh *DirectionalLightHelper) RotateOnAxis(axis *Vector3, angle float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("rotateOnAxis", axis, angle)}
}
func (dlh *DirectionalLightHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("rotateOnWorldAxis", axis, angle)}
}
func (dlh *DirectionalLightHelper) RotateX(angle float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("rotateX", angle)}
}
func (dlh *DirectionalLightHelper) RotateY(angle float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("rotateY", angle)}
}
func (dlh *DirectionalLightHelper) RotateZ(angle float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("rotateZ", angle)}
}
func (dlh *DirectionalLightHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	dlh.Call("setRotationFromAxisAngle", axis, angle)
}
func (dlh *DirectionalLightHelper) SetRotationFromEuler(euler *Euler) {
	dlh.Call("setRotationFromEuler", euler)
}
func (dlh *DirectionalLightHelper) SetRotationFromMatrix(m *Matrix4) {
	dlh.Call("setRotationFromMatrix", m)
}
func (dlh *DirectionalLightHelper) SetRotationFromQuaternion(q *Quaternion) {
	dlh.Call("setRotationFromQuaternion", q)
}
func (dlh *DirectionalLightHelper) ToJSON(meta js.Value) js.Value {
	return dlh.Call("toJSON", meta)
}
func (dlh *DirectionalLightHelper) TranslateOnAxis(axis *Vector3, distance float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("translateOnAxis", axis, distance)}
}
func (dlh *DirectionalLightHelper) TranslateX(distance float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("translateX", distance)}
}
func (dlh *DirectionalLightHelper) TranslateY(distance float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("translateY", distance)}
}
func (dlh *DirectionalLightHelper) TranslateZ(distance float64) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: dlh.Call("translateZ", distance)}
}
func (dlh *DirectionalLightHelper) Traverse(callback js.Value) {
	dlh.Call("traverse", callback)
}
func (dlh *DirectionalLightHelper) TraverseAncestors(callback js.Value) {
	dlh.Call("traverseAncestors", callback)
}
func (dlh *DirectionalLightHelper) TraverseVisible(callback js.Value) {
	dlh.Call("traverseVisible", callback)
}
func (dlh *DirectionalLightHelper) Update() {
	dlh.Call("update")
}
func (dlh *DirectionalLightHelper) UpdateMatrix() {
	dlh.Call("updateMatrix")
}
func (dlh *DirectionalLightHelper) UpdateMatrixWorld(force bool) {
	dlh.Call("updateMatrixWorld", force)
}
func (dlh *DirectionalLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	dlh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (dlh *DirectionalLightHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: dlh.Call("worldToLocal", vector)}
}

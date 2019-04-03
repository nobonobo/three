// Code generated from "helpers/SpotLightHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type SpotLightHelper struct {
	js.Value
}

func NewSpotLightHelper(light *Light, color *Color) *SpotLightHelper {
	return &SpotLightHelper{Value: get("SpotLightHelper").New(light, color)}
}
func (slh *SpotLightHelper) CastShadow() bool {
	return slh.Get("castShadow").Bool()
}
func (slh *SpotLightHelper) SetCastShadow(v bool) {
	slh.Set("castShadow", v)
}
func (slh *SpotLightHelper) Children() js.Value {
	return slh.Get("children")
}
func (slh *SpotLightHelper) SetChildren(v js.Value) {
	slh.Set("children", v)
}
func (slh *SpotLightHelper) Color() *Color {
	return &Color{Value: slh.Get("color")}
}
func (slh *SpotLightHelper) SetColor(v *Color) {
	slh.Set("color", v)
}
func (slh *SpotLightHelper) FrustumCulled() bool {
	return slh.Get("frustumCulled").Bool()
}
func (slh *SpotLightHelper) SetFrustumCulled(v bool) {
	slh.Set("frustumCulled", v)
}
func (slh *SpotLightHelper) Id() int {
	return slh.Get("id").Int()
}
func (slh *SpotLightHelper) SetId(v int) {
	slh.Set("id", v)
}
func (slh *SpotLightHelper) IsObject3D() bool {
	return slh.Get("isObject3D").Bool()
}
func (slh *SpotLightHelper) SetIsObject3D(v bool) {
	slh.Set("isObject3D", v)
}
func (slh *SpotLightHelper) Layers() *Layers {
	return &Layers{Value: slh.Get("layers")}
}
func (slh *SpotLightHelper) SetLayers(v *Layers) {
	slh.Set("layers", v)
}
func (slh *SpotLightHelper) Light() *Light {
	return &Light{Value: slh.Get("light")}
}
func (slh *SpotLightHelper) SetLight(v *Light) {
	slh.Set("light", v)
}
func (slh *SpotLightHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: slh.Get("matrix")}
}
func (slh *SpotLightHelper) SetMatrix(v *Matrix4) {
	slh.Set("matrix", v)
}
func (slh *SpotLightHelper) MatrixAutoUpdate() bool {
	return slh.Get("matrixAutoUpdate").Bool()
}
func (slh *SpotLightHelper) SetMatrixAutoUpdate(v bool) {
	slh.Set("matrixAutoUpdate", v)
}
func (slh *SpotLightHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: slh.Get("matrixWorld")}
}
func (slh *SpotLightHelper) SetMatrixWorld(v *Matrix4) {
	slh.Set("matrixWorld", v)
}
func (slh *SpotLightHelper) MatrixWorldNeedsUpdate() bool {
	return slh.Get("matrixWorldNeedsUpdate").Bool()
}
func (slh *SpotLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	slh.Set("matrixWorldNeedsUpdate", v)
}
func (slh *SpotLightHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: slh.Get("modelViewMatrix")}
}
func (slh *SpotLightHelper) SetModelViewMatrix(v *Matrix4) {
	slh.Set("modelViewMatrix", v)
}
func (slh *SpotLightHelper) Name() string {
	return slh.Get("name").String()
}
func (slh *SpotLightHelper) SetName(v string) {
	slh.Set("name", v)
}
func (slh *SpotLightHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: slh.Get("normalMatrix")}
}
func (slh *SpotLightHelper) SetNormalMatrix(v *Matrix3) {
	slh.Set("normalMatrix", v)
}
func (slh *SpotLightHelper) OnAfterRender() js.Value {
	return slh.Get("onAfterRender")
}
func (slh *SpotLightHelper) SetOnAfterRender(v js.Value) {
	slh.Set("onAfterRender", v)
}
func (slh *SpotLightHelper) OnBeforeRender() js.Value {
	return slh.Get("onBeforeRender")
}
func (slh *SpotLightHelper) SetOnBeforeRender(v js.Value) {
	slh.Set("onBeforeRender", v)
}
func (slh *SpotLightHelper) Parent() *Object3D {
	return &Object3D{Value: slh.Get("parent")}
}
func (slh *SpotLightHelper) SetParent(v *Object3D) {
	slh.Set("parent", v)
}
func (slh *SpotLightHelper) Position() *Vector3 {
	return &Vector3{Value: slh.Get("position")}
}
func (slh *SpotLightHelper) SetPosition(v *Vector3) {
	slh.Set("position", v)
}
func (slh *SpotLightHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: slh.Get("quaternion")}
}
func (slh *SpotLightHelper) SetQuaternion(v *Quaternion) {
	slh.Set("quaternion", v)
}
func (slh *SpotLightHelper) ReceiveShadow() bool {
	return slh.Get("receiveShadow").Bool()
}
func (slh *SpotLightHelper) SetReceiveShadow(v bool) {
	slh.Set("receiveShadow", v)
}
func (slh *SpotLightHelper) RenderOrder() float64 {
	return slh.Get("renderOrder").Float()
}
func (slh *SpotLightHelper) SetRenderOrder(v float64) {
	slh.Set("renderOrder", v)
}
func (slh *SpotLightHelper) Rotation() *Euler {
	return &Euler{Value: slh.Get("rotation")}
}
func (slh *SpotLightHelper) SetRotation(v *Euler) {
	slh.Set("rotation", v)
}
func (slh *SpotLightHelper) Scale() *Vector3 {
	return &Vector3{Value: slh.Get("scale")}
}
func (slh *SpotLightHelper) SetScale(v *Vector3) {
	slh.Set("scale", v)
}
func (slh *SpotLightHelper) Type() string {
	return slh.Get("type").String()
}
func (slh *SpotLightHelper) SetType(v string) {
	slh.Set("type", v)
}
func (slh *SpotLightHelper) Up() *Vector3 {
	return &Vector3{Value: slh.Get("up")}
}
func (slh *SpotLightHelper) SetUp(v *Vector3) {
	slh.Set("up", v)
}
func (slh *SpotLightHelper) UserData() js.Value {
	return slh.Get("userData")
}
func (slh *SpotLightHelper) SetUserData(v js.Value) {
	slh.Set("userData", v)
}
func (slh *SpotLightHelper) Uuid() string {
	return slh.Get("uuid").String()
}
func (slh *SpotLightHelper) SetUuid(v string) {
	slh.Set("uuid", v)
}
func (slh *SpotLightHelper) Visible() bool {
	return slh.Get("visible").Bool()
}
func (slh *SpotLightHelper) SetVisible(v bool) {
	slh.Set("visible", v)
}
func (slh *SpotLightHelper) DefaultMatrixAutoUpdate() bool {
	return slh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (slh *SpotLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	slh.Set("DefaultMatrixAutoUpdate", v)
}
func (slh *SpotLightHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: slh.Get("DefaultUp")}
}
func (slh *SpotLightHelper) SetDefaultUp(v *Vector3) {
	slh.Set("DefaultUp", v)
}
func (slh *SpotLightHelper) Add(object js.Value) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("add", object)}
}
func (slh *SpotLightHelper) AddEventListener(typ string, listener js.Value) {
	slh.Call("addEventListener", typ, listener)
}
func (slh *SpotLightHelper) ApplyMatrix(matrix *Matrix4) {
	slh.Call("applyMatrix", matrix)
}
func (slh *SpotLightHelper) ApplyQuaternion(quaternion *Quaternion) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("applyQuaternion", quaternion)}
}
func (slh *SpotLightHelper) Clone(recursive bool) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("clone", recursive)}
}
func (slh *SpotLightHelper) Copy(source *Object3D, recursive bool) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("copy", source, recursive)}
}
func (slh *SpotLightHelper) DispatchEvent(event js.Value) {
	slh.Call("dispatchEvent", event)
}
func (slh *SpotLightHelper) Dispose() {
	slh.Call("dispose")
}
func (slh *SpotLightHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: slh.Call("getObjectById", id)}
}
func (slh *SpotLightHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: slh.Call("getObjectByName", name)}
}
func (slh *SpotLightHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: slh.Call("getObjectByProperty", name, value)}
}
func (slh *SpotLightHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: slh.Call("getWorldDirection", target)}
}
func (slh *SpotLightHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: slh.Call("getWorldPosition", target)}
}
func (slh *SpotLightHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: slh.Call("getWorldQuaternion", target)}
}
func (slh *SpotLightHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: slh.Call("getWorldScale", target)}
}
func (slh *SpotLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return slh.Call("hasEventListener", typ, listener).Bool()
}
func (slh *SpotLightHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: slh.Call("localToWorld", vector)}
}
func (slh *SpotLightHelper) LookAt(vector *Vector3, y float64, z float64) {
	slh.Call("lookAt", vector, y, z)
}
func (slh *SpotLightHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	slh.Call("raycast", raycaster, intersects)
}
func (slh *SpotLightHelper) Remove(object js.Value) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("remove", object)}
}
func (slh *SpotLightHelper) RemoveEventListener(typ string, listener js.Value) {
	slh.Call("removeEventListener", typ, listener)
}
func (slh *SpotLightHelper) RotateOnAxis(axis *Vector3, angle float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("rotateOnAxis", axis, angle)}
}
func (slh *SpotLightHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("rotateOnWorldAxis", axis, angle)}
}
func (slh *SpotLightHelper) RotateX(angle float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("rotateX", angle)}
}
func (slh *SpotLightHelper) RotateY(angle float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("rotateY", angle)}
}
func (slh *SpotLightHelper) RotateZ(angle float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("rotateZ", angle)}
}
func (slh *SpotLightHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	slh.Call("setRotationFromAxisAngle", axis, angle)
}
func (slh *SpotLightHelper) SetRotationFromEuler(euler *Euler) {
	slh.Call("setRotationFromEuler", euler)
}
func (slh *SpotLightHelper) SetRotationFromMatrix(m *Matrix4) {
	slh.Call("setRotationFromMatrix", m)
}
func (slh *SpotLightHelper) SetRotationFromQuaternion(q *Quaternion) {
	slh.Call("setRotationFromQuaternion", q)
}
func (slh *SpotLightHelper) ToJSON(meta js.Value) js.Value {
	return slh.Call("toJSON", meta)
}
func (slh *SpotLightHelper) TranslateOnAxis(axis *Vector3, distance float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("translateOnAxis", axis, distance)}
}
func (slh *SpotLightHelper) TranslateX(distance float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("translateX", distance)}
}
func (slh *SpotLightHelper) TranslateY(distance float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("translateY", distance)}
}
func (slh *SpotLightHelper) TranslateZ(distance float64) *SpotLightHelper {
	return &SpotLightHelper{Value: slh.Call("translateZ", distance)}
}
func (slh *SpotLightHelper) Traverse(callback js.Value) {
	slh.Call("traverse", callback)
}
func (slh *SpotLightHelper) TraverseAncestors(callback js.Value) {
	slh.Call("traverseAncestors", callback)
}
func (slh *SpotLightHelper) TraverseVisible(callback js.Value) {
	slh.Call("traverseVisible", callback)
}
func (slh *SpotLightHelper) Update() {
	slh.Call("update")
}
func (slh *SpotLightHelper) UpdateMatrix() {
	slh.Call("updateMatrix")
}
func (slh *SpotLightHelper) UpdateMatrixWorld(force bool) {
	slh.Call("updateMatrixWorld", force)
}
func (slh *SpotLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	slh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (slh *SpotLightHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: slh.Call("worldToLocal", vector)}
}

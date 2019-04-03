// Code generated from "helpers/HemisphereLightHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type HemisphereLightHelper struct {
	js.Value
}

func NewHemisphereLightHelper(light *HemisphereLight, size float64, color *Color) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: get("HemisphereLightHelper").New(light, size, color)}
}
func (hlh *HemisphereLightHelper) CastShadow() bool {
	return hlh.Get("castShadow").Bool()
}
func (hlh *HemisphereLightHelper) SetCastShadow(v bool) {
	hlh.Set("castShadow", v)
}
func (hlh *HemisphereLightHelper) Children() js.Value {
	return hlh.Get("children")
}
func (hlh *HemisphereLightHelper) SetChildren(v js.Value) {
	hlh.Set("children", v)
}
func (hlh *HemisphereLightHelper) Color() *Color {
	return &Color{Value: hlh.Get("color")}
}
func (hlh *HemisphereLightHelper) SetColor(v *Color) {
	hlh.Set("color", v)
}
func (hlh *HemisphereLightHelper) FrustumCulled() bool {
	return hlh.Get("frustumCulled").Bool()
}
func (hlh *HemisphereLightHelper) SetFrustumCulled(v bool) {
	hlh.Set("frustumCulled", v)
}
func (hlh *HemisphereLightHelper) Id() int {
	return hlh.Get("id").Int()
}
func (hlh *HemisphereLightHelper) SetId(v int) {
	hlh.Set("id", v)
}
func (hlh *HemisphereLightHelper) IsObject3D() bool {
	return hlh.Get("isObject3D").Bool()
}
func (hlh *HemisphereLightHelper) SetIsObject3D(v bool) {
	hlh.Set("isObject3D", v)
}
func (hlh *HemisphereLightHelper) Layers() *Layers {
	return &Layers{Value: hlh.Get("layers")}
}
func (hlh *HemisphereLightHelper) SetLayers(v *Layers) {
	hlh.Set("layers", v)
}
func (hlh *HemisphereLightHelper) Light() *HemisphereLight {
	return &HemisphereLight{Value: hlh.Get("light")}
}
func (hlh *HemisphereLightHelper) SetLight(v *HemisphereLight) {
	hlh.Set("light", v)
}
func (hlh *HemisphereLightHelper) Material() *MeshBasicMaterial {
	return &MeshBasicMaterial{Value: hlh.Get("material")}
}
func (hlh *HemisphereLightHelper) SetMaterial(v *MeshBasicMaterial) {
	hlh.Set("material", v)
}
func (hlh *HemisphereLightHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: hlh.Get("matrix")}
}
func (hlh *HemisphereLightHelper) SetMatrix(v *Matrix4) {
	hlh.Set("matrix", v)
}
func (hlh *HemisphereLightHelper) MatrixAutoUpdate() bool {
	return hlh.Get("matrixAutoUpdate").Bool()
}
func (hlh *HemisphereLightHelper) SetMatrixAutoUpdate(v bool) {
	hlh.Set("matrixAutoUpdate", v)
}
func (hlh *HemisphereLightHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: hlh.Get("matrixWorld")}
}
func (hlh *HemisphereLightHelper) SetMatrixWorld(v *Matrix4) {
	hlh.Set("matrixWorld", v)
}
func (hlh *HemisphereLightHelper) MatrixWorldNeedsUpdate() bool {
	return hlh.Get("matrixWorldNeedsUpdate").Bool()
}
func (hlh *HemisphereLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	hlh.Set("matrixWorldNeedsUpdate", v)
}
func (hlh *HemisphereLightHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: hlh.Get("modelViewMatrix")}
}
func (hlh *HemisphereLightHelper) SetModelViewMatrix(v *Matrix4) {
	hlh.Set("modelViewMatrix", v)
}
func (hlh *HemisphereLightHelper) Name() string {
	return hlh.Get("name").String()
}
func (hlh *HemisphereLightHelper) SetName(v string) {
	hlh.Set("name", v)
}
func (hlh *HemisphereLightHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: hlh.Get("normalMatrix")}
}
func (hlh *HemisphereLightHelper) SetNormalMatrix(v *Matrix3) {
	hlh.Set("normalMatrix", v)
}
func (hlh *HemisphereLightHelper) OnAfterRender() js.Value {
	return hlh.Get("onAfterRender")
}
func (hlh *HemisphereLightHelper) SetOnAfterRender(v js.Value) {
	hlh.Set("onAfterRender", v)
}
func (hlh *HemisphereLightHelper) OnBeforeRender() js.Value {
	return hlh.Get("onBeforeRender")
}
func (hlh *HemisphereLightHelper) SetOnBeforeRender(v js.Value) {
	hlh.Set("onBeforeRender", v)
}
func (hlh *HemisphereLightHelper) Parent() *Object3D {
	return &Object3D{Value: hlh.Get("parent")}
}
func (hlh *HemisphereLightHelper) SetParent(v *Object3D) {
	hlh.Set("parent", v)
}
func (hlh *HemisphereLightHelper) Position() *Vector3 {
	return &Vector3{Value: hlh.Get("position")}
}
func (hlh *HemisphereLightHelper) SetPosition(v *Vector3) {
	hlh.Set("position", v)
}
func (hlh *HemisphereLightHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: hlh.Get("quaternion")}
}
func (hlh *HemisphereLightHelper) SetQuaternion(v *Quaternion) {
	hlh.Set("quaternion", v)
}
func (hlh *HemisphereLightHelper) ReceiveShadow() bool {
	return hlh.Get("receiveShadow").Bool()
}
func (hlh *HemisphereLightHelper) SetReceiveShadow(v bool) {
	hlh.Set("receiveShadow", v)
}
func (hlh *HemisphereLightHelper) RenderOrder() float64 {
	return hlh.Get("renderOrder").Float()
}
func (hlh *HemisphereLightHelper) SetRenderOrder(v float64) {
	hlh.Set("renderOrder", v)
}
func (hlh *HemisphereLightHelper) Rotation() *Euler {
	return &Euler{Value: hlh.Get("rotation")}
}
func (hlh *HemisphereLightHelper) SetRotation(v *Euler) {
	hlh.Set("rotation", v)
}
func (hlh *HemisphereLightHelper) Scale() *Vector3 {
	return &Vector3{Value: hlh.Get("scale")}
}
func (hlh *HemisphereLightHelper) SetScale(v *Vector3) {
	hlh.Set("scale", v)
}
func (hlh *HemisphereLightHelper) Type() string {
	return hlh.Get("type").String()
}
func (hlh *HemisphereLightHelper) SetType(v string) {
	hlh.Set("type", v)
}
func (hlh *HemisphereLightHelper) Up() *Vector3 {
	return &Vector3{Value: hlh.Get("up")}
}
func (hlh *HemisphereLightHelper) SetUp(v *Vector3) {
	hlh.Set("up", v)
}
func (hlh *HemisphereLightHelper) UserData() js.Value {
	return hlh.Get("userData")
}
func (hlh *HemisphereLightHelper) SetUserData(v js.Value) {
	hlh.Set("userData", v)
}
func (hlh *HemisphereLightHelper) Uuid() string {
	return hlh.Get("uuid").String()
}
func (hlh *HemisphereLightHelper) SetUuid(v string) {
	hlh.Set("uuid", v)
}
func (hlh *HemisphereLightHelper) Visible() bool {
	return hlh.Get("visible").Bool()
}
func (hlh *HemisphereLightHelper) SetVisible(v bool) {
	hlh.Set("visible", v)
}
func (hlh *HemisphereLightHelper) DefaultMatrixAutoUpdate() bool {
	return hlh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (hlh *HemisphereLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	hlh.Set("DefaultMatrixAutoUpdate", v)
}
func (hlh *HemisphereLightHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: hlh.Get("DefaultUp")}
}
func (hlh *HemisphereLightHelper) SetDefaultUp(v *Vector3) {
	hlh.Set("DefaultUp", v)
}
func (hlh *HemisphereLightHelper) Add(object js.Value) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("add", object)}
}
func (hlh *HemisphereLightHelper) AddEventListener(typ string, listener js.Value) {
	hlh.Call("addEventListener", typ, listener)
}
func (hlh *HemisphereLightHelper) ApplyMatrix(matrix *Matrix4) {
	hlh.Call("applyMatrix", matrix)
}
func (hlh *HemisphereLightHelper) ApplyQuaternion(quaternion *Quaternion) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("applyQuaternion", quaternion)}
}
func (hlh *HemisphereLightHelper) Clone(recursive bool) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("clone", recursive)}
}
func (hlh *HemisphereLightHelper) Copy(source *Object3D, recursive bool) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("copy", source, recursive)}
}
func (hlh *HemisphereLightHelper) DispatchEvent(event js.Value) {
	hlh.Call("dispatchEvent", event)
}
func (hlh *HemisphereLightHelper) Dispose() {
	hlh.Call("dispose")
}
func (hlh *HemisphereLightHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: hlh.Call("getObjectById", id)}
}
func (hlh *HemisphereLightHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: hlh.Call("getObjectByName", name)}
}
func (hlh *HemisphereLightHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: hlh.Call("getObjectByProperty", name, value)}
}
func (hlh *HemisphereLightHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: hlh.Call("getWorldDirection", target)}
}
func (hlh *HemisphereLightHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: hlh.Call("getWorldPosition", target)}
}
func (hlh *HemisphereLightHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: hlh.Call("getWorldQuaternion", target)}
}
func (hlh *HemisphereLightHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: hlh.Call("getWorldScale", target)}
}
func (hlh *HemisphereLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return hlh.Call("hasEventListener", typ, listener).Bool()
}
func (hlh *HemisphereLightHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: hlh.Call("localToWorld", vector)}
}
func (hlh *HemisphereLightHelper) LookAt(vector *Vector3, y float64, z float64) {
	hlh.Call("lookAt", vector, y, z)
}
func (hlh *HemisphereLightHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	hlh.Call("raycast", raycaster, intersects)
}
func (hlh *HemisphereLightHelper) Remove(object js.Value) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("remove", object)}
}
func (hlh *HemisphereLightHelper) RemoveEventListener(typ string, listener js.Value) {
	hlh.Call("removeEventListener", typ, listener)
}
func (hlh *HemisphereLightHelper) RotateOnAxis(axis *Vector3, angle float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("rotateOnAxis", axis, angle)}
}
func (hlh *HemisphereLightHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("rotateOnWorldAxis", axis, angle)}
}
func (hlh *HemisphereLightHelper) RotateX(angle float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("rotateX", angle)}
}
func (hlh *HemisphereLightHelper) RotateY(angle float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("rotateY", angle)}
}
func (hlh *HemisphereLightHelper) RotateZ(angle float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("rotateZ", angle)}
}
func (hlh *HemisphereLightHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	hlh.Call("setRotationFromAxisAngle", axis, angle)
}
func (hlh *HemisphereLightHelper) SetRotationFromEuler(euler *Euler) {
	hlh.Call("setRotationFromEuler", euler)
}
func (hlh *HemisphereLightHelper) SetRotationFromMatrix(m *Matrix4) {
	hlh.Call("setRotationFromMatrix", m)
}
func (hlh *HemisphereLightHelper) SetRotationFromQuaternion(q *Quaternion) {
	hlh.Call("setRotationFromQuaternion", q)
}
func (hlh *HemisphereLightHelper) ToJSON(meta js.Value) js.Value {
	return hlh.Call("toJSON", meta)
}
func (hlh *HemisphereLightHelper) TranslateOnAxis(axis *Vector3, distance float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("translateOnAxis", axis, distance)}
}
func (hlh *HemisphereLightHelper) TranslateX(distance float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("translateX", distance)}
}
func (hlh *HemisphereLightHelper) TranslateY(distance float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("translateY", distance)}
}
func (hlh *HemisphereLightHelper) TranslateZ(distance float64) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: hlh.Call("translateZ", distance)}
}
func (hlh *HemisphereLightHelper) Traverse(callback js.Value) {
	hlh.Call("traverse", callback)
}
func (hlh *HemisphereLightHelper) TraverseAncestors(callback js.Value) {
	hlh.Call("traverseAncestors", callback)
}
func (hlh *HemisphereLightHelper) TraverseVisible(callback js.Value) {
	hlh.Call("traverseVisible", callback)
}
func (hlh *HemisphereLightHelper) Update() {
	hlh.Call("update")
}
func (hlh *HemisphereLightHelper) UpdateMatrix() {
	hlh.Call("updateMatrix")
}
func (hlh *HemisphereLightHelper) UpdateMatrixWorld(force bool) {
	hlh.Call("updateMatrixWorld", force)
}
func (hlh *HemisphereLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	hlh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (hlh *HemisphereLightHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: hlh.Call("worldToLocal", vector)}
}

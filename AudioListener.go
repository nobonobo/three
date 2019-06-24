// Code generated from "audio/AudioListener.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type AudioListener struct {
	js.Value
}

func NewAudioListener() *AudioListener {
	return &AudioListener{Value: get("AudioListener").New()}
}
func (al *AudioListener) CastShadow() bool {
	return al.Get("castShadow").Bool()
}
func (al *AudioListener) SetCastShadow(v bool) {
	al.Set("castShadow", v)
}
func (al *AudioListener) Children() js.Value {
	return al.Get("children")
}
func (al *AudioListener) SetChildren(v js.Value) {
	al.Set("children", v)
}
func (al *AudioListener) Context() js.Value {
	return al.Get("context")
}
func (al *AudioListener) SetContext(v js.Value) {
	al.Set("context", v)
}
func (al *AudioListener) Filter() js.Value {
	return al.Get("filter")
}
func (al *AudioListener) SetFilter(v js.Value) {
	al.Set("filter", v)
}
func (al *AudioListener) FrustumCulled() bool {
	return al.Get("frustumCulled").Bool()
}
func (al *AudioListener) SetFrustumCulled(v bool) {
	al.Set("frustumCulled", v)
}
func (al *AudioListener) Gain() js.Value {
	return al.Get("gain")
}
func (al *AudioListener) SetGain(v js.Value) {
	al.Set("gain", v)
}
func (al *AudioListener) Id() int {
	return al.Get("id").Int()
}
func (al *AudioListener) SetId(v int) {
	al.Set("id", v)
}
func (al *AudioListener) IsObject3D() bool {
	return al.Get("isObject3D").Bool()
}
func (al *AudioListener) SetIsObject3D(v bool) {
	al.Set("isObject3D", v)
}
func (al *AudioListener) Layers() *Layers {
	return &Layers{Value: al.Get("layers")}
}
func (al *AudioListener) SetLayers(v *Layers) {
	al.Set("layers", v)
}
func (al *AudioListener) Matrix() *Matrix4 {
	return &Matrix4{Value: al.Get("matrix")}
}
func (al *AudioListener) SetMatrix(v *Matrix4) {
	al.Set("matrix", v)
}
func (al *AudioListener) MatrixAutoUpdate() bool {
	return al.Get("matrixAutoUpdate").Bool()
}
func (al *AudioListener) SetMatrixAutoUpdate(v bool) {
	al.Set("matrixAutoUpdate", v)
}
func (al *AudioListener) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: al.Get("matrixWorld")}
}
func (al *AudioListener) SetMatrixWorld(v *Matrix4) {
	al.Set("matrixWorld", v)
}
func (al *AudioListener) MatrixWorldNeedsUpdate() bool {
	return al.Get("matrixWorldNeedsUpdate").Bool()
}
func (al *AudioListener) SetMatrixWorldNeedsUpdate(v bool) {
	al.Set("matrixWorldNeedsUpdate", v)
}
func (al *AudioListener) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: al.Get("modelViewMatrix")}
}
func (al *AudioListener) SetModelViewMatrix(v *Matrix4) {
	al.Set("modelViewMatrix", v)
}
func (al *AudioListener) Name() string {
	return al.Get("name").String()
}
func (al *AudioListener) SetName(v string) {
	al.Set("name", v)
}
func (al *AudioListener) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: al.Get("normalMatrix")}
}
func (al *AudioListener) SetNormalMatrix(v *Matrix3) {
	al.Set("normalMatrix", v)
}
func (al *AudioListener) OnAfterRender() js.Value {
	return al.Get("onAfterRender")
}
func (al *AudioListener) SetOnAfterRender(v js.Value) {
	al.Set("onAfterRender", v)
}
func (al *AudioListener) OnBeforeRender() js.Value {
	return al.Get("onBeforeRender")
}
func (al *AudioListener) SetOnBeforeRender(v js.Value) {
	al.Set("onBeforeRender", v)
}
func (al *AudioListener) Parent() *Object3D {
	return &Object3D{Value: al.Get("parent")}
}
func (al *AudioListener) SetParent(v *Object3D) {
	al.Set("parent", v)
}
func (al *AudioListener) Position() *Vector3 {
	return &Vector3{Value: al.Get("position")}
}
func (al *AudioListener) SetPosition(v *Vector3) {
	al.Set("position", v)
}
func (al *AudioListener) Quaternion() *Quaternion {
	return &Quaternion{Value: al.Get("quaternion")}
}
func (al *AudioListener) SetQuaternion(v *Quaternion) {
	al.Set("quaternion", v)
}
func (al *AudioListener) ReceiveShadow() bool {
	return al.Get("receiveShadow").Bool()
}
func (al *AudioListener) SetReceiveShadow(v bool) {
	al.Set("receiveShadow", v)
}
func (al *AudioListener) RenderOrder() float64 {
	return al.Get("renderOrder").Float()
}
func (al *AudioListener) SetRenderOrder(v float64) {
	al.Set("renderOrder", v)
}
func (al *AudioListener) Rotation() *Euler {
	return &Euler{Value: al.Get("rotation")}
}
func (al *AudioListener) SetRotation(v *Euler) {
	al.Set("rotation", v)
}
func (al *AudioListener) Scale() *Vector3 {
	return &Vector3{Value: al.Get("scale")}
}
func (al *AudioListener) SetScale(v *Vector3) {
	al.Set("scale", v)
}
func (al *AudioListener) Type() string {
	return al.Get("type").String()
}
func (al *AudioListener) SetType(v string) {
	al.Set("type", v)
}
func (al *AudioListener) Up() *Vector3 {
	return &Vector3{Value: al.Get("up")}
}
func (al *AudioListener) SetUp(v *Vector3) {
	al.Set("up", v)
}
func (al *AudioListener) UserData() js.Value {
	return al.Get("userData")
}
func (al *AudioListener) SetUserData(v js.Value) {
	al.Set("userData", v)
}
func (al *AudioListener) Uuid() string {
	return al.Get("uuid").String()
}
func (al *AudioListener) SetUuid(v string) {
	al.Set("uuid", v)
}
func (al *AudioListener) Visible() bool {
	return al.Get("visible").Bool()
}
func (al *AudioListener) SetVisible(v bool) {
	al.Set("visible", v)
}
func (al *AudioListener) DefaultMatrixAutoUpdate() bool {
	return al.Get("DefaultMatrixAutoUpdate").Bool()
}
func (al *AudioListener) SetDefaultMatrixAutoUpdate(v bool) {
	al.Set("DefaultMatrixAutoUpdate", v)
}
func (al *AudioListener) DefaultUp() *Vector3 {
	return &Vector3{Value: al.Get("DefaultUp")}
}
func (al *AudioListener) SetDefaultUp(v *Vector3) {
	al.Set("DefaultUp", v)
}
func (al *AudioListener) Add(object js.Value) *AudioListener {
	return &AudioListener{Value: al.Call("add", object)}
}
func (al *AudioListener) AddEventListener(typ string, listener js.Value) {
	al.Call("addEventListener", typ, listener)
}
func (al *AudioListener) ApplyMatrix(matrix *Matrix4) {
	al.Call("applyMatrix", matrix)
}
func (al *AudioListener) ApplyQuaternion(quaternion *Quaternion) *AudioListener {
	return &AudioListener{Value: al.Call("applyQuaternion", quaternion)}
}
func (al *AudioListener) Clone(recursive bool) *AudioListener {
	return &AudioListener{Value: al.Call("clone", recursive)}
}
func (al *AudioListener) Copy(source *Object3D, recursive bool) *AudioListener {
	return &AudioListener{Value: al.Call("copy", source, recursive)}
}
func (al *AudioListener) DispatchEvent(event js.Value) {
	al.Call("dispatchEvent", event)
}
func (al *AudioListener) GetFilter() js.Value {
	return al.Call("getFilter")
}
func (al *AudioListener) GetInput() js.Value {
	return al.Call("getInput")
}
func (al *AudioListener) GetMasterVolume() float64 {
	return al.Call("getMasterVolume").Float()
}
func (al *AudioListener) GetObjectById(id int) *Object3D {
	return &Object3D{Value: al.Call("getObjectById", id)}
}
func (al *AudioListener) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: al.Call("getObjectByName", name)}
}
func (al *AudioListener) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: al.Call("getObjectByProperty", name, value)}
}
func (al *AudioListener) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("getWorldDirection", target)}
}
func (al *AudioListener) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("getWorldPosition", target)}
}
func (al *AudioListener) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: al.Call("getWorldQuaternion", target)}
}
func (al *AudioListener) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("getWorldScale", target)}
}
func (al *AudioListener) HasEventListener(typ string, listener js.Value) bool {
	return al.Call("hasEventListener", typ, listener).Bool()
}
func (al *AudioListener) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("localToWorld", vector)}
}
func (al *AudioListener) LookAt(vector *Vector3, y float64, z float64) {
	al.Call("lookAt", vector, y, z)
}
func (al *AudioListener) Raycast(raycaster *Raycaster, intersects js.Value) {
	al.Call("raycast", raycaster, intersects)
}
func (al *AudioListener) Remove(object js.Value) *AudioListener {
	return &AudioListener{Value: al.Call("remove", object)}
}
func (al *AudioListener) RemoveEventListener(typ string, listener js.Value) {
	al.Call("removeEventListener", typ, listener)
}
func (al *AudioListener) RemoveFilter() {
	al.Call("removeFilter")
}
func (al *AudioListener) RotateOnAxis(axis *Vector3, angle float64) *AudioListener {
	return &AudioListener{Value: al.Call("rotateOnAxis", axis, angle)}
}
func (al *AudioListener) RotateOnWorldAxis(axis *Vector3, angle float64) *AudioListener {
	return &AudioListener{Value: al.Call("rotateOnWorldAxis", axis, angle)}
}
func (al *AudioListener) RotateX(angle float64) *AudioListener {
	return &AudioListener{Value: al.Call("rotateX", angle)}
}
func (al *AudioListener) RotateY(angle float64) *AudioListener {
	return &AudioListener{Value: al.Call("rotateY", angle)}
}
func (al *AudioListener) RotateZ(angle float64) *AudioListener {
	return &AudioListener{Value: al.Call("rotateZ", angle)}
}
func (al *AudioListener) SetFilter2(value js.Value) {
	al.Call("setFilter", value)
}
func (al *AudioListener) SetMasterVolume(value float64) {
	al.Call("setMasterVolume", value)
}
func (al *AudioListener) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	al.Call("setRotationFromAxisAngle", axis, angle)
}
func (al *AudioListener) SetRotationFromEuler(euler *Euler) {
	al.Call("setRotationFromEuler", euler)
}
func (al *AudioListener) SetRotationFromMatrix(m *Matrix4) {
	al.Call("setRotationFromMatrix", m)
}
func (al *AudioListener) SetRotationFromQuaternion(q *Quaternion) {
	al.Call("setRotationFromQuaternion", q)
}
func (al *AudioListener) ToJSON(meta js.Value) js.Value {
	return al.Call("toJSON", meta)
}
func (al *AudioListener) TranslateOnAxis(axis *Vector3, distance float64) *AudioListener {
	return &AudioListener{Value: al.Call("translateOnAxis", axis, distance)}
}
func (al *AudioListener) TranslateX(distance float64) *AudioListener {
	return &AudioListener{Value: al.Call("translateX", distance)}
}
func (al *AudioListener) TranslateY(distance float64) *AudioListener {
	return &AudioListener{Value: al.Call("translateY", distance)}
}
func (al *AudioListener) TranslateZ(distance float64) *AudioListener {
	return &AudioListener{Value: al.Call("translateZ", distance)}
}
func (al *AudioListener) Traverse(callback js.Value) {
	al.Call("traverse", callback)
}
func (al *AudioListener) TraverseAncestors(callback js.Value) {
	al.Call("traverseAncestors", callback)
}
func (al *AudioListener) TraverseVisible(callback js.Value) {
	al.Call("traverseVisible", callback)
}
func (al *AudioListener) UpdateMatrix() {
	al.Call("updateMatrix")
}
func (al *AudioListener) UpdateMatrixWorld(force bool) {
	al.Call("updateMatrixWorld", force)
}
func (al *AudioListener) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	al.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (al *AudioListener) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("worldToLocal", vector)}
}

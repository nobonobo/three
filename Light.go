// Code generated from "lights/Light.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Light extend: [Object3D]
type Light struct {
	js.Value
}

func NewLight(hex int, intensity float64) *Light {
	return &Light{Value: get("Light").New(hex, intensity)}
}
func (ll *Light) JSValue() js.Value {
	return ll.Value
}
func (ll *Light) CastShadow() bool {
	return ll.Get("castShadow").Bool()
}
func (ll *Light) SetCastShadow(v bool) {
	ll.Set("castShadow", v)
}
func (ll *Light) Children() js.Value {
	return ll.Get("children")
}
func (ll *Light) SetChildren(v js.Value) {
	ll.Set("children", v)
}
func (ll *Light) Color() *Color {
	return &Color{Value: ll.Get("color")}
}
func (ll *Light) SetColor(v *Color) {
	ll.Set("color", v.Value)
}
func (ll *Light) FrustumCulled() bool {
	return ll.Get("frustumCulled").Bool()
}
func (ll *Light) SetFrustumCulled(v bool) {
	ll.Set("frustumCulled", v)
}
func (ll *Light) Id() int {
	return ll.Get("id").Int()
}
func (ll *Light) SetId(v int) {
	ll.Set("id", v)
}
func (ll *Light) Intensity() float64 {
	return ll.Get("intensity").Float()
}
func (ll *Light) SetIntensity(v float64) {
	ll.Set("intensity", v)
}
func (ll *Light) IsLight() bool {
	return ll.Get("isLight").Bool()
}
func (ll *Light) SetIsLight(v bool) {
	ll.Set("isLight", v)
}
func (ll *Light) IsObject3D() bool {
	return ll.Get("isObject3D").Bool()
}
func (ll *Light) SetIsObject3D(v bool) {
	ll.Set("isObject3D", v)
}
func (ll *Light) Layers() *Layers {
	return &Layers{Value: ll.Get("layers")}
}
func (ll *Light) SetLayers(v *Layers) {
	ll.Set("layers", v.Value)
}
func (ll *Light) Matrix() *Matrix4 {
	return &Matrix4{Value: ll.Get("matrix")}
}
func (ll *Light) SetMatrix(v *Matrix4) {
	ll.Set("matrix", v.Value)
}
func (ll *Light) MatrixAutoUpdate() bool {
	return ll.Get("matrixAutoUpdate").Bool()
}
func (ll *Light) SetMatrixAutoUpdate(v bool) {
	ll.Set("matrixAutoUpdate", v)
}
func (ll *Light) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ll.Get("matrixWorld")}
}
func (ll *Light) SetMatrixWorld(v *Matrix4) {
	ll.Set("matrixWorld", v.Value)
}
func (ll *Light) MatrixWorldNeedsUpdate() bool {
	return ll.Get("matrixWorldNeedsUpdate").Bool()
}
func (ll *Light) SetMatrixWorldNeedsUpdate(v bool) {
	ll.Set("matrixWorldNeedsUpdate", v)
}
func (ll *Light) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ll.Get("modelViewMatrix")}
}
func (ll *Light) SetModelViewMatrix(v *Matrix4) {
	ll.Set("modelViewMatrix", v.Value)
}
func (ll *Light) Name() string {
	return ll.Get("name").String()
}
func (ll *Light) SetName(v string) {
	ll.Set("name", v)
}
func (ll *Light) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ll.Get("normalMatrix")}
}
func (ll *Light) SetNormalMatrix(v *Matrix3) {
	ll.Set("normalMatrix", v.Value)
}
func (ll *Light) OnAfterRender() js.Value {
	return ll.Get("onAfterRender")
}
func (ll *Light) SetOnAfterRender(v js.Value) {
	ll.Set("onAfterRender", v)
}
func (ll *Light) OnBeforeRender() js.Value {
	return ll.Get("onBeforeRender")
}
func (ll *Light) SetOnBeforeRender(v js.Value) {
	ll.Set("onBeforeRender", v)
}
func (ll *Light) Parent() *Object3D {
	return &Object3D{Value: ll.Get("parent")}
}
func (ll *Light) SetParent(v *Object3D) {
	ll.Set("parent", v.Value)
}
func (ll *Light) Position() *Vector3 {
	return &Vector3{Value: ll.Get("position")}
}
func (ll *Light) SetPosition(v *Vector3) {
	ll.Set("position", v.Value)
}
func (ll *Light) Quaternion() *Quaternion {
	return &Quaternion{Value: ll.Get("quaternion")}
}
func (ll *Light) SetQuaternion(v *Quaternion) {
	ll.Set("quaternion", v.Value)
}
func (ll *Light) ReceiveShadow() bool {
	return ll.Get("receiveShadow").Bool()
}
func (ll *Light) SetReceiveShadow(v bool) {
	ll.Set("receiveShadow", v)
}
func (ll *Light) RenderOrder() float64 {
	return ll.Get("renderOrder").Float()
}
func (ll *Light) SetRenderOrder(v float64) {
	ll.Set("renderOrder", v)
}
func (ll *Light) Rotation() *Euler {
	return &Euler{Value: ll.Get("rotation")}
}
func (ll *Light) SetRotation(v *Euler) {
	ll.Set("rotation", v.Value)
}
func (ll *Light) Scale() *Vector3 {
	return &Vector3{Value: ll.Get("scale")}
}
func (ll *Light) SetScale(v *Vector3) {
	ll.Set("scale", v.Value)
}
func (ll *Light) Shadow() *LightShadow {
	return &LightShadow{Value: ll.Get("shadow")}
}
func (ll *Light) SetShadow(v *LightShadow) {
	ll.Set("shadow", v.Value)
}
func (ll *Light) ShadowBias() js.Value {
	return ll.Get("shadowBias")
}
func (ll *Light) SetShadowBias(v js.Value) {
	ll.Set("shadowBias", v)
}
func (ll *Light) ShadowCameraBottom() js.Value {
	return ll.Get("shadowCameraBottom")
}
func (ll *Light) SetShadowCameraBottom(v js.Value) {
	ll.Set("shadowCameraBottom", v)
}
func (ll *Light) ShadowCameraFar() js.Value {
	return ll.Get("shadowCameraFar")
}
func (ll *Light) SetShadowCameraFar(v js.Value) {
	ll.Set("shadowCameraFar", v)
}
func (ll *Light) ShadowCameraFov() js.Value {
	return ll.Get("shadowCameraFov")
}
func (ll *Light) SetShadowCameraFov(v js.Value) {
	ll.Set("shadowCameraFov", v)
}
func (ll *Light) ShadowCameraLeft() js.Value {
	return ll.Get("shadowCameraLeft")
}
func (ll *Light) SetShadowCameraLeft(v js.Value) {
	ll.Set("shadowCameraLeft", v)
}
func (ll *Light) ShadowCameraNear() js.Value {
	return ll.Get("shadowCameraNear")
}
func (ll *Light) SetShadowCameraNear(v js.Value) {
	ll.Set("shadowCameraNear", v)
}
func (ll *Light) ShadowCameraRight() js.Value {
	return ll.Get("shadowCameraRight")
}
func (ll *Light) SetShadowCameraRight(v js.Value) {
	ll.Set("shadowCameraRight", v)
}
func (ll *Light) ShadowCameraTop() js.Value {
	return ll.Get("shadowCameraTop")
}
func (ll *Light) SetShadowCameraTop(v js.Value) {
	ll.Set("shadowCameraTop", v)
}
func (ll *Light) ShadowMapHeight() js.Value {
	return ll.Get("shadowMapHeight")
}
func (ll *Light) SetShadowMapHeight(v js.Value) {
	ll.Set("shadowMapHeight", v)
}
func (ll *Light) ShadowMapWidth() js.Value {
	return ll.Get("shadowMapWidth")
}
func (ll *Light) SetShadowMapWidth(v js.Value) {
	ll.Set("shadowMapWidth", v)
}
func (ll *Light) Type() string {
	return ll.Get("type").String()
}
func (ll *Light) SetType(v string) {
	ll.Set("type", v)
}
func (ll *Light) Up() *Vector3 {
	return &Vector3{Value: ll.Get("up")}
}
func (ll *Light) SetUp(v *Vector3) {
	ll.Set("up", v.Value)
}
func (ll *Light) UserData() js.Value {
	return ll.Get("userData")
}
func (ll *Light) SetUserData(v js.Value) {
	ll.Set("userData", v)
}
func (ll *Light) Uuid() string {
	return ll.Get("uuid").String()
}
func (ll *Light) SetUuid(v string) {
	ll.Set("uuid", v)
}
func (ll *Light) Visible() bool {
	return ll.Get("visible").Bool()
}
func (ll *Light) SetVisible(v bool) {
	ll.Set("visible", v)
}
func (ll *Light) DefaultMatrixAutoUpdate() bool {
	return ll.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ll *Light) SetDefaultMatrixAutoUpdate(v bool) {
	ll.Set("DefaultMatrixAutoUpdate", v)
}
func (ll *Light) DefaultUp() *Vector3 {
	return &Vector3{Value: ll.Get("DefaultUp")}
}
func (ll *Light) SetDefaultUp(v *Vector3) {
	ll.Set("DefaultUp", v.Value)
}
func (ll *Light) Add(object js.Value) *Light {
	return &Light{Value: ll.Call("add", object)}
}
func (ll *Light) AddEventListener(typ string, listener js.Value) {
	ll.Call("addEventListener", typ, listener)
}
func (ll *Light) ApplyMatrix(matrix *Matrix4) {
	ll.Call("applyMatrix", matrix)
}
func (ll *Light) ApplyQuaternion(quaternion *Quaternion) *Light {
	return &Light{Value: ll.Call("applyQuaternion", quaternion)}
}
func (ll *Light) Clone(recursive bool) *Light {
	return &Light{Value: ll.Call("clone", recursive)}
}
func (ll *Light) Copy(source *Object3D, recursive bool) *Light {
	return &Light{Value: ll.Call("copy", source, recursive)}
}
func (ll *Light) DispatchEvent(event js.Value) {
	ll.Call("dispatchEvent", event)
}
func (ll *Light) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ll.Call("getObjectById", id)}
}
func (ll *Light) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ll.Call("getObjectByName", name)}
}
func (ll *Light) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ll.Call("getObjectByProperty", name, value)}
}
func (ll *Light) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldDirection", target)}
}
func (ll *Light) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldPosition", target)}
}
func (ll *Light) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ll.Call("getWorldQuaternion", target)}
}
func (ll *Light) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldScale", target)}
}
func (ll *Light) HasEventListener(typ string, listener js.Value) bool {
	return ll.Call("hasEventListener", typ, listener).Bool()
}
func (ll *Light) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("localToWorld", vector)}
}
func (ll *Light) LookAt(vector *Vector3, y float64, z float64) {
	ll.Call("lookAt", vector, y, z)
}
func (ll *Light) Raycast(raycaster *Raycaster, intersects js.Value) {
	ll.Call("raycast", raycaster, intersects)
}
func (ll *Light) Remove(object js.Value) *Light {
	return &Light{Value: ll.Call("remove", object)}
}
func (ll *Light) RemoveEventListener(typ string, listener js.Value) {
	ll.Call("removeEventListener", typ, listener)
}
func (ll *Light) RotateOnAxis(axis *Vector3, angle float64) *Light {
	return &Light{Value: ll.Call("rotateOnAxis", axis, angle)}
}
func (ll *Light) RotateOnWorldAxis(axis *Vector3, angle float64) *Light {
	return &Light{Value: ll.Call("rotateOnWorldAxis", axis, angle)}
}
func (ll *Light) RotateX(angle float64) *Light {
	return &Light{Value: ll.Call("rotateX", angle)}
}
func (ll *Light) RotateY(angle float64) *Light {
	return &Light{Value: ll.Call("rotateY", angle)}
}
func (ll *Light) RotateZ(angle float64) *Light {
	return &Light{Value: ll.Call("rotateZ", angle)}
}
func (ll *Light) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ll.Call("setRotationFromAxisAngle", axis, angle)
}
func (ll *Light) SetRotationFromEuler(euler *Euler) {
	ll.Call("setRotationFromEuler", euler)
}
func (ll *Light) SetRotationFromMatrix(m *Matrix4) {
	ll.Call("setRotationFromMatrix", m)
}
func (ll *Light) SetRotationFromQuaternion(q *Quaternion) {
	ll.Call("setRotationFromQuaternion", q)
}
func (ll *Light) ToJSON(meta js.Value) js.Value {
	return ll.Call("toJSON", meta)
}
func (ll *Light) TranslateOnAxis(axis *Vector3, distance float64) *Light {
	return &Light{Value: ll.Call("translateOnAxis", axis, distance)}
}
func (ll *Light) TranslateX(distance float64) *Light {
	return &Light{Value: ll.Call("translateX", distance)}
}
func (ll *Light) TranslateY(distance float64) *Light {
	return &Light{Value: ll.Call("translateY", distance)}
}
func (ll *Light) TranslateZ(distance float64) *Light {
	return &Light{Value: ll.Call("translateZ", distance)}
}
func (ll *Light) Traverse(callback js.Value) {
	ll.Call("traverse", callback)
}
func (ll *Light) TraverseAncestors(callback js.Value) {
	ll.Call("traverseAncestors", callback)
}
func (ll *Light) TraverseVisible(callback js.Value) {
	ll.Call("traverseVisible", callback)
}
func (ll *Light) UpdateMatrix() {
	ll.Call("updateMatrix")
}
func (ll *Light) UpdateMatrixWorld(force bool) {
	ll.Call("updateMatrixWorld", force)
}
func (ll *Light) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ll.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ll *Light) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("worldToLocal", vector)}
}

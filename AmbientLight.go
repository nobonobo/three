// Code generated from "lights/AmbientLight.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type AmbientLight struct {
	js.Value
}

func NewAmbientLight(color *Color, intensity float64) *AmbientLight {
	return &AmbientLight{Value: get("AmbientLight").New(color, intensity)}
}
func (al *AmbientLight) CastShadow() bool {
	return al.Get("castShadow").Bool()
}
func (al *AmbientLight) SetCastShadow(v bool) {
	al.Set("castShadow", v)
}
func (al *AmbientLight) Children() js.Value {
	return al.Get("children")
}
func (al *AmbientLight) SetChildren(v js.Value) {
	al.Set("children", v)
}
func (al *AmbientLight) Color() *Color {
	return &Color{Value: al.Get("color")}
}
func (al *AmbientLight) SetColor(v *Color) {
	al.Set("color", v)
}
func (al *AmbientLight) FrustumCulled() bool {
	return al.Get("frustumCulled").Bool()
}
func (al *AmbientLight) SetFrustumCulled(v bool) {
	al.Set("frustumCulled", v)
}
func (al *AmbientLight) Id() int {
	return al.Get("id").Int()
}
func (al *AmbientLight) SetId(v int) {
	al.Set("id", v)
}
func (al *AmbientLight) Intensity() float64 {
	return al.Get("intensity").Float()
}
func (al *AmbientLight) SetIntensity(v float64) {
	al.Set("intensity", v)
}
func (al *AmbientLight) IsLight() bool {
	return al.Get("isLight").Bool()
}
func (al *AmbientLight) SetIsLight(v bool) {
	al.Set("isLight", v)
}
func (al *AmbientLight) IsObject3D() bool {
	return al.Get("isObject3D").Bool()
}
func (al *AmbientLight) SetIsObject3D(v bool) {
	al.Set("isObject3D", v)
}
func (al *AmbientLight) Layers() *Layers {
	return &Layers{Value: al.Get("layers")}
}
func (al *AmbientLight) SetLayers(v *Layers) {
	al.Set("layers", v)
}
func (al *AmbientLight) Matrix() *Matrix4 {
	return &Matrix4{Value: al.Get("matrix")}
}
func (al *AmbientLight) SetMatrix(v *Matrix4) {
	al.Set("matrix", v)
}
func (al *AmbientLight) MatrixAutoUpdate() bool {
	return al.Get("matrixAutoUpdate").Bool()
}
func (al *AmbientLight) SetMatrixAutoUpdate(v bool) {
	al.Set("matrixAutoUpdate", v)
}
func (al *AmbientLight) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: al.Get("matrixWorld")}
}
func (al *AmbientLight) SetMatrixWorld(v *Matrix4) {
	al.Set("matrixWorld", v)
}
func (al *AmbientLight) MatrixWorldNeedsUpdate() bool {
	return al.Get("matrixWorldNeedsUpdate").Bool()
}
func (al *AmbientLight) SetMatrixWorldNeedsUpdate(v bool) {
	al.Set("matrixWorldNeedsUpdate", v)
}
func (al *AmbientLight) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: al.Get("modelViewMatrix")}
}
func (al *AmbientLight) SetModelViewMatrix(v *Matrix4) {
	al.Set("modelViewMatrix", v)
}
func (al *AmbientLight) Name() string {
	return al.Get("name").String()
}
func (al *AmbientLight) SetName(v string) {
	al.Set("name", v)
}
func (al *AmbientLight) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: al.Get("normalMatrix")}
}
func (al *AmbientLight) SetNormalMatrix(v *Matrix3) {
	al.Set("normalMatrix", v)
}
func (al *AmbientLight) OnAfterRender() js.Value {
	return al.Get("onAfterRender")
}
func (al *AmbientLight) SetOnAfterRender(v js.Value) {
	al.Set("onAfterRender", v)
}
func (al *AmbientLight) OnBeforeRender() js.Value {
	return al.Get("onBeforeRender")
}
func (al *AmbientLight) SetOnBeforeRender(v js.Value) {
	al.Set("onBeforeRender", v)
}
func (al *AmbientLight) Parent() *Object3D {
	return &Object3D{Value: al.Get("parent")}
}
func (al *AmbientLight) SetParent(v *Object3D) {
	al.Set("parent", v)
}
func (al *AmbientLight) Position() *Vector3 {
	return &Vector3{Value: al.Get("position")}
}
func (al *AmbientLight) SetPosition(v *Vector3) {
	al.Set("position", v)
}
func (al *AmbientLight) Quaternion() *Quaternion {
	return &Quaternion{Value: al.Get("quaternion")}
}
func (al *AmbientLight) SetQuaternion(v *Quaternion) {
	al.Set("quaternion", v)
}
func (al *AmbientLight) ReceiveShadow() bool {
	return al.Get("receiveShadow").Bool()
}
func (al *AmbientLight) SetReceiveShadow(v bool) {
	al.Set("receiveShadow", v)
}
func (al *AmbientLight) RenderOrder() float64 {
	return al.Get("renderOrder").Float()
}
func (al *AmbientLight) SetRenderOrder(v float64) {
	al.Set("renderOrder", v)
}
func (al *AmbientLight) Rotation() *Euler {
	return &Euler{Value: al.Get("rotation")}
}
func (al *AmbientLight) SetRotation(v *Euler) {
	al.Set("rotation", v)
}
func (al *AmbientLight) Scale() *Vector3 {
	return &Vector3{Value: al.Get("scale")}
}
func (al *AmbientLight) SetScale(v *Vector3) {
	al.Set("scale", v)
}
func (al *AmbientLight) Shadow() *LightShadow {
	return &LightShadow{Value: al.Get("shadow")}
}
func (al *AmbientLight) SetShadow(v *LightShadow) {
	al.Set("shadow", v)
}
func (al *AmbientLight) ShadowBias() js.Value {
	return al.Get("shadowBias")
}
func (al *AmbientLight) SetShadowBias(v js.Value) {
	al.Set("shadowBias", v)
}
func (al *AmbientLight) ShadowCameraBottom() js.Value {
	return al.Get("shadowCameraBottom")
}
func (al *AmbientLight) SetShadowCameraBottom(v js.Value) {
	al.Set("shadowCameraBottom", v)
}
func (al *AmbientLight) ShadowCameraFar() js.Value {
	return al.Get("shadowCameraFar")
}
func (al *AmbientLight) SetShadowCameraFar(v js.Value) {
	al.Set("shadowCameraFar", v)
}
func (al *AmbientLight) ShadowCameraFov() js.Value {
	return al.Get("shadowCameraFov")
}
func (al *AmbientLight) SetShadowCameraFov(v js.Value) {
	al.Set("shadowCameraFov", v)
}
func (al *AmbientLight) ShadowCameraLeft() js.Value {
	return al.Get("shadowCameraLeft")
}
func (al *AmbientLight) SetShadowCameraLeft(v js.Value) {
	al.Set("shadowCameraLeft", v)
}
func (al *AmbientLight) ShadowCameraNear() js.Value {
	return al.Get("shadowCameraNear")
}
func (al *AmbientLight) SetShadowCameraNear(v js.Value) {
	al.Set("shadowCameraNear", v)
}
func (al *AmbientLight) ShadowCameraRight() js.Value {
	return al.Get("shadowCameraRight")
}
func (al *AmbientLight) SetShadowCameraRight(v js.Value) {
	al.Set("shadowCameraRight", v)
}
func (al *AmbientLight) ShadowCameraTop() js.Value {
	return al.Get("shadowCameraTop")
}
func (al *AmbientLight) SetShadowCameraTop(v js.Value) {
	al.Set("shadowCameraTop", v)
}
func (al *AmbientLight) ShadowMapHeight() js.Value {
	return al.Get("shadowMapHeight")
}
func (al *AmbientLight) SetShadowMapHeight(v js.Value) {
	al.Set("shadowMapHeight", v)
}
func (al *AmbientLight) ShadowMapWidth() js.Value {
	return al.Get("shadowMapWidth")
}
func (al *AmbientLight) SetShadowMapWidth(v js.Value) {
	al.Set("shadowMapWidth", v)
}
func (al *AmbientLight) Type() string {
	return al.Get("type").String()
}
func (al *AmbientLight) SetType(v string) {
	al.Set("type", v)
}
func (al *AmbientLight) Up() *Vector3 {
	return &Vector3{Value: al.Get("up")}
}
func (al *AmbientLight) SetUp(v *Vector3) {
	al.Set("up", v)
}
func (al *AmbientLight) UserData() js.Value {
	return al.Get("userData")
}
func (al *AmbientLight) SetUserData(v js.Value) {
	al.Set("userData", v)
}
func (al *AmbientLight) Uuid() string {
	return al.Get("uuid").String()
}
func (al *AmbientLight) SetUuid(v string) {
	al.Set("uuid", v)
}
func (al *AmbientLight) Visible() bool {
	return al.Get("visible").Bool()
}
func (al *AmbientLight) SetVisible(v bool) {
	al.Set("visible", v)
}
func (al *AmbientLight) DefaultMatrixAutoUpdate() bool {
	return al.Get("DefaultMatrixAutoUpdate").Bool()
}
func (al *AmbientLight) SetDefaultMatrixAutoUpdate(v bool) {
	al.Set("DefaultMatrixAutoUpdate", v)
}
func (al *AmbientLight) DefaultUp() *Vector3 {
	return &Vector3{Value: al.Get("DefaultUp")}
}
func (al *AmbientLight) SetDefaultUp(v *Vector3) {
	al.Set("DefaultUp", v)
}
func (al *AmbientLight) Add(object js.Value) *AmbientLight {
	return &AmbientLight{Value: al.Call("add", object)}
}
func (al *AmbientLight) AddEventListener(typ string, listener js.Value) {
	al.Call("addEventListener", typ, listener)
}
func (al *AmbientLight) ApplyMatrix(matrix *Matrix4) {
	al.Call("applyMatrix", matrix)
}
func (al *AmbientLight) ApplyQuaternion(quaternion *Quaternion) *AmbientLight {
	return &AmbientLight{Value: al.Call("applyQuaternion", quaternion)}
}
func (al *AmbientLight) Clone(recursive bool) *AmbientLight {
	return &AmbientLight{Value: al.Call("clone", recursive)}
}
func (al *AmbientLight) Copy(source *Object3D, recursive bool) *AmbientLight {
	return &AmbientLight{Value: al.Call("copy", source, recursive)}
}
func (al *AmbientLight) DispatchEvent(event js.Value) {
	al.Call("dispatchEvent", event)
}
func (al *AmbientLight) GetObjectById(id int) *Object3D {
	return &Object3D{Value: al.Call("getObjectById", id)}
}
func (al *AmbientLight) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: al.Call("getObjectByName", name)}
}
func (al *AmbientLight) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: al.Call("getObjectByProperty", name, value)}
}
func (al *AmbientLight) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("getWorldDirection", target)}
}
func (al *AmbientLight) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("getWorldPosition", target)}
}
func (al *AmbientLight) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: al.Call("getWorldQuaternion", target)}
}
func (al *AmbientLight) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("getWorldScale", target)}
}
func (al *AmbientLight) HasEventListener(typ string, listener js.Value) bool {
	return al.Call("hasEventListener", typ, listener).Bool()
}
func (al *AmbientLight) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("localToWorld", vector)}
}
func (al *AmbientLight) LookAt(vector *Vector3, y float64, z float64) {
	al.Call("lookAt", vector, y, z)
}
func (al *AmbientLight) Raycast(raycaster *Raycaster, intersects js.Value) {
	al.Call("raycast", raycaster, intersects)
}
func (al *AmbientLight) Remove(object js.Value) *AmbientLight {
	return &AmbientLight{Value: al.Call("remove", object)}
}
func (al *AmbientLight) RemoveEventListener(typ string, listener js.Value) {
	al.Call("removeEventListener", typ, listener)
}
func (al *AmbientLight) RotateOnAxis(axis *Vector3, angle float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("rotateOnAxis", axis, angle)}
}
func (al *AmbientLight) RotateOnWorldAxis(axis *Vector3, angle float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("rotateOnWorldAxis", axis, angle)}
}
func (al *AmbientLight) RotateX(angle float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("rotateX", angle)}
}
func (al *AmbientLight) RotateY(angle float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("rotateY", angle)}
}
func (al *AmbientLight) RotateZ(angle float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("rotateZ", angle)}
}
func (al *AmbientLight) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	al.Call("setRotationFromAxisAngle", axis, angle)
}
func (al *AmbientLight) SetRotationFromEuler(euler *Euler) {
	al.Call("setRotationFromEuler", euler)
}
func (al *AmbientLight) SetRotationFromMatrix(m *Matrix4) {
	al.Call("setRotationFromMatrix", m)
}
func (al *AmbientLight) SetRotationFromQuaternion(q *Quaternion) {
	al.Call("setRotationFromQuaternion", q)
}
func (al *AmbientLight) ToJSON(meta js.Value) js.Value {
	return al.Call("toJSON", meta)
}
func (al *AmbientLight) TranslateOnAxis(axis *Vector3, distance float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("translateOnAxis", axis, distance)}
}
func (al *AmbientLight) TranslateX(distance float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("translateX", distance)}
}
func (al *AmbientLight) TranslateY(distance float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("translateY", distance)}
}
func (al *AmbientLight) TranslateZ(distance float64) *AmbientLight {
	return &AmbientLight{Value: al.Call("translateZ", distance)}
}
func (al *AmbientLight) Traverse(callback js.Value) {
	al.Call("traverse", callback)
}
func (al *AmbientLight) TraverseAncestors(callback js.Value) {
	al.Call("traverseAncestors", callback)
}
func (al *AmbientLight) TraverseVisible(callback js.Value) {
	al.Call("traverseVisible", callback)
}
func (al *AmbientLight) UpdateMatrix() {
	al.Call("updateMatrix")
}
func (al *AmbientLight) UpdateMatrixWorld(force bool) {
	al.Call("updateMatrixWorld", force)
}
func (al *AmbientLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	al.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (al *AmbientLight) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: al.Call("worldToLocal", vector)}
}

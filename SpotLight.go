// Code generated from "lights/SpotLight.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type SpotLight struct {
	js.Value
}

func NewSpotLight(color *Color, intensity float64, distance float64, angle float64, exponent float64, decay float64) *SpotLight {
	return &SpotLight{Value: get("SpotLight").New(color, intensity, distance, angle, exponent, decay)}
}
func (sl *SpotLight) Angle() float64 {
	return sl.Get("angle").Float()
}
func (sl *SpotLight) SetAngle(v float64) {
	sl.Set("angle", v)
}
func (sl *SpotLight) CastShadow() bool {
	return sl.Get("castShadow").Bool()
}
func (sl *SpotLight) SetCastShadow(v bool) {
	sl.Set("castShadow", v)
}
func (sl *SpotLight) Children() js.Value {
	return sl.Get("children")
}
func (sl *SpotLight) SetChildren(v js.Value) {
	sl.Set("children", v)
}
func (sl *SpotLight) Color() *Color {
	return &Color{Value: sl.Get("color")}
}
func (sl *SpotLight) SetColor(v *Color) {
	sl.Set("color", v)
}
func (sl *SpotLight) Decay() float64 {
	return sl.Get("decay").Float()
}
func (sl *SpotLight) SetDecay(v float64) {
	sl.Set("decay", v)
}
func (sl *SpotLight) Distance() float64 {
	return sl.Get("distance").Float()
}
func (sl *SpotLight) SetDistance(v float64) {
	sl.Set("distance", v)
}
func (sl *SpotLight) Exponent() float64 {
	return sl.Get("exponent").Float()
}
func (sl *SpotLight) SetExponent(v float64) {
	sl.Set("exponent", v)
}
func (sl *SpotLight) FrustumCulled() bool {
	return sl.Get("frustumCulled").Bool()
}
func (sl *SpotLight) SetFrustumCulled(v bool) {
	sl.Set("frustumCulled", v)
}
func (sl *SpotLight) Id() int {
	return sl.Get("id").Int()
}
func (sl *SpotLight) SetId(v int) {
	sl.Set("id", v)
}
func (sl *SpotLight) Intensity() float64 {
	return sl.Get("intensity").Float()
}
func (sl *SpotLight) SetIntensity(v float64) {
	sl.Set("intensity", v)
}
func (sl *SpotLight) IsLight() bool {
	return sl.Get("isLight").Bool()
}
func (sl *SpotLight) SetIsLight(v bool) {
	sl.Set("isLight", v)
}
func (sl *SpotLight) IsObject3D() bool {
	return sl.Get("isObject3D").Bool()
}
func (sl *SpotLight) SetIsObject3D(v bool) {
	sl.Set("isObject3D", v)
}
func (sl *SpotLight) Layers() *Layers {
	return &Layers{Value: sl.Get("layers")}
}
func (sl *SpotLight) SetLayers(v *Layers) {
	sl.Set("layers", v)
}
func (sl *SpotLight) Matrix() *Matrix4 {
	return &Matrix4{Value: sl.Get("matrix")}
}
func (sl *SpotLight) SetMatrix(v *Matrix4) {
	sl.Set("matrix", v)
}
func (sl *SpotLight) MatrixAutoUpdate() bool {
	return sl.Get("matrixAutoUpdate").Bool()
}
func (sl *SpotLight) SetMatrixAutoUpdate(v bool) {
	sl.Set("matrixAutoUpdate", v)
}
func (sl *SpotLight) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: sl.Get("matrixWorld")}
}
func (sl *SpotLight) SetMatrixWorld(v *Matrix4) {
	sl.Set("matrixWorld", v)
}
func (sl *SpotLight) MatrixWorldNeedsUpdate() bool {
	return sl.Get("matrixWorldNeedsUpdate").Bool()
}
func (sl *SpotLight) SetMatrixWorldNeedsUpdate(v bool) {
	sl.Set("matrixWorldNeedsUpdate", v)
}
func (sl *SpotLight) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: sl.Get("modelViewMatrix")}
}
func (sl *SpotLight) SetModelViewMatrix(v *Matrix4) {
	sl.Set("modelViewMatrix", v)
}
func (sl *SpotLight) Name() string {
	return sl.Get("name").String()
}
func (sl *SpotLight) SetName(v string) {
	sl.Set("name", v)
}
func (sl *SpotLight) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: sl.Get("normalMatrix")}
}
func (sl *SpotLight) SetNormalMatrix(v *Matrix3) {
	sl.Set("normalMatrix", v)
}
func (sl *SpotLight) OnAfterRender() js.Value {
	return sl.Get("onAfterRender")
}
func (sl *SpotLight) SetOnAfterRender(v js.Value) {
	sl.Set("onAfterRender", v)
}
func (sl *SpotLight) OnBeforeRender() js.Value {
	return sl.Get("onBeforeRender")
}
func (sl *SpotLight) SetOnBeforeRender(v js.Value) {
	sl.Set("onBeforeRender", v)
}
func (sl *SpotLight) Parent() *Object3D {
	return &Object3D{Value: sl.Get("parent")}
}
func (sl *SpotLight) SetParent(v *Object3D) {
	sl.Set("parent", v)
}
func (sl *SpotLight) Penumbra() float64 {
	return sl.Get("penumbra").Float()
}
func (sl *SpotLight) SetPenumbra(v float64) {
	sl.Set("penumbra", v)
}
func (sl *SpotLight) Position() *Vector3 {
	return &Vector3{Value: sl.Get("position")}
}
func (sl *SpotLight) SetPosition(v *Vector3) {
	sl.Set("position", v)
}
func (sl *SpotLight) Power() float64 {
	return sl.Get("power").Float()
}
func (sl *SpotLight) SetPower(v float64) {
	sl.Set("power", v)
}
func (sl *SpotLight) Quaternion() *Quaternion {
	return &Quaternion{Value: sl.Get("quaternion")}
}
func (sl *SpotLight) SetQuaternion(v *Quaternion) {
	sl.Set("quaternion", v)
}
func (sl *SpotLight) ReceiveShadow() bool {
	return sl.Get("receiveShadow").Bool()
}
func (sl *SpotLight) SetReceiveShadow(v bool) {
	sl.Set("receiveShadow", v)
}
func (sl *SpotLight) RenderOrder() float64 {
	return sl.Get("renderOrder").Float()
}
func (sl *SpotLight) SetRenderOrder(v float64) {
	sl.Set("renderOrder", v)
}
func (sl *SpotLight) Rotation() *Euler {
	return &Euler{Value: sl.Get("rotation")}
}
func (sl *SpotLight) SetRotation(v *Euler) {
	sl.Set("rotation", v)
}
func (sl *SpotLight) Scale() *Vector3 {
	return &Vector3{Value: sl.Get("scale")}
}
func (sl *SpotLight) SetScale(v *Vector3) {
	sl.Set("scale", v)
}
func (sl *SpotLight) Shadow() *SpotLightShadow {
	return &SpotLightShadow{Value: sl.Get("shadow")}
}
func (sl *SpotLight) SetShadow(v *SpotLightShadow) {
	sl.Set("shadow", v)
}
func (sl *SpotLight) ShadowBias() js.Value {
	return sl.Get("shadowBias")
}
func (sl *SpotLight) SetShadowBias(v js.Value) {
	sl.Set("shadowBias", v)
}
func (sl *SpotLight) ShadowCameraBottom() js.Value {
	return sl.Get("shadowCameraBottom")
}
func (sl *SpotLight) SetShadowCameraBottom(v js.Value) {
	sl.Set("shadowCameraBottom", v)
}
func (sl *SpotLight) ShadowCameraFar() js.Value {
	return sl.Get("shadowCameraFar")
}
func (sl *SpotLight) SetShadowCameraFar(v js.Value) {
	sl.Set("shadowCameraFar", v)
}
func (sl *SpotLight) ShadowCameraFov() js.Value {
	return sl.Get("shadowCameraFov")
}
func (sl *SpotLight) SetShadowCameraFov(v js.Value) {
	sl.Set("shadowCameraFov", v)
}
func (sl *SpotLight) ShadowCameraLeft() js.Value {
	return sl.Get("shadowCameraLeft")
}
func (sl *SpotLight) SetShadowCameraLeft(v js.Value) {
	sl.Set("shadowCameraLeft", v)
}
func (sl *SpotLight) ShadowCameraNear() js.Value {
	return sl.Get("shadowCameraNear")
}
func (sl *SpotLight) SetShadowCameraNear(v js.Value) {
	sl.Set("shadowCameraNear", v)
}
func (sl *SpotLight) ShadowCameraRight() js.Value {
	return sl.Get("shadowCameraRight")
}
func (sl *SpotLight) SetShadowCameraRight(v js.Value) {
	sl.Set("shadowCameraRight", v)
}
func (sl *SpotLight) ShadowCameraTop() js.Value {
	return sl.Get("shadowCameraTop")
}
func (sl *SpotLight) SetShadowCameraTop(v js.Value) {
	sl.Set("shadowCameraTop", v)
}
func (sl *SpotLight) ShadowMapHeight() js.Value {
	return sl.Get("shadowMapHeight")
}
func (sl *SpotLight) SetShadowMapHeight(v js.Value) {
	sl.Set("shadowMapHeight", v)
}
func (sl *SpotLight) ShadowMapWidth() js.Value {
	return sl.Get("shadowMapWidth")
}
func (sl *SpotLight) SetShadowMapWidth(v js.Value) {
	sl.Set("shadowMapWidth", v)
}
func (sl *SpotLight) Target() *Object3D {
	return &Object3D{Value: sl.Get("target")}
}
func (sl *SpotLight) SetTarget(v *Object3D) {
	sl.Set("target", v)
}
func (sl *SpotLight) Type() string {
	return sl.Get("type").String()
}
func (sl *SpotLight) SetType(v string) {
	sl.Set("type", v)
}
func (sl *SpotLight) Up() *Vector3 {
	return &Vector3{Value: sl.Get("up")}
}
func (sl *SpotLight) SetUp(v *Vector3) {
	sl.Set("up", v)
}
func (sl *SpotLight) UserData() js.Value {
	return sl.Get("userData")
}
func (sl *SpotLight) SetUserData(v js.Value) {
	sl.Set("userData", v)
}
func (sl *SpotLight) Uuid() string {
	return sl.Get("uuid").String()
}
func (sl *SpotLight) SetUuid(v string) {
	sl.Set("uuid", v)
}
func (sl *SpotLight) Visible() bool {
	return sl.Get("visible").Bool()
}
func (sl *SpotLight) SetVisible(v bool) {
	sl.Set("visible", v)
}
func (sl *SpotLight) DefaultMatrixAutoUpdate() bool {
	return sl.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sl *SpotLight) SetDefaultMatrixAutoUpdate(v bool) {
	sl.Set("DefaultMatrixAutoUpdate", v)
}
func (sl *SpotLight) DefaultUp() *Vector3 {
	return &Vector3{Value: sl.Get("DefaultUp")}
}
func (sl *SpotLight) SetDefaultUp(v *Vector3) {
	sl.Set("DefaultUp", v)
}
func (sl *SpotLight) Add(object js.Value) *SpotLight {
	return &SpotLight{Value: sl.Call("add", object)}
}
func (sl *SpotLight) AddEventListener(typ string, listener js.Value) {
	sl.Call("addEventListener", typ, listener)
}
func (sl *SpotLight) ApplyMatrix(matrix *Matrix4) {
	sl.Call("applyMatrix", matrix)
}
func (sl *SpotLight) ApplyQuaternion(quaternion *Quaternion) *SpotLight {
	return &SpotLight{Value: sl.Call("applyQuaternion", quaternion)}
}
func (sl *SpotLight) Clone(recursive bool) *SpotLight {
	return &SpotLight{Value: sl.Call("clone", recursive)}
}
func (sl *SpotLight) Copy(source *Object3D, recursive bool) *SpotLight {
	return &SpotLight{Value: sl.Call("copy", source, recursive)}
}
func (sl *SpotLight) DispatchEvent(event js.Value) {
	sl.Call("dispatchEvent", event)
}
func (sl *SpotLight) GetObjectById(id int) *Object3D {
	return &Object3D{Value: sl.Call("getObjectById", id)}
}
func (sl *SpotLight) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: sl.Call("getObjectByName", name)}
}
func (sl *SpotLight) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: sl.Call("getObjectByProperty", name, value)}
}
func (sl *SpotLight) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: sl.Call("getWorldDirection", target)}
}
func (sl *SpotLight) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: sl.Call("getWorldPosition", target)}
}
func (sl *SpotLight) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: sl.Call("getWorldQuaternion", target)}
}
func (sl *SpotLight) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: sl.Call("getWorldScale", target)}
}
func (sl *SpotLight) HasEventListener(typ string, listener js.Value) bool {
	return sl.Call("hasEventListener", typ, listener).Bool()
}
func (sl *SpotLight) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: sl.Call("localToWorld", vector)}
}
func (sl *SpotLight) LookAt(vector *Vector3, y float64, z float64) {
	sl.Call("lookAt", vector, y, z)
}
func (sl *SpotLight) Raycast(raycaster *Raycaster, intersects js.Value) {
	sl.Call("raycast", raycaster, intersects)
}
func (sl *SpotLight) Remove(object js.Value) *SpotLight {
	return &SpotLight{Value: sl.Call("remove", object)}
}
func (sl *SpotLight) RemoveEventListener(typ string, listener js.Value) {
	sl.Call("removeEventListener", typ, listener)
}
func (sl *SpotLight) RotateOnAxis(axis *Vector3, angle float64) *SpotLight {
	return &SpotLight{Value: sl.Call("rotateOnAxis", axis, angle)}
}
func (sl *SpotLight) RotateOnWorldAxis(axis *Vector3, angle float64) *SpotLight {
	return &SpotLight{Value: sl.Call("rotateOnWorldAxis", axis, angle)}
}
func (sl *SpotLight) RotateX(angle float64) *SpotLight {
	return &SpotLight{Value: sl.Call("rotateX", angle)}
}
func (sl *SpotLight) RotateY(angle float64) *SpotLight {
	return &SpotLight{Value: sl.Call("rotateY", angle)}
}
func (sl *SpotLight) RotateZ(angle float64) *SpotLight {
	return &SpotLight{Value: sl.Call("rotateZ", angle)}
}
func (sl *SpotLight) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	sl.Call("setRotationFromAxisAngle", axis, angle)
}
func (sl *SpotLight) SetRotationFromEuler(euler *Euler) {
	sl.Call("setRotationFromEuler", euler)
}
func (sl *SpotLight) SetRotationFromMatrix(m *Matrix4) {
	sl.Call("setRotationFromMatrix", m)
}
func (sl *SpotLight) SetRotationFromQuaternion(q *Quaternion) {
	sl.Call("setRotationFromQuaternion", q)
}
func (sl *SpotLight) ToJSON(meta js.Value) js.Value {
	return sl.Call("toJSON", meta)
}
func (sl *SpotLight) TranslateOnAxis(axis *Vector3, distance float64) *SpotLight {
	return &SpotLight{Value: sl.Call("translateOnAxis", axis, distance)}
}
func (sl *SpotLight) TranslateX(distance float64) *SpotLight {
	return &SpotLight{Value: sl.Call("translateX", distance)}
}
func (sl *SpotLight) TranslateY(distance float64) *SpotLight {
	return &SpotLight{Value: sl.Call("translateY", distance)}
}
func (sl *SpotLight) TranslateZ(distance float64) *SpotLight {
	return &SpotLight{Value: sl.Call("translateZ", distance)}
}
func (sl *SpotLight) Traverse(callback js.Value) {
	sl.Call("traverse", callback)
}
func (sl *SpotLight) TraverseAncestors(callback js.Value) {
	sl.Call("traverseAncestors", callback)
}
func (sl *SpotLight) TraverseVisible(callback js.Value) {
	sl.Call("traverseVisible", callback)
}
func (sl *SpotLight) UpdateMatrix() {
	sl.Call("updateMatrix")
}
func (sl *SpotLight) UpdateMatrixWorld(force bool) {
	sl.Call("updateMatrixWorld", force)
}
func (sl *SpotLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sl.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sl *SpotLight) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: sl.Call("worldToLocal", vector)}
}

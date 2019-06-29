// Code generated from "lights/RectAreaLight.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// RectAreaLight extend: [Light]
type RectAreaLight struct {
	js.Value
}

func NewRectAreaLight(color *Color, intensity float64, width float64, height float64) *RectAreaLight {
	return &RectAreaLight{Value: get("RectAreaLight").New(color, intensity, width, height)}
}
func (ral *RectAreaLight) JSValue() js.Value {
	return ral.Value
}
func (ral *RectAreaLight) CastShadow() bool {
	return ral.Get("castShadow").Bool()
}
func (ral *RectAreaLight) SetCastShadow(v bool) {
	ral.Set("castShadow", v)
}
func (ral *RectAreaLight) Children() js.Value {
	return ral.Get("children")
}
func (ral *RectAreaLight) SetChildren(v js.Value) {
	ral.Set("children", v)
}
func (ral *RectAreaLight) Color() *Color {
	return &Color{Value: ral.Get("color")}
}
func (ral *RectAreaLight) SetColor(v *Color) {
	ral.Set("color", v.JSValue())
}
func (ral *RectAreaLight) FrustumCulled() bool {
	return ral.Get("frustumCulled").Bool()
}
func (ral *RectAreaLight) SetFrustumCulled(v bool) {
	ral.Set("frustumCulled", v)
}
func (ral *RectAreaLight) Height() float64 {
	return ral.Get("height").Float()
}
func (ral *RectAreaLight) SetHeight(v float64) {
	ral.Set("height", v)
}
func (ral *RectAreaLight) Id() int {
	return ral.Get("id").Int()
}
func (ral *RectAreaLight) SetId(v int) {
	ral.Set("id", v)
}
func (ral *RectAreaLight) Intensity() float64 {
	return ral.Get("intensity").Float()
}
func (ral *RectAreaLight) SetIntensity(v float64) {
	ral.Set("intensity", v)
}
func (ral *RectAreaLight) IsLight() bool {
	return ral.Get("isLight").Bool()
}
func (ral *RectAreaLight) SetIsLight(v bool) {
	ral.Set("isLight", v)
}
func (ral *RectAreaLight) IsObject3D() bool {
	return ral.Get("isObject3D").Bool()
}
func (ral *RectAreaLight) SetIsObject3D(v bool) {
	ral.Set("isObject3D", v)
}
func (ral *RectAreaLight) Layers() *Layers {
	return &Layers{Value: ral.Get("layers")}
}
func (ral *RectAreaLight) SetLayers(v *Layers) {
	ral.Set("layers", v.JSValue())
}
func (ral *RectAreaLight) Matrix() *Matrix4 {
	return &Matrix4{Value: ral.Get("matrix")}
}
func (ral *RectAreaLight) SetMatrix(v *Matrix4) {
	ral.Set("matrix", v.JSValue())
}
func (ral *RectAreaLight) MatrixAutoUpdate() bool {
	return ral.Get("matrixAutoUpdate").Bool()
}
func (ral *RectAreaLight) SetMatrixAutoUpdate(v bool) {
	ral.Set("matrixAutoUpdate", v)
}
func (ral *RectAreaLight) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ral.Get("matrixWorld")}
}
func (ral *RectAreaLight) SetMatrixWorld(v *Matrix4) {
	ral.Set("matrixWorld", v.JSValue())
}
func (ral *RectAreaLight) MatrixWorldNeedsUpdate() bool {
	return ral.Get("matrixWorldNeedsUpdate").Bool()
}
func (ral *RectAreaLight) SetMatrixWorldNeedsUpdate(v bool) {
	ral.Set("matrixWorldNeedsUpdate", v)
}
func (ral *RectAreaLight) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ral.Get("modelViewMatrix")}
}
func (ral *RectAreaLight) SetModelViewMatrix(v *Matrix4) {
	ral.Set("modelViewMatrix", v.JSValue())
}
func (ral *RectAreaLight) Name() string {
	return ral.Get("name").String()
}
func (ral *RectAreaLight) SetName(v string) {
	ral.Set("name", v)
}
func (ral *RectAreaLight) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ral.Get("normalMatrix")}
}
func (ral *RectAreaLight) SetNormalMatrix(v *Matrix3) {
	ral.Set("normalMatrix", v.JSValue())
}
func (ral *RectAreaLight) OnAfterRender() js.Value {
	return ral.Get("onAfterRender")
}
func (ral *RectAreaLight) SetOnAfterRender(v js.Value) {
	ral.Set("onAfterRender", v)
}
func (ral *RectAreaLight) OnBeforeRender() js.Value {
	return ral.Get("onBeforeRender")
}
func (ral *RectAreaLight) SetOnBeforeRender(v js.Value) {
	ral.Set("onBeforeRender", v)
}
func (ral *RectAreaLight) Parent() *Object3D {
	return &Object3D{Value: ral.Get("parent")}
}
func (ral *RectAreaLight) SetParent(v *Object3D) {
	ral.Set("parent", v.JSValue())
}
func (ral *RectAreaLight) Position() *Vector3 {
	return &Vector3{Value: ral.Get("position")}
}
func (ral *RectAreaLight) SetPosition(v *Vector3) {
	ral.Set("position", v.JSValue())
}
func (ral *RectAreaLight) Quaternion() *Quaternion {
	return &Quaternion{Value: ral.Get("quaternion")}
}
func (ral *RectAreaLight) SetQuaternion(v *Quaternion) {
	ral.Set("quaternion", v.JSValue())
}
func (ral *RectAreaLight) ReceiveShadow() bool {
	return ral.Get("receiveShadow").Bool()
}
func (ral *RectAreaLight) SetReceiveShadow(v bool) {
	ral.Set("receiveShadow", v)
}
func (ral *RectAreaLight) RenderOrder() float64 {
	return ral.Get("renderOrder").Float()
}
func (ral *RectAreaLight) SetRenderOrder(v float64) {
	ral.Set("renderOrder", v)
}
func (ral *RectAreaLight) Rotation() *Euler {
	return &Euler{Value: ral.Get("rotation")}
}
func (ral *RectAreaLight) SetRotation(v *Euler) {
	ral.Set("rotation", v.JSValue())
}
func (ral *RectAreaLight) Scale() *Vector3 {
	return &Vector3{Value: ral.Get("scale")}
}
func (ral *RectAreaLight) SetScale(v *Vector3) {
	ral.Set("scale", v.JSValue())
}
func (ral *RectAreaLight) Shadow() *LightShadow {
	return &LightShadow{Value: ral.Get("shadow")}
}
func (ral *RectAreaLight) SetShadow(v *LightShadow) {
	ral.Set("shadow", v.JSValue())
}
func (ral *RectAreaLight) ShadowBias() js.Value {
	return ral.Get("shadowBias")
}
func (ral *RectAreaLight) SetShadowBias(v js.Value) {
	ral.Set("shadowBias", v)
}
func (ral *RectAreaLight) ShadowCameraBottom() js.Value {
	return ral.Get("shadowCameraBottom")
}
func (ral *RectAreaLight) SetShadowCameraBottom(v js.Value) {
	ral.Set("shadowCameraBottom", v)
}
func (ral *RectAreaLight) ShadowCameraFar() js.Value {
	return ral.Get("shadowCameraFar")
}
func (ral *RectAreaLight) SetShadowCameraFar(v js.Value) {
	ral.Set("shadowCameraFar", v)
}
func (ral *RectAreaLight) ShadowCameraFov() js.Value {
	return ral.Get("shadowCameraFov")
}
func (ral *RectAreaLight) SetShadowCameraFov(v js.Value) {
	ral.Set("shadowCameraFov", v)
}
func (ral *RectAreaLight) ShadowCameraLeft() js.Value {
	return ral.Get("shadowCameraLeft")
}
func (ral *RectAreaLight) SetShadowCameraLeft(v js.Value) {
	ral.Set("shadowCameraLeft", v)
}
func (ral *RectAreaLight) ShadowCameraNear() js.Value {
	return ral.Get("shadowCameraNear")
}
func (ral *RectAreaLight) SetShadowCameraNear(v js.Value) {
	ral.Set("shadowCameraNear", v)
}
func (ral *RectAreaLight) ShadowCameraRight() js.Value {
	return ral.Get("shadowCameraRight")
}
func (ral *RectAreaLight) SetShadowCameraRight(v js.Value) {
	ral.Set("shadowCameraRight", v)
}
func (ral *RectAreaLight) ShadowCameraTop() js.Value {
	return ral.Get("shadowCameraTop")
}
func (ral *RectAreaLight) SetShadowCameraTop(v js.Value) {
	ral.Set("shadowCameraTop", v)
}
func (ral *RectAreaLight) ShadowMapHeight() js.Value {
	return ral.Get("shadowMapHeight")
}
func (ral *RectAreaLight) SetShadowMapHeight(v js.Value) {
	ral.Set("shadowMapHeight", v)
}
func (ral *RectAreaLight) ShadowMapWidth() js.Value {
	return ral.Get("shadowMapWidth")
}
func (ral *RectAreaLight) SetShadowMapWidth(v js.Value) {
	ral.Set("shadowMapWidth", v)
}
func (ral *RectAreaLight) Type() string {
	return ral.Get("type").String()
}
func (ral *RectAreaLight) SetType(v string) {
	ral.Set("type", v)
}
func (ral *RectAreaLight) Up() *Vector3 {
	return &Vector3{Value: ral.Get("up")}
}
func (ral *RectAreaLight) SetUp(v *Vector3) {
	ral.Set("up", v.JSValue())
}
func (ral *RectAreaLight) UserData() js.Value {
	return ral.Get("userData")
}
func (ral *RectAreaLight) SetUserData(v js.Value) {
	ral.Set("userData", v)
}
func (ral *RectAreaLight) Uuid() string {
	return ral.Get("uuid").String()
}
func (ral *RectAreaLight) SetUuid(v string) {
	ral.Set("uuid", v)
}
func (ral *RectAreaLight) Visible() bool {
	return ral.Get("visible").Bool()
}
func (ral *RectAreaLight) SetVisible(v bool) {
	ral.Set("visible", v)
}
func (ral *RectAreaLight) Width() float64 {
	return ral.Get("width").Float()
}
func (ral *RectAreaLight) SetWidth(v float64) {
	ral.Set("width", v)
}
func (ral *RectAreaLight) DefaultMatrixAutoUpdate() bool {
	return ral.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ral *RectAreaLight) SetDefaultMatrixAutoUpdate(v bool) {
	ral.Set("DefaultMatrixAutoUpdate", v)
}
func (ral *RectAreaLight) DefaultUp() *Vector3 {
	return &Vector3{Value: ral.Get("DefaultUp")}
}
func (ral *RectAreaLight) SetDefaultUp(v *Vector3) {
	ral.Set("DefaultUp", v.JSValue())
}
func (ral *RectAreaLight) Add(object js.Value) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("add", object)}
}
func (ral *RectAreaLight) AddEventListener(typ string, listener js.Value) {
	ral.Call("addEventListener", typ, listener)
}
func (ral *RectAreaLight) ApplyMatrix(matrix *Matrix4) {
	ral.Call("applyMatrix", matrix.JSValue())
}
func (ral *RectAreaLight) ApplyQuaternion(quaternion *Quaternion) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("applyQuaternion", quaternion)}
}
func (ral *RectAreaLight) Clone(recursive bool) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("clone", recursive)}
}
func (ral *RectAreaLight) Copy(source *Object3D, recursive bool) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("copy", source, recursive)}
}
func (ral *RectAreaLight) DispatchEvent(event js.Value) {
	ral.Call("dispatchEvent", event)
}
func (ral *RectAreaLight) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ral.Call("getObjectById", id)}
}
func (ral *RectAreaLight) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ral.Call("getObjectByName", name)}
}
func (ral *RectAreaLight) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ral.Call("getObjectByProperty", name, value)}
}
func (ral *RectAreaLight) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ral.Call("getWorldDirection", target)}
}
func (ral *RectAreaLight) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ral.Call("getWorldPosition", target)}
}
func (ral *RectAreaLight) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ral.Call("getWorldQuaternion", target)}
}
func (ral *RectAreaLight) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ral.Call("getWorldScale", target)}
}
func (ral *RectAreaLight) HasEventListener(typ string, listener js.Value) bool {
	return ral.Call("hasEventListener", typ, listener).Bool()
}
func (ral *RectAreaLight) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ral.Call("localToWorld", vector)}
}
func (ral *RectAreaLight) LookAt(vector *Vector3, y float64, z float64) {
	ral.Call("lookAt", vector, y, z)
}
func (ral *RectAreaLight) Raycast(raycaster *Raycaster, intersects js.Value) {
	ral.Call("raycast", raycaster.JSValue(), intersects)
}
func (ral *RectAreaLight) Remove(object js.Value) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("remove", object)}
}
func (ral *RectAreaLight) RemoveEventListener(typ string, listener js.Value) {
	ral.Call("removeEventListener", typ, listener)
}
func (ral *RectAreaLight) RotateOnAxis(axis *Vector3, angle float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("rotateOnAxis", axis, angle)}
}
func (ral *RectAreaLight) RotateOnWorldAxis(axis *Vector3, angle float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("rotateOnWorldAxis", axis, angle)}
}
func (ral *RectAreaLight) RotateX(angle float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("rotateX", angle)}
}
func (ral *RectAreaLight) RotateY(angle float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("rotateY", angle)}
}
func (ral *RectAreaLight) RotateZ(angle float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("rotateZ", angle)}
}
func (ral *RectAreaLight) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ral.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (ral *RectAreaLight) SetRotationFromEuler(euler *Euler) {
	ral.Call("setRotationFromEuler", euler.JSValue())
}
func (ral *RectAreaLight) SetRotationFromMatrix(m *Matrix4) {
	ral.Call("setRotationFromMatrix", m.JSValue())
}
func (ral *RectAreaLight) SetRotationFromQuaternion(q *Quaternion) {
	ral.Call("setRotationFromQuaternion", q.JSValue())
}
func (ral *RectAreaLight) ToJSON(meta js.Value) js.Value {
	return ral.Call("toJSON", meta)
}
func (ral *RectAreaLight) TranslateOnAxis(axis *Vector3, distance float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("translateOnAxis", axis, distance)}
}
func (ral *RectAreaLight) TranslateX(distance float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("translateX", distance)}
}
func (ral *RectAreaLight) TranslateY(distance float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("translateY", distance)}
}
func (ral *RectAreaLight) TranslateZ(distance float64) *RectAreaLight {
	return &RectAreaLight{Value: ral.Call("translateZ", distance)}
}
func (ral *RectAreaLight) Traverse(callback js.Value) {
	ral.Call("traverse", callback)
}
func (ral *RectAreaLight) TraverseAncestors(callback js.Value) {
	ral.Call("traverseAncestors", callback)
}
func (ral *RectAreaLight) TraverseVisible(callback js.Value) {
	ral.Call("traverseVisible", callback)
}
func (ral *RectAreaLight) UpdateMatrix() {
	ral.Call("updateMatrix")
}
func (ral *RectAreaLight) UpdateMatrixWorld(force bool) {
	ral.Call("updateMatrixWorld", force)
}
func (ral *RectAreaLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ral.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ral *RectAreaLight) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ral.Call("worldToLocal", vector)}
}

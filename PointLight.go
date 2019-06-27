// Code generated from "lights/PointLight.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PointLight extend: [Light]
type PointLight struct {
	js.Value
}

func NewPointLight(color *Color, intensity float64, distance float64, decay float64) *PointLight {
	return &PointLight{Value: get("PointLight").New(color, intensity, distance, decay)}
}
func (pl *PointLight) JSValue() js.Value {
	return pl.Value
}
func (pl *PointLight) CastShadow() bool {
	return pl.Get("castShadow").Bool()
}
func (pl *PointLight) SetCastShadow(v bool) {
	pl.Set("castShadow", v)
}
func (pl *PointLight) Children() js.Value {
	return pl.Get("children")
}
func (pl *PointLight) SetChildren(v js.Value) {
	pl.Set("children", v)
}
func (pl *PointLight) Color() *Color {
	return &Color{Value: pl.Get("color")}
}
func (pl *PointLight) SetColor(v *Color) {
	pl.Set("color", v.Value)
}
func (pl *PointLight) Decay() float64 {
	return pl.Get("decay").Float()
}
func (pl *PointLight) SetDecay(v float64) {
	pl.Set("decay", v)
}
func (pl *PointLight) Distance() float64 {
	return pl.Get("distance").Float()
}
func (pl *PointLight) SetDistance(v float64) {
	pl.Set("distance", v)
}
func (pl *PointLight) FrustumCulled() bool {
	return pl.Get("frustumCulled").Bool()
}
func (pl *PointLight) SetFrustumCulled(v bool) {
	pl.Set("frustumCulled", v)
}
func (pl *PointLight) Id() int {
	return pl.Get("id").Int()
}
func (pl *PointLight) SetId(v int) {
	pl.Set("id", v)
}
func (pl *PointLight) Intensity() float64 {
	return pl.Get("intensity").Float()
}
func (pl *PointLight) SetIntensity(v float64) {
	pl.Set("intensity", v)
}
func (pl *PointLight) IsLight() bool {
	return pl.Get("isLight").Bool()
}
func (pl *PointLight) SetIsLight(v bool) {
	pl.Set("isLight", v)
}
func (pl *PointLight) IsObject3D() bool {
	return pl.Get("isObject3D").Bool()
}
func (pl *PointLight) SetIsObject3D(v bool) {
	pl.Set("isObject3D", v)
}
func (pl *PointLight) Layers() *Layers {
	return &Layers{Value: pl.Get("layers")}
}
func (pl *PointLight) SetLayers(v *Layers) {
	pl.Set("layers", v.Value)
}
func (pl *PointLight) Matrix() *Matrix4 {
	return &Matrix4{Value: pl.Get("matrix")}
}
func (pl *PointLight) SetMatrix(v *Matrix4) {
	pl.Set("matrix", v.Value)
}
func (pl *PointLight) MatrixAutoUpdate() bool {
	return pl.Get("matrixAutoUpdate").Bool()
}
func (pl *PointLight) SetMatrixAutoUpdate(v bool) {
	pl.Set("matrixAutoUpdate", v)
}
func (pl *PointLight) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: pl.Get("matrixWorld")}
}
func (pl *PointLight) SetMatrixWorld(v *Matrix4) {
	pl.Set("matrixWorld", v.Value)
}
func (pl *PointLight) MatrixWorldNeedsUpdate() bool {
	return pl.Get("matrixWorldNeedsUpdate").Bool()
}
func (pl *PointLight) SetMatrixWorldNeedsUpdate(v bool) {
	pl.Set("matrixWorldNeedsUpdate", v)
}
func (pl *PointLight) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: pl.Get("modelViewMatrix")}
}
func (pl *PointLight) SetModelViewMatrix(v *Matrix4) {
	pl.Set("modelViewMatrix", v.Value)
}
func (pl *PointLight) Name() string {
	return pl.Get("name").String()
}
func (pl *PointLight) SetName(v string) {
	pl.Set("name", v)
}
func (pl *PointLight) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: pl.Get("normalMatrix")}
}
func (pl *PointLight) SetNormalMatrix(v *Matrix3) {
	pl.Set("normalMatrix", v.Value)
}
func (pl *PointLight) OnAfterRender() js.Value {
	return pl.Get("onAfterRender")
}
func (pl *PointLight) SetOnAfterRender(v js.Value) {
	pl.Set("onAfterRender", v)
}
func (pl *PointLight) OnBeforeRender() js.Value {
	return pl.Get("onBeforeRender")
}
func (pl *PointLight) SetOnBeforeRender(v js.Value) {
	pl.Set("onBeforeRender", v)
}
func (pl *PointLight) Parent() *Object3D {
	return &Object3D{Value: pl.Get("parent")}
}
func (pl *PointLight) SetParent(v *Object3D) {
	pl.Set("parent", v.Value)
}
func (pl *PointLight) Position() *Vector3 {
	return &Vector3{Value: pl.Get("position")}
}
func (pl *PointLight) SetPosition(v *Vector3) {
	pl.Set("position", v.Value)
}
func (pl *PointLight) Power() float64 {
	return pl.Get("power").Float()
}
func (pl *PointLight) SetPower(v float64) {
	pl.Set("power", v)
}
func (pl *PointLight) Quaternion() *Quaternion {
	return &Quaternion{Value: pl.Get("quaternion")}
}
func (pl *PointLight) SetQuaternion(v *Quaternion) {
	pl.Set("quaternion", v.Value)
}
func (pl *PointLight) ReceiveShadow() bool {
	return pl.Get("receiveShadow").Bool()
}
func (pl *PointLight) SetReceiveShadow(v bool) {
	pl.Set("receiveShadow", v)
}
func (pl *PointLight) RenderOrder() float64 {
	return pl.Get("renderOrder").Float()
}
func (pl *PointLight) SetRenderOrder(v float64) {
	pl.Set("renderOrder", v)
}
func (pl *PointLight) Rotation() *Euler {
	return &Euler{Value: pl.Get("rotation")}
}
func (pl *PointLight) SetRotation(v *Euler) {
	pl.Set("rotation", v.Value)
}
func (pl *PointLight) Scale() *Vector3 {
	return &Vector3{Value: pl.Get("scale")}
}
func (pl *PointLight) SetScale(v *Vector3) {
	pl.Set("scale", v.Value)
}
func (pl *PointLight) Shadow() *PointLightShadow {
	return &PointLightShadow{Value: pl.Get("shadow")}
}
func (pl *PointLight) SetShadow(v *PointLightShadow) {
	pl.Set("shadow", v.Value)
}
func (pl *PointLight) ShadowBias() js.Value {
	return pl.Get("shadowBias")
}
func (pl *PointLight) SetShadowBias(v js.Value) {
	pl.Set("shadowBias", v)
}
func (pl *PointLight) ShadowCameraBottom() js.Value {
	return pl.Get("shadowCameraBottom")
}
func (pl *PointLight) SetShadowCameraBottom(v js.Value) {
	pl.Set("shadowCameraBottom", v)
}
func (pl *PointLight) ShadowCameraFar() js.Value {
	return pl.Get("shadowCameraFar")
}
func (pl *PointLight) SetShadowCameraFar(v js.Value) {
	pl.Set("shadowCameraFar", v)
}
func (pl *PointLight) ShadowCameraFov() js.Value {
	return pl.Get("shadowCameraFov")
}
func (pl *PointLight) SetShadowCameraFov(v js.Value) {
	pl.Set("shadowCameraFov", v)
}
func (pl *PointLight) ShadowCameraLeft() js.Value {
	return pl.Get("shadowCameraLeft")
}
func (pl *PointLight) SetShadowCameraLeft(v js.Value) {
	pl.Set("shadowCameraLeft", v)
}
func (pl *PointLight) ShadowCameraNear() js.Value {
	return pl.Get("shadowCameraNear")
}
func (pl *PointLight) SetShadowCameraNear(v js.Value) {
	pl.Set("shadowCameraNear", v)
}
func (pl *PointLight) ShadowCameraRight() js.Value {
	return pl.Get("shadowCameraRight")
}
func (pl *PointLight) SetShadowCameraRight(v js.Value) {
	pl.Set("shadowCameraRight", v)
}
func (pl *PointLight) ShadowCameraTop() js.Value {
	return pl.Get("shadowCameraTop")
}
func (pl *PointLight) SetShadowCameraTop(v js.Value) {
	pl.Set("shadowCameraTop", v)
}
func (pl *PointLight) ShadowMapHeight() js.Value {
	return pl.Get("shadowMapHeight")
}
func (pl *PointLight) SetShadowMapHeight(v js.Value) {
	pl.Set("shadowMapHeight", v)
}
func (pl *PointLight) ShadowMapWidth() js.Value {
	return pl.Get("shadowMapWidth")
}
func (pl *PointLight) SetShadowMapWidth(v js.Value) {
	pl.Set("shadowMapWidth", v)
}
func (pl *PointLight) Type() string {
	return pl.Get("type").String()
}
func (pl *PointLight) SetType(v string) {
	pl.Set("type", v)
}
func (pl *PointLight) Up() *Vector3 {
	return &Vector3{Value: pl.Get("up")}
}
func (pl *PointLight) SetUp(v *Vector3) {
	pl.Set("up", v.Value)
}
func (pl *PointLight) UserData() js.Value {
	return pl.Get("userData")
}
func (pl *PointLight) SetUserData(v js.Value) {
	pl.Set("userData", v)
}
func (pl *PointLight) Uuid() string {
	return pl.Get("uuid").String()
}
func (pl *PointLight) SetUuid(v string) {
	pl.Set("uuid", v)
}
func (pl *PointLight) Visible() bool {
	return pl.Get("visible").Bool()
}
func (pl *PointLight) SetVisible(v bool) {
	pl.Set("visible", v)
}
func (pl *PointLight) DefaultMatrixAutoUpdate() bool {
	return pl.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pl *PointLight) SetDefaultMatrixAutoUpdate(v bool) {
	pl.Set("DefaultMatrixAutoUpdate", v)
}
func (pl *PointLight) DefaultUp() *Vector3 {
	return &Vector3{Value: pl.Get("DefaultUp")}
}
func (pl *PointLight) SetDefaultUp(v *Vector3) {
	pl.Set("DefaultUp", v.Value)
}
func (pl *PointLight) Add(object js.Value) *PointLight {
	return &PointLight{Value: pl.Call("add", object)}
}
func (pl *PointLight) AddEventListener(typ string, listener js.Value) {
	pl.Call("addEventListener", typ, listener)
}
func (pl *PointLight) ApplyMatrix(matrix *Matrix4) {
	pl.Call("applyMatrix", matrix)
}
func (pl *PointLight) ApplyQuaternion(quaternion *Quaternion) *PointLight {
	return &PointLight{Value: pl.Call("applyQuaternion", quaternion)}
}
func (pl *PointLight) Clone(recursive bool) *PointLight {
	return &PointLight{Value: pl.Call("clone", recursive)}
}
func (pl *PointLight) Copy(source *Object3D, recursive bool) *PointLight {
	return &PointLight{Value: pl.Call("copy", source, recursive)}
}
func (pl *PointLight) DispatchEvent(event js.Value) {
	pl.Call("dispatchEvent", event)
}
func (pl *PointLight) GetObjectById(id int) *Object3D {
	return &Object3D{Value: pl.Call("getObjectById", id)}
}
func (pl *PointLight) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: pl.Call("getObjectByName", name)}
}
func (pl *PointLight) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: pl.Call("getObjectByProperty", name, value)}
}
func (pl *PointLight) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: pl.Call("getWorldDirection", target)}
}
func (pl *PointLight) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: pl.Call("getWorldPosition", target)}
}
func (pl *PointLight) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: pl.Call("getWorldQuaternion", target)}
}
func (pl *PointLight) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: pl.Call("getWorldScale", target)}
}
func (pl *PointLight) HasEventListener(typ string, listener js.Value) bool {
	return pl.Call("hasEventListener", typ, listener).Bool()
}
func (pl *PointLight) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: pl.Call("localToWorld", vector)}
}
func (pl *PointLight) LookAt(vector *Vector3, y float64, z float64) {
	pl.Call("lookAt", vector, y, z)
}
func (pl *PointLight) Raycast(raycaster *Raycaster, intersects js.Value) {
	pl.Call("raycast", raycaster, intersects)
}
func (pl *PointLight) Remove(object js.Value) *PointLight {
	return &PointLight{Value: pl.Call("remove", object)}
}
func (pl *PointLight) RemoveEventListener(typ string, listener js.Value) {
	pl.Call("removeEventListener", typ, listener)
}
func (pl *PointLight) RotateOnAxis(axis *Vector3, angle float64) *PointLight {
	return &PointLight{Value: pl.Call("rotateOnAxis", axis, angle)}
}
func (pl *PointLight) RotateOnWorldAxis(axis *Vector3, angle float64) *PointLight {
	return &PointLight{Value: pl.Call("rotateOnWorldAxis", axis, angle)}
}
func (pl *PointLight) RotateX(angle float64) *PointLight {
	return &PointLight{Value: pl.Call("rotateX", angle)}
}
func (pl *PointLight) RotateY(angle float64) *PointLight {
	return &PointLight{Value: pl.Call("rotateY", angle)}
}
func (pl *PointLight) RotateZ(angle float64) *PointLight {
	return &PointLight{Value: pl.Call("rotateZ", angle)}
}
func (pl *PointLight) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	pl.Call("setRotationFromAxisAngle", axis, angle)
}
func (pl *PointLight) SetRotationFromEuler(euler *Euler) {
	pl.Call("setRotationFromEuler", euler)
}
func (pl *PointLight) SetRotationFromMatrix(m *Matrix4) {
	pl.Call("setRotationFromMatrix", m)
}
func (pl *PointLight) SetRotationFromQuaternion(q *Quaternion) {
	pl.Call("setRotationFromQuaternion", q)
}
func (pl *PointLight) ToJSON(meta js.Value) js.Value {
	return pl.Call("toJSON", meta)
}
func (pl *PointLight) TranslateOnAxis(axis *Vector3, distance float64) *PointLight {
	return &PointLight{Value: pl.Call("translateOnAxis", axis, distance)}
}
func (pl *PointLight) TranslateX(distance float64) *PointLight {
	return &PointLight{Value: pl.Call("translateX", distance)}
}
func (pl *PointLight) TranslateY(distance float64) *PointLight {
	return &PointLight{Value: pl.Call("translateY", distance)}
}
func (pl *PointLight) TranslateZ(distance float64) *PointLight {
	return &PointLight{Value: pl.Call("translateZ", distance)}
}
func (pl *PointLight) Traverse(callback js.Value) {
	pl.Call("traverse", callback)
}
func (pl *PointLight) TraverseAncestors(callback js.Value) {
	pl.Call("traverseAncestors", callback)
}
func (pl *PointLight) TraverseVisible(callback js.Value) {
	pl.Call("traverseVisible", callback)
}
func (pl *PointLight) UpdateMatrix() {
	pl.Call("updateMatrix")
}
func (pl *PointLight) UpdateMatrixWorld(force bool) {
	pl.Call("updateMatrixWorld", force)
}
func (pl *PointLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pl.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pl *PointLight) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: pl.Call("worldToLocal", vector)}
}

// PointLightShadow extend: [LightShadow]
type PointLightShadow struct {
	js.Value
}

func NewPointLightShadow(camera Camera) *PointLightShadow {
	return &PointLightShadow{Value: get("PointLightShadow").New(camera.JSValue())}
}
func (pls *PointLightShadow) JSValue() js.Value {
	return pls.Value
}
func (pls *PointLightShadow) Bias() float64 {
	return pls.Get("bias").Float()
}
func (pls *PointLightShadow) SetBias(v float64) {
	pls.Set("bias", v)
}
func (pls *PointLightShadow) Camera() *PerspectiveCamera {
	return &PerspectiveCamera{Value: pls.Get("camera")}
}
func (pls *PointLightShadow) SetCamera(v *PerspectiveCamera) {
	pls.Set("camera", v.Value)
}
func (pls *PointLightShadow) Map() js.Value {
	return pls.Get("map")
}
func (pls *PointLightShadow) SetMap(v js.Value) {
	pls.Set("map", v)
}
func (pls *PointLightShadow) MapSize() *Vector2 {
	return &Vector2{Value: pls.Get("mapSize")}
}
func (pls *PointLightShadow) SetMapSize(v *Vector2) {
	pls.Set("mapSize", v.Value)
}
func (pls *PointLightShadow) Matrix() *Matrix4 {
	return &Matrix4{Value: pls.Get("matrix")}
}
func (pls *PointLightShadow) SetMatrix(v *Matrix4) {
	pls.Set("matrix", v.Value)
}
func (pls *PointLightShadow) Radius() float64 {
	return pls.Get("radius").Float()
}
func (pls *PointLightShadow) SetRadius(v float64) {
	pls.Set("radius", v)
}
func (pls *PointLightShadow) Clone(recursive bool) *PointLightShadow {
	return &PointLightShadow{Value: pls.Call("clone", recursive)}
}
func (pls *PointLightShadow) Copy(source *LightShadow) *PointLightShadow {
	return &PointLightShadow{Value: pls.Call("copy", source)}
}
func (pls *PointLightShadow) ToJSON() js.Value {
	return pls.Call("toJSON")
}

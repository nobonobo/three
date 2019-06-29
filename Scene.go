// Code generated from "scenes/Scene.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Scene interface {
	JSValue() js.Value
	AutoUpdate() bool
	SetAutoUpdate(v bool)
	Background() js.Value
	SetBackground(v js.Value)
	CastShadow() bool
	SetCastShadow(v bool)
	Children() js.Value
	SetChildren(v js.Value)
	Fog() js.Value
	SetFog(v js.Value)
	FrustumCulled() bool
	SetFrustumCulled(v bool)
	Id() int
	SetId(v int)
	IsObject3D() bool
	SetIsObject3D(v bool)
	Layers() *Layers
	SetLayers(v *Layers)
	Matrix() *Matrix4
	SetMatrix(v *Matrix4)
	MatrixAutoUpdate() bool
	SetMatrixAutoUpdate(v bool)
	MatrixWorld() *Matrix4
	SetMatrixWorld(v *Matrix4)
	MatrixWorldNeedsUpdate() bool
	SetMatrixWorldNeedsUpdate(v bool)
	ModelViewMatrix() *Matrix4
	SetModelViewMatrix(v *Matrix4)
	Name() string
	SetName(v string)
	NormalMatrix() *Matrix3
	SetNormalMatrix(v *Matrix3)
	OnAfterRender() js.Value
	SetOnAfterRender(v js.Value)
	OnBeforeRender() js.Value
	SetOnBeforeRender(v js.Value)
	OverrideMaterial() Material
	SetOverrideMaterial(v Material)
	Parent() *Object3D
	SetParent(v *Object3D)
	Position() *Vector3
	SetPosition(v *Vector3)
	Quaternion() *Quaternion
	SetQuaternion(v *Quaternion)
	ReceiveShadow() bool
	SetReceiveShadow(v bool)
	RenderOrder() float64
	SetRenderOrder(v float64)
	Rotation() *Euler
	SetRotation(v *Euler)
	Scale() *Vector3
	SetScale(v *Vector3)
	Type() string
	SetType(v string)
	Up() *Vector3
	SetUp(v *Vector3)
	UserData() js.Value
	SetUserData(v js.Value)
	Uuid() string
	SetUuid(v string)
	Visible() bool
	SetVisible(v bool)
	DefaultMatrixAutoUpdate() bool
	SetDefaultMatrixAutoUpdate(v bool)
	DefaultUp() *Vector3
	SetDefaultUp(v *Vector3)
	Add(object js.Value) Scene
	AddEventListener(typ string, listener js.Value)
	ApplyMatrix(matrix *Matrix4)
	ApplyQuaternion(quaternion *Quaternion) Scene
	Clone(recursive bool) Scene
	Copy(source Scene, recursive bool) Scene
	DispatchEvent(event js.Value)
	GetObjectById(id int) *Object3D
	GetObjectByName(name string) *Object3D
	GetObjectByProperty(name string, value string) *Object3D
	GetWorldDirection(target *Vector3) *Vector3
	GetWorldPosition(target *Vector3) *Vector3
	GetWorldQuaternion(target *Quaternion) *Quaternion
	GetWorldScale(target *Vector3) *Vector3
	HasEventListener(typ string, listener js.Value) bool
	LocalToWorld(vector *Vector3) *Vector3
	LookAt(vector *Vector3, y float64, z float64)
	Raycast(raycaster *Raycaster, intersects js.Value)
	Remove(object js.Value) Scene
	RemoveEventListener(typ string, listener js.Value)
	RotateOnAxis(axis *Vector3, angle float64) Scene
	RotateOnWorldAxis(axis *Vector3, angle float64) Scene
	RotateX(angle float64) Scene
	RotateY(angle float64) Scene
	RotateZ(angle float64) Scene
	SetRotationFromAxisAngle(axis *Vector3, angle float64)
	SetRotationFromEuler(euler *Euler)
	SetRotationFromMatrix(m *Matrix4)
	SetRotationFromQuaternion(q *Quaternion)
	ToJSON(meta js.Value) js.Value
	TranslateOnAxis(axis *Vector3, distance float64) Scene
	TranslateX(distance float64) Scene
	TranslateY(distance float64) Scene
	TranslateZ(distance float64) Scene
	Traverse(callback js.Value)
	TraverseAncestors(callback js.Value)
	TraverseVisible(callback js.Value)
	UpdateMatrix()
	UpdateMatrixWorld(force bool)
	UpdateWorldMatrix(updateParents bool, updateChildren bool)
	WorldToLocal(vector *Vector3) *Vector3
}

// SceneImpl extend: [Object3D]
type SceneImpl struct {
	js.Value
}

func NewScene() *SceneImpl {
	return &SceneImpl{Value: get("Scene").New()}
}
func (ss *SceneImpl) JSValue() js.Value {
	return ss.Value
}
func (ss *SceneImpl) AutoUpdate() bool {
	return ss.Get("autoUpdate").Bool()
}
func (ss *SceneImpl) SetAutoUpdate(v bool) {
	ss.Set("autoUpdate", v)
}
func (ss *SceneImpl) Background() js.Value {
	return ss.Get("background")
}
func (ss *SceneImpl) SetBackground(v js.Value) {
	ss.Set("background", v)
}
func (ss *SceneImpl) CastShadow() bool {
	return ss.Get("castShadow").Bool()
}
func (ss *SceneImpl) SetCastShadow(v bool) {
	ss.Set("castShadow", v)
}
func (ss *SceneImpl) Children() js.Value {
	return ss.Get("children")
}
func (ss *SceneImpl) SetChildren(v js.Value) {
	ss.Set("children", v)
}
func (ss *SceneImpl) Fog() js.Value {
	return ss.Get("fog")
}
func (ss *SceneImpl) SetFog(v js.Value) {
	ss.Set("fog", v)
}
func (ss *SceneImpl) FrustumCulled() bool {
	return ss.Get("frustumCulled").Bool()
}
func (ss *SceneImpl) SetFrustumCulled(v bool) {
	ss.Set("frustumCulled", v)
}
func (ss *SceneImpl) Id() int {
	return ss.Get("id").Int()
}
func (ss *SceneImpl) SetId(v int) {
	ss.Set("id", v)
}
func (ss *SceneImpl) IsObject3D() bool {
	return ss.Get("isObject3D").Bool()
}
func (ss *SceneImpl) SetIsObject3D(v bool) {
	ss.Set("isObject3D", v)
}
func (ss *SceneImpl) Layers() *Layers {
	return &Layers{Value: ss.Get("layers")}
}
func (ss *SceneImpl) SetLayers(v *Layers) {
	ss.Set("layers", v.JSValue())
}
func (ss *SceneImpl) Matrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrix")}
}
func (ss *SceneImpl) SetMatrix(v *Matrix4) {
	ss.Set("matrix", v.JSValue())
}
func (ss *SceneImpl) MatrixAutoUpdate() bool {
	return ss.Get("matrixAutoUpdate").Bool()
}
func (ss *SceneImpl) SetMatrixAutoUpdate(v bool) {
	ss.Set("matrixAutoUpdate", v)
}
func (ss *SceneImpl) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrixWorld")}
}
func (ss *SceneImpl) SetMatrixWorld(v *Matrix4) {
	ss.Set("matrixWorld", v.JSValue())
}
func (ss *SceneImpl) MatrixWorldNeedsUpdate() bool {
	return ss.Get("matrixWorldNeedsUpdate").Bool()
}
func (ss *SceneImpl) SetMatrixWorldNeedsUpdate(v bool) {
	ss.Set("matrixWorldNeedsUpdate", v)
}
func (ss *SceneImpl) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("modelViewMatrix")}
}
func (ss *SceneImpl) SetModelViewMatrix(v *Matrix4) {
	ss.Set("modelViewMatrix", v.JSValue())
}
func (ss *SceneImpl) Name() string {
	return ss.Get("name").String()
}
func (ss *SceneImpl) SetName(v string) {
	ss.Set("name", v)
}
func (ss *SceneImpl) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ss.Get("normalMatrix")}
}
func (ss *SceneImpl) SetNormalMatrix(v *Matrix3) {
	ss.Set("normalMatrix", v.JSValue())
}
func (ss *SceneImpl) OnAfterRender() js.Value {
	return ss.Get("onAfterRender")
}
func (ss *SceneImpl) SetOnAfterRender(v js.Value) {
	ss.Set("onAfterRender", v)
}
func (ss *SceneImpl) OnBeforeRender() js.Value {
	return ss.Get("onBeforeRender")
}
func (ss *SceneImpl) SetOnBeforeRender(v js.Value) {
	ss.Set("onBeforeRender", v)
}
func (ss *SceneImpl) OverrideMaterial() Material {
	return &MaterialImpl{Value: ss.Get("overrideMaterial")}
}
func (ss *SceneImpl) SetOverrideMaterial(v Material) {
	ss.Set("overrideMaterial", v.JSValue())
}
func (ss *SceneImpl) Parent() *Object3D {
	return &Object3D{Value: ss.Get("parent")}
}
func (ss *SceneImpl) SetParent(v *Object3D) {
	ss.Set("parent", v.JSValue())
}
func (ss *SceneImpl) Position() *Vector3 {
	return &Vector3{Value: ss.Get("position")}
}
func (ss *SceneImpl) SetPosition(v *Vector3) {
	ss.Set("position", v.JSValue())
}
func (ss *SceneImpl) Quaternion() *Quaternion {
	return &Quaternion{Value: ss.Get("quaternion")}
}
func (ss *SceneImpl) SetQuaternion(v *Quaternion) {
	ss.Set("quaternion", v.JSValue())
}
func (ss *SceneImpl) ReceiveShadow() bool {
	return ss.Get("receiveShadow").Bool()
}
func (ss *SceneImpl) SetReceiveShadow(v bool) {
	ss.Set("receiveShadow", v)
}
func (ss *SceneImpl) RenderOrder() float64 {
	return ss.Get("renderOrder").Float()
}
func (ss *SceneImpl) SetRenderOrder(v float64) {
	ss.Set("renderOrder", v)
}
func (ss *SceneImpl) Rotation() *Euler {
	return &Euler{Value: ss.Get("rotation")}
}
func (ss *SceneImpl) SetRotation(v *Euler) {
	ss.Set("rotation", v.JSValue())
}
func (ss *SceneImpl) Scale() *Vector3 {
	return &Vector3{Value: ss.Get("scale")}
}
func (ss *SceneImpl) SetScale(v *Vector3) {
	ss.Set("scale", v.JSValue())
}
func (ss *SceneImpl) Type() string {
	return ss.Get("type").String()
}
func (ss *SceneImpl) SetType(v string) {
	ss.Set("type", v)
}
func (ss *SceneImpl) Up() *Vector3 {
	return &Vector3{Value: ss.Get("up")}
}
func (ss *SceneImpl) SetUp(v *Vector3) {
	ss.Set("up", v.JSValue())
}
func (ss *SceneImpl) UserData() js.Value {
	return ss.Get("userData")
}
func (ss *SceneImpl) SetUserData(v js.Value) {
	ss.Set("userData", v)
}
func (ss *SceneImpl) Uuid() string {
	return ss.Get("uuid").String()
}
func (ss *SceneImpl) SetUuid(v string) {
	ss.Set("uuid", v)
}
func (ss *SceneImpl) Visible() bool {
	return ss.Get("visible").Bool()
}
func (ss *SceneImpl) SetVisible(v bool) {
	ss.Set("visible", v)
}
func (ss *SceneImpl) DefaultMatrixAutoUpdate() bool {
	return ss.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ss *SceneImpl) SetDefaultMatrixAutoUpdate(v bool) {
	ss.Set("DefaultMatrixAutoUpdate", v)
}
func (ss *SceneImpl) DefaultUp() *Vector3 {
	return &Vector3{Value: ss.Get("DefaultUp")}
}
func (ss *SceneImpl) SetDefaultUp(v *Vector3) {
	ss.Set("DefaultUp", v.JSValue())
}
func (ss *SceneImpl) Add(object js.Value) Scene {
	return &SceneImpl{Value: ss.Call("add", object)}
}
func (ss *SceneImpl) AddEventListener(typ string, listener js.Value) {
	ss.Call("addEventListener", typ, listener)
}
func (ss *SceneImpl) ApplyMatrix(matrix *Matrix4) {
	ss.Call("applyMatrix", matrix.JSValue())
}
func (ss *SceneImpl) ApplyQuaternion(quaternion *Quaternion) Scene {
	return &SceneImpl{Value: ss.Call("applyQuaternion", quaternion)}
}
func (ss *SceneImpl) Clone(recursive bool) Scene {
	return &SceneImpl{Value: ss.Call("clone", recursive)}
}
func (ss *SceneImpl) Copy(source Scene, recursive bool) Scene {
	return &SceneImpl{Value: ss.Call("copy", source, recursive)}
}
func (ss *SceneImpl) DispatchEvent(event js.Value) {
	ss.Call("dispatchEvent", event)
}
func (ss *SceneImpl) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ss.Call("getObjectById", id)}
}
func (ss *SceneImpl) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByName", name)}
}
func (ss *SceneImpl) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByProperty", name, value)}
}
func (ss *SceneImpl) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldDirection", target)}
}
func (ss *SceneImpl) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldPosition", target)}
}
func (ss *SceneImpl) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ss.Call("getWorldQuaternion", target)}
}
func (ss *SceneImpl) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldScale", target)}
}
func (ss *SceneImpl) HasEventListener(typ string, listener js.Value) bool {
	return ss.Call("hasEventListener", typ, listener).Bool()
}
func (ss *SceneImpl) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("localToWorld", vector)}
}
func (ss *SceneImpl) LookAt(vector *Vector3, y float64, z float64) {
	ss.Call("lookAt", vector, y, z)
}
func (ss *SceneImpl) Raycast(raycaster *Raycaster, intersects js.Value) {
	ss.Call("raycast", raycaster.JSValue(), intersects)
}
func (ss *SceneImpl) Remove(object js.Value) Scene {
	return &SceneImpl{Value: ss.Call("remove", object)}
}
func (ss *SceneImpl) RemoveEventListener(typ string, listener js.Value) {
	ss.Call("removeEventListener", typ, listener)
}
func (ss *SceneImpl) RotateOnAxis(axis *Vector3, angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateOnAxis", axis, angle)}
}
func (ss *SceneImpl) RotateOnWorldAxis(axis *Vector3, angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateOnWorldAxis", axis, angle)}
}
func (ss *SceneImpl) RotateX(angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateX", angle)}
}
func (ss *SceneImpl) RotateY(angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateY", angle)}
}
func (ss *SceneImpl) RotateZ(angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateZ", angle)}
}
func (ss *SceneImpl) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ss.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (ss *SceneImpl) SetRotationFromEuler(euler *Euler) {
	ss.Call("setRotationFromEuler", euler.JSValue())
}
func (ss *SceneImpl) SetRotationFromMatrix(m *Matrix4) {
	ss.Call("setRotationFromMatrix", m.JSValue())
}
func (ss *SceneImpl) SetRotationFromQuaternion(q *Quaternion) {
	ss.Call("setRotationFromQuaternion", q.JSValue())
}
func (ss *SceneImpl) ToJSON(meta js.Value) js.Value {
	return ss.Call("toJSON", meta)
}
func (ss *SceneImpl) TranslateOnAxis(axis *Vector3, distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateOnAxis", axis, distance)}
}
func (ss *SceneImpl) TranslateX(distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateX", distance)}
}
func (ss *SceneImpl) TranslateY(distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateY", distance)}
}
func (ss *SceneImpl) TranslateZ(distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateZ", distance)}
}
func (ss *SceneImpl) Traverse(callback js.Value) {
	ss.Call("traverse", callback)
}
func (ss *SceneImpl) TraverseAncestors(callback js.Value) {
	ss.Call("traverseAncestors", callback)
}
func (ss *SceneImpl) TraverseVisible(callback js.Value) {
	ss.Call("traverseVisible", callback)
}
func (ss *SceneImpl) UpdateMatrix() {
	ss.Call("updateMatrix")
}
func (ss *SceneImpl) UpdateMatrixWorld(force bool) {
	ss.Call("updateMatrixWorld", force)
}
func (ss *SceneImpl) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ss.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ss *SceneImpl) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("worldToLocal", vector)}
}

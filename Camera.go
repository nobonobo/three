// Code generated from "cameras/Camera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Camera interface {
	JSValue() js.Value
	CastShadow() bool
	SetCastShadow(v bool)
	Children() js.Value
	SetChildren(v js.Value)
	FrustumCulled() bool
	SetFrustumCulled(v bool)
	Id() int
	SetId(v int)
	IsCamera() bool
	SetIsCamera(v bool)
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
	MatrixWorldInverse() *Matrix4
	SetMatrixWorldInverse(v *Matrix4)
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
	Parent() *Object3D
	SetParent(v *Object3D)
	Position() *Vector3
	SetPosition(v *Vector3)
	ProjectionMatrix() *Matrix4
	SetProjectionMatrix(v *Matrix4)
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
	Add(object js.Value) Camera
	AddEventListener(typ string, listener js.Value)
	ApplyMatrix(matrix *Matrix4)
	ApplyQuaternion(quaternion *Quaternion) Camera
	Clone(recursive bool) Camera
	Copy(source Camera, recursive bool) Camera
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
	Remove(object js.Value) Camera
	RemoveEventListener(typ string, listener js.Value)
	RotateOnAxis(axis *Vector3, angle float64) Camera
	RotateOnWorldAxis(axis *Vector3, angle float64) Camera
	RotateX(angle float64) Camera
	RotateY(angle float64) Camera
	RotateZ(angle float64) Camera
	SetRotationFromAxisAngle(axis *Vector3, angle float64)
	SetRotationFromEuler(euler *Euler)
	SetRotationFromMatrix(m *Matrix4)
	SetRotationFromQuaternion(q *Quaternion)
	ToJSON(meta js.Value) js.Value
	TranslateOnAxis(axis *Vector3, distance float64) Camera
	TranslateX(distance float64) Camera
	TranslateY(distance float64) Camera
	TranslateZ(distance float64) Camera
	Traverse(callback js.Value)
	TraverseAncestors(callback js.Value)
	TraverseVisible(callback js.Value)
	UpdateMatrix()
	UpdateMatrixWorld(force bool)
	UpdateWorldMatrix(updateParents bool, updateChildren bool)
	WorldToLocal(vector *Vector3) *Vector3
}

// CameraImpl extend: [Object3D]
type CameraImpl struct {
	js.Value
}

func NewCamera() *CameraImpl {
	return &CameraImpl{Value: get("Camera").New()}
}
func (cc *CameraImpl) JSValue() js.Value {
	return cc.Value
}
func (cc *CameraImpl) CastShadow() bool {
	return cc.Get("castShadow").Bool()
}
func (cc *CameraImpl) SetCastShadow(v bool) {
	cc.Set("castShadow", v)
}
func (cc *CameraImpl) Children() js.Value {
	return cc.Get("children")
}
func (cc *CameraImpl) SetChildren(v js.Value) {
	cc.Set("children", v)
}
func (cc *CameraImpl) FrustumCulled() bool {
	return cc.Get("frustumCulled").Bool()
}
func (cc *CameraImpl) SetFrustumCulled(v bool) {
	cc.Set("frustumCulled", v)
}
func (cc *CameraImpl) Id() int {
	return cc.Get("id").Int()
}
func (cc *CameraImpl) SetId(v int) {
	cc.Set("id", v)
}
func (cc *CameraImpl) IsCamera() bool {
	return cc.Get("isCamera").Bool()
}
func (cc *CameraImpl) SetIsCamera(v bool) {
	cc.Set("isCamera", v)
}
func (cc *CameraImpl) IsObject3D() bool {
	return cc.Get("isObject3D").Bool()
}
func (cc *CameraImpl) SetIsObject3D(v bool) {
	cc.Set("isObject3D", v)
}
func (cc *CameraImpl) Layers() *Layers {
	return &Layers{Value: cc.Get("layers")}
}
func (cc *CameraImpl) SetLayers(v *Layers) {
	cc.Set("layers", v.Value)
}
func (cc *CameraImpl) Matrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrix")}
}
func (cc *CameraImpl) SetMatrix(v *Matrix4) {
	cc.Set("matrix", v.Value)
}
func (cc *CameraImpl) MatrixAutoUpdate() bool {
	return cc.Get("matrixAutoUpdate").Bool()
}
func (cc *CameraImpl) SetMatrixAutoUpdate(v bool) {
	cc.Set("matrixAutoUpdate", v)
}
func (cc *CameraImpl) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorld")}
}
func (cc *CameraImpl) SetMatrixWorld(v *Matrix4) {
	cc.Set("matrixWorld", v.Value)
}
func (cc *CameraImpl) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorldInverse")}
}
func (cc *CameraImpl) SetMatrixWorldInverse(v *Matrix4) {
	cc.Set("matrixWorldInverse", v.Value)
}
func (cc *CameraImpl) MatrixWorldNeedsUpdate() bool {
	return cc.Get("matrixWorldNeedsUpdate").Bool()
}
func (cc *CameraImpl) SetMatrixWorldNeedsUpdate(v bool) {
	cc.Set("matrixWorldNeedsUpdate", v)
}
func (cc *CameraImpl) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("modelViewMatrix")}
}
func (cc *CameraImpl) SetModelViewMatrix(v *Matrix4) {
	cc.Set("modelViewMatrix", v.Value)
}
func (cc *CameraImpl) Name() string {
	return cc.Get("name").String()
}
func (cc *CameraImpl) SetName(v string) {
	cc.Set("name", v)
}
func (cc *CameraImpl) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: cc.Get("normalMatrix")}
}
func (cc *CameraImpl) SetNormalMatrix(v *Matrix3) {
	cc.Set("normalMatrix", v.Value)
}
func (cc *CameraImpl) OnAfterRender() js.Value {
	return cc.Get("onAfterRender")
}
func (cc *CameraImpl) SetOnAfterRender(v js.Value) {
	cc.Set("onAfterRender", v)
}
func (cc *CameraImpl) OnBeforeRender() js.Value {
	return cc.Get("onBeforeRender")
}
func (cc *CameraImpl) SetOnBeforeRender(v js.Value) {
	cc.Set("onBeforeRender", v)
}
func (cc *CameraImpl) Parent() *Object3D {
	return &Object3D{Value: cc.Get("parent")}
}
func (cc *CameraImpl) SetParent(v *Object3D) {
	cc.Set("parent", v.Value)
}
func (cc *CameraImpl) Position() *Vector3 {
	return &Vector3{Value: cc.Get("position")}
}
func (cc *CameraImpl) SetPosition(v *Vector3) {
	cc.Set("position", v.Value)
}
func (cc *CameraImpl) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("projectionMatrix")}
}
func (cc *CameraImpl) SetProjectionMatrix(v *Matrix4) {
	cc.Set("projectionMatrix", v.Value)
}
func (cc *CameraImpl) Quaternion() *Quaternion {
	return &Quaternion{Value: cc.Get("quaternion")}
}
func (cc *CameraImpl) SetQuaternion(v *Quaternion) {
	cc.Set("quaternion", v.Value)
}
func (cc *CameraImpl) ReceiveShadow() bool {
	return cc.Get("receiveShadow").Bool()
}
func (cc *CameraImpl) SetReceiveShadow(v bool) {
	cc.Set("receiveShadow", v)
}
func (cc *CameraImpl) RenderOrder() float64 {
	return cc.Get("renderOrder").Float()
}
func (cc *CameraImpl) SetRenderOrder(v float64) {
	cc.Set("renderOrder", v)
}
func (cc *CameraImpl) Rotation() *Euler {
	return &Euler{Value: cc.Get("rotation")}
}
func (cc *CameraImpl) SetRotation(v *Euler) {
	cc.Set("rotation", v.Value)
}
func (cc *CameraImpl) Scale() *Vector3 {
	return &Vector3{Value: cc.Get("scale")}
}
func (cc *CameraImpl) SetScale(v *Vector3) {
	cc.Set("scale", v.Value)
}
func (cc *CameraImpl) Type() string {
	return cc.Get("type").String()
}
func (cc *CameraImpl) SetType(v string) {
	cc.Set("type", v)
}
func (cc *CameraImpl) Up() *Vector3 {
	return &Vector3{Value: cc.Get("up")}
}
func (cc *CameraImpl) SetUp(v *Vector3) {
	cc.Set("up", v.Value)
}
func (cc *CameraImpl) UserData() js.Value {
	return cc.Get("userData")
}
func (cc *CameraImpl) SetUserData(v js.Value) {
	cc.Set("userData", v)
}
func (cc *CameraImpl) Uuid() string {
	return cc.Get("uuid").String()
}
func (cc *CameraImpl) SetUuid(v string) {
	cc.Set("uuid", v)
}
func (cc *CameraImpl) Visible() bool {
	return cc.Get("visible").Bool()
}
func (cc *CameraImpl) SetVisible(v bool) {
	cc.Set("visible", v)
}
func (cc *CameraImpl) DefaultMatrixAutoUpdate() bool {
	return cc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (cc *CameraImpl) SetDefaultMatrixAutoUpdate(v bool) {
	cc.Set("DefaultMatrixAutoUpdate", v)
}
func (cc *CameraImpl) DefaultUp() *Vector3 {
	return &Vector3{Value: cc.Get("DefaultUp")}
}
func (cc *CameraImpl) SetDefaultUp(v *Vector3) {
	cc.Set("DefaultUp", v.Value)
}
func (cc *CameraImpl) Add(object js.Value) Camera {
	return &CameraImpl{Value: cc.Call("add", object)}
}
func (cc *CameraImpl) AddEventListener(typ string, listener js.Value) {
	cc.Call("addEventListener", typ, listener)
}
func (cc *CameraImpl) ApplyMatrix(matrix *Matrix4) {
	cc.Call("applyMatrix", matrix)
}
func (cc *CameraImpl) ApplyQuaternion(quaternion *Quaternion) Camera {
	return &CameraImpl{Value: cc.Call("applyQuaternion", quaternion)}
}
func (cc *CameraImpl) Clone(recursive bool) Camera {
	return &CameraImpl{Value: cc.Call("clone", recursive)}
}
func (cc *CameraImpl) Copy(source Camera, recursive bool) Camera {
	return &CameraImpl{Value: cc.Call("copy", source.JSValue(), recursive)}
}
func (cc *CameraImpl) DispatchEvent(event js.Value) {
	cc.Call("dispatchEvent", event)
}
func (cc *CameraImpl) GetObjectById(id int) *Object3D {
	return &Object3D{Value: cc.Call("getObjectById", id)}
}
func (cc *CameraImpl) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByName", name)}
}
func (cc *CameraImpl) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByProperty", name, value)}
}
func (cc *CameraImpl) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldDirection", target)}
}
func (cc *CameraImpl) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldPosition", target)}
}
func (cc *CameraImpl) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: cc.Call("getWorldQuaternion", target)}
}
func (cc *CameraImpl) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldScale", target)}
}
func (cc *CameraImpl) HasEventListener(typ string, listener js.Value) bool {
	return cc.Call("hasEventListener", typ, listener).Bool()
}
func (cc *CameraImpl) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("localToWorld", vector)}
}
func (cc *CameraImpl) LookAt(vector *Vector3, y float64, z float64) {
	cc.Call("lookAt", vector, y, z)
}
func (cc *CameraImpl) Raycast(raycaster *Raycaster, intersects js.Value) {
	cc.Call("raycast", raycaster, intersects)
}
func (cc *CameraImpl) Remove(object js.Value) Camera {
	return &CameraImpl{Value: cc.Call("remove", object)}
}
func (cc *CameraImpl) RemoveEventListener(typ string, listener js.Value) {
	cc.Call("removeEventListener", typ, listener)
}
func (cc *CameraImpl) RotateOnAxis(axis *Vector3, angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateOnAxis", axis, angle)}
}
func (cc *CameraImpl) RotateOnWorldAxis(axis *Vector3, angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateOnWorldAxis", axis, angle)}
}
func (cc *CameraImpl) RotateX(angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateX", angle)}
}
func (cc *CameraImpl) RotateY(angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateY", angle)}
}
func (cc *CameraImpl) RotateZ(angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateZ", angle)}
}
func (cc *CameraImpl) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	cc.Call("setRotationFromAxisAngle", axis, angle)
}
func (cc *CameraImpl) SetRotationFromEuler(euler *Euler) {
	cc.Call("setRotationFromEuler", euler)
}
func (cc *CameraImpl) SetRotationFromMatrix(m *Matrix4) {
	cc.Call("setRotationFromMatrix", m)
}
func (cc *CameraImpl) SetRotationFromQuaternion(q *Quaternion) {
	cc.Call("setRotationFromQuaternion", q)
}
func (cc *CameraImpl) ToJSON(meta js.Value) js.Value {
	return cc.Call("toJSON", meta)
}
func (cc *CameraImpl) TranslateOnAxis(axis *Vector3, distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateOnAxis", axis, distance)}
}
func (cc *CameraImpl) TranslateX(distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateX", distance)}
}
func (cc *CameraImpl) TranslateY(distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateY", distance)}
}
func (cc *CameraImpl) TranslateZ(distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateZ", distance)}
}
func (cc *CameraImpl) Traverse(callback js.Value) {
	cc.Call("traverse", callback)
}
func (cc *CameraImpl) TraverseAncestors(callback js.Value) {
	cc.Call("traverseAncestors", callback)
}
func (cc *CameraImpl) TraverseVisible(callback js.Value) {
	cc.Call("traverseVisible", callback)
}
func (cc *CameraImpl) UpdateMatrix() {
	cc.Call("updateMatrix")
}
func (cc *CameraImpl) UpdateMatrixWorld(force bool) {
	cc.Call("updateMatrixWorld", force)
}
func (cc *CameraImpl) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	cc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (cc *CameraImpl) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("worldToLocal", vector)}
}

// Code generated from "extras/objects/ImmediateRenderObject.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type ImmediateRenderObject struct {
	js.Value
}

func NewImmediateRenderObject(material *Material) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: get("ImmediateRenderObject").New(material)}
}
func (iro *ImmediateRenderObject) CastShadow() bool {
	return iro.Get("castShadow").Bool()
}
func (iro *ImmediateRenderObject) SetCastShadow(v bool) {
	iro.Set("castShadow", v)
}
func (iro *ImmediateRenderObject) Children() js.Value {
	return iro.Get("children")
}
func (iro *ImmediateRenderObject) SetChildren(v js.Value) {
	iro.Set("children", v)
}
func (iro *ImmediateRenderObject) FrustumCulled() bool {
	return iro.Get("frustumCulled").Bool()
}
func (iro *ImmediateRenderObject) SetFrustumCulled(v bool) {
	iro.Set("frustumCulled", v)
}
func (iro *ImmediateRenderObject) Id() int {
	return iro.Get("id").Int()
}
func (iro *ImmediateRenderObject) SetId(v int) {
	iro.Set("id", v)
}
func (iro *ImmediateRenderObject) IsObject3D() bool {
	return iro.Get("isObject3D").Bool()
}
func (iro *ImmediateRenderObject) SetIsObject3D(v bool) {
	iro.Set("isObject3D", v)
}
func (iro *ImmediateRenderObject) Layers() *Layers {
	return &Layers{Value: iro.Get("layers")}
}
func (iro *ImmediateRenderObject) SetLayers(v *Layers) {
	iro.Set("layers", v)
}
func (iro *ImmediateRenderObject) Material() *Material {
	return &Material{Value: iro.Get("material")}
}
func (iro *ImmediateRenderObject) SetMaterial(v *Material) {
	iro.Set("material", v)
}
func (iro *ImmediateRenderObject) Matrix() *Matrix4 {
	return &Matrix4{Value: iro.Get("matrix")}
}
func (iro *ImmediateRenderObject) SetMatrix(v *Matrix4) {
	iro.Set("matrix", v)
}
func (iro *ImmediateRenderObject) MatrixAutoUpdate() bool {
	return iro.Get("matrixAutoUpdate").Bool()
}
func (iro *ImmediateRenderObject) SetMatrixAutoUpdate(v bool) {
	iro.Set("matrixAutoUpdate", v)
}
func (iro *ImmediateRenderObject) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: iro.Get("matrixWorld")}
}
func (iro *ImmediateRenderObject) SetMatrixWorld(v *Matrix4) {
	iro.Set("matrixWorld", v)
}
func (iro *ImmediateRenderObject) MatrixWorldNeedsUpdate() bool {
	return iro.Get("matrixWorldNeedsUpdate").Bool()
}
func (iro *ImmediateRenderObject) SetMatrixWorldNeedsUpdate(v bool) {
	iro.Set("matrixWorldNeedsUpdate", v)
}
func (iro *ImmediateRenderObject) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: iro.Get("modelViewMatrix")}
}
func (iro *ImmediateRenderObject) SetModelViewMatrix(v *Matrix4) {
	iro.Set("modelViewMatrix", v)
}
func (iro *ImmediateRenderObject) Name() string {
	return iro.Get("name").String()
}
func (iro *ImmediateRenderObject) SetName(v string) {
	iro.Set("name", v)
}
func (iro *ImmediateRenderObject) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: iro.Get("normalMatrix")}
}
func (iro *ImmediateRenderObject) SetNormalMatrix(v *Matrix3) {
	iro.Set("normalMatrix", v)
}
func (iro *ImmediateRenderObject) OnAfterRender() js.Value {
	return iro.Get("onAfterRender")
}
func (iro *ImmediateRenderObject) SetOnAfterRender(v js.Value) {
	iro.Set("onAfterRender", v)
}
func (iro *ImmediateRenderObject) OnBeforeRender() js.Value {
	return iro.Get("onBeforeRender")
}
func (iro *ImmediateRenderObject) SetOnBeforeRender(v js.Value) {
	iro.Set("onBeforeRender", v)
}
func (iro *ImmediateRenderObject) Parent() *Object3D {
	return &Object3D{Value: iro.Get("parent")}
}
func (iro *ImmediateRenderObject) SetParent(v *Object3D) {
	iro.Set("parent", v)
}
func (iro *ImmediateRenderObject) Position() *Vector3 {
	return &Vector3{Value: iro.Get("position")}
}
func (iro *ImmediateRenderObject) SetPosition(v *Vector3) {
	iro.Set("position", v)
}
func (iro *ImmediateRenderObject) Quaternion() *Quaternion {
	return &Quaternion{Value: iro.Get("quaternion")}
}
func (iro *ImmediateRenderObject) SetQuaternion(v *Quaternion) {
	iro.Set("quaternion", v)
}
func (iro *ImmediateRenderObject) ReceiveShadow() bool {
	return iro.Get("receiveShadow").Bool()
}
func (iro *ImmediateRenderObject) SetReceiveShadow(v bool) {
	iro.Set("receiveShadow", v)
}
func (iro *ImmediateRenderObject) RenderOrder() float64 {
	return iro.Get("renderOrder").Float()
}
func (iro *ImmediateRenderObject) SetRenderOrder(v float64) {
	iro.Set("renderOrder", v)
}
func (iro *ImmediateRenderObject) Rotation() *Euler {
	return &Euler{Value: iro.Get("rotation")}
}
func (iro *ImmediateRenderObject) SetRotation(v *Euler) {
	iro.Set("rotation", v)
}
func (iro *ImmediateRenderObject) Scale() *Vector3 {
	return &Vector3{Value: iro.Get("scale")}
}
func (iro *ImmediateRenderObject) SetScale(v *Vector3) {
	iro.Set("scale", v)
}
func (iro *ImmediateRenderObject) Type() string {
	return iro.Get("type").String()
}
func (iro *ImmediateRenderObject) SetType(v string) {
	iro.Set("type", v)
}
func (iro *ImmediateRenderObject) Up() *Vector3 {
	return &Vector3{Value: iro.Get("up")}
}
func (iro *ImmediateRenderObject) SetUp(v *Vector3) {
	iro.Set("up", v)
}
func (iro *ImmediateRenderObject) UserData() js.Value {
	return iro.Get("userData")
}
func (iro *ImmediateRenderObject) SetUserData(v js.Value) {
	iro.Set("userData", v)
}
func (iro *ImmediateRenderObject) Uuid() string {
	return iro.Get("uuid").String()
}
func (iro *ImmediateRenderObject) SetUuid(v string) {
	iro.Set("uuid", v)
}
func (iro *ImmediateRenderObject) Visible() bool {
	return iro.Get("visible").Bool()
}
func (iro *ImmediateRenderObject) SetVisible(v bool) {
	iro.Set("visible", v)
}
func (iro *ImmediateRenderObject) DefaultMatrixAutoUpdate() bool {
	return iro.Get("DefaultMatrixAutoUpdate").Bool()
}
func (iro *ImmediateRenderObject) SetDefaultMatrixAutoUpdate(v bool) {
	iro.Set("DefaultMatrixAutoUpdate", v)
}
func (iro *ImmediateRenderObject) DefaultUp() *Vector3 {
	return &Vector3{Value: iro.Get("DefaultUp")}
}
func (iro *ImmediateRenderObject) SetDefaultUp(v *Vector3) {
	iro.Set("DefaultUp", v)
}
func (iro *ImmediateRenderObject) Add(object js.Value) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("add", object)}
}
func (iro *ImmediateRenderObject) AddEventListener(typ string, listener js.Value) {
	iro.Call("addEventListener", typ, listener)
}
func (iro *ImmediateRenderObject) ApplyMatrix(matrix *Matrix4) {
	iro.Call("applyMatrix", matrix)
}
func (iro *ImmediateRenderObject) ApplyQuaternion(quaternion *Quaternion) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("applyQuaternion", quaternion)}
}
func (iro *ImmediateRenderObject) Clone(recursive bool) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("clone", recursive)}
}
func (iro *ImmediateRenderObject) Copy(source *Object3D, recursive bool) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("copy", source, recursive)}
}
func (iro *ImmediateRenderObject) DispatchEvent(event js.Value) {
	iro.Call("dispatchEvent", event)
}
func (iro *ImmediateRenderObject) GetObjectById(id int) *Object3D {
	return &Object3D{Value: iro.Call("getObjectById", id)}
}
func (iro *ImmediateRenderObject) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: iro.Call("getObjectByName", name)}
}
func (iro *ImmediateRenderObject) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: iro.Call("getObjectByProperty", name, value)}
}
func (iro *ImmediateRenderObject) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: iro.Call("getWorldDirection", target)}
}
func (iro *ImmediateRenderObject) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: iro.Call("getWorldPosition", target)}
}
func (iro *ImmediateRenderObject) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: iro.Call("getWorldQuaternion", target)}
}
func (iro *ImmediateRenderObject) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: iro.Call("getWorldScale", target)}
}
func (iro *ImmediateRenderObject) HasEventListener(typ string, listener js.Value) bool {
	return iro.Call("hasEventListener", typ, listener).Bool()
}
func (iro *ImmediateRenderObject) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: iro.Call("localToWorld", vector)}
}
func (iro *ImmediateRenderObject) LookAt(vector *Vector3, y float64, z float64) {
	iro.Call("lookAt", vector, y, z)
}
func (iro *ImmediateRenderObject) Raycast(raycaster *Raycaster, intersects js.Value) {
	iro.Call("raycast", raycaster, intersects)
}
func (iro *ImmediateRenderObject) Remove(object js.Value) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("remove", object)}
}
func (iro *ImmediateRenderObject) RemoveEventListener(typ string, listener js.Value) {
	iro.Call("removeEventListener", typ, listener)
}
func (iro *ImmediateRenderObject) Render(renderCallback js.Value) {
	iro.Call("render", renderCallback)
}
func (iro *ImmediateRenderObject) RotateOnAxis(axis *Vector3, angle float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("rotateOnAxis", axis, angle)}
}
func (iro *ImmediateRenderObject) RotateOnWorldAxis(axis *Vector3, angle float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("rotateOnWorldAxis", axis, angle)}
}
func (iro *ImmediateRenderObject) RotateX(angle float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("rotateX", angle)}
}
func (iro *ImmediateRenderObject) RotateY(angle float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("rotateY", angle)}
}
func (iro *ImmediateRenderObject) RotateZ(angle float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("rotateZ", angle)}
}
func (iro *ImmediateRenderObject) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	iro.Call("setRotationFromAxisAngle", axis, angle)
}
func (iro *ImmediateRenderObject) SetRotationFromEuler(euler *Euler) {
	iro.Call("setRotationFromEuler", euler)
}
func (iro *ImmediateRenderObject) SetRotationFromMatrix(m *Matrix4) {
	iro.Call("setRotationFromMatrix", m)
}
func (iro *ImmediateRenderObject) SetRotationFromQuaternion(q *Quaternion) {
	iro.Call("setRotationFromQuaternion", q)
}
func (iro *ImmediateRenderObject) ToJSON(meta js.Value) js.Value {
	return iro.Call("toJSON", meta)
}
func (iro *ImmediateRenderObject) TranslateOnAxis(axis *Vector3, distance float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("translateOnAxis", axis, distance)}
}
func (iro *ImmediateRenderObject) TranslateX(distance float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("translateX", distance)}
}
func (iro *ImmediateRenderObject) TranslateY(distance float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("translateY", distance)}
}
func (iro *ImmediateRenderObject) TranslateZ(distance float64) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: iro.Call("translateZ", distance)}
}
func (iro *ImmediateRenderObject) Traverse(callback js.Value) {
	iro.Call("traverse", callback)
}
func (iro *ImmediateRenderObject) TraverseAncestors(callback js.Value) {
	iro.Call("traverseAncestors", callback)
}
func (iro *ImmediateRenderObject) TraverseVisible(callback js.Value) {
	iro.Call("traverseVisible", callback)
}
func (iro *ImmediateRenderObject) UpdateMatrix() {
	iro.Call("updateMatrix")
}
func (iro *ImmediateRenderObject) UpdateMatrixWorld(force bool) {
	iro.Call("updateMatrixWorld", force)
}
func (iro *ImmediateRenderObject) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	iro.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (iro *ImmediateRenderObject) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: iro.Call("worldToLocal", vector)}
}

// Code generated from "objects/Sprite.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Sprite extend: [Object3D]
type Sprite struct {
	js.Value
}

func NewSprite(material *SpriteMaterial) *Sprite {
	return &Sprite{Value: get("Sprite").New(material)}
}
func (ss *Sprite) JSValue() js.Value {
	return ss.Value
}
func (ss *Sprite) CastShadow() bool {
	return ss.Get("castShadow").Bool()
}
func (ss *Sprite) SetCastShadow(v bool) {
	ss.Set("castShadow", v)
}
func (ss *Sprite) Center() *Vector2 {
	return &Vector2{Value: ss.Get("center")}
}
func (ss *Sprite) SetCenter(v *Vector2) {
	ss.Set("center", v.Value)
}
func (ss *Sprite) Children() js.Value {
	return ss.Get("children")
}
func (ss *Sprite) SetChildren(v js.Value) {
	ss.Set("children", v)
}
func (ss *Sprite) FrustumCulled() bool {
	return ss.Get("frustumCulled").Bool()
}
func (ss *Sprite) SetFrustumCulled(v bool) {
	ss.Set("frustumCulled", v)
}
func (ss *Sprite) Geometry() *BufferGeometry {
	return &BufferGeometry{Value: ss.Get("geometry")}
}
func (ss *Sprite) SetGeometry(v *BufferGeometry) {
	ss.Set("geometry", v.Value)
}
func (ss *Sprite) Id() int {
	return ss.Get("id").Int()
}
func (ss *Sprite) SetId(v int) {
	ss.Set("id", v)
}
func (ss *Sprite) IsObject3D() bool {
	return ss.Get("isObject3D").Bool()
}
func (ss *Sprite) SetIsObject3D(v bool) {
	ss.Set("isObject3D", v)
}
func (ss *Sprite) IsSprite() bool {
	return ss.Get("isSprite").Bool()
}
func (ss *Sprite) SetIsSprite(v bool) {
	ss.Set("isSprite", v)
}
func (ss *Sprite) Layers() *Layers {
	return &Layers{Value: ss.Get("layers")}
}
func (ss *Sprite) SetLayers(v *Layers) {
	ss.Set("layers", v.Value)
}
func (ss *Sprite) Material() *SpriteMaterial {
	return &SpriteMaterial{Value: ss.Get("material")}
}
func (ss *Sprite) SetMaterial(v *SpriteMaterial) {
	ss.Set("material", v.Value)
}
func (ss *Sprite) Matrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrix")}
}
func (ss *Sprite) SetMatrix(v *Matrix4) {
	ss.Set("matrix", v.Value)
}
func (ss *Sprite) MatrixAutoUpdate() bool {
	return ss.Get("matrixAutoUpdate").Bool()
}
func (ss *Sprite) SetMatrixAutoUpdate(v bool) {
	ss.Set("matrixAutoUpdate", v)
}
func (ss *Sprite) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrixWorld")}
}
func (ss *Sprite) SetMatrixWorld(v *Matrix4) {
	ss.Set("matrixWorld", v.Value)
}
func (ss *Sprite) MatrixWorldNeedsUpdate() bool {
	return ss.Get("matrixWorldNeedsUpdate").Bool()
}
func (ss *Sprite) SetMatrixWorldNeedsUpdate(v bool) {
	ss.Set("matrixWorldNeedsUpdate", v)
}
func (ss *Sprite) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("modelViewMatrix")}
}
func (ss *Sprite) SetModelViewMatrix(v *Matrix4) {
	ss.Set("modelViewMatrix", v.Value)
}
func (ss *Sprite) Name() string {
	return ss.Get("name").String()
}
func (ss *Sprite) SetName(v string) {
	ss.Set("name", v)
}
func (ss *Sprite) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ss.Get("normalMatrix")}
}
func (ss *Sprite) SetNormalMatrix(v *Matrix3) {
	ss.Set("normalMatrix", v.Value)
}
func (ss *Sprite) OnAfterRender() js.Value {
	return ss.Get("onAfterRender")
}
func (ss *Sprite) SetOnAfterRender(v js.Value) {
	ss.Set("onAfterRender", v)
}
func (ss *Sprite) OnBeforeRender() js.Value {
	return ss.Get("onBeforeRender")
}
func (ss *Sprite) SetOnBeforeRender(v js.Value) {
	ss.Set("onBeforeRender", v)
}
func (ss *Sprite) Parent() *Object3D {
	return &Object3D{Value: ss.Get("parent")}
}
func (ss *Sprite) SetParent(v *Object3D) {
	ss.Set("parent", v.Value)
}
func (ss *Sprite) Position() *Vector3 {
	return &Vector3{Value: ss.Get("position")}
}
func (ss *Sprite) SetPosition(v *Vector3) {
	ss.Set("position", v.Value)
}
func (ss *Sprite) Quaternion() *Quaternion {
	return &Quaternion{Value: ss.Get("quaternion")}
}
func (ss *Sprite) SetQuaternion(v *Quaternion) {
	ss.Set("quaternion", v.Value)
}
func (ss *Sprite) ReceiveShadow() bool {
	return ss.Get("receiveShadow").Bool()
}
func (ss *Sprite) SetReceiveShadow(v bool) {
	ss.Set("receiveShadow", v)
}
func (ss *Sprite) RenderOrder() float64 {
	return ss.Get("renderOrder").Float()
}
func (ss *Sprite) SetRenderOrder(v float64) {
	ss.Set("renderOrder", v)
}
func (ss *Sprite) Rotation() *Euler {
	return &Euler{Value: ss.Get("rotation")}
}
func (ss *Sprite) SetRotation(v *Euler) {
	ss.Set("rotation", v.Value)
}
func (ss *Sprite) Scale() *Vector3 {
	return &Vector3{Value: ss.Get("scale")}
}
func (ss *Sprite) SetScale(v *Vector3) {
	ss.Set("scale", v.Value)
}
func (ss *Sprite) Type() string {
	return ss.Get("type").String()
}
func (ss *Sprite) SetType(v string) {
	ss.Set("type", v)
}
func (ss *Sprite) Up() *Vector3 {
	return &Vector3{Value: ss.Get("up")}
}
func (ss *Sprite) SetUp(v *Vector3) {
	ss.Set("up", v.Value)
}
func (ss *Sprite) UserData() js.Value {
	return ss.Get("userData")
}
func (ss *Sprite) SetUserData(v js.Value) {
	ss.Set("userData", v)
}
func (ss *Sprite) Uuid() string {
	return ss.Get("uuid").String()
}
func (ss *Sprite) SetUuid(v string) {
	ss.Set("uuid", v)
}
func (ss *Sprite) Visible() bool {
	return ss.Get("visible").Bool()
}
func (ss *Sprite) SetVisible(v bool) {
	ss.Set("visible", v)
}
func (ss *Sprite) DefaultMatrixAutoUpdate() bool {
	return ss.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ss *Sprite) SetDefaultMatrixAutoUpdate(v bool) {
	ss.Set("DefaultMatrixAutoUpdate", v)
}
func (ss *Sprite) DefaultUp() *Vector3 {
	return &Vector3{Value: ss.Get("DefaultUp")}
}
func (ss *Sprite) SetDefaultUp(v *Vector3) {
	ss.Set("DefaultUp", v.Value)
}
func (ss *Sprite) Add(object js.Value) *Sprite {
	return &Sprite{Value: ss.Call("add", object)}
}
func (ss *Sprite) AddEventListener(typ string, listener js.Value) {
	ss.Call("addEventListener", typ, listener)
}
func (ss *Sprite) ApplyMatrix(matrix *Matrix4) {
	ss.Call("applyMatrix", matrix)
}
func (ss *Sprite) ApplyQuaternion(quaternion *Quaternion) *Sprite {
	return &Sprite{Value: ss.Call("applyQuaternion", quaternion)}
}
func (ss *Sprite) Clone(recursive bool) *Sprite {
	return &Sprite{Value: ss.Call("clone", recursive)}
}
func (ss *Sprite) Copy(source *Sprite) *Sprite {
	return &Sprite{Value: ss.Call("copy", source)}
}
func (ss *Sprite) DispatchEvent(event js.Value) {
	ss.Call("dispatchEvent", event)
}
func (ss *Sprite) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ss.Call("getObjectById", id)}
}
func (ss *Sprite) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByName", name)}
}
func (ss *Sprite) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByProperty", name, value)}
}
func (ss *Sprite) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldDirection", target)}
}
func (ss *Sprite) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldPosition", target)}
}
func (ss *Sprite) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ss.Call("getWorldQuaternion", target)}
}
func (ss *Sprite) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldScale", target)}
}
func (ss *Sprite) HasEventListener(typ string, listener js.Value) bool {
	return ss.Call("hasEventListener", typ, listener).Bool()
}
func (ss *Sprite) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("localToWorld", vector)}
}
func (ss *Sprite) LookAt(vector *Vector3, y float64, z float64) {
	ss.Call("lookAt", vector, y, z)
}
func (ss *Sprite) Raycast(raycaster *Raycaster, intersects js.Value) {
	ss.Call("raycast", raycaster, intersects)
}
func (ss *Sprite) Remove(object js.Value) *Sprite {
	return &Sprite{Value: ss.Call("remove", object)}
}
func (ss *Sprite) RemoveEventListener(typ string, listener js.Value) {
	ss.Call("removeEventListener", typ, listener)
}
func (ss *Sprite) RotateOnAxis(axis *Vector3, angle float64) *Sprite {
	return &Sprite{Value: ss.Call("rotateOnAxis", axis, angle)}
}
func (ss *Sprite) RotateOnWorldAxis(axis *Vector3, angle float64) *Sprite {
	return &Sprite{Value: ss.Call("rotateOnWorldAxis", axis, angle)}
}
func (ss *Sprite) RotateX(angle float64) *Sprite {
	return &Sprite{Value: ss.Call("rotateX", angle)}
}
func (ss *Sprite) RotateY(angle float64) *Sprite {
	return &Sprite{Value: ss.Call("rotateY", angle)}
}
func (ss *Sprite) RotateZ(angle float64) *Sprite {
	return &Sprite{Value: ss.Call("rotateZ", angle)}
}
func (ss *Sprite) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ss.Call("setRotationFromAxisAngle", axis, angle)
}
func (ss *Sprite) SetRotationFromEuler(euler *Euler) {
	ss.Call("setRotationFromEuler", euler)
}
func (ss *Sprite) SetRotationFromMatrix(m *Matrix4) {
	ss.Call("setRotationFromMatrix", m)
}
func (ss *Sprite) SetRotationFromQuaternion(q *Quaternion) {
	ss.Call("setRotationFromQuaternion", q)
}
func (ss *Sprite) ToJSON(meta js.Value) js.Value {
	return ss.Call("toJSON", meta)
}
func (ss *Sprite) TranslateOnAxis(axis *Vector3, distance float64) *Sprite {
	return &Sprite{Value: ss.Call("translateOnAxis", axis, distance)}
}
func (ss *Sprite) TranslateX(distance float64) *Sprite {
	return &Sprite{Value: ss.Call("translateX", distance)}
}
func (ss *Sprite) TranslateY(distance float64) *Sprite {
	return &Sprite{Value: ss.Call("translateY", distance)}
}
func (ss *Sprite) TranslateZ(distance float64) *Sprite {
	return &Sprite{Value: ss.Call("translateZ", distance)}
}
func (ss *Sprite) Traverse(callback js.Value) {
	ss.Call("traverse", callback)
}
func (ss *Sprite) TraverseAncestors(callback js.Value) {
	ss.Call("traverseAncestors", callback)
}
func (ss *Sprite) TraverseVisible(callback js.Value) {
	ss.Call("traverseVisible", callback)
}
func (ss *Sprite) UpdateMatrix() {
	ss.Call("updateMatrix")
}
func (ss *Sprite) UpdateMatrixWorld(force bool) {
	ss.Call("updateMatrixWorld", force)
}
func (ss *Sprite) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ss.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ss *Sprite) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("worldToLocal", vector)}
}

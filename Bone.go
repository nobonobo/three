// Code generated from "objects/Bone.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Bone struct {
	js.Value
}

func NewBone() *Bone {
	return &Bone{Value: get("Bone").New()}
}
func (bb *Bone) CastShadow() bool {
	return bb.Get("castShadow").Bool()
}
func (bb *Bone) SetCastShadow(v bool) {
	bb.Set("castShadow", v)
}
func (bb *Bone) Children() js.Value {
	return bb.Get("children")
}
func (bb *Bone) SetChildren(v js.Value) {
	bb.Set("children", v)
}
func (bb *Bone) FrustumCulled() bool {
	return bb.Get("frustumCulled").Bool()
}
func (bb *Bone) SetFrustumCulled(v bool) {
	bb.Set("frustumCulled", v)
}
func (bb *Bone) Id() int {
	return bb.Get("id").Int()
}
func (bb *Bone) SetId(v int) {
	bb.Set("id", v)
}
func (bb *Bone) IsBone() bool {
	return bb.Get("isBone").Bool()
}
func (bb *Bone) SetIsBone(v bool) {
	bb.Set("isBone", v)
}
func (bb *Bone) IsObject3D() bool {
	return bb.Get("isObject3D").Bool()
}
func (bb *Bone) SetIsObject3D(v bool) {
	bb.Set("isObject3D", v)
}
func (bb *Bone) Layers() *Layers {
	return &Layers{Value: bb.Get("layers")}
}
func (bb *Bone) SetLayers(v *Layers) {
	bb.Set("layers", v)
}
func (bb *Bone) Matrix() *Matrix4 {
	return &Matrix4{Value: bb.Get("matrix")}
}
func (bb *Bone) SetMatrix(v *Matrix4) {
	bb.Set("matrix", v)
}
func (bb *Bone) MatrixAutoUpdate() bool {
	return bb.Get("matrixAutoUpdate").Bool()
}
func (bb *Bone) SetMatrixAutoUpdate(v bool) {
	bb.Set("matrixAutoUpdate", v)
}
func (bb *Bone) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: bb.Get("matrixWorld")}
}
func (bb *Bone) SetMatrixWorld(v *Matrix4) {
	bb.Set("matrixWorld", v)
}
func (bb *Bone) MatrixWorldNeedsUpdate() bool {
	return bb.Get("matrixWorldNeedsUpdate").Bool()
}
func (bb *Bone) SetMatrixWorldNeedsUpdate(v bool) {
	bb.Set("matrixWorldNeedsUpdate", v)
}
func (bb *Bone) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: bb.Get("modelViewMatrix")}
}
func (bb *Bone) SetModelViewMatrix(v *Matrix4) {
	bb.Set("modelViewMatrix", v)
}
func (bb *Bone) Name() string {
	return bb.Get("name").String()
}
func (bb *Bone) SetName(v string) {
	bb.Set("name", v)
}
func (bb *Bone) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: bb.Get("normalMatrix")}
}
func (bb *Bone) SetNormalMatrix(v *Matrix3) {
	bb.Set("normalMatrix", v)
}
func (bb *Bone) OnAfterRender() js.Value {
	return bb.Get("onAfterRender")
}
func (bb *Bone) SetOnAfterRender(v js.Value) {
	bb.Set("onAfterRender", v)
}
func (bb *Bone) OnBeforeRender() js.Value {
	return bb.Get("onBeforeRender")
}
func (bb *Bone) SetOnBeforeRender(v js.Value) {
	bb.Set("onBeforeRender", v)
}
func (bb *Bone) Parent() *Object3D {
	return &Object3D{Value: bb.Get("parent")}
}
func (bb *Bone) SetParent(v *Object3D) {
	bb.Set("parent", v)
}
func (bb *Bone) Position() *Vector3 {
	return &Vector3{Value: bb.Get("position")}
}
func (bb *Bone) SetPosition(v *Vector3) {
	bb.Set("position", v)
}
func (bb *Bone) Quaternion() *Quaternion {
	return &Quaternion{Value: bb.Get("quaternion")}
}
func (bb *Bone) SetQuaternion(v *Quaternion) {
	bb.Set("quaternion", v)
}
func (bb *Bone) ReceiveShadow() bool {
	return bb.Get("receiveShadow").Bool()
}
func (bb *Bone) SetReceiveShadow(v bool) {
	bb.Set("receiveShadow", v)
}
func (bb *Bone) RenderOrder() float64 {
	return bb.Get("renderOrder").Float()
}
func (bb *Bone) SetRenderOrder(v float64) {
	bb.Set("renderOrder", v)
}
func (bb *Bone) Rotation() *Euler {
	return &Euler{Value: bb.Get("rotation")}
}
func (bb *Bone) SetRotation(v *Euler) {
	bb.Set("rotation", v)
}
func (bb *Bone) Scale() *Vector3 {
	return &Vector3{Value: bb.Get("scale")}
}
func (bb *Bone) SetScale(v *Vector3) {
	bb.Set("scale", v)
}
func (bb *Bone) Type() string {
	return bb.Get("type").String()
}
func (bb *Bone) SetType(v string) {
	bb.Set("type", v)
}
func (bb *Bone) Up() *Vector3 {
	return &Vector3{Value: bb.Get("up")}
}
func (bb *Bone) SetUp(v *Vector3) {
	bb.Set("up", v)
}
func (bb *Bone) UserData() js.Value {
	return bb.Get("userData")
}
func (bb *Bone) SetUserData(v js.Value) {
	bb.Set("userData", v)
}
func (bb *Bone) Uuid() string {
	return bb.Get("uuid").String()
}
func (bb *Bone) SetUuid(v string) {
	bb.Set("uuid", v)
}
func (bb *Bone) Visible() bool {
	return bb.Get("visible").Bool()
}
func (bb *Bone) SetVisible(v bool) {
	bb.Set("visible", v)
}
func (bb *Bone) DefaultMatrixAutoUpdate() bool {
	return bb.Get("DefaultMatrixAutoUpdate").Bool()
}
func (bb *Bone) SetDefaultMatrixAutoUpdate(v bool) {
	bb.Set("DefaultMatrixAutoUpdate", v)
}
func (bb *Bone) DefaultUp() *Vector3 {
	return &Vector3{Value: bb.Get("DefaultUp")}
}
func (bb *Bone) SetDefaultUp(v *Vector3) {
	bb.Set("DefaultUp", v)
}
func (bb *Bone) Add(object js.Value) *Bone {
	return &Bone{Value: bb.Call("add", object)}
}
func (bb *Bone) AddEventListener(typ string, listener js.Value) {
	bb.Call("addEventListener", typ, listener)
}
func (bb *Bone) ApplyMatrix(matrix *Matrix4) {
	bb.Call("applyMatrix", matrix)
}
func (bb *Bone) ApplyQuaternion(quaternion *Quaternion) *Bone {
	return &Bone{Value: bb.Call("applyQuaternion", quaternion)}
}
func (bb *Bone) Clone(recursive bool) *Bone {
	return &Bone{Value: bb.Call("clone", recursive)}
}
func (bb *Bone) Copy(source *Object3D, recursive bool) *Bone {
	return &Bone{Value: bb.Call("copy", source, recursive)}
}
func (bb *Bone) DispatchEvent(event js.Value) {
	bb.Call("dispatchEvent", event)
}
func (bb *Bone) GetObjectById(id int) *Object3D {
	return &Object3D{Value: bb.Call("getObjectById", id)}
}
func (bb *Bone) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: bb.Call("getObjectByName", name)}
}
func (bb *Bone) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: bb.Call("getObjectByProperty", name, value)}
}
func (bb *Bone) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("getWorldDirection", target)}
}
func (bb *Bone) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("getWorldPosition", target)}
}
func (bb *Bone) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: bb.Call("getWorldQuaternion", target)}
}
func (bb *Bone) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("getWorldScale", target)}
}
func (bb *Bone) HasEventListener(typ string, listener js.Value) bool {
	return bb.Call("hasEventListener", typ, listener).Bool()
}
func (bb *Bone) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("localToWorld", vector)}
}
func (bb *Bone) LookAt(vector *Vector3, y float64, z float64) {
	bb.Call("lookAt", vector, y, z)
}
func (bb *Bone) Raycast(raycaster *Raycaster, intersects js.Value) {
	bb.Call("raycast", raycaster, intersects)
}
func (bb *Bone) Remove(object js.Value) *Bone {
	return &Bone{Value: bb.Call("remove", object)}
}
func (bb *Bone) RemoveEventListener(typ string, listener js.Value) {
	bb.Call("removeEventListener", typ, listener)
}
func (bb *Bone) RotateOnAxis(axis *Vector3, angle float64) *Bone {
	return &Bone{Value: bb.Call("rotateOnAxis", axis, angle)}
}
func (bb *Bone) RotateOnWorldAxis(axis *Vector3, angle float64) *Bone {
	return &Bone{Value: bb.Call("rotateOnWorldAxis", axis, angle)}
}
func (bb *Bone) RotateX(angle float64) *Bone {
	return &Bone{Value: bb.Call("rotateX", angle)}
}
func (bb *Bone) RotateY(angle float64) *Bone {
	return &Bone{Value: bb.Call("rotateY", angle)}
}
func (bb *Bone) RotateZ(angle float64) *Bone {
	return &Bone{Value: bb.Call("rotateZ", angle)}
}
func (bb *Bone) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	bb.Call("setRotationFromAxisAngle", axis, angle)
}
func (bb *Bone) SetRotationFromEuler(euler *Euler) {
	bb.Call("setRotationFromEuler", euler)
}
func (bb *Bone) SetRotationFromMatrix(m *Matrix4) {
	bb.Call("setRotationFromMatrix", m)
}
func (bb *Bone) SetRotationFromQuaternion(q *Quaternion) {
	bb.Call("setRotationFromQuaternion", q)
}
func (bb *Bone) ToJSON(meta js.Value) js.Value {
	return bb.Call("toJSON", meta)
}
func (bb *Bone) TranslateOnAxis(axis *Vector3, distance float64) *Bone {
	return &Bone{Value: bb.Call("translateOnAxis", axis, distance)}
}
func (bb *Bone) TranslateX(distance float64) *Bone {
	return &Bone{Value: bb.Call("translateX", distance)}
}
func (bb *Bone) TranslateY(distance float64) *Bone {
	return &Bone{Value: bb.Call("translateY", distance)}
}
func (bb *Bone) TranslateZ(distance float64) *Bone {
	return &Bone{Value: bb.Call("translateZ", distance)}
}
func (bb *Bone) Traverse(callback js.Value) {
	bb.Call("traverse", callback)
}
func (bb *Bone) TraverseAncestors(callback js.Value) {
	bb.Call("traverseAncestors", callback)
}
func (bb *Bone) TraverseVisible(callback js.Value) {
	bb.Call("traverseVisible", callback)
}
func (bb *Bone) UpdateMatrix() {
	bb.Call("updateMatrix")
}
func (bb *Bone) UpdateMatrixWorld(force bool) {
	bb.Call("updateMatrixWorld", force)
}
func (bb *Bone) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	bb.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (bb *Bone) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: bb.Call("worldToLocal", vector)}
}

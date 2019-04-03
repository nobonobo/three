// Code generated from "objects/Group.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Group struct {
	js.Value
}

func NewGroup() *Group {
	return &Group{Value: get("Group").New()}
}
func (gg *Group) CastShadow() bool {
	return gg.Get("castShadow").Bool()
}
func (gg *Group) SetCastShadow(v bool) {
	gg.Set("castShadow", v)
}
func (gg *Group) Children() js.Value {
	return gg.Get("children")
}
func (gg *Group) SetChildren(v js.Value) {
	gg.Set("children", v)
}
func (gg *Group) FrustumCulled() bool {
	return gg.Get("frustumCulled").Bool()
}
func (gg *Group) SetFrustumCulled(v bool) {
	gg.Set("frustumCulled", v)
}
func (gg *Group) Id() int {
	return gg.Get("id").Int()
}
func (gg *Group) SetId(v int) {
	gg.Set("id", v)
}
func (gg *Group) IsGroup() bool {
	return gg.Get("isGroup").Bool()
}
func (gg *Group) SetIsGroup(v bool) {
	gg.Set("isGroup", v)
}
func (gg *Group) IsObject3D() bool {
	return gg.Get("isObject3D").Bool()
}
func (gg *Group) SetIsObject3D(v bool) {
	gg.Set("isObject3D", v)
}
func (gg *Group) Layers() *Layers {
	return &Layers{Value: gg.Get("layers")}
}
func (gg *Group) SetLayers(v *Layers) {
	gg.Set("layers", v)
}
func (gg *Group) Matrix() *Matrix4 {
	return &Matrix4{Value: gg.Get("matrix")}
}
func (gg *Group) SetMatrix(v *Matrix4) {
	gg.Set("matrix", v)
}
func (gg *Group) MatrixAutoUpdate() bool {
	return gg.Get("matrixAutoUpdate").Bool()
}
func (gg *Group) SetMatrixAutoUpdate(v bool) {
	gg.Set("matrixAutoUpdate", v)
}
func (gg *Group) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: gg.Get("matrixWorld")}
}
func (gg *Group) SetMatrixWorld(v *Matrix4) {
	gg.Set("matrixWorld", v)
}
func (gg *Group) MatrixWorldNeedsUpdate() bool {
	return gg.Get("matrixWorldNeedsUpdate").Bool()
}
func (gg *Group) SetMatrixWorldNeedsUpdate(v bool) {
	gg.Set("matrixWorldNeedsUpdate", v)
}
func (gg *Group) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: gg.Get("modelViewMatrix")}
}
func (gg *Group) SetModelViewMatrix(v *Matrix4) {
	gg.Set("modelViewMatrix", v)
}
func (gg *Group) Name() string {
	return gg.Get("name").String()
}
func (gg *Group) SetName(v string) {
	gg.Set("name", v)
}
func (gg *Group) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: gg.Get("normalMatrix")}
}
func (gg *Group) SetNormalMatrix(v *Matrix3) {
	gg.Set("normalMatrix", v)
}
func (gg *Group) OnAfterRender() js.Value {
	return gg.Get("onAfterRender")
}
func (gg *Group) SetOnAfterRender(v js.Value) {
	gg.Set("onAfterRender", v)
}
func (gg *Group) OnBeforeRender() js.Value {
	return gg.Get("onBeforeRender")
}
func (gg *Group) SetOnBeforeRender(v js.Value) {
	gg.Set("onBeforeRender", v)
}
func (gg *Group) Parent() *Object3D {
	return &Object3D{Value: gg.Get("parent")}
}
func (gg *Group) SetParent(v *Object3D) {
	gg.Set("parent", v)
}
func (gg *Group) Position() *Vector3 {
	return &Vector3{Value: gg.Get("position")}
}
func (gg *Group) SetPosition(v *Vector3) {
	gg.Set("position", v)
}
func (gg *Group) Quaternion() *Quaternion {
	return &Quaternion{Value: gg.Get("quaternion")}
}
func (gg *Group) SetQuaternion(v *Quaternion) {
	gg.Set("quaternion", v)
}
func (gg *Group) ReceiveShadow() bool {
	return gg.Get("receiveShadow").Bool()
}
func (gg *Group) SetReceiveShadow(v bool) {
	gg.Set("receiveShadow", v)
}
func (gg *Group) RenderOrder() float64 {
	return gg.Get("renderOrder").Float()
}
func (gg *Group) SetRenderOrder(v float64) {
	gg.Set("renderOrder", v)
}
func (gg *Group) Rotation() *Euler {
	return &Euler{Value: gg.Get("rotation")}
}
func (gg *Group) SetRotation(v *Euler) {
	gg.Set("rotation", v)
}
func (gg *Group) Scale() *Vector3 {
	return &Vector3{Value: gg.Get("scale")}
}
func (gg *Group) SetScale(v *Vector3) {
	gg.Set("scale", v)
}
func (gg *Group) Type() string {
	return gg.Get("type").String()
}
func (gg *Group) SetType(v string) {
	gg.Set("type", v)
}
func (gg *Group) Up() *Vector3 {
	return &Vector3{Value: gg.Get("up")}
}
func (gg *Group) SetUp(v *Vector3) {
	gg.Set("up", v)
}
func (gg *Group) UserData() js.Value {
	return gg.Get("userData")
}
func (gg *Group) SetUserData(v js.Value) {
	gg.Set("userData", v)
}
func (gg *Group) Uuid() string {
	return gg.Get("uuid").String()
}
func (gg *Group) SetUuid(v string) {
	gg.Set("uuid", v)
}
func (gg *Group) Visible() bool {
	return gg.Get("visible").Bool()
}
func (gg *Group) SetVisible(v bool) {
	gg.Set("visible", v)
}
func (gg *Group) DefaultMatrixAutoUpdate() bool {
	return gg.Get("DefaultMatrixAutoUpdate").Bool()
}
func (gg *Group) SetDefaultMatrixAutoUpdate(v bool) {
	gg.Set("DefaultMatrixAutoUpdate", v)
}
func (gg *Group) DefaultUp() *Vector3 {
	return &Vector3{Value: gg.Get("DefaultUp")}
}
func (gg *Group) SetDefaultUp(v *Vector3) {
	gg.Set("DefaultUp", v)
}
func (gg *Group) Add(object js.Value) *Group {
	return &Group{Value: gg.Call("add", object)}
}
func (gg *Group) AddEventListener(typ string, listener js.Value) {
	gg.Call("addEventListener", typ, listener)
}
func (gg *Group) ApplyMatrix(matrix *Matrix4) {
	gg.Call("applyMatrix", matrix)
}
func (gg *Group) ApplyQuaternion(quaternion *Quaternion) *Group {
	return &Group{Value: gg.Call("applyQuaternion", quaternion)}
}
func (gg *Group) Clone(recursive bool) *Group {
	return &Group{Value: gg.Call("clone", recursive)}
}
func (gg *Group) Copy(source *Object3D, recursive bool) *Group {
	return &Group{Value: gg.Call("copy", source, recursive)}
}
func (gg *Group) DispatchEvent(event js.Value) {
	gg.Call("dispatchEvent", event)
}
func (gg *Group) GetObjectById(id int) *Object3D {
	return &Object3D{Value: gg.Call("getObjectById", id)}
}
func (gg *Group) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: gg.Call("getObjectByName", name)}
}
func (gg *Group) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: gg.Call("getObjectByProperty", name, value)}
}
func (gg *Group) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: gg.Call("getWorldDirection", target)}
}
func (gg *Group) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: gg.Call("getWorldPosition", target)}
}
func (gg *Group) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: gg.Call("getWorldQuaternion", target)}
}
func (gg *Group) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: gg.Call("getWorldScale", target)}
}
func (gg *Group) HasEventListener(typ string, listener js.Value) bool {
	return gg.Call("hasEventListener", typ, listener).Bool()
}
func (gg *Group) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: gg.Call("localToWorld", vector)}
}
func (gg *Group) LookAt(vector *Vector3, y float64, z float64) {
	gg.Call("lookAt", vector, y, z)
}
func (gg *Group) Raycast(raycaster *Raycaster, intersects js.Value) {
	gg.Call("raycast", raycaster, intersects)
}
func (gg *Group) Remove(object js.Value) *Group {
	return &Group{Value: gg.Call("remove", object)}
}
func (gg *Group) RemoveEventListener(typ string, listener js.Value) {
	gg.Call("removeEventListener", typ, listener)
}
func (gg *Group) RotateOnAxis(axis *Vector3, angle float64) *Group {
	return &Group{Value: gg.Call("rotateOnAxis", axis, angle)}
}
func (gg *Group) RotateOnWorldAxis(axis *Vector3, angle float64) *Group {
	return &Group{Value: gg.Call("rotateOnWorldAxis", axis, angle)}
}
func (gg *Group) RotateX(angle float64) *Group {
	return &Group{Value: gg.Call("rotateX", angle)}
}
func (gg *Group) RotateY(angle float64) *Group {
	return &Group{Value: gg.Call("rotateY", angle)}
}
func (gg *Group) RotateZ(angle float64) *Group {
	return &Group{Value: gg.Call("rotateZ", angle)}
}
func (gg *Group) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	gg.Call("setRotationFromAxisAngle", axis, angle)
}
func (gg *Group) SetRotationFromEuler(euler *Euler) {
	gg.Call("setRotationFromEuler", euler)
}
func (gg *Group) SetRotationFromMatrix(m *Matrix4) {
	gg.Call("setRotationFromMatrix", m)
}
func (gg *Group) SetRotationFromQuaternion(q *Quaternion) {
	gg.Call("setRotationFromQuaternion", q)
}
func (gg *Group) ToJSON(meta js.Value) js.Value {
	return gg.Call("toJSON", meta)
}
func (gg *Group) TranslateOnAxis(axis *Vector3, distance float64) *Group {
	return &Group{Value: gg.Call("translateOnAxis", axis, distance)}
}
func (gg *Group) TranslateX(distance float64) *Group {
	return &Group{Value: gg.Call("translateX", distance)}
}
func (gg *Group) TranslateY(distance float64) *Group {
	return &Group{Value: gg.Call("translateY", distance)}
}
func (gg *Group) TranslateZ(distance float64) *Group {
	return &Group{Value: gg.Call("translateZ", distance)}
}
func (gg *Group) Traverse(callback js.Value) {
	gg.Call("traverse", callback)
}
func (gg *Group) TraverseAncestors(callback js.Value) {
	gg.Call("traverseAncestors", callback)
}
func (gg *Group) TraverseVisible(callback js.Value) {
	gg.Call("traverseVisible", callback)
}
func (gg *Group) UpdateMatrix() {
	gg.Call("updateMatrix")
}
func (gg *Group) UpdateMatrixWorld(force bool) {
	gg.Call("updateMatrixWorld", force)
}
func (gg *Group) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	gg.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (gg *Group) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: gg.Call("worldToLocal", vector)}
}

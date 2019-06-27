// Code generated from "objects/Points.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Points extend: [Object3D]
type Points struct {
	js.Value
}

func NewPoints(geometry Geometry, material Material) *Points {
	return &Points{Value: get("Points").New(geometry.JSValue(), material.JSValue())}
}
func (pp *Points) JSValue() js.Value {
	return pp.Value
}
func (pp *Points) CastShadow() bool {
	return pp.Get("castShadow").Bool()
}
func (pp *Points) SetCastShadow(v bool) {
	pp.Set("castShadow", v)
}
func (pp *Points) Children() js.Value {
	return pp.Get("children")
}
func (pp *Points) SetChildren(v js.Value) {
	pp.Set("children", v)
}
func (pp *Points) FrustumCulled() bool {
	return pp.Get("frustumCulled").Bool()
}
func (pp *Points) SetFrustumCulled(v bool) {
	pp.Set("frustumCulled", v)
}
func (pp *Points) Geometry() Geometry {
	return &GeometryImpl{Value: pp.Get("geometry")}
}
func (pp *Points) SetGeometry(v Geometry) {
	pp.Set("geometry", v.JSValue())
}
func (pp *Points) Id() int {
	return pp.Get("id").Int()
}
func (pp *Points) SetId(v int) {
	pp.Set("id", v)
}
func (pp *Points) IsObject3D() bool {
	return pp.Get("isObject3D").Bool()
}
func (pp *Points) SetIsObject3D(v bool) {
	pp.Set("isObject3D", v)
}
func (pp *Points) IsPoints() bool {
	return pp.Get("isPoints").Bool()
}
func (pp *Points) SetIsPoints(v bool) {
	pp.Set("isPoints", v)
}
func (pp *Points) Layers() *Layers {
	return &Layers{Value: pp.Get("layers")}
}
func (pp *Points) SetLayers(v *Layers) {
	pp.Set("layers", v.Value)
}
func (pp *Points) Material() Material {
	return &MaterialImpl{Value: pp.Get("material")}
}
func (pp *Points) SetMaterial(v Material) {
	pp.Set("material", v.JSValue())
}
func (pp *Points) Matrix() *Matrix4 {
	return &Matrix4{Value: pp.Get("matrix")}
}
func (pp *Points) SetMatrix(v *Matrix4) {
	pp.Set("matrix", v.Value)
}
func (pp *Points) MatrixAutoUpdate() bool {
	return pp.Get("matrixAutoUpdate").Bool()
}
func (pp *Points) SetMatrixAutoUpdate(v bool) {
	pp.Set("matrixAutoUpdate", v)
}
func (pp *Points) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: pp.Get("matrixWorld")}
}
func (pp *Points) SetMatrixWorld(v *Matrix4) {
	pp.Set("matrixWorld", v.Value)
}
func (pp *Points) MatrixWorldNeedsUpdate() bool {
	return pp.Get("matrixWorldNeedsUpdate").Bool()
}
func (pp *Points) SetMatrixWorldNeedsUpdate(v bool) {
	pp.Set("matrixWorldNeedsUpdate", v)
}
func (pp *Points) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: pp.Get("modelViewMatrix")}
}
func (pp *Points) SetModelViewMatrix(v *Matrix4) {
	pp.Set("modelViewMatrix", v.Value)
}
func (pp *Points) Name() string {
	return pp.Get("name").String()
}
func (pp *Points) SetName(v string) {
	pp.Set("name", v)
}
func (pp *Points) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: pp.Get("normalMatrix")}
}
func (pp *Points) SetNormalMatrix(v *Matrix3) {
	pp.Set("normalMatrix", v.Value)
}
func (pp *Points) OnAfterRender() js.Value {
	return pp.Get("onAfterRender")
}
func (pp *Points) SetOnAfterRender(v js.Value) {
	pp.Set("onAfterRender", v)
}
func (pp *Points) OnBeforeRender() js.Value {
	return pp.Get("onBeforeRender")
}
func (pp *Points) SetOnBeforeRender(v js.Value) {
	pp.Set("onBeforeRender", v)
}
func (pp *Points) Parent() *Object3D {
	return &Object3D{Value: pp.Get("parent")}
}
func (pp *Points) SetParent(v *Object3D) {
	pp.Set("parent", v.Value)
}
func (pp *Points) Position() *Vector3 {
	return &Vector3{Value: pp.Get("position")}
}
func (pp *Points) SetPosition(v *Vector3) {
	pp.Set("position", v.Value)
}
func (pp *Points) Quaternion() *Quaternion {
	return &Quaternion{Value: pp.Get("quaternion")}
}
func (pp *Points) SetQuaternion(v *Quaternion) {
	pp.Set("quaternion", v.Value)
}
func (pp *Points) ReceiveShadow() bool {
	return pp.Get("receiveShadow").Bool()
}
func (pp *Points) SetReceiveShadow(v bool) {
	pp.Set("receiveShadow", v)
}
func (pp *Points) RenderOrder() float64 {
	return pp.Get("renderOrder").Float()
}
func (pp *Points) SetRenderOrder(v float64) {
	pp.Set("renderOrder", v)
}
func (pp *Points) Rotation() *Euler {
	return &Euler{Value: pp.Get("rotation")}
}
func (pp *Points) SetRotation(v *Euler) {
	pp.Set("rotation", v.Value)
}
func (pp *Points) Scale() *Vector3 {
	return &Vector3{Value: pp.Get("scale")}
}
func (pp *Points) SetScale(v *Vector3) {
	pp.Set("scale", v.Value)
}
func (pp *Points) Type() string {
	return pp.Get("type").String()
}
func (pp *Points) SetType(v string) {
	pp.Set("type", v)
}
func (pp *Points) Up() *Vector3 {
	return &Vector3{Value: pp.Get("up")}
}
func (pp *Points) SetUp(v *Vector3) {
	pp.Set("up", v.Value)
}
func (pp *Points) UserData() js.Value {
	return pp.Get("userData")
}
func (pp *Points) SetUserData(v js.Value) {
	pp.Set("userData", v)
}
func (pp *Points) Uuid() string {
	return pp.Get("uuid").String()
}
func (pp *Points) SetUuid(v string) {
	pp.Set("uuid", v)
}
func (pp *Points) Visible() bool {
	return pp.Get("visible").Bool()
}
func (pp *Points) SetVisible(v bool) {
	pp.Set("visible", v)
}
func (pp *Points) DefaultMatrixAutoUpdate() bool {
	return pp.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pp *Points) SetDefaultMatrixAutoUpdate(v bool) {
	pp.Set("DefaultMatrixAutoUpdate", v)
}
func (pp *Points) DefaultUp() *Vector3 {
	return &Vector3{Value: pp.Get("DefaultUp")}
}
func (pp *Points) SetDefaultUp(v *Vector3) {
	pp.Set("DefaultUp", v.Value)
}
func (pp *Points) Add(object js.Value) *Points {
	return &Points{Value: pp.Call("add", object)}
}
func (pp *Points) AddEventListener(typ string, listener js.Value) {
	pp.Call("addEventListener", typ, listener)
}
func (pp *Points) ApplyMatrix(matrix *Matrix4) {
	pp.Call("applyMatrix", matrix)
}
func (pp *Points) ApplyQuaternion(quaternion *Quaternion) *Points {
	return &Points{Value: pp.Call("applyQuaternion", quaternion)}
}
func (pp *Points) Clone(recursive bool) *Points {
	return &Points{Value: pp.Call("clone", recursive)}
}
func (pp *Points) Copy(source *Object3D, recursive bool) *Points {
	return &Points{Value: pp.Call("copy", source, recursive)}
}
func (pp *Points) DispatchEvent(event js.Value) {
	pp.Call("dispatchEvent", event)
}
func (pp *Points) GetObjectById(id int) *Object3D {
	return &Object3D{Value: pp.Call("getObjectById", id)}
}
func (pp *Points) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: pp.Call("getObjectByName", name)}
}
func (pp *Points) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: pp.Call("getObjectByProperty", name, value)}
}
func (pp *Points) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("getWorldDirection", target)}
}
func (pp *Points) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("getWorldPosition", target)}
}
func (pp *Points) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: pp.Call("getWorldQuaternion", target)}
}
func (pp *Points) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("getWorldScale", target)}
}
func (pp *Points) HasEventListener(typ string, listener js.Value) bool {
	return pp.Call("hasEventListener", typ, listener).Bool()
}
func (pp *Points) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("localToWorld", vector)}
}
func (pp *Points) LookAt(vector *Vector3, y float64, z float64) {
	pp.Call("lookAt", vector, y, z)
}
func (pp *Points) Raycast(raycaster *Raycaster, intersects js.Value) {
	pp.Call("raycast", raycaster, intersects)
}
func (pp *Points) Remove(object js.Value) *Points {
	return &Points{Value: pp.Call("remove", object)}
}
func (pp *Points) RemoveEventListener(typ string, listener js.Value) {
	pp.Call("removeEventListener", typ, listener)
}
func (pp *Points) RotateOnAxis(axis *Vector3, angle float64) *Points {
	return &Points{Value: pp.Call("rotateOnAxis", axis, angle)}
}
func (pp *Points) RotateOnWorldAxis(axis *Vector3, angle float64) *Points {
	return &Points{Value: pp.Call("rotateOnWorldAxis", axis, angle)}
}
func (pp *Points) RotateX(angle float64) *Points {
	return &Points{Value: pp.Call("rotateX", angle)}
}
func (pp *Points) RotateY(angle float64) *Points {
	return &Points{Value: pp.Call("rotateY", angle)}
}
func (pp *Points) RotateZ(angle float64) *Points {
	return &Points{Value: pp.Call("rotateZ", angle)}
}
func (pp *Points) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	pp.Call("setRotationFromAxisAngle", axis, angle)
}
func (pp *Points) SetRotationFromEuler(euler *Euler) {
	pp.Call("setRotationFromEuler", euler)
}
func (pp *Points) SetRotationFromMatrix(m *Matrix4) {
	pp.Call("setRotationFromMatrix", m)
}
func (pp *Points) SetRotationFromQuaternion(q *Quaternion) {
	pp.Call("setRotationFromQuaternion", q)
}
func (pp *Points) ToJSON(meta js.Value) js.Value {
	return pp.Call("toJSON", meta)
}
func (pp *Points) TranslateOnAxis(axis *Vector3, distance float64) *Points {
	return &Points{Value: pp.Call("translateOnAxis", axis, distance)}
}
func (pp *Points) TranslateX(distance float64) *Points {
	return &Points{Value: pp.Call("translateX", distance)}
}
func (pp *Points) TranslateY(distance float64) *Points {
	return &Points{Value: pp.Call("translateY", distance)}
}
func (pp *Points) TranslateZ(distance float64) *Points {
	return &Points{Value: pp.Call("translateZ", distance)}
}
func (pp *Points) Traverse(callback js.Value) {
	pp.Call("traverse", callback)
}
func (pp *Points) TraverseAncestors(callback js.Value) {
	pp.Call("traverseAncestors", callback)
}
func (pp *Points) TraverseVisible(callback js.Value) {
	pp.Call("traverseVisible", callback)
}
func (pp *Points) UpdateMatrix() {
	pp.Call("updateMatrix")
}
func (pp *Points) UpdateMatrixWorld(force bool) {
	pp.Call("updateMatrixWorld", force)
}
func (pp *Points) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pp.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pp *Points) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: pp.Call("worldToLocal", vector)}
}

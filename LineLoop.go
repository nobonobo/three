// Code generated from "objects/LineLoop.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// LineLoop extend: [Line]
type LineLoop struct {
	js.Value
}

func NewLineLoop(geometry Geometry, material Material) *LineLoop {
	return &LineLoop{Value: get("LineLoop").New(geometry.JSValue(), material.JSValue())}
}
func (ll *LineLoop) JSValue() js.Value {
	return ll.Value
}
func (ll *LineLoop) CastShadow() bool {
	return ll.Get("castShadow").Bool()
}
func (ll *LineLoop) SetCastShadow(v bool) {
	ll.Set("castShadow", v)
}
func (ll *LineLoop) Children() js.Value {
	return ll.Get("children")
}
func (ll *LineLoop) SetChildren(v js.Value) {
	ll.Set("children", v)
}
func (ll *LineLoop) FrustumCulled() bool {
	return ll.Get("frustumCulled").Bool()
}
func (ll *LineLoop) SetFrustumCulled(v bool) {
	ll.Set("frustumCulled", v)
}
func (ll *LineLoop) Geometry() Geometry {
	return &GeometryImpl{Value: ll.Get("geometry")}
}
func (ll *LineLoop) SetGeometry(v Geometry) {
	ll.Set("geometry", v.JSValue())
}
func (ll *LineLoop) Id() int {
	return ll.Get("id").Int()
}
func (ll *LineLoop) SetId(v int) {
	ll.Set("id", v)
}
func (ll *LineLoop) IsLine() bool {
	return ll.Get("isLine").Bool()
}
func (ll *LineLoop) SetIsLine(v bool) {
	ll.Set("isLine", v)
}
func (ll *LineLoop) IsLineLoop() bool {
	return ll.Get("isLineLoop").Bool()
}
func (ll *LineLoop) SetIsLineLoop(v bool) {
	ll.Set("isLineLoop", v)
}
func (ll *LineLoop) IsObject3D() bool {
	return ll.Get("isObject3D").Bool()
}
func (ll *LineLoop) SetIsObject3D(v bool) {
	ll.Set("isObject3D", v)
}
func (ll *LineLoop) Layers() *Layers {
	return &Layers{Value: ll.Get("layers")}
}
func (ll *LineLoop) SetLayers(v *Layers) {
	ll.Set("layers", v.JSValue())
}
func (ll *LineLoop) Material() Material {
	return &MaterialImpl{Value: ll.Get("material")}
}
func (ll *LineLoop) SetMaterial(v Material) {
	ll.Set("material", v.JSValue())
}
func (ll *LineLoop) Matrix() *Matrix4 {
	return &Matrix4{Value: ll.Get("matrix")}
}
func (ll *LineLoop) SetMatrix(v *Matrix4) {
	ll.Set("matrix", v.JSValue())
}
func (ll *LineLoop) MatrixAutoUpdate() bool {
	return ll.Get("matrixAutoUpdate").Bool()
}
func (ll *LineLoop) SetMatrixAutoUpdate(v bool) {
	ll.Set("matrixAutoUpdate", v)
}
func (ll *LineLoop) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ll.Get("matrixWorld")}
}
func (ll *LineLoop) SetMatrixWorld(v *Matrix4) {
	ll.Set("matrixWorld", v.JSValue())
}
func (ll *LineLoop) MatrixWorldNeedsUpdate() bool {
	return ll.Get("matrixWorldNeedsUpdate").Bool()
}
func (ll *LineLoop) SetMatrixWorldNeedsUpdate(v bool) {
	ll.Set("matrixWorldNeedsUpdate", v)
}
func (ll *LineLoop) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ll.Get("modelViewMatrix")}
}
func (ll *LineLoop) SetModelViewMatrix(v *Matrix4) {
	ll.Set("modelViewMatrix", v.JSValue())
}
func (ll *LineLoop) Name() string {
	return ll.Get("name").String()
}
func (ll *LineLoop) SetName(v string) {
	ll.Set("name", v)
}
func (ll *LineLoop) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ll.Get("normalMatrix")}
}
func (ll *LineLoop) SetNormalMatrix(v *Matrix3) {
	ll.Set("normalMatrix", v.JSValue())
}
func (ll *LineLoop) OnAfterRender() js.Value {
	return ll.Get("onAfterRender")
}
func (ll *LineLoop) SetOnAfterRender(v js.Value) {
	ll.Set("onAfterRender", v)
}
func (ll *LineLoop) OnBeforeRender() js.Value {
	return ll.Get("onBeforeRender")
}
func (ll *LineLoop) SetOnBeforeRender(v js.Value) {
	ll.Set("onBeforeRender", v)
}
func (ll *LineLoop) Parent() *Object3D {
	return &Object3D{Value: ll.Get("parent")}
}
func (ll *LineLoop) SetParent(v *Object3D) {
	ll.Set("parent", v.JSValue())
}
func (ll *LineLoop) Position() *Vector3 {
	return &Vector3{Value: ll.Get("position")}
}
func (ll *LineLoop) SetPosition(v *Vector3) {
	ll.Set("position", v.JSValue())
}
func (ll *LineLoop) Quaternion() *Quaternion {
	return &Quaternion{Value: ll.Get("quaternion")}
}
func (ll *LineLoop) SetQuaternion(v *Quaternion) {
	ll.Set("quaternion", v.JSValue())
}
func (ll *LineLoop) ReceiveShadow() bool {
	return ll.Get("receiveShadow").Bool()
}
func (ll *LineLoop) SetReceiveShadow(v bool) {
	ll.Set("receiveShadow", v)
}
func (ll *LineLoop) RenderOrder() float64 {
	return ll.Get("renderOrder").Float()
}
func (ll *LineLoop) SetRenderOrder(v float64) {
	ll.Set("renderOrder", v)
}
func (ll *LineLoop) Rotation() *Euler {
	return &Euler{Value: ll.Get("rotation")}
}
func (ll *LineLoop) SetRotation(v *Euler) {
	ll.Set("rotation", v.JSValue())
}
func (ll *LineLoop) Scale() *Vector3 {
	return &Vector3{Value: ll.Get("scale")}
}
func (ll *LineLoop) SetScale(v *Vector3) {
	ll.Set("scale", v.JSValue())
}
func (ll *LineLoop) Type() string {
	return ll.Get("type").String()
}
func (ll *LineLoop) SetType(v string) {
	ll.Set("type", v)
}
func (ll *LineLoop) Up() *Vector3 {
	return &Vector3{Value: ll.Get("up")}
}
func (ll *LineLoop) SetUp(v *Vector3) {
	ll.Set("up", v.JSValue())
}
func (ll *LineLoop) UserData() js.Value {
	return ll.Get("userData")
}
func (ll *LineLoop) SetUserData(v js.Value) {
	ll.Set("userData", v)
}
func (ll *LineLoop) Uuid() string {
	return ll.Get("uuid").String()
}
func (ll *LineLoop) SetUuid(v string) {
	ll.Set("uuid", v)
}
func (ll *LineLoop) Visible() bool {
	return ll.Get("visible").Bool()
}
func (ll *LineLoop) SetVisible(v bool) {
	ll.Set("visible", v)
}
func (ll *LineLoop) DefaultMatrixAutoUpdate() bool {
	return ll.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ll *LineLoop) SetDefaultMatrixAutoUpdate(v bool) {
	ll.Set("DefaultMatrixAutoUpdate", v)
}
func (ll *LineLoop) DefaultUp() *Vector3 {
	return &Vector3{Value: ll.Get("DefaultUp")}
}
func (ll *LineLoop) SetDefaultUp(v *Vector3) {
	ll.Set("DefaultUp", v.JSValue())
}
func (ll *LineLoop) Add(object js.Value) *LineLoop {
	return &LineLoop{Value: ll.Call("add", object)}
}
func (ll *LineLoop) AddEventListener(typ string, listener js.Value) {
	ll.Call("addEventListener", typ, listener)
}
func (ll *LineLoop) ApplyMatrix(matrix *Matrix4) {
	ll.Call("applyMatrix", matrix.JSValue())
}
func (ll *LineLoop) ApplyQuaternion(quaternion *Quaternion) *LineLoop {
	return &LineLoop{Value: ll.Call("applyQuaternion", quaternion)}
}
func (ll *LineLoop) Clone(recursive bool) *LineLoop {
	return &LineLoop{Value: ll.Call("clone", recursive)}
}
func (ll *LineLoop) ComputeLineDistances() *LineLoop {
	return &LineLoop{Value: ll.Call("computeLineDistances")}
}
func (ll *LineLoop) Copy(source *Object3D, recursive bool) *LineLoop {
	return &LineLoop{Value: ll.Call("copy", source, recursive)}
}
func (ll *LineLoop) DispatchEvent(event js.Value) {
	ll.Call("dispatchEvent", event)
}
func (ll *LineLoop) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ll.Call("getObjectById", id)}
}
func (ll *LineLoop) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ll.Call("getObjectByName", name)}
}
func (ll *LineLoop) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ll.Call("getObjectByProperty", name, value)}
}
func (ll *LineLoop) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldDirection", target)}
}
func (ll *LineLoop) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldPosition", target)}
}
func (ll *LineLoop) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ll.Call("getWorldQuaternion", target)}
}
func (ll *LineLoop) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldScale", target)}
}
func (ll *LineLoop) HasEventListener(typ string, listener js.Value) bool {
	return ll.Call("hasEventListener", typ, listener).Bool()
}
func (ll *LineLoop) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("localToWorld", vector)}
}
func (ll *LineLoop) LookAt(vector *Vector3, y float64, z float64) {
	ll.Call("lookAt", vector, y, z)
}
func (ll *LineLoop) Raycast(raycaster *Raycaster, intersects js.Value) {
	ll.Call("raycast", raycaster.JSValue(), intersects)
}
func (ll *LineLoop) Remove(object js.Value) *LineLoop {
	return &LineLoop{Value: ll.Call("remove", object)}
}
func (ll *LineLoop) RemoveEventListener(typ string, listener js.Value) {
	ll.Call("removeEventListener", typ, listener)
}
func (ll *LineLoop) RotateOnAxis(axis *Vector3, angle float64) *LineLoop {
	return &LineLoop{Value: ll.Call("rotateOnAxis", axis, angle)}
}
func (ll *LineLoop) RotateOnWorldAxis(axis *Vector3, angle float64) *LineLoop {
	return &LineLoop{Value: ll.Call("rotateOnWorldAxis", axis, angle)}
}
func (ll *LineLoop) RotateX(angle float64) *LineLoop {
	return &LineLoop{Value: ll.Call("rotateX", angle)}
}
func (ll *LineLoop) RotateY(angle float64) *LineLoop {
	return &LineLoop{Value: ll.Call("rotateY", angle)}
}
func (ll *LineLoop) RotateZ(angle float64) *LineLoop {
	return &LineLoop{Value: ll.Call("rotateZ", angle)}
}
func (ll *LineLoop) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ll.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (ll *LineLoop) SetRotationFromEuler(euler *Euler) {
	ll.Call("setRotationFromEuler", euler.JSValue())
}
func (ll *LineLoop) SetRotationFromMatrix(m *Matrix4) {
	ll.Call("setRotationFromMatrix", m.JSValue())
}
func (ll *LineLoop) SetRotationFromQuaternion(q *Quaternion) {
	ll.Call("setRotationFromQuaternion", q.JSValue())
}
func (ll *LineLoop) ToJSON(meta js.Value) js.Value {
	return ll.Call("toJSON", meta)
}
func (ll *LineLoop) TranslateOnAxis(axis *Vector3, distance float64) *LineLoop {
	return &LineLoop{Value: ll.Call("translateOnAxis", axis, distance)}
}
func (ll *LineLoop) TranslateX(distance float64) *LineLoop {
	return &LineLoop{Value: ll.Call("translateX", distance)}
}
func (ll *LineLoop) TranslateY(distance float64) *LineLoop {
	return &LineLoop{Value: ll.Call("translateY", distance)}
}
func (ll *LineLoop) TranslateZ(distance float64) *LineLoop {
	return &LineLoop{Value: ll.Call("translateZ", distance)}
}
func (ll *LineLoop) Traverse(callback js.Value) {
	ll.Call("traverse", callback)
}
func (ll *LineLoop) TraverseAncestors(callback js.Value) {
	ll.Call("traverseAncestors", callback)
}
func (ll *LineLoop) TraverseVisible(callback js.Value) {
	ll.Call("traverseVisible", callback)
}
func (ll *LineLoop) UpdateMatrix() {
	ll.Call("updateMatrix")
}
func (ll *LineLoop) UpdateMatrixWorld(force bool) {
	ll.Call("updateMatrixWorld", force)
}
func (ll *LineLoop) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ll.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ll *LineLoop) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("worldToLocal", vector)}
}

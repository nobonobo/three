// Code generated from "objects/Line.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Line extend: [Object3D]
type Line struct {
	js.Value
}

func NewLine(geometry Geometry, material Material, mode float64) *Line {
	return &Line{Value: get("Line").New(geometry.JSValue(), material.JSValue(), mode)}
}
func (ll *Line) JSValue() js.Value {
	return ll.Value
}
func (ll *Line) CastShadow() bool {
	return ll.Get("castShadow").Bool()
}
func (ll *Line) SetCastShadow(v bool) {
	ll.Set("castShadow", v)
}
func (ll *Line) Children() js.Value {
	return ll.Get("children")
}
func (ll *Line) SetChildren(v js.Value) {
	ll.Set("children", v)
}
func (ll *Line) FrustumCulled() bool {
	return ll.Get("frustumCulled").Bool()
}
func (ll *Line) SetFrustumCulled(v bool) {
	ll.Set("frustumCulled", v)
}
func (ll *Line) Geometry() Geometry {
	return &GeometryImpl{Value: ll.Get("geometry")}
}
func (ll *Line) SetGeometry(v Geometry) {
	ll.Set("geometry", v.JSValue())
}
func (ll *Line) Id() int {
	return ll.Get("id").Int()
}
func (ll *Line) SetId(v int) {
	ll.Set("id", v)
}
func (ll *Line) IsLine() bool {
	return ll.Get("isLine").Bool()
}
func (ll *Line) SetIsLine(v bool) {
	ll.Set("isLine", v)
}
func (ll *Line) IsObject3D() bool {
	return ll.Get("isObject3D").Bool()
}
func (ll *Line) SetIsObject3D(v bool) {
	ll.Set("isObject3D", v)
}
func (ll *Line) Layers() *Layers {
	return &Layers{Value: ll.Get("layers")}
}
func (ll *Line) SetLayers(v *Layers) {
	ll.Set("layers", v.JSValue())
}
func (ll *Line) Material() Material {
	return &MaterialImpl{Value: ll.Get("material")}
}
func (ll *Line) SetMaterial(v Material) {
	ll.Set("material", v.JSValue())
}
func (ll *Line) Matrix() *Matrix4 {
	return &Matrix4{Value: ll.Get("matrix")}
}
func (ll *Line) SetMatrix(v *Matrix4) {
	ll.Set("matrix", v.JSValue())
}
func (ll *Line) MatrixAutoUpdate() bool {
	return ll.Get("matrixAutoUpdate").Bool()
}
func (ll *Line) SetMatrixAutoUpdate(v bool) {
	ll.Set("matrixAutoUpdate", v)
}
func (ll *Line) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ll.Get("matrixWorld")}
}
func (ll *Line) SetMatrixWorld(v *Matrix4) {
	ll.Set("matrixWorld", v.JSValue())
}
func (ll *Line) MatrixWorldNeedsUpdate() bool {
	return ll.Get("matrixWorldNeedsUpdate").Bool()
}
func (ll *Line) SetMatrixWorldNeedsUpdate(v bool) {
	ll.Set("matrixWorldNeedsUpdate", v)
}
func (ll *Line) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ll.Get("modelViewMatrix")}
}
func (ll *Line) SetModelViewMatrix(v *Matrix4) {
	ll.Set("modelViewMatrix", v.JSValue())
}
func (ll *Line) Name() string {
	return ll.Get("name").String()
}
func (ll *Line) SetName(v string) {
	ll.Set("name", v)
}
func (ll *Line) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ll.Get("normalMatrix")}
}
func (ll *Line) SetNormalMatrix(v *Matrix3) {
	ll.Set("normalMatrix", v.JSValue())
}
func (ll *Line) OnAfterRender() js.Value {
	return ll.Get("onAfterRender")
}
func (ll *Line) SetOnAfterRender(v js.Value) {
	ll.Set("onAfterRender", v)
}
func (ll *Line) OnBeforeRender() js.Value {
	return ll.Get("onBeforeRender")
}
func (ll *Line) SetOnBeforeRender(v js.Value) {
	ll.Set("onBeforeRender", v)
}
func (ll *Line) Parent() *Object3D {
	return &Object3D{Value: ll.Get("parent")}
}
func (ll *Line) SetParent(v *Object3D) {
	ll.Set("parent", v.JSValue())
}
func (ll *Line) Position() *Vector3 {
	return &Vector3{Value: ll.Get("position")}
}
func (ll *Line) SetPosition(v *Vector3) {
	ll.Set("position", v.JSValue())
}
func (ll *Line) Quaternion() *Quaternion {
	return &Quaternion{Value: ll.Get("quaternion")}
}
func (ll *Line) SetQuaternion(v *Quaternion) {
	ll.Set("quaternion", v.JSValue())
}
func (ll *Line) ReceiveShadow() bool {
	return ll.Get("receiveShadow").Bool()
}
func (ll *Line) SetReceiveShadow(v bool) {
	ll.Set("receiveShadow", v)
}
func (ll *Line) RenderOrder() float64 {
	return ll.Get("renderOrder").Float()
}
func (ll *Line) SetRenderOrder(v float64) {
	ll.Set("renderOrder", v)
}
func (ll *Line) Rotation() *Euler {
	return &Euler{Value: ll.Get("rotation")}
}
func (ll *Line) SetRotation(v *Euler) {
	ll.Set("rotation", v.JSValue())
}
func (ll *Line) Scale() *Vector3 {
	return &Vector3{Value: ll.Get("scale")}
}
func (ll *Line) SetScale(v *Vector3) {
	ll.Set("scale", v.JSValue())
}
func (ll *Line) Type() string {
	return ll.Get("type").String()
}
func (ll *Line) SetType(v string) {
	ll.Set("type", v)
}
func (ll *Line) Up() *Vector3 {
	return &Vector3{Value: ll.Get("up")}
}
func (ll *Line) SetUp(v *Vector3) {
	ll.Set("up", v.JSValue())
}
func (ll *Line) UserData() js.Value {
	return ll.Get("userData")
}
func (ll *Line) SetUserData(v js.Value) {
	ll.Set("userData", v)
}
func (ll *Line) Uuid() string {
	return ll.Get("uuid").String()
}
func (ll *Line) SetUuid(v string) {
	ll.Set("uuid", v)
}
func (ll *Line) Visible() bool {
	return ll.Get("visible").Bool()
}
func (ll *Line) SetVisible(v bool) {
	ll.Set("visible", v)
}
func (ll *Line) DefaultMatrixAutoUpdate() bool {
	return ll.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ll *Line) SetDefaultMatrixAutoUpdate(v bool) {
	ll.Set("DefaultMatrixAutoUpdate", v)
}
func (ll *Line) DefaultUp() *Vector3 {
	return &Vector3{Value: ll.Get("DefaultUp")}
}
func (ll *Line) SetDefaultUp(v *Vector3) {
	ll.Set("DefaultUp", v.JSValue())
}
func (ll *Line) Add(object js.Value) *Line {
	return &Line{Value: ll.Call("add", object)}
}
func (ll *Line) AddEventListener(typ string, listener js.Value) {
	ll.Call("addEventListener", typ, listener)
}
func (ll *Line) ApplyMatrix(matrix *Matrix4) {
	ll.Call("applyMatrix", matrix.JSValue())
}
func (ll *Line) ApplyQuaternion(quaternion *Quaternion) *Line {
	return &Line{Value: ll.Call("applyQuaternion", quaternion)}
}
func (ll *Line) Clone(recursive bool) *Line {
	return &Line{Value: ll.Call("clone", recursive)}
}
func (ll *Line) ComputeLineDistances() *Line {
	return &Line{Value: ll.Call("computeLineDistances")}
}
func (ll *Line) Copy(source *Object3D, recursive bool) *Line {
	return &Line{Value: ll.Call("copy", source, recursive)}
}
func (ll *Line) DispatchEvent(event js.Value) {
	ll.Call("dispatchEvent", event)
}
func (ll *Line) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ll.Call("getObjectById", id)}
}
func (ll *Line) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ll.Call("getObjectByName", name)}
}
func (ll *Line) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ll.Call("getObjectByProperty", name, value)}
}
func (ll *Line) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldDirection", target)}
}
func (ll *Line) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldPosition", target)}
}
func (ll *Line) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ll.Call("getWorldQuaternion", target)}
}
func (ll *Line) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("getWorldScale", target)}
}
func (ll *Line) HasEventListener(typ string, listener js.Value) bool {
	return ll.Call("hasEventListener", typ, listener).Bool()
}
func (ll *Line) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("localToWorld", vector)}
}
func (ll *Line) LookAt(vector *Vector3, y float64, z float64) {
	ll.Call("lookAt", vector, y, z)
}
func (ll *Line) Raycast(raycaster *Raycaster, intersects js.Value) {
	ll.Call("raycast", raycaster.JSValue(), intersects)
}
func (ll *Line) Remove(object js.Value) *Line {
	return &Line{Value: ll.Call("remove", object)}
}
func (ll *Line) RemoveEventListener(typ string, listener js.Value) {
	ll.Call("removeEventListener", typ, listener)
}
func (ll *Line) RotateOnAxis(axis *Vector3, angle float64) *Line {
	return &Line{Value: ll.Call("rotateOnAxis", axis, angle)}
}
func (ll *Line) RotateOnWorldAxis(axis *Vector3, angle float64) *Line {
	return &Line{Value: ll.Call("rotateOnWorldAxis", axis, angle)}
}
func (ll *Line) RotateX(angle float64) *Line {
	return &Line{Value: ll.Call("rotateX", angle)}
}
func (ll *Line) RotateY(angle float64) *Line {
	return &Line{Value: ll.Call("rotateY", angle)}
}
func (ll *Line) RotateZ(angle float64) *Line {
	return &Line{Value: ll.Call("rotateZ", angle)}
}
func (ll *Line) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ll.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (ll *Line) SetRotationFromEuler(euler *Euler) {
	ll.Call("setRotationFromEuler", euler.JSValue())
}
func (ll *Line) SetRotationFromMatrix(m *Matrix4) {
	ll.Call("setRotationFromMatrix", m.JSValue())
}
func (ll *Line) SetRotationFromQuaternion(q *Quaternion) {
	ll.Call("setRotationFromQuaternion", q.JSValue())
}
func (ll *Line) ToJSON(meta js.Value) js.Value {
	return ll.Call("toJSON", meta)
}
func (ll *Line) TranslateOnAxis(axis *Vector3, distance float64) *Line {
	return &Line{Value: ll.Call("translateOnAxis", axis, distance)}
}
func (ll *Line) TranslateX(distance float64) *Line {
	return &Line{Value: ll.Call("translateX", distance)}
}
func (ll *Line) TranslateY(distance float64) *Line {
	return &Line{Value: ll.Call("translateY", distance)}
}
func (ll *Line) TranslateZ(distance float64) *Line {
	return &Line{Value: ll.Call("translateZ", distance)}
}
func (ll *Line) Traverse(callback js.Value) {
	ll.Call("traverse", callback)
}
func (ll *Line) TraverseAncestors(callback js.Value) {
	ll.Call("traverseAncestors", callback)
}
func (ll *Line) TraverseVisible(callback js.Value) {
	ll.Call("traverseVisible", callback)
}
func (ll *Line) UpdateMatrix() {
	ll.Call("updateMatrix")
}
func (ll *Line) UpdateMatrixWorld(force bool) {
	ll.Call("updateMatrixWorld", force)
}
func (ll *Line) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ll.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ll *Line) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ll.Call("worldToLocal", vector)}
}

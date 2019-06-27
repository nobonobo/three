// Code generated from "helpers/AxesHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// AxesHelper extend: [LineSegments]
type AxesHelper struct {
	js.Value
}

func NewAxesHelper(size float64) *AxesHelper {
	return &AxesHelper{Value: get("AxesHelper").New(size)}
}
func (ah *AxesHelper) JSValue() js.Value {
	return ah.Value
}
func (ah *AxesHelper) CastShadow() bool {
	return ah.Get("castShadow").Bool()
}
func (ah *AxesHelper) SetCastShadow(v bool) {
	ah.Set("castShadow", v)
}
func (ah *AxesHelper) Children() js.Value {
	return ah.Get("children")
}
func (ah *AxesHelper) SetChildren(v js.Value) {
	ah.Set("children", v)
}
func (ah *AxesHelper) FrustumCulled() bool {
	return ah.Get("frustumCulled").Bool()
}
func (ah *AxesHelper) SetFrustumCulled(v bool) {
	ah.Set("frustumCulled", v)
}
func (ah *AxesHelper) Geometry() Geometry {
	return &GeometryImpl{Value: ah.Get("geometry")}
}
func (ah *AxesHelper) SetGeometry(v Geometry) {
	ah.Set("geometry", v.JSValue())
}
func (ah *AxesHelper) Id() int {
	return ah.Get("id").Int()
}
func (ah *AxesHelper) SetId(v int) {
	ah.Set("id", v)
}
func (ah *AxesHelper) IsLine() bool {
	return ah.Get("isLine").Bool()
}
func (ah *AxesHelper) SetIsLine(v bool) {
	ah.Set("isLine", v)
}
func (ah *AxesHelper) IsLineSegments() bool {
	return ah.Get("isLineSegments").Bool()
}
func (ah *AxesHelper) SetIsLineSegments(v bool) {
	ah.Set("isLineSegments", v)
}
func (ah *AxesHelper) IsObject3D() bool {
	return ah.Get("isObject3D").Bool()
}
func (ah *AxesHelper) SetIsObject3D(v bool) {
	ah.Set("isObject3D", v)
}
func (ah *AxesHelper) Layers() *Layers {
	return &Layers{Value: ah.Get("layers")}
}
func (ah *AxesHelper) SetLayers(v *Layers) {
	ah.Set("layers", v.Value)
}
func (ah *AxesHelper) Material() Material {
	return &MaterialImpl{Value: ah.Get("material")}
}
func (ah *AxesHelper) SetMaterial(v Material) {
	ah.Set("material", v.JSValue())
}
func (ah *AxesHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: ah.Get("matrix")}
}
func (ah *AxesHelper) SetMatrix(v *Matrix4) {
	ah.Set("matrix", v.Value)
}
func (ah *AxesHelper) MatrixAutoUpdate() bool {
	return ah.Get("matrixAutoUpdate").Bool()
}
func (ah *AxesHelper) SetMatrixAutoUpdate(v bool) {
	ah.Set("matrixAutoUpdate", v)
}
func (ah *AxesHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ah.Get("matrixWorld")}
}
func (ah *AxesHelper) SetMatrixWorld(v *Matrix4) {
	ah.Set("matrixWorld", v.Value)
}
func (ah *AxesHelper) MatrixWorldNeedsUpdate() bool {
	return ah.Get("matrixWorldNeedsUpdate").Bool()
}
func (ah *AxesHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ah.Set("matrixWorldNeedsUpdate", v)
}
func (ah *AxesHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ah.Get("modelViewMatrix")}
}
func (ah *AxesHelper) SetModelViewMatrix(v *Matrix4) {
	ah.Set("modelViewMatrix", v.Value)
}
func (ah *AxesHelper) Name() string {
	return ah.Get("name").String()
}
func (ah *AxesHelper) SetName(v string) {
	ah.Set("name", v)
}
func (ah *AxesHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ah.Get("normalMatrix")}
}
func (ah *AxesHelper) SetNormalMatrix(v *Matrix3) {
	ah.Set("normalMatrix", v.Value)
}
func (ah *AxesHelper) OnAfterRender() js.Value {
	return ah.Get("onAfterRender")
}
func (ah *AxesHelper) SetOnAfterRender(v js.Value) {
	ah.Set("onAfterRender", v)
}
func (ah *AxesHelper) OnBeforeRender() js.Value {
	return ah.Get("onBeforeRender")
}
func (ah *AxesHelper) SetOnBeforeRender(v js.Value) {
	ah.Set("onBeforeRender", v)
}
func (ah *AxesHelper) Parent() *Object3D {
	return &Object3D{Value: ah.Get("parent")}
}
func (ah *AxesHelper) SetParent(v *Object3D) {
	ah.Set("parent", v.Value)
}
func (ah *AxesHelper) Position() *Vector3 {
	return &Vector3{Value: ah.Get("position")}
}
func (ah *AxesHelper) SetPosition(v *Vector3) {
	ah.Set("position", v.Value)
}
func (ah *AxesHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: ah.Get("quaternion")}
}
func (ah *AxesHelper) SetQuaternion(v *Quaternion) {
	ah.Set("quaternion", v.Value)
}
func (ah *AxesHelper) ReceiveShadow() bool {
	return ah.Get("receiveShadow").Bool()
}
func (ah *AxesHelper) SetReceiveShadow(v bool) {
	ah.Set("receiveShadow", v)
}
func (ah *AxesHelper) RenderOrder() float64 {
	return ah.Get("renderOrder").Float()
}
func (ah *AxesHelper) SetRenderOrder(v float64) {
	ah.Set("renderOrder", v)
}
func (ah *AxesHelper) Rotation() *Euler {
	return &Euler{Value: ah.Get("rotation")}
}
func (ah *AxesHelper) SetRotation(v *Euler) {
	ah.Set("rotation", v.Value)
}
func (ah *AxesHelper) Scale() *Vector3 {
	return &Vector3{Value: ah.Get("scale")}
}
func (ah *AxesHelper) SetScale(v *Vector3) {
	ah.Set("scale", v.Value)
}
func (ah *AxesHelper) Type() string {
	return ah.Get("type").String()
}
func (ah *AxesHelper) SetType(v string) {
	ah.Set("type", v)
}
func (ah *AxesHelper) Up() *Vector3 {
	return &Vector3{Value: ah.Get("up")}
}
func (ah *AxesHelper) SetUp(v *Vector3) {
	ah.Set("up", v.Value)
}
func (ah *AxesHelper) UserData() js.Value {
	return ah.Get("userData")
}
func (ah *AxesHelper) SetUserData(v js.Value) {
	ah.Set("userData", v)
}
func (ah *AxesHelper) Uuid() string {
	return ah.Get("uuid").String()
}
func (ah *AxesHelper) SetUuid(v string) {
	ah.Set("uuid", v)
}
func (ah *AxesHelper) Visible() bool {
	return ah.Get("visible").Bool()
}
func (ah *AxesHelper) SetVisible(v bool) {
	ah.Set("visible", v)
}
func (ah *AxesHelper) DefaultMatrixAutoUpdate() bool {
	return ah.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ah *AxesHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ah.Set("DefaultMatrixAutoUpdate", v)
}
func (ah *AxesHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: ah.Get("DefaultUp")}
}
func (ah *AxesHelper) SetDefaultUp(v *Vector3) {
	ah.Set("DefaultUp", v.Value)
}
func (ah *AxesHelper) Add(object js.Value) *AxesHelper {
	return &AxesHelper{Value: ah.Call("add", object)}
}
func (ah *AxesHelper) AddEventListener(typ string, listener js.Value) {
	ah.Call("addEventListener", typ, listener)
}
func (ah *AxesHelper) ApplyMatrix(matrix *Matrix4) {
	ah.Call("applyMatrix", matrix)
}
func (ah *AxesHelper) ApplyQuaternion(quaternion *Quaternion) *AxesHelper {
	return &AxesHelper{Value: ah.Call("applyQuaternion", quaternion)}
}
func (ah *AxesHelper) Clone(recursive bool) *AxesHelper {
	return &AxesHelper{Value: ah.Call("clone", recursive)}
}
func (ah *AxesHelper) ComputeLineDistances() *AxesHelper {
	return &AxesHelper{Value: ah.Call("computeLineDistances")}
}
func (ah *AxesHelper) Copy(source *Object3D, recursive bool) *AxesHelper {
	return &AxesHelper{Value: ah.Call("copy", source, recursive)}
}
func (ah *AxesHelper) DispatchEvent(event js.Value) {
	ah.Call("dispatchEvent", event)
}
func (ah *AxesHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ah.Call("getObjectById", id)}
}
func (ah *AxesHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ah.Call("getObjectByName", name)}
}
func (ah *AxesHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ah.Call("getObjectByProperty", name, value)}
}
func (ah *AxesHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("getWorldDirection", target)}
}
func (ah *AxesHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("getWorldPosition", target)}
}
func (ah *AxesHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ah.Call("getWorldQuaternion", target)}
}
func (ah *AxesHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("getWorldScale", target)}
}
func (ah *AxesHelper) HasEventListener(typ string, listener js.Value) bool {
	return ah.Call("hasEventListener", typ, listener).Bool()
}
func (ah *AxesHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("localToWorld", vector)}
}
func (ah *AxesHelper) LookAt(vector *Vector3, y float64, z float64) {
	ah.Call("lookAt", vector, y, z)
}
func (ah *AxesHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	ah.Call("raycast", raycaster, intersects)
}
func (ah *AxesHelper) Remove(object js.Value) *AxesHelper {
	return &AxesHelper{Value: ah.Call("remove", object)}
}
func (ah *AxesHelper) RemoveEventListener(typ string, listener js.Value) {
	ah.Call("removeEventListener", typ, listener)
}
func (ah *AxesHelper) RotateOnAxis(axis *Vector3, angle float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("rotateOnAxis", axis, angle)}
}
func (ah *AxesHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("rotateOnWorldAxis", axis, angle)}
}
func (ah *AxesHelper) RotateX(angle float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("rotateX", angle)}
}
func (ah *AxesHelper) RotateY(angle float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("rotateY", angle)}
}
func (ah *AxesHelper) RotateZ(angle float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("rotateZ", angle)}
}
func (ah *AxesHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ah.Call("setRotationFromAxisAngle", axis, angle)
}
func (ah *AxesHelper) SetRotationFromEuler(euler *Euler) {
	ah.Call("setRotationFromEuler", euler)
}
func (ah *AxesHelper) SetRotationFromMatrix(m *Matrix4) {
	ah.Call("setRotationFromMatrix", m)
}
func (ah *AxesHelper) SetRotationFromQuaternion(q *Quaternion) {
	ah.Call("setRotationFromQuaternion", q)
}
func (ah *AxesHelper) ToJSON(meta js.Value) js.Value {
	return ah.Call("toJSON", meta)
}
func (ah *AxesHelper) TranslateOnAxis(axis *Vector3, distance float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("translateOnAxis", axis, distance)}
}
func (ah *AxesHelper) TranslateX(distance float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("translateX", distance)}
}
func (ah *AxesHelper) TranslateY(distance float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("translateY", distance)}
}
func (ah *AxesHelper) TranslateZ(distance float64) *AxesHelper {
	return &AxesHelper{Value: ah.Call("translateZ", distance)}
}
func (ah *AxesHelper) Traverse(callback js.Value) {
	ah.Call("traverse", callback)
}
func (ah *AxesHelper) TraverseAncestors(callback js.Value) {
	ah.Call("traverseAncestors", callback)
}
func (ah *AxesHelper) TraverseVisible(callback js.Value) {
	ah.Call("traverseVisible", callback)
}
func (ah *AxesHelper) UpdateMatrix() {
	ah.Call("updateMatrix")
}
func (ah *AxesHelper) UpdateMatrixWorld(force bool) {
	ah.Call("updateMatrixWorld", force)
}
func (ah *AxesHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ah.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ah *AxesHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("worldToLocal", vector)}
}

// Code generated from "helpers/BoxHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type BoxHelper struct {
	js.Value
}

func NewBoxHelper(object *Object3D, color *Color) *BoxHelper {
	return &BoxHelper{Value: get("BoxHelper").New(object, color)}
}
func (bh *BoxHelper) CastShadow() bool {
	return bh.Get("castShadow").Bool()
}
func (bh *BoxHelper) SetCastShadow(v bool) {
	bh.Set("castShadow", v)
}
func (bh *BoxHelper) Children() js.Value {
	return bh.Get("children")
}
func (bh *BoxHelper) SetChildren(v js.Value) {
	bh.Set("children", v)
}
func (bh *BoxHelper) FrustumCulled() bool {
	return bh.Get("frustumCulled").Bool()
}
func (bh *BoxHelper) SetFrustumCulled(v bool) {
	bh.Set("frustumCulled", v)
}
func (bh *BoxHelper) Geometry() *Geometry {
	return &Geometry{Value: bh.Get("geometry")}
}
func (bh *BoxHelper) SetGeometry(v *Geometry) {
	bh.Set("geometry", v)
}
func (bh *BoxHelper) Id() int {
	return bh.Get("id").Int()
}
func (bh *BoxHelper) SetId(v int) {
	bh.Set("id", v)
}
func (bh *BoxHelper) IsLine() bool {
	return bh.Get("isLine").Bool()
}
func (bh *BoxHelper) SetIsLine(v bool) {
	bh.Set("isLine", v)
}
func (bh *BoxHelper) IsLineSegments() bool {
	return bh.Get("isLineSegments").Bool()
}
func (bh *BoxHelper) SetIsLineSegments(v bool) {
	bh.Set("isLineSegments", v)
}
func (bh *BoxHelper) IsObject3D() bool {
	return bh.Get("isObject3D").Bool()
}
func (bh *BoxHelper) SetIsObject3D(v bool) {
	bh.Set("isObject3D", v)
}
func (bh *BoxHelper) Layers() *Layers {
	return &Layers{Value: bh.Get("layers")}
}
func (bh *BoxHelper) SetLayers(v *Layers) {
	bh.Set("layers", v)
}
func (bh *BoxHelper) Material() *Material {
	return &Material{Value: bh.Get("material")}
}
func (bh *BoxHelper) SetMaterial(v *Material) {
	bh.Set("material", v)
}
func (bh *BoxHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: bh.Get("matrix")}
}
func (bh *BoxHelper) SetMatrix(v *Matrix4) {
	bh.Set("matrix", v)
}
func (bh *BoxHelper) MatrixAutoUpdate() bool {
	return bh.Get("matrixAutoUpdate").Bool()
}
func (bh *BoxHelper) SetMatrixAutoUpdate(v bool) {
	bh.Set("matrixAutoUpdate", v)
}
func (bh *BoxHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: bh.Get("matrixWorld")}
}
func (bh *BoxHelper) SetMatrixWorld(v *Matrix4) {
	bh.Set("matrixWorld", v)
}
func (bh *BoxHelper) MatrixWorldNeedsUpdate() bool {
	return bh.Get("matrixWorldNeedsUpdate").Bool()
}
func (bh *BoxHelper) SetMatrixWorldNeedsUpdate(v bool) {
	bh.Set("matrixWorldNeedsUpdate", v)
}
func (bh *BoxHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: bh.Get("modelViewMatrix")}
}
func (bh *BoxHelper) SetModelViewMatrix(v *Matrix4) {
	bh.Set("modelViewMatrix", v)
}
func (bh *BoxHelper) Name() string {
	return bh.Get("name").String()
}
func (bh *BoxHelper) SetName(v string) {
	bh.Set("name", v)
}
func (bh *BoxHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: bh.Get("normalMatrix")}
}
func (bh *BoxHelper) SetNormalMatrix(v *Matrix3) {
	bh.Set("normalMatrix", v)
}
func (bh *BoxHelper) OnAfterRender() js.Value {
	return bh.Get("onAfterRender")
}
func (bh *BoxHelper) SetOnAfterRender(v js.Value) {
	bh.Set("onAfterRender", v)
}
func (bh *BoxHelper) OnBeforeRender() js.Value {
	return bh.Get("onBeforeRender")
}
func (bh *BoxHelper) SetOnBeforeRender(v js.Value) {
	bh.Set("onBeforeRender", v)
}
func (bh *BoxHelper) Parent() *Object3D {
	return &Object3D{Value: bh.Get("parent")}
}
func (bh *BoxHelper) SetParent(v *Object3D) {
	bh.Set("parent", v)
}
func (bh *BoxHelper) Position() *Vector3 {
	return &Vector3{Value: bh.Get("position")}
}
func (bh *BoxHelper) SetPosition(v *Vector3) {
	bh.Set("position", v)
}
func (bh *BoxHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: bh.Get("quaternion")}
}
func (bh *BoxHelper) SetQuaternion(v *Quaternion) {
	bh.Set("quaternion", v)
}
func (bh *BoxHelper) ReceiveShadow() bool {
	return bh.Get("receiveShadow").Bool()
}
func (bh *BoxHelper) SetReceiveShadow(v bool) {
	bh.Set("receiveShadow", v)
}
func (bh *BoxHelper) RenderOrder() float64 {
	return bh.Get("renderOrder").Float()
}
func (bh *BoxHelper) SetRenderOrder(v float64) {
	bh.Set("renderOrder", v)
}
func (bh *BoxHelper) Rotation() *Euler {
	return &Euler{Value: bh.Get("rotation")}
}
func (bh *BoxHelper) SetRotation(v *Euler) {
	bh.Set("rotation", v)
}
func (bh *BoxHelper) Scale() *Vector3 {
	return &Vector3{Value: bh.Get("scale")}
}
func (bh *BoxHelper) SetScale(v *Vector3) {
	bh.Set("scale", v)
}
func (bh *BoxHelper) Type() string {
	return bh.Get("type").String()
}
func (bh *BoxHelper) SetType(v string) {
	bh.Set("type", v)
}
func (bh *BoxHelper) Up() *Vector3 {
	return &Vector3{Value: bh.Get("up")}
}
func (bh *BoxHelper) SetUp(v *Vector3) {
	bh.Set("up", v)
}
func (bh *BoxHelper) UserData() js.Value {
	return bh.Get("userData")
}
func (bh *BoxHelper) SetUserData(v js.Value) {
	bh.Set("userData", v)
}
func (bh *BoxHelper) Uuid() string {
	return bh.Get("uuid").String()
}
func (bh *BoxHelper) SetUuid(v string) {
	bh.Set("uuid", v)
}
func (bh *BoxHelper) Visible() bool {
	return bh.Get("visible").Bool()
}
func (bh *BoxHelper) SetVisible(v bool) {
	bh.Set("visible", v)
}
func (bh *BoxHelper) DefaultMatrixAutoUpdate() bool {
	return bh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (bh *BoxHelper) SetDefaultMatrixAutoUpdate(v bool) {
	bh.Set("DefaultMatrixAutoUpdate", v)
}
func (bh *BoxHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: bh.Get("DefaultUp")}
}
func (bh *BoxHelper) SetDefaultUp(v *Vector3) {
	bh.Set("DefaultUp", v)
}
func (bh *BoxHelper) Add(object js.Value) *BoxHelper {
	return &BoxHelper{Value: bh.Call("add", object)}
}
func (bh *BoxHelper) AddEventListener(typ string, listener js.Value) {
	bh.Call("addEventListener", typ, listener)
}
func (bh *BoxHelper) ApplyMatrix(matrix *Matrix4) {
	bh.Call("applyMatrix", matrix)
}
func (bh *BoxHelper) ApplyQuaternion(quaternion *Quaternion) *BoxHelper {
	return &BoxHelper{Value: bh.Call("applyQuaternion", quaternion)}
}
func (bh *BoxHelper) Clone(recursive bool) *BoxHelper {
	return &BoxHelper{Value: bh.Call("clone", recursive)}
}
func (bh *BoxHelper) ComputeLineDistances() *BoxHelper {
	return &BoxHelper{Value: bh.Call("computeLineDistances")}
}
func (bh *BoxHelper) Copy(source *Object3D, recursive bool) *BoxHelper {
	return &BoxHelper{Value: bh.Call("copy", source, recursive)}
}
func (bh *BoxHelper) DispatchEvent(event js.Value) {
	bh.Call("dispatchEvent", event)
}
func (bh *BoxHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: bh.Call("getObjectById", id)}
}
func (bh *BoxHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: bh.Call("getObjectByName", name)}
}
func (bh *BoxHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: bh.Call("getObjectByProperty", name, value)}
}
func (bh *BoxHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: bh.Call("getWorldDirection", target)}
}
func (bh *BoxHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: bh.Call("getWorldPosition", target)}
}
func (bh *BoxHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: bh.Call("getWorldQuaternion", target)}
}
func (bh *BoxHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: bh.Call("getWorldScale", target)}
}
func (bh *BoxHelper) HasEventListener(typ string, listener js.Value) bool {
	return bh.Call("hasEventListener", typ, listener).Bool()
}
func (bh *BoxHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: bh.Call("localToWorld", vector)}
}
func (bh *BoxHelper) LookAt(vector *Vector3, y float64, z float64) {
	bh.Call("lookAt", vector, y, z)
}
func (bh *BoxHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	bh.Call("raycast", raycaster, intersects)
}
func (bh *BoxHelper) Remove(object js.Value) *BoxHelper {
	return &BoxHelper{Value: bh.Call("remove", object)}
}
func (bh *BoxHelper) RemoveEventListener(typ string, listener js.Value) {
	bh.Call("removeEventListener", typ, listener)
}
func (bh *BoxHelper) RotateOnAxis(axis *Vector3, angle float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("rotateOnAxis", axis, angle)}
}
func (bh *BoxHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("rotateOnWorldAxis", axis, angle)}
}
func (bh *BoxHelper) RotateX(angle float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("rotateX", angle)}
}
func (bh *BoxHelper) RotateY(angle float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("rotateY", angle)}
}
func (bh *BoxHelper) RotateZ(angle float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("rotateZ", angle)}
}
func (bh *BoxHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	bh.Call("setRotationFromAxisAngle", axis, angle)
}
func (bh *BoxHelper) SetRotationFromEuler(euler *Euler) {
	bh.Call("setRotationFromEuler", euler)
}
func (bh *BoxHelper) SetRotationFromMatrix(m *Matrix4) {
	bh.Call("setRotationFromMatrix", m)
}
func (bh *BoxHelper) SetRotationFromQuaternion(q *Quaternion) {
	bh.Call("setRotationFromQuaternion", q)
}
func (bh *BoxHelper) ToJSON(meta js.Value) js.Value {
	return bh.Call("toJSON", meta)
}
func (bh *BoxHelper) TranslateOnAxis(axis *Vector3, distance float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("translateOnAxis", axis, distance)}
}
func (bh *BoxHelper) TranslateX(distance float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("translateX", distance)}
}
func (bh *BoxHelper) TranslateY(distance float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("translateY", distance)}
}
func (bh *BoxHelper) TranslateZ(distance float64) *BoxHelper {
	return &BoxHelper{Value: bh.Call("translateZ", distance)}
}
func (bh *BoxHelper) Traverse(callback js.Value) {
	bh.Call("traverse", callback)
}
func (bh *BoxHelper) TraverseAncestors(callback js.Value) {
	bh.Call("traverseAncestors", callback)
}
func (bh *BoxHelper) TraverseVisible(callback js.Value) {
	bh.Call("traverseVisible", callback)
}
func (bh *BoxHelper) Update(object *Object3D) {
	bh.Call("update", object)
}
func (bh *BoxHelper) UpdateMatrix() {
	bh.Call("updateMatrix")
}
func (bh *BoxHelper) UpdateMatrixWorld(force bool) {
	bh.Call("updateMatrixWorld", force)
}
func (bh *BoxHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	bh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (bh *BoxHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: bh.Call("worldToLocal", vector)}
}

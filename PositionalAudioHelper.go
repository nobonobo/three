// Code generated from "helpers/PositionalAudioHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PositionalAudioHelper extend: [Line]
type PositionalAudioHelper struct {
	js.Value
}

func NewPositionalAudioHelper(audio *PositionalAudio, rng float64, divisionsInnerAngle float64, divisionsOuterAngle float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: get("PositionalAudioHelper").New(audio.JSValue(), rng, divisionsInnerAngle, divisionsOuterAngle)}
}
func (pah *PositionalAudioHelper) JSValue() js.Value {
	return pah.Value
}
func (pah *PositionalAudioHelper) Audio() *PositionalAudio {
	return &PositionalAudio{Value: pah.Get("audio")}
}
func (pah *PositionalAudioHelper) SetAudio(v *PositionalAudio) {
	pah.Set("audio", v.JSValue())
}
func (pah *PositionalAudioHelper) CastShadow() bool {
	return pah.Get("castShadow").Bool()
}
func (pah *PositionalAudioHelper) SetCastShadow(v bool) {
	pah.Set("castShadow", v)
}
func (pah *PositionalAudioHelper) Children() js.Value {
	return pah.Get("children")
}
func (pah *PositionalAudioHelper) SetChildren(v js.Value) {
	pah.Set("children", v)
}
func (pah *PositionalAudioHelper) DivisionsInnerAngle() float64 {
	return pah.Get("divisionsInnerAngle").Float()
}
func (pah *PositionalAudioHelper) SetDivisionsInnerAngle(v float64) {
	pah.Set("divisionsInnerAngle", v)
}
func (pah *PositionalAudioHelper) DivisionsOuterAngle() float64 {
	return pah.Get("divisionsOuterAngle").Float()
}
func (pah *PositionalAudioHelper) SetDivisionsOuterAngle(v float64) {
	pah.Set("divisionsOuterAngle", v)
}
func (pah *PositionalAudioHelper) FrustumCulled() bool {
	return pah.Get("frustumCulled").Bool()
}
func (pah *PositionalAudioHelper) SetFrustumCulled(v bool) {
	pah.Set("frustumCulled", v)
}
func (pah *PositionalAudioHelper) Geometry() Geometry {
	return &GeometryImpl{Value: pah.Get("geometry")}
}
func (pah *PositionalAudioHelper) SetGeometry(v Geometry) {
	pah.Set("geometry", v.JSValue())
}
func (pah *PositionalAudioHelper) Id() int {
	return pah.Get("id").Int()
}
func (pah *PositionalAudioHelper) SetId(v int) {
	pah.Set("id", v)
}
func (pah *PositionalAudioHelper) IsLine() bool {
	return pah.Get("isLine").Bool()
}
func (pah *PositionalAudioHelper) SetIsLine(v bool) {
	pah.Set("isLine", v)
}
func (pah *PositionalAudioHelper) IsObject3D() bool {
	return pah.Get("isObject3D").Bool()
}
func (pah *PositionalAudioHelper) SetIsObject3D(v bool) {
	pah.Set("isObject3D", v)
}
func (pah *PositionalAudioHelper) Layers() *Layers {
	return &Layers{Value: pah.Get("layers")}
}
func (pah *PositionalAudioHelper) SetLayers(v *Layers) {
	pah.Set("layers", v.JSValue())
}
func (pah *PositionalAudioHelper) Material() Material {
	return &MaterialImpl{Value: pah.Get("material")}
}
func (pah *PositionalAudioHelper) SetMaterial(v Material) {
	pah.Set("material", v.JSValue())
}
func (pah *PositionalAudioHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: pah.Get("matrix")}
}
func (pah *PositionalAudioHelper) SetMatrix(v *Matrix4) {
	pah.Set("matrix", v.JSValue())
}
func (pah *PositionalAudioHelper) MatrixAutoUpdate() bool {
	return pah.Get("matrixAutoUpdate").Bool()
}
func (pah *PositionalAudioHelper) SetMatrixAutoUpdate(v bool) {
	pah.Set("matrixAutoUpdate", v)
}
func (pah *PositionalAudioHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: pah.Get("matrixWorld")}
}
func (pah *PositionalAudioHelper) SetMatrixWorld(v *Matrix4) {
	pah.Set("matrixWorld", v.JSValue())
}
func (pah *PositionalAudioHelper) MatrixWorldNeedsUpdate() bool {
	return pah.Get("matrixWorldNeedsUpdate").Bool()
}
func (pah *PositionalAudioHelper) SetMatrixWorldNeedsUpdate(v bool) {
	pah.Set("matrixWorldNeedsUpdate", v)
}
func (pah *PositionalAudioHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: pah.Get("modelViewMatrix")}
}
func (pah *PositionalAudioHelper) SetModelViewMatrix(v *Matrix4) {
	pah.Set("modelViewMatrix", v.JSValue())
}
func (pah *PositionalAudioHelper) Name() string {
	return pah.Get("name").String()
}
func (pah *PositionalAudioHelper) SetName(v string) {
	pah.Set("name", v)
}
func (pah *PositionalAudioHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: pah.Get("normalMatrix")}
}
func (pah *PositionalAudioHelper) SetNormalMatrix(v *Matrix3) {
	pah.Set("normalMatrix", v.JSValue())
}
func (pah *PositionalAudioHelper) OnAfterRender() js.Value {
	return pah.Get("onAfterRender")
}
func (pah *PositionalAudioHelper) SetOnAfterRender(v js.Value) {
	pah.Set("onAfterRender", v)
}
func (pah *PositionalAudioHelper) OnBeforeRender() js.Value {
	return pah.Get("onBeforeRender")
}
func (pah *PositionalAudioHelper) SetOnBeforeRender(v js.Value) {
	pah.Set("onBeforeRender", v)
}
func (pah *PositionalAudioHelper) Parent() *Object3D {
	return &Object3D{Value: pah.Get("parent")}
}
func (pah *PositionalAudioHelper) SetParent(v *Object3D) {
	pah.Set("parent", v.JSValue())
}
func (pah *PositionalAudioHelper) Position() *Vector3 {
	return &Vector3{Value: pah.Get("position")}
}
func (pah *PositionalAudioHelper) SetPosition(v *Vector3) {
	pah.Set("position", v.JSValue())
}
func (pah *PositionalAudioHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: pah.Get("quaternion")}
}
func (pah *PositionalAudioHelper) SetQuaternion(v *Quaternion) {
	pah.Set("quaternion", v.JSValue())
}
func (pah *PositionalAudioHelper) Range() float64 {
	return pah.Get("range").Float()
}
func (pah *PositionalAudioHelper) SetRange(v float64) {
	pah.Set("range", v)
}
func (pah *PositionalAudioHelper) ReceiveShadow() bool {
	return pah.Get("receiveShadow").Bool()
}
func (pah *PositionalAudioHelper) SetReceiveShadow(v bool) {
	pah.Set("receiveShadow", v)
}
func (pah *PositionalAudioHelper) RenderOrder() float64 {
	return pah.Get("renderOrder").Float()
}
func (pah *PositionalAudioHelper) SetRenderOrder(v float64) {
	pah.Set("renderOrder", v)
}
func (pah *PositionalAudioHelper) Rotation() *Euler {
	return &Euler{Value: pah.Get("rotation")}
}
func (pah *PositionalAudioHelper) SetRotation(v *Euler) {
	pah.Set("rotation", v.JSValue())
}
func (pah *PositionalAudioHelper) Scale() *Vector3 {
	return &Vector3{Value: pah.Get("scale")}
}
func (pah *PositionalAudioHelper) SetScale(v *Vector3) {
	pah.Set("scale", v.JSValue())
}
func (pah *PositionalAudioHelper) Type() string {
	return pah.Get("type").String()
}
func (pah *PositionalAudioHelper) SetType(v string) {
	pah.Set("type", v)
}
func (pah *PositionalAudioHelper) Up() *Vector3 {
	return &Vector3{Value: pah.Get("up")}
}
func (pah *PositionalAudioHelper) SetUp(v *Vector3) {
	pah.Set("up", v.JSValue())
}
func (pah *PositionalAudioHelper) UserData() js.Value {
	return pah.Get("userData")
}
func (pah *PositionalAudioHelper) SetUserData(v js.Value) {
	pah.Set("userData", v)
}
func (pah *PositionalAudioHelper) Uuid() string {
	return pah.Get("uuid").String()
}
func (pah *PositionalAudioHelper) SetUuid(v string) {
	pah.Set("uuid", v)
}
func (pah *PositionalAudioHelper) Visible() bool {
	return pah.Get("visible").Bool()
}
func (pah *PositionalAudioHelper) SetVisible(v bool) {
	pah.Set("visible", v)
}
func (pah *PositionalAudioHelper) DefaultMatrixAutoUpdate() bool {
	return pah.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pah *PositionalAudioHelper) SetDefaultMatrixAutoUpdate(v bool) {
	pah.Set("DefaultMatrixAutoUpdate", v)
}
func (pah *PositionalAudioHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: pah.Get("DefaultUp")}
}
func (pah *PositionalAudioHelper) SetDefaultUp(v *Vector3) {
	pah.Set("DefaultUp", v.JSValue())
}
func (pah *PositionalAudioHelper) Add(object js.Value) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("add", object)}
}
func (pah *PositionalAudioHelper) AddEventListener(typ string, listener js.Value) {
	pah.Call("addEventListener", typ, listener)
}
func (pah *PositionalAudioHelper) ApplyMatrix(matrix *Matrix4) {
	pah.Call("applyMatrix", matrix.JSValue())
}
func (pah *PositionalAudioHelper) ApplyQuaternion(quaternion *Quaternion) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("applyQuaternion", quaternion)}
}
func (pah *PositionalAudioHelper) Clone(recursive bool) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("clone", recursive)}
}
func (pah *PositionalAudioHelper) ComputeLineDistances() *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("computeLineDistances")}
}
func (pah *PositionalAudioHelper) Copy(source *Object3D, recursive bool) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("copy", source, recursive)}
}
func (pah *PositionalAudioHelper) DispatchEvent(event js.Value) {
	pah.Call("dispatchEvent", event)
}
func (pah *PositionalAudioHelper) Dispose() {
	pah.Call("dispose")
}
func (pah *PositionalAudioHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: pah.Call("getObjectById", id)}
}
func (pah *PositionalAudioHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: pah.Call("getObjectByName", name)}
}
func (pah *PositionalAudioHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: pah.Call("getObjectByProperty", name, value)}
}
func (pah *PositionalAudioHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: pah.Call("getWorldDirection", target)}
}
func (pah *PositionalAudioHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: pah.Call("getWorldPosition", target)}
}
func (pah *PositionalAudioHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: pah.Call("getWorldQuaternion", target)}
}
func (pah *PositionalAudioHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: pah.Call("getWorldScale", target)}
}
func (pah *PositionalAudioHelper) HasEventListener(typ string, listener js.Value) bool {
	return pah.Call("hasEventListener", typ, listener).Bool()
}
func (pah *PositionalAudioHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: pah.Call("localToWorld", vector)}
}
func (pah *PositionalAudioHelper) LookAt(vector *Vector3, y float64, z float64) {
	pah.Call("lookAt", vector, y, z)
}
func (pah *PositionalAudioHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	pah.Call("raycast", raycaster.JSValue(), intersects)
}
func (pah *PositionalAudioHelper) Remove(object js.Value) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("remove", object)}
}
func (pah *PositionalAudioHelper) RemoveEventListener(typ string, listener js.Value) {
	pah.Call("removeEventListener", typ, listener)
}
func (pah *PositionalAudioHelper) RotateOnAxis(axis *Vector3, angle float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("rotateOnAxis", axis, angle)}
}
func (pah *PositionalAudioHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("rotateOnWorldAxis", axis, angle)}
}
func (pah *PositionalAudioHelper) RotateX(angle float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("rotateX", angle)}
}
func (pah *PositionalAudioHelper) RotateY(angle float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("rotateY", angle)}
}
func (pah *PositionalAudioHelper) RotateZ(angle float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("rotateZ", angle)}
}
func (pah *PositionalAudioHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	pah.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (pah *PositionalAudioHelper) SetRotationFromEuler(euler *Euler) {
	pah.Call("setRotationFromEuler", euler.JSValue())
}
func (pah *PositionalAudioHelper) SetRotationFromMatrix(m *Matrix4) {
	pah.Call("setRotationFromMatrix", m.JSValue())
}
func (pah *PositionalAudioHelper) SetRotationFromQuaternion(q *Quaternion) {
	pah.Call("setRotationFromQuaternion", q.JSValue())
}
func (pah *PositionalAudioHelper) ToJSON(meta js.Value) js.Value {
	return pah.Call("toJSON", meta)
}
func (pah *PositionalAudioHelper) TranslateOnAxis(axis *Vector3, distance float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("translateOnAxis", axis, distance)}
}
func (pah *PositionalAudioHelper) TranslateX(distance float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("translateX", distance)}
}
func (pah *PositionalAudioHelper) TranslateY(distance float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("translateY", distance)}
}
func (pah *PositionalAudioHelper) TranslateZ(distance float64) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: pah.Call("translateZ", distance)}
}
func (pah *PositionalAudioHelper) Traverse(callback js.Value) {
	pah.Call("traverse", callback)
}
func (pah *PositionalAudioHelper) TraverseAncestors(callback js.Value) {
	pah.Call("traverseAncestors", callback)
}
func (pah *PositionalAudioHelper) TraverseVisible(callback js.Value) {
	pah.Call("traverseVisible", callback)
}
func (pah *PositionalAudioHelper) Update() {
	pah.Call("update")
}
func (pah *PositionalAudioHelper) UpdateMatrix() {
	pah.Call("updateMatrix")
}
func (pah *PositionalAudioHelper) UpdateMatrixWorld(force bool) {
	pah.Call("updateMatrixWorld", force)
}
func (pah *PositionalAudioHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pah.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pah *PositionalAudioHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: pah.Call("worldToLocal", vector)}
}

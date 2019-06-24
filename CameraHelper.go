// Code generated from "helpers/CameraHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type CameraHelper struct {
	js.Value
}

func NewCameraHelper(camera *Camera) *CameraHelper {
	return &CameraHelper{Value: get("CameraHelper").New(camera)}
}
func (ch *CameraHelper) Camera() *Camera {
	return &Camera{Value: ch.Get("camera")}
}
func (ch *CameraHelper) SetCamera(v *Camera) {
	ch.Set("camera", v)
}
func (ch *CameraHelper) CastShadow() bool {
	return ch.Get("castShadow").Bool()
}
func (ch *CameraHelper) SetCastShadow(v bool) {
	ch.Set("castShadow", v)
}
func (ch *CameraHelper) Children() js.Value {
	return ch.Get("children")
}
func (ch *CameraHelper) SetChildren(v js.Value) {
	ch.Set("children", v)
}
func (ch *CameraHelper) FrustumCulled() bool {
	return ch.Get("frustumCulled").Bool()
}
func (ch *CameraHelper) SetFrustumCulled(v bool) {
	ch.Set("frustumCulled", v)
}
func (ch *CameraHelper) Geometry() *Geometry {
	return &Geometry{Value: ch.Get("geometry")}
}
func (ch *CameraHelper) SetGeometry(v *Geometry) {
	ch.Set("geometry", v)
}
func (ch *CameraHelper) Id() int {
	return ch.Get("id").Int()
}
func (ch *CameraHelper) SetId(v int) {
	ch.Set("id", v)
}
func (ch *CameraHelper) IsLine() bool {
	return ch.Get("isLine").Bool()
}
func (ch *CameraHelper) SetIsLine(v bool) {
	ch.Set("isLine", v)
}
func (ch *CameraHelper) IsLineSegments() bool {
	return ch.Get("isLineSegments").Bool()
}
func (ch *CameraHelper) SetIsLineSegments(v bool) {
	ch.Set("isLineSegments", v)
}
func (ch *CameraHelper) IsObject3D() bool {
	return ch.Get("isObject3D").Bool()
}
func (ch *CameraHelper) SetIsObject3D(v bool) {
	ch.Set("isObject3D", v)
}
func (ch *CameraHelper) Layers() *Layers {
	return &Layers{Value: ch.Get("layers")}
}
func (ch *CameraHelper) SetLayers(v *Layers) {
	ch.Set("layers", v)
}
func (ch *CameraHelper) Material() *Material {
	return &Material{Value: ch.Get("material")}
}
func (ch *CameraHelper) SetMaterial(v *Material) {
	ch.Set("material", v)
}
func (ch *CameraHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: ch.Get("matrix")}
}
func (ch *CameraHelper) SetMatrix(v *Matrix4) {
	ch.Set("matrix", v)
}
func (ch *CameraHelper) MatrixAutoUpdate() bool {
	return ch.Get("matrixAutoUpdate").Bool()
}
func (ch *CameraHelper) SetMatrixAutoUpdate(v bool) {
	ch.Set("matrixAutoUpdate", v)
}
func (ch *CameraHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ch.Get("matrixWorld")}
}
func (ch *CameraHelper) SetMatrixWorld(v *Matrix4) {
	ch.Set("matrixWorld", v)
}
func (ch *CameraHelper) MatrixWorldNeedsUpdate() bool {
	return ch.Get("matrixWorldNeedsUpdate").Bool()
}
func (ch *CameraHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ch.Set("matrixWorldNeedsUpdate", v)
}
func (ch *CameraHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ch.Get("modelViewMatrix")}
}
func (ch *CameraHelper) SetModelViewMatrix(v *Matrix4) {
	ch.Set("modelViewMatrix", v)
}
func (ch *CameraHelper) Name() string {
	return ch.Get("name").String()
}
func (ch *CameraHelper) SetName(v string) {
	ch.Set("name", v)
}
func (ch *CameraHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ch.Get("normalMatrix")}
}
func (ch *CameraHelper) SetNormalMatrix(v *Matrix3) {
	ch.Set("normalMatrix", v)
}
func (ch *CameraHelper) OnAfterRender() js.Value {
	return ch.Get("onAfterRender")
}
func (ch *CameraHelper) SetOnAfterRender(v js.Value) {
	ch.Set("onAfterRender", v)
}
func (ch *CameraHelper) OnBeforeRender() js.Value {
	return ch.Get("onBeforeRender")
}
func (ch *CameraHelper) SetOnBeforeRender(v js.Value) {
	ch.Set("onBeforeRender", v)
}
func (ch *CameraHelper) Parent() *Object3D {
	return &Object3D{Value: ch.Get("parent")}
}
func (ch *CameraHelper) SetParent(v *Object3D) {
	ch.Set("parent", v)
}
func (ch *CameraHelper) PointMap() js.Value {
	return ch.Get("pointMap")
}
func (ch *CameraHelper) SetPointMap(v js.Value) {
	ch.Set("pointMap", v)
}
func (ch *CameraHelper) Position() *Vector3 {
	return &Vector3{Value: ch.Get("position")}
}
func (ch *CameraHelper) SetPosition(v *Vector3) {
	ch.Set("position", v)
}
func (ch *CameraHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: ch.Get("quaternion")}
}
func (ch *CameraHelper) SetQuaternion(v *Quaternion) {
	ch.Set("quaternion", v)
}
func (ch *CameraHelper) ReceiveShadow() bool {
	return ch.Get("receiveShadow").Bool()
}
func (ch *CameraHelper) SetReceiveShadow(v bool) {
	ch.Set("receiveShadow", v)
}
func (ch *CameraHelper) RenderOrder() float64 {
	return ch.Get("renderOrder").Float()
}
func (ch *CameraHelper) SetRenderOrder(v float64) {
	ch.Set("renderOrder", v)
}
func (ch *CameraHelper) Rotation() *Euler {
	return &Euler{Value: ch.Get("rotation")}
}
func (ch *CameraHelper) SetRotation(v *Euler) {
	ch.Set("rotation", v)
}
func (ch *CameraHelper) Scale() *Vector3 {
	return &Vector3{Value: ch.Get("scale")}
}
func (ch *CameraHelper) SetScale(v *Vector3) {
	ch.Set("scale", v)
}
func (ch *CameraHelper) Type() string {
	return ch.Get("type").String()
}
func (ch *CameraHelper) SetType(v string) {
	ch.Set("type", v)
}
func (ch *CameraHelper) Up() *Vector3 {
	return &Vector3{Value: ch.Get("up")}
}
func (ch *CameraHelper) SetUp(v *Vector3) {
	ch.Set("up", v)
}
func (ch *CameraHelper) UserData() js.Value {
	return ch.Get("userData")
}
func (ch *CameraHelper) SetUserData(v js.Value) {
	ch.Set("userData", v)
}
func (ch *CameraHelper) Uuid() string {
	return ch.Get("uuid").String()
}
func (ch *CameraHelper) SetUuid(v string) {
	ch.Set("uuid", v)
}
func (ch *CameraHelper) Visible() bool {
	return ch.Get("visible").Bool()
}
func (ch *CameraHelper) SetVisible(v bool) {
	ch.Set("visible", v)
}
func (ch *CameraHelper) DefaultMatrixAutoUpdate() bool {
	return ch.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ch *CameraHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ch.Set("DefaultMatrixAutoUpdate", v)
}
func (ch *CameraHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: ch.Get("DefaultUp")}
}
func (ch *CameraHelper) SetDefaultUp(v *Vector3) {
	ch.Set("DefaultUp", v)
}
func (ch *CameraHelper) Add(object js.Value) *CameraHelper {
	return &CameraHelper{Value: ch.Call("add", object)}
}
func (ch *CameraHelper) AddEventListener(typ string, listener js.Value) {
	ch.Call("addEventListener", typ, listener)
}
func (ch *CameraHelper) ApplyMatrix(matrix *Matrix4) {
	ch.Call("applyMatrix", matrix)
}
func (ch *CameraHelper) ApplyQuaternion(quaternion *Quaternion) *CameraHelper {
	return &CameraHelper{Value: ch.Call("applyQuaternion", quaternion)}
}
func (ch *CameraHelper) Clone(recursive bool) *CameraHelper {
	return &CameraHelper{Value: ch.Call("clone", recursive)}
}
func (ch *CameraHelper) ComputeLineDistances() *CameraHelper {
	return &CameraHelper{Value: ch.Call("computeLineDistances")}
}
func (ch *CameraHelper) Copy(source *Object3D, recursive bool) *CameraHelper {
	return &CameraHelper{Value: ch.Call("copy", source, recursive)}
}
func (ch *CameraHelper) DispatchEvent(event js.Value) {
	ch.Call("dispatchEvent", event)
}
func (ch *CameraHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ch.Call("getObjectById", id)}
}
func (ch *CameraHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ch.Call("getObjectByName", name)}
}
func (ch *CameraHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ch.Call("getObjectByProperty", name, value)}
}
func (ch *CameraHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ch.Call("getWorldDirection", target)}
}
func (ch *CameraHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ch.Call("getWorldPosition", target)}
}
func (ch *CameraHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ch.Call("getWorldQuaternion", target)}
}
func (ch *CameraHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ch.Call("getWorldScale", target)}
}
func (ch *CameraHelper) HasEventListener(typ string, listener js.Value) bool {
	return ch.Call("hasEventListener", typ, listener).Bool()
}
func (ch *CameraHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ch.Call("localToWorld", vector)}
}
func (ch *CameraHelper) LookAt(vector *Vector3, y float64, z float64) {
	ch.Call("lookAt", vector, y, z)
}
func (ch *CameraHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	ch.Call("raycast", raycaster, intersects)
}
func (ch *CameraHelper) Remove(object js.Value) *CameraHelper {
	return &CameraHelper{Value: ch.Call("remove", object)}
}
func (ch *CameraHelper) RemoveEventListener(typ string, listener js.Value) {
	ch.Call("removeEventListener", typ, listener)
}
func (ch *CameraHelper) RotateOnAxis(axis *Vector3, angle float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("rotateOnAxis", axis, angle)}
}
func (ch *CameraHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("rotateOnWorldAxis", axis, angle)}
}
func (ch *CameraHelper) RotateX(angle float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("rotateX", angle)}
}
func (ch *CameraHelper) RotateY(angle float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("rotateY", angle)}
}
func (ch *CameraHelper) RotateZ(angle float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("rotateZ", angle)}
}
func (ch *CameraHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ch.Call("setRotationFromAxisAngle", axis, angle)
}
func (ch *CameraHelper) SetRotationFromEuler(euler *Euler) {
	ch.Call("setRotationFromEuler", euler)
}
func (ch *CameraHelper) SetRotationFromMatrix(m *Matrix4) {
	ch.Call("setRotationFromMatrix", m)
}
func (ch *CameraHelper) SetRotationFromQuaternion(q *Quaternion) {
	ch.Call("setRotationFromQuaternion", q)
}
func (ch *CameraHelper) ToJSON(meta js.Value) js.Value {
	return ch.Call("toJSON", meta)
}
func (ch *CameraHelper) TranslateOnAxis(axis *Vector3, distance float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("translateOnAxis", axis, distance)}
}
func (ch *CameraHelper) TranslateX(distance float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("translateX", distance)}
}
func (ch *CameraHelper) TranslateY(distance float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("translateY", distance)}
}
func (ch *CameraHelper) TranslateZ(distance float64) *CameraHelper {
	return &CameraHelper{Value: ch.Call("translateZ", distance)}
}
func (ch *CameraHelper) Traverse(callback js.Value) {
	ch.Call("traverse", callback)
}
func (ch *CameraHelper) TraverseAncestors(callback js.Value) {
	ch.Call("traverseAncestors", callback)
}
func (ch *CameraHelper) TraverseVisible(callback js.Value) {
	ch.Call("traverseVisible", callback)
}
func (ch *CameraHelper) Update() {
	ch.Call("update")
}
func (ch *CameraHelper) UpdateMatrix() {
	ch.Call("updateMatrix")
}
func (ch *CameraHelper) UpdateMatrixWorld(force bool) {
	ch.Call("updateMatrixWorld", force)
}
func (ch *CameraHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ch.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ch *CameraHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ch.Call("worldToLocal", vector)}
}

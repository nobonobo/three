// Code generated from "helpers/FaceNormalsHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type FaceNormalsHelper struct {
	js.Value
}

func NewFaceNormalsHelper(object *Object3D, size int, hex int, linewidth float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: get("FaceNormalsHelper").New(object, size, hex, linewidth)}
}
func (fnh *FaceNormalsHelper) CastShadow() bool {
	return fnh.Get("castShadow").Bool()
}
func (fnh *FaceNormalsHelper) SetCastShadow(v bool) {
	fnh.Set("castShadow", v)
}
func (fnh *FaceNormalsHelper) Children() js.Value {
	return fnh.Get("children")
}
func (fnh *FaceNormalsHelper) SetChildren(v js.Value) {
	fnh.Set("children", v)
}
func (fnh *FaceNormalsHelper) FrustumCulled() bool {
	return fnh.Get("frustumCulled").Bool()
}
func (fnh *FaceNormalsHelper) SetFrustumCulled(v bool) {
	fnh.Set("frustumCulled", v)
}
func (fnh *FaceNormalsHelper) Geometry() *Geometry {
	return &Geometry{Value: fnh.Get("geometry")}
}
func (fnh *FaceNormalsHelper) SetGeometry(v *Geometry) {
	fnh.Set("geometry", v)
}
func (fnh *FaceNormalsHelper) Id() int {
	return fnh.Get("id").Int()
}
func (fnh *FaceNormalsHelper) SetId(v int) {
	fnh.Set("id", v)
}
func (fnh *FaceNormalsHelper) IsLine() bool {
	return fnh.Get("isLine").Bool()
}
func (fnh *FaceNormalsHelper) SetIsLine(v bool) {
	fnh.Set("isLine", v)
}
func (fnh *FaceNormalsHelper) IsLineSegments() bool {
	return fnh.Get("isLineSegments").Bool()
}
func (fnh *FaceNormalsHelper) SetIsLineSegments(v bool) {
	fnh.Set("isLineSegments", v)
}
func (fnh *FaceNormalsHelper) IsObject3D() bool {
	return fnh.Get("isObject3D").Bool()
}
func (fnh *FaceNormalsHelper) SetIsObject3D(v bool) {
	fnh.Set("isObject3D", v)
}
func (fnh *FaceNormalsHelper) Layers() *Layers {
	return &Layers{Value: fnh.Get("layers")}
}
func (fnh *FaceNormalsHelper) SetLayers(v *Layers) {
	fnh.Set("layers", v)
}
func (fnh *FaceNormalsHelper) Material() *Material {
	return &Material{Value: fnh.Get("material")}
}
func (fnh *FaceNormalsHelper) SetMaterial(v *Material) {
	fnh.Set("material", v)
}
func (fnh *FaceNormalsHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: fnh.Get("matrix")}
}
func (fnh *FaceNormalsHelper) SetMatrix(v *Matrix4) {
	fnh.Set("matrix", v)
}
func (fnh *FaceNormalsHelper) MatrixAutoUpdate() bool {
	return fnh.Get("matrixAutoUpdate").Bool()
}
func (fnh *FaceNormalsHelper) SetMatrixAutoUpdate(v bool) {
	fnh.Set("matrixAutoUpdate", v)
}
func (fnh *FaceNormalsHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: fnh.Get("matrixWorld")}
}
func (fnh *FaceNormalsHelper) SetMatrixWorld(v *Matrix4) {
	fnh.Set("matrixWorld", v)
}
func (fnh *FaceNormalsHelper) MatrixWorldNeedsUpdate() bool {
	return fnh.Get("matrixWorldNeedsUpdate").Bool()
}
func (fnh *FaceNormalsHelper) SetMatrixWorldNeedsUpdate(v bool) {
	fnh.Set("matrixWorldNeedsUpdate", v)
}
func (fnh *FaceNormalsHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: fnh.Get("modelViewMatrix")}
}
func (fnh *FaceNormalsHelper) SetModelViewMatrix(v *Matrix4) {
	fnh.Set("modelViewMatrix", v)
}
func (fnh *FaceNormalsHelper) Name() string {
	return fnh.Get("name").String()
}
func (fnh *FaceNormalsHelper) SetName(v string) {
	fnh.Set("name", v)
}
func (fnh *FaceNormalsHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: fnh.Get("normalMatrix")}
}
func (fnh *FaceNormalsHelper) SetNormalMatrix(v *Matrix3) {
	fnh.Set("normalMatrix", v)
}
func (fnh *FaceNormalsHelper) Object() *Object3D {
	return &Object3D{Value: fnh.Get("object")}
}
func (fnh *FaceNormalsHelper) SetObject(v *Object3D) {
	fnh.Set("object", v)
}
func (fnh *FaceNormalsHelper) OnAfterRender() js.Value {
	return fnh.Get("onAfterRender")
}
func (fnh *FaceNormalsHelper) SetOnAfterRender(v js.Value) {
	fnh.Set("onAfterRender", v)
}
func (fnh *FaceNormalsHelper) OnBeforeRender() js.Value {
	return fnh.Get("onBeforeRender")
}
func (fnh *FaceNormalsHelper) SetOnBeforeRender(v js.Value) {
	fnh.Set("onBeforeRender", v)
}
func (fnh *FaceNormalsHelper) Parent() *Object3D {
	return &Object3D{Value: fnh.Get("parent")}
}
func (fnh *FaceNormalsHelper) SetParent(v *Object3D) {
	fnh.Set("parent", v)
}
func (fnh *FaceNormalsHelper) Position() *Vector3 {
	return &Vector3{Value: fnh.Get("position")}
}
func (fnh *FaceNormalsHelper) SetPosition(v *Vector3) {
	fnh.Set("position", v)
}
func (fnh *FaceNormalsHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: fnh.Get("quaternion")}
}
func (fnh *FaceNormalsHelper) SetQuaternion(v *Quaternion) {
	fnh.Set("quaternion", v)
}
func (fnh *FaceNormalsHelper) ReceiveShadow() bool {
	return fnh.Get("receiveShadow").Bool()
}
func (fnh *FaceNormalsHelper) SetReceiveShadow(v bool) {
	fnh.Set("receiveShadow", v)
}
func (fnh *FaceNormalsHelper) RenderOrder() float64 {
	return fnh.Get("renderOrder").Float()
}
func (fnh *FaceNormalsHelper) SetRenderOrder(v float64) {
	fnh.Set("renderOrder", v)
}
func (fnh *FaceNormalsHelper) Rotation() *Euler {
	return &Euler{Value: fnh.Get("rotation")}
}
func (fnh *FaceNormalsHelper) SetRotation(v *Euler) {
	fnh.Set("rotation", v)
}
func (fnh *FaceNormalsHelper) Scale() *Vector3 {
	return &Vector3{Value: fnh.Get("scale")}
}
func (fnh *FaceNormalsHelper) SetScale(v *Vector3) {
	fnh.Set("scale", v)
}
func (fnh *FaceNormalsHelper) Size() float64 {
	return fnh.Get("size").Float()
}
func (fnh *FaceNormalsHelper) SetSize(v float64) {
	fnh.Set("size", v)
}
func (fnh *FaceNormalsHelper) Type() string {
	return fnh.Get("type").String()
}
func (fnh *FaceNormalsHelper) SetType(v string) {
	fnh.Set("type", v)
}
func (fnh *FaceNormalsHelper) Up() *Vector3 {
	return &Vector3{Value: fnh.Get("up")}
}
func (fnh *FaceNormalsHelper) SetUp(v *Vector3) {
	fnh.Set("up", v)
}
func (fnh *FaceNormalsHelper) UserData() js.Value {
	return fnh.Get("userData")
}
func (fnh *FaceNormalsHelper) SetUserData(v js.Value) {
	fnh.Set("userData", v)
}
func (fnh *FaceNormalsHelper) Uuid() string {
	return fnh.Get("uuid").String()
}
func (fnh *FaceNormalsHelper) SetUuid(v string) {
	fnh.Set("uuid", v)
}
func (fnh *FaceNormalsHelper) Visible() bool {
	return fnh.Get("visible").Bool()
}
func (fnh *FaceNormalsHelper) SetVisible(v bool) {
	fnh.Set("visible", v)
}
func (fnh *FaceNormalsHelper) DefaultMatrixAutoUpdate() bool {
	return fnh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (fnh *FaceNormalsHelper) SetDefaultMatrixAutoUpdate(v bool) {
	fnh.Set("DefaultMatrixAutoUpdate", v)
}
func (fnh *FaceNormalsHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: fnh.Get("DefaultUp")}
}
func (fnh *FaceNormalsHelper) SetDefaultUp(v *Vector3) {
	fnh.Set("DefaultUp", v)
}
func (fnh *FaceNormalsHelper) Add(object js.Value) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("add", object)}
}
func (fnh *FaceNormalsHelper) AddEventListener(typ string, listener js.Value) {
	fnh.Call("addEventListener", typ, listener)
}
func (fnh *FaceNormalsHelper) ApplyMatrix(matrix *Matrix4) {
	fnh.Call("applyMatrix", matrix)
}
func (fnh *FaceNormalsHelper) ApplyQuaternion(quaternion *Quaternion) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("applyQuaternion", quaternion)}
}
func (fnh *FaceNormalsHelper) Clone(recursive bool) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("clone", recursive)}
}
func (fnh *FaceNormalsHelper) ComputeLineDistances() *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("computeLineDistances")}
}
func (fnh *FaceNormalsHelper) Copy(source *Object3D, recursive bool) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("copy", source, recursive)}
}
func (fnh *FaceNormalsHelper) DispatchEvent(event js.Value) {
	fnh.Call("dispatchEvent", event)
}
func (fnh *FaceNormalsHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: fnh.Call("getObjectById", id)}
}
func (fnh *FaceNormalsHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: fnh.Call("getObjectByName", name)}
}
func (fnh *FaceNormalsHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: fnh.Call("getObjectByProperty", name, value)}
}
func (fnh *FaceNormalsHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: fnh.Call("getWorldDirection", target)}
}
func (fnh *FaceNormalsHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: fnh.Call("getWorldPosition", target)}
}
func (fnh *FaceNormalsHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: fnh.Call("getWorldQuaternion", target)}
}
func (fnh *FaceNormalsHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: fnh.Call("getWorldScale", target)}
}
func (fnh *FaceNormalsHelper) HasEventListener(typ string, listener js.Value) bool {
	return fnh.Call("hasEventListener", typ, listener).Bool()
}
func (fnh *FaceNormalsHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: fnh.Call("localToWorld", vector)}
}
func (fnh *FaceNormalsHelper) LookAt(vector *Vector3, y float64, z float64) {
	fnh.Call("lookAt", vector, y, z)
}
func (fnh *FaceNormalsHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	fnh.Call("raycast", raycaster, intersects)
}
func (fnh *FaceNormalsHelper) Remove(object js.Value) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("remove", object)}
}
func (fnh *FaceNormalsHelper) RemoveEventListener(typ string, listener js.Value) {
	fnh.Call("removeEventListener", typ, listener)
}
func (fnh *FaceNormalsHelper) RotateOnAxis(axis *Vector3, angle float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("rotateOnAxis", axis, angle)}
}
func (fnh *FaceNormalsHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("rotateOnWorldAxis", axis, angle)}
}
func (fnh *FaceNormalsHelper) RotateX(angle float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("rotateX", angle)}
}
func (fnh *FaceNormalsHelper) RotateY(angle float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("rotateY", angle)}
}
func (fnh *FaceNormalsHelper) RotateZ(angle float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("rotateZ", angle)}
}
func (fnh *FaceNormalsHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	fnh.Call("setRotationFromAxisAngle", axis, angle)
}
func (fnh *FaceNormalsHelper) SetRotationFromEuler(euler *Euler) {
	fnh.Call("setRotationFromEuler", euler)
}
func (fnh *FaceNormalsHelper) SetRotationFromMatrix(m *Matrix4) {
	fnh.Call("setRotationFromMatrix", m)
}
func (fnh *FaceNormalsHelper) SetRotationFromQuaternion(q *Quaternion) {
	fnh.Call("setRotationFromQuaternion", q)
}
func (fnh *FaceNormalsHelper) ToJSON(meta js.Value) js.Value {
	return fnh.Call("toJSON", meta)
}
func (fnh *FaceNormalsHelper) TranslateOnAxis(axis *Vector3, distance float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("translateOnAxis", axis, distance)}
}
func (fnh *FaceNormalsHelper) TranslateX(distance float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("translateX", distance)}
}
func (fnh *FaceNormalsHelper) TranslateY(distance float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("translateY", distance)}
}
func (fnh *FaceNormalsHelper) TranslateZ(distance float64) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: fnh.Call("translateZ", distance)}
}
func (fnh *FaceNormalsHelper) Traverse(callback js.Value) {
	fnh.Call("traverse", callback)
}
func (fnh *FaceNormalsHelper) TraverseAncestors(callback js.Value) {
	fnh.Call("traverseAncestors", callback)
}
func (fnh *FaceNormalsHelper) TraverseVisible(callback js.Value) {
	fnh.Call("traverseVisible", callback)
}
func (fnh *FaceNormalsHelper) Update(object *Object3D) {
	fnh.Call("update", object)
}
func (fnh *FaceNormalsHelper) UpdateMatrix() {
	fnh.Call("updateMatrix")
}
func (fnh *FaceNormalsHelper) UpdateMatrixWorld(force bool) {
	fnh.Call("updateMatrixWorld", force)
}
func (fnh *FaceNormalsHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	fnh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (fnh *FaceNormalsHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: fnh.Call("worldToLocal", vector)}
}

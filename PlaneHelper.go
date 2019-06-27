// Code generated from "helpers/PlaneHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// PlaneHelper extend: [LineSegments]
type PlaneHelper struct {
	js.Value
}

func NewPlaneHelper(plane *Plane, size int, hex int) *PlaneHelper {
	return &PlaneHelper{Value: get("PlaneHelper").New(plane, size, hex)}
}
func (ph *PlaneHelper) JSValue() js.Value {
	return ph.Value
}
func (ph *PlaneHelper) CastShadow() bool {
	return ph.Get("castShadow").Bool()
}
func (ph *PlaneHelper) SetCastShadow(v bool) {
	ph.Set("castShadow", v)
}
func (ph *PlaneHelper) Children() js.Value {
	return ph.Get("children")
}
func (ph *PlaneHelper) SetChildren(v js.Value) {
	ph.Set("children", v)
}
func (ph *PlaneHelper) FrustumCulled() bool {
	return ph.Get("frustumCulled").Bool()
}
func (ph *PlaneHelper) SetFrustumCulled(v bool) {
	ph.Set("frustumCulled", v)
}
func (ph *PlaneHelper) Geometry() Geometry {
	return &GeometryImpl{Value: ph.Get("geometry")}
}
func (ph *PlaneHelper) SetGeometry(v Geometry) {
	ph.Set("geometry", v.JSValue())
}
func (ph *PlaneHelper) Id() int {
	return ph.Get("id").Int()
}
func (ph *PlaneHelper) SetId(v int) {
	ph.Set("id", v)
}
func (ph *PlaneHelper) IsLine() bool {
	return ph.Get("isLine").Bool()
}
func (ph *PlaneHelper) SetIsLine(v bool) {
	ph.Set("isLine", v)
}
func (ph *PlaneHelper) IsLineSegments() bool {
	return ph.Get("isLineSegments").Bool()
}
func (ph *PlaneHelper) SetIsLineSegments(v bool) {
	ph.Set("isLineSegments", v)
}
func (ph *PlaneHelper) IsObject3D() bool {
	return ph.Get("isObject3D").Bool()
}
func (ph *PlaneHelper) SetIsObject3D(v bool) {
	ph.Set("isObject3D", v)
}
func (ph *PlaneHelper) Layers() *Layers {
	return &Layers{Value: ph.Get("layers")}
}
func (ph *PlaneHelper) SetLayers(v *Layers) {
	ph.Set("layers", v.Value)
}
func (ph *PlaneHelper) Material() Material {
	return &MaterialImpl{Value: ph.Get("material")}
}
func (ph *PlaneHelper) SetMaterial(v Material) {
	ph.Set("material", v.JSValue())
}
func (ph *PlaneHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: ph.Get("matrix")}
}
func (ph *PlaneHelper) SetMatrix(v *Matrix4) {
	ph.Set("matrix", v.Value)
}
func (ph *PlaneHelper) MatrixAutoUpdate() bool {
	return ph.Get("matrixAutoUpdate").Bool()
}
func (ph *PlaneHelper) SetMatrixAutoUpdate(v bool) {
	ph.Set("matrixAutoUpdate", v)
}
func (ph *PlaneHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ph.Get("matrixWorld")}
}
func (ph *PlaneHelper) SetMatrixWorld(v *Matrix4) {
	ph.Set("matrixWorld", v.Value)
}
func (ph *PlaneHelper) MatrixWorldNeedsUpdate() bool {
	return ph.Get("matrixWorldNeedsUpdate").Bool()
}
func (ph *PlaneHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ph.Set("matrixWorldNeedsUpdate", v)
}
func (ph *PlaneHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ph.Get("modelViewMatrix")}
}
func (ph *PlaneHelper) SetModelViewMatrix(v *Matrix4) {
	ph.Set("modelViewMatrix", v.Value)
}
func (ph *PlaneHelper) Name() string {
	return ph.Get("name").String()
}
func (ph *PlaneHelper) SetName(v string) {
	ph.Set("name", v)
}
func (ph *PlaneHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ph.Get("normalMatrix")}
}
func (ph *PlaneHelper) SetNormalMatrix(v *Matrix3) {
	ph.Set("normalMatrix", v.Value)
}
func (ph *PlaneHelper) OnAfterRender() js.Value {
	return ph.Get("onAfterRender")
}
func (ph *PlaneHelper) SetOnAfterRender(v js.Value) {
	ph.Set("onAfterRender", v)
}
func (ph *PlaneHelper) OnBeforeRender() js.Value {
	return ph.Get("onBeforeRender")
}
func (ph *PlaneHelper) SetOnBeforeRender(v js.Value) {
	ph.Set("onBeforeRender", v)
}
func (ph *PlaneHelper) Parent() *Object3D {
	return &Object3D{Value: ph.Get("parent")}
}
func (ph *PlaneHelper) SetParent(v *Object3D) {
	ph.Set("parent", v.Value)
}
func (ph *PlaneHelper) Plane() *Plane {
	return &Plane{Value: ph.Get("plane")}
}
func (ph *PlaneHelper) SetPlane(v *Plane) {
	ph.Set("plane", v.Value)
}
func (ph *PlaneHelper) Position() *Vector3 {
	return &Vector3{Value: ph.Get("position")}
}
func (ph *PlaneHelper) SetPosition(v *Vector3) {
	ph.Set("position", v.Value)
}
func (ph *PlaneHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: ph.Get("quaternion")}
}
func (ph *PlaneHelper) SetQuaternion(v *Quaternion) {
	ph.Set("quaternion", v.Value)
}
func (ph *PlaneHelper) ReceiveShadow() bool {
	return ph.Get("receiveShadow").Bool()
}
func (ph *PlaneHelper) SetReceiveShadow(v bool) {
	ph.Set("receiveShadow", v)
}
func (ph *PlaneHelper) RenderOrder() float64 {
	return ph.Get("renderOrder").Float()
}
func (ph *PlaneHelper) SetRenderOrder(v float64) {
	ph.Set("renderOrder", v)
}
func (ph *PlaneHelper) Rotation() *Euler {
	return &Euler{Value: ph.Get("rotation")}
}
func (ph *PlaneHelper) SetRotation(v *Euler) {
	ph.Set("rotation", v.Value)
}
func (ph *PlaneHelper) Scale() *Vector3 {
	return &Vector3{Value: ph.Get("scale")}
}
func (ph *PlaneHelper) SetScale(v *Vector3) {
	ph.Set("scale", v.Value)
}
func (ph *PlaneHelper) Size() float64 {
	return ph.Get("size").Float()
}
func (ph *PlaneHelper) SetSize(v float64) {
	ph.Set("size", v)
}
func (ph *PlaneHelper) Type() string {
	return ph.Get("type").String()
}
func (ph *PlaneHelper) SetType(v string) {
	ph.Set("type", v)
}
func (ph *PlaneHelper) Up() *Vector3 {
	return &Vector3{Value: ph.Get("up")}
}
func (ph *PlaneHelper) SetUp(v *Vector3) {
	ph.Set("up", v.Value)
}
func (ph *PlaneHelper) UserData() js.Value {
	return ph.Get("userData")
}
func (ph *PlaneHelper) SetUserData(v js.Value) {
	ph.Set("userData", v)
}
func (ph *PlaneHelper) Uuid() string {
	return ph.Get("uuid").String()
}
func (ph *PlaneHelper) SetUuid(v string) {
	ph.Set("uuid", v)
}
func (ph *PlaneHelper) Visible() bool {
	return ph.Get("visible").Bool()
}
func (ph *PlaneHelper) SetVisible(v bool) {
	ph.Set("visible", v)
}
func (ph *PlaneHelper) DefaultMatrixAutoUpdate() bool {
	return ph.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ph *PlaneHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ph.Set("DefaultMatrixAutoUpdate", v)
}
func (ph *PlaneHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: ph.Get("DefaultUp")}
}
func (ph *PlaneHelper) SetDefaultUp(v *Vector3) {
	ph.Set("DefaultUp", v.Value)
}
func (ph *PlaneHelper) Add(object js.Value) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("add", object)}
}
func (ph *PlaneHelper) AddEventListener(typ string, listener js.Value) {
	ph.Call("addEventListener", typ, listener)
}
func (ph *PlaneHelper) ApplyMatrix(matrix *Matrix4) {
	ph.Call("applyMatrix", matrix)
}
func (ph *PlaneHelper) ApplyQuaternion(quaternion *Quaternion) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("applyQuaternion", quaternion)}
}
func (ph *PlaneHelper) Clone(recursive bool) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("clone", recursive)}
}
func (ph *PlaneHelper) ComputeLineDistances() *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("computeLineDistances")}
}
func (ph *PlaneHelper) Copy(source *Object3D, recursive bool) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("copy", source, recursive)}
}
func (ph *PlaneHelper) DispatchEvent(event js.Value) {
	ph.Call("dispatchEvent", event)
}
func (ph *PlaneHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ph.Call("getObjectById", id)}
}
func (ph *PlaneHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ph.Call("getObjectByName", name)}
}
func (ph *PlaneHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ph.Call("getObjectByProperty", name, value)}
}
func (ph *PlaneHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ph.Call("getWorldDirection", target)}
}
func (ph *PlaneHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ph.Call("getWorldPosition", target)}
}
func (ph *PlaneHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ph.Call("getWorldQuaternion", target)}
}
func (ph *PlaneHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ph.Call("getWorldScale", target)}
}
func (ph *PlaneHelper) HasEventListener(typ string, listener js.Value) bool {
	return ph.Call("hasEventListener", typ, listener).Bool()
}
func (ph *PlaneHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ph.Call("localToWorld", vector)}
}
func (ph *PlaneHelper) LookAt(vector *Vector3, y float64, z float64) {
	ph.Call("lookAt", vector, y, z)
}
func (ph *PlaneHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	ph.Call("raycast", raycaster, intersects)
}
func (ph *PlaneHelper) Remove(object js.Value) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("remove", object)}
}
func (ph *PlaneHelper) RemoveEventListener(typ string, listener js.Value) {
	ph.Call("removeEventListener", typ, listener)
}
func (ph *PlaneHelper) RotateOnAxis(axis *Vector3, angle float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("rotateOnAxis", axis, angle)}
}
func (ph *PlaneHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("rotateOnWorldAxis", axis, angle)}
}
func (ph *PlaneHelper) RotateX(angle float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("rotateX", angle)}
}
func (ph *PlaneHelper) RotateY(angle float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("rotateY", angle)}
}
func (ph *PlaneHelper) RotateZ(angle float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("rotateZ", angle)}
}
func (ph *PlaneHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ph.Call("setRotationFromAxisAngle", axis, angle)
}
func (ph *PlaneHelper) SetRotationFromEuler(euler *Euler) {
	ph.Call("setRotationFromEuler", euler)
}
func (ph *PlaneHelper) SetRotationFromMatrix(m *Matrix4) {
	ph.Call("setRotationFromMatrix", m)
}
func (ph *PlaneHelper) SetRotationFromQuaternion(q *Quaternion) {
	ph.Call("setRotationFromQuaternion", q)
}
func (ph *PlaneHelper) ToJSON(meta js.Value) js.Value {
	return ph.Call("toJSON", meta)
}
func (ph *PlaneHelper) TranslateOnAxis(axis *Vector3, distance float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("translateOnAxis", axis, distance)}
}
func (ph *PlaneHelper) TranslateX(distance float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("translateX", distance)}
}
func (ph *PlaneHelper) TranslateY(distance float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("translateY", distance)}
}
func (ph *PlaneHelper) TranslateZ(distance float64) *PlaneHelper {
	return &PlaneHelper{Value: ph.Call("translateZ", distance)}
}
func (ph *PlaneHelper) Traverse(callback js.Value) {
	ph.Call("traverse", callback)
}
func (ph *PlaneHelper) TraverseAncestors(callback js.Value) {
	ph.Call("traverseAncestors", callback)
}
func (ph *PlaneHelper) TraverseVisible(callback js.Value) {
	ph.Call("traverseVisible", callback)
}
func (ph *PlaneHelper) UpdateMatrix() {
	ph.Call("updateMatrix")
}
func (ph *PlaneHelper) UpdateMatrixWorld(force bool) {
	ph.Call("updateMatrixWorld", force)
}
func (ph *PlaneHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ph.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ph *PlaneHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ph.Call("worldToLocal", vector)}
}

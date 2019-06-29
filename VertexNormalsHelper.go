// Code generated from "helpers/VertexNormalsHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// VertexNormalsHelper extend: [LineSegments]
type VertexNormalsHelper struct {
	js.Value
}

func NewVertexNormalsHelper(object *Object3D, size int, hex int, linewidth float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: get("VertexNormalsHelper").New(object.JSValue(), size, hex, linewidth)}
}
func (vnh *VertexNormalsHelper) JSValue() js.Value {
	return vnh.Value
}
func (vnh *VertexNormalsHelper) CastShadow() bool {
	return vnh.Get("castShadow").Bool()
}
func (vnh *VertexNormalsHelper) SetCastShadow(v bool) {
	vnh.Set("castShadow", v)
}
func (vnh *VertexNormalsHelper) Children() js.Value {
	return vnh.Get("children")
}
func (vnh *VertexNormalsHelper) SetChildren(v js.Value) {
	vnh.Set("children", v)
}
func (vnh *VertexNormalsHelper) FrustumCulled() bool {
	return vnh.Get("frustumCulled").Bool()
}
func (vnh *VertexNormalsHelper) SetFrustumCulled(v bool) {
	vnh.Set("frustumCulled", v)
}
func (vnh *VertexNormalsHelper) Geometry() Geometry {
	return &GeometryImpl{Value: vnh.Get("geometry")}
}
func (vnh *VertexNormalsHelper) SetGeometry(v Geometry) {
	vnh.Set("geometry", v.JSValue())
}
func (vnh *VertexNormalsHelper) Id() int {
	return vnh.Get("id").Int()
}
func (vnh *VertexNormalsHelper) SetId(v int) {
	vnh.Set("id", v)
}
func (vnh *VertexNormalsHelper) IsLine() bool {
	return vnh.Get("isLine").Bool()
}
func (vnh *VertexNormalsHelper) SetIsLine(v bool) {
	vnh.Set("isLine", v)
}
func (vnh *VertexNormalsHelper) IsLineSegments() bool {
	return vnh.Get("isLineSegments").Bool()
}
func (vnh *VertexNormalsHelper) SetIsLineSegments(v bool) {
	vnh.Set("isLineSegments", v)
}
func (vnh *VertexNormalsHelper) IsObject3D() bool {
	return vnh.Get("isObject3D").Bool()
}
func (vnh *VertexNormalsHelper) SetIsObject3D(v bool) {
	vnh.Set("isObject3D", v)
}
func (vnh *VertexNormalsHelper) Layers() *Layers {
	return &Layers{Value: vnh.Get("layers")}
}
func (vnh *VertexNormalsHelper) SetLayers(v *Layers) {
	vnh.Set("layers", v.JSValue())
}
func (vnh *VertexNormalsHelper) Material() Material {
	return &MaterialImpl{Value: vnh.Get("material")}
}
func (vnh *VertexNormalsHelper) SetMaterial(v Material) {
	vnh.Set("material", v.JSValue())
}
func (vnh *VertexNormalsHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: vnh.Get("matrix")}
}
func (vnh *VertexNormalsHelper) SetMatrix(v *Matrix4) {
	vnh.Set("matrix", v.JSValue())
}
func (vnh *VertexNormalsHelper) MatrixAutoUpdate() bool {
	return vnh.Get("matrixAutoUpdate").Bool()
}
func (vnh *VertexNormalsHelper) SetMatrixAutoUpdate(v bool) {
	vnh.Set("matrixAutoUpdate", v)
}
func (vnh *VertexNormalsHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: vnh.Get("matrixWorld")}
}
func (vnh *VertexNormalsHelper) SetMatrixWorld(v *Matrix4) {
	vnh.Set("matrixWorld", v.JSValue())
}
func (vnh *VertexNormalsHelper) MatrixWorldNeedsUpdate() bool {
	return vnh.Get("matrixWorldNeedsUpdate").Bool()
}
func (vnh *VertexNormalsHelper) SetMatrixWorldNeedsUpdate(v bool) {
	vnh.Set("matrixWorldNeedsUpdate", v)
}
func (vnh *VertexNormalsHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: vnh.Get("modelViewMatrix")}
}
func (vnh *VertexNormalsHelper) SetModelViewMatrix(v *Matrix4) {
	vnh.Set("modelViewMatrix", v.JSValue())
}
func (vnh *VertexNormalsHelper) Name() string {
	return vnh.Get("name").String()
}
func (vnh *VertexNormalsHelper) SetName(v string) {
	vnh.Set("name", v)
}
func (vnh *VertexNormalsHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: vnh.Get("normalMatrix")}
}
func (vnh *VertexNormalsHelper) SetNormalMatrix(v *Matrix3) {
	vnh.Set("normalMatrix", v.JSValue())
}
func (vnh *VertexNormalsHelper) Object() *Object3D {
	return &Object3D{Value: vnh.Get("object")}
}
func (vnh *VertexNormalsHelper) SetObject(v *Object3D) {
	vnh.Set("object", v.JSValue())
}
func (vnh *VertexNormalsHelper) OnAfterRender() js.Value {
	return vnh.Get("onAfterRender")
}
func (vnh *VertexNormalsHelper) SetOnAfterRender(v js.Value) {
	vnh.Set("onAfterRender", v)
}
func (vnh *VertexNormalsHelper) OnBeforeRender() js.Value {
	return vnh.Get("onBeforeRender")
}
func (vnh *VertexNormalsHelper) SetOnBeforeRender(v js.Value) {
	vnh.Set("onBeforeRender", v)
}
func (vnh *VertexNormalsHelper) Parent() *Object3D {
	return &Object3D{Value: vnh.Get("parent")}
}
func (vnh *VertexNormalsHelper) SetParent(v *Object3D) {
	vnh.Set("parent", v.JSValue())
}
func (vnh *VertexNormalsHelper) Position() *Vector3 {
	return &Vector3{Value: vnh.Get("position")}
}
func (vnh *VertexNormalsHelper) SetPosition(v *Vector3) {
	vnh.Set("position", v.JSValue())
}
func (vnh *VertexNormalsHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: vnh.Get("quaternion")}
}
func (vnh *VertexNormalsHelper) SetQuaternion(v *Quaternion) {
	vnh.Set("quaternion", v.JSValue())
}
func (vnh *VertexNormalsHelper) ReceiveShadow() bool {
	return vnh.Get("receiveShadow").Bool()
}
func (vnh *VertexNormalsHelper) SetReceiveShadow(v bool) {
	vnh.Set("receiveShadow", v)
}
func (vnh *VertexNormalsHelper) RenderOrder() float64 {
	return vnh.Get("renderOrder").Float()
}
func (vnh *VertexNormalsHelper) SetRenderOrder(v float64) {
	vnh.Set("renderOrder", v)
}
func (vnh *VertexNormalsHelper) Rotation() *Euler {
	return &Euler{Value: vnh.Get("rotation")}
}
func (vnh *VertexNormalsHelper) SetRotation(v *Euler) {
	vnh.Set("rotation", v.JSValue())
}
func (vnh *VertexNormalsHelper) Scale() *Vector3 {
	return &Vector3{Value: vnh.Get("scale")}
}
func (vnh *VertexNormalsHelper) SetScale(v *Vector3) {
	vnh.Set("scale", v.JSValue())
}
func (vnh *VertexNormalsHelper) Size() float64 {
	return vnh.Get("size").Float()
}
func (vnh *VertexNormalsHelper) SetSize(v float64) {
	vnh.Set("size", v)
}
func (vnh *VertexNormalsHelper) Type() string {
	return vnh.Get("type").String()
}
func (vnh *VertexNormalsHelper) SetType(v string) {
	vnh.Set("type", v)
}
func (vnh *VertexNormalsHelper) Up() *Vector3 {
	return &Vector3{Value: vnh.Get("up")}
}
func (vnh *VertexNormalsHelper) SetUp(v *Vector3) {
	vnh.Set("up", v.JSValue())
}
func (vnh *VertexNormalsHelper) UserData() js.Value {
	return vnh.Get("userData")
}
func (vnh *VertexNormalsHelper) SetUserData(v js.Value) {
	vnh.Set("userData", v)
}
func (vnh *VertexNormalsHelper) Uuid() string {
	return vnh.Get("uuid").String()
}
func (vnh *VertexNormalsHelper) SetUuid(v string) {
	vnh.Set("uuid", v)
}
func (vnh *VertexNormalsHelper) Visible() bool {
	return vnh.Get("visible").Bool()
}
func (vnh *VertexNormalsHelper) SetVisible(v bool) {
	vnh.Set("visible", v)
}
func (vnh *VertexNormalsHelper) DefaultMatrixAutoUpdate() bool {
	return vnh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (vnh *VertexNormalsHelper) SetDefaultMatrixAutoUpdate(v bool) {
	vnh.Set("DefaultMatrixAutoUpdate", v)
}
func (vnh *VertexNormalsHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: vnh.Get("DefaultUp")}
}
func (vnh *VertexNormalsHelper) SetDefaultUp(v *Vector3) {
	vnh.Set("DefaultUp", v.JSValue())
}
func (vnh *VertexNormalsHelper) Add(object js.Value) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("add", object)}
}
func (vnh *VertexNormalsHelper) AddEventListener(typ string, listener js.Value) {
	vnh.Call("addEventListener", typ, listener)
}
func (vnh *VertexNormalsHelper) ApplyMatrix(matrix *Matrix4) {
	vnh.Call("applyMatrix", matrix.JSValue())
}
func (vnh *VertexNormalsHelper) ApplyQuaternion(quaternion *Quaternion) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("applyQuaternion", quaternion)}
}
func (vnh *VertexNormalsHelper) Clone(recursive bool) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("clone", recursive)}
}
func (vnh *VertexNormalsHelper) ComputeLineDistances() *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("computeLineDistances")}
}
func (vnh *VertexNormalsHelper) Copy(source *Object3D, recursive bool) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("copy", source, recursive)}
}
func (vnh *VertexNormalsHelper) DispatchEvent(event js.Value) {
	vnh.Call("dispatchEvent", event)
}
func (vnh *VertexNormalsHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: vnh.Call("getObjectById", id)}
}
func (vnh *VertexNormalsHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: vnh.Call("getObjectByName", name)}
}
func (vnh *VertexNormalsHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: vnh.Call("getObjectByProperty", name, value)}
}
func (vnh *VertexNormalsHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: vnh.Call("getWorldDirection", target)}
}
func (vnh *VertexNormalsHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: vnh.Call("getWorldPosition", target)}
}
func (vnh *VertexNormalsHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: vnh.Call("getWorldQuaternion", target)}
}
func (vnh *VertexNormalsHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: vnh.Call("getWorldScale", target)}
}
func (vnh *VertexNormalsHelper) HasEventListener(typ string, listener js.Value) bool {
	return vnh.Call("hasEventListener", typ, listener).Bool()
}
func (vnh *VertexNormalsHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: vnh.Call("localToWorld", vector)}
}
func (vnh *VertexNormalsHelper) LookAt(vector *Vector3, y float64, z float64) {
	vnh.Call("lookAt", vector, y, z)
}
func (vnh *VertexNormalsHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	vnh.Call("raycast", raycaster.JSValue(), intersects)
}
func (vnh *VertexNormalsHelper) Remove(object js.Value) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("remove", object)}
}
func (vnh *VertexNormalsHelper) RemoveEventListener(typ string, listener js.Value) {
	vnh.Call("removeEventListener", typ, listener)
}
func (vnh *VertexNormalsHelper) RotateOnAxis(axis *Vector3, angle float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("rotateOnAxis", axis, angle)}
}
func (vnh *VertexNormalsHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("rotateOnWorldAxis", axis, angle)}
}
func (vnh *VertexNormalsHelper) RotateX(angle float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("rotateX", angle)}
}
func (vnh *VertexNormalsHelper) RotateY(angle float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("rotateY", angle)}
}
func (vnh *VertexNormalsHelper) RotateZ(angle float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("rotateZ", angle)}
}
func (vnh *VertexNormalsHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	vnh.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (vnh *VertexNormalsHelper) SetRotationFromEuler(euler *Euler) {
	vnh.Call("setRotationFromEuler", euler.JSValue())
}
func (vnh *VertexNormalsHelper) SetRotationFromMatrix(m *Matrix4) {
	vnh.Call("setRotationFromMatrix", m.JSValue())
}
func (vnh *VertexNormalsHelper) SetRotationFromQuaternion(q *Quaternion) {
	vnh.Call("setRotationFromQuaternion", q.JSValue())
}
func (vnh *VertexNormalsHelper) ToJSON(meta js.Value) js.Value {
	return vnh.Call("toJSON", meta)
}
func (vnh *VertexNormalsHelper) TranslateOnAxis(axis *Vector3, distance float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("translateOnAxis", axis, distance)}
}
func (vnh *VertexNormalsHelper) TranslateX(distance float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("translateX", distance)}
}
func (vnh *VertexNormalsHelper) TranslateY(distance float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("translateY", distance)}
}
func (vnh *VertexNormalsHelper) TranslateZ(distance float64) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: vnh.Call("translateZ", distance)}
}
func (vnh *VertexNormalsHelper) Traverse(callback js.Value) {
	vnh.Call("traverse", callback)
}
func (vnh *VertexNormalsHelper) TraverseAncestors(callback js.Value) {
	vnh.Call("traverseAncestors", callback)
}
func (vnh *VertexNormalsHelper) TraverseVisible(callback js.Value) {
	vnh.Call("traverseVisible", callback)
}
func (vnh *VertexNormalsHelper) Update(object *Object3D) {
	vnh.Call("update", object.JSValue())
}
func (vnh *VertexNormalsHelper) UpdateMatrix() {
	vnh.Call("updateMatrix")
}
func (vnh *VertexNormalsHelper) UpdateMatrixWorld(force bool) {
	vnh.Call("updateMatrixWorld", force)
}
func (vnh *VertexNormalsHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	vnh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (vnh *VertexNormalsHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: vnh.Call("worldToLocal", vector)}
}

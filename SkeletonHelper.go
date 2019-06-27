// Code generated from "helpers/SkeletonHelper.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// SkeletonHelper extend: [LineSegments]
type SkeletonHelper struct {
	js.Value
}

func NewSkeletonHelper(bone *Object3D) *SkeletonHelper {
	return &SkeletonHelper{Value: get("SkeletonHelper").New(bone)}
}
func (sh *SkeletonHelper) JSValue() js.Value {
	return sh.Value
}
func (sh *SkeletonHelper) Bones() js.Value {
	return sh.Get("bones")
}
func (sh *SkeletonHelper) SetBones(v js.Value) {
	sh.Set("bones", v)
}
func (sh *SkeletonHelper) CastShadow() bool {
	return sh.Get("castShadow").Bool()
}
func (sh *SkeletonHelper) SetCastShadow(v bool) {
	sh.Set("castShadow", v)
}
func (sh *SkeletonHelper) Children() js.Value {
	return sh.Get("children")
}
func (sh *SkeletonHelper) SetChildren(v js.Value) {
	sh.Set("children", v)
}
func (sh *SkeletonHelper) FrustumCulled() bool {
	return sh.Get("frustumCulled").Bool()
}
func (sh *SkeletonHelper) SetFrustumCulled(v bool) {
	sh.Set("frustumCulled", v)
}
func (sh *SkeletonHelper) Geometry() Geometry {
	return &GeometryImpl{Value: sh.Get("geometry")}
}
func (sh *SkeletonHelper) SetGeometry(v Geometry) {
	sh.Set("geometry", v.JSValue())
}
func (sh *SkeletonHelper) Id() int {
	return sh.Get("id").Int()
}
func (sh *SkeletonHelper) SetId(v int) {
	sh.Set("id", v)
}
func (sh *SkeletonHelper) IsLine() bool {
	return sh.Get("isLine").Bool()
}
func (sh *SkeletonHelper) SetIsLine(v bool) {
	sh.Set("isLine", v)
}
func (sh *SkeletonHelper) IsLineSegments() bool {
	return sh.Get("isLineSegments").Bool()
}
func (sh *SkeletonHelper) SetIsLineSegments(v bool) {
	sh.Set("isLineSegments", v)
}
func (sh *SkeletonHelper) IsObject3D() bool {
	return sh.Get("isObject3D").Bool()
}
func (sh *SkeletonHelper) SetIsObject3D(v bool) {
	sh.Set("isObject3D", v)
}
func (sh *SkeletonHelper) Layers() *Layers {
	return &Layers{Value: sh.Get("layers")}
}
func (sh *SkeletonHelper) SetLayers(v *Layers) {
	sh.Set("layers", v.Value)
}
func (sh *SkeletonHelper) Material() Material {
	return &MaterialImpl{Value: sh.Get("material")}
}
func (sh *SkeletonHelper) SetMaterial(v Material) {
	sh.Set("material", v.JSValue())
}
func (sh *SkeletonHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: sh.Get("matrix")}
}
func (sh *SkeletonHelper) SetMatrix(v *Matrix4) {
	sh.Set("matrix", v.Value)
}
func (sh *SkeletonHelper) MatrixAutoUpdate() bool {
	return sh.Get("matrixAutoUpdate").Bool()
}
func (sh *SkeletonHelper) SetMatrixAutoUpdate(v bool) {
	sh.Set("matrixAutoUpdate", v)
}
func (sh *SkeletonHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: sh.Get("matrixWorld")}
}
func (sh *SkeletonHelper) SetMatrixWorld(v *Matrix4) {
	sh.Set("matrixWorld", v.Value)
}
func (sh *SkeletonHelper) MatrixWorldNeedsUpdate() bool {
	return sh.Get("matrixWorldNeedsUpdate").Bool()
}
func (sh *SkeletonHelper) SetMatrixWorldNeedsUpdate(v bool) {
	sh.Set("matrixWorldNeedsUpdate", v)
}
func (sh *SkeletonHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: sh.Get("modelViewMatrix")}
}
func (sh *SkeletonHelper) SetModelViewMatrix(v *Matrix4) {
	sh.Set("modelViewMatrix", v.Value)
}
func (sh *SkeletonHelper) Name() string {
	return sh.Get("name").String()
}
func (sh *SkeletonHelper) SetName(v string) {
	sh.Set("name", v)
}
func (sh *SkeletonHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: sh.Get("normalMatrix")}
}
func (sh *SkeletonHelper) SetNormalMatrix(v *Matrix3) {
	sh.Set("normalMatrix", v.Value)
}
func (sh *SkeletonHelper) OnAfterRender() js.Value {
	return sh.Get("onAfterRender")
}
func (sh *SkeletonHelper) SetOnAfterRender(v js.Value) {
	sh.Set("onAfterRender", v)
}
func (sh *SkeletonHelper) OnBeforeRender() js.Value {
	return sh.Get("onBeforeRender")
}
func (sh *SkeletonHelper) SetOnBeforeRender(v js.Value) {
	sh.Set("onBeforeRender", v)
}
func (sh *SkeletonHelper) Parent() *Object3D {
	return &Object3D{Value: sh.Get("parent")}
}
func (sh *SkeletonHelper) SetParent(v *Object3D) {
	sh.Set("parent", v.Value)
}
func (sh *SkeletonHelper) Position() *Vector3 {
	return &Vector3{Value: sh.Get("position")}
}
func (sh *SkeletonHelper) SetPosition(v *Vector3) {
	sh.Set("position", v.Value)
}
func (sh *SkeletonHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: sh.Get("quaternion")}
}
func (sh *SkeletonHelper) SetQuaternion(v *Quaternion) {
	sh.Set("quaternion", v.Value)
}
func (sh *SkeletonHelper) ReceiveShadow() bool {
	return sh.Get("receiveShadow").Bool()
}
func (sh *SkeletonHelper) SetReceiveShadow(v bool) {
	sh.Set("receiveShadow", v)
}
func (sh *SkeletonHelper) RenderOrder() float64 {
	return sh.Get("renderOrder").Float()
}
func (sh *SkeletonHelper) SetRenderOrder(v float64) {
	sh.Set("renderOrder", v)
}
func (sh *SkeletonHelper) Root() *Object3D {
	return &Object3D{Value: sh.Get("root")}
}
func (sh *SkeletonHelper) SetRoot(v *Object3D) {
	sh.Set("root", v.Value)
}
func (sh *SkeletonHelper) Rotation() *Euler {
	return &Euler{Value: sh.Get("rotation")}
}
func (sh *SkeletonHelper) SetRotation(v *Euler) {
	sh.Set("rotation", v.Value)
}
func (sh *SkeletonHelper) Scale() *Vector3 {
	return &Vector3{Value: sh.Get("scale")}
}
func (sh *SkeletonHelper) SetScale(v *Vector3) {
	sh.Set("scale", v.Value)
}
func (sh *SkeletonHelper) Type() string {
	return sh.Get("type").String()
}
func (sh *SkeletonHelper) SetType(v string) {
	sh.Set("type", v)
}
func (sh *SkeletonHelper) Up() *Vector3 {
	return &Vector3{Value: sh.Get("up")}
}
func (sh *SkeletonHelper) SetUp(v *Vector3) {
	sh.Set("up", v.Value)
}
func (sh *SkeletonHelper) UserData() js.Value {
	return sh.Get("userData")
}
func (sh *SkeletonHelper) SetUserData(v js.Value) {
	sh.Set("userData", v)
}
func (sh *SkeletonHelper) Uuid() string {
	return sh.Get("uuid").String()
}
func (sh *SkeletonHelper) SetUuid(v string) {
	sh.Set("uuid", v)
}
func (sh *SkeletonHelper) Visible() bool {
	return sh.Get("visible").Bool()
}
func (sh *SkeletonHelper) SetVisible(v bool) {
	sh.Set("visible", v)
}
func (sh *SkeletonHelper) DefaultMatrixAutoUpdate() bool {
	return sh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sh *SkeletonHelper) SetDefaultMatrixAutoUpdate(v bool) {
	sh.Set("DefaultMatrixAutoUpdate", v)
}
func (sh *SkeletonHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: sh.Get("DefaultUp")}
}
func (sh *SkeletonHelper) SetDefaultUp(v *Vector3) {
	sh.Set("DefaultUp", v.Value)
}
func (sh *SkeletonHelper) Add(object js.Value) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("add", object)}
}
func (sh *SkeletonHelper) AddEventListener(typ string, listener js.Value) {
	sh.Call("addEventListener", typ, listener)
}
func (sh *SkeletonHelper) ApplyMatrix(matrix *Matrix4) {
	sh.Call("applyMatrix", matrix)
}
func (sh *SkeletonHelper) ApplyQuaternion(quaternion *Quaternion) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("applyQuaternion", quaternion)}
}
func (sh *SkeletonHelper) Clone(recursive bool) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("clone", recursive)}
}
func (sh *SkeletonHelper) ComputeLineDistances() *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("computeLineDistances")}
}
func (sh *SkeletonHelper) Copy(source *Object3D, recursive bool) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("copy", source, recursive)}
}
func (sh *SkeletonHelper) DispatchEvent(event js.Value) {
	sh.Call("dispatchEvent", event)
}
func (sh *SkeletonHelper) GetBoneList(object *Object3D) js.Value {
	return sh.Call("getBoneList", object)
}
func (sh *SkeletonHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: sh.Call("getObjectById", id)}
}
func (sh *SkeletonHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: sh.Call("getObjectByName", name)}
}
func (sh *SkeletonHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: sh.Call("getObjectByProperty", name, value)}
}
func (sh *SkeletonHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: sh.Call("getWorldDirection", target)}
}
func (sh *SkeletonHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: sh.Call("getWorldPosition", target)}
}
func (sh *SkeletonHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: sh.Call("getWorldQuaternion", target)}
}
func (sh *SkeletonHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: sh.Call("getWorldScale", target)}
}
func (sh *SkeletonHelper) HasEventListener(typ string, listener js.Value) bool {
	return sh.Call("hasEventListener", typ, listener).Bool()
}
func (sh *SkeletonHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: sh.Call("localToWorld", vector)}
}
func (sh *SkeletonHelper) LookAt(vector *Vector3, y float64, z float64) {
	sh.Call("lookAt", vector, y, z)
}
func (sh *SkeletonHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	sh.Call("raycast", raycaster, intersects)
}
func (sh *SkeletonHelper) Remove(object js.Value) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("remove", object)}
}
func (sh *SkeletonHelper) RemoveEventListener(typ string, listener js.Value) {
	sh.Call("removeEventListener", typ, listener)
}
func (sh *SkeletonHelper) RotateOnAxis(axis *Vector3, angle float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("rotateOnAxis", axis, angle)}
}
func (sh *SkeletonHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("rotateOnWorldAxis", axis, angle)}
}
func (sh *SkeletonHelper) RotateX(angle float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("rotateX", angle)}
}
func (sh *SkeletonHelper) RotateY(angle float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("rotateY", angle)}
}
func (sh *SkeletonHelper) RotateZ(angle float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("rotateZ", angle)}
}
func (sh *SkeletonHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	sh.Call("setRotationFromAxisAngle", axis, angle)
}
func (sh *SkeletonHelper) SetRotationFromEuler(euler *Euler) {
	sh.Call("setRotationFromEuler", euler)
}
func (sh *SkeletonHelper) SetRotationFromMatrix(m *Matrix4) {
	sh.Call("setRotationFromMatrix", m)
}
func (sh *SkeletonHelper) SetRotationFromQuaternion(q *Quaternion) {
	sh.Call("setRotationFromQuaternion", q)
}
func (sh *SkeletonHelper) ToJSON(meta js.Value) js.Value {
	return sh.Call("toJSON", meta)
}
func (sh *SkeletonHelper) TranslateOnAxis(axis *Vector3, distance float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("translateOnAxis", axis, distance)}
}
func (sh *SkeletonHelper) TranslateX(distance float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("translateX", distance)}
}
func (sh *SkeletonHelper) TranslateY(distance float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("translateY", distance)}
}
func (sh *SkeletonHelper) TranslateZ(distance float64) *SkeletonHelper {
	return &SkeletonHelper{Value: sh.Call("translateZ", distance)}
}
func (sh *SkeletonHelper) Traverse(callback js.Value) {
	sh.Call("traverse", callback)
}
func (sh *SkeletonHelper) TraverseAncestors(callback js.Value) {
	sh.Call("traverseAncestors", callback)
}
func (sh *SkeletonHelper) TraverseVisible(callback js.Value) {
	sh.Call("traverseVisible", callback)
}
func (sh *SkeletonHelper) Update() {
	sh.Call("update")
}
func (sh *SkeletonHelper) UpdateMatrix() {
	sh.Call("updateMatrix")
}
func (sh *SkeletonHelper) UpdateMatrixWorld(force bool) {
	sh.Call("updateMatrixWorld", force)
}
func (sh *SkeletonHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sh *SkeletonHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: sh.Call("worldToLocal", vector)}
}

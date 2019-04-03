// Code generated from "helpers/GridHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type GridHelper struct {
	js.Value
}

func NewGridHelper(size float64, divisions float64, color1 *Color, color2 *Color) *GridHelper {
	return &GridHelper{Value: get("GridHelper").New(size, divisions, color1, color2)}
}
func (gh *GridHelper) CastShadow() bool {
	return gh.Get("castShadow").Bool()
}
func (gh *GridHelper) SetCastShadow(v bool) {
	gh.Set("castShadow", v)
}
func (gh *GridHelper) Children() js.Value {
	return gh.Get("children")
}
func (gh *GridHelper) SetChildren(v js.Value) {
	gh.Set("children", v)
}
func (gh *GridHelper) FrustumCulled() bool {
	return gh.Get("frustumCulled").Bool()
}
func (gh *GridHelper) SetFrustumCulled(v bool) {
	gh.Set("frustumCulled", v)
}
func (gh *GridHelper) Geometry() *Geometry {
	return &Geometry{Value: gh.Get("geometry")}
}
func (gh *GridHelper) SetGeometry(v *Geometry) {
	gh.Set("geometry", v)
}
func (gh *GridHelper) Id() int {
	return gh.Get("id").Int()
}
func (gh *GridHelper) SetId(v int) {
	gh.Set("id", v)
}
func (gh *GridHelper) IsLine() bool {
	return gh.Get("isLine").Bool()
}
func (gh *GridHelper) SetIsLine(v bool) {
	gh.Set("isLine", v)
}
func (gh *GridHelper) IsLineSegments() bool {
	return gh.Get("isLineSegments").Bool()
}
func (gh *GridHelper) SetIsLineSegments(v bool) {
	gh.Set("isLineSegments", v)
}
func (gh *GridHelper) IsObject3D() bool {
	return gh.Get("isObject3D").Bool()
}
func (gh *GridHelper) SetIsObject3D(v bool) {
	gh.Set("isObject3D", v)
}
func (gh *GridHelper) Layers() *Layers {
	return &Layers{Value: gh.Get("layers")}
}
func (gh *GridHelper) SetLayers(v *Layers) {
	gh.Set("layers", v)
}
func (gh *GridHelper) Material() *Material {
	return &Material{Value: gh.Get("material")}
}
func (gh *GridHelper) SetMaterial(v *Material) {
	gh.Set("material", v)
}
func (gh *GridHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: gh.Get("matrix")}
}
func (gh *GridHelper) SetMatrix(v *Matrix4) {
	gh.Set("matrix", v)
}
func (gh *GridHelper) MatrixAutoUpdate() bool {
	return gh.Get("matrixAutoUpdate").Bool()
}
func (gh *GridHelper) SetMatrixAutoUpdate(v bool) {
	gh.Set("matrixAutoUpdate", v)
}
func (gh *GridHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: gh.Get("matrixWorld")}
}
func (gh *GridHelper) SetMatrixWorld(v *Matrix4) {
	gh.Set("matrixWorld", v)
}
func (gh *GridHelper) MatrixWorldNeedsUpdate() bool {
	return gh.Get("matrixWorldNeedsUpdate").Bool()
}
func (gh *GridHelper) SetMatrixWorldNeedsUpdate(v bool) {
	gh.Set("matrixWorldNeedsUpdate", v)
}
func (gh *GridHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: gh.Get("modelViewMatrix")}
}
func (gh *GridHelper) SetModelViewMatrix(v *Matrix4) {
	gh.Set("modelViewMatrix", v)
}
func (gh *GridHelper) Name() string {
	return gh.Get("name").String()
}
func (gh *GridHelper) SetName(v string) {
	gh.Set("name", v)
}
func (gh *GridHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: gh.Get("normalMatrix")}
}
func (gh *GridHelper) SetNormalMatrix(v *Matrix3) {
	gh.Set("normalMatrix", v)
}
func (gh *GridHelper) OnAfterRender() js.Value {
	return gh.Get("onAfterRender")
}
func (gh *GridHelper) SetOnAfterRender(v js.Value) {
	gh.Set("onAfterRender", v)
}
func (gh *GridHelper) OnBeforeRender() js.Value {
	return gh.Get("onBeforeRender")
}
func (gh *GridHelper) SetOnBeforeRender(v js.Value) {
	gh.Set("onBeforeRender", v)
}
func (gh *GridHelper) Parent() *Object3D {
	return &Object3D{Value: gh.Get("parent")}
}
func (gh *GridHelper) SetParent(v *Object3D) {
	gh.Set("parent", v)
}
func (gh *GridHelper) Position() *Vector3 {
	return &Vector3{Value: gh.Get("position")}
}
func (gh *GridHelper) SetPosition(v *Vector3) {
	gh.Set("position", v)
}
func (gh *GridHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: gh.Get("quaternion")}
}
func (gh *GridHelper) SetQuaternion(v *Quaternion) {
	gh.Set("quaternion", v)
}
func (gh *GridHelper) ReceiveShadow() bool {
	return gh.Get("receiveShadow").Bool()
}
func (gh *GridHelper) SetReceiveShadow(v bool) {
	gh.Set("receiveShadow", v)
}
func (gh *GridHelper) RenderOrder() float64 {
	return gh.Get("renderOrder").Float()
}
func (gh *GridHelper) SetRenderOrder(v float64) {
	gh.Set("renderOrder", v)
}
func (gh *GridHelper) Rotation() *Euler {
	return &Euler{Value: gh.Get("rotation")}
}
func (gh *GridHelper) SetRotation(v *Euler) {
	gh.Set("rotation", v)
}
func (gh *GridHelper) Scale() *Vector3 {
	return &Vector3{Value: gh.Get("scale")}
}
func (gh *GridHelper) SetScale(v *Vector3) {
	gh.Set("scale", v)
}
func (gh *GridHelper) Type() string {
	return gh.Get("type").String()
}
func (gh *GridHelper) SetType(v string) {
	gh.Set("type", v)
}
func (gh *GridHelper) Up() *Vector3 {
	return &Vector3{Value: gh.Get("up")}
}
func (gh *GridHelper) SetUp(v *Vector3) {
	gh.Set("up", v)
}
func (gh *GridHelper) UserData() js.Value {
	return gh.Get("userData")
}
func (gh *GridHelper) SetUserData(v js.Value) {
	gh.Set("userData", v)
}
func (gh *GridHelper) Uuid() string {
	return gh.Get("uuid").String()
}
func (gh *GridHelper) SetUuid(v string) {
	gh.Set("uuid", v)
}
func (gh *GridHelper) Visible() bool {
	return gh.Get("visible").Bool()
}
func (gh *GridHelper) SetVisible(v bool) {
	gh.Set("visible", v)
}
func (gh *GridHelper) DefaultMatrixAutoUpdate() bool {
	return gh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (gh *GridHelper) SetDefaultMatrixAutoUpdate(v bool) {
	gh.Set("DefaultMatrixAutoUpdate", v)
}
func (gh *GridHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: gh.Get("DefaultUp")}
}
func (gh *GridHelper) SetDefaultUp(v *Vector3) {
	gh.Set("DefaultUp", v)
}
func (gh *GridHelper) Add(object js.Value) *GridHelper {
	return &GridHelper{Value: gh.Call("add", object)}
}
func (gh *GridHelper) AddEventListener(typ string, listener js.Value) {
	gh.Call("addEventListener", typ, listener)
}
func (gh *GridHelper) ApplyMatrix(matrix *Matrix4) {
	gh.Call("applyMatrix", matrix)
}
func (gh *GridHelper) ApplyQuaternion(quaternion *Quaternion) *GridHelper {
	return &GridHelper{Value: gh.Call("applyQuaternion", quaternion)}
}
func (gh *GridHelper) Clone(recursive bool) *GridHelper {
	return &GridHelper{Value: gh.Call("clone", recursive)}
}
func (gh *GridHelper) ComputeLineDistances() *GridHelper {
	return &GridHelper{Value: gh.Call("computeLineDistances")}
}
func (gh *GridHelper) Copy(source *Object3D, recursive bool) *GridHelper {
	return &GridHelper{Value: gh.Call("copy", source, recursive)}
}
func (gh *GridHelper) DispatchEvent(event js.Value) {
	gh.Call("dispatchEvent", event)
}
func (gh *GridHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: gh.Call("getObjectById", id)}
}
func (gh *GridHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: gh.Call("getObjectByName", name)}
}
func (gh *GridHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: gh.Call("getObjectByProperty", name, value)}
}
func (gh *GridHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: gh.Call("getWorldDirection", target)}
}
func (gh *GridHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: gh.Call("getWorldPosition", target)}
}
func (gh *GridHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: gh.Call("getWorldQuaternion", target)}
}
func (gh *GridHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: gh.Call("getWorldScale", target)}
}
func (gh *GridHelper) HasEventListener(typ string, listener js.Value) bool {
	return gh.Call("hasEventListener", typ, listener).Bool()
}
func (gh *GridHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: gh.Call("localToWorld", vector)}
}
func (gh *GridHelper) LookAt(vector *Vector3, y float64, z float64) {
	gh.Call("lookAt", vector, y, z)
}
func (gh *GridHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	gh.Call("raycast", raycaster, intersects)
}
func (gh *GridHelper) Remove(object js.Value) *GridHelper {
	return &GridHelper{Value: gh.Call("remove", object)}
}
func (gh *GridHelper) RemoveEventListener(typ string, listener js.Value) {
	gh.Call("removeEventListener", typ, listener)
}
func (gh *GridHelper) RotateOnAxis(axis *Vector3, angle float64) *GridHelper {
	return &GridHelper{Value: gh.Call("rotateOnAxis", axis, angle)}
}
func (gh *GridHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *GridHelper {
	return &GridHelper{Value: gh.Call("rotateOnWorldAxis", axis, angle)}
}
func (gh *GridHelper) RotateX(angle float64) *GridHelper {
	return &GridHelper{Value: gh.Call("rotateX", angle)}
}
func (gh *GridHelper) RotateY(angle float64) *GridHelper {
	return &GridHelper{Value: gh.Call("rotateY", angle)}
}
func (gh *GridHelper) RotateZ(angle float64) *GridHelper {
	return &GridHelper{Value: gh.Call("rotateZ", angle)}
}
func (gh *GridHelper) SetColors(color1 *Color, color2 *Color) {
	gh.Call("setColors", color1, color2)
}
func (gh *GridHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	gh.Call("setRotationFromAxisAngle", axis, angle)
}
func (gh *GridHelper) SetRotationFromEuler(euler *Euler) {
	gh.Call("setRotationFromEuler", euler)
}
func (gh *GridHelper) SetRotationFromMatrix(m *Matrix4) {
	gh.Call("setRotationFromMatrix", m)
}
func (gh *GridHelper) SetRotationFromQuaternion(q *Quaternion) {
	gh.Call("setRotationFromQuaternion", q)
}
func (gh *GridHelper) ToJSON(meta js.Value) js.Value {
	return gh.Call("toJSON", meta)
}
func (gh *GridHelper) TranslateOnAxis(axis *Vector3, distance float64) *GridHelper {
	return &GridHelper{Value: gh.Call("translateOnAxis", axis, distance)}
}
func (gh *GridHelper) TranslateX(distance float64) *GridHelper {
	return &GridHelper{Value: gh.Call("translateX", distance)}
}
func (gh *GridHelper) TranslateY(distance float64) *GridHelper {
	return &GridHelper{Value: gh.Call("translateY", distance)}
}
func (gh *GridHelper) TranslateZ(distance float64) *GridHelper {
	return &GridHelper{Value: gh.Call("translateZ", distance)}
}
func (gh *GridHelper) Traverse(callback js.Value) {
	gh.Call("traverse", callback)
}
func (gh *GridHelper) TraverseAncestors(callback js.Value) {
	gh.Call("traverseAncestors", callback)
}
func (gh *GridHelper) TraverseVisible(callback js.Value) {
	gh.Call("traverseVisible", callback)
}
func (gh *GridHelper) UpdateMatrix() {
	gh.Call("updateMatrix")
}
func (gh *GridHelper) UpdateMatrixWorld(force bool) {
	gh.Call("updateMatrixWorld", force)
}
func (gh *GridHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	gh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (gh *GridHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: gh.Call("worldToLocal", vector)}
}

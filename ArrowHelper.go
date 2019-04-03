// Code generated from "helpers/ArrowHelper.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type ArrowHelper struct {
	js.Value
}

func NewArrowHelper(dir *Vector3, origin *Vector3, length float64, hex int, headLength float64, headWidth float64) *ArrowHelper {
	return &ArrowHelper{Value: get("ArrowHelper").New(dir, origin, length, hex, headLength, headWidth)}
}
func (ah *ArrowHelper) CastShadow() bool {
	return ah.Get("castShadow").Bool()
}
func (ah *ArrowHelper) SetCastShadow(v bool) {
	ah.Set("castShadow", v)
}
func (ah *ArrowHelper) Children() js.Value {
	return ah.Get("children")
}
func (ah *ArrowHelper) SetChildren(v js.Value) {
	ah.Set("children", v)
}
func (ah *ArrowHelper) Cone() *Mesh {
	return &Mesh{Value: ah.Get("cone")}
}
func (ah *ArrowHelper) SetCone(v *Mesh) {
	ah.Set("cone", v)
}
func (ah *ArrowHelper) FrustumCulled() bool {
	return ah.Get("frustumCulled").Bool()
}
func (ah *ArrowHelper) SetFrustumCulled(v bool) {
	ah.Set("frustumCulled", v)
}
func (ah *ArrowHelper) Id() int {
	return ah.Get("id").Int()
}
func (ah *ArrowHelper) SetId(v int) {
	ah.Set("id", v)
}
func (ah *ArrowHelper) IsObject3D() bool {
	return ah.Get("isObject3D").Bool()
}
func (ah *ArrowHelper) SetIsObject3D(v bool) {
	ah.Set("isObject3D", v)
}
func (ah *ArrowHelper) Layers() *Layers {
	return &Layers{Value: ah.Get("layers")}
}
func (ah *ArrowHelper) SetLayers(v *Layers) {
	ah.Set("layers", v)
}
func (ah *ArrowHelper) Line() *Line {
	return &Line{Value: ah.Get("line")}
}
func (ah *ArrowHelper) SetLine(v *Line) {
	ah.Set("line", v)
}
func (ah *ArrowHelper) Matrix() *Matrix4 {
	return &Matrix4{Value: ah.Get("matrix")}
}
func (ah *ArrowHelper) SetMatrix(v *Matrix4) {
	ah.Set("matrix", v)
}
func (ah *ArrowHelper) MatrixAutoUpdate() bool {
	return ah.Get("matrixAutoUpdate").Bool()
}
func (ah *ArrowHelper) SetMatrixAutoUpdate(v bool) {
	ah.Set("matrixAutoUpdate", v)
}
func (ah *ArrowHelper) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ah.Get("matrixWorld")}
}
func (ah *ArrowHelper) SetMatrixWorld(v *Matrix4) {
	ah.Set("matrixWorld", v)
}
func (ah *ArrowHelper) MatrixWorldNeedsUpdate() bool {
	return ah.Get("matrixWorldNeedsUpdate").Bool()
}
func (ah *ArrowHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ah.Set("matrixWorldNeedsUpdate", v)
}
func (ah *ArrowHelper) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ah.Get("modelViewMatrix")}
}
func (ah *ArrowHelper) SetModelViewMatrix(v *Matrix4) {
	ah.Set("modelViewMatrix", v)
}
func (ah *ArrowHelper) Name() string {
	return ah.Get("name").String()
}
func (ah *ArrowHelper) SetName(v string) {
	ah.Set("name", v)
}
func (ah *ArrowHelper) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ah.Get("normalMatrix")}
}
func (ah *ArrowHelper) SetNormalMatrix(v *Matrix3) {
	ah.Set("normalMatrix", v)
}
func (ah *ArrowHelper) OnAfterRender() js.Value {
	return ah.Get("onAfterRender")
}
func (ah *ArrowHelper) SetOnAfterRender(v js.Value) {
	ah.Set("onAfterRender", v)
}
func (ah *ArrowHelper) OnBeforeRender() js.Value {
	return ah.Get("onBeforeRender")
}
func (ah *ArrowHelper) SetOnBeforeRender(v js.Value) {
	ah.Set("onBeforeRender", v)
}
func (ah *ArrowHelper) Parent() *Object3D {
	return &Object3D{Value: ah.Get("parent")}
}
func (ah *ArrowHelper) SetParent(v *Object3D) {
	ah.Set("parent", v)
}
func (ah *ArrowHelper) Position() *Vector3 {
	return &Vector3{Value: ah.Get("position")}
}
func (ah *ArrowHelper) SetPosition(v *Vector3) {
	ah.Set("position", v)
}
func (ah *ArrowHelper) Quaternion() *Quaternion {
	return &Quaternion{Value: ah.Get("quaternion")}
}
func (ah *ArrowHelper) SetQuaternion(v *Quaternion) {
	ah.Set("quaternion", v)
}
func (ah *ArrowHelper) ReceiveShadow() bool {
	return ah.Get("receiveShadow").Bool()
}
func (ah *ArrowHelper) SetReceiveShadow(v bool) {
	ah.Set("receiveShadow", v)
}
func (ah *ArrowHelper) RenderOrder() float64 {
	return ah.Get("renderOrder").Float()
}
func (ah *ArrowHelper) SetRenderOrder(v float64) {
	ah.Set("renderOrder", v)
}
func (ah *ArrowHelper) Rotation() *Euler {
	return &Euler{Value: ah.Get("rotation")}
}
func (ah *ArrowHelper) SetRotation(v *Euler) {
	ah.Set("rotation", v)
}
func (ah *ArrowHelper) Scale() *Vector3 {
	return &Vector3{Value: ah.Get("scale")}
}
func (ah *ArrowHelper) SetScale(v *Vector3) {
	ah.Set("scale", v)
}
func (ah *ArrowHelper) Type() string {
	return ah.Get("type").String()
}
func (ah *ArrowHelper) SetType(v string) {
	ah.Set("type", v)
}
func (ah *ArrowHelper) Up() *Vector3 {
	return &Vector3{Value: ah.Get("up")}
}
func (ah *ArrowHelper) SetUp(v *Vector3) {
	ah.Set("up", v)
}
func (ah *ArrowHelper) UserData() js.Value {
	return ah.Get("userData")
}
func (ah *ArrowHelper) SetUserData(v js.Value) {
	ah.Set("userData", v)
}
func (ah *ArrowHelper) Uuid() string {
	return ah.Get("uuid").String()
}
func (ah *ArrowHelper) SetUuid(v string) {
	ah.Set("uuid", v)
}
func (ah *ArrowHelper) Visible() bool {
	return ah.Get("visible").Bool()
}
func (ah *ArrowHelper) SetVisible(v bool) {
	ah.Set("visible", v)
}
func (ah *ArrowHelper) DefaultMatrixAutoUpdate() bool {
	return ah.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ah *ArrowHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ah.Set("DefaultMatrixAutoUpdate", v)
}
func (ah *ArrowHelper) DefaultUp() *Vector3 {
	return &Vector3{Value: ah.Get("DefaultUp")}
}
func (ah *ArrowHelper) SetDefaultUp(v *Vector3) {
	ah.Set("DefaultUp", v)
}
func (ah *ArrowHelper) Add(object js.Value) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("add", object)}
}
func (ah *ArrowHelper) AddEventListener(typ string, listener js.Value) {
	ah.Call("addEventListener", typ, listener)
}
func (ah *ArrowHelper) ApplyMatrix(matrix *Matrix4) {
	ah.Call("applyMatrix", matrix)
}
func (ah *ArrowHelper) ApplyQuaternion(quaternion *Quaternion) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("applyQuaternion", quaternion)}
}
func (ah *ArrowHelper) Clone(recursive bool) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("clone", recursive)}
}
func (ah *ArrowHelper) Copy(source *Object3D, recursive bool) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("copy", source, recursive)}
}
func (ah *ArrowHelper) DispatchEvent(event js.Value) {
	ah.Call("dispatchEvent", event)
}
func (ah *ArrowHelper) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ah.Call("getObjectById", id)}
}
func (ah *ArrowHelper) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ah.Call("getObjectByName", name)}
}
func (ah *ArrowHelper) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ah.Call("getObjectByProperty", name, value)}
}
func (ah *ArrowHelper) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("getWorldDirection", target)}
}
func (ah *ArrowHelper) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("getWorldPosition", target)}
}
func (ah *ArrowHelper) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ah.Call("getWorldQuaternion", target)}
}
func (ah *ArrowHelper) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("getWorldScale", target)}
}
func (ah *ArrowHelper) HasEventListener(typ string, listener js.Value) bool {
	return ah.Call("hasEventListener", typ, listener).Bool()
}
func (ah *ArrowHelper) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("localToWorld", vector)}
}
func (ah *ArrowHelper) LookAt(vector *Vector3, y float64, z float64) {
	ah.Call("lookAt", vector, y, z)
}
func (ah *ArrowHelper) Raycast(raycaster *Raycaster, intersects js.Value) {
	ah.Call("raycast", raycaster, intersects)
}
func (ah *ArrowHelper) Remove(object js.Value) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("remove", object)}
}
func (ah *ArrowHelper) RemoveEventListener(typ string, listener js.Value) {
	ah.Call("removeEventListener", typ, listener)
}
func (ah *ArrowHelper) RotateOnAxis(axis *Vector3, angle float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("rotateOnAxis", axis, angle)}
}
func (ah *ArrowHelper) RotateOnWorldAxis(axis *Vector3, angle float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("rotateOnWorldAxis", axis, angle)}
}
func (ah *ArrowHelper) RotateX(angle float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("rotateX", angle)}
}
func (ah *ArrowHelper) RotateY(angle float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("rotateY", angle)}
}
func (ah *ArrowHelper) RotateZ(angle float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("rotateZ", angle)}
}
func (ah *ArrowHelper) SetColor(color *Color) {
	ah.Call("setColor", color)
}
func (ah *ArrowHelper) SetDirection(dir *Vector3) {
	ah.Call("setDirection", dir)
}
func (ah *ArrowHelper) SetLength(length float64, headLength float64, headWidth float64) {
	ah.Call("setLength", length, headLength, headWidth)
}
func (ah *ArrowHelper) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ah.Call("setRotationFromAxisAngle", axis, angle)
}
func (ah *ArrowHelper) SetRotationFromEuler(euler *Euler) {
	ah.Call("setRotationFromEuler", euler)
}
func (ah *ArrowHelper) SetRotationFromMatrix(m *Matrix4) {
	ah.Call("setRotationFromMatrix", m)
}
func (ah *ArrowHelper) SetRotationFromQuaternion(q *Quaternion) {
	ah.Call("setRotationFromQuaternion", q)
}
func (ah *ArrowHelper) ToJSON(meta js.Value) js.Value {
	return ah.Call("toJSON", meta)
}
func (ah *ArrowHelper) TranslateOnAxis(axis *Vector3, distance float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("translateOnAxis", axis, distance)}
}
func (ah *ArrowHelper) TranslateX(distance float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("translateX", distance)}
}
func (ah *ArrowHelper) TranslateY(distance float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("translateY", distance)}
}
func (ah *ArrowHelper) TranslateZ(distance float64) *ArrowHelper {
	return &ArrowHelper{Value: ah.Call("translateZ", distance)}
}
func (ah *ArrowHelper) Traverse(callback js.Value) {
	ah.Call("traverse", callback)
}
func (ah *ArrowHelper) TraverseAncestors(callback js.Value) {
	ah.Call("traverseAncestors", callback)
}
func (ah *ArrowHelper) TraverseVisible(callback js.Value) {
	ah.Call("traverseVisible", callback)
}
func (ah *ArrowHelper) UpdateMatrix() {
	ah.Call("updateMatrix")
}
func (ah *ArrowHelper) UpdateMatrixWorld(force bool) {
	ah.Call("updateMatrixWorld", force)
}
func (ah *ArrowHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ah.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ah *ArrowHelper) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ah.Call("worldToLocal", vector)}
}

// Code generated from "objects/LOD.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type LOD struct {
	js.Value
}

func NewLOD() *LOD {
	return &LOD{Value: get("LOD").New()}
}
func (lod *LOD) CastShadow() bool {
	return lod.Get("castShadow").Bool()
}
func (lod *LOD) SetCastShadow(v bool) {
	lod.Set("castShadow", v)
}
func (lod *LOD) Children() js.Value {
	return lod.Get("children")
}
func (lod *LOD) SetChildren(v js.Value) {
	lod.Set("children", v)
}
func (lod *LOD) FrustumCulled() bool {
	return lod.Get("frustumCulled").Bool()
}
func (lod *LOD) SetFrustumCulled(v bool) {
	lod.Set("frustumCulled", v)
}
func (lod *LOD) Id() int {
	return lod.Get("id").Int()
}
func (lod *LOD) SetId(v int) {
	lod.Set("id", v)
}
func (lod *LOD) IsObject3D() bool {
	return lod.Get("isObject3D").Bool()
}
func (lod *LOD) SetIsObject3D(v bool) {
	lod.Set("isObject3D", v)
}
func (lod *LOD) Layers() *Layers {
	return &Layers{Value: lod.Get("layers")}
}
func (lod *LOD) SetLayers(v *Layers) {
	lod.Set("layers", v)
}
func (lod *LOD) Levels() js.Value {
	return lod.Get("levels")
}
func (lod *LOD) SetLevels(v js.Value) {
	lod.Set("levels", v)
}
func (lod *LOD) Matrix() *Matrix4 {
	return &Matrix4{Value: lod.Get("matrix")}
}
func (lod *LOD) SetMatrix(v *Matrix4) {
	lod.Set("matrix", v)
}
func (lod *LOD) MatrixAutoUpdate() bool {
	return lod.Get("matrixAutoUpdate").Bool()
}
func (lod *LOD) SetMatrixAutoUpdate(v bool) {
	lod.Set("matrixAutoUpdate", v)
}
func (lod *LOD) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: lod.Get("matrixWorld")}
}
func (lod *LOD) SetMatrixWorld(v *Matrix4) {
	lod.Set("matrixWorld", v)
}
func (lod *LOD) MatrixWorldNeedsUpdate() bool {
	return lod.Get("matrixWorldNeedsUpdate").Bool()
}
func (lod *LOD) SetMatrixWorldNeedsUpdate(v bool) {
	lod.Set("matrixWorldNeedsUpdate", v)
}
func (lod *LOD) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: lod.Get("modelViewMatrix")}
}
func (lod *LOD) SetModelViewMatrix(v *Matrix4) {
	lod.Set("modelViewMatrix", v)
}
func (lod *LOD) Name() string {
	return lod.Get("name").String()
}
func (lod *LOD) SetName(v string) {
	lod.Set("name", v)
}
func (lod *LOD) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: lod.Get("normalMatrix")}
}
func (lod *LOD) SetNormalMatrix(v *Matrix3) {
	lod.Set("normalMatrix", v)
}
func (lod *LOD) Objects() js.Value {
	return lod.Get("objects")
}
func (lod *LOD) SetObjects(v js.Value) {
	lod.Set("objects", v)
}
func (lod *LOD) OnAfterRender() js.Value {
	return lod.Get("onAfterRender")
}
func (lod *LOD) SetOnAfterRender(v js.Value) {
	lod.Set("onAfterRender", v)
}
func (lod *LOD) OnBeforeRender() js.Value {
	return lod.Get("onBeforeRender")
}
func (lod *LOD) SetOnBeforeRender(v js.Value) {
	lod.Set("onBeforeRender", v)
}
func (lod *LOD) Parent() *Object3D {
	return &Object3D{Value: lod.Get("parent")}
}
func (lod *LOD) SetParent(v *Object3D) {
	lod.Set("parent", v)
}
func (lod *LOD) Position() *Vector3 {
	return &Vector3{Value: lod.Get("position")}
}
func (lod *LOD) SetPosition(v *Vector3) {
	lod.Set("position", v)
}
func (lod *LOD) Quaternion() *Quaternion {
	return &Quaternion{Value: lod.Get("quaternion")}
}
func (lod *LOD) SetQuaternion(v *Quaternion) {
	lod.Set("quaternion", v)
}
func (lod *LOD) ReceiveShadow() bool {
	return lod.Get("receiveShadow").Bool()
}
func (lod *LOD) SetReceiveShadow(v bool) {
	lod.Set("receiveShadow", v)
}
func (lod *LOD) RenderOrder() float64 {
	return lod.Get("renderOrder").Float()
}
func (lod *LOD) SetRenderOrder(v float64) {
	lod.Set("renderOrder", v)
}
func (lod *LOD) Rotation() *Euler {
	return &Euler{Value: lod.Get("rotation")}
}
func (lod *LOD) SetRotation(v *Euler) {
	lod.Set("rotation", v)
}
func (lod *LOD) Scale() *Vector3 {
	return &Vector3{Value: lod.Get("scale")}
}
func (lod *LOD) SetScale(v *Vector3) {
	lod.Set("scale", v)
}
func (lod *LOD) Type() string {
	return lod.Get("type").String()
}
func (lod *LOD) SetType(v string) {
	lod.Set("type", v)
}
func (lod *LOD) Up() *Vector3 {
	return &Vector3{Value: lod.Get("up")}
}
func (lod *LOD) SetUp(v *Vector3) {
	lod.Set("up", v)
}
func (lod *LOD) UserData() js.Value {
	return lod.Get("userData")
}
func (lod *LOD) SetUserData(v js.Value) {
	lod.Set("userData", v)
}
func (lod *LOD) Uuid() string {
	return lod.Get("uuid").String()
}
func (lod *LOD) SetUuid(v string) {
	lod.Set("uuid", v)
}
func (lod *LOD) Visible() bool {
	return lod.Get("visible").Bool()
}
func (lod *LOD) SetVisible(v bool) {
	lod.Set("visible", v)
}
func (lod *LOD) DefaultMatrixAutoUpdate() bool {
	return lod.Get("DefaultMatrixAutoUpdate").Bool()
}
func (lod *LOD) SetDefaultMatrixAutoUpdate(v bool) {
	lod.Set("DefaultMatrixAutoUpdate", v)
}
func (lod *LOD) DefaultUp() *Vector3 {
	return &Vector3{Value: lod.Get("DefaultUp")}
}
func (lod *LOD) SetDefaultUp(v *Vector3) {
	lod.Set("DefaultUp", v)
}
func (lod *LOD) Add(object js.Value) *LOD {
	return &LOD{Value: lod.Call("add", object)}
}
func (lod *LOD) AddEventListener(typ string, listener js.Value) {
	lod.Call("addEventListener", typ, listener)
}
func (lod *LOD) AddLevel(object *Object3D, distance float64) *LOD {
	return &LOD{Value: lod.Call("addLevel", object, distance)}
}
func (lod *LOD) ApplyMatrix(matrix *Matrix4) {
	lod.Call("applyMatrix", matrix)
}
func (lod *LOD) ApplyQuaternion(quaternion *Quaternion) *LOD {
	return &LOD{Value: lod.Call("applyQuaternion", quaternion)}
}
func (lod *LOD) Clone(recursive bool) *LOD {
	return &LOD{Value: lod.Call("clone", recursive)}
}
func (lod *LOD) Copy(source *Object3D, recursive bool) *LOD {
	return &LOD{Value: lod.Call("copy", source, recursive)}
}
func (lod *LOD) DispatchEvent(event js.Value) {
	lod.Call("dispatchEvent", event)
}
func (lod *LOD) GetObjectById(id int) *Object3D {
	return &Object3D{Value: lod.Call("getObjectById", id)}
}
func (lod *LOD) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: lod.Call("getObjectByName", name)}
}
func (lod *LOD) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: lod.Call("getObjectByProperty", name, value)}
}
func (lod *LOD) GetObjectForDistance(distance float64) *Object3D {
	return &Object3D{Value: lod.Call("getObjectForDistance", distance)}
}
func (lod *LOD) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: lod.Call("getWorldDirection", target)}
}
func (lod *LOD) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: lod.Call("getWorldPosition", target)}
}
func (lod *LOD) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: lod.Call("getWorldQuaternion", target)}
}
func (lod *LOD) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: lod.Call("getWorldScale", target)}
}
func (lod *LOD) HasEventListener(typ string, listener js.Value) bool {
	return lod.Call("hasEventListener", typ, listener).Bool()
}
func (lod *LOD) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: lod.Call("localToWorld", vector)}
}
func (lod *LOD) LookAt(vector *Vector3, y float64, z float64) {
	lod.Call("lookAt", vector, y, z)
}
func (lod *LOD) Raycast(raycaster *Raycaster, intersects js.Value) {
	lod.Call("raycast", raycaster, intersects)
}
func (lod *LOD) Remove(object js.Value) *LOD {
	return &LOD{Value: lod.Call("remove", object)}
}
func (lod *LOD) RemoveEventListener(typ string, listener js.Value) {
	lod.Call("removeEventListener", typ, listener)
}
func (lod *LOD) RotateOnAxis(axis *Vector3, angle float64) *LOD {
	return &LOD{Value: lod.Call("rotateOnAxis", axis, angle)}
}
func (lod *LOD) RotateOnWorldAxis(axis *Vector3, angle float64) *LOD {
	return &LOD{Value: lod.Call("rotateOnWorldAxis", axis, angle)}
}
func (lod *LOD) RotateX(angle float64) *LOD {
	return &LOD{Value: lod.Call("rotateX", angle)}
}
func (lod *LOD) RotateY(angle float64) *LOD {
	return &LOD{Value: lod.Call("rotateY", angle)}
}
func (lod *LOD) RotateZ(angle float64) *LOD {
	return &LOD{Value: lod.Call("rotateZ", angle)}
}
func (lod *LOD) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	lod.Call("setRotationFromAxisAngle", axis, angle)
}
func (lod *LOD) SetRotationFromEuler(euler *Euler) {
	lod.Call("setRotationFromEuler", euler)
}
func (lod *LOD) SetRotationFromMatrix(m *Matrix4) {
	lod.Call("setRotationFromMatrix", m)
}
func (lod *LOD) SetRotationFromQuaternion(q *Quaternion) {
	lod.Call("setRotationFromQuaternion", q)
}
func (lod *LOD) ToJSON(meta js.Value) js.Value {
	return lod.Call("toJSON", meta)
}
func (lod *LOD) TranslateOnAxis(axis *Vector3, distance float64) *LOD {
	return &LOD{Value: lod.Call("translateOnAxis", axis, distance)}
}
func (lod *LOD) TranslateX(distance float64) *LOD {
	return &LOD{Value: lod.Call("translateX", distance)}
}
func (lod *LOD) TranslateY(distance float64) *LOD {
	return &LOD{Value: lod.Call("translateY", distance)}
}
func (lod *LOD) TranslateZ(distance float64) *LOD {
	return &LOD{Value: lod.Call("translateZ", distance)}
}
func (lod *LOD) Traverse(callback js.Value) {
	lod.Call("traverse", callback)
}
func (lod *LOD) TraverseAncestors(callback js.Value) {
	lod.Call("traverseAncestors", callback)
}
func (lod *LOD) TraverseVisible(callback js.Value) {
	lod.Call("traverseVisible", callback)
}
func (lod *LOD) Update(camera *Camera) {
	lod.Call("update", camera)
}
func (lod *LOD) UpdateMatrix() {
	lod.Call("updateMatrix")
}
func (lod *LOD) UpdateMatrixWorld(force bool) {
	lod.Call("updateMatrixWorld", force)
}
func (lod *LOD) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	lod.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (lod *LOD) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: lod.Call("worldToLocal", vector)}
}

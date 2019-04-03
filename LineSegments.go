// Code generated from "objects/LineSegments.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

func LinePieces() float64 {
	return get("LinePieces").Float()
}
func SetLinePieces(v float64) {
	set("LinePieces", v)
}
func LineStrip() float64 {
	return get("LineStrip").Float()
}
func SetLineStrip(v float64) {
	set("LineStrip", v)
}

type LineSegments struct {
	js.Value
}

func NewLineSegments(geometry *Geometry, material *Material, mode float64) *LineSegments {
	return &LineSegments{Value: get("LineSegments").New(geometry, material, mode)}
}
func (ls *LineSegments) CastShadow() bool {
	return ls.Get("castShadow").Bool()
}
func (ls *LineSegments) SetCastShadow(v bool) {
	ls.Set("castShadow", v)
}
func (ls *LineSegments) Children() js.Value {
	return ls.Get("children")
}
func (ls *LineSegments) SetChildren(v js.Value) {
	ls.Set("children", v)
}
func (ls *LineSegments) FrustumCulled() bool {
	return ls.Get("frustumCulled").Bool()
}
func (ls *LineSegments) SetFrustumCulled(v bool) {
	ls.Set("frustumCulled", v)
}
func (ls *LineSegments) Geometry() *Geometry {
	return &Geometry{Value: ls.Get("geometry")}
}
func (ls *LineSegments) SetGeometry(v *Geometry) {
	ls.Set("geometry", v)
}
func (ls *LineSegments) Id() int {
	return ls.Get("id").Int()
}
func (ls *LineSegments) SetId(v int) {
	ls.Set("id", v)
}
func (ls *LineSegments) IsLine() bool {
	return ls.Get("isLine").Bool()
}
func (ls *LineSegments) SetIsLine(v bool) {
	ls.Set("isLine", v)
}
func (ls *LineSegments) IsLineSegments() bool {
	return ls.Get("isLineSegments").Bool()
}
func (ls *LineSegments) SetIsLineSegments(v bool) {
	ls.Set("isLineSegments", v)
}
func (ls *LineSegments) IsObject3D() bool {
	return ls.Get("isObject3D").Bool()
}
func (ls *LineSegments) SetIsObject3D(v bool) {
	ls.Set("isObject3D", v)
}
func (ls *LineSegments) Layers() *Layers {
	return &Layers{Value: ls.Get("layers")}
}
func (ls *LineSegments) SetLayers(v *Layers) {
	ls.Set("layers", v)
}
func (ls *LineSegments) Material() *Material {
	return &Material{Value: ls.Get("material")}
}
func (ls *LineSegments) SetMaterial(v *Material) {
	ls.Set("material", v)
}
func (ls *LineSegments) Matrix() *Matrix4 {
	return &Matrix4{Value: ls.Get("matrix")}
}
func (ls *LineSegments) SetMatrix(v *Matrix4) {
	ls.Set("matrix", v)
}
func (ls *LineSegments) MatrixAutoUpdate() bool {
	return ls.Get("matrixAutoUpdate").Bool()
}
func (ls *LineSegments) SetMatrixAutoUpdate(v bool) {
	ls.Set("matrixAutoUpdate", v)
}
func (ls *LineSegments) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ls.Get("matrixWorld")}
}
func (ls *LineSegments) SetMatrixWorld(v *Matrix4) {
	ls.Set("matrixWorld", v)
}
func (ls *LineSegments) MatrixWorldNeedsUpdate() bool {
	return ls.Get("matrixWorldNeedsUpdate").Bool()
}
func (ls *LineSegments) SetMatrixWorldNeedsUpdate(v bool) {
	ls.Set("matrixWorldNeedsUpdate", v)
}
func (ls *LineSegments) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ls.Get("modelViewMatrix")}
}
func (ls *LineSegments) SetModelViewMatrix(v *Matrix4) {
	ls.Set("modelViewMatrix", v)
}
func (ls *LineSegments) Name() string {
	return ls.Get("name").String()
}
func (ls *LineSegments) SetName(v string) {
	ls.Set("name", v)
}
func (ls *LineSegments) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ls.Get("normalMatrix")}
}
func (ls *LineSegments) SetNormalMatrix(v *Matrix3) {
	ls.Set("normalMatrix", v)
}
func (ls *LineSegments) OnAfterRender() js.Value {
	return ls.Get("onAfterRender")
}
func (ls *LineSegments) SetOnAfterRender(v js.Value) {
	ls.Set("onAfterRender", v)
}
func (ls *LineSegments) OnBeforeRender() js.Value {
	return ls.Get("onBeforeRender")
}
func (ls *LineSegments) SetOnBeforeRender(v js.Value) {
	ls.Set("onBeforeRender", v)
}
func (ls *LineSegments) Parent() *Object3D {
	return &Object3D{Value: ls.Get("parent")}
}
func (ls *LineSegments) SetParent(v *Object3D) {
	ls.Set("parent", v)
}
func (ls *LineSegments) Position() *Vector3 {
	return &Vector3{Value: ls.Get("position")}
}
func (ls *LineSegments) SetPosition(v *Vector3) {
	ls.Set("position", v)
}
func (ls *LineSegments) Quaternion() *Quaternion {
	return &Quaternion{Value: ls.Get("quaternion")}
}
func (ls *LineSegments) SetQuaternion(v *Quaternion) {
	ls.Set("quaternion", v)
}
func (ls *LineSegments) ReceiveShadow() bool {
	return ls.Get("receiveShadow").Bool()
}
func (ls *LineSegments) SetReceiveShadow(v bool) {
	ls.Set("receiveShadow", v)
}
func (ls *LineSegments) RenderOrder() float64 {
	return ls.Get("renderOrder").Float()
}
func (ls *LineSegments) SetRenderOrder(v float64) {
	ls.Set("renderOrder", v)
}
func (ls *LineSegments) Rotation() *Euler {
	return &Euler{Value: ls.Get("rotation")}
}
func (ls *LineSegments) SetRotation(v *Euler) {
	ls.Set("rotation", v)
}
func (ls *LineSegments) Scale() *Vector3 {
	return &Vector3{Value: ls.Get("scale")}
}
func (ls *LineSegments) SetScale(v *Vector3) {
	ls.Set("scale", v)
}
func (ls *LineSegments) Type() string {
	return ls.Get("type").String()
}
func (ls *LineSegments) SetType(v string) {
	ls.Set("type", v)
}
func (ls *LineSegments) Up() *Vector3 {
	return &Vector3{Value: ls.Get("up")}
}
func (ls *LineSegments) SetUp(v *Vector3) {
	ls.Set("up", v)
}
func (ls *LineSegments) UserData() js.Value {
	return ls.Get("userData")
}
func (ls *LineSegments) SetUserData(v js.Value) {
	ls.Set("userData", v)
}
func (ls *LineSegments) Uuid() string {
	return ls.Get("uuid").String()
}
func (ls *LineSegments) SetUuid(v string) {
	ls.Set("uuid", v)
}
func (ls *LineSegments) Visible() bool {
	return ls.Get("visible").Bool()
}
func (ls *LineSegments) SetVisible(v bool) {
	ls.Set("visible", v)
}
func (ls *LineSegments) DefaultMatrixAutoUpdate() bool {
	return ls.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ls *LineSegments) SetDefaultMatrixAutoUpdate(v bool) {
	ls.Set("DefaultMatrixAutoUpdate", v)
}
func (ls *LineSegments) DefaultUp() *Vector3 {
	return &Vector3{Value: ls.Get("DefaultUp")}
}
func (ls *LineSegments) SetDefaultUp(v *Vector3) {
	ls.Set("DefaultUp", v)
}
func (ls *LineSegments) Add(object js.Value) *LineSegments {
	return &LineSegments{Value: ls.Call("add", object)}
}
func (ls *LineSegments) AddEventListener(typ string, listener js.Value) {
	ls.Call("addEventListener", typ, listener)
}
func (ls *LineSegments) ApplyMatrix(matrix *Matrix4) {
	ls.Call("applyMatrix", matrix)
}
func (ls *LineSegments) ApplyQuaternion(quaternion *Quaternion) *LineSegments {
	return &LineSegments{Value: ls.Call("applyQuaternion", quaternion)}
}
func (ls *LineSegments) Clone(recursive bool) *LineSegments {
	return &LineSegments{Value: ls.Call("clone", recursive)}
}
func (ls *LineSegments) ComputeLineDistances() *LineSegments {
	return &LineSegments{Value: ls.Call("computeLineDistances")}
}
func (ls *LineSegments) Copy(source *Object3D, recursive bool) *LineSegments {
	return &LineSegments{Value: ls.Call("copy", source, recursive)}
}
func (ls *LineSegments) DispatchEvent(event js.Value) {
	ls.Call("dispatchEvent", event)
}
func (ls *LineSegments) GetObjectById(id int) *Object3D {
	return &Object3D{Value: ls.Call("getObjectById", id)}
}
func (ls *LineSegments) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ls.Call("getObjectByName", name)}
}
func (ls *LineSegments) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ls.Call("getObjectByProperty", name, value)}
}
func (ls *LineSegments) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ls.Call("getWorldDirection", target)}
}
func (ls *LineSegments) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ls.Call("getWorldPosition", target)}
}
func (ls *LineSegments) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ls.Call("getWorldQuaternion", target)}
}
func (ls *LineSegments) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ls.Call("getWorldScale", target)}
}
func (ls *LineSegments) HasEventListener(typ string, listener js.Value) bool {
	return ls.Call("hasEventListener", typ, listener).Bool()
}
func (ls *LineSegments) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ls.Call("localToWorld", vector)}
}
func (ls *LineSegments) LookAt(vector *Vector3, y float64, z float64) {
	ls.Call("lookAt", vector, y, z)
}
func (ls *LineSegments) Raycast(raycaster *Raycaster, intersects js.Value) {
	ls.Call("raycast", raycaster, intersects)
}
func (ls *LineSegments) Remove(object js.Value) *LineSegments {
	return &LineSegments{Value: ls.Call("remove", object)}
}
func (ls *LineSegments) RemoveEventListener(typ string, listener js.Value) {
	ls.Call("removeEventListener", typ, listener)
}
func (ls *LineSegments) RotateOnAxis(axis *Vector3, angle float64) *LineSegments {
	return &LineSegments{Value: ls.Call("rotateOnAxis", axis, angle)}
}
func (ls *LineSegments) RotateOnWorldAxis(axis *Vector3, angle float64) *LineSegments {
	return &LineSegments{Value: ls.Call("rotateOnWorldAxis", axis, angle)}
}
func (ls *LineSegments) RotateX(angle float64) *LineSegments {
	return &LineSegments{Value: ls.Call("rotateX", angle)}
}
func (ls *LineSegments) RotateY(angle float64) *LineSegments {
	return &LineSegments{Value: ls.Call("rotateY", angle)}
}
func (ls *LineSegments) RotateZ(angle float64) *LineSegments {
	return &LineSegments{Value: ls.Call("rotateZ", angle)}
}
func (ls *LineSegments) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ls.Call("setRotationFromAxisAngle", axis, angle)
}
func (ls *LineSegments) SetRotationFromEuler(euler *Euler) {
	ls.Call("setRotationFromEuler", euler)
}
func (ls *LineSegments) SetRotationFromMatrix(m *Matrix4) {
	ls.Call("setRotationFromMatrix", m)
}
func (ls *LineSegments) SetRotationFromQuaternion(q *Quaternion) {
	ls.Call("setRotationFromQuaternion", q)
}
func (ls *LineSegments) ToJSON(meta js.Value) js.Value {
	return ls.Call("toJSON", meta)
}
func (ls *LineSegments) TranslateOnAxis(axis *Vector3, distance float64) *LineSegments {
	return &LineSegments{Value: ls.Call("translateOnAxis", axis, distance)}
}
func (ls *LineSegments) TranslateX(distance float64) *LineSegments {
	return &LineSegments{Value: ls.Call("translateX", distance)}
}
func (ls *LineSegments) TranslateY(distance float64) *LineSegments {
	return &LineSegments{Value: ls.Call("translateY", distance)}
}
func (ls *LineSegments) TranslateZ(distance float64) *LineSegments {
	return &LineSegments{Value: ls.Call("translateZ", distance)}
}
func (ls *LineSegments) Traverse(callback js.Value) {
	ls.Call("traverse", callback)
}
func (ls *LineSegments) TraverseAncestors(callback js.Value) {
	ls.Call("traverseAncestors", callback)
}
func (ls *LineSegments) TraverseVisible(callback js.Value) {
	ls.Call("traverseVisible", callback)
}
func (ls *LineSegments) UpdateMatrix() {
	ls.Call("updateMatrix")
}
func (ls *LineSegments) UpdateMatrixWorld(force bool) {
	ls.Call("updateMatrixWorld", force)
}
func (ls *LineSegments) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ls.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ls *LineSegments) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ls.Call("worldToLocal", vector)}
}

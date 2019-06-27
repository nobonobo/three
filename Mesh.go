// Code generated from "objects/Mesh.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Mesh extend: [Object3D]
type Mesh struct {
	js.Value
}

func NewMesh(geometry Geometry, material Material) *Mesh {
	return &Mesh{Value: get("Mesh").New(geometry.JSValue(), material.JSValue())}
}
func (mm *Mesh) JSValue() js.Value {
	return mm.Value
}
func (mm *Mesh) CastShadow() bool {
	return mm.Get("castShadow").Bool()
}
func (mm *Mesh) SetCastShadow(v bool) {
	mm.Set("castShadow", v)
}
func (mm *Mesh) Children() js.Value {
	return mm.Get("children")
}
func (mm *Mesh) SetChildren(v js.Value) {
	mm.Set("children", v)
}
func (mm *Mesh) DrawMode() TrianglesDrawModes {
	return TrianglesDrawModes(mm.Get("drawMode"))
}
func (mm *Mesh) SetDrawMode(v TrianglesDrawModes) {
	mm.Set("drawMode", v)
}
func (mm *Mesh) FrustumCulled() bool {
	return mm.Get("frustumCulled").Bool()
}
func (mm *Mesh) SetFrustumCulled(v bool) {
	mm.Set("frustumCulled", v)
}
func (mm *Mesh) Geometry() Geometry {
	return &GeometryImpl{Value: mm.Get("geometry")}
}
func (mm *Mesh) SetGeometry(v Geometry) {
	mm.Set("geometry", v.JSValue())
}
func (mm *Mesh) Id() int {
	return mm.Get("id").Int()
}
func (mm *Mesh) SetId(v int) {
	mm.Set("id", v)
}
func (mm *Mesh) IsMesh() bool {
	return mm.Get("isMesh").Bool()
}
func (mm *Mesh) SetIsMesh(v bool) {
	mm.Set("isMesh", v)
}
func (mm *Mesh) IsObject3D() bool {
	return mm.Get("isObject3D").Bool()
}
func (mm *Mesh) SetIsObject3D(v bool) {
	mm.Set("isObject3D", v)
}
func (mm *Mesh) Layers() *Layers {
	return &Layers{Value: mm.Get("layers")}
}
func (mm *Mesh) SetLayers(v *Layers) {
	mm.Set("layers", v.Value)
}
func (mm *Mesh) Material() Material {
	return &MaterialImpl{Value: mm.Get("material")}
}
func (mm *Mesh) SetMaterial(v Material) {
	mm.Set("material", v.JSValue())
}
func (mm *Mesh) Matrix() *Matrix4 {
	return &Matrix4{Value: mm.Get("matrix")}
}
func (mm *Mesh) SetMatrix(v *Matrix4) {
	mm.Set("matrix", v.Value)
}
func (mm *Mesh) MatrixAutoUpdate() bool {
	return mm.Get("matrixAutoUpdate").Bool()
}
func (mm *Mesh) SetMatrixAutoUpdate(v bool) {
	mm.Set("matrixAutoUpdate", v)
}
func (mm *Mesh) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: mm.Get("matrixWorld")}
}
func (mm *Mesh) SetMatrixWorld(v *Matrix4) {
	mm.Set("matrixWorld", v.Value)
}
func (mm *Mesh) MatrixWorldNeedsUpdate() bool {
	return mm.Get("matrixWorldNeedsUpdate").Bool()
}
func (mm *Mesh) SetMatrixWorldNeedsUpdate(v bool) {
	mm.Set("matrixWorldNeedsUpdate", v)
}
func (mm *Mesh) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: mm.Get("modelViewMatrix")}
}
func (mm *Mesh) SetModelViewMatrix(v *Matrix4) {
	mm.Set("modelViewMatrix", v.Value)
}
func (mm *Mesh) MorphTargetDictionary() js.Value {
	return mm.Get("morphTargetDictionary")
}
func (mm *Mesh) SetMorphTargetDictionary(v js.Value) {
	mm.Set("morphTargetDictionary", v)
}
func (mm *Mesh) MorphTargetInfluences() js.Value {
	return mm.Get("morphTargetInfluences")
}
func (mm *Mesh) SetMorphTargetInfluences(v js.Value) {
	mm.Set("morphTargetInfluences", v)
}
func (mm *Mesh) Name() string {
	return mm.Get("name").String()
}
func (mm *Mesh) SetName(v string) {
	mm.Set("name", v)
}
func (mm *Mesh) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: mm.Get("normalMatrix")}
}
func (mm *Mesh) SetNormalMatrix(v *Matrix3) {
	mm.Set("normalMatrix", v.Value)
}
func (mm *Mesh) OnAfterRender() js.Value {
	return mm.Get("onAfterRender")
}
func (mm *Mesh) SetOnAfterRender(v js.Value) {
	mm.Set("onAfterRender", v)
}
func (mm *Mesh) OnBeforeRender() js.Value {
	return mm.Get("onBeforeRender")
}
func (mm *Mesh) SetOnBeforeRender(v js.Value) {
	mm.Set("onBeforeRender", v)
}
func (mm *Mesh) Parent() *Object3D {
	return &Object3D{Value: mm.Get("parent")}
}
func (mm *Mesh) SetParent(v *Object3D) {
	mm.Set("parent", v.Value)
}
func (mm *Mesh) Position() *Vector3 {
	return &Vector3{Value: mm.Get("position")}
}
func (mm *Mesh) SetPosition(v *Vector3) {
	mm.Set("position", v.Value)
}
func (mm *Mesh) Quaternion() *Quaternion {
	return &Quaternion{Value: mm.Get("quaternion")}
}
func (mm *Mesh) SetQuaternion(v *Quaternion) {
	mm.Set("quaternion", v.Value)
}
func (mm *Mesh) ReceiveShadow() bool {
	return mm.Get("receiveShadow").Bool()
}
func (mm *Mesh) SetReceiveShadow(v bool) {
	mm.Set("receiveShadow", v)
}
func (mm *Mesh) RenderOrder() float64 {
	return mm.Get("renderOrder").Float()
}
func (mm *Mesh) SetRenderOrder(v float64) {
	mm.Set("renderOrder", v)
}
func (mm *Mesh) Rotation() *Euler {
	return &Euler{Value: mm.Get("rotation")}
}
func (mm *Mesh) SetRotation(v *Euler) {
	mm.Set("rotation", v.Value)
}
func (mm *Mesh) Scale() *Vector3 {
	return &Vector3{Value: mm.Get("scale")}
}
func (mm *Mesh) SetScale(v *Vector3) {
	mm.Set("scale", v.Value)
}
func (mm *Mesh) Type() string {
	return mm.Get("type").String()
}
func (mm *Mesh) SetType(v string) {
	mm.Set("type", v)
}
func (mm *Mesh) Up() *Vector3 {
	return &Vector3{Value: mm.Get("up")}
}
func (mm *Mesh) SetUp(v *Vector3) {
	mm.Set("up", v.Value)
}
func (mm *Mesh) UserData() js.Value {
	return mm.Get("userData")
}
func (mm *Mesh) SetUserData(v js.Value) {
	mm.Set("userData", v)
}
func (mm *Mesh) Uuid() string {
	return mm.Get("uuid").String()
}
func (mm *Mesh) SetUuid(v string) {
	mm.Set("uuid", v)
}
func (mm *Mesh) Visible() bool {
	return mm.Get("visible").Bool()
}
func (mm *Mesh) SetVisible(v bool) {
	mm.Set("visible", v)
}
func (mm *Mesh) DefaultMatrixAutoUpdate() bool {
	return mm.Get("DefaultMatrixAutoUpdate").Bool()
}
func (mm *Mesh) SetDefaultMatrixAutoUpdate(v bool) {
	mm.Set("DefaultMatrixAutoUpdate", v)
}
func (mm *Mesh) DefaultUp() *Vector3 {
	return &Vector3{Value: mm.Get("DefaultUp")}
}
func (mm *Mesh) SetDefaultUp(v *Vector3) {
	mm.Set("DefaultUp", v.Value)
}
func (mm *Mesh) Add(object js.Value) *Mesh {
	return &Mesh{Value: mm.Call("add", object)}
}
func (mm *Mesh) AddEventListener(typ string, listener js.Value) {
	mm.Call("addEventListener", typ, listener)
}
func (mm *Mesh) ApplyMatrix(matrix *Matrix4) {
	mm.Call("applyMatrix", matrix)
}
func (mm *Mesh) ApplyQuaternion(quaternion *Quaternion) *Mesh {
	return &Mesh{Value: mm.Call("applyQuaternion", quaternion)}
}
func (mm *Mesh) Clone(recursive bool) *Mesh {
	return &Mesh{Value: mm.Call("clone", recursive)}
}
func (mm *Mesh) Copy(source *Mesh, recursive bool) *Mesh {
	return &Mesh{Value: mm.Call("copy", source, recursive)}
}
func (mm *Mesh) DispatchEvent(event js.Value) {
	mm.Call("dispatchEvent", event)
}
func (mm *Mesh) GetObjectById(id int) *Object3D {
	return &Object3D{Value: mm.Call("getObjectById", id)}
}
func (mm *Mesh) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: mm.Call("getObjectByName", name)}
}
func (mm *Mesh) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: mm.Call("getObjectByProperty", name, value)}
}
func (mm *Mesh) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("getWorldDirection", target)}
}
func (mm *Mesh) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("getWorldPosition", target)}
}
func (mm *Mesh) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: mm.Call("getWorldQuaternion", target)}
}
func (mm *Mesh) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("getWorldScale", target)}
}
func (mm *Mesh) HasEventListener(typ string, listener js.Value) bool {
	return mm.Call("hasEventListener", typ, listener).Bool()
}
func (mm *Mesh) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("localToWorld", vector)}
}
func (mm *Mesh) LookAt(vector *Vector3, y float64, z float64) {
	mm.Call("lookAt", vector, y, z)
}
func (mm *Mesh) Raycast(raycaster *Raycaster, intersects js.Value) {
	mm.Call("raycast", raycaster, intersects)
}
func (mm *Mesh) Remove(object js.Value) *Mesh {
	return &Mesh{Value: mm.Call("remove", object)}
}
func (mm *Mesh) RemoveEventListener(typ string, listener js.Value) {
	mm.Call("removeEventListener", typ, listener)
}
func (mm *Mesh) RotateOnAxis(axis *Vector3, angle float64) *Mesh {
	return &Mesh{Value: mm.Call("rotateOnAxis", axis, angle)}
}
func (mm *Mesh) RotateOnWorldAxis(axis *Vector3, angle float64) *Mesh {
	return &Mesh{Value: mm.Call("rotateOnWorldAxis", axis, angle)}
}
func (mm *Mesh) RotateX(angle float64) *Mesh {
	return &Mesh{Value: mm.Call("rotateX", angle)}
}
func (mm *Mesh) RotateY(angle float64) *Mesh {
	return &Mesh{Value: mm.Call("rotateY", angle)}
}
func (mm *Mesh) RotateZ(angle float64) *Mesh {
	return &Mesh{Value: mm.Call("rotateZ", angle)}
}
func (mm *Mesh) SetDrawMode2(drawMode TrianglesDrawModes) {
	mm.Call("setDrawMode", drawMode)
}
func (mm *Mesh) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	mm.Call("setRotationFromAxisAngle", axis, angle)
}
func (mm *Mesh) SetRotationFromEuler(euler *Euler) {
	mm.Call("setRotationFromEuler", euler)
}
func (mm *Mesh) SetRotationFromMatrix(m *Matrix4) {
	mm.Call("setRotationFromMatrix", m)
}
func (mm *Mesh) SetRotationFromQuaternion(q *Quaternion) {
	mm.Call("setRotationFromQuaternion", q)
}
func (mm *Mesh) ToJSON(meta js.Value) js.Value {
	return mm.Call("toJSON", meta)
}
func (mm *Mesh) TranslateOnAxis(axis *Vector3, distance float64) *Mesh {
	return &Mesh{Value: mm.Call("translateOnAxis", axis, distance)}
}
func (mm *Mesh) TranslateX(distance float64) *Mesh {
	return &Mesh{Value: mm.Call("translateX", distance)}
}
func (mm *Mesh) TranslateY(distance float64) *Mesh {
	return &Mesh{Value: mm.Call("translateY", distance)}
}
func (mm *Mesh) TranslateZ(distance float64) *Mesh {
	return &Mesh{Value: mm.Call("translateZ", distance)}
}
func (mm *Mesh) Traverse(callback js.Value) {
	mm.Call("traverse", callback)
}
func (mm *Mesh) TraverseAncestors(callback js.Value) {
	mm.Call("traverseAncestors", callback)
}
func (mm *Mesh) TraverseVisible(callback js.Value) {
	mm.Call("traverseVisible", callback)
}
func (mm *Mesh) UpdateMatrix() {
	mm.Call("updateMatrix")
}
func (mm *Mesh) UpdateMatrixWorld(force bool) {
	mm.Call("updateMatrixWorld", force)
}
func (mm *Mesh) UpdateMorphTargets() {
	mm.Call("updateMorphTargets")
}
func (mm *Mesh) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	mm.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (mm *Mesh) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("worldToLocal", vector)}
}

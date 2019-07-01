// Code generated from "objects/Mesh.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

type Mesh interface {
	JSValue() js.Value
	CastShadow() bool
	SetCastShadow(v bool)
	Children() js.Value
	SetChildren(v js.Value)
	DrawMode() TrianglesDrawModes
	SetDrawMode(v TrianglesDrawModes)
	FrustumCulled() bool
	SetFrustumCulled(v bool)
	Geometry() Geometry
	SetGeometry(v Geometry)
	Id() int
	SetId(v int)
	IsMesh() bool
	SetIsMesh(v bool)
	IsObject3D() bool
	SetIsObject3D(v bool)
	Layers() *Layers
	SetLayers(v *Layers)
	Material() Material
	SetMaterial(v Material)
	Matrix() *Matrix4
	SetMatrix(v *Matrix4)
	MatrixAutoUpdate() bool
	SetMatrixAutoUpdate(v bool)
	MatrixWorld() *Matrix4
	SetMatrixWorld(v *Matrix4)
	MatrixWorldNeedsUpdate() bool
	SetMatrixWorldNeedsUpdate(v bool)
	ModelViewMatrix() *Matrix4
	SetModelViewMatrix(v *Matrix4)
	MorphTargetDictionary() js.Value
	SetMorphTargetDictionary(v js.Value)
	MorphTargetInfluences() js.Value
	SetMorphTargetInfluences(v js.Value)
	Name() string
	SetName(v string)
	NormalMatrix() *Matrix3
	SetNormalMatrix(v *Matrix3)
	OnAfterRender() js.Value
	SetOnAfterRender(v js.Value)
	OnBeforeRender() js.Value
	SetOnBeforeRender(v js.Value)
	Parent() *Object3D
	SetParent(v *Object3D)
	Position() *Vector3
	SetPosition(v *Vector3)
	Quaternion() *Quaternion
	SetQuaternion(v *Quaternion)
	ReceiveShadow() bool
	SetReceiveShadow(v bool)
	RenderOrder() float64
	SetRenderOrder(v float64)
	Rotation() *Euler
	SetRotation(v *Euler)
	Scale() *Vector3
	SetScale(v *Vector3)
	Type() string
	SetType(v string)
	Up() *Vector3
	SetUp(v *Vector3)
	UserData() js.Value
	SetUserData(v js.Value)
	Uuid() string
	SetUuid(v string)
	Visible() bool
	SetVisible(v bool)
	DefaultMatrixAutoUpdate() bool
	SetDefaultMatrixAutoUpdate(v bool)
	DefaultUp() *Vector3
	SetDefaultUp(v *Vector3)
	Add(object js.Value) Mesh
	AddEventListener(typ string, listener js.Value)
	ApplyMatrix(matrix *Matrix4)
	ApplyQuaternion(quaternion *Quaternion) Mesh
	Clone(recursive bool) Mesh
	Copy(source Mesh, recursive bool) Mesh
	DispatchEvent(event js.Value)
	GetObjectById(id int) *Object3D
	GetObjectByName(name string) *Object3D
	GetObjectByProperty(name string, value string) *Object3D
	GetWorldDirection(target *Vector3) *Vector3
	GetWorldPosition(target *Vector3) *Vector3
	GetWorldQuaternion(target *Quaternion) *Quaternion
	GetWorldScale(target *Vector3) *Vector3
	HasEventListener(typ string, listener js.Value) bool
	LocalToWorld(vector *Vector3) *Vector3
	LookAt(vector *Vector3, y float64, z float64)
	Raycast(raycaster *Raycaster, intersects js.Value)
	Remove(object js.Value) Mesh
	RemoveEventListener(typ string, listener js.Value)
	RotateOnAxis(axis *Vector3, angle float64) Mesh
	RotateOnWorldAxis(axis *Vector3, angle float64) Mesh
	RotateX(angle float64) Mesh
	RotateY(angle float64) Mesh
	RotateZ(angle float64) Mesh
	SetDrawMode2(drawMode TrianglesDrawModes)
	SetRotationFromAxisAngle(axis *Vector3, angle float64)
	SetRotationFromEuler(euler *Euler)
	SetRotationFromMatrix(m *Matrix4)
	SetRotationFromQuaternion(q *Quaternion)
	ToJSON(meta js.Value) js.Value
	TranslateOnAxis(axis *Vector3, distance float64) Mesh
	TranslateX(distance float64) Mesh
	TranslateY(distance float64) Mesh
	TranslateZ(distance float64) Mesh
	Traverse(callback js.Value)
	TraverseAncestors(callback js.Value)
	TraverseVisible(callback js.Value)
	UpdateMatrix()
	UpdateMatrixWorld(force bool)
	UpdateMorphTargets()
	UpdateWorldMatrix(updateParents bool, updateChildren bool)
	WorldToLocal(vector *Vector3) *Vector3
}

// MeshImpl extend: [Object3D]
type MeshImpl struct {
	js.Value
}

func NewMesh(geometry js.Wrapper, material Material) *MeshImpl {
	return &MeshImpl{Value: get("Mesh").New(geometry.JSValue(), material.JSValue())}
}
func (mm *MeshImpl) JSValue() js.Value {
	return mm.Value
}
func (mm *MeshImpl) CastShadow() bool {
	return mm.Get("castShadow").Bool()
}
func (mm *MeshImpl) SetCastShadow(v bool) {
	mm.Set("castShadow", v)
}
func (mm *MeshImpl) Children() js.Value {
	return mm.Get("children")
}
func (mm *MeshImpl) SetChildren(v js.Value) {
	mm.Set("children", v)
}
func (mm *MeshImpl) DrawMode() TrianglesDrawModes {
	return TrianglesDrawModes(mm.Get("drawMode"))
}
func (mm *MeshImpl) SetDrawMode(v TrianglesDrawModes) {
	mm.Set("drawMode", v)
}
func (mm *MeshImpl) FrustumCulled() bool {
	return mm.Get("frustumCulled").Bool()
}
func (mm *MeshImpl) SetFrustumCulled(v bool) {
	mm.Set("frustumCulled", v)
}
func (mm *MeshImpl) Geometry() Geometry {
	return &GeometryImpl{Value: mm.Get("geometry")}
}
func (mm *MeshImpl) SetGeometry(v Geometry) {
	mm.Set("geometry", v.JSValue())
}
func (mm *MeshImpl) Id() int {
	return mm.Get("id").Int()
}
func (mm *MeshImpl) SetId(v int) {
	mm.Set("id", v)
}
func (mm *MeshImpl) IsMesh() bool {
	return mm.Get("isMesh").Bool()
}
func (mm *MeshImpl) SetIsMesh(v bool) {
	mm.Set("isMesh", v)
}
func (mm *MeshImpl) IsObject3D() bool {
	return mm.Get("isObject3D").Bool()
}
func (mm *MeshImpl) SetIsObject3D(v bool) {
	mm.Set("isObject3D", v)
}
func (mm *MeshImpl) Layers() *Layers {
	return &Layers{Value: mm.Get("layers")}
}
func (mm *MeshImpl) SetLayers(v *Layers) {
	mm.Set("layers", v.JSValue())
}
func (mm *MeshImpl) Material() Material {
	return &MaterialImpl{Value: mm.Get("material")}
}
func (mm *MeshImpl) SetMaterial(v Material) {
	mm.Set("material", v.JSValue())
}
func (mm *MeshImpl) Matrix() *Matrix4 {
	return &Matrix4{Value: mm.Get("matrix")}
}
func (mm *MeshImpl) SetMatrix(v *Matrix4) {
	mm.Set("matrix", v.JSValue())
}
func (mm *MeshImpl) MatrixAutoUpdate() bool {
	return mm.Get("matrixAutoUpdate").Bool()
}
func (mm *MeshImpl) SetMatrixAutoUpdate(v bool) {
	mm.Set("matrixAutoUpdate", v)
}
func (mm *MeshImpl) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: mm.Get("matrixWorld")}
}
func (mm *MeshImpl) SetMatrixWorld(v *Matrix4) {
	mm.Set("matrixWorld", v.JSValue())
}
func (mm *MeshImpl) MatrixWorldNeedsUpdate() bool {
	return mm.Get("matrixWorldNeedsUpdate").Bool()
}
func (mm *MeshImpl) SetMatrixWorldNeedsUpdate(v bool) {
	mm.Set("matrixWorldNeedsUpdate", v)
}
func (mm *MeshImpl) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: mm.Get("modelViewMatrix")}
}
func (mm *MeshImpl) SetModelViewMatrix(v *Matrix4) {
	mm.Set("modelViewMatrix", v.JSValue())
}
func (mm *MeshImpl) MorphTargetDictionary() js.Value {
	return mm.Get("morphTargetDictionary")
}
func (mm *MeshImpl) SetMorphTargetDictionary(v js.Value) {
	mm.Set("morphTargetDictionary", v)
}
func (mm *MeshImpl) MorphTargetInfluences() js.Value {
	return mm.Get("morphTargetInfluences")
}
func (mm *MeshImpl) SetMorphTargetInfluences(v js.Value) {
	mm.Set("morphTargetInfluences", v)
}
func (mm *MeshImpl) Name() string {
	return mm.Get("name").String()
}
func (mm *MeshImpl) SetName(v string) {
	mm.Set("name", v)
}
func (mm *MeshImpl) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: mm.Get("normalMatrix")}
}
func (mm *MeshImpl) SetNormalMatrix(v *Matrix3) {
	mm.Set("normalMatrix", v.JSValue())
}
func (mm *MeshImpl) OnAfterRender() js.Value {
	return mm.Get("onAfterRender")
}
func (mm *MeshImpl) SetOnAfterRender(v js.Value) {
	mm.Set("onAfterRender", v)
}
func (mm *MeshImpl) OnBeforeRender() js.Value {
	return mm.Get("onBeforeRender")
}
func (mm *MeshImpl) SetOnBeforeRender(v js.Value) {
	mm.Set("onBeforeRender", v)
}
func (mm *MeshImpl) Parent() *Object3D {
	return &Object3D{Value: mm.Get("parent")}
}
func (mm *MeshImpl) SetParent(v *Object3D) {
	mm.Set("parent", v.JSValue())
}
func (mm *MeshImpl) Position() *Vector3 {
	return &Vector3{Value: mm.Get("position")}
}
func (mm *MeshImpl) SetPosition(v *Vector3) {
	mm.Set("position", v.JSValue())
}
func (mm *MeshImpl) Quaternion() *Quaternion {
	return &Quaternion{Value: mm.Get("quaternion")}
}
func (mm *MeshImpl) SetQuaternion(v *Quaternion) {
	mm.Set("quaternion", v.JSValue())
}
func (mm *MeshImpl) ReceiveShadow() bool {
	return mm.Get("receiveShadow").Bool()
}
func (mm *MeshImpl) SetReceiveShadow(v bool) {
	mm.Set("receiveShadow", v)
}
func (mm *MeshImpl) RenderOrder() float64 {
	return mm.Get("renderOrder").Float()
}
func (mm *MeshImpl) SetRenderOrder(v float64) {
	mm.Set("renderOrder", v)
}
func (mm *MeshImpl) Rotation() *Euler {
	return &Euler{Value: mm.Get("rotation")}
}
func (mm *MeshImpl) SetRotation(v *Euler) {
	mm.Set("rotation", v.JSValue())
}
func (mm *MeshImpl) Scale() *Vector3 {
	return &Vector3{Value: mm.Get("scale")}
}
func (mm *MeshImpl) SetScale(v *Vector3) {
	mm.Set("scale", v.JSValue())
}
func (mm *MeshImpl) Type() string {
	return mm.Get("type").String()
}
func (mm *MeshImpl) SetType(v string) {
	mm.Set("type", v)
}
func (mm *MeshImpl) Up() *Vector3 {
	return &Vector3{Value: mm.Get("up")}
}
func (mm *MeshImpl) SetUp(v *Vector3) {
	mm.Set("up", v.JSValue())
}
func (mm *MeshImpl) UserData() js.Value {
	return mm.Get("userData")
}
func (mm *MeshImpl) SetUserData(v js.Value) {
	mm.Set("userData", v)
}
func (mm *MeshImpl) Uuid() string {
	return mm.Get("uuid").String()
}
func (mm *MeshImpl) SetUuid(v string) {
	mm.Set("uuid", v)
}
func (mm *MeshImpl) Visible() bool {
	return mm.Get("visible").Bool()
}
func (mm *MeshImpl) SetVisible(v bool) {
	mm.Set("visible", v)
}
func (mm *MeshImpl) DefaultMatrixAutoUpdate() bool {
	return mm.Get("DefaultMatrixAutoUpdate").Bool()
}
func (mm *MeshImpl) SetDefaultMatrixAutoUpdate(v bool) {
	mm.Set("DefaultMatrixAutoUpdate", v)
}
func (mm *MeshImpl) DefaultUp() *Vector3 {
	return &Vector3{Value: mm.Get("DefaultUp")}
}
func (mm *MeshImpl) SetDefaultUp(v *Vector3) {
	mm.Set("DefaultUp", v.JSValue())
}
func (mm *MeshImpl) Add(object js.Value) Mesh {
	return &MeshImpl{Value: mm.Call("add", object)}
}
func (mm *MeshImpl) AddEventListener(typ string, listener js.Value) {
	mm.Call("addEventListener", typ, listener)
}
func (mm *MeshImpl) ApplyMatrix(matrix *Matrix4) {
	mm.Call("applyMatrix", matrix.JSValue())
}
func (mm *MeshImpl) ApplyQuaternion(quaternion *Quaternion) Mesh {
	return &MeshImpl{Value: mm.Call("applyQuaternion", quaternion)}
}
func (mm *MeshImpl) Clone(recursive bool) Mesh {
	return &MeshImpl{Value: mm.Call("clone", recursive)}
}
func (mm *MeshImpl) Copy(source Mesh, recursive bool) Mesh {
	return &MeshImpl{Value: mm.Call("copy", source, recursive)}
}
func (mm *MeshImpl) DispatchEvent(event js.Value) {
	mm.Call("dispatchEvent", event)
}
func (mm *MeshImpl) GetObjectById(id int) *Object3D {
	return &Object3D{Value: mm.Call("getObjectById", id)}
}
func (mm *MeshImpl) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: mm.Call("getObjectByName", name)}
}
func (mm *MeshImpl) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: mm.Call("getObjectByProperty", name, value)}
}
func (mm *MeshImpl) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("getWorldDirection", target)}
}
func (mm *MeshImpl) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("getWorldPosition", target)}
}
func (mm *MeshImpl) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: mm.Call("getWorldQuaternion", target)}
}
func (mm *MeshImpl) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("getWorldScale", target)}
}
func (mm *MeshImpl) HasEventListener(typ string, listener js.Value) bool {
	return mm.Call("hasEventListener", typ, listener).Bool()
}
func (mm *MeshImpl) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("localToWorld", vector)}
}
func (mm *MeshImpl) LookAt(vector *Vector3, y float64, z float64) {
	mm.Call("lookAt", vector, y, z)
}
func (mm *MeshImpl) Raycast(raycaster *Raycaster, intersects js.Value) {
	mm.Call("raycast", raycaster.JSValue(), intersects)
}
func (mm *MeshImpl) Remove(object js.Value) Mesh {
	return &MeshImpl{Value: mm.Call("remove", object)}
}
func (mm *MeshImpl) RemoveEventListener(typ string, listener js.Value) {
	mm.Call("removeEventListener", typ, listener)
}
func (mm *MeshImpl) RotateOnAxis(axis *Vector3, angle float64) Mesh {
	return &MeshImpl{Value: mm.Call("rotateOnAxis", axis, angle)}
}
func (mm *MeshImpl) RotateOnWorldAxis(axis *Vector3, angle float64) Mesh {
	return &MeshImpl{Value: mm.Call("rotateOnWorldAxis", axis, angle)}
}
func (mm *MeshImpl) RotateX(angle float64) Mesh {
	return &MeshImpl{Value: mm.Call("rotateX", angle)}
}
func (mm *MeshImpl) RotateY(angle float64) Mesh {
	return &MeshImpl{Value: mm.Call("rotateY", angle)}
}
func (mm *MeshImpl) RotateZ(angle float64) Mesh {
	return &MeshImpl{Value: mm.Call("rotateZ", angle)}
}
func (mm *MeshImpl) SetDrawMode2(drawMode TrianglesDrawModes) {
	mm.Call("setDrawMode", drawMode)
}
func (mm *MeshImpl) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	mm.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (mm *MeshImpl) SetRotationFromEuler(euler *Euler) {
	mm.Call("setRotationFromEuler", euler.JSValue())
}
func (mm *MeshImpl) SetRotationFromMatrix(m *Matrix4) {
	mm.Call("setRotationFromMatrix", m.JSValue())
}
func (mm *MeshImpl) SetRotationFromQuaternion(q *Quaternion) {
	mm.Call("setRotationFromQuaternion", q.JSValue())
}
func (mm *MeshImpl) ToJSON(meta js.Value) js.Value {
	return mm.Call("toJSON", meta)
}
func (mm *MeshImpl) TranslateOnAxis(axis *Vector3, distance float64) Mesh {
	return &MeshImpl{Value: mm.Call("translateOnAxis", axis, distance)}
}
func (mm *MeshImpl) TranslateX(distance float64) Mesh {
	return &MeshImpl{Value: mm.Call("translateX", distance)}
}
func (mm *MeshImpl) TranslateY(distance float64) Mesh {
	return &MeshImpl{Value: mm.Call("translateY", distance)}
}
func (mm *MeshImpl) TranslateZ(distance float64) Mesh {
	return &MeshImpl{Value: mm.Call("translateZ", distance)}
}
func (mm *MeshImpl) Traverse(callback js.Value) {
	mm.Call("traverse", callback)
}
func (mm *MeshImpl) TraverseAncestors(callback js.Value) {
	mm.Call("traverseAncestors", callback)
}
func (mm *MeshImpl) TraverseVisible(callback js.Value) {
	mm.Call("traverseVisible", callback)
}
func (mm *MeshImpl) UpdateMatrix() {
	mm.Call("updateMatrix")
}
func (mm *MeshImpl) UpdateMatrixWorld(force bool) {
	mm.Call("updateMatrixWorld", force)
}
func (mm *MeshImpl) UpdateMorphTargets() {
	mm.Call("updateMorphTargets")
}
func (mm *MeshImpl) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	mm.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (mm *MeshImpl) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: mm.Call("worldToLocal", vector)}
}

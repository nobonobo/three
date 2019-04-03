// Code generated from "objects/SkinnedMesh.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type SkinnedMesh struct {
	js.Value
}

func NewSkinnedMesh(geometry *Geometry, material *Material, useVertexTexture bool) *SkinnedMesh {
	return &SkinnedMesh{Value: get("SkinnedMesh").New(geometry, material, useVertexTexture)}
}
func (sm *SkinnedMesh) BindMatrix() *Matrix4 {
	return &Matrix4{Value: sm.Get("bindMatrix")}
}
func (sm *SkinnedMesh) SetBindMatrix(v *Matrix4) {
	sm.Set("bindMatrix", v)
}
func (sm *SkinnedMesh) BindMatrixInverse() *Matrix4 {
	return &Matrix4{Value: sm.Get("bindMatrixInverse")}
}
func (sm *SkinnedMesh) SetBindMatrixInverse(v *Matrix4) {
	sm.Set("bindMatrixInverse", v)
}
func (sm *SkinnedMesh) BindMode() string {
	return sm.Get("bindMode").String()
}
func (sm *SkinnedMesh) SetBindMode(v string) {
	sm.Set("bindMode", v)
}
func (sm *SkinnedMesh) CastShadow() bool {
	return sm.Get("castShadow").Bool()
}
func (sm *SkinnedMesh) SetCastShadow(v bool) {
	sm.Set("castShadow", v)
}
func (sm *SkinnedMesh) Children() js.Value {
	return sm.Get("children")
}
func (sm *SkinnedMesh) SetChildren(v js.Value) {
	sm.Set("children", v)
}
func (sm *SkinnedMesh) DrawMode() TrianglesDrawModes {
	return TrianglesDrawModes(sm.Get("drawMode"))
}
func (sm *SkinnedMesh) SetDrawMode(v TrianglesDrawModes) {
	sm.Set("drawMode", v)
}
func (sm *SkinnedMesh) FrustumCulled() bool {
	return sm.Get("frustumCulled").Bool()
}
func (sm *SkinnedMesh) SetFrustumCulled(v bool) {
	sm.Set("frustumCulled", v)
}
func (sm *SkinnedMesh) Geometry() *Geometry {
	return &Geometry{Value: sm.Get("geometry")}
}
func (sm *SkinnedMesh) SetGeometry(v *Geometry) {
	sm.Set("geometry", v)
}
func (sm *SkinnedMesh) Id() int {
	return sm.Get("id").Int()
}
func (sm *SkinnedMesh) SetId(v int) {
	sm.Set("id", v)
}
func (sm *SkinnedMesh) IsMesh() bool {
	return sm.Get("isMesh").Bool()
}
func (sm *SkinnedMesh) SetIsMesh(v bool) {
	sm.Set("isMesh", v)
}
func (sm *SkinnedMesh) IsObject3D() bool {
	return sm.Get("isObject3D").Bool()
}
func (sm *SkinnedMesh) SetIsObject3D(v bool) {
	sm.Set("isObject3D", v)
}
func (sm *SkinnedMesh) Layers() *Layers {
	return &Layers{Value: sm.Get("layers")}
}
func (sm *SkinnedMesh) SetLayers(v *Layers) {
	sm.Set("layers", v)
}
func (sm *SkinnedMesh) Material() *Material {
	return &Material{Value: sm.Get("material")}
}
func (sm *SkinnedMesh) SetMaterial(v *Material) {
	sm.Set("material", v)
}
func (sm *SkinnedMesh) Matrix() *Matrix4 {
	return &Matrix4{Value: sm.Get("matrix")}
}
func (sm *SkinnedMesh) SetMatrix(v *Matrix4) {
	sm.Set("matrix", v)
}
func (sm *SkinnedMesh) MatrixAutoUpdate() bool {
	return sm.Get("matrixAutoUpdate").Bool()
}
func (sm *SkinnedMesh) SetMatrixAutoUpdate(v bool) {
	sm.Set("matrixAutoUpdate", v)
}
func (sm *SkinnedMesh) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: sm.Get("matrixWorld")}
}
func (sm *SkinnedMesh) SetMatrixWorld(v *Matrix4) {
	sm.Set("matrixWorld", v)
}
func (sm *SkinnedMesh) MatrixWorldNeedsUpdate() bool {
	return sm.Get("matrixWorldNeedsUpdate").Bool()
}
func (sm *SkinnedMesh) SetMatrixWorldNeedsUpdate(v bool) {
	sm.Set("matrixWorldNeedsUpdate", v)
}
func (sm *SkinnedMesh) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: sm.Get("modelViewMatrix")}
}
func (sm *SkinnedMesh) SetModelViewMatrix(v *Matrix4) {
	sm.Set("modelViewMatrix", v)
}
func (sm *SkinnedMesh) MorphTargetDictionary() js.Value {
	return sm.Get("morphTargetDictionary")
}
func (sm *SkinnedMesh) SetMorphTargetDictionary(v js.Value) {
	sm.Set("morphTargetDictionary", v)
}
func (sm *SkinnedMesh) MorphTargetInfluences() js.Value {
	return sm.Get("morphTargetInfluences")
}
func (sm *SkinnedMesh) SetMorphTargetInfluences(v js.Value) {
	sm.Set("morphTargetInfluences", v)
}
func (sm *SkinnedMesh) Name() string {
	return sm.Get("name").String()
}
func (sm *SkinnedMesh) SetName(v string) {
	sm.Set("name", v)
}
func (sm *SkinnedMesh) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: sm.Get("normalMatrix")}
}
func (sm *SkinnedMesh) SetNormalMatrix(v *Matrix3) {
	sm.Set("normalMatrix", v)
}
func (sm *SkinnedMesh) OnAfterRender() js.Value {
	return sm.Get("onAfterRender")
}
func (sm *SkinnedMesh) SetOnAfterRender(v js.Value) {
	sm.Set("onAfterRender", v)
}
func (sm *SkinnedMesh) OnBeforeRender() js.Value {
	return sm.Get("onBeforeRender")
}
func (sm *SkinnedMesh) SetOnBeforeRender(v js.Value) {
	sm.Set("onBeforeRender", v)
}
func (sm *SkinnedMesh) Parent() *Object3D {
	return &Object3D{Value: sm.Get("parent")}
}
func (sm *SkinnedMesh) SetParent(v *Object3D) {
	sm.Set("parent", v)
}
func (sm *SkinnedMesh) Position() *Vector3 {
	return &Vector3{Value: sm.Get("position")}
}
func (sm *SkinnedMesh) SetPosition(v *Vector3) {
	sm.Set("position", v)
}
func (sm *SkinnedMesh) Quaternion() *Quaternion {
	return &Quaternion{Value: sm.Get("quaternion")}
}
func (sm *SkinnedMesh) SetQuaternion(v *Quaternion) {
	sm.Set("quaternion", v)
}
func (sm *SkinnedMesh) ReceiveShadow() bool {
	return sm.Get("receiveShadow").Bool()
}
func (sm *SkinnedMesh) SetReceiveShadow(v bool) {
	sm.Set("receiveShadow", v)
}
func (sm *SkinnedMesh) RenderOrder() float64 {
	return sm.Get("renderOrder").Float()
}
func (sm *SkinnedMesh) SetRenderOrder(v float64) {
	sm.Set("renderOrder", v)
}
func (sm *SkinnedMesh) Rotation() *Euler {
	return &Euler{Value: sm.Get("rotation")}
}
func (sm *SkinnedMesh) SetRotation(v *Euler) {
	sm.Set("rotation", v)
}
func (sm *SkinnedMesh) Scale() *Vector3 {
	return &Vector3{Value: sm.Get("scale")}
}
func (sm *SkinnedMesh) SetScale(v *Vector3) {
	sm.Set("scale", v)
}
func (sm *SkinnedMesh) Skeleton() *Skeleton {
	return &Skeleton{Value: sm.Get("skeleton")}
}
func (sm *SkinnedMesh) SetSkeleton(v *Skeleton) {
	sm.Set("skeleton", v)
}
func (sm *SkinnedMesh) Type() string {
	return sm.Get("type").String()
}
func (sm *SkinnedMesh) SetType(v string) {
	sm.Set("type", v)
}
func (sm *SkinnedMesh) Up() *Vector3 {
	return &Vector3{Value: sm.Get("up")}
}
func (sm *SkinnedMesh) SetUp(v *Vector3) {
	sm.Set("up", v)
}
func (sm *SkinnedMesh) UserData() js.Value {
	return sm.Get("userData")
}
func (sm *SkinnedMesh) SetUserData(v js.Value) {
	sm.Set("userData", v)
}
func (sm *SkinnedMesh) Uuid() string {
	return sm.Get("uuid").String()
}
func (sm *SkinnedMesh) SetUuid(v string) {
	sm.Set("uuid", v)
}
func (sm *SkinnedMesh) Visible() bool {
	return sm.Get("visible").Bool()
}
func (sm *SkinnedMesh) SetVisible(v bool) {
	sm.Set("visible", v)
}
func (sm *SkinnedMesh) DefaultMatrixAutoUpdate() bool {
	return sm.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sm *SkinnedMesh) SetDefaultMatrixAutoUpdate(v bool) {
	sm.Set("DefaultMatrixAutoUpdate", v)
}
func (sm *SkinnedMesh) DefaultUp() *Vector3 {
	return &Vector3{Value: sm.Get("DefaultUp")}
}
func (sm *SkinnedMesh) SetDefaultUp(v *Vector3) {
	sm.Set("DefaultUp", v)
}
func (sm *SkinnedMesh) Add(object js.Value) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("add", object)}
}
func (sm *SkinnedMesh) AddEventListener(typ string, listener js.Value) {
	sm.Call("addEventListener", typ, listener)
}
func (sm *SkinnedMesh) ApplyMatrix(matrix *Matrix4) {
	sm.Call("applyMatrix", matrix)
}
func (sm *SkinnedMesh) ApplyQuaternion(quaternion *Quaternion) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("applyQuaternion", quaternion)}
}
func (sm *SkinnedMesh) Bind(skeleton *Skeleton, bindMatrix *Matrix4) {
	sm.Call("bind", skeleton, bindMatrix)
}
func (sm *SkinnedMesh) Clone(recursive bool) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("clone", recursive)}
}
func (sm *SkinnedMesh) Copy(source *SkinnedMesh, recursive bool) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("copy", source, recursive)}
}
func (sm *SkinnedMesh) DispatchEvent(event js.Value) {
	sm.Call("dispatchEvent", event)
}
func (sm *SkinnedMesh) GetObjectById(id int) *Object3D {
	return &Object3D{Value: sm.Call("getObjectById", id)}
}
func (sm *SkinnedMesh) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: sm.Call("getObjectByName", name)}
}
func (sm *SkinnedMesh) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: sm.Call("getObjectByProperty", name, value)}
}
func (sm *SkinnedMesh) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: sm.Call("getWorldDirection", target)}
}
func (sm *SkinnedMesh) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: sm.Call("getWorldPosition", target)}
}
func (sm *SkinnedMesh) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: sm.Call("getWorldQuaternion", target)}
}
func (sm *SkinnedMesh) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: sm.Call("getWorldScale", target)}
}
func (sm *SkinnedMesh) HasEventListener(typ string, listener js.Value) bool {
	return sm.Call("hasEventListener", typ, listener).Bool()
}
func (sm *SkinnedMesh) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: sm.Call("localToWorld", vector)}
}
func (sm *SkinnedMesh) LookAt(vector *Vector3, y float64, z float64) {
	sm.Call("lookAt", vector, y, z)
}
func (sm *SkinnedMesh) NormalizeSkinWeights() {
	sm.Call("normalizeSkinWeights")
}
func (sm *SkinnedMesh) Pose() {
	sm.Call("pose")
}
func (sm *SkinnedMesh) Raycast(raycaster *Raycaster, intersects js.Value) {
	sm.Call("raycast", raycaster, intersects)
}
func (sm *SkinnedMesh) Remove(object js.Value) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("remove", object)}
}
func (sm *SkinnedMesh) RemoveEventListener(typ string, listener js.Value) {
	sm.Call("removeEventListener", typ, listener)
}
func (sm *SkinnedMesh) RotateOnAxis(axis *Vector3, angle float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("rotateOnAxis", axis, angle)}
}
func (sm *SkinnedMesh) RotateOnWorldAxis(axis *Vector3, angle float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("rotateOnWorldAxis", axis, angle)}
}
func (sm *SkinnedMesh) RotateX(angle float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("rotateX", angle)}
}
func (sm *SkinnedMesh) RotateY(angle float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("rotateY", angle)}
}
func (sm *SkinnedMesh) RotateZ(angle float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("rotateZ", angle)}
}
func (sm *SkinnedMesh) SetDrawMode2(drawMode TrianglesDrawModes) {
	sm.Call("setDrawMode", drawMode)
}
func (sm *SkinnedMesh) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	sm.Call("setRotationFromAxisAngle", axis, angle)
}
func (sm *SkinnedMesh) SetRotationFromEuler(euler *Euler) {
	sm.Call("setRotationFromEuler", euler)
}
func (sm *SkinnedMesh) SetRotationFromMatrix(m *Matrix4) {
	sm.Call("setRotationFromMatrix", m)
}
func (sm *SkinnedMesh) SetRotationFromQuaternion(q *Quaternion) {
	sm.Call("setRotationFromQuaternion", q)
}
func (sm *SkinnedMesh) ToJSON(meta js.Value) js.Value {
	return sm.Call("toJSON", meta)
}
func (sm *SkinnedMesh) TranslateOnAxis(axis *Vector3, distance float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("translateOnAxis", axis, distance)}
}
func (sm *SkinnedMesh) TranslateX(distance float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("translateX", distance)}
}
func (sm *SkinnedMesh) TranslateY(distance float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("translateY", distance)}
}
func (sm *SkinnedMesh) TranslateZ(distance float64) *SkinnedMesh {
	return &SkinnedMesh{Value: sm.Call("translateZ", distance)}
}
func (sm *SkinnedMesh) Traverse(callback js.Value) {
	sm.Call("traverse", callback)
}
func (sm *SkinnedMesh) TraverseAncestors(callback js.Value) {
	sm.Call("traverseAncestors", callback)
}
func (sm *SkinnedMesh) TraverseVisible(callback js.Value) {
	sm.Call("traverseVisible", callback)
}
func (sm *SkinnedMesh) UpdateMatrix() {
	sm.Call("updateMatrix")
}
func (sm *SkinnedMesh) UpdateMatrixWorld(force bool) {
	sm.Call("updateMatrixWorld", force)
}
func (sm *SkinnedMesh) UpdateMorphTargets() {
	sm.Call("updateMorphTargets")
}
func (sm *SkinnedMesh) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sm.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sm *SkinnedMesh) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: sm.Call("worldToLocal", vector)}
}

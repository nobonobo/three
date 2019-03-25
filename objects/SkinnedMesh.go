package objects

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type SkinnedMesh struct {
	js.Value
}

func NewSkinnedMesh(geometry core.Geometry, material materials.Material, useVertexTexture bool) *SkinnedMesh {
	return &SkinnedMesh{Value: get("SkinnedMesh").New(geometry, material, useVertexTexture)}
}
func (sm *SkinnedMesh) BindMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: sm.Get("bindMatrix")}
}
func (sm *SkinnedMesh) SetBindMatrix(v *math.Matrix4) {
	sm.Set("bindMatrix", v)
}
func (sm *SkinnedMesh) BindMatrixInverse() *math.Matrix4 {
	return &math.Matrix4{Value: sm.Get("bindMatrixInverse")}
}
func (sm *SkinnedMesh) SetBindMatrixInverse(v *math.Matrix4) {
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
func (sm *SkinnedMesh) Children() []core.Object3D {
	return []core.Object3D(sm.Get("children"))
}
func (sm *SkinnedMesh) SetChildren(v []core.Object3D) {
	sm.Set("children", v)
}
func (sm *SkinnedMesh) DrawMode() *TrianglesDrawModes {
	return &TrianglesDrawModes{Value: sm.Get("drawMode")}
}
func (sm *SkinnedMesh) SetDrawMode(v *TrianglesDrawModes) {
	sm.Set("drawMode", v)
}
func (sm *SkinnedMesh) FrustumCulled() bool {
	return sm.Get("frustumCulled").Bool()
}
func (sm *SkinnedMesh) SetFrustumCulled(v bool) {
	sm.Set("frustumCulled", v)
}
func (sm *SkinnedMesh) Geometry() core.Geometry {
	return core.Geometry(sm.Get("geometry"))
}
func (sm *SkinnedMesh) SetGeometry(v core.Geometry) {
	sm.Set("geometry", v)
}
func (sm *SkinnedMesh) Id() int {
	return sm.Get("id").Int()
}
func (sm *SkinnedMesh) SetId(v int) {
	sm.Set("id", v)
}
func (sm *SkinnedMesh) IsMesh() true {
	return true(sm.Get("isMesh"))
}
func (sm *SkinnedMesh) SetIsMesh(v true) {
	sm.Set("isMesh", v)
}
func (sm *SkinnedMesh) IsObject3D() true {
	return true(sm.Get("isObject3D"))
}
func (sm *SkinnedMesh) SetIsObject3D(v true) {
	sm.Set("isObject3D", v)
}
func (sm *SkinnedMesh) Layers() *core.Layers {
	return &core.Layers{Value: sm.Get("layers")}
}
func (sm *SkinnedMesh) SetLayers(v *core.Layers) {
	sm.Set("layers", v)
}
func (sm *SkinnedMesh) Material() materials.Material {
	return materials.Material(sm.Get("material"))
}
func (sm *SkinnedMesh) SetMaterial(v materials.Material) {
	sm.Set("material", v)
}
func (sm *SkinnedMesh) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: sm.Get("matrix")}
}
func (sm *SkinnedMesh) SetMatrix(v *math.Matrix4) {
	sm.Set("matrix", v)
}
func (sm *SkinnedMesh) MatrixAutoUpdate() bool {
	return sm.Get("matrixAutoUpdate").Bool()
}
func (sm *SkinnedMesh) SetMatrixAutoUpdate(v bool) {
	sm.Set("matrixAutoUpdate", v)
}
func (sm *SkinnedMesh) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: sm.Get("matrixWorld")}
}
func (sm *SkinnedMesh) SetMatrixWorld(v *math.Matrix4) {
	sm.Set("matrixWorld", v)
}
func (sm *SkinnedMesh) MatrixWorldNeedsUpdate() bool {
	return sm.Get("matrixWorldNeedsUpdate").Bool()
}
func (sm *SkinnedMesh) SetMatrixWorldNeedsUpdate(v bool) {
	sm.Set("matrixWorldNeedsUpdate", v)
}
func (sm *SkinnedMesh) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: sm.Get("modelViewMatrix")}
}
func (sm *SkinnedMesh) SetModelViewMatrix(v *math.Matrix4) {
	sm.Set("modelViewMatrix", v)
}
func (sm *SkinnedMesh) MorphTargetDictionary() js.Value {
	return sm.Get("morphTargetDictionary")
}
func (sm *SkinnedMesh) SetMorphTargetDictionary(v js.Value) {
	sm.Set("morphTargetDictionary", v)
}
func (sm *SkinnedMesh) MorphTargetInfluences() []int {
	return []int(sm.Get("morphTargetInfluences"))
}
func (sm *SkinnedMesh) SetMorphTargetInfluences(v []int) {
	sm.Set("morphTargetInfluences", v)
}
func (sm *SkinnedMesh) Name() string {
	return sm.Get("name").String()
}
func (sm *SkinnedMesh) SetName(v string) {
	sm.Set("name", v)
}
func (sm *SkinnedMesh) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: sm.Get("normalMatrix")}
}
func (sm *SkinnedMesh) SetNormalMatrix(v *math.Matrix3) {
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
func (sm *SkinnedMesh) Parent() core.Object3D {
	return core.Object3D(sm.Get("parent"))
}
func (sm *SkinnedMesh) SetParent(v core.Object3D) {
	sm.Set("parent", v)
}
func (sm *SkinnedMesh) Position() *math.Vector3 {
	return &math.Vector3{Value: sm.Get("position")}
}
func (sm *SkinnedMesh) SetPosition(v *math.Vector3) {
	sm.Set("position", v)
}
func (sm *SkinnedMesh) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: sm.Get("quaternion")}
}
func (sm *SkinnedMesh) SetQuaternion(v *math.Quaternion) {
	sm.Set("quaternion", v)
}
func (sm *SkinnedMesh) ReceiveShadow() bool {
	return sm.Get("receiveShadow").Bool()
}
func (sm *SkinnedMesh) SetReceiveShadow(v bool) {
	sm.Set("receiveShadow", v)
}
func (sm *SkinnedMesh) RenderOrder() int {
	return sm.Get("renderOrder").Int()
}
func (sm *SkinnedMesh) SetRenderOrder(v int) {
	sm.Set("renderOrder", v)
}
func (sm *SkinnedMesh) Rotation() *math.Euler {
	return &math.Euler{Value: sm.Get("rotation")}
}
func (sm *SkinnedMesh) SetRotation(v *math.Euler) {
	sm.Set("rotation", v)
}
func (sm *SkinnedMesh) Scale() *math.Vector3 {
	return &math.Vector3{Value: sm.Get("scale")}
}
func (sm *SkinnedMesh) SetScale(v *math.Vector3) {
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
func (sm *SkinnedMesh) Up() *math.Vector3 {
	return &math.Vector3{Value: sm.Get("up")}
}
func (sm *SkinnedMesh) SetUp(v *math.Vector3) {
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
func (sm *SkinnedMesh) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: sm.Get("DefaultUp")}
}
func (sm *SkinnedMesh) SetDefaultUp(v *math.Vector3) {
	sm.Set("DefaultUp", v)
}
func (sm *SkinnedMesh) Add(object []core.Object3D) this {
	return this(sm.Call("add", object))
}
func (sm *SkinnedMesh) AddEventListener(typ string, listener js.Value) {
	sm.Call("addEventListener", typ, listener)
}
func (sm *SkinnedMesh) ApplyMatrix(matrix *math.Matrix4) {
	sm.Call("applyMatrix", matrix)
}
func (sm *SkinnedMesh) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(sm.Call("applyQuaternion", quaternion))
}
func (sm *SkinnedMesh) Bind(skeleton *Skeleton, bindMatrix *math.Matrix4) {
	sm.Call("bind", skeleton, bindMatrix)
}
func (sm *SkinnedMesh) Clone(recursive bool) this {
	return this(sm.Call("clone", recursive))
}
func (sm *SkinnedMesh) Copy(source this, recursive bool) this {
	return this(sm.Call("copy", source, recursive))
}
func (sm *SkinnedMesh) DispatchEvent(event js.Value) {
	sm.Call("dispatchEvent", event)
}
func (sm *SkinnedMesh) GetObjectById(id int) core.Object3D {
	return core.Object3D(sm.Call("getObjectById", id))
}
func (sm *SkinnedMesh) GetObjectByName(name string) core.Object3D {
	return core.Object3D(sm.Call("getObjectByName", name))
}
func (sm *SkinnedMesh) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(sm.Call("getObjectByProperty", name, value))
}
func (sm *SkinnedMesh) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sm.Call("getWorldDirection", target)}
}
func (sm *SkinnedMesh) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sm.Call("getWorldPosition", target)}
}
func (sm *SkinnedMesh) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: sm.Call("getWorldQuaternion", target)}
}
func (sm *SkinnedMesh) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sm.Call("getWorldScale", target)}
}
func (sm *SkinnedMesh) HasEventListener(typ string, listener js.Value) bool {
	return sm.Call("hasEventListener", typ, listener).Bool()
}
func (sm *SkinnedMesh) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sm.Call("localToWorld", vector)}
}
func (sm *SkinnedMesh) LookAt(vector math.Vector3, y int, z int) {
	sm.Call("lookAt", vector, y, z)
}
func (sm *SkinnedMesh) NormalizeSkinWeights() {
	sm.Call("normalizeSkinWeights")
}
func (sm *SkinnedMesh) Pose() {
	sm.Call("pose")
}
func (sm *SkinnedMesh) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	sm.Call("raycast", raycaster, intersects)
}
func (sm *SkinnedMesh) Remove(object []core.Object3D) this {
	return this(sm.Call("remove", object))
}
func (sm *SkinnedMesh) RemoveEventListener(typ string, listener js.Value) {
	sm.Call("removeEventListener", typ, listener)
}
func (sm *SkinnedMesh) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(sm.Call("rotateOnAxis", axis, angle))
}
func (sm *SkinnedMesh) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(sm.Call("rotateOnWorldAxis", axis, angle))
}
func (sm *SkinnedMesh) RotateX(angle int) this {
	return this(sm.Call("rotateX", angle))
}
func (sm *SkinnedMesh) RotateY(angle int) this {
	return this(sm.Call("rotateY", angle))
}
func (sm *SkinnedMesh) RotateZ(angle int) this {
	return this(sm.Call("rotateZ", angle))
}
func (sm *SkinnedMesh) SetDrawMode(drawMode *TrianglesDrawModes) {
	sm.Call("setDrawMode", drawMode)
}
func (sm *SkinnedMesh) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	sm.Call("setRotationFromAxisAngle", axis, angle)
}
func (sm *SkinnedMesh) SetRotationFromEuler(euler *math.Euler) {
	sm.Call("setRotationFromEuler", euler)
}
func (sm *SkinnedMesh) SetRotationFromMatrix(m *math.Matrix4) {
	sm.Call("setRotationFromMatrix", m)
}
func (sm *SkinnedMesh) SetRotationFromQuaternion(q *math.Quaternion) {
	sm.Call("setRotationFromQuaternion", q)
}
func (sm *SkinnedMesh) ToJSON(meta js.Value) js.Value {
	return sm.Call("toJSON", meta)
}
func (sm *SkinnedMesh) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(sm.Call("translateOnAxis", axis, distance))
}
func (sm *SkinnedMesh) TranslateX(distance int) this {
	return this(sm.Call("translateX", distance))
}
func (sm *SkinnedMesh) TranslateY(distance int) this {
	return this(sm.Call("translateY", distance))
}
func (sm *SkinnedMesh) TranslateZ(distance int) this {
	return this(sm.Call("translateZ", distance))
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
func (sm *SkinnedMesh) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sm.Call("worldToLocal", vector)}
}

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

type Mesh struct {
	js.Value
}

func NewMesh(geometry core.Geometry, material materials.Material) *Mesh {
	return &Mesh{Value: get("Mesh").New(geometry, material)}
}
func (m *Mesh) CastShadow() bool {
	return m.Get("castShadow").Bool()
}
func (m *Mesh) SetCastShadow(v bool) {
	m.Set("castShadow", v)
}
func (m *Mesh) Children() []core.Object3D {
	return []core.Object3D(m.Get("children"))
}
func (m *Mesh) SetChildren(v []core.Object3D) {
	m.Set("children", v)
}
func (m *Mesh) DrawMode() *TrianglesDrawModes {
	return &TrianglesDrawModes{Value: m.Get("drawMode")}
}
func (m *Mesh) SetDrawMode(v *TrianglesDrawModes) {
	m.Set("drawMode", v)
}
func (m *Mesh) FrustumCulled() bool {
	return m.Get("frustumCulled").Bool()
}
func (m *Mesh) SetFrustumCulled(v bool) {
	m.Set("frustumCulled", v)
}
func (m *Mesh) Geometry() core.Geometry {
	return core.Geometry(m.Get("geometry"))
}
func (m *Mesh) SetGeometry(v core.Geometry) {
	m.Set("geometry", v)
}
func (m *Mesh) Id() int {
	return m.Get("id").Int()
}
func (m *Mesh) SetId(v int) {
	m.Set("id", v)
}
func (m *Mesh) IsMesh() true {
	return true(m.Get("isMesh"))
}
func (m *Mesh) SetIsMesh(v true) {
	m.Set("isMesh", v)
}
func (m *Mesh) IsObject3D() true {
	return true(m.Get("isObject3D"))
}
func (m *Mesh) SetIsObject3D(v true) {
	m.Set("isObject3D", v)
}
func (m *Mesh) Layers() *core.Layers {
	return &core.Layers{Value: m.Get("layers")}
}
func (m *Mesh) SetLayers(v *core.Layers) {
	m.Set("layers", v)
}
func (m *Mesh) Material() materials.Material {
	return materials.Material(m.Get("material"))
}
func (m *Mesh) SetMaterial(v materials.Material) {
	m.Set("material", v)
}
func (m *Mesh) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: m.Get("matrix")}
}
func (m *Mesh) SetMatrix(v *math.Matrix4) {
	m.Set("matrix", v)
}
func (m *Mesh) MatrixAutoUpdate() bool {
	return m.Get("matrixAutoUpdate").Bool()
}
func (m *Mesh) SetMatrixAutoUpdate(v bool) {
	m.Set("matrixAutoUpdate", v)
}
func (m *Mesh) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: m.Get("matrixWorld")}
}
func (m *Mesh) SetMatrixWorld(v *math.Matrix4) {
	m.Set("matrixWorld", v)
}
func (m *Mesh) MatrixWorldNeedsUpdate() bool {
	return m.Get("matrixWorldNeedsUpdate").Bool()
}
func (m *Mesh) SetMatrixWorldNeedsUpdate(v bool) {
	m.Set("matrixWorldNeedsUpdate", v)
}
func (m *Mesh) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: m.Get("modelViewMatrix")}
}
func (m *Mesh) SetModelViewMatrix(v *math.Matrix4) {
	m.Set("modelViewMatrix", v)
}
func (m *Mesh) MorphTargetDictionary() js.Value {
	return m.Get("morphTargetDictionary")
}
func (m *Mesh) SetMorphTargetDictionary(v js.Value) {
	m.Set("morphTargetDictionary", v)
}
func (m *Mesh) MorphTargetInfluences() []int {
	return []int(m.Get("morphTargetInfluences"))
}
func (m *Mesh) SetMorphTargetInfluences(v []int) {
	m.Set("morphTargetInfluences", v)
}
func (m *Mesh) Name() string {
	return m.Get("name").String()
}
func (m *Mesh) SetName(v string) {
	m.Set("name", v)
}
func (m *Mesh) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: m.Get("normalMatrix")}
}
func (m *Mesh) SetNormalMatrix(v *math.Matrix3) {
	m.Set("normalMatrix", v)
}
func (m *Mesh) OnAfterRender() js.Value {
	return m.Get("onAfterRender")
}
func (m *Mesh) SetOnAfterRender(v js.Value) {
	m.Set("onAfterRender", v)
}
func (m *Mesh) OnBeforeRender() js.Value {
	return m.Get("onBeforeRender")
}
func (m *Mesh) SetOnBeforeRender(v js.Value) {
	m.Set("onBeforeRender", v)
}
func (m *Mesh) Parent() core.Object3D {
	return core.Object3D(m.Get("parent"))
}
func (m *Mesh) SetParent(v core.Object3D) {
	m.Set("parent", v)
}
func (m *Mesh) Position() *math.Vector3 {
	return &math.Vector3{Value: m.Get("position")}
}
func (m *Mesh) SetPosition(v *math.Vector3) {
	m.Set("position", v)
}
func (m *Mesh) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: m.Get("quaternion")}
}
func (m *Mesh) SetQuaternion(v *math.Quaternion) {
	m.Set("quaternion", v)
}
func (m *Mesh) ReceiveShadow() bool {
	return m.Get("receiveShadow").Bool()
}
func (m *Mesh) SetReceiveShadow(v bool) {
	m.Set("receiveShadow", v)
}
func (m *Mesh) RenderOrder() int {
	return m.Get("renderOrder").Int()
}
func (m *Mesh) SetRenderOrder(v int) {
	m.Set("renderOrder", v)
}
func (m *Mesh) Rotation() *math.Euler {
	return &math.Euler{Value: m.Get("rotation")}
}
func (m *Mesh) SetRotation(v *math.Euler) {
	m.Set("rotation", v)
}
func (m *Mesh) Scale() *math.Vector3 {
	return &math.Vector3{Value: m.Get("scale")}
}
func (m *Mesh) SetScale(v *math.Vector3) {
	m.Set("scale", v)
}
func (m *Mesh) Type() string {
	return m.Get("type").String()
}
func (m *Mesh) SetType(v string) {
	m.Set("type", v)
}
func (m *Mesh) Up() *math.Vector3 {
	return &math.Vector3{Value: m.Get("up")}
}
func (m *Mesh) SetUp(v *math.Vector3) {
	m.Set("up", v)
}
func (m *Mesh) UserData() js.Value {
	return m.Get("userData")
}
func (m *Mesh) SetUserData(v js.Value) {
	m.Set("userData", v)
}
func (m *Mesh) Uuid() string {
	return m.Get("uuid").String()
}
func (m *Mesh) SetUuid(v string) {
	m.Set("uuid", v)
}
func (m *Mesh) Visible() bool {
	return m.Get("visible").Bool()
}
func (m *Mesh) SetVisible(v bool) {
	m.Set("visible", v)
}
func (m *Mesh) DefaultMatrixAutoUpdate() bool {
	return m.Get("DefaultMatrixAutoUpdate").Bool()
}
func (m *Mesh) SetDefaultMatrixAutoUpdate(v bool) {
	m.Set("DefaultMatrixAutoUpdate", v)
}
func (m *Mesh) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: m.Get("DefaultUp")}
}
func (m *Mesh) SetDefaultUp(v *math.Vector3) {
	m.Set("DefaultUp", v)
}
func (m *Mesh) Add(object []core.Object3D) this {
	return this(m.Call("add", object))
}
func (m *Mesh) AddEventListener(typ string, listener js.Value) {
	m.Call("addEventListener", typ, listener)
}
func (m *Mesh) ApplyMatrix(matrix *math.Matrix4) {
	m.Call("applyMatrix", matrix)
}
func (m *Mesh) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(m.Call("applyQuaternion", quaternion))
}
func (m *Mesh) Clone(recursive bool) this {
	return this(m.Call("clone", recursive))
}
func (m *Mesh) Copy(source this, recursive bool) this {
	return this(m.Call("copy", source, recursive))
}
func (m *Mesh) DispatchEvent(event js.Value) {
	m.Call("dispatchEvent", event)
}
func (m *Mesh) GetObjectById(id int) core.Object3D {
	return core.Object3D(m.Call("getObjectById", id))
}
func (m *Mesh) GetObjectByName(name string) core.Object3D {
	return core.Object3D(m.Call("getObjectByName", name))
}
func (m *Mesh) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(m.Call("getObjectByProperty", name, value))
}
func (m *Mesh) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: m.Call("getWorldDirection", target)}
}
func (m *Mesh) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: m.Call("getWorldPosition", target)}
}
func (m *Mesh) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: m.Call("getWorldQuaternion", target)}
}
func (m *Mesh) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: m.Call("getWorldScale", target)}
}
func (m *Mesh) HasEventListener(typ string, listener js.Value) bool {
	return m.Call("hasEventListener", typ, listener).Bool()
}
func (m *Mesh) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: m.Call("localToWorld", vector)}
}
func (m *Mesh) LookAt(vector math.Vector3, y int, z int) {
	m.Call("lookAt", vector, y, z)
}
func (m *Mesh) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	m.Call("raycast", raycaster, intersects)
}
func (m *Mesh) Remove(object []core.Object3D) this {
	return this(m.Call("remove", object))
}
func (m *Mesh) RemoveEventListener(typ string, listener js.Value) {
	m.Call("removeEventListener", typ, listener)
}
func (m *Mesh) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(m.Call("rotateOnAxis", axis, angle))
}
func (m *Mesh) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(m.Call("rotateOnWorldAxis", axis, angle))
}
func (m *Mesh) RotateX(angle int) this {
	return this(m.Call("rotateX", angle))
}
func (m *Mesh) RotateY(angle int) this {
	return this(m.Call("rotateY", angle))
}
func (m *Mesh) RotateZ(angle int) this {
	return this(m.Call("rotateZ", angle))
}
func (m *Mesh) SetDrawMode(drawMode *TrianglesDrawModes) {
	m.Call("setDrawMode", drawMode)
}
func (m *Mesh) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	m.Call("setRotationFromAxisAngle", axis, angle)
}
func (m *Mesh) SetRotationFromEuler(euler *math.Euler) {
	m.Call("setRotationFromEuler", euler)
}
func (m *Mesh) SetRotationFromMatrix(m *math.Matrix4) {
	m.Call("setRotationFromMatrix", m)
}
func (m *Mesh) SetRotationFromQuaternion(q *math.Quaternion) {
	m.Call("setRotationFromQuaternion", q)
}
func (m *Mesh) ToJSON(meta js.Value) js.Value {
	return m.Call("toJSON", meta)
}
func (m *Mesh) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(m.Call("translateOnAxis", axis, distance))
}
func (m *Mesh) TranslateX(distance int) this {
	return this(m.Call("translateX", distance))
}
func (m *Mesh) TranslateY(distance int) this {
	return this(m.Call("translateY", distance))
}
func (m *Mesh) TranslateZ(distance int) this {
	return this(m.Call("translateZ", distance))
}
func (m *Mesh) Traverse(callback js.Value) {
	m.Call("traverse", callback)
}
func (m *Mesh) TraverseAncestors(callback js.Value) {
	m.Call("traverseAncestors", callback)
}
func (m *Mesh) TraverseVisible(callback js.Value) {
	m.Call("traverseVisible", callback)
}
func (m *Mesh) UpdateMatrix() {
	m.Call("updateMatrix")
}
func (m *Mesh) UpdateMatrixWorld(force bool) {
	m.Call("updateMatrixWorld", force)
}
func (m *Mesh) UpdateMorphTargets() {
	m.Call("updateMorphTargets")
}
func (m *Mesh) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	m.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (m *Mesh) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: m.Call("worldToLocal", vector)}
}

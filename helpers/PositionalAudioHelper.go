package helpers

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/audio"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type PositionalAudioHelper struct {
	js.Value
}

func NewPositionalAudioHelper(audio *audio.PositionalAudio, rng int, divisionsInnerAngle int, divisionsOuterAngle int) *PositionalAudioHelper {
	return &PositionalAudioHelper{Value: get("PositionalAudioHelper").New(audio, rng, divisionsInnerAngle, divisionsOuterAngle)}
}
func (pah *PositionalAudioHelper) Audio() *audio.PositionalAudio {
	return &audio.PositionalAudio{Value: pah.Get("audio")}
}
func (pah *PositionalAudioHelper) SetAudio(v *audio.PositionalAudio) {
	pah.Set("audio", v)
}
func (pah *PositionalAudioHelper) CastShadow() bool {
	return pah.Get("castShadow").Bool()
}
func (pah *PositionalAudioHelper) SetCastShadow(v bool) {
	pah.Set("castShadow", v)
}
func (pah *PositionalAudioHelper) Children() []core.Object3D {
	return []core.Object3D(pah.Get("children"))
}
func (pah *PositionalAudioHelper) SetChildren(v []core.Object3D) {
	pah.Set("children", v)
}
func (pah *PositionalAudioHelper) DivisionsInnerAngle() int {
	return pah.Get("divisionsInnerAngle").Int()
}
func (pah *PositionalAudioHelper) SetDivisionsInnerAngle(v int) {
	pah.Set("divisionsInnerAngle", v)
}
func (pah *PositionalAudioHelper) DivisionsOuterAngle() int {
	return pah.Get("divisionsOuterAngle").Int()
}
func (pah *PositionalAudioHelper) SetDivisionsOuterAngle(v int) {
	pah.Set("divisionsOuterAngle", v)
}
func (pah *PositionalAudioHelper) FrustumCulled() bool {
	return pah.Get("frustumCulled").Bool()
}
func (pah *PositionalAudioHelper) SetFrustumCulled(v bool) {
	pah.Set("frustumCulled", v)
}
func (pah *PositionalAudioHelper) Geometry() core.Geometry {
	return core.Geometry(pah.Get("geometry"))
}
func (pah *PositionalAudioHelper) SetGeometry(v core.Geometry) {
	pah.Set("geometry", v)
}
func (pah *PositionalAudioHelper) Id() int {
	return pah.Get("id").Int()
}
func (pah *PositionalAudioHelper) SetId(v int) {
	pah.Set("id", v)
}
func (pah *PositionalAudioHelper) IsLine() true {
	return true(pah.Get("isLine"))
}
func (pah *PositionalAudioHelper) SetIsLine(v true) {
	pah.Set("isLine", v)
}
func (pah *PositionalAudioHelper) IsObject3D() true {
	return true(pah.Get("isObject3D"))
}
func (pah *PositionalAudioHelper) SetIsObject3D(v true) {
	pah.Set("isObject3D", v)
}
func (pah *PositionalAudioHelper) Layers() *core.Layers {
	return &core.Layers{Value: pah.Get("layers")}
}
func (pah *PositionalAudioHelper) SetLayers(v *core.Layers) {
	pah.Set("layers", v)
}
func (pah *PositionalAudioHelper) Material() materials.Material {
	return materials.Material(pah.Get("material"))
}
func (pah *PositionalAudioHelper) SetMaterial(v materials.Material) {
	pah.Set("material", v)
}
func (pah *PositionalAudioHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: pah.Get("matrix")}
}
func (pah *PositionalAudioHelper) SetMatrix(v *math.Matrix4) {
	pah.Set("matrix", v)
}
func (pah *PositionalAudioHelper) MatrixAutoUpdate() bool {
	return pah.Get("matrixAutoUpdate").Bool()
}
func (pah *PositionalAudioHelper) SetMatrixAutoUpdate(v bool) {
	pah.Set("matrixAutoUpdate", v)
}
func (pah *PositionalAudioHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: pah.Get("matrixWorld")}
}
func (pah *PositionalAudioHelper) SetMatrixWorld(v *math.Matrix4) {
	pah.Set("matrixWorld", v)
}
func (pah *PositionalAudioHelper) MatrixWorldNeedsUpdate() bool {
	return pah.Get("matrixWorldNeedsUpdate").Bool()
}
func (pah *PositionalAudioHelper) SetMatrixWorldNeedsUpdate(v bool) {
	pah.Set("matrixWorldNeedsUpdate", v)
}
func (pah *PositionalAudioHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: pah.Get("modelViewMatrix")}
}
func (pah *PositionalAudioHelper) SetModelViewMatrix(v *math.Matrix4) {
	pah.Set("modelViewMatrix", v)
}
func (pah *PositionalAudioHelper) Name() string {
	return pah.Get("name").String()
}
func (pah *PositionalAudioHelper) SetName(v string) {
	pah.Set("name", v)
}
func (pah *PositionalAudioHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: pah.Get("normalMatrix")}
}
func (pah *PositionalAudioHelper) SetNormalMatrix(v *math.Matrix3) {
	pah.Set("normalMatrix", v)
}
func (pah *PositionalAudioHelper) OnAfterRender() js.Value {
	return pah.Get("onAfterRender")
}
func (pah *PositionalAudioHelper) SetOnAfterRender(v js.Value) {
	pah.Set("onAfterRender", v)
}
func (pah *PositionalAudioHelper) OnBeforeRender() js.Value {
	return pah.Get("onBeforeRender")
}
func (pah *PositionalAudioHelper) SetOnBeforeRender(v js.Value) {
	pah.Set("onBeforeRender", v)
}
func (pah *PositionalAudioHelper) Parent() core.Object3D {
	return core.Object3D(pah.Get("parent"))
}
func (pah *PositionalAudioHelper) SetParent(v core.Object3D) {
	pah.Set("parent", v)
}
func (pah *PositionalAudioHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: pah.Get("position")}
}
func (pah *PositionalAudioHelper) SetPosition(v *math.Vector3) {
	pah.Set("position", v)
}
func (pah *PositionalAudioHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: pah.Get("quaternion")}
}
func (pah *PositionalAudioHelper) SetQuaternion(v *math.Quaternion) {
	pah.Set("quaternion", v)
}
func (pah *PositionalAudioHelper) Range() int {
	return pah.Get("range").Int()
}
func (pah *PositionalAudioHelper) SetRange(v int) {
	pah.Set("range", v)
}
func (pah *PositionalAudioHelper) ReceiveShadow() bool {
	return pah.Get("receiveShadow").Bool()
}
func (pah *PositionalAudioHelper) SetReceiveShadow(v bool) {
	pah.Set("receiveShadow", v)
}
func (pah *PositionalAudioHelper) RenderOrder() int {
	return pah.Get("renderOrder").Int()
}
func (pah *PositionalAudioHelper) SetRenderOrder(v int) {
	pah.Set("renderOrder", v)
}
func (pah *PositionalAudioHelper) Rotation() *math.Euler {
	return &math.Euler{Value: pah.Get("rotation")}
}
func (pah *PositionalAudioHelper) SetRotation(v *math.Euler) {
	pah.Set("rotation", v)
}
func (pah *PositionalAudioHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: pah.Get("scale")}
}
func (pah *PositionalAudioHelper) SetScale(v *math.Vector3) {
	pah.Set("scale", v)
}
func (pah *PositionalAudioHelper) Type() string {
	return pah.Get("type").String()
}
func (pah *PositionalAudioHelper) SetType(v string) {
	pah.Set("type", v)
}
func (pah *PositionalAudioHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: pah.Get("up")}
}
func (pah *PositionalAudioHelper) SetUp(v *math.Vector3) {
	pah.Set("up", v)
}
func (pah *PositionalAudioHelper) UserData() js.Value {
	return pah.Get("userData")
}
func (pah *PositionalAudioHelper) SetUserData(v js.Value) {
	pah.Set("userData", v)
}
func (pah *PositionalAudioHelper) Uuid() string {
	return pah.Get("uuid").String()
}
func (pah *PositionalAudioHelper) SetUuid(v string) {
	pah.Set("uuid", v)
}
func (pah *PositionalAudioHelper) Visible() bool {
	return pah.Get("visible").Bool()
}
func (pah *PositionalAudioHelper) SetVisible(v bool) {
	pah.Set("visible", v)
}
func (pah *PositionalAudioHelper) DefaultMatrixAutoUpdate() bool {
	return pah.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pah *PositionalAudioHelper) SetDefaultMatrixAutoUpdate(v bool) {
	pah.Set("DefaultMatrixAutoUpdate", v)
}
func (pah *PositionalAudioHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: pah.Get("DefaultUp")}
}
func (pah *PositionalAudioHelper) SetDefaultUp(v *math.Vector3) {
	pah.Set("DefaultUp", v)
}
func (pah *PositionalAudioHelper) Add(object []core.Object3D) this {
	return this(pah.Call("add", object))
}
func (pah *PositionalAudioHelper) AddEventListener(typ string, listener js.Value) {
	pah.Call("addEventListener", typ, listener)
}
func (pah *PositionalAudioHelper) ApplyMatrix(matrix *math.Matrix4) {
	pah.Call("applyMatrix", matrix)
}
func (pah *PositionalAudioHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(pah.Call("applyQuaternion", quaternion))
}
func (pah *PositionalAudioHelper) Clone(recursive bool) this {
	return this(pah.Call("clone", recursive))
}
func (pah *PositionalAudioHelper) ComputeLineDistances() this {
	return this(pah.Call("computeLineDistances"))
}
func (pah *PositionalAudioHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(pah.Call("copy", source, recursive))
}
func (pah *PositionalAudioHelper) DispatchEvent(event js.Value) {
	pah.Call("dispatchEvent", event)
}
func (pah *PositionalAudioHelper) Dispose() {
	pah.Call("dispose")
}
func (pah *PositionalAudioHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(pah.Call("getObjectById", id))
}
func (pah *PositionalAudioHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(pah.Call("getObjectByName", name))
}
func (pah *PositionalAudioHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(pah.Call("getObjectByProperty", name, value))
}
func (pah *PositionalAudioHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pah.Call("getWorldDirection", target)}
}
func (pah *PositionalAudioHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pah.Call("getWorldPosition", target)}
}
func (pah *PositionalAudioHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: pah.Call("getWorldQuaternion", target)}
}
func (pah *PositionalAudioHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pah.Call("getWorldScale", target)}
}
func (pah *PositionalAudioHelper) HasEventListener(typ string, listener js.Value) bool {
	return pah.Call("hasEventListener", typ, listener).Bool()
}
func (pah *PositionalAudioHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pah.Call("localToWorld", vector)}
}
func (pah *PositionalAudioHelper) LookAt(vector math.Vector3, y int, z int) {
	pah.Call("lookAt", vector, y, z)
}
func (pah *PositionalAudioHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	pah.Call("raycast", raycaster, intersects)
}
func (pah *PositionalAudioHelper) Remove(object []core.Object3D) this {
	return this(pah.Call("remove", object))
}
func (pah *PositionalAudioHelper) RemoveEventListener(typ string, listener js.Value) {
	pah.Call("removeEventListener", typ, listener)
}
func (pah *PositionalAudioHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(pah.Call("rotateOnAxis", axis, angle))
}
func (pah *PositionalAudioHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(pah.Call("rotateOnWorldAxis", axis, angle))
}
func (pah *PositionalAudioHelper) RotateX(angle int) this {
	return this(pah.Call("rotateX", angle))
}
func (pah *PositionalAudioHelper) RotateY(angle int) this {
	return this(pah.Call("rotateY", angle))
}
func (pah *PositionalAudioHelper) RotateZ(angle int) this {
	return this(pah.Call("rotateZ", angle))
}
func (pah *PositionalAudioHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	pah.Call("setRotationFromAxisAngle", axis, angle)
}
func (pah *PositionalAudioHelper) SetRotationFromEuler(euler *math.Euler) {
	pah.Call("setRotationFromEuler", euler)
}
func (pah *PositionalAudioHelper) SetRotationFromMatrix(m *math.Matrix4) {
	pah.Call("setRotationFromMatrix", m)
}
func (pah *PositionalAudioHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	pah.Call("setRotationFromQuaternion", q)
}
func (pah *PositionalAudioHelper) ToJSON(meta js.Value) js.Value {
	return pah.Call("toJSON", meta)
}
func (pah *PositionalAudioHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(pah.Call("translateOnAxis", axis, distance))
}
func (pah *PositionalAudioHelper) TranslateX(distance int) this {
	return this(pah.Call("translateX", distance))
}
func (pah *PositionalAudioHelper) TranslateY(distance int) this {
	return this(pah.Call("translateY", distance))
}
func (pah *PositionalAudioHelper) TranslateZ(distance int) this {
	return this(pah.Call("translateZ", distance))
}
func (pah *PositionalAudioHelper) Traverse(callback js.Value) {
	pah.Call("traverse", callback)
}
func (pah *PositionalAudioHelper) TraverseAncestors(callback js.Value) {
	pah.Call("traverseAncestors", callback)
}
func (pah *PositionalAudioHelper) TraverseVisible(callback js.Value) {
	pah.Call("traverseVisible", callback)
}
func (pah *PositionalAudioHelper) Update() {
	pah.Call("update")
}
func (pah *PositionalAudioHelper) UpdateMatrix() {
	pah.Call("updateMatrix")
}
func (pah *PositionalAudioHelper) UpdateMatrixWorld(force bool) {
	pah.Call("updateMatrixWorld", force)
}
func (pah *PositionalAudioHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pah.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pah *PositionalAudioHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pah.Call("worldToLocal", vector)}
}

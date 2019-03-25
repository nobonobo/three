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

type LineLoop struct {
	js.Value
}

func NewLineLoop(geometry core.Geometry, material materials.Material) *LineLoop {
	return &LineLoop{Value: get("LineLoop").New(geometry, material)}
}
func (ll *LineLoop) CastShadow() bool {
	return ll.Get("castShadow").Bool()
}
func (ll *LineLoop) SetCastShadow(v bool) {
	ll.Set("castShadow", v)
}
func (ll *LineLoop) Children() []core.Object3D {
	return []core.Object3D(ll.Get("children"))
}
func (ll *LineLoop) SetChildren(v []core.Object3D) {
	ll.Set("children", v)
}
func (ll *LineLoop) FrustumCulled() bool {
	return ll.Get("frustumCulled").Bool()
}
func (ll *LineLoop) SetFrustumCulled(v bool) {
	ll.Set("frustumCulled", v)
}
func (ll *LineLoop) Geometry() core.Geometry {
	return core.Geometry(ll.Get("geometry"))
}
func (ll *LineLoop) SetGeometry(v core.Geometry) {
	ll.Set("geometry", v)
}
func (ll *LineLoop) Id() int {
	return ll.Get("id").Int()
}
func (ll *LineLoop) SetId(v int) {
	ll.Set("id", v)
}
func (ll *LineLoop) IsLine() true {
	return true(ll.Get("isLine"))
}
func (ll *LineLoop) SetIsLine(v true) {
	ll.Set("isLine", v)
}
func (ll *LineLoop) IsLineLoop() true {
	return true(ll.Get("isLineLoop"))
}
func (ll *LineLoop) SetIsLineLoop(v true) {
	ll.Set("isLineLoop", v)
}
func (ll *LineLoop) IsObject3D() true {
	return true(ll.Get("isObject3D"))
}
func (ll *LineLoop) SetIsObject3D(v true) {
	ll.Set("isObject3D", v)
}
func (ll *LineLoop) Layers() *core.Layers {
	return &core.Layers{Value: ll.Get("layers")}
}
func (ll *LineLoop) SetLayers(v *core.Layers) {
	ll.Set("layers", v)
}
func (ll *LineLoop) Material() materials.Material {
	return materials.Material(ll.Get("material"))
}
func (ll *LineLoop) SetMaterial(v materials.Material) {
	ll.Set("material", v)
}
func (ll *LineLoop) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ll.Get("matrix")}
}
func (ll *LineLoop) SetMatrix(v *math.Matrix4) {
	ll.Set("matrix", v)
}
func (ll *LineLoop) MatrixAutoUpdate() bool {
	return ll.Get("matrixAutoUpdate").Bool()
}
func (ll *LineLoop) SetMatrixAutoUpdate(v bool) {
	ll.Set("matrixAutoUpdate", v)
}
func (ll *LineLoop) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ll.Get("matrixWorld")}
}
func (ll *LineLoop) SetMatrixWorld(v *math.Matrix4) {
	ll.Set("matrixWorld", v)
}
func (ll *LineLoop) MatrixWorldNeedsUpdate() bool {
	return ll.Get("matrixWorldNeedsUpdate").Bool()
}
func (ll *LineLoop) SetMatrixWorldNeedsUpdate(v bool) {
	ll.Set("matrixWorldNeedsUpdate", v)
}
func (ll *LineLoop) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ll.Get("modelViewMatrix")}
}
func (ll *LineLoop) SetModelViewMatrix(v *math.Matrix4) {
	ll.Set("modelViewMatrix", v)
}
func (ll *LineLoop) Name() string {
	return ll.Get("name").String()
}
func (ll *LineLoop) SetName(v string) {
	ll.Set("name", v)
}
func (ll *LineLoop) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ll.Get("normalMatrix")}
}
func (ll *LineLoop) SetNormalMatrix(v *math.Matrix3) {
	ll.Set("normalMatrix", v)
}
func (ll *LineLoop) OnAfterRender() js.Value {
	return ll.Get("onAfterRender")
}
func (ll *LineLoop) SetOnAfterRender(v js.Value) {
	ll.Set("onAfterRender", v)
}
func (ll *LineLoop) OnBeforeRender() js.Value {
	return ll.Get("onBeforeRender")
}
func (ll *LineLoop) SetOnBeforeRender(v js.Value) {
	ll.Set("onBeforeRender", v)
}
func (ll *LineLoop) Parent() core.Object3D {
	return core.Object3D(ll.Get("parent"))
}
func (ll *LineLoop) SetParent(v core.Object3D) {
	ll.Set("parent", v)
}
func (ll *LineLoop) Position() *math.Vector3 {
	return &math.Vector3{Value: ll.Get("position")}
}
func (ll *LineLoop) SetPosition(v *math.Vector3) {
	ll.Set("position", v)
}
func (ll *LineLoop) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ll.Get("quaternion")}
}
func (ll *LineLoop) SetQuaternion(v *math.Quaternion) {
	ll.Set("quaternion", v)
}
func (ll *LineLoop) ReceiveShadow() bool {
	return ll.Get("receiveShadow").Bool()
}
func (ll *LineLoop) SetReceiveShadow(v bool) {
	ll.Set("receiveShadow", v)
}
func (ll *LineLoop) RenderOrder() int {
	return ll.Get("renderOrder").Int()
}
func (ll *LineLoop) SetRenderOrder(v int) {
	ll.Set("renderOrder", v)
}
func (ll *LineLoop) Rotation() *math.Euler {
	return &math.Euler{Value: ll.Get("rotation")}
}
func (ll *LineLoop) SetRotation(v *math.Euler) {
	ll.Set("rotation", v)
}
func (ll *LineLoop) Scale() *math.Vector3 {
	return &math.Vector3{Value: ll.Get("scale")}
}
func (ll *LineLoop) SetScale(v *math.Vector3) {
	ll.Set("scale", v)
}
func (ll *LineLoop) Type() string {
	return ll.Get("type").String()
}
func (ll *LineLoop) SetType(v string) {
	ll.Set("type", v)
}
func (ll *LineLoop) Up() *math.Vector3 {
	return &math.Vector3{Value: ll.Get("up")}
}
func (ll *LineLoop) SetUp(v *math.Vector3) {
	ll.Set("up", v)
}
func (ll *LineLoop) UserData() js.Value {
	return ll.Get("userData")
}
func (ll *LineLoop) SetUserData(v js.Value) {
	ll.Set("userData", v)
}
func (ll *LineLoop) Uuid() string {
	return ll.Get("uuid").String()
}
func (ll *LineLoop) SetUuid(v string) {
	ll.Set("uuid", v)
}
func (ll *LineLoop) Visible() bool {
	return ll.Get("visible").Bool()
}
func (ll *LineLoop) SetVisible(v bool) {
	ll.Set("visible", v)
}
func (ll *LineLoop) DefaultMatrixAutoUpdate() bool {
	return ll.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ll *LineLoop) SetDefaultMatrixAutoUpdate(v bool) {
	ll.Set("DefaultMatrixAutoUpdate", v)
}
func (ll *LineLoop) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ll.Get("DefaultUp")}
}
func (ll *LineLoop) SetDefaultUp(v *math.Vector3) {
	ll.Set("DefaultUp", v)
}
func (ll *LineLoop) Add(object []core.Object3D) this {
	return this(ll.Call("add", object))
}
func (ll *LineLoop) AddEventListener(typ string, listener js.Value) {
	ll.Call("addEventListener", typ, listener)
}
func (ll *LineLoop) ApplyMatrix(matrix *math.Matrix4) {
	ll.Call("applyMatrix", matrix)
}
func (ll *LineLoop) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ll.Call("applyQuaternion", quaternion))
}
func (ll *LineLoop) Clone(recursive bool) this {
	return this(ll.Call("clone", recursive))
}
func (ll *LineLoop) ComputeLineDistances() this {
	return this(ll.Call("computeLineDistances"))
}
func (ll *LineLoop) Copy(source *core.Object3D, recursive bool) this {
	return this(ll.Call("copy", source, recursive))
}
func (ll *LineLoop) DispatchEvent(event js.Value) {
	ll.Call("dispatchEvent", event)
}
func (ll *LineLoop) GetObjectById(id int) core.Object3D {
	return core.Object3D(ll.Call("getObjectById", id))
}
func (ll *LineLoop) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ll.Call("getObjectByName", name))
}
func (ll *LineLoop) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ll.Call("getObjectByProperty", name, value))
}
func (ll *LineLoop) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ll.Call("getWorldDirection", target)}
}
func (ll *LineLoop) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ll.Call("getWorldPosition", target)}
}
func (ll *LineLoop) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ll.Call("getWorldQuaternion", target)}
}
func (ll *LineLoop) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ll.Call("getWorldScale", target)}
}
func (ll *LineLoop) HasEventListener(typ string, listener js.Value) bool {
	return ll.Call("hasEventListener", typ, listener).Bool()
}
func (ll *LineLoop) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ll.Call("localToWorld", vector)}
}
func (ll *LineLoop) LookAt(vector math.Vector3, y int, z int) {
	ll.Call("lookAt", vector, y, z)
}
func (ll *LineLoop) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ll.Call("raycast", raycaster, intersects)
}
func (ll *LineLoop) Remove(object []core.Object3D) this {
	return this(ll.Call("remove", object))
}
func (ll *LineLoop) RemoveEventListener(typ string, listener js.Value) {
	ll.Call("removeEventListener", typ, listener)
}
func (ll *LineLoop) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ll.Call("rotateOnAxis", axis, angle))
}
func (ll *LineLoop) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ll.Call("rotateOnWorldAxis", axis, angle))
}
func (ll *LineLoop) RotateX(angle int) this {
	return this(ll.Call("rotateX", angle))
}
func (ll *LineLoop) RotateY(angle int) this {
	return this(ll.Call("rotateY", angle))
}
func (ll *LineLoop) RotateZ(angle int) this {
	return this(ll.Call("rotateZ", angle))
}
func (ll *LineLoop) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ll.Call("setRotationFromAxisAngle", axis, angle)
}
func (ll *LineLoop) SetRotationFromEuler(euler *math.Euler) {
	ll.Call("setRotationFromEuler", euler)
}
func (ll *LineLoop) SetRotationFromMatrix(m *math.Matrix4) {
	ll.Call("setRotationFromMatrix", m)
}
func (ll *LineLoop) SetRotationFromQuaternion(q *math.Quaternion) {
	ll.Call("setRotationFromQuaternion", q)
}
func (ll *LineLoop) ToJSON(meta js.Value) js.Value {
	return ll.Call("toJSON", meta)
}
func (ll *LineLoop) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ll.Call("translateOnAxis", axis, distance))
}
func (ll *LineLoop) TranslateX(distance int) this {
	return this(ll.Call("translateX", distance))
}
func (ll *LineLoop) TranslateY(distance int) this {
	return this(ll.Call("translateY", distance))
}
func (ll *LineLoop) TranslateZ(distance int) this {
	return this(ll.Call("translateZ", distance))
}
func (ll *LineLoop) Traverse(callback js.Value) {
	ll.Call("traverse", callback)
}
func (ll *LineLoop) TraverseAncestors(callback js.Value) {
	ll.Call("traverseAncestors", callback)
}
func (ll *LineLoop) TraverseVisible(callback js.Value) {
	ll.Call("traverseVisible", callback)
}
func (ll *LineLoop) UpdateMatrix() {
	ll.Call("updateMatrix")
}
func (ll *LineLoop) UpdateMatrixWorld(force bool) {
	ll.Call("updateMatrixWorld", force)
}
func (ll *LineLoop) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ll.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ll *LineLoop) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ll.Call("worldToLocal", vector)}
}

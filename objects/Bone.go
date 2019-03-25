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

type Bone struct {
	js.Value
}

func NewBone() *Bone {
	return &Bone{Value: get("Bone").New()}
}
func (b *Bone) CastShadow() bool {
	return b.Get("castShadow").Bool()
}
func (b *Bone) SetCastShadow(v bool) {
	b.Set("castShadow", v)
}
func (b *Bone) Children() []core.Object3D {
	return []core.Object3D(b.Get("children"))
}
func (b *Bone) SetChildren(v []core.Object3D) {
	b.Set("children", v)
}
func (b *Bone) FrustumCulled() bool {
	return b.Get("frustumCulled").Bool()
}
func (b *Bone) SetFrustumCulled(v bool) {
	b.Set("frustumCulled", v)
}
func (b *Bone) Id() int {
	return b.Get("id").Int()
}
func (b *Bone) SetId(v int) {
	b.Set("id", v)
}
func (b *Bone) IsBone() true {
	return true(b.Get("isBone"))
}
func (b *Bone) SetIsBone(v true) {
	b.Set("isBone", v)
}
func (b *Bone) IsObject3D() true {
	return true(b.Get("isObject3D"))
}
func (b *Bone) SetIsObject3D(v true) {
	b.Set("isObject3D", v)
}
func (b *Bone) Layers() *core.Layers {
	return &core.Layers{Value: b.Get("layers")}
}
func (b *Bone) SetLayers(v *core.Layers) {
	b.Set("layers", v)
}
func (b *Bone) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: b.Get("matrix")}
}
func (b *Bone) SetMatrix(v *math.Matrix4) {
	b.Set("matrix", v)
}
func (b *Bone) MatrixAutoUpdate() bool {
	return b.Get("matrixAutoUpdate").Bool()
}
func (b *Bone) SetMatrixAutoUpdate(v bool) {
	b.Set("matrixAutoUpdate", v)
}
func (b *Bone) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: b.Get("matrixWorld")}
}
func (b *Bone) SetMatrixWorld(v *math.Matrix4) {
	b.Set("matrixWorld", v)
}
func (b *Bone) MatrixWorldNeedsUpdate() bool {
	return b.Get("matrixWorldNeedsUpdate").Bool()
}
func (b *Bone) SetMatrixWorldNeedsUpdate(v bool) {
	b.Set("matrixWorldNeedsUpdate", v)
}
func (b *Bone) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: b.Get("modelViewMatrix")}
}
func (b *Bone) SetModelViewMatrix(v *math.Matrix4) {
	b.Set("modelViewMatrix", v)
}
func (b *Bone) Name() string {
	return b.Get("name").String()
}
func (b *Bone) SetName(v string) {
	b.Set("name", v)
}
func (b *Bone) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: b.Get("normalMatrix")}
}
func (b *Bone) SetNormalMatrix(v *math.Matrix3) {
	b.Set("normalMatrix", v)
}
func (b *Bone) OnAfterRender() js.Value {
	return b.Get("onAfterRender")
}
func (b *Bone) SetOnAfterRender(v js.Value) {
	b.Set("onAfterRender", v)
}
func (b *Bone) OnBeforeRender() js.Value {
	return b.Get("onBeforeRender")
}
func (b *Bone) SetOnBeforeRender(v js.Value) {
	b.Set("onBeforeRender", v)
}
func (b *Bone) Parent() core.Object3D {
	return core.Object3D(b.Get("parent"))
}
func (b *Bone) SetParent(v core.Object3D) {
	b.Set("parent", v)
}
func (b *Bone) Position() *math.Vector3 {
	return &math.Vector3{Value: b.Get("position")}
}
func (b *Bone) SetPosition(v *math.Vector3) {
	b.Set("position", v)
}
func (b *Bone) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: b.Get("quaternion")}
}
func (b *Bone) SetQuaternion(v *math.Quaternion) {
	b.Set("quaternion", v)
}
func (b *Bone) ReceiveShadow() bool {
	return b.Get("receiveShadow").Bool()
}
func (b *Bone) SetReceiveShadow(v bool) {
	b.Set("receiveShadow", v)
}
func (b *Bone) RenderOrder() int {
	return b.Get("renderOrder").Int()
}
func (b *Bone) SetRenderOrder(v int) {
	b.Set("renderOrder", v)
}
func (b *Bone) Rotation() *math.Euler {
	return &math.Euler{Value: b.Get("rotation")}
}
func (b *Bone) SetRotation(v *math.Euler) {
	b.Set("rotation", v)
}
func (b *Bone) Scale() *math.Vector3 {
	return &math.Vector3{Value: b.Get("scale")}
}
func (b *Bone) SetScale(v *math.Vector3) {
	b.Set("scale", v)
}
func (b *Bone) Type() string {
	return b.Get("type").String()
}
func (b *Bone) SetType(v string) {
	b.Set("type", v)
}
func (b *Bone) Up() *math.Vector3 {
	return &math.Vector3{Value: b.Get("up")}
}
func (b *Bone) SetUp(v *math.Vector3) {
	b.Set("up", v)
}
func (b *Bone) UserData() js.Value {
	return b.Get("userData")
}
func (b *Bone) SetUserData(v js.Value) {
	b.Set("userData", v)
}
func (b *Bone) Uuid() string {
	return b.Get("uuid").String()
}
func (b *Bone) SetUuid(v string) {
	b.Set("uuid", v)
}
func (b *Bone) Visible() bool {
	return b.Get("visible").Bool()
}
func (b *Bone) SetVisible(v bool) {
	b.Set("visible", v)
}
func (b *Bone) DefaultMatrixAutoUpdate() bool {
	return b.Get("DefaultMatrixAutoUpdate").Bool()
}
func (b *Bone) SetDefaultMatrixAutoUpdate(v bool) {
	b.Set("DefaultMatrixAutoUpdate", v)
}
func (b *Bone) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: b.Get("DefaultUp")}
}
func (b *Bone) SetDefaultUp(v *math.Vector3) {
	b.Set("DefaultUp", v)
}
func (b *Bone) Add(object []core.Object3D) this {
	return this(b.Call("add", object))
}
func (b *Bone) AddEventListener(typ string, listener js.Value) {
	b.Call("addEventListener", typ, listener)
}
func (b *Bone) ApplyMatrix(matrix *math.Matrix4) {
	b.Call("applyMatrix", matrix)
}
func (b *Bone) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(b.Call("applyQuaternion", quaternion))
}
func (b *Bone) Clone(recursive bool) this {
	return this(b.Call("clone", recursive))
}
func (b *Bone) Copy(source *core.Object3D, recursive bool) this {
	return this(b.Call("copy", source, recursive))
}
func (b *Bone) DispatchEvent(event js.Value) {
	b.Call("dispatchEvent", event)
}
func (b *Bone) GetObjectById(id int) core.Object3D {
	return core.Object3D(b.Call("getObjectById", id))
}
func (b *Bone) GetObjectByName(name string) core.Object3D {
	return core.Object3D(b.Call("getObjectByName", name))
}
func (b *Bone) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(b.Call("getObjectByProperty", name, value))
}
func (b *Bone) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: b.Call("getWorldDirection", target)}
}
func (b *Bone) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: b.Call("getWorldPosition", target)}
}
func (b *Bone) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: b.Call("getWorldQuaternion", target)}
}
func (b *Bone) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: b.Call("getWorldScale", target)}
}
func (b *Bone) HasEventListener(typ string, listener js.Value) bool {
	return b.Call("hasEventListener", typ, listener).Bool()
}
func (b *Bone) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: b.Call("localToWorld", vector)}
}
func (b *Bone) LookAt(vector math.Vector3, y int, z int) {
	b.Call("lookAt", vector, y, z)
}
func (b *Bone) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	b.Call("raycast", raycaster, intersects)
}
func (b *Bone) Remove(object []core.Object3D) this {
	return this(b.Call("remove", object))
}
func (b *Bone) RemoveEventListener(typ string, listener js.Value) {
	b.Call("removeEventListener", typ, listener)
}
func (b *Bone) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(b.Call("rotateOnAxis", axis, angle))
}
func (b *Bone) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(b.Call("rotateOnWorldAxis", axis, angle))
}
func (b *Bone) RotateX(angle int) this {
	return this(b.Call("rotateX", angle))
}
func (b *Bone) RotateY(angle int) this {
	return this(b.Call("rotateY", angle))
}
func (b *Bone) RotateZ(angle int) this {
	return this(b.Call("rotateZ", angle))
}
func (b *Bone) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	b.Call("setRotationFromAxisAngle", axis, angle)
}
func (b *Bone) SetRotationFromEuler(euler *math.Euler) {
	b.Call("setRotationFromEuler", euler)
}
func (b *Bone) SetRotationFromMatrix(m *math.Matrix4) {
	b.Call("setRotationFromMatrix", m)
}
func (b *Bone) SetRotationFromQuaternion(q *math.Quaternion) {
	b.Call("setRotationFromQuaternion", q)
}
func (b *Bone) ToJSON(meta js.Value) js.Value {
	return b.Call("toJSON", meta)
}
func (b *Bone) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(b.Call("translateOnAxis", axis, distance))
}
func (b *Bone) TranslateX(distance int) this {
	return this(b.Call("translateX", distance))
}
func (b *Bone) TranslateY(distance int) this {
	return this(b.Call("translateY", distance))
}
func (b *Bone) TranslateZ(distance int) this {
	return this(b.Call("translateZ", distance))
}
func (b *Bone) Traverse(callback js.Value) {
	b.Call("traverse", callback)
}
func (b *Bone) TraverseAncestors(callback js.Value) {
	b.Call("traverseAncestors", callback)
}
func (b *Bone) TraverseVisible(callback js.Value) {
	b.Call("traverseVisible", callback)
}
func (b *Bone) UpdateMatrix() {
	b.Call("updateMatrix")
}
func (b *Bone) UpdateMatrixWorld(force bool) {
	b.Call("updateMatrixWorld", force)
}
func (b *Bone) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	b.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (b *Bone) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: b.Call("worldToLocal", vector)}
}

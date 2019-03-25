package objects

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type ImmediateRenderObject struct {
	js.Value
}

func NewImmediateRenderObject(material *materials.Material) *ImmediateRenderObject {
	return &ImmediateRenderObject{Value: get("ImmediateRenderObject").New(material)}
}
func (iro *ImmediateRenderObject) CastShadow() bool {
	return iro.Get("castShadow").Bool()
}
func (iro *ImmediateRenderObject) SetCastShadow(v bool) {
	iro.Set("castShadow", v)
}
func (iro *ImmediateRenderObject) Children() []core.Object3D {
	return []core.Object3D(iro.Get("children"))
}
func (iro *ImmediateRenderObject) SetChildren(v []core.Object3D) {
	iro.Set("children", v)
}
func (iro *ImmediateRenderObject) FrustumCulled() bool {
	return iro.Get("frustumCulled").Bool()
}
func (iro *ImmediateRenderObject) SetFrustumCulled(v bool) {
	iro.Set("frustumCulled", v)
}
func (iro *ImmediateRenderObject) Id() int {
	return iro.Get("id").Int()
}
func (iro *ImmediateRenderObject) SetId(v int) {
	iro.Set("id", v)
}
func (iro *ImmediateRenderObject) IsObject3D() true {
	return true(iro.Get("isObject3D"))
}
func (iro *ImmediateRenderObject) SetIsObject3D(v true) {
	iro.Set("isObject3D", v)
}
func (iro *ImmediateRenderObject) Layers() *core.Layers {
	return &core.Layers{Value: iro.Get("layers")}
}
func (iro *ImmediateRenderObject) SetLayers(v *core.Layers) {
	iro.Set("layers", v)
}
func (iro *ImmediateRenderObject) Material() *materials.Material {
	return &materials.Material{Value: iro.Get("material")}
}
func (iro *ImmediateRenderObject) SetMaterial(v *materials.Material) {
	iro.Set("material", v)
}
func (iro *ImmediateRenderObject) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: iro.Get("matrix")}
}
func (iro *ImmediateRenderObject) SetMatrix(v *math.Matrix4) {
	iro.Set("matrix", v)
}
func (iro *ImmediateRenderObject) MatrixAutoUpdate() bool {
	return iro.Get("matrixAutoUpdate").Bool()
}
func (iro *ImmediateRenderObject) SetMatrixAutoUpdate(v bool) {
	iro.Set("matrixAutoUpdate", v)
}
func (iro *ImmediateRenderObject) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: iro.Get("matrixWorld")}
}
func (iro *ImmediateRenderObject) SetMatrixWorld(v *math.Matrix4) {
	iro.Set("matrixWorld", v)
}
func (iro *ImmediateRenderObject) MatrixWorldNeedsUpdate() bool {
	return iro.Get("matrixWorldNeedsUpdate").Bool()
}
func (iro *ImmediateRenderObject) SetMatrixWorldNeedsUpdate(v bool) {
	iro.Set("matrixWorldNeedsUpdate", v)
}
func (iro *ImmediateRenderObject) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: iro.Get("modelViewMatrix")}
}
func (iro *ImmediateRenderObject) SetModelViewMatrix(v *math.Matrix4) {
	iro.Set("modelViewMatrix", v)
}
func (iro *ImmediateRenderObject) Name() string {
	return iro.Get("name").String()
}
func (iro *ImmediateRenderObject) SetName(v string) {
	iro.Set("name", v)
}
func (iro *ImmediateRenderObject) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: iro.Get("normalMatrix")}
}
func (iro *ImmediateRenderObject) SetNormalMatrix(v *math.Matrix3) {
	iro.Set("normalMatrix", v)
}
func (iro *ImmediateRenderObject) OnAfterRender() js.Value {
	return iro.Get("onAfterRender")
}
func (iro *ImmediateRenderObject) SetOnAfterRender(v js.Value) {
	iro.Set("onAfterRender", v)
}
func (iro *ImmediateRenderObject) OnBeforeRender() js.Value {
	return iro.Get("onBeforeRender")
}
func (iro *ImmediateRenderObject) SetOnBeforeRender(v js.Value) {
	iro.Set("onBeforeRender", v)
}
func (iro *ImmediateRenderObject) Parent() core.Object3D {
	return core.Object3D(iro.Get("parent"))
}
func (iro *ImmediateRenderObject) SetParent(v core.Object3D) {
	iro.Set("parent", v)
}
func (iro *ImmediateRenderObject) Position() *math.Vector3 {
	return &math.Vector3{Value: iro.Get("position")}
}
func (iro *ImmediateRenderObject) SetPosition(v *math.Vector3) {
	iro.Set("position", v)
}
func (iro *ImmediateRenderObject) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: iro.Get("quaternion")}
}
func (iro *ImmediateRenderObject) SetQuaternion(v *math.Quaternion) {
	iro.Set("quaternion", v)
}
func (iro *ImmediateRenderObject) ReceiveShadow() bool {
	return iro.Get("receiveShadow").Bool()
}
func (iro *ImmediateRenderObject) SetReceiveShadow(v bool) {
	iro.Set("receiveShadow", v)
}
func (iro *ImmediateRenderObject) RenderOrder() int {
	return iro.Get("renderOrder").Int()
}
func (iro *ImmediateRenderObject) SetRenderOrder(v int) {
	iro.Set("renderOrder", v)
}
func (iro *ImmediateRenderObject) Rotation() *math.Euler {
	return &math.Euler{Value: iro.Get("rotation")}
}
func (iro *ImmediateRenderObject) SetRotation(v *math.Euler) {
	iro.Set("rotation", v)
}
func (iro *ImmediateRenderObject) Scale() *math.Vector3 {
	return &math.Vector3{Value: iro.Get("scale")}
}
func (iro *ImmediateRenderObject) SetScale(v *math.Vector3) {
	iro.Set("scale", v)
}
func (iro *ImmediateRenderObject) Type() string {
	return iro.Get("type").String()
}
func (iro *ImmediateRenderObject) SetType(v string) {
	iro.Set("type", v)
}
func (iro *ImmediateRenderObject) Up() *math.Vector3 {
	return &math.Vector3{Value: iro.Get("up")}
}
func (iro *ImmediateRenderObject) SetUp(v *math.Vector3) {
	iro.Set("up", v)
}
func (iro *ImmediateRenderObject) UserData() js.Value {
	return iro.Get("userData")
}
func (iro *ImmediateRenderObject) SetUserData(v js.Value) {
	iro.Set("userData", v)
}
func (iro *ImmediateRenderObject) Uuid() string {
	return iro.Get("uuid").String()
}
func (iro *ImmediateRenderObject) SetUuid(v string) {
	iro.Set("uuid", v)
}
func (iro *ImmediateRenderObject) Visible() bool {
	return iro.Get("visible").Bool()
}
func (iro *ImmediateRenderObject) SetVisible(v bool) {
	iro.Set("visible", v)
}
func (iro *ImmediateRenderObject) DefaultMatrixAutoUpdate() bool {
	return iro.Get("DefaultMatrixAutoUpdate").Bool()
}
func (iro *ImmediateRenderObject) SetDefaultMatrixAutoUpdate(v bool) {
	iro.Set("DefaultMatrixAutoUpdate", v)
}
func (iro *ImmediateRenderObject) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: iro.Get("DefaultUp")}
}
func (iro *ImmediateRenderObject) SetDefaultUp(v *math.Vector3) {
	iro.Set("DefaultUp", v)
}
func (iro *ImmediateRenderObject) Add(object []core.Object3D) this {
	return this(iro.Call("add", object))
}
func (iro *ImmediateRenderObject) AddEventListener(typ string, listener js.Value) {
	iro.Call("addEventListener", typ, listener)
}
func (iro *ImmediateRenderObject) ApplyMatrix(matrix *math.Matrix4) {
	iro.Call("applyMatrix", matrix)
}
func (iro *ImmediateRenderObject) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(iro.Call("applyQuaternion", quaternion))
}
func (iro *ImmediateRenderObject) Clone(recursive bool) this {
	return this(iro.Call("clone", recursive))
}
func (iro *ImmediateRenderObject) Copy(source *core.Object3D, recursive bool) this {
	return this(iro.Call("copy", source, recursive))
}
func (iro *ImmediateRenderObject) DispatchEvent(event js.Value) {
	iro.Call("dispatchEvent", event)
}
func (iro *ImmediateRenderObject) GetObjectById(id int) core.Object3D {
	return core.Object3D(iro.Call("getObjectById", id))
}
func (iro *ImmediateRenderObject) GetObjectByName(name string) core.Object3D {
	return core.Object3D(iro.Call("getObjectByName", name))
}
func (iro *ImmediateRenderObject) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(iro.Call("getObjectByProperty", name, value))
}
func (iro *ImmediateRenderObject) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: iro.Call("getWorldDirection", target)}
}
func (iro *ImmediateRenderObject) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: iro.Call("getWorldPosition", target)}
}
func (iro *ImmediateRenderObject) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: iro.Call("getWorldQuaternion", target)}
}
func (iro *ImmediateRenderObject) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: iro.Call("getWorldScale", target)}
}
func (iro *ImmediateRenderObject) HasEventListener(typ string, listener js.Value) bool {
	return iro.Call("hasEventListener", typ, listener).Bool()
}
func (iro *ImmediateRenderObject) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: iro.Call("localToWorld", vector)}
}
func (iro *ImmediateRenderObject) LookAt(vector math.Vector3, y int, z int) {
	iro.Call("lookAt", vector, y, z)
}
func (iro *ImmediateRenderObject) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	iro.Call("raycast", raycaster, intersects)
}
func (iro *ImmediateRenderObject) Remove(object []core.Object3D) this {
	return this(iro.Call("remove", object))
}
func (iro *ImmediateRenderObject) RemoveEventListener(typ string, listener js.Value) {
	iro.Call("removeEventListener", typ, listener)
}
func (iro *ImmediateRenderObject) Render(renderCallback js.Value) {
	iro.Call("render", renderCallback)
}
func (iro *ImmediateRenderObject) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(iro.Call("rotateOnAxis", axis, angle))
}
func (iro *ImmediateRenderObject) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(iro.Call("rotateOnWorldAxis", axis, angle))
}
func (iro *ImmediateRenderObject) RotateX(angle int) this {
	return this(iro.Call("rotateX", angle))
}
func (iro *ImmediateRenderObject) RotateY(angle int) this {
	return this(iro.Call("rotateY", angle))
}
func (iro *ImmediateRenderObject) RotateZ(angle int) this {
	return this(iro.Call("rotateZ", angle))
}
func (iro *ImmediateRenderObject) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	iro.Call("setRotationFromAxisAngle", axis, angle)
}
func (iro *ImmediateRenderObject) SetRotationFromEuler(euler *math.Euler) {
	iro.Call("setRotationFromEuler", euler)
}
func (iro *ImmediateRenderObject) SetRotationFromMatrix(m *math.Matrix4) {
	iro.Call("setRotationFromMatrix", m)
}
func (iro *ImmediateRenderObject) SetRotationFromQuaternion(q *math.Quaternion) {
	iro.Call("setRotationFromQuaternion", q)
}
func (iro *ImmediateRenderObject) ToJSON(meta js.Value) js.Value {
	return iro.Call("toJSON", meta)
}
func (iro *ImmediateRenderObject) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(iro.Call("translateOnAxis", axis, distance))
}
func (iro *ImmediateRenderObject) TranslateX(distance int) this {
	return this(iro.Call("translateX", distance))
}
func (iro *ImmediateRenderObject) TranslateY(distance int) this {
	return this(iro.Call("translateY", distance))
}
func (iro *ImmediateRenderObject) TranslateZ(distance int) this {
	return this(iro.Call("translateZ", distance))
}
func (iro *ImmediateRenderObject) Traverse(callback js.Value) {
	iro.Call("traverse", callback)
}
func (iro *ImmediateRenderObject) TraverseAncestors(callback js.Value) {
	iro.Call("traverseAncestors", callback)
}
func (iro *ImmediateRenderObject) TraverseVisible(callback js.Value) {
	iro.Call("traverseVisible", callback)
}
func (iro *ImmediateRenderObject) UpdateMatrix() {
	iro.Call("updateMatrix")
}
func (iro *ImmediateRenderObject) UpdateMatrixWorld(force bool) {
	iro.Call("updateMatrixWorld", force)
}
func (iro *ImmediateRenderObject) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	iro.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (iro *ImmediateRenderObject) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: iro.Call("worldToLocal", vector)}
}

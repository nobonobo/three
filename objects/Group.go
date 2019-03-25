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

type Group struct {
	js.Value
}

func NewGroup() *Group {
	return &Group{Value: get("Group").New()}
}
func (g *Group) CastShadow() bool {
	return g.Get("castShadow").Bool()
}
func (g *Group) SetCastShadow(v bool) {
	g.Set("castShadow", v)
}
func (g *Group) Children() []core.Object3D {
	return []core.Object3D(g.Get("children"))
}
func (g *Group) SetChildren(v []core.Object3D) {
	g.Set("children", v)
}
func (g *Group) FrustumCulled() bool {
	return g.Get("frustumCulled").Bool()
}
func (g *Group) SetFrustumCulled(v bool) {
	g.Set("frustumCulled", v)
}
func (g *Group) Id() int {
	return g.Get("id").Int()
}
func (g *Group) SetId(v int) {
	g.Set("id", v)
}
func (g *Group) IsGroup() true {
	return true(g.Get("isGroup"))
}
func (g *Group) SetIsGroup(v true) {
	g.Set("isGroup", v)
}
func (g *Group) IsObject3D() true {
	return true(g.Get("isObject3D"))
}
func (g *Group) SetIsObject3D(v true) {
	g.Set("isObject3D", v)
}
func (g *Group) Layers() *core.Layers {
	return &core.Layers{Value: g.Get("layers")}
}
func (g *Group) SetLayers(v *core.Layers) {
	g.Set("layers", v)
}
func (g *Group) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: g.Get("matrix")}
}
func (g *Group) SetMatrix(v *math.Matrix4) {
	g.Set("matrix", v)
}
func (g *Group) MatrixAutoUpdate() bool {
	return g.Get("matrixAutoUpdate").Bool()
}
func (g *Group) SetMatrixAutoUpdate(v bool) {
	g.Set("matrixAutoUpdate", v)
}
func (g *Group) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: g.Get("matrixWorld")}
}
func (g *Group) SetMatrixWorld(v *math.Matrix4) {
	g.Set("matrixWorld", v)
}
func (g *Group) MatrixWorldNeedsUpdate() bool {
	return g.Get("matrixWorldNeedsUpdate").Bool()
}
func (g *Group) SetMatrixWorldNeedsUpdate(v bool) {
	g.Set("matrixWorldNeedsUpdate", v)
}
func (g *Group) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: g.Get("modelViewMatrix")}
}
func (g *Group) SetModelViewMatrix(v *math.Matrix4) {
	g.Set("modelViewMatrix", v)
}
func (g *Group) Name() string {
	return g.Get("name").String()
}
func (g *Group) SetName(v string) {
	g.Set("name", v)
}
func (g *Group) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: g.Get("normalMatrix")}
}
func (g *Group) SetNormalMatrix(v *math.Matrix3) {
	g.Set("normalMatrix", v)
}
func (g *Group) OnAfterRender() js.Value {
	return g.Get("onAfterRender")
}
func (g *Group) SetOnAfterRender(v js.Value) {
	g.Set("onAfterRender", v)
}
func (g *Group) OnBeforeRender() js.Value {
	return g.Get("onBeforeRender")
}
func (g *Group) SetOnBeforeRender(v js.Value) {
	g.Set("onBeforeRender", v)
}
func (g *Group) Parent() core.Object3D {
	return core.Object3D(g.Get("parent"))
}
func (g *Group) SetParent(v core.Object3D) {
	g.Set("parent", v)
}
func (g *Group) Position() *math.Vector3 {
	return &math.Vector3{Value: g.Get("position")}
}
func (g *Group) SetPosition(v *math.Vector3) {
	g.Set("position", v)
}
func (g *Group) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: g.Get("quaternion")}
}
func (g *Group) SetQuaternion(v *math.Quaternion) {
	g.Set("quaternion", v)
}
func (g *Group) ReceiveShadow() bool {
	return g.Get("receiveShadow").Bool()
}
func (g *Group) SetReceiveShadow(v bool) {
	g.Set("receiveShadow", v)
}
func (g *Group) RenderOrder() int {
	return g.Get("renderOrder").Int()
}
func (g *Group) SetRenderOrder(v int) {
	g.Set("renderOrder", v)
}
func (g *Group) Rotation() *math.Euler {
	return &math.Euler{Value: g.Get("rotation")}
}
func (g *Group) SetRotation(v *math.Euler) {
	g.Set("rotation", v)
}
func (g *Group) Scale() *math.Vector3 {
	return &math.Vector3{Value: g.Get("scale")}
}
func (g *Group) SetScale(v *math.Vector3) {
	g.Set("scale", v)
}
func (g *Group) Type() string {
	return g.Get("type").String()
}
func (g *Group) SetType(v string) {
	g.Set("type", v)
}
func (g *Group) Up() *math.Vector3 {
	return &math.Vector3{Value: g.Get("up")}
}
func (g *Group) SetUp(v *math.Vector3) {
	g.Set("up", v)
}
func (g *Group) UserData() js.Value {
	return g.Get("userData")
}
func (g *Group) SetUserData(v js.Value) {
	g.Set("userData", v)
}
func (g *Group) Uuid() string {
	return g.Get("uuid").String()
}
func (g *Group) SetUuid(v string) {
	g.Set("uuid", v)
}
func (g *Group) Visible() bool {
	return g.Get("visible").Bool()
}
func (g *Group) SetVisible(v bool) {
	g.Set("visible", v)
}
func (g *Group) DefaultMatrixAutoUpdate() bool {
	return g.Get("DefaultMatrixAutoUpdate").Bool()
}
func (g *Group) SetDefaultMatrixAutoUpdate(v bool) {
	g.Set("DefaultMatrixAutoUpdate", v)
}
func (g *Group) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: g.Get("DefaultUp")}
}
func (g *Group) SetDefaultUp(v *math.Vector3) {
	g.Set("DefaultUp", v)
}
func (g *Group) Add(object []core.Object3D) this {
	return this(g.Call("add", object))
}
func (g *Group) AddEventListener(typ string, listener js.Value) {
	g.Call("addEventListener", typ, listener)
}
func (g *Group) ApplyMatrix(matrix *math.Matrix4) {
	g.Call("applyMatrix", matrix)
}
func (g *Group) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(g.Call("applyQuaternion", quaternion))
}
func (g *Group) Clone(recursive bool) this {
	return this(g.Call("clone", recursive))
}
func (g *Group) Copy(source *core.Object3D, recursive bool) this {
	return this(g.Call("copy", source, recursive))
}
func (g *Group) DispatchEvent(event js.Value) {
	g.Call("dispatchEvent", event)
}
func (g *Group) GetObjectById(id int) core.Object3D {
	return core.Object3D(g.Call("getObjectById", id))
}
func (g *Group) GetObjectByName(name string) core.Object3D {
	return core.Object3D(g.Call("getObjectByName", name))
}
func (g *Group) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(g.Call("getObjectByProperty", name, value))
}
func (g *Group) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: g.Call("getWorldDirection", target)}
}
func (g *Group) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: g.Call("getWorldPosition", target)}
}
func (g *Group) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: g.Call("getWorldQuaternion", target)}
}
func (g *Group) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: g.Call("getWorldScale", target)}
}
func (g *Group) HasEventListener(typ string, listener js.Value) bool {
	return g.Call("hasEventListener", typ, listener).Bool()
}
func (g *Group) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: g.Call("localToWorld", vector)}
}
func (g *Group) LookAt(vector math.Vector3, y int, z int) {
	g.Call("lookAt", vector, y, z)
}
func (g *Group) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	g.Call("raycast", raycaster, intersects)
}
func (g *Group) Remove(object []core.Object3D) this {
	return this(g.Call("remove", object))
}
func (g *Group) RemoveEventListener(typ string, listener js.Value) {
	g.Call("removeEventListener", typ, listener)
}
func (g *Group) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(g.Call("rotateOnAxis", axis, angle))
}
func (g *Group) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(g.Call("rotateOnWorldAxis", axis, angle))
}
func (g *Group) RotateX(angle int) this {
	return this(g.Call("rotateX", angle))
}
func (g *Group) RotateY(angle int) this {
	return this(g.Call("rotateY", angle))
}
func (g *Group) RotateZ(angle int) this {
	return this(g.Call("rotateZ", angle))
}
func (g *Group) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	g.Call("setRotationFromAxisAngle", axis, angle)
}
func (g *Group) SetRotationFromEuler(euler *math.Euler) {
	g.Call("setRotationFromEuler", euler)
}
func (g *Group) SetRotationFromMatrix(m *math.Matrix4) {
	g.Call("setRotationFromMatrix", m)
}
func (g *Group) SetRotationFromQuaternion(q *math.Quaternion) {
	g.Call("setRotationFromQuaternion", q)
}
func (g *Group) ToJSON(meta js.Value) js.Value {
	return g.Call("toJSON", meta)
}
func (g *Group) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(g.Call("translateOnAxis", axis, distance))
}
func (g *Group) TranslateX(distance int) this {
	return this(g.Call("translateX", distance))
}
func (g *Group) TranslateY(distance int) this {
	return this(g.Call("translateY", distance))
}
func (g *Group) TranslateZ(distance int) this {
	return this(g.Call("translateZ", distance))
}
func (g *Group) Traverse(callback js.Value) {
	g.Call("traverse", callback)
}
func (g *Group) TraverseAncestors(callback js.Value) {
	g.Call("traverseAncestors", callback)
}
func (g *Group) TraverseVisible(callback js.Value) {
	g.Call("traverseVisible", callback)
}
func (g *Group) UpdateMatrix() {
	g.Call("updateMatrix")
}
func (g *Group) UpdateMatrixWorld(force bool) {
	g.Call("updateMatrixWorld", force)
}
func (g *Group) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	g.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (g *Group) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: g.Call("worldToLocal", vector)}
}

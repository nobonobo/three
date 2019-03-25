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

type Points struct {
	js.Value
}

func NewPoints(geometry core.Geometry, material materials.Material) *Points {
	return &Points{Value: get("Points").New(geometry, material)}
}
func (p *Points) CastShadow() bool {
	return p.Get("castShadow").Bool()
}
func (p *Points) SetCastShadow(v bool) {
	p.Set("castShadow", v)
}
func (p *Points) Children() []core.Object3D {
	return []core.Object3D(p.Get("children"))
}
func (p *Points) SetChildren(v []core.Object3D) {
	p.Set("children", v)
}
func (p *Points) FrustumCulled() bool {
	return p.Get("frustumCulled").Bool()
}
func (p *Points) SetFrustumCulled(v bool) {
	p.Set("frustumCulled", v)
}
func (p *Points) Geometry() core.Geometry {
	return core.Geometry(p.Get("geometry"))
}
func (p *Points) SetGeometry(v core.Geometry) {
	p.Set("geometry", v)
}
func (p *Points) Id() int {
	return p.Get("id").Int()
}
func (p *Points) SetId(v int) {
	p.Set("id", v)
}
func (p *Points) IsObject3D() true {
	return true(p.Get("isObject3D"))
}
func (p *Points) SetIsObject3D(v true) {
	p.Set("isObject3D", v)
}
func (p *Points) IsPoints() true {
	return true(p.Get("isPoints"))
}
func (p *Points) SetIsPoints(v true) {
	p.Set("isPoints", v)
}
func (p *Points) Layers() *core.Layers {
	return &core.Layers{Value: p.Get("layers")}
}
func (p *Points) SetLayers(v *core.Layers) {
	p.Set("layers", v)
}
func (p *Points) Material() materials.Material {
	return materials.Material(p.Get("material"))
}
func (p *Points) SetMaterial(v materials.Material) {
	p.Set("material", v)
}
func (p *Points) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: p.Get("matrix")}
}
func (p *Points) SetMatrix(v *math.Matrix4) {
	p.Set("matrix", v)
}
func (p *Points) MatrixAutoUpdate() bool {
	return p.Get("matrixAutoUpdate").Bool()
}
func (p *Points) SetMatrixAutoUpdate(v bool) {
	p.Set("matrixAutoUpdate", v)
}
func (p *Points) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: p.Get("matrixWorld")}
}
func (p *Points) SetMatrixWorld(v *math.Matrix4) {
	p.Set("matrixWorld", v)
}
func (p *Points) MatrixWorldNeedsUpdate() bool {
	return p.Get("matrixWorldNeedsUpdate").Bool()
}
func (p *Points) SetMatrixWorldNeedsUpdate(v bool) {
	p.Set("matrixWorldNeedsUpdate", v)
}
func (p *Points) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: p.Get("modelViewMatrix")}
}
func (p *Points) SetModelViewMatrix(v *math.Matrix4) {
	p.Set("modelViewMatrix", v)
}
func (p *Points) Name() string {
	return p.Get("name").String()
}
func (p *Points) SetName(v string) {
	p.Set("name", v)
}
func (p *Points) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: p.Get("normalMatrix")}
}
func (p *Points) SetNormalMatrix(v *math.Matrix3) {
	p.Set("normalMatrix", v)
}
func (p *Points) OnAfterRender() js.Value {
	return p.Get("onAfterRender")
}
func (p *Points) SetOnAfterRender(v js.Value) {
	p.Set("onAfterRender", v)
}
func (p *Points) OnBeforeRender() js.Value {
	return p.Get("onBeforeRender")
}
func (p *Points) SetOnBeforeRender(v js.Value) {
	p.Set("onBeforeRender", v)
}
func (p *Points) Parent() core.Object3D {
	return core.Object3D(p.Get("parent"))
}
func (p *Points) SetParent(v core.Object3D) {
	p.Set("parent", v)
}
func (p *Points) Position() *math.Vector3 {
	return &math.Vector3{Value: p.Get("position")}
}
func (p *Points) SetPosition(v *math.Vector3) {
	p.Set("position", v)
}
func (p *Points) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: p.Get("quaternion")}
}
func (p *Points) SetQuaternion(v *math.Quaternion) {
	p.Set("quaternion", v)
}
func (p *Points) ReceiveShadow() bool {
	return p.Get("receiveShadow").Bool()
}
func (p *Points) SetReceiveShadow(v bool) {
	p.Set("receiveShadow", v)
}
func (p *Points) RenderOrder() int {
	return p.Get("renderOrder").Int()
}
func (p *Points) SetRenderOrder(v int) {
	p.Set("renderOrder", v)
}
func (p *Points) Rotation() *math.Euler {
	return &math.Euler{Value: p.Get("rotation")}
}
func (p *Points) SetRotation(v *math.Euler) {
	p.Set("rotation", v)
}
func (p *Points) Scale() *math.Vector3 {
	return &math.Vector3{Value: p.Get("scale")}
}
func (p *Points) SetScale(v *math.Vector3) {
	p.Set("scale", v)
}
func (p *Points) Type() string {
	return p.Get("type").String()
}
func (p *Points) SetType(v string) {
	p.Set("type", v)
}
func (p *Points) Up() *math.Vector3 {
	return &math.Vector3{Value: p.Get("up")}
}
func (p *Points) SetUp(v *math.Vector3) {
	p.Set("up", v)
}
func (p *Points) UserData() js.Value {
	return p.Get("userData")
}
func (p *Points) SetUserData(v js.Value) {
	p.Set("userData", v)
}
func (p *Points) Uuid() string {
	return p.Get("uuid").String()
}
func (p *Points) SetUuid(v string) {
	p.Set("uuid", v)
}
func (p *Points) Visible() bool {
	return p.Get("visible").Bool()
}
func (p *Points) SetVisible(v bool) {
	p.Set("visible", v)
}
func (p *Points) DefaultMatrixAutoUpdate() bool {
	return p.Get("DefaultMatrixAutoUpdate").Bool()
}
func (p *Points) SetDefaultMatrixAutoUpdate(v bool) {
	p.Set("DefaultMatrixAutoUpdate", v)
}
func (p *Points) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: p.Get("DefaultUp")}
}
func (p *Points) SetDefaultUp(v *math.Vector3) {
	p.Set("DefaultUp", v)
}
func (p *Points) Add(object []core.Object3D) this {
	return this(p.Call("add", object))
}
func (p *Points) AddEventListener(typ string, listener js.Value) {
	p.Call("addEventListener", typ, listener)
}
func (p *Points) ApplyMatrix(matrix *math.Matrix4) {
	p.Call("applyMatrix", matrix)
}
func (p *Points) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(p.Call("applyQuaternion", quaternion))
}
func (p *Points) Clone(recursive bool) this {
	return this(p.Call("clone", recursive))
}
func (p *Points) Copy(source *core.Object3D, recursive bool) this {
	return this(p.Call("copy", source, recursive))
}
func (p *Points) DispatchEvent(event js.Value) {
	p.Call("dispatchEvent", event)
}
func (p *Points) GetObjectById(id int) core.Object3D {
	return core.Object3D(p.Call("getObjectById", id))
}
func (p *Points) GetObjectByName(name string) core.Object3D {
	return core.Object3D(p.Call("getObjectByName", name))
}
func (p *Points) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(p.Call("getObjectByProperty", name, value))
}
func (p *Points) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: p.Call("getWorldDirection", target)}
}
func (p *Points) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: p.Call("getWorldPosition", target)}
}
func (p *Points) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: p.Call("getWorldQuaternion", target)}
}
func (p *Points) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: p.Call("getWorldScale", target)}
}
func (p *Points) HasEventListener(typ string, listener js.Value) bool {
	return p.Call("hasEventListener", typ, listener).Bool()
}
func (p *Points) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: p.Call("localToWorld", vector)}
}
func (p *Points) LookAt(vector math.Vector3, y int, z int) {
	p.Call("lookAt", vector, y, z)
}
func (p *Points) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	p.Call("raycast", raycaster, intersects)
}
func (p *Points) Remove(object []core.Object3D) this {
	return this(p.Call("remove", object))
}
func (p *Points) RemoveEventListener(typ string, listener js.Value) {
	p.Call("removeEventListener", typ, listener)
}
func (p *Points) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(p.Call("rotateOnAxis", axis, angle))
}
func (p *Points) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(p.Call("rotateOnWorldAxis", axis, angle))
}
func (p *Points) RotateX(angle int) this {
	return this(p.Call("rotateX", angle))
}
func (p *Points) RotateY(angle int) this {
	return this(p.Call("rotateY", angle))
}
func (p *Points) RotateZ(angle int) this {
	return this(p.Call("rotateZ", angle))
}
func (p *Points) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	p.Call("setRotationFromAxisAngle", axis, angle)
}
func (p *Points) SetRotationFromEuler(euler *math.Euler) {
	p.Call("setRotationFromEuler", euler)
}
func (p *Points) SetRotationFromMatrix(m *math.Matrix4) {
	p.Call("setRotationFromMatrix", m)
}
func (p *Points) SetRotationFromQuaternion(q *math.Quaternion) {
	p.Call("setRotationFromQuaternion", q)
}
func (p *Points) ToJSON(meta js.Value) js.Value {
	return p.Call("toJSON", meta)
}
func (p *Points) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(p.Call("translateOnAxis", axis, distance))
}
func (p *Points) TranslateX(distance int) this {
	return this(p.Call("translateX", distance))
}
func (p *Points) TranslateY(distance int) this {
	return this(p.Call("translateY", distance))
}
func (p *Points) TranslateZ(distance int) this {
	return this(p.Call("translateZ", distance))
}
func (p *Points) Traverse(callback js.Value) {
	p.Call("traverse", callback)
}
func (p *Points) TraverseAncestors(callback js.Value) {
	p.Call("traverseAncestors", callback)
}
func (p *Points) TraverseVisible(callback js.Value) {
	p.Call("traverseVisible", callback)
}
func (p *Points) UpdateMatrix() {
	p.Call("updateMatrix")
}
func (p *Points) UpdateMatrixWorld(force bool) {
	p.Call("updateMatrixWorld", force)
}
func (p *Points) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	p.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (p *Points) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: p.Call("worldToLocal", vector)}
}

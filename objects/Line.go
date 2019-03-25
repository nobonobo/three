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

type Line struct {
	js.Value
}

func NewLine(geometry core.Geometry, material materials.Material, mode int) *Line {
	return &Line{Value: get("Line").New(geometry, material, mode)}
}
func (l *Line) CastShadow() bool {
	return l.Get("castShadow").Bool()
}
func (l *Line) SetCastShadow(v bool) {
	l.Set("castShadow", v)
}
func (l *Line) Children() []core.Object3D {
	return []core.Object3D(l.Get("children"))
}
func (l *Line) SetChildren(v []core.Object3D) {
	l.Set("children", v)
}
func (l *Line) FrustumCulled() bool {
	return l.Get("frustumCulled").Bool()
}
func (l *Line) SetFrustumCulled(v bool) {
	l.Set("frustumCulled", v)
}
func (l *Line) Geometry() core.Geometry {
	return core.Geometry(l.Get("geometry"))
}
func (l *Line) SetGeometry(v core.Geometry) {
	l.Set("geometry", v)
}
func (l *Line) Id() int {
	return l.Get("id").Int()
}
func (l *Line) SetId(v int) {
	l.Set("id", v)
}
func (l *Line) IsLine() true {
	return true(l.Get("isLine"))
}
func (l *Line) SetIsLine(v true) {
	l.Set("isLine", v)
}
func (l *Line) IsObject3D() true {
	return true(l.Get("isObject3D"))
}
func (l *Line) SetIsObject3D(v true) {
	l.Set("isObject3D", v)
}
func (l *Line) Layers() *core.Layers {
	return &core.Layers{Value: l.Get("layers")}
}
func (l *Line) SetLayers(v *core.Layers) {
	l.Set("layers", v)
}
func (l *Line) Material() materials.Material {
	return materials.Material(l.Get("material"))
}
func (l *Line) SetMaterial(v materials.Material) {
	l.Set("material", v)
}
func (l *Line) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: l.Get("matrix")}
}
func (l *Line) SetMatrix(v *math.Matrix4) {
	l.Set("matrix", v)
}
func (l *Line) MatrixAutoUpdate() bool {
	return l.Get("matrixAutoUpdate").Bool()
}
func (l *Line) SetMatrixAutoUpdate(v bool) {
	l.Set("matrixAutoUpdate", v)
}
func (l *Line) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: l.Get("matrixWorld")}
}
func (l *Line) SetMatrixWorld(v *math.Matrix4) {
	l.Set("matrixWorld", v)
}
func (l *Line) MatrixWorldNeedsUpdate() bool {
	return l.Get("matrixWorldNeedsUpdate").Bool()
}
func (l *Line) SetMatrixWorldNeedsUpdate(v bool) {
	l.Set("matrixWorldNeedsUpdate", v)
}
func (l *Line) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: l.Get("modelViewMatrix")}
}
func (l *Line) SetModelViewMatrix(v *math.Matrix4) {
	l.Set("modelViewMatrix", v)
}
func (l *Line) Name() string {
	return l.Get("name").String()
}
func (l *Line) SetName(v string) {
	l.Set("name", v)
}
func (l *Line) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: l.Get("normalMatrix")}
}
func (l *Line) SetNormalMatrix(v *math.Matrix3) {
	l.Set("normalMatrix", v)
}
func (l *Line) OnAfterRender() js.Value {
	return l.Get("onAfterRender")
}
func (l *Line) SetOnAfterRender(v js.Value) {
	l.Set("onAfterRender", v)
}
func (l *Line) OnBeforeRender() js.Value {
	return l.Get("onBeforeRender")
}
func (l *Line) SetOnBeforeRender(v js.Value) {
	l.Set("onBeforeRender", v)
}
func (l *Line) Parent() core.Object3D {
	return core.Object3D(l.Get("parent"))
}
func (l *Line) SetParent(v core.Object3D) {
	l.Set("parent", v)
}
func (l *Line) Position() *math.Vector3 {
	return &math.Vector3{Value: l.Get("position")}
}
func (l *Line) SetPosition(v *math.Vector3) {
	l.Set("position", v)
}
func (l *Line) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: l.Get("quaternion")}
}
func (l *Line) SetQuaternion(v *math.Quaternion) {
	l.Set("quaternion", v)
}
func (l *Line) ReceiveShadow() bool {
	return l.Get("receiveShadow").Bool()
}
func (l *Line) SetReceiveShadow(v bool) {
	l.Set("receiveShadow", v)
}
func (l *Line) RenderOrder() int {
	return l.Get("renderOrder").Int()
}
func (l *Line) SetRenderOrder(v int) {
	l.Set("renderOrder", v)
}
func (l *Line) Rotation() *math.Euler {
	return &math.Euler{Value: l.Get("rotation")}
}
func (l *Line) SetRotation(v *math.Euler) {
	l.Set("rotation", v)
}
func (l *Line) Scale() *math.Vector3 {
	return &math.Vector3{Value: l.Get("scale")}
}
func (l *Line) SetScale(v *math.Vector3) {
	l.Set("scale", v)
}
func (l *Line) Type() string {
	return l.Get("type").String()
}
func (l *Line) SetType(v string) {
	l.Set("type", v)
}
func (l *Line) Up() *math.Vector3 {
	return &math.Vector3{Value: l.Get("up")}
}
func (l *Line) SetUp(v *math.Vector3) {
	l.Set("up", v)
}
func (l *Line) UserData() js.Value {
	return l.Get("userData")
}
func (l *Line) SetUserData(v js.Value) {
	l.Set("userData", v)
}
func (l *Line) Uuid() string {
	return l.Get("uuid").String()
}
func (l *Line) SetUuid(v string) {
	l.Set("uuid", v)
}
func (l *Line) Visible() bool {
	return l.Get("visible").Bool()
}
func (l *Line) SetVisible(v bool) {
	l.Set("visible", v)
}
func (l *Line) DefaultMatrixAutoUpdate() bool {
	return l.Get("DefaultMatrixAutoUpdate").Bool()
}
func (l *Line) SetDefaultMatrixAutoUpdate(v bool) {
	l.Set("DefaultMatrixAutoUpdate", v)
}
func (l *Line) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: l.Get("DefaultUp")}
}
func (l *Line) SetDefaultUp(v *math.Vector3) {
	l.Set("DefaultUp", v)
}
func (l *Line) Add(object []core.Object3D) this {
	return this(l.Call("add", object))
}
func (l *Line) AddEventListener(typ string, listener js.Value) {
	l.Call("addEventListener", typ, listener)
}
func (l *Line) ApplyMatrix(matrix *math.Matrix4) {
	l.Call("applyMatrix", matrix)
}
func (l *Line) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(l.Call("applyQuaternion", quaternion))
}
func (l *Line) Clone(recursive bool) this {
	return this(l.Call("clone", recursive))
}
func (l *Line) ComputeLineDistances() this {
	return this(l.Call("computeLineDistances"))
}
func (l *Line) Copy(source *core.Object3D, recursive bool) this {
	return this(l.Call("copy", source, recursive))
}
func (l *Line) DispatchEvent(event js.Value) {
	l.Call("dispatchEvent", event)
}
func (l *Line) GetObjectById(id int) core.Object3D {
	return core.Object3D(l.Call("getObjectById", id))
}
func (l *Line) GetObjectByName(name string) core.Object3D {
	return core.Object3D(l.Call("getObjectByName", name))
}
func (l *Line) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(l.Call("getObjectByProperty", name, value))
}
func (l *Line) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("getWorldDirection", target)}
}
func (l *Line) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("getWorldPosition", target)}
}
func (l *Line) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: l.Call("getWorldQuaternion", target)}
}
func (l *Line) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("getWorldScale", target)}
}
func (l *Line) HasEventListener(typ string, listener js.Value) bool {
	return l.Call("hasEventListener", typ, listener).Bool()
}
func (l *Line) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("localToWorld", vector)}
}
func (l *Line) LookAt(vector math.Vector3, y int, z int) {
	l.Call("lookAt", vector, y, z)
}
func (l *Line) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	l.Call("raycast", raycaster, intersects)
}
func (l *Line) Remove(object []core.Object3D) this {
	return this(l.Call("remove", object))
}
func (l *Line) RemoveEventListener(typ string, listener js.Value) {
	l.Call("removeEventListener", typ, listener)
}
func (l *Line) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(l.Call("rotateOnAxis", axis, angle))
}
func (l *Line) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(l.Call("rotateOnWorldAxis", axis, angle))
}
func (l *Line) RotateX(angle int) this {
	return this(l.Call("rotateX", angle))
}
func (l *Line) RotateY(angle int) this {
	return this(l.Call("rotateY", angle))
}
func (l *Line) RotateZ(angle int) this {
	return this(l.Call("rotateZ", angle))
}
func (l *Line) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	l.Call("setRotationFromAxisAngle", axis, angle)
}
func (l *Line) SetRotationFromEuler(euler *math.Euler) {
	l.Call("setRotationFromEuler", euler)
}
func (l *Line) SetRotationFromMatrix(m *math.Matrix4) {
	l.Call("setRotationFromMatrix", m)
}
func (l *Line) SetRotationFromQuaternion(q *math.Quaternion) {
	l.Call("setRotationFromQuaternion", q)
}
func (l *Line) ToJSON(meta js.Value) js.Value {
	return l.Call("toJSON", meta)
}
func (l *Line) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(l.Call("translateOnAxis", axis, distance))
}
func (l *Line) TranslateX(distance int) this {
	return this(l.Call("translateX", distance))
}
func (l *Line) TranslateY(distance int) this {
	return this(l.Call("translateY", distance))
}
func (l *Line) TranslateZ(distance int) this {
	return this(l.Call("translateZ", distance))
}
func (l *Line) Traverse(callback js.Value) {
	l.Call("traverse", callback)
}
func (l *Line) TraverseAncestors(callback js.Value) {
	l.Call("traverseAncestors", callback)
}
func (l *Line) TraverseVisible(callback js.Value) {
	l.Call("traverseVisible", callback)
}
func (l *Line) UpdateMatrix() {
	l.Call("updateMatrix")
}
func (l *Line) UpdateMatrixWorld(force bool) {
	l.Call("updateMatrixWorld", force)
}
func (l *Line) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	l.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (l *Line) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("worldToLocal", vector)}
}

package helpers

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/lights"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type DirectionalLightHelper struct {
	js.Value
}

func NewDirectionalLightHelper(light *lights.DirectionalLight, size int, color math.Color) *DirectionalLightHelper {
	return &DirectionalLightHelper{Value: get("DirectionalLightHelper").New(light, size, color)}
}
func (dlh *DirectionalLightHelper) CastShadow() bool {
	return dlh.Get("castShadow").Bool()
}
func (dlh *DirectionalLightHelper) SetCastShadow(v bool) {
	dlh.Set("castShadow", v)
}
func (dlh *DirectionalLightHelper) Children() []core.Object3D {
	return []core.Object3D(dlh.Get("children"))
}
func (dlh *DirectionalLightHelper) SetChildren(v []core.Object3D) {
	dlh.Set("children", v)
}
func (dlh *DirectionalLightHelper) Color() math.Color {
	return math.Color(dlh.Get("color"))
}
func (dlh *DirectionalLightHelper) SetColor(v math.Color) {
	dlh.Set("color", v)
}
func (dlh *DirectionalLightHelper) FrustumCulled() bool {
	return dlh.Get("frustumCulled").Bool()
}
func (dlh *DirectionalLightHelper) SetFrustumCulled(v bool) {
	dlh.Set("frustumCulled", v)
}
func (dlh *DirectionalLightHelper) Id() int {
	return dlh.Get("id").Int()
}
func (dlh *DirectionalLightHelper) SetId(v int) {
	dlh.Set("id", v)
}
func (dlh *DirectionalLightHelper) IsObject3D() true {
	return true(dlh.Get("isObject3D"))
}
func (dlh *DirectionalLightHelper) SetIsObject3D(v true) {
	dlh.Set("isObject3D", v)
}
func (dlh *DirectionalLightHelper) Layers() *core.Layers {
	return &core.Layers{Value: dlh.Get("layers")}
}
func (dlh *DirectionalLightHelper) SetLayers(v *core.Layers) {
	dlh.Set("layers", v)
}
func (dlh *DirectionalLightHelper) Light() *lights.DirectionalLight {
	return &lights.DirectionalLight{Value: dlh.Get("light")}
}
func (dlh *DirectionalLightHelper) SetLight(v *lights.DirectionalLight) {
	dlh.Set("light", v)
}
func (dlh *DirectionalLightHelper) LightPlane() *objects.Line {
	return &objects.Line{Value: dlh.Get("lightPlane")}
}
func (dlh *DirectionalLightHelper) SetLightPlane(v *objects.Line) {
	dlh.Set("lightPlane", v)
}
func (dlh *DirectionalLightHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: dlh.Get("matrix")}
}
func (dlh *DirectionalLightHelper) SetMatrix(v *math.Matrix4) {
	dlh.Set("matrix", v)
}
func (dlh *DirectionalLightHelper) MatrixAutoUpdate() bool {
	return dlh.Get("matrixAutoUpdate").Bool()
}
func (dlh *DirectionalLightHelper) SetMatrixAutoUpdate(v bool) {
	dlh.Set("matrixAutoUpdate", v)
}
func (dlh *DirectionalLightHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: dlh.Get("matrixWorld")}
}
func (dlh *DirectionalLightHelper) SetMatrixWorld(v *math.Matrix4) {
	dlh.Set("matrixWorld", v)
}
func (dlh *DirectionalLightHelper) MatrixWorldNeedsUpdate() bool {
	return dlh.Get("matrixWorldNeedsUpdate").Bool()
}
func (dlh *DirectionalLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	dlh.Set("matrixWorldNeedsUpdate", v)
}
func (dlh *DirectionalLightHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: dlh.Get("modelViewMatrix")}
}
func (dlh *DirectionalLightHelper) SetModelViewMatrix(v *math.Matrix4) {
	dlh.Set("modelViewMatrix", v)
}
func (dlh *DirectionalLightHelper) Name() string {
	return dlh.Get("name").String()
}
func (dlh *DirectionalLightHelper) SetName(v string) {
	dlh.Set("name", v)
}
func (dlh *DirectionalLightHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: dlh.Get("normalMatrix")}
}
func (dlh *DirectionalLightHelper) SetNormalMatrix(v *math.Matrix3) {
	dlh.Set("normalMatrix", v)
}
func (dlh *DirectionalLightHelper) OnAfterRender() js.Value {
	return dlh.Get("onAfterRender")
}
func (dlh *DirectionalLightHelper) SetOnAfterRender(v js.Value) {
	dlh.Set("onAfterRender", v)
}
func (dlh *DirectionalLightHelper) OnBeforeRender() js.Value {
	return dlh.Get("onBeforeRender")
}
func (dlh *DirectionalLightHelper) SetOnBeforeRender(v js.Value) {
	dlh.Set("onBeforeRender", v)
}
func (dlh *DirectionalLightHelper) Parent() core.Object3D {
	return core.Object3D(dlh.Get("parent"))
}
func (dlh *DirectionalLightHelper) SetParent(v core.Object3D) {
	dlh.Set("parent", v)
}
func (dlh *DirectionalLightHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: dlh.Get("position")}
}
func (dlh *DirectionalLightHelper) SetPosition(v *math.Vector3) {
	dlh.Set("position", v)
}
func (dlh *DirectionalLightHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: dlh.Get("quaternion")}
}
func (dlh *DirectionalLightHelper) SetQuaternion(v *math.Quaternion) {
	dlh.Set("quaternion", v)
}
func (dlh *DirectionalLightHelper) ReceiveShadow() bool {
	return dlh.Get("receiveShadow").Bool()
}
func (dlh *DirectionalLightHelper) SetReceiveShadow(v bool) {
	dlh.Set("receiveShadow", v)
}
func (dlh *DirectionalLightHelper) RenderOrder() int {
	return dlh.Get("renderOrder").Int()
}
func (dlh *DirectionalLightHelper) SetRenderOrder(v int) {
	dlh.Set("renderOrder", v)
}
func (dlh *DirectionalLightHelper) Rotation() *math.Euler {
	return &math.Euler{Value: dlh.Get("rotation")}
}
func (dlh *DirectionalLightHelper) SetRotation(v *math.Euler) {
	dlh.Set("rotation", v)
}
func (dlh *DirectionalLightHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: dlh.Get("scale")}
}
func (dlh *DirectionalLightHelper) SetScale(v *math.Vector3) {
	dlh.Set("scale", v)
}
func (dlh *DirectionalLightHelper) TargetPlane() *objects.Line {
	return &objects.Line{Value: dlh.Get("targetPlane")}
}
func (dlh *DirectionalLightHelper) SetTargetPlane(v *objects.Line) {
	dlh.Set("targetPlane", v)
}
func (dlh *DirectionalLightHelper) Type() string {
	return dlh.Get("type").String()
}
func (dlh *DirectionalLightHelper) SetType(v string) {
	dlh.Set("type", v)
}
func (dlh *DirectionalLightHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: dlh.Get("up")}
}
func (dlh *DirectionalLightHelper) SetUp(v *math.Vector3) {
	dlh.Set("up", v)
}
func (dlh *DirectionalLightHelper) UserData() js.Value {
	return dlh.Get("userData")
}
func (dlh *DirectionalLightHelper) SetUserData(v js.Value) {
	dlh.Set("userData", v)
}
func (dlh *DirectionalLightHelper) Uuid() string {
	return dlh.Get("uuid").String()
}
func (dlh *DirectionalLightHelper) SetUuid(v string) {
	dlh.Set("uuid", v)
}
func (dlh *DirectionalLightHelper) Visible() bool {
	return dlh.Get("visible").Bool()
}
func (dlh *DirectionalLightHelper) SetVisible(v bool) {
	dlh.Set("visible", v)
}
func (dlh *DirectionalLightHelper) DefaultMatrixAutoUpdate() bool {
	return dlh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (dlh *DirectionalLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	dlh.Set("DefaultMatrixAutoUpdate", v)
}
func (dlh *DirectionalLightHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: dlh.Get("DefaultUp")}
}
func (dlh *DirectionalLightHelper) SetDefaultUp(v *math.Vector3) {
	dlh.Set("DefaultUp", v)
}
func (dlh *DirectionalLightHelper) Add(object []core.Object3D) this {
	return this(dlh.Call("add", object))
}
func (dlh *DirectionalLightHelper) AddEventListener(typ string, listener js.Value) {
	dlh.Call("addEventListener", typ, listener)
}
func (dlh *DirectionalLightHelper) ApplyMatrix(matrix *math.Matrix4) {
	dlh.Call("applyMatrix", matrix)
}
func (dlh *DirectionalLightHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(dlh.Call("applyQuaternion", quaternion))
}
func (dlh *DirectionalLightHelper) Clone(recursive bool) this {
	return this(dlh.Call("clone", recursive))
}
func (dlh *DirectionalLightHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(dlh.Call("copy", source, recursive))
}
func (dlh *DirectionalLightHelper) DispatchEvent(event js.Value) {
	dlh.Call("dispatchEvent", event)
}
func (dlh *DirectionalLightHelper) Dispose() {
	dlh.Call("dispose")
}
func (dlh *DirectionalLightHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(dlh.Call("getObjectById", id))
}
func (dlh *DirectionalLightHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(dlh.Call("getObjectByName", name))
}
func (dlh *DirectionalLightHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(dlh.Call("getObjectByProperty", name, value))
}
func (dlh *DirectionalLightHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dlh.Call("getWorldDirection", target)}
}
func (dlh *DirectionalLightHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dlh.Call("getWorldPosition", target)}
}
func (dlh *DirectionalLightHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: dlh.Call("getWorldQuaternion", target)}
}
func (dlh *DirectionalLightHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dlh.Call("getWorldScale", target)}
}
func (dlh *DirectionalLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return dlh.Call("hasEventListener", typ, listener).Bool()
}
func (dlh *DirectionalLightHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dlh.Call("localToWorld", vector)}
}
func (dlh *DirectionalLightHelper) LookAt(vector math.Vector3, y int, z int) {
	dlh.Call("lookAt", vector, y, z)
}
func (dlh *DirectionalLightHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	dlh.Call("raycast", raycaster, intersects)
}
func (dlh *DirectionalLightHelper) Remove(object []core.Object3D) this {
	return this(dlh.Call("remove", object))
}
func (dlh *DirectionalLightHelper) RemoveEventListener(typ string, listener js.Value) {
	dlh.Call("removeEventListener", typ, listener)
}
func (dlh *DirectionalLightHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(dlh.Call("rotateOnAxis", axis, angle))
}
func (dlh *DirectionalLightHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(dlh.Call("rotateOnWorldAxis", axis, angle))
}
func (dlh *DirectionalLightHelper) RotateX(angle int) this {
	return this(dlh.Call("rotateX", angle))
}
func (dlh *DirectionalLightHelper) RotateY(angle int) this {
	return this(dlh.Call("rotateY", angle))
}
func (dlh *DirectionalLightHelper) RotateZ(angle int) this {
	return this(dlh.Call("rotateZ", angle))
}
func (dlh *DirectionalLightHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	dlh.Call("setRotationFromAxisAngle", axis, angle)
}
func (dlh *DirectionalLightHelper) SetRotationFromEuler(euler *math.Euler) {
	dlh.Call("setRotationFromEuler", euler)
}
func (dlh *DirectionalLightHelper) SetRotationFromMatrix(m *math.Matrix4) {
	dlh.Call("setRotationFromMatrix", m)
}
func (dlh *DirectionalLightHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	dlh.Call("setRotationFromQuaternion", q)
}
func (dlh *DirectionalLightHelper) ToJSON(meta js.Value) js.Value {
	return dlh.Call("toJSON", meta)
}
func (dlh *DirectionalLightHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(dlh.Call("translateOnAxis", axis, distance))
}
func (dlh *DirectionalLightHelper) TranslateX(distance int) this {
	return this(dlh.Call("translateX", distance))
}
func (dlh *DirectionalLightHelper) TranslateY(distance int) this {
	return this(dlh.Call("translateY", distance))
}
func (dlh *DirectionalLightHelper) TranslateZ(distance int) this {
	return this(dlh.Call("translateZ", distance))
}
func (dlh *DirectionalLightHelper) Traverse(callback js.Value) {
	dlh.Call("traverse", callback)
}
func (dlh *DirectionalLightHelper) TraverseAncestors(callback js.Value) {
	dlh.Call("traverseAncestors", callback)
}
func (dlh *DirectionalLightHelper) TraverseVisible(callback js.Value) {
	dlh.Call("traverseVisible", callback)
}
func (dlh *DirectionalLightHelper) Update() {
	dlh.Call("update")
}
func (dlh *DirectionalLightHelper) UpdateMatrix() {
	dlh.Call("updateMatrix")
}
func (dlh *DirectionalLightHelper) UpdateMatrixWorld(force bool) {
	dlh.Call("updateMatrixWorld", force)
}
func (dlh *DirectionalLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	dlh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (dlh *DirectionalLightHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dlh.Call("worldToLocal", vector)}
}

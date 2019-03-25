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

type HemisphereLightHelper struct {
	js.Value
}

func NewHemisphereLightHelper(light *lights.HemisphereLight, size int, color math.Color) *HemisphereLightHelper {
	return &HemisphereLightHelper{Value: get("HemisphereLightHelper").New(light, size, color)}
}
func (hlh *HemisphereLightHelper) CastShadow() bool {
	return hlh.Get("castShadow").Bool()
}
func (hlh *HemisphereLightHelper) SetCastShadow(v bool) {
	hlh.Set("castShadow", v)
}
func (hlh *HemisphereLightHelper) Children() []core.Object3D {
	return []core.Object3D(hlh.Get("children"))
}
func (hlh *HemisphereLightHelper) SetChildren(v []core.Object3D) {
	hlh.Set("children", v)
}
func (hlh *HemisphereLightHelper) Color() math.Color {
	return math.Color(hlh.Get("color"))
}
func (hlh *HemisphereLightHelper) SetColor(v math.Color) {
	hlh.Set("color", v)
}
func (hlh *HemisphereLightHelper) FrustumCulled() bool {
	return hlh.Get("frustumCulled").Bool()
}
func (hlh *HemisphereLightHelper) SetFrustumCulled(v bool) {
	hlh.Set("frustumCulled", v)
}
func (hlh *HemisphereLightHelper) Id() int {
	return hlh.Get("id").Int()
}
func (hlh *HemisphereLightHelper) SetId(v int) {
	hlh.Set("id", v)
}
func (hlh *HemisphereLightHelper) IsObject3D() true {
	return true(hlh.Get("isObject3D"))
}
func (hlh *HemisphereLightHelper) SetIsObject3D(v true) {
	hlh.Set("isObject3D", v)
}
func (hlh *HemisphereLightHelper) Layers() *core.Layers {
	return &core.Layers{Value: hlh.Get("layers")}
}
func (hlh *HemisphereLightHelper) SetLayers(v *core.Layers) {
	hlh.Set("layers", v)
}
func (hlh *HemisphereLightHelper) Light() *lights.HemisphereLight {
	return &lights.HemisphereLight{Value: hlh.Get("light")}
}
func (hlh *HemisphereLightHelper) SetLight(v *lights.HemisphereLight) {
	hlh.Set("light", v)
}
func (hlh *HemisphereLightHelper) Material() *materials.MeshBasicMaterial {
	return &materials.MeshBasicMaterial{Value: hlh.Get("material")}
}
func (hlh *HemisphereLightHelper) SetMaterial(v *materials.MeshBasicMaterial) {
	hlh.Set("material", v)
}
func (hlh *HemisphereLightHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: hlh.Get("matrix")}
}
func (hlh *HemisphereLightHelper) SetMatrix(v *math.Matrix4) {
	hlh.Set("matrix", v)
}
func (hlh *HemisphereLightHelper) MatrixAutoUpdate() bool {
	return hlh.Get("matrixAutoUpdate").Bool()
}
func (hlh *HemisphereLightHelper) SetMatrixAutoUpdate(v bool) {
	hlh.Set("matrixAutoUpdate", v)
}
func (hlh *HemisphereLightHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: hlh.Get("matrixWorld")}
}
func (hlh *HemisphereLightHelper) SetMatrixWorld(v *math.Matrix4) {
	hlh.Set("matrixWorld", v)
}
func (hlh *HemisphereLightHelper) MatrixWorldNeedsUpdate() bool {
	return hlh.Get("matrixWorldNeedsUpdate").Bool()
}
func (hlh *HemisphereLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	hlh.Set("matrixWorldNeedsUpdate", v)
}
func (hlh *HemisphereLightHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: hlh.Get("modelViewMatrix")}
}
func (hlh *HemisphereLightHelper) SetModelViewMatrix(v *math.Matrix4) {
	hlh.Set("modelViewMatrix", v)
}
func (hlh *HemisphereLightHelper) Name() string {
	return hlh.Get("name").String()
}
func (hlh *HemisphereLightHelper) SetName(v string) {
	hlh.Set("name", v)
}
func (hlh *HemisphereLightHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: hlh.Get("normalMatrix")}
}
func (hlh *HemisphereLightHelper) SetNormalMatrix(v *math.Matrix3) {
	hlh.Set("normalMatrix", v)
}
func (hlh *HemisphereLightHelper) OnAfterRender() js.Value {
	return hlh.Get("onAfterRender")
}
func (hlh *HemisphereLightHelper) SetOnAfterRender(v js.Value) {
	hlh.Set("onAfterRender", v)
}
func (hlh *HemisphereLightHelper) OnBeforeRender() js.Value {
	return hlh.Get("onBeforeRender")
}
func (hlh *HemisphereLightHelper) SetOnBeforeRender(v js.Value) {
	hlh.Set("onBeforeRender", v)
}
func (hlh *HemisphereLightHelper) Parent() core.Object3D {
	return core.Object3D(hlh.Get("parent"))
}
func (hlh *HemisphereLightHelper) SetParent(v core.Object3D) {
	hlh.Set("parent", v)
}
func (hlh *HemisphereLightHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: hlh.Get("position")}
}
func (hlh *HemisphereLightHelper) SetPosition(v *math.Vector3) {
	hlh.Set("position", v)
}
func (hlh *HemisphereLightHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: hlh.Get("quaternion")}
}
func (hlh *HemisphereLightHelper) SetQuaternion(v *math.Quaternion) {
	hlh.Set("quaternion", v)
}
func (hlh *HemisphereLightHelper) ReceiveShadow() bool {
	return hlh.Get("receiveShadow").Bool()
}
func (hlh *HemisphereLightHelper) SetReceiveShadow(v bool) {
	hlh.Set("receiveShadow", v)
}
func (hlh *HemisphereLightHelper) RenderOrder() int {
	return hlh.Get("renderOrder").Int()
}
func (hlh *HemisphereLightHelper) SetRenderOrder(v int) {
	hlh.Set("renderOrder", v)
}
func (hlh *HemisphereLightHelper) Rotation() *math.Euler {
	return &math.Euler{Value: hlh.Get("rotation")}
}
func (hlh *HemisphereLightHelper) SetRotation(v *math.Euler) {
	hlh.Set("rotation", v)
}
func (hlh *HemisphereLightHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: hlh.Get("scale")}
}
func (hlh *HemisphereLightHelper) SetScale(v *math.Vector3) {
	hlh.Set("scale", v)
}
func (hlh *HemisphereLightHelper) Type() string {
	return hlh.Get("type").String()
}
func (hlh *HemisphereLightHelper) SetType(v string) {
	hlh.Set("type", v)
}
func (hlh *HemisphereLightHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: hlh.Get("up")}
}
func (hlh *HemisphereLightHelper) SetUp(v *math.Vector3) {
	hlh.Set("up", v)
}
func (hlh *HemisphereLightHelper) UserData() js.Value {
	return hlh.Get("userData")
}
func (hlh *HemisphereLightHelper) SetUserData(v js.Value) {
	hlh.Set("userData", v)
}
func (hlh *HemisphereLightHelper) Uuid() string {
	return hlh.Get("uuid").String()
}
func (hlh *HemisphereLightHelper) SetUuid(v string) {
	hlh.Set("uuid", v)
}
func (hlh *HemisphereLightHelper) Visible() bool {
	return hlh.Get("visible").Bool()
}
func (hlh *HemisphereLightHelper) SetVisible(v bool) {
	hlh.Set("visible", v)
}
func (hlh *HemisphereLightHelper) DefaultMatrixAutoUpdate() bool {
	return hlh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (hlh *HemisphereLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	hlh.Set("DefaultMatrixAutoUpdate", v)
}
func (hlh *HemisphereLightHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: hlh.Get("DefaultUp")}
}
func (hlh *HemisphereLightHelper) SetDefaultUp(v *math.Vector3) {
	hlh.Set("DefaultUp", v)
}
func (hlh *HemisphereLightHelper) Add(object []core.Object3D) this {
	return this(hlh.Call("add", object))
}
func (hlh *HemisphereLightHelper) AddEventListener(typ string, listener js.Value) {
	hlh.Call("addEventListener", typ, listener)
}
func (hlh *HemisphereLightHelper) ApplyMatrix(matrix *math.Matrix4) {
	hlh.Call("applyMatrix", matrix)
}
func (hlh *HemisphereLightHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(hlh.Call("applyQuaternion", quaternion))
}
func (hlh *HemisphereLightHelper) Clone(recursive bool) this {
	return this(hlh.Call("clone", recursive))
}
func (hlh *HemisphereLightHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(hlh.Call("copy", source, recursive))
}
func (hlh *HemisphereLightHelper) DispatchEvent(event js.Value) {
	hlh.Call("dispatchEvent", event)
}
func (hlh *HemisphereLightHelper) Dispose() {
	hlh.Call("dispose")
}
func (hlh *HemisphereLightHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(hlh.Call("getObjectById", id))
}
func (hlh *HemisphereLightHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(hlh.Call("getObjectByName", name))
}
func (hlh *HemisphereLightHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(hlh.Call("getObjectByProperty", name, value))
}
func (hlh *HemisphereLightHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hlh.Call("getWorldDirection", target)}
}
func (hlh *HemisphereLightHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hlh.Call("getWorldPosition", target)}
}
func (hlh *HemisphereLightHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: hlh.Call("getWorldQuaternion", target)}
}
func (hlh *HemisphereLightHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hlh.Call("getWorldScale", target)}
}
func (hlh *HemisphereLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return hlh.Call("hasEventListener", typ, listener).Bool()
}
func (hlh *HemisphereLightHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hlh.Call("localToWorld", vector)}
}
func (hlh *HemisphereLightHelper) LookAt(vector math.Vector3, y int, z int) {
	hlh.Call("lookAt", vector, y, z)
}
func (hlh *HemisphereLightHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	hlh.Call("raycast", raycaster, intersects)
}
func (hlh *HemisphereLightHelper) Remove(object []core.Object3D) this {
	return this(hlh.Call("remove", object))
}
func (hlh *HemisphereLightHelper) RemoveEventListener(typ string, listener js.Value) {
	hlh.Call("removeEventListener", typ, listener)
}
func (hlh *HemisphereLightHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(hlh.Call("rotateOnAxis", axis, angle))
}
func (hlh *HemisphereLightHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(hlh.Call("rotateOnWorldAxis", axis, angle))
}
func (hlh *HemisphereLightHelper) RotateX(angle int) this {
	return this(hlh.Call("rotateX", angle))
}
func (hlh *HemisphereLightHelper) RotateY(angle int) this {
	return this(hlh.Call("rotateY", angle))
}
func (hlh *HemisphereLightHelper) RotateZ(angle int) this {
	return this(hlh.Call("rotateZ", angle))
}
func (hlh *HemisphereLightHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	hlh.Call("setRotationFromAxisAngle", axis, angle)
}
func (hlh *HemisphereLightHelper) SetRotationFromEuler(euler *math.Euler) {
	hlh.Call("setRotationFromEuler", euler)
}
func (hlh *HemisphereLightHelper) SetRotationFromMatrix(m *math.Matrix4) {
	hlh.Call("setRotationFromMatrix", m)
}
func (hlh *HemisphereLightHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	hlh.Call("setRotationFromQuaternion", q)
}
func (hlh *HemisphereLightHelper) ToJSON(meta js.Value) js.Value {
	return hlh.Call("toJSON", meta)
}
func (hlh *HemisphereLightHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(hlh.Call("translateOnAxis", axis, distance))
}
func (hlh *HemisphereLightHelper) TranslateX(distance int) this {
	return this(hlh.Call("translateX", distance))
}
func (hlh *HemisphereLightHelper) TranslateY(distance int) this {
	return this(hlh.Call("translateY", distance))
}
func (hlh *HemisphereLightHelper) TranslateZ(distance int) this {
	return this(hlh.Call("translateZ", distance))
}
func (hlh *HemisphereLightHelper) Traverse(callback js.Value) {
	hlh.Call("traverse", callback)
}
func (hlh *HemisphereLightHelper) TraverseAncestors(callback js.Value) {
	hlh.Call("traverseAncestors", callback)
}
func (hlh *HemisphereLightHelper) TraverseVisible(callback js.Value) {
	hlh.Call("traverseVisible", callback)
}
func (hlh *HemisphereLightHelper) Update() {
	hlh.Call("update")
}
func (hlh *HemisphereLightHelper) UpdateMatrix() {
	hlh.Call("updateMatrix")
}
func (hlh *HemisphereLightHelper) UpdateMatrixWorld(force bool) {
	hlh.Call("updateMatrixWorld", force)
}
func (hlh *HemisphereLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	hlh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (hlh *HemisphereLightHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hlh.Call("worldToLocal", vector)}
}

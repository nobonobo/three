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

type SpotLightHelper struct {
	js.Value
}

func NewSpotLightHelper(light *lights.Light, color math.Color) *SpotLightHelper {
	return &SpotLightHelper{Value: get("SpotLightHelper").New(light, color)}
}
func (slh *SpotLightHelper) CastShadow() bool {
	return slh.Get("castShadow").Bool()
}
func (slh *SpotLightHelper) SetCastShadow(v bool) {
	slh.Set("castShadow", v)
}
func (slh *SpotLightHelper) Children() []core.Object3D {
	return []core.Object3D(slh.Get("children"))
}
func (slh *SpotLightHelper) SetChildren(v []core.Object3D) {
	slh.Set("children", v)
}
func (slh *SpotLightHelper) Color() math.Color {
	return math.Color(slh.Get("color"))
}
func (slh *SpotLightHelper) SetColor(v math.Color) {
	slh.Set("color", v)
}
func (slh *SpotLightHelper) FrustumCulled() bool {
	return slh.Get("frustumCulled").Bool()
}
func (slh *SpotLightHelper) SetFrustumCulled(v bool) {
	slh.Set("frustumCulled", v)
}
func (slh *SpotLightHelper) Id() int {
	return slh.Get("id").Int()
}
func (slh *SpotLightHelper) SetId(v int) {
	slh.Set("id", v)
}
func (slh *SpotLightHelper) IsObject3D() true {
	return true(slh.Get("isObject3D"))
}
func (slh *SpotLightHelper) SetIsObject3D(v true) {
	slh.Set("isObject3D", v)
}
func (slh *SpotLightHelper) Layers() *core.Layers {
	return &core.Layers{Value: slh.Get("layers")}
}
func (slh *SpotLightHelper) SetLayers(v *core.Layers) {
	slh.Set("layers", v)
}
func (slh *SpotLightHelper) Light() *lights.Light {
	return &lights.Light{Value: slh.Get("light")}
}
func (slh *SpotLightHelper) SetLight(v *lights.Light) {
	slh.Set("light", v)
}
func (slh *SpotLightHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: slh.Get("matrix")}
}
func (slh *SpotLightHelper) SetMatrix(v *math.Matrix4) {
	slh.Set("matrix", v)
}
func (slh *SpotLightHelper) MatrixAutoUpdate() bool {
	return slh.Get("matrixAutoUpdate").Bool()
}
func (slh *SpotLightHelper) SetMatrixAutoUpdate(v bool) {
	slh.Set("matrixAutoUpdate", v)
}
func (slh *SpotLightHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: slh.Get("matrixWorld")}
}
func (slh *SpotLightHelper) SetMatrixWorld(v *math.Matrix4) {
	slh.Set("matrixWorld", v)
}
func (slh *SpotLightHelper) MatrixWorldNeedsUpdate() bool {
	return slh.Get("matrixWorldNeedsUpdate").Bool()
}
func (slh *SpotLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	slh.Set("matrixWorldNeedsUpdate", v)
}
func (slh *SpotLightHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: slh.Get("modelViewMatrix")}
}
func (slh *SpotLightHelper) SetModelViewMatrix(v *math.Matrix4) {
	slh.Set("modelViewMatrix", v)
}
func (slh *SpotLightHelper) Name() string {
	return slh.Get("name").String()
}
func (slh *SpotLightHelper) SetName(v string) {
	slh.Set("name", v)
}
func (slh *SpotLightHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: slh.Get("normalMatrix")}
}
func (slh *SpotLightHelper) SetNormalMatrix(v *math.Matrix3) {
	slh.Set("normalMatrix", v)
}
func (slh *SpotLightHelper) OnAfterRender() js.Value {
	return slh.Get("onAfterRender")
}
func (slh *SpotLightHelper) SetOnAfterRender(v js.Value) {
	slh.Set("onAfterRender", v)
}
func (slh *SpotLightHelper) OnBeforeRender() js.Value {
	return slh.Get("onBeforeRender")
}
func (slh *SpotLightHelper) SetOnBeforeRender(v js.Value) {
	slh.Set("onBeforeRender", v)
}
func (slh *SpotLightHelper) Parent() core.Object3D {
	return core.Object3D(slh.Get("parent"))
}
func (slh *SpotLightHelper) SetParent(v core.Object3D) {
	slh.Set("parent", v)
}
func (slh *SpotLightHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: slh.Get("position")}
}
func (slh *SpotLightHelper) SetPosition(v *math.Vector3) {
	slh.Set("position", v)
}
func (slh *SpotLightHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: slh.Get("quaternion")}
}
func (slh *SpotLightHelper) SetQuaternion(v *math.Quaternion) {
	slh.Set("quaternion", v)
}
func (slh *SpotLightHelper) ReceiveShadow() bool {
	return slh.Get("receiveShadow").Bool()
}
func (slh *SpotLightHelper) SetReceiveShadow(v bool) {
	slh.Set("receiveShadow", v)
}
func (slh *SpotLightHelper) RenderOrder() int {
	return slh.Get("renderOrder").Int()
}
func (slh *SpotLightHelper) SetRenderOrder(v int) {
	slh.Set("renderOrder", v)
}
func (slh *SpotLightHelper) Rotation() *math.Euler {
	return &math.Euler{Value: slh.Get("rotation")}
}
func (slh *SpotLightHelper) SetRotation(v *math.Euler) {
	slh.Set("rotation", v)
}
func (slh *SpotLightHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: slh.Get("scale")}
}
func (slh *SpotLightHelper) SetScale(v *math.Vector3) {
	slh.Set("scale", v)
}
func (slh *SpotLightHelper) Type() string {
	return slh.Get("type").String()
}
func (slh *SpotLightHelper) SetType(v string) {
	slh.Set("type", v)
}
func (slh *SpotLightHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: slh.Get("up")}
}
func (slh *SpotLightHelper) SetUp(v *math.Vector3) {
	slh.Set("up", v)
}
func (slh *SpotLightHelper) UserData() js.Value {
	return slh.Get("userData")
}
func (slh *SpotLightHelper) SetUserData(v js.Value) {
	slh.Set("userData", v)
}
func (slh *SpotLightHelper) Uuid() string {
	return slh.Get("uuid").String()
}
func (slh *SpotLightHelper) SetUuid(v string) {
	slh.Set("uuid", v)
}
func (slh *SpotLightHelper) Visible() bool {
	return slh.Get("visible").Bool()
}
func (slh *SpotLightHelper) SetVisible(v bool) {
	slh.Set("visible", v)
}
func (slh *SpotLightHelper) DefaultMatrixAutoUpdate() bool {
	return slh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (slh *SpotLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	slh.Set("DefaultMatrixAutoUpdate", v)
}
func (slh *SpotLightHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: slh.Get("DefaultUp")}
}
func (slh *SpotLightHelper) SetDefaultUp(v *math.Vector3) {
	slh.Set("DefaultUp", v)
}
func (slh *SpotLightHelper) Add(object []core.Object3D) this {
	return this(slh.Call("add", object))
}
func (slh *SpotLightHelper) AddEventListener(typ string, listener js.Value) {
	slh.Call("addEventListener", typ, listener)
}
func (slh *SpotLightHelper) ApplyMatrix(matrix *math.Matrix4) {
	slh.Call("applyMatrix", matrix)
}
func (slh *SpotLightHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(slh.Call("applyQuaternion", quaternion))
}
func (slh *SpotLightHelper) Clone(recursive bool) this {
	return this(slh.Call("clone", recursive))
}
func (slh *SpotLightHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(slh.Call("copy", source, recursive))
}
func (slh *SpotLightHelper) DispatchEvent(event js.Value) {
	slh.Call("dispatchEvent", event)
}
func (slh *SpotLightHelper) Dispose() {
	slh.Call("dispose")
}
func (slh *SpotLightHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(slh.Call("getObjectById", id))
}
func (slh *SpotLightHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(slh.Call("getObjectByName", name))
}
func (slh *SpotLightHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(slh.Call("getObjectByProperty", name, value))
}
func (slh *SpotLightHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: slh.Call("getWorldDirection", target)}
}
func (slh *SpotLightHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: slh.Call("getWorldPosition", target)}
}
func (slh *SpotLightHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: slh.Call("getWorldQuaternion", target)}
}
func (slh *SpotLightHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: slh.Call("getWorldScale", target)}
}
func (slh *SpotLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return slh.Call("hasEventListener", typ, listener).Bool()
}
func (slh *SpotLightHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: slh.Call("localToWorld", vector)}
}
func (slh *SpotLightHelper) LookAt(vector math.Vector3, y int, z int) {
	slh.Call("lookAt", vector, y, z)
}
func (slh *SpotLightHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	slh.Call("raycast", raycaster, intersects)
}
func (slh *SpotLightHelper) Remove(object []core.Object3D) this {
	return this(slh.Call("remove", object))
}
func (slh *SpotLightHelper) RemoveEventListener(typ string, listener js.Value) {
	slh.Call("removeEventListener", typ, listener)
}
func (slh *SpotLightHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(slh.Call("rotateOnAxis", axis, angle))
}
func (slh *SpotLightHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(slh.Call("rotateOnWorldAxis", axis, angle))
}
func (slh *SpotLightHelper) RotateX(angle int) this {
	return this(slh.Call("rotateX", angle))
}
func (slh *SpotLightHelper) RotateY(angle int) this {
	return this(slh.Call("rotateY", angle))
}
func (slh *SpotLightHelper) RotateZ(angle int) this {
	return this(slh.Call("rotateZ", angle))
}
func (slh *SpotLightHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	slh.Call("setRotationFromAxisAngle", axis, angle)
}
func (slh *SpotLightHelper) SetRotationFromEuler(euler *math.Euler) {
	slh.Call("setRotationFromEuler", euler)
}
func (slh *SpotLightHelper) SetRotationFromMatrix(m *math.Matrix4) {
	slh.Call("setRotationFromMatrix", m)
}
func (slh *SpotLightHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	slh.Call("setRotationFromQuaternion", q)
}
func (slh *SpotLightHelper) ToJSON(meta js.Value) js.Value {
	return slh.Call("toJSON", meta)
}
func (slh *SpotLightHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(slh.Call("translateOnAxis", axis, distance))
}
func (slh *SpotLightHelper) TranslateX(distance int) this {
	return this(slh.Call("translateX", distance))
}
func (slh *SpotLightHelper) TranslateY(distance int) this {
	return this(slh.Call("translateY", distance))
}
func (slh *SpotLightHelper) TranslateZ(distance int) this {
	return this(slh.Call("translateZ", distance))
}
func (slh *SpotLightHelper) Traverse(callback js.Value) {
	slh.Call("traverse", callback)
}
func (slh *SpotLightHelper) TraverseAncestors(callback js.Value) {
	slh.Call("traverseAncestors", callback)
}
func (slh *SpotLightHelper) TraverseVisible(callback js.Value) {
	slh.Call("traverseVisible", callback)
}
func (slh *SpotLightHelper) Update() {
	slh.Call("update")
}
func (slh *SpotLightHelper) UpdateMatrix() {
	slh.Call("updateMatrix")
}
func (slh *SpotLightHelper) UpdateMatrixWorld(force bool) {
	slh.Call("updateMatrixWorld", force)
}
func (slh *SpotLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	slh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (slh *SpotLightHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: slh.Call("worldToLocal", vector)}
}

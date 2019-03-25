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

type PointLightHelper struct {
	js.Value
}

func NewPointLightHelper(light *lights.PointLight, sphereSize int, color math.Color) *PointLightHelper {
	return &PointLightHelper{Value: get("PointLightHelper").New(light, sphereSize, color)}
}
func (plh *PointLightHelper) CastShadow() bool {
	return plh.Get("castShadow").Bool()
}
func (plh *PointLightHelper) SetCastShadow(v bool) {
	plh.Set("castShadow", v)
}
func (plh *PointLightHelper) Children() []core.Object3D {
	return []core.Object3D(plh.Get("children"))
}
func (plh *PointLightHelper) SetChildren(v []core.Object3D) {
	plh.Set("children", v)
}
func (plh *PointLightHelper) Color() math.Color {
	return math.Color(plh.Get("color"))
}
func (plh *PointLightHelper) SetColor(v math.Color) {
	plh.Set("color", v)
}
func (plh *PointLightHelper) FrustumCulled() bool {
	return plh.Get("frustumCulled").Bool()
}
func (plh *PointLightHelper) SetFrustumCulled(v bool) {
	plh.Set("frustumCulled", v)
}
func (plh *PointLightHelper) Id() int {
	return plh.Get("id").Int()
}
func (plh *PointLightHelper) SetId(v int) {
	plh.Set("id", v)
}
func (plh *PointLightHelper) IsObject3D() true {
	return true(plh.Get("isObject3D"))
}
func (plh *PointLightHelper) SetIsObject3D(v true) {
	plh.Set("isObject3D", v)
}
func (plh *PointLightHelper) Layers() *core.Layers {
	return &core.Layers{Value: plh.Get("layers")}
}
func (plh *PointLightHelper) SetLayers(v *core.Layers) {
	plh.Set("layers", v)
}
func (plh *PointLightHelper) Light() *lights.PointLight {
	return &lights.PointLight{Value: plh.Get("light")}
}
func (plh *PointLightHelper) SetLight(v *lights.PointLight) {
	plh.Set("light", v)
}
func (plh *PointLightHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: plh.Get("matrix")}
}
func (plh *PointLightHelper) SetMatrix(v *math.Matrix4) {
	plh.Set("matrix", v)
}
func (plh *PointLightHelper) MatrixAutoUpdate() bool {
	return plh.Get("matrixAutoUpdate").Bool()
}
func (plh *PointLightHelper) SetMatrixAutoUpdate(v bool) {
	plh.Set("matrixAutoUpdate", v)
}
func (plh *PointLightHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: plh.Get("matrixWorld")}
}
func (plh *PointLightHelper) SetMatrixWorld(v *math.Matrix4) {
	plh.Set("matrixWorld", v)
}
func (plh *PointLightHelper) MatrixWorldNeedsUpdate() bool {
	return plh.Get("matrixWorldNeedsUpdate").Bool()
}
func (plh *PointLightHelper) SetMatrixWorldNeedsUpdate(v bool) {
	plh.Set("matrixWorldNeedsUpdate", v)
}
func (plh *PointLightHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: plh.Get("modelViewMatrix")}
}
func (plh *PointLightHelper) SetModelViewMatrix(v *math.Matrix4) {
	plh.Set("modelViewMatrix", v)
}
func (plh *PointLightHelper) Name() string {
	return plh.Get("name").String()
}
func (plh *PointLightHelper) SetName(v string) {
	plh.Set("name", v)
}
func (plh *PointLightHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: plh.Get("normalMatrix")}
}
func (plh *PointLightHelper) SetNormalMatrix(v *math.Matrix3) {
	plh.Set("normalMatrix", v)
}
func (plh *PointLightHelper) OnAfterRender() js.Value {
	return plh.Get("onAfterRender")
}
func (plh *PointLightHelper) SetOnAfterRender(v js.Value) {
	plh.Set("onAfterRender", v)
}
func (plh *PointLightHelper) OnBeforeRender() js.Value {
	return plh.Get("onBeforeRender")
}
func (plh *PointLightHelper) SetOnBeforeRender(v js.Value) {
	plh.Set("onBeforeRender", v)
}
func (plh *PointLightHelper) Parent() core.Object3D {
	return core.Object3D(plh.Get("parent"))
}
func (plh *PointLightHelper) SetParent(v core.Object3D) {
	plh.Set("parent", v)
}
func (plh *PointLightHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: plh.Get("position")}
}
func (plh *PointLightHelper) SetPosition(v *math.Vector3) {
	plh.Set("position", v)
}
func (plh *PointLightHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: plh.Get("quaternion")}
}
func (plh *PointLightHelper) SetQuaternion(v *math.Quaternion) {
	plh.Set("quaternion", v)
}
func (plh *PointLightHelper) ReceiveShadow() bool {
	return plh.Get("receiveShadow").Bool()
}
func (plh *PointLightHelper) SetReceiveShadow(v bool) {
	plh.Set("receiveShadow", v)
}
func (plh *PointLightHelper) RenderOrder() int {
	return plh.Get("renderOrder").Int()
}
func (plh *PointLightHelper) SetRenderOrder(v int) {
	plh.Set("renderOrder", v)
}
func (plh *PointLightHelper) Rotation() *math.Euler {
	return &math.Euler{Value: plh.Get("rotation")}
}
func (plh *PointLightHelper) SetRotation(v *math.Euler) {
	plh.Set("rotation", v)
}
func (plh *PointLightHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: plh.Get("scale")}
}
func (plh *PointLightHelper) SetScale(v *math.Vector3) {
	plh.Set("scale", v)
}
func (plh *PointLightHelper) Type() string {
	return plh.Get("type").String()
}
func (plh *PointLightHelper) SetType(v string) {
	plh.Set("type", v)
}
func (plh *PointLightHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: plh.Get("up")}
}
func (plh *PointLightHelper) SetUp(v *math.Vector3) {
	plh.Set("up", v)
}
func (plh *PointLightHelper) UserData() js.Value {
	return plh.Get("userData")
}
func (plh *PointLightHelper) SetUserData(v js.Value) {
	plh.Set("userData", v)
}
func (plh *PointLightHelper) Uuid() string {
	return plh.Get("uuid").String()
}
func (plh *PointLightHelper) SetUuid(v string) {
	plh.Set("uuid", v)
}
func (plh *PointLightHelper) Visible() bool {
	return plh.Get("visible").Bool()
}
func (plh *PointLightHelper) SetVisible(v bool) {
	plh.Set("visible", v)
}
func (plh *PointLightHelper) DefaultMatrixAutoUpdate() bool {
	return plh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (plh *PointLightHelper) SetDefaultMatrixAutoUpdate(v bool) {
	plh.Set("DefaultMatrixAutoUpdate", v)
}
func (plh *PointLightHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: plh.Get("DefaultUp")}
}
func (plh *PointLightHelper) SetDefaultUp(v *math.Vector3) {
	plh.Set("DefaultUp", v)
}
func (plh *PointLightHelper) Add(object []core.Object3D) this {
	return this(plh.Call("add", object))
}
func (plh *PointLightHelper) AddEventListener(typ string, listener js.Value) {
	plh.Call("addEventListener", typ, listener)
}
func (plh *PointLightHelper) ApplyMatrix(matrix *math.Matrix4) {
	plh.Call("applyMatrix", matrix)
}
func (plh *PointLightHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(plh.Call("applyQuaternion", quaternion))
}
func (plh *PointLightHelper) Clone(recursive bool) this {
	return this(plh.Call("clone", recursive))
}
func (plh *PointLightHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(plh.Call("copy", source, recursive))
}
func (plh *PointLightHelper) DispatchEvent(event js.Value) {
	plh.Call("dispatchEvent", event)
}
func (plh *PointLightHelper) Dispose() {
	plh.Call("dispose")
}
func (plh *PointLightHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(plh.Call("getObjectById", id))
}
func (plh *PointLightHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(plh.Call("getObjectByName", name))
}
func (plh *PointLightHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(plh.Call("getObjectByProperty", name, value))
}
func (plh *PointLightHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: plh.Call("getWorldDirection", target)}
}
func (plh *PointLightHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: plh.Call("getWorldPosition", target)}
}
func (plh *PointLightHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: plh.Call("getWorldQuaternion", target)}
}
func (plh *PointLightHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: plh.Call("getWorldScale", target)}
}
func (plh *PointLightHelper) HasEventListener(typ string, listener js.Value) bool {
	return plh.Call("hasEventListener", typ, listener).Bool()
}
func (plh *PointLightHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: plh.Call("localToWorld", vector)}
}
func (plh *PointLightHelper) LookAt(vector math.Vector3, y int, z int) {
	plh.Call("lookAt", vector, y, z)
}
func (plh *PointLightHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	plh.Call("raycast", raycaster, intersects)
}
func (plh *PointLightHelper) Remove(object []core.Object3D) this {
	return this(plh.Call("remove", object))
}
func (plh *PointLightHelper) RemoveEventListener(typ string, listener js.Value) {
	plh.Call("removeEventListener", typ, listener)
}
func (plh *PointLightHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(plh.Call("rotateOnAxis", axis, angle))
}
func (plh *PointLightHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(plh.Call("rotateOnWorldAxis", axis, angle))
}
func (plh *PointLightHelper) RotateX(angle int) this {
	return this(plh.Call("rotateX", angle))
}
func (plh *PointLightHelper) RotateY(angle int) this {
	return this(plh.Call("rotateY", angle))
}
func (plh *PointLightHelper) RotateZ(angle int) this {
	return this(plh.Call("rotateZ", angle))
}
func (plh *PointLightHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	plh.Call("setRotationFromAxisAngle", axis, angle)
}
func (plh *PointLightHelper) SetRotationFromEuler(euler *math.Euler) {
	plh.Call("setRotationFromEuler", euler)
}
func (plh *PointLightHelper) SetRotationFromMatrix(m *math.Matrix4) {
	plh.Call("setRotationFromMatrix", m)
}
func (plh *PointLightHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	plh.Call("setRotationFromQuaternion", q)
}
func (plh *PointLightHelper) ToJSON(meta js.Value) js.Value {
	return plh.Call("toJSON", meta)
}
func (plh *PointLightHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(plh.Call("translateOnAxis", axis, distance))
}
func (plh *PointLightHelper) TranslateX(distance int) this {
	return this(plh.Call("translateX", distance))
}
func (plh *PointLightHelper) TranslateY(distance int) this {
	return this(plh.Call("translateY", distance))
}
func (plh *PointLightHelper) TranslateZ(distance int) this {
	return this(plh.Call("translateZ", distance))
}
func (plh *PointLightHelper) Traverse(callback js.Value) {
	plh.Call("traverse", callback)
}
func (plh *PointLightHelper) TraverseAncestors(callback js.Value) {
	plh.Call("traverseAncestors", callback)
}
func (plh *PointLightHelper) TraverseVisible(callback js.Value) {
	plh.Call("traverseVisible", callback)
}
func (plh *PointLightHelper) Update() {
	plh.Call("update")
}
func (plh *PointLightHelper) UpdateMatrix() {
	plh.Call("updateMatrix")
}
func (plh *PointLightHelper) UpdateMatrixWorld(force bool) {
	plh.Call("updateMatrixWorld", force)
}
func (plh *PointLightHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	plh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (plh *PointLightHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: plh.Call("worldToLocal", vector)}
}

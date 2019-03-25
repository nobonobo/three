package lights

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

type Light struct {
	js.Value
}

func NewLight(hex int, intensity int) *Light {
	return &Light{Value: get("Light").New(hex, intensity)}
}
func (l *Light) CastShadow() bool {
	return l.Get("castShadow").Bool()
}
func (l *Light) SetCastShadow(v bool) {
	l.Set("castShadow", v)
}
func (l *Light) Children() []core.Object3D {
	return []core.Object3D(l.Get("children"))
}
func (l *Light) SetChildren(v []core.Object3D) {
	l.Set("children", v)
}
func (l *Light) Color() *math.Color {
	return &math.Color{Value: l.Get("color")}
}
func (l *Light) SetColor(v *math.Color) {
	l.Set("color", v)
}
func (l *Light) FrustumCulled() bool {
	return l.Get("frustumCulled").Bool()
}
func (l *Light) SetFrustumCulled(v bool) {
	l.Set("frustumCulled", v)
}
func (l *Light) Id() int {
	return l.Get("id").Int()
}
func (l *Light) SetId(v int) {
	l.Set("id", v)
}
func (l *Light) Intensity() int {
	return l.Get("intensity").Int()
}
func (l *Light) SetIntensity(v int) {
	l.Set("intensity", v)
}
func (l *Light) IsObject3D() true {
	return true(l.Get("isObject3D"))
}
func (l *Light) SetIsObject3D(v true) {
	l.Set("isObject3D", v)
}
func (l *Light) Layers() *core.Layers {
	return &core.Layers{Value: l.Get("layers")}
}
func (l *Light) SetLayers(v *core.Layers) {
	l.Set("layers", v)
}
func (l *Light) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: l.Get("matrix")}
}
func (l *Light) SetMatrix(v *math.Matrix4) {
	l.Set("matrix", v)
}
func (l *Light) MatrixAutoUpdate() bool {
	return l.Get("matrixAutoUpdate").Bool()
}
func (l *Light) SetMatrixAutoUpdate(v bool) {
	l.Set("matrixAutoUpdate", v)
}
func (l *Light) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: l.Get("matrixWorld")}
}
func (l *Light) SetMatrixWorld(v *math.Matrix4) {
	l.Set("matrixWorld", v)
}
func (l *Light) MatrixWorldNeedsUpdate() bool {
	return l.Get("matrixWorldNeedsUpdate").Bool()
}
func (l *Light) SetMatrixWorldNeedsUpdate(v bool) {
	l.Set("matrixWorldNeedsUpdate", v)
}
func (l *Light) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: l.Get("modelViewMatrix")}
}
func (l *Light) SetModelViewMatrix(v *math.Matrix4) {
	l.Set("modelViewMatrix", v)
}
func (l *Light) Name() string {
	return l.Get("name").String()
}
func (l *Light) SetName(v string) {
	l.Set("name", v)
}
func (l *Light) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: l.Get("normalMatrix")}
}
func (l *Light) SetNormalMatrix(v *math.Matrix3) {
	l.Set("normalMatrix", v)
}
func (l *Light) OnAfterRender() js.Value {
	return l.Get("onAfterRender")
}
func (l *Light) SetOnAfterRender(v js.Value) {
	l.Set("onAfterRender", v)
}
func (l *Light) OnBeforeRender() js.Value {
	return l.Get("onBeforeRender")
}
func (l *Light) SetOnBeforeRender(v js.Value) {
	l.Set("onBeforeRender", v)
}
func (l *Light) Parent() core.Object3D {
	return core.Object3D(l.Get("parent"))
}
func (l *Light) SetParent(v core.Object3D) {
	l.Set("parent", v)
}
func (l *Light) Position() *math.Vector3 {
	return &math.Vector3{Value: l.Get("position")}
}
func (l *Light) SetPosition(v *math.Vector3) {
	l.Set("position", v)
}
func (l *Light) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: l.Get("quaternion")}
}
func (l *Light) SetQuaternion(v *math.Quaternion) {
	l.Set("quaternion", v)
}
func (l *Light) ReceiveShadow() bool {
	return l.Get("receiveShadow").Bool()
}
func (l *Light) SetReceiveShadow(v bool) {
	l.Set("receiveShadow", v)
}
func (l *Light) RenderOrder() int {
	return l.Get("renderOrder").Int()
}
func (l *Light) SetRenderOrder(v int) {
	l.Set("renderOrder", v)
}
func (l *Light) Rotation() *math.Euler {
	return &math.Euler{Value: l.Get("rotation")}
}
func (l *Light) SetRotation(v *math.Euler) {
	l.Set("rotation", v)
}
func (l *Light) Scale() *math.Vector3 {
	return &math.Vector3{Value: l.Get("scale")}
}
func (l *Light) SetScale(v *math.Vector3) {
	l.Set("scale", v)
}
func (l *Light) Shadow() *LightShadow {
	return &LightShadow{Value: l.Get("shadow")}
}
func (l *Light) SetShadow(v *LightShadow) {
	l.Set("shadow", v)
}
func (l *Light) ShadowBias() js.Value {
	return l.Get("shadowBias")
}
func (l *Light) SetShadowBias(v js.Value) {
	l.Set("shadowBias", v)
}
func (l *Light) ShadowCameraBottom() js.Value {
	return l.Get("shadowCameraBottom")
}
func (l *Light) SetShadowCameraBottom(v js.Value) {
	l.Set("shadowCameraBottom", v)
}
func (l *Light) ShadowCameraFar() js.Value {
	return l.Get("shadowCameraFar")
}
func (l *Light) SetShadowCameraFar(v js.Value) {
	l.Set("shadowCameraFar", v)
}
func (l *Light) ShadowCameraFov() js.Value {
	return l.Get("shadowCameraFov")
}
func (l *Light) SetShadowCameraFov(v js.Value) {
	l.Set("shadowCameraFov", v)
}
func (l *Light) ShadowCameraLeft() js.Value {
	return l.Get("shadowCameraLeft")
}
func (l *Light) SetShadowCameraLeft(v js.Value) {
	l.Set("shadowCameraLeft", v)
}
func (l *Light) ShadowCameraNear() js.Value {
	return l.Get("shadowCameraNear")
}
func (l *Light) SetShadowCameraNear(v js.Value) {
	l.Set("shadowCameraNear", v)
}
func (l *Light) ShadowCameraRight() js.Value {
	return l.Get("shadowCameraRight")
}
func (l *Light) SetShadowCameraRight(v js.Value) {
	l.Set("shadowCameraRight", v)
}
func (l *Light) ShadowCameraTop() js.Value {
	return l.Get("shadowCameraTop")
}
func (l *Light) SetShadowCameraTop(v js.Value) {
	l.Set("shadowCameraTop", v)
}
func (l *Light) ShadowMapHeight() js.Value {
	return l.Get("shadowMapHeight")
}
func (l *Light) SetShadowMapHeight(v js.Value) {
	l.Set("shadowMapHeight", v)
}
func (l *Light) ShadowMapWidth() js.Value {
	return l.Get("shadowMapWidth")
}
func (l *Light) SetShadowMapWidth(v js.Value) {
	l.Set("shadowMapWidth", v)
}
func (l *Light) Type() string {
	return l.Get("type").String()
}
func (l *Light) SetType(v string) {
	l.Set("type", v)
}
func (l *Light) Up() *math.Vector3 {
	return &math.Vector3{Value: l.Get("up")}
}
func (l *Light) SetUp(v *math.Vector3) {
	l.Set("up", v)
}
func (l *Light) UserData() js.Value {
	return l.Get("userData")
}
func (l *Light) SetUserData(v js.Value) {
	l.Set("userData", v)
}
func (l *Light) Uuid() string {
	return l.Get("uuid").String()
}
func (l *Light) SetUuid(v string) {
	l.Set("uuid", v)
}
func (l *Light) Visible() bool {
	return l.Get("visible").Bool()
}
func (l *Light) SetVisible(v bool) {
	l.Set("visible", v)
}
func (l *Light) DefaultMatrixAutoUpdate() bool {
	return l.Get("DefaultMatrixAutoUpdate").Bool()
}
func (l *Light) SetDefaultMatrixAutoUpdate(v bool) {
	l.Set("DefaultMatrixAutoUpdate", v)
}
func (l *Light) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: l.Get("DefaultUp")}
}
func (l *Light) SetDefaultUp(v *math.Vector3) {
	l.Set("DefaultUp", v)
}
func (l *Light) Add(object []core.Object3D) this {
	return this(l.Call("add", object))
}
func (l *Light) AddEventListener(typ string, listener js.Value) {
	l.Call("addEventListener", typ, listener)
}
func (l *Light) ApplyMatrix(matrix *math.Matrix4) {
	l.Call("applyMatrix", matrix)
}
func (l *Light) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(l.Call("applyQuaternion", quaternion))
}
func (l *Light) Clone(recursive bool) this {
	return this(l.Call("clone", recursive))
}
func (l *Light) Copy(source *core.Object3D, recursive bool) this {
	return this(l.Call("copy", source, recursive))
}
func (l *Light) DispatchEvent(event js.Value) {
	l.Call("dispatchEvent", event)
}
func (l *Light) GetObjectById(id int) core.Object3D {
	return core.Object3D(l.Call("getObjectById", id))
}
func (l *Light) GetObjectByName(name string) core.Object3D {
	return core.Object3D(l.Call("getObjectByName", name))
}
func (l *Light) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(l.Call("getObjectByProperty", name, value))
}
func (l *Light) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("getWorldDirection", target)}
}
func (l *Light) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("getWorldPosition", target)}
}
func (l *Light) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: l.Call("getWorldQuaternion", target)}
}
func (l *Light) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("getWorldScale", target)}
}
func (l *Light) HasEventListener(typ string, listener js.Value) bool {
	return l.Call("hasEventListener", typ, listener).Bool()
}
func (l *Light) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("localToWorld", vector)}
}
func (l *Light) LookAt(vector math.Vector3, y int, z int) {
	l.Call("lookAt", vector, y, z)
}
func (l *Light) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	l.Call("raycast", raycaster, intersects)
}
func (l *Light) Remove(object []core.Object3D) this {
	return this(l.Call("remove", object))
}
func (l *Light) RemoveEventListener(typ string, listener js.Value) {
	l.Call("removeEventListener", typ, listener)
}
func (l *Light) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(l.Call("rotateOnAxis", axis, angle))
}
func (l *Light) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(l.Call("rotateOnWorldAxis", axis, angle))
}
func (l *Light) RotateX(angle int) this {
	return this(l.Call("rotateX", angle))
}
func (l *Light) RotateY(angle int) this {
	return this(l.Call("rotateY", angle))
}
func (l *Light) RotateZ(angle int) this {
	return this(l.Call("rotateZ", angle))
}
func (l *Light) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	l.Call("setRotationFromAxisAngle", axis, angle)
}
func (l *Light) SetRotationFromEuler(euler *math.Euler) {
	l.Call("setRotationFromEuler", euler)
}
func (l *Light) SetRotationFromMatrix(m *math.Matrix4) {
	l.Call("setRotationFromMatrix", m)
}
func (l *Light) SetRotationFromQuaternion(q *math.Quaternion) {
	l.Call("setRotationFromQuaternion", q)
}
func (l *Light) ToJSON(meta js.Value) js.Value {
	return l.Call("toJSON", meta)
}
func (l *Light) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(l.Call("translateOnAxis", axis, distance))
}
func (l *Light) TranslateX(distance int) this {
	return this(l.Call("translateX", distance))
}
func (l *Light) TranslateY(distance int) this {
	return this(l.Call("translateY", distance))
}
func (l *Light) TranslateZ(distance int) this {
	return this(l.Call("translateZ", distance))
}
func (l *Light) Traverse(callback js.Value) {
	l.Call("traverse", callback)
}
func (l *Light) TraverseAncestors(callback js.Value) {
	l.Call("traverseAncestors", callback)
}
func (l *Light) TraverseVisible(callback js.Value) {
	l.Call("traverseVisible", callback)
}
func (l *Light) UpdateMatrix() {
	l.Call("updateMatrix")
}
func (l *Light) UpdateMatrixWorld(force bool) {
	l.Call("updateMatrixWorld", force)
}
func (l *Light) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	l.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (l *Light) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: l.Call("worldToLocal", vector)}
}

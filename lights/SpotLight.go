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

type SpotLight struct {
	js.Value
}

func NewSpotLight(color math.Color, intensity int, distance int, angle int, exponent int, decay int) *SpotLight {
	return &SpotLight{Value: get("SpotLight").New(color, intensity, distance, angle, exponent, decay)}
}
func (sl *SpotLight) Angle() int {
	return sl.Get("angle").Int()
}
func (sl *SpotLight) SetAngle(v int) {
	sl.Set("angle", v)
}
func (sl *SpotLight) CastShadow() bool {
	return sl.Get("castShadow").Bool()
}
func (sl *SpotLight) SetCastShadow(v bool) {
	sl.Set("castShadow", v)
}
func (sl *SpotLight) Children() []core.Object3D {
	return []core.Object3D(sl.Get("children"))
}
func (sl *SpotLight) SetChildren(v []core.Object3D) {
	sl.Set("children", v)
}
func (sl *SpotLight) Color() *math.Color {
	return &math.Color{Value: sl.Get("color")}
}
func (sl *SpotLight) SetColor(v *math.Color) {
	sl.Set("color", v)
}
func (sl *SpotLight) Decay() int {
	return sl.Get("decay").Int()
}
func (sl *SpotLight) SetDecay(v int) {
	sl.Set("decay", v)
}
func (sl *SpotLight) Distance() int {
	return sl.Get("distance").Int()
}
func (sl *SpotLight) SetDistance(v int) {
	sl.Set("distance", v)
}
func (sl *SpotLight) Exponent() int {
	return sl.Get("exponent").Int()
}
func (sl *SpotLight) SetExponent(v int) {
	sl.Set("exponent", v)
}
func (sl *SpotLight) FrustumCulled() bool {
	return sl.Get("frustumCulled").Bool()
}
func (sl *SpotLight) SetFrustumCulled(v bool) {
	sl.Set("frustumCulled", v)
}
func (sl *SpotLight) Id() int {
	return sl.Get("id").Int()
}
func (sl *SpotLight) SetId(v int) {
	sl.Set("id", v)
}
func (sl *SpotLight) Intensity() int {
	return sl.Get("intensity").Int()
}
func (sl *SpotLight) SetIntensity(v int) {
	sl.Set("intensity", v)
}
func (sl *SpotLight) IsObject3D() true {
	return true(sl.Get("isObject3D"))
}
func (sl *SpotLight) SetIsObject3D(v true) {
	sl.Set("isObject3D", v)
}
func (sl *SpotLight) Layers() *core.Layers {
	return &core.Layers{Value: sl.Get("layers")}
}
func (sl *SpotLight) SetLayers(v *core.Layers) {
	sl.Set("layers", v)
}
func (sl *SpotLight) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: sl.Get("matrix")}
}
func (sl *SpotLight) SetMatrix(v *math.Matrix4) {
	sl.Set("matrix", v)
}
func (sl *SpotLight) MatrixAutoUpdate() bool {
	return sl.Get("matrixAutoUpdate").Bool()
}
func (sl *SpotLight) SetMatrixAutoUpdate(v bool) {
	sl.Set("matrixAutoUpdate", v)
}
func (sl *SpotLight) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: sl.Get("matrixWorld")}
}
func (sl *SpotLight) SetMatrixWorld(v *math.Matrix4) {
	sl.Set("matrixWorld", v)
}
func (sl *SpotLight) MatrixWorldNeedsUpdate() bool {
	return sl.Get("matrixWorldNeedsUpdate").Bool()
}
func (sl *SpotLight) SetMatrixWorldNeedsUpdate(v bool) {
	sl.Set("matrixWorldNeedsUpdate", v)
}
func (sl *SpotLight) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: sl.Get("modelViewMatrix")}
}
func (sl *SpotLight) SetModelViewMatrix(v *math.Matrix4) {
	sl.Set("modelViewMatrix", v)
}
func (sl *SpotLight) Name() string {
	return sl.Get("name").String()
}
func (sl *SpotLight) SetName(v string) {
	sl.Set("name", v)
}
func (sl *SpotLight) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: sl.Get("normalMatrix")}
}
func (sl *SpotLight) SetNormalMatrix(v *math.Matrix3) {
	sl.Set("normalMatrix", v)
}
func (sl *SpotLight) OnAfterRender() js.Value {
	return sl.Get("onAfterRender")
}
func (sl *SpotLight) SetOnAfterRender(v js.Value) {
	sl.Set("onAfterRender", v)
}
func (sl *SpotLight) OnBeforeRender() js.Value {
	return sl.Get("onBeforeRender")
}
func (sl *SpotLight) SetOnBeforeRender(v js.Value) {
	sl.Set("onBeforeRender", v)
}
func (sl *SpotLight) Parent() core.Object3D {
	return core.Object3D(sl.Get("parent"))
}
func (sl *SpotLight) SetParent(v core.Object3D) {
	sl.Set("parent", v)
}
func (sl *SpotLight) Penumbra() int {
	return sl.Get("penumbra").Int()
}
func (sl *SpotLight) SetPenumbra(v int) {
	sl.Set("penumbra", v)
}
func (sl *SpotLight) Position() *math.Vector3 {
	return &math.Vector3{Value: sl.Get("position")}
}
func (sl *SpotLight) SetPosition(v *math.Vector3) {
	sl.Set("position", v)
}
func (sl *SpotLight) Power() int {
	return sl.Get("power").Int()
}
func (sl *SpotLight) SetPower(v int) {
	sl.Set("power", v)
}
func (sl *SpotLight) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: sl.Get("quaternion")}
}
func (sl *SpotLight) SetQuaternion(v *math.Quaternion) {
	sl.Set("quaternion", v)
}
func (sl *SpotLight) ReceiveShadow() bool {
	return sl.Get("receiveShadow").Bool()
}
func (sl *SpotLight) SetReceiveShadow(v bool) {
	sl.Set("receiveShadow", v)
}
func (sl *SpotLight) RenderOrder() int {
	return sl.Get("renderOrder").Int()
}
func (sl *SpotLight) SetRenderOrder(v int) {
	sl.Set("renderOrder", v)
}
func (sl *SpotLight) Rotation() *math.Euler {
	return &math.Euler{Value: sl.Get("rotation")}
}
func (sl *SpotLight) SetRotation(v *math.Euler) {
	sl.Set("rotation", v)
}
func (sl *SpotLight) Scale() *math.Vector3 {
	return &math.Vector3{Value: sl.Get("scale")}
}
func (sl *SpotLight) SetScale(v *math.Vector3) {
	sl.Set("scale", v)
}
func (sl *SpotLight) Shadow() *SpotLightShadow {
	return &SpotLightShadow{Value: sl.Get("shadow")}
}
func (sl *SpotLight) SetShadow(v *SpotLightShadow) {
	sl.Set("shadow", v)
}
func (sl *SpotLight) ShadowBias() js.Value {
	return sl.Get("shadowBias")
}
func (sl *SpotLight) SetShadowBias(v js.Value) {
	sl.Set("shadowBias", v)
}
func (sl *SpotLight) ShadowCameraBottom() js.Value {
	return sl.Get("shadowCameraBottom")
}
func (sl *SpotLight) SetShadowCameraBottom(v js.Value) {
	sl.Set("shadowCameraBottom", v)
}
func (sl *SpotLight) ShadowCameraFar() js.Value {
	return sl.Get("shadowCameraFar")
}
func (sl *SpotLight) SetShadowCameraFar(v js.Value) {
	sl.Set("shadowCameraFar", v)
}
func (sl *SpotLight) ShadowCameraFov() js.Value {
	return sl.Get("shadowCameraFov")
}
func (sl *SpotLight) SetShadowCameraFov(v js.Value) {
	sl.Set("shadowCameraFov", v)
}
func (sl *SpotLight) ShadowCameraLeft() js.Value {
	return sl.Get("shadowCameraLeft")
}
func (sl *SpotLight) SetShadowCameraLeft(v js.Value) {
	sl.Set("shadowCameraLeft", v)
}
func (sl *SpotLight) ShadowCameraNear() js.Value {
	return sl.Get("shadowCameraNear")
}
func (sl *SpotLight) SetShadowCameraNear(v js.Value) {
	sl.Set("shadowCameraNear", v)
}
func (sl *SpotLight) ShadowCameraRight() js.Value {
	return sl.Get("shadowCameraRight")
}
func (sl *SpotLight) SetShadowCameraRight(v js.Value) {
	sl.Set("shadowCameraRight", v)
}
func (sl *SpotLight) ShadowCameraTop() js.Value {
	return sl.Get("shadowCameraTop")
}
func (sl *SpotLight) SetShadowCameraTop(v js.Value) {
	sl.Set("shadowCameraTop", v)
}
func (sl *SpotLight) ShadowMapHeight() js.Value {
	return sl.Get("shadowMapHeight")
}
func (sl *SpotLight) SetShadowMapHeight(v js.Value) {
	sl.Set("shadowMapHeight", v)
}
func (sl *SpotLight) ShadowMapWidth() js.Value {
	return sl.Get("shadowMapWidth")
}
func (sl *SpotLight) SetShadowMapWidth(v js.Value) {
	sl.Set("shadowMapWidth", v)
}
func (sl *SpotLight) Target() *core.Object3D {
	return &core.Object3D{Value: sl.Get("target")}
}
func (sl *SpotLight) SetTarget(v *core.Object3D) {
	sl.Set("target", v)
}
func (sl *SpotLight) Type() string {
	return sl.Get("type").String()
}
func (sl *SpotLight) SetType(v string) {
	sl.Set("type", v)
}
func (sl *SpotLight) Up() *math.Vector3 {
	return &math.Vector3{Value: sl.Get("up")}
}
func (sl *SpotLight) SetUp(v *math.Vector3) {
	sl.Set("up", v)
}
func (sl *SpotLight) UserData() js.Value {
	return sl.Get("userData")
}
func (sl *SpotLight) SetUserData(v js.Value) {
	sl.Set("userData", v)
}
func (sl *SpotLight) Uuid() string {
	return sl.Get("uuid").String()
}
func (sl *SpotLight) SetUuid(v string) {
	sl.Set("uuid", v)
}
func (sl *SpotLight) Visible() bool {
	return sl.Get("visible").Bool()
}
func (sl *SpotLight) SetVisible(v bool) {
	sl.Set("visible", v)
}
func (sl *SpotLight) DefaultMatrixAutoUpdate() bool {
	return sl.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sl *SpotLight) SetDefaultMatrixAutoUpdate(v bool) {
	sl.Set("DefaultMatrixAutoUpdate", v)
}
func (sl *SpotLight) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: sl.Get("DefaultUp")}
}
func (sl *SpotLight) SetDefaultUp(v *math.Vector3) {
	sl.Set("DefaultUp", v)
}
func (sl *SpotLight) Add(object []core.Object3D) this {
	return this(sl.Call("add", object))
}
func (sl *SpotLight) AddEventListener(typ string, listener js.Value) {
	sl.Call("addEventListener", typ, listener)
}
func (sl *SpotLight) ApplyMatrix(matrix *math.Matrix4) {
	sl.Call("applyMatrix", matrix)
}
func (sl *SpotLight) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(sl.Call("applyQuaternion", quaternion))
}
func (sl *SpotLight) Clone(recursive bool) this {
	return this(sl.Call("clone", recursive))
}
func (sl *SpotLight) Copy(source *core.Object3D, recursive bool) this {
	return this(sl.Call("copy", source, recursive))
}
func (sl *SpotLight) DispatchEvent(event js.Value) {
	sl.Call("dispatchEvent", event)
}
func (sl *SpotLight) GetObjectById(id int) core.Object3D {
	return core.Object3D(sl.Call("getObjectById", id))
}
func (sl *SpotLight) GetObjectByName(name string) core.Object3D {
	return core.Object3D(sl.Call("getObjectByName", name))
}
func (sl *SpotLight) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(sl.Call("getObjectByProperty", name, value))
}
func (sl *SpotLight) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sl.Call("getWorldDirection", target)}
}
func (sl *SpotLight) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sl.Call("getWorldPosition", target)}
}
func (sl *SpotLight) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: sl.Call("getWorldQuaternion", target)}
}
func (sl *SpotLight) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sl.Call("getWorldScale", target)}
}
func (sl *SpotLight) HasEventListener(typ string, listener js.Value) bool {
	return sl.Call("hasEventListener", typ, listener).Bool()
}
func (sl *SpotLight) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sl.Call("localToWorld", vector)}
}
func (sl *SpotLight) LookAt(vector math.Vector3, y int, z int) {
	sl.Call("lookAt", vector, y, z)
}
func (sl *SpotLight) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	sl.Call("raycast", raycaster, intersects)
}
func (sl *SpotLight) Remove(object []core.Object3D) this {
	return this(sl.Call("remove", object))
}
func (sl *SpotLight) RemoveEventListener(typ string, listener js.Value) {
	sl.Call("removeEventListener", typ, listener)
}
func (sl *SpotLight) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(sl.Call("rotateOnAxis", axis, angle))
}
func (sl *SpotLight) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(sl.Call("rotateOnWorldAxis", axis, angle))
}
func (sl *SpotLight) RotateX(angle int) this {
	return this(sl.Call("rotateX", angle))
}
func (sl *SpotLight) RotateY(angle int) this {
	return this(sl.Call("rotateY", angle))
}
func (sl *SpotLight) RotateZ(angle int) this {
	return this(sl.Call("rotateZ", angle))
}
func (sl *SpotLight) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	sl.Call("setRotationFromAxisAngle", axis, angle)
}
func (sl *SpotLight) SetRotationFromEuler(euler *math.Euler) {
	sl.Call("setRotationFromEuler", euler)
}
func (sl *SpotLight) SetRotationFromMatrix(m *math.Matrix4) {
	sl.Call("setRotationFromMatrix", m)
}
func (sl *SpotLight) SetRotationFromQuaternion(q *math.Quaternion) {
	sl.Call("setRotationFromQuaternion", q)
}
func (sl *SpotLight) ToJSON(meta js.Value) js.Value {
	return sl.Call("toJSON", meta)
}
func (sl *SpotLight) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(sl.Call("translateOnAxis", axis, distance))
}
func (sl *SpotLight) TranslateX(distance int) this {
	return this(sl.Call("translateX", distance))
}
func (sl *SpotLight) TranslateY(distance int) this {
	return this(sl.Call("translateY", distance))
}
func (sl *SpotLight) TranslateZ(distance int) this {
	return this(sl.Call("translateZ", distance))
}
func (sl *SpotLight) Traverse(callback js.Value) {
	sl.Call("traverse", callback)
}
func (sl *SpotLight) TraverseAncestors(callback js.Value) {
	sl.Call("traverseAncestors", callback)
}
func (sl *SpotLight) TraverseVisible(callback js.Value) {
	sl.Call("traverseVisible", callback)
}
func (sl *SpotLight) UpdateMatrix() {
	sl.Call("updateMatrix")
}
func (sl *SpotLight) UpdateMatrixWorld(force bool) {
	sl.Call("updateMatrixWorld", force)
}
func (sl *SpotLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sl.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sl *SpotLight) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sl.Call("worldToLocal", vector)}
}

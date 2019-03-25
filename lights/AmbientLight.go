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

type AmbientLight struct {
	js.Value
}

func NewAmbientLight(color math.Color, intensity int) *AmbientLight {
	return &AmbientLight{Value: get("AmbientLight").New(color, intensity)}
}
func (al *AmbientLight) CastShadow() bool {
	return al.Get("castShadow").Bool()
}
func (al *AmbientLight) SetCastShadow(v bool) {
	al.Set("castShadow", v)
}
func (al *AmbientLight) Children() []core.Object3D {
	return []core.Object3D(al.Get("children"))
}
func (al *AmbientLight) SetChildren(v []core.Object3D) {
	al.Set("children", v)
}
func (al *AmbientLight) Color() *math.Color {
	return &math.Color{Value: al.Get("color")}
}
func (al *AmbientLight) SetColor(v *math.Color) {
	al.Set("color", v)
}
func (al *AmbientLight) FrustumCulled() bool {
	return al.Get("frustumCulled").Bool()
}
func (al *AmbientLight) SetFrustumCulled(v bool) {
	al.Set("frustumCulled", v)
}
func (al *AmbientLight) Id() int {
	return al.Get("id").Int()
}
func (al *AmbientLight) SetId(v int) {
	al.Set("id", v)
}
func (al *AmbientLight) Intensity() int {
	return al.Get("intensity").Int()
}
func (al *AmbientLight) SetIntensity(v int) {
	al.Set("intensity", v)
}
func (al *AmbientLight) IsObject3D() true {
	return true(al.Get("isObject3D"))
}
func (al *AmbientLight) SetIsObject3D(v true) {
	al.Set("isObject3D", v)
}
func (al *AmbientLight) Layers() *core.Layers {
	return &core.Layers{Value: al.Get("layers")}
}
func (al *AmbientLight) SetLayers(v *core.Layers) {
	al.Set("layers", v)
}
func (al *AmbientLight) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: al.Get("matrix")}
}
func (al *AmbientLight) SetMatrix(v *math.Matrix4) {
	al.Set("matrix", v)
}
func (al *AmbientLight) MatrixAutoUpdate() bool {
	return al.Get("matrixAutoUpdate").Bool()
}
func (al *AmbientLight) SetMatrixAutoUpdate(v bool) {
	al.Set("matrixAutoUpdate", v)
}
func (al *AmbientLight) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: al.Get("matrixWorld")}
}
func (al *AmbientLight) SetMatrixWorld(v *math.Matrix4) {
	al.Set("matrixWorld", v)
}
func (al *AmbientLight) MatrixWorldNeedsUpdate() bool {
	return al.Get("matrixWorldNeedsUpdate").Bool()
}
func (al *AmbientLight) SetMatrixWorldNeedsUpdate(v bool) {
	al.Set("matrixWorldNeedsUpdate", v)
}
func (al *AmbientLight) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: al.Get("modelViewMatrix")}
}
func (al *AmbientLight) SetModelViewMatrix(v *math.Matrix4) {
	al.Set("modelViewMatrix", v)
}
func (al *AmbientLight) Name() string {
	return al.Get("name").String()
}
func (al *AmbientLight) SetName(v string) {
	al.Set("name", v)
}
func (al *AmbientLight) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: al.Get("normalMatrix")}
}
func (al *AmbientLight) SetNormalMatrix(v *math.Matrix3) {
	al.Set("normalMatrix", v)
}
func (al *AmbientLight) OnAfterRender() js.Value {
	return al.Get("onAfterRender")
}
func (al *AmbientLight) SetOnAfterRender(v js.Value) {
	al.Set("onAfterRender", v)
}
func (al *AmbientLight) OnBeforeRender() js.Value {
	return al.Get("onBeforeRender")
}
func (al *AmbientLight) SetOnBeforeRender(v js.Value) {
	al.Set("onBeforeRender", v)
}
func (al *AmbientLight) Parent() core.Object3D {
	return core.Object3D(al.Get("parent"))
}
func (al *AmbientLight) SetParent(v core.Object3D) {
	al.Set("parent", v)
}
func (al *AmbientLight) Position() *math.Vector3 {
	return &math.Vector3{Value: al.Get("position")}
}
func (al *AmbientLight) SetPosition(v *math.Vector3) {
	al.Set("position", v)
}
func (al *AmbientLight) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: al.Get("quaternion")}
}
func (al *AmbientLight) SetQuaternion(v *math.Quaternion) {
	al.Set("quaternion", v)
}
func (al *AmbientLight) ReceiveShadow() bool {
	return al.Get("receiveShadow").Bool()
}
func (al *AmbientLight) SetReceiveShadow(v bool) {
	al.Set("receiveShadow", v)
}
func (al *AmbientLight) RenderOrder() int {
	return al.Get("renderOrder").Int()
}
func (al *AmbientLight) SetRenderOrder(v int) {
	al.Set("renderOrder", v)
}
func (al *AmbientLight) Rotation() *math.Euler {
	return &math.Euler{Value: al.Get("rotation")}
}
func (al *AmbientLight) SetRotation(v *math.Euler) {
	al.Set("rotation", v)
}
func (al *AmbientLight) Scale() *math.Vector3 {
	return &math.Vector3{Value: al.Get("scale")}
}
func (al *AmbientLight) SetScale(v *math.Vector3) {
	al.Set("scale", v)
}
func (al *AmbientLight) Shadow() *LightShadow {
	return &LightShadow{Value: al.Get("shadow")}
}
func (al *AmbientLight) SetShadow(v *LightShadow) {
	al.Set("shadow", v)
}
func (al *AmbientLight) ShadowBias() js.Value {
	return al.Get("shadowBias")
}
func (al *AmbientLight) SetShadowBias(v js.Value) {
	al.Set("shadowBias", v)
}
func (al *AmbientLight) ShadowCameraBottom() js.Value {
	return al.Get("shadowCameraBottom")
}
func (al *AmbientLight) SetShadowCameraBottom(v js.Value) {
	al.Set("shadowCameraBottom", v)
}
func (al *AmbientLight) ShadowCameraFar() js.Value {
	return al.Get("shadowCameraFar")
}
func (al *AmbientLight) SetShadowCameraFar(v js.Value) {
	al.Set("shadowCameraFar", v)
}
func (al *AmbientLight) ShadowCameraFov() js.Value {
	return al.Get("shadowCameraFov")
}
func (al *AmbientLight) SetShadowCameraFov(v js.Value) {
	al.Set("shadowCameraFov", v)
}
func (al *AmbientLight) ShadowCameraLeft() js.Value {
	return al.Get("shadowCameraLeft")
}
func (al *AmbientLight) SetShadowCameraLeft(v js.Value) {
	al.Set("shadowCameraLeft", v)
}
func (al *AmbientLight) ShadowCameraNear() js.Value {
	return al.Get("shadowCameraNear")
}
func (al *AmbientLight) SetShadowCameraNear(v js.Value) {
	al.Set("shadowCameraNear", v)
}
func (al *AmbientLight) ShadowCameraRight() js.Value {
	return al.Get("shadowCameraRight")
}
func (al *AmbientLight) SetShadowCameraRight(v js.Value) {
	al.Set("shadowCameraRight", v)
}
func (al *AmbientLight) ShadowCameraTop() js.Value {
	return al.Get("shadowCameraTop")
}
func (al *AmbientLight) SetShadowCameraTop(v js.Value) {
	al.Set("shadowCameraTop", v)
}
func (al *AmbientLight) ShadowMapHeight() js.Value {
	return al.Get("shadowMapHeight")
}
func (al *AmbientLight) SetShadowMapHeight(v js.Value) {
	al.Set("shadowMapHeight", v)
}
func (al *AmbientLight) ShadowMapWidth() js.Value {
	return al.Get("shadowMapWidth")
}
func (al *AmbientLight) SetShadowMapWidth(v js.Value) {
	al.Set("shadowMapWidth", v)
}
func (al *AmbientLight) Type() string {
	return al.Get("type").String()
}
func (al *AmbientLight) SetType(v string) {
	al.Set("type", v)
}
func (al *AmbientLight) Up() *math.Vector3 {
	return &math.Vector3{Value: al.Get("up")}
}
func (al *AmbientLight) SetUp(v *math.Vector3) {
	al.Set("up", v)
}
func (al *AmbientLight) UserData() js.Value {
	return al.Get("userData")
}
func (al *AmbientLight) SetUserData(v js.Value) {
	al.Set("userData", v)
}
func (al *AmbientLight) Uuid() string {
	return al.Get("uuid").String()
}
func (al *AmbientLight) SetUuid(v string) {
	al.Set("uuid", v)
}
func (al *AmbientLight) Visible() bool {
	return al.Get("visible").Bool()
}
func (al *AmbientLight) SetVisible(v bool) {
	al.Set("visible", v)
}
func (al *AmbientLight) DefaultMatrixAutoUpdate() bool {
	return al.Get("DefaultMatrixAutoUpdate").Bool()
}
func (al *AmbientLight) SetDefaultMatrixAutoUpdate(v bool) {
	al.Set("DefaultMatrixAutoUpdate", v)
}
func (al *AmbientLight) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: al.Get("DefaultUp")}
}
func (al *AmbientLight) SetDefaultUp(v *math.Vector3) {
	al.Set("DefaultUp", v)
}
func (al *AmbientLight) Add(object []core.Object3D) this {
	return this(al.Call("add", object))
}
func (al *AmbientLight) AddEventListener(typ string, listener js.Value) {
	al.Call("addEventListener", typ, listener)
}
func (al *AmbientLight) ApplyMatrix(matrix *math.Matrix4) {
	al.Call("applyMatrix", matrix)
}
func (al *AmbientLight) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(al.Call("applyQuaternion", quaternion))
}
func (al *AmbientLight) Clone(recursive bool) this {
	return this(al.Call("clone", recursive))
}
func (al *AmbientLight) Copy(source *core.Object3D, recursive bool) this {
	return this(al.Call("copy", source, recursive))
}
func (al *AmbientLight) DispatchEvent(event js.Value) {
	al.Call("dispatchEvent", event)
}
func (al *AmbientLight) GetObjectById(id int) core.Object3D {
	return core.Object3D(al.Call("getObjectById", id))
}
func (al *AmbientLight) GetObjectByName(name string) core.Object3D {
	return core.Object3D(al.Call("getObjectByName", name))
}
func (al *AmbientLight) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(al.Call("getObjectByProperty", name, value))
}
func (al *AmbientLight) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("getWorldDirection", target)}
}
func (al *AmbientLight) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("getWorldPosition", target)}
}
func (al *AmbientLight) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: al.Call("getWorldQuaternion", target)}
}
func (al *AmbientLight) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("getWorldScale", target)}
}
func (al *AmbientLight) HasEventListener(typ string, listener js.Value) bool {
	return al.Call("hasEventListener", typ, listener).Bool()
}
func (al *AmbientLight) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("localToWorld", vector)}
}
func (al *AmbientLight) LookAt(vector math.Vector3, y int, z int) {
	al.Call("lookAt", vector, y, z)
}
func (al *AmbientLight) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	al.Call("raycast", raycaster, intersects)
}
func (al *AmbientLight) Remove(object []core.Object3D) this {
	return this(al.Call("remove", object))
}
func (al *AmbientLight) RemoveEventListener(typ string, listener js.Value) {
	al.Call("removeEventListener", typ, listener)
}
func (al *AmbientLight) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(al.Call("rotateOnAxis", axis, angle))
}
func (al *AmbientLight) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(al.Call("rotateOnWorldAxis", axis, angle))
}
func (al *AmbientLight) RotateX(angle int) this {
	return this(al.Call("rotateX", angle))
}
func (al *AmbientLight) RotateY(angle int) this {
	return this(al.Call("rotateY", angle))
}
func (al *AmbientLight) RotateZ(angle int) this {
	return this(al.Call("rotateZ", angle))
}
func (al *AmbientLight) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	al.Call("setRotationFromAxisAngle", axis, angle)
}
func (al *AmbientLight) SetRotationFromEuler(euler *math.Euler) {
	al.Call("setRotationFromEuler", euler)
}
func (al *AmbientLight) SetRotationFromMatrix(m *math.Matrix4) {
	al.Call("setRotationFromMatrix", m)
}
func (al *AmbientLight) SetRotationFromQuaternion(q *math.Quaternion) {
	al.Call("setRotationFromQuaternion", q)
}
func (al *AmbientLight) ToJSON(meta js.Value) js.Value {
	return al.Call("toJSON", meta)
}
func (al *AmbientLight) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(al.Call("translateOnAxis", axis, distance))
}
func (al *AmbientLight) TranslateX(distance int) this {
	return this(al.Call("translateX", distance))
}
func (al *AmbientLight) TranslateY(distance int) this {
	return this(al.Call("translateY", distance))
}
func (al *AmbientLight) TranslateZ(distance int) this {
	return this(al.Call("translateZ", distance))
}
func (al *AmbientLight) Traverse(callback js.Value) {
	al.Call("traverse", callback)
}
func (al *AmbientLight) TraverseAncestors(callback js.Value) {
	al.Call("traverseAncestors", callback)
}
func (al *AmbientLight) TraverseVisible(callback js.Value) {
	al.Call("traverseVisible", callback)
}
func (al *AmbientLight) UpdateMatrix() {
	al.Call("updateMatrix")
}
func (al *AmbientLight) UpdateMatrixWorld(force bool) {
	al.Call("updateMatrixWorld", force)
}
func (al *AmbientLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	al.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (al *AmbientLight) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("worldToLocal", vector)}
}

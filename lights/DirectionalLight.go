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

type DirectionalLight struct {
	js.Value
}

func NewDirectionalLight(color math.Color, intensity int) *DirectionalLight {
	return &DirectionalLight{Value: get("DirectionalLight").New(color, intensity)}
}
func (dl *DirectionalLight) CastShadow() bool {
	return dl.Get("castShadow").Bool()
}
func (dl *DirectionalLight) SetCastShadow(v bool) {
	dl.Set("castShadow", v)
}
func (dl *DirectionalLight) Children() []core.Object3D {
	return []core.Object3D(dl.Get("children"))
}
func (dl *DirectionalLight) SetChildren(v []core.Object3D) {
	dl.Set("children", v)
}
func (dl *DirectionalLight) Color() *math.Color {
	return &math.Color{Value: dl.Get("color")}
}
func (dl *DirectionalLight) SetColor(v *math.Color) {
	dl.Set("color", v)
}
func (dl *DirectionalLight) FrustumCulled() bool {
	return dl.Get("frustumCulled").Bool()
}
func (dl *DirectionalLight) SetFrustumCulled(v bool) {
	dl.Set("frustumCulled", v)
}
func (dl *DirectionalLight) Id() int {
	return dl.Get("id").Int()
}
func (dl *DirectionalLight) SetId(v int) {
	dl.Set("id", v)
}
func (dl *DirectionalLight) Intensity() int {
	return dl.Get("intensity").Int()
}
func (dl *DirectionalLight) SetIntensity(v int) {
	dl.Set("intensity", v)
}
func (dl *DirectionalLight) IsObject3D() true {
	return true(dl.Get("isObject3D"))
}
func (dl *DirectionalLight) SetIsObject3D(v true) {
	dl.Set("isObject3D", v)
}
func (dl *DirectionalLight) Layers() *core.Layers {
	return &core.Layers{Value: dl.Get("layers")}
}
func (dl *DirectionalLight) SetLayers(v *core.Layers) {
	dl.Set("layers", v)
}
func (dl *DirectionalLight) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: dl.Get("matrix")}
}
func (dl *DirectionalLight) SetMatrix(v *math.Matrix4) {
	dl.Set("matrix", v)
}
func (dl *DirectionalLight) MatrixAutoUpdate() bool {
	return dl.Get("matrixAutoUpdate").Bool()
}
func (dl *DirectionalLight) SetMatrixAutoUpdate(v bool) {
	dl.Set("matrixAutoUpdate", v)
}
func (dl *DirectionalLight) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: dl.Get("matrixWorld")}
}
func (dl *DirectionalLight) SetMatrixWorld(v *math.Matrix4) {
	dl.Set("matrixWorld", v)
}
func (dl *DirectionalLight) MatrixWorldNeedsUpdate() bool {
	return dl.Get("matrixWorldNeedsUpdate").Bool()
}
func (dl *DirectionalLight) SetMatrixWorldNeedsUpdate(v bool) {
	dl.Set("matrixWorldNeedsUpdate", v)
}
func (dl *DirectionalLight) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: dl.Get("modelViewMatrix")}
}
func (dl *DirectionalLight) SetModelViewMatrix(v *math.Matrix4) {
	dl.Set("modelViewMatrix", v)
}
func (dl *DirectionalLight) Name() string {
	return dl.Get("name").String()
}
func (dl *DirectionalLight) SetName(v string) {
	dl.Set("name", v)
}
func (dl *DirectionalLight) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: dl.Get("normalMatrix")}
}
func (dl *DirectionalLight) SetNormalMatrix(v *math.Matrix3) {
	dl.Set("normalMatrix", v)
}
func (dl *DirectionalLight) OnAfterRender() js.Value {
	return dl.Get("onAfterRender")
}
func (dl *DirectionalLight) SetOnAfterRender(v js.Value) {
	dl.Set("onAfterRender", v)
}
func (dl *DirectionalLight) OnBeforeRender() js.Value {
	return dl.Get("onBeforeRender")
}
func (dl *DirectionalLight) SetOnBeforeRender(v js.Value) {
	dl.Set("onBeforeRender", v)
}
func (dl *DirectionalLight) Parent() core.Object3D {
	return core.Object3D(dl.Get("parent"))
}
func (dl *DirectionalLight) SetParent(v core.Object3D) {
	dl.Set("parent", v)
}
func (dl *DirectionalLight) Position() *math.Vector3 {
	return &math.Vector3{Value: dl.Get("position")}
}
func (dl *DirectionalLight) SetPosition(v *math.Vector3) {
	dl.Set("position", v)
}
func (dl *DirectionalLight) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: dl.Get("quaternion")}
}
func (dl *DirectionalLight) SetQuaternion(v *math.Quaternion) {
	dl.Set("quaternion", v)
}
func (dl *DirectionalLight) ReceiveShadow() bool {
	return dl.Get("receiveShadow").Bool()
}
func (dl *DirectionalLight) SetReceiveShadow(v bool) {
	dl.Set("receiveShadow", v)
}
func (dl *DirectionalLight) RenderOrder() int {
	return dl.Get("renderOrder").Int()
}
func (dl *DirectionalLight) SetRenderOrder(v int) {
	dl.Set("renderOrder", v)
}
func (dl *DirectionalLight) Rotation() *math.Euler {
	return &math.Euler{Value: dl.Get("rotation")}
}
func (dl *DirectionalLight) SetRotation(v *math.Euler) {
	dl.Set("rotation", v)
}
func (dl *DirectionalLight) Scale() *math.Vector3 {
	return &math.Vector3{Value: dl.Get("scale")}
}
func (dl *DirectionalLight) SetScale(v *math.Vector3) {
	dl.Set("scale", v)
}
func (dl *DirectionalLight) Shadow() *DirectionalLightShadow {
	return &DirectionalLightShadow{Value: dl.Get("shadow")}
}
func (dl *DirectionalLight) SetShadow(v *DirectionalLightShadow) {
	dl.Set("shadow", v)
}
func (dl *DirectionalLight) ShadowBias() js.Value {
	return dl.Get("shadowBias")
}
func (dl *DirectionalLight) SetShadowBias(v js.Value) {
	dl.Set("shadowBias", v)
}
func (dl *DirectionalLight) ShadowCameraBottom() js.Value {
	return dl.Get("shadowCameraBottom")
}
func (dl *DirectionalLight) SetShadowCameraBottom(v js.Value) {
	dl.Set("shadowCameraBottom", v)
}
func (dl *DirectionalLight) ShadowCameraFar() js.Value {
	return dl.Get("shadowCameraFar")
}
func (dl *DirectionalLight) SetShadowCameraFar(v js.Value) {
	dl.Set("shadowCameraFar", v)
}
func (dl *DirectionalLight) ShadowCameraFov() js.Value {
	return dl.Get("shadowCameraFov")
}
func (dl *DirectionalLight) SetShadowCameraFov(v js.Value) {
	dl.Set("shadowCameraFov", v)
}
func (dl *DirectionalLight) ShadowCameraLeft() js.Value {
	return dl.Get("shadowCameraLeft")
}
func (dl *DirectionalLight) SetShadowCameraLeft(v js.Value) {
	dl.Set("shadowCameraLeft", v)
}
func (dl *DirectionalLight) ShadowCameraNear() js.Value {
	return dl.Get("shadowCameraNear")
}
func (dl *DirectionalLight) SetShadowCameraNear(v js.Value) {
	dl.Set("shadowCameraNear", v)
}
func (dl *DirectionalLight) ShadowCameraRight() js.Value {
	return dl.Get("shadowCameraRight")
}
func (dl *DirectionalLight) SetShadowCameraRight(v js.Value) {
	dl.Set("shadowCameraRight", v)
}
func (dl *DirectionalLight) ShadowCameraTop() js.Value {
	return dl.Get("shadowCameraTop")
}
func (dl *DirectionalLight) SetShadowCameraTop(v js.Value) {
	dl.Set("shadowCameraTop", v)
}
func (dl *DirectionalLight) ShadowMapHeight() js.Value {
	return dl.Get("shadowMapHeight")
}
func (dl *DirectionalLight) SetShadowMapHeight(v js.Value) {
	dl.Set("shadowMapHeight", v)
}
func (dl *DirectionalLight) ShadowMapWidth() js.Value {
	return dl.Get("shadowMapWidth")
}
func (dl *DirectionalLight) SetShadowMapWidth(v js.Value) {
	dl.Set("shadowMapWidth", v)
}
func (dl *DirectionalLight) Target() *core.Object3D {
	return &core.Object3D{Value: dl.Get("target")}
}
func (dl *DirectionalLight) SetTarget(v *core.Object3D) {
	dl.Set("target", v)
}
func (dl *DirectionalLight) Type() string {
	return dl.Get("type").String()
}
func (dl *DirectionalLight) SetType(v string) {
	dl.Set("type", v)
}
func (dl *DirectionalLight) Up() *math.Vector3 {
	return &math.Vector3{Value: dl.Get("up")}
}
func (dl *DirectionalLight) SetUp(v *math.Vector3) {
	dl.Set("up", v)
}
func (dl *DirectionalLight) UserData() js.Value {
	return dl.Get("userData")
}
func (dl *DirectionalLight) SetUserData(v js.Value) {
	dl.Set("userData", v)
}
func (dl *DirectionalLight) Uuid() string {
	return dl.Get("uuid").String()
}
func (dl *DirectionalLight) SetUuid(v string) {
	dl.Set("uuid", v)
}
func (dl *DirectionalLight) Visible() bool {
	return dl.Get("visible").Bool()
}
func (dl *DirectionalLight) SetVisible(v bool) {
	dl.Set("visible", v)
}
func (dl *DirectionalLight) DefaultMatrixAutoUpdate() bool {
	return dl.Get("DefaultMatrixAutoUpdate").Bool()
}
func (dl *DirectionalLight) SetDefaultMatrixAutoUpdate(v bool) {
	dl.Set("DefaultMatrixAutoUpdate", v)
}
func (dl *DirectionalLight) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: dl.Get("DefaultUp")}
}
func (dl *DirectionalLight) SetDefaultUp(v *math.Vector3) {
	dl.Set("DefaultUp", v)
}
func (dl *DirectionalLight) Add(object []core.Object3D) this {
	return this(dl.Call("add", object))
}
func (dl *DirectionalLight) AddEventListener(typ string, listener js.Value) {
	dl.Call("addEventListener", typ, listener)
}
func (dl *DirectionalLight) ApplyMatrix(matrix *math.Matrix4) {
	dl.Call("applyMatrix", matrix)
}
func (dl *DirectionalLight) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(dl.Call("applyQuaternion", quaternion))
}
func (dl *DirectionalLight) Clone(recursive bool) this {
	return this(dl.Call("clone", recursive))
}
func (dl *DirectionalLight) Copy(source *core.Object3D, recursive bool) this {
	return this(dl.Call("copy", source, recursive))
}
func (dl *DirectionalLight) DispatchEvent(event js.Value) {
	dl.Call("dispatchEvent", event)
}
func (dl *DirectionalLight) GetObjectById(id int) core.Object3D {
	return core.Object3D(dl.Call("getObjectById", id))
}
func (dl *DirectionalLight) GetObjectByName(name string) core.Object3D {
	return core.Object3D(dl.Call("getObjectByName", name))
}
func (dl *DirectionalLight) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(dl.Call("getObjectByProperty", name, value))
}
func (dl *DirectionalLight) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dl.Call("getWorldDirection", target)}
}
func (dl *DirectionalLight) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dl.Call("getWorldPosition", target)}
}
func (dl *DirectionalLight) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: dl.Call("getWorldQuaternion", target)}
}
func (dl *DirectionalLight) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dl.Call("getWorldScale", target)}
}
func (dl *DirectionalLight) HasEventListener(typ string, listener js.Value) bool {
	return dl.Call("hasEventListener", typ, listener).Bool()
}
func (dl *DirectionalLight) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dl.Call("localToWorld", vector)}
}
func (dl *DirectionalLight) LookAt(vector math.Vector3, y int, z int) {
	dl.Call("lookAt", vector, y, z)
}
func (dl *DirectionalLight) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	dl.Call("raycast", raycaster, intersects)
}
func (dl *DirectionalLight) Remove(object []core.Object3D) this {
	return this(dl.Call("remove", object))
}
func (dl *DirectionalLight) RemoveEventListener(typ string, listener js.Value) {
	dl.Call("removeEventListener", typ, listener)
}
func (dl *DirectionalLight) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(dl.Call("rotateOnAxis", axis, angle))
}
func (dl *DirectionalLight) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(dl.Call("rotateOnWorldAxis", axis, angle))
}
func (dl *DirectionalLight) RotateX(angle int) this {
	return this(dl.Call("rotateX", angle))
}
func (dl *DirectionalLight) RotateY(angle int) this {
	return this(dl.Call("rotateY", angle))
}
func (dl *DirectionalLight) RotateZ(angle int) this {
	return this(dl.Call("rotateZ", angle))
}
func (dl *DirectionalLight) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	dl.Call("setRotationFromAxisAngle", axis, angle)
}
func (dl *DirectionalLight) SetRotationFromEuler(euler *math.Euler) {
	dl.Call("setRotationFromEuler", euler)
}
func (dl *DirectionalLight) SetRotationFromMatrix(m *math.Matrix4) {
	dl.Call("setRotationFromMatrix", m)
}
func (dl *DirectionalLight) SetRotationFromQuaternion(q *math.Quaternion) {
	dl.Call("setRotationFromQuaternion", q)
}
func (dl *DirectionalLight) ToJSON(meta js.Value) js.Value {
	return dl.Call("toJSON", meta)
}
func (dl *DirectionalLight) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(dl.Call("translateOnAxis", axis, distance))
}
func (dl *DirectionalLight) TranslateX(distance int) this {
	return this(dl.Call("translateX", distance))
}
func (dl *DirectionalLight) TranslateY(distance int) this {
	return this(dl.Call("translateY", distance))
}
func (dl *DirectionalLight) TranslateZ(distance int) this {
	return this(dl.Call("translateZ", distance))
}
func (dl *DirectionalLight) Traverse(callback js.Value) {
	dl.Call("traverse", callback)
}
func (dl *DirectionalLight) TraverseAncestors(callback js.Value) {
	dl.Call("traverseAncestors", callback)
}
func (dl *DirectionalLight) TraverseVisible(callback js.Value) {
	dl.Call("traverseVisible", callback)
}
func (dl *DirectionalLight) UpdateMatrix() {
	dl.Call("updateMatrix")
}
func (dl *DirectionalLight) UpdateMatrixWorld(force bool) {
	dl.Call("updateMatrixWorld", force)
}
func (dl *DirectionalLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	dl.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (dl *DirectionalLight) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: dl.Call("worldToLocal", vector)}
}

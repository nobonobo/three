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

type RectAreaLight struct {
	js.Value
}

func NewRectAreaLight(color math.Color, intensity int, width int, height int) *RectAreaLight {
	return &RectAreaLight{Value: get("RectAreaLight").New(color, intensity, width, height)}
}
func (ral *RectAreaLight) CastShadow() bool {
	return ral.Get("castShadow").Bool()
}
func (ral *RectAreaLight) SetCastShadow(v bool) {
	ral.Set("castShadow", v)
}
func (ral *RectAreaLight) Children() []core.Object3D {
	return []core.Object3D(ral.Get("children"))
}
func (ral *RectAreaLight) SetChildren(v []core.Object3D) {
	ral.Set("children", v)
}
func (ral *RectAreaLight) Color() *math.Color {
	return &math.Color{Value: ral.Get("color")}
}
func (ral *RectAreaLight) SetColor(v *math.Color) {
	ral.Set("color", v)
}
func (ral *RectAreaLight) FrustumCulled() bool {
	return ral.Get("frustumCulled").Bool()
}
func (ral *RectAreaLight) SetFrustumCulled(v bool) {
	ral.Set("frustumCulled", v)
}
func (ral *RectAreaLight) Height() int {
	return ral.Get("height").Int()
}
func (ral *RectAreaLight) SetHeight(v int) {
	ral.Set("height", v)
}
func (ral *RectAreaLight) Id() int {
	return ral.Get("id").Int()
}
func (ral *RectAreaLight) SetId(v int) {
	ral.Set("id", v)
}
func (ral *RectAreaLight) Intensity() int {
	return ral.Get("intensity").Int()
}
func (ral *RectAreaLight) SetIntensity(v int) {
	ral.Set("intensity", v)
}
func (ral *RectAreaLight) IsObject3D() true {
	return true(ral.Get("isObject3D"))
}
func (ral *RectAreaLight) SetIsObject3D(v true) {
	ral.Set("isObject3D", v)
}
func (ral *RectAreaLight) Layers() *core.Layers {
	return &core.Layers{Value: ral.Get("layers")}
}
func (ral *RectAreaLight) SetLayers(v *core.Layers) {
	ral.Set("layers", v)
}
func (ral *RectAreaLight) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ral.Get("matrix")}
}
func (ral *RectAreaLight) SetMatrix(v *math.Matrix4) {
	ral.Set("matrix", v)
}
func (ral *RectAreaLight) MatrixAutoUpdate() bool {
	return ral.Get("matrixAutoUpdate").Bool()
}
func (ral *RectAreaLight) SetMatrixAutoUpdate(v bool) {
	ral.Set("matrixAutoUpdate", v)
}
func (ral *RectAreaLight) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ral.Get("matrixWorld")}
}
func (ral *RectAreaLight) SetMatrixWorld(v *math.Matrix4) {
	ral.Set("matrixWorld", v)
}
func (ral *RectAreaLight) MatrixWorldNeedsUpdate() bool {
	return ral.Get("matrixWorldNeedsUpdate").Bool()
}
func (ral *RectAreaLight) SetMatrixWorldNeedsUpdate(v bool) {
	ral.Set("matrixWorldNeedsUpdate", v)
}
func (ral *RectAreaLight) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ral.Get("modelViewMatrix")}
}
func (ral *RectAreaLight) SetModelViewMatrix(v *math.Matrix4) {
	ral.Set("modelViewMatrix", v)
}
func (ral *RectAreaLight) Name() string {
	return ral.Get("name").String()
}
func (ral *RectAreaLight) SetName(v string) {
	ral.Set("name", v)
}
func (ral *RectAreaLight) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ral.Get("normalMatrix")}
}
func (ral *RectAreaLight) SetNormalMatrix(v *math.Matrix3) {
	ral.Set("normalMatrix", v)
}
func (ral *RectAreaLight) OnAfterRender() js.Value {
	return ral.Get("onAfterRender")
}
func (ral *RectAreaLight) SetOnAfterRender(v js.Value) {
	ral.Set("onAfterRender", v)
}
func (ral *RectAreaLight) OnBeforeRender() js.Value {
	return ral.Get("onBeforeRender")
}
func (ral *RectAreaLight) SetOnBeforeRender(v js.Value) {
	ral.Set("onBeforeRender", v)
}
func (ral *RectAreaLight) Parent() core.Object3D {
	return core.Object3D(ral.Get("parent"))
}
func (ral *RectAreaLight) SetParent(v core.Object3D) {
	ral.Set("parent", v)
}
func (ral *RectAreaLight) Position() *math.Vector3 {
	return &math.Vector3{Value: ral.Get("position")}
}
func (ral *RectAreaLight) SetPosition(v *math.Vector3) {
	ral.Set("position", v)
}
func (ral *RectAreaLight) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ral.Get("quaternion")}
}
func (ral *RectAreaLight) SetQuaternion(v *math.Quaternion) {
	ral.Set("quaternion", v)
}
func (ral *RectAreaLight) ReceiveShadow() bool {
	return ral.Get("receiveShadow").Bool()
}
func (ral *RectAreaLight) SetReceiveShadow(v bool) {
	ral.Set("receiveShadow", v)
}
func (ral *RectAreaLight) RenderOrder() int {
	return ral.Get("renderOrder").Int()
}
func (ral *RectAreaLight) SetRenderOrder(v int) {
	ral.Set("renderOrder", v)
}
func (ral *RectAreaLight) Rotation() *math.Euler {
	return &math.Euler{Value: ral.Get("rotation")}
}
func (ral *RectAreaLight) SetRotation(v *math.Euler) {
	ral.Set("rotation", v)
}
func (ral *RectAreaLight) Scale() *math.Vector3 {
	return &math.Vector3{Value: ral.Get("scale")}
}
func (ral *RectAreaLight) SetScale(v *math.Vector3) {
	ral.Set("scale", v)
}
func (ral *RectAreaLight) Shadow() *LightShadow {
	return &LightShadow{Value: ral.Get("shadow")}
}
func (ral *RectAreaLight) SetShadow(v *LightShadow) {
	ral.Set("shadow", v)
}
func (ral *RectAreaLight) ShadowBias() js.Value {
	return ral.Get("shadowBias")
}
func (ral *RectAreaLight) SetShadowBias(v js.Value) {
	ral.Set("shadowBias", v)
}
func (ral *RectAreaLight) ShadowCameraBottom() js.Value {
	return ral.Get("shadowCameraBottom")
}
func (ral *RectAreaLight) SetShadowCameraBottom(v js.Value) {
	ral.Set("shadowCameraBottom", v)
}
func (ral *RectAreaLight) ShadowCameraFar() js.Value {
	return ral.Get("shadowCameraFar")
}
func (ral *RectAreaLight) SetShadowCameraFar(v js.Value) {
	ral.Set("shadowCameraFar", v)
}
func (ral *RectAreaLight) ShadowCameraFov() js.Value {
	return ral.Get("shadowCameraFov")
}
func (ral *RectAreaLight) SetShadowCameraFov(v js.Value) {
	ral.Set("shadowCameraFov", v)
}
func (ral *RectAreaLight) ShadowCameraLeft() js.Value {
	return ral.Get("shadowCameraLeft")
}
func (ral *RectAreaLight) SetShadowCameraLeft(v js.Value) {
	ral.Set("shadowCameraLeft", v)
}
func (ral *RectAreaLight) ShadowCameraNear() js.Value {
	return ral.Get("shadowCameraNear")
}
func (ral *RectAreaLight) SetShadowCameraNear(v js.Value) {
	ral.Set("shadowCameraNear", v)
}
func (ral *RectAreaLight) ShadowCameraRight() js.Value {
	return ral.Get("shadowCameraRight")
}
func (ral *RectAreaLight) SetShadowCameraRight(v js.Value) {
	ral.Set("shadowCameraRight", v)
}
func (ral *RectAreaLight) ShadowCameraTop() js.Value {
	return ral.Get("shadowCameraTop")
}
func (ral *RectAreaLight) SetShadowCameraTop(v js.Value) {
	ral.Set("shadowCameraTop", v)
}
func (ral *RectAreaLight) ShadowMapHeight() js.Value {
	return ral.Get("shadowMapHeight")
}
func (ral *RectAreaLight) SetShadowMapHeight(v js.Value) {
	ral.Set("shadowMapHeight", v)
}
func (ral *RectAreaLight) ShadowMapWidth() js.Value {
	return ral.Get("shadowMapWidth")
}
func (ral *RectAreaLight) SetShadowMapWidth(v js.Value) {
	ral.Set("shadowMapWidth", v)
}
func (ral *RectAreaLight) Type() string {
	return ral.Get("type").String()
}
func (ral *RectAreaLight) SetType(v string) {
	ral.Set("type", v)
}
func (ral *RectAreaLight) Up() *math.Vector3 {
	return &math.Vector3{Value: ral.Get("up")}
}
func (ral *RectAreaLight) SetUp(v *math.Vector3) {
	ral.Set("up", v)
}
func (ral *RectAreaLight) UserData() js.Value {
	return ral.Get("userData")
}
func (ral *RectAreaLight) SetUserData(v js.Value) {
	ral.Set("userData", v)
}
func (ral *RectAreaLight) Uuid() string {
	return ral.Get("uuid").String()
}
func (ral *RectAreaLight) SetUuid(v string) {
	ral.Set("uuid", v)
}
func (ral *RectAreaLight) Visible() bool {
	return ral.Get("visible").Bool()
}
func (ral *RectAreaLight) SetVisible(v bool) {
	ral.Set("visible", v)
}
func (ral *RectAreaLight) Width() int {
	return ral.Get("width").Int()
}
func (ral *RectAreaLight) SetWidth(v int) {
	ral.Set("width", v)
}
func (ral *RectAreaLight) DefaultMatrixAutoUpdate() bool {
	return ral.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ral *RectAreaLight) SetDefaultMatrixAutoUpdate(v bool) {
	ral.Set("DefaultMatrixAutoUpdate", v)
}
func (ral *RectAreaLight) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ral.Get("DefaultUp")}
}
func (ral *RectAreaLight) SetDefaultUp(v *math.Vector3) {
	ral.Set("DefaultUp", v)
}
func (ral *RectAreaLight) Add(object []core.Object3D) this {
	return this(ral.Call("add", object))
}
func (ral *RectAreaLight) AddEventListener(typ string, listener js.Value) {
	ral.Call("addEventListener", typ, listener)
}
func (ral *RectAreaLight) ApplyMatrix(matrix *math.Matrix4) {
	ral.Call("applyMatrix", matrix)
}
func (ral *RectAreaLight) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ral.Call("applyQuaternion", quaternion))
}
func (ral *RectAreaLight) Clone(recursive bool) this {
	return this(ral.Call("clone", recursive))
}
func (ral *RectAreaLight) Copy(source *core.Object3D, recursive bool) this {
	return this(ral.Call("copy", source, recursive))
}
func (ral *RectAreaLight) DispatchEvent(event js.Value) {
	ral.Call("dispatchEvent", event)
}
func (ral *RectAreaLight) GetObjectById(id int) core.Object3D {
	return core.Object3D(ral.Call("getObjectById", id))
}
func (ral *RectAreaLight) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ral.Call("getObjectByName", name))
}
func (ral *RectAreaLight) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ral.Call("getObjectByProperty", name, value))
}
func (ral *RectAreaLight) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ral.Call("getWorldDirection", target)}
}
func (ral *RectAreaLight) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ral.Call("getWorldPosition", target)}
}
func (ral *RectAreaLight) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ral.Call("getWorldQuaternion", target)}
}
func (ral *RectAreaLight) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ral.Call("getWorldScale", target)}
}
func (ral *RectAreaLight) HasEventListener(typ string, listener js.Value) bool {
	return ral.Call("hasEventListener", typ, listener).Bool()
}
func (ral *RectAreaLight) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ral.Call("localToWorld", vector)}
}
func (ral *RectAreaLight) LookAt(vector math.Vector3, y int, z int) {
	ral.Call("lookAt", vector, y, z)
}
func (ral *RectAreaLight) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ral.Call("raycast", raycaster, intersects)
}
func (ral *RectAreaLight) Remove(object []core.Object3D) this {
	return this(ral.Call("remove", object))
}
func (ral *RectAreaLight) RemoveEventListener(typ string, listener js.Value) {
	ral.Call("removeEventListener", typ, listener)
}
func (ral *RectAreaLight) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ral.Call("rotateOnAxis", axis, angle))
}
func (ral *RectAreaLight) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ral.Call("rotateOnWorldAxis", axis, angle))
}
func (ral *RectAreaLight) RotateX(angle int) this {
	return this(ral.Call("rotateX", angle))
}
func (ral *RectAreaLight) RotateY(angle int) this {
	return this(ral.Call("rotateY", angle))
}
func (ral *RectAreaLight) RotateZ(angle int) this {
	return this(ral.Call("rotateZ", angle))
}
func (ral *RectAreaLight) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ral.Call("setRotationFromAxisAngle", axis, angle)
}
func (ral *RectAreaLight) SetRotationFromEuler(euler *math.Euler) {
	ral.Call("setRotationFromEuler", euler)
}
func (ral *RectAreaLight) SetRotationFromMatrix(m *math.Matrix4) {
	ral.Call("setRotationFromMatrix", m)
}
func (ral *RectAreaLight) SetRotationFromQuaternion(q *math.Quaternion) {
	ral.Call("setRotationFromQuaternion", q)
}
func (ral *RectAreaLight) ToJSON(meta js.Value) js.Value {
	return ral.Call("toJSON", meta)
}
func (ral *RectAreaLight) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ral.Call("translateOnAxis", axis, distance))
}
func (ral *RectAreaLight) TranslateX(distance int) this {
	return this(ral.Call("translateX", distance))
}
func (ral *RectAreaLight) TranslateY(distance int) this {
	return this(ral.Call("translateY", distance))
}
func (ral *RectAreaLight) TranslateZ(distance int) this {
	return this(ral.Call("translateZ", distance))
}
func (ral *RectAreaLight) Traverse(callback js.Value) {
	ral.Call("traverse", callback)
}
func (ral *RectAreaLight) TraverseAncestors(callback js.Value) {
	ral.Call("traverseAncestors", callback)
}
func (ral *RectAreaLight) TraverseVisible(callback js.Value) {
	ral.Call("traverseVisible", callback)
}
func (ral *RectAreaLight) UpdateMatrix() {
	ral.Call("updateMatrix")
}
func (ral *RectAreaLight) UpdateMatrixWorld(force bool) {
	ral.Call("updateMatrixWorld", force)
}
func (ral *RectAreaLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ral.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ral *RectAreaLight) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ral.Call("worldToLocal", vector)}
}

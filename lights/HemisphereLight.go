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

type HemisphereLight struct {
	js.Value
}

func NewHemisphereLight(skyColor math.Color, groundColor math.Color, intensity int) *HemisphereLight {
	return &HemisphereLight{Value: get("HemisphereLight").New(skyColor, groundColor, intensity)}
}
func (hl *HemisphereLight) CastShadow() bool {
	return hl.Get("castShadow").Bool()
}
func (hl *HemisphereLight) SetCastShadow(v bool) {
	hl.Set("castShadow", v)
}
func (hl *HemisphereLight) Children() []core.Object3D {
	return []core.Object3D(hl.Get("children"))
}
func (hl *HemisphereLight) SetChildren(v []core.Object3D) {
	hl.Set("children", v)
}
func (hl *HemisphereLight) Color() *math.Color {
	return &math.Color{Value: hl.Get("color")}
}
func (hl *HemisphereLight) SetColor(v *math.Color) {
	hl.Set("color", v)
}
func (hl *HemisphereLight) FrustumCulled() bool {
	return hl.Get("frustumCulled").Bool()
}
func (hl *HemisphereLight) SetFrustumCulled(v bool) {
	hl.Set("frustumCulled", v)
}
func (hl *HemisphereLight) GroundColor() *math.Color {
	return &math.Color{Value: hl.Get("groundColor")}
}
func (hl *HemisphereLight) SetGroundColor(v *math.Color) {
	hl.Set("groundColor", v)
}
func (hl *HemisphereLight) Id() int {
	return hl.Get("id").Int()
}
func (hl *HemisphereLight) SetId(v int) {
	hl.Set("id", v)
}
func (hl *HemisphereLight) Intensity() int {
	return hl.Get("intensity").Int()
}
func (hl *HemisphereLight) SetIntensity(v int) {
	hl.Set("intensity", v)
}
func (hl *HemisphereLight) IsObject3D() true {
	return true(hl.Get("isObject3D"))
}
func (hl *HemisphereLight) SetIsObject3D(v true) {
	hl.Set("isObject3D", v)
}
func (hl *HemisphereLight) Layers() *core.Layers {
	return &core.Layers{Value: hl.Get("layers")}
}
func (hl *HemisphereLight) SetLayers(v *core.Layers) {
	hl.Set("layers", v)
}
func (hl *HemisphereLight) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: hl.Get("matrix")}
}
func (hl *HemisphereLight) SetMatrix(v *math.Matrix4) {
	hl.Set("matrix", v)
}
func (hl *HemisphereLight) MatrixAutoUpdate() bool {
	return hl.Get("matrixAutoUpdate").Bool()
}
func (hl *HemisphereLight) SetMatrixAutoUpdate(v bool) {
	hl.Set("matrixAutoUpdate", v)
}
func (hl *HemisphereLight) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: hl.Get("matrixWorld")}
}
func (hl *HemisphereLight) SetMatrixWorld(v *math.Matrix4) {
	hl.Set("matrixWorld", v)
}
func (hl *HemisphereLight) MatrixWorldNeedsUpdate() bool {
	return hl.Get("matrixWorldNeedsUpdate").Bool()
}
func (hl *HemisphereLight) SetMatrixWorldNeedsUpdate(v bool) {
	hl.Set("matrixWorldNeedsUpdate", v)
}
func (hl *HemisphereLight) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: hl.Get("modelViewMatrix")}
}
func (hl *HemisphereLight) SetModelViewMatrix(v *math.Matrix4) {
	hl.Set("modelViewMatrix", v)
}
func (hl *HemisphereLight) Name() string {
	return hl.Get("name").String()
}
func (hl *HemisphereLight) SetName(v string) {
	hl.Set("name", v)
}
func (hl *HemisphereLight) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: hl.Get("normalMatrix")}
}
func (hl *HemisphereLight) SetNormalMatrix(v *math.Matrix3) {
	hl.Set("normalMatrix", v)
}
func (hl *HemisphereLight) OnAfterRender() js.Value {
	return hl.Get("onAfterRender")
}
func (hl *HemisphereLight) SetOnAfterRender(v js.Value) {
	hl.Set("onAfterRender", v)
}
func (hl *HemisphereLight) OnBeforeRender() js.Value {
	return hl.Get("onBeforeRender")
}
func (hl *HemisphereLight) SetOnBeforeRender(v js.Value) {
	hl.Set("onBeforeRender", v)
}
func (hl *HemisphereLight) Parent() core.Object3D {
	return core.Object3D(hl.Get("parent"))
}
func (hl *HemisphereLight) SetParent(v core.Object3D) {
	hl.Set("parent", v)
}
func (hl *HemisphereLight) Position() *math.Vector3 {
	return &math.Vector3{Value: hl.Get("position")}
}
func (hl *HemisphereLight) SetPosition(v *math.Vector3) {
	hl.Set("position", v)
}
func (hl *HemisphereLight) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: hl.Get("quaternion")}
}
func (hl *HemisphereLight) SetQuaternion(v *math.Quaternion) {
	hl.Set("quaternion", v)
}
func (hl *HemisphereLight) ReceiveShadow() bool {
	return hl.Get("receiveShadow").Bool()
}
func (hl *HemisphereLight) SetReceiveShadow(v bool) {
	hl.Set("receiveShadow", v)
}
func (hl *HemisphereLight) RenderOrder() int {
	return hl.Get("renderOrder").Int()
}
func (hl *HemisphereLight) SetRenderOrder(v int) {
	hl.Set("renderOrder", v)
}
func (hl *HemisphereLight) Rotation() *math.Euler {
	return &math.Euler{Value: hl.Get("rotation")}
}
func (hl *HemisphereLight) SetRotation(v *math.Euler) {
	hl.Set("rotation", v)
}
func (hl *HemisphereLight) Scale() *math.Vector3 {
	return &math.Vector3{Value: hl.Get("scale")}
}
func (hl *HemisphereLight) SetScale(v *math.Vector3) {
	hl.Set("scale", v)
}
func (hl *HemisphereLight) Shadow() *LightShadow {
	return &LightShadow{Value: hl.Get("shadow")}
}
func (hl *HemisphereLight) SetShadow(v *LightShadow) {
	hl.Set("shadow", v)
}
func (hl *HemisphereLight) ShadowBias() js.Value {
	return hl.Get("shadowBias")
}
func (hl *HemisphereLight) SetShadowBias(v js.Value) {
	hl.Set("shadowBias", v)
}
func (hl *HemisphereLight) ShadowCameraBottom() js.Value {
	return hl.Get("shadowCameraBottom")
}
func (hl *HemisphereLight) SetShadowCameraBottom(v js.Value) {
	hl.Set("shadowCameraBottom", v)
}
func (hl *HemisphereLight) ShadowCameraFar() js.Value {
	return hl.Get("shadowCameraFar")
}
func (hl *HemisphereLight) SetShadowCameraFar(v js.Value) {
	hl.Set("shadowCameraFar", v)
}
func (hl *HemisphereLight) ShadowCameraFov() js.Value {
	return hl.Get("shadowCameraFov")
}
func (hl *HemisphereLight) SetShadowCameraFov(v js.Value) {
	hl.Set("shadowCameraFov", v)
}
func (hl *HemisphereLight) ShadowCameraLeft() js.Value {
	return hl.Get("shadowCameraLeft")
}
func (hl *HemisphereLight) SetShadowCameraLeft(v js.Value) {
	hl.Set("shadowCameraLeft", v)
}
func (hl *HemisphereLight) ShadowCameraNear() js.Value {
	return hl.Get("shadowCameraNear")
}
func (hl *HemisphereLight) SetShadowCameraNear(v js.Value) {
	hl.Set("shadowCameraNear", v)
}
func (hl *HemisphereLight) ShadowCameraRight() js.Value {
	return hl.Get("shadowCameraRight")
}
func (hl *HemisphereLight) SetShadowCameraRight(v js.Value) {
	hl.Set("shadowCameraRight", v)
}
func (hl *HemisphereLight) ShadowCameraTop() js.Value {
	return hl.Get("shadowCameraTop")
}
func (hl *HemisphereLight) SetShadowCameraTop(v js.Value) {
	hl.Set("shadowCameraTop", v)
}
func (hl *HemisphereLight) ShadowMapHeight() js.Value {
	return hl.Get("shadowMapHeight")
}
func (hl *HemisphereLight) SetShadowMapHeight(v js.Value) {
	hl.Set("shadowMapHeight", v)
}
func (hl *HemisphereLight) ShadowMapWidth() js.Value {
	return hl.Get("shadowMapWidth")
}
func (hl *HemisphereLight) SetShadowMapWidth(v js.Value) {
	hl.Set("shadowMapWidth", v)
}
func (hl *HemisphereLight) SkyColor() *math.Color {
	return &math.Color{Value: hl.Get("skyColor")}
}
func (hl *HemisphereLight) SetSkyColor(v *math.Color) {
	hl.Set("skyColor", v)
}
func (hl *HemisphereLight) Type() string {
	return hl.Get("type").String()
}
func (hl *HemisphereLight) SetType(v string) {
	hl.Set("type", v)
}
func (hl *HemisphereLight) Up() *math.Vector3 {
	return &math.Vector3{Value: hl.Get("up")}
}
func (hl *HemisphereLight) SetUp(v *math.Vector3) {
	hl.Set("up", v)
}
func (hl *HemisphereLight) UserData() js.Value {
	return hl.Get("userData")
}
func (hl *HemisphereLight) SetUserData(v js.Value) {
	hl.Set("userData", v)
}
func (hl *HemisphereLight) Uuid() string {
	return hl.Get("uuid").String()
}
func (hl *HemisphereLight) SetUuid(v string) {
	hl.Set("uuid", v)
}
func (hl *HemisphereLight) Visible() bool {
	return hl.Get("visible").Bool()
}
func (hl *HemisphereLight) SetVisible(v bool) {
	hl.Set("visible", v)
}
func (hl *HemisphereLight) DefaultMatrixAutoUpdate() bool {
	return hl.Get("DefaultMatrixAutoUpdate").Bool()
}
func (hl *HemisphereLight) SetDefaultMatrixAutoUpdate(v bool) {
	hl.Set("DefaultMatrixAutoUpdate", v)
}
func (hl *HemisphereLight) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: hl.Get("DefaultUp")}
}
func (hl *HemisphereLight) SetDefaultUp(v *math.Vector3) {
	hl.Set("DefaultUp", v)
}
func (hl *HemisphereLight) Add(object []core.Object3D) this {
	return this(hl.Call("add", object))
}
func (hl *HemisphereLight) AddEventListener(typ string, listener js.Value) {
	hl.Call("addEventListener", typ, listener)
}
func (hl *HemisphereLight) ApplyMatrix(matrix *math.Matrix4) {
	hl.Call("applyMatrix", matrix)
}
func (hl *HemisphereLight) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(hl.Call("applyQuaternion", quaternion))
}
func (hl *HemisphereLight) Clone(recursive bool) this {
	return this(hl.Call("clone", recursive))
}
func (hl *HemisphereLight) Copy(source *core.Object3D, recursive bool) this {
	return this(hl.Call("copy", source, recursive))
}
func (hl *HemisphereLight) DispatchEvent(event js.Value) {
	hl.Call("dispatchEvent", event)
}
func (hl *HemisphereLight) GetObjectById(id int) core.Object3D {
	return core.Object3D(hl.Call("getObjectById", id))
}
func (hl *HemisphereLight) GetObjectByName(name string) core.Object3D {
	return core.Object3D(hl.Call("getObjectByName", name))
}
func (hl *HemisphereLight) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(hl.Call("getObjectByProperty", name, value))
}
func (hl *HemisphereLight) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hl.Call("getWorldDirection", target)}
}
func (hl *HemisphereLight) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hl.Call("getWorldPosition", target)}
}
func (hl *HemisphereLight) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: hl.Call("getWorldQuaternion", target)}
}
func (hl *HemisphereLight) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hl.Call("getWorldScale", target)}
}
func (hl *HemisphereLight) HasEventListener(typ string, listener js.Value) bool {
	return hl.Call("hasEventListener", typ, listener).Bool()
}
func (hl *HemisphereLight) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hl.Call("localToWorld", vector)}
}
func (hl *HemisphereLight) LookAt(vector math.Vector3, y int, z int) {
	hl.Call("lookAt", vector, y, z)
}
func (hl *HemisphereLight) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	hl.Call("raycast", raycaster, intersects)
}
func (hl *HemisphereLight) Remove(object []core.Object3D) this {
	return this(hl.Call("remove", object))
}
func (hl *HemisphereLight) RemoveEventListener(typ string, listener js.Value) {
	hl.Call("removeEventListener", typ, listener)
}
func (hl *HemisphereLight) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(hl.Call("rotateOnAxis", axis, angle))
}
func (hl *HemisphereLight) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(hl.Call("rotateOnWorldAxis", axis, angle))
}
func (hl *HemisphereLight) RotateX(angle int) this {
	return this(hl.Call("rotateX", angle))
}
func (hl *HemisphereLight) RotateY(angle int) this {
	return this(hl.Call("rotateY", angle))
}
func (hl *HemisphereLight) RotateZ(angle int) this {
	return this(hl.Call("rotateZ", angle))
}
func (hl *HemisphereLight) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	hl.Call("setRotationFromAxisAngle", axis, angle)
}
func (hl *HemisphereLight) SetRotationFromEuler(euler *math.Euler) {
	hl.Call("setRotationFromEuler", euler)
}
func (hl *HemisphereLight) SetRotationFromMatrix(m *math.Matrix4) {
	hl.Call("setRotationFromMatrix", m)
}
func (hl *HemisphereLight) SetRotationFromQuaternion(q *math.Quaternion) {
	hl.Call("setRotationFromQuaternion", q)
}
func (hl *HemisphereLight) ToJSON(meta js.Value) js.Value {
	return hl.Call("toJSON", meta)
}
func (hl *HemisphereLight) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(hl.Call("translateOnAxis", axis, distance))
}
func (hl *HemisphereLight) TranslateX(distance int) this {
	return this(hl.Call("translateX", distance))
}
func (hl *HemisphereLight) TranslateY(distance int) this {
	return this(hl.Call("translateY", distance))
}
func (hl *HemisphereLight) TranslateZ(distance int) this {
	return this(hl.Call("translateZ", distance))
}
func (hl *HemisphereLight) Traverse(callback js.Value) {
	hl.Call("traverse", callback)
}
func (hl *HemisphereLight) TraverseAncestors(callback js.Value) {
	hl.Call("traverseAncestors", callback)
}
func (hl *HemisphereLight) TraverseVisible(callback js.Value) {
	hl.Call("traverseVisible", callback)
}
func (hl *HemisphereLight) UpdateMatrix() {
	hl.Call("updateMatrix")
}
func (hl *HemisphereLight) UpdateMatrixWorld(force bool) {
	hl.Call("updateMatrixWorld", force)
}
func (hl *HemisphereLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	hl.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (hl *HemisphereLight) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: hl.Call("worldToLocal", vector)}
}

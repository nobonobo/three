package lights

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/renderers/webgl"
	"github.com/nobonobo/three/scenes"
)

type PointLight struct {
	js.Value
}

func NewPointLight(color math.Color, intensity int, distance int, decay int) *PointLight {
	return &PointLight{Value: get("PointLight").New(color, intensity, distance, decay)}
}
func (pl *PointLight) CastShadow() bool {
	return pl.Get("castShadow").Bool()
}
func (pl *PointLight) SetCastShadow(v bool) {
	pl.Set("castShadow", v)
}
func (pl *PointLight) Children() []core.Object3D {
	return []core.Object3D(pl.Get("children"))
}
func (pl *PointLight) SetChildren(v []core.Object3D) {
	pl.Set("children", v)
}
func (pl *PointLight) Color() *math.Color {
	return &math.Color{Value: pl.Get("color")}
}
func (pl *PointLight) SetColor(v *math.Color) {
	pl.Set("color", v)
}
func (pl *PointLight) Decay() int {
	return pl.Get("decay").Int()
}
func (pl *PointLight) SetDecay(v int) {
	pl.Set("decay", v)
}
func (pl *PointLight) Distance() int {
	return pl.Get("distance").Int()
}
func (pl *PointLight) SetDistance(v int) {
	pl.Set("distance", v)
}
func (pl *PointLight) FrustumCulled() bool {
	return pl.Get("frustumCulled").Bool()
}
func (pl *PointLight) SetFrustumCulled(v bool) {
	pl.Set("frustumCulled", v)
}
func (pl *PointLight) Id() int {
	return pl.Get("id").Int()
}
func (pl *PointLight) SetId(v int) {
	pl.Set("id", v)
}
func (pl *PointLight) Intensity() int {
	return pl.Get("intensity").Int()
}
func (pl *PointLight) SetIntensity(v int) {
	pl.Set("intensity", v)
}
func (pl *PointLight) IsObject3D() true {
	return true(pl.Get("isObject3D"))
}
func (pl *PointLight) SetIsObject3D(v true) {
	pl.Set("isObject3D", v)
}
func (pl *PointLight) Layers() *core.Layers {
	return &core.Layers{Value: pl.Get("layers")}
}
func (pl *PointLight) SetLayers(v *core.Layers) {
	pl.Set("layers", v)
}
func (pl *PointLight) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: pl.Get("matrix")}
}
func (pl *PointLight) SetMatrix(v *math.Matrix4) {
	pl.Set("matrix", v)
}
func (pl *PointLight) MatrixAutoUpdate() bool {
	return pl.Get("matrixAutoUpdate").Bool()
}
func (pl *PointLight) SetMatrixAutoUpdate(v bool) {
	pl.Set("matrixAutoUpdate", v)
}
func (pl *PointLight) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: pl.Get("matrixWorld")}
}
func (pl *PointLight) SetMatrixWorld(v *math.Matrix4) {
	pl.Set("matrixWorld", v)
}
func (pl *PointLight) MatrixWorldNeedsUpdate() bool {
	return pl.Get("matrixWorldNeedsUpdate").Bool()
}
func (pl *PointLight) SetMatrixWorldNeedsUpdate(v bool) {
	pl.Set("matrixWorldNeedsUpdate", v)
}
func (pl *PointLight) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: pl.Get("modelViewMatrix")}
}
func (pl *PointLight) SetModelViewMatrix(v *math.Matrix4) {
	pl.Set("modelViewMatrix", v)
}
func (pl *PointLight) Name() string {
	return pl.Get("name").String()
}
func (pl *PointLight) SetName(v string) {
	pl.Set("name", v)
}
func (pl *PointLight) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: pl.Get("normalMatrix")}
}
func (pl *PointLight) SetNormalMatrix(v *math.Matrix3) {
	pl.Set("normalMatrix", v)
}
func (pl *PointLight) OnAfterRender() js.Value {
	return pl.Get("onAfterRender")
}
func (pl *PointLight) SetOnAfterRender(v js.Value) {
	pl.Set("onAfterRender", v)
}
func (pl *PointLight) OnBeforeRender() js.Value {
	return pl.Get("onBeforeRender")
}
func (pl *PointLight) SetOnBeforeRender(v js.Value) {
	pl.Set("onBeforeRender", v)
}
func (pl *PointLight) Parent() core.Object3D {
	return core.Object3D(pl.Get("parent"))
}
func (pl *PointLight) SetParent(v core.Object3D) {
	pl.Set("parent", v)
}
func (pl *PointLight) Position() *math.Vector3 {
	return &math.Vector3{Value: pl.Get("position")}
}
func (pl *PointLight) SetPosition(v *math.Vector3) {
	pl.Set("position", v)
}
func (pl *PointLight) Power() int {
	return pl.Get("power").Int()
}
func (pl *PointLight) SetPower(v int) {
	pl.Set("power", v)
}
func (pl *PointLight) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: pl.Get("quaternion")}
}
func (pl *PointLight) SetQuaternion(v *math.Quaternion) {
	pl.Set("quaternion", v)
}
func (pl *PointLight) ReceiveShadow() bool {
	return pl.Get("receiveShadow").Bool()
}
func (pl *PointLight) SetReceiveShadow(v bool) {
	pl.Set("receiveShadow", v)
}
func (pl *PointLight) RenderOrder() int {
	return pl.Get("renderOrder").Int()
}
func (pl *PointLight) SetRenderOrder(v int) {
	pl.Set("renderOrder", v)
}
func (pl *PointLight) Rotation() *math.Euler {
	return &math.Euler{Value: pl.Get("rotation")}
}
func (pl *PointLight) SetRotation(v *math.Euler) {
	pl.Set("rotation", v)
}
func (pl *PointLight) Scale() *math.Vector3 {
	return &math.Vector3{Value: pl.Get("scale")}
}
func (pl *PointLight) SetScale(v *math.Vector3) {
	pl.Set("scale", v)
}
func (pl *PointLight) Shadow() *PointLightShadow {
	return &PointLightShadow{Value: pl.Get("shadow")}
}
func (pl *PointLight) SetShadow(v *PointLightShadow) {
	pl.Set("shadow", v)
}
func (pl *PointLight) ShadowBias() js.Value {
	return pl.Get("shadowBias")
}
func (pl *PointLight) SetShadowBias(v js.Value) {
	pl.Set("shadowBias", v)
}
func (pl *PointLight) ShadowCameraBottom() js.Value {
	return pl.Get("shadowCameraBottom")
}
func (pl *PointLight) SetShadowCameraBottom(v js.Value) {
	pl.Set("shadowCameraBottom", v)
}
func (pl *PointLight) ShadowCameraFar() js.Value {
	return pl.Get("shadowCameraFar")
}
func (pl *PointLight) SetShadowCameraFar(v js.Value) {
	pl.Set("shadowCameraFar", v)
}
func (pl *PointLight) ShadowCameraFov() js.Value {
	return pl.Get("shadowCameraFov")
}
func (pl *PointLight) SetShadowCameraFov(v js.Value) {
	pl.Set("shadowCameraFov", v)
}
func (pl *PointLight) ShadowCameraLeft() js.Value {
	return pl.Get("shadowCameraLeft")
}
func (pl *PointLight) SetShadowCameraLeft(v js.Value) {
	pl.Set("shadowCameraLeft", v)
}
func (pl *PointLight) ShadowCameraNear() js.Value {
	return pl.Get("shadowCameraNear")
}
func (pl *PointLight) SetShadowCameraNear(v js.Value) {
	pl.Set("shadowCameraNear", v)
}
func (pl *PointLight) ShadowCameraRight() js.Value {
	return pl.Get("shadowCameraRight")
}
func (pl *PointLight) SetShadowCameraRight(v js.Value) {
	pl.Set("shadowCameraRight", v)
}
func (pl *PointLight) ShadowCameraTop() js.Value {
	return pl.Get("shadowCameraTop")
}
func (pl *PointLight) SetShadowCameraTop(v js.Value) {
	pl.Set("shadowCameraTop", v)
}
func (pl *PointLight) ShadowMapHeight() js.Value {
	return pl.Get("shadowMapHeight")
}
func (pl *PointLight) SetShadowMapHeight(v js.Value) {
	pl.Set("shadowMapHeight", v)
}
func (pl *PointLight) ShadowMapWidth() js.Value {
	return pl.Get("shadowMapWidth")
}
func (pl *PointLight) SetShadowMapWidth(v js.Value) {
	pl.Set("shadowMapWidth", v)
}
func (pl *PointLight) Type() string {
	return pl.Get("type").String()
}
func (pl *PointLight) SetType(v string) {
	pl.Set("type", v)
}
func (pl *PointLight) Up() *math.Vector3 {
	return &math.Vector3{Value: pl.Get("up")}
}
func (pl *PointLight) SetUp(v *math.Vector3) {
	pl.Set("up", v)
}
func (pl *PointLight) UserData() js.Value {
	return pl.Get("userData")
}
func (pl *PointLight) SetUserData(v js.Value) {
	pl.Set("userData", v)
}
func (pl *PointLight) Uuid() string {
	return pl.Get("uuid").String()
}
func (pl *PointLight) SetUuid(v string) {
	pl.Set("uuid", v)
}
func (pl *PointLight) Visible() bool {
	return pl.Get("visible").Bool()
}
func (pl *PointLight) SetVisible(v bool) {
	pl.Set("visible", v)
}
func (pl *PointLight) DefaultMatrixAutoUpdate() bool {
	return pl.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pl *PointLight) SetDefaultMatrixAutoUpdate(v bool) {
	pl.Set("DefaultMatrixAutoUpdate", v)
}
func (pl *PointLight) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: pl.Get("DefaultUp")}
}
func (pl *PointLight) SetDefaultUp(v *math.Vector3) {
	pl.Set("DefaultUp", v)
}
func (pl *PointLight) Add(object []core.Object3D) this {
	return this(pl.Call("add", object))
}
func (pl *PointLight) AddEventListener(typ string, listener js.Value) {
	pl.Call("addEventListener", typ, listener)
}
func (pl *PointLight) ApplyMatrix(matrix *math.Matrix4) {
	pl.Call("applyMatrix", matrix)
}
func (pl *PointLight) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(pl.Call("applyQuaternion", quaternion))
}
func (pl *PointLight) Clone(recursive bool) this {
	return this(pl.Call("clone", recursive))
}
func (pl *PointLight) Copy(source *core.Object3D, recursive bool) this {
	return this(pl.Call("copy", source, recursive))
}
func (pl *PointLight) DispatchEvent(event js.Value) {
	pl.Call("dispatchEvent", event)
}
func (pl *PointLight) GetObjectById(id int) core.Object3D {
	return core.Object3D(pl.Call("getObjectById", id))
}
func (pl *PointLight) GetObjectByName(name string) core.Object3D {
	return core.Object3D(pl.Call("getObjectByName", name))
}
func (pl *PointLight) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(pl.Call("getObjectByProperty", name, value))
}
func (pl *PointLight) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pl.Call("getWorldDirection", target)}
}
func (pl *PointLight) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pl.Call("getWorldPosition", target)}
}
func (pl *PointLight) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: pl.Call("getWorldQuaternion", target)}
}
func (pl *PointLight) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pl.Call("getWorldScale", target)}
}
func (pl *PointLight) HasEventListener(typ string, listener js.Value) bool {
	return pl.Call("hasEventListener", typ, listener).Bool()
}
func (pl *PointLight) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pl.Call("localToWorld", vector)}
}
func (pl *PointLight) LookAt(vector math.Vector3, y int, z int) {
	pl.Call("lookAt", vector, y, z)
}
func (pl *PointLight) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	pl.Call("raycast", raycaster, intersects)
}
func (pl *PointLight) Remove(object []core.Object3D) this {
	return this(pl.Call("remove", object))
}
func (pl *PointLight) RemoveEventListener(typ string, listener js.Value) {
	pl.Call("removeEventListener", typ, listener)
}
func (pl *PointLight) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(pl.Call("rotateOnAxis", axis, angle))
}
func (pl *PointLight) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(pl.Call("rotateOnWorldAxis", axis, angle))
}
func (pl *PointLight) RotateX(angle int) this {
	return this(pl.Call("rotateX", angle))
}
func (pl *PointLight) RotateY(angle int) this {
	return this(pl.Call("rotateY", angle))
}
func (pl *PointLight) RotateZ(angle int) this {
	return this(pl.Call("rotateZ", angle))
}
func (pl *PointLight) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	pl.Call("setRotationFromAxisAngle", axis, angle)
}
func (pl *PointLight) SetRotationFromEuler(euler *math.Euler) {
	pl.Call("setRotationFromEuler", euler)
}
func (pl *PointLight) SetRotationFromMatrix(m *math.Matrix4) {
	pl.Call("setRotationFromMatrix", m)
}
func (pl *PointLight) SetRotationFromQuaternion(q *math.Quaternion) {
	pl.Call("setRotationFromQuaternion", q)
}
func (pl *PointLight) ToJSON(meta js.Value) js.Value {
	return pl.Call("toJSON", meta)
}
func (pl *PointLight) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(pl.Call("translateOnAxis", axis, distance))
}
func (pl *PointLight) TranslateX(distance int) this {
	return this(pl.Call("translateX", distance))
}
func (pl *PointLight) TranslateY(distance int) this {
	return this(pl.Call("translateY", distance))
}
func (pl *PointLight) TranslateZ(distance int) this {
	return this(pl.Call("translateZ", distance))
}
func (pl *PointLight) Traverse(callback js.Value) {
	pl.Call("traverse", callback)
}
func (pl *PointLight) TraverseAncestors(callback js.Value) {
	pl.Call("traverseAncestors", callback)
}
func (pl *PointLight) TraverseVisible(callback js.Value) {
	pl.Call("traverseVisible", callback)
}
func (pl *PointLight) UpdateMatrix() {
	pl.Call("updateMatrix")
}
func (pl *PointLight) UpdateMatrixWorld(force bool) {
	pl.Call("updateMatrixWorld", force)
}
func (pl *PointLight) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pl.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pl *PointLight) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pl.Call("worldToLocal", vector)}
}

type PointLightShadow struct {
	js.Value
}

func NewPointLightShadow(camera *cameras.Camera) *PointLightShadow {
	return &PointLightShadow{Value: get("PointLightShadow").New(camera)}
}
func (pls *PointLightShadow) Bias() int {
	return pls.Get("bias").Int()
}
func (pls *PointLightShadow) SetBias(v int) {
	pls.Set("bias", v)
}
func (pls *PointLightShadow) Camera() *cameras.PerspectiveCamera {
	return &cameras.PerspectiveCamera{Value: pls.Get("camera")}
}
func (pls *PointLightShadow) SetCamera(v *cameras.PerspectiveCamera) {
	pls.Set("camera", v)
}
func (pls *PointLightShadow) Map() webgl.RenderTarget {
	return webgl.RenderTarget(pls.Get("map"))
}
func (pls *PointLightShadow) SetMap(v webgl.RenderTarget) {
	pls.Set("map", v)
}
func (pls *PointLightShadow) MapSize() *math.Vector2 {
	return &math.Vector2{Value: pls.Get("mapSize")}
}
func (pls *PointLightShadow) SetMapSize(v *math.Vector2) {
	pls.Set("mapSize", v)
}
func (pls *PointLightShadow) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: pls.Get("matrix")}
}
func (pls *PointLightShadow) SetMatrix(v *math.Matrix4) {
	pls.Set("matrix", v)
}
func (pls *PointLightShadow) Radius() int {
	return pls.Get("radius").Int()
}
func (pls *PointLightShadow) SetRadius(v int) {
	pls.Set("radius", v)
}
func (pls *PointLightShadow) Clone(recursive bool) this {
	return this(pls.Call("clone", recursive))
}
func (pls *PointLightShadow) Copy(source *LightShadow) this {
	return this(pls.Call("copy", source))
}
func (pls *PointLightShadow) ToJSON() js.Value {
	return pls.Call("toJSON")
}

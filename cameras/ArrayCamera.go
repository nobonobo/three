package cameras

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type ArrayCamera struct {
	js.Value
}

func NewArrayCamera(cameras []PerspectiveCamera) *ArrayCamera {
	return &ArrayCamera{Value: get("ArrayCamera").New(cameras)}
}
func (ac *ArrayCamera) Aspect() int {
	return ac.Get("aspect").Int()
}
func (ac *ArrayCamera) SetAspect(v int) {
	ac.Set("aspect", v)
}
func (ac *ArrayCamera) Cameras() []PerspectiveCamera {
	return []PerspectiveCamera(ac.Get("cameras"))
}
func (ac *ArrayCamera) SetCameras(v []PerspectiveCamera) {
	ac.Set("cameras", v)
}
func (ac *ArrayCamera) CastShadow() bool {
	return ac.Get("castShadow").Bool()
}
func (ac *ArrayCamera) SetCastShadow(v bool) {
	ac.Set("castShadow", v)
}
func (ac *ArrayCamera) Children() []core.Object3D {
	return []core.Object3D(ac.Get("children"))
}
func (ac *ArrayCamera) SetChildren(v []core.Object3D) {
	ac.Set("children", v)
}
func (ac *ArrayCamera) Far() int {
	return ac.Get("far").Int()
}
func (ac *ArrayCamera) SetFar(v int) {
	ac.Set("far", v)
}
func (ac *ArrayCamera) FilmGauge() int {
	return ac.Get("filmGauge").Int()
}
func (ac *ArrayCamera) SetFilmGauge(v int) {
	ac.Set("filmGauge", v)
}
func (ac *ArrayCamera) FilmOffset() int {
	return ac.Get("filmOffset").Int()
}
func (ac *ArrayCamera) SetFilmOffset(v int) {
	ac.Set("filmOffset", v)
}
func (ac *ArrayCamera) Focus() int {
	return ac.Get("focus").Int()
}
func (ac *ArrayCamera) SetFocus(v int) {
	ac.Set("focus", v)
}
func (ac *ArrayCamera) Fov() int {
	return ac.Get("fov").Int()
}
func (ac *ArrayCamera) SetFov(v int) {
	ac.Set("fov", v)
}
func (ac *ArrayCamera) FrustumCulled() bool {
	return ac.Get("frustumCulled").Bool()
}
func (ac *ArrayCamera) SetFrustumCulled(v bool) {
	ac.Set("frustumCulled", v)
}
func (ac *ArrayCamera) Id() int {
	return ac.Get("id").Int()
}
func (ac *ArrayCamera) SetId(v int) {
	ac.Set("id", v)
}
func (ac *ArrayCamera) IsArrayCamera() true {
	return true(ac.Get("isArrayCamera"))
}
func (ac *ArrayCamera) SetIsArrayCamera(v true) {
	ac.Set("isArrayCamera", v)
}
func (ac *ArrayCamera) IsCamera() true {
	return true(ac.Get("isCamera"))
}
func (ac *ArrayCamera) SetIsCamera(v true) {
	ac.Set("isCamera", v)
}
func (ac *ArrayCamera) IsObject3D() true {
	return true(ac.Get("isObject3D"))
}
func (ac *ArrayCamera) SetIsObject3D(v true) {
	ac.Set("isObject3D", v)
}
func (ac *ArrayCamera) IsPerspectiveCamera() true {
	return true(ac.Get("isPerspectiveCamera"))
}
func (ac *ArrayCamera) SetIsPerspectiveCamera(v true) {
	ac.Set("isPerspectiveCamera", v)
}
func (ac *ArrayCamera) Layers() *core.Layers {
	return &core.Layers{Value: ac.Get("layers")}
}
func (ac *ArrayCamera) SetLayers(v *core.Layers) {
	ac.Set("layers", v)
}
func (ac *ArrayCamera) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ac.Get("matrix")}
}
func (ac *ArrayCamera) SetMatrix(v *math.Matrix4) {
	ac.Set("matrix", v)
}
func (ac *ArrayCamera) MatrixAutoUpdate() bool {
	return ac.Get("matrixAutoUpdate").Bool()
}
func (ac *ArrayCamera) SetMatrixAutoUpdate(v bool) {
	ac.Set("matrixAutoUpdate", v)
}
func (ac *ArrayCamera) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ac.Get("matrixWorld")}
}
func (ac *ArrayCamera) SetMatrixWorld(v *math.Matrix4) {
	ac.Set("matrixWorld", v)
}
func (ac *ArrayCamera) MatrixWorldInverse() *math.Matrix4 {
	return &math.Matrix4{Value: ac.Get("matrixWorldInverse")}
}
func (ac *ArrayCamera) SetMatrixWorldInverse(v *math.Matrix4) {
	ac.Set("matrixWorldInverse", v)
}
func (ac *ArrayCamera) MatrixWorldNeedsUpdate() bool {
	return ac.Get("matrixWorldNeedsUpdate").Bool()
}
func (ac *ArrayCamera) SetMatrixWorldNeedsUpdate(v bool) {
	ac.Set("matrixWorldNeedsUpdate", v)
}
func (ac *ArrayCamera) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ac.Get("modelViewMatrix")}
}
func (ac *ArrayCamera) SetModelViewMatrix(v *math.Matrix4) {
	ac.Set("modelViewMatrix", v)
}
func (ac *ArrayCamera) Name() string {
	return ac.Get("name").String()
}
func (ac *ArrayCamera) SetName(v string) {
	ac.Set("name", v)
}
func (ac *ArrayCamera) Near() int {
	return ac.Get("near").Int()
}
func (ac *ArrayCamera) SetNear(v int) {
	ac.Set("near", v)
}
func (ac *ArrayCamera) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ac.Get("normalMatrix")}
}
func (ac *ArrayCamera) SetNormalMatrix(v *math.Matrix3) {
	ac.Set("normalMatrix", v)
}
func (ac *ArrayCamera) OnAfterRender() js.Value {
	return ac.Get("onAfterRender")
}
func (ac *ArrayCamera) SetOnAfterRender(v js.Value) {
	ac.Set("onAfterRender", v)
}
func (ac *ArrayCamera) OnBeforeRender() js.Value {
	return ac.Get("onBeforeRender")
}
func (ac *ArrayCamera) SetOnBeforeRender(v js.Value) {
	ac.Set("onBeforeRender", v)
}
func (ac *ArrayCamera) Parent() core.Object3D {
	return core.Object3D(ac.Get("parent"))
}
func (ac *ArrayCamera) SetParent(v core.Object3D) {
	ac.Set("parent", v)
}
func (ac *ArrayCamera) Position() *math.Vector3 {
	return &math.Vector3{Value: ac.Get("position")}
}
func (ac *ArrayCamera) SetPosition(v *math.Vector3) {
	ac.Set("position", v)
}
func (ac *ArrayCamera) ProjectionMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ac.Get("projectionMatrix")}
}
func (ac *ArrayCamera) SetProjectionMatrix(v *math.Matrix4) {
	ac.Set("projectionMatrix", v)
}
func (ac *ArrayCamera) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ac.Get("quaternion")}
}
func (ac *ArrayCamera) SetQuaternion(v *math.Quaternion) {
	ac.Set("quaternion", v)
}
func (ac *ArrayCamera) ReceiveShadow() bool {
	return ac.Get("receiveShadow").Bool()
}
func (ac *ArrayCamera) SetReceiveShadow(v bool) {
	ac.Set("receiveShadow", v)
}
func (ac *ArrayCamera) RenderOrder() int {
	return ac.Get("renderOrder").Int()
}
func (ac *ArrayCamera) SetRenderOrder(v int) {
	ac.Set("renderOrder", v)
}
func (ac *ArrayCamera) Rotation() *math.Euler {
	return &math.Euler{Value: ac.Get("rotation")}
}
func (ac *ArrayCamera) SetRotation(v *math.Euler) {
	ac.Set("rotation", v)
}
func (ac *ArrayCamera) Scale() *math.Vector3 {
	return &math.Vector3{Value: ac.Get("scale")}
}
func (ac *ArrayCamera) SetScale(v *math.Vector3) {
	ac.Set("scale", v)
}
func (ac *ArrayCamera) Type() string {
	return ac.Get("type").String()
}
func (ac *ArrayCamera) SetType(v string) {
	ac.Set("type", v)
}
func (ac *ArrayCamera) Up() *math.Vector3 {
	return &math.Vector3{Value: ac.Get("up")}
}
func (ac *ArrayCamera) SetUp(v *math.Vector3) {
	ac.Set("up", v)
}
func (ac *ArrayCamera) UserData() js.Value {
	return ac.Get("userData")
}
func (ac *ArrayCamera) SetUserData(v js.Value) {
	ac.Set("userData", v)
}
func (ac *ArrayCamera) Uuid() string {
	return ac.Get("uuid").String()
}
func (ac *ArrayCamera) SetUuid(v string) {
	ac.Set("uuid", v)
}
func (ac *ArrayCamera) View() null {
	return null(ac.Get("view"))
}
func (ac *ArrayCamera) SetView(v null) {
	ac.Set("view", v)
}
func (ac *ArrayCamera) Visible() bool {
	return ac.Get("visible").Bool()
}
func (ac *ArrayCamera) SetVisible(v bool) {
	ac.Set("visible", v)
}
func (ac *ArrayCamera) Zoom() int {
	return ac.Get("zoom").Int()
}
func (ac *ArrayCamera) SetZoom(v int) {
	ac.Set("zoom", v)
}
func (ac *ArrayCamera) DefaultMatrixAutoUpdate() bool {
	return ac.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ac *ArrayCamera) SetDefaultMatrixAutoUpdate(v bool) {
	ac.Set("DefaultMatrixAutoUpdate", v)
}
func (ac *ArrayCamera) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ac.Get("DefaultUp")}
}
func (ac *ArrayCamera) SetDefaultUp(v *math.Vector3) {
	ac.Set("DefaultUp", v)
}
func (ac *ArrayCamera) Add(object []core.Object3D) this {
	return this(ac.Call("add", object))
}
func (ac *ArrayCamera) AddEventListener(typ string, listener js.Value) {
	ac.Call("addEventListener", typ, listener)
}
func (ac *ArrayCamera) ApplyMatrix(matrix *math.Matrix4) {
	ac.Call("applyMatrix", matrix)
}
func (ac *ArrayCamera) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ac.Call("applyQuaternion", quaternion))
}
func (ac *ArrayCamera) ClearViewOffset() {
	ac.Call("clearViewOffset")
}
func (ac *ArrayCamera) Clone(recursive bool) this {
	return this(ac.Call("clone", recursive))
}
func (ac *ArrayCamera) Copy(source *Camera, recursive bool) this {
	return this(ac.Call("copy", source, recursive))
}
func (ac *ArrayCamera) DispatchEvent(event js.Value) {
	ac.Call("dispatchEvent", event)
}
func (ac *ArrayCamera) GetEffectiveFOV() int {
	return ac.Call("getEffectiveFOV").Int()
}
func (ac *ArrayCamera) GetFilmHeight() int {
	return ac.Call("getFilmHeight").Int()
}
func (ac *ArrayCamera) GetFilmWidth() int {
	return ac.Call("getFilmWidth").Int()
}
func (ac *ArrayCamera) GetFocalLength() int {
	return ac.Call("getFocalLength").Int()
}
func (ac *ArrayCamera) GetObjectById(id int) core.Object3D {
	return core.Object3D(ac.Call("getObjectById", id))
}
func (ac *ArrayCamera) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ac.Call("getObjectByName", name))
}
func (ac *ArrayCamera) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ac.Call("getObjectByProperty", name, value))
}
func (ac *ArrayCamera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ac.Call("getWorldDirection", target)}
}
func (ac *ArrayCamera) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ac.Call("getWorldPosition", target)}
}
func (ac *ArrayCamera) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ac.Call("getWorldQuaternion", target)}
}
func (ac *ArrayCamera) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ac.Call("getWorldScale", target)}
}
func (ac *ArrayCamera) HasEventListener(typ string, listener js.Value) bool {
	return ac.Call("hasEventListener", typ, listener).Bool()
}
func (ac *ArrayCamera) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ac.Call("localToWorld", vector)}
}
func (ac *ArrayCamera) LookAt(vector math.Vector3, y int, z int) {
	ac.Call("lookAt", vector, y, z)
}
func (ac *ArrayCamera) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ac.Call("raycast", raycaster, intersects)
}
func (ac *ArrayCamera) Remove(object []core.Object3D) this {
	return this(ac.Call("remove", object))
}
func (ac *ArrayCamera) RemoveEventListener(typ string, listener js.Value) {
	ac.Call("removeEventListener", typ, listener)
}
func (ac *ArrayCamera) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ac.Call("rotateOnAxis", axis, angle))
}
func (ac *ArrayCamera) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ac.Call("rotateOnWorldAxis", axis, angle))
}
func (ac *ArrayCamera) RotateX(angle int) this {
	return this(ac.Call("rotateX", angle))
}
func (ac *ArrayCamera) RotateY(angle int) this {
	return this(ac.Call("rotateY", angle))
}
func (ac *ArrayCamera) RotateZ(angle int) this {
	return this(ac.Call("rotateZ", angle))
}
func (ac *ArrayCamera) SetFocalLength(focalLength int) {
	ac.Call("setFocalLength", focalLength)
}
func (ac *ArrayCamera) SetLens(focalLength int, frameHeight int) {
	ac.Call("setLens", focalLength, frameHeight)
}
func (ac *ArrayCamera) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ac.Call("setRotationFromAxisAngle", axis, angle)
}
func (ac *ArrayCamera) SetRotationFromEuler(euler *math.Euler) {
	ac.Call("setRotationFromEuler", euler)
}
func (ac *ArrayCamera) SetRotationFromMatrix(m *math.Matrix4) {
	ac.Call("setRotationFromMatrix", m)
}
func (ac *ArrayCamera) SetRotationFromQuaternion(q *math.Quaternion) {
	ac.Call("setRotationFromQuaternion", q)
}
func (ac *ArrayCamera) SetViewOffset(fullWidth int, fullHeight int, x int, y int, width int, height int) {
	ac.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}
func (ac *ArrayCamera) ToJSON(meta js.Value) js.Value {
	return ac.Call("toJSON", meta)
}
func (ac *ArrayCamera) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ac.Call("translateOnAxis", axis, distance))
}
func (ac *ArrayCamera) TranslateX(distance int) this {
	return this(ac.Call("translateX", distance))
}
func (ac *ArrayCamera) TranslateY(distance int) this {
	return this(ac.Call("translateY", distance))
}
func (ac *ArrayCamera) TranslateZ(distance int) this {
	return this(ac.Call("translateZ", distance))
}
func (ac *ArrayCamera) Traverse(callback js.Value) {
	ac.Call("traverse", callback)
}
func (ac *ArrayCamera) TraverseAncestors(callback js.Value) {
	ac.Call("traverseAncestors", callback)
}
func (ac *ArrayCamera) TraverseVisible(callback js.Value) {
	ac.Call("traverseVisible", callback)
}
func (ac *ArrayCamera) UpdateMatrix() {
	ac.Call("updateMatrix")
}
func (ac *ArrayCamera) UpdateMatrixWorld(force bool) {
	ac.Call("updateMatrixWorld", force)
}
func (ac *ArrayCamera) UpdateProjectionMatrix() {
	ac.Call("updateProjectionMatrix")
}
func (ac *ArrayCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ac.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ac *ArrayCamera) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ac.Call("worldToLocal", vector)}
}

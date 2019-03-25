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

type PerspectiveCamera struct {
	js.Value
}

func NewPerspectiveCamera(fov int, aspect int, near int, far int) *PerspectiveCamera {
	return &PerspectiveCamera{Value: get("PerspectiveCamera").New(fov, aspect, near, far)}
}
func (pc *PerspectiveCamera) Aspect() int {
	return pc.Get("aspect").Int()
}
func (pc *PerspectiveCamera) SetAspect(v int) {
	pc.Set("aspect", v)
}
func (pc *PerspectiveCamera) CastShadow() bool {
	return pc.Get("castShadow").Bool()
}
func (pc *PerspectiveCamera) SetCastShadow(v bool) {
	pc.Set("castShadow", v)
}
func (pc *PerspectiveCamera) Children() []core.Object3D {
	return []core.Object3D(pc.Get("children"))
}
func (pc *PerspectiveCamera) SetChildren(v []core.Object3D) {
	pc.Set("children", v)
}
func (pc *PerspectiveCamera) Far() int {
	return pc.Get("far").Int()
}
func (pc *PerspectiveCamera) SetFar(v int) {
	pc.Set("far", v)
}
func (pc *PerspectiveCamera) FilmGauge() int {
	return pc.Get("filmGauge").Int()
}
func (pc *PerspectiveCamera) SetFilmGauge(v int) {
	pc.Set("filmGauge", v)
}
func (pc *PerspectiveCamera) FilmOffset() int {
	return pc.Get("filmOffset").Int()
}
func (pc *PerspectiveCamera) SetFilmOffset(v int) {
	pc.Set("filmOffset", v)
}
func (pc *PerspectiveCamera) Focus() int {
	return pc.Get("focus").Int()
}
func (pc *PerspectiveCamera) SetFocus(v int) {
	pc.Set("focus", v)
}
func (pc *PerspectiveCamera) Fov() int {
	return pc.Get("fov").Int()
}
func (pc *PerspectiveCamera) SetFov(v int) {
	pc.Set("fov", v)
}
func (pc *PerspectiveCamera) FrustumCulled() bool {
	return pc.Get("frustumCulled").Bool()
}
func (pc *PerspectiveCamera) SetFrustumCulled(v bool) {
	pc.Set("frustumCulled", v)
}
func (pc *PerspectiveCamera) Id() int {
	return pc.Get("id").Int()
}
func (pc *PerspectiveCamera) SetId(v int) {
	pc.Set("id", v)
}
func (pc *PerspectiveCamera) IsCamera() true {
	return true(pc.Get("isCamera"))
}
func (pc *PerspectiveCamera) SetIsCamera(v true) {
	pc.Set("isCamera", v)
}
func (pc *PerspectiveCamera) IsObject3D() true {
	return true(pc.Get("isObject3D"))
}
func (pc *PerspectiveCamera) SetIsObject3D(v true) {
	pc.Set("isObject3D", v)
}
func (pc *PerspectiveCamera) IsPerspectiveCamera() true {
	return true(pc.Get("isPerspectiveCamera"))
}
func (pc *PerspectiveCamera) SetIsPerspectiveCamera(v true) {
	pc.Set("isPerspectiveCamera", v)
}
func (pc *PerspectiveCamera) Layers() *core.Layers {
	return &core.Layers{Value: pc.Get("layers")}
}
func (pc *PerspectiveCamera) SetLayers(v *core.Layers) {
	pc.Set("layers", v)
}
func (pc *PerspectiveCamera) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: pc.Get("matrix")}
}
func (pc *PerspectiveCamera) SetMatrix(v *math.Matrix4) {
	pc.Set("matrix", v)
}
func (pc *PerspectiveCamera) MatrixAutoUpdate() bool {
	return pc.Get("matrixAutoUpdate").Bool()
}
func (pc *PerspectiveCamera) SetMatrixAutoUpdate(v bool) {
	pc.Set("matrixAutoUpdate", v)
}
func (pc *PerspectiveCamera) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: pc.Get("matrixWorld")}
}
func (pc *PerspectiveCamera) SetMatrixWorld(v *math.Matrix4) {
	pc.Set("matrixWorld", v)
}
func (pc *PerspectiveCamera) MatrixWorldInverse() *math.Matrix4 {
	return &math.Matrix4{Value: pc.Get("matrixWorldInverse")}
}
func (pc *PerspectiveCamera) SetMatrixWorldInverse(v *math.Matrix4) {
	pc.Set("matrixWorldInverse", v)
}
func (pc *PerspectiveCamera) MatrixWorldNeedsUpdate() bool {
	return pc.Get("matrixWorldNeedsUpdate").Bool()
}
func (pc *PerspectiveCamera) SetMatrixWorldNeedsUpdate(v bool) {
	pc.Set("matrixWorldNeedsUpdate", v)
}
func (pc *PerspectiveCamera) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: pc.Get("modelViewMatrix")}
}
func (pc *PerspectiveCamera) SetModelViewMatrix(v *math.Matrix4) {
	pc.Set("modelViewMatrix", v)
}
func (pc *PerspectiveCamera) Name() string {
	return pc.Get("name").String()
}
func (pc *PerspectiveCamera) SetName(v string) {
	pc.Set("name", v)
}
func (pc *PerspectiveCamera) Near() int {
	return pc.Get("near").Int()
}
func (pc *PerspectiveCamera) SetNear(v int) {
	pc.Set("near", v)
}
func (pc *PerspectiveCamera) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: pc.Get("normalMatrix")}
}
func (pc *PerspectiveCamera) SetNormalMatrix(v *math.Matrix3) {
	pc.Set("normalMatrix", v)
}
func (pc *PerspectiveCamera) OnAfterRender() js.Value {
	return pc.Get("onAfterRender")
}
func (pc *PerspectiveCamera) SetOnAfterRender(v js.Value) {
	pc.Set("onAfterRender", v)
}
func (pc *PerspectiveCamera) OnBeforeRender() js.Value {
	return pc.Get("onBeforeRender")
}
func (pc *PerspectiveCamera) SetOnBeforeRender(v js.Value) {
	pc.Set("onBeforeRender", v)
}
func (pc *PerspectiveCamera) Parent() core.Object3D {
	return core.Object3D(pc.Get("parent"))
}
func (pc *PerspectiveCamera) SetParent(v core.Object3D) {
	pc.Set("parent", v)
}
func (pc *PerspectiveCamera) Position() *math.Vector3 {
	return &math.Vector3{Value: pc.Get("position")}
}
func (pc *PerspectiveCamera) SetPosition(v *math.Vector3) {
	pc.Set("position", v)
}
func (pc *PerspectiveCamera) ProjectionMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: pc.Get("projectionMatrix")}
}
func (pc *PerspectiveCamera) SetProjectionMatrix(v *math.Matrix4) {
	pc.Set("projectionMatrix", v)
}
func (pc *PerspectiveCamera) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: pc.Get("quaternion")}
}
func (pc *PerspectiveCamera) SetQuaternion(v *math.Quaternion) {
	pc.Set("quaternion", v)
}
func (pc *PerspectiveCamera) ReceiveShadow() bool {
	return pc.Get("receiveShadow").Bool()
}
func (pc *PerspectiveCamera) SetReceiveShadow(v bool) {
	pc.Set("receiveShadow", v)
}
func (pc *PerspectiveCamera) RenderOrder() int {
	return pc.Get("renderOrder").Int()
}
func (pc *PerspectiveCamera) SetRenderOrder(v int) {
	pc.Set("renderOrder", v)
}
func (pc *PerspectiveCamera) Rotation() *math.Euler {
	return &math.Euler{Value: pc.Get("rotation")}
}
func (pc *PerspectiveCamera) SetRotation(v *math.Euler) {
	pc.Set("rotation", v)
}
func (pc *PerspectiveCamera) Scale() *math.Vector3 {
	return &math.Vector3{Value: pc.Get("scale")}
}
func (pc *PerspectiveCamera) SetScale(v *math.Vector3) {
	pc.Set("scale", v)
}
func (pc *PerspectiveCamera) Type() string {
	return pc.Get("type").String()
}
func (pc *PerspectiveCamera) SetType(v string) {
	pc.Set("type", v)
}
func (pc *PerspectiveCamera) Up() *math.Vector3 {
	return &math.Vector3{Value: pc.Get("up")}
}
func (pc *PerspectiveCamera) SetUp(v *math.Vector3) {
	pc.Set("up", v)
}
func (pc *PerspectiveCamera) UserData() js.Value {
	return pc.Get("userData")
}
func (pc *PerspectiveCamera) SetUserData(v js.Value) {
	pc.Set("userData", v)
}
func (pc *PerspectiveCamera) Uuid() string {
	return pc.Get("uuid").String()
}
func (pc *PerspectiveCamera) SetUuid(v string) {
	pc.Set("uuid", v)
}
func (pc *PerspectiveCamera) View() null {
	return null(pc.Get("view"))
}
func (pc *PerspectiveCamera) SetView(v null) {
	pc.Set("view", v)
}
func (pc *PerspectiveCamera) Visible() bool {
	return pc.Get("visible").Bool()
}
func (pc *PerspectiveCamera) SetVisible(v bool) {
	pc.Set("visible", v)
}
func (pc *PerspectiveCamera) Zoom() int {
	return pc.Get("zoom").Int()
}
func (pc *PerspectiveCamera) SetZoom(v int) {
	pc.Set("zoom", v)
}
func (pc *PerspectiveCamera) DefaultMatrixAutoUpdate() bool {
	return pc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pc *PerspectiveCamera) SetDefaultMatrixAutoUpdate(v bool) {
	pc.Set("DefaultMatrixAutoUpdate", v)
}
func (pc *PerspectiveCamera) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: pc.Get("DefaultUp")}
}
func (pc *PerspectiveCamera) SetDefaultUp(v *math.Vector3) {
	pc.Set("DefaultUp", v)
}
func (pc *PerspectiveCamera) Add(object []core.Object3D) this {
	return this(pc.Call("add", object))
}
func (pc *PerspectiveCamera) AddEventListener(typ string, listener js.Value) {
	pc.Call("addEventListener", typ, listener)
}
func (pc *PerspectiveCamera) ApplyMatrix(matrix *math.Matrix4) {
	pc.Call("applyMatrix", matrix)
}
func (pc *PerspectiveCamera) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(pc.Call("applyQuaternion", quaternion))
}
func (pc *PerspectiveCamera) ClearViewOffset() {
	pc.Call("clearViewOffset")
}
func (pc *PerspectiveCamera) Clone(recursive bool) this {
	return this(pc.Call("clone", recursive))
}
func (pc *PerspectiveCamera) Copy(source *Camera, recursive bool) this {
	return this(pc.Call("copy", source, recursive))
}
func (pc *PerspectiveCamera) DispatchEvent(event js.Value) {
	pc.Call("dispatchEvent", event)
}
func (pc *PerspectiveCamera) GetEffectiveFOV() int {
	return pc.Call("getEffectiveFOV").Int()
}
func (pc *PerspectiveCamera) GetFilmHeight() int {
	return pc.Call("getFilmHeight").Int()
}
func (pc *PerspectiveCamera) GetFilmWidth() int {
	return pc.Call("getFilmWidth").Int()
}
func (pc *PerspectiveCamera) GetFocalLength() int {
	return pc.Call("getFocalLength").Int()
}
func (pc *PerspectiveCamera) GetObjectById(id int) core.Object3D {
	return core.Object3D(pc.Call("getObjectById", id))
}
func (pc *PerspectiveCamera) GetObjectByName(name string) core.Object3D {
	return core.Object3D(pc.Call("getObjectByName", name))
}
func (pc *PerspectiveCamera) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(pc.Call("getObjectByProperty", name, value))
}
func (pc *PerspectiveCamera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pc.Call("getWorldDirection", target)}
}
func (pc *PerspectiveCamera) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pc.Call("getWorldPosition", target)}
}
func (pc *PerspectiveCamera) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: pc.Call("getWorldQuaternion", target)}
}
func (pc *PerspectiveCamera) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pc.Call("getWorldScale", target)}
}
func (pc *PerspectiveCamera) HasEventListener(typ string, listener js.Value) bool {
	return pc.Call("hasEventListener", typ, listener).Bool()
}
func (pc *PerspectiveCamera) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pc.Call("localToWorld", vector)}
}
func (pc *PerspectiveCamera) LookAt(vector math.Vector3, y int, z int) {
	pc.Call("lookAt", vector, y, z)
}
func (pc *PerspectiveCamera) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	pc.Call("raycast", raycaster, intersects)
}
func (pc *PerspectiveCamera) Remove(object []core.Object3D) this {
	return this(pc.Call("remove", object))
}
func (pc *PerspectiveCamera) RemoveEventListener(typ string, listener js.Value) {
	pc.Call("removeEventListener", typ, listener)
}
func (pc *PerspectiveCamera) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(pc.Call("rotateOnAxis", axis, angle))
}
func (pc *PerspectiveCamera) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(pc.Call("rotateOnWorldAxis", axis, angle))
}
func (pc *PerspectiveCamera) RotateX(angle int) this {
	return this(pc.Call("rotateX", angle))
}
func (pc *PerspectiveCamera) RotateY(angle int) this {
	return this(pc.Call("rotateY", angle))
}
func (pc *PerspectiveCamera) RotateZ(angle int) this {
	return this(pc.Call("rotateZ", angle))
}
func (pc *PerspectiveCamera) SetFocalLength(focalLength int) {
	pc.Call("setFocalLength", focalLength)
}
func (pc *PerspectiveCamera) SetLens(focalLength int, frameHeight int) {
	pc.Call("setLens", focalLength, frameHeight)
}
func (pc *PerspectiveCamera) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	pc.Call("setRotationFromAxisAngle", axis, angle)
}
func (pc *PerspectiveCamera) SetRotationFromEuler(euler *math.Euler) {
	pc.Call("setRotationFromEuler", euler)
}
func (pc *PerspectiveCamera) SetRotationFromMatrix(m *math.Matrix4) {
	pc.Call("setRotationFromMatrix", m)
}
func (pc *PerspectiveCamera) SetRotationFromQuaternion(q *math.Quaternion) {
	pc.Call("setRotationFromQuaternion", q)
}
func (pc *PerspectiveCamera) SetViewOffset(fullWidth int, fullHeight int, x int, y int, width int, height int) {
	pc.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}
func (pc *PerspectiveCamera) ToJSON(meta js.Value) js.Value {
	return pc.Call("toJSON", meta)
}
func (pc *PerspectiveCamera) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(pc.Call("translateOnAxis", axis, distance))
}
func (pc *PerspectiveCamera) TranslateX(distance int) this {
	return this(pc.Call("translateX", distance))
}
func (pc *PerspectiveCamera) TranslateY(distance int) this {
	return this(pc.Call("translateY", distance))
}
func (pc *PerspectiveCamera) TranslateZ(distance int) this {
	return this(pc.Call("translateZ", distance))
}
func (pc *PerspectiveCamera) Traverse(callback js.Value) {
	pc.Call("traverse", callback)
}
func (pc *PerspectiveCamera) TraverseAncestors(callback js.Value) {
	pc.Call("traverseAncestors", callback)
}
func (pc *PerspectiveCamera) TraverseVisible(callback js.Value) {
	pc.Call("traverseVisible", callback)
}
func (pc *PerspectiveCamera) UpdateMatrix() {
	pc.Call("updateMatrix")
}
func (pc *PerspectiveCamera) UpdateMatrixWorld(force bool) {
	pc.Call("updateMatrixWorld", force)
}
func (pc *PerspectiveCamera) UpdateProjectionMatrix() {
	pc.Call("updateProjectionMatrix")
}
func (pc *PerspectiveCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pc *PerspectiveCamera) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pc.Call("worldToLocal", vector)}
}

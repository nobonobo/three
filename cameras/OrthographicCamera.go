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

type OrthographicCamera struct {
	js.Value
}

func NewOrthographicCamera(left int, right int, top int, bottom int, near int, far int) *OrthographicCamera {
	return &OrthographicCamera{Value: get("OrthographicCamera").New(left, right, top, bottom, near, far)}
}
func (oc *OrthographicCamera) Bottom() int {
	return oc.Get("bottom").Int()
}
func (oc *OrthographicCamera) SetBottom(v int) {
	oc.Set("bottom", v)
}
func (oc *OrthographicCamera) CastShadow() bool {
	return oc.Get("castShadow").Bool()
}
func (oc *OrthographicCamera) SetCastShadow(v bool) {
	oc.Set("castShadow", v)
}
func (oc *OrthographicCamera) Children() []core.Object3D {
	return []core.Object3D(oc.Get("children"))
}
func (oc *OrthographicCamera) SetChildren(v []core.Object3D) {
	oc.Set("children", v)
}
func (oc *OrthographicCamera) Far() int {
	return oc.Get("far").Int()
}
func (oc *OrthographicCamera) SetFar(v int) {
	oc.Set("far", v)
}
func (oc *OrthographicCamera) FrustumCulled() bool {
	return oc.Get("frustumCulled").Bool()
}
func (oc *OrthographicCamera) SetFrustumCulled(v bool) {
	oc.Set("frustumCulled", v)
}
func (oc *OrthographicCamera) Id() int {
	return oc.Get("id").Int()
}
func (oc *OrthographicCamera) SetId(v int) {
	oc.Set("id", v)
}
func (oc *OrthographicCamera) IsCamera() true {
	return true(oc.Get("isCamera"))
}
func (oc *OrthographicCamera) SetIsCamera(v true) {
	oc.Set("isCamera", v)
}
func (oc *OrthographicCamera) IsObject3D() true {
	return true(oc.Get("isObject3D"))
}
func (oc *OrthographicCamera) SetIsObject3D(v true) {
	oc.Set("isObject3D", v)
}
func (oc *OrthographicCamera) IsOrthographicCamera() true {
	return true(oc.Get("isOrthographicCamera"))
}
func (oc *OrthographicCamera) SetIsOrthographicCamera(v true) {
	oc.Set("isOrthographicCamera", v)
}
func (oc *OrthographicCamera) Layers() *core.Layers {
	return &core.Layers{Value: oc.Get("layers")}
}
func (oc *OrthographicCamera) SetLayers(v *core.Layers) {
	oc.Set("layers", v)
}
func (oc *OrthographicCamera) Left() int {
	return oc.Get("left").Int()
}
func (oc *OrthographicCamera) SetLeft(v int) {
	oc.Set("left", v)
}
func (oc *OrthographicCamera) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: oc.Get("matrix")}
}
func (oc *OrthographicCamera) SetMatrix(v *math.Matrix4) {
	oc.Set("matrix", v)
}
func (oc *OrthographicCamera) MatrixAutoUpdate() bool {
	return oc.Get("matrixAutoUpdate").Bool()
}
func (oc *OrthographicCamera) SetMatrixAutoUpdate(v bool) {
	oc.Set("matrixAutoUpdate", v)
}
func (oc *OrthographicCamera) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: oc.Get("matrixWorld")}
}
func (oc *OrthographicCamera) SetMatrixWorld(v *math.Matrix4) {
	oc.Set("matrixWorld", v)
}
func (oc *OrthographicCamera) MatrixWorldInverse() *math.Matrix4 {
	return &math.Matrix4{Value: oc.Get("matrixWorldInverse")}
}
func (oc *OrthographicCamera) SetMatrixWorldInverse(v *math.Matrix4) {
	oc.Set("matrixWorldInverse", v)
}
func (oc *OrthographicCamera) MatrixWorldNeedsUpdate() bool {
	return oc.Get("matrixWorldNeedsUpdate").Bool()
}
func (oc *OrthographicCamera) SetMatrixWorldNeedsUpdate(v bool) {
	oc.Set("matrixWorldNeedsUpdate", v)
}
func (oc *OrthographicCamera) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: oc.Get("modelViewMatrix")}
}
func (oc *OrthographicCamera) SetModelViewMatrix(v *math.Matrix4) {
	oc.Set("modelViewMatrix", v)
}
func (oc *OrthographicCamera) Name() string {
	return oc.Get("name").String()
}
func (oc *OrthographicCamera) SetName(v string) {
	oc.Set("name", v)
}
func (oc *OrthographicCamera) Near() int {
	return oc.Get("near").Int()
}
func (oc *OrthographicCamera) SetNear(v int) {
	oc.Set("near", v)
}
func (oc *OrthographicCamera) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: oc.Get("normalMatrix")}
}
func (oc *OrthographicCamera) SetNormalMatrix(v *math.Matrix3) {
	oc.Set("normalMatrix", v)
}
func (oc *OrthographicCamera) OnAfterRender() js.Value {
	return oc.Get("onAfterRender")
}
func (oc *OrthographicCamera) SetOnAfterRender(v js.Value) {
	oc.Set("onAfterRender", v)
}
func (oc *OrthographicCamera) OnBeforeRender() js.Value {
	return oc.Get("onBeforeRender")
}
func (oc *OrthographicCamera) SetOnBeforeRender(v js.Value) {
	oc.Set("onBeforeRender", v)
}
func (oc *OrthographicCamera) Parent() core.Object3D {
	return core.Object3D(oc.Get("parent"))
}
func (oc *OrthographicCamera) SetParent(v core.Object3D) {
	oc.Set("parent", v)
}
func (oc *OrthographicCamera) Position() *math.Vector3 {
	return &math.Vector3{Value: oc.Get("position")}
}
func (oc *OrthographicCamera) SetPosition(v *math.Vector3) {
	oc.Set("position", v)
}
func (oc *OrthographicCamera) ProjectionMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: oc.Get("projectionMatrix")}
}
func (oc *OrthographicCamera) SetProjectionMatrix(v *math.Matrix4) {
	oc.Set("projectionMatrix", v)
}
func (oc *OrthographicCamera) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: oc.Get("quaternion")}
}
func (oc *OrthographicCamera) SetQuaternion(v *math.Quaternion) {
	oc.Set("quaternion", v)
}
func (oc *OrthographicCamera) ReceiveShadow() bool {
	return oc.Get("receiveShadow").Bool()
}
func (oc *OrthographicCamera) SetReceiveShadow(v bool) {
	oc.Set("receiveShadow", v)
}
func (oc *OrthographicCamera) RenderOrder() int {
	return oc.Get("renderOrder").Int()
}
func (oc *OrthographicCamera) SetRenderOrder(v int) {
	oc.Set("renderOrder", v)
}
func (oc *OrthographicCamera) Right() int {
	return oc.Get("right").Int()
}
func (oc *OrthographicCamera) SetRight(v int) {
	oc.Set("right", v)
}
func (oc *OrthographicCamera) Rotation() *math.Euler {
	return &math.Euler{Value: oc.Get("rotation")}
}
func (oc *OrthographicCamera) SetRotation(v *math.Euler) {
	oc.Set("rotation", v)
}
func (oc *OrthographicCamera) Scale() *math.Vector3 {
	return &math.Vector3{Value: oc.Get("scale")}
}
func (oc *OrthographicCamera) SetScale(v *math.Vector3) {
	oc.Set("scale", v)
}
func (oc *OrthographicCamera) Top() int {
	return oc.Get("top").Int()
}
func (oc *OrthographicCamera) SetTop(v int) {
	oc.Set("top", v)
}
func (oc *OrthographicCamera) Type() string {
	return oc.Get("type").String()
}
func (oc *OrthographicCamera) SetType(v string) {
	oc.Set("type", v)
}
func (oc *OrthographicCamera) Up() *math.Vector3 {
	return &math.Vector3{Value: oc.Get("up")}
}
func (oc *OrthographicCamera) SetUp(v *math.Vector3) {
	oc.Set("up", v)
}
func (oc *OrthographicCamera) UserData() js.Value {
	return oc.Get("userData")
}
func (oc *OrthographicCamera) SetUserData(v js.Value) {
	oc.Set("userData", v)
}
func (oc *OrthographicCamera) Uuid() string {
	return oc.Get("uuid").String()
}
func (oc *OrthographicCamera) SetUuid(v string) {
	oc.Set("uuid", v)
}
func (oc *OrthographicCamera) View() null {
	return null(oc.Get("view"))
}
func (oc *OrthographicCamera) SetView(v null) {
	oc.Set("view", v)
}
func (oc *OrthographicCamera) Visible() bool {
	return oc.Get("visible").Bool()
}
func (oc *OrthographicCamera) SetVisible(v bool) {
	oc.Set("visible", v)
}
func (oc *OrthographicCamera) Zoom() int {
	return oc.Get("zoom").Int()
}
func (oc *OrthographicCamera) SetZoom(v int) {
	oc.Set("zoom", v)
}
func (oc *OrthographicCamera) DefaultMatrixAutoUpdate() bool {
	return oc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (oc *OrthographicCamera) SetDefaultMatrixAutoUpdate(v bool) {
	oc.Set("DefaultMatrixAutoUpdate", v)
}
func (oc *OrthographicCamera) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: oc.Get("DefaultUp")}
}
func (oc *OrthographicCamera) SetDefaultUp(v *math.Vector3) {
	oc.Set("DefaultUp", v)
}
func (oc *OrthographicCamera) Add(object []core.Object3D) this {
	return this(oc.Call("add", object))
}
func (oc *OrthographicCamera) AddEventListener(typ string, listener js.Value) {
	oc.Call("addEventListener", typ, listener)
}
func (oc *OrthographicCamera) ApplyMatrix(matrix *math.Matrix4) {
	oc.Call("applyMatrix", matrix)
}
func (oc *OrthographicCamera) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(oc.Call("applyQuaternion", quaternion))
}
func (oc *OrthographicCamera) ClearViewOffset() {
	oc.Call("clearViewOffset")
}
func (oc *OrthographicCamera) Clone(recursive bool) this {
	return this(oc.Call("clone", recursive))
}
func (oc *OrthographicCamera) Copy(source *Camera, recursive bool) this {
	return this(oc.Call("copy", source, recursive))
}
func (oc *OrthographicCamera) DispatchEvent(event js.Value) {
	oc.Call("dispatchEvent", event)
}
func (oc *OrthographicCamera) GetObjectById(id int) core.Object3D {
	return core.Object3D(oc.Call("getObjectById", id))
}
func (oc *OrthographicCamera) GetObjectByName(name string) core.Object3D {
	return core.Object3D(oc.Call("getObjectByName", name))
}
func (oc *OrthographicCamera) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(oc.Call("getObjectByProperty", name, value))
}
func (oc *OrthographicCamera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: oc.Call("getWorldDirection", target)}
}
func (oc *OrthographicCamera) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: oc.Call("getWorldPosition", target)}
}
func (oc *OrthographicCamera) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: oc.Call("getWorldQuaternion", target)}
}
func (oc *OrthographicCamera) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: oc.Call("getWorldScale", target)}
}
func (oc *OrthographicCamera) HasEventListener(typ string, listener js.Value) bool {
	return oc.Call("hasEventListener", typ, listener).Bool()
}
func (oc *OrthographicCamera) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: oc.Call("localToWorld", vector)}
}
func (oc *OrthographicCamera) LookAt(vector math.Vector3, y int, z int) {
	oc.Call("lookAt", vector, y, z)
}
func (oc *OrthographicCamera) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	oc.Call("raycast", raycaster, intersects)
}
func (oc *OrthographicCamera) Remove(object []core.Object3D) this {
	return this(oc.Call("remove", object))
}
func (oc *OrthographicCamera) RemoveEventListener(typ string, listener js.Value) {
	oc.Call("removeEventListener", typ, listener)
}
func (oc *OrthographicCamera) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(oc.Call("rotateOnAxis", axis, angle))
}
func (oc *OrthographicCamera) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(oc.Call("rotateOnWorldAxis", axis, angle))
}
func (oc *OrthographicCamera) RotateX(angle int) this {
	return this(oc.Call("rotateX", angle))
}
func (oc *OrthographicCamera) RotateY(angle int) this {
	return this(oc.Call("rotateY", angle))
}
func (oc *OrthographicCamera) RotateZ(angle int) this {
	return this(oc.Call("rotateZ", angle))
}
func (oc *OrthographicCamera) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	oc.Call("setRotationFromAxisAngle", axis, angle)
}
func (oc *OrthographicCamera) SetRotationFromEuler(euler *math.Euler) {
	oc.Call("setRotationFromEuler", euler)
}
func (oc *OrthographicCamera) SetRotationFromMatrix(m *math.Matrix4) {
	oc.Call("setRotationFromMatrix", m)
}
func (oc *OrthographicCamera) SetRotationFromQuaternion(q *math.Quaternion) {
	oc.Call("setRotationFromQuaternion", q)
}
func (oc *OrthographicCamera) SetViewOffset(fullWidth int, fullHeight int, offsetX int, offsetY int, width int, height int) {
	oc.Call("setViewOffset", fullWidth, fullHeight, offsetX, offsetY, width, height)
}
func (oc *OrthographicCamera) ToJSON(meta js.Value) js.Value {
	return oc.Call("toJSON", meta)
}
func (oc *OrthographicCamera) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(oc.Call("translateOnAxis", axis, distance))
}
func (oc *OrthographicCamera) TranslateX(distance int) this {
	return this(oc.Call("translateX", distance))
}
func (oc *OrthographicCamera) TranslateY(distance int) this {
	return this(oc.Call("translateY", distance))
}
func (oc *OrthographicCamera) TranslateZ(distance int) this {
	return this(oc.Call("translateZ", distance))
}
func (oc *OrthographicCamera) Traverse(callback js.Value) {
	oc.Call("traverse", callback)
}
func (oc *OrthographicCamera) TraverseAncestors(callback js.Value) {
	oc.Call("traverseAncestors", callback)
}
func (oc *OrthographicCamera) TraverseVisible(callback js.Value) {
	oc.Call("traverseVisible", callback)
}
func (oc *OrthographicCamera) UpdateMatrix() {
	oc.Call("updateMatrix")
}
func (oc *OrthographicCamera) UpdateMatrixWorld(force bool) {
	oc.Call("updateMatrixWorld", force)
}
func (oc *OrthographicCamera) UpdateProjectionMatrix() {
	oc.Call("updateProjectionMatrix")
}
func (oc *OrthographicCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	oc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (oc *OrthographicCamera) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: oc.Call("worldToLocal", vector)}
}

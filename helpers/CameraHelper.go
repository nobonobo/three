package helpers

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

type CameraHelper struct {
	js.Value
}

func NewCameraHelper(camera *cameras.Camera) *CameraHelper {
	return &CameraHelper{Value: get("CameraHelper").New(camera)}
}
func (ch *CameraHelper) Camera() *cameras.Camera {
	return &cameras.Camera{Value: ch.Get("camera")}
}
func (ch *CameraHelper) SetCamera(v *cameras.Camera) {
	ch.Set("camera", v)
}
func (ch *CameraHelper) CastShadow() bool {
	return ch.Get("castShadow").Bool()
}
func (ch *CameraHelper) SetCastShadow(v bool) {
	ch.Set("castShadow", v)
}
func (ch *CameraHelper) Children() []core.Object3D {
	return []core.Object3D(ch.Get("children"))
}
func (ch *CameraHelper) SetChildren(v []core.Object3D) {
	ch.Set("children", v)
}
func (ch *CameraHelper) FrustumCulled() bool {
	return ch.Get("frustumCulled").Bool()
}
func (ch *CameraHelper) SetFrustumCulled(v bool) {
	ch.Set("frustumCulled", v)
}
func (ch *CameraHelper) Geometry() core.Geometry {
	return core.Geometry(ch.Get("geometry"))
}
func (ch *CameraHelper) SetGeometry(v core.Geometry) {
	ch.Set("geometry", v)
}
func (ch *CameraHelper) Id() int {
	return ch.Get("id").Int()
}
func (ch *CameraHelper) SetId(v int) {
	ch.Set("id", v)
}
func (ch *CameraHelper) IsLine() true {
	return true(ch.Get("isLine"))
}
func (ch *CameraHelper) SetIsLine(v true) {
	ch.Set("isLine", v)
}
func (ch *CameraHelper) IsLineSegments() true {
	return true(ch.Get("isLineSegments"))
}
func (ch *CameraHelper) SetIsLineSegments(v true) {
	ch.Set("isLineSegments", v)
}
func (ch *CameraHelper) IsObject3D() true {
	return true(ch.Get("isObject3D"))
}
func (ch *CameraHelper) SetIsObject3D(v true) {
	ch.Set("isObject3D", v)
}
func (ch *CameraHelper) Layers() *core.Layers {
	return &core.Layers{Value: ch.Get("layers")}
}
func (ch *CameraHelper) SetLayers(v *core.Layers) {
	ch.Set("layers", v)
}
func (ch *CameraHelper) Material() materials.Material {
	return materials.Material(ch.Get("material"))
}
func (ch *CameraHelper) SetMaterial(v materials.Material) {
	ch.Set("material", v)
}
func (ch *CameraHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ch.Get("matrix")}
}
func (ch *CameraHelper) SetMatrix(v *math.Matrix4) {
	ch.Set("matrix", v)
}
func (ch *CameraHelper) MatrixAutoUpdate() bool {
	return ch.Get("matrixAutoUpdate").Bool()
}
func (ch *CameraHelper) SetMatrixAutoUpdate(v bool) {
	ch.Set("matrixAutoUpdate", v)
}
func (ch *CameraHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ch.Get("matrixWorld")}
}
func (ch *CameraHelper) SetMatrixWorld(v *math.Matrix4) {
	ch.Set("matrixWorld", v)
}
func (ch *CameraHelper) MatrixWorldNeedsUpdate() bool {
	return ch.Get("matrixWorldNeedsUpdate").Bool()
}
func (ch *CameraHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ch.Set("matrixWorldNeedsUpdate", v)
}
func (ch *CameraHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ch.Get("modelViewMatrix")}
}
func (ch *CameraHelper) SetModelViewMatrix(v *math.Matrix4) {
	ch.Set("modelViewMatrix", v)
}
func (ch *CameraHelper) Name() string {
	return ch.Get("name").String()
}
func (ch *CameraHelper) SetName(v string) {
	ch.Set("name", v)
}
func (ch *CameraHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ch.Get("normalMatrix")}
}
func (ch *CameraHelper) SetNormalMatrix(v *math.Matrix3) {
	ch.Set("normalMatrix", v)
}
func (ch *CameraHelper) OnAfterRender() js.Value {
	return ch.Get("onAfterRender")
}
func (ch *CameraHelper) SetOnAfterRender(v js.Value) {
	ch.Set("onAfterRender", v)
}
func (ch *CameraHelper) OnBeforeRender() js.Value {
	return ch.Get("onBeforeRender")
}
func (ch *CameraHelper) SetOnBeforeRender(v js.Value) {
	ch.Set("onBeforeRender", v)
}
func (ch *CameraHelper) Parent() core.Object3D {
	return core.Object3D(ch.Get("parent"))
}
func (ch *CameraHelper) SetParent(v core.Object3D) {
	ch.Set("parent", v)
}
func (ch *CameraHelper) PointMap() js.Value {
	return ch.Get("pointMap")
}
func (ch *CameraHelper) SetPointMap(v js.Value) {
	ch.Set("pointMap", v)
}
func (ch *CameraHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: ch.Get("position")}
}
func (ch *CameraHelper) SetPosition(v *math.Vector3) {
	ch.Set("position", v)
}
func (ch *CameraHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ch.Get("quaternion")}
}
func (ch *CameraHelper) SetQuaternion(v *math.Quaternion) {
	ch.Set("quaternion", v)
}
func (ch *CameraHelper) ReceiveShadow() bool {
	return ch.Get("receiveShadow").Bool()
}
func (ch *CameraHelper) SetReceiveShadow(v bool) {
	ch.Set("receiveShadow", v)
}
func (ch *CameraHelper) RenderOrder() int {
	return ch.Get("renderOrder").Int()
}
func (ch *CameraHelper) SetRenderOrder(v int) {
	ch.Set("renderOrder", v)
}
func (ch *CameraHelper) Rotation() *math.Euler {
	return &math.Euler{Value: ch.Get("rotation")}
}
func (ch *CameraHelper) SetRotation(v *math.Euler) {
	ch.Set("rotation", v)
}
func (ch *CameraHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: ch.Get("scale")}
}
func (ch *CameraHelper) SetScale(v *math.Vector3) {
	ch.Set("scale", v)
}
func (ch *CameraHelper) Type() string {
	return ch.Get("type").String()
}
func (ch *CameraHelper) SetType(v string) {
	ch.Set("type", v)
}
func (ch *CameraHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: ch.Get("up")}
}
func (ch *CameraHelper) SetUp(v *math.Vector3) {
	ch.Set("up", v)
}
func (ch *CameraHelper) UserData() js.Value {
	return ch.Get("userData")
}
func (ch *CameraHelper) SetUserData(v js.Value) {
	ch.Set("userData", v)
}
func (ch *CameraHelper) Uuid() string {
	return ch.Get("uuid").String()
}
func (ch *CameraHelper) SetUuid(v string) {
	ch.Set("uuid", v)
}
func (ch *CameraHelper) Visible() bool {
	return ch.Get("visible").Bool()
}
func (ch *CameraHelper) SetVisible(v bool) {
	ch.Set("visible", v)
}
func (ch *CameraHelper) DefaultMatrixAutoUpdate() bool {
	return ch.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ch *CameraHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ch.Set("DefaultMatrixAutoUpdate", v)
}
func (ch *CameraHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ch.Get("DefaultUp")}
}
func (ch *CameraHelper) SetDefaultUp(v *math.Vector3) {
	ch.Set("DefaultUp", v)
}
func (ch *CameraHelper) Add(object []core.Object3D) this {
	return this(ch.Call("add", object))
}
func (ch *CameraHelper) AddEventListener(typ string, listener js.Value) {
	ch.Call("addEventListener", typ, listener)
}
func (ch *CameraHelper) ApplyMatrix(matrix *math.Matrix4) {
	ch.Call("applyMatrix", matrix)
}
func (ch *CameraHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ch.Call("applyQuaternion", quaternion))
}
func (ch *CameraHelper) Clone(recursive bool) this {
	return this(ch.Call("clone", recursive))
}
func (ch *CameraHelper) ComputeLineDistances() this {
	return this(ch.Call("computeLineDistances"))
}
func (ch *CameraHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(ch.Call("copy", source, recursive))
}
func (ch *CameraHelper) DispatchEvent(event js.Value) {
	ch.Call("dispatchEvent", event)
}
func (ch *CameraHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(ch.Call("getObjectById", id))
}
func (ch *CameraHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ch.Call("getObjectByName", name))
}
func (ch *CameraHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ch.Call("getObjectByProperty", name, value))
}
func (ch *CameraHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ch.Call("getWorldDirection", target)}
}
func (ch *CameraHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ch.Call("getWorldPosition", target)}
}
func (ch *CameraHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ch.Call("getWorldQuaternion", target)}
}
func (ch *CameraHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ch.Call("getWorldScale", target)}
}
func (ch *CameraHelper) HasEventListener(typ string, listener js.Value) bool {
	return ch.Call("hasEventListener", typ, listener).Bool()
}
func (ch *CameraHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ch.Call("localToWorld", vector)}
}
func (ch *CameraHelper) LookAt(vector math.Vector3, y int, z int) {
	ch.Call("lookAt", vector, y, z)
}
func (ch *CameraHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ch.Call("raycast", raycaster, intersects)
}
func (ch *CameraHelper) Remove(object []core.Object3D) this {
	return this(ch.Call("remove", object))
}
func (ch *CameraHelper) RemoveEventListener(typ string, listener js.Value) {
	ch.Call("removeEventListener", typ, listener)
}
func (ch *CameraHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ch.Call("rotateOnAxis", axis, angle))
}
func (ch *CameraHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ch.Call("rotateOnWorldAxis", axis, angle))
}
func (ch *CameraHelper) RotateX(angle int) this {
	return this(ch.Call("rotateX", angle))
}
func (ch *CameraHelper) RotateY(angle int) this {
	return this(ch.Call("rotateY", angle))
}
func (ch *CameraHelper) RotateZ(angle int) this {
	return this(ch.Call("rotateZ", angle))
}
func (ch *CameraHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ch.Call("setRotationFromAxisAngle", axis, angle)
}
func (ch *CameraHelper) SetRotationFromEuler(euler *math.Euler) {
	ch.Call("setRotationFromEuler", euler)
}
func (ch *CameraHelper) SetRotationFromMatrix(m *math.Matrix4) {
	ch.Call("setRotationFromMatrix", m)
}
func (ch *CameraHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	ch.Call("setRotationFromQuaternion", q)
}
func (ch *CameraHelper) ToJSON(meta js.Value) js.Value {
	return ch.Call("toJSON", meta)
}
func (ch *CameraHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ch.Call("translateOnAxis", axis, distance))
}
func (ch *CameraHelper) TranslateX(distance int) this {
	return this(ch.Call("translateX", distance))
}
func (ch *CameraHelper) TranslateY(distance int) this {
	return this(ch.Call("translateY", distance))
}
func (ch *CameraHelper) TranslateZ(distance int) this {
	return this(ch.Call("translateZ", distance))
}
func (ch *CameraHelper) Traverse(callback js.Value) {
	ch.Call("traverse", callback)
}
func (ch *CameraHelper) TraverseAncestors(callback js.Value) {
	ch.Call("traverseAncestors", callback)
}
func (ch *CameraHelper) TraverseVisible(callback js.Value) {
	ch.Call("traverseVisible", callback)
}
func (ch *CameraHelper) Update() {
	ch.Call("update")
}
func (ch *CameraHelper) UpdateMatrix() {
	ch.Call("updateMatrix")
}
func (ch *CameraHelper) UpdateMatrixWorld(force bool) {
	ch.Call("updateMatrixWorld", force)
}
func (ch *CameraHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ch.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ch *CameraHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ch.Call("worldToLocal", vector)}
}

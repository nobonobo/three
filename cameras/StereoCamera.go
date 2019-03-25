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

type StereoCamera struct {
	js.Value
}

func NewStereoCamera() *StereoCamera {
	return &StereoCamera{Value: get("StereoCamera").New()}
}
func (sc *StereoCamera) Aspect() int {
	return sc.Get("aspect").Int()
}
func (sc *StereoCamera) SetAspect(v int) {
	sc.Set("aspect", v)
}
func (sc *StereoCamera) CameraL() *PerspectiveCamera {
	return &PerspectiveCamera{Value: sc.Get("cameraL")}
}
func (sc *StereoCamera) SetCameraL(v *PerspectiveCamera) {
	sc.Set("cameraL", v)
}
func (sc *StereoCamera) CameraR() *PerspectiveCamera {
	return &PerspectiveCamera{Value: sc.Get("cameraR")}
}
func (sc *StereoCamera) SetCameraR(v *PerspectiveCamera) {
	sc.Set("cameraR", v)
}
func (sc *StereoCamera) CastShadow() bool {
	return sc.Get("castShadow").Bool()
}
func (sc *StereoCamera) SetCastShadow(v bool) {
	sc.Set("castShadow", v)
}
func (sc *StereoCamera) Children() []core.Object3D {
	return []core.Object3D(sc.Get("children"))
}
func (sc *StereoCamera) SetChildren(v []core.Object3D) {
	sc.Set("children", v)
}
func (sc *StereoCamera) EyeSep() int {
	return sc.Get("eyeSep").Int()
}
func (sc *StereoCamera) SetEyeSep(v int) {
	sc.Set("eyeSep", v)
}
func (sc *StereoCamera) FrustumCulled() bool {
	return sc.Get("frustumCulled").Bool()
}
func (sc *StereoCamera) SetFrustumCulled(v bool) {
	sc.Set("frustumCulled", v)
}
func (sc *StereoCamera) Id() int {
	return sc.Get("id").Int()
}
func (sc *StereoCamera) SetId(v int) {
	sc.Set("id", v)
}
func (sc *StereoCamera) IsCamera() true {
	return true(sc.Get("isCamera"))
}
func (sc *StereoCamera) SetIsCamera(v true) {
	sc.Set("isCamera", v)
}
func (sc *StereoCamera) IsObject3D() true {
	return true(sc.Get("isObject3D"))
}
func (sc *StereoCamera) SetIsObject3D(v true) {
	sc.Set("isObject3D", v)
}
func (sc *StereoCamera) Layers() *core.Layers {
	return &core.Layers{Value: sc.Get("layers")}
}
func (sc *StereoCamera) SetLayers(v *core.Layers) {
	sc.Set("layers", v)
}
func (sc *StereoCamera) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: sc.Get("matrix")}
}
func (sc *StereoCamera) SetMatrix(v *math.Matrix4) {
	sc.Set("matrix", v)
}
func (sc *StereoCamera) MatrixAutoUpdate() bool {
	return sc.Get("matrixAutoUpdate").Bool()
}
func (sc *StereoCamera) SetMatrixAutoUpdate(v bool) {
	sc.Set("matrixAutoUpdate", v)
}
func (sc *StereoCamera) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: sc.Get("matrixWorld")}
}
func (sc *StereoCamera) SetMatrixWorld(v *math.Matrix4) {
	sc.Set("matrixWorld", v)
}
func (sc *StereoCamera) MatrixWorldInverse() *math.Matrix4 {
	return &math.Matrix4{Value: sc.Get("matrixWorldInverse")}
}
func (sc *StereoCamera) SetMatrixWorldInverse(v *math.Matrix4) {
	sc.Set("matrixWorldInverse", v)
}
func (sc *StereoCamera) MatrixWorldNeedsUpdate() bool {
	return sc.Get("matrixWorldNeedsUpdate").Bool()
}
func (sc *StereoCamera) SetMatrixWorldNeedsUpdate(v bool) {
	sc.Set("matrixWorldNeedsUpdate", v)
}
func (sc *StereoCamera) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: sc.Get("modelViewMatrix")}
}
func (sc *StereoCamera) SetModelViewMatrix(v *math.Matrix4) {
	sc.Set("modelViewMatrix", v)
}
func (sc *StereoCamera) Name() string {
	return sc.Get("name").String()
}
func (sc *StereoCamera) SetName(v string) {
	sc.Set("name", v)
}
func (sc *StereoCamera) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: sc.Get("normalMatrix")}
}
func (sc *StereoCamera) SetNormalMatrix(v *math.Matrix3) {
	sc.Set("normalMatrix", v)
}
func (sc *StereoCamera) OnAfterRender() js.Value {
	return sc.Get("onAfterRender")
}
func (sc *StereoCamera) SetOnAfterRender(v js.Value) {
	sc.Set("onAfterRender", v)
}
func (sc *StereoCamera) OnBeforeRender() js.Value {
	return sc.Get("onBeforeRender")
}
func (sc *StereoCamera) SetOnBeforeRender(v js.Value) {
	sc.Set("onBeforeRender", v)
}
func (sc *StereoCamera) Parent() core.Object3D {
	return core.Object3D(sc.Get("parent"))
}
func (sc *StereoCamera) SetParent(v core.Object3D) {
	sc.Set("parent", v)
}
func (sc *StereoCamera) Position() *math.Vector3 {
	return &math.Vector3{Value: sc.Get("position")}
}
func (sc *StereoCamera) SetPosition(v *math.Vector3) {
	sc.Set("position", v)
}
func (sc *StereoCamera) ProjectionMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: sc.Get("projectionMatrix")}
}
func (sc *StereoCamera) SetProjectionMatrix(v *math.Matrix4) {
	sc.Set("projectionMatrix", v)
}
func (sc *StereoCamera) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: sc.Get("quaternion")}
}
func (sc *StereoCamera) SetQuaternion(v *math.Quaternion) {
	sc.Set("quaternion", v)
}
func (sc *StereoCamera) ReceiveShadow() bool {
	return sc.Get("receiveShadow").Bool()
}
func (sc *StereoCamera) SetReceiveShadow(v bool) {
	sc.Set("receiveShadow", v)
}
func (sc *StereoCamera) RenderOrder() int {
	return sc.Get("renderOrder").Int()
}
func (sc *StereoCamera) SetRenderOrder(v int) {
	sc.Set("renderOrder", v)
}
func (sc *StereoCamera) Rotation() *math.Euler {
	return &math.Euler{Value: sc.Get("rotation")}
}
func (sc *StereoCamera) SetRotation(v *math.Euler) {
	sc.Set("rotation", v)
}
func (sc *StereoCamera) Scale() *math.Vector3 {
	return &math.Vector3{Value: sc.Get("scale")}
}
func (sc *StereoCamera) SetScale(v *math.Vector3) {
	sc.Set("scale", v)
}
func (sc *StereoCamera) Type() string {
	return sc.Get("type").String()
}
func (sc *StereoCamera) SetType(v string) {
	sc.Set("type", v)
}
func (sc *StereoCamera) Up() *math.Vector3 {
	return &math.Vector3{Value: sc.Get("up")}
}
func (sc *StereoCamera) SetUp(v *math.Vector3) {
	sc.Set("up", v)
}
func (sc *StereoCamera) UserData() js.Value {
	return sc.Get("userData")
}
func (sc *StereoCamera) SetUserData(v js.Value) {
	sc.Set("userData", v)
}
func (sc *StereoCamera) Uuid() string {
	return sc.Get("uuid").String()
}
func (sc *StereoCamera) SetUuid(v string) {
	sc.Set("uuid", v)
}
func (sc *StereoCamera) Visible() bool {
	return sc.Get("visible").Bool()
}
func (sc *StereoCamera) SetVisible(v bool) {
	sc.Set("visible", v)
}
func (sc *StereoCamera) DefaultMatrixAutoUpdate() bool {
	return sc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sc *StereoCamera) SetDefaultMatrixAutoUpdate(v bool) {
	sc.Set("DefaultMatrixAutoUpdate", v)
}
func (sc *StereoCamera) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: sc.Get("DefaultUp")}
}
func (sc *StereoCamera) SetDefaultUp(v *math.Vector3) {
	sc.Set("DefaultUp", v)
}
func (sc *StereoCamera) Add(object []core.Object3D) this {
	return this(sc.Call("add", object))
}
func (sc *StereoCamera) AddEventListener(typ string, listener js.Value) {
	sc.Call("addEventListener", typ, listener)
}
func (sc *StereoCamera) ApplyMatrix(matrix *math.Matrix4) {
	sc.Call("applyMatrix", matrix)
}
func (sc *StereoCamera) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(sc.Call("applyQuaternion", quaternion))
}
func (sc *StereoCamera) Clone(recursive bool) this {
	return this(sc.Call("clone", recursive))
}
func (sc *StereoCamera) Copy(source *Camera, recursive bool) this {
	return this(sc.Call("copy", source, recursive))
}
func (sc *StereoCamera) DispatchEvent(event js.Value) {
	sc.Call("dispatchEvent", event)
}
func (sc *StereoCamera) GetObjectById(id int) core.Object3D {
	return core.Object3D(sc.Call("getObjectById", id))
}
func (sc *StereoCamera) GetObjectByName(name string) core.Object3D {
	return core.Object3D(sc.Call("getObjectByName", name))
}
func (sc *StereoCamera) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(sc.Call("getObjectByProperty", name, value))
}
func (sc *StereoCamera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sc.Call("getWorldDirection", target)}
}
func (sc *StereoCamera) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sc.Call("getWorldPosition", target)}
}
func (sc *StereoCamera) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: sc.Call("getWorldQuaternion", target)}
}
func (sc *StereoCamera) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sc.Call("getWorldScale", target)}
}
func (sc *StereoCamera) HasEventListener(typ string, listener js.Value) bool {
	return sc.Call("hasEventListener", typ, listener).Bool()
}
func (sc *StereoCamera) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sc.Call("localToWorld", vector)}
}
func (sc *StereoCamera) LookAt(vector math.Vector3, y int, z int) {
	sc.Call("lookAt", vector, y, z)
}
func (sc *StereoCamera) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	sc.Call("raycast", raycaster, intersects)
}
func (sc *StereoCamera) Remove(object []core.Object3D) this {
	return this(sc.Call("remove", object))
}
func (sc *StereoCamera) RemoveEventListener(typ string, listener js.Value) {
	sc.Call("removeEventListener", typ, listener)
}
func (sc *StereoCamera) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(sc.Call("rotateOnAxis", axis, angle))
}
func (sc *StereoCamera) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(sc.Call("rotateOnWorldAxis", axis, angle))
}
func (sc *StereoCamera) RotateX(angle int) this {
	return this(sc.Call("rotateX", angle))
}
func (sc *StereoCamera) RotateY(angle int) this {
	return this(sc.Call("rotateY", angle))
}
func (sc *StereoCamera) RotateZ(angle int) this {
	return this(sc.Call("rotateZ", angle))
}
func (sc *StereoCamera) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	sc.Call("setRotationFromAxisAngle", axis, angle)
}
func (sc *StereoCamera) SetRotationFromEuler(euler *math.Euler) {
	sc.Call("setRotationFromEuler", euler)
}
func (sc *StereoCamera) SetRotationFromMatrix(m *math.Matrix4) {
	sc.Call("setRotationFromMatrix", m)
}
func (sc *StereoCamera) SetRotationFromQuaternion(q *math.Quaternion) {
	sc.Call("setRotationFromQuaternion", q)
}
func (sc *StereoCamera) ToJSON(meta js.Value) js.Value {
	return sc.Call("toJSON", meta)
}
func (sc *StereoCamera) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(sc.Call("translateOnAxis", axis, distance))
}
func (sc *StereoCamera) TranslateX(distance int) this {
	return this(sc.Call("translateX", distance))
}
func (sc *StereoCamera) TranslateY(distance int) this {
	return this(sc.Call("translateY", distance))
}
func (sc *StereoCamera) TranslateZ(distance int) this {
	return this(sc.Call("translateZ", distance))
}
func (sc *StereoCamera) Traverse(callback js.Value) {
	sc.Call("traverse", callback)
}
func (sc *StereoCamera) TraverseAncestors(callback js.Value) {
	sc.Call("traverseAncestors", callback)
}
func (sc *StereoCamera) TraverseVisible(callback js.Value) {
	sc.Call("traverseVisible", callback)
}
func (sc *StereoCamera) Update(camera *PerspectiveCamera) {
	sc.Call("update", camera)
}
func (sc *StereoCamera) UpdateMatrix() {
	sc.Call("updateMatrix")
}
func (sc *StereoCamera) UpdateMatrixWorld(force bool) {
	sc.Call("updateMatrixWorld", force)
}
func (sc *StereoCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sc *StereoCamera) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sc.Call("worldToLocal", vector)}
}

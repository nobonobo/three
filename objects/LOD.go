package objects

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type LOD struct {
	js.Value
}

func NewLOD() *LOD {
	return &LOD{Value: get("LOD").New()}
}
func (lod *LOD) CastShadow() bool {
	return lod.Get("castShadow").Bool()
}
func (lod *LOD) SetCastShadow(v bool) {
	lod.Set("castShadow", v)
}
func (lod *LOD) Children() []core.Object3D {
	return []core.Object3D(lod.Get("children"))
}
func (lod *LOD) SetChildren(v []core.Object3D) {
	lod.Set("children", v)
}
func (lod *LOD) FrustumCulled() bool {
	return lod.Get("frustumCulled").Bool()
}
func (lod *LOD) SetFrustumCulled(v bool) {
	lod.Set("frustumCulled", v)
}
func (lod *LOD) Id() int {
	return lod.Get("id").Int()
}
func (lod *LOD) SetId(v int) {
	lod.Set("id", v)
}
func (lod *LOD) IsObject3D() true {
	return true(lod.Get("isObject3D"))
}
func (lod *LOD) SetIsObject3D(v true) {
	lod.Set("isObject3D", v)
}
func (lod *LOD) Layers() *core.Layers {
	return &core.Layers{Value: lod.Get("layers")}
}
func (lod *LOD) SetLayers(v *core.Layers) {
	lod.Set("layers", v)
}
func (lod *LOD) Levels() []js.Value {
	return []js.Value(lod.Get("levels"))
}
func (lod *LOD) SetLevels(v []js.Value) {
	lod.Set("levels", v)
}
func (lod *LOD) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: lod.Get("matrix")}
}
func (lod *LOD) SetMatrix(v *math.Matrix4) {
	lod.Set("matrix", v)
}
func (lod *LOD) MatrixAutoUpdate() bool {
	return lod.Get("matrixAutoUpdate").Bool()
}
func (lod *LOD) SetMatrixAutoUpdate(v bool) {
	lod.Set("matrixAutoUpdate", v)
}
func (lod *LOD) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: lod.Get("matrixWorld")}
}
func (lod *LOD) SetMatrixWorld(v *math.Matrix4) {
	lod.Set("matrixWorld", v)
}
func (lod *LOD) MatrixWorldNeedsUpdate() bool {
	return lod.Get("matrixWorldNeedsUpdate").Bool()
}
func (lod *LOD) SetMatrixWorldNeedsUpdate(v bool) {
	lod.Set("matrixWorldNeedsUpdate", v)
}
func (lod *LOD) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: lod.Get("modelViewMatrix")}
}
func (lod *LOD) SetModelViewMatrix(v *math.Matrix4) {
	lod.Set("modelViewMatrix", v)
}
func (lod *LOD) Name() string {
	return lod.Get("name").String()
}
func (lod *LOD) SetName(v string) {
	lod.Set("name", v)
}
func (lod *LOD) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: lod.Get("normalMatrix")}
}
func (lod *LOD) SetNormalMatrix(v *math.Matrix3) {
	lod.Set("normalMatrix", v)
}
func (lod *LOD) Objects() []js.Value {
	return []js.Value(lod.Get("objects"))
}
func (lod *LOD) SetObjects(v []js.Value) {
	lod.Set("objects", v)
}
func (lod *LOD) OnAfterRender() js.Value {
	return lod.Get("onAfterRender")
}
func (lod *LOD) SetOnAfterRender(v js.Value) {
	lod.Set("onAfterRender", v)
}
func (lod *LOD) OnBeforeRender() js.Value {
	return lod.Get("onBeforeRender")
}
func (lod *LOD) SetOnBeforeRender(v js.Value) {
	lod.Set("onBeforeRender", v)
}
func (lod *LOD) Parent() core.Object3D {
	return core.Object3D(lod.Get("parent"))
}
func (lod *LOD) SetParent(v core.Object3D) {
	lod.Set("parent", v)
}
func (lod *LOD) Position() *math.Vector3 {
	return &math.Vector3{Value: lod.Get("position")}
}
func (lod *LOD) SetPosition(v *math.Vector3) {
	lod.Set("position", v)
}
func (lod *LOD) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: lod.Get("quaternion")}
}
func (lod *LOD) SetQuaternion(v *math.Quaternion) {
	lod.Set("quaternion", v)
}
func (lod *LOD) ReceiveShadow() bool {
	return lod.Get("receiveShadow").Bool()
}
func (lod *LOD) SetReceiveShadow(v bool) {
	lod.Set("receiveShadow", v)
}
func (lod *LOD) RenderOrder() int {
	return lod.Get("renderOrder").Int()
}
func (lod *LOD) SetRenderOrder(v int) {
	lod.Set("renderOrder", v)
}
func (lod *LOD) Rotation() *math.Euler {
	return &math.Euler{Value: lod.Get("rotation")}
}
func (lod *LOD) SetRotation(v *math.Euler) {
	lod.Set("rotation", v)
}
func (lod *LOD) Scale() *math.Vector3 {
	return &math.Vector3{Value: lod.Get("scale")}
}
func (lod *LOD) SetScale(v *math.Vector3) {
	lod.Set("scale", v)
}
func (lod *LOD) Type() string {
	return lod.Get("type").String()
}
func (lod *LOD) SetType(v string) {
	lod.Set("type", v)
}
func (lod *LOD) Up() *math.Vector3 {
	return &math.Vector3{Value: lod.Get("up")}
}
func (lod *LOD) SetUp(v *math.Vector3) {
	lod.Set("up", v)
}
func (lod *LOD) UserData() js.Value {
	return lod.Get("userData")
}
func (lod *LOD) SetUserData(v js.Value) {
	lod.Set("userData", v)
}
func (lod *LOD) Uuid() string {
	return lod.Get("uuid").String()
}
func (lod *LOD) SetUuid(v string) {
	lod.Set("uuid", v)
}
func (lod *LOD) Visible() bool {
	return lod.Get("visible").Bool()
}
func (lod *LOD) SetVisible(v bool) {
	lod.Set("visible", v)
}
func (lod *LOD) DefaultMatrixAutoUpdate() bool {
	return lod.Get("DefaultMatrixAutoUpdate").Bool()
}
func (lod *LOD) SetDefaultMatrixAutoUpdate(v bool) {
	lod.Set("DefaultMatrixAutoUpdate", v)
}
func (lod *LOD) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: lod.Get("DefaultUp")}
}
func (lod *LOD) SetDefaultUp(v *math.Vector3) {
	lod.Set("DefaultUp", v)
}
func (lod *LOD) Add(object []core.Object3D) this {
	return this(lod.Call("add", object))
}
func (lod *LOD) AddEventListener(typ string, listener js.Value) {
	lod.Call("addEventListener", typ, listener)
}
func (lod *LOD) AddLevel(object *core.Object3D, distance int) {
	lod.Call("addLevel", object, distance)
}
func (lod *LOD) ApplyMatrix(matrix *math.Matrix4) {
	lod.Call("applyMatrix", matrix)
}
func (lod *LOD) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(lod.Call("applyQuaternion", quaternion))
}
func (lod *LOD) Clone(recursive bool) this {
	return this(lod.Call("clone", recursive))
}
func (lod *LOD) Copy(source *core.Object3D, recursive bool) this {
	return this(lod.Call("copy", source, recursive))
}
func (lod *LOD) DispatchEvent(event js.Value) {
	lod.Call("dispatchEvent", event)
}
func (lod *LOD) GetObjectById(id int) core.Object3D {
	return core.Object3D(lod.Call("getObjectById", id))
}
func (lod *LOD) GetObjectByName(name string) core.Object3D {
	return core.Object3D(lod.Call("getObjectByName", name))
}
func (lod *LOD) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(lod.Call("getObjectByProperty", name, value))
}
func (lod *LOD) GetObjectForDistance(distance int) *core.Object3D {
	return &core.Object3D{Value: lod.Call("getObjectForDistance", distance)}
}
func (lod *LOD) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: lod.Call("getWorldDirection", target)}
}
func (lod *LOD) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: lod.Call("getWorldPosition", target)}
}
func (lod *LOD) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: lod.Call("getWorldQuaternion", target)}
}
func (lod *LOD) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: lod.Call("getWorldScale", target)}
}
func (lod *LOD) HasEventListener(typ string, listener js.Value) bool {
	return lod.Call("hasEventListener", typ, listener).Bool()
}
func (lod *LOD) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: lod.Call("localToWorld", vector)}
}
func (lod *LOD) LookAt(vector math.Vector3, y int, z int) {
	lod.Call("lookAt", vector, y, z)
}
func (lod *LOD) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	lod.Call("raycast", raycaster, intersects)
}
func (lod *LOD) Remove(object []core.Object3D) this {
	return this(lod.Call("remove", object))
}
func (lod *LOD) RemoveEventListener(typ string, listener js.Value) {
	lod.Call("removeEventListener", typ, listener)
}
func (lod *LOD) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(lod.Call("rotateOnAxis", axis, angle))
}
func (lod *LOD) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(lod.Call("rotateOnWorldAxis", axis, angle))
}
func (lod *LOD) RotateX(angle int) this {
	return this(lod.Call("rotateX", angle))
}
func (lod *LOD) RotateY(angle int) this {
	return this(lod.Call("rotateY", angle))
}
func (lod *LOD) RotateZ(angle int) this {
	return this(lod.Call("rotateZ", angle))
}
func (lod *LOD) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	lod.Call("setRotationFromAxisAngle", axis, angle)
}
func (lod *LOD) SetRotationFromEuler(euler *math.Euler) {
	lod.Call("setRotationFromEuler", euler)
}
func (lod *LOD) SetRotationFromMatrix(m *math.Matrix4) {
	lod.Call("setRotationFromMatrix", m)
}
func (lod *LOD) SetRotationFromQuaternion(q *math.Quaternion) {
	lod.Call("setRotationFromQuaternion", q)
}
func (lod *LOD) ToJSON(meta js.Value) js.Value {
	return lod.Call("toJSON", meta)
}
func (lod *LOD) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(lod.Call("translateOnAxis", axis, distance))
}
func (lod *LOD) TranslateX(distance int) this {
	return this(lod.Call("translateX", distance))
}
func (lod *LOD) TranslateY(distance int) this {
	return this(lod.Call("translateY", distance))
}
func (lod *LOD) TranslateZ(distance int) this {
	return this(lod.Call("translateZ", distance))
}
func (lod *LOD) Traverse(callback js.Value) {
	lod.Call("traverse", callback)
}
func (lod *LOD) TraverseAncestors(callback js.Value) {
	lod.Call("traverseAncestors", callback)
}
func (lod *LOD) TraverseVisible(callback js.Value) {
	lod.Call("traverseVisible", callback)
}
func (lod *LOD) Update(camera *cameras.Camera) {
	lod.Call("update", camera)
}
func (lod *LOD) UpdateMatrix() {
	lod.Call("updateMatrix")
}
func (lod *LOD) UpdateMatrixWorld(force bool) {
	lod.Call("updateMatrixWorld", force)
}
func (lod *LOD) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	lod.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (lod *LOD) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: lod.Call("worldToLocal", vector)}
}

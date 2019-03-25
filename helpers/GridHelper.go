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

type GridHelper struct {
	js.Value
}

func NewGridHelper(size int, divisions int, color1 math.Color, color2 math.Color) *GridHelper {
	return &GridHelper{Value: get("GridHelper").New(size, divisions, color1, color2)}
}
func (gh *GridHelper) CastShadow() bool {
	return gh.Get("castShadow").Bool()
}
func (gh *GridHelper) SetCastShadow(v bool) {
	gh.Set("castShadow", v)
}
func (gh *GridHelper) Children() []core.Object3D {
	return []core.Object3D(gh.Get("children"))
}
func (gh *GridHelper) SetChildren(v []core.Object3D) {
	gh.Set("children", v)
}
func (gh *GridHelper) FrustumCulled() bool {
	return gh.Get("frustumCulled").Bool()
}
func (gh *GridHelper) SetFrustumCulled(v bool) {
	gh.Set("frustumCulled", v)
}
func (gh *GridHelper) Geometry() core.Geometry {
	return core.Geometry(gh.Get("geometry"))
}
func (gh *GridHelper) SetGeometry(v core.Geometry) {
	gh.Set("geometry", v)
}
func (gh *GridHelper) Id() int {
	return gh.Get("id").Int()
}
func (gh *GridHelper) SetId(v int) {
	gh.Set("id", v)
}
func (gh *GridHelper) IsLine() true {
	return true(gh.Get("isLine"))
}
func (gh *GridHelper) SetIsLine(v true) {
	gh.Set("isLine", v)
}
func (gh *GridHelper) IsLineSegments() true {
	return true(gh.Get("isLineSegments"))
}
func (gh *GridHelper) SetIsLineSegments(v true) {
	gh.Set("isLineSegments", v)
}
func (gh *GridHelper) IsObject3D() true {
	return true(gh.Get("isObject3D"))
}
func (gh *GridHelper) SetIsObject3D(v true) {
	gh.Set("isObject3D", v)
}
func (gh *GridHelper) Layers() *core.Layers {
	return &core.Layers{Value: gh.Get("layers")}
}
func (gh *GridHelper) SetLayers(v *core.Layers) {
	gh.Set("layers", v)
}
func (gh *GridHelper) Material() materials.Material {
	return materials.Material(gh.Get("material"))
}
func (gh *GridHelper) SetMaterial(v materials.Material) {
	gh.Set("material", v)
}
func (gh *GridHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: gh.Get("matrix")}
}
func (gh *GridHelper) SetMatrix(v *math.Matrix4) {
	gh.Set("matrix", v)
}
func (gh *GridHelper) MatrixAutoUpdate() bool {
	return gh.Get("matrixAutoUpdate").Bool()
}
func (gh *GridHelper) SetMatrixAutoUpdate(v bool) {
	gh.Set("matrixAutoUpdate", v)
}
func (gh *GridHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: gh.Get("matrixWorld")}
}
func (gh *GridHelper) SetMatrixWorld(v *math.Matrix4) {
	gh.Set("matrixWorld", v)
}
func (gh *GridHelper) MatrixWorldNeedsUpdate() bool {
	return gh.Get("matrixWorldNeedsUpdate").Bool()
}
func (gh *GridHelper) SetMatrixWorldNeedsUpdate(v bool) {
	gh.Set("matrixWorldNeedsUpdate", v)
}
func (gh *GridHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: gh.Get("modelViewMatrix")}
}
func (gh *GridHelper) SetModelViewMatrix(v *math.Matrix4) {
	gh.Set("modelViewMatrix", v)
}
func (gh *GridHelper) Name() string {
	return gh.Get("name").String()
}
func (gh *GridHelper) SetName(v string) {
	gh.Set("name", v)
}
func (gh *GridHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: gh.Get("normalMatrix")}
}
func (gh *GridHelper) SetNormalMatrix(v *math.Matrix3) {
	gh.Set("normalMatrix", v)
}
func (gh *GridHelper) OnAfterRender() js.Value {
	return gh.Get("onAfterRender")
}
func (gh *GridHelper) SetOnAfterRender(v js.Value) {
	gh.Set("onAfterRender", v)
}
func (gh *GridHelper) OnBeforeRender() js.Value {
	return gh.Get("onBeforeRender")
}
func (gh *GridHelper) SetOnBeforeRender(v js.Value) {
	gh.Set("onBeforeRender", v)
}
func (gh *GridHelper) Parent() core.Object3D {
	return core.Object3D(gh.Get("parent"))
}
func (gh *GridHelper) SetParent(v core.Object3D) {
	gh.Set("parent", v)
}
func (gh *GridHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: gh.Get("position")}
}
func (gh *GridHelper) SetPosition(v *math.Vector3) {
	gh.Set("position", v)
}
func (gh *GridHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: gh.Get("quaternion")}
}
func (gh *GridHelper) SetQuaternion(v *math.Quaternion) {
	gh.Set("quaternion", v)
}
func (gh *GridHelper) ReceiveShadow() bool {
	return gh.Get("receiveShadow").Bool()
}
func (gh *GridHelper) SetReceiveShadow(v bool) {
	gh.Set("receiveShadow", v)
}
func (gh *GridHelper) RenderOrder() int {
	return gh.Get("renderOrder").Int()
}
func (gh *GridHelper) SetRenderOrder(v int) {
	gh.Set("renderOrder", v)
}
func (gh *GridHelper) Rotation() *math.Euler {
	return &math.Euler{Value: gh.Get("rotation")}
}
func (gh *GridHelper) SetRotation(v *math.Euler) {
	gh.Set("rotation", v)
}
func (gh *GridHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: gh.Get("scale")}
}
func (gh *GridHelper) SetScale(v *math.Vector3) {
	gh.Set("scale", v)
}
func (gh *GridHelper) Type() string {
	return gh.Get("type").String()
}
func (gh *GridHelper) SetType(v string) {
	gh.Set("type", v)
}
func (gh *GridHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: gh.Get("up")}
}
func (gh *GridHelper) SetUp(v *math.Vector3) {
	gh.Set("up", v)
}
func (gh *GridHelper) UserData() js.Value {
	return gh.Get("userData")
}
func (gh *GridHelper) SetUserData(v js.Value) {
	gh.Set("userData", v)
}
func (gh *GridHelper) Uuid() string {
	return gh.Get("uuid").String()
}
func (gh *GridHelper) SetUuid(v string) {
	gh.Set("uuid", v)
}
func (gh *GridHelper) Visible() bool {
	return gh.Get("visible").Bool()
}
func (gh *GridHelper) SetVisible(v bool) {
	gh.Set("visible", v)
}
func (gh *GridHelper) DefaultMatrixAutoUpdate() bool {
	return gh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (gh *GridHelper) SetDefaultMatrixAutoUpdate(v bool) {
	gh.Set("DefaultMatrixAutoUpdate", v)
}
func (gh *GridHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: gh.Get("DefaultUp")}
}
func (gh *GridHelper) SetDefaultUp(v *math.Vector3) {
	gh.Set("DefaultUp", v)
}
func (gh *GridHelper) Add(object []core.Object3D) this {
	return this(gh.Call("add", object))
}
func (gh *GridHelper) AddEventListener(typ string, listener js.Value) {
	gh.Call("addEventListener", typ, listener)
}
func (gh *GridHelper) ApplyMatrix(matrix *math.Matrix4) {
	gh.Call("applyMatrix", matrix)
}
func (gh *GridHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(gh.Call("applyQuaternion", quaternion))
}
func (gh *GridHelper) Clone(recursive bool) this {
	return this(gh.Call("clone", recursive))
}
func (gh *GridHelper) ComputeLineDistances() this {
	return this(gh.Call("computeLineDistances"))
}
func (gh *GridHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(gh.Call("copy", source, recursive))
}
func (gh *GridHelper) DispatchEvent(event js.Value) {
	gh.Call("dispatchEvent", event)
}
func (gh *GridHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(gh.Call("getObjectById", id))
}
func (gh *GridHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(gh.Call("getObjectByName", name))
}
func (gh *GridHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(gh.Call("getObjectByProperty", name, value))
}
func (gh *GridHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: gh.Call("getWorldDirection", target)}
}
func (gh *GridHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: gh.Call("getWorldPosition", target)}
}
func (gh *GridHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: gh.Call("getWorldQuaternion", target)}
}
func (gh *GridHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: gh.Call("getWorldScale", target)}
}
func (gh *GridHelper) HasEventListener(typ string, listener js.Value) bool {
	return gh.Call("hasEventListener", typ, listener).Bool()
}
func (gh *GridHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: gh.Call("localToWorld", vector)}
}
func (gh *GridHelper) LookAt(vector math.Vector3, y int, z int) {
	gh.Call("lookAt", vector, y, z)
}
func (gh *GridHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	gh.Call("raycast", raycaster, intersects)
}
func (gh *GridHelper) Remove(object []core.Object3D) this {
	return this(gh.Call("remove", object))
}
func (gh *GridHelper) RemoveEventListener(typ string, listener js.Value) {
	gh.Call("removeEventListener", typ, listener)
}
func (gh *GridHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(gh.Call("rotateOnAxis", axis, angle))
}
func (gh *GridHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(gh.Call("rotateOnWorldAxis", axis, angle))
}
func (gh *GridHelper) RotateX(angle int) this {
	return this(gh.Call("rotateX", angle))
}
func (gh *GridHelper) RotateY(angle int) this {
	return this(gh.Call("rotateY", angle))
}
func (gh *GridHelper) RotateZ(angle int) this {
	return this(gh.Call("rotateZ", angle))
}
func (gh *GridHelper) SetColors(color1 math.Color, color2 math.Color) {
	gh.Call("setColors", color1, color2)
}
func (gh *GridHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	gh.Call("setRotationFromAxisAngle", axis, angle)
}
func (gh *GridHelper) SetRotationFromEuler(euler *math.Euler) {
	gh.Call("setRotationFromEuler", euler)
}
func (gh *GridHelper) SetRotationFromMatrix(m *math.Matrix4) {
	gh.Call("setRotationFromMatrix", m)
}
func (gh *GridHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	gh.Call("setRotationFromQuaternion", q)
}
func (gh *GridHelper) ToJSON(meta js.Value) js.Value {
	return gh.Call("toJSON", meta)
}
func (gh *GridHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(gh.Call("translateOnAxis", axis, distance))
}
func (gh *GridHelper) TranslateX(distance int) this {
	return this(gh.Call("translateX", distance))
}
func (gh *GridHelper) TranslateY(distance int) this {
	return this(gh.Call("translateY", distance))
}
func (gh *GridHelper) TranslateZ(distance int) this {
	return this(gh.Call("translateZ", distance))
}
func (gh *GridHelper) Traverse(callback js.Value) {
	gh.Call("traverse", callback)
}
func (gh *GridHelper) TraverseAncestors(callback js.Value) {
	gh.Call("traverseAncestors", callback)
}
func (gh *GridHelper) TraverseVisible(callback js.Value) {
	gh.Call("traverseVisible", callback)
}
func (gh *GridHelper) UpdateMatrix() {
	gh.Call("updateMatrix")
}
func (gh *GridHelper) UpdateMatrixWorld(force bool) {
	gh.Call("updateMatrixWorld", force)
}
func (gh *GridHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	gh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (gh *GridHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: gh.Call("worldToLocal", vector)}
}

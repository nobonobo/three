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

type AxesHelper struct {
	js.Value
}

func NewAxesHelper(size int) *AxesHelper {
	return &AxesHelper{Value: get("AxesHelper").New(size)}
}
func (ah *AxesHelper) CastShadow() bool {
	return ah.Get("castShadow").Bool()
}
func (ah *AxesHelper) SetCastShadow(v bool) {
	ah.Set("castShadow", v)
}
func (ah *AxesHelper) Children() []core.Object3D {
	return []core.Object3D(ah.Get("children"))
}
func (ah *AxesHelper) SetChildren(v []core.Object3D) {
	ah.Set("children", v)
}
func (ah *AxesHelper) FrustumCulled() bool {
	return ah.Get("frustumCulled").Bool()
}
func (ah *AxesHelper) SetFrustumCulled(v bool) {
	ah.Set("frustumCulled", v)
}
func (ah *AxesHelper) Geometry() core.Geometry {
	return core.Geometry(ah.Get("geometry"))
}
func (ah *AxesHelper) SetGeometry(v core.Geometry) {
	ah.Set("geometry", v)
}
func (ah *AxesHelper) Id() int {
	return ah.Get("id").Int()
}
func (ah *AxesHelper) SetId(v int) {
	ah.Set("id", v)
}
func (ah *AxesHelper) IsLine() true {
	return true(ah.Get("isLine"))
}
func (ah *AxesHelper) SetIsLine(v true) {
	ah.Set("isLine", v)
}
func (ah *AxesHelper) IsLineSegments() true {
	return true(ah.Get("isLineSegments"))
}
func (ah *AxesHelper) SetIsLineSegments(v true) {
	ah.Set("isLineSegments", v)
}
func (ah *AxesHelper) IsObject3D() true {
	return true(ah.Get("isObject3D"))
}
func (ah *AxesHelper) SetIsObject3D(v true) {
	ah.Set("isObject3D", v)
}
func (ah *AxesHelper) Layers() *core.Layers {
	return &core.Layers{Value: ah.Get("layers")}
}
func (ah *AxesHelper) SetLayers(v *core.Layers) {
	ah.Set("layers", v)
}
func (ah *AxesHelper) Material() materials.Material {
	return materials.Material(ah.Get("material"))
}
func (ah *AxesHelper) SetMaterial(v materials.Material) {
	ah.Set("material", v)
}
func (ah *AxesHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ah.Get("matrix")}
}
func (ah *AxesHelper) SetMatrix(v *math.Matrix4) {
	ah.Set("matrix", v)
}
func (ah *AxesHelper) MatrixAutoUpdate() bool {
	return ah.Get("matrixAutoUpdate").Bool()
}
func (ah *AxesHelper) SetMatrixAutoUpdate(v bool) {
	ah.Set("matrixAutoUpdate", v)
}
func (ah *AxesHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ah.Get("matrixWorld")}
}
func (ah *AxesHelper) SetMatrixWorld(v *math.Matrix4) {
	ah.Set("matrixWorld", v)
}
func (ah *AxesHelper) MatrixWorldNeedsUpdate() bool {
	return ah.Get("matrixWorldNeedsUpdate").Bool()
}
func (ah *AxesHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ah.Set("matrixWorldNeedsUpdate", v)
}
func (ah *AxesHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ah.Get("modelViewMatrix")}
}
func (ah *AxesHelper) SetModelViewMatrix(v *math.Matrix4) {
	ah.Set("modelViewMatrix", v)
}
func (ah *AxesHelper) Name() string {
	return ah.Get("name").String()
}
func (ah *AxesHelper) SetName(v string) {
	ah.Set("name", v)
}
func (ah *AxesHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ah.Get("normalMatrix")}
}
func (ah *AxesHelper) SetNormalMatrix(v *math.Matrix3) {
	ah.Set("normalMatrix", v)
}
func (ah *AxesHelper) OnAfterRender() js.Value {
	return ah.Get("onAfterRender")
}
func (ah *AxesHelper) SetOnAfterRender(v js.Value) {
	ah.Set("onAfterRender", v)
}
func (ah *AxesHelper) OnBeforeRender() js.Value {
	return ah.Get("onBeforeRender")
}
func (ah *AxesHelper) SetOnBeforeRender(v js.Value) {
	ah.Set("onBeforeRender", v)
}
func (ah *AxesHelper) Parent() core.Object3D {
	return core.Object3D(ah.Get("parent"))
}
func (ah *AxesHelper) SetParent(v core.Object3D) {
	ah.Set("parent", v)
}
func (ah *AxesHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("position")}
}
func (ah *AxesHelper) SetPosition(v *math.Vector3) {
	ah.Set("position", v)
}
func (ah *AxesHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ah.Get("quaternion")}
}
func (ah *AxesHelper) SetQuaternion(v *math.Quaternion) {
	ah.Set("quaternion", v)
}
func (ah *AxesHelper) ReceiveShadow() bool {
	return ah.Get("receiveShadow").Bool()
}
func (ah *AxesHelper) SetReceiveShadow(v bool) {
	ah.Set("receiveShadow", v)
}
func (ah *AxesHelper) RenderOrder() int {
	return ah.Get("renderOrder").Int()
}
func (ah *AxesHelper) SetRenderOrder(v int) {
	ah.Set("renderOrder", v)
}
func (ah *AxesHelper) Rotation() *math.Euler {
	return &math.Euler{Value: ah.Get("rotation")}
}
func (ah *AxesHelper) SetRotation(v *math.Euler) {
	ah.Set("rotation", v)
}
func (ah *AxesHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("scale")}
}
func (ah *AxesHelper) SetScale(v *math.Vector3) {
	ah.Set("scale", v)
}
func (ah *AxesHelper) Type() string {
	return ah.Get("type").String()
}
func (ah *AxesHelper) SetType(v string) {
	ah.Set("type", v)
}
func (ah *AxesHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("up")}
}
func (ah *AxesHelper) SetUp(v *math.Vector3) {
	ah.Set("up", v)
}
func (ah *AxesHelper) UserData() js.Value {
	return ah.Get("userData")
}
func (ah *AxesHelper) SetUserData(v js.Value) {
	ah.Set("userData", v)
}
func (ah *AxesHelper) Uuid() string {
	return ah.Get("uuid").String()
}
func (ah *AxesHelper) SetUuid(v string) {
	ah.Set("uuid", v)
}
func (ah *AxesHelper) Visible() bool {
	return ah.Get("visible").Bool()
}
func (ah *AxesHelper) SetVisible(v bool) {
	ah.Set("visible", v)
}
func (ah *AxesHelper) DefaultMatrixAutoUpdate() bool {
	return ah.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ah *AxesHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ah.Set("DefaultMatrixAutoUpdate", v)
}
func (ah *AxesHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("DefaultUp")}
}
func (ah *AxesHelper) SetDefaultUp(v *math.Vector3) {
	ah.Set("DefaultUp", v)
}
func (ah *AxesHelper) Add(object []core.Object3D) this {
	return this(ah.Call("add", object))
}
func (ah *AxesHelper) AddEventListener(typ string, listener js.Value) {
	ah.Call("addEventListener", typ, listener)
}
func (ah *AxesHelper) ApplyMatrix(matrix *math.Matrix4) {
	ah.Call("applyMatrix", matrix)
}
func (ah *AxesHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ah.Call("applyQuaternion", quaternion))
}
func (ah *AxesHelper) Clone(recursive bool) this {
	return this(ah.Call("clone", recursive))
}
func (ah *AxesHelper) ComputeLineDistances() this {
	return this(ah.Call("computeLineDistances"))
}
func (ah *AxesHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(ah.Call("copy", source, recursive))
}
func (ah *AxesHelper) DispatchEvent(event js.Value) {
	ah.Call("dispatchEvent", event)
}
func (ah *AxesHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(ah.Call("getObjectById", id))
}
func (ah *AxesHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ah.Call("getObjectByName", name))
}
func (ah *AxesHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ah.Call("getObjectByProperty", name, value))
}
func (ah *AxesHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("getWorldDirection", target)}
}
func (ah *AxesHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("getWorldPosition", target)}
}
func (ah *AxesHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ah.Call("getWorldQuaternion", target)}
}
func (ah *AxesHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("getWorldScale", target)}
}
func (ah *AxesHelper) HasEventListener(typ string, listener js.Value) bool {
	return ah.Call("hasEventListener", typ, listener).Bool()
}
func (ah *AxesHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("localToWorld", vector)}
}
func (ah *AxesHelper) LookAt(vector math.Vector3, y int, z int) {
	ah.Call("lookAt", vector, y, z)
}
func (ah *AxesHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ah.Call("raycast", raycaster, intersects)
}
func (ah *AxesHelper) Remove(object []core.Object3D) this {
	return this(ah.Call("remove", object))
}
func (ah *AxesHelper) RemoveEventListener(typ string, listener js.Value) {
	ah.Call("removeEventListener", typ, listener)
}
func (ah *AxesHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ah.Call("rotateOnAxis", axis, angle))
}
func (ah *AxesHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ah.Call("rotateOnWorldAxis", axis, angle))
}
func (ah *AxesHelper) RotateX(angle int) this {
	return this(ah.Call("rotateX", angle))
}
func (ah *AxesHelper) RotateY(angle int) this {
	return this(ah.Call("rotateY", angle))
}
func (ah *AxesHelper) RotateZ(angle int) this {
	return this(ah.Call("rotateZ", angle))
}
func (ah *AxesHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ah.Call("setRotationFromAxisAngle", axis, angle)
}
func (ah *AxesHelper) SetRotationFromEuler(euler *math.Euler) {
	ah.Call("setRotationFromEuler", euler)
}
func (ah *AxesHelper) SetRotationFromMatrix(m *math.Matrix4) {
	ah.Call("setRotationFromMatrix", m)
}
func (ah *AxesHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	ah.Call("setRotationFromQuaternion", q)
}
func (ah *AxesHelper) ToJSON(meta js.Value) js.Value {
	return ah.Call("toJSON", meta)
}
func (ah *AxesHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ah.Call("translateOnAxis", axis, distance))
}
func (ah *AxesHelper) TranslateX(distance int) this {
	return this(ah.Call("translateX", distance))
}
func (ah *AxesHelper) TranslateY(distance int) this {
	return this(ah.Call("translateY", distance))
}
func (ah *AxesHelper) TranslateZ(distance int) this {
	return this(ah.Call("translateZ", distance))
}
func (ah *AxesHelper) Traverse(callback js.Value) {
	ah.Call("traverse", callback)
}
func (ah *AxesHelper) TraverseAncestors(callback js.Value) {
	ah.Call("traverseAncestors", callback)
}
func (ah *AxesHelper) TraverseVisible(callback js.Value) {
	ah.Call("traverseVisible", callback)
}
func (ah *AxesHelper) UpdateMatrix() {
	ah.Call("updateMatrix")
}
func (ah *AxesHelper) UpdateMatrixWorld(force bool) {
	ah.Call("updateMatrixWorld", force)
}
func (ah *AxesHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ah.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ah *AxesHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("worldToLocal", vector)}
}

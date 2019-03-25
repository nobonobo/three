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

type BoxHelper struct {
	js.Value
}

func NewBoxHelper(object *core.Object3D, color *math.Color) *BoxHelper {
	return &BoxHelper{Value: get("BoxHelper").New(object, color)}
}
func (bh *BoxHelper) CastShadow() bool {
	return bh.Get("castShadow").Bool()
}
func (bh *BoxHelper) SetCastShadow(v bool) {
	bh.Set("castShadow", v)
}
func (bh *BoxHelper) Children() []core.Object3D {
	return []core.Object3D(bh.Get("children"))
}
func (bh *BoxHelper) SetChildren(v []core.Object3D) {
	bh.Set("children", v)
}
func (bh *BoxHelper) FrustumCulled() bool {
	return bh.Get("frustumCulled").Bool()
}
func (bh *BoxHelper) SetFrustumCulled(v bool) {
	bh.Set("frustumCulled", v)
}
func (bh *BoxHelper) Geometry() core.Geometry {
	return core.Geometry(bh.Get("geometry"))
}
func (bh *BoxHelper) SetGeometry(v core.Geometry) {
	bh.Set("geometry", v)
}
func (bh *BoxHelper) Id() int {
	return bh.Get("id").Int()
}
func (bh *BoxHelper) SetId(v int) {
	bh.Set("id", v)
}
func (bh *BoxHelper) IsLine() true {
	return true(bh.Get("isLine"))
}
func (bh *BoxHelper) SetIsLine(v true) {
	bh.Set("isLine", v)
}
func (bh *BoxHelper) IsLineSegments() true {
	return true(bh.Get("isLineSegments"))
}
func (bh *BoxHelper) SetIsLineSegments(v true) {
	bh.Set("isLineSegments", v)
}
func (bh *BoxHelper) IsObject3D() true {
	return true(bh.Get("isObject3D"))
}
func (bh *BoxHelper) SetIsObject3D(v true) {
	bh.Set("isObject3D", v)
}
func (bh *BoxHelper) Layers() *core.Layers {
	return &core.Layers{Value: bh.Get("layers")}
}
func (bh *BoxHelper) SetLayers(v *core.Layers) {
	bh.Set("layers", v)
}
func (bh *BoxHelper) Material() materials.Material {
	return materials.Material(bh.Get("material"))
}
func (bh *BoxHelper) SetMaterial(v materials.Material) {
	bh.Set("material", v)
}
func (bh *BoxHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: bh.Get("matrix")}
}
func (bh *BoxHelper) SetMatrix(v *math.Matrix4) {
	bh.Set("matrix", v)
}
func (bh *BoxHelper) MatrixAutoUpdate() bool {
	return bh.Get("matrixAutoUpdate").Bool()
}
func (bh *BoxHelper) SetMatrixAutoUpdate(v bool) {
	bh.Set("matrixAutoUpdate", v)
}
func (bh *BoxHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: bh.Get("matrixWorld")}
}
func (bh *BoxHelper) SetMatrixWorld(v *math.Matrix4) {
	bh.Set("matrixWorld", v)
}
func (bh *BoxHelper) MatrixWorldNeedsUpdate() bool {
	return bh.Get("matrixWorldNeedsUpdate").Bool()
}
func (bh *BoxHelper) SetMatrixWorldNeedsUpdate(v bool) {
	bh.Set("matrixWorldNeedsUpdate", v)
}
func (bh *BoxHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: bh.Get("modelViewMatrix")}
}
func (bh *BoxHelper) SetModelViewMatrix(v *math.Matrix4) {
	bh.Set("modelViewMatrix", v)
}
func (bh *BoxHelper) Name() string {
	return bh.Get("name").String()
}
func (bh *BoxHelper) SetName(v string) {
	bh.Set("name", v)
}
func (bh *BoxHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: bh.Get("normalMatrix")}
}
func (bh *BoxHelper) SetNormalMatrix(v *math.Matrix3) {
	bh.Set("normalMatrix", v)
}
func (bh *BoxHelper) OnAfterRender() js.Value {
	return bh.Get("onAfterRender")
}
func (bh *BoxHelper) SetOnAfterRender(v js.Value) {
	bh.Set("onAfterRender", v)
}
func (bh *BoxHelper) OnBeforeRender() js.Value {
	return bh.Get("onBeforeRender")
}
func (bh *BoxHelper) SetOnBeforeRender(v js.Value) {
	bh.Set("onBeforeRender", v)
}
func (bh *BoxHelper) Parent() core.Object3D {
	return core.Object3D(bh.Get("parent"))
}
func (bh *BoxHelper) SetParent(v core.Object3D) {
	bh.Set("parent", v)
}
func (bh *BoxHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: bh.Get("position")}
}
func (bh *BoxHelper) SetPosition(v *math.Vector3) {
	bh.Set("position", v)
}
func (bh *BoxHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: bh.Get("quaternion")}
}
func (bh *BoxHelper) SetQuaternion(v *math.Quaternion) {
	bh.Set("quaternion", v)
}
func (bh *BoxHelper) ReceiveShadow() bool {
	return bh.Get("receiveShadow").Bool()
}
func (bh *BoxHelper) SetReceiveShadow(v bool) {
	bh.Set("receiveShadow", v)
}
func (bh *BoxHelper) RenderOrder() int {
	return bh.Get("renderOrder").Int()
}
func (bh *BoxHelper) SetRenderOrder(v int) {
	bh.Set("renderOrder", v)
}
func (bh *BoxHelper) Rotation() *math.Euler {
	return &math.Euler{Value: bh.Get("rotation")}
}
func (bh *BoxHelper) SetRotation(v *math.Euler) {
	bh.Set("rotation", v)
}
func (bh *BoxHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: bh.Get("scale")}
}
func (bh *BoxHelper) SetScale(v *math.Vector3) {
	bh.Set("scale", v)
}
func (bh *BoxHelper) Type() string {
	return bh.Get("type").String()
}
func (bh *BoxHelper) SetType(v string) {
	bh.Set("type", v)
}
func (bh *BoxHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: bh.Get("up")}
}
func (bh *BoxHelper) SetUp(v *math.Vector3) {
	bh.Set("up", v)
}
func (bh *BoxHelper) UserData() js.Value {
	return bh.Get("userData")
}
func (bh *BoxHelper) SetUserData(v js.Value) {
	bh.Set("userData", v)
}
func (bh *BoxHelper) Uuid() string {
	return bh.Get("uuid").String()
}
func (bh *BoxHelper) SetUuid(v string) {
	bh.Set("uuid", v)
}
func (bh *BoxHelper) Visible() bool {
	return bh.Get("visible").Bool()
}
func (bh *BoxHelper) SetVisible(v bool) {
	bh.Set("visible", v)
}
func (bh *BoxHelper) DefaultMatrixAutoUpdate() bool {
	return bh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (bh *BoxHelper) SetDefaultMatrixAutoUpdate(v bool) {
	bh.Set("DefaultMatrixAutoUpdate", v)
}
func (bh *BoxHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: bh.Get("DefaultUp")}
}
func (bh *BoxHelper) SetDefaultUp(v *math.Vector3) {
	bh.Set("DefaultUp", v)
}
func (bh *BoxHelper) Add(object []core.Object3D) this {
	return this(bh.Call("add", object))
}
func (bh *BoxHelper) AddEventListener(typ string, listener js.Value) {
	bh.Call("addEventListener", typ, listener)
}
func (bh *BoxHelper) ApplyMatrix(matrix *math.Matrix4) {
	bh.Call("applyMatrix", matrix)
}
func (bh *BoxHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(bh.Call("applyQuaternion", quaternion))
}
func (bh *BoxHelper) Clone(recursive bool) this {
	return this(bh.Call("clone", recursive))
}
func (bh *BoxHelper) ComputeLineDistances() this {
	return this(bh.Call("computeLineDistances"))
}
func (bh *BoxHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(bh.Call("copy", source, recursive))
}
func (bh *BoxHelper) DispatchEvent(event js.Value) {
	bh.Call("dispatchEvent", event)
}
func (bh *BoxHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(bh.Call("getObjectById", id))
}
func (bh *BoxHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(bh.Call("getObjectByName", name))
}
func (bh *BoxHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(bh.Call("getObjectByProperty", name, value))
}
func (bh *BoxHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: bh.Call("getWorldDirection", target)}
}
func (bh *BoxHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: bh.Call("getWorldPosition", target)}
}
func (bh *BoxHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: bh.Call("getWorldQuaternion", target)}
}
func (bh *BoxHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: bh.Call("getWorldScale", target)}
}
func (bh *BoxHelper) HasEventListener(typ string, listener js.Value) bool {
	return bh.Call("hasEventListener", typ, listener).Bool()
}
func (bh *BoxHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: bh.Call("localToWorld", vector)}
}
func (bh *BoxHelper) LookAt(vector math.Vector3, y int, z int) {
	bh.Call("lookAt", vector, y, z)
}
func (bh *BoxHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	bh.Call("raycast", raycaster, intersects)
}
func (bh *BoxHelper) Remove(object []core.Object3D) this {
	return this(bh.Call("remove", object))
}
func (bh *BoxHelper) RemoveEventListener(typ string, listener js.Value) {
	bh.Call("removeEventListener", typ, listener)
}
func (bh *BoxHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(bh.Call("rotateOnAxis", axis, angle))
}
func (bh *BoxHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(bh.Call("rotateOnWorldAxis", axis, angle))
}
func (bh *BoxHelper) RotateX(angle int) this {
	return this(bh.Call("rotateX", angle))
}
func (bh *BoxHelper) RotateY(angle int) this {
	return this(bh.Call("rotateY", angle))
}
func (bh *BoxHelper) RotateZ(angle int) this {
	return this(bh.Call("rotateZ", angle))
}
func (bh *BoxHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	bh.Call("setRotationFromAxisAngle", axis, angle)
}
func (bh *BoxHelper) SetRotationFromEuler(euler *math.Euler) {
	bh.Call("setRotationFromEuler", euler)
}
func (bh *BoxHelper) SetRotationFromMatrix(m *math.Matrix4) {
	bh.Call("setRotationFromMatrix", m)
}
func (bh *BoxHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	bh.Call("setRotationFromQuaternion", q)
}
func (bh *BoxHelper) ToJSON(meta js.Value) js.Value {
	return bh.Call("toJSON", meta)
}
func (bh *BoxHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(bh.Call("translateOnAxis", axis, distance))
}
func (bh *BoxHelper) TranslateX(distance int) this {
	return this(bh.Call("translateX", distance))
}
func (bh *BoxHelper) TranslateY(distance int) this {
	return this(bh.Call("translateY", distance))
}
func (bh *BoxHelper) TranslateZ(distance int) this {
	return this(bh.Call("translateZ", distance))
}
func (bh *BoxHelper) Traverse(callback js.Value) {
	bh.Call("traverse", callback)
}
func (bh *BoxHelper) TraverseAncestors(callback js.Value) {
	bh.Call("traverseAncestors", callback)
}
func (bh *BoxHelper) TraverseVisible(callback js.Value) {
	bh.Call("traverseVisible", callback)
}
func (bh *BoxHelper) Update(object *core.Object3D) {
	bh.Call("update", object)
}
func (bh *BoxHelper) UpdateMatrix() {
	bh.Call("updateMatrix")
}
func (bh *BoxHelper) UpdateMatrixWorld(force bool) {
	bh.Call("updateMatrixWorld", force)
}
func (bh *BoxHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	bh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (bh *BoxHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: bh.Call("worldToLocal", vector)}
}

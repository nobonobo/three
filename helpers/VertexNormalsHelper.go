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

type VertexNormalsHelper struct {
	js.Value
}

func NewVertexNormalsHelper(object *core.Object3D, size int, hex int, linewidth int) *VertexNormalsHelper {
	return &VertexNormalsHelper{Value: get("VertexNormalsHelper").New(object, size, hex, linewidth)}
}
func (vnh *VertexNormalsHelper) CastShadow() bool {
	return vnh.Get("castShadow").Bool()
}
func (vnh *VertexNormalsHelper) SetCastShadow(v bool) {
	vnh.Set("castShadow", v)
}
func (vnh *VertexNormalsHelper) Children() []core.Object3D {
	return []core.Object3D(vnh.Get("children"))
}
func (vnh *VertexNormalsHelper) SetChildren(v []core.Object3D) {
	vnh.Set("children", v)
}
func (vnh *VertexNormalsHelper) FrustumCulled() bool {
	return vnh.Get("frustumCulled").Bool()
}
func (vnh *VertexNormalsHelper) SetFrustumCulled(v bool) {
	vnh.Set("frustumCulled", v)
}
func (vnh *VertexNormalsHelper) Geometry() core.Geometry {
	return core.Geometry(vnh.Get("geometry"))
}
func (vnh *VertexNormalsHelper) SetGeometry(v core.Geometry) {
	vnh.Set("geometry", v)
}
func (vnh *VertexNormalsHelper) Id() int {
	return vnh.Get("id").Int()
}
func (vnh *VertexNormalsHelper) SetId(v int) {
	vnh.Set("id", v)
}
func (vnh *VertexNormalsHelper) IsLine() true {
	return true(vnh.Get("isLine"))
}
func (vnh *VertexNormalsHelper) SetIsLine(v true) {
	vnh.Set("isLine", v)
}
func (vnh *VertexNormalsHelper) IsLineSegments() true {
	return true(vnh.Get("isLineSegments"))
}
func (vnh *VertexNormalsHelper) SetIsLineSegments(v true) {
	vnh.Set("isLineSegments", v)
}
func (vnh *VertexNormalsHelper) IsObject3D() true {
	return true(vnh.Get("isObject3D"))
}
func (vnh *VertexNormalsHelper) SetIsObject3D(v true) {
	vnh.Set("isObject3D", v)
}
func (vnh *VertexNormalsHelper) Layers() *core.Layers {
	return &core.Layers{Value: vnh.Get("layers")}
}
func (vnh *VertexNormalsHelper) SetLayers(v *core.Layers) {
	vnh.Set("layers", v)
}
func (vnh *VertexNormalsHelper) Material() materials.Material {
	return materials.Material(vnh.Get("material"))
}
func (vnh *VertexNormalsHelper) SetMaterial(v materials.Material) {
	vnh.Set("material", v)
}
func (vnh *VertexNormalsHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: vnh.Get("matrix")}
}
func (vnh *VertexNormalsHelper) SetMatrix(v *math.Matrix4) {
	vnh.Set("matrix", v)
}
func (vnh *VertexNormalsHelper) MatrixAutoUpdate() bool {
	return vnh.Get("matrixAutoUpdate").Bool()
}
func (vnh *VertexNormalsHelper) SetMatrixAutoUpdate(v bool) {
	vnh.Set("matrixAutoUpdate", v)
}
func (vnh *VertexNormalsHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: vnh.Get("matrixWorld")}
}
func (vnh *VertexNormalsHelper) SetMatrixWorld(v *math.Matrix4) {
	vnh.Set("matrixWorld", v)
}
func (vnh *VertexNormalsHelper) MatrixWorldNeedsUpdate() bool {
	return vnh.Get("matrixWorldNeedsUpdate").Bool()
}
func (vnh *VertexNormalsHelper) SetMatrixWorldNeedsUpdate(v bool) {
	vnh.Set("matrixWorldNeedsUpdate", v)
}
func (vnh *VertexNormalsHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: vnh.Get("modelViewMatrix")}
}
func (vnh *VertexNormalsHelper) SetModelViewMatrix(v *math.Matrix4) {
	vnh.Set("modelViewMatrix", v)
}
func (vnh *VertexNormalsHelper) Name() string {
	return vnh.Get("name").String()
}
func (vnh *VertexNormalsHelper) SetName(v string) {
	vnh.Set("name", v)
}
func (vnh *VertexNormalsHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: vnh.Get("normalMatrix")}
}
func (vnh *VertexNormalsHelper) SetNormalMatrix(v *math.Matrix3) {
	vnh.Set("normalMatrix", v)
}
func (vnh *VertexNormalsHelper) Object() *core.Object3D {
	return &core.Object3D{Value: vnh.Get("object")}
}
func (vnh *VertexNormalsHelper) SetObject(v *core.Object3D) {
	vnh.Set("object", v)
}
func (vnh *VertexNormalsHelper) OnAfterRender() js.Value {
	return vnh.Get("onAfterRender")
}
func (vnh *VertexNormalsHelper) SetOnAfterRender(v js.Value) {
	vnh.Set("onAfterRender", v)
}
func (vnh *VertexNormalsHelper) OnBeforeRender() js.Value {
	return vnh.Get("onBeforeRender")
}
func (vnh *VertexNormalsHelper) SetOnBeforeRender(v js.Value) {
	vnh.Set("onBeforeRender", v)
}
func (vnh *VertexNormalsHelper) Parent() core.Object3D {
	return core.Object3D(vnh.Get("parent"))
}
func (vnh *VertexNormalsHelper) SetParent(v core.Object3D) {
	vnh.Set("parent", v)
}
func (vnh *VertexNormalsHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: vnh.Get("position")}
}
func (vnh *VertexNormalsHelper) SetPosition(v *math.Vector3) {
	vnh.Set("position", v)
}
func (vnh *VertexNormalsHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: vnh.Get("quaternion")}
}
func (vnh *VertexNormalsHelper) SetQuaternion(v *math.Quaternion) {
	vnh.Set("quaternion", v)
}
func (vnh *VertexNormalsHelper) ReceiveShadow() bool {
	return vnh.Get("receiveShadow").Bool()
}
func (vnh *VertexNormalsHelper) SetReceiveShadow(v bool) {
	vnh.Set("receiveShadow", v)
}
func (vnh *VertexNormalsHelper) RenderOrder() int {
	return vnh.Get("renderOrder").Int()
}
func (vnh *VertexNormalsHelper) SetRenderOrder(v int) {
	vnh.Set("renderOrder", v)
}
func (vnh *VertexNormalsHelper) Rotation() *math.Euler {
	return &math.Euler{Value: vnh.Get("rotation")}
}
func (vnh *VertexNormalsHelper) SetRotation(v *math.Euler) {
	vnh.Set("rotation", v)
}
func (vnh *VertexNormalsHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: vnh.Get("scale")}
}
func (vnh *VertexNormalsHelper) SetScale(v *math.Vector3) {
	vnh.Set("scale", v)
}
func (vnh *VertexNormalsHelper) Size() int {
	return vnh.Get("size").Int()
}
func (vnh *VertexNormalsHelper) SetSize(v int) {
	vnh.Set("size", v)
}
func (vnh *VertexNormalsHelper) Type() string {
	return vnh.Get("type").String()
}
func (vnh *VertexNormalsHelper) SetType(v string) {
	vnh.Set("type", v)
}
func (vnh *VertexNormalsHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: vnh.Get("up")}
}
func (vnh *VertexNormalsHelper) SetUp(v *math.Vector3) {
	vnh.Set("up", v)
}
func (vnh *VertexNormalsHelper) UserData() js.Value {
	return vnh.Get("userData")
}
func (vnh *VertexNormalsHelper) SetUserData(v js.Value) {
	vnh.Set("userData", v)
}
func (vnh *VertexNormalsHelper) Uuid() string {
	return vnh.Get("uuid").String()
}
func (vnh *VertexNormalsHelper) SetUuid(v string) {
	vnh.Set("uuid", v)
}
func (vnh *VertexNormalsHelper) Visible() bool {
	return vnh.Get("visible").Bool()
}
func (vnh *VertexNormalsHelper) SetVisible(v bool) {
	vnh.Set("visible", v)
}
func (vnh *VertexNormalsHelper) DefaultMatrixAutoUpdate() bool {
	return vnh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (vnh *VertexNormalsHelper) SetDefaultMatrixAutoUpdate(v bool) {
	vnh.Set("DefaultMatrixAutoUpdate", v)
}
func (vnh *VertexNormalsHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: vnh.Get("DefaultUp")}
}
func (vnh *VertexNormalsHelper) SetDefaultUp(v *math.Vector3) {
	vnh.Set("DefaultUp", v)
}
func (vnh *VertexNormalsHelper) Add(object []core.Object3D) this {
	return this(vnh.Call("add", object))
}
func (vnh *VertexNormalsHelper) AddEventListener(typ string, listener js.Value) {
	vnh.Call("addEventListener", typ, listener)
}
func (vnh *VertexNormalsHelper) ApplyMatrix(matrix *math.Matrix4) {
	vnh.Call("applyMatrix", matrix)
}
func (vnh *VertexNormalsHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(vnh.Call("applyQuaternion", quaternion))
}
func (vnh *VertexNormalsHelper) Clone(recursive bool) this {
	return this(vnh.Call("clone", recursive))
}
func (vnh *VertexNormalsHelper) ComputeLineDistances() this {
	return this(vnh.Call("computeLineDistances"))
}
func (vnh *VertexNormalsHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(vnh.Call("copy", source, recursive))
}
func (vnh *VertexNormalsHelper) DispatchEvent(event js.Value) {
	vnh.Call("dispatchEvent", event)
}
func (vnh *VertexNormalsHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(vnh.Call("getObjectById", id))
}
func (vnh *VertexNormalsHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(vnh.Call("getObjectByName", name))
}
func (vnh *VertexNormalsHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(vnh.Call("getObjectByProperty", name, value))
}
func (vnh *VertexNormalsHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: vnh.Call("getWorldDirection", target)}
}
func (vnh *VertexNormalsHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: vnh.Call("getWorldPosition", target)}
}
func (vnh *VertexNormalsHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: vnh.Call("getWorldQuaternion", target)}
}
func (vnh *VertexNormalsHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: vnh.Call("getWorldScale", target)}
}
func (vnh *VertexNormalsHelper) HasEventListener(typ string, listener js.Value) bool {
	return vnh.Call("hasEventListener", typ, listener).Bool()
}
func (vnh *VertexNormalsHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: vnh.Call("localToWorld", vector)}
}
func (vnh *VertexNormalsHelper) LookAt(vector math.Vector3, y int, z int) {
	vnh.Call("lookAt", vector, y, z)
}
func (vnh *VertexNormalsHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	vnh.Call("raycast", raycaster, intersects)
}
func (vnh *VertexNormalsHelper) Remove(object []core.Object3D) this {
	return this(vnh.Call("remove", object))
}
func (vnh *VertexNormalsHelper) RemoveEventListener(typ string, listener js.Value) {
	vnh.Call("removeEventListener", typ, listener)
}
func (vnh *VertexNormalsHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(vnh.Call("rotateOnAxis", axis, angle))
}
func (vnh *VertexNormalsHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(vnh.Call("rotateOnWorldAxis", axis, angle))
}
func (vnh *VertexNormalsHelper) RotateX(angle int) this {
	return this(vnh.Call("rotateX", angle))
}
func (vnh *VertexNormalsHelper) RotateY(angle int) this {
	return this(vnh.Call("rotateY", angle))
}
func (vnh *VertexNormalsHelper) RotateZ(angle int) this {
	return this(vnh.Call("rotateZ", angle))
}
func (vnh *VertexNormalsHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	vnh.Call("setRotationFromAxisAngle", axis, angle)
}
func (vnh *VertexNormalsHelper) SetRotationFromEuler(euler *math.Euler) {
	vnh.Call("setRotationFromEuler", euler)
}
func (vnh *VertexNormalsHelper) SetRotationFromMatrix(m *math.Matrix4) {
	vnh.Call("setRotationFromMatrix", m)
}
func (vnh *VertexNormalsHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	vnh.Call("setRotationFromQuaternion", q)
}
func (vnh *VertexNormalsHelper) ToJSON(meta js.Value) js.Value {
	return vnh.Call("toJSON", meta)
}
func (vnh *VertexNormalsHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(vnh.Call("translateOnAxis", axis, distance))
}
func (vnh *VertexNormalsHelper) TranslateX(distance int) this {
	return this(vnh.Call("translateX", distance))
}
func (vnh *VertexNormalsHelper) TranslateY(distance int) this {
	return this(vnh.Call("translateY", distance))
}
func (vnh *VertexNormalsHelper) TranslateZ(distance int) this {
	return this(vnh.Call("translateZ", distance))
}
func (vnh *VertexNormalsHelper) Traverse(callback js.Value) {
	vnh.Call("traverse", callback)
}
func (vnh *VertexNormalsHelper) TraverseAncestors(callback js.Value) {
	vnh.Call("traverseAncestors", callback)
}
func (vnh *VertexNormalsHelper) TraverseVisible(callback js.Value) {
	vnh.Call("traverseVisible", callback)
}
func (vnh *VertexNormalsHelper) Update(object *core.Object3D) {
	vnh.Call("update", object)
}
func (vnh *VertexNormalsHelper) UpdateMatrix() {
	vnh.Call("updateMatrix")
}
func (vnh *VertexNormalsHelper) UpdateMatrixWorld(force bool) {
	vnh.Call("updateMatrixWorld", force)
}
func (vnh *VertexNormalsHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	vnh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (vnh *VertexNormalsHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: vnh.Call("worldToLocal", vector)}
}

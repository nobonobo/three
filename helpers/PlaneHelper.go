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

type PlaneHelper struct {
	js.Value
}

func NewPlaneHelper(plane *math.Plane, size int, hex int) *PlaneHelper {
	return &PlaneHelper{Value: get("PlaneHelper").New(plane, size, hex)}
}
func (ph *PlaneHelper) CastShadow() bool {
	return ph.Get("castShadow").Bool()
}
func (ph *PlaneHelper) SetCastShadow(v bool) {
	ph.Set("castShadow", v)
}
func (ph *PlaneHelper) Children() []core.Object3D {
	return []core.Object3D(ph.Get("children"))
}
func (ph *PlaneHelper) SetChildren(v []core.Object3D) {
	ph.Set("children", v)
}
func (ph *PlaneHelper) FrustumCulled() bool {
	return ph.Get("frustumCulled").Bool()
}
func (ph *PlaneHelper) SetFrustumCulled(v bool) {
	ph.Set("frustumCulled", v)
}
func (ph *PlaneHelper) Geometry() core.Geometry {
	return core.Geometry(ph.Get("geometry"))
}
func (ph *PlaneHelper) SetGeometry(v core.Geometry) {
	ph.Set("geometry", v)
}
func (ph *PlaneHelper) Id() int {
	return ph.Get("id").Int()
}
func (ph *PlaneHelper) SetId(v int) {
	ph.Set("id", v)
}
func (ph *PlaneHelper) IsLine() true {
	return true(ph.Get("isLine"))
}
func (ph *PlaneHelper) SetIsLine(v true) {
	ph.Set("isLine", v)
}
func (ph *PlaneHelper) IsLineSegments() true {
	return true(ph.Get("isLineSegments"))
}
func (ph *PlaneHelper) SetIsLineSegments(v true) {
	ph.Set("isLineSegments", v)
}
func (ph *PlaneHelper) IsObject3D() true {
	return true(ph.Get("isObject3D"))
}
func (ph *PlaneHelper) SetIsObject3D(v true) {
	ph.Set("isObject3D", v)
}
func (ph *PlaneHelper) Layers() *core.Layers {
	return &core.Layers{Value: ph.Get("layers")}
}
func (ph *PlaneHelper) SetLayers(v *core.Layers) {
	ph.Set("layers", v)
}
func (ph *PlaneHelper) Material() materials.Material {
	return materials.Material(ph.Get("material"))
}
func (ph *PlaneHelper) SetMaterial(v materials.Material) {
	ph.Set("material", v)
}
func (ph *PlaneHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ph.Get("matrix")}
}
func (ph *PlaneHelper) SetMatrix(v *math.Matrix4) {
	ph.Set("matrix", v)
}
func (ph *PlaneHelper) MatrixAutoUpdate() bool {
	return ph.Get("matrixAutoUpdate").Bool()
}
func (ph *PlaneHelper) SetMatrixAutoUpdate(v bool) {
	ph.Set("matrixAutoUpdate", v)
}
func (ph *PlaneHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ph.Get("matrixWorld")}
}
func (ph *PlaneHelper) SetMatrixWorld(v *math.Matrix4) {
	ph.Set("matrixWorld", v)
}
func (ph *PlaneHelper) MatrixWorldNeedsUpdate() bool {
	return ph.Get("matrixWorldNeedsUpdate").Bool()
}
func (ph *PlaneHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ph.Set("matrixWorldNeedsUpdate", v)
}
func (ph *PlaneHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ph.Get("modelViewMatrix")}
}
func (ph *PlaneHelper) SetModelViewMatrix(v *math.Matrix4) {
	ph.Set("modelViewMatrix", v)
}
func (ph *PlaneHelper) Name() string {
	return ph.Get("name").String()
}
func (ph *PlaneHelper) SetName(v string) {
	ph.Set("name", v)
}
func (ph *PlaneHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ph.Get("normalMatrix")}
}
func (ph *PlaneHelper) SetNormalMatrix(v *math.Matrix3) {
	ph.Set("normalMatrix", v)
}
func (ph *PlaneHelper) OnAfterRender() js.Value {
	return ph.Get("onAfterRender")
}
func (ph *PlaneHelper) SetOnAfterRender(v js.Value) {
	ph.Set("onAfterRender", v)
}
func (ph *PlaneHelper) OnBeforeRender() js.Value {
	return ph.Get("onBeforeRender")
}
func (ph *PlaneHelper) SetOnBeforeRender(v js.Value) {
	ph.Set("onBeforeRender", v)
}
func (ph *PlaneHelper) Parent() core.Object3D {
	return core.Object3D(ph.Get("parent"))
}
func (ph *PlaneHelper) SetParent(v core.Object3D) {
	ph.Set("parent", v)
}
func (ph *PlaneHelper) Plane() *math.Plane {
	return &math.Plane{Value: ph.Get("plane")}
}
func (ph *PlaneHelper) SetPlane(v *math.Plane) {
	ph.Set("plane", v)
}
func (ph *PlaneHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: ph.Get("position")}
}
func (ph *PlaneHelper) SetPosition(v *math.Vector3) {
	ph.Set("position", v)
}
func (ph *PlaneHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ph.Get("quaternion")}
}
func (ph *PlaneHelper) SetQuaternion(v *math.Quaternion) {
	ph.Set("quaternion", v)
}
func (ph *PlaneHelper) ReceiveShadow() bool {
	return ph.Get("receiveShadow").Bool()
}
func (ph *PlaneHelper) SetReceiveShadow(v bool) {
	ph.Set("receiveShadow", v)
}
func (ph *PlaneHelper) RenderOrder() int {
	return ph.Get("renderOrder").Int()
}
func (ph *PlaneHelper) SetRenderOrder(v int) {
	ph.Set("renderOrder", v)
}
func (ph *PlaneHelper) Rotation() *math.Euler {
	return &math.Euler{Value: ph.Get("rotation")}
}
func (ph *PlaneHelper) SetRotation(v *math.Euler) {
	ph.Set("rotation", v)
}
func (ph *PlaneHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: ph.Get("scale")}
}
func (ph *PlaneHelper) SetScale(v *math.Vector3) {
	ph.Set("scale", v)
}
func (ph *PlaneHelper) Size() int {
	return ph.Get("size").Int()
}
func (ph *PlaneHelper) SetSize(v int) {
	ph.Set("size", v)
}
func (ph *PlaneHelper) Type() string {
	return ph.Get("type").String()
}
func (ph *PlaneHelper) SetType(v string) {
	ph.Set("type", v)
}
func (ph *PlaneHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: ph.Get("up")}
}
func (ph *PlaneHelper) SetUp(v *math.Vector3) {
	ph.Set("up", v)
}
func (ph *PlaneHelper) UserData() js.Value {
	return ph.Get("userData")
}
func (ph *PlaneHelper) SetUserData(v js.Value) {
	ph.Set("userData", v)
}
func (ph *PlaneHelper) Uuid() string {
	return ph.Get("uuid").String()
}
func (ph *PlaneHelper) SetUuid(v string) {
	ph.Set("uuid", v)
}
func (ph *PlaneHelper) Visible() bool {
	return ph.Get("visible").Bool()
}
func (ph *PlaneHelper) SetVisible(v bool) {
	ph.Set("visible", v)
}
func (ph *PlaneHelper) DefaultMatrixAutoUpdate() bool {
	return ph.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ph *PlaneHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ph.Set("DefaultMatrixAutoUpdate", v)
}
func (ph *PlaneHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ph.Get("DefaultUp")}
}
func (ph *PlaneHelper) SetDefaultUp(v *math.Vector3) {
	ph.Set("DefaultUp", v)
}
func (ph *PlaneHelper) Add(object []core.Object3D) this {
	return this(ph.Call("add", object))
}
func (ph *PlaneHelper) AddEventListener(typ string, listener js.Value) {
	ph.Call("addEventListener", typ, listener)
}
func (ph *PlaneHelper) ApplyMatrix(matrix *math.Matrix4) {
	ph.Call("applyMatrix", matrix)
}
func (ph *PlaneHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ph.Call("applyQuaternion", quaternion))
}
func (ph *PlaneHelper) Clone(recursive bool) this {
	return this(ph.Call("clone", recursive))
}
func (ph *PlaneHelper) ComputeLineDistances() this {
	return this(ph.Call("computeLineDistances"))
}
func (ph *PlaneHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(ph.Call("copy", source, recursive))
}
func (ph *PlaneHelper) DispatchEvent(event js.Value) {
	ph.Call("dispatchEvent", event)
}
func (ph *PlaneHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(ph.Call("getObjectById", id))
}
func (ph *PlaneHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ph.Call("getObjectByName", name))
}
func (ph *PlaneHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ph.Call("getObjectByProperty", name, value))
}
func (ph *PlaneHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ph.Call("getWorldDirection", target)}
}
func (ph *PlaneHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ph.Call("getWorldPosition", target)}
}
func (ph *PlaneHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ph.Call("getWorldQuaternion", target)}
}
func (ph *PlaneHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ph.Call("getWorldScale", target)}
}
func (ph *PlaneHelper) HasEventListener(typ string, listener js.Value) bool {
	return ph.Call("hasEventListener", typ, listener).Bool()
}
func (ph *PlaneHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ph.Call("localToWorld", vector)}
}
func (ph *PlaneHelper) LookAt(vector math.Vector3, y int, z int) {
	ph.Call("lookAt", vector, y, z)
}
func (ph *PlaneHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ph.Call("raycast", raycaster, intersects)
}
func (ph *PlaneHelper) Remove(object []core.Object3D) this {
	return this(ph.Call("remove", object))
}
func (ph *PlaneHelper) RemoveEventListener(typ string, listener js.Value) {
	ph.Call("removeEventListener", typ, listener)
}
func (ph *PlaneHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ph.Call("rotateOnAxis", axis, angle))
}
func (ph *PlaneHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ph.Call("rotateOnWorldAxis", axis, angle))
}
func (ph *PlaneHelper) RotateX(angle int) this {
	return this(ph.Call("rotateX", angle))
}
func (ph *PlaneHelper) RotateY(angle int) this {
	return this(ph.Call("rotateY", angle))
}
func (ph *PlaneHelper) RotateZ(angle int) this {
	return this(ph.Call("rotateZ", angle))
}
func (ph *PlaneHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ph.Call("setRotationFromAxisAngle", axis, angle)
}
func (ph *PlaneHelper) SetRotationFromEuler(euler *math.Euler) {
	ph.Call("setRotationFromEuler", euler)
}
func (ph *PlaneHelper) SetRotationFromMatrix(m *math.Matrix4) {
	ph.Call("setRotationFromMatrix", m)
}
func (ph *PlaneHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	ph.Call("setRotationFromQuaternion", q)
}
func (ph *PlaneHelper) ToJSON(meta js.Value) js.Value {
	return ph.Call("toJSON", meta)
}
func (ph *PlaneHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ph.Call("translateOnAxis", axis, distance))
}
func (ph *PlaneHelper) TranslateX(distance int) this {
	return this(ph.Call("translateX", distance))
}
func (ph *PlaneHelper) TranslateY(distance int) this {
	return this(ph.Call("translateY", distance))
}
func (ph *PlaneHelper) TranslateZ(distance int) this {
	return this(ph.Call("translateZ", distance))
}
func (ph *PlaneHelper) Traverse(callback js.Value) {
	ph.Call("traverse", callback)
}
func (ph *PlaneHelper) TraverseAncestors(callback js.Value) {
	ph.Call("traverseAncestors", callback)
}
func (ph *PlaneHelper) TraverseVisible(callback js.Value) {
	ph.Call("traverseVisible", callback)
}
func (ph *PlaneHelper) UpdateMatrix() {
	ph.Call("updateMatrix")
}
func (ph *PlaneHelper) UpdateMatrixWorld(force bool) {
	ph.Call("updateMatrixWorld", force)
}
func (ph *PlaneHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ph.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ph *PlaneHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ph.Call("worldToLocal", vector)}
}

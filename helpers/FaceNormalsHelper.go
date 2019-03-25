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

type FaceNormalsHelper struct {
	js.Value
}

func NewFaceNormalsHelper(object *core.Object3D, size int, hex int, linewidth int) *FaceNormalsHelper {
	return &FaceNormalsHelper{Value: get("FaceNormalsHelper").New(object, size, hex, linewidth)}
}
func (fnh *FaceNormalsHelper) CastShadow() bool {
	return fnh.Get("castShadow").Bool()
}
func (fnh *FaceNormalsHelper) SetCastShadow(v bool) {
	fnh.Set("castShadow", v)
}
func (fnh *FaceNormalsHelper) Children() []core.Object3D {
	return []core.Object3D(fnh.Get("children"))
}
func (fnh *FaceNormalsHelper) SetChildren(v []core.Object3D) {
	fnh.Set("children", v)
}
func (fnh *FaceNormalsHelper) FrustumCulled() bool {
	return fnh.Get("frustumCulled").Bool()
}
func (fnh *FaceNormalsHelper) SetFrustumCulled(v bool) {
	fnh.Set("frustumCulled", v)
}
func (fnh *FaceNormalsHelper) Geometry() core.Geometry {
	return core.Geometry(fnh.Get("geometry"))
}
func (fnh *FaceNormalsHelper) SetGeometry(v core.Geometry) {
	fnh.Set("geometry", v)
}
func (fnh *FaceNormalsHelper) Id() int {
	return fnh.Get("id").Int()
}
func (fnh *FaceNormalsHelper) SetId(v int) {
	fnh.Set("id", v)
}
func (fnh *FaceNormalsHelper) IsLine() true {
	return true(fnh.Get("isLine"))
}
func (fnh *FaceNormalsHelper) SetIsLine(v true) {
	fnh.Set("isLine", v)
}
func (fnh *FaceNormalsHelper) IsLineSegments() true {
	return true(fnh.Get("isLineSegments"))
}
func (fnh *FaceNormalsHelper) SetIsLineSegments(v true) {
	fnh.Set("isLineSegments", v)
}
func (fnh *FaceNormalsHelper) IsObject3D() true {
	return true(fnh.Get("isObject3D"))
}
func (fnh *FaceNormalsHelper) SetIsObject3D(v true) {
	fnh.Set("isObject3D", v)
}
func (fnh *FaceNormalsHelper) Layers() *core.Layers {
	return &core.Layers{Value: fnh.Get("layers")}
}
func (fnh *FaceNormalsHelper) SetLayers(v *core.Layers) {
	fnh.Set("layers", v)
}
func (fnh *FaceNormalsHelper) Material() materials.Material {
	return materials.Material(fnh.Get("material"))
}
func (fnh *FaceNormalsHelper) SetMaterial(v materials.Material) {
	fnh.Set("material", v)
}
func (fnh *FaceNormalsHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: fnh.Get("matrix")}
}
func (fnh *FaceNormalsHelper) SetMatrix(v *math.Matrix4) {
	fnh.Set("matrix", v)
}
func (fnh *FaceNormalsHelper) MatrixAutoUpdate() bool {
	return fnh.Get("matrixAutoUpdate").Bool()
}
func (fnh *FaceNormalsHelper) SetMatrixAutoUpdate(v bool) {
	fnh.Set("matrixAutoUpdate", v)
}
func (fnh *FaceNormalsHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: fnh.Get("matrixWorld")}
}
func (fnh *FaceNormalsHelper) SetMatrixWorld(v *math.Matrix4) {
	fnh.Set("matrixWorld", v)
}
func (fnh *FaceNormalsHelper) MatrixWorldNeedsUpdate() bool {
	return fnh.Get("matrixWorldNeedsUpdate").Bool()
}
func (fnh *FaceNormalsHelper) SetMatrixWorldNeedsUpdate(v bool) {
	fnh.Set("matrixWorldNeedsUpdate", v)
}
func (fnh *FaceNormalsHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: fnh.Get("modelViewMatrix")}
}
func (fnh *FaceNormalsHelper) SetModelViewMatrix(v *math.Matrix4) {
	fnh.Set("modelViewMatrix", v)
}
func (fnh *FaceNormalsHelper) Name() string {
	return fnh.Get("name").String()
}
func (fnh *FaceNormalsHelper) SetName(v string) {
	fnh.Set("name", v)
}
func (fnh *FaceNormalsHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: fnh.Get("normalMatrix")}
}
func (fnh *FaceNormalsHelper) SetNormalMatrix(v *math.Matrix3) {
	fnh.Set("normalMatrix", v)
}
func (fnh *FaceNormalsHelper) Object() *core.Object3D {
	return &core.Object3D{Value: fnh.Get("object")}
}
func (fnh *FaceNormalsHelper) SetObject(v *core.Object3D) {
	fnh.Set("object", v)
}
func (fnh *FaceNormalsHelper) OnAfterRender() js.Value {
	return fnh.Get("onAfterRender")
}
func (fnh *FaceNormalsHelper) SetOnAfterRender(v js.Value) {
	fnh.Set("onAfterRender", v)
}
func (fnh *FaceNormalsHelper) OnBeforeRender() js.Value {
	return fnh.Get("onBeforeRender")
}
func (fnh *FaceNormalsHelper) SetOnBeforeRender(v js.Value) {
	fnh.Set("onBeforeRender", v)
}
func (fnh *FaceNormalsHelper) Parent() core.Object3D {
	return core.Object3D(fnh.Get("parent"))
}
func (fnh *FaceNormalsHelper) SetParent(v core.Object3D) {
	fnh.Set("parent", v)
}
func (fnh *FaceNormalsHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: fnh.Get("position")}
}
func (fnh *FaceNormalsHelper) SetPosition(v *math.Vector3) {
	fnh.Set("position", v)
}
func (fnh *FaceNormalsHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: fnh.Get("quaternion")}
}
func (fnh *FaceNormalsHelper) SetQuaternion(v *math.Quaternion) {
	fnh.Set("quaternion", v)
}
func (fnh *FaceNormalsHelper) ReceiveShadow() bool {
	return fnh.Get("receiveShadow").Bool()
}
func (fnh *FaceNormalsHelper) SetReceiveShadow(v bool) {
	fnh.Set("receiveShadow", v)
}
func (fnh *FaceNormalsHelper) RenderOrder() int {
	return fnh.Get("renderOrder").Int()
}
func (fnh *FaceNormalsHelper) SetRenderOrder(v int) {
	fnh.Set("renderOrder", v)
}
func (fnh *FaceNormalsHelper) Rotation() *math.Euler {
	return &math.Euler{Value: fnh.Get("rotation")}
}
func (fnh *FaceNormalsHelper) SetRotation(v *math.Euler) {
	fnh.Set("rotation", v)
}
func (fnh *FaceNormalsHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: fnh.Get("scale")}
}
func (fnh *FaceNormalsHelper) SetScale(v *math.Vector3) {
	fnh.Set("scale", v)
}
func (fnh *FaceNormalsHelper) Size() int {
	return fnh.Get("size").Int()
}
func (fnh *FaceNormalsHelper) SetSize(v int) {
	fnh.Set("size", v)
}
func (fnh *FaceNormalsHelper) Type() string {
	return fnh.Get("type").String()
}
func (fnh *FaceNormalsHelper) SetType(v string) {
	fnh.Set("type", v)
}
func (fnh *FaceNormalsHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: fnh.Get("up")}
}
func (fnh *FaceNormalsHelper) SetUp(v *math.Vector3) {
	fnh.Set("up", v)
}
func (fnh *FaceNormalsHelper) UserData() js.Value {
	return fnh.Get("userData")
}
func (fnh *FaceNormalsHelper) SetUserData(v js.Value) {
	fnh.Set("userData", v)
}
func (fnh *FaceNormalsHelper) Uuid() string {
	return fnh.Get("uuid").String()
}
func (fnh *FaceNormalsHelper) SetUuid(v string) {
	fnh.Set("uuid", v)
}
func (fnh *FaceNormalsHelper) Visible() bool {
	return fnh.Get("visible").Bool()
}
func (fnh *FaceNormalsHelper) SetVisible(v bool) {
	fnh.Set("visible", v)
}
func (fnh *FaceNormalsHelper) DefaultMatrixAutoUpdate() bool {
	return fnh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (fnh *FaceNormalsHelper) SetDefaultMatrixAutoUpdate(v bool) {
	fnh.Set("DefaultMatrixAutoUpdate", v)
}
func (fnh *FaceNormalsHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: fnh.Get("DefaultUp")}
}
func (fnh *FaceNormalsHelper) SetDefaultUp(v *math.Vector3) {
	fnh.Set("DefaultUp", v)
}
func (fnh *FaceNormalsHelper) Add(object []core.Object3D) this {
	return this(fnh.Call("add", object))
}
func (fnh *FaceNormalsHelper) AddEventListener(typ string, listener js.Value) {
	fnh.Call("addEventListener", typ, listener)
}
func (fnh *FaceNormalsHelper) ApplyMatrix(matrix *math.Matrix4) {
	fnh.Call("applyMatrix", matrix)
}
func (fnh *FaceNormalsHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(fnh.Call("applyQuaternion", quaternion))
}
func (fnh *FaceNormalsHelper) Clone(recursive bool) this {
	return this(fnh.Call("clone", recursive))
}
func (fnh *FaceNormalsHelper) ComputeLineDistances() this {
	return this(fnh.Call("computeLineDistances"))
}
func (fnh *FaceNormalsHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(fnh.Call("copy", source, recursive))
}
func (fnh *FaceNormalsHelper) DispatchEvent(event js.Value) {
	fnh.Call("dispatchEvent", event)
}
func (fnh *FaceNormalsHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(fnh.Call("getObjectById", id))
}
func (fnh *FaceNormalsHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(fnh.Call("getObjectByName", name))
}
func (fnh *FaceNormalsHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(fnh.Call("getObjectByProperty", name, value))
}
func (fnh *FaceNormalsHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: fnh.Call("getWorldDirection", target)}
}
func (fnh *FaceNormalsHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: fnh.Call("getWorldPosition", target)}
}
func (fnh *FaceNormalsHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: fnh.Call("getWorldQuaternion", target)}
}
func (fnh *FaceNormalsHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: fnh.Call("getWorldScale", target)}
}
func (fnh *FaceNormalsHelper) HasEventListener(typ string, listener js.Value) bool {
	return fnh.Call("hasEventListener", typ, listener).Bool()
}
func (fnh *FaceNormalsHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: fnh.Call("localToWorld", vector)}
}
func (fnh *FaceNormalsHelper) LookAt(vector math.Vector3, y int, z int) {
	fnh.Call("lookAt", vector, y, z)
}
func (fnh *FaceNormalsHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	fnh.Call("raycast", raycaster, intersects)
}
func (fnh *FaceNormalsHelper) Remove(object []core.Object3D) this {
	return this(fnh.Call("remove", object))
}
func (fnh *FaceNormalsHelper) RemoveEventListener(typ string, listener js.Value) {
	fnh.Call("removeEventListener", typ, listener)
}
func (fnh *FaceNormalsHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(fnh.Call("rotateOnAxis", axis, angle))
}
func (fnh *FaceNormalsHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(fnh.Call("rotateOnWorldAxis", axis, angle))
}
func (fnh *FaceNormalsHelper) RotateX(angle int) this {
	return this(fnh.Call("rotateX", angle))
}
func (fnh *FaceNormalsHelper) RotateY(angle int) this {
	return this(fnh.Call("rotateY", angle))
}
func (fnh *FaceNormalsHelper) RotateZ(angle int) this {
	return this(fnh.Call("rotateZ", angle))
}
func (fnh *FaceNormalsHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	fnh.Call("setRotationFromAxisAngle", axis, angle)
}
func (fnh *FaceNormalsHelper) SetRotationFromEuler(euler *math.Euler) {
	fnh.Call("setRotationFromEuler", euler)
}
func (fnh *FaceNormalsHelper) SetRotationFromMatrix(m *math.Matrix4) {
	fnh.Call("setRotationFromMatrix", m)
}
func (fnh *FaceNormalsHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	fnh.Call("setRotationFromQuaternion", q)
}
func (fnh *FaceNormalsHelper) ToJSON(meta js.Value) js.Value {
	return fnh.Call("toJSON", meta)
}
func (fnh *FaceNormalsHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(fnh.Call("translateOnAxis", axis, distance))
}
func (fnh *FaceNormalsHelper) TranslateX(distance int) this {
	return this(fnh.Call("translateX", distance))
}
func (fnh *FaceNormalsHelper) TranslateY(distance int) this {
	return this(fnh.Call("translateY", distance))
}
func (fnh *FaceNormalsHelper) TranslateZ(distance int) this {
	return this(fnh.Call("translateZ", distance))
}
func (fnh *FaceNormalsHelper) Traverse(callback js.Value) {
	fnh.Call("traverse", callback)
}
func (fnh *FaceNormalsHelper) TraverseAncestors(callback js.Value) {
	fnh.Call("traverseAncestors", callback)
}
func (fnh *FaceNormalsHelper) TraverseVisible(callback js.Value) {
	fnh.Call("traverseVisible", callback)
}
func (fnh *FaceNormalsHelper) Update(object *core.Object3D) {
	fnh.Call("update", object)
}
func (fnh *FaceNormalsHelper) UpdateMatrix() {
	fnh.Call("updateMatrix")
}
func (fnh *FaceNormalsHelper) UpdateMatrixWorld(force bool) {
	fnh.Call("updateMatrixWorld", force)
}
func (fnh *FaceNormalsHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	fnh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (fnh *FaceNormalsHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: fnh.Call("worldToLocal", vector)}
}

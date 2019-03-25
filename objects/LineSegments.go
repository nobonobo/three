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

func LinePieces() int {
	return get("LinePieces").Int()
}
func SetLinePieces(v int) {
	set("LinePieces", v)
}
func LineStrip() int {
	return get("LineStrip").Int()
}
func SetLineStrip(v int) {
	set("LineStrip", v)
}

type LineSegments struct {
	js.Value
}

func NewLineSegments(geometry core.Geometry, material materials.Material, mode int) *LineSegments {
	return &LineSegments{Value: get("LineSegments").New(geometry, material, mode)}
}
func (ls *LineSegments) CastShadow() bool {
	return ls.Get("castShadow").Bool()
}
func (ls *LineSegments) SetCastShadow(v bool) {
	ls.Set("castShadow", v)
}
func (ls *LineSegments) Children() []core.Object3D {
	return []core.Object3D(ls.Get("children"))
}
func (ls *LineSegments) SetChildren(v []core.Object3D) {
	ls.Set("children", v)
}
func (ls *LineSegments) FrustumCulled() bool {
	return ls.Get("frustumCulled").Bool()
}
func (ls *LineSegments) SetFrustumCulled(v bool) {
	ls.Set("frustumCulled", v)
}
func (ls *LineSegments) Geometry() core.Geometry {
	return core.Geometry(ls.Get("geometry"))
}
func (ls *LineSegments) SetGeometry(v core.Geometry) {
	ls.Set("geometry", v)
}
func (ls *LineSegments) Id() int {
	return ls.Get("id").Int()
}
func (ls *LineSegments) SetId(v int) {
	ls.Set("id", v)
}
func (ls *LineSegments) IsLine() true {
	return true(ls.Get("isLine"))
}
func (ls *LineSegments) SetIsLine(v true) {
	ls.Set("isLine", v)
}
func (ls *LineSegments) IsLineSegments() true {
	return true(ls.Get("isLineSegments"))
}
func (ls *LineSegments) SetIsLineSegments(v true) {
	ls.Set("isLineSegments", v)
}
func (ls *LineSegments) IsObject3D() true {
	return true(ls.Get("isObject3D"))
}
func (ls *LineSegments) SetIsObject3D(v true) {
	ls.Set("isObject3D", v)
}
func (ls *LineSegments) Layers() *core.Layers {
	return &core.Layers{Value: ls.Get("layers")}
}
func (ls *LineSegments) SetLayers(v *core.Layers) {
	ls.Set("layers", v)
}
func (ls *LineSegments) Material() materials.Material {
	return materials.Material(ls.Get("material"))
}
func (ls *LineSegments) SetMaterial(v materials.Material) {
	ls.Set("material", v)
}
func (ls *LineSegments) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ls.Get("matrix")}
}
func (ls *LineSegments) SetMatrix(v *math.Matrix4) {
	ls.Set("matrix", v)
}
func (ls *LineSegments) MatrixAutoUpdate() bool {
	return ls.Get("matrixAutoUpdate").Bool()
}
func (ls *LineSegments) SetMatrixAutoUpdate(v bool) {
	ls.Set("matrixAutoUpdate", v)
}
func (ls *LineSegments) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ls.Get("matrixWorld")}
}
func (ls *LineSegments) SetMatrixWorld(v *math.Matrix4) {
	ls.Set("matrixWorld", v)
}
func (ls *LineSegments) MatrixWorldNeedsUpdate() bool {
	return ls.Get("matrixWorldNeedsUpdate").Bool()
}
func (ls *LineSegments) SetMatrixWorldNeedsUpdate(v bool) {
	ls.Set("matrixWorldNeedsUpdate", v)
}
func (ls *LineSegments) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ls.Get("modelViewMatrix")}
}
func (ls *LineSegments) SetModelViewMatrix(v *math.Matrix4) {
	ls.Set("modelViewMatrix", v)
}
func (ls *LineSegments) Name() string {
	return ls.Get("name").String()
}
func (ls *LineSegments) SetName(v string) {
	ls.Set("name", v)
}
func (ls *LineSegments) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ls.Get("normalMatrix")}
}
func (ls *LineSegments) SetNormalMatrix(v *math.Matrix3) {
	ls.Set("normalMatrix", v)
}
func (ls *LineSegments) OnAfterRender() js.Value {
	return ls.Get("onAfterRender")
}
func (ls *LineSegments) SetOnAfterRender(v js.Value) {
	ls.Set("onAfterRender", v)
}
func (ls *LineSegments) OnBeforeRender() js.Value {
	return ls.Get("onBeforeRender")
}
func (ls *LineSegments) SetOnBeforeRender(v js.Value) {
	ls.Set("onBeforeRender", v)
}
func (ls *LineSegments) Parent() core.Object3D {
	return core.Object3D(ls.Get("parent"))
}
func (ls *LineSegments) SetParent(v core.Object3D) {
	ls.Set("parent", v)
}
func (ls *LineSegments) Position() *math.Vector3 {
	return &math.Vector3{Value: ls.Get("position")}
}
func (ls *LineSegments) SetPosition(v *math.Vector3) {
	ls.Set("position", v)
}
func (ls *LineSegments) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ls.Get("quaternion")}
}
func (ls *LineSegments) SetQuaternion(v *math.Quaternion) {
	ls.Set("quaternion", v)
}
func (ls *LineSegments) ReceiveShadow() bool {
	return ls.Get("receiveShadow").Bool()
}
func (ls *LineSegments) SetReceiveShadow(v bool) {
	ls.Set("receiveShadow", v)
}
func (ls *LineSegments) RenderOrder() int {
	return ls.Get("renderOrder").Int()
}
func (ls *LineSegments) SetRenderOrder(v int) {
	ls.Set("renderOrder", v)
}
func (ls *LineSegments) Rotation() *math.Euler {
	return &math.Euler{Value: ls.Get("rotation")}
}
func (ls *LineSegments) SetRotation(v *math.Euler) {
	ls.Set("rotation", v)
}
func (ls *LineSegments) Scale() *math.Vector3 {
	return &math.Vector3{Value: ls.Get("scale")}
}
func (ls *LineSegments) SetScale(v *math.Vector3) {
	ls.Set("scale", v)
}
func (ls *LineSegments) Type() string {
	return ls.Get("type").String()
}
func (ls *LineSegments) SetType(v string) {
	ls.Set("type", v)
}
func (ls *LineSegments) Up() *math.Vector3 {
	return &math.Vector3{Value: ls.Get("up")}
}
func (ls *LineSegments) SetUp(v *math.Vector3) {
	ls.Set("up", v)
}
func (ls *LineSegments) UserData() js.Value {
	return ls.Get("userData")
}
func (ls *LineSegments) SetUserData(v js.Value) {
	ls.Set("userData", v)
}
func (ls *LineSegments) Uuid() string {
	return ls.Get("uuid").String()
}
func (ls *LineSegments) SetUuid(v string) {
	ls.Set("uuid", v)
}
func (ls *LineSegments) Visible() bool {
	return ls.Get("visible").Bool()
}
func (ls *LineSegments) SetVisible(v bool) {
	ls.Set("visible", v)
}
func (ls *LineSegments) DefaultMatrixAutoUpdate() bool {
	return ls.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ls *LineSegments) SetDefaultMatrixAutoUpdate(v bool) {
	ls.Set("DefaultMatrixAutoUpdate", v)
}
func (ls *LineSegments) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ls.Get("DefaultUp")}
}
func (ls *LineSegments) SetDefaultUp(v *math.Vector3) {
	ls.Set("DefaultUp", v)
}
func (ls *LineSegments) Add(object []core.Object3D) this {
	return this(ls.Call("add", object))
}
func (ls *LineSegments) AddEventListener(typ string, listener js.Value) {
	ls.Call("addEventListener", typ, listener)
}
func (ls *LineSegments) ApplyMatrix(matrix *math.Matrix4) {
	ls.Call("applyMatrix", matrix)
}
func (ls *LineSegments) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ls.Call("applyQuaternion", quaternion))
}
func (ls *LineSegments) Clone(recursive bool) this {
	return this(ls.Call("clone", recursive))
}
func (ls *LineSegments) ComputeLineDistances() this {
	return this(ls.Call("computeLineDistances"))
}
func (ls *LineSegments) Copy(source *core.Object3D, recursive bool) this {
	return this(ls.Call("copy", source, recursive))
}
func (ls *LineSegments) DispatchEvent(event js.Value) {
	ls.Call("dispatchEvent", event)
}
func (ls *LineSegments) GetObjectById(id int) core.Object3D {
	return core.Object3D(ls.Call("getObjectById", id))
}
func (ls *LineSegments) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ls.Call("getObjectByName", name))
}
func (ls *LineSegments) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ls.Call("getObjectByProperty", name, value))
}
func (ls *LineSegments) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ls.Call("getWorldDirection", target)}
}
func (ls *LineSegments) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ls.Call("getWorldPosition", target)}
}
func (ls *LineSegments) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ls.Call("getWorldQuaternion", target)}
}
func (ls *LineSegments) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ls.Call("getWorldScale", target)}
}
func (ls *LineSegments) HasEventListener(typ string, listener js.Value) bool {
	return ls.Call("hasEventListener", typ, listener).Bool()
}
func (ls *LineSegments) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ls.Call("localToWorld", vector)}
}
func (ls *LineSegments) LookAt(vector math.Vector3, y int, z int) {
	ls.Call("lookAt", vector, y, z)
}
func (ls *LineSegments) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ls.Call("raycast", raycaster, intersects)
}
func (ls *LineSegments) Remove(object []core.Object3D) this {
	return this(ls.Call("remove", object))
}
func (ls *LineSegments) RemoveEventListener(typ string, listener js.Value) {
	ls.Call("removeEventListener", typ, listener)
}
func (ls *LineSegments) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ls.Call("rotateOnAxis", axis, angle))
}
func (ls *LineSegments) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ls.Call("rotateOnWorldAxis", axis, angle))
}
func (ls *LineSegments) RotateX(angle int) this {
	return this(ls.Call("rotateX", angle))
}
func (ls *LineSegments) RotateY(angle int) this {
	return this(ls.Call("rotateY", angle))
}
func (ls *LineSegments) RotateZ(angle int) this {
	return this(ls.Call("rotateZ", angle))
}
func (ls *LineSegments) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ls.Call("setRotationFromAxisAngle", axis, angle)
}
func (ls *LineSegments) SetRotationFromEuler(euler *math.Euler) {
	ls.Call("setRotationFromEuler", euler)
}
func (ls *LineSegments) SetRotationFromMatrix(m *math.Matrix4) {
	ls.Call("setRotationFromMatrix", m)
}
func (ls *LineSegments) SetRotationFromQuaternion(q *math.Quaternion) {
	ls.Call("setRotationFromQuaternion", q)
}
func (ls *LineSegments) ToJSON(meta js.Value) js.Value {
	return ls.Call("toJSON", meta)
}
func (ls *LineSegments) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ls.Call("translateOnAxis", axis, distance))
}
func (ls *LineSegments) TranslateX(distance int) this {
	return this(ls.Call("translateX", distance))
}
func (ls *LineSegments) TranslateY(distance int) this {
	return this(ls.Call("translateY", distance))
}
func (ls *LineSegments) TranslateZ(distance int) this {
	return this(ls.Call("translateZ", distance))
}
func (ls *LineSegments) Traverse(callback js.Value) {
	ls.Call("traverse", callback)
}
func (ls *LineSegments) TraverseAncestors(callback js.Value) {
	ls.Call("traverseAncestors", callback)
}
func (ls *LineSegments) TraverseVisible(callback js.Value) {
	ls.Call("traverseVisible", callback)
}
func (ls *LineSegments) UpdateMatrix() {
	ls.Call("updateMatrix")
}
func (ls *LineSegments) UpdateMatrixWorld(force bool) {
	ls.Call("updateMatrixWorld", force)
}
func (ls *LineSegments) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ls.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ls *LineSegments) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ls.Call("worldToLocal", vector)}
}

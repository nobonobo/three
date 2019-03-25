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

type ArrowHelper struct {
	js.Value
}

func NewArrowHelper(dir *math.Vector3, origin *math.Vector3, length int, hex int, headLength int, headWidth int) *ArrowHelper {
	return &ArrowHelper{Value: get("ArrowHelper").New(dir, origin, length, hex, headLength, headWidth)}
}
func (ah *ArrowHelper) CastShadow() bool {
	return ah.Get("castShadow").Bool()
}
func (ah *ArrowHelper) SetCastShadow(v bool) {
	ah.Set("castShadow", v)
}
func (ah *ArrowHelper) Children() []core.Object3D {
	return []core.Object3D(ah.Get("children"))
}
func (ah *ArrowHelper) SetChildren(v []core.Object3D) {
	ah.Set("children", v)
}
func (ah *ArrowHelper) Cone() *objects.Mesh {
	return &objects.Mesh{Value: ah.Get("cone")}
}
func (ah *ArrowHelper) SetCone(v *objects.Mesh) {
	ah.Set("cone", v)
}
func (ah *ArrowHelper) FrustumCulled() bool {
	return ah.Get("frustumCulled").Bool()
}
func (ah *ArrowHelper) SetFrustumCulled(v bool) {
	ah.Set("frustumCulled", v)
}
func (ah *ArrowHelper) Id() int {
	return ah.Get("id").Int()
}
func (ah *ArrowHelper) SetId(v int) {
	ah.Set("id", v)
}
func (ah *ArrowHelper) IsObject3D() true {
	return true(ah.Get("isObject3D"))
}
func (ah *ArrowHelper) SetIsObject3D(v true) {
	ah.Set("isObject3D", v)
}
func (ah *ArrowHelper) Layers() *core.Layers {
	return &core.Layers{Value: ah.Get("layers")}
}
func (ah *ArrowHelper) SetLayers(v *core.Layers) {
	ah.Set("layers", v)
}
func (ah *ArrowHelper) Line() *objects.Line {
	return &objects.Line{Value: ah.Get("line")}
}
func (ah *ArrowHelper) SetLine(v *objects.Line) {
	ah.Set("line", v)
}
func (ah *ArrowHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: ah.Get("matrix")}
}
func (ah *ArrowHelper) SetMatrix(v *math.Matrix4) {
	ah.Set("matrix", v)
}
func (ah *ArrowHelper) MatrixAutoUpdate() bool {
	return ah.Get("matrixAutoUpdate").Bool()
}
func (ah *ArrowHelper) SetMatrixAutoUpdate(v bool) {
	ah.Set("matrixAutoUpdate", v)
}
func (ah *ArrowHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: ah.Get("matrixWorld")}
}
func (ah *ArrowHelper) SetMatrixWorld(v *math.Matrix4) {
	ah.Set("matrixWorld", v)
}
func (ah *ArrowHelper) MatrixWorldNeedsUpdate() bool {
	return ah.Get("matrixWorldNeedsUpdate").Bool()
}
func (ah *ArrowHelper) SetMatrixWorldNeedsUpdate(v bool) {
	ah.Set("matrixWorldNeedsUpdate", v)
}
func (ah *ArrowHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: ah.Get("modelViewMatrix")}
}
func (ah *ArrowHelper) SetModelViewMatrix(v *math.Matrix4) {
	ah.Set("modelViewMatrix", v)
}
func (ah *ArrowHelper) Name() string {
	return ah.Get("name").String()
}
func (ah *ArrowHelper) SetName(v string) {
	ah.Set("name", v)
}
func (ah *ArrowHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: ah.Get("normalMatrix")}
}
func (ah *ArrowHelper) SetNormalMatrix(v *math.Matrix3) {
	ah.Set("normalMatrix", v)
}
func (ah *ArrowHelper) OnAfterRender() js.Value {
	return ah.Get("onAfterRender")
}
func (ah *ArrowHelper) SetOnAfterRender(v js.Value) {
	ah.Set("onAfterRender", v)
}
func (ah *ArrowHelper) OnBeforeRender() js.Value {
	return ah.Get("onBeforeRender")
}
func (ah *ArrowHelper) SetOnBeforeRender(v js.Value) {
	ah.Set("onBeforeRender", v)
}
func (ah *ArrowHelper) Parent() core.Object3D {
	return core.Object3D(ah.Get("parent"))
}
func (ah *ArrowHelper) SetParent(v core.Object3D) {
	ah.Set("parent", v)
}
func (ah *ArrowHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("position")}
}
func (ah *ArrowHelper) SetPosition(v *math.Vector3) {
	ah.Set("position", v)
}
func (ah *ArrowHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: ah.Get("quaternion")}
}
func (ah *ArrowHelper) SetQuaternion(v *math.Quaternion) {
	ah.Set("quaternion", v)
}
func (ah *ArrowHelper) ReceiveShadow() bool {
	return ah.Get("receiveShadow").Bool()
}
func (ah *ArrowHelper) SetReceiveShadow(v bool) {
	ah.Set("receiveShadow", v)
}
func (ah *ArrowHelper) RenderOrder() int {
	return ah.Get("renderOrder").Int()
}
func (ah *ArrowHelper) SetRenderOrder(v int) {
	ah.Set("renderOrder", v)
}
func (ah *ArrowHelper) Rotation() *math.Euler {
	return &math.Euler{Value: ah.Get("rotation")}
}
func (ah *ArrowHelper) SetRotation(v *math.Euler) {
	ah.Set("rotation", v)
}
func (ah *ArrowHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("scale")}
}
func (ah *ArrowHelper) SetScale(v *math.Vector3) {
	ah.Set("scale", v)
}
func (ah *ArrowHelper) Type() string {
	return ah.Get("type").String()
}
func (ah *ArrowHelper) SetType(v string) {
	ah.Set("type", v)
}
func (ah *ArrowHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("up")}
}
func (ah *ArrowHelper) SetUp(v *math.Vector3) {
	ah.Set("up", v)
}
func (ah *ArrowHelper) UserData() js.Value {
	return ah.Get("userData")
}
func (ah *ArrowHelper) SetUserData(v js.Value) {
	ah.Set("userData", v)
}
func (ah *ArrowHelper) Uuid() string {
	return ah.Get("uuid").String()
}
func (ah *ArrowHelper) SetUuid(v string) {
	ah.Set("uuid", v)
}
func (ah *ArrowHelper) Visible() bool {
	return ah.Get("visible").Bool()
}
func (ah *ArrowHelper) SetVisible(v bool) {
	ah.Set("visible", v)
}
func (ah *ArrowHelper) DefaultMatrixAutoUpdate() bool {
	return ah.Get("DefaultMatrixAutoUpdate").Bool()
}
func (ah *ArrowHelper) SetDefaultMatrixAutoUpdate(v bool) {
	ah.Set("DefaultMatrixAutoUpdate", v)
}
func (ah *ArrowHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: ah.Get("DefaultUp")}
}
func (ah *ArrowHelper) SetDefaultUp(v *math.Vector3) {
	ah.Set("DefaultUp", v)
}
func (ah *ArrowHelper) Add(object []core.Object3D) this {
	return this(ah.Call("add", object))
}
func (ah *ArrowHelper) AddEventListener(typ string, listener js.Value) {
	ah.Call("addEventListener", typ, listener)
}
func (ah *ArrowHelper) ApplyMatrix(matrix *math.Matrix4) {
	ah.Call("applyMatrix", matrix)
}
func (ah *ArrowHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(ah.Call("applyQuaternion", quaternion))
}
func (ah *ArrowHelper) Clone(recursive bool) this {
	return this(ah.Call("clone", recursive))
}
func (ah *ArrowHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(ah.Call("copy", source, recursive))
}
func (ah *ArrowHelper) DispatchEvent(event js.Value) {
	ah.Call("dispatchEvent", event)
}
func (ah *ArrowHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(ah.Call("getObjectById", id))
}
func (ah *ArrowHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(ah.Call("getObjectByName", name))
}
func (ah *ArrowHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(ah.Call("getObjectByProperty", name, value))
}
func (ah *ArrowHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("getWorldDirection", target)}
}
func (ah *ArrowHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("getWorldPosition", target)}
}
func (ah *ArrowHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: ah.Call("getWorldQuaternion", target)}
}
func (ah *ArrowHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("getWorldScale", target)}
}
func (ah *ArrowHelper) HasEventListener(typ string, listener js.Value) bool {
	return ah.Call("hasEventListener", typ, listener).Bool()
}
func (ah *ArrowHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("localToWorld", vector)}
}
func (ah *ArrowHelper) LookAt(vector math.Vector3, y int, z int) {
	ah.Call("lookAt", vector, y, z)
}
func (ah *ArrowHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	ah.Call("raycast", raycaster, intersects)
}
func (ah *ArrowHelper) Remove(object []core.Object3D) this {
	return this(ah.Call("remove", object))
}
func (ah *ArrowHelper) RemoveEventListener(typ string, listener js.Value) {
	ah.Call("removeEventListener", typ, listener)
}
func (ah *ArrowHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(ah.Call("rotateOnAxis", axis, angle))
}
func (ah *ArrowHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(ah.Call("rotateOnWorldAxis", axis, angle))
}
func (ah *ArrowHelper) RotateX(angle int) this {
	return this(ah.Call("rotateX", angle))
}
func (ah *ArrowHelper) RotateY(angle int) this {
	return this(ah.Call("rotateY", angle))
}
func (ah *ArrowHelper) RotateZ(angle int) this {
	return this(ah.Call("rotateZ", angle))
}
func (ah *ArrowHelper) SetColor(color *math.Color) {
	ah.Call("setColor", color)
}
func (ah *ArrowHelper) SetDirection(dir *math.Vector3) {
	ah.Call("setDirection", dir)
}
func (ah *ArrowHelper) SetLength(length int, headLength int, headWidth int) {
	ah.Call("setLength", length, headLength, headWidth)
}
func (ah *ArrowHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	ah.Call("setRotationFromAxisAngle", axis, angle)
}
func (ah *ArrowHelper) SetRotationFromEuler(euler *math.Euler) {
	ah.Call("setRotationFromEuler", euler)
}
func (ah *ArrowHelper) SetRotationFromMatrix(m *math.Matrix4) {
	ah.Call("setRotationFromMatrix", m)
}
func (ah *ArrowHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	ah.Call("setRotationFromQuaternion", q)
}
func (ah *ArrowHelper) ToJSON(meta js.Value) js.Value {
	return ah.Call("toJSON", meta)
}
func (ah *ArrowHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(ah.Call("translateOnAxis", axis, distance))
}
func (ah *ArrowHelper) TranslateX(distance int) this {
	return this(ah.Call("translateX", distance))
}
func (ah *ArrowHelper) TranslateY(distance int) this {
	return this(ah.Call("translateY", distance))
}
func (ah *ArrowHelper) TranslateZ(distance int) this {
	return this(ah.Call("translateZ", distance))
}
func (ah *ArrowHelper) Traverse(callback js.Value) {
	ah.Call("traverse", callback)
}
func (ah *ArrowHelper) TraverseAncestors(callback js.Value) {
	ah.Call("traverseAncestors", callback)
}
func (ah *ArrowHelper) TraverseVisible(callback js.Value) {
	ah.Call("traverseVisible", callback)
}
func (ah *ArrowHelper) UpdateMatrix() {
	ah.Call("updateMatrix")
}
func (ah *ArrowHelper) UpdateMatrixWorld(force bool) {
	ah.Call("updateMatrixWorld", force)
}
func (ah *ArrowHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ah.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (ah *ArrowHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: ah.Call("worldToLocal", vector)}
}

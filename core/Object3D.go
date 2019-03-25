package core

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

func Object3DIdCount() int {
	return get("Object3DIdCount").Int()
}
func SetObject3DIdCount(v int) {
	set("Object3DIdCount", v)
}

type Object3D struct {
	js.Value
}

func NewObject3D() *Object3D {
	return &Object3D{Value: get("Object3D").New()}
}
func (od *Object3D) CastShadow() bool {
	return od.Get("castShadow").Bool()
}
func (od *Object3D) SetCastShadow(v bool) {
	od.Set("castShadow", v)
}
func (od *Object3D) Children() []Object3D {
	return []Object3D(od.Get("children"))
}
func (od *Object3D) SetChildren(v []Object3D) {
	od.Set("children", v)
}
func (od *Object3D) FrustumCulled() bool {
	return od.Get("frustumCulled").Bool()
}
func (od *Object3D) SetFrustumCulled(v bool) {
	od.Set("frustumCulled", v)
}
func (od *Object3D) Id() int {
	return od.Get("id").Int()
}
func (od *Object3D) SetId(v int) {
	od.Set("id", v)
}
func (od *Object3D) IsObject3D() true {
	return true(od.Get("isObject3D"))
}
func (od *Object3D) SetIsObject3D(v true) {
	od.Set("isObject3D", v)
}
func (od *Object3D) Layers() *Layers {
	return &Layers{Value: od.Get("layers")}
}
func (od *Object3D) SetLayers(v *Layers) {
	od.Set("layers", v)
}
func (od *Object3D) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: od.Get("matrix")}
}
func (od *Object3D) SetMatrix(v *math.Matrix4) {
	od.Set("matrix", v)
}
func (od *Object3D) MatrixAutoUpdate() bool {
	return od.Get("matrixAutoUpdate").Bool()
}
func (od *Object3D) SetMatrixAutoUpdate(v bool) {
	od.Set("matrixAutoUpdate", v)
}
func (od *Object3D) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: od.Get("matrixWorld")}
}
func (od *Object3D) SetMatrixWorld(v *math.Matrix4) {
	od.Set("matrixWorld", v)
}
func (od *Object3D) MatrixWorldNeedsUpdate() bool {
	return od.Get("matrixWorldNeedsUpdate").Bool()
}
func (od *Object3D) SetMatrixWorldNeedsUpdate(v bool) {
	od.Set("matrixWorldNeedsUpdate", v)
}
func (od *Object3D) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: od.Get("modelViewMatrix")}
}
func (od *Object3D) SetModelViewMatrix(v *math.Matrix4) {
	od.Set("modelViewMatrix", v)
}
func (od *Object3D) Name() string {
	return od.Get("name").String()
}
func (od *Object3D) SetName(v string) {
	od.Set("name", v)
}
func (od *Object3D) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: od.Get("normalMatrix")}
}
func (od *Object3D) SetNormalMatrix(v *math.Matrix3) {
	od.Set("normalMatrix", v)
}
func (od *Object3D) OnAfterRender() js.Value {
	return od.Get("onAfterRender")
}
func (od *Object3D) SetOnAfterRender(v js.Value) {
	od.Set("onAfterRender", v)
}
func (od *Object3D) OnBeforeRender() js.Value {
	return od.Get("onBeforeRender")
}
func (od *Object3D) SetOnBeforeRender(v js.Value) {
	od.Set("onBeforeRender", v)
}
func (od *Object3D) Parent() Object3D {
	return Object3D(od.Get("parent"))
}
func (od *Object3D) SetParent(v Object3D) {
	od.Set("parent", v)
}
func (od *Object3D) Position() *math.Vector3 {
	return &math.Vector3{Value: od.Get("position")}
}
func (od *Object3D) SetPosition(v *math.Vector3) {
	od.Set("position", v)
}
func (od *Object3D) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: od.Get("quaternion")}
}
func (od *Object3D) SetQuaternion(v *math.Quaternion) {
	od.Set("quaternion", v)
}
func (od *Object3D) ReceiveShadow() bool {
	return od.Get("receiveShadow").Bool()
}
func (od *Object3D) SetReceiveShadow(v bool) {
	od.Set("receiveShadow", v)
}
func (od *Object3D) RenderOrder() int {
	return od.Get("renderOrder").Int()
}
func (od *Object3D) SetRenderOrder(v int) {
	od.Set("renderOrder", v)
}
func (od *Object3D) Rotation() *math.Euler {
	return &math.Euler{Value: od.Get("rotation")}
}
func (od *Object3D) SetRotation(v *math.Euler) {
	od.Set("rotation", v)
}
func (od *Object3D) Scale() *math.Vector3 {
	return &math.Vector3{Value: od.Get("scale")}
}
func (od *Object3D) SetScale(v *math.Vector3) {
	od.Set("scale", v)
}
func (od *Object3D) Type() string {
	return od.Get("type").String()
}
func (od *Object3D) SetType(v string) {
	od.Set("type", v)
}
func (od *Object3D) Up() *math.Vector3 {
	return &math.Vector3{Value: od.Get("up")}
}
func (od *Object3D) SetUp(v *math.Vector3) {
	od.Set("up", v)
}
func (od *Object3D) UserData() js.Value {
	return od.Get("userData")
}
func (od *Object3D) SetUserData(v js.Value) {
	od.Set("userData", v)
}
func (od *Object3D) Uuid() string {
	return od.Get("uuid").String()
}
func (od *Object3D) SetUuid(v string) {
	od.Set("uuid", v)
}
func (od *Object3D) Visible() bool {
	return od.Get("visible").Bool()
}
func (od *Object3D) SetVisible(v bool) {
	od.Set("visible", v)
}
func (od *Object3D) DefaultMatrixAutoUpdate() bool {
	return od.Get("DefaultMatrixAutoUpdate").Bool()
}
func (od *Object3D) SetDefaultMatrixAutoUpdate(v bool) {
	od.Set("DefaultMatrixAutoUpdate", v)
}
func (od *Object3D) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: od.Get("DefaultUp")}
}
func (od *Object3D) SetDefaultUp(v *math.Vector3) {
	od.Set("DefaultUp", v)
}
func (od *Object3D) Add(object []Object3D) this {
	return this(od.Call("add", object))
}
func (od *Object3D) AddEventListener(typ string, listener js.Value) {
	od.Call("addEventListener", typ, listener)
}
func (od *Object3D) ApplyMatrix(matrix *math.Matrix4) {
	od.Call("applyMatrix", matrix)
}
func (od *Object3D) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(od.Call("applyQuaternion", quaternion))
}
func (od *Object3D) Clone(recursive bool) this {
	return this(od.Call("clone", recursive))
}
func (od *Object3D) Copy(source *Object3D, recursive bool) this {
	return this(od.Call("copy", source, recursive))
}
func (od *Object3D) DispatchEvent(event js.Value) {
	od.Call("dispatchEvent", event)
}
func (od *Object3D) GetObjectById(id int) Object3D {
	return Object3D(od.Call("getObjectById", id))
}
func (od *Object3D) GetObjectByName(name string) Object3D {
	return Object3D(od.Call("getObjectByName", name))
}
func (od *Object3D) GetObjectByProperty(name string, value string) Object3D {
	return Object3D(od.Call("getObjectByProperty", name, value))
}
func (od *Object3D) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: od.Call("getWorldDirection", target)}
}
func (od *Object3D) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: od.Call("getWorldPosition", target)}
}
func (od *Object3D) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: od.Call("getWorldQuaternion", target)}
}
func (od *Object3D) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: od.Call("getWorldScale", target)}
}
func (od *Object3D) HasEventListener(typ string, listener js.Value) bool {
	return od.Call("hasEventListener", typ, listener).Bool()
}
func (od *Object3D) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: od.Call("localToWorld", vector)}
}
func (od *Object3D) LookAt(vector math.Vector3, y int, z int) {
	od.Call("lookAt", vector, y, z)
}
func (od *Object3D) Raycast(raycaster *Raycaster, intersects []Intersection) {
	od.Call("raycast", raycaster, intersects)
}
func (od *Object3D) Remove(object []Object3D) this {
	return this(od.Call("remove", object))
}
func (od *Object3D) RemoveEventListener(typ string, listener js.Value) {
	od.Call("removeEventListener", typ, listener)
}
func (od *Object3D) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(od.Call("rotateOnAxis", axis, angle))
}
func (od *Object3D) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(od.Call("rotateOnWorldAxis", axis, angle))
}
func (od *Object3D) RotateX(angle int) this {
	return this(od.Call("rotateX", angle))
}
func (od *Object3D) RotateY(angle int) this {
	return this(od.Call("rotateY", angle))
}
func (od *Object3D) RotateZ(angle int) this {
	return this(od.Call("rotateZ", angle))
}
func (od *Object3D) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	od.Call("setRotationFromAxisAngle", axis, angle)
}
func (od *Object3D) SetRotationFromEuler(euler *math.Euler) {
	od.Call("setRotationFromEuler", euler)
}
func (od *Object3D) SetRotationFromMatrix(m *math.Matrix4) {
	od.Call("setRotationFromMatrix", m)
}
func (od *Object3D) SetRotationFromQuaternion(q *math.Quaternion) {
	od.Call("setRotationFromQuaternion", q)
}
func (od *Object3D) ToJSON(meta js.Value) js.Value {
	return od.Call("toJSON", meta)
}
func (od *Object3D) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(od.Call("translateOnAxis", axis, distance))
}
func (od *Object3D) TranslateX(distance int) this {
	return this(od.Call("translateX", distance))
}
func (od *Object3D) TranslateY(distance int) this {
	return this(od.Call("translateY", distance))
}
func (od *Object3D) TranslateZ(distance int) this {
	return this(od.Call("translateZ", distance))
}
func (od *Object3D) Traverse(callback js.Value) {
	od.Call("traverse", callback)
}
func (od *Object3D) TraverseAncestors(callback js.Value) {
	od.Call("traverseAncestors", callback)
}
func (od *Object3D) TraverseVisible(callback js.Value) {
	od.Call("traverseVisible", callback)
}
func (od *Object3D) UpdateMatrix() {
	od.Call("updateMatrix")
}
func (od *Object3D) UpdateMatrixWorld(force bool) {
	od.Call("updateMatrixWorld", force)
}
func (od *Object3D) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	od.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (od *Object3D) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: od.Call("worldToLocal", vector)}
}

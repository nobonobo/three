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

type Sprite struct {
	js.Value
}

func NewSprite(material *materials.Material) *Sprite {
	return &Sprite{Value: get("Sprite").New(material)}
}
func (s *Sprite) CastShadow() bool {
	return s.Get("castShadow").Bool()
}
func (s *Sprite) SetCastShadow(v bool) {
	s.Set("castShadow", v)
}
func (s *Sprite) Center() *math.Vector2 {
	return &math.Vector2{Value: s.Get("center")}
}
func (s *Sprite) SetCenter(v *math.Vector2) {
	s.Set("center", v)
}
func (s *Sprite) Children() []core.Object3D {
	return []core.Object3D(s.Get("children"))
}
func (s *Sprite) SetChildren(v []core.Object3D) {
	s.Set("children", v)
}
func (s *Sprite) FrustumCulled() bool {
	return s.Get("frustumCulled").Bool()
}
func (s *Sprite) SetFrustumCulled(v bool) {
	s.Set("frustumCulled", v)
}
func (s *Sprite) Id() int {
	return s.Get("id").Int()
}
func (s *Sprite) SetId(v int) {
	s.Set("id", v)
}
func (s *Sprite) IsObject3D() true {
	return true(s.Get("isObject3D"))
}
func (s *Sprite) SetIsObject3D(v true) {
	s.Set("isObject3D", v)
}
func (s *Sprite) IsSprite() true {
	return true(s.Get("isSprite"))
}
func (s *Sprite) SetIsSprite(v true) {
	s.Set("isSprite", v)
}
func (s *Sprite) Layers() *core.Layers {
	return &core.Layers{Value: s.Get("layers")}
}
func (s *Sprite) SetLayers(v *core.Layers) {
	s.Set("layers", v)
}
func (s *Sprite) Material() *materials.Material {
	return &materials.Material{Value: s.Get("material")}
}
func (s *Sprite) SetMaterial(v *materials.Material) {
	s.Set("material", v)
}
func (s *Sprite) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("matrix")}
}
func (s *Sprite) SetMatrix(v *math.Matrix4) {
	s.Set("matrix", v)
}
func (s *Sprite) MatrixAutoUpdate() bool {
	return s.Get("matrixAutoUpdate").Bool()
}
func (s *Sprite) SetMatrixAutoUpdate(v bool) {
	s.Set("matrixAutoUpdate", v)
}
func (s *Sprite) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("matrixWorld")}
}
func (s *Sprite) SetMatrixWorld(v *math.Matrix4) {
	s.Set("matrixWorld", v)
}
func (s *Sprite) MatrixWorldNeedsUpdate() bool {
	return s.Get("matrixWorldNeedsUpdate").Bool()
}
func (s *Sprite) SetMatrixWorldNeedsUpdate(v bool) {
	s.Set("matrixWorldNeedsUpdate", v)
}
func (s *Sprite) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("modelViewMatrix")}
}
func (s *Sprite) SetModelViewMatrix(v *math.Matrix4) {
	s.Set("modelViewMatrix", v)
}
func (s *Sprite) Name() string {
	return s.Get("name").String()
}
func (s *Sprite) SetName(v string) {
	s.Set("name", v)
}
func (s *Sprite) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: s.Get("normalMatrix")}
}
func (s *Sprite) SetNormalMatrix(v *math.Matrix3) {
	s.Set("normalMatrix", v)
}
func (s *Sprite) OnAfterRender() js.Value {
	return s.Get("onAfterRender")
}
func (s *Sprite) SetOnAfterRender(v js.Value) {
	s.Set("onAfterRender", v)
}
func (s *Sprite) OnBeforeRender() js.Value {
	return s.Get("onBeforeRender")
}
func (s *Sprite) SetOnBeforeRender(v js.Value) {
	s.Set("onBeforeRender", v)
}
func (s *Sprite) Parent() core.Object3D {
	return core.Object3D(s.Get("parent"))
}
func (s *Sprite) SetParent(v core.Object3D) {
	s.Set("parent", v)
}
func (s *Sprite) Position() *math.Vector3 {
	return &math.Vector3{Value: s.Get("position")}
}
func (s *Sprite) SetPosition(v *math.Vector3) {
	s.Set("position", v)
}
func (s *Sprite) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: s.Get("quaternion")}
}
func (s *Sprite) SetQuaternion(v *math.Quaternion) {
	s.Set("quaternion", v)
}
func (s *Sprite) ReceiveShadow() bool {
	return s.Get("receiveShadow").Bool()
}
func (s *Sprite) SetReceiveShadow(v bool) {
	s.Set("receiveShadow", v)
}
func (s *Sprite) RenderOrder() int {
	return s.Get("renderOrder").Int()
}
func (s *Sprite) SetRenderOrder(v int) {
	s.Set("renderOrder", v)
}
func (s *Sprite) Rotation() *math.Euler {
	return &math.Euler{Value: s.Get("rotation")}
}
func (s *Sprite) SetRotation(v *math.Euler) {
	s.Set("rotation", v)
}
func (s *Sprite) Scale() *math.Vector3 {
	return &math.Vector3{Value: s.Get("scale")}
}
func (s *Sprite) SetScale(v *math.Vector3) {
	s.Set("scale", v)
}
func (s *Sprite) Type() string {
	return s.Get("type").String()
}
func (s *Sprite) SetType(v string) {
	s.Set("type", v)
}
func (s *Sprite) Up() *math.Vector3 {
	return &math.Vector3{Value: s.Get("up")}
}
func (s *Sprite) SetUp(v *math.Vector3) {
	s.Set("up", v)
}
func (s *Sprite) UserData() js.Value {
	return s.Get("userData")
}
func (s *Sprite) SetUserData(v js.Value) {
	s.Set("userData", v)
}
func (s *Sprite) Uuid() string {
	return s.Get("uuid").String()
}
func (s *Sprite) SetUuid(v string) {
	s.Set("uuid", v)
}
func (s *Sprite) Visible() bool {
	return s.Get("visible").Bool()
}
func (s *Sprite) SetVisible(v bool) {
	s.Set("visible", v)
}
func (s *Sprite) DefaultMatrixAutoUpdate() bool {
	return s.Get("DefaultMatrixAutoUpdate").Bool()
}
func (s *Sprite) SetDefaultMatrixAutoUpdate(v bool) {
	s.Set("DefaultMatrixAutoUpdate", v)
}
func (s *Sprite) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: s.Get("DefaultUp")}
}
func (s *Sprite) SetDefaultUp(v *math.Vector3) {
	s.Set("DefaultUp", v)
}
func (s *Sprite) Add(object []core.Object3D) this {
	return this(s.Call("add", object))
}
func (s *Sprite) AddEventListener(typ string, listener js.Value) {
	s.Call("addEventListener", typ, listener)
}
func (s *Sprite) ApplyMatrix(matrix *math.Matrix4) {
	s.Call("applyMatrix", matrix)
}
func (s *Sprite) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(s.Call("applyQuaternion", quaternion))
}
func (s *Sprite) Clone(recursive bool) this {
	return this(s.Call("clone", recursive))
}
func (s *Sprite) Copy(source this, recursive bool) this {
	return this(s.Call("copy", source, recursive))
}
func (s *Sprite) DispatchEvent(event js.Value) {
	s.Call("dispatchEvent", event)
}
func (s *Sprite) GetObjectById(id int) core.Object3D {
	return core.Object3D(s.Call("getObjectById", id))
}
func (s *Sprite) GetObjectByName(name string) core.Object3D {
	return core.Object3D(s.Call("getObjectByName", name))
}
func (s *Sprite) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(s.Call("getObjectByProperty", name, value))
}
func (s *Sprite) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("getWorldDirection", target)}
}
func (s *Sprite) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("getWorldPosition", target)}
}
func (s *Sprite) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: s.Call("getWorldQuaternion", target)}
}
func (s *Sprite) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("getWorldScale", target)}
}
func (s *Sprite) HasEventListener(typ string, listener js.Value) bool {
	return s.Call("hasEventListener", typ, listener).Bool()
}
func (s *Sprite) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("localToWorld", vector)}
}
func (s *Sprite) LookAt(vector math.Vector3, y int, z int) {
	s.Call("lookAt", vector, y, z)
}
func (s *Sprite) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	s.Call("raycast", raycaster, intersects)
}
func (s *Sprite) Remove(object []core.Object3D) this {
	return this(s.Call("remove", object))
}
func (s *Sprite) RemoveEventListener(typ string, listener js.Value) {
	s.Call("removeEventListener", typ, listener)
}
func (s *Sprite) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(s.Call("rotateOnAxis", axis, angle))
}
func (s *Sprite) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(s.Call("rotateOnWorldAxis", axis, angle))
}
func (s *Sprite) RotateX(angle int) this {
	return this(s.Call("rotateX", angle))
}
func (s *Sprite) RotateY(angle int) this {
	return this(s.Call("rotateY", angle))
}
func (s *Sprite) RotateZ(angle int) this {
	return this(s.Call("rotateZ", angle))
}
func (s *Sprite) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	s.Call("setRotationFromAxisAngle", axis, angle)
}
func (s *Sprite) SetRotationFromEuler(euler *math.Euler) {
	s.Call("setRotationFromEuler", euler)
}
func (s *Sprite) SetRotationFromMatrix(m *math.Matrix4) {
	s.Call("setRotationFromMatrix", m)
}
func (s *Sprite) SetRotationFromQuaternion(q *math.Quaternion) {
	s.Call("setRotationFromQuaternion", q)
}
func (s *Sprite) ToJSON(meta js.Value) js.Value {
	return s.Call("toJSON", meta)
}
func (s *Sprite) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(s.Call("translateOnAxis", axis, distance))
}
func (s *Sprite) TranslateX(distance int) this {
	return this(s.Call("translateX", distance))
}
func (s *Sprite) TranslateY(distance int) this {
	return this(s.Call("translateY", distance))
}
func (s *Sprite) TranslateZ(distance int) this {
	return this(s.Call("translateZ", distance))
}
func (s *Sprite) Traverse(callback js.Value) {
	s.Call("traverse", callback)
}
func (s *Sprite) TraverseAncestors(callback js.Value) {
	s.Call("traverseAncestors", callback)
}
func (s *Sprite) TraverseVisible(callback js.Value) {
	s.Call("traverseVisible", callback)
}
func (s *Sprite) UpdateMatrix() {
	s.Call("updateMatrix")
}
func (s *Sprite) UpdateMatrixWorld(force bool) {
	s.Call("updateMatrixWorld", force)
}
func (s *Sprite) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	s.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (s *Sprite) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("worldToLocal", vector)}
}

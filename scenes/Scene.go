package scenes

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/cameras"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
)

type Scene struct {
	js.Value
}

func NewScene() *Scene {
	return &Scene{Value: get("Scene").New()}
}
func (s *Scene) AutoUpdate() bool {
	return s.Get("autoUpdate").Bool()
}
func (s *Scene) SetAutoUpdate(v bool) {
	s.Set("autoUpdate", v)
}
func (s *Scene) Background() null {
	return null(s.Get("background"))
}
func (s *Scene) SetBackground(v null) {
	s.Set("background", v)
}
func (s *Scene) CastShadow() bool {
	return s.Get("castShadow").Bool()
}
func (s *Scene) SetCastShadow(v bool) {
	s.Set("castShadow", v)
}
func (s *Scene) Children() []core.Object3D {
	return []core.Object3D(s.Get("children"))
}
func (s *Scene) SetChildren(v []core.Object3D) {
	s.Set("children", v)
}
func (s *Scene) Fog() IFog {
	return IFog(s.Get("fog"))
}
func (s *Scene) SetFog(v IFog) {
	s.Set("fog", v)
}
func (s *Scene) FrustumCulled() bool {
	return s.Get("frustumCulled").Bool()
}
func (s *Scene) SetFrustumCulled(v bool) {
	s.Set("frustumCulled", v)
}
func (s *Scene) Id() int {
	return s.Get("id").Int()
}
func (s *Scene) SetId(v int) {
	s.Set("id", v)
}
func (s *Scene) IsObject3D() true {
	return true(s.Get("isObject3D"))
}
func (s *Scene) SetIsObject3D(v true) {
	s.Set("isObject3D", v)
}
func (s *Scene) Layers() *core.Layers {
	return &core.Layers{Value: s.Get("layers")}
}
func (s *Scene) SetLayers(v *core.Layers) {
	s.Set("layers", v)
}
func (s *Scene) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("matrix")}
}
func (s *Scene) SetMatrix(v *math.Matrix4) {
	s.Set("matrix", v)
}
func (s *Scene) MatrixAutoUpdate() bool {
	return s.Get("matrixAutoUpdate").Bool()
}
func (s *Scene) SetMatrixAutoUpdate(v bool) {
	s.Set("matrixAutoUpdate", v)
}
func (s *Scene) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("matrixWorld")}
}
func (s *Scene) SetMatrixWorld(v *math.Matrix4) {
	s.Set("matrixWorld", v)
}
func (s *Scene) MatrixWorldNeedsUpdate() bool {
	return s.Get("matrixWorldNeedsUpdate").Bool()
}
func (s *Scene) SetMatrixWorldNeedsUpdate(v bool) {
	s.Set("matrixWorldNeedsUpdate", v)
}
func (s *Scene) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("modelViewMatrix")}
}
func (s *Scene) SetModelViewMatrix(v *math.Matrix4) {
	s.Set("modelViewMatrix", v)
}
func (s *Scene) Name() string {
	return s.Get("name").String()
}
func (s *Scene) SetName(v string) {
	s.Set("name", v)
}
func (s *Scene) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: s.Get("normalMatrix")}
}
func (s *Scene) SetNormalMatrix(v *math.Matrix3) {
	s.Set("normalMatrix", v)
}
func (s *Scene) OnAfterRender() js.Value {
	return s.Get("onAfterRender")
}
func (s *Scene) SetOnAfterRender(v js.Value) {
	s.Set("onAfterRender", v)
}
func (s *Scene) OnBeforeRender() js.Value {
	return s.Get("onBeforeRender")
}
func (s *Scene) SetOnBeforeRender(v js.Value) {
	s.Set("onBeforeRender", v)
}
func (s *Scene) OverrideMaterial() materials.Material {
	return materials.Material(s.Get("overrideMaterial"))
}
func (s *Scene) SetOverrideMaterial(v materials.Material) {
	s.Set("overrideMaterial", v)
}
func (s *Scene) Parent() core.Object3D {
	return core.Object3D(s.Get("parent"))
}
func (s *Scene) SetParent(v core.Object3D) {
	s.Set("parent", v)
}
func (s *Scene) Position() *math.Vector3 {
	return &math.Vector3{Value: s.Get("position")}
}
func (s *Scene) SetPosition(v *math.Vector3) {
	s.Set("position", v)
}
func (s *Scene) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: s.Get("quaternion")}
}
func (s *Scene) SetQuaternion(v *math.Quaternion) {
	s.Set("quaternion", v)
}
func (s *Scene) ReceiveShadow() bool {
	return s.Get("receiveShadow").Bool()
}
func (s *Scene) SetReceiveShadow(v bool) {
	s.Set("receiveShadow", v)
}
func (s *Scene) RenderOrder() int {
	return s.Get("renderOrder").Int()
}
func (s *Scene) SetRenderOrder(v int) {
	s.Set("renderOrder", v)
}
func (s *Scene) Rotation() *math.Euler {
	return &math.Euler{Value: s.Get("rotation")}
}
func (s *Scene) SetRotation(v *math.Euler) {
	s.Set("rotation", v)
}
func (s *Scene) Scale() *math.Vector3 {
	return &math.Vector3{Value: s.Get("scale")}
}
func (s *Scene) SetScale(v *math.Vector3) {
	s.Set("scale", v)
}
func (s *Scene) Type() string {
	return s.Get("type").String()
}
func (s *Scene) SetType(v string) {
	s.Set("type", v)
}
func (s *Scene) Up() *math.Vector3 {
	return &math.Vector3{Value: s.Get("up")}
}
func (s *Scene) SetUp(v *math.Vector3) {
	s.Set("up", v)
}
func (s *Scene) UserData() js.Value {
	return s.Get("userData")
}
func (s *Scene) SetUserData(v js.Value) {
	s.Set("userData", v)
}
func (s *Scene) Uuid() string {
	return s.Get("uuid").String()
}
func (s *Scene) SetUuid(v string) {
	s.Set("uuid", v)
}
func (s *Scene) Visible() bool {
	return s.Get("visible").Bool()
}
func (s *Scene) SetVisible(v bool) {
	s.Set("visible", v)
}
func (s *Scene) DefaultMatrixAutoUpdate() bool {
	return s.Get("DefaultMatrixAutoUpdate").Bool()
}
func (s *Scene) SetDefaultMatrixAutoUpdate(v bool) {
	s.Set("DefaultMatrixAutoUpdate", v)
}
func (s *Scene) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: s.Get("DefaultUp")}
}
func (s *Scene) SetDefaultUp(v *math.Vector3) {
	s.Set("DefaultUp", v)
}
func (s *Scene) Add(object []core.Object3D) this {
	return this(s.Call("add", object))
}
func (s *Scene) AddEventListener(typ string, listener js.Value) {
	s.Call("addEventListener", typ, listener)
}
func (s *Scene) ApplyMatrix(matrix *math.Matrix4) {
	s.Call("applyMatrix", matrix)
}
func (s *Scene) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(s.Call("applyQuaternion", quaternion))
}
func (s *Scene) Clone(recursive bool) this {
	return this(s.Call("clone", recursive))
}
func (s *Scene) Copy(source this, recursive bool) this {
	return this(s.Call("copy", source, recursive))
}
func (s *Scene) DispatchEvent(event js.Value) {
	s.Call("dispatchEvent", event)
}
func (s *Scene) GetObjectById(id int) core.Object3D {
	return core.Object3D(s.Call("getObjectById", id))
}
func (s *Scene) GetObjectByName(name string) core.Object3D {
	return core.Object3D(s.Call("getObjectByName", name))
}
func (s *Scene) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(s.Call("getObjectByProperty", name, value))
}
func (s *Scene) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("getWorldDirection", target)}
}
func (s *Scene) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("getWorldPosition", target)}
}
func (s *Scene) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: s.Call("getWorldQuaternion", target)}
}
func (s *Scene) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("getWorldScale", target)}
}
func (s *Scene) HasEventListener(typ string, listener js.Value) bool {
	return s.Call("hasEventListener", typ, listener).Bool()
}
func (s *Scene) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("localToWorld", vector)}
}
func (s *Scene) LookAt(vector math.Vector3, y int, z int) {
	s.Call("lookAt", vector, y, z)
}
func (s *Scene) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	s.Call("raycast", raycaster, intersects)
}
func (s *Scene) Remove(object []core.Object3D) this {
	return this(s.Call("remove", object))
}
func (s *Scene) RemoveEventListener(typ string, listener js.Value) {
	s.Call("removeEventListener", typ, listener)
}
func (s *Scene) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(s.Call("rotateOnAxis", axis, angle))
}
func (s *Scene) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(s.Call("rotateOnWorldAxis", axis, angle))
}
func (s *Scene) RotateX(angle int) this {
	return this(s.Call("rotateX", angle))
}
func (s *Scene) RotateY(angle int) this {
	return this(s.Call("rotateY", angle))
}
func (s *Scene) RotateZ(angle int) this {
	return this(s.Call("rotateZ", angle))
}
func (s *Scene) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	s.Call("setRotationFromAxisAngle", axis, angle)
}
func (s *Scene) SetRotationFromEuler(euler *math.Euler) {
	s.Call("setRotationFromEuler", euler)
}
func (s *Scene) SetRotationFromMatrix(m *math.Matrix4) {
	s.Call("setRotationFromMatrix", m)
}
func (s *Scene) SetRotationFromQuaternion(q *math.Quaternion) {
	s.Call("setRotationFromQuaternion", q)
}
func (s *Scene) ToJSON(meta js.Value) js.Value {
	return s.Call("toJSON", meta)
}
func (s *Scene) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(s.Call("translateOnAxis", axis, distance))
}
func (s *Scene) TranslateX(distance int) this {
	return this(s.Call("translateX", distance))
}
func (s *Scene) TranslateY(distance int) this {
	return this(s.Call("translateY", distance))
}
func (s *Scene) TranslateZ(distance int) this {
	return this(s.Call("translateZ", distance))
}
func (s *Scene) Traverse(callback js.Value) {
	s.Call("traverse", callback)
}
func (s *Scene) TraverseAncestors(callback js.Value) {
	s.Call("traverseAncestors", callback)
}
func (s *Scene) TraverseVisible(callback js.Value) {
	s.Call("traverseVisible", callback)
}
func (s *Scene) UpdateMatrix() {
	s.Call("updateMatrix")
}
func (s *Scene) UpdateMatrixWorld(force bool) {
	s.Call("updateMatrixWorld", force)
}
func (s *Scene) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	s.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (s *Scene) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: s.Call("worldToLocal", vector)}
}

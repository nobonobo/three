package audio

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

type AudioListener struct {
	js.Value
}

func NewAudioListener() *AudioListener {
	return &AudioListener{Value: get("AudioListener").New()}
}
func (al *AudioListener) CastShadow() bool {
	return al.Get("castShadow").Bool()
}
func (al *AudioListener) SetCastShadow(v bool) {
	al.Set("castShadow", v)
}
func (al *AudioListener) Children() []core.Object3D {
	return []core.Object3D(al.Get("children"))
}
func (al *AudioListener) SetChildren(v []core.Object3D) {
	al.Set("children", v)
}
func (al *AudioListener) Context() *AudioContext {
	return &AudioContext{Value: al.Get("context")}
}
func (al *AudioListener) SetContext(v *AudioContext) {
	al.Set("context", v)
}
func (al *AudioListener) Filter() null {
	return null(al.Get("filter"))
}
func (al *AudioListener) SetFilter(v null) {
	al.Set("filter", v)
}
func (al *AudioListener) FrustumCulled() bool {
	return al.Get("frustumCulled").Bool()
}
func (al *AudioListener) SetFrustumCulled(v bool) {
	al.Set("frustumCulled", v)
}
func (al *AudioListener) Gain() *GainNode {
	return &GainNode{Value: al.Get("gain")}
}
func (al *AudioListener) SetGain(v *GainNode) {
	al.Set("gain", v)
}
func (al *AudioListener) Id() int {
	return al.Get("id").Int()
}
func (al *AudioListener) SetId(v int) {
	al.Set("id", v)
}
func (al *AudioListener) IsObject3D() true {
	return true(al.Get("isObject3D"))
}
func (al *AudioListener) SetIsObject3D(v true) {
	al.Set("isObject3D", v)
}
func (al *AudioListener) Layers() *core.Layers {
	return &core.Layers{Value: al.Get("layers")}
}
func (al *AudioListener) SetLayers(v *core.Layers) {
	al.Set("layers", v)
}
func (al *AudioListener) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: al.Get("matrix")}
}
func (al *AudioListener) SetMatrix(v *math.Matrix4) {
	al.Set("matrix", v)
}
func (al *AudioListener) MatrixAutoUpdate() bool {
	return al.Get("matrixAutoUpdate").Bool()
}
func (al *AudioListener) SetMatrixAutoUpdate(v bool) {
	al.Set("matrixAutoUpdate", v)
}
func (al *AudioListener) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: al.Get("matrixWorld")}
}
func (al *AudioListener) SetMatrixWorld(v *math.Matrix4) {
	al.Set("matrixWorld", v)
}
func (al *AudioListener) MatrixWorldNeedsUpdate() bool {
	return al.Get("matrixWorldNeedsUpdate").Bool()
}
func (al *AudioListener) SetMatrixWorldNeedsUpdate(v bool) {
	al.Set("matrixWorldNeedsUpdate", v)
}
func (al *AudioListener) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: al.Get("modelViewMatrix")}
}
func (al *AudioListener) SetModelViewMatrix(v *math.Matrix4) {
	al.Set("modelViewMatrix", v)
}
func (al *AudioListener) Name() string {
	return al.Get("name").String()
}
func (al *AudioListener) SetName(v string) {
	al.Set("name", v)
}
func (al *AudioListener) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: al.Get("normalMatrix")}
}
func (al *AudioListener) SetNormalMatrix(v *math.Matrix3) {
	al.Set("normalMatrix", v)
}
func (al *AudioListener) OnAfterRender() js.Value {
	return al.Get("onAfterRender")
}
func (al *AudioListener) SetOnAfterRender(v js.Value) {
	al.Set("onAfterRender", v)
}
func (al *AudioListener) OnBeforeRender() js.Value {
	return al.Get("onBeforeRender")
}
func (al *AudioListener) SetOnBeforeRender(v js.Value) {
	al.Set("onBeforeRender", v)
}
func (al *AudioListener) Parent() core.Object3D {
	return core.Object3D(al.Get("parent"))
}
func (al *AudioListener) SetParent(v core.Object3D) {
	al.Set("parent", v)
}
func (al *AudioListener) Position() *math.Vector3 {
	return &math.Vector3{Value: al.Get("position")}
}
func (al *AudioListener) SetPosition(v *math.Vector3) {
	al.Set("position", v)
}
func (al *AudioListener) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: al.Get("quaternion")}
}
func (al *AudioListener) SetQuaternion(v *math.Quaternion) {
	al.Set("quaternion", v)
}
func (al *AudioListener) ReceiveShadow() bool {
	return al.Get("receiveShadow").Bool()
}
func (al *AudioListener) SetReceiveShadow(v bool) {
	al.Set("receiveShadow", v)
}
func (al *AudioListener) RenderOrder() int {
	return al.Get("renderOrder").Int()
}
func (al *AudioListener) SetRenderOrder(v int) {
	al.Set("renderOrder", v)
}
func (al *AudioListener) Rotation() *math.Euler {
	return &math.Euler{Value: al.Get("rotation")}
}
func (al *AudioListener) SetRotation(v *math.Euler) {
	al.Set("rotation", v)
}
func (al *AudioListener) Scale() *math.Vector3 {
	return &math.Vector3{Value: al.Get("scale")}
}
func (al *AudioListener) SetScale(v *math.Vector3) {
	al.Set("scale", v)
}
func (al *AudioListener) Type() string {
	return al.Get("type").String()
}
func (al *AudioListener) SetType(v string) {
	al.Set("type", v)
}
func (al *AudioListener) Up() *math.Vector3 {
	return &math.Vector3{Value: al.Get("up")}
}
func (al *AudioListener) SetUp(v *math.Vector3) {
	al.Set("up", v)
}
func (al *AudioListener) UserData() js.Value {
	return al.Get("userData")
}
func (al *AudioListener) SetUserData(v js.Value) {
	al.Set("userData", v)
}
func (al *AudioListener) Uuid() string {
	return al.Get("uuid").String()
}
func (al *AudioListener) SetUuid(v string) {
	al.Set("uuid", v)
}
func (al *AudioListener) Visible() bool {
	return al.Get("visible").Bool()
}
func (al *AudioListener) SetVisible(v bool) {
	al.Set("visible", v)
}
func (al *AudioListener) DefaultMatrixAutoUpdate() bool {
	return al.Get("DefaultMatrixAutoUpdate").Bool()
}
func (al *AudioListener) SetDefaultMatrixAutoUpdate(v bool) {
	al.Set("DefaultMatrixAutoUpdate", v)
}
func (al *AudioListener) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: al.Get("DefaultUp")}
}
func (al *AudioListener) SetDefaultUp(v *math.Vector3) {
	al.Set("DefaultUp", v)
}
func (al *AudioListener) Add(object []core.Object3D) this {
	return this(al.Call("add", object))
}
func (al *AudioListener) AddEventListener(typ string, listener js.Value) {
	al.Call("addEventListener", typ, listener)
}
func (al *AudioListener) ApplyMatrix(matrix *math.Matrix4) {
	al.Call("applyMatrix", matrix)
}
func (al *AudioListener) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(al.Call("applyQuaternion", quaternion))
}
func (al *AudioListener) Clone(recursive bool) this {
	return this(al.Call("clone", recursive))
}
func (al *AudioListener) Copy(source *core.Object3D, recursive bool) this {
	return this(al.Call("copy", source, recursive))
}
func (al *AudioListener) DispatchEvent(event js.Value) {
	al.Call("dispatchEvent", event)
}
func (al *AudioListener) GetFilter() js.Value {
	return al.Call("getFilter")
}
func (al *AudioListener) GetInput() *GainNode {
	return &GainNode{Value: al.Call("getInput")}
}
func (al *AudioListener) GetMasterVolume() int {
	return al.Call("getMasterVolume").Int()
}
func (al *AudioListener) GetObjectById(id int) core.Object3D {
	return core.Object3D(al.Call("getObjectById", id))
}
func (al *AudioListener) GetObjectByName(name string) core.Object3D {
	return core.Object3D(al.Call("getObjectByName", name))
}
func (al *AudioListener) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(al.Call("getObjectByProperty", name, value))
}
func (al *AudioListener) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("getWorldDirection", target)}
}
func (al *AudioListener) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("getWorldPosition", target)}
}
func (al *AudioListener) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: al.Call("getWorldQuaternion", target)}
}
func (al *AudioListener) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("getWorldScale", target)}
}
func (al *AudioListener) HasEventListener(typ string, listener js.Value) bool {
	return al.Call("hasEventListener", typ, listener).Bool()
}
func (al *AudioListener) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("localToWorld", vector)}
}
func (al *AudioListener) LookAt(vector math.Vector3, y int, z int) {
	al.Call("lookAt", vector, y, z)
}
func (al *AudioListener) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	al.Call("raycast", raycaster, intersects)
}
func (al *AudioListener) Remove(object []core.Object3D) this {
	return this(al.Call("remove", object))
}
func (al *AudioListener) RemoveEventListener(typ string, listener js.Value) {
	al.Call("removeEventListener", typ, listener)
}
func (al *AudioListener) RemoveFilter() {
	al.Call("removeFilter")
}
func (al *AudioListener) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(al.Call("rotateOnAxis", axis, angle))
}
func (al *AudioListener) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(al.Call("rotateOnWorldAxis", axis, angle))
}
func (al *AudioListener) RotateX(angle int) this {
	return this(al.Call("rotateX", angle))
}
func (al *AudioListener) RotateY(angle int) this {
	return this(al.Call("rotateY", angle))
}
func (al *AudioListener) RotateZ(angle int) this {
	return this(al.Call("rotateZ", angle))
}
func (al *AudioListener) SetFilter(value js.Value) {
	al.Call("setFilter", value)
}
func (al *AudioListener) SetMasterVolume(value int) {
	al.Call("setMasterVolume", value)
}
func (al *AudioListener) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	al.Call("setRotationFromAxisAngle", axis, angle)
}
func (al *AudioListener) SetRotationFromEuler(euler *math.Euler) {
	al.Call("setRotationFromEuler", euler)
}
func (al *AudioListener) SetRotationFromMatrix(m *math.Matrix4) {
	al.Call("setRotationFromMatrix", m)
}
func (al *AudioListener) SetRotationFromQuaternion(q *math.Quaternion) {
	al.Call("setRotationFromQuaternion", q)
}
func (al *AudioListener) ToJSON(meta js.Value) js.Value {
	return al.Call("toJSON", meta)
}
func (al *AudioListener) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(al.Call("translateOnAxis", axis, distance))
}
func (al *AudioListener) TranslateX(distance int) this {
	return this(al.Call("translateX", distance))
}
func (al *AudioListener) TranslateY(distance int) this {
	return this(al.Call("translateY", distance))
}
func (al *AudioListener) TranslateZ(distance int) this {
	return this(al.Call("translateZ", distance))
}
func (al *AudioListener) Traverse(callback js.Value) {
	al.Call("traverse", callback)
}
func (al *AudioListener) TraverseAncestors(callback js.Value) {
	al.Call("traverseAncestors", callback)
}
func (al *AudioListener) TraverseVisible(callback js.Value) {
	al.Call("traverseVisible", callback)
}
func (al *AudioListener) UpdateMatrix() {
	al.Call("updateMatrix")
}
func (al *AudioListener) UpdateMatrixWorld(force bool) {
	al.Call("updateMatrixWorld", force)
}
func (al *AudioListener) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	al.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (al *AudioListener) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: al.Call("worldToLocal", vector)}
}

package cameras

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/materials"
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/objects"
	"github.com/nobonobo/three/renderers"
	"github.com/nobonobo/three/scenes"
)

type Camera struct {
	js.Value
}

func NewCamera() *Camera {
	return &Camera{Value: get("Camera").New()}
}
func (c *Camera) CastShadow() bool {
	return c.Get("castShadow").Bool()
}
func (c *Camera) SetCastShadow(v bool) {
	c.Set("castShadow", v)
}
func (c *Camera) Children() []core.Object3D {
	return []core.Object3D(c.Get("children"))
}
func (c *Camera) SetChildren(v []core.Object3D) {
	c.Set("children", v)
}
func (c *Camera) FrustumCulled() bool {
	return c.Get("frustumCulled").Bool()
}
func (c *Camera) SetFrustumCulled(v bool) {
	c.Set("frustumCulled", v)
}
func (c *Camera) Id() int {
	return c.Get("id").Int()
}
func (c *Camera) SetId(v int) {
	c.Set("id", v)
}
func (c *Camera) IsCamera() true {
	return true(c.Get("isCamera"))
}
func (c *Camera) SetIsCamera(v true) {
	c.Set("isCamera", v)
}
func (c *Camera) IsObject3D() true {
	return true(c.Get("isObject3D"))
}
func (c *Camera) SetIsObject3D(v true) {
	c.Set("isObject3D", v)
}
func (c *Camera) Layers() *core.Layers {
	return &core.Layers{Value: c.Get("layers")}
}
func (c *Camera) SetLayers(v *core.Layers) {
	c.Set("layers", v)
}
func (c *Camera) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: c.Get("matrix")}
}
func (c *Camera) SetMatrix(v *math.Matrix4) {
	c.Set("matrix", v)
}
func (c *Camera) MatrixAutoUpdate() bool {
	return c.Get("matrixAutoUpdate").Bool()
}
func (c *Camera) SetMatrixAutoUpdate(v bool) {
	c.Set("matrixAutoUpdate", v)
}
func (c *Camera) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: c.Get("matrixWorld")}
}
func (c *Camera) SetMatrixWorld(v *math.Matrix4) {
	c.Set("matrixWorld", v)
}
func (c *Camera) MatrixWorldInverse() *math.Matrix4 {
	return &math.Matrix4{Value: c.Get("matrixWorldInverse")}
}
func (c *Camera) SetMatrixWorldInverse(v *math.Matrix4) {
	c.Set("matrixWorldInverse", v)
}
func (c *Camera) MatrixWorldNeedsUpdate() bool {
	return c.Get("matrixWorldNeedsUpdate").Bool()
}
func (c *Camera) SetMatrixWorldNeedsUpdate(v bool) {
	c.Set("matrixWorldNeedsUpdate", v)
}
func (c *Camera) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: c.Get("modelViewMatrix")}
}
func (c *Camera) SetModelViewMatrix(v *math.Matrix4) {
	c.Set("modelViewMatrix", v)
}
func (c *Camera) Name() string {
	return c.Get("name").String()
}
func (c *Camera) SetName(v string) {
	c.Set("name", v)
}
func (c *Camera) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: c.Get("normalMatrix")}
}
func (c *Camera) SetNormalMatrix(v *math.Matrix3) {
	c.Set("normalMatrix", v)
}
func (c *Camera) OnAfterRender() js.Value {
	return c.Get("onAfterRender")
}
func (c *Camera) SetOnAfterRender(v js.Value) {
	c.Set("onAfterRender", v)
}
func (c *Camera) OnBeforeRender() js.Value {
	return c.Get("onBeforeRender")
}
func (c *Camera) SetOnBeforeRender(v js.Value) {
	c.Set("onBeforeRender", v)
}
func (c *Camera) Parent() core.Object3D {
	return core.Object3D(c.Get("parent"))
}
func (c *Camera) SetParent(v core.Object3D) {
	c.Set("parent", v)
}
func (c *Camera) Position() *math.Vector3 {
	return &math.Vector3{Value: c.Get("position")}
}
func (c *Camera) SetPosition(v *math.Vector3) {
	c.Set("position", v)
}
func (c *Camera) ProjectionMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: c.Get("projectionMatrix")}
}
func (c *Camera) SetProjectionMatrix(v *math.Matrix4) {
	c.Set("projectionMatrix", v)
}
func (c *Camera) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: c.Get("quaternion")}
}
func (c *Camera) SetQuaternion(v *math.Quaternion) {
	c.Set("quaternion", v)
}
func (c *Camera) ReceiveShadow() bool {
	return c.Get("receiveShadow").Bool()
}
func (c *Camera) SetReceiveShadow(v bool) {
	c.Set("receiveShadow", v)
}
func (c *Camera) RenderOrder() int {
	return c.Get("renderOrder").Int()
}
func (c *Camera) SetRenderOrder(v int) {
	c.Set("renderOrder", v)
}
func (c *Camera) Rotation() *math.Euler {
	return &math.Euler{Value: c.Get("rotation")}
}
func (c *Camera) SetRotation(v *math.Euler) {
	c.Set("rotation", v)
}
func (c *Camera) Scale() *math.Vector3 {
	return &math.Vector3{Value: c.Get("scale")}
}
func (c *Camera) SetScale(v *math.Vector3) {
	c.Set("scale", v)
}
func (c *Camera) Type() string {
	return c.Get("type").String()
}
func (c *Camera) SetType(v string) {
	c.Set("type", v)
}
func (c *Camera) Up() *math.Vector3 {
	return &math.Vector3{Value: c.Get("up")}
}
func (c *Camera) SetUp(v *math.Vector3) {
	c.Set("up", v)
}
func (c *Camera) UserData() js.Value {
	return c.Get("userData")
}
func (c *Camera) SetUserData(v js.Value) {
	c.Set("userData", v)
}
func (c *Camera) Uuid() string {
	return c.Get("uuid").String()
}
func (c *Camera) SetUuid(v string) {
	c.Set("uuid", v)
}
func (c *Camera) Visible() bool {
	return c.Get("visible").Bool()
}
func (c *Camera) SetVisible(v bool) {
	c.Set("visible", v)
}
func (c *Camera) DefaultMatrixAutoUpdate() bool {
	return c.Get("DefaultMatrixAutoUpdate").Bool()
}
func (c *Camera) SetDefaultMatrixAutoUpdate(v bool) {
	c.Set("DefaultMatrixAutoUpdate", v)
}
func (c *Camera) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: c.Get("DefaultUp")}
}
func (c *Camera) SetDefaultUp(v *math.Vector3) {
	c.Set("DefaultUp", v)
}
func (c *Camera) Add(object []core.Object3D) this {
	return this(c.Call("add", object))
}
func (c *Camera) AddEventListener(typ string, listener js.Value) {
	c.Call("addEventListener", typ, listener)
}
func (c *Camera) ApplyMatrix(matrix *math.Matrix4) {
	c.Call("applyMatrix", matrix)
}
func (c *Camera) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(c.Call("applyQuaternion", quaternion))
}
func (c *Camera) Clone(recursive bool) this {
	return this(c.Call("clone", recursive))
}
func (c *Camera) Copy(source *Camera, recursive bool) this {
	return this(c.Call("copy", source, recursive))
}
func (c *Camera) DispatchEvent(event js.Value) {
	c.Call("dispatchEvent", event)
}
func (c *Camera) GetObjectById(id int) core.Object3D {
	return core.Object3D(c.Call("getObjectById", id))
}
func (c *Camera) GetObjectByName(name string) core.Object3D {
	return core.Object3D(c.Call("getObjectByName", name))
}
func (c *Camera) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(c.Call("getObjectByProperty", name, value))
}
func (c *Camera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: c.Call("getWorldDirection", target)}
}
func (c *Camera) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: c.Call("getWorldPosition", target)}
}
func (c *Camera) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: c.Call("getWorldQuaternion", target)}
}
func (c *Camera) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: c.Call("getWorldScale", target)}
}
func (c *Camera) HasEventListener(typ string, listener js.Value) bool {
	return c.Call("hasEventListener", typ, listener).Bool()
}
func (c *Camera) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: c.Call("localToWorld", vector)}
}
func (c *Camera) LookAt(vector math.Vector3, y int, z int) {
	c.Call("lookAt", vector, y, z)
}
func (c *Camera) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	c.Call("raycast", raycaster, intersects)
}
func (c *Camera) Remove(object []core.Object3D) this {
	return this(c.Call("remove", object))
}
func (c *Camera) RemoveEventListener(typ string, listener js.Value) {
	c.Call("removeEventListener", typ, listener)
}
func (c *Camera) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(c.Call("rotateOnAxis", axis, angle))
}
func (c *Camera) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(c.Call("rotateOnWorldAxis", axis, angle))
}
func (c *Camera) RotateX(angle int) this {
	return this(c.Call("rotateX", angle))
}
func (c *Camera) RotateY(angle int) this {
	return this(c.Call("rotateY", angle))
}
func (c *Camera) RotateZ(angle int) this {
	return this(c.Call("rotateZ", angle))
}
func (c *Camera) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	c.Call("setRotationFromAxisAngle", axis, angle)
}
func (c *Camera) SetRotationFromEuler(euler *math.Euler) {
	c.Call("setRotationFromEuler", euler)
}
func (c *Camera) SetRotationFromMatrix(m *math.Matrix4) {
	c.Call("setRotationFromMatrix", m)
}
func (c *Camera) SetRotationFromQuaternion(q *math.Quaternion) {
	c.Call("setRotationFromQuaternion", q)
}
func (c *Camera) ToJSON(meta js.Value) js.Value {
	return c.Call("toJSON", meta)
}
func (c *Camera) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(c.Call("translateOnAxis", axis, distance))
}
func (c *Camera) TranslateX(distance int) this {
	return this(c.Call("translateX", distance))
}
func (c *Camera) TranslateY(distance int) this {
	return this(c.Call("translateY", distance))
}
func (c *Camera) TranslateZ(distance int) this {
	return this(c.Call("translateZ", distance))
}
func (c *Camera) Traverse(callback js.Value) {
	c.Call("traverse", callback)
}
func (c *Camera) TraverseAncestors(callback js.Value) {
	c.Call("traverseAncestors", callback)
}
func (c *Camera) TraverseVisible(callback js.Value) {
	c.Call("traverseVisible", callback)
}
func (c *Camera) UpdateMatrix() {
	c.Call("updateMatrix")
}
func (c *Camera) UpdateMatrixWorld(force bool) {
	c.Call("updateMatrixWorld", force)
}
func (c *Camera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	c.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (c *Camera) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: c.Call("worldToLocal", vector)}
}

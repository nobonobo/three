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

type CubeCamera struct {
	js.Value
}

func NewCubeCamera(near int, far int, cubeResolution int) *CubeCamera {
	return &CubeCamera{Value: get("CubeCamera").New(near, far, cubeResolution)}
}
func (cc *CubeCamera) CastShadow() bool {
	return cc.Get("castShadow").Bool()
}
func (cc *CubeCamera) SetCastShadow(v bool) {
	cc.Set("castShadow", v)
}
func (cc *CubeCamera) Children() []core.Object3D {
	return []core.Object3D(cc.Get("children"))
}
func (cc *CubeCamera) SetChildren(v []core.Object3D) {
	cc.Set("children", v)
}
func (cc *CubeCamera) FrustumCulled() bool {
	return cc.Get("frustumCulled").Bool()
}
func (cc *CubeCamera) SetFrustumCulled(v bool) {
	cc.Set("frustumCulled", v)
}
func (cc *CubeCamera) Id() int {
	return cc.Get("id").Int()
}
func (cc *CubeCamera) SetId(v int) {
	cc.Set("id", v)
}
func (cc *CubeCamera) IsObject3D() true {
	return true(cc.Get("isObject3D"))
}
func (cc *CubeCamera) SetIsObject3D(v true) {
	cc.Set("isObject3D", v)
}
func (cc *CubeCamera) Layers() *core.Layers {
	return &core.Layers{Value: cc.Get("layers")}
}
func (cc *CubeCamera) SetLayers(v *core.Layers) {
	cc.Set("layers", v)
}
func (cc *CubeCamera) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: cc.Get("matrix")}
}
func (cc *CubeCamera) SetMatrix(v *math.Matrix4) {
	cc.Set("matrix", v)
}
func (cc *CubeCamera) MatrixAutoUpdate() bool {
	return cc.Get("matrixAutoUpdate").Bool()
}
func (cc *CubeCamera) SetMatrixAutoUpdate(v bool) {
	cc.Set("matrixAutoUpdate", v)
}
func (cc *CubeCamera) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: cc.Get("matrixWorld")}
}
func (cc *CubeCamera) SetMatrixWorld(v *math.Matrix4) {
	cc.Set("matrixWorld", v)
}
func (cc *CubeCamera) MatrixWorldNeedsUpdate() bool {
	return cc.Get("matrixWorldNeedsUpdate").Bool()
}
func (cc *CubeCamera) SetMatrixWorldNeedsUpdate(v bool) {
	cc.Set("matrixWorldNeedsUpdate", v)
}
func (cc *CubeCamera) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: cc.Get("modelViewMatrix")}
}
func (cc *CubeCamera) SetModelViewMatrix(v *math.Matrix4) {
	cc.Set("modelViewMatrix", v)
}
func (cc *CubeCamera) Name() string {
	return cc.Get("name").String()
}
func (cc *CubeCamera) SetName(v string) {
	cc.Set("name", v)
}
func (cc *CubeCamera) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: cc.Get("normalMatrix")}
}
func (cc *CubeCamera) SetNormalMatrix(v *math.Matrix3) {
	cc.Set("normalMatrix", v)
}
func (cc *CubeCamera) OnAfterRender() js.Value {
	return cc.Get("onAfterRender")
}
func (cc *CubeCamera) SetOnAfterRender(v js.Value) {
	cc.Set("onAfterRender", v)
}
func (cc *CubeCamera) OnBeforeRender() js.Value {
	return cc.Get("onBeforeRender")
}
func (cc *CubeCamera) SetOnBeforeRender(v js.Value) {
	cc.Set("onBeforeRender", v)
}
func (cc *CubeCamera) Parent() core.Object3D {
	return core.Object3D(cc.Get("parent"))
}
func (cc *CubeCamera) SetParent(v core.Object3D) {
	cc.Set("parent", v)
}
func (cc *CubeCamera) Position() *math.Vector3 {
	return &math.Vector3{Value: cc.Get("position")}
}
func (cc *CubeCamera) SetPosition(v *math.Vector3) {
	cc.Set("position", v)
}
func (cc *CubeCamera) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: cc.Get("quaternion")}
}
func (cc *CubeCamera) SetQuaternion(v *math.Quaternion) {
	cc.Set("quaternion", v)
}
func (cc *CubeCamera) ReceiveShadow() bool {
	return cc.Get("receiveShadow").Bool()
}
func (cc *CubeCamera) SetReceiveShadow(v bool) {
	cc.Set("receiveShadow", v)
}
func (cc *CubeCamera) RenderOrder() int {
	return cc.Get("renderOrder").Int()
}
func (cc *CubeCamera) SetRenderOrder(v int) {
	cc.Set("renderOrder", v)
}
func (cc *CubeCamera) RenderTarget() *renderers.WebGLRenderTargetCube {
	return &renderers.WebGLRenderTargetCube{Value: cc.Get("renderTarget")}
}
func (cc *CubeCamera) SetRenderTarget(v *renderers.WebGLRenderTargetCube) {
	cc.Set("renderTarget", v)
}
func (cc *CubeCamera) Rotation() *math.Euler {
	return &math.Euler{Value: cc.Get("rotation")}
}
func (cc *CubeCamera) SetRotation(v *math.Euler) {
	cc.Set("rotation", v)
}
func (cc *CubeCamera) Scale() *math.Vector3 {
	return &math.Vector3{Value: cc.Get("scale")}
}
func (cc *CubeCamera) SetScale(v *math.Vector3) {
	cc.Set("scale", v)
}
func (cc *CubeCamera) Type() string {
	return cc.Get("type").String()
}
func (cc *CubeCamera) SetType(v string) {
	cc.Set("type", v)
}
func (cc *CubeCamera) Up() *math.Vector3 {
	return &math.Vector3{Value: cc.Get("up")}
}
func (cc *CubeCamera) SetUp(v *math.Vector3) {
	cc.Set("up", v)
}
func (cc *CubeCamera) UserData() js.Value {
	return cc.Get("userData")
}
func (cc *CubeCamera) SetUserData(v js.Value) {
	cc.Set("userData", v)
}
func (cc *CubeCamera) Uuid() string {
	return cc.Get("uuid").String()
}
func (cc *CubeCamera) SetUuid(v string) {
	cc.Set("uuid", v)
}
func (cc *CubeCamera) Visible() bool {
	return cc.Get("visible").Bool()
}
func (cc *CubeCamera) SetVisible(v bool) {
	cc.Set("visible", v)
}
func (cc *CubeCamera) DefaultMatrixAutoUpdate() bool {
	return cc.Get("DefaultMatrixAutoUpdate").Bool()
}
func (cc *CubeCamera) SetDefaultMatrixAutoUpdate(v bool) {
	cc.Set("DefaultMatrixAutoUpdate", v)
}
func (cc *CubeCamera) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: cc.Get("DefaultUp")}
}
func (cc *CubeCamera) SetDefaultUp(v *math.Vector3) {
	cc.Set("DefaultUp", v)
}
func (cc *CubeCamera) Add(object []core.Object3D) this {
	return this(cc.Call("add", object))
}
func (cc *CubeCamera) AddEventListener(typ string, listener js.Value) {
	cc.Call("addEventListener", typ, listener)
}
func (cc *CubeCamera) ApplyMatrix(matrix *math.Matrix4) {
	cc.Call("applyMatrix", matrix)
}
func (cc *CubeCamera) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(cc.Call("applyQuaternion", quaternion))
}
func (cc *CubeCamera) Clone(recursive bool) this {
	return this(cc.Call("clone", recursive))
}
func (cc *CubeCamera) Copy(source *core.Object3D, recursive bool) this {
	return this(cc.Call("copy", source, recursive))
}
func (cc *CubeCamera) DispatchEvent(event js.Value) {
	cc.Call("dispatchEvent", event)
}
func (cc *CubeCamera) GetObjectById(id int) core.Object3D {
	return core.Object3D(cc.Call("getObjectById", id))
}
func (cc *CubeCamera) GetObjectByName(name string) core.Object3D {
	return core.Object3D(cc.Call("getObjectByName", name))
}
func (cc *CubeCamera) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(cc.Call("getObjectByProperty", name, value))
}
func (cc *CubeCamera) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: cc.Call("getWorldDirection", target)}
}
func (cc *CubeCamera) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: cc.Call("getWorldPosition", target)}
}
func (cc *CubeCamera) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: cc.Call("getWorldQuaternion", target)}
}
func (cc *CubeCamera) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: cc.Call("getWorldScale", target)}
}
func (cc *CubeCamera) HasEventListener(typ string, listener js.Value) bool {
	return cc.Call("hasEventListener", typ, listener).Bool()
}
func (cc *CubeCamera) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: cc.Call("localToWorld", vector)}
}
func (cc *CubeCamera) LookAt(vector math.Vector3, y int, z int) {
	cc.Call("lookAt", vector, y, z)
}
func (cc *CubeCamera) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	cc.Call("raycast", raycaster, intersects)
}
func (cc *CubeCamera) Remove(object []core.Object3D) this {
	return this(cc.Call("remove", object))
}
func (cc *CubeCamera) RemoveEventListener(typ string, listener js.Value) {
	cc.Call("removeEventListener", typ, listener)
}
func (cc *CubeCamera) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(cc.Call("rotateOnAxis", axis, angle))
}
func (cc *CubeCamera) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(cc.Call("rotateOnWorldAxis", axis, angle))
}
func (cc *CubeCamera) RotateX(angle int) this {
	return this(cc.Call("rotateX", angle))
}
func (cc *CubeCamera) RotateY(angle int) this {
	return this(cc.Call("rotateY", angle))
}
func (cc *CubeCamera) RotateZ(angle int) this {
	return this(cc.Call("rotateZ", angle))
}
func (cc *CubeCamera) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	cc.Call("setRotationFromAxisAngle", axis, angle)
}
func (cc *CubeCamera) SetRotationFromEuler(euler *math.Euler) {
	cc.Call("setRotationFromEuler", euler)
}
func (cc *CubeCamera) SetRotationFromMatrix(m *math.Matrix4) {
	cc.Call("setRotationFromMatrix", m)
}
func (cc *CubeCamera) SetRotationFromQuaternion(q *math.Quaternion) {
	cc.Call("setRotationFromQuaternion", q)
}
func (cc *CubeCamera) ToJSON(meta js.Value) js.Value {
	return cc.Call("toJSON", meta)
}
func (cc *CubeCamera) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(cc.Call("translateOnAxis", axis, distance))
}
func (cc *CubeCamera) TranslateX(distance int) this {
	return this(cc.Call("translateX", distance))
}
func (cc *CubeCamera) TranslateY(distance int) this {
	return this(cc.Call("translateY", distance))
}
func (cc *CubeCamera) TranslateZ(distance int) this {
	return this(cc.Call("translateZ", distance))
}
func (cc *CubeCamera) Traverse(callback js.Value) {
	cc.Call("traverse", callback)
}
func (cc *CubeCamera) TraverseAncestors(callback js.Value) {
	cc.Call("traverseAncestors", callback)
}
func (cc *CubeCamera) TraverseVisible(callback js.Value) {
	cc.Call("traverseVisible", callback)
}
func (cc *CubeCamera) Update(renderer *renderers.WebGLRenderer, scene *scenes.Scene) {
	cc.Call("update", renderer, scene)
}
func (cc *CubeCamera) UpdateMatrix() {
	cc.Call("updateMatrix")
}
func (cc *CubeCamera) UpdateMatrixWorld(force bool) {
	cc.Call("updateMatrixWorld", force)
}
func (cc *CubeCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	cc.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (cc *CubeCamera) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: cc.Call("worldToLocal", vector)}
}

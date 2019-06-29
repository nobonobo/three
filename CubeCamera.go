// Code generated from "cameras/CubeCamera.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// CubeCamera extend: [Object3D]
type CubeCamera struct {
	js.Value
}

func NewCubeCamera(near float64, far float64, cubeResolution float64) *CubeCamera {
	return &CubeCamera{Value: get("CubeCamera").New(near, far, cubeResolution)}
}
func (cc *CubeCamera) JSValue() js.Value {
	return cc.Value
}
func (cc *CubeCamera) CastShadow() bool {
	return cc.Get("castShadow").Bool()
}
func (cc *CubeCamera) SetCastShadow(v bool) {
	cc.Set("castShadow", v)
}
func (cc *CubeCamera) Children() js.Value {
	return cc.Get("children")
}
func (cc *CubeCamera) SetChildren(v js.Value) {
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
func (cc *CubeCamera) IsObject3D() bool {
	return cc.Get("isObject3D").Bool()
}
func (cc *CubeCamera) SetIsObject3D(v bool) {
	cc.Set("isObject3D", v)
}
func (cc *CubeCamera) Layers() *Layers {
	return &Layers{Value: cc.Get("layers")}
}
func (cc *CubeCamera) SetLayers(v *Layers) {
	cc.Set("layers", v.JSValue())
}
func (cc *CubeCamera) Matrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrix")}
}
func (cc *CubeCamera) SetMatrix(v *Matrix4) {
	cc.Set("matrix", v.JSValue())
}
func (cc *CubeCamera) MatrixAutoUpdate() bool {
	return cc.Get("matrixAutoUpdate").Bool()
}
func (cc *CubeCamera) SetMatrixAutoUpdate(v bool) {
	cc.Set("matrixAutoUpdate", v)
}
func (cc *CubeCamera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorld")}
}
func (cc *CubeCamera) SetMatrixWorld(v *Matrix4) {
	cc.Set("matrixWorld", v.JSValue())
}
func (cc *CubeCamera) MatrixWorldNeedsUpdate() bool {
	return cc.Get("matrixWorldNeedsUpdate").Bool()
}
func (cc *CubeCamera) SetMatrixWorldNeedsUpdate(v bool) {
	cc.Set("matrixWorldNeedsUpdate", v)
}
func (cc *CubeCamera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("modelViewMatrix")}
}
func (cc *CubeCamera) SetModelViewMatrix(v *Matrix4) {
	cc.Set("modelViewMatrix", v.JSValue())
}
func (cc *CubeCamera) Name() string {
	return cc.Get("name").String()
}
func (cc *CubeCamera) SetName(v string) {
	cc.Set("name", v)
}
func (cc *CubeCamera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: cc.Get("normalMatrix")}
}
func (cc *CubeCamera) SetNormalMatrix(v *Matrix3) {
	cc.Set("normalMatrix", v.JSValue())
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
func (cc *CubeCamera) Parent() *Object3D {
	return &Object3D{Value: cc.Get("parent")}
}
func (cc *CubeCamera) SetParent(v *Object3D) {
	cc.Set("parent", v.JSValue())
}
func (cc *CubeCamera) Position() *Vector3 {
	return &Vector3{Value: cc.Get("position")}
}
func (cc *CubeCamera) SetPosition(v *Vector3) {
	cc.Set("position", v.JSValue())
}
func (cc *CubeCamera) Quaternion() *Quaternion {
	return &Quaternion{Value: cc.Get("quaternion")}
}
func (cc *CubeCamera) SetQuaternion(v *Quaternion) {
	cc.Set("quaternion", v.JSValue())
}
func (cc *CubeCamera) ReceiveShadow() bool {
	return cc.Get("receiveShadow").Bool()
}
func (cc *CubeCamera) SetReceiveShadow(v bool) {
	cc.Set("receiveShadow", v)
}
func (cc *CubeCamera) RenderOrder() float64 {
	return cc.Get("renderOrder").Float()
}
func (cc *CubeCamera) SetRenderOrder(v float64) {
	cc.Set("renderOrder", v)
}
func (cc *CubeCamera) RenderTarget() *WebGLRenderTargetCube {
	return &WebGLRenderTargetCube{Value: cc.Get("renderTarget")}
}
func (cc *CubeCamera) SetRenderTarget(v *WebGLRenderTargetCube) {
	cc.Set("renderTarget", v.JSValue())
}
func (cc *CubeCamera) Rotation() *Euler {
	return &Euler{Value: cc.Get("rotation")}
}
func (cc *CubeCamera) SetRotation(v *Euler) {
	cc.Set("rotation", v.JSValue())
}
func (cc *CubeCamera) Scale() *Vector3 {
	return &Vector3{Value: cc.Get("scale")}
}
func (cc *CubeCamera) SetScale(v *Vector3) {
	cc.Set("scale", v.JSValue())
}
func (cc *CubeCamera) Type() string {
	return cc.Get("type").String()
}
func (cc *CubeCamera) SetType(v string) {
	cc.Set("type", v)
}
func (cc *CubeCamera) Up() *Vector3 {
	return &Vector3{Value: cc.Get("up")}
}
func (cc *CubeCamera) SetUp(v *Vector3) {
	cc.Set("up", v.JSValue())
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
func (cc *CubeCamera) DefaultUp() *Vector3 {
	return &Vector3{Value: cc.Get("DefaultUp")}
}
func (cc *CubeCamera) SetDefaultUp(v *Vector3) {
	cc.Set("DefaultUp", v.JSValue())
}
func (cc *CubeCamera) Add(object js.Value) *CubeCamera {
	return &CubeCamera{Value: cc.Call("add", object)}
}
func (cc *CubeCamera) AddEventListener(typ string, listener js.Value) {
	cc.Call("addEventListener", typ, listener)
}
func (cc *CubeCamera) ApplyMatrix(matrix *Matrix4) {
	cc.Call("applyMatrix", matrix.JSValue())
}
func (cc *CubeCamera) ApplyQuaternion(quaternion *Quaternion) *CubeCamera {
	return &CubeCamera{Value: cc.Call("applyQuaternion", quaternion)}
}
func (cc *CubeCamera) Clone(recursive bool) *CubeCamera {
	return &CubeCamera{Value: cc.Call("clone", recursive)}
}
func (cc *CubeCamera) Copy(source *Object3D, recursive bool) *CubeCamera {
	return &CubeCamera{Value: cc.Call("copy", source, recursive)}
}
func (cc *CubeCamera) DispatchEvent(event js.Value) {
	cc.Call("dispatchEvent", event)
}
func (cc *CubeCamera) GetObjectById(id int) *Object3D {
	return &Object3D{Value: cc.Call("getObjectById", id)}
}
func (cc *CubeCamera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByName", name)}
}
func (cc *CubeCamera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByProperty", name, value)}
}
func (cc *CubeCamera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldDirection", target)}
}
func (cc *CubeCamera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldPosition", target)}
}
func (cc *CubeCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: cc.Call("getWorldQuaternion", target)}
}
func (cc *CubeCamera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldScale", target)}
}
func (cc *CubeCamera) HasEventListener(typ string, listener js.Value) bool {
	return cc.Call("hasEventListener", typ, listener).Bool()
}
func (cc *CubeCamera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("localToWorld", vector)}
}
func (cc *CubeCamera) LookAt(vector *Vector3, y float64, z float64) {
	cc.Call("lookAt", vector, y, z)
}
func (cc *CubeCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
	cc.Call("raycast", raycaster.JSValue(), intersects)
}
func (cc *CubeCamera) Remove(object js.Value) *CubeCamera {
	return &CubeCamera{Value: cc.Call("remove", object)}
}
func (cc *CubeCamera) RemoveEventListener(typ string, listener js.Value) {
	cc.Call("removeEventListener", typ, listener)
}
func (cc *CubeCamera) RotateOnAxis(axis *Vector3, angle float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("rotateOnAxis", axis, angle)}
}
func (cc *CubeCamera) RotateOnWorldAxis(axis *Vector3, angle float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("rotateOnWorldAxis", axis, angle)}
}
func (cc *CubeCamera) RotateX(angle float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("rotateX", angle)}
}
func (cc *CubeCamera) RotateY(angle float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("rotateY", angle)}
}
func (cc *CubeCamera) RotateZ(angle float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("rotateZ", angle)}
}
func (cc *CubeCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	cc.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (cc *CubeCamera) SetRotationFromEuler(euler *Euler) {
	cc.Call("setRotationFromEuler", euler.JSValue())
}
func (cc *CubeCamera) SetRotationFromMatrix(m *Matrix4) {
	cc.Call("setRotationFromMatrix", m.JSValue())
}
func (cc *CubeCamera) SetRotationFromQuaternion(q *Quaternion) {
	cc.Call("setRotationFromQuaternion", q.JSValue())
}
func (cc *CubeCamera) ToJSON(meta js.Value) js.Value {
	return cc.Call("toJSON", meta)
}
func (cc *CubeCamera) TranslateOnAxis(axis *Vector3, distance float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("translateOnAxis", axis, distance)}
}
func (cc *CubeCamera) TranslateX(distance float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("translateX", distance)}
}
func (cc *CubeCamera) TranslateY(distance float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("translateY", distance)}
}
func (cc *CubeCamera) TranslateZ(distance float64) *CubeCamera {
	return &CubeCamera{Value: cc.Call("translateZ", distance)}
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
func (cc *CubeCamera) Update(renderer *WebGLRenderer, scene Scene) {
	cc.Call("update", renderer.JSValue(), scene.JSValue())
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
func (cc *CubeCamera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("worldToLocal", vector)}
}

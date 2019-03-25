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

type SkeletonHelper struct {
	js.Value
}

func NewSkeletonHelper(bone *core.Object3D) *SkeletonHelper {
	return &SkeletonHelper{Value: get("SkeletonHelper").New(bone)}
}
func (sh *SkeletonHelper) Bones() []objects.Bone {
	return []objects.Bone(sh.Get("bones"))
}
func (sh *SkeletonHelper) SetBones(v []objects.Bone) {
	sh.Set("bones", v)
}
func (sh *SkeletonHelper) CastShadow() bool {
	return sh.Get("castShadow").Bool()
}
func (sh *SkeletonHelper) SetCastShadow(v bool) {
	sh.Set("castShadow", v)
}
func (sh *SkeletonHelper) Children() []core.Object3D {
	return []core.Object3D(sh.Get("children"))
}
func (sh *SkeletonHelper) SetChildren(v []core.Object3D) {
	sh.Set("children", v)
}
func (sh *SkeletonHelper) FrustumCulled() bool {
	return sh.Get("frustumCulled").Bool()
}
func (sh *SkeletonHelper) SetFrustumCulled(v bool) {
	sh.Set("frustumCulled", v)
}
func (sh *SkeletonHelper) Geometry() core.Geometry {
	return core.Geometry(sh.Get("geometry"))
}
func (sh *SkeletonHelper) SetGeometry(v core.Geometry) {
	sh.Set("geometry", v)
}
func (sh *SkeletonHelper) Id() int {
	return sh.Get("id").Int()
}
func (sh *SkeletonHelper) SetId(v int) {
	sh.Set("id", v)
}
func (sh *SkeletonHelper) IsLine() true {
	return true(sh.Get("isLine"))
}
func (sh *SkeletonHelper) SetIsLine(v true) {
	sh.Set("isLine", v)
}
func (sh *SkeletonHelper) IsLineSegments() true {
	return true(sh.Get("isLineSegments"))
}
func (sh *SkeletonHelper) SetIsLineSegments(v true) {
	sh.Set("isLineSegments", v)
}
func (sh *SkeletonHelper) IsObject3D() true {
	return true(sh.Get("isObject3D"))
}
func (sh *SkeletonHelper) SetIsObject3D(v true) {
	sh.Set("isObject3D", v)
}
func (sh *SkeletonHelper) Layers() *core.Layers {
	return &core.Layers{Value: sh.Get("layers")}
}
func (sh *SkeletonHelper) SetLayers(v *core.Layers) {
	sh.Set("layers", v)
}
func (sh *SkeletonHelper) Material() materials.Material {
	return materials.Material(sh.Get("material"))
}
func (sh *SkeletonHelper) SetMaterial(v materials.Material) {
	sh.Set("material", v)
}
func (sh *SkeletonHelper) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: sh.Get("matrix")}
}
func (sh *SkeletonHelper) SetMatrix(v *math.Matrix4) {
	sh.Set("matrix", v)
}
func (sh *SkeletonHelper) MatrixAutoUpdate() bool {
	return sh.Get("matrixAutoUpdate").Bool()
}
func (sh *SkeletonHelper) SetMatrixAutoUpdate(v bool) {
	sh.Set("matrixAutoUpdate", v)
}
func (sh *SkeletonHelper) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: sh.Get("matrixWorld")}
}
func (sh *SkeletonHelper) SetMatrixWorld(v *math.Matrix4) {
	sh.Set("matrixWorld", v)
}
func (sh *SkeletonHelper) MatrixWorldNeedsUpdate() bool {
	return sh.Get("matrixWorldNeedsUpdate").Bool()
}
func (sh *SkeletonHelper) SetMatrixWorldNeedsUpdate(v bool) {
	sh.Set("matrixWorldNeedsUpdate", v)
}
func (sh *SkeletonHelper) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: sh.Get("modelViewMatrix")}
}
func (sh *SkeletonHelper) SetModelViewMatrix(v *math.Matrix4) {
	sh.Set("modelViewMatrix", v)
}
func (sh *SkeletonHelper) Name() string {
	return sh.Get("name").String()
}
func (sh *SkeletonHelper) SetName(v string) {
	sh.Set("name", v)
}
func (sh *SkeletonHelper) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: sh.Get("normalMatrix")}
}
func (sh *SkeletonHelper) SetNormalMatrix(v *math.Matrix3) {
	sh.Set("normalMatrix", v)
}
func (sh *SkeletonHelper) OnAfterRender() js.Value {
	return sh.Get("onAfterRender")
}
func (sh *SkeletonHelper) SetOnAfterRender(v js.Value) {
	sh.Set("onAfterRender", v)
}
func (sh *SkeletonHelper) OnBeforeRender() js.Value {
	return sh.Get("onBeforeRender")
}
func (sh *SkeletonHelper) SetOnBeforeRender(v js.Value) {
	sh.Set("onBeforeRender", v)
}
func (sh *SkeletonHelper) Parent() core.Object3D {
	return core.Object3D(sh.Get("parent"))
}
func (sh *SkeletonHelper) SetParent(v core.Object3D) {
	sh.Set("parent", v)
}
func (sh *SkeletonHelper) Position() *math.Vector3 {
	return &math.Vector3{Value: sh.Get("position")}
}
func (sh *SkeletonHelper) SetPosition(v *math.Vector3) {
	sh.Set("position", v)
}
func (sh *SkeletonHelper) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: sh.Get("quaternion")}
}
func (sh *SkeletonHelper) SetQuaternion(v *math.Quaternion) {
	sh.Set("quaternion", v)
}
func (sh *SkeletonHelper) ReceiveShadow() bool {
	return sh.Get("receiveShadow").Bool()
}
func (sh *SkeletonHelper) SetReceiveShadow(v bool) {
	sh.Set("receiveShadow", v)
}
func (sh *SkeletonHelper) RenderOrder() int {
	return sh.Get("renderOrder").Int()
}
func (sh *SkeletonHelper) SetRenderOrder(v int) {
	sh.Set("renderOrder", v)
}
func (sh *SkeletonHelper) Root() *core.Object3D {
	return &core.Object3D{Value: sh.Get("root")}
}
func (sh *SkeletonHelper) SetRoot(v *core.Object3D) {
	sh.Set("root", v)
}
func (sh *SkeletonHelper) Rotation() *math.Euler {
	return &math.Euler{Value: sh.Get("rotation")}
}
func (sh *SkeletonHelper) SetRotation(v *math.Euler) {
	sh.Set("rotation", v)
}
func (sh *SkeletonHelper) Scale() *math.Vector3 {
	return &math.Vector3{Value: sh.Get("scale")}
}
func (sh *SkeletonHelper) SetScale(v *math.Vector3) {
	sh.Set("scale", v)
}
func (sh *SkeletonHelper) Type() string {
	return sh.Get("type").String()
}
func (sh *SkeletonHelper) SetType(v string) {
	sh.Set("type", v)
}
func (sh *SkeletonHelper) Up() *math.Vector3 {
	return &math.Vector3{Value: sh.Get("up")}
}
func (sh *SkeletonHelper) SetUp(v *math.Vector3) {
	sh.Set("up", v)
}
func (sh *SkeletonHelper) UserData() js.Value {
	return sh.Get("userData")
}
func (sh *SkeletonHelper) SetUserData(v js.Value) {
	sh.Set("userData", v)
}
func (sh *SkeletonHelper) Uuid() string {
	return sh.Get("uuid").String()
}
func (sh *SkeletonHelper) SetUuid(v string) {
	sh.Set("uuid", v)
}
func (sh *SkeletonHelper) Visible() bool {
	return sh.Get("visible").Bool()
}
func (sh *SkeletonHelper) SetVisible(v bool) {
	sh.Set("visible", v)
}
func (sh *SkeletonHelper) DefaultMatrixAutoUpdate() bool {
	return sh.Get("DefaultMatrixAutoUpdate").Bool()
}
func (sh *SkeletonHelper) SetDefaultMatrixAutoUpdate(v bool) {
	sh.Set("DefaultMatrixAutoUpdate", v)
}
func (sh *SkeletonHelper) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: sh.Get("DefaultUp")}
}
func (sh *SkeletonHelper) SetDefaultUp(v *math.Vector3) {
	sh.Set("DefaultUp", v)
}
func (sh *SkeletonHelper) Add(object []core.Object3D) this {
	return this(sh.Call("add", object))
}
func (sh *SkeletonHelper) AddEventListener(typ string, listener js.Value) {
	sh.Call("addEventListener", typ, listener)
}
func (sh *SkeletonHelper) ApplyMatrix(matrix *math.Matrix4) {
	sh.Call("applyMatrix", matrix)
}
func (sh *SkeletonHelper) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(sh.Call("applyQuaternion", quaternion))
}
func (sh *SkeletonHelper) Clone(recursive bool) this {
	return this(sh.Call("clone", recursive))
}
func (sh *SkeletonHelper) ComputeLineDistances() this {
	return this(sh.Call("computeLineDistances"))
}
func (sh *SkeletonHelper) Copy(source *core.Object3D, recursive bool) this {
	return this(sh.Call("copy", source, recursive))
}
func (sh *SkeletonHelper) DispatchEvent(event js.Value) {
	sh.Call("dispatchEvent", event)
}
func (sh *SkeletonHelper) GetBoneList(object *core.Object3D) []objects.Bone {
	return []objects.Bone(sh.Call("getBoneList", object))
}
func (sh *SkeletonHelper) GetObjectById(id int) core.Object3D {
	return core.Object3D(sh.Call("getObjectById", id))
}
func (sh *SkeletonHelper) GetObjectByName(name string) core.Object3D {
	return core.Object3D(sh.Call("getObjectByName", name))
}
func (sh *SkeletonHelper) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(sh.Call("getObjectByProperty", name, value))
}
func (sh *SkeletonHelper) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sh.Call("getWorldDirection", target)}
}
func (sh *SkeletonHelper) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sh.Call("getWorldPosition", target)}
}
func (sh *SkeletonHelper) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: sh.Call("getWorldQuaternion", target)}
}
func (sh *SkeletonHelper) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sh.Call("getWorldScale", target)}
}
func (sh *SkeletonHelper) HasEventListener(typ string, listener js.Value) bool {
	return sh.Call("hasEventListener", typ, listener).Bool()
}
func (sh *SkeletonHelper) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sh.Call("localToWorld", vector)}
}
func (sh *SkeletonHelper) LookAt(vector math.Vector3, y int, z int) {
	sh.Call("lookAt", vector, y, z)
}
func (sh *SkeletonHelper) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	sh.Call("raycast", raycaster, intersects)
}
func (sh *SkeletonHelper) Remove(object []core.Object3D) this {
	return this(sh.Call("remove", object))
}
func (sh *SkeletonHelper) RemoveEventListener(typ string, listener js.Value) {
	sh.Call("removeEventListener", typ, listener)
}
func (sh *SkeletonHelper) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(sh.Call("rotateOnAxis", axis, angle))
}
func (sh *SkeletonHelper) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(sh.Call("rotateOnWorldAxis", axis, angle))
}
func (sh *SkeletonHelper) RotateX(angle int) this {
	return this(sh.Call("rotateX", angle))
}
func (sh *SkeletonHelper) RotateY(angle int) this {
	return this(sh.Call("rotateY", angle))
}
func (sh *SkeletonHelper) RotateZ(angle int) this {
	return this(sh.Call("rotateZ", angle))
}
func (sh *SkeletonHelper) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	sh.Call("setRotationFromAxisAngle", axis, angle)
}
func (sh *SkeletonHelper) SetRotationFromEuler(euler *math.Euler) {
	sh.Call("setRotationFromEuler", euler)
}
func (sh *SkeletonHelper) SetRotationFromMatrix(m *math.Matrix4) {
	sh.Call("setRotationFromMatrix", m)
}
func (sh *SkeletonHelper) SetRotationFromQuaternion(q *math.Quaternion) {
	sh.Call("setRotationFromQuaternion", q)
}
func (sh *SkeletonHelper) ToJSON(meta js.Value) js.Value {
	return sh.Call("toJSON", meta)
}
func (sh *SkeletonHelper) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(sh.Call("translateOnAxis", axis, distance))
}
func (sh *SkeletonHelper) TranslateX(distance int) this {
	return this(sh.Call("translateX", distance))
}
func (sh *SkeletonHelper) TranslateY(distance int) this {
	return this(sh.Call("translateY", distance))
}
func (sh *SkeletonHelper) TranslateZ(distance int) this {
	return this(sh.Call("translateZ", distance))
}
func (sh *SkeletonHelper) Traverse(callback js.Value) {
	sh.Call("traverse", callback)
}
func (sh *SkeletonHelper) TraverseAncestors(callback js.Value) {
	sh.Call("traverseAncestors", callback)
}
func (sh *SkeletonHelper) TraverseVisible(callback js.Value) {
	sh.Call("traverseVisible", callback)
}
func (sh *SkeletonHelper) Update() {
	sh.Call("update")
}
func (sh *SkeletonHelper) UpdateMatrix() {
	sh.Call("updateMatrix")
}
func (sh *SkeletonHelper) UpdateMatrixWorld(force bool) {
	sh.Call("updateMatrixWorld", force)
}
func (sh *SkeletonHelper) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	sh.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (sh *SkeletonHelper) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: sh.Call("worldToLocal", vector)}
}

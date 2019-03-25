package geometries

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/math"
)

type WireframeGeometry struct {
	js.Value
}

func NewWireframeGeometry(geometry core.Geometry) *WireframeGeometry {
	return &WireframeGeometry{Value: get("WireframeGeometry").New(geometry)}
}
func (wg *WireframeGeometry) Attributes() js.Value {
	return wg.Get("attributes")
}
func (wg *WireframeGeometry) SetAttributes(v js.Value) {
	wg.Set("attributes", v)
}
func (wg *WireframeGeometry) BoundingBox() *math.Box3 {
	return &math.Box3{Value: wg.Get("boundingBox")}
}
func (wg *WireframeGeometry) SetBoundingBox(v *math.Box3) {
	wg.Set("boundingBox", v)
}
func (wg *WireframeGeometry) BoundingSphere() *math.Sphere {
	return &math.Sphere{Value: wg.Get("boundingSphere")}
}
func (wg *WireframeGeometry) SetBoundingSphere(v *math.Sphere) {
	wg.Set("boundingSphere", v)
}
func (wg *WireframeGeometry) DrawRange() js.Value {
	return wg.Get("drawRange")
}
func (wg *WireframeGeometry) SetDrawRange(v js.Value) {
	wg.Set("drawRange", v)
}
func (wg *WireframeGeometry) Drawcalls() js.Value {
	return wg.Get("drawcalls")
}
func (wg *WireframeGeometry) SetDrawcalls(v js.Value) {
	wg.Set("drawcalls", v)
}
func (wg *WireframeGeometry) Groups() []js.Value {
	return []js.Value(wg.Get("groups"))
}
func (wg *WireframeGeometry) SetGroups(v []js.Value) {
	wg.Set("groups", v)
}
func (wg *WireframeGeometry) Id() int {
	return wg.Get("id").Int()
}
func (wg *WireframeGeometry) SetId(v int) {
	wg.Set("id", v)
}
func (wg *WireframeGeometry) Index() *core.BufferAttribute {
	return &core.BufferAttribute{Value: wg.Get("index")}
}
func (wg *WireframeGeometry) SetIndex(v *core.BufferAttribute) {
	wg.Set("index", v)
}
func (wg *WireframeGeometry) MorphAttributes() js.Value {
	return wg.Get("morphAttributes")
}
func (wg *WireframeGeometry) SetMorphAttributes(v js.Value) {
	wg.Set("morphAttributes", v)
}
func (wg *WireframeGeometry) Name() string {
	return wg.Get("name").String()
}
func (wg *WireframeGeometry) SetName(v string) {
	wg.Set("name", v)
}
func (wg *WireframeGeometry) Offsets() js.Value {
	return wg.Get("offsets")
}
func (wg *WireframeGeometry) SetOffsets(v js.Value) {
	wg.Set("offsets", v)
}
func (wg *WireframeGeometry) Type() string {
	return wg.Get("type").String()
}
func (wg *WireframeGeometry) SetType(v string) {
	wg.Set("type", v)
}
func (wg *WireframeGeometry) UserData() js.Value {
	return wg.Get("userData")
}
func (wg *WireframeGeometry) SetUserData(v js.Value) {
	wg.Set("userData", v)
}
func (wg *WireframeGeometry) Uuid() string {
	return wg.Get("uuid").String()
}
func (wg *WireframeGeometry) SetUuid(v string) {
	wg.Set("uuid", v)
}
func (wg *WireframeGeometry) MaxIndex() int {
	return wg.Get("MaxIndex").Int()
}
func (wg *WireframeGeometry) SetMaxIndex(v int) {
	wg.Set("MaxIndex", v)
}
func (wg *WireframeGeometry) AddAttribute(name string, attribute core.BufferAttribute) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("addAttribute", name, attribute)}
}
func (wg *WireframeGeometry) AddAttribute2(name js.Value, array js.Value, itemSize js.Value) js.Value {
	return wg.Call("addAttribute", name, array, itemSize)
}
func (wg *WireframeGeometry) AddDrawCall(start js.Value, count js.Value, indexOffset js.Value) {
	wg.Call("addDrawCall", start, count, indexOffset)
}
func (wg *WireframeGeometry) AddEventListener(typ string, listener js.Value) {
	wg.Call("addEventListener", typ, listener)
}
func (wg *WireframeGeometry) AddGroup(start int, count int, materialIndex int) {
	wg.Call("addGroup", start, count, materialIndex)
}
func (wg *WireframeGeometry) AddIndex(index js.Value) {
	wg.Call("addIndex", index)
}
func (wg *WireframeGeometry) ApplyMatrix(matrix *math.Matrix4) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("applyMatrix", matrix)}
}
func (wg *WireframeGeometry) Center() *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("center")}
}
func (wg *WireframeGeometry) ClearDrawCalls() {
	wg.Call("clearDrawCalls")
}
func (wg *WireframeGeometry) ClearGroups() {
	wg.Call("clearGroups")
}
func (wg *WireframeGeometry) Clone() this {
	return this(wg.Call("clone"))
}
func (wg *WireframeGeometry) ComputeBoundingBox() {
	wg.Call("computeBoundingBox")
}
func (wg *WireframeGeometry) ComputeBoundingSphere() {
	wg.Call("computeBoundingSphere")
}
func (wg *WireframeGeometry) ComputeVertexNormals() {
	wg.Call("computeVertexNormals")
}
func (wg *WireframeGeometry) Copy(source *core.BufferGeometry) this {
	return this(wg.Call("copy", source))
}
func (wg *WireframeGeometry) DispatchEvent(event js.Value) {
	wg.Call("dispatchEvent", event)
}
func (wg *WireframeGeometry) Dispose() {
	wg.Call("dispose")
}
func (wg *WireframeGeometry) FromDirectGeometry(geometry *core.DirectGeometry) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("fromDirectGeometry", geometry)}
}
func (wg *WireframeGeometry) FromGeometry(geometry *core.Geometry, settings js.Value) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("fromGeometry", geometry, settings)}
}
func (wg *WireframeGeometry) GetAttribute(name string) core.BufferAttribute {
	return core.BufferAttribute(wg.Call("getAttribute", name))
}
func (wg *WireframeGeometry) GetIndex() *core.BufferAttribute {
	return &core.BufferAttribute{Value: wg.Call("getIndex")}
}
func (wg *WireframeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return wg.Call("hasEventListener", typ, listener).Bool()
}
func (wg *WireframeGeometry) LookAt(v *math.Vector3) {
	wg.Call("lookAt", v)
}
func (wg *WireframeGeometry) Merge(geometry *core.BufferGeometry, offset int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("merge", geometry, offset)}
}
func (wg *WireframeGeometry) NormalizeNormals() {
	wg.Call("normalizeNormals")
}
func (wg *WireframeGeometry) RemoveAttribute(name string) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("removeAttribute", name)}
}
func (wg *WireframeGeometry) RemoveEventListener(typ string, listener js.Value) {
	wg.Call("removeEventListener", typ, listener)
}
func (wg *WireframeGeometry) RotateX(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("rotateX", angle)}
}
func (wg *WireframeGeometry) RotateY(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("rotateY", angle)}
}
func (wg *WireframeGeometry) RotateZ(angle int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("rotateZ", angle)}
}
func (wg *WireframeGeometry) Scale(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("scale", x, y, z)}
}
func (wg *WireframeGeometry) SetDrawRange(start int, count int) {
	wg.Call("setDrawRange", start, count)
}
func (wg *WireframeGeometry) SetFromObject(object *core.Object3D) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("setFromObject", object)}
}
func (wg *WireframeGeometry) SetFromPoints(points []math.Vector3) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("setFromPoints", points)}
}
func (wg *WireframeGeometry) SetIndex(index core.BufferAttribute) {
	wg.Call("setIndex", index)
}
func (wg *WireframeGeometry) ToJSON() js.Value {
	return wg.Call("toJSON")
}
func (wg *WireframeGeometry) ToNonIndexed() *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("toNonIndexed")}
}
func (wg *WireframeGeometry) Translate(x int, y int, z int) *core.BufferGeometry {
	return &core.BufferGeometry{Value: wg.Call("translate", x, y, z)}
}
func (wg *WireframeGeometry) UpdateFromObject(object *core.Object3D) {
	wg.Call("updateFromObject", object)
}

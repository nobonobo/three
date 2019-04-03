// Code generated from "geometries/WireframeGeometry.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type WireframeGeometry struct {
	js.Value
}

func NewWireframeGeometry(geometry *Geometry) *WireframeGeometry {
	return &WireframeGeometry{Value: get("WireframeGeometry").New(geometry)}
}
func (wg *WireframeGeometry) Attributes() js.Value {
	return wg.Get("attributes")
}
func (wg *WireframeGeometry) SetAttributes(v js.Value) {
	wg.Set("attributes", v)
}
func (wg *WireframeGeometry) BoundingBox() *Box3 {
	return &Box3{Value: wg.Get("boundingBox")}
}
func (wg *WireframeGeometry) SetBoundingBox(v *Box3) {
	wg.Set("boundingBox", v)
}
func (wg *WireframeGeometry) BoundingSphere() *Sphere {
	return &Sphere{Value: wg.Get("boundingSphere")}
}
func (wg *WireframeGeometry) SetBoundingSphere(v *Sphere) {
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
func (wg *WireframeGeometry) Groups() js.Value {
	return wg.Get("groups")
}
func (wg *WireframeGeometry) SetGroups(v js.Value) {
	wg.Set("groups", v)
}
func (wg *WireframeGeometry) Id() int {
	return wg.Get("id").Int()
}
func (wg *WireframeGeometry) SetId(v int) {
	wg.Set("id", v)
}
func (wg *WireframeGeometry) Index() *BufferAttribute {
	return &BufferAttribute{Value: wg.Get("index")}
}
func (wg *WireframeGeometry) SetIndex(v *BufferAttribute) {
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
func (wg *WireframeGeometry) AddAttribute(name string, attribute *BufferAttribute) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("addAttribute", name, attribute)}
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
func (wg *WireframeGeometry) ApplyMatrix(matrix *Matrix4) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("applyMatrix", matrix)}
}
func (wg *WireframeGeometry) Center() *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("center")}
}
func (wg *WireframeGeometry) ClearDrawCalls() {
	wg.Call("clearDrawCalls")
}
func (wg *WireframeGeometry) ClearGroups() {
	wg.Call("clearGroups")
}
func (wg *WireframeGeometry) Clone() *WireframeGeometry {
	return &WireframeGeometry{Value: wg.Call("clone")}
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
func (wg *WireframeGeometry) Copy(source *BufferGeometry) *WireframeGeometry {
	return &WireframeGeometry{Value: wg.Call("copy", source)}
}
func (wg *WireframeGeometry) DispatchEvent(event js.Value) {
	wg.Call("dispatchEvent", event)
}
func (wg *WireframeGeometry) Dispose() {
	wg.Call("dispose")
}
func (wg *WireframeGeometry) FromDirectGeometry(geometry *DirectGeometry) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("fromDirectGeometry", geometry)}
}
func (wg *WireframeGeometry) FromGeometry(geometry *Geometry, settings js.Value) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("fromGeometry", geometry, settings)}
}
func (wg *WireframeGeometry) GetAttribute(name string) *BufferAttribute {
	return &BufferAttribute{Value: wg.Call("getAttribute", name)}
}
func (wg *WireframeGeometry) GetIndex() *BufferAttribute {
	return &BufferAttribute{Value: wg.Call("getIndex")}
}
func (wg *WireframeGeometry) HasEventListener(typ string, listener js.Value) bool {
	return wg.Call("hasEventListener", typ, listener).Bool()
}
func (wg *WireframeGeometry) LookAt(v *Vector3) {
	wg.Call("lookAt", v)
}
func (wg *WireframeGeometry) Merge(geometry *BufferGeometry, offset int) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("merge", geometry, offset)}
}
func (wg *WireframeGeometry) NormalizeNormals() {
	wg.Call("normalizeNormals")
}
func (wg *WireframeGeometry) RemoveAttribute(name string) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("removeAttribute", name)}
}
func (wg *WireframeGeometry) RemoveEventListener(typ string, listener js.Value) {
	wg.Call("removeEventListener", typ, listener)
}
func (wg *WireframeGeometry) RotateX(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("rotateX", angle)}
}
func (wg *WireframeGeometry) RotateY(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("rotateY", angle)}
}
func (wg *WireframeGeometry) RotateZ(angle float64) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("rotateZ", angle)}
}
func (wg *WireframeGeometry) Scale(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("scale", x, y, z)}
}
func (wg *WireframeGeometry) SetDrawRange2(start int, count int) {
	wg.Call("setDrawRange", start, count)
}
func (wg *WireframeGeometry) SetFromObject(object *Object3D) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("setFromObject", object)}
}
func (wg *WireframeGeometry) SetFromPoints(points js.Value) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("setFromPoints", points)}
}
func (wg *WireframeGeometry) SetIndex2(index *BufferAttribute) {
	wg.Call("setIndex", index)
}
func (wg *WireframeGeometry) ToJSON() js.Value {
	return wg.Call("toJSON")
}
func (wg *WireframeGeometry) ToNonIndexed() *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("toNonIndexed")}
}
func (wg *WireframeGeometry) Translate(x float64, y float64, z float64) *BufferGeometry {
	return &BufferGeometry{Value: wg.Call("translate", x, y, z)}
}
func (wg *WireframeGeometry) UpdateFromObject(object *Object3D) {
	wg.Call("updateFromObject", object)
}

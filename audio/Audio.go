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

type Audio struct {
	js.Value
}

func NewAudio(listener *AudioListener) *Audio {
	return &Audio{Value: get("Audio").New(listener)}
}
func (a *Audio) Autoplay() bool {
	return a.Get("autoplay").Bool()
}
func (a *Audio) SetAutoplay(v bool) {
	a.Set("autoplay", v)
}
func (a *Audio) Buffer() null {
	return null(a.Get("buffer"))
}
func (a *Audio) SetBuffer(v null) {
	a.Set("buffer", v)
}
func (a *Audio) CastShadow() bool {
	return a.Get("castShadow").Bool()
}
func (a *Audio) SetCastShadow(v bool) {
	a.Set("castShadow", v)
}
func (a *Audio) Children() []core.Object3D {
	return []core.Object3D(a.Get("children"))
}
func (a *Audio) SetChildren(v []core.Object3D) {
	a.Set("children", v)
}
func (a *Audio) Context() *AudioContext {
	return &AudioContext{Value: a.Get("context")}
}
func (a *Audio) SetContext(v *AudioContext) {
	a.Set("context", v)
}
func (a *Audio) Detune() int {
	return a.Get("detune").Int()
}
func (a *Audio) SetDetune(v int) {
	a.Set("detune", v)
}
func (a *Audio) Filters() []js.Value {
	return []js.Value(a.Get("filters"))
}
func (a *Audio) SetFilters(v []js.Value) {
	a.Set("filters", v)
}
func (a *Audio) FrustumCulled() bool {
	return a.Get("frustumCulled").Bool()
}
func (a *Audio) SetFrustumCulled(v bool) {
	a.Set("frustumCulled", v)
}
func (a *Audio) Gain() *GainNode {
	return &GainNode{Value: a.Get("gain")}
}
func (a *Audio) SetGain(v *GainNode) {
	a.Set("gain", v)
}
func (a *Audio) HasPlaybackControl() bool {
	return a.Get("hasPlaybackControl").Bool()
}
func (a *Audio) SetHasPlaybackControl(v bool) {
	a.Set("hasPlaybackControl", v)
}
func (a *Audio) Id() int {
	return a.Get("id").Int()
}
func (a *Audio) SetId(v int) {
	a.Set("id", v)
}
func (a *Audio) IsObject3D() true {
	return true(a.Get("isObject3D"))
}
func (a *Audio) SetIsObject3D(v true) {
	a.Set("isObject3D", v)
}
func (a *Audio) IsPlaying() bool {
	return a.Get("isPlaying").Bool()
}
func (a *Audio) SetIsPlaying(v bool) {
	a.Set("isPlaying", v)
}
func (a *Audio) Layers() *core.Layers {
	return &core.Layers{Value: a.Get("layers")}
}
func (a *Audio) SetLayers(v *core.Layers) {
	a.Set("layers", v)
}
func (a *Audio) Loop() bool {
	return a.Get("loop").Bool()
}
func (a *Audio) SetLoop(v bool) {
	a.Set("loop", v)
}
func (a *Audio) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: a.Get("matrix")}
}
func (a *Audio) SetMatrix(v *math.Matrix4) {
	a.Set("matrix", v)
}
func (a *Audio) MatrixAutoUpdate() bool {
	return a.Get("matrixAutoUpdate").Bool()
}
func (a *Audio) SetMatrixAutoUpdate(v bool) {
	a.Set("matrixAutoUpdate", v)
}
func (a *Audio) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: a.Get("matrixWorld")}
}
func (a *Audio) SetMatrixWorld(v *math.Matrix4) {
	a.Set("matrixWorld", v)
}
func (a *Audio) MatrixWorldNeedsUpdate() bool {
	return a.Get("matrixWorldNeedsUpdate").Bool()
}
func (a *Audio) SetMatrixWorldNeedsUpdate(v bool) {
	a.Set("matrixWorldNeedsUpdate", v)
}
func (a *Audio) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: a.Get("modelViewMatrix")}
}
func (a *Audio) SetModelViewMatrix(v *math.Matrix4) {
	a.Set("modelViewMatrix", v)
}
func (a *Audio) Name() string {
	return a.Get("name").String()
}
func (a *Audio) SetName(v string) {
	a.Set("name", v)
}
func (a *Audio) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: a.Get("normalMatrix")}
}
func (a *Audio) SetNormalMatrix(v *math.Matrix3) {
	a.Set("normalMatrix", v)
}
func (a *Audio) Offset() int {
	return a.Get("offset").Int()
}
func (a *Audio) SetOffset(v int) {
	a.Set("offset", v)
}
func (a *Audio) OnAfterRender() js.Value {
	return a.Get("onAfterRender")
}
func (a *Audio) SetOnAfterRender(v js.Value) {
	a.Set("onAfterRender", v)
}
func (a *Audio) OnBeforeRender() js.Value {
	return a.Get("onBeforeRender")
}
func (a *Audio) SetOnBeforeRender(v js.Value) {
	a.Set("onBeforeRender", v)
}
func (a *Audio) Parent() core.Object3D {
	return core.Object3D(a.Get("parent"))
}
func (a *Audio) SetParent(v core.Object3D) {
	a.Set("parent", v)
}
func (a *Audio) PlaybackRate() int {
	return a.Get("playbackRate").Int()
}
func (a *Audio) SetPlaybackRate(v int) {
	a.Set("playbackRate", v)
}
func (a *Audio) Position() *math.Vector3 {
	return &math.Vector3{Value: a.Get("position")}
}
func (a *Audio) SetPosition(v *math.Vector3) {
	a.Set("position", v)
}
func (a *Audio) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: a.Get("quaternion")}
}
func (a *Audio) SetQuaternion(v *math.Quaternion) {
	a.Set("quaternion", v)
}
func (a *Audio) ReceiveShadow() bool {
	return a.Get("receiveShadow").Bool()
}
func (a *Audio) SetReceiveShadow(v bool) {
	a.Set("receiveShadow", v)
}
func (a *Audio) RenderOrder() int {
	return a.Get("renderOrder").Int()
}
func (a *Audio) SetRenderOrder(v int) {
	a.Set("renderOrder", v)
}
func (a *Audio) Rotation() *math.Euler {
	return &math.Euler{Value: a.Get("rotation")}
}
func (a *Audio) SetRotation(v *math.Euler) {
	a.Set("rotation", v)
}
func (a *Audio) Scale() *math.Vector3 {
	return &math.Vector3{Value: a.Get("scale")}
}
func (a *Audio) SetScale(v *math.Vector3) {
	a.Set("scale", v)
}
func (a *Audio) Source() *AudioBufferSourceNode {
	return &AudioBufferSourceNode{Value: a.Get("source")}
}
func (a *Audio) SetSource(v *AudioBufferSourceNode) {
	a.Set("source", v)
}
func (a *Audio) SourceType() string {
	return a.Get("sourceType").String()
}
func (a *Audio) SetSourceType(v string) {
	a.Set("sourceType", v)
}
func (a *Audio) StartTime() int {
	return a.Get("startTime").Int()
}
func (a *Audio) SetStartTime(v int) {
	a.Set("startTime", v)
}
func (a *Audio) Type() string {
	return a.Get("type").String()
}
func (a *Audio) SetType(v string) {
	a.Set("type", v)
}
func (a *Audio) Up() *math.Vector3 {
	return &math.Vector3{Value: a.Get("up")}
}
func (a *Audio) SetUp(v *math.Vector3) {
	a.Set("up", v)
}
func (a *Audio) UserData() js.Value {
	return a.Get("userData")
}
func (a *Audio) SetUserData(v js.Value) {
	a.Set("userData", v)
}
func (a *Audio) Uuid() string {
	return a.Get("uuid").String()
}
func (a *Audio) SetUuid(v string) {
	a.Set("uuid", v)
}
func (a *Audio) Visible() bool {
	return a.Get("visible").Bool()
}
func (a *Audio) SetVisible(v bool) {
	a.Set("visible", v)
}
func (a *Audio) DefaultMatrixAutoUpdate() bool {
	return a.Get("DefaultMatrixAutoUpdate").Bool()
}
func (a *Audio) SetDefaultMatrixAutoUpdate(v bool) {
	a.Set("DefaultMatrixAutoUpdate", v)
}
func (a *Audio) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: a.Get("DefaultUp")}
}
func (a *Audio) SetDefaultUp(v *math.Vector3) {
	a.Set("DefaultUp", v)
}
func (a *Audio) Add(object []core.Object3D) this {
	return this(a.Call("add", object))
}
func (a *Audio) AddEventListener(typ string, listener js.Value) {
	a.Call("addEventListener", typ, listener)
}
func (a *Audio) ApplyMatrix(matrix *math.Matrix4) {
	a.Call("applyMatrix", matrix)
}
func (a *Audio) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(a.Call("applyQuaternion", quaternion))
}
func (a *Audio) Clone(recursive bool) this {
	return this(a.Call("clone", recursive))
}
func (a *Audio) Connect() this {
	return this(a.Call("connect"))
}
func (a *Audio) Copy(source *core.Object3D, recursive bool) this {
	return this(a.Call("copy", source, recursive))
}
func (a *Audio) Disconnect() this {
	return this(a.Call("disconnect"))
}
func (a *Audio) DispatchEvent(event js.Value) {
	a.Call("dispatchEvent", event)
}
func (a *Audio) GetDetune() int {
	return a.Call("getDetune").Int()
}
func (a *Audio) GetFilter() js.Value {
	return a.Call("getFilter")
}
func (a *Audio) GetFilters() []js.Value {
	return []js.Value(a.Call("getFilters"))
}
func (a *Audio) GetLoop() bool {
	return a.Call("getLoop").Bool()
}
func (a *Audio) GetObjectById(id int) core.Object3D {
	return core.Object3D(a.Call("getObjectById", id))
}
func (a *Audio) GetObjectByName(name string) core.Object3D {
	return core.Object3D(a.Call("getObjectByName", name))
}
func (a *Audio) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(a.Call("getObjectByProperty", name, value))
}
func (a *Audio) GetOutput() *GainNode {
	return &GainNode{Value: a.Call("getOutput")}
}
func (a *Audio) GetPlaybackRate() int {
	return a.Call("getPlaybackRate").Int()
}
func (a *Audio) GetVolume() int {
	return a.Call("getVolume").Int()
}
func (a *Audio) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: a.Call("getWorldDirection", target)}
}
func (a *Audio) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: a.Call("getWorldPosition", target)}
}
func (a *Audio) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: a.Call("getWorldQuaternion", target)}
}
func (a *Audio) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: a.Call("getWorldScale", target)}
}
func (a *Audio) HasEventListener(typ string, listener js.Value) bool {
	return a.Call("hasEventListener", typ, listener).Bool()
}
func (a *Audio) Load(file string) *Audio {
	return &Audio{Value: a.Call("load", file)}
}
func (a *Audio) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: a.Call("localToWorld", vector)}
}
func (a *Audio) LookAt(vector math.Vector3, y int, z int) {
	a.Call("lookAt", vector, y, z)
}
func (a *Audio) OnEnded() {
	a.Call("onEnded")
}
func (a *Audio) Pause() this {
	return this(a.Call("pause"))
}
func (a *Audio) Play() this {
	return this(a.Call("play"))
}
func (a *Audio) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	a.Call("raycast", raycaster, intersects)
}
func (a *Audio) Remove(object []core.Object3D) this {
	return this(a.Call("remove", object))
}
func (a *Audio) RemoveEventListener(typ string, listener js.Value) {
	a.Call("removeEventListener", typ, listener)
}
func (a *Audio) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(a.Call("rotateOnAxis", axis, angle))
}
func (a *Audio) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(a.Call("rotateOnWorldAxis", axis, angle))
}
func (a *Audio) RotateX(angle int) this {
	return this(a.Call("rotateX", angle))
}
func (a *Audio) RotateY(angle int) this {
	return this(a.Call("rotateY", angle))
}
func (a *Audio) RotateZ(angle int) this {
	return this(a.Call("rotateZ", angle))
}
func (a *Audio) SetBuffer(audioBuffer *AudioBuffer) this {
	return this(a.Call("setBuffer", audioBuffer))
}
func (a *Audio) SetDetune(value int) this {
	return this(a.Call("setDetune", value))
}
func (a *Audio) SetFilter(value []js.Value) this {
	return this(a.Call("setFilter", value))
}
func (a *Audio) SetFilter2(filter js.Value) this {
	return this(a.Call("setFilter", filter))
}
func (a *Audio) SetLoop(value bool) {
	a.Call("setLoop", value)
}
func (a *Audio) SetMediaElementSource(mediaElement *MediaElementAudioSourceNode) this {
	return this(a.Call("setMediaElementSource", mediaElement))
}
func (a *Audio) SetNodeSource(audioNode *AudioBufferSourceNode) this {
	return this(a.Call("setNodeSource", audioNode))
}
func (a *Audio) SetPlaybackRate(value int) this {
	return this(a.Call("setPlaybackRate", value))
}
func (a *Audio) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	a.Call("setRotationFromAxisAngle", axis, angle)
}
func (a *Audio) SetRotationFromEuler(euler *math.Euler) {
	a.Call("setRotationFromEuler", euler)
}
func (a *Audio) SetRotationFromMatrix(m *math.Matrix4) {
	a.Call("setRotationFromMatrix", m)
}
func (a *Audio) SetRotationFromQuaternion(q *math.Quaternion) {
	a.Call("setRotationFromQuaternion", q)
}
func (a *Audio) SetVolume(value int) this {
	return this(a.Call("setVolume", value))
}
func (a *Audio) Stop() this {
	return this(a.Call("stop"))
}
func (a *Audio) ToJSON(meta js.Value) js.Value {
	return a.Call("toJSON", meta)
}
func (a *Audio) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(a.Call("translateOnAxis", axis, distance))
}
func (a *Audio) TranslateX(distance int) this {
	return this(a.Call("translateX", distance))
}
func (a *Audio) TranslateY(distance int) this {
	return this(a.Call("translateY", distance))
}
func (a *Audio) TranslateZ(distance int) this {
	return this(a.Call("translateZ", distance))
}
func (a *Audio) Traverse(callback js.Value) {
	a.Call("traverse", callback)
}
func (a *Audio) TraverseAncestors(callback js.Value) {
	a.Call("traverseAncestors", callback)
}
func (a *Audio) TraverseVisible(callback js.Value) {
	a.Call("traverseVisible", callback)
}
func (a *Audio) UpdateMatrix() {
	a.Call("updateMatrix")
}
func (a *Audio) UpdateMatrixWorld(force bool) {
	a.Call("updateMatrixWorld", force)
}
func (a *Audio) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	a.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (a *Audio) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: a.Call("worldToLocal", vector)}
}

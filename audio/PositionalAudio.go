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

type AudioBuffer struct {
	js.Value
}

func NewAudioBuffer(context js.Value) *AudioBuffer {
	return &AudioBuffer{Value: get("AudioBuffer").New(context)}
}
func (ab *AudioBuffer) Context() js.Value {
	return ab.Get("context")
}
func (ab *AudioBuffer) SetContext(v js.Value) {
	ab.Set("context", v)
}
func (ab *AudioBuffer) Ready() bool {
	return ab.Get("ready").Bool()
}
func (ab *AudioBuffer) SetReady(v bool) {
	ab.Set("ready", v)
}
func (ab *AudioBuffer) ReadyCallbacks() []js.Value {
	return []js.Value(ab.Get("readyCallbacks"))
}
func (ab *AudioBuffer) SetReadyCallbacks(v []js.Value) {
	ab.Set("readyCallbacks", v)
}
func (ab *AudioBuffer) Load(file string) *AudioBuffer {
	return &AudioBuffer{Value: ab.Call("load", file)}
}
func (ab *AudioBuffer) OnReady(callback js.Value) {
	ab.Call("onReady", callback)
}

type PositionalAudio struct {
	js.Value
}

func NewPositionalAudio(listener *AudioListener) *PositionalAudio {
	return &PositionalAudio{Value: get("PositionalAudio").New(listener)}
}
func (pa *PositionalAudio) Autoplay() bool {
	return pa.Get("autoplay").Bool()
}
func (pa *PositionalAudio) SetAutoplay(v bool) {
	pa.Set("autoplay", v)
}
func (pa *PositionalAudio) Buffer() null {
	return null(pa.Get("buffer"))
}
func (pa *PositionalAudio) SetBuffer(v null) {
	pa.Set("buffer", v)
}
func (pa *PositionalAudio) CastShadow() bool {
	return pa.Get("castShadow").Bool()
}
func (pa *PositionalAudio) SetCastShadow(v bool) {
	pa.Set("castShadow", v)
}
func (pa *PositionalAudio) Children() []core.Object3D {
	return []core.Object3D(pa.Get("children"))
}
func (pa *PositionalAudio) SetChildren(v []core.Object3D) {
	pa.Set("children", v)
}
func (pa *PositionalAudio) Context() *AudioContext {
	return &AudioContext{Value: pa.Get("context")}
}
func (pa *PositionalAudio) SetContext(v *AudioContext) {
	pa.Set("context", v)
}
func (pa *PositionalAudio) Detune() int {
	return pa.Get("detune").Int()
}
func (pa *PositionalAudio) SetDetune(v int) {
	pa.Set("detune", v)
}
func (pa *PositionalAudio) Filters() []js.Value {
	return []js.Value(pa.Get("filters"))
}
func (pa *PositionalAudio) SetFilters(v []js.Value) {
	pa.Set("filters", v)
}
func (pa *PositionalAudio) FrustumCulled() bool {
	return pa.Get("frustumCulled").Bool()
}
func (pa *PositionalAudio) SetFrustumCulled(v bool) {
	pa.Set("frustumCulled", v)
}
func (pa *PositionalAudio) Gain() *GainNode {
	return &GainNode{Value: pa.Get("gain")}
}
func (pa *PositionalAudio) SetGain(v *GainNode) {
	pa.Set("gain", v)
}
func (pa *PositionalAudio) HasPlaybackControl() bool {
	return pa.Get("hasPlaybackControl").Bool()
}
func (pa *PositionalAudio) SetHasPlaybackControl(v bool) {
	pa.Set("hasPlaybackControl", v)
}
func (pa *PositionalAudio) Id() int {
	return pa.Get("id").Int()
}
func (pa *PositionalAudio) SetId(v int) {
	pa.Set("id", v)
}
func (pa *PositionalAudio) IsObject3D() true {
	return true(pa.Get("isObject3D"))
}
func (pa *PositionalAudio) SetIsObject3D(v true) {
	pa.Set("isObject3D", v)
}
func (pa *PositionalAudio) IsPlaying() bool {
	return pa.Get("isPlaying").Bool()
}
func (pa *PositionalAudio) SetIsPlaying(v bool) {
	pa.Set("isPlaying", v)
}
func (pa *PositionalAudio) Layers() *core.Layers {
	return &core.Layers{Value: pa.Get("layers")}
}
func (pa *PositionalAudio) SetLayers(v *core.Layers) {
	pa.Set("layers", v)
}
func (pa *PositionalAudio) Loop() bool {
	return pa.Get("loop").Bool()
}
func (pa *PositionalAudio) SetLoop(v bool) {
	pa.Set("loop", v)
}
func (pa *PositionalAudio) Matrix() *math.Matrix4 {
	return &math.Matrix4{Value: pa.Get("matrix")}
}
func (pa *PositionalAudio) SetMatrix(v *math.Matrix4) {
	pa.Set("matrix", v)
}
func (pa *PositionalAudio) MatrixAutoUpdate() bool {
	return pa.Get("matrixAutoUpdate").Bool()
}
func (pa *PositionalAudio) SetMatrixAutoUpdate(v bool) {
	pa.Set("matrixAutoUpdate", v)
}
func (pa *PositionalAudio) MatrixWorld() *math.Matrix4 {
	return &math.Matrix4{Value: pa.Get("matrixWorld")}
}
func (pa *PositionalAudio) SetMatrixWorld(v *math.Matrix4) {
	pa.Set("matrixWorld", v)
}
func (pa *PositionalAudio) MatrixWorldNeedsUpdate() bool {
	return pa.Get("matrixWorldNeedsUpdate").Bool()
}
func (pa *PositionalAudio) SetMatrixWorldNeedsUpdate(v bool) {
	pa.Set("matrixWorldNeedsUpdate", v)
}
func (pa *PositionalAudio) ModelViewMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: pa.Get("modelViewMatrix")}
}
func (pa *PositionalAudio) SetModelViewMatrix(v *math.Matrix4) {
	pa.Set("modelViewMatrix", v)
}
func (pa *PositionalAudio) Name() string {
	return pa.Get("name").String()
}
func (pa *PositionalAudio) SetName(v string) {
	pa.Set("name", v)
}
func (pa *PositionalAudio) NormalMatrix() *math.Matrix3 {
	return &math.Matrix3{Value: pa.Get("normalMatrix")}
}
func (pa *PositionalAudio) SetNormalMatrix(v *math.Matrix3) {
	pa.Set("normalMatrix", v)
}
func (pa *PositionalAudio) Offset() int {
	return pa.Get("offset").Int()
}
func (pa *PositionalAudio) SetOffset(v int) {
	pa.Set("offset", v)
}
func (pa *PositionalAudio) OnAfterRender() js.Value {
	return pa.Get("onAfterRender")
}
func (pa *PositionalAudio) SetOnAfterRender(v js.Value) {
	pa.Set("onAfterRender", v)
}
func (pa *PositionalAudio) OnBeforeRender() js.Value {
	return pa.Get("onBeforeRender")
}
func (pa *PositionalAudio) SetOnBeforeRender(v js.Value) {
	pa.Set("onBeforeRender", v)
}
func (pa *PositionalAudio) Panner() *PannerNode {
	return &PannerNode{Value: pa.Get("panner")}
}
func (pa *PositionalAudio) SetPanner(v *PannerNode) {
	pa.Set("panner", v)
}
func (pa *PositionalAudio) Parent() core.Object3D {
	return core.Object3D(pa.Get("parent"))
}
func (pa *PositionalAudio) SetParent(v core.Object3D) {
	pa.Set("parent", v)
}
func (pa *PositionalAudio) PlaybackRate() int {
	return pa.Get("playbackRate").Int()
}
func (pa *PositionalAudio) SetPlaybackRate(v int) {
	pa.Set("playbackRate", v)
}
func (pa *PositionalAudio) Position() *math.Vector3 {
	return &math.Vector3{Value: pa.Get("position")}
}
func (pa *PositionalAudio) SetPosition(v *math.Vector3) {
	pa.Set("position", v)
}
func (pa *PositionalAudio) Quaternion() *math.Quaternion {
	return &math.Quaternion{Value: pa.Get("quaternion")}
}
func (pa *PositionalAudio) SetQuaternion(v *math.Quaternion) {
	pa.Set("quaternion", v)
}
func (pa *PositionalAudio) ReceiveShadow() bool {
	return pa.Get("receiveShadow").Bool()
}
func (pa *PositionalAudio) SetReceiveShadow(v bool) {
	pa.Set("receiveShadow", v)
}
func (pa *PositionalAudio) RenderOrder() int {
	return pa.Get("renderOrder").Int()
}
func (pa *PositionalAudio) SetRenderOrder(v int) {
	pa.Set("renderOrder", v)
}
func (pa *PositionalAudio) Rotation() *math.Euler {
	return &math.Euler{Value: pa.Get("rotation")}
}
func (pa *PositionalAudio) SetRotation(v *math.Euler) {
	pa.Set("rotation", v)
}
func (pa *PositionalAudio) Scale() *math.Vector3 {
	return &math.Vector3{Value: pa.Get("scale")}
}
func (pa *PositionalAudio) SetScale(v *math.Vector3) {
	pa.Set("scale", v)
}
func (pa *PositionalAudio) Source() *AudioBufferSourceNode {
	return &AudioBufferSourceNode{Value: pa.Get("source")}
}
func (pa *PositionalAudio) SetSource(v *AudioBufferSourceNode) {
	pa.Set("source", v)
}
func (pa *PositionalAudio) SourceType() string {
	return pa.Get("sourceType").String()
}
func (pa *PositionalAudio) SetSourceType(v string) {
	pa.Set("sourceType", v)
}
func (pa *PositionalAudio) StartTime() int {
	return pa.Get("startTime").Int()
}
func (pa *PositionalAudio) SetStartTime(v int) {
	pa.Set("startTime", v)
}
func (pa *PositionalAudio) Type() string {
	return pa.Get("type").String()
}
func (pa *PositionalAudio) SetType(v string) {
	pa.Set("type", v)
}
func (pa *PositionalAudio) Up() *math.Vector3 {
	return &math.Vector3{Value: pa.Get("up")}
}
func (pa *PositionalAudio) SetUp(v *math.Vector3) {
	pa.Set("up", v)
}
func (pa *PositionalAudio) UserData() js.Value {
	return pa.Get("userData")
}
func (pa *PositionalAudio) SetUserData(v js.Value) {
	pa.Set("userData", v)
}
func (pa *PositionalAudio) Uuid() string {
	return pa.Get("uuid").String()
}
func (pa *PositionalAudio) SetUuid(v string) {
	pa.Set("uuid", v)
}
func (pa *PositionalAudio) Visible() bool {
	return pa.Get("visible").Bool()
}
func (pa *PositionalAudio) SetVisible(v bool) {
	pa.Set("visible", v)
}
func (pa *PositionalAudio) DefaultMatrixAutoUpdate() bool {
	return pa.Get("DefaultMatrixAutoUpdate").Bool()
}
func (pa *PositionalAudio) SetDefaultMatrixAutoUpdate(v bool) {
	pa.Set("DefaultMatrixAutoUpdate", v)
}
func (pa *PositionalAudio) DefaultUp() *math.Vector3 {
	return &math.Vector3{Value: pa.Get("DefaultUp")}
}
func (pa *PositionalAudio) SetDefaultUp(v *math.Vector3) {
	pa.Set("DefaultUp", v)
}
func (pa *PositionalAudio) Add(object []core.Object3D) this {
	return this(pa.Call("add", object))
}
func (pa *PositionalAudio) AddEventListener(typ string, listener js.Value) {
	pa.Call("addEventListener", typ, listener)
}
func (pa *PositionalAudio) ApplyMatrix(matrix *math.Matrix4) {
	pa.Call("applyMatrix", matrix)
}
func (pa *PositionalAudio) ApplyQuaternion(quaternion *math.Quaternion) this {
	return this(pa.Call("applyQuaternion", quaternion))
}
func (pa *PositionalAudio) Clone(recursive bool) this {
	return this(pa.Call("clone", recursive))
}
func (pa *PositionalAudio) Connect() this {
	return this(pa.Call("connect"))
}
func (pa *PositionalAudio) Copy(source *core.Object3D, recursive bool) this {
	return this(pa.Call("copy", source, recursive))
}
func (pa *PositionalAudio) Disconnect() this {
	return this(pa.Call("disconnect"))
}
func (pa *PositionalAudio) DispatchEvent(event js.Value) {
	pa.Call("dispatchEvent", event)
}
func (pa *PositionalAudio) GetDetune() int {
	return pa.Call("getDetune").Int()
}
func (pa *PositionalAudio) GetDistanceModel() string {
	return pa.Call("getDistanceModel").String()
}
func (pa *PositionalAudio) GetFilter() js.Value {
	return pa.Call("getFilter")
}
func (pa *PositionalAudio) GetFilters() []js.Value {
	return []js.Value(pa.Call("getFilters"))
}
func (pa *PositionalAudio) GetLoop() bool {
	return pa.Call("getLoop").Bool()
}
func (pa *PositionalAudio) GetMaxDistance() int {
	return pa.Call("getMaxDistance").Int()
}
func (pa *PositionalAudio) GetObjectById(id int) core.Object3D {
	return core.Object3D(pa.Call("getObjectById", id))
}
func (pa *PositionalAudio) GetObjectByName(name string) core.Object3D {
	return core.Object3D(pa.Call("getObjectByName", name))
}
func (pa *PositionalAudio) GetObjectByProperty(name string, value string) core.Object3D {
	return core.Object3D(pa.Call("getObjectByProperty", name, value))
}
func (pa *PositionalAudio) GetOutput() *GainNode {
	return &GainNode{Value: pa.Call("getOutput")}
}
func (pa *PositionalAudio) GetPlaybackRate() int {
	return pa.Call("getPlaybackRate").Int()
}
func (pa *PositionalAudio) GetRefDistance() int {
	return pa.Call("getRefDistance").Int()
}
func (pa *PositionalAudio) GetRolloffFactor() int {
	return pa.Call("getRolloffFactor").Int()
}
func (pa *PositionalAudio) GetVolume() int {
	return pa.Call("getVolume").Int()
}
func (pa *PositionalAudio) GetWorldDirection(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pa.Call("getWorldDirection", target)}
}
func (pa *PositionalAudio) GetWorldPosition(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pa.Call("getWorldPosition", target)}
}
func (pa *PositionalAudio) GetWorldQuaternion(target *math.Quaternion) *math.Quaternion {
	return &math.Quaternion{Value: pa.Call("getWorldQuaternion", target)}
}
func (pa *PositionalAudio) GetWorldScale(target *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pa.Call("getWorldScale", target)}
}
func (pa *PositionalAudio) HasEventListener(typ string, listener js.Value) bool {
	return pa.Call("hasEventListener", typ, listener).Bool()
}
func (pa *PositionalAudio) Load(file string) *Audio {
	return &Audio{Value: pa.Call("load", file)}
}
func (pa *PositionalAudio) LocalToWorld(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pa.Call("localToWorld", vector)}
}
func (pa *PositionalAudio) LookAt(vector math.Vector3, y int, z int) {
	pa.Call("lookAt", vector, y, z)
}
func (pa *PositionalAudio) OnEnded() {
	pa.Call("onEnded")
}
func (pa *PositionalAudio) Pause() this {
	return this(pa.Call("pause"))
}
func (pa *PositionalAudio) Play() this {
	return this(pa.Call("play"))
}
func (pa *PositionalAudio) Raycast(raycaster *core.Raycaster, intersects []core.Intersection) {
	pa.Call("raycast", raycaster, intersects)
}
func (pa *PositionalAudio) Remove(object []core.Object3D) this {
	return this(pa.Call("remove", object))
}
func (pa *PositionalAudio) RemoveEventListener(typ string, listener js.Value) {
	pa.Call("removeEventListener", typ, listener)
}
func (pa *PositionalAudio) RotateOnAxis(axis *math.Vector3, angle int) this {
	return this(pa.Call("rotateOnAxis", axis, angle))
}
func (pa *PositionalAudio) RotateOnWorldAxis(axis *math.Vector3, angle int) this {
	return this(pa.Call("rotateOnWorldAxis", axis, angle))
}
func (pa *PositionalAudio) RotateX(angle int) this {
	return this(pa.Call("rotateX", angle))
}
func (pa *PositionalAudio) RotateY(angle int) this {
	return this(pa.Call("rotateY", angle))
}
func (pa *PositionalAudio) RotateZ(angle int) this {
	return this(pa.Call("rotateZ", angle))
}
func (pa *PositionalAudio) SetBuffer(audioBuffer *AudioBuffer) this {
	return this(pa.Call("setBuffer", audioBuffer))
}
func (pa *PositionalAudio) SetDetune(value int) this {
	return this(pa.Call("setDetune", value))
}
func (pa *PositionalAudio) SetDirectionalCone(coneInnerAngle int, coneOuterAngle int, coneOuterGain int) this {
	return this(pa.Call("setDirectionalCone", coneInnerAngle, coneOuterAngle, coneOuterGain))
}
func (pa *PositionalAudio) SetDistanceModel(value string) this {
	return this(pa.Call("setDistanceModel", value))
}
func (pa *PositionalAudio) SetFilter(value []js.Value) this {
	return this(pa.Call("setFilter", value))
}
func (pa *PositionalAudio) SetFilter2(filter js.Value) this {
	return this(pa.Call("setFilter", filter))
}
func (pa *PositionalAudio) SetLoop(value bool) {
	pa.Call("setLoop", value)
}
func (pa *PositionalAudio) SetMaxDistance(value int) this {
	return this(pa.Call("setMaxDistance", value))
}
func (pa *PositionalAudio) SetMediaElementSource(mediaElement *MediaElementAudioSourceNode) this {
	return this(pa.Call("setMediaElementSource", mediaElement))
}
func (pa *PositionalAudio) SetNodeSource(audioNode *AudioBufferSourceNode) this {
	return this(pa.Call("setNodeSource", audioNode))
}
func (pa *PositionalAudio) SetPlaybackRate(value int) this {
	return this(pa.Call("setPlaybackRate", value))
}
func (pa *PositionalAudio) SetRefDistance(value int) this {
	return this(pa.Call("setRefDistance", value))
}
func (pa *PositionalAudio) SetRolloffFactor(value int) this {
	return this(pa.Call("setRolloffFactor", value))
}
func (pa *PositionalAudio) SetRotationFromAxisAngle(axis *math.Vector3, angle int) {
	pa.Call("setRotationFromAxisAngle", axis, angle)
}
func (pa *PositionalAudio) SetRotationFromEuler(euler *math.Euler) {
	pa.Call("setRotationFromEuler", euler)
}
func (pa *PositionalAudio) SetRotationFromMatrix(m *math.Matrix4) {
	pa.Call("setRotationFromMatrix", m)
}
func (pa *PositionalAudio) SetRotationFromQuaternion(q *math.Quaternion) {
	pa.Call("setRotationFromQuaternion", q)
}
func (pa *PositionalAudio) SetVolume(value int) this {
	return this(pa.Call("setVolume", value))
}
func (pa *PositionalAudio) Stop() this {
	return this(pa.Call("stop"))
}
func (pa *PositionalAudio) ToJSON(meta js.Value) js.Value {
	return pa.Call("toJSON", meta)
}
func (pa *PositionalAudio) TranslateOnAxis(axis *math.Vector3, distance int) this {
	return this(pa.Call("translateOnAxis", axis, distance))
}
func (pa *PositionalAudio) TranslateX(distance int) this {
	return this(pa.Call("translateX", distance))
}
func (pa *PositionalAudio) TranslateY(distance int) this {
	return this(pa.Call("translateY", distance))
}
func (pa *PositionalAudio) TranslateZ(distance int) this {
	return this(pa.Call("translateZ", distance))
}
func (pa *PositionalAudio) Traverse(callback js.Value) {
	pa.Call("traverse", callback)
}
func (pa *PositionalAudio) TraverseAncestors(callback js.Value) {
	pa.Call("traverseAncestors", callback)
}
func (pa *PositionalAudio) TraverseVisible(callback js.Value) {
	pa.Call("traverseVisible", callback)
}
func (pa *PositionalAudio) UpdateMatrix() {
	pa.Call("updateMatrix")
}
func (pa *PositionalAudio) UpdateMatrixWorld(force bool) {
	pa.Call("updateMatrixWorld", force)
}
func (pa *PositionalAudio) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pa.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (pa *PositionalAudio) WorldToLocal(vector *math.Vector3) *math.Vector3 {
	return &math.Vector3{Value: pa.Call("worldToLocal", vector)}
}

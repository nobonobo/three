// Code generated from "audio/Audio.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// Audio extend: [Object3D]
type Audio struct {
	js.Value
}

func NewAudio(listener *AudioListener) *Audio {
	return &Audio{Value: get("Audio").New(listener.JSValue())}
}
func (aa *Audio) JSValue() js.Value {
	return aa.Value
}
func (aa *Audio) Autoplay() bool {
	return aa.Get("autoplay").Bool()
}
func (aa *Audio) SetAutoplay(v bool) {
	aa.Set("autoplay", v)
}
func (aa *Audio) Buffer() js.Value {
	return aa.Get("buffer")
}
func (aa *Audio) SetBuffer(v js.Value) {
	aa.Set("buffer", v)
}
func (aa *Audio) CastShadow() bool {
	return aa.Get("castShadow").Bool()
}
func (aa *Audio) SetCastShadow(v bool) {
	aa.Set("castShadow", v)
}
func (aa *Audio) Children() js.Value {
	return aa.Get("children")
}
func (aa *Audio) SetChildren(v js.Value) {
	aa.Set("children", v)
}
func (aa *Audio) Context() js.Value {
	return aa.Get("context")
}
func (aa *Audio) SetContext(v js.Value) {
	aa.Set("context", v)
}
func (aa *Audio) Detune() float64 {
	return aa.Get("detune").Float()
}
func (aa *Audio) SetDetune(v float64) {
	aa.Set("detune", v)
}
func (aa *Audio) Filters() js.Value {
	return aa.Get("filters")
}
func (aa *Audio) SetFilters(v js.Value) {
	aa.Set("filters", v)
}
func (aa *Audio) FrustumCulled() bool {
	return aa.Get("frustumCulled").Bool()
}
func (aa *Audio) SetFrustumCulled(v bool) {
	aa.Set("frustumCulled", v)
}
func (aa *Audio) Gain() js.Value {
	return aa.Get("gain")
}
func (aa *Audio) SetGain(v js.Value) {
	aa.Set("gain", v)
}
func (aa *Audio) HasPlaybackControl() bool {
	return aa.Get("hasPlaybackControl").Bool()
}
func (aa *Audio) SetHasPlaybackControl(v bool) {
	aa.Set("hasPlaybackControl", v)
}
func (aa *Audio) Id() int {
	return aa.Get("id").Int()
}
func (aa *Audio) SetId(v int) {
	aa.Set("id", v)
}
func (aa *Audio) IsObject3D() bool {
	return aa.Get("isObject3D").Bool()
}
func (aa *Audio) SetIsObject3D(v bool) {
	aa.Set("isObject3D", v)
}
func (aa *Audio) IsPlaying() bool {
	return aa.Get("isPlaying").Bool()
}
func (aa *Audio) SetIsPlaying(v bool) {
	aa.Set("isPlaying", v)
}
func (aa *Audio) Layers() *Layers {
	return &Layers{Value: aa.Get("layers")}
}
func (aa *Audio) SetLayers(v *Layers) {
	aa.Set("layers", v.JSValue())
}
func (aa *Audio) Loop() bool {
	return aa.Get("loop").Bool()
}
func (aa *Audio) SetLoop(v bool) {
	aa.Set("loop", v)
}
func (aa *Audio) Matrix() *Matrix4 {
	return &Matrix4{Value: aa.Get("matrix")}
}
func (aa *Audio) SetMatrix(v *Matrix4) {
	aa.Set("matrix", v.JSValue())
}
func (aa *Audio) MatrixAutoUpdate() bool {
	return aa.Get("matrixAutoUpdate").Bool()
}
func (aa *Audio) SetMatrixAutoUpdate(v bool) {
	aa.Set("matrixAutoUpdate", v)
}
func (aa *Audio) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: aa.Get("matrixWorld")}
}
func (aa *Audio) SetMatrixWorld(v *Matrix4) {
	aa.Set("matrixWorld", v.JSValue())
}
func (aa *Audio) MatrixWorldNeedsUpdate() bool {
	return aa.Get("matrixWorldNeedsUpdate").Bool()
}
func (aa *Audio) SetMatrixWorldNeedsUpdate(v bool) {
	aa.Set("matrixWorldNeedsUpdate", v)
}
func (aa *Audio) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: aa.Get("modelViewMatrix")}
}
func (aa *Audio) SetModelViewMatrix(v *Matrix4) {
	aa.Set("modelViewMatrix", v.JSValue())
}
func (aa *Audio) Name() string {
	return aa.Get("name").String()
}
func (aa *Audio) SetName(v string) {
	aa.Set("name", v)
}
func (aa *Audio) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: aa.Get("normalMatrix")}
}
func (aa *Audio) SetNormalMatrix(v *Matrix3) {
	aa.Set("normalMatrix", v.JSValue())
}
func (aa *Audio) Offset() float64 {
	return aa.Get("offset").Float()
}
func (aa *Audio) SetOffset(v float64) {
	aa.Set("offset", v)
}
func (aa *Audio) OnAfterRender() js.Value {
	return aa.Get("onAfterRender")
}
func (aa *Audio) SetOnAfterRender(v js.Value) {
	aa.Set("onAfterRender", v)
}
func (aa *Audio) OnBeforeRender() js.Value {
	return aa.Get("onBeforeRender")
}
func (aa *Audio) SetOnBeforeRender(v js.Value) {
	aa.Set("onBeforeRender", v)
}
func (aa *Audio) Parent() *Object3D {
	return &Object3D{Value: aa.Get("parent")}
}
func (aa *Audio) SetParent(v *Object3D) {
	aa.Set("parent", v.JSValue())
}
func (aa *Audio) PlaybackRate() float64 {
	return aa.Get("playbackRate").Float()
}
func (aa *Audio) SetPlaybackRate(v float64) {
	aa.Set("playbackRate", v)
}
func (aa *Audio) Position() *Vector3 {
	return &Vector3{Value: aa.Get("position")}
}
func (aa *Audio) SetPosition(v *Vector3) {
	aa.Set("position", v.JSValue())
}
func (aa *Audio) Quaternion() *Quaternion {
	return &Quaternion{Value: aa.Get("quaternion")}
}
func (aa *Audio) SetQuaternion(v *Quaternion) {
	aa.Set("quaternion", v.JSValue())
}
func (aa *Audio) ReceiveShadow() bool {
	return aa.Get("receiveShadow").Bool()
}
func (aa *Audio) SetReceiveShadow(v bool) {
	aa.Set("receiveShadow", v)
}
func (aa *Audio) RenderOrder() float64 {
	return aa.Get("renderOrder").Float()
}
func (aa *Audio) SetRenderOrder(v float64) {
	aa.Set("renderOrder", v)
}
func (aa *Audio) Rotation() *Euler {
	return &Euler{Value: aa.Get("rotation")}
}
func (aa *Audio) SetRotation(v *Euler) {
	aa.Set("rotation", v.JSValue())
}
func (aa *Audio) Scale() *Vector3 {
	return &Vector3{Value: aa.Get("scale")}
}
func (aa *Audio) SetScale(v *Vector3) {
	aa.Set("scale", v.JSValue())
}
func (aa *Audio) Source() js.Value {
	return aa.Get("source")
}
func (aa *Audio) SetSource(v js.Value) {
	aa.Set("source", v)
}
func (aa *Audio) SourceType() string {
	return aa.Get("sourceType").String()
}
func (aa *Audio) SetSourceType(v string) {
	aa.Set("sourceType", v)
}
func (aa *Audio) StartTime() float64 {
	return aa.Get("startTime").Float()
}
func (aa *Audio) SetStartTime(v float64) {
	aa.Set("startTime", v)
}
func (aa *Audio) Type() string {
	return aa.Get("type").String()
}
func (aa *Audio) SetType(v string) {
	aa.Set("type", v)
}
func (aa *Audio) Up() *Vector3 {
	return &Vector3{Value: aa.Get("up")}
}
func (aa *Audio) SetUp(v *Vector3) {
	aa.Set("up", v.JSValue())
}
func (aa *Audio) UserData() js.Value {
	return aa.Get("userData")
}
func (aa *Audio) SetUserData(v js.Value) {
	aa.Set("userData", v)
}
func (aa *Audio) Uuid() string {
	return aa.Get("uuid").String()
}
func (aa *Audio) SetUuid(v string) {
	aa.Set("uuid", v)
}
func (aa *Audio) Visible() bool {
	return aa.Get("visible").Bool()
}
func (aa *Audio) SetVisible(v bool) {
	aa.Set("visible", v)
}
func (aa *Audio) DefaultMatrixAutoUpdate() bool {
	return aa.Get("DefaultMatrixAutoUpdate").Bool()
}
func (aa *Audio) SetDefaultMatrixAutoUpdate(v bool) {
	aa.Set("DefaultMatrixAutoUpdate", v)
}
func (aa *Audio) DefaultUp() *Vector3 {
	return &Vector3{Value: aa.Get("DefaultUp")}
}
func (aa *Audio) SetDefaultUp(v *Vector3) {
	aa.Set("DefaultUp", v.JSValue())
}
func (aa *Audio) Add(object js.Value) *Audio {
	return &Audio{Value: aa.Call("add", object)}
}
func (aa *Audio) AddEventListener(typ string, listener js.Value) {
	aa.Call("addEventListener", typ, listener)
}
func (aa *Audio) ApplyMatrix(matrix *Matrix4) {
	aa.Call("applyMatrix", matrix.JSValue())
}
func (aa *Audio) ApplyQuaternion(quaternion *Quaternion) *Audio {
	return &Audio{Value: aa.Call("applyQuaternion", quaternion)}
}
func (aa *Audio) Clone(recursive bool) *Audio {
	return &Audio{Value: aa.Call("clone", recursive)}
}
func (aa *Audio) Connect() *Audio {
	return &Audio{Value: aa.Call("connect")}
}
func (aa *Audio) Copy(source *Object3D, recursive bool) *Audio {
	return &Audio{Value: aa.Call("copy", source, recursive)}
}
func (aa *Audio) Disconnect() *Audio {
	return &Audio{Value: aa.Call("disconnect")}
}
func (aa *Audio) DispatchEvent(event js.Value) {
	aa.Call("dispatchEvent", event)
}
func (aa *Audio) GetDetune() float64 {
	return aa.Call("getDetune").Float()
}
func (aa *Audio) GetFilter() js.Value {
	return aa.Call("getFilter")
}
func (aa *Audio) GetFilters() js.Value {
	return aa.Call("getFilters")
}
func (aa *Audio) GetLoop() bool {
	return aa.Call("getLoop").Bool()
}
func (aa *Audio) GetObjectById(id int) *Object3D {
	return &Object3D{Value: aa.Call("getObjectById", id)}
}
func (aa *Audio) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: aa.Call("getObjectByName", name)}
}
func (aa *Audio) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: aa.Call("getObjectByProperty", name, value)}
}
func (aa *Audio) GetOutput() js.Value {
	return aa.Call("getOutput")
}
func (aa *Audio) GetPlaybackRate() float64 {
	return aa.Call("getPlaybackRate").Float()
}
func (aa *Audio) GetVolume() float64 {
	return aa.Call("getVolume").Float()
}
func (aa *Audio) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: aa.Call("getWorldDirection", target)}
}
func (aa *Audio) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: aa.Call("getWorldPosition", target)}
}
func (aa *Audio) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: aa.Call("getWorldQuaternion", target)}
}
func (aa *Audio) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: aa.Call("getWorldScale", target)}
}
func (aa *Audio) HasEventListener(typ string, listener js.Value) bool {
	return aa.Call("hasEventListener", typ, listener).Bool()
}
func (aa *Audio) Load(file string) *Audio {
	return &Audio{Value: aa.Call("load", file)}
}
func (aa *Audio) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: aa.Call("localToWorld", vector)}
}
func (aa *Audio) LookAt(vector *Vector3, y float64, z float64) {
	aa.Call("lookAt", vector, y, z)
}
func (aa *Audio) OnEnded() {
	aa.Call("onEnded")
}
func (aa *Audio) Pause() *Audio {
	return &Audio{Value: aa.Call("pause")}
}
func (aa *Audio) Play() *Audio {
	return &Audio{Value: aa.Call("play")}
}
func (aa *Audio) Raycast(raycaster *Raycaster, intersects js.Value) {
	aa.Call("raycast", raycaster.JSValue(), intersects)
}
func (aa *Audio) Remove(object js.Value) *Audio {
	return &Audio{Value: aa.Call("remove", object)}
}
func (aa *Audio) RemoveEventListener(typ string, listener js.Value) {
	aa.Call("removeEventListener", typ, listener)
}
func (aa *Audio) RotateOnAxis(axis *Vector3, angle float64) *Audio {
	return &Audio{Value: aa.Call("rotateOnAxis", axis, angle)}
}
func (aa *Audio) RotateOnWorldAxis(axis *Vector3, angle float64) *Audio {
	return &Audio{Value: aa.Call("rotateOnWorldAxis", axis, angle)}
}
func (aa *Audio) RotateX(angle float64) *Audio {
	return &Audio{Value: aa.Call("rotateX", angle)}
}
func (aa *Audio) RotateY(angle float64) *Audio {
	return &Audio{Value: aa.Call("rotateY", angle)}
}
func (aa *Audio) RotateZ(angle float64) *Audio {
	return &Audio{Value: aa.Call("rotateZ", angle)}
}
func (aa *Audio) SetBuffer2(audioBuffer *AudioBuffer) *Audio {
	return &Audio{Value: aa.Call("setBuffer", audioBuffer)}
}
func (aa *Audio) SetDetune2(value float64) *Audio {
	return &Audio{Value: aa.Call("setDetune", value)}
}
func (aa *Audio) SetFilter(value js.Value) *Audio {
	return &Audio{Value: aa.Call("setFilter", value)}
}
func (aa *Audio) SetFilter2(filter js.Value) *Audio {
	return &Audio{Value: aa.Call("setFilter", filter)}
}
func (aa *Audio) SetLoop2(value bool) {
	aa.Call("setLoop", value)
}
func (aa *Audio) SetMediaElementSource(mediaElement js.Value) *Audio {
	return &Audio{Value: aa.Call("setMediaElementSource", mediaElement)}
}
func (aa *Audio) SetNodeSource(audioNode js.Value) *Audio {
	return &Audio{Value: aa.Call("setNodeSource", audioNode)}
}
func (aa *Audio) SetPlaybackRate2(value float64) *Audio {
	return &Audio{Value: aa.Call("setPlaybackRate", value)}
}
func (aa *Audio) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	aa.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}
func (aa *Audio) SetRotationFromEuler(euler *Euler) {
	aa.Call("setRotationFromEuler", euler.JSValue())
}
func (aa *Audio) SetRotationFromMatrix(m *Matrix4) {
	aa.Call("setRotationFromMatrix", m.JSValue())
}
func (aa *Audio) SetRotationFromQuaternion(q *Quaternion) {
	aa.Call("setRotationFromQuaternion", q.JSValue())
}
func (aa *Audio) SetVolume(value float64) *Audio {
	return &Audio{Value: aa.Call("setVolume", value)}
}
func (aa *Audio) Stop() *Audio {
	return &Audio{Value: aa.Call("stop")}
}
func (aa *Audio) ToJSON(meta js.Value) js.Value {
	return aa.Call("toJSON", meta)
}
func (aa *Audio) TranslateOnAxis(axis *Vector3, distance float64) *Audio {
	return &Audio{Value: aa.Call("translateOnAxis", axis, distance)}
}
func (aa *Audio) TranslateX(distance float64) *Audio {
	return &Audio{Value: aa.Call("translateX", distance)}
}
func (aa *Audio) TranslateY(distance float64) *Audio {
	return &Audio{Value: aa.Call("translateY", distance)}
}
func (aa *Audio) TranslateZ(distance float64) *Audio {
	return &Audio{Value: aa.Call("translateZ", distance)}
}
func (aa *Audio) Traverse(callback js.Value) {
	aa.Call("traverse", callback)
}
func (aa *Audio) TraverseAncestors(callback js.Value) {
	aa.Call("traverseAncestors", callback)
}
func (aa *Audio) TraverseVisible(callback js.Value) {
	aa.Call("traverseVisible", callback)
}
func (aa *Audio) UpdateMatrix() {
	aa.Call("updateMatrix")
}
func (aa *Audio) UpdateMatrixWorld(force bool) {
	aa.Call("updateMatrixWorld", force)
}
func (aa *Audio) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	aa.Call("updateWorldMatrix", updateParents, updateChildren)
}
func (aa *Audio) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: aa.Call("worldToLocal", vector)}
}

// AudioBuffer extend: []
type AudioBuffer struct {
	js.Value
}

func NewAudioBuffer(context js.Value) *AudioBuffer {
	return &AudioBuffer{Value: get("AudioBuffer").New(context)}
}
func (ab *AudioBuffer) JSValue() js.Value {
	return ab.Value
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
func (ab *AudioBuffer) ReadyCallbacks() js.Value {
	return ab.Get("readyCallbacks")
}
func (ab *AudioBuffer) SetReadyCallbacks(v js.Value) {
	ab.Set("readyCallbacks", v)
}
func (ab *AudioBuffer) Load(file string) *AudioBuffer {
	return &AudioBuffer{Value: ab.Call("load", file)}
}
func (ab *AudioBuffer) OnReady(callback js.Value) {
	ab.Call("onReady", callback)
}

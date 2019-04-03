// Code generated from "audio/PositionalAudio.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

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
func (pa *PositionalAudio) Buffer() js.Value {
	return pa.Get("buffer")
}
func (pa *PositionalAudio) SetBuffer(v js.Value) {
	pa.Set("buffer", v)
}
func (pa *PositionalAudio) CastShadow() bool {
	return pa.Get("castShadow").Bool()
}
func (pa *PositionalAudio) SetCastShadow(v bool) {
	pa.Set("castShadow", v)
}
func (pa *PositionalAudio) Children() js.Value {
	return pa.Get("children")
}
func (pa *PositionalAudio) SetChildren(v js.Value) {
	pa.Set("children", v)
}
func (pa *PositionalAudio) Context() js.Value {
	return pa.Get("context")
}
func (pa *PositionalAudio) SetContext(v js.Value) {
	pa.Set("context", v)
}
func (pa *PositionalAudio) Detune() float64 {
	return pa.Get("detune").Float()
}
func (pa *PositionalAudio) SetDetune(v float64) {
	pa.Set("detune", v)
}
func (pa *PositionalAudio) Filters() js.Value {
	return pa.Get("filters")
}
func (pa *PositionalAudio) SetFilters(v js.Value) {
	pa.Set("filters", v)
}
func (pa *PositionalAudio) FrustumCulled() bool {
	return pa.Get("frustumCulled").Bool()
}
func (pa *PositionalAudio) SetFrustumCulled(v bool) {
	pa.Set("frustumCulled", v)
}
func (pa *PositionalAudio) Gain() js.Value {
	return pa.Get("gain")
}
func (pa *PositionalAudio) SetGain(v js.Value) {
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
func (pa *PositionalAudio) IsObject3D() bool {
	return pa.Get("isObject3D").Bool()
}
func (pa *PositionalAudio) SetIsObject3D(v bool) {
	pa.Set("isObject3D", v)
}
func (pa *PositionalAudio) IsPlaying() bool {
	return pa.Get("isPlaying").Bool()
}
func (pa *PositionalAudio) SetIsPlaying(v bool) {
	pa.Set("isPlaying", v)
}
func (pa *PositionalAudio) Layers() *Layers {
	return &Layers{Value: pa.Get("layers")}
}
func (pa *PositionalAudio) SetLayers(v *Layers) {
	pa.Set("layers", v)
}
func (pa *PositionalAudio) Loop() bool {
	return pa.Get("loop").Bool()
}
func (pa *PositionalAudio) SetLoop(v bool) {
	pa.Set("loop", v)
}
func (pa *PositionalAudio) Matrix() *Matrix4 {
	return &Matrix4{Value: pa.Get("matrix")}
}
func (pa *PositionalAudio) SetMatrix(v *Matrix4) {
	pa.Set("matrix", v)
}
func (pa *PositionalAudio) MatrixAutoUpdate() bool {
	return pa.Get("matrixAutoUpdate").Bool()
}
func (pa *PositionalAudio) SetMatrixAutoUpdate(v bool) {
	pa.Set("matrixAutoUpdate", v)
}
func (pa *PositionalAudio) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: pa.Get("matrixWorld")}
}
func (pa *PositionalAudio) SetMatrixWorld(v *Matrix4) {
	pa.Set("matrixWorld", v)
}
func (pa *PositionalAudio) MatrixWorldNeedsUpdate() bool {
	return pa.Get("matrixWorldNeedsUpdate").Bool()
}
func (pa *PositionalAudio) SetMatrixWorldNeedsUpdate(v bool) {
	pa.Set("matrixWorldNeedsUpdate", v)
}
func (pa *PositionalAudio) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: pa.Get("modelViewMatrix")}
}
func (pa *PositionalAudio) SetModelViewMatrix(v *Matrix4) {
	pa.Set("modelViewMatrix", v)
}
func (pa *PositionalAudio) Name() string {
	return pa.Get("name").String()
}
func (pa *PositionalAudio) SetName(v string) {
	pa.Set("name", v)
}
func (pa *PositionalAudio) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: pa.Get("normalMatrix")}
}
func (pa *PositionalAudio) SetNormalMatrix(v *Matrix3) {
	pa.Set("normalMatrix", v)
}
func (pa *PositionalAudio) Offset() float64 {
	return pa.Get("offset").Float()
}
func (pa *PositionalAudio) SetOffset(v float64) {
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
func (pa *PositionalAudio) Panner() js.Value {
	return pa.Get("panner")
}
func (pa *PositionalAudio) SetPanner(v js.Value) {
	pa.Set("panner", v)
}
func (pa *PositionalAudio) Parent() *Object3D {
	return &Object3D{Value: pa.Get("parent")}
}
func (pa *PositionalAudio) SetParent(v *Object3D) {
	pa.Set("parent", v)
}
func (pa *PositionalAudio) PlaybackRate() float64 {
	return pa.Get("playbackRate").Float()
}
func (pa *PositionalAudio) SetPlaybackRate(v float64) {
	pa.Set("playbackRate", v)
}
func (pa *PositionalAudio) Position() *Vector3 {
	return &Vector3{Value: pa.Get("position")}
}
func (pa *PositionalAudio) SetPosition(v *Vector3) {
	pa.Set("position", v)
}
func (pa *PositionalAudio) Quaternion() *Quaternion {
	return &Quaternion{Value: pa.Get("quaternion")}
}
func (pa *PositionalAudio) SetQuaternion(v *Quaternion) {
	pa.Set("quaternion", v)
}
func (pa *PositionalAudio) ReceiveShadow() bool {
	return pa.Get("receiveShadow").Bool()
}
func (pa *PositionalAudio) SetReceiveShadow(v bool) {
	pa.Set("receiveShadow", v)
}
func (pa *PositionalAudio) RenderOrder() float64 {
	return pa.Get("renderOrder").Float()
}
func (pa *PositionalAudio) SetRenderOrder(v float64) {
	pa.Set("renderOrder", v)
}
func (pa *PositionalAudio) Rotation() *Euler {
	return &Euler{Value: pa.Get("rotation")}
}
func (pa *PositionalAudio) SetRotation(v *Euler) {
	pa.Set("rotation", v)
}
func (pa *PositionalAudio) Scale() *Vector3 {
	return &Vector3{Value: pa.Get("scale")}
}
func (pa *PositionalAudio) SetScale(v *Vector3) {
	pa.Set("scale", v)
}
func (pa *PositionalAudio) Source() js.Value {
	return pa.Get("source")
}
func (pa *PositionalAudio) SetSource(v js.Value) {
	pa.Set("source", v)
}
func (pa *PositionalAudio) SourceType() string {
	return pa.Get("sourceType").String()
}
func (pa *PositionalAudio) SetSourceType(v string) {
	pa.Set("sourceType", v)
}
func (pa *PositionalAudio) StartTime() float64 {
	return pa.Get("startTime").Float()
}
func (pa *PositionalAudio) SetStartTime(v float64) {
	pa.Set("startTime", v)
}
func (pa *PositionalAudio) Type() string {
	return pa.Get("type").String()
}
func (pa *PositionalAudio) SetType(v string) {
	pa.Set("type", v)
}
func (pa *PositionalAudio) Up() *Vector3 {
	return &Vector3{Value: pa.Get("up")}
}
func (pa *PositionalAudio) SetUp(v *Vector3) {
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
func (pa *PositionalAudio) DefaultUp() *Vector3 {
	return &Vector3{Value: pa.Get("DefaultUp")}
}
func (pa *PositionalAudio) SetDefaultUp(v *Vector3) {
	pa.Set("DefaultUp", v)
}
func (pa *PositionalAudio) Add(object js.Value) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("add", object)}
}
func (pa *PositionalAudio) AddEventListener(typ string, listener js.Value) {
	pa.Call("addEventListener", typ, listener)
}
func (pa *PositionalAudio) ApplyMatrix(matrix *Matrix4) {
	pa.Call("applyMatrix", matrix)
}
func (pa *PositionalAudio) ApplyQuaternion(quaternion *Quaternion) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("applyQuaternion", quaternion)}
}
func (pa *PositionalAudio) Clone(recursive bool) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("clone", recursive)}
}
func (pa *PositionalAudio) Connect() *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("connect")}
}
func (pa *PositionalAudio) Copy(source *Object3D, recursive bool) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("copy", source, recursive)}
}
func (pa *PositionalAudio) Disconnect() *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("disconnect")}
}
func (pa *PositionalAudio) DispatchEvent(event js.Value) {
	pa.Call("dispatchEvent", event)
}
func (pa *PositionalAudio) GetDetune() float64 {
	return pa.Call("getDetune").Float()
}
func (pa *PositionalAudio) GetDistanceModel() string {
	return pa.Call("getDistanceModel").String()
}
func (pa *PositionalAudio) GetFilter() js.Value {
	return pa.Call("getFilter")
}
func (pa *PositionalAudio) GetFilters() js.Value {
	return pa.Call("getFilters")
}
func (pa *PositionalAudio) GetLoop() bool {
	return pa.Call("getLoop").Bool()
}
func (pa *PositionalAudio) GetMaxDistance() float64 {
	return pa.Call("getMaxDistance").Float()
}
func (pa *PositionalAudio) GetObjectById(id int) *Object3D {
	return &Object3D{Value: pa.Call("getObjectById", id)}
}
func (pa *PositionalAudio) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: pa.Call("getObjectByName", name)}
}
func (pa *PositionalAudio) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: pa.Call("getObjectByProperty", name, value)}
}
func (pa *PositionalAudio) GetOutput() js.Value {
	return pa.Call("getOutput")
}
func (pa *PositionalAudio) GetPlaybackRate() float64 {
	return pa.Call("getPlaybackRate").Float()
}
func (pa *PositionalAudio) GetRefDistance() float64 {
	return pa.Call("getRefDistance").Float()
}
func (pa *PositionalAudio) GetRolloffFactor() float64 {
	return pa.Call("getRolloffFactor").Float()
}
func (pa *PositionalAudio) GetVolume() float64 {
	return pa.Call("getVolume").Float()
}
func (pa *PositionalAudio) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: pa.Call("getWorldDirection", target)}
}
func (pa *PositionalAudio) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: pa.Call("getWorldPosition", target)}
}
func (pa *PositionalAudio) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: pa.Call("getWorldQuaternion", target)}
}
func (pa *PositionalAudio) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: pa.Call("getWorldScale", target)}
}
func (pa *PositionalAudio) HasEventListener(typ string, listener js.Value) bool {
	return pa.Call("hasEventListener", typ, listener).Bool()
}
func (pa *PositionalAudio) Load(file string) *Audio {
	return &Audio{Value: pa.Call("load", file)}
}
func (pa *PositionalAudio) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: pa.Call("localToWorld", vector)}
}
func (pa *PositionalAudio) LookAt(vector *Vector3, y float64, z float64) {
	pa.Call("lookAt", vector, y, z)
}
func (pa *PositionalAudio) OnEnded() {
	pa.Call("onEnded")
}
func (pa *PositionalAudio) Pause() *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("pause")}
}
func (pa *PositionalAudio) Play() *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("play")}
}
func (pa *PositionalAudio) Raycast(raycaster *Raycaster, intersects js.Value) {
	pa.Call("raycast", raycaster, intersects)
}
func (pa *PositionalAudio) Remove(object js.Value) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("remove", object)}
}
func (pa *PositionalAudio) RemoveEventListener(typ string, listener js.Value) {
	pa.Call("removeEventListener", typ, listener)
}
func (pa *PositionalAudio) RotateOnAxis(axis *Vector3, angle float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("rotateOnAxis", axis, angle)}
}
func (pa *PositionalAudio) RotateOnWorldAxis(axis *Vector3, angle float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("rotateOnWorldAxis", axis, angle)}
}
func (pa *PositionalAudio) RotateX(angle float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("rotateX", angle)}
}
func (pa *PositionalAudio) RotateY(angle float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("rotateY", angle)}
}
func (pa *PositionalAudio) RotateZ(angle float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("rotateZ", angle)}
}
func (pa *PositionalAudio) SetBuffer2(audioBuffer *AudioBuffer) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setBuffer", audioBuffer)}
}
func (pa *PositionalAudio) SetDetune2(value float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setDetune", value)}
}
func (pa *PositionalAudio) SetDirectionalCone(coneInnerAngle float64, coneOuterAngle float64, coneOuterGain float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setDirectionalCone", coneInnerAngle, coneOuterAngle, coneOuterGain)}
}
func (pa *PositionalAudio) SetDistanceModel(value string) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setDistanceModel", value)}
}
func (pa *PositionalAudio) SetFilter(value js.Value) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setFilter", value)}
}
func (pa *PositionalAudio) SetFilter2(filter js.Value) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setFilter", filter)}
}
func (pa *PositionalAudio) SetLoop2(value bool) {
	pa.Call("setLoop", value)
}
func (pa *PositionalAudio) SetMaxDistance(value float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setMaxDistance", value)}
}
func (pa *PositionalAudio) SetMediaElementSource(mediaElement js.Value) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setMediaElementSource", mediaElement)}
}
func (pa *PositionalAudio) SetNodeSource(audioNode js.Value) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setNodeSource", audioNode)}
}
func (pa *PositionalAudio) SetPlaybackRate2(value float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setPlaybackRate", value)}
}
func (pa *PositionalAudio) SetRefDistance(value float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setRefDistance", value)}
}
func (pa *PositionalAudio) SetRolloffFactor(value float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setRolloffFactor", value)}
}
func (pa *PositionalAudio) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	pa.Call("setRotationFromAxisAngle", axis, angle)
}
func (pa *PositionalAudio) SetRotationFromEuler(euler *Euler) {
	pa.Call("setRotationFromEuler", euler)
}
func (pa *PositionalAudio) SetRotationFromMatrix(m *Matrix4) {
	pa.Call("setRotationFromMatrix", m)
}
func (pa *PositionalAudio) SetRotationFromQuaternion(q *Quaternion) {
	pa.Call("setRotationFromQuaternion", q)
}
func (pa *PositionalAudio) SetVolume(value float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("setVolume", value)}
}
func (pa *PositionalAudio) Stop() *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("stop")}
}
func (pa *PositionalAudio) ToJSON(meta js.Value) js.Value {
	return pa.Call("toJSON", meta)
}
func (pa *PositionalAudio) TranslateOnAxis(axis *Vector3, distance float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("translateOnAxis", axis, distance)}
}
func (pa *PositionalAudio) TranslateX(distance float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("translateX", distance)}
}
func (pa *PositionalAudio) TranslateY(distance float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("translateY", distance)}
}
func (pa *PositionalAudio) TranslateZ(distance float64) *PositionalAudio {
	return &PositionalAudio{Value: pa.Call("translateZ", distance)}
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
func (pa *PositionalAudio) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: pa.Call("worldToLocal", vector)}
}

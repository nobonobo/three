package animation

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/three/core"
	"github.com/nobonobo/three/objects"
)

type AnimationClip struct {
	js.Value
}

func NewAnimationClip(name string, duration int, tracks []KeyframeTrack) *AnimationClip {
	return &AnimationClip{Value: get("AnimationClip").New(name, duration, tracks)}
}
func (ac *AnimationClip) Duration() int {
	return ac.Get("duration").Int()
}
func (ac *AnimationClip) SetDuration(v int) {
	ac.Set("duration", v)
}
func (ac *AnimationClip) Name() string {
	return ac.Get("name").String()
}
func (ac *AnimationClip) SetName(v string) {
	ac.Set("name", v)
}
func (ac *AnimationClip) Results() []js.Value {
	return []js.Value(ac.Get("results"))
}
func (ac *AnimationClip) SetResults(v []js.Value) {
	ac.Set("results", v)
}
func (ac *AnimationClip) Tracks() []KeyframeTrack {
	return []KeyframeTrack(ac.Get("tracks"))
}
func (ac *AnimationClip) SetTracks(v []KeyframeTrack) {
	ac.Set("tracks", v)
}
func (ac *AnimationClip) Uuid() string {
	return ac.Get("uuid").String()
}
func (ac *AnimationClip) SetUuid(v string) {
	ac.Set("uuid", v)
}
func (ac *AnimationClip) Optimize() *AnimationClip {
	return &AnimationClip{Value: ac.Call("optimize")}
}
func (ac *AnimationClip) ResetDuration() {
	ac.Call("resetDuration")
}
func (ac *AnimationClip) Trim() *AnimationClip {
	return &AnimationClip{Value: ac.Call("trim")}
}
func (ac *AnimationClip) CreateClipsFromMorphTargetSequences(morphTargets []core.MorphTarget, fps int, noLoop bool) []AnimationClip {
	return []AnimationClip(ac.Call("CreateClipsFromMorphTargetSequences", morphTargets, fps, noLoop))
}
func (ac *AnimationClip) CreateFromMorphTargetSequence(name string, morphTargetSequence []core.MorphTarget, fps int, noLoop bool) *AnimationClip {
	return &AnimationClip{Value: ac.Call("CreateFromMorphTargetSequence", name, morphTargetSequence, fps, noLoop)}
}
func (ac *AnimationClip) FindByName(clipArray []AnimationClip, name string) *AnimationClip {
	return &AnimationClip{Value: ac.Call("findByName", clipArray, name)}
}
func (ac *AnimationClip) Parse(json js.Value) *AnimationClip {
	return &AnimationClip{Value: ac.Call("parse", json)}
}
func (ac *AnimationClip) ParseAnimation(animation js.Value, bones []objects.Bone, nodeName string) *AnimationClip {
	return &AnimationClip{Value: ac.Call("parseAnimation", animation, bones, nodeName)}
}
func (ac *AnimationClip) ToJSON() js.Value {
	return ac.Call("toJSON")
}

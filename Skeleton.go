// Code generated from "objects/Skeleton.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type Skeleton struct {
	js.Value
}

func NewSkeleton(bones js.Value, boneInverses js.Value) *Skeleton {
	return &Skeleton{Value: get("Skeleton").New(bones, boneInverses)}
}
func (ss *Skeleton) BoneInverses() js.Value {
	return ss.Get("boneInverses")
}
func (ss *Skeleton) SetBoneInverses(v js.Value) {
	ss.Set("boneInverses", v)
}
func (ss *Skeleton) BoneMatrices() js.Value {
	return ss.Get("boneMatrices")
}
func (ss *Skeleton) SetBoneMatrices(v js.Value) {
	ss.Set("boneMatrices", v)
}
func (ss *Skeleton) BoneTexture() *DataTexture {
	return &DataTexture{Value: ss.Get("boneTexture")}
}
func (ss *Skeleton) SetBoneTexture(v *DataTexture) {
	ss.Set("boneTexture", v)
}
func (ss *Skeleton) BoneTextureHeight() float64 {
	return ss.Get("boneTextureHeight").Float()
}
func (ss *Skeleton) SetBoneTextureHeight(v float64) {
	ss.Set("boneTextureHeight", v)
}
func (ss *Skeleton) BoneTextureWidth() float64 {
	return ss.Get("boneTextureWidth").Float()
}
func (ss *Skeleton) SetBoneTextureWidth(v float64) {
	ss.Set("boneTextureWidth", v)
}
func (ss *Skeleton) Bones() js.Value {
	return ss.Get("bones")
}
func (ss *Skeleton) SetBones(v js.Value) {
	ss.Set("bones", v)
}
func (ss *Skeleton) IdentityMatrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("identityMatrix")}
}
func (ss *Skeleton) SetIdentityMatrix(v *Matrix4) {
	ss.Set("identityMatrix", v)
}
func (ss *Skeleton) UseVertexTexture() bool {
	return ss.Get("useVertexTexture").Bool()
}
func (ss *Skeleton) SetUseVertexTexture(v bool) {
	ss.Set("useVertexTexture", v)
}
func (ss *Skeleton) CalculateInverses(bone *Bone) {
	ss.Call("calculateInverses", bone)
}
func (ss *Skeleton) Clone() *Skeleton {
	return &Skeleton{Value: ss.Call("clone")}
}
func (ss *Skeleton) Pose() {
	ss.Call("pose")
}
func (ss *Skeleton) Update() {
	ss.Call("update")
}

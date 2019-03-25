package objects

import (
	"github.com/nobonobo/three/math"
	"github.com/nobonobo/three/textures"
)

type Skeleton struct {
	js.Value
}

func NewSkeleton(bones []Bone, boneInverses []math.Matrix4) *Skeleton {
	return &Skeleton{Value: get("Skeleton").New(bones, boneInverses)}
}
func (s *Skeleton) BoneInverses() []math.Matrix4 {
	return []math.Matrix4(s.Get("boneInverses"))
}
func (s *Skeleton) SetBoneInverses(v []math.Matrix4) {
	s.Set("boneInverses", v)
}
func (s *Skeleton) BoneMatrices() *Float32Array {
	return &Float32Array{Value: s.Get("boneMatrices")}
}
func (s *Skeleton) SetBoneMatrices(v *Float32Array) {
	s.Set("boneMatrices", v)
}
func (s *Skeleton) BoneTexture() *textures.DataTexture {
	return &textures.DataTexture{Value: s.Get("boneTexture")}
}
func (s *Skeleton) SetBoneTexture(v *textures.DataTexture) {
	s.Set("boneTexture", v)
}
func (s *Skeleton) BoneTextureHeight() int {
	return s.Get("boneTextureHeight").Int()
}
func (s *Skeleton) SetBoneTextureHeight(v int) {
	s.Set("boneTextureHeight", v)
}
func (s *Skeleton) BoneTextureWidth() int {
	return s.Get("boneTextureWidth").Int()
}
func (s *Skeleton) SetBoneTextureWidth(v int) {
	s.Set("boneTextureWidth", v)
}
func (s *Skeleton) Bones() []Bone {
	return []Bone(s.Get("bones"))
}
func (s *Skeleton) SetBones(v []Bone) {
	s.Set("bones", v)
}
func (s *Skeleton) IdentityMatrix() *math.Matrix4 {
	return &math.Matrix4{Value: s.Get("identityMatrix")}
}
func (s *Skeleton) SetIdentityMatrix(v *math.Matrix4) {
	s.Set("identityMatrix", v)
}
func (s *Skeleton) UseVertexTexture() bool {
	return s.Get("useVertexTexture").Bool()
}
func (s *Skeleton) SetUseVertexTexture(v bool) {
	s.Set("useVertexTexture", v)
}
func (s *Skeleton) CalculateInverses(bone *Bone) {
	s.Call("calculateInverses", bone)
}
func (s *Skeleton) Clone() this {
	return this(s.Call("clone"))
}
func (s *Skeleton) Pose() {
	s.Call("pose")
}
func (s *Skeleton) Update() {
	s.Call("update")
}

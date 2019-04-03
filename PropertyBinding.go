// Code generated from "animation/PropertyBinding.d.ts"; DO NOT EDIT.

package three

import (
	"github.com/gopherjs/gopherwasm/js"
)

type PropertyBinding struct {
	js.Value
}

func NewPropertyBinding(rootNode js.Value, path string, parsedPath js.Value) *PropertyBinding {
	return &PropertyBinding{Value: get("PropertyBinding").New(rootNode, path, parsedPath)}
}
func (pb *PropertyBinding) BindingType() js.Value {
	return pb.Get("BindingType")
}
func (pb *PropertyBinding) SetBindingType(v js.Value) {
	pb.Set("BindingType", v)
}
func (pb *PropertyBinding) GetterByBindingType() js.Value {
	return pb.Get("GetterByBindingType")
}
func (pb *PropertyBinding) SetGetterByBindingType(v js.Value) {
	pb.Set("GetterByBindingType", v)
}
func (pb *PropertyBinding) SetterByBindingTypeAndVersioning() js.Value {
	return pb.Get("SetterByBindingTypeAndVersioning")
}
func (pb *PropertyBinding) SetSetterByBindingTypeAndVersioning(v js.Value) {
	pb.Set("SetterByBindingTypeAndVersioning", v)
}
func (pb *PropertyBinding) Versioning() js.Value {
	return pb.Get("Versioning")
}
func (pb *PropertyBinding) SetVersioning(v js.Value) {
	pb.Set("Versioning", v)
}
func (pb *PropertyBinding) Node() js.Value {
	return pb.Get("node")
}
func (pb *PropertyBinding) SetNode(v js.Value) {
	pb.Set("node", v)
}
func (pb *PropertyBinding) ParsedPath() js.Value {
	return pb.Get("parsedPath")
}
func (pb *PropertyBinding) SetParsedPath(v js.Value) {
	pb.Set("parsedPath", v)
}
func (pb *PropertyBinding) Path() string {
	return pb.Get("path").String()
}
func (pb *PropertyBinding) SetPath(v string) {
	pb.Set("path", v)
}
func (pb *PropertyBinding) RootNode() js.Value {
	return pb.Get("rootNode")
}
func (pb *PropertyBinding) SetRootNode(v js.Value) {
	pb.Set("rootNode", v)
}
func (pb *PropertyBinding) Bind() {
	pb.Call("bind")
}
func (pb *PropertyBinding) GetValue(targetArray js.Value, offset int) js.Value {
	return pb.Call("getValue", targetArray, offset)
}
func (pb *PropertyBinding) SetValue2(sourceArray js.Value, offset int) {
	pb.Call("setValue", sourceArray, offset)
}
func (pb *PropertyBinding) Unbind() {
	pb.Call("unbind")
}
func (pb *PropertyBinding) Create(root js.Value, path js.Value, parsedPath js.Value) *PropertyBinding {
	return &PropertyBinding{Value: pb.Call("create", root, path, parsedPath)}
}
func (pb *PropertyBinding) FindNode(root js.Value, nodeName string) js.Value {
	return pb.Call("findNode", root, nodeName)
}
func (pb *PropertyBinding) ParseTrackName(trackName string) js.Value {
	return pb.Call("parseTrackName", trackName)
}

type Composite struct {
	js.Value
}

func NewComposite(targetGroup js.Value, path js.Value, parsedPath js.Value) *Composite {
	return &Composite{Value: get("Composite").New(targetGroup, path, parsedPath)}
}
func (cc *Composite) Bind() {
	cc.Call("bind")
}
func (cc *Composite) GetValue(array js.Value, offset int) js.Value {
	return cc.Call("getValue", array, offset)
}
func (cc *Composite) SetValue2(array js.Value, offset int) {
	cc.Call("setValue", array, offset)
}
func (cc *Composite) Unbind() {
	cc.Call("unbind")
}

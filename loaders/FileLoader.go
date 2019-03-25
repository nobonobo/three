package loaders

import (
	"github.com/gopherjs/gopherwasm/js"
)

type AnyLoader interface {
	Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value
}
type LoaderHandler interface {
	Add(regex *RegExp, loader AnyLoader)
	Get(file string) AnyLoader
}
type FileLoader struct {
	js.Value
}

func NewFileLoader(manager *LoadingManager) *FileLoader {
	return &FileLoader{Value: get("FileLoader").New(manager)}
}
func (fl *FileLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: fl.Get("manager")}
}
func (fl *FileLoader) SetManager(v *LoadingManager) {
	fl.Set("manager", v)
}
func (fl *FileLoader) MimeType() *MimeType {
	return &MimeType{Value: fl.Get("mimeType")}
}
func (fl *FileLoader) SetMimeType(v *MimeType) {
	fl.Set("mimeType", v)
}
func (fl *FileLoader) Path() string {
	return fl.Get("path").String()
}
func (fl *FileLoader) SetPath(v string) {
	fl.Set("path", v)
}
func (fl *FileLoader) ResponseType() string {
	return fl.Get("responseType").String()
}
func (fl *FileLoader) SetResponseType(v string) {
	fl.Set("responseType", v)
}
func (fl *FileLoader) WithCredentials() string {
	return fl.Get("withCredentials").String()
}
func (fl *FileLoader) SetWithCredentials(v string) {
	fl.Set("withCredentials", v)
}
func (fl *FileLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) js.Value {
	return fl.Call("load", url, onLoad, onProgress, onError)
}
func (fl *FileLoader) SetMimeType(mimeType *MimeType) *FileLoader {
	return &FileLoader{Value: fl.Call("setMimeType", mimeType)}
}
func (fl *FileLoader) SetPath(path string) *FileLoader {
	return &FileLoader{Value: fl.Call("setPath", path)}
}
func (fl *FileLoader) SetRequestHeader(value js.Value) *FileLoader {
	return &FileLoader{Value: fl.Call("setRequestHeader", value)}
}
func (fl *FileLoader) SetResponseType(responseType string) *FileLoader {
	return &FileLoader{Value: fl.Call("setResponseType", responseType)}
}
func (fl *FileLoader) SetWithCredentials(value string) *FileLoader {
	return &FileLoader{Value: fl.Call("setWithCredentials", value)}
}

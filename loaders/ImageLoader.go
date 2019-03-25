package loaders

type ImageLoader struct {
	js.Value
}

func NewImageLoader(manager *LoadingManager) *ImageLoader {
	return &ImageLoader{Value: get("ImageLoader").New(manager)}
}
func (il *ImageLoader) CrossOrigin() string {
	return il.Get("crossOrigin").String()
}
func (il *ImageLoader) SetCrossOrigin(v string) {
	il.Set("crossOrigin", v)
}
func (il *ImageLoader) Manager() *LoadingManager {
	return &LoadingManager{Value: il.Get("manager")}
}
func (il *ImageLoader) SetManager(v *LoadingManager) {
	il.Set("manager", v)
}
func (il *ImageLoader) Path() string {
	return il.Get("path").String()
}
func (il *ImageLoader) SetPath(v string) {
	il.Set("path", v)
}
func (il *ImageLoader) WithCredentials() string {
	return il.Get("withCredentials").String()
}
func (il *ImageLoader) SetWithCredentials(v string) {
	il.Set("withCredentials", v)
}
func (il *ImageLoader) Load(url string, onLoad js.Value, onProgress js.Value, onError js.Value) *HTMLImageElement {
	return &HTMLImageElement{Value: il.Call("load", url, onLoad, onProgress, onError)}
}
func (il *ImageLoader) SetCrossOrigin(crossOrigin string) *ImageLoader {
	return &ImageLoader{Value: il.Call("setCrossOrigin", crossOrigin)}
}
func (il *ImageLoader) SetPath(value string) *ImageLoader {
	return &ImageLoader{Value: il.Call("setPath", value)}
}
func (il *ImageLoader) SetWithCredentials(value string) *ImageLoader {
	return &ImageLoader{Value: il.Call("setWithCredentials", value)}
}

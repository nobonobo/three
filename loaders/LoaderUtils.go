package loaders

type LoaderUtils struct {
	js.Value
}

func (lu *LoaderUtils) DecodeText(array *TypedArray) string {
	return lu.Call("decodeText", array).String()
}
func (lu *LoaderUtils) ExtractUrlBase(url string) string {
	return lu.Call("extractUrlBase", url).String()
}

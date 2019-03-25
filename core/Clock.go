package core

type Clock struct {
	js.Value
}

func NewClock(autoStart bool) *Clock {
	return &Clock{Value: get("Clock").New(autoStart)}
}
func (c *Clock) AutoStart() bool {
	return c.Get("autoStart").Bool()
}
func (c *Clock) SetAutoStart(v bool) {
	c.Set("autoStart", v)
}
func (c *Clock) ElapsedTime() int {
	return c.Get("elapsedTime").Int()
}
func (c *Clock) SetElapsedTime(v int) {
	c.Set("elapsedTime", v)
}
func (c *Clock) OldTime() int {
	return c.Get("oldTime").Int()
}
func (c *Clock) SetOldTime(v int) {
	c.Set("oldTime", v)
}
func (c *Clock) Running() bool {
	return c.Get("running").Bool()
}
func (c *Clock) SetRunning(v bool) {
	c.Set("running", v)
}
func (c *Clock) StartTime() int {
	return c.Get("startTime").Int()
}
func (c *Clock) SetStartTime(v int) {
	c.Set("startTime", v)
}
func (c *Clock) GetDelta() int {
	return c.Call("getDelta").Int()
}
func (c *Clock) GetElapsedTime() int {
	return c.Call("getElapsedTime").Int()
}
func (c *Clock) Start() {
	c.Call("start")
}
func (c *Clock) Stop() {
	c.Call("stop")
}

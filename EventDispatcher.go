// Code generated from "core/EventDispatcher.d.ts"; DO NOT EDIT.

// +build go1.12 js

package three

import (
	"syscall/js"
)

// EventDispatcher extend: []
type EventDispatcher struct {
	js.Value
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{Value: get("EventDispatcher").New()}
}
func (ed *EventDispatcher) JSValue() js.Value {
	return ed.Value
}
func (ed *EventDispatcher) AddEventListener(typ string, listener js.Value) {
	ed.Call("addEventListener", typ, listener)
}
func (ed *EventDispatcher) DispatchEvent(event js.Value) {
	ed.Call("dispatchEvent", event)
}
func (ed *EventDispatcher) HasEventListener(typ string, listener js.Value) bool {
	return ed.Call("hasEventListener", typ, listener).Bool()
}
func (ed *EventDispatcher) RemoveEventListener(typ string, listener js.Value) {
	ed.Call("removeEventListener", typ, listener)
}

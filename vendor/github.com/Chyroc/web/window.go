// +build js,wasm

package web

import "syscall/js"

type WindowInterface interface {
	EventTarget

	RequestAnimationFrame(callback js.Callback)
	Alert(s string)
}

type implWindow struct {
	implEventTarget

	v js.Value
}

func newimplWindow() implWindow {
	v := js.Global().Get("window")
	return implWindow{newImplEventTarget(v), v}
}

func (r *implWindow) RequestAnimationFrame(callback js.Callback) {
	r.v.Call("requestAnimationFrame", callback)
}

func (r *implWindow) Alert(s string) {
	r.v.Call("alert", s)
}

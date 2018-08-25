// +build js,wasm

package web

import "syscall/js"

type Location interface {
	Reload()
}

type implLocation struct {
	v js.Value
}

func newImplLocation(v js.Value) implLocation {
	return implLocation{
		v: v,
	}
}

func (r *implLocation) Reload() {
	r.v.Call("reload")
}

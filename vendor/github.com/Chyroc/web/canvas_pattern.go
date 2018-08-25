// +build js,wasm

package web

import "syscall/js"

type CanvasPattern interface {
}

type implCanvasPattern struct {
	v js.Value
}

func newImplCanvasPattern(v js.Value) implCanvasPattern {
	return implCanvasPattern{v: v}
}

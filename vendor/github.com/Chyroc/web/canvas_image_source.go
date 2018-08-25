// +build js,wasm

package web

import "syscall/js"

// TODO: impl
type CanvasImageSource interface {
}

type implCanvasImageSource struct {
	v js.Value
}

func newImplCanvasImageSource(v js.Value) implCanvasImageSource {
	return implCanvasImageSource{v: v}
}

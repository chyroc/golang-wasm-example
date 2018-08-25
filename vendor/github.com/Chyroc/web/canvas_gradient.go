// +build js,wasm

package web

import "syscall/js"

type CanvasGradient interface {
	AddColorStop(offset float64, color string)
}

type implCanvasGradient struct {
	v js.Value
}

func newImplCanvasGradient(v js.Value) implCanvasGradient {
	return implCanvasGradient{v: v}
}

// The CanvasGradient.addColorStop() method adds a new stop, defined by an
// offset and a color, to the gradient. If the offset is not between 0 and 1, an
// INDEX_SIZE_ERR is raised, if the color can't be parsed as a CSS <color>, a
// SYNTAX_ERR is raised.
//
// offset范围0到1，表示百分比，color是颜色
func (r *implCanvasGradient) AddColorStop(offset float64, color string) {
	r.v.Call("addColorStop", offset, color)
}

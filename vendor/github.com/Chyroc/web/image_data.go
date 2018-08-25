// +build js,wasm

package web

import "syscall/js"

type ImageData interface {
	Data() js.Value
	Height() float64
	Width() float64
}

type implImageData struct {
	v js.Value
}

func newImplImageData(v js.Value) implImageData {
	return implImageData{v}
}

func NewImageData(width, height float64) ImageData {
	v := js.Global().Get("ImageData").New(width, height)
	t := newImplImageData(v)
	return &t
}

//js.Global().Get("ImageData").New("100", "100")
func (r *implImageData) Data() js.Value {
	return r.v.Get("height")
}

func (r *implImageData) Height() float64 {
	return r.v.Get("height").Float()
}

func (r *implImageData) Width() float64 {
	return r.v.Get("width").Float()
}

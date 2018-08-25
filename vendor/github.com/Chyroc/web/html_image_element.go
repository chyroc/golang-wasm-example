// +build js,wasm

package web

import "syscall/js"

type HTMLImageElement interface {
	HTMLElement
}

type implHTMLImageElement struct {
	implHTMLElement

	v js.Value
}

func newImplHTMLImageElement(v js.Value) implHTMLImageElement {
	return implHTMLImageElement{newImplHTMLElement(v), v,}
}

func newHTMLImageElement(v js.Value) HTMLImageElement {
	d := newImplHTMLImageElement(v)
	return &d
}

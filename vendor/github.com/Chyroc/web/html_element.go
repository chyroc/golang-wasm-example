// +build js,wasm

package web

import "syscall/js"

type HTMLElement interface {
	Element
}

type implHTMLElement struct {
	implElement
}

func newImplHTMLElement(v js.Value) implHTMLElement {
	return implHTMLElement{
		implElement: newImplElement(v),
	}
}

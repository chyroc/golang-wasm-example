// +build js,wasm

package web

import (
	"strings"
	"syscall/js"
)

type DocumentInterface interface {
	EventTarget

	Location() Location
	GetElementById(id string) HTMLElement
}

// https://developer.mozilla.org/en-US/docs/Web/API/Document
type implDocument struct {
	implEventTarget

	v js.Value
}

func newImplDocument() implDocument {
	v := js.Global().Get("document")
	return implDocument{
		newImplEventTarget(v),
		v,
	}
}

func (r *implDocument) Location() Location {
	t := newImplLocation(r.v.Get("location"))
	return &t
}

func (r *implDocument) GetElementById(id string) HTMLElement {
	v := r.v.Call("getElementById", id)

	switch tagName := strings.ToLower(v.Get("tagName").String()); tagName {
	case "canvas":
		return newHTMLCanvasElement(v)
	case "img":
		return newHTMLImageElement(v)
	case "select":
		return newHTMLSelectElement(v)
	default:
		t := newImplHTMLElement(v)
		return &t
	}
}

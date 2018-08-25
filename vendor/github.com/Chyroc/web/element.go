// +build js,wasm

package web

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/element
type Element interface {
	EventTarget

	// innerHTML
	SetInnerHTML(s string)
	InnerHTML() string

	SetAttribute(name string, value interface{})
}

type implElement struct {
	implEventTarget

	v js.Value
}

func newImplElement(v js.Value) implElement {
	return implElement{
		implEventTarget: newImplEventTarget(v),
		v:               v,
	}
}

func (r *implElement) SetInnerHTML(s string) {
	r.v.Set("innerHTML", s)
}

func (r *implElement) InnerHTML() string {
	return r.v.Get("innerHTML").String()
}

func (r *implElement) SetAttribute(name string, value interface{}) {
	r.v.Set(name, value)
}

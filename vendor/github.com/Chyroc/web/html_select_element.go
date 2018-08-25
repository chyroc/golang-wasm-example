// +build js,wasm

package web

import "syscall/js"

type HTMLSelectElement interface {
	HTMLElement

	// value
	Value() string
	SetValue(s string)
	// multiple
	Multiple() bool
	SetMultiple(s bool)
	// required
	Required() bool
	SetRequired(s bool)
}

func newImplHTMLSelectElement(v js.Value) implHTMLSelectElement {
	return implHTMLSelectElement{newImplHTMLElement(v), v}
}

func newHTMLSelectElement(v js.Value) HTMLSelectElement {
	d := newImplHTMLSelectElement(v)
	return &d
}

type implHTMLSelectElement struct {
	implHTMLElement

	v js.Value
}

func (r *implHTMLSelectElement) Value() string {
	return r.v.Get("value").String()
}

func (r *implHTMLSelectElement) SetValue(s string) {
	r.v.Set("value", s)
}

// Sets or retrieves the Boolean value indicating whether multiple items can be selected from a list.
func (r *implHTMLSelectElement) Multiple() bool {
	return r.v.Get("multiple").Bool()
}

func (r *implHTMLSelectElement) SetMultiple(s bool) {
	r.v.Set("multiple", s)
}

// When present, marks an element that can't be submitted without a value.
func (r *implHTMLSelectElement) Required() bool {
	return r.v.Get("required").Bool()
}

func (r *implHTMLSelectElement) SetRequired(s bool) {
	r.v.Set("required", s)
}

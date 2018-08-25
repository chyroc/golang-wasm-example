// +build js,wasm

package web

import "syscall/js"

var (
	Console  *console
	Document DocumentInterface
	Window   WindowInterface
)

func init() {
	// document
	d := newImplDocument()
	Document = &d

	// console
	Console = &console{v: js.Global().Get("console")}

	// window
	w := newimplWindow()
	Window = &w
}

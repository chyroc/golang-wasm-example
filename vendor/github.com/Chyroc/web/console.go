// +build js,wasm

package web

import (
	"fmt"
	"syscall/js"
)

type console struct {
	v js.Value
}

func (r *console) Log(msg string, a ...interface{}) {
	r.v.Call("log", fmt.Sprintf(msg+"\n", a...))
}

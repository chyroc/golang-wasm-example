package web

import "syscall/js"

func Wait() {
	js.Global().Get("document").Call("addEventListener", "resize", js.NewCallback(func(args []js.Value) {}))
	select {}
}

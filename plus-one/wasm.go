// +build js,wasm

package main

import (
	"strconv"
	"syscall/js"
)

func main() {
	var numberDoc = js.Global().Get("document").Call("getElementById", "number")
	var plus = js.Global().Get("document").Call("getElementById", "plus")
	var minus = js.Global().Get("document").Call("getElementById", "minus")
	var number int

	plus.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		println("press +")
		number++
		numberDoc.Set("innerHTML", strconv.Itoa(number))
	}))

	minus.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		number--
		numberDoc.Set("innerHTML", strconv.Itoa(number))
		println("press -")
	}))

	plus.Set("disabled", false)
	minus.Set("disabled", false)

	select {}
}

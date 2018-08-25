// +build js,wasm

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/Chyroc/web"
	"github.com/fogleman/gg"
)

func DrawEllipse(S float64) (string, error) {
	dc := gg.NewContext(int(S), int(S))
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	var buf = new(bytes.Buffer)
	err := dc.EncodePNG(buf)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func draw(imgDom web.HTMLElement, number int) {
	img, err := DrawEllipse(float64(number * 10))
	if err != nil {
		web.Console.Log(err.Error())
	} else {
		imgDom.SetAttribute("src", fmt.Sprintf("data:image/png;base64,%s", img))
	}
}

func main() {
	var numberDoc = js.Global().Get("document").Call("getElementById", "number")
	var plus = js.Global().Get("document").Call("getElementById", "plus")
	var minus = js.Global().Get("document").Call("getElementById", "minus")
	var imgDom = web.Document.GetElementById("img")
	var number = 50

	plus.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		plus.Set("disabled", true)
		minus.Set("disabled", true)

		number++
		fmt.Printf("change to %v\n", number)
		go numberDoc.Set("innerHTML", strconv.Itoa(number))
		draw(imgDom, number)

		plus.Set("disabled", false)
		minus.Set("disabled", false)
	}))

	minus.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		plus.Set("disabled", true)
		minus.Set("disabled", true)

		number--
		fmt.Printf("change to %v\n", number)
		go numberDoc.Set("innerHTML", strconv.Itoa(number))
		draw(imgDom, number)

		plus.Set("disabled", false)
		minus.Set("disabled", false)
	}))

	draw(imgDom, number)
	plus.Set("disabled", false)
	minus.Set("disabled", false)

	web.Wait()
}

// +build js,wasm

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"strings"
	"syscall/js"

	"github.com/o1egl/govatar"
	"github.com/satori/go.uuid"

	"github.com/Chyroc/web"
)

var gender = govatar.MALE

func generateAvator(name string) (string, error) {
	if name == "" {
		name = uuid.NewV4().String()
	}

	img, err := govatar.GenerateFromUsername(gender, name)
	if err != nil {
		return "", err
	}

	var buf = new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func changeGender(genderDom web.HTMLSelectElement) {
	if strings.ToLower(genderDom.Value()) == "male" {
		gender = govatar.MALE
	} else {
		gender = govatar.FEMALE
	}
}

func main() {
	generateButtonDom := web.Document.GetElementById("generate")
	imgDom := web.Document.GetElementById("img")
	genderDom := web.Document.GetElementById("gender").(web.HTMLSelectElement)

	generateButtonDom.AddEventListener("click", func(args []js.Value) {
		web.Console.Log("click")

		img, err := generateAvator("")
		if err != nil {
			web.Console.Log(err.Error())
		} else {
			imgDom.SetAttribute("src", fmt.Sprintf("data:image/png;base64,%s", img))
		}
	})
	genderDom.AddEventListener("change", func(args []js.Value) {
		fmt.Printf("change gender value to %v\n", genderDom.Value())

		changeGender(genderDom)
	})

	changeGender(genderDom)
	generateButtonDom.SetAttribute("disabled", false)
	genderDom.SetAttribute("disabled", false)
	fmt.Printf("default gender value is %v\n", genderDom.Value())

	web.Wait()
}

package main

import (
	"image"
	"log"

	"github.com/toelsiba/fopix"
	"github.com/toelsiba/fopix/imutil"
)

func main() {

	filename := "../../fonts/tom-thumb-new.json"

	var fi fopix.FontInfo
	err := fopix.ReadFileJSON(filename, &fi)
	checkError(err)

	d, err := fopix.NewDrawer(fi)
	checkError(err)

	d.SetScale(5)

	text := "Hello, World!"

	m := image.NewRGBA(d.TextBounds(text))

	d.DrawText(m, image.ZP, text)

	err = imutil.ImageSaveToPNG("hello-world.png", m)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

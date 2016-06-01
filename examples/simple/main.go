package main

import (
	"image"
	"log"

	"github.com/toelsiba/fopix"
	"github.com/toelsiba/fopix/examples/imutil"
)

func main() {

	f, err := fopix.NewFromFile("../../fonts/tom-thumb-new.json")
	if err != nil {
		log.Fatal(err)
	}
	f.Scale(5)

	const text = "Hello, World!"

	m := image.NewRGBA(f.GetTextBounds(text))

	f.DrawText(m, image.ZP, text)

	if err = imutil.ImageSaveToPNG("test.png", m); err != nil {
		log.Fatal(err)
	}
}

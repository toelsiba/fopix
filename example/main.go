package main

import (
	"bufio"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/toelsiba/fopix"
)

func main() {

	f, err := fopix.NewFromFile("../fonts/tom-thumb-new.json")
	if err != nil {
		log.Fatal(err)
	}
	f.Scale(5)

	const text = "Hello, World!"

	m := image.NewRGBA(f.GetTextBounds(text))

	f.DrawText(m, image.ZP, text)

	if err = imageSaveToPNG("test.png", m); err != nil {
		log.Fatal(err)
	}
}

func imageSaveToPNG(fileName string, i image.Image) error {

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	return png.Encode(w, i)
}

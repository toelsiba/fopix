package main

import (
	"image"
	"image/color"
	"log"

	"github.com/toelsiba/fopix"
	"github.com/toelsiba/fopix/imutil"
)

// custom font
var gopherFont = fopix.FontInfo{
	Name:        "Go font",
	Author:      "Gopher",
	Description: "something ...",
	Size:        fopix.Point{X: 6, Y: 7},
	AnchorPos:   fopix.Point{X: 0, Y: 0},
	TargetChar:  '0',
	CharSet: []fopix.RuneInfo{
		fopix.RuneInfo{
			Character: 'G',
			Bitmap: []string{
				"-000-",
				"0---0",
				"0----",
				"0-000",
				"0---0",
				"-000-",
			},
		},
		fopix.RuneInfo{
			Character: 'o',
			Bitmap: []string{
				"-----",
				"-----",
				"-000-",
				"0---0",
				"0---0",
				"-000-",
			},
		},
	},
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	d, err := fopix.NewDrawer(gopherFont)
	checkError(err)

	d.SetScale(10)
	d.SetColor(color.RGBA{0, 0, 0xFF, 0xFF})

	text := "Go"

	m := image.NewRGBA(d.TextBounds(text))

	imutil.ImageSolidFill(m, color.White)

	d.DrawText(m, image.ZP, text)

	err = imutil.ImageSaveToPNG("go-font.png", m)
	checkError(err)
}

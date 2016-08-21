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
	Size:        fopix.Size{Dx: 6, Dy: 7},
	AnchorPos:   image.Point{0, 0},
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

func main() {

	f, err := fopix.New(gopherFont)
	if err != nil {
		log.Fatal(err)
	}
	f.Scale(10)
	f.Color(color.RGBA{0, 0, 0xFF, 0xFF})

	text := "Go"

	m := image.NewRGBA(f.GetTextBounds(text))

	imutil.ImageSolidFill(m, color.White)

	f.DrawText(m, image.ZP, text)

	if err = imutil.ImageSaveToPNG("go-font.png", m); err != nil {
		log.Fatal(err)
	}
}

# fopix
Simple monospaced pixel font package for golang

Partly idea taken from: [pixfont](https://github.com/pbnjay/pixfont)

## Installation
```bash
go get github.com/toelsiba/fopix
```

## Fonts

Font files are available in the directory [fonts](fonts)

####digits 3x3
![digits-3x3](samples/images/digits-3x3.png)

####Digits 3x4
![digits-3x4](samples/images/digits-3x4.png)

####Digits 3x5
![digits-3x5](samples/images/digits-3x5.png)

####[3x3 Font for Nerds](http://cargocollective.com/slowercase/3x3-Font-for-Nerds)
![font-3x3](samples/images/font-3x3-multiline.png)

###Victor
![victor-ascii](samples/images/victor-ascii.png)

![victor-multiline](samples/images/victor-multiline.png)

####[Miniwi](https://github.com/sshbio/miniwi)
![miniwi-ascii](samples/images/miniwi-ascii.png)

![miniwi-multiline](samples/images/miniwi-multiline.png)

####[Tom Thumb](http://robey.lag.net/2010/01/23/tiny-monospace-font.html#comment-1526952840)
![tom-thumb-ascii](samples/images/tom-thumb-ascii.png)

![tom-thumb-multiline](samples/images/tom-thumb-multiline.png)

---

## Example

###use an existing font
```go
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
```

###use an custom font
```go
package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/toelsiba/fopix"
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

	draw.Draw(m, m.Bounds(), &image.Uniform{color.White}, image.ZP, draw.Src)

	f.DrawText(m, image.ZP, text)

	file, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if err = png.Encode(file, m); err != nil {
		log.Fatal(err)
	}
}
```

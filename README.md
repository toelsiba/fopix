# fopix
Simple monospaced pixel font package for golang

Partly idea taken from: [pixfont](https://github.com/pbnjay/pixfont)

## Installation
```bash
go get github.com/toelsiba/fopix
```

## Fonts

Font files are available in the directory [fonts](fonts). Fonts is saved in JSON format.

#### Digits 3x3
![digits-3x3](images/digits-3x3.png)

#### Digits 3x4
![digits-3x4](images/digits-3x4.png)

#### Digits 3x5
![digits-3x5](images/digits-3x5.png)

#### [3x3 Font for Nerds](http://cargocollective.com/slowercase/3x3-Font-for-Nerds)
![font-3x3-ascii](images/font-3x3-ascii.png)

![font-3x3-multiline](images/font-3x3-multiline.png)

#### Victor
![victor-ascii](images/victor-ascii.png)

![victor-multiline](images/victor-multiline.png)

#### [Miniwi](https://github.com/sshbio/miniwi)
![miniwi-ascii](images/miniwi-ascii.png)

![miniwi-multiline](images/miniwi-multiline.png)

#### [Tom Thumb](http://robey.lag.net/2010/01/23/tiny-monospace-font.html#comment-1526952840)
![tom-thumb-ascii](images/tom-thumb-ascii.png)

![tom-thumb-multiline](images/tom-thumb-multiline.png)

#### Tom Thumb New
![tom-thumb-new-ascii](images/tom-thumb-new-ascii.png)

![tom-thumb-new-multiline](images/tom-thumb-new-multiline.png)

#### Pixefon
![pixefon-ascii](images/pixefon-4x5-ascii.png)

![pixefon-multiline](images/pixefon-4x5-multiline.png)

---

## Examples

### using existing font
```go
package main

import (
	"image"
	"log"

	"github.com/toelsiba/fopix"
	"github.com/toelsiba/fopix/imutil"
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

	if err = imutil.ImageSaveToPNG("hello-world.png", m); err != nil {
		log.Fatal(err)
	}
}
```
#### Result image
![hello-world](images/hello-world.png)


### using custom font
```go
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
```
#### Result image
![go-font](images/go-font.png)

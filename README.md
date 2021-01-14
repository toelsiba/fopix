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

#### [VGA CP437](https://en.wikipedia.org/wiki/Code_page_437)
![cp437-table](images/cp437-table.png)

![cp437-text](images/cp437-text.png)

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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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
```
#### Result image
![go-font](images/go-font.png)

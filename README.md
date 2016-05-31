# fopix
Simple monospaced pixel font package for golang

Partly idea taken from: [pixfont](https://github.com/pbnjay/pixfont)

## Installation
As usual:
```bash
go get github.com/toelsiba/fopix
```

## Fonts

Font files are available in the [fonts](fonts)

- [Tom Thumb] (http://robey.lag.net/2010/01/23/tiny-monospace-font.html#comment-1526952840)
- 3x3 Font for Nerds [3x3](http://cargocollective.com/slowercase/3x3-Font-for-Nerds)
- Victor monospaced
- [Miniwi](https://github.com/sshbio/miniwi)
and other

## Example

```go
package main

import (
	"bufio"
	"image"
	"image/color"
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
	f.Color(color.NRGBA{0, 0, 255, 255})

	const text = "Hello, World!"

	m := image.NewRGBA(f.GetTextBounds(text))
	f.DrawText(m, image.Point{0, 0}, text)

	err = imageSaveToPNG("test.png", m)
	if err != nil {
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

package imutil

import (
	"bufio"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func ImageSaveToPNG(fileName string, i image.Image) error {

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	return png.Encode(w, i)
}

func ImageSolidFill(m draw.Image, c color.Color) {
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
}

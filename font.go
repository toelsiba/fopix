package fopix

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"unicode/utf8"

	"github.com/toelsiba/fopix/bitmap"
)

type ColorSetter interface {
	Set(x, y int, c color.Color)
}

type Font struct {
	size  Size
	scale int
	c     color.Color
	m     map[rune]*bitmap.Bitmap
}

func New(fi FontInfo) (*Font, error) {
	m, err := newMapFromFontInfo(fi)
	if err != nil {
		return nil, err
	}
	return &Font{
		size:  fi.Size,
		scale: 1,
		c:     color.Black,
		m:     m,
	}, nil
}

func NewFromFile(fileName string) (*Font, error) {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var fi FontInfo

	err = json.Unmarshal(data, &fi)
	if err != nil {
		return nil, err
	}

	return New(fi)
}

func newMapFromFontInfo(fi FontInfo) (map[rune]*bitmap.Bitmap, error) {
	m := make(map[rune]*bitmap.Bitmap)
	for _, ri := range fi.CharSet {
		r := rune(ri.Character)
		if _, ok := m[r]; ok {
			return nil, fmt.Errorf("duplicate rune '%c'", r)
		}
		bm, err := newBitmapFromLines(ri.Bitmap, rune(fi.TargetChar), fi.AnchorPos, fi.Size)
		if err != nil {
			return nil, err
		}
		m[r] = bm
	}
	return m, nil
}

func newBitmapFromLines(lines []string, target rune, pos image.Point, size Size) (*bitmap.Bitmap, error) {
	var (
		nX = size.Dx
		nY = size.Dy
	)
	bm, err := bitmap.New(nX * nY)
	if err != nil {
		return nil, err
	}
	for iY := 0; iY < nY; iY++ {
		if iY >= len(lines) {
			break
		}
		data := []byte(lines[iY])
		for iX := 0; iX < nX; iX++ {
			r, size := utf8.DecodeRune(data)
			if size == 0 {
				break
			}
			data = data[size:]
			if r == target {
				var (
					x = pos.X + iX
					y = pos.Y + iY
				)
				if (x >= 0) && (x < nX) {
					if (y >= 0) && (y < nY) {
						bm.Set(y*nX+x, true)
					}
				}
			}
		}
	}
	return bm, nil
}

func (f *Font) Scale(scale int) {
	if scale > 0 {
		f.scale = scale
	}
}

func (f *Font) GetScale() int {
	return f.scale
}

func (f *Font) Color(c color.Color) {
	f.c = c
}

func (f *Font) GetColor() color.Color {
	return f.c
}

func (f *Font) GetTextBounds(text string) image.Rectangle {

	data := []byte(text)

	maxY := 1
	var x, maxX int

	for {
		r, size := utf8.DecodeRune(data)
		if size == 0 {
			break
		}
		data = data[size:]

		if r == '\n' {
			x = 0
			maxY++
			continue
		}
		if r == '\t' {
			x += 4
			if maxX < x {
				maxX = x
			}
			continue
		}
		x++
		if maxX < x {
			maxX = x
		}
	}

	var (
		x1 = f.size.Dx * f.scale * maxX
		y1 = f.size.Dy * f.scale * maxY
	)

	return image.Rect(0, 0, x1, y1)
}

func (f *Font) DrawText(cs ColorSetter, pos image.Point, text string) {

	data := []byte(text)

	sizeX := f.size.Dx * f.scale
	sizeY := f.size.Dy * f.scale

	x := 0
	y := 0

	for {
		r, size := utf8.DecodeRune(data)
		if size == 0 {
			break
		}
		data = data[size:]

		if r == '\n' {
			x = 0
			y++
			continue
		}
		if r == '\t' {
			x += 4
			continue
		}

		p := image.Point{
			X: pos.X + x*sizeX,
			Y: pos.Y + y*sizeY,
		}

		f.DrawRune(cs, p, r)

		x++
	}
}

func (f *Font) DrawRune(cs ColorSetter, pos image.Point, r rune) {
	bm, ok := f.m[r]
	if !ok {
		return
	}

	if m, ok := cs.(*image.RGBA); ok {
		f.drawBitmapRGBA(m, pos, bm)
	} else {
		f.drawBitmap(cs, pos, bm)
	}
}

func (f *Font) drawBitmap(cs ColorSetter, pos image.Point, bm *bitmap.Bitmap) {

	var (
		nX = f.size.Dx
		nY = f.size.Dy
	)

	if f.scale == 1 {
		y := pos.Y
		for iY := 0; iY < nY; iY++ {
			x := pos.X
			for iX := 0; iX < nX; iX++ {
				if bit, _ := bm.Get(iY*nX + iX); bit {
					cs.Set(x, y, f.c)
				}
				x++
			}
			y++
		}
	} else if f.scale > 1 {
		y := pos.Y
		for iY := 0; iY < nY; iY++ {
			x := pos.X
			for iX := 0; iX < nX; iX++ {
				if bit, _ := bm.Get(iY*nX + iX); bit {
					fillRect(cs, pixelRect(x, y, f.scale), f.c)
				}
				x += f.scale
			}
			y += f.scale
		}
	}
}

func pixelRect(x, y int, scale int) image.Rectangle {
	return image.Rect(x, y, x+scale, y+scale)
}

func fillRect(cs ColorSetter, r image.Rectangle, c color.Color) {
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			cs.Set(x, y, c)
		}
	}
}

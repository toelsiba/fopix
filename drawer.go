package fopix

import (
	"fmt"
	"image"
	"image/color"
	"unicode/utf8"

	"github.com/toelsiba/fopix/bitmap"
)

type ColorSetter interface {
	Set(x, y int, c color.Color)
}

type Drawer struct {
	size  Point
	scale int
	c     color.Color
	m     map[rune]*bitmap.Bitmap
}

func NewDrawer(fi FontInfo) (*Drawer, error) {
	m, err := newMapFromFontInfo(fi)
	if err != nil {
		return nil, err
	}
	return &Drawer{
		size:  fi.Size,
		scale: 1,
		c:     color.Black,
		m:     m,
	}, nil
}

func newMapFromFontInfo(fi FontInfo) (map[rune]*bitmap.Bitmap, error) {
	m := make(map[rune]*bitmap.Bitmap)
	for _, ri := range fi.CharSet {
		r := rune(ri.Character)
		if _, ok := m[r]; ok {
			return nil, fmt.Errorf("there is duplicate of rune '%c'", r)
		}
		bm, err := newBitmapFromLines(ri.Bitmap, rune(fi.TargetChar), fi.AnchorPos, fi.Size)
		if err != nil {
			return nil, err
		}
		m[r] = bm
	}
	return m, nil
}

func newBitmapFromLines(lines []string, target rune, pos Point, size Point) (*bitmap.Bitmap, error) {
	var (
		nX = size.X
		nY = size.Y
	)
	bm := bitmap.New(nX * nY)
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
						bm.SetBit(y*nX+x, 1)
					}
				}
			}
		}
	}
	return bm, nil
}

func (d *Drawer) SetScale(scale int) {
	if scale > 0 {
		d.scale = scale
	}
}

func (d *Drawer) Scale() int {
	return d.scale
}

func (d *Drawer) SetColor(c color.Color) {
	d.c = c
}

func (d *Drawer) Color() color.Color {
	return d.c
}

func (d *Drawer) RuneBounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{
			X: d.size.X * d.scale,
			Y: d.size.Y * d.scale,
		},
	}
}

func (d *Drawer) TextBounds(text string) image.Rectangle {

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
		x1 = d.size.X * d.scale * maxX
		y1 = d.size.Y * d.scale * maxY
	)

	return image.Rect(0, 0, x1, y1)
}

func (d *Drawer) DrawRune(cs ColorSetter, pos image.Point, r rune) {

	bm, ok := d.m[r]
	if !ok {
		return
	}

	if m, ok := cs.(*image.RGBA); ok {
		d.drawBitmapRGBA(m, pos, bm)
	} else {
		d.drawBitmap(cs, pos, bm)
	}
}

func (d *Drawer) DrawText(cs ColorSetter, pos image.Point, text string) {

	data := []byte(text)

	sizeX := d.size.X * d.scale
	sizeY := d.size.Y * d.scale

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

		d.DrawRune(cs, p, r)

		x++
	}
}

func (d *Drawer) drawBitmap(cs ColorSetter, pos image.Point, bm *bitmap.Bitmap) {

	var (
		nX = d.size.X
		nY = d.size.Y
	)

	if d.scale == 1 {
		y := pos.Y
		for iY := 0; iY < nY; iY++ {
			x := pos.X
			for iX := 0; iX < nX; iX++ {
				if bit, _ := bm.GetBit(iY*nX + iX); bit == 1 {
					cs.Set(x, y, d.c)
				}
				x++
			}
			y++
		}
	} else if d.scale > 1 {
		y := pos.Y
		for iY := 0; iY < nY; iY++ {
			x := pos.X
			for iX := 0; iX < nX; iX++ {
				if bit, _ := bm.GetBit(iY*nX + iX); bit == 1 {
					fillRect(cs, pixelRect(x, y, d.scale), d.c)
				}
				x += d.scale
			}
			y += d.scale
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

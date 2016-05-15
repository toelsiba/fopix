package fopix

import (
	"image"
	"image/color"

	"github.com/toelsiba/fopix/bitmap"
)

func (f *Font) drawBitmapRGBA(m *image.RGBA, pos image.Point, bm *bitmap.Bitmap) {

	q := color.RGBAModel.Convert(f.c).(color.RGBA)
	colorData := []byte{q.R, q.G, q.B, q.A}

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
					m.Set(x, y, f.c)
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
					fillRectRGBA(m, pixelRect(x, y, f.scale), colorData)
				}
				x += f.scale
			}
			y += f.scale
		}
	}
}

func fillRectRGBA(m *image.RGBA, r image.Rectangle, colorData []byte) {

	r = m.Bounds().Intersect(r)
	if r.Empty() {
		return
	}

	for y := r.Min.Y; y < r.Max.Y; y++ {
		data := m.Pix[m.PixOffset(r.Min.X, y):]
		for x := r.Min.X; x < r.Max.X; x++ {
			copy(data, colorData)
			data = data[4:]
		}
	}
}

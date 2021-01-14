package fopix

import (
	"image"
	"image/color"

	"github.com/toelsiba/fopix/bitmap"
)

func (d *Drawer) drawBitmapRGBA(m *image.RGBA, pos image.Point, bm *bitmap.Bitmap) {

	q := color.RGBAModel.Convert(d.c).(color.RGBA)
	colorData := []byte{q.R, q.G, q.B, q.A}

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
					m.Set(x, y, d.c)
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
					fillRectRGBA(m, pixelRect(x, y, d.scale), colorData)
				}
				x += d.scale
			}
			y += d.scale
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

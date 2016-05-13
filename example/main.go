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

const (
	font3x3         = "../fonts/font3x3.json"
	fontVictor      = "../fonts/victor.json"
	fontMiniwi      = "../fonts/miniwi.json"
	fontTomThumb    = "../fonts/tom-thumb.json"
	fontTomThumbNew = "../fonts/tom-thumb-new.json"
	fontDigits3x3   = "../fonts/digits3x3.json"
	fontDigits3x4   = "../fonts/digits3x4.json"
	fontDigits3x5   = "../fonts/digits3x5.json"
)

var textMultiline = `During the 20th century, the field of professional astronomy
split into observational and theoretical branches. Observational astronomy
is focused on acquiring data from observations of astronomical objects, which
is then analyzed using basic principles of physics. Theoretical astronomy is
oriented toward the development of computer or analytical models to describe
astronomical objects and phenomena. The two fields complement each other, with
theoretical astronomy seeking to explain the observational results and
observations being used to confirm theoretical results.
Astronomy is one of the few sciences where amateurs can still play an active
role, especially in the discovery and observation of transient phenomena.
Amateur astronomers have made and contributed to many important astronomical
discoveries.`

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

func fillImage(m *image.RGBA, c color.Color) {
	var (
		data = m.Pix
		q    = color.RGBAModel.Convert(c).(color.RGBA)
		temp = []uint8{q.R, q.G, q.B, q.A}
	)
	for len(data) >= 4 {
		copy(data, temp)
		data = data[4:]
	}
}

func main() {
	f, err := fopix.NewFromFile(fontTomThumbNew)
	if err != nil {
		log.Fatal(err)
	}

	f.Scale(2)
	f.Color(color.RGBA{0, 0, 0xFF, 0xFF})

	//text := "abcdefghijklmnopqrstuvwxyz"
	//text := "`1234567890-=[]\\;',./"
	//text := "~!@#$%^&*()_+{}|:\"<>?"
	text := textMultiline

	bounds := f.GetTextBounds(text)
	if bounds.Empty() {
		log.Fatal("bounds is empty")
	}
	m := image.NewRGBA(bounds)

	fillImage(m, color.Black)
	//fillImage(m, color.White)

	f.DrawText(m, image.Point{0, 0}, text)

	err = imageSaveToPNG("test.png", m)
	if err != nil {
		log.Fatal(err)
	}
}

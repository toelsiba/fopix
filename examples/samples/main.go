package main

import (
	"errors"
	"image"
	"image/color"
	"log"
	"os"
	"path/filepath"

	"github.com/toelsiba/fopix"
	"github.com/toelsiba/fopix/examples/imutil"
)

type Sample struct {
	FontFile  string
	Text      string
	ImageFile string
	Scale     int
}

func makeSample(s Sample, outDir string) error {

	f, err := fopix.NewFromFile(s.FontFile)
	if err != nil {
		return err
	}

	f.Scale(s.Scale)
	f.Color(color.RGBA{0, 0, 0xFF, 0xFF})

	bounds := f.GetTextBounds(s.Text)
	if bounds.Empty() {
		return errors.New("bounds is empty")
	}
	m := image.NewRGBA(bounds)

	imutil.ImageSolidFill(m, color.RGBA{0x24, 0x24, 0x24, 0xFF})

	f.DrawText(m, image.ZP, s.Text)

	return imutil.ImageSaveToPNG(filepath.Join(outDir, s.ImageFile), m)
}

func drawableASCII() string {
	var data []byte
	for y := 1; y < 4; y++ {
		if y > 1 {
			data = append(data, '\n')
		}
		for x := 0; x < 32; x++ {
			data = append(data, byte(y*32+x))
		}
	}
	return string(data)
}

func main() {

	dirFonts := "../../fonts"

	var (
		fontFile3x3         = filepath.Join(dirFonts, "font-3x3.json")
		fontFileVictor      = filepath.Join(dirFonts, "victor.json")
		fontFileMiniwi      = filepath.Join(dirFonts, "miniwi.json")
		fontFileTomThumb    = filepath.Join(dirFonts, "tom-thumb.json")
		fontFileTomThumbNew = filepath.Join(dirFonts, "tom-thumb-new.json")
		fontFileDigits3x3   = filepath.Join(dirFonts, "digits-3x3.json")
		fontFileDigits3x4   = filepath.Join(dirFonts, "digits-3x4.json")
		fontFileDigits3x5   = filepath.Join(dirFonts, "digits-3x5.json")
		fontFilePixefon4x5  = filepath.Join(dirFonts, "pixefon-4x5.json")
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

	var textASCII = drawableASCII()

	samples := []Sample{
		Sample{
			FontFile:  fontFileDigits3x3,
			Text:      "0123456789",
			ImageFile: "digits-3x3.png",
			Scale:     10,
		},
		Sample{
			FontFile:  fontFileDigits3x4,
			Text:      "0123456789",
			ImageFile: "digits-3x4.png",
			Scale:     10,
		},
		Sample{
			FontFile:  fontFileDigits3x5,
			Text:      "0123456789",
			ImageFile: "digits-3x5.png",
			Scale:     10,
		},
		Sample{
			FontFile:  fontFile3x3,
			Text:      textASCII,
			ImageFile: "font-3x3-ascii.png",
			Scale:     5,
		},
		Sample{
			FontFile:  fontFile3x3,
			Text:      textMultiline,
			ImageFile: "font-3x3-multiline.png",
			Scale:     2,
		},
		Sample{
			FontFile:  fontFileTomThumb,
			Text:      textASCII,
			ImageFile: "tom-thumb-ascii.png",
			Scale:     5,
		},
		Sample{
			FontFile:  fontFileTomThumb,
			Text:      textMultiline,
			ImageFile: "tom-thumb-multiline.png",
			Scale:     2,
		},
		Sample{
			FontFile:  fontFileTomThumbNew,
			Text:      textASCII,
			ImageFile: "tom-thumb-new-ascii.png",
			Scale:     5,
		},
		Sample{
			FontFile:  fontFileTomThumbNew,
			Text:      textMultiline,
			ImageFile: "tom-thumb-new-multiline.png",
			Scale:     2,
		},
		Sample{
			FontFile:  fontFileVictor,
			Text:      textASCII,
			ImageFile: "victor-ascii.png",
			Scale:     3,
		},
		Sample{
			FontFile:  fontFileVictor,
			Text:      textMultiline,
			ImageFile: "victor-multiline.png",
			Scale:     1,
		},
		Sample{
			FontFile:  fontFileMiniwi,
			Text:      textASCII,
			ImageFile: "miniwi-ascii.png",
			Scale:     5,
		},
		Sample{
			FontFile:  fontFileMiniwi,
			Text:      textMultiline,
			ImageFile: "miniwi-multiline.png",
			Scale:     2,
		},
		Sample{
			FontFile:  fontFilePixefon4x5,
			Text:      textASCII,
			ImageFile: "pixefon-4x5-ascii.png",
			Scale:     5,
		},
		Sample{
			FontFile:  fontFilePixefon4x5,
			Text:      textMultiline,
			ImageFile: "pixefon-4x5-multiline.png",
			Scale:     2,
		},
	}

	const dirImages = "images"
	os.Mkdir(dirImages, os.ModePerm)

	for _, s := range samples {
		if err := makeSample(s, dirImages); err != nil {
			log.Fatal(err)
		}
	}
}

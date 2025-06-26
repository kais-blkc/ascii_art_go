package imageutils

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type ColorsBg struct {
	black image.Image
	white image.Image
	red   image.Image
	green image.Image
	blue  image.Image
}

func AsciiToImage(asciiArt []string, outputFile string) error {
	fontface := basicfont.Face7x13

	charWidth := 7
	charHeight := 13

	imgWidth := len(asciiArt[0]) * charWidth
	imgHeight := len(asciiArt) * charHeight

	myBgColors := &ColorsBg{
		black: &image.Uniform{color.Black},
		white: &image.Uniform{color.White},
		red:   &image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 255}},
		green: &image.Uniform{color.RGBA{R: 0, G: 255, B: 0, A: 255}},
		blue:  &image.Uniform{color.RGBA{R: 0, G: 0, B: 255, A: 255}},
	}

	colorBg := myBgColors.black
	colorDraw := myBgColors.green

	rgba := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(rgba, rgba.Bounds(), colorBg, image.Point{}, draw.Src)

	d := &font.Drawer{
		Dst:  rgba,
		Src:  colorDraw,
		Face: fontface,
	}

	for i, line := range asciiArt {
		d.Dot = fixed.Point26_6{
			X: fixed.I(0),
			Y: fixed.I((i + 1) * charHeight),
		}
		d.DrawString(line)
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, rgba, &jpeg.Options{Quality: 90})
	if err != nil {
		return err
	}

	log.Printf("Image saved to %s", outputFile)
	return nil
}

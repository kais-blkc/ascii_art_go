package imageutils

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func AsciiToImageRGB(img image.Image, outFile string, asciiRamp string) error {
	fontface := basicfont.Face7x13
	charWidth := 7
	charHeight := 13

	cols := img.Bounds().Dx() / charWidth
	rows := img.Bounds().Dy() / charHeight

	newImg := image.NewRGBA(image.Rect(0, 0, cols*charWidth, rows*charHeight))
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.Black}, image.Point{}, draw.Src)

	drawer := &font.Drawer{
		Dst:  newImg,
		Face: fontface,
	}

	for y := range rows {
		for x := range cols {
			px := x * charWidth
			py := y * charHeight

			r, g, b, _ := img.At(px, py).RGBA()
			gray := PixelToGray(r, g, b, float64(1.9))
			char := GrayToAscii(gray, asciiRamp)

			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
			drawer.Src = image.NewUniform(color.RGBA{r8, g8, b8, 255})

			drawer.Dot = fixed.Point26_6{
				X: fixed.I(x * charWidth),
				Y: fixed.I(y * charHeight),
			}
			drawer.DrawString(char)
		}
	}

	file, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, newImg, nil)
	if err != nil {
		return err
	}

	return nil
}

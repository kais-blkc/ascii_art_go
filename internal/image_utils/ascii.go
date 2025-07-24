package imageutils

import (
	"image"
	"strings"

	"github.com/kais-blkc/ascii_art/internal/shared/constants"
)

func ImageToASCII(img image.Image) string {
	bounds := img.Bounds()
	var sb strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := PixelToGray(r, g, b, float64(1.0))
			sb.WriteString(GrayToAscii(gray, constants.AsciiRampDefault))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

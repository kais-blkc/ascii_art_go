package imageutils

import "github.com/kais-blkc/ascii_art/internal/shared/helpers"

// var asciiRamp = " .:-=+*#%@" // Чем левее, тем темнее символ
var localAsciiRamp = "░▒▓█" // Чем левее, тем темнее символ

func PixelToGray(r, g, b uint32, factor float64) uint8 {
	// RGB -> Y по формуле NTSC (весовые коэффициенты)
	r8, g8, b8 := float64(r>>8)*factor, float64(g>>8)*factor, float64(b>>8)*factor
	gray := 0.299*r8 + 0.587*g8 + 0.114*b8

	// return uint8(gray)
	return uint8(helpers.Clamp(gray, 0, 255))
}

func GrayToAscii(gray uint8, asciiRamp string) string {
	scale := float64(gray) / 255.0
	index := int(scale * float64(len(asciiRamp)-1))
	return string(asciiRamp[index])
}

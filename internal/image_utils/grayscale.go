package imageutils

func PixelToGray(r, g, b uint32) uint8 {
	// RGB -> Y по формуле NTSC (весовые коэффициенты)
	r8, g8, b8 := float64(r>>8), float64(g>>8), float64(b>>8)
	gray := 0.299*r8 + 0.587*g8 + 0.114*b8

	return uint8(gray)
}

var asciiRamp = "@%#*+=-:. " // Чем левее, тем темнее символ

func GrayToAscii(gray uint8) string {
	scale := float64(gray) / 255.0
	index := int(scale * float64(len(asciiRamp)-1))
	return string(asciiRamp[index])
}

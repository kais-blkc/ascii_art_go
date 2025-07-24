package convert

import (
	"image"

	"github.com/nfnt/resize"
)

func ResizeImage(img image.Image, width uint, forTerminal bool) image.Image {
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	aspectRatio := float64(originalHeight) / float64(originalWidth)
	var newHeight uint

	if forTerminal {
		ratio := 0.5
		adjustedAspectRatio := aspectRatio * ratio
		newHeight = uint(float64(width) * adjustedAspectRatio)
	} else {
		newHeight = uint(float64(width) * aspectRatio)
	}

	return resize.Resize(width, newHeight, img, resize.Lanczos3)
}

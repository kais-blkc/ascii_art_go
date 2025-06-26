package convert

import (
	"image"

	"github.com/nfnt/resize"
)

func ResizeImage(img image.Image, width uint) image.Image {
	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()

	ratio := 0.5
	aspectRatio := float64(originalHeight) / float64(originalWidth)
	adjustedAspectRatio := aspectRatio * ratio
	newHeight := uint(float64(width) * adjustedAspectRatio)

	return resize.Resize(width, newHeight, img, resize.Lanczos3)
}

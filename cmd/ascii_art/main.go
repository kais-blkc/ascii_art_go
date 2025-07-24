package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kais-blkc/ascii_art/internal/convert"
	imageutils "github.com/kais-blkc/ascii_art/internal/image_utils"
	"github.com/kais-blkc/ascii_art/internal/shared/constants"
	"github.com/kais-blkc/ascii_art/internal/ui"
)

func getWidth() uint {
	width := os.Args[2]

	widthInt, err := strconv.Atoi(width)
	if err != nil {
		log.Println("Invalid width argument:", width)
		return uint(80)
	}

	return uint(widthInt)
}

func convertImgToAscii() {
	width := getWidth()
	path := os.Args[1]
	img := convert.LoadImage(path)
	fmt.Printf("Image original size: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

	img = convert.ResizeImage(img, width, false)
	fmt.Printf("Image resized size: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

	err := imageutils.AsciiToImageRGB(img, "output.jpg", constants.AsciiRampDefault)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}
}

func main() {
	if len(os.Args) >= 2 {
		convertImgToAscii()
		return
	}

	ui.StartUI()
}

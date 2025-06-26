package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kais-blkc/ascii_art/internal/convert"
	imageutils "github.com/kais-blkc/ascii_art/internal/image_utils"
)

func getWidth() uint {
	if os.Args[2] != "" {
		width, err := strconv.Atoi(os.Args[2])

		if err == nil {
			return uint(width)
		}
	}

	return uint(80)
}

func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: go run main.go <image_path> [width]")
		return
	}

	width := getWidth()
	path := os.Args[1]
	img := convert.LoadImage(path)
	fmt.Printf("Image original size: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

	img = convert.ResizeImage(img, width)
	fmt.Printf("Image resized size: %dx%d\n", img.Bounds().Dx(), img.Bounds().Dy())

	ascii := imageutils.ImageToASCII(img)
	fmt.Println(ascii)
}

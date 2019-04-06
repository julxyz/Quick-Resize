package main

import (
	"fmt"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	// open image from origin
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// filetype filter

	// decoders

	// decode jpg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// gui

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(100, 0, img, resize.Lanczos3)

	newFileName := fmt.Sprintf("%s resized.png", os.Args[2])
	newFileName = strings.Replace(newFileName, ".png", "", 1)
	out, err := os.Create(newFileName) // change to var
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

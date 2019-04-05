package main

import (
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// open "test.jpg"
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// decode jpg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(100, 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

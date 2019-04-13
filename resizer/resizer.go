package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func resizepng(file *os.File) {

	// decode png into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	width, height, saveaspng := sizeDialog(img, "png")

	if saveaspng {
	} else {
	}

	m := resize.Resize(width, height, img, resize.NearestNeighbor)

	newFileName := fmt.Sprintf("%s resized.png", strings.Replace(os.Args[2], ".png", "", -1))
	out, err := os.Create(newFileName) // change to var
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}

func resizejpg(file *os.File) {

	// decode jpg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	width, height, saveaspng := sizeDialog(img, "jpg")

	m := resize.Resize(width, height, img, resize.NearestNeighbor)

	if saveaspng {
		newFileName := fmt.Sprintf("%s resized.png", strings.Replace(os.Args[2], ".jpg", "", -1))
		out, err := os.Create(newFileName) // change to var
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		png.Encode(out, m)

	} else {
		newFileName := fmt.Sprintf("%s resized.jpg", strings.Replace(os.Args[2], ".jpg", "", -1))
		out, err := os.Create(newFileName) // change to var
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
	}
}

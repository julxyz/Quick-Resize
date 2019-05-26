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

func resizeimg(file *os.File, filetype string) {

	// decode png into image.Image
	if filetype == "png" {
		img, err := png.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	} else if filetype == "jpg" {
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		img := "None"
	}
	file.Close()

	width, height := sizeDialog(img)

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

	width, height := sizeDialog(img)

	m := resize.Resize(width, height, img, resize.NearestNeighbor)

	newFileName := fmt.Sprintf("%s resized.jpg", strings.Replace(os.Args[2], ".jpg", "", -1))
	out, err := os.Create(newFileName) // change to var
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

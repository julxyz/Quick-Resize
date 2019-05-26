package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// open image from origin
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// filetype filter

	if strings.HasSuffix(strings.ToLower(os.Args[2]), ".png") {
		resizeimg(file, "png")
	}

	if strings.HasSuffix(strings.ToLower(os.Args[2]), ".jpg") {
		resizeimg(file, "jpg")
	}

	if strings.HasSuffix(strings.ToLower(os.Args[2]), ".jpeg") {
		resizeimg(file, "jpg")
	}
}

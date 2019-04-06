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

	if strings.HasSuffix(os.Args[2], ".png") {
		resizepng(file)
	}

	if strings.HasSuffix(os.Args[2], ".jpg") {
		resizejpg(file)
	}
}

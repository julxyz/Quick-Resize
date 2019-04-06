package main

import (
	"log"
	"os"
)

func main() {

	// create new directory

	mainDirNewPath := "C:/Program Files/Quick Resize"
	resDirNewPath := "C:/Program Files/Quick Resize/res"
	os.MkdirAll(mainDirNewPath, os.ModePerm)
	os.MkdirAll(resDirNewPath, os.ModePerm)

	// move files to program files

	execOldLocation := "src/resizer.exe"
	execNewLocation := "C:/Program Files/Quick Resize/resizer.exe"
	iconOldLocation := "src/icon.ico"
	iconNewLocation := "C:/Program Files/Quick Resize/res/icon.ico"
	err := os.Rename(execOldLocation, execNewLocation)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Rename(iconOldLocation, iconNewLocation)
	if err != nil {
		log.Fatal(err)
	}

	// call reg editor
	// editregistry()
}

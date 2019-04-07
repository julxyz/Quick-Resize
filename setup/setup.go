package main

import (
	"log"
	"os"
	"os/exec"
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
	err := os.Rename(execOldLocation, execNewLocation)
	if err != nil {
		log.Fatal(err)
	}

	cmnd := exec.Command("regedit.exe", "src/menu.reg")
	cmnd.Start()
	log.Println("log")
}

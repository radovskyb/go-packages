package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	dirPath := "tmp"

	// Create a new directory.
	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		log.Fatalln(err)
	}

	// Remove the directory.
	//
	// Rmdir deletes a directory, which must be empty.
	if err := syscall.Rmdir(dirPath); err != nil {
		log.Fatalln(err)
	}
}

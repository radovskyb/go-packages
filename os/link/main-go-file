package main

import (
	"log"
	"os"
)

func main() {
	// Create a hardlink, linking main-go-file to main.go
	//
	// Link creates newname as a hard link to the oldname file.
	// If there is an error, it will be of type *LinkError.
	err := os.Link("main.go", "main-go-file")

	// Log any errors
	if err != nil {
		log.Fatalln(err)
	}
}

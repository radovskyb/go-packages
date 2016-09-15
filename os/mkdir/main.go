package main

import (
	"log"
	"os"
)

func main() {
	// Use os.Mkdir to create a new directory tmp
	//
	// Mkdir creates a new directory with the specified name and
	// permission bits. If there is an error, it will be of type *PathError.
	err := os.Mkdir("tmp", 0755)

	// Log any errors
	if err != nil {
		log.Fatalln(err)
	}
}

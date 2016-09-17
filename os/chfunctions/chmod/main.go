package main

import (
	"log"
	"os"
)

func main() {
	// Create a new file with the permissions 0644
	if _, err := os.Create("file.txt"); err != nil {
		log.Fatalln(err)
	}

	// Change the file file.txt's permissions to 0755
	//
	// Chmod changes the mode of the named file to mode. If the file
	// is a symbolic link, it changes the mode of the link's target.
	// If there is an error, it will be of type *PathError.
	if err := os.Chmod("file.txt", 0755); err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Get the current working directory before using os.Chdir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the current working directory
	fmt.Println(dir)

	// Create a file handle to the /tmp directory.
	f, err := os.Open("/tmp")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Change the current working directory to the file handle's (tmp folder) directory
	//
	// Chdir changes the current working directory to the file, which
	// must be a directory. If there is an error, it will be of type *PathError.
	// err := os.Chdir("tmp")
	if err := f.Chdir(); err != nil {
		log.Fatalln(err)
	}

	// Get the current working directory
	dir, err = os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// Once again print the current working directory
	fmt.Println(dir)
}

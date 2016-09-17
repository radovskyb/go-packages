package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// Open the file file.txt for reading
	//
	// Open opens the named file for reading. If successful, methods on
	// the returned file can be used for reading; the associated file
	// descriptor has mode O_RDONLY. If there is an error, it will
	// be of type *PathError.
	file, err := os.Open("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Copy the contents of file.txt to os.Stdout
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatalln(err)
	}
}

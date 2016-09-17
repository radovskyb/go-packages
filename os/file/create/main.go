package main

import (
	"log"
	"os"
)

func main() {
	// Create a new file called file.txt with 0666 permissions
	//
	// Create creates the named file with mode 0666 (before umask),
	// truncating it if it already exists. If successful, methods on
	// the returned File can be used for I/O; the associated file descriptor
	// has mode O_RDWR. If there is an error, it will be of type *PathError.
	file, err := os.Create("file.txt")

	// Check if there were any errors creating file.txt and if
	// there was any, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Close the file handle.
	if err := file.Close(); err != nil {
		log.Fatalln(err)
	}
}

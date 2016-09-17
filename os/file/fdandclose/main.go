package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new file for reading
	f, err := os.Create("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Close and then remove the file after main finishes execution
	defer func() {
		// Close closes the File, rendering it unusable for I/O.
		// It returns an error, if any.
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}

		// Remove the file file.txt
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	// Get and print the files description (fd)
	//
	// Fd returns the integer Unix file descriptor referencing the
	// open file. The file descriptor is valid only until f.Close is
	// called or f is garbage collected.
	fmt.Println(f.Fd())
}

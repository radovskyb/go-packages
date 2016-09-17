package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a new file
	f, err := os.Create("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Close and then remove the file after main finishes execution
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	// Get and print the files name
	//
	// Name returns the name of the file as presented to Open.
	fmt.Println(f.Name())
}

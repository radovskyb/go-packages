package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Write the bytes `Hello, World!` to a file named file.txt
	//
	// WriteFile writes data to a file named by filename. If the file does
	// not exist, WriteFile creates it with permissions perm; otherwise
	// WriteFile truncates it before writing.
	err := ioutil.WriteFile("file.txt", []byte("Hello, World!"), 0644)

	// Check if there were any errors and if so log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print success message
	fmt.Println("Successfully created new file and wrote to it.")
}

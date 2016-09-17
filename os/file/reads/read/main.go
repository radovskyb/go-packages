package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open file.txt for reading
	f, err := os.Open("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Get the file.txt's info
	info, err := os.Stat("file.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Use the file's info to get it's size in bytes
	fSize := info.Size()

	// Create a new byte slice `data` of the size of file.txt
	data := make([]byte, fSize)

	// Read from the file into `data`
	//
	// Read reads up to len(b) bytes from the File. It returns the number
	// of bytes read and an error, if any. EOF is signaled by a zero
	// count with err set to io.EOF.
	_, err = f.Read(data)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the data received from file.txt as a string
	fmt.Println(string(data))
}

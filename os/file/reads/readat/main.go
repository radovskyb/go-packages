package main

import (
	"fmt"
	"io"
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

	// Read from the file into `data` starting at the offset of 7
	//
	// ReadAt reads len(b) bytes from the File starting at byte offset off.
	// It returns the number of bytes read and the error, if any. ReadAt always
	// returns a non-nil error when n < len(b). At end of file, that error is io.EOF.
	_, err = f.ReadAt(data, 7)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}

	// Print out the data received from file.txt as a string
	fmt.Println(string(data))
}

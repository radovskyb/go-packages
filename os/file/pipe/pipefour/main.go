package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const filename = "file.txt"

func main() {
	// Create a new file called file.txt with read/write permissons
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Defer to close the file handle and then remove file.txt
	defer func() {
		file.Close()
		os.Remove(filename)
	}()

	// Create a new pipe reader and writer
	r, w, err := os.Pipe()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Write a string to the pipe writer
	w.WriteString("Hello, World!")

	// Now close the pipe writer
	w.Close()

	// Write into file.txt what was received in the pipe reader
	// from the pipe writer (Hello, World!)
	io.Copy(file, r)

	// Get the contents of file.txt
	contents, err := ioutil.ReadFile(filename)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out file.txt's contents as a string
	fmt.Println(string(contents)) // Hello, World!
}

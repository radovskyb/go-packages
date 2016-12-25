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
	// Create a new file called file.txt with read/write permissions
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Defer to close the file handle and then remove file.txt
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := os.Remove(filename); err != nil {
			log.Fatalln(err)
		}
	}()

	// Create a new pipe reader and writer
	r, w, err := os.Pipe()

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Write a string to the pipe writer
	_, err = w.WriteString("Hello, World!")
	if err != nil {
		log.Fatalln(err)
	}

	// Now close the pipe writer
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}

	// Write into file.txt what was received in the pipe reader
	// from the pipe writer (Hello, World!)
	_, err = io.Copy(file, r)
	if err != nil {
		log.Fatalln(err)
	}

	// Get the contents of file.txt
	contents, err := ioutil.ReadFile(filename)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out file.txt's contents as a string
	fmt.Println(string(contents)) // Hello, World!
}

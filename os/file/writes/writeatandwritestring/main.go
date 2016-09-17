package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Open a new file for reading and writing
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR, 0755)

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Close and remove the file after main finishes execution
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
		if err := os.Remove("file.txt"); err != nil {
			log.Fatalln(err)
		}
	}()

	// Write the string sup to file.txt
	//
	// WriteString is like Write, but writes the contents of string
	// s rather than a slice of bytes.
	if _, err := file.WriteString("Sup?"); err != nil {
		log.Fatalln(err)
	}

	// Write to the file at offset 0 therefore replacing Sup?
	//
	// Write writes len(b) bytes to the File. It returns the number
	// of bytes written and an error, if any. Write returns a
	// non-nil error when n != len(b).
	_, err = file.WriteAt([]byte("Hello, World!"), 0)
	if err != nil {
		log.Fatalln(err)
	}

	// Retrieve the files contents
	contents, err := ioutil.ReadFile("file.txt")

	// Check if there were any errors and if so, log them
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the contents that were written to file.txt as a string
	fmt.Println(string(contents))
}

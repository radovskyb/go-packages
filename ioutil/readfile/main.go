package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read all the contents from file.txt into byte slice contents
	//
	// ReadFile reads the file named by filename and returns the contents.
	// A successful call returns err == nil, not err == EOF. Because ReadFile
	// reads the whole file, it does not treat an EOF from Read as an
	// error to be reported.
	contents, err := ioutil.ReadFile("file.txt")

	// Check if there were any errors from ReadFile
	if err != nil {
		// Print any errors that occurred
		log.Fatalln(err)
	}

	// Print out contents as a string.
	fmt.Println(string(contents))
}

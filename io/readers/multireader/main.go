package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	// Create a string reader that reads the string "Hello, World!\n"
	sr := strings.NewReader("Hello, World!\n")

	// Now create a new bytes reader that reads the bytes "What's Up?"
	br := bytes.NewReader([]byte("What's Up?"))

	// Create an io.Reader slice of 2 readers, br and sr
	readers := []io.Reader{sr, br}

	// Create a new multi reader that reads from both byte reader br
	// and string reader sr, combined into a single reader m
	//
	// MultiReader returns a Reader that's the logical concatenation
	// of the provided input readers. They're read sequentially. Once
	// all inputs have returned EOF, Read will return EOF. If any of
	// the readers return a non-nil, non-EOF error, Read will return that error.
	m := io.MultiReader(readers...)

	// Read all the data from both readers in multi reader m
	// using ioutil.ReadAll which reads all data from a reader
	// until either an error or EOF occurs
	data, err := ioutil.ReadAll(m)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the combined data read from multi reader m
	// in stringified format as opposed to a slice of bytes
	fmt.Println(string(data))
}

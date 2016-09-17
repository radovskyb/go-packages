package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Copy from stdin to stdout but only the allowed number
	// of bytes, passed in as a parameter to io.CopyN()
	//
	// CopyN copies n bytes (or until an error) from src to dst.
	// It returns the number of bytes copied and the earliest error
	// encountered while copying. On return, written == n if and
	// only if err == nil.
	//
	// If dst implements the ReaderFrom interface, the copy
	// is implemented using it.
	n, err := io.CopyN(os.Stdout, os.Stdin, 10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Copied %d bytes from os.Stdin to os.Stdout", n)
}

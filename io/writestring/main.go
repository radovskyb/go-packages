package main

import (
	"io"
	"os"
)

func main() {
	// Write a string to stdout
	//
	// WriteString writes the contents of the string s to w,
	// which accepts a slice of bytes. If w implements a
	// WriteString method, it is invoked directly.
	io.WriteString(os.Stdout, "Hello, World!")
}

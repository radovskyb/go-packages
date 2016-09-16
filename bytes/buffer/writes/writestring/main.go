package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes buffer buf with some contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c'})

	// Print buf's contents
	fmt.Println(buf) // abc

	// Append a string to buf
	//
	// WriteString appends the contents of s to the buffer, growing
	// the buffer as needed. The return value n is the length of s;
	// err is always nil. If the buffer becomes too large, WriteString
	// will panic with ErrTooLarge.
	_, err := buf.WriteString("defghijk")
	if err != nil {
		log.Fatalln(err)
	}

	// Once again print buf's contents
	fmt.Println(buf) // abcdefghijk
}

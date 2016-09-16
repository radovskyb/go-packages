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

	// Append 1 rune to buf
	//
	// WriteRune appends the UTF-8 encoding of Unicode code point r
	// to the buffer, returning its length and an error, which is always
	// nil but is included to match bufio.Writer's WriteRune. The buffer
	// is grown as needed; if it becomes too large, WriteRune will
	// panic with ErrTooLarge.
	_, err := buf.WriteRune('d')
	if err != nil {
		log.Fatalln(err)
	}

	// Once again print buf's contents
	fmt.Println(buf) // abcd
}

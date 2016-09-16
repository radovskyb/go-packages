package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes buffer with some contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c', 'd', 'e', 'f'})

	// Read one rune from buf into r
	r, _, err := buf.ReadRune()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the rune we read
	fmt.Println(string(r))

	// Now print the remaining contents of buf
	fmt.Println(buf) // bcdef

	// Now unread the rune that we read before
	//
	// UnreadRune unreads the last rune returned by ReadRune. If the
	// most recent read or write operation on the buffer was not a
	// ReadRune, UnreadRune returns an error. (In this regard it is
	// stricter than UnreadByte, which will unread the last byte
	// from any read operation.)
	if err := buf.UnreadRune(); err != nil {
		log.Fatalln(err)
	}

	// Once again print the remaining contents of buf
	fmt.Println(buf) // abcdef
}

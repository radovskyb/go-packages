package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	// Create a new bytes buffer with some contents
	buf := bytes.NewBuffer([]byte{'a', 'b', 'c', 'd', 'e', 'f'})

	// Read one byte from buf into c
	c, err := buf.ReadByte()
	if err != nil {
		log.Fatalln(err)
	}

	// Print the character we read
	fmt.Println(string(c))

	// Now print the remaining contents of buf
	fmt.Println(buf) // bcdef

	// Now unread the byte that we read before
	//
	// UnreadByte unreads the last byte returned by the most
	// recent read operation. If write has happened since the
	// last read, UnreadByte returns an error.
	if err := buf.UnreadByte(); err != nil {
		log.Fatalln(err)
	}

	// Once again print the remaining contents of buf
	fmt.Println(buf) // abcdef
}

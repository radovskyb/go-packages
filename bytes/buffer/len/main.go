package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create a byte slice b with 2 byte elements inside of it
	b := []byte{'a', 'b'}

	// Create a new bytes buffer buf with b as it's initial contents
	buf := bytes.NewBuffer(b)

	// Print the current length of buffer buf
	//
	// Len returns the number of bytes of the unread portion
	// of the buffer; b.Len() == len(b.Bytes()).
	fmt.Println(buf.Len()) // 2

	// Write one more byte element `c` to the buffer buf
	_, err := buf.Write([]byte{'c'})
	if err != nil {
		log.Fatalln(err)
	}

	// Once again print the current length of buffer buf
	fmt.Println(buf.Len()) // 3

	// Write out the buffers contents to os.Stdout
	_, err = buf.WriteTo(os.Stdout) // abc
	if err != nil {
		log.Fatalln(err)
	}
}

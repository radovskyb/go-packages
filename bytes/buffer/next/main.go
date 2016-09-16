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

	// Print out the next 2 bytes in the buffer buf using buf.Next(2)
	//
	// Next returns a slice containing the next n bytes from the buffer,
	// advancing the buffer as if the bytes had been returned by Read. If
	// there are fewer than n bytes in the buffer, Next returns the entire buffer.
	// The slice is only valid until the next call to a read or write method.
	fmt.Println(string(buf.Next(2))) // ab

	// Add some more bytes to buffer buf
	fmt.Fprintf(buf, "cdefghijklmnopqrstuv")

	// Now print out the next 10 bytes in the buffer buf using buf.Next(10)
	fmt.Println(string(buf.Next(10))) // cdefghijkl

	// Write the rest of the buffer out to os.Stdout
	_, err := buf.WriteTo(os.Stdout) // mnopqrstuv
	if err != nil {
		log.Fatalln(err)
	}
}
